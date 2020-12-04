/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/RHsyseng/operator-utils/pkg/utils/openshift"
	"github.com/go-logr/logr"
	"github.com/openshift/library-go/pkg/manifest"
	"github.com/pkg/errors"
	admregv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	operatorv1alpha1 "github.com/gatekeeper/gatekeeper-operator/api/v1alpha1"
	"github.com/gatekeeper/gatekeeper-operator/controllers/merge"
	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
)

var (
	defaultGatekeeperCrName        = "gatekeeper"
	openshiftAssetsDir             = "openshift/"
	RoleFile                       = "rbac.authorization.k8s.io_v1_role_gatekeeper-manager-role.yaml"
	AuditFile                      = "apps_v1_deployment_gatekeeper-audit.yaml"
	WebhookFile                    = "apps_v1_deployment_gatekeeper-controller-manager.yaml"
	ClusterRoleBindingFile         = "rbac.authorization.k8s.io_v1_clusterrolebinding_gatekeeper-manager-rolebinding.yaml"
	RoleBindingFile                = "rbac.authorization.k8s.io_v1_rolebinding_gatekeeper-manager-rolebinding.yaml"
	ValidatingWebhookConfiguration = "admissionregistration.k8s.io_v1beta1_validatingwebhookconfiguration_gatekeeper-validating-webhook-configuration.yaml"
	orderedStaticAssets            = []string{
		"apiextensions.k8s.io_v1beta1_customresourcedefinition_configs.config.gatekeeper.sh.yaml",
		"apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplates.templates.gatekeeper.sh.yaml",
		"apiextensions.k8s.io_v1beta1_customresourcedefinition_constrainttemplatepodstatuses.status.gatekeeper.sh.yaml",
		"apiextensions.k8s.io_v1beta1_customresourcedefinition_constraintpodstatuses.status.gatekeeper.sh.yaml",
		"v1_secret_gatekeeper-webhook-server-cert.yaml",
		"v1_serviceaccount_gatekeeper-admin.yaml",
		"policy_v1beta1_podsecuritypolicy_gatekeeper-admin.yaml",
		"rbac.authorization.k8s.io_v1_clusterrole_gatekeeper-manager-role.yaml",
		ClusterRoleBindingFile,
		RoleFile,
		RoleBindingFile,
		AuditFile,
		WebhookFile,
		"v1_service_gatekeeper-webhook-service.yaml",
		ValidatingWebhookConfiguration,
	}
	ValidationGatekeeperWebhook = "validation.gatekeeper.sh"
)

const (
	gatekeeperFinalizer         = "finalizer.operator.gatekeeper.sh"
	managerContainer            = "manager"
	LogLevelArg                 = "--log-level"
	AuditIntervalArg            = "--audit-interval"
	ConstraintViolationLimitArg = "--constraint-violations-limit"
	AuditFromCacheArg           = "--audit-from-cache"
	AuditChunkSizeArg           = "--audit-chunk-size"
	EmitAuditEventsArg          = "--emit-audit-events"
	EmitAdmissionEventsArg      = "--emit-admission-events"
	ExemptNamespaceArg          = "--exempt-namespace"
)

// GatekeeperReconciler reconciles a Gatekeeper object
type GatekeeperReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// Gatekeeper Operator RBAC permissions to manager Gatekeeper custom resource
// +kubebuilder:rbac:groups=operator.gatekeeper.sh,namespace="system",resources=gatekeepers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.gatekeeper.sh,namespace="system",resources=gatekeepers/finalizers,verbs=get;update;patch;delete
// +kubebuilder:rbac:groups=operator.gatekeeper.sh,namespace="system",resources=gatekeepers/status,verbs=get;update;patch

// Gatekeeper Operator RBAC permissions to deploy Gatekeeper. Many of these
// RBAC permissions are needed because the operator must have the permissions
// to grant Gatekeeper its required RBAC permissions.

// Cluster Scoped
// +kubebuilder:rbac:groups=*,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.gatekeeper.sh,resources=configs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.gatekeeper.sh,resources=configs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=constraints.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=policy,resources=podsecuritypolicies,verbs=create;delete;update;use
// +kubebuilder:rbac:groups=status.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=templates.gatekeeper.sh,resources=constrainttemplates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=templates.gatekeeper.sh,resources=constrainttemplates/finalizers,verbs=get;update;patch;delete
// +kubebuilder:rbac:groups=templates.gatekeeper.sh,resources=constrainttemplates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

