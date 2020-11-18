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
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/library-go/pkg/manifest"
	"github.com/pkg/errors"
	admregv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/gatekeeper/gatekeeper-operator/api/v1alpha1"
	"github.com/gatekeeper/gatekeeper-operator/pkg/bindata"
)

const (
	// The length of time between polls.
	pollInterval = 50 * time.Millisecond
	// How long to try before giving up.
	waitTimeout = 30 * time.Second
	// Gatekeeper name and namespace
	gkName      = "gatekeeper"
	gkNamespace = "gatekeeper-system"
)

var (
	auditReplicas   = int32(1)
	webhookReplicas = int32(3)
	ctx             = context.Background()
	auditName       = types.NamespacedName{
		Namespace: gkNamespace,
		Name:      "gatekeeper-audit",
	}
	controllerManagerName = types.NamespacedName{
		Namespace: gkNamespace,
		Name:      "gatekeeper-controller-manager",
	}
	gatekeeperName = types.NamespacedName{
		Namespace: gkNamespace,
		Name:      gkName,
	}
	validatingWebhookName = types.NamespacedName{
		Namespace: gkNamespace,
		Name:      "gatekeeper-validating-webhook-configuration",
	}
	// duplicated variables from the controllers package
	auditFile   = "apps_v1_deployment_gatekeeper-audit.yaml"
	webhookFile = "apps_v1_deployment_gatekeeper-controller-manager.yaml"
)

