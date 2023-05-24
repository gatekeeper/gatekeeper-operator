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

package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	admregv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gatekeeper/gatekeeper-operator/api/v1alpha1"
	"github.com/gatekeeper/gatekeeper-operator/controllers"
	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
	test "github.com/gatekeeper/gatekeeper-operator/test/e2e/util"
)

const (
	// Gatekeeper name and namespace
	gkName                      = "gatekeeper"
	gatekeeperWithAllValuesFile = "gatekeeper_with_all_values.yaml"
)

var (
	ctx                   = context.Background()
	globalsInitialized    = false
	auditName             = types.NamespacedName{}
	controllerManagerName = types.NamespacedName{}
	gatekeeperName        = types.NamespacedName{
		Name: gkName,
	}
	validatingWebhookName = types.NamespacedName{}
	mutatingWebhookName   = types.NamespacedName{}
)

func initializeGlobals() {
	auditName = types.NamespacedName{
		Namespace: gatekeeperNamespace,
		Name:      "gatekeeper-audit",
	}
	controllerManagerName = types.NamespacedName{
		Namespace: gatekeeperNamespace,
		Name:      "gatekeeper-controller-manager",
	}
	validatingWebhookName = types.NamespacedName{
		Namespace: gatekeeperNamespace,
		Name:      "gatekeeper-validating-webhook-configuration",
	}
	mutatingWebhookName = types.NamespacedName{
		Namespace: gatekeeperNamespace,
		Name:      "gatekeeper-mutating-webhook-configuration",
	}
}

