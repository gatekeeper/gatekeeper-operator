SHELL := /bin/bash
# Detect the OS to set per-OS defaults
OS_NAME = $(shell uname -s)
# Current Operator version
VERSION ?= v0.1.1
# Current Gatekeeper version
GATEKEEPER_VERSION ?= v3.3.0
# Default image repo
REPO ?= quay.io/gatekeeper
# Default bundle image tag
BUNDLE_IMG ?= $(REPO)/gatekeeper-operator-bundle:$(VERSION)
# Default bundle index image tag
BUNDLE_INDEX_IMG ?= $(REPO)/gatekeeper-operator-bundle-index:$(VERSION)
# Default namespace
NAMESPACE ?= gatekeeper-system
# Default Kubernetes distribution
KUBE_DISTRIBUTION ?= vanilla
# Options for 'bundle-build'
CHANNELS ?= stable
DEFAULT_CHANNEL ?= stable
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# Image URL to use all building/pushing image targets
IMG ?= $(REPO)/gatekeeper-operator:$(VERSION)
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true,crdVersions=v1beta1"

GATEKEEPER_MANIFEST_DIR ?= config/gatekeeper
OPENSHIFT_RBAC_DIR = config/rbac/overlays/openshift

ifeq (openshift, $(KUBE_DISTRIBUTION))
RBAC_DIR=$(OPENSHIFT_RBAC_DIR)
else
RBAC_DIR=config/rbac/base
endif

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Get the current controller-gen binary. If there isn't any, we'll use the
# GOBIN path
ifeq (, $(shell which controller-gen))
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# Get the current kustomize binary. If there isn't any, we'll use the
# GOBIN path
ifeq (, $(shell which kustomize))
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Get the current opm binary. If there isn't any, we'll use the
# GOBIN path
ifeq (, $(shell which opm))
OPM=$(GOBIN)/opm
else
OPM=$(shell which opm)
endif

# operator-sdk variables
# ======================
OPERATOR_SDK_VERSION ?= v1.2.0
ifeq ($(OS_NAME), Linux)
    OPERATOR_SDK_URL=https://github.com/operator-framework/operator-sdk/releases/download/$(OPERATOR_SDK_VERSION)/operator-sdk-$(OPERATOR_SDK_VERSION)-x86_64-linux-gnu
else ifeq ($(OS_NAME), Darwin)
    OPERATOR_SDK_URL=https://github.com/operator-framework/operator-sdk/releases/download/$(OPERATOR_SDK_VERSION)/operator-sdk-$(OPERATOR_SDK_VERSION)-x86_64-apple-darwin
endif

# Get the current operator-sdk binary. If there isn't any, we'll use the
# GOBIN path
ifeq (, $(shell which operator-sdk))
OPERATOR_SDK=$(GOBIN)/operator-sdk
else
OPERATOR_SDK=$(shell which operator-sdk)
endif

# Use the vendored directory
GOFLAGS = -mod=vendor

# Set version variables for LDFLAGS
GIT_VERSION ?= $(shell git describe --match='v*' --always --dirty)
GIT_HASH ?= $(shell git rev-parse HEAD)
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_TREESTATE = "clean"
DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(DIFF), 1)
    GIT_TREESTATE = "dirty"
endif
LDFLAGS = "-X github.com/gatekeeper/gatekeeper-operator/pkg/version.gitVersion=$(GIT_VERSION) \
             -X github.com/gatekeeper/gatekeeper-operator/pkg/version.gitCommit=$(GIT_HASH) \
             -X github.com/gatekeeper/gatekeeper-operator/pkg/version.gitTreeState=$(GIT_TREESTATE) \
             -X github.com/gatekeeper/gatekeeper-operator/pkg/version.buildDate=$(BUILDDATE)"

.PHONY: all
all: manager

# Run tests
# Set SKIP_FETCH_TOOLS=y to use tools in your own environment
ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
.PHONY: test
test: generate fmt vet manifests
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/master/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); GOFLAGS=$(GOFLAGS) go test -v ./... -coverprofile cover.out

.PHONY: test-e2e
test-e2e: generate fmt vet
	GOFLAGS=$(GOFLAGS) USE_EXISTING_CLUSTER=true go test -v ./test/e2e -coverprofile cover.out -race -args -ginkgo.v -ginkgo.trace -namespace $(NAMESPACE)

