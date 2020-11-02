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
	"github.com/openshift/library-go/pkg/manifest"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestReplicas(t *testing.T) {
	auditReplicas := int64(4)
	webhookReplicas := int64(7)
	gatekeeper := &operatorv1alpha1.Gatekeeper{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "testns",
		},
	}
	// test default audit replicas
	auditManifest, err := getManifest(auditFile)
	assert.Nil(t, err)
	testManifestReplicas(t, auditManifest, int64(1))
	// test nil audit replicas
	err = crOverrides(gatekeeper, auditFile, auditManifest)
	assert.Nil(t, err)
	testManifestReplicas(t, auditManifest, int64(1))
	// test audit replicas override
	gatekeeper.Spec.Audit = &operatorv1alpha1.AuditConfig{Replicas: &auditReplicas}
	err = crOverrides(gatekeeper, auditFile, auditManifest)
	assert.Nil(t, err)
	testManifestReplicas(t, auditManifest, auditReplicas)

	// test default webhook replicas
	webhookManifest, err := getManifest(webhookFile)
	assert.Nil(t, err)
	testManifestReplicas(t, webhookManifest, int64(3))
	// test nil webhook replicas
	err = crOverrides(gatekeeper, webhookFile, webhookManifest)
	assert.Nil(t, err)
	testManifestReplicas(t, webhookManifest, int64(3))
	// test webhook replicas override
	gatekeeper.Spec.Webhook = &operatorv1alpha1.WebhookConfig{Replicas: &webhookReplicas}
	err = crOverrides(gatekeeper, webhookFile, webhookManifest)
	assert.Nil(t, err)
	testManifestReplicas(t, webhookManifest, webhookReplicas)
}

func testManifestReplicas(t *testing.T, manifest *manifest.Manifest, expectedReplicas int64) {
	replicas, found, err := unstructured.NestedInt64(manifest.Obj.Object, "spec", "replicas")
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, expectedReplicas, replicas)
}
