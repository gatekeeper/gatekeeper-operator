# Archived: Use [stolostron/gatekeeper-operator](https://github.com/stolostron/gatekeeper-operator)

# OPA Gatekeeper Operator
[![CI-Tests](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/ci_tests.yaml/badge.svg)](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/ci_tests.yaml)
[![OLM-Tests](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/olm_tests.yaml/badge.svg)](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/olm_tests.yaml)
[![Create Release](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/release.yaml/badge.svg)](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/release.yaml)
[![Image](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/image.yaml/badge.svg)](https://github.com/gatekeeper/gatekeeper-operator/actions/workflows/image.yaml)
[![Docker Repository on Quay](https://img.shields.io/:Image-Quay-blue.svg)](https://quay.io/repository/gatekeeper/gatekeeper-operator)

Operator for OPA Gatekeeper

## Design

Please see the Gatekeeper Operator design document located at
https://docs.google.com/document/d/1Nxw4Agq6nJrPL24fJPiTXtjtLQRsLJtHo9x5urwYB_I/edit?usp=sharing
for some background information.

## Installation

To install the Gatekeeper Operator, you can either run it outside the cluster,
for faster iteration during development, or inside the cluster.

But first we require installing the Operator CRD:

```shell
make install
```

Then proceed to the installation method you prefer below.

Note: By default, `docker` is used to build and push images in Makefile. To use `podman` instead, set `DOCKER=podman` when calling `make`.

### Outside the Cluster

If you would like to run the Gatekeeper Operator outside the cluster, the
operator will use the default namespace to deploy Gatekeeper. If instead you
would like to have the operator deploy Gatekeeper to a different namespace,
then set the `NAMESPACE` environment variable. To do that just execute:

```shell
make run NAMESPACE=<namespace>
```

### Inside the Cluster

If you would like to run the Operator inside the cluster, you'll need to build
a container image. You can use a local private registry, or host it on a public
registry service like [quay.io](https://quay.io).

1. Build your image:
    ```shell
    make docker-build IMG=<registry>/<imagename>:<tag>
    ```
1. Push the image:
    ```shell
    make docker-push IMG=<registry>/<imagename>:<tag>
    ```
1. Deploy the Operator:
    ```shell
    make deploy IMG=<registry>/<imagename>:<tag>
    ```

You can also specify in which namespace you want the operator to be deployed to by
providing the `NAMESPACE` variable. If not provided the default namespace will be 
`gatekeeper-system`.

```shell
make deploy IMG=<registry>/<imagename>:tag NAMESPACE=mygatekeeper
```

### Deploy Operator using OLM

If you would like to deploy Operator using OLM, you'll need to build and push the bundle image and index image. You need to host the images on a public registry service like [quay.io](https://quay.io).

1. Build your bundle image
    ```shell
    make bundle-build REPO=<registry>
    ```
1. Push the bundle image
    ```shell
    make docker-push IMG=<bundle image name>
    ```
1. Build the index image

    This `make` target will install `opm` if it is not already installed. If
    you would like to install it in your `PATH` manually instead, get it from
    [here](https://github.com/operator-framework/operator-registry/releases).
    ```shell
    make bundle-index-build REPO=<registry>
    ```
1. Push the index image
    ```shell
    make docker-push IMG=<index image name>
    ```
1. Create the CatalogSource/OperatorGroup/Subscription
    ```yaml
    ---
    apiVersion: operators.coreos.com/v1alpha1
    kind: CatalogSource
    metadata:
      name: gatekeeper-operator
      namespace: gatekeeper-system
    spec:
      displayName: Gatekeeper Operator Upstream
      image: <index image name>
      publisher: github.com/gatekeeper/gatekeeper-operator
      sourceType: grpc
    ---
    apiVersion: operators.coreos.com/v1
    kind: OperatorGroup
    metadata:
      name: gatekeeper-operator
      namespace: gatekeeper-system
    ---
    apiVersion: operators.coreos.com/v1alpha1
    kind: Subscription
    metadata:
      name: gatekeeper-operator-sub
      namespace: gatekeeper-system
    spec:
      name: gatekeeper-operator
      channel: alpha
      source: gatekeeper-operator
      sourceNamespace: gatekeeper-system
    ```

## Usage

Before using Gatekeeper you have to create a `gatekeeper` resource that will be consumed by the operator and create all the necessary resources for you.

Here you can find an example of a `gatekeeper` resource definition:

```yaml
apiVersion: operator.gatekeeper.sh/v1alpha1
kind: Gatekeeper
metadata:
  name: gatekeeper
spec:
  # Add fields here
  audit:
    replicas: 1
    logLevel: ERROR
```

If nothing is defined in the `spec`, the default values will be used. In the example above the number of replicas for the audit pod is set to `1` and the logLevel to `ERROR` where the default is `INFO`.

The default behaviour for the `ValidatingWebhookConfiguration` is `ENABLED`, that means that it will be created. To disable the `ValidatingWebhookConfiguration` deployment, set the `validatingWebhook` spec property to `DISABLED`.

In order to create an instance of gatekeeper in the specified namespace you can start from one of the [sample configurations](config/samples).

```shell
kubectl create -f config/samples/operator_v1alpha1_gatekeeper.yaml
```