.PHONY: deploy-olm
deploy-olm:
	kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.17.0/crds.yaml
	kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.17.0/olm.yaml

.PHONY: deploy-with-olm
deploy-with-olm:
	sed -i 's#quay.io/gatekeeper/gatekeeper-operator-bundle-index:latest#$(BUNDLE_INDEX_IMG)#g' config/olm-install/install-resources.yaml
	(cd config/olm-install && $(KUSTOMIZE) edit set namespace $(NAMESPACE))
	$(KUSTOMIZE) build config/olm-install | kubectl apply -f -

# Build manager binary
.PHONY: manager
manager: generate manifests
	GOFLAGS=$(GOFLAGS) go build -ldflags $(LDFLAGS) -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate fmt vet manifests
	GOFLAGS=$(GOFLAGS) GATEKEEPER_TARGET_NAMESPACE=$(NAMESPACE) go run -ldflags $(LDFLAGS) ./main.go

# Install CRDs into a cluster
.PHONY: install
install: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
.PHONY: uninstall
uninstall: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests kustomize
	cd config/default && $(KUSTOMIZE) edit set namespace $(NAMESPACE)
	cd $(RBAC_DIR) && $(KUSTOMIZE) edit set namespace $(NAMESPACE)
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	{ $(KUSTOMIZE) build config/default ; echo "---" ; $(KUSTOMIZE) build $(RBAC_DIR) ; } | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases output:rbac:dir=config/rbac/base

# Path used to import Gatekeeper manifests. For example, this could be a local
# file system directory if kustomize has errors using the GitHub URL. See
# https://github.com/kubernetes-sigs/kustomize/issues/3515 for details.
IMPORT_MANIFESTS_PATH ?= https://github.com/open-policy-agent/gatekeeper

# Import Gatekeeper manifests
.PHONY: import-manifests
import-manifests: kustomize
	if [[ $(IMPORT_MANIFESTS_PATH) =~ https://* ]]; then \
		$(KUSTOMIZE) build $(IMPORT_MANIFESTS_PATH)/config/overlays/mutation_webhook/?ref=$(GATEKEEPER_VERSION) -o $(GATEKEEPER_MANIFEST_DIR); \
	else \
		$(KUSTOMIZE) build $(IMPORT_MANIFESTS_PATH)/config/overlays/mutation_webhook -o $(GATEKEEPER_MANIFEST_DIR); \
		$(KUSTOMIZE) build --load_restrictor LoadRestrictionsNone $(IMPORT_MANIFESTS_PATH)/config/overlays/mutation -o $(GATEKEEPER_MANIFEST_DIR); \
	fi

# Run go fmt against code
.PHONY: fmt
fmt:
	GOFLAGS=$(GOFLAGS) go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	GOFLAGS=$(GOFLAGS) go vet ./...

# Generate code
.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

BINDATA_OUTPUT_FILE := ./pkg/bindata/bindata.go
.ensure-go-bindata:
	ln -s $(abspath ./vendor) "$${TMP_GOPATH}/src"
	export GO111MODULE=off && export GOPATH=$${TMP_GOPATH} && export GOBIN=$${TMP_GOPATH}/bin && GOFLAGS=$(GOFLAGS) go install "./vendor/github.com/go-bindata/go-bindata/..."
.PHONY: .ensure-go-bindata

.run-bindata: .ensure-go-bindata
	$${TMP_GOPATH}/bin/go-bindata -nocompress -nometadata \
		-prefix "bindata" \
		-pkg "bindata" \
		-o "$${BINDATA_OUTPUT_PREFIX}$(BINDATA_OUTPUT_FILE)" \
		-ignore "OWNERS" \
		./$(GATEKEEPER_MANIFEST_DIR)/... && \
	gofmt -s -w "$${BINDATA_OUTPUT_PREFIX}$(BINDATA_OUTPUT_FILE)"
.PHONY: .run-bindata

update-bindata:
	export TMP_GOPATH=$$(mktemp -d) ;\
	$(MAKE) .run-bindata ;\
	rm -rf "$${TMP_GOPATH}"
.PHONY: update-bindata

verify-bindata:
	export TMP_GOPATH=$$(mktemp -d) ;\
	export TMP_DIR=$$(mktemp -d) ;\
	export BINDATA_OUTPUT_PREFIX="$${TMP_DIR}/" ;\
	$(MAKE) .run-bindata ;\
	diff -Naup {.,$${TMP_DIR}}/$(BINDATA_OUTPUT_FILE) ;\
	rm -rf "$${TMP_DIR}" ;\
	rm -rf "$${TMP_GOPATH}"
.PHONY: verify-bindata

# Build the docker image
.PHONY: docker-build
docker-build:
	docker build . --build-arg LDFLAGS=${LDFLAGS} -t ${IMG}

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

# find or download controller-gen
# download controller-gen if necessary
.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN)

$(CONTROLLER_GEN):
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}

