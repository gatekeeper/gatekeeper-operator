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
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	admregv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	operatorv1alpha1 "github.com/gatekeeper/gatekeeper-operator/api/v1alpha1"
	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
	test "github.com/gatekeeper/gatekeeper-operator/test/e2e/util"
)

var namespace = "mygatekeeper"

func TestDeployWebhookConfigs(t *testing.T) {
	g := NewWithT(t)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// Test default (nil) webhook configurations
	// ValidatingWebhookConfiguration nil
	// MutatingWebhookConfiguration nil
	deleteWebhookAssets, applyOrderedAssets, applyWebhookAssets, deleteCRDAssets := getStaticAssets(gatekeeper)
	g.Expect(applyWebhookAssets).To(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).To(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).NotTo(ContainElements(MutatingCRDs))
	g.Expect(applyOrderedAssets).To(ContainElements(MutatingCRDs))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingCRDs))
	g.Expect(deleteCRDAssets).NotTo(ContainElements(MutatingCRDs))

	webhookEnabled := operatorv1alpha1.WebhookEnabled
	webhookDisabled := operatorv1alpha1.WebhookDisabled

	// ValidatingWebhookConfiguration enabled
	// MutatingWebhookConfiguration enabled
	gatekeeper.Spec.ValidatingWebhook = &webhookEnabled
	gatekeeper.Spec.MutatingWebhook = &webhookEnabled
	deleteWebhookAssets, applyOrderedAssets, applyWebhookAssets, deleteCRDAssets = getStaticAssets(gatekeeper)
	g.Expect(applyOrderedAssets).To(ContainElements(MutatingCRDs))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).To(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).To(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingCRDs))
	g.Expect(deleteCRDAssets).NotTo(ContainElements(MutatingCRDs))

	// ValidatingWebhookConfiguration enabled
	// MutatingWebhookConfiguration disabled
	gatekeeper.Spec.ValidatingWebhook = &webhookEnabled
	gatekeeper.Spec.MutatingWebhook = &webhookDisabled
	deleteWebhookAssets, applyOrderedAssets, applyWebhookAssets, deleteCRDAssets = getStaticAssets(gatekeeper)
	g.Expect(applyWebhookAssets).To(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElements(MutatingCRDs))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).To(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingCRDs))
	g.Expect(deleteCRDAssets).To(ContainElements(MutatingCRDs))

	// ValidatingWebhookConfiguration disabled
	// MutatingWebhookConfiguration enabled
	gatekeeper.Spec.ValidatingWebhook = &webhookDisabled
	gatekeeper.Spec.MutatingWebhook = &webhookEnabled
	deleteWebhookAssets, applyOrderedAssets, applyWebhookAssets, deleteCRDAssets = getStaticAssets(gatekeeper)
	g.Expect(applyWebhookAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).To(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).To(ContainElements(MutatingCRDs))
	g.Expect(deleteWebhookAssets).To(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingCRDs))
	g.Expect(deleteCRDAssets).NotTo(ContainElements(MutatingCRDs))

	// ValidatingWebhookConfiguration disabled
	// MutatingWebhookConfiguration disabled
	gatekeeper.Spec.ValidatingWebhook = &webhookDisabled
	gatekeeper.Spec.MutatingWebhook = &webhookDisabled
	deleteWebhookAssets, applyOrderedAssets, applyWebhookAssets, deleteCRDAssets = getStaticAssets(gatekeeper)
	g.Expect(applyWebhookAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyWebhookAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElement(MutatingWebhookConfiguration))
	g.Expect(applyOrderedAssets).NotTo(ContainElements(MutatingCRDs))
	g.Expect(deleteWebhookAssets).To(ContainElement(ValidatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).To(ContainElement(MutatingWebhookConfiguration))
	g.Expect(deleteWebhookAssets).NotTo(ContainElement(MutatingCRDs))
	g.Expect(deleteCRDAssets).To(ContainElements(MutatingCRDs))
}

func TestGetSubsetOfAssets(t *testing.T) {
	g := NewWithT(t)
	g.Expect(getSubsetOfAssets(orderedStaticAssets)).To(Equal(orderedStaticAssets))
	g.Expect(getSubsetOfAssets(orderedStaticAssets, orderedStaticAssets...)).To(HaveLen(0))
	g.Expect(getSubsetOfAssets(orderedStaticAssets, MutatingCRDs...)).To(HaveLen(len(orderedStaticAssets) - len(MutatingCRDs)))
}

