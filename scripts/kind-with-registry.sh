#!/usr/bin/env bash
set -o errexit

# desired cluster name; default is "kind"
KIND_CLUSTER_NAME="${KIND_CLUSTER_NAME:-kind}"
KIND_IMG_TAG="${KIND_IMG_TAG:-}"

# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

kind version

KIND_CMD=
if [[ -z "${KIND_IMG_TAG}" ]]; then
  KIND_CMD="kind create cluster ${IMG} --name ${KIND_CLUSTER_NAME} --wait=5m --config=-"
else
  KIND_CMD="kind create cluster --image kindest/node:${KIND_IMG_TAG} --name ${KIND_CLUSTER_NAME} --wait=5m --config=-"
fi

# create a cluster with the local registry enabled in containerd
cat <<EOF | ${KIND_CMD}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF

# connect the registry to the cluster network
docker network connect "kind" "${reg_name}" || true

# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