.PHONY: kustomize
kustomize: $(KUSTOMIZE)

$(KUSTOMIZE):
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.8.6 ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}

.PHONY: opm
opm: $(OPM)

$(OPM):
	@{ \
	set -e ;\
	OPM_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$OPM_GEN_TMP_DIR ;\
	export GOPATH=$${OPM_GEN_TMP_DIR} ;\
	go get github.com/operator-framework/operator-registry || true;\
	cd src/github.com/operator-framework/operator-registry ;\
	git checkout -b v1.15.1 ;\
	make bin/opm ;\
	mv bin/opm $@ ;\
	rm -rf $$OPM_GEN_TMP_DIR ;\
	}

.PHONY: operator-sdk
operator-sdk: $(OPERATOR_SDK)

$(OPERATOR_SDK):
	curl -L $(OPERATOR_SDK_URL) -o $(OPERATOR_SDK) || (echo "curl returned $$? trying to fetch operator-sdk. Please install operator-sdk and try again"; exit 1)
	chmod +x $(OPERATOR_SDK)

# Generate bundle manifests and metadata, then validate generated files.
.PHONY: bundle
bundle: operator-sdk manifests
	$(OPERATOR_SDK) generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	VERSION=$(VERSION) ;\
	{ $(KUSTOMIZE) build config/manifests ; echo "---" ; $(KUSTOMIZE) build $(OPENSHIFT_RBAC_DIR) ; } | $(OPERATOR_SDK) generate bundle -q --overwrite --version $${VERSION/v/} $(BUNDLE_METADATA_OPTS)
	sed -i 's/base64data: \"\"/base64data: \"PHN2ZyBpZD0iZjc0ZTM5ZDEtODA2Yy00M2E0LTgyZGQtZjM3ZjM1NWQ4YWYzIiBkYXRhLW5hbWU9Ikljb24iIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgdmlld0JveD0iMCAwIDM2IDM2Ij4KICA8ZGVmcz4KICAgIDxzdHlsZT4KICAgICAgLmE0MWM1MjM0LWExNGEtNGYzZC05MTYwLTQ0NzJiNzZkMDA0MCB7CiAgICAgICAgZmlsbDogI2UwMDsKICAgICAgfQogICAgPC9zdHlsZT4KICA8L2RlZnM+CiAgPGc+CiAgICA8cGF0aCBjbGFzcz0iYTQxYzUyMzQtYTE0YS00ZjNkLTkxNjAtNDQ3MmI3NmQwMDQwIiBkPSJNMjUsMTcuMzhIMjMuMjNhNS4yNyw1LjI3LDAsMCwwLTEuMDktMi42NGwxLjI1LTEuMjVhLjYyLjYyLDAsMSwwLS44OC0uODhsLTEuMjUsMS4yNWE1LjI3LDUuMjcsMCwwLDAtMi42NC0xLjA5VjExYS42Mi42MiwwLDEsMC0xLjI0LDB2MS43N2E1LjI3LDUuMjcsMCwwLDAtMi42NCwxLjA5bC0xLjI1LTEuMjVhLjYyLjYyLDAsMCwwLS44OC44OGwxLjI1LDEuMjVhNS4yNyw1LjI3LDAsMCwwLTEuMDksMi42NEgxMWEuNjIuNjIsMCwwLDAsMCwxLjI0aDEuNzdhNS4yNyw1LjI3LDAsMCwwLDEuMDksMi42NGwtMS4yNSwxLjI1YS42MS42MSwwLDAsMCwwLC44OC42My42MywwLDAsMCwuODgsMGwxLjI1LTEuMjVhNS4yNyw1LjI3LDAsMCwwLDIuNjQsMS4wOVYyNWEuNjIuNjIsMCwwLDAsMS4yNCwwVjIzLjIzYTUuMjcsNS4yNywwLDAsMCwyLjY0LTEuMDlsMS4yNSwxLjI1YS42My42MywwLDAsMCwuODgsMCwuNjEuNjEsMCwwLDAsMC0uODhsLTEuMjUtMS4yNWE1LjI3LDUuMjcsMCwwLDAsMS4wOS0yLjY0SDI1YS42Mi42MiwwLDAsMCwwLTEuMjRabS03LDQuNjhBNC4wNiw0LjA2LDAsMSwxLDIyLjA2LDE4LDQuMDYsNC4wNiwwLDAsMSwxOCwyMi4wNloiLz4KICAgIDxwYXRoIGNsYXNzPSJhNDFjNTIzNC1hMTRhLTRmM2QtOTE2MC00NDcyYjc2ZDAwNDAiIGQ9Ik0yNy45LDI4LjUyYS42Mi42MiwwLDAsMS0uNDQtLjE4LjYxLjYxLDAsMCwxLDAtLjg4LDEzLjQyLDEzLjQyLDAsMCwwLDIuNjMtMTUuMTkuNjEuNjEsMCwwLDEsLjMtLjgzLjYyLjYyLDAsMCwxLC44My4yOSwxNC42NywxNC42NywwLDAsMS0yLjg4LDE2LjYxQS42Mi42MiwwLDAsMSwyNy45LDI4LjUyWiIvPgogICAgPHBhdGggY2xhc3M9ImE0MWM1MjM0LWExNGEtNGYzZC05MTYwLTQ0NzJiNzZkMDA0MCIgZD0iTTI3LjksOC43M2EuNjMuNjMsMCwwLDEtLjQ0LS4xOUExMy40LDEzLjQsMCwwLDAsMTIuMjcsNS45MWEuNjEuNjEsMCwwLDEtLjgzLS4zLjYyLjYyLDAsMCwxLC4yOS0uODNBMTQuNjcsMTQuNjcsMCwwLDEsMjguMzQsNy42NmEuNjMuNjMsMCwwLDEtLjQ0LDEuMDdaIi8+CiAgICA8cGF0aCBjbGFzcz0iYTQxYzUyMzQtYTE0YS00ZjNkLTkxNjAtNDQ3MmI3NmQwMDQwIiBkPSJNNS4zNSwyNC42MmEuNjMuNjMsMCwwLDEtLjU3LS4zNUExNC42NywxNC42NywwLDAsMSw3LjY2LDcuNjZhLjYyLjYyLDAsMCwxLC44OC44OEExMy40MiwxMy40MiwwLDAsMCw1LjkxLDIzLjczYS42MS42MSwwLDAsMS0uMy44M0EuNDguNDgsMCwwLDEsNS4zNSwyNC42MloiLz4KICAgIDxwYXRoIGNsYXNzPSJhNDFjNTIzNC1hMTRhLTRmM2QtOTE2MC00NDcyYjc2ZDAwNDAiIGQ9Ik0xOCwzMi42MkExNC42NCwxNC42NCwwLDAsMSw3LjY2LDI4LjM0YS42My42MywwLDAsMSwwLS44OC42MS42MSwwLDAsMSwuODgsMCwxMy40MiwxMy40MiwwLDAsMCwxNS4xOSwyLjYzLjYxLjYxLDAsMCwxLC44My4zLjYyLjYyLDAsMCwxLS4yOS44M0ExNC42NywxNC42NywwLDAsMSwxOCwzMi42MloiLz4KICAgIDxwYXRoIGNsYXNzPSJhNDFjNTIzNC1hMTRhLTRmM2QtOTE2MC00NDcyYjc2ZDAwNDAiIGQ9Ik0zMCwyOS42MkgyN2EuNjIuNjIsMCwwLDEtLjYyLS42MlYyNmEuNjIuNjIsMCwwLDEsMS4yNCwwdjIuMzhIMzBhLjYyLjYyLDAsMCwxLDAsMS4yNFoiLz4KICAgIDxwYXRoIGNsYXNzPSJhNDFjNTIzNC1hMTRhLTRmM2QtOTE2MC00NDcyYjc2ZDAwNDAiIGQ9Ik03LDMwLjYyQS42Mi42MiwwLDAsMSw2LjM4LDMwVjI3QS42Mi42MiwwLDAsMSw3LDI2LjM4aDNhLjYyLjYyLDAsMCwxLDAsMS4yNEg3LjYyVjMwQS42Mi42MiwwLDAsMSw3LDMwLjYyWiIvPgogICAgPHBhdGggY2xhc3M9ImE0MWM1MjM0LWExNGEtNGYzZC05MTYwLTQ0NzJiNzZkMDA0MCIgZD0iTTI5LDkuNjJIMjZhLjYyLjYyLDAsMCwxLDAtMS4yNGgyLjM4VjZhLjYyLjYyLDAsMCwxLDEuMjQsMFY5QS42Mi42MiwwLDAsMSwyOSw5LjYyWiIvPgogICAgPHBhdGggY2xhc3M9ImE0MWM1MjM0LWExNGEtNGYzZC05MTYwLTQ0NzJiNzZkMDA0MCIgZD0iTTksMTAuNjJBLjYyLjYyLDAsMCwxLDguMzgsMTBWNy42Mkg2QS42Mi42MiwwLDAsMSw2LDYuMzhIOUEuNjIuNjIsMCwwLDEsOS42Miw3djNBLjYyLjYyLDAsMCwxLDksMTAuNjJaIi8+CiAgPC9nPgo8L3N2Zz4K\"/g' bundle/manifests/gatekeeper-operator.clusterserviceversion.yaml
	sed -i 's/mediatype: \"\"/mediatype: \"image\/svg+xml\"/g' bundle/manifests/gatekeeper-operator.clusterserviceversion.yaml
	$(OPERATOR_SDK) bundle validate ./bundle

