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
	"testing"

	operatorv1alpha1 "github.com/font/gatekeeper-operator/api/v1alpha1"
	. "github.com/onsi/gomega"
	"github.com/openshift/library-go/pkg/manifest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	auditReplicas   = int64(1)
	webhookReplicas = int64(3)
)

func TestReplicas(t *testing.T) {
	g := NewWithT(t)
	auditReplicaOverride := int64(4)
	webhookReplicaOverride := int64(7)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "testns",
		},
	}
	// test default audit replicas
	auditManifest, err := getManifest(auditFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(auditManifest).ToNot(BeNil())
	testManifestReplicas(t, auditManifest, auditReplicas)
	// test nil audit replicas
	err = crOverrides(gatekeeper, auditFile, auditManifest)
	g.Expect(err).ToNot(HaveOccurred())
	testManifestReplicas(t, auditManifest, auditReplicas)
	// test audit replicas override
	gatekeeper.Spec.Audit = &operatorv1alpha1.AuditConfig{Replicas: &auditReplicaOverride}
	err = crOverrides(gatekeeper, auditFile, auditManifest)
	g.Expect(err).ToNot(HaveOccurred())
	testManifestReplicas(t, auditManifest, auditReplicaOverride)

	// test default webhook replicas
	webhookManifest, err := getManifest(webhookFile)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(webhookManifest).ToNot(BeNil())
	testManifestReplicas(t, webhookManifest, webhookReplicas)
	// test nil webhook replicas
	err = crOverrides(gatekeeper, webhookFile, webhookManifest)
	g.Expect(err).ToNot(HaveOccurred())
	testManifestReplicas(t, webhookManifest, webhookReplicas)
	// test webhook replicas override
	gatekeeper.Spec.Webhook = &operatorv1alpha1.WebhookConfig{Replicas: &webhookReplicaOverride}
	err = crOverrides(gatekeeper, webhookFile, webhookManifest)
	g.Expect(err).ToNot(HaveOccurred())
	testManifestReplicas(t, webhookManifest, webhookReplicaOverride)
}

func testManifestReplicas(t *testing.T, manifest *manifest.Manifest, expectedReplicas int64) {
	g := NewWithT(t)
	replicas, found, err := unstructured.NestedInt64(manifest.Obj.Object, "spec", "replicas")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(found).To(BeTrue())
	g.Expect(replicas).To(BeIdenticalTo(expectedReplicas))
}