func TestCustomNamespace(t *testing.T) {
	g := NewWithT(t)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	expectedNamespace := "otherNamespace"
	// Rolebinding namespace overrides
	rolebindingObj, err := util.GetManifestObject(RoleBindingFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rolebindingObj).ToNot(BeNil())
	err = crOverrides(gatekeeper, RoleBindingFile, rolebindingObj, expectedNamespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	subjects, found, err := unstructured.NestedSlice(rolebindingObj.Object, "subjects")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	for _, s := range subjects {
		subject := s.(map[string]interface{})
		g.Expect(subject).NotTo(BeNil())
		ns, found, err := unstructured.NestedString(subject, "namespace")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeTrue())
		g.Expect(ns).To(Equal(expectedNamespace))
	}

	// WebhookFile namespace overrides
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, expectedNamespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(ExemptNamespaceArg, expectedNamespace))

	// ValidatingWebhookConfiguration and MutatingWebhookConfiguration namespace overrides
	webhookConfigs := []string{
		ValidatingWebhookConfiguration,
		MutatingWebhookConfiguration,
	}
	for _, webhookConfigFile := range webhookConfigs {
		webhookConfig, err := util.GetManifestObject(webhookConfigFile)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(webhookConfig).ToNot(BeNil())
		err = crOverrides(gatekeeper, webhookConfigFile, webhookConfig, expectedNamespace, false, false)
		g.Expect(err).ToNot(HaveOccurred())

		assertWebhooksWithFn(g, webhookConfig, func(webhook map[string]interface{}) {
			ns, found, err := unstructured.NestedString(webhook, "clientConfig", "service", "namespace")
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(found).To(BeTrue())
			g.Expect(ns).To(Equal(expectedNamespace))
		})
	}
}

func TestReplicas(t *testing.T) {
	g := NewWithT(t)
	auditReplicaOverride := int32(4)
	webhookReplicaOverride := int32(7)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default audit replicas
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	testObjReplicas(g, auditObj, test.DefaultDeployment.AuditReplicas)
	// test nil audit replicas
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	testObjReplicas(g, auditObj, test.DefaultDeployment.AuditReplicas)
	// test audit replicas override
	gatekeeper.Spec.Audit = &operatorv1alpha1.AuditConfig{Replicas: &auditReplicaOverride}
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	testObjReplicas(g, auditObj, auditReplicaOverride)

	// test default webhook replicas
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	testObjReplicas(g, webhookObj, test.DefaultDeployment.WebhookReplicas)
	// test nil webhook replicas
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	testObjReplicas(g, webhookObj, test.DefaultDeployment.WebhookReplicas)
	// test webhook replicas override
	gatekeeper.Spec.Webhook = &operatorv1alpha1.WebhookConfig{Replicas: &webhookReplicaOverride}
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	testObjReplicas(g, webhookObj, webhookReplicaOverride)
}

func testObjReplicas(g *WithT, obj *unstructured.Unstructured, expectedReplicas int32) {
	replicas, found, err := unstructured.NestedInt64(obj.Object, "spec", "replicas")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(int32(replicas)).To(BeIdenticalTo(expectedReplicas))
}

func TestAffinity(t *testing.T) {
	g := NewWithT(t)
	affinity := &corev1.Affinity{
		PodAffinity: &corev1.PodAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: []corev1.PodAffinityTerm{
				{
					LabelSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"auditKey": "auditValue",
						},
					},
				},
			},
		},
	}
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default affinity
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertAuditAffinity(g, auditObj, nil)
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertWebhookAffinity(g, webhookObj, nil)

	// test nil affinity
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertAuditAffinity(g, auditObj, nil)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertWebhookAffinity(g, webhookObj, nil)

	// test affinity override
	gatekeeper.Spec.Affinity = affinity

	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertAuditAffinity(g, auditObj, affinity)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertWebhookAffinity(g, webhookObj, affinity)
}

func assertAuditAffinity(g *WithT, obj *unstructured.Unstructured, expected *corev1.Affinity) {
	g.Expect(obj).ToNot(BeNil())
	current, found, err := unstructured.NestedFieldCopy(obj.Object, "spec", "template", "spec", "affinity")
	g.Expect(err).ToNot(HaveOccurred())
	if expected == nil {
		g.Expect(found).To(BeFalse())
	} else {
		g.Expect(found).To(BeTrue())
		assertAffinity(g, expected, current)
	}
}

func assertWebhookAffinity(g *WithT, obj *unstructured.Unstructured, expected *corev1.Affinity) {
	g.Expect(obj).ToNot(BeNil())
	current, found, err := unstructured.NestedFieldCopy(obj.Object, "spec", "template", "spec", "affinity")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	if expected == nil {
		assertAffinity(g, test.DefaultDeployment.Affinity, current)
	} else {
		assertAffinity(g, expected, current)
	}
}

func assertAffinity(g *WithT, expected *corev1.Affinity, current interface{}) {
	g.Expect(util.ToMap(expected)).To(BeEquivalentTo(util.ToMap(current)))
}

