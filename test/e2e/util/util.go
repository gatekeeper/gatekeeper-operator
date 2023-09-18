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

package util

import (
	admregv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type defaultConfig struct {
	AuditReplicas     int32
	WebhookReplicas   int32
	Affinity          *corev1.Affinity
	NodeSelector      map[string]string
	PodAnnotations    map[string]string
	Resources         *corev1.ResourceRequirements
	FailurePolicy     admregv1.FailurePolicyType
	NamespaceSelector *metav1.LabelSelector
}

// DefaultDeployment is the expected default configuration to be deployed
var DefaultDeployment = defaultConfig{
	AuditReplicas:   int32(1),
	WebhookReplicas: int32(3),
	Affinity: &corev1.Affinity{
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
	},
	NodeSelector: map[string]string{
		"kubernetes.io/os": "linux",
	},
	Resources: &corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1000m"),
			corev1.ResourceMemory: resource.MustParse("512Mi"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("100m"),
			corev1.ResourceMemory: resource.MustParse("512Mi"),
		},
	},
	FailurePolicy: admregv1.Ignore,
	NamespaceSelector: &metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "admission.gatekeeper.sh/ignore",
				Operator: metav1.LabelSelectorOpDoesNotExist,
			},
			{
				Key:      "kubernetes.io/metadata.name",
				Operator: metav1.LabelSelectorOpNotIn,
				Values:   []string{"mygatekeeper"},
			},
		},
	},
}
