# Releasing Gatekeeper Operator

The following steps need to be done in order to successfully automate the
release of the Gatekeeper Operator using the GitHub Actions release workflow.

**NOTE: This assumes that your git remote name for this repository is named
`upstream` and that the remote name for your fork is named `origin`.**

1. Make sure your clone is up-to-date:
    ```shell
    git fetch --prune upstream
    ```
1. Set the desired version being released:
    ```shell
    export RELEASE_VERSION=v0.0.1
    ```
1. Checkout a new branch based on `upstream/master`:
    ```shell
    git checkout -b release-${RELEASE_VERSION} --no-track upstream/master
    ```
1. Update the version of the operator in the Makefile:
    ```shell
    sed -i "s/^VERSION ?= .*/VERSION ?= ${RELEASE_VERSION}/" Makefile
    ```
1. Update the release manifest:
    ```shell
    make release VERSION=${RELEASE_VERSION}
    ```
1. Update bundle:
    ```shell
    make bundle
    ```
1. Commit above changes:
    ```shell
    git commit -m "Release ${RELEASE_VERSION}" Makefile ./deploy/ ./bundle/ ./config/manager/kustomization.yaml
    ```
1. Push the changes in the branch to your fork:
    ```shell
    git push -u origin release-${RELEASE_VERSION}
    ```
1. Create a PR with the above changes and merge it. If using the `gh` [GitHub
   CLI](https://cli.github.com/) you can create the PR using:
   ```shell
   gh pr create --repo gatekeeper/gatekeeper-operator --title "Release ${RELEASE_VERSION}" --body ""
   ```
1. After the PR is merged, fetch the new commits:
    ```shell
    git fetch --prune upstream
    ```
1. Create and push a new release tag. This will trigger the GitHub actions
   release workflow to build and push the release image and create a new
   release on GitHub. Note that `upstream` is used as the remote name for this
   repository:
    ```shell
    git tag -a -m "Release ${RELEASE_VERSION}" ${RELEASE_VERSION} upstream/master
    git push upstream ${RELEASE_VERSION}
    ```