func TestOpenShiftOverrides(t *testing.T) {
	g := NewWithT(t)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}

	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())

	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())

	// Test that no OpenShift overrides take place when it's not OpenShift
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertOverrides(g, auditObj, true)

	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertOverrides(g, webhookObj, true)

	// Test that OpenShift overrides take place
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, true, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertOverrides(g, auditObj, false)

	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, true, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertOverrides(g, webhookObj, false)
}

func assertOverrides(g *WithT, current *unstructured.Unstructured, isSet bool) {
	containers, _, err := unstructured.NestedSlice(current.Object, "spec", "template", "spec", "containers")
	g.ExpectWithOffset(1, err).ToNot(HaveOccurred())
	g.ExpectWithOffset(1, containers).ToNot(BeEmpty())

	for i := range containers {
		container, ok := containers[i].(map[string]interface{})
		if !ok {
			continue
		}

		_, runAsUserFound, err := unstructured.NestedInt64(container, "securityContext", "runAsUser")
		g.ExpectWithOffset(1, err).ToNot(HaveOccurred())
		g.ExpectWithOffset(1, runAsUserFound).To(Equal(isSet))

		_, runAsGroupFound, err := unstructured.NestedInt64(container, "securityContext", "runAsGroup")
		g.ExpectWithOffset(1, err).ToNot(HaveOccurred())
		g.ExpectWithOffset(1, runAsGroupFound).To(Equal(isSet))

		_, seccompProfileFound, err := unstructured.NestedMap(container, "securityContext", "seccompProfile")
		g.ExpectWithOffset(1, err).ToNot(HaveOccurred())
		g.ExpectWithOffset(1, seccompProfileFound).To(Equal(isSet))
	}
}

func TestNodeSelector(t *testing.T) {
	g := NewWithT(t)
	nodeSelector := map[string]string{
		"region": "emea",
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default nodeSelector
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, auditObj, nil)
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, auditObj, nil)
	assertNodeSelector(g, webhookObj, nil)

	// test nil nodeSelector
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, auditObj, nil)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, webhookObj, nil)

	// test nodeSelector override
	gatekeeper.Spec.NodeSelector = nodeSelector
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, auditObj, nodeSelector)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNodeSelector(g, webhookObj, nodeSelector)
}

func assertNodeSelector(g *WithT, obj *unstructured.Unstructured, expected map[string]string) {
	g.Expect(obj).NotTo(BeNil())
	current, found, err := unstructured.NestedStringMap(obj.Object, "spec", "template", "spec", "nodeSelector")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	if expected == nil {
		g.Expect(test.DefaultDeployment.NodeSelector).To(BeEquivalentTo(current))
	} else {
		g.Expect(expected).To(BeEquivalentTo(current))
	}
}

func TestPodAnnotations(t *testing.T) {
	g := NewWithT(t)
	podAnnotations := map[string]string{
		"my.annotation/foo":         "example",
		"some.other.annotation/bar": "baz",
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default podAnnotations
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, auditObj, nil)
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, webhookObj, nil)

	// test nil podAnnotations
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, auditObj, nil)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, webhookObj, nil)

	// test podAnnotations override
	gatekeeper.Spec.PodAnnotations = podAnnotations
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, auditObj, podAnnotations)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertPodAnnotations(g, webhookObj, podAnnotations)
}

func assertPodAnnotations(g *WithT, obj *unstructured.Unstructured, expected map[string]string) {
	g.Expect(obj).NotTo(BeNil())
	current, found, err := unstructured.NestedStringMap(obj.Object, "spec", "template", "metadata", "annotations")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	if expected == nil {
		g.Expect(test.DefaultDeployment.PodAnnotations).To(BeEquivalentTo(current))
	} else {
		g.Expect(expected).To(BeEquivalentTo(current))
	}
}

func TestTolerations(t *testing.T) {
	g := NewWithT(t)
	tolerations := []corev1.Toleration{
		{
			Key:      "example",
			Operator: corev1.TolerationOpExists,
			Effect:   corev1.TaintEffectNoSchedule,
		},
		{
			Key:      "example2",
			Operator: corev1.TolerationOpExists,
			Effect:   corev1.TaintEffectPreferNoSchedule,
		},
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default tolerations
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, auditObj, nil)
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, webhookObj, nil)

	// test nil tolerations
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, auditObj, nil)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, webhookObj, nil)

	// test tolerations override
	gatekeeper.Spec.Tolerations = tolerations
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, auditObj, tolerations)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertTolerations(g, webhookObj, tolerations)
}

func assertTolerations(g *WithT, obj *unstructured.Unstructured, expected []corev1.Toleration) {
	g.Expect(obj).NotTo(BeNil())
	current, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "tolerations")
	g.Expect(err).ToNot(HaveOccurred())
	if expected == nil {
		g.Expect(found).To(BeFalse())
	} else {
		for i, toleration := range expected {
			g.Expect(util.ToMap(toleration)).To(BeEquivalentTo(current[i]))
		}
	}
}