var _ = Describe("Gatekeeper", func() {
	BeforeEach(func() {
		if !useExistingCluster() {
			Skip("Test requires existing cluster. Set environment variable USE_EXISTING_CLUSTER=true and try again.")
		}

		if !globalsInitialized {
			initializeGlobals()
			globalsInitialized = true
		}
	})

	AfterEach(func() {
		Expect(K8sClient.Delete(ctx, emptyGatekeeper(), client.PropagationPolicy(v1.DeletePropagationForeground))).Should(Succeed())

		// Once this succeeds, clean up has happened for all owned resources.
		Eventually(func() bool {
			err := K8sClient.Get(ctx, gatekeeperName, &v1alpha1.Gatekeeper{})
			if err == nil {
				return false
			}
			return apierrors.IsNotFound(err)
		}, deleteTimeout, pollInterval).Should(BeTrue())
	})

	Describe("Overriding CR", Ordered, func() {
		It("Creating an empty gatekeeper contains default values", func() {
			gatekeeper := emptyGatekeeper()
			err := loadGatekeeperFromFile(gatekeeper, "gatekeeper_empty.yaml")
			Expect(err).ToNot(HaveOccurred())

			By("Creating Gatekeeper resource", func() {
				Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			})

			auditDeployment, webhookDeployment := gatekeeperDeployments()

			By("Checking default replicas", func() {
				Expect(auditDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(*auditDeployment.Spec.Replicas).To(Equal(test.DefaultDeployment.AuditReplicas))
				Expect(webhookDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(*webhookDeployment.Spec.Replicas).To(Equal(test.DefaultDeployment.WebhookReplicas))
			})

			By("Checking gatekeeper-controller-manager readiness", func() {
				gkDeployment := &appsv1.Deployment{}
				Eventually(func() (int32, error) {
					return getDeploymentReadyReplicas(ctx, controllerManagerName, gkDeployment)
				}, timeout, pollInterval).Should(Equal(test.DefaultDeployment.WebhookReplicas))
			})

			By("Checking gatekeeper-audit readiness", func() {
				gkDeployment := &appsv1.Deployment{}
				Eventually(func() (int32, error) {
					return getDeploymentReadyReplicas(ctx, auditName, gkDeployment)
				}, timeout, pollInterval).Should(Equal(test.DefaultDeployment.AuditReplicas))
			})

			By("Checking validatingWebhookConfiguration is deployed", func() {
				validatingWebhookConfiguration := &admregv1.ValidatingWebhookConfiguration{}
				Eventually(func() error {
					return K8sClient.Get(ctx, validatingWebhookName, validatingWebhookConfiguration)
				}, timeout, pollInterval).ShouldNot(HaveOccurred())
				Expect(validatingWebhookConfiguration.OwnerReferences).To(HaveLen(1))
				Expect(validatingWebhookConfiguration.OwnerReferences[0].Kind).To(Equal("Gatekeeper"))
				Expect(validatingWebhookConfiguration.OwnerReferences[0].Name).To(Equal(gkName))
			})

			By("Checking mutatingWebhookConfiguration is deployed", func() {
				mutatingWebhookConfiguration := &admregv1.MutatingWebhookConfiguration{}
				Eventually(func() error {
					return K8sClient.Get(ctx, mutatingWebhookName, mutatingWebhookConfiguration)
				}, timeout, pollInterval).ShouldNot(HaveOccurred())
				Expect(mutatingWebhookConfiguration.OwnerReferences).To(HaveLen(1))
				Expect(mutatingWebhookConfiguration.OwnerReferences[0].Kind).To(Equal("Gatekeeper"))
				Expect(mutatingWebhookConfiguration.OwnerReferences[0].Name).To(Equal(gkName))
			})

			By("Checking default pod affinity", func() {
				Expect(auditDeployment.Spec.Template.Spec.Affinity).To(BeNil())
				Expect(webhookDeployment.Spec.Template.Spec.Affinity).To(BeEquivalentTo(test.DefaultDeployment.Affinity))
			})

			By("Checking default node selector", func() {
				Expect(auditDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(test.DefaultDeployment.NodeSelector))
				Expect(webhookDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(test.DefaultDeployment.NodeSelector))
			})

			By("Checking default pod annotations", func() {
				Expect(auditDeployment.Spec.Template.Annotations).To(BeEquivalentTo(test.DefaultDeployment.PodAnnotations))
				Expect(webhookDeployment.Spec.Template.Annotations).To(BeEquivalentTo(test.DefaultDeployment.PodAnnotations))
			})

			By("Checking default tolerations", func() {
				Expect(auditDeployment.Spec.Template.Spec.Tolerations).To(BeNil())
				Expect(webhookDeployment.Spec.Template.Spec.Tolerations).To(BeNil())
			})

			By("Checking default resource limits and requests", func() {
				assertResources(*test.DefaultDeployment.Resources, auditDeployment.Spec.Template.Spec.Containers[0].Resources)
				assertResources(*test.DefaultDeployment.Resources, webhookDeployment.Spec.Template.Spec.Containers[0].Resources)
			})

			By("Checking default image", func() {
				auditImage, auditImagePullPolicy, err := getDefaultImage(controllers.AuditFile)
				Expect(err).NotTo(HaveOccurred())
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(auditImage))
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(auditImagePullPolicy))
				webhookImage, webhookImagePullPolicy, err := getDefaultImage(controllers.WebhookFile)
				Expect(err).NotTo(HaveOccurred())
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(webhookImage))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(webhookImagePullPolicy))
			})

			byCheckingFailurePolicy(&validatingWebhookName, "default",
				util.ValidatingWebhookConfigurationKind,
				controllers.ValidationGatekeeperWebhook,
				&test.DefaultDeployment.FailurePolicy)

			byCheckingNamespaceSelector(&validatingWebhookName, "default",
				util.ValidatingWebhookConfigurationKind,
				controllers.ValidationGatekeeperWebhook,
				test.DefaultDeployment.NamespaceSelector)

			byCheckingFailurePolicy(&mutatingWebhookName, "default",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				&test.DefaultDeployment.FailurePolicy)

			byCheckingNamespaceSelector(&mutatingWebhookName, "default",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				test.DefaultDeployment.NamespaceSelector)

			By("Checking default audit interval", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditIntervalArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default audit log level", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.LogLevelArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default audit constraint violation limit", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.ConstraintViolationLimitArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default audit chunk size", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditChunkSizeArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default audit from cache", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditFromCacheArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default emit audit events", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.EmitAuditEventsArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default emit admission events", func() {
				_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.EmitAdmissionEventsArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default webhook log level", func() {
				_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.LogLevelArg)
				Expect(found).To(BeFalse())
			})

			By("Checking default disabled builtins", func() {
				_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.DisabledBuiltinArg)
				Expect(found).To(BeTrue())
			})
		})

		It("Contains the configured values", func() {
			gatekeeper := &v1alpha1.Gatekeeper{}
			gatekeeper.Namespace = gatekeeperNamespace
			err := loadGatekeeperFromFile(gatekeeper, gatekeeperWithAllValuesFile)
			Expect(err).ToNot(HaveOccurred())
			Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			auditDeployment, webhookDeployment := gatekeeperDeployments()

			By("Checking expected replicas", func() {
				Expect(auditDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(auditDeployment.Spec.Replicas).To(Equal(gatekeeper.Spec.Audit.Replicas))
				Expect(webhookDeployment.Spec.Replicas).NotTo(BeNil())
				// TODO: Remove once flake has been fixed. See
				// https://github.com/gatekeeper/gatekeeper-operator/pull/168/checks?check_run_id=2918723659
				// for example failure.
				fmt.Fprint(GinkgoWriter, "webhookDeployment", webhookDeployment)
				Expect(webhookDeployment.Spec.Replicas).To(Equal(gatekeeper.Spec.Webhook.Replicas))
			})

			By("Checking expected pod affinity", func() {
				Expect(auditDeployment.Spec.Template.Spec.Affinity).To(BeEquivalentTo(gatekeeper.Spec.Affinity))
				Expect(webhookDeployment.Spec.Template.Spec.Affinity).To(BeEquivalentTo(gatekeeper.Spec.Affinity))
			})

			By("Checking expected node selector", func() {
				Expect(auditDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(gatekeeper.Spec.NodeSelector))
				Expect(webhookDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(gatekeeper.Spec.NodeSelector))
			})

			By("Checking expected pod annotations", func() {
				Expect(auditDeployment.Spec.Template.Annotations).To(BeEquivalentTo(gatekeeper.Spec.PodAnnotations))
				Expect(webhookDeployment.Spec.Template.Annotations).To(BeEquivalentTo(gatekeeper.Spec.PodAnnotations))
			})

			By("Checking expected tolerations", func() {
				Expect(auditDeployment.Spec.Template.Spec.Tolerations).To(BeEquivalentTo(gatekeeper.Spec.Tolerations))
				Expect(webhookDeployment.Spec.Template.Spec.Tolerations).To(BeEquivalentTo(gatekeeper.Spec.Tolerations))
			})

			By("Checking expected resource limits and requests", func() {
				assertResources(*gatekeeper.Spec.Audit.Resources, auditDeployment.Spec.Template.Spec.Containers[0].Resources)
				assertResources(*gatekeeper.Spec.Webhook.Resources, webhookDeployment.Spec.Template.Spec.Containers[0].Resources)
			})

			By("Checking expected image", func() {
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].Image).ToNot(Equal(*gatekeeper.Spec.Image.Image))
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(*gatekeeper.Spec.Image.ImagePullPolicy))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].Image).ToNot(Equal(*gatekeeper.Spec.Image.Image))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(*gatekeeper.Spec.Image.ImagePullPolicy))
			})

			By("Checking ready replicas", func() {
				gkDeployment := &appsv1.Deployment{}
				Eventually(func() (int32, error) {
					return getDeploymentReadyReplicas(ctx, controllerManagerName, gkDeployment)
				}, timeout, pollInterval).Should(Equal(*gatekeeper.Spec.Webhook.Replicas))
			})

			By("Checking webhook is available", func() {
				byCheckingValidationEnabled()
			})

			byCheckingFailurePolicy(&validatingWebhookName, "expected",
				util.ValidatingWebhookConfigurationKind,
				controllers.ValidationGatekeeperWebhook,
				gatekeeper.Spec.Webhook.FailurePolicy)

			byCheckingNamespaceSelector(&validatingWebhookName, "expected",
				util.ValidatingWebhookConfigurationKind,
				controllers.ValidationGatekeeperWebhook,
				gatekeeper.Spec.Webhook.NamespaceSelector)

			By("Checking expected audit interval", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditIntervalArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.AuditIntervalArg, "10")))
			})

			By("Checking expected audit log level", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.LogLevelArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.LogLevelArg, "DEBUG")))
			})

			By("Checking expected audit constraint violation limit", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.ConstraintViolationLimitArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.ConstraintViolationLimitArg, "55")))
			})

			By("Checking expected audit chunk size", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditChunkSizeArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.AuditChunkSizeArg, "66")))
			})

			By("Checking expected audit from cache", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.AuditFromCacheArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.AuditFromCacheArg, "true")))
			})

			By("Checking expected emit audit events", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, controllers.EmitAuditEventsArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.EmitAuditEventsArg, "true")))
			})

			By("Checking expected emit admission events", func() {
				value, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.EmitAdmissionEventsArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.EmitAdmissionEventsArg, "true")))
			})

			By("Checking expected webhook log level", func() {
				value, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.LogLevelArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.LogLevelArg, "ERROR")))
			})

			By("Checking expected disabled builtins", func() {
				value, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.DisabledBuiltinArg)
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(util.ToArg(controllers.DisabledBuiltinArg, "{http.send}")))
			})
		})

		It("Does not deploy the ValidatingWebhookConfiguration", func() {
			gatekeeper := emptyGatekeeper()
			By("First creating Gatekeeper CR with validation enabled", func() {
				Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			})

			gatekeeperDeployments()
			byCheckingValidationEnabled()

			By("Getting Gatekeeper CR for updating", func() {
				err := K8sClient.Get(ctx, gatekeeperName, gatekeeper)
				Expect(err).ToNot(HaveOccurred())
			})

			By("Updating Gatekeeper CR with validation disabled", func() {
				webhookMode := v1alpha1.WebhookDisabled
				gatekeeper.Spec.ValidatingWebhook = &webhookMode
				Expect(K8sClient.Update(ctx, gatekeeper)).Should(Succeed())
			})

			gatekeeperDeployments()
			byCheckingValidationDisabled()
		})

		It("Enables Gatekeeper mutation with default values", func() {
			gatekeeper := emptyGatekeeper()
			webhookMode := v1alpha1.WebhookEnabled
			gatekeeper.Spec.MutatingWebhook = &webhookMode
			Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			auditDeployment, webhookDeployment := gatekeeperDeployments()

			byCheckingMutationEnabled(auditDeployment, webhookDeployment)

			byCheckingFailurePolicy(&mutatingWebhookName, "default",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				&test.DefaultDeployment.FailurePolicy)

			byCheckingNamespaceSelector(&mutatingWebhookName, "default",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				test.DefaultDeployment.NamespaceSelector)
		})

		It("Enables Gatekeeper mutation with configured values", func() {
			gatekeeper := emptyGatekeeper()
			err := loadGatekeeperFromFile(gatekeeper, gatekeeperWithAllValuesFile)
			Expect(err).ToNot(HaveOccurred())
			webhookMode := v1alpha1.WebhookEnabled
			gatekeeper.Spec.MutatingWebhook = &webhookMode
			Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			auditDeployment, webhookDeployment := gatekeeperDeployments()

			byCheckingMutationEnabled(auditDeployment, webhookDeployment)

			byCheckingFailurePolicy(&mutatingWebhookName, "expected",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				gatekeeper.Spec.Webhook.FailurePolicy)

			byCheckingNamespaceSelector(&mutatingWebhookName, "expected",
				util.MutatingWebhookConfigurationKind,
				controllers.MutationGatekeeperWebhook,
				gatekeeper.Spec.Webhook.NamespaceSelector)
		})

		It("Enables then disables Gatekeeper mutation", func() {
			gatekeeper := emptyGatekeeper()
			By("First creating Gatekeeper CR with mutation enabled", func() {
				webhookMode := v1alpha1.WebhookEnabled
				gatekeeper.Spec.MutatingWebhook = &webhookMode
				Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
			})

			auditDeployment, webhookDeployment := gatekeeperDeployments()
			byCheckingMutationEnabled(auditDeployment, webhookDeployment)

			By("Getting Gatekeeper CR for updating", func() {
				err := K8sClient.Get(ctx, gatekeeperName, gatekeeper)
				Expect(err).ToNot(HaveOccurred())
			})

			By("Updating Gatekeeper CR with mutation disabled", func() {
				webhookMode := v1alpha1.WebhookDisabled
				gatekeeper.Spec.MutatingWebhook = &webhookMode
				Expect(K8sClient.Update(ctx, gatekeeper)).Should(Succeed())
			})

			auditDeployment, webhookDeployment = gatekeeperDeployments()
			byCheckingMutationDisabled(auditDeployment, webhookDeployment)
		})
	})
})

