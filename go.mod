module github.com/gatekeeper/gatekeeper-operator

go 1.15

require (
	github.com/RHsyseng/operator-utils v1.4.6-0.20201116165605-e3e8466ece23
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-logr/logr v0.2.1
	github.com/go-logr/zapr v0.2.0 // indirect
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/openshift/library-go v0.0.0-20201006230840-f360b9835cc8
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	sigs.k8s.io/controller-runtime v0.6.2
)