func TestResources(t *testing.T) {
	g := NewWithT(t)
	audit := &operatorv1alpha1.AuditConfig{
		Resources: &corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("2000M"),
				corev1.ResourceMemory: resource.MustParse("1024Mi"),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("200m"),
				corev1.ResourceMemory: resource.MustParse("512"),
			},
		},
	}
	webhook := &operatorv1alpha1.WebhookConfig{
		Resources: &corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("2001M"),
				corev1.ResourceMemory: resource.MustParse("1025Mi"),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("201m"),
				corev1.ResourceMemory: resource.MustParse("513"),
			},
		},
	}
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default resources
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, auditObj, nil)
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, webhookObj, nil)

	// test nil resources
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, auditObj, nil)
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, webhookObj, nil)

	// test resources override
	gatekeeper.Spec.Audit = audit
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, auditObj, audit.Resources)

	gatekeeper.Spec.Webhook = webhook
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertResources(g, webhookObj, webhook.Resources)
}

func assertResources(g *WithT, obj *unstructured.Unstructured, expected *corev1.ResourceRequirements) {
	g.Expect(obj).NotTo(BeNil())
	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())

	for _, c := range containers {
		current, found, err := unstructured.NestedMap(util.ToMap(c), "resources")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeTrue())
		if expected == nil {
			assertResource(g, test.DefaultDeployment.Resources, current)
		} else {
			assertResource(g, expected, current)
		}
	}
}

func assertResource(g *WithT, expected *corev1.ResourceRequirements, current map[string]interface{}) {
	g.Expect(expected.Limits.Cpu().Cmp(resource.MustParse(current["limits"].(map[string]interface{})["cpu"].(string)))).To(BeZero())
	g.Expect(expected.Limits.Memory().Cmp(resource.MustParse(current["limits"].(map[string]interface{})["memory"].(string)))).To(BeZero())
	g.Expect(expected.Requests.Cpu().Cmp(resource.MustParse(current["requests"].(map[string]interface{})["cpu"].(string)))).To(BeZero())
	g.Expect(expected.Requests.Memory().Cmp(resource.MustParse(current["requests"].(map[string]interface{})["memory"].(string)))).To(BeZero())
}

func TestImage(t *testing.T) {
	g := NewWithT(t)
	imagePullPolicy := corev1.PullIfNotPresent
	imageConfig := &operatorv1alpha1.ImageConfig{
		ImagePullPolicy: &imagePullPolicy,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}

	// test default image
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	auditObjCopy := auditObj.DeepCopy()
	assertImage(g, auditObj, auditObjCopy, nil, "")
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	webhookObjCopy := webhookObj.DeepCopy()
	assertImage(g, webhookObj, webhookObjCopy, nil, "")

	// test nil image
	auditObjCopy = auditObj.DeepCopy()
	err = crOverrides(gatekeeper, AuditFile, auditObjCopy, namespace, false, false)
	assertImage(g, auditObj, auditObjCopy, nil, "")
	webhookObjCopy = webhookObj.DeepCopy()
	err = crOverrides(gatekeeper, WebhookFile, webhookObjCopy, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertImage(g, webhookObj, webhookObjCopy, nil, "")

	// test image override
	gatekeeper.Spec.Image = imageConfig
	image := "mycustom-image/the-gatekeeper:v1.0.0"
	err = os.Setenv(GatekeeperImageEnvVar, image)
	g.Expect(err).ToNot(HaveOccurred())
	auditObjCopy = auditObj.DeepCopy()
	err = crOverrides(gatekeeper, AuditFile, auditObjCopy, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertImage(g, auditObj, auditObjCopy, imageConfig, image)
	webhookObjCopy = webhookObj.DeepCopy()
	err = crOverrides(gatekeeper, WebhookFile, webhookObjCopy, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertImage(g, webhookObj, webhookObjCopy, imageConfig, image)
}

func assertImage(g *WithT, obj, objCopy *unstructured.Unstructured, expected *operatorv1alpha1.ImageConfig, image string) {
	g.Expect(obj).NotTo(BeNil())
	defaultImage, defaultImagePullPolicy := getDefaultImageConfig(g, obj)
	containers, found, err := unstructured.NestedSlice(objCopy.Object, "spec", "template", "spec", "containers")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())

	for _, c := range containers {
		currentImage, found, err := unstructured.NestedString(util.ToMap(c), "image")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeTrue())
		if expected == nil {
			g.Expect(defaultImage).To(BeEquivalentTo(currentImage))
		} else {
			g.Expect(image).To(BeEquivalentTo(currentImage))
		}
		currentImagePullPolicy, found, err := unstructured.NestedString(util.ToMap(c), "imagePullPolicy")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeTrue())
		if expected == nil {
			g.Expect(defaultImagePullPolicy).To(BeEquivalentTo(currentImagePullPolicy))
		} else {
			g.Expect(*expected.ImagePullPolicy).To(BeEquivalentTo(currentImagePullPolicy))
		}
	}
}

