// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/gatekeeper/admissionregistration.k8s.io_v1beta1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml
// config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_configs.config.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml
// config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml
// config/gatekeeper/apps_v1_deployment_gatekeeper-audit.yaml
// config/gatekeeper/apps_v1_deployment_gatekeeper-controller-manager.yaml
// config/gatekeeper/openshift/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml
// config/gatekeeper/policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml
// config/gatekeeper/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml
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

var _configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  labels:
    gatekeeper.sh/system: "yes"
  name: gatekeeper-validating-webhook-configuration
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admit
  failurePolicy: Ignore
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
- clientConfig:
    caBundle: Cg==
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admitlabel
  failurePolicy: Fail
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

func configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes() ([]byte, error) {
	return _configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, nil
}

func configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml() (*asset, error) {
	bytes, err := configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/admissionregistration.k8s.io_v1beta1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.3.0
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
  validation:
    openAPIV3Schema:
      description: Config is the Schema for the configs API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
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
                  description: If non-empty, only entries on this list will be replicated into OPA
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
                  description: List of requests to trace. Both "user" and "kinds" must be specified
                  items:
                    properties:
                      dump:
                        description: Also dump the state of OPA with the trace. Set to ` + "`" + `All` + "`" + ` to dump everything.
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
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_configs.config.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.3.0
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
  validation:
    openAPIV3Schema:
      description: ConstraintPodStatus is the Schema for the constraintpodstatuses API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        status:
          description: ConstraintPodStatusStatus defines the observed state of ConstraintPodStatus
          properties:
            constraintUID:
              description: Storing the constraint UID allows us to detect drift, such as when a constraint has been recreated after its CRD was deleted out from under it, interrupting the watch
              type: string
            enforced:
              type: boolean
            errors:
              items:
                description: Error represents a single error caught while adding a constraint to OPA
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
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.3.0
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
  validation:
    openAPIV3Schema:
      description: ConstraintTemplatePodStatus is the Schema for the constrainttemplatepodstatuses API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        status:
          description: ConstraintTemplatePodStatusStatus defines the observed state of ConstraintTemplatePodStatus
          properties:
            errors:
              items:
                description: CreateCRDError represents a single error caught during parsing, compiling, etc.
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
              description: 'Important: Run "make" to regenerate code after modifying this file'
              type: string
            observedGeneration:
              format: int64
              type: integer
            operations:
              items:
                type: string
              type: array
            templateUID:
              description: UID is a type that holds unique ID values, including UUIDs.  Because we don't ONLY use UUIDs, this is an alias to string.  Being a type captures intent and helps make sure that UIDs and names do not get conflated.
              type: string
          type: object
      type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
    gatekeeper.sh/system: "yes"
  name: constrainttemplates.templates.gatekeeper.sh
spec:
  group: templates.gatekeeper.sh
  names:
    kind: ConstraintTemplate
    plural: constrainttemplates
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
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
          properties:
            byPod:
              items:
                properties:
                  errors:
                    items:
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
                    description: a unique identifier for the pod that wrote the status
                    type: string
                  observedGeneration:
                    format: int64
                    type: integer
                type: object
              type: array
            created:
              type: boolean
          type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
  - name: v1alpha1
    served: true
    storage: false
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes() ([]byte, error) {
	return _configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, nil
}

func configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml() (*asset, error) {
	bytes, err := configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configGatekeeperApps_v1_deployment_gatekeeperAuditYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
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
        image: openpolicyagent/gatekeeper:v3.2.1
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
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
      nodeSelector:
        kubernetes.io/os: linux
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
        image: openpolicyagent/gatekeeper:v3.2.1
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
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
        volumeMounts:
        - mountPath: /certs
          name: cert
          readOnly: true
      nodeSelector:
        kubernetes.io/os: linux
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
	"config/gatekeeper/admissionregistration.k8s.io_v1beta1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml,
	"config/gatekeeper/apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml,
	"config/gatekeeper/apps_v1_deployment_gatekeeper-audit.yaml":                                                                             configGatekeeperApps_v1_deployment_gatekeeperAuditYaml,
	"config/gatekeeper/apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                                configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml,
	"config/gatekeeper/openshift/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                             configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml,
	"config/gatekeeper/policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml":                                                               configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":                                                configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml":                                  configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                                                       configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml,
	"config/gatekeeper/rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":                                         configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml,
	"config/gatekeeper/v1_secret_gatekeeper-webhook-server-cert.yaml":                                                                        configGatekeeperV1_secret_gatekeeperWebhookServerCertYaml,
	"config/gatekeeper/v1_service_gatekeeper-webhook-service.yaml":                                                                           configGatekeeperV1_service_gatekeeperWebhookServiceYaml,
	"config/gatekeeper/v1_serviceaccount_gatekeeper-admin.yaml":                                                                              configGatekeeperV1_serviceaccount_gatekeeperAdminYaml,
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
			"admissionregistration.k8s.io_v1beta1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml": {configGatekeeperAdmissionregistrationK8sIo_v1beta1_validatingwebhookconfiguration_gatekeeperValidatingWebhookConfigurationYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1beta1_customresourcedefinition_configs.config.gatekeeper.sh.yaml":                              {configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_configsConfigGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1beta1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml":                {configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constraintpodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml":        {configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatepodstatusesStatusGatekeeperShYaml, map[string]*bintree{}},
			"apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml":               {configGatekeeperApiextensionsK8sIo_v1beta1_customresourcedefinition_constrainttemplatesTemplatesGatekeeperShYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-audit.yaml":                                                                             {configGatekeeperApps_v1_deployment_gatekeeperAuditYaml, map[string]*bintree{}},
			"apps_v1_deployment_gatekeeper-controller-manager.yaml":                                                                {configGatekeeperApps_v1_deployment_gatekeeperControllerManagerYaml, map[string]*bintree{}},
			"openshift": {nil, map[string]*bintree{
				"rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml": {configGatekeeperOpenshiftRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			}},
			"policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml":                              {configGatekeeperPolicy_v1beta1_podsecuritypolicy_gatekeeperAdminYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml":               {configGatekeeperRbacAuthorizationK8sIo_v1_clusterrole_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml": {configGatekeeperRbacAuthorizationK8sIo_v1_clusterrolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml":                      {configGatekeeperRbacAuthorizationK8sIo_v1_role_gatekeeperManagerRoleYaml, map[string]*bintree{}},
			"rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml":        {configGatekeeperRbacAuthorizationK8sIo_v1_rolebinding_gatekeeperManagerRolebindingYaml, map[string]*bintree{}},
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