func gatekeeperDeployments() (auditDeployment, webhookDeployment *appsv1.Deployment) {
	return gatekeeperAuditDeployment(), gatekeeperWebhookDeployment()
}

func gatekeeperAuditDeployment() (auditDeployment *appsv1.Deployment) {
	auditDeployment = &appsv1.Deployment{}
	Eventually(func() error {
		return K8sClient.Get(ctx, auditName, auditDeployment)
	}, timeout, pollInterval).ShouldNot(HaveOccurred())
	return
}

func gatekeeperWebhookDeployment() (webhookDeployment *appsv1.Deployment) {
	webhookDeployment = &appsv1.Deployment{}
	Eventually(func() error {
		return K8sClient.Get(ctx, controllerManagerName, webhookDeployment)
	}, timeout, pollInterval).ShouldNot(HaveOccurred())
	return
}

func assertResources(expected, current corev1.ResourceRequirements) {
	Expect(expected.Limits.Cpu().Cmp(*current.Limits.Cpu())).To(BeZero())
	Expect(expected.Limits.Memory().Cmp(*current.Limits.Memory())).To(BeZero())
	Expect(expected.Requests.Cpu().Cmp(*current.Requests.Cpu())).To(BeZero())
	Expect(expected.Requests.Memory().Cmp(*current.Requests.Memory())).To(BeZero())
}