func getDefaultImageConfig(g *WithT, obj *unstructured.Unstructured) (image string, imagePullPolicy corev1.PullPolicy) {
	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(containers).NotTo(BeNil())
	g.Expect(containers).To(HaveLen(1))
	image, found, err = unstructured.NestedString(containers[0].(map[string]interface{}), "image")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(image).NotTo(BeNil())
	policy, found, err := unstructured.NestedString(containers[0].(map[string]interface{}), "imagePullPolicy")
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(policy).NotTo(BeNil())
	imagePullPolicy = corev1.PullPolicy(policy)
	return image, imagePullPolicy
}

func TestFailurePolicy(t *testing.T) {
	g := NewWithT(t)

	failurePolicy := admregv1.Fail
	webhook := operatorv1alpha1.WebhookConfig{
		FailurePolicy: &failurePolicy,
	}
	mutatingWebhook := operatorv1alpha1.WebhookEnabled
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: operatorv1alpha1.GatekeeperSpec{
			MutatingWebhook: &mutatingWebhook,
		},
	}
	// test default failurePolicy
	valObj, err := util.GetManifestObject(ValidatingWebhookConfiguration)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, valObj, ValidationGatekeeperWebhook, nil)
	mutObj, err := util.GetManifestObject(MutatingWebhookConfiguration)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, mutObj, MutatingWebhookConfiguration, nil)

	// test nil failurePolicy
	err = crOverrides(gatekeeper, ValidatingWebhookConfiguration, valObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, valObj, ValidationGatekeeperWebhook, nil)
	err = crOverrides(gatekeeper, MutatingWebhookConfiguration, mutObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, mutObj, MutationGatekeeperWebhook, nil)

	// test failurePolicy override
	gatekeeper.Spec.Webhook = &webhook
	err = crOverrides(gatekeeper, ValidatingWebhookConfiguration, valObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, valObj, ValidationGatekeeperWebhook, &failurePolicy)
	err = crOverrides(gatekeeper, MutatingWebhookConfiguration, mutObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, mutObj, MutationGatekeeperWebhook, &failurePolicy)

	// test controllerDeploymentPending override
	failurePolicyIgnore := admregv1.Ignore
	gatekeeper.Spec.Webhook = &webhook
	err = crOverrides(gatekeeper, ValidatingWebhookConfiguration, valObj, namespace, false, true)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, valObj, ValidationGatekeeperWebhook, &failurePolicyIgnore)
	err = crOverrides(gatekeeper, CheckIgnoreLabelGatekeeperWebhook, valObj, namespace, false, true)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, valObj, CheckIgnoreLabelGatekeeperWebhook, &failurePolicyIgnore)
	err = crOverrides(gatekeeper, MutatingWebhookConfiguration, mutObj, namespace, false, true)
	g.Expect(err).ToNot(HaveOccurred())
	assertFailurePolicy(g, mutObj, MutationGatekeeperWebhook, &failurePolicyIgnore)
}

func assertFailurePolicy(g *WithT, obj *unstructured.Unstructured, webhookName string, expected *admregv1.FailurePolicyType) {
	assertWebhooksWithFn(g, obj, func(webhook map[string]interface{}) {
		if webhook["name"] == webhookName {
			current, found, err := unstructured.NestedString(webhook, "failurePolicy")
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(found).To(BeTrue())
			if expected == nil {
				g.Expect(test.DefaultDeployment.FailurePolicy).To(BeEquivalentTo(current))
			} else {
				g.Expect(*expected).To(BeEquivalentTo(current))
			}
		}
	})
}

func TestNamespaceSelector(t *testing.T) {
	g := NewWithT(t)

	defExpectedNamespaceSelector := metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "admission.gatekeeper.sh/ignore",
				Operator: metav1.LabelSelectorOpDoesNotExist,
			},
			{
				Key:      "kubernetes.io/metadata.name",
				Operator: metav1.LabelSelectorOpNotIn,
				Values:   []string{"gatekeeper-system"},
			},
		},
	}

	namespaceSelector := metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "admission.gatekeeper.sh/enabled",
				Operator: metav1.LabelSelectorOpExists,
			},
		},
	}
	webhook := operatorv1alpha1.WebhookConfig{
		NamespaceSelector: &namespaceSelector,
	}
	mutatingWebhook := operatorv1alpha1.WebhookEnabled
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: operatorv1alpha1.GatekeeperSpec{
			MutatingWebhook: &mutatingWebhook,
		},
	}
	// test default namespaceSelector
	valObj, err := util.GetManifestObject(ValidatingWebhookConfiguration)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, valObj, ValidationGatekeeperWebhook, &defExpectedNamespaceSelector)
	mutObj, err := util.GetManifestObject(MutatingWebhookConfiguration)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, mutObj, MutationGatekeeperWebhook, &defExpectedNamespaceSelector)

	// test nil namespaceSelector
	err = crOverrides(gatekeeper, ValidatingWebhookConfiguration, valObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, valObj, ValidationGatekeeperWebhook, nil)
	err = crOverrides(gatekeeper, MutatingWebhookConfiguration, mutObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, mutObj, MutationGatekeeperWebhook, nil)

	// test namespaceSelector override
	gatekeeper.Spec.Webhook = &webhook
	err = crOverrides(gatekeeper, ValidatingWebhookConfiguration, valObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, valObj, ValidatingWebhookConfiguration, &namespaceSelector)
	err = crOverrides(gatekeeper, MutatingWebhookConfiguration, mutObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	assertNamespaceSelector(g, mutObj, MutatingWebhookConfiguration, &namespaceSelector)
}