# Build the bundle image.
.PHONY: bundle-build
bundle-build:
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

# Build the bundle index image.
.PHONY: bundle-index-build
bundle-index-build: opm
	$(OPM) index add --bundles $(BUNDLE_IMG) --tag $(BUNDLE_INDEX_IMG) -c docker

# Generate and push bundle image and bundle index image
# Note: OPERATOR_VERSION is an arbitrary number and does not need to match any official versions
.PHONY: build-and-push-bundle-images
build-and-push-bundle-images: docker-build docker-push 
	$(MAKE) bundle VERSION=$(OPERATOR_VERSION)
	$(MAKE) bundle-build
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)
	$(MAKE) bundle-index-build
	$(MAKE) docker-push IMG=$(BUNDLE_INDEX_IMG)

.PHONY: vendor
vendor:
	GO111MODULE=on GOFLAGS=$(GOFLAGS) go mod vendor

.PHONY: tidy
tidy:
	GO111MODULE=on GOFLAGS=$(GOFLAGS) go mod tidy

.PHONY: test-cluster
test-cluster:
	./scripts/kind-with-registry.sh

.PHONY: test-gatekeeper-e2e
test-gatekeeper-e2e:
	kubectl -n $(NAMESPACE) apply -f ./config/samples/gatekeeper_e2e_test.yaml
	bats -t test/bats/test.bats

.PHONY: deploy-ci
deploy-ci: deploy-ci-namespace deploy

.PHONY: deploy-ci-namespace
deploy-ci-namespace: install
	kubectl create namespace --dry-run=client -o yaml $(NAMESPACE) | kubectl apply -f-
	sed -i 's/imagePullPolicy: Always/imagePullPolicy: IfNotPresent/g' config/manager/manager.yaml

.PHONY: release
release: manifests kustomize
	cd config/default && $(KUSTOMIZE) edit set namespace $(NAMESPACE)
	cd $(RBAC_DIR) && $(KUSTOMIZE) edit set namespace $(NAMESPACE)
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	{ $(KUSTOMIZE) build config/default ; echo "---" ; $(KUSTOMIZE) build $(RBAC_DIR) ; } > ./deploy/gatekeeper-operator.yaml