func byCheckingValidationEnabled() {
	By("Checking validation is enabled", func() {
		validatingWebhookConfiguration := &admregv1.ValidatingWebhookConfiguration{}
		Eventually(func() error {
			return K8sClient.Get(ctx, validatingWebhookName, validatingWebhookConfiguration)
		}, timeout, pollInterval).ShouldNot(HaveOccurred())
	})
}

type getCRDFunc func(types.NamespacedName, *extv1.CustomResourceDefinition)

func byCheckingMutationEnabled(auditDeployment, webhookDeployment *appsv1.Deployment) {
	By(fmt.Sprintf("Checking %s=%s argument is set", controllers.OperationArg, controllers.OperationMutationWebhook), func() {
		Eventually(func() bool {
			return findContainerArgValue(auditDeployment.Spec.Template.Spec.Containers[0].Args,
				controllers.OperationArg, controllers.OperationMutationStatus)
		}, timeout, pollInterval).Should(BeTrue())
	})

	By(fmt.Sprintf("Checking %s=%s argument is set", controllers.OperationArg, controllers.OperationMutationStatus), func() {
		Eventually(func() bool {
			return findContainerArgValue(auditDeployment.Spec.Template.Spec.Containers[0].Args,
				controllers.OperationArg, controllers.OperationMutationStatus)
		}, timeout, pollInterval).Should(BeTrue())
	})

	By("Checking MutatingWebhookConfiguration deployed", func() {
		mutatingWebhookConfiguration := &admregv1.MutatingWebhookConfiguration{}
		Eventually(func() error {
			return K8sClient.Get(ctx, mutatingWebhookName, mutatingWebhookConfiguration)
		}, timeout, pollInterval).ShouldNot(HaveOccurred())
	})

	var crdFn getCRDFunc
	crdFn = func(crdName types.NamespacedName, mutatingCRD *extv1.CustomResourceDefinition) {
		Eventually(func() error {
			return K8sClient.Get(ctx, crdName, mutatingCRD)
		}, timeout, pollInterval).ShouldNot(HaveOccurred())
	}
	byCheckingMutatingCRDs("deployed", crdFn)
}