func assertNamespaceSelector(g *WithT, obj *unstructured.Unstructured, webhookName string, expected *metav1.LabelSelector) {
	assertWebhooksWithFn(g, obj, func(webhook map[string]interface{}) {
		if webhook["name"] == webhookName {
			current, found, err := unstructured.NestedFieldCopy(webhook, "namespaceSelector")
			g.Expect(err).ToNot(HaveOccurred())
			if expected == nil {
				// ValidatingWebhookConfiguration and
				// MutatingWebhookConfiguration have the same defaults.
				g.Expect(found).To(BeTrue())
				g.Expect(util.ToMap(test.DefaultDeployment.NamespaceSelector)).To(BeEquivalentTo(current))
			} else {
				g.Expect(util.ToMap(*expected)).To(BeEquivalentTo(current))
			}
		}
	})
}

func assertWebhooksWithFn(g *WithT, obj *unstructured.Unstructured, webhookFn func(map[string]interface{})) {
	g.Expect(obj).NotTo(BeNil())
	webhooks, found, err := unstructured.NestedSlice(obj.Object, "webhooks")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(webhooks).ToNot(BeNil())

	for _, w := range webhooks {
		webhook := w.(map[string]interface{})
		webhookFn(webhook)
	}
}

func TestAuditInterval(t *testing.T) {
	g := NewWithT(t)
	auditInterval := metav1.Duration{
		Duration: time.Hour,
	}
	auditOverride := operatorv1alpha1.AuditConfig{
		AuditInterval: &auditInterval,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditIntervalArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditIntervalArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditIntervalArg, "3600"))
}