// Namespace Scoped
// +kubebuilder:rbac:groups=core,namespace="system",resources=secrets;serviceaccounts;services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,namespace="system",resources=roles;rolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,namespace="system",resources=deployments,verbs=get;list;watch;create;update;patch;delete

func (r *GatekeeperReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("gatekeeper", req.NamespacedName)
	logger.Info("Reconciling Gatekeeper")

	cfg, err := config.GetConfig()
	if err != nil {
		return ctrl.Result{}, err
	}
	platformName, err := openshift.GetPlatformName(cfg)
	if err != nil {
		return ctrl.Result{}, err
	}

	if req.Name != defaultGatekeeperCrName {
		err := fmt.Errorf("Gatekeeper resource name must be '%s'", defaultGatekeeperCrName)
		logger.Error(err, "Invalid Gatekeeper resource name")
		// Return success to avoid requeue
		return ctrl.Result{}, nil
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{}
	err = r.Get(ctx, req.NamespacedName, gatekeeper)
	if err != nil {
		if apierrors.IsNotFound(err) {

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	isGatekeeperMarkedToBeDeleted := gatekeeper.GetDeletionTimestamp() != nil
	if isGatekeeperMarkedToBeDeleted {
		if sets.NewString(gatekeeper.GetFinalizers()...).Has(gatekeeperFinalizer) {

			if err := r.finalizeGatekeeper(logger, gatekeeper); err != nil {
				return ctrl.Result{}, err
			}

			controllerutil.RemoveFinalizer(gatekeeper, gatekeeperFinalizer)
			err := r.Update(ctx, gatekeeper)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !sets.NewString(gatekeeper.GetFinalizers()...).Has(gatekeeperFinalizer) {
		if err := r.addFinalizer(logger, gatekeeper); err != nil {
			return ctrl.Result{}, err
		}
	}

	err = r.deployGatekeeperResources(gatekeeper, platformName)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "Unable to deploy Gatekeeper resources")
	}

	return ctrl.Result{}, nil
}

func (r *GatekeeperReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.Gatekeeper{}).
		WithEventFilter(predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {
				oldGeneration := e.MetaOld.GetGeneration()
				newGeneration := e.MetaNew.GetGeneration()

				return oldGeneration != newGeneration
			},
			DeleteFunc: func(e event.DeleteEvent) bool {

				return false
			},
		}).
		Complete(r)
}

func (r *GatekeeperReconciler) deployGatekeeperResources(gatekeeper *operatorv1alpha1.Gatekeeper, platformName string) error {
	for _, a := range getStaticAssets(gatekeeper) {
		if a == RoleFile && platformName == "OpenShift" {
			a = openshiftAssetsDir + a
		}
		manifest, err := util.GetManifest(a)
		if err != nil {
			return err
		}
		if err = crOverrides(gatekeeper, a, manifest, (platformName == "OpenShift")); err != nil {
			return err
		}

		if err = r.updateOrCreateResource(manifest, gatekeeper); err != nil {
			return err
		}
	}
	return nil
}

func getStaticAssets(gatekeeper *operatorv1alpha1.Gatekeeper) []string {
	if gatekeeper.Spec.ValidatingWebhook == nil || *gatekeeper.Spec.ValidatingWebhook == operatorv1alpha1.WebhookEnabled {
		return orderedStaticAssets
	}
	assets := make([]string, 0)
	for _, a := range orderedStaticAssets {
		if a != ValidatingWebhookConfiguration {
			assets = append(assets, a)
		}
	}
	return assets
}

func (r *GatekeeperReconciler) updateOrCreateResource(manifest *manifest.Manifest, gatekeeper *operatorv1alpha1.Gatekeeper) error {
	var err error
	ctx := context.Background()
	clusterObj := &unstructured.Unstructured{}
	clusterObj.SetAPIVersion(manifest.Obj.GetAPIVersion())
	clusterObj.SetKind(manifest.Obj.GetKind())

	namespacedName := types.NamespacedName{
		Namespace: manifest.Obj.GetNamespace(),
		Name:      manifest.Obj.GetName(),
	}

	logger := r.Log.WithValues("Gatekeeper resource", namespacedName)

	if manifest.Obj.GetNamespace() != "" {
		err = ctrl.SetControllerReference(gatekeeper, manifest.Obj, r.Scheme)
		if err != nil {
			return errors.Wrapf(err, "Unable to set controller reference for %s", namespacedName)
		}
	}

	err = r.Get(ctx, namespacedName, clusterObj)

	switch {
	case err == nil:
		err = merge.RetainClusterObjectFields(manifest.Obj, clusterObj)
		if err != nil {
			return errors.Wrapf(err, "Unable to retain cluster object fields from %s", namespacedName)
		}

		err = r.Update(ctx, manifest.Obj)
		if err != nil {
			return errors.Wrapf(err, "Error attempting to update resource %s", namespacedName)
		}

		logger.Info(fmt.Sprintf("Updated Gatekeeper resource"))

	case apierrors.IsNotFound(err):
		err = r.Create(ctx, manifest.Obj)
		if err != nil {
			return errors.Wrapf(err, "Error attempting to create resource %s", namespacedName)
		}
		logger.Info(fmt.Sprintf("Created Gatekeeper resource"))

	case err != nil:
		return errors.Wrapf(err, "Error attempting to get resource %s", namespacedName)
	}

	return err
}

func (r *GatekeeperReconciler) finalizeGatekeeper(reqLogger logr.Logger, gatekeeper *operatorv1alpha1.Gatekeeper) error {
	ctx := context.Background()

	var err error
	for _, a := range orderedStaticAssets {
		manifest, err := util.GetManifest(a)
		if err != nil {
			return err
		}

		// Delete cluster scoped resource not owned by the CR
		if manifest.Obj.GetNamespace() == "" {

			err = r.Delete(ctx, manifest.Obj)
			if err != nil && !apierrors.IsNotFound(err) {
				return errors.Wrapf(err, "Error Deleting Finalizer Resource. Kind: '%s'. Name: '%s'", manifest.GVK.Kind, manifest.Obj.GetName())
			}
		}

	}

	reqLogger.Info("Successfully finalized Gatekeeper")
	return err
}

func (r *GatekeeperReconciler) addFinalizer(reqLogger logr.Logger, g *operatorv1alpha1.Gatekeeper) error {
	ctx := context.Background()

	controllerutil.AddFinalizer(g, gatekeeperFinalizer)

	// Update CR
	err := r.Update(ctx, g)
	if err != nil {
		return errors.Wrapf(err, "Failed to update Gatekeeper with finalizer. Name: '%s'", g.Name)
	}
	return nil
}

var commonSpecOverridesFn = []func(*unstructured.Unstructured, operatorv1alpha1.GatekeeperSpec) error{
	setAffinity,
	setNodeSelector,
	setPodAnnotations,
	setTolerations,
	containerOverrides,
}
var commonContainerOverridesFn = []func(map[string]interface{}, operatorv1alpha1.GatekeeperSpec) error{
	setImage,
}

// crOverrides
func crOverrides(gatekeeper *operatorv1alpha1.Gatekeeper, asset string, manifest *manifest.Manifest, isOpenshift bool) error {
	// set current namespace
	if err := setCurrentNamespace(manifest.Obj, asset, gatekeeper.Namespace); err != nil {
		return err
	}
	// audit overrides
	if asset == AuditFile {
		if err := commonOverrides(manifest.Obj, gatekeeper.Spec); err != nil {
			return err
		}
		if err := auditOverrides(manifest.Obj, gatekeeper.Spec.Audit); err != nil {
			return err
		}
		if isOpenshift {
			if err := removeAnnotations(manifest.Obj); err != nil {
				return err
			}
		}
	}
	// webhook overrides
	if asset == WebhookFile {
		if err := commonOverrides(manifest.Obj, gatekeeper.Spec); err != nil {
			return err
		}
		if err := webhookOverrides(manifest.Obj, gatekeeper.Spec.Webhook); err != nil {
			return err
		}
		if isOpenshift {
			if err := removeAnnotations(manifest.Obj); err != nil {
				return err
			}
		}
	}
	// ValidatingWebhookConfiguration overrides
	if asset == ValidatingWebhookConfiguration {
		if err := validatingWebhookConfigurationOverrides(manifest.Obj, gatekeeper.Spec.Webhook); err != nil {
			return err
		}
	}
	return nil
}

func commonOverrides(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	for _, f := range commonSpecOverridesFn {
		if err := f(obj, spec); err != nil {
			return err
		}
	}
	return nil
}

func auditOverrides(obj *unstructured.Unstructured, audit *operatorv1alpha1.AuditConfig) error {
	if audit != nil {
		if err := setReplicas(obj, audit.Replicas); err != nil {
			return err
		}
		if err := setLogLevel(obj, audit.LogLevel); err != nil {
			return err
		}
		if err := setAuditInterval(obj, audit.AuditInterval); err != nil {
			return err
		}
		if err := setConstraintViolationLimit(obj, audit.ConstraintViolationLimit); err != nil {
			return err
		}
		if err := setAuditFromCache(obj, audit.AuditFromCache); err != nil {
			return err
		}
		if err := setAuditChunkSize(obj, audit.AuditChunkSize); err != nil {
			return err
		}
		if err := setEmitEvents(obj, EmitAuditEventsArg, audit.EmitAuditEvents); err != nil {
			return err
		}
		if err := setResources(obj, audit.Resources); err != nil {
			return err
		}
	}
	return nil
}

func webhookOverrides(obj *unstructured.Unstructured, webhook *operatorv1alpha1.WebhookConfig) error {
	if webhook != nil {
		if err := setReplicas(obj, webhook.Replicas); err != nil {
			return err
		}
		if err := setLogLevel(obj, webhook.LogLevel); err != nil {
			return err
		}
		if err := setEmitEvents(obj, EmitAdmissionEventsArg, webhook.EmitAdmissionEvents); err != nil {
			return err
		}
		if err := setResources(obj, webhook.Resources); err != nil {
			return err
		}
	}
	return nil
}

func validatingWebhookConfigurationOverrides(obj *unstructured.Unstructured, webhook *operatorv1alpha1.WebhookConfig) error {
	if webhook != nil {
		if err := setFailurePolicy(obj, webhook.FailurePolicy); err != nil {
			return err
		}
	}
	return nil
}

func containerOverrides(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	for _, f := range commonContainerOverridesFn {
		err := setContainerAttrWithFn(obj, managerContainer, func(container map[string]interface{}) error {
			return f(container, spec)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// setReplicas
func setReplicas(obj *unstructured.Unstructured, replicas *int32) error {
	if replicas != nil {
		if err := unstructured.SetNestedField(obj.Object, int64(*replicas), "spec", "replicas"); err != nil {
			return errors.Wrapf(err, "Failed to set replica value")
		}
	}
	return nil
}

func removeAnnotations(obj *unstructured.Unstructured) error {
	if err := unstructured.SetNestedField(obj.Object, map[string]interface{}{}, "spec", "template", "metadata", "annotations"); err != nil {
		return errors.Wrapf(err, "Failed to remove annotations")
	}
	return nil
}

func setLogLevel(obj *unstructured.Unstructured, logLevel *operatorv1alpha1.LogLevelMode) error {
	if logLevel != nil {
		return setContainerArg(obj, managerContainer, LogLevelArg, string(*logLevel))
	}
	return nil
}

func setAuditInterval(obj *unstructured.Unstructured, auditInterval *metav1.Duration) error {
	if auditInterval != nil {
		return setContainerArg(obj, managerContainer, AuditIntervalArg, fmt.Sprint(auditInterval.Round(time.Second).Seconds()))
	}
	return nil
}

func setConstraintViolationLimit(obj *unstructured.Unstructured, constraintViolationLimit *uint64) error {
	if constraintViolationLimit != nil {
		return setContainerArg(obj, managerContainer, ConstraintViolationLimitArg, strconv.FormatUint(*constraintViolationLimit, 10))
	}
	return nil
}

func setAuditFromCache(obj *unstructured.Unstructured, auditFromCache *operatorv1alpha1.AuditFromCacheMode) error {
	if auditFromCache != nil {
		auditFromCacheValue := "false"
		if *auditFromCache == operatorv1alpha1.AuditFromCacheEnabled {
			auditFromCacheValue = "true"
		}
		return setContainerArg(obj, managerContainer, AuditFromCacheArg, auditFromCacheValue)
	}
	return nil
}

func setAuditChunkSize(obj *unstructured.Unstructured, auditChunkSize *uint64) error {
	if auditChunkSize != nil {
		return setContainerArg(obj, managerContainer, AuditChunkSizeArg, strconv.FormatUint(*auditChunkSize, 10))
	}
	return nil
}

func setEmitEvents(obj *unstructured.Unstructured, argName string, emitEvents *operatorv1alpha1.EmitEventsMode) error {
	if emitEvents != nil {
		emitArgValue := "false"
		if *emitEvents == operatorv1alpha1.EmitEventsEnabled {
			emitArgValue = "true"
		}
		return setContainerArg(obj, managerContainer, argName, emitArgValue)
	}
	return nil
}

func setFailurePolicy(obj *unstructured.Unstructured, failurePolicy *admregv1.FailurePolicyType) error {
	if failurePolicy != nil {
		webhooks, found, err := unstructured.NestedSlice(obj.Object, "webhooks")
		if err != nil || !found {
			return errors.Wrapf(err, "Failed to retrieve webhooks definition")
		}
		for _, w := range webhooks {
			webhook := w.(map[string]interface{})
			if webhook["name"] == ValidationGatekeeperWebhook {
				if err := unstructured.SetNestedField(webhook, string(*failurePolicy), "failurePolicy"); err != nil {
					return errors.Wrapf(err, "Failed to set webhook failure policy")
				}
			}
		}
		if err := unstructured.SetNestedSlice(obj.Object, webhooks, "webhooks"); err != nil {
			return errors.Wrapf(err, "Failed to set webhooks")
		}
	}
	return nil
}

// Generic setters

func setAffinity(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	if spec.Affinity != nil {
		if err := unstructured.SetNestedField(obj.Object, util.ToMap(spec.Affinity), "spec", "template", "spec", "affinity"); err != nil {
			return errors.Wrapf(err, "Failed to set affinity value")
		}
	}
	return nil
}

func setNodeSelector(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	if spec.NodeSelector != nil {
		if err := unstructured.SetNestedStringMap(obj.Object, spec.NodeSelector, "spec", "template", "spec", "nodeSelector"); err != nil {
			return errors.Wrapf(err, "Failed to set nodeSelector value")
		}
	}
	return nil
}

func setPodAnnotations(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	if spec.PodAnnotations != nil {
		if err := unstructured.SetNestedStringMap(obj.Object, spec.PodAnnotations, "spec", "template", "metadata", "annotations"); err != nil {
			return errors.Wrapf(err, "Failed to set podAnnotations")
		}
	}
	return nil
}

func setTolerations(obj *unstructured.Unstructured, spec operatorv1alpha1.GatekeeperSpec) error {
	if spec.Tolerations != nil {
		tolerations := make([]interface{}, len(spec.Tolerations))
		for i, t := range spec.Tolerations {
			tolerations[i] = util.ToMap(t)
		}
		if err := unstructured.SetNestedSlice(obj.Object, tolerations, "spec", "template", "spec", "tolerations"); err != nil {
			return errors.Wrapf(err, "Failed to set container tolerations")
		}
	}
	return nil
}

// Container specific setters

func setResources(obj *unstructured.Unstructured, resources *corev1.ResourceRequirements) error {
	if resources != nil {
		return setContainerAttrWithFn(obj, managerContainer, func(container map[string]interface{}) error {
			if err := unstructured.SetNestedField(container, util.ToMap(resources), "resources"); err != nil {
				return errors.Wrapf(err, "Failed to set container resources")
			}
			return nil
		})
	}
	return nil
}

func setImage(container map[string]interface{}, spec operatorv1alpha1.GatekeeperSpec) error {
	if spec.Image == nil {
		return nil
	}
	if spec.Image.Image != nil {
		if err := unstructured.SetNestedField(container, *spec.Image.Image, "image"); err != nil {
			return errors.Wrapf(err, "Failed to set container image")
		}
	}
	if spec.Image.ImagePullPolicy != nil {
		if err := unstructured.SetNestedField(container, string(*spec.Image.ImagePullPolicy), "imagePullPolicy"); err != nil {
			return errors.Wrapf(err, "Failed to set container image pull policy")
		}
	}
	return nil
}

func setContainerAttrWithFn(obj *unstructured.Unstructured, containerName string, containerFn func(map[string]interface{}) error) error {
	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	if err != nil || !found {
		return errors.Wrapf(err, "Failed to retrieve containers")
	}
	for _, c := range containers {
		container := c.(map[string]interface{})
		if name, found, err := unstructured.NestedString(container, "name"); err != nil || !found {
			return errors.Wrapf(err, "Unable to retrieve container: %s", name)
		} else if name == containerName {
			if err := containerFn(container); err != nil {
				return err
			}
		}
	}
	if err := unstructured.SetNestedSlice(obj.Object, containers, "spec", "template", "spec", "containers"); err != nil {
		return errors.Wrapf(err, "Failed to set containers")
	}
	return nil
}

func setContainerArg(obj *unstructured.Unstructured, containerName, argName string, argValue string) error {
	return setContainerAttrWithFn(obj, containerName, func(container map[string]interface{}) error {
		args, found, err := unstructured.NestedStringSlice(container, "args")
		if !found || err != nil {
			return errors.Wrapf(err, "Unable to retrieve container arguments for: %s", containerName)
		}
		exists := false
		for i, arg := range args {
			n, _ := util.FromArg(arg)
			if n == argName {
				args[i] = util.ToArg(argName, argValue)
				exists = true
			}
		}
		if !exists {
			args = append(args, util.ToArg(argName, argValue))
		}
		return unstructured.SetNestedStringSlice(container, args, "args")
	})
}

func setCurrentNamespace(obj *unstructured.Unstructured, asset, namespace string) error {
	if obj.GetNamespace() != "" {
		obj.SetNamespace(namespace)
	}
	if err := setClientConfigNamespace(obj, asset, namespace); err != nil {
		return err
	}
	if err := setControllerManagerExceptNamespace(obj, asset, namespace); err != nil {
		return err
	}
	if err := setRoleBindingSubjectNamespace(obj, asset, namespace); err != nil {

	}
	return nil
}

func setClientConfigNamespace(obj *unstructured.Unstructured, asset, namespace string) error {
	if asset != ValidatingWebhookConfiguration {
		return nil
	}
	webhooks, found, err := unstructured.NestedSlice(obj.Object, "webhooks")
	if err != nil || !found {
		return errors.Wrapf(err, "Failed to retrieve webhooks definition")
	}
	for _, w := range webhooks {
		webhook := w.(map[string]interface{})
		if err := unstructured.SetNestedField(webhook, namespace, "clientConfig", "service", "namespace"); err != nil {
			return errors.Wrapf(err, "Failed to set webhook clientConfig.service.namespace")
		}
	}
	if err := unstructured.SetNestedSlice(obj.Object, webhooks, "webhooks"); err != nil {
		return errors.Wrapf(err, "Failed to set webhooks")
	}
	return nil
}

func setControllerManagerExceptNamespace(obj *unstructured.Unstructured, asset, namespace string) error {
	if asset != WebhookFile {
		return nil
	}
	return setContainerArg(obj, managerContainer, ExemptNamespaceArg, namespace)
}

func setRoleBindingSubjectNamespace(obj *unstructured.Unstructured, asset, namespace string) error {
	if asset != ClusterRoleBindingFile && asset != RoleBindingFile {
		return nil
	}
	subjects, found, err := unstructured.NestedSlice(obj.Object, "subjects")
	if !found || err != nil {
		return errors.Wrapf(err, "Failed to retrieve subjects from roleBinding")
	}
	for _, s := range subjects {
		subject := s.(map[string]interface{})
		if err := unstructured.SetNestedField(subject, namespace, "namespace"); err != nil {
			return errors.Wrapf(err, "Failed to set namespace for rolebinding subject")
		}
	}
	if err := unstructured.SetNestedSlice(obj.Object, subjects, "subjects"); err != nil {
		return errors.Wrapf(err, "Failed to set updated subjects in rolebinding")
	}
	return nil
}