func byCheckingValidationDisabled() {
	By("Checking validation is disabled", func() {
		validatingWebhookConfiguration := &admregv1.ValidatingWebhookConfiguration{}
		Eventually(func() bool {
			err := K8sClient.Get(ctx, validatingWebhookName, validatingWebhookConfiguration)
			return apierrors.IsNotFound(err)
		}, timeout, pollInterval).Should(BeTrue())
	})
}

func byCheckingMutationDisabled(auditDeployment, webhookDeployment *appsv1.Deployment) {
	By(fmt.Sprintf("Checking %s argument is not set", controllers.EnableMutationArg), func() {
		Eventually(func() bool {
			webhookDeployment = gatekeeperWebhookDeployment()
			_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, controllers.EnableMutationArg)
			return found
		}, timeout, pollInterval).Should(BeFalse())
	})

	By(fmt.Sprintf("Checking %s=%s argument is set", controllers.OperationArg, controllers.OperationMutationStatus), func() {
		Eventually(func() bool {
			auditDeployment = gatekeeperAuditDeployment()
			found := findContainerArgValue(auditDeployment.Spec.Template.Spec.Containers[0].Args,
				controllers.OperationArg, controllers.OperationMutationStatus)
			return found
		}, timeout, pollInterval).Should(BeTrue())
	})

	By("Checking MutatingWebhookConfiguration not deployed", func() {
		mutatingWebhookConfiguration := &admregv1.MutatingWebhookConfiguration{}
		Eventually(func() bool {
			err := K8sClient.Get(ctx, mutatingWebhookName, mutatingWebhookConfiguration)
			return apierrors.IsNotFound(err)
		}, timeout, pollInterval).Should(BeTrue())
	})

	var crdFn getCRDFunc
	crdFn = func(crdName types.NamespacedName, mutatingCRD *extv1.CustomResourceDefinition) {
		Eventually(func() bool {
			err := K8sClient.Get(ctx, crdName, mutatingCRD)
			return apierrors.IsNotFound(err)
		}, timeout, pollInterval).Should(BeTrue())
	}
	byCheckingMutatingCRDs("not deployed", crdFn)
}