func TestAuditLogLevel(t *testing.T) {
	g := NewWithT(t)
	logLevel := operatorv1alpha1.LogLevelDEBUG
	auditOverride := operatorv1alpha1.AuditConfig{
		LogLevel: &logLevel,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(LogLevelArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(LogLevelArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
}

func TestAuditConstraintViolationLimit(t *testing.T) {
	g := NewWithT(t)
	constraintViolationLimit := uint64(20)
	auditOverride := operatorv1alpha1.AuditConfig{
		ConstraintViolationLimit: &constraintViolationLimit,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(ConstraintViolationLimitArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(ConstraintViolationLimitArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(ConstraintViolationLimitArg, "20"))
}

func TestAuditChunkSize(t *testing.T) {
	g := NewWithT(t)
	auditChunkSize := uint64(10)
	auditOverride := operatorv1alpha1.AuditConfig{
		AuditChunkSize: &auditChunkSize,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditChunkSizeArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditChunkSizeArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditChunkSizeArg, "10"))
}

func TestAuditFromCache(t *testing.T) {
	g := NewWithT(t)
	auditFromCache := operatorv1alpha1.AuditFromCacheEnabled
	auditOverride := operatorv1alpha1.AuditConfig{
		AuditFromCache: &auditFromCache,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditFromCacheArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditFromCacheArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditFromCacheArg, "true"))
}

func TestEmitAuditEvents(t *testing.T) {
	g := NewWithT(t)
	emitEvents := operatorv1alpha1.EmitEventsEnabled
	auditOverride := operatorv1alpha1.AuditConfig{
		EmitAuditEvents: &emitEvents,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(EmitAuditEventsArg))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(EmitAuditEventsArg))
	// test override
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(EmitAuditEventsArg, "true"))
}

func TestAllAuditArgs(t *testing.T) {
	g := NewWithT(t)
	auditChunkSize := uint64(10)
	auditFromCache := operatorv1alpha1.AuditFromCacheEnabled
	constraintViolationLimit := uint64(20)
	emitEvents := operatorv1alpha1.EmitEventsEnabled
	logLevel := operatorv1alpha1.LogLevelDEBUG
	auditInterval := metav1.Duration{
		Duration: time.Hour,
	}
	auditOverride := operatorv1alpha1.AuditConfig{
		AuditChunkSize:           &auditChunkSize,
		AuditFromCache:           &auditFromCache,
		ConstraintViolationLimit: &constraintViolationLimit,
		EmitAuditEvents:          &emitEvents,
		LogLevel:                 &logLevel,
		AuditInterval:            &auditInterval,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditChunkSizeArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditFromCacheArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(ConstraintViolationLimitArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(LogLevelArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(EmitAuditEventsArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditIntervalArg))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test nil
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditChunkSizeArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditFromCacheArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(ConstraintViolationLimitArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(LogLevelArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(EmitAuditEventsArg))
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKey(AuditIntervalArg))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test override without mutation
	gatekeeper.Spec.Audit = &auditOverride
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditChunkSizeArg, "10"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditFromCacheArg, "true"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(ConstraintViolationLimitArg, "20"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(EmitAuditEventsArg, "true"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditIntervalArg, "3600"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test override with mutation
	mutatingWebhook := operatorv1alpha1.WebhookEnabled
	gatekeeper.Spec.MutatingWebhook = &mutatingWebhook
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditChunkSizeArg, "10"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditFromCacheArg, "true"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(ConstraintViolationLimitArg, "20"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(EmitAuditEventsArg, "true"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(AuditIntervalArg, "3600"))
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
}

func TestEmitAdmissionEvents(t *testing.T) {
	g := NewWithT(t)
	emitEvents := operatorv1alpha1.EmitEventsEnabled
	webhookOverride := operatorv1alpha1.WebhookConfig{
		EmitAdmissionEvents: &emitEvents,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EmitAdmissionEventsArg))
	// test nil
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EmitAdmissionEventsArg))
	// test override
	gatekeeper.Spec.Webhook = &webhookOverride
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(EmitAdmissionEventsArg, "true"))
}

func TestWebhookLogLevel(t *testing.T) {
	g := NewWithT(t)
	logLevel := operatorv1alpha1.LogLevelDEBUG
	webhookOverride := operatorv1alpha1.WebhookConfig{
		LogLevel: &logLevel,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(LogLevelArg))
	// test nil
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(LogLevelArg))
	// test override
	gatekeeper.Spec.Webhook = &webhookOverride
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
}

func TestMutationArg(t *testing.T) {
	g := NewWithT(t)

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}

	// test default
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	auditObj, err := util.GetManifestObject(AuditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test nil
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test disabled override
	mutation := operatorv1alpha1.WebhookDisabled
	gatekeeper.Spec.MutatingWebhook = &mutation
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).NotTo(HaveKeyWithValue(OperationArg, OperationMutationStatus))
	// test enabled override
	mutation = operatorv1alpha1.WebhookEnabled
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(OperationArg, OperationMutationWebhook))
	err = crOverrides(gatekeeper, AuditFile, auditObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, auditObj).To(HaveKeyWithValue(OperationArg, OperationMutationStatus))
}

func TestDisabledBuiltins(t *testing.T) {
	g := NewWithT(t)
	webhookOverride := operatorv1alpha1.WebhookConfig{
		DisabledBuiltins: []string{
			"http.send",
			"crypto.sha1",
		},
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}

	// test default
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKey(DisabledBuiltinArg))
	g.Expect(getContainerArgumentsSlice(g, managerContainer, webhookObj)).To(ContainElements(DisabledBuiltinArg + "={http.send}"))
	// test nil
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKey(DisabledBuiltinArg))
	g.Expect(getContainerArgumentsSlice(g, managerContainer, webhookObj)).To(ContainElements(DisabledBuiltinArg + "={http.send}"))
	// test override
	gatekeeper.Spec.Webhook = &webhookOverride
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(getContainerArgumentsSlice(g, managerContainer, webhookObj)).To(ContainElements(DisabledBuiltinArg+"=http.send", DisabledBuiltinArg+"=crypto.sha1"))
}

func TestAllWebhookArgs(t *testing.T) {
	g := NewWithT(t)
	emitEvents := operatorv1alpha1.EmitEventsEnabled
	logLevel := operatorv1alpha1.LogLevelDEBUG
	webhookOverride := operatorv1alpha1.WebhookConfig{
		EmitAdmissionEvents: &emitEvents,
		LogLevel:            &logLevel,
	}

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	webhookObj, err := util.GetManifestObject(WebhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookObj).ToNot(BeNil())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EmitAdmissionEventsArg))
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(LogLevelArg))
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	// test nil
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EmitAdmissionEventsArg))
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(LogLevelArg))
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	// test override without mutation
	gatekeeper.Spec.Webhook = &webhookOverride
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(EmitAdmissionEventsArg, "true"))
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
	expectObjContainerArgument(g, managerContainer, webhookObj).NotTo(HaveKey(EnableMutationArg))
	// test override with mutation
	mutatingWebhook := operatorv1alpha1.WebhookEnabled
	gatekeeper.Spec.MutatingWebhook = &mutatingWebhook
	err = crOverrides(gatekeeper, WebhookFile, webhookObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(EmitAdmissionEventsArg, "true"))
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(LogLevelArg, "DEBUG"))
	expectObjContainerArgument(g, managerContainer, webhookObj).To(HaveKeyWithValue(OperationArg, OperationMutationWebhook))
}

