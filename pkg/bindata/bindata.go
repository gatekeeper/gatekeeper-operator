// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/gatekeeper/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml
// config/gatekeeper/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper/apps_v1_deployment_gatekeeper-audit.yaml
// config/gatekeeper/apps_v1_deployment_gatekeeper-controller-manager.yaml
// config/gatekeeper/openshift/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml
// config/gatekeeper/policy_v1beta1_poddisruptionbudget_gatekeeper-controller-manager.yaml
// config/gatekeeper/policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml
// config/gatekeeper/v1_namespace_gatekeeper-system.yaml
// config/gatekeeper/v1_resourcequota_gatekeeper-critical-pods.yaml
// config/gatekeeper/v1_secret_gatekeeper-webhook-server-cert.yaml
// config/gatekeeper/v1_service_gatekeeper-webhook-service.yaml
// config/gatekeeper/v1_serviceaccount_gatekeeper-admin.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  creationTimestamp: null
  name: gatekeeper-mutating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/mutate
  failurePolicy: Ignore
  matchPolicy: Exact
  name: mutation.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: admission.gatekeeper.sh/ignore
      operator: DoesNotExist
  rules:
  - apiGroups:
    - '*'
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - '*'
  sideEffects: None
  timeoutSeconds: 3
`)

func configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYamlBytes() ([]byte, error) {
	return _configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml, nil
}

func configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml() (*asset, error) {
	bytes, err := configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-validating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admit
  failurePolicy: Ignore
  matchPolicy: Exact
  name: validation.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: admission.gatekeeper.sh/ignore
      operator: DoesNotExist
  rules:
  - apiGroups:
    - '*'
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - '*'
  sideEffects: None
  timeoutSeconds: 3
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admitlabel
  failurePolicy: Fail
  matchPolicy: Exact
  name: check-ignore-label.gatekeeper.sh
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - namespaces
  sideEffects: None
  timeoutSeconds: 3
`)

func configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes() ([]byte, error) {
	return _configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, nil
}

func configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml() (*asset, error) {
	bytes, err := configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  name: assign.mutations.gatekeeper.sh
spec:
  group: mutations.gatekeeper.sh
  names:
    kind: Assign
    listKind: AssignList
    plural: assign
    singular: assign
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Assign is the Schema for the assign API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignSpec defines the desired state of Assign
            properties:
              applyTo:
                description: 'INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
                  Important: Run "make" to regenerate code after modifying this file'
                items:
                  description: ApplyTo determines what GVKs items the mutation should
                    apply to. Globs are not allowed.
                  properties:
                    groups:
                      items:
                        type: string
                      type: array
                    kinds:
                      items:
                        type: string
                      type: array
                    versions:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              location:
                type: string
              match:
                description: Match selects objects to apply mutations to.
                properties:
                  excludedNamespaces:
                    items:
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: A label selector is a label query over a set of resources.
                      The result of matchLabels and matchExpressions are ANDed. An
                      empty label selector matches all objects. A null label selector
                      matches no objects.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaceSelector:
                    description: A label selector is a label query over a set of resources.
                      The result of matchLabels and matchExpressions are ANDed. An
                      empty label selector matches all objects. A null label selector
                      matches no objects.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    items:
                      type: string
                    type: array
                  scope:
                    description: ResourceScope is an enum defining the different scopes
                      available to a custom resource
                    type: string
                type: object
              parameters:
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                  assignIf:
                    description: once https://github.com/kubernetes-sigs/controller-tools/pull/528
                      is merged, we can use an actual object
                    type: object
                  pathTests:
                    items:
                      description: "PathTest allows the user to customize how the
                        mutation works if parent paths are missing. It traverses the
                        list in order. All sub paths are tested against the provided
                        condition, if the test fails, the mutation is not applied.
                        All ` + "`" + `subPath` + "`" + ` entries must be a prefix of ` + "`" + `location` + "`" + `. Any
                        glob characters will take on the same value as was used to
                        expand the matching glob in ` + "`" + `location` + "`" + `. \n Available Tests:
                        * MustExist    - the path must exist or do not mutate * MustNotExist
                        - the path must not exist or do not mutate"
                      properties:
                        condition:
                          description: Condition describes whether the path either
                            MustExist or MustNotExist in the original object
                          enum:
                          - MustExist
                          - MustNotExist
                          type: string
                        subPath:
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: AssignStatus defines the observed state of Assign
            properties:
              byPod:
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system
                        properties:
                          message:
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  name: assignmetadata.mutations.gatekeeper.sh
spec:
  group: mutations.gatekeeper.sh
  names:
    kind: AssignMetadata
    listKind: AssignMetadataList
    plural: assignmetadata
    singular: assignmetadata
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AssignMetadata is the Schema for the assignmetadata API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: AssignMetadataSpec defines the desired state of AssignMetadata
            properties:
              location:
                type: string
              match:
                description: Match selects objects to apply mutations to.
                properties:
                  excludedNamespaces:
                    items:
                      type: string
                    type: array
                  kinds:
                    items:
                      description: Kinds accepts a list of objects with apiGroups
                        and kinds fields that list the groups/kinds of objects to
                        which the mutation will apply. If multiple groups/kinds objects
                        are specified, only one match is needed for the resource to
                        be in scope.
                      properties:
                        apiGroups:
                          description: APIGroups is the API groups the resources belong
                            to. '*' is all groups. If '*' is present, the length of
                            the slice must be one. Required.
                          items:
                            type: string
                          type: array
                        kinds:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  labelSelector:
                    description: A label selector is a label query over a set of resources.
                      The result of matchLabels and matchExpressions are ANDed. An
                      empty label selector matches all objects. A null label selector
                      matches no objects.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaceSelector:
                    description: A label selector is a label query over a set of resources.
                      The result of matchLabels and matchExpressions are ANDed. An
                      empty label selector matches all objects. A null label selector
                      matches no objects.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                  namespaces:
                    items:
                      type: string
                    type: array
                  scope:
                    description: ResourceScope is an enum defining the different scopes
                      available to a custom resource
                    type: string
                type: object
              parameters:
                properties:
                  assign:
                    description: Assign.value holds the value to be assigned
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                type: object
            type: object
          status:
            description: AssignMetadataStatus defines the observed state of AssignMetadata
            properties:
              byPod:
                description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
                  of cluster Important: Run "make" to regenerate code after modifying
                  this file'
                items:
                  description: MutatorPodStatusStatus defines the observed state of
                    MutatorPodStatus
                  properties:
                    enforced:
                      type: boolean
                    errors:
                      items:
                        description: MutatorError represents a single error caught
                          while adding a mutator to a system
                        properties:
                          message:
                            type: string
                        required:
                        - message
                        type: object
                      type: array
                    id:
                      type: string
                    mutatorUID:
                      description: Storing the mutator UID allows us to detect drift,
                        such as when a mutator has been recreated after its CRD was
                        deleted out from under it, interrupting the watch
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                    operations:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: configs.config.gatekeeper.sh
spec:
  group: config.gatekeeper.sh
  names:
    kind: Config
    listKind: ConfigList
    plural: configs
    singular: config
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Config is the Schema for the configs API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConfigSpec defines the desired state of Config
            properties:
              match:
                description: Configuration for namespace exclusion
                items:
                  properties:
                    excludedNamespaces:
                      items:
                        type: string
                      type: array
                    processes:
                      items:
                        type: string
                      type: array
                  type: object
                type: array
              readiness:
                description: Configuration for readiness tracker
                properties:
                  statsEnabled:
                    type: boolean
                type: object
              sync:
                description: Configuration for syncing k8s objects
                properties:
                  syncOnly:
                    description: If non-empty, only entries on this list will be replicated
                      into OPA
                    items:
                      properties:
                        group:
                          type: string
                        kind:
                          type: string
                        version:
                          type: string
                      type: object
                    type: array
                type: object
              validation:
                description: Configuration for validation
                properties:
                  traces:
                    description: List of requests to trace. Both "user" and "kinds"
                      must be specified
                    items:
                      properties:
                        dump:
                          description: Also dump the state of OPA with the trace.
                            Set to ` + "`" + `All` + "`" + ` to dump everything.
                          type: string
                        kind:
                          description: Only trace requests of the following GroupVersionKind
                          properties:
                            group:
                              type: string
                            kind:
                              type: string
                            version:
                              type: string
                          type: object
                        user:
                          description: Only trace requests from the specified user
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: ConfigStatus defines the observed state of Config
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: constraintpodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintPodStatus
    listKind: ConstraintPodStatusList
    plural: constraintpodstatuses
    singular: constraintpodstatus
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintPodStatus is the Schema for the constraintpodstatuses
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: ConstraintPodStatusStatus defines the observed state of ConstraintPodStatus
            properties:
              constraintUID:
                description: Storing the constraint UID allows us to detect drift,
                  such as when a constraint has been recreated after its CRD was deleted
                  out from under it, interrupting the watch
                type: string
              enforced:
                type: boolean
              errors:
                items:
                  description: Error represents a single error caught while adding
                    a constraint to OPA
                  properties:
                    code:
                      type: string
                    location:
                      type: string
                    message:
                      type: string
                  required:
                  - code
                  - message
                  type: object
                type: array
              id:
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: constrainttemplatepodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintTemplatePodStatus
    listKind: ConstraintTemplatePodStatusList
    plural: constrainttemplatepodstatuses
    singular: constrainttemplatepodstatus
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplatePodStatus is the Schema for the constrainttemplatepodstatuses
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: ConstraintTemplatePodStatusStatus defines the observed state
              of ConstraintTemplatePodStatus
            properties:
              errors:
                items:
                  description: CreateCRDError represents a single error caught during
                    parsing, compiling, etc.
                  properties:
                    code:
                      type: string
                    location:
                      type: string
                    message:
                      type: string
                  required:
                  - code
                  - message
                  type: object
                type: array
              id:
                description: 'Important: Run "make" to regenerate code after modifying
                  this file'
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
              templateUID:
                description: UID is a type that holds unique ID values, including
                  UUIDs.  Because we don't ONLY use UUIDs, this is an alias to string.  Being
                  a type captures intent and helps make sure that UIDs and names do
                  not get conflated.
                type: string
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: constrainttemplates.templates.gatekeeper.sh
spec:
  group: templates.gatekeeper.sh
  names:
    kind: ConstraintTemplate
    listKind: ConstraintTemplateList
    plural: constrainttemplates
    singular: constrainttemplate
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplate is the Schema for the constrainttemplates
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConstraintTemplateSpec defines the desired state of ConstraintTemplate
            properties:
              crd:
                properties:
                  spec:
                    properties:
                      names:
                        properties:
                          kind:
                            type: string
                          shortNames:
                            items:
                              type: string
                            type: array
                        type: object
                      validation:
                        properties:
                          openAPIV3Schema:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                    type: object
                type: object
              targets:
                items:
                  properties:
                    libs:
                      items:
                        type: string
                      type: array
                    rego:
                      type: string
                    target:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: ConstraintTemplateStatus defines the observed state of ConstraintTemplate
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the observed state of ConstraintTemplate
                    as seen by an individual controller
                  properties:
                    errors:
                      items:
                        description: CreateCRDError represents a single error caught
                          during parsing, compiling, etc.
                        properties:
                          code:
                            type: string
                          location:
                            type: string
                          message:
                            type: string
                        required:
                        - code
                        - message
                        type: object
                      type: array
                    id:
                      description: a unique identifier for the pod that wrote the
                        status
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                type: array
              created:
                type: boolean
            type: object
        type: object
    served: true
    storage: false
    subresources:
      status: {}
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: ConstraintTemplate is the Schema for the constrainttemplates
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ConstraintTemplateSpec defines the desired state of ConstraintTemplate
            properties:
              crd:
                properties:
                  spec:
                    properties:
                      names:
                        properties:
                          kind:
                            type: string
                          shortNames:
                            items:
                              type: string
                            type: array
                        type: object
                      validation:
                        properties:
                          openAPIV3Schema:
                            type: object
                            x-kubernetes-preserve-unknown-fields: true
                        type: object
                    type: object
                type: object
              targets:
                items:
                  properties:
                    libs:
                      items:
                        type: string
                      type: array
                    rego:
                      type: string
                    target:
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: ConstraintTemplateStatus defines the observed state of ConstraintTemplate
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the observed state of ConstraintTemplate
                    as seen by an individual controller
                  properties:
                    errors:
                      items:
                        description: CreateCRDError represents a single error caught
                          during parsing, compiling, etc.
                        properties:
                          code:
                            type: string
                          location:
                            type: string
                          message:
                            type: string
                        required:
                        - code
                        - message
                        type: object
                      type: array
                    id:
                      description: a unique identifier for the pod that wrote the
                        status
                      type: string
                    observedGeneration:
                      format: int64
                      type: integer
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                type: array
              created:
                type: boolean
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  name: mutatorpodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: MutatorPodStatus
    listKind: MutatorPodStatusList
    plural: mutatorpodstatuses
    singular: mutatorpodstatus
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: MutatorPodStatus is the Schema for the mutationpodstatuses API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: MutatorPodStatusStatus defines the observed state of MutatorPodStatus
            properties:
              enforced:
                type: boolean
              errors:
                items:
                  description: MutatorError represents a single error caught while
                    adding a mutator to a system
                  properties:
                    message:
                      type: string
                  required:
                  - message
                  type: object
                type: array
              id:
                type: string
              mutatorUID:
                description: Storing the mutator UID allows us to detect drift, such
                  as when a mutator has been recreated after its CRD was deleted out
                  from under it, interrupting the watch
                type: string
              observedGeneration:
                format: int64
                type: integer
              operations:
                items:
                  type: string
                type: array
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApps_v1_deployment_gatekeeperAuditYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: audit-controller
    gatekeeper.sh/operation: audit
    gatekeeper.sh/system: "yes"
  name: gatekeeper-audit
  namespace: gatekeeper-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: audit-controller
      gatekeeper.sh/operation: audit
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: audit-controller
        gatekeeper.sh/operation: audit
        gatekeeper.sh/system: "yes"
    spec:
      automountServiceAccountToken: true
      containers:
      - args:
        - --operation=audit
        - --operation=status
        - --logtostderr
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        image: openpolicyagent/gatekeeper:v3.6.0
        imagePullPolicy: Always
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
        name: manager
        ports:
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 9090
          name: healthz
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 256Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - all
          readOnlyRootFilesystem: true
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
`)

func configGatekeeperApps_v1_deployment_gatekeeperAuditYamlBytes() ([]byte, error) {
	return _configGatekeeperApps_v1_deployment_gatekeeperAuditYaml, nil
}

func configGatekeeperApps_v1_deployment_gatekeeperAuditYaml() (*asset, error) {
	bytes, err := configGatekeeperApps_v1_deployment_gatekeeperAuditYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apps_v1_deployment_gatekeeper-audit.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  replicas: 3
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/operation: webhook
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: controller-manager
        gatekeeper.sh/operation: webhook
        gatekeeper.sh/system: "yes"
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: gatekeeper.sh/operation
                  operator: In
                  values:
                  - webhook
              topologyKey: kubernetes.io/hostname
            weight: 100
      automountServiceAccountToken: true
      containers:
      - args:
        - --port=8443
        - --logtostderr
        - --exempt-namespace=gatekeeper-system
        - --operation=webhook
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        image: openpolicyagent/gatekeeper:v3.6.0
        imagePullPolicy: Always
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
        name: manager
        ports:
        - containerPort: 8443
          name: webhook-server
          protocol: TCP
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 9090
          name: healthz
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
        resources:
          limits:
            cpu: 1000m
            memory: 512Mi
          requests:
            cpu: 100m
            memory: 256Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - all
          readOnlyRootFilesystem: true
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
        volumeMounts:
        - mountPath: /certs
          name: cert
          readOnly: true
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: gatekeeper-webhook-server-cert
`)

func configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYamlBytes() ([]byte, error) {
	return _configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml, nil
}

func configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml() (*asset, error) {
	bytes, err := configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apps_v1_deployment_gatekeeper-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-role
  namespace: gatekeeper-system
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - security.openshift.io
  resourceNames:
    - anyuid
  resources:
    - securitycontextconstraints
  verbs:
    - use
`)

func configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes() ([]byte, error) {
	return _configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, nil
}

func configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml() (*asset, error) {
	bytes, err := configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/openshift/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYaml = []byte(`apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  minAvailable: 1
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/operation: webhook
      gatekeeper.sh/system: "yes"
`)

func configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYamlBytes() ([]byte, error) {
	return _configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYaml, nil
}

func configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYaml() (*asset, error) {
	bytes, err := configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/policy_v1beta1_poddisruptionbudget_gatekeeper-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml = []byte(`apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-admin
spec:
  allowPrivilegeEscalation: false
  fsGroup:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  requiredDropCapabilities:
  - ALL
  runAsUser:
    rule: MustRunAsNonRoot
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  volumes:
  - configMap
  - projected
  - secret
  - downwardAPI
`)

func configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYamlBytes() ([]byte, error) {
	return _configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml, nil
}

func configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml() (*asset, error) {
	bytes, err := configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-role
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - constraints.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - mutations.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - policy
  resourceNames:
  - gatekeeper-admin
  resources:
  - podsecuritypolicies
  verbs:
  - use
- apiGroups:
  - status.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates/finalizers
  verbs:
  - delete
  - get
  - patch
  - update
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - gatekeeper-validating-webhook-configuration
  resources:
  - validatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - gatekeeper-mutating-webhook-configuration
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
`)

func configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYamlBytes() ([]byte, error) {
	return _configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml, nil
}

func configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml() (*asset, error) {
	bytes, err := configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYamlBytes() ([]byte, error) {
	return _configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml, nil
}

func configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml() (*asset, error) {
	bytes, err := configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-role
  namespace: gatekeeper-system
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
`)

func configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes() ([]byte, error) {
	return _configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, nil
}

func configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml() (*asset, error) {
	bytes, err := configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-manager-rolebinding
  namespace: gatekeeper-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYamlBytes() ([]byte, error) {
	return _configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml, nil
}

func configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml() (*asset, error) {
	bytes, err := configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperV1_namespace_gatekeeperSystemYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    admission.gatekeeper.sh/ignore: no-self-managing
    control-plane: controller-manager
    gatekeeper.sh/system: "yes"
  name: gatekeeper-system
`)

func configGatekeeperV1_namespace_gatekeeperSystemYamlBytes() ([]byte, error) {
	return _configGatekeeperV1_namespace_gatekeeperSystemYaml, nil
}

func configGatekeeperV1_namespace_gatekeeperSystemYaml() (*asset, error) {
	bytes, err := configGatekeeperV1_namespace_gatekeeperSystemYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/v1_namespace_gatekeeper-system.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYaml = []byte(`apiVersion: v1
kind: ResourceQuota
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-critical-pods
  namespace: gatekeeper-system
spec:
  hard:
    pods: 100
  scopeSelector:
    matchExpressions:
    - operator: In
      scopeName: PriorityClass
      values:
      - system-cluster-critical
`)

func configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYamlBytes() ([]byte, error) {
	return _configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYaml, nil
}

func configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYaml() (*asset, error) {
	bytes, err := configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/v1_resourcequota_gatekeeper-critical-pods.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-webhook-server-cert
  namespace: gatekeeper-system
`)

func configGatekeeperV1_secret_gatekeeperWebhookServerCertYamlBytes() ([]byte, error) {
	return _configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml, nil
}

func configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml() (*asset, error) {
	bytes, err := configGatekeeperV1_secret_gatekeeperWebhookServerCertYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/v1_secret_gatekeeper-webhook-server-cert.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperV1_service_gatekeeperWebhookServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-webhook-service
  namespace: gatekeeper-system
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    control-plane: controller-manager
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
`)

func configGatekeeperV1_service_gatekeeperWebhookServiceYamlBytes() ([]byte, error) {
	return _configGatekeeperV1_service_gatekeeperWebhookServiceYaml, nil
}

func configGatekeeperV1_service_gatekeeperWebhookServiceYaml() (*asset, error) {
	bytes, err := configGatekeeperV1_service_gatekeeperWebhookServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/v1_service_gatekeeper-webhook-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperV1_serviceaccount_gatekeeperAdminYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-admin
  namespace: gatekeeper-system
`)

func configGatekeeperV1_serviceaccount_gatekeeperAdminYamlBytes() ([]byte, error) {
	return _configGatekeeperV1_serviceaccount_gatekeeperAdminYaml, nil
}

func configGatekeeperV1_serviceaccount_gatekeeperAdminYaml() (*asset, error) {
	bytes, err := configGatekeeperV1_serviceaccount_gatekeeperAdminYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/v1_serviceaccount_gatekeeper-admin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/gatekeeper/admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml":     configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml,
	"config/gatekeeper/admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml":                            configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml":                    configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml":                   configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper/apps_v1_deployment_gatekeeper-audit.yaml":                                                                        configGatekeeperApps_v1_deployment_gatekeeperAuditYaml,
	"config/gatekeeper/apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                           configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml,
	"config/gatekeeper/openshift/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                        configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml,
	"config/gatekeeper/policy_v1beta1_poddisruptionbudget_gatekeeper-controller-manager.yaml":                                           configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYaml,
	"config/gatekeeper/policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml":                                                          configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":                                           configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml":                             configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                                  configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":                                    configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper/v1_namespace_gatekeeper-system.yaml":                                                                             configGatekeeperV1_namespace_gatekeeperSystemYaml,
	"config/gatekeeper/v1_resourcequota_gatekeeper-critical-pods.yaml":                                                                  configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYaml,
	"config/gatekeeper/v1_secret_gatekeeper-webhook-server-cert.yaml":                                                                   configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml,
	"config/gatekeeper/v1_service_gatekeeper-webhook-service.yaml":                                                                      configGatekeeperV1_service_gatekeeperWebhookServiceYaml,
	"config/gatekeeper/v1_serviceaccount_gatekeeper-admin.yaml":                                                                         configGatekeeperV1_serviceaccount_gatekeeperAdminYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": {nil, map[string]*bintree{
		"gatekeeper": {nil, map[string]*bintree{
			"admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_gatekeeper-mutating-webhook-configuration.yaml":     {configGatekeeperAdmissionregistrationK8sIo_v1_mutatingwebhookconfiguration_gatekeeperMutatingWebhookConfigurationYaml, map[string]*bintree{}},
			"admissionregistration.k8s.io_v1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": {configGatekeeperAdmissionregistrationK8sIo_v1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_assign.mutations.gatekeeper.sh.yaml":                            {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignMutationsGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_assignmetadata.mutations.gatekeeper.sh.yaml":                    {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_assignmetadataMutationsGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_configsConfigGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1_customresourcedefinition_mutatorpodstatuses.status.gatekeeper.sh.yaml":                   {configGatekeeperApiextensionsK8sIo_v1_customresourcedefinition_mutatorpodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-audit.yaml":                                                                        {configGatekeeperApps_v1_deployment_gatekeeperAuditYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                           {configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml, map[string]*bintree{}},
			"openshift": {nil, map[string]*bintree{
				"rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml": {configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			}},
			"policy_v1beta1_poddisruptionbudget_gatekeeper-controller-manager.yaml":               {configGatekeeperPolicy_v1beta1_poddisruptionbudget_gatekeeperControllerManagerYaml, map[string]*bintree{}},
			"policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml":                              {configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":               {configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml": {configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                      {configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":        {configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
			"v1_namespace_gatekeeper-system.yaml":                                                 {configGatekeeperV1_namespace_gatekeeperSystemYaml, map[string]*bintree{}},
			"v1_resourcequota_gatekeeper-critical-pods.yaml":                                      {configGatekeeperV1_resourcequota_gatekeeperCriticalPodsYaml, map[string]*bintree{}},
			"v1_secret_gatekeeper-webhook-server-cert.yaml":                                       {configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml, map[string]*bintree{}},
			"v1_service_gatekeeper-webhook-service.yaml":                                          {configGatekeeperV1_service_gatekeeperWebhookServiceYaml, map[string]*bintree{}},
			"v1_serviceaccount_gatekeeper-admin.yaml":                                             {configGatekeeperV1_serviceaccount_gatekeeperAdminYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