func byCheckingMutatingCRDs(deployMsg string, f getCRDFunc) {
	for _, asset := range controllers.MutatingCRDs {
		obj, err := util.GetManifestObject(asset)
		Expect(err).ToNot(HaveOccurred())

		crdNamespacedName := types.NamespacedName{
			Name: obj.GetName(),
		}
		By(fmt.Sprintf("Checking %s Mutating CRD %s", obj.GetName(), deployMsg), func() {
			mutatingAssignCRD := &extv1.CustomResourceDefinition{}
			f(crdNamespacedName, mutatingAssignCRD)
		})
	}
}

func byCheckingFailurePolicy(webhookNamespacedName *types.NamespacedName,
	testName, kind, webhookName string, failurePolicy *admregv1.FailurePolicyType,
) {
	By(fmt.Sprintf("Checking %s failure policy", testName), func() {
		webhookConfiguration := &unstructured.Unstructured{}
		webhookConfiguration.SetAPIVersion(admregv1.SchemeGroupVersion.String())
		webhookConfiguration.SetKind(kind)
		Eventually(func() error {
			return K8sClient.Get(ctx, *webhookNamespacedName, webhookConfiguration)
		}, timeout, pollInterval).ShouldNot(HaveOccurred())
		assertFailurePolicy(webhookConfiguration, webhookName, failurePolicy)
	})
}

func assertFailurePolicy(obj *unstructured.Unstructured, webhookName string, expected *admregv1.FailurePolicyType) {
	assertWebhook(obj, webhookName, func(webhook map[string]interface{}) {
		Expect(webhook["failurePolicy"]).To(BeEquivalentTo(string(*expected)))
	})
}

func byCheckingNamespaceSelector(webhookNamespacedName *types.NamespacedName,
	testName, kind, webhookName string, namespaceSelector *metav1.LabelSelector,
) {
	By(fmt.Sprintf("Checking %s namespace selector", testName), func() {
		webhookConfiguration := &unstructured.Unstructured{}
		webhookConfiguration.SetAPIVersion(admregv1.SchemeGroupVersion.String())
		webhookConfiguration.SetKind(kind)
		Eventually(func() error {
			return K8sClient.Get(ctx, *webhookNamespacedName, webhookConfiguration)
		}, timeout, pollInterval).ShouldNot(HaveOccurred())
		assertNamespaceSelector(webhookConfiguration, webhookName, namespaceSelector)
	})
}