var _ = Describe("Gatekeeper", func() {

	BeforeEach(func() {
		if !useExistingCluster() {
			Skip("Test requires existing cluster. Set environment variable USE_EXISTING_CLUSTER=true and try again.")
		}
	})

	AfterEach(func() {
		Expect(K8sClient.Delete(ctx, emptyGatekeeper())).Should(Succeed())
		Eventually(func() bool {
			err := K8sClient.Get(ctx, gatekeeperName, &v1alpha1.Gatekeeper{})
			if err == nil {
				return false
			}
			return apierrors.IsNotFound(err)
		}, waitTimeout, pollInterval).Should(BeTrue())

		//wait for reconciliation clean up
		Eventually(func() bool {
			err := K8sClient.Get(ctx, auditName, &appsv1.Deployment{})
			if err == nil {
				return false
			}
			return apierrors.IsNotFound(err)
		}, waitTimeout, pollInterval).Should(BeTrue())
		Eventually(func() bool {
			err := K8sClient.Get(ctx, controllerManagerName, &appsv1.Deployment{})
			if err == nil {
				return false
			}
			return apierrors.IsNotFound(err)
		}, waitTimeout, pollInterval).Should(BeTrue())
	})

	Describe("Install", func() {
		Context("Creating Gatekeeper custom resource", func() {
			It("Should install Gatekeeper", func() {
				gatekeeper := &v1alpha1.Gatekeeper{}
				gatekeeper.Namespace = "gatekeeper-system"
				err := loadGatekeeperFromFile(gatekeeper, "operator_v1alpha1_gatekeeper.yaml")
				Expect(err).ToNot(HaveOccurred())
				gkDeployment := &appsv1.Deployment{}

				By("Creating Gatekeeper resource", func() {
					Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
				})

				By("Checking gatekeeper-controller-manager readiness", func() {
					Eventually(func() (int32, error) {
						return getDeploymentReadyReplicas(ctx, controllerManagerName, gkDeployment)
					}, waitTimeout, pollInterval).Should(Equal(*gatekeeper.Spec.Webhook.Replicas))
				})

				By("Checking gatekeeper-audit readiness", func() {
					Eventually(func() (int32, error) {
						return getDeploymentReadyReplicas(ctx, auditName, gkDeployment)
					}, waitTimeout, pollInterval).Should(Equal(*gatekeeper.Spec.Audit.Replicas))
				})
			})
		})
	})

	Describe("Overriding CR", func() {
		It("Contains default values", func() {
			gatekeeper := emptyGatekeeper()
			Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())

			auditDeployment := &appsv1.Deployment{}
			Eventually(func() error {
				return K8sClient.Get(ctx, auditName, auditDeployment)
			}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
			Expect(auditDeployment).NotTo(BeNil())

			webhookDeployment := &appsv1.Deployment{}
			Eventually(func() error {
				return K8sClient.Get(ctx, controllerManagerName, webhookDeployment)
			}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
			Expect(webhookDeployment).NotTo(BeNil())

			By("Checking default replicas", func() {
				Expect(auditDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(*auditDeployment.Spec.Replicas).To(Equal(auditReplicas))
				Expect(webhookDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(*webhookDeployment.Spec.Replicas).To(Equal(webhookReplicas))
			})

			By("Checking default pod affinity", func() {
				Expect(auditDeployment.Spec.Template.Spec.Affinity).To(BeNil())
				defaultWebhookAffinity := &corev1.Affinity{
					PodAntiAffinity: &corev1.PodAntiAffinity{
						PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
							{
								Weight: 100,
								PodAffinityTerm: corev1.PodAffinityTerm{
									LabelSelector: &metav1.LabelSelector{
										MatchExpressions: []metav1.LabelSelectorRequirement{
											{
												Key:      "gatekeeper.sh/operation",
												Operator: metav1.LabelSelectorOpIn,
												Values: []string{
													"webhook",
												},
											},
										},
									},
									TopologyKey: "kubernetes.io/hostname",
								},
							},
						},
					},
				}
				Expect(webhookDeployment.Spec.Template.Spec.Affinity).To(BeEquivalentTo(defaultWebhookAffinity))
			})

			By("Checking default node selector", func() {
				defaultConfig := map[string]string{
					"kubernetes.io/os": "linux",
				}
				Expect(auditDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(defaultConfig))
				Expect(webhookDeployment.Spec.Template.Spec.NodeSelector).To(BeEquivalentTo(defaultConfig))
			})

			By("Checking default pod annotations", func() {
				defaultConfig := map[string]string{
					"container.seccomp.security.alpha.kubernetes.io/manager": "runtime/default",
				}
				Expect(auditDeployment.Spec.Template.Annotations).To(BeEquivalentTo(defaultConfig))
				Expect(webhookDeployment.Spec.Template.Annotations).To(BeEquivalentTo(defaultConfig))
			})

			By("Checking default tolerations", func() {
				Expect(auditDeployment.Spec.Template.Spec.Tolerations).To(BeNil())
				Expect(webhookDeployment.Spec.Template.Spec.Tolerations).To(BeNil())
			})

			By("Checking default resource limits and requests", func() {
				defaultConfig := corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("1000m"),
						corev1.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("100m"),
						corev1.ResourceMemory: resource.MustParse("256Mi"),
					},
				}
				assertResources(defaultConfig, auditDeployment.Spec.Template.Spec.Containers[0].Resources)
				assertResources(defaultConfig, webhookDeployment.Spec.Template.Spec.Containers[0].Resources)
			})

			By("Checking default image", func() {
				auditImage, auditImagePullPolicy, err := getDefaultImage(auditFile)
				Expect(err).NotTo(HaveOccurred())
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(auditImage))
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(auditImagePullPolicy))
				webhookImage, webhookImagePullPolicy, err := getDefaultImage(webhookFile)
				Expect(err).NotTo(HaveOccurred())
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(webhookImage))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(webhookImagePullPolicy))
			})

			By("Checking default failure policy", func() {
				defaultPolicy := admregv1.Ignore
				validatingWebhookConfiguration := &admregv1.ValidatingWebhookConfiguration{}
				Eventually(func() error {
					return K8sClient.Get(ctx, validatingWebhookName, validatingWebhookConfiguration)
				}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
				Expect(auditDeployment).NotTo(BeNil())
				for _, wh := range validatingWebhookConfiguration.Webhooks {
					if wh.Name == "validation.gatekeeper.sh" {
						Expect(wh.FailurePolicy).To(Equal(&defaultPolicy))
					}
				}
			})

			By("Checking default audit interval", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-interval")
				Expect(found).To(BeFalse())
			})

			By("Checking default audit log level", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--log-level")
				Expect(found).To(BeFalse())
			})

			By("Checking default audit constraint violation limit", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--constraint-violations-limit")
				Expect(found).To(BeFalse())
			})

			By("Checking default audit chunk size", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-chunk-size")
				Expect(found).To(BeFalse())
			})

			By("Checking default audit from cache", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-from-cache")
				Expect(found).To(BeFalse())
			})

			By("Checking default emit audit events", func() {
				_, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--emit-audit-events")
				Expect(found).To(BeFalse())
			})

			By("Checking default emit admission events", func() {
				_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, "--emit-admission-events")
				Expect(found).To(BeFalse())
			})

			By("Checking default webhook log level", func() {
				_, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, "--log-level")
				Expect(found).To(BeFalse())
			})
		})

		It("Contains the configured values", func() {
			gatekeeper := &v1alpha1.Gatekeeper{}
			gatekeeper.Namespace = "gatekeeper-system"
			err := loadGatekeeperFromFile(gatekeeper, "gatekeeper_with_all_values.yaml")
			Expect(err).ToNot(HaveOccurred())
			Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())

			auditDeployment := &appsv1.Deployment{}
			Eventually(func() error {
				return K8sClient.Get(ctx, auditName, auditDeployment)
			}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
			Expect(auditDeployment).NotTo(BeNil())

			webhookDeployment := &appsv1.Deployment{}
			Eventually(func() error {
				return K8sClient.Get(ctx, controllerManagerName, webhookDeployment)
			}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
			Expect(webhookDeployment).NotTo(BeNil())

			By("Checking expected replicas", func() {
				Expect(auditDeployment.Spec.Replicas).NotTo(BeNil())
				Expect(auditDeployment.Spec.Replicas).To(Equal(gatekeeper.Spec.Audit.Replicas))
				Expect(webhookDeployment.Spec.Replicas).NotTo(BeNil())
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
				assertResources(*gatekeeper.Spec.Resources, auditDeployment.Spec.Template.Spec.Containers[0].Resources)
				assertResources(*gatekeeper.Spec.Resources, webhookDeployment.Spec.Template.Spec.Containers[0].Resources)
			})

			By("Checking expected image", func() {
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(*gatekeeper.Spec.Image.Image))
				Expect(auditDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(*gatekeeper.Spec.Image.ImagePullPolicy))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].Image).To(Equal(*gatekeeper.Spec.Image.Image))
				Expect(webhookDeployment.Spec.Template.Spec.Containers[0].ImagePullPolicy).To(Equal(*gatekeeper.Spec.Image.ImagePullPolicy))
			})

			By("Checking expected failure policy", func() {
				validatingWebhookConfiguration := &admregv1.ValidatingWebhookConfiguration{}
				Eventually(func() error {
					return K8sClient.Get(ctx, validatingWebhookName, validatingWebhookConfiguration)
				}, waitTimeout, pollInterval).ShouldNot(HaveOccurred())
				Expect(validatingWebhookConfiguration).NotTo(BeNil())
				for _, wh := range validatingWebhookConfiguration.Webhooks {
					if wh.Name == "validation.gatekeeper.sh" {
						Expect(wh.FailurePolicy).To(Equal(gatekeeper.Spec.Webhook.FailurePolicy))
					}
				}
			})

			By("Checking expected audit interval", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-interval")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--audit-interval=10"))
			})

			By("Checking expected audit log level", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--log-level")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--log-level=DEBUG"))
			})

			By("Checking expected audit constraint violation limit", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--constraint-violations-limit")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--constraint-violations-limit=55"))
			})

			By("Checking expected audit chunk size", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-chunk-size")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--audit-chunk-size=66"))
			})

			By("Checking expected audit from cache", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--audit-from-cache")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--audit-from-cache=true"))
			})

			By("Checking expected emit audit events", func() {
				value, found := getContainerArg(auditDeployment.Spec.Template.Spec.Containers[0].Args, "--emit-audit-events")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--emit-audit-events=true"))
			})

			By("Checking expected emit admission events", func() {
				value, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, "--emit-admission-events")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--emit-admission-events=true"))
			})

			By("Checking expected webhook log level", func() {
				value, found := getContainerArg(webhookDeployment.Spec.Template.Spec.Containers[0].Args, "--log-level")
				Expect(found).To(BeTrue())
				Expect(value).To(Equal("--log-level=ERROR"))
			})
		})
	})
})

