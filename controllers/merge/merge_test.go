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

package merge

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
)

func TestRetainClusterObjectFields(t *testing.T) {
	g := NewWithT(t)

	testCases := map[string]struct {
		desiredCABundle string
		clusterCABundle string
	}{
		"cluster CABundle is empty": {
			desiredCABundle: "ZGVzaXJlZCBkZWZhdWx0IHZhbHVlCg==",
			clusterCABundle: "Cg==",
		},
		"cluster CABundle is set": {
			desiredCABundle: "Cg==",
			clusterCABundle: "Y2x1c3RlciBDQUJ1bmRsZSBpcyBzZXQK",
		},
	}

	webhookConfigKinds := []string{
		util.ValidatingWebhookConfigurationKind,
		util.MutatingWebhookConfigurationKind,
	}
	for testName, testCase := range testCases {
		for _, kind := range webhookConfigKinds {
			t.Run(testName, func(t *testing.T) {
				desiredObj := &unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind": kind,
						"webhooks": []interface{}{
							map[string]interface{}{
								"clientConfig": map[string]interface{}{
									"caBundle": testCase.desiredCABundle,
								},
							},
						},
					},
				}
				clusterObj := &unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind": kind,
						"webhooks": []interface{}{
							map[string]interface{}{
								"clientConfig": map[string]interface{}{
									"caBundle": testCase.clusterCABundle,
								},
							},
						},
					},
				}

				err := RetainClusterObjectFields(desiredObj, clusterObj)
				g.Expect(err).ToNot(HaveOccurred())

				desiredWebhooks, found, err := unstructured.NestedSlice(desiredObj.Object, "webhooks")
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(found).To(BeTrue())
				g.Expect(desiredWebhooks).ToNot(BeNil())

				desiredCABundle, found, err := unstructured.NestedString(desiredWebhooks[0].(map[string]interface{}), "clientConfig", "caBundle")
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(found).To(BeTrue())
				g.Expect(desiredCABundle).To(Equal(testCase.clusterCABundle))
			})
		}
	}
}
