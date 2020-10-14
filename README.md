# OPA Gatekeeper Operator
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

### Outside the Cluster

If you would like to run the Operator outside the cluster, you'll have to set the
`WATCH_NAMESPACE` environment variable to the namespace you want the
Operator to monitor:

1. Set the WATCH_NAMESPACE environment variable:
    ```shell
    export WATCH_NAMESPACE=gatekeeper-system
    ```
1. You then run the Operator with:
    ```shell
    make run
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