func assertResources(expected, current corev1.ResourceRequirements) {
	Expect(expected.Limits.Cpu().Cmp(*current.Limits.Cpu())).To(BeZero())
	Expect(expected.Limits.Memory().Cmp(*current.Limits.Memory())).To(BeZero())
	Expect(expected.Requests.Cpu().Cmp(*current.Requests.Cpu())).To(BeZero())
	Expect(expected.Requests.Memory().Cmp(*current.Requests.Memory())).To(BeZero())
}

func getContainerArg(args []string, argPrefix string) (arg string, found bool) {
	for _, arg := range args {
		if strings.HasPrefix(arg, argPrefix) {
			return arg, true
		}
	}
	return "", false
}

func loadGatekeeperFromFile(gatekeeper *v1alpha1.Gatekeeper, fileName string) error {
	f, err := os.Open(fmt.Sprintf("../config/samples/%s", fileName))
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
	deploy *appsv1.Deployment) (int32, error) {
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
			Namespace: gkNamespace,
		},
	}
}

func ptrToInt(num int32) *int32 {
	return &num
}

func getManifest(asset string) (*manifest.Manifest, error) {
	manifest := &manifest.Manifest{}
	assetName := "config/gatekeeper/" + asset
	bytes, err := bindata.Asset(assetName)
	if err != nil {
		return manifest, errors.Wrapf(err, "Unable to retrieve bindata asset %s", assetName)
	}

	err = manifest.UnmarshalJSON(bytes)
	if err != nil {
		return manifest, errors.Wrapf(err, "Unable to unmarshal YAML bytes for asset name %s", assetName)
	}
	return manifest, nil
}

func getDefaultImage(file string) (image string, imagePullPolicy corev1.PullPolicy, err error) {
	manifest, err := getManifest(file)
	if err != nil {
		return "", "", err
	}
	containers, found, err := unstructured.NestedSlice(manifest.Obj.Object, "spec", "template", "spec", "containers")
	if !found {
		return "", "", fmt.Errorf("Containers not found")
	}
	if err != nil {
		return "", "", err
	}
	image, found, err = unstructured.NestedString(containers[0].(map[string]interface{}), "image")
	if !found {
		return "", "", fmt.Errorf("Image not found")
	}
	if err != nil {
		return "", "", err
	}
	policy, found, err := unstructured.NestedString(containers[0].(map[string]interface{}), "imagePullPolicy")
	if !found {
		return "", "", fmt.Errorf("ImagePullPolicy not found")
	}
	return image, corev1.PullPolicy(policy), err
}