func expectObjContainerArgument(g *WithT, containerName string, obj *unstructured.Unstructured) Assertion {
	args := getContainerArgumentsMap(g, containerName, obj)
	return g.Expect(args)
}

func getContainerArgumentsMap(g *WithT, containerName string, obj *unstructured.Unstructured) map[string]string {
	argsMap := make(map[string]string)
	args := getContainerArgumentsSlice(g, containerName, obj)
	for _, arg := range args {
		key, value := util.FromArg(arg)
		argsMap[key] = value
	}
	return argsMap
}

func getContainerArgumentsSlice(g *WithT, containerName string, obj *unstructured.Unstructured) []string {
	containers, found, err := unstructured.NestedSlice(obj.Object, "spec", "template", "spec", "containers")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())

	for _, c := range containers {
		cName, found, err := unstructured.NestedString(util.ToMap(c), "name")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeTrue())
		if cName == containerName {
			args, found, err := unstructured.NestedStringSlice(util.ToMap(c), "args")
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(found).To(BeTrue())
			return args
		}
	}
	return nil
}

func TestSetCertNamespace(t *testing.T) {
	g := NewWithT(t)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}
	// test default
	serverCertObj, err := util.GetManifestObject(ServerCertFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(serverCertObj).ToNot(BeNil())
	err = crOverrides(gatekeeper, ServerCertFile, serverCertObj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(serverCertObj.GetNamespace()).To(Equal(namespace))
}

func TestMutationRBACConfig(t *testing.T) {
	g := NewWithT(t)

	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}

	clusterRoleObj, err := util.GetManifestObject(ClusterRoleFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(clusterRoleObj).ToNot(BeNil())
	obj := &unstructured.Unstructured{}

	// Test default RBAC config
	obj = clusterRoleObj.DeepCopy()
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err := unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	matchCount := 0
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			if found {
				matchCount++
			}
		}
	}
	g.Expect(matchCount).To(Equal(len(matchMutatingRBACRuleFns)))

	// Test RBAC config when mutating webhook mode is nil
	obj = clusterRoleObj.DeepCopy()
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	matchCount = 0
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			if found {
				matchCount++
			}
		}
	}
	g.Expect(matchCount).To(Equal(len(matchMutatingRBACRuleFns)))

	// Test RBAC config when mutating webhook mode is nil
	obj = clusterRoleObj.DeepCopy()
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())

	// Test RBAC config when mutating webhook mode is disabled
	obj = clusterRoleObj.DeepCopy()
	mutation := operatorv1alpha1.WebhookDisabled
	gatekeeper.Spec.MutatingWebhook = &mutation
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(found).To(BeFalse())
		}
	}

	// Test RBAC config when mutating webhook mode is enabled
	obj = clusterRoleObj.DeepCopy()
	mutation = operatorv1alpha1.WebhookEnabled
	gatekeeper.Spec.MutatingWebhook = &mutation
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	matchCount = 0
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			if found {
				matchCount++
			}
		}
	}
	g.Expect(matchCount).To(Equal(len(matchMutatingRBACRuleFns)))

	// Test RBAC config when mutating webhook mode is disabled
	obj = clusterRoleObj.DeepCopy()
	mutation = operatorv1alpha1.WebhookDisabled
	gatekeeper.Spec.MutatingWebhook = &mutation
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(found).To(BeFalse())
		}
	}

	// Test RBAC config when mutating webhook mode is enabled
	obj = clusterRoleObj.DeepCopy()
	mutation = operatorv1alpha1.WebhookEnabled
	gatekeeper.Spec.MutatingWebhook = &mutation
	err = crOverrides(gatekeeper, ClusterRoleFile, obj, namespace, false, false)
	g.Expect(err).ToNot(HaveOccurred())
	rules, found, err = unstructured.NestedSlice(obj.Object, "rules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(rules).NotTo(BeEmpty())
	matchCount = 0
	for _, rule := range rules {
		r := rule.(map[string]interface{})
		for _, f := range matchMutatingRBACRuleFns {
			found, err := f(r)
			g.Expect(err).ToNot(HaveOccurred())
			if found {
				matchCount++
			}
		}
	}
	g.Expect(matchCount).To(Equal(len(matchMutatingRBACRuleFns)))
}
