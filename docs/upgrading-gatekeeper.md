# Upgrading Gatekeeper

The following are steps to perform when upgrading the operator to support a new
version of Gatekeeper. This guide is not intended to be all-encompassing, but a
good starting point for the general steps needed to be done assuming no new
significant functionality is being introduced.

As a general guideline to upgrading Gatekeeper, you may want to consider
upgrading one minor version at a time. This can be time consuming, but it may
help better deduce the incremental changes that were made in each version.
However, note that this is generally not required, especially if Gatekeeper did
not introduce any significant i.e. breaking, functionality in intermediate
versions such that you can jump straight to the desired version. For example,
this was done when supporting Gatekeeper `v3.3.0` and then supporting `v3.5.1`
and skipping `v3.4.Z`.

## 1. Set desired Gatekeeper version and commit

```shell
GATEKEEPER_PREV_VERSION=$(awk '/^GATEKEEPER_VERSION \?= .*/ {print $3}' Makefile)
GATEKEEPER_VERSION=<DESIRED_VERSION>
sed -i "s/GATEKEEPER_VERSION ?= .*/GATEKEEPER_VERSION ?= ${GATEKEEPER_VERSION}/" Makefile
git commit -m "Set Gatekeeper version to ${GATEKEEPER_VERSION}" Makefile
```

## 2. Update Operator deployment's Gatekeeper image environment variable

```shell
sed -Ei "s|(value: openpolicyagent/gatekeeper:)${GATEKEEPER_PREV_VERSION}|\1${GATEKEEPER_VERSION}|" ./config/manager/manager.yaml
git commit -m "Update deployed Gatekeeper image to ${GATEKEEPER_VERSION}" ./config/manager/manager.yaml
```

## 3. Update existing sample configs

```shell
sed -i "s/${GATEKEEPER_PREV_VERSION}/${GATEKEEPER_VERSION}/" ./config/samples/*
git commit -m "Update sample configs to use Gatekeeper ${GATEKEEPER_VERSION}" ./config/samples/*
```

## 4. Import Gatekeeper manifests and commit

Use the `Makefile` target `import-manifests` to import manifests
directly from Gatekeeper's repository and then commit the changes.

```shell
make import-manifests
```

## 5. Assess changes and if any new Gatekeeper manifests have been added

Assess any changes that have been introduced by Gatekeeper that may impact how
the operator needs to handle it. These are things like new option flags,
different defaults, etc.

Also see if git detects that there are untracked files. If so, then Gatekeeper
has added some new manifests that will need to be version controlled and added
to the [operator controller's list of Gatekeeper manifests to
manage](https://github.com/gatekeeper/gatekeeper-operator/blob/44420118530000ff25264adc4229b4490013abed/controllers/gatekeeper_controller.go#L83-L114).

Add and commit the changes once you've made a note of any new Gatekeeper
manifests.

```shell
git add ./config/gatekeeper/*
git commit -m "make import-manifests" ./config/gatekeeper/
```

## 6. Update static binary data

Use the `Makefile` target `update-bindata` to use the `go-bindata` tool to
generate the static data that contains all the Gatekeeper manifests the
operator manages. This data is then compiled into the resulting statically
linked `manager` binary.

```shell
make update-bindata
git commit -m "make update-bindata" ./pkg/bindata/bindata.go
```

## 7. Update any necessary RBAC permissions

If Gatekeeper has introduced any new or modified existing RBAC permissions
required or no longer required for it to run successfully, the operator will
similarly need to make changes to add or remove the same permissions. This is
done via updates to the [kubebuilder RBAC markers specified in the operator
controller](https://github.com/gatekeeper/gatekeeper-operator/blob/44420118530000ff25264adc4229b4490013abed/controllers/gatekeeper_controller.go#L133-L163).

Once completed, make sure to update the operator's RBAC manifests by running:

```shell
make manifests
git commit -m "make manifests" ./config/rbac/
```

## 8. Update OLM bundle

After previous changes have been made, you'll need to update the OLM bundle
manifests to incorporate those changes. To do this, run:

```shell
make bundle
git commit -m "make bundle" ./bundle/
```

## 9. Make any necessary operator controller changes

Update any necessary logic in the operator controller as a result of the newly
imported Gatekeeper manifests.

## 10. Update unit and e2e tests

Make sure to add or modify any unit and e2e tests as a result of any operator
controller changes.