func assertNamespaceSelector(obj *unstructured.Unstructured, webhookName string, expected *metav1.LabelSelector) {
	assertWebhook(obj, webhookName, func(webhook map[string]interface{}) {
		nsSelector, found, err := unstructured.NestedFieldNoCopy(webhook, "namespaceSelector")
		Expect(err).NotTo(HaveOccurred())
		Expect(found).To(BeTrue())

		nsSelectorBytes, err := json.Marshal(nsSelector)
		Expect(err).NotTo(HaveOccurred())

		nsSelectorTyped := &metav1.LabelSelector{}
		err = json.Unmarshal(nsSelectorBytes, nsSelectorTyped)
		Expect(err).NotTo(HaveOccurred())

		Expect(nsSelectorTyped).To(BeEquivalentTo(expected))
	})
}

func assertWebhook(obj *unstructured.Unstructured, webhookName string, webhookFn func(map[string]interface{})) {
	webhooks, found, err := unstructured.NestedSlice(obj.Object, "webhooks")
	Expect(err).NotTo(HaveOccurred())
	Expect(found).To(BeTrue())
	for _, webhook := range webhooks {
		w, ok := webhook.(map[string]interface{})
		Expect(ok).To(BeTrue())
		if w["name"] == webhookName {
			webhookFn(w)
		}
	}
}

func getContainerArg(args []string, argPrefix string) (arg string, found bool) {
	for _, arg := range args {
		if strings.HasPrefix(arg, argPrefix) {
			return arg, true
		}
	}
	return "", false
}

func findContainerArgValue(args []string, argKey, argValue string) bool {
	argKeyValue := fmt.Sprintf("%s=%s", argKey, argValue)
	for _, arg := range args {
		if strings.Compare(arg, argKeyValue) == 0 {
			return true
		}
	}
	return false
}

func loadGatekeeperFromFile(gatekeeper *v1alpha1.Gatekeeper, fileName string) error {
	f, err := os.Open(fmt.Sprintf("../../config/samples/%s", fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	return decodeYAML(f, gatekeeper)
}

func decodeYAML(r io.Reader, obj interface{}) error {
	decoder := yaml.NewYAMLToJSONDecoder(r)
	return decoder.Decode(obj)
}

func useExistingCluster() bool {
	return strings.ToLower(os.Getenv("USE_EXISTING_CLUSTER")) == "true"
}

func getDeploymentReadyReplicas(ctx context.Context, name types.NamespacedName,
	deploy *appsv1.Deployment,
) (int32, error) {
	err := K8sClient.Get(ctx, name, deploy)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}

	return deploy.Status.ReadyReplicas, nil
}

func emptyGatekeeper() *v1alpha1.Gatekeeper {
	return &v1alpha1.Gatekeeper{
		ObjectMeta: v1.ObjectMeta{
			Name:      gkName,
			Namespace: gatekeeperNamespace,
		},
	}
}

func getDefaultImage(file string) (image string, imagePullPolicy corev1.PullPolicy, err error) {
	obj, err := util.GetManifestObject(file)
	if err != nil {
		return "", "", err
	}
	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	if err != nil {
		return "", "", err
	}
	if !found {
		return "", "", fmt.Errorf("Containers not found")
	}
	image, found, err = unstructured.NestedString(containers[0].(map[string]interface{}), "image")
	if err != nil {
		return "", "", err
	}
	if !found {
		return "", "", fmt.Errorf("Image not found")
	}
	policy, found, err := unstructured.NestedString(containers[0].(map[string]interface{}), "imagePullPolicy")
	if err != nil {
		return "", "", err
	}
	if !found {
		return "", "", fmt.Errorf("ImagePullPolicy not found")
	}
	return image, corev1.PullPolicy(policy), nil
}
