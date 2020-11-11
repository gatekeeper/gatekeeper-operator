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
	"io"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/font/gatekeeper-operator/api/v1alpha1"
)

const (
	// The length of time between polls.
	pollInterval = 50 * time.Millisecond
	// How long to try before giving up.
	waitTimeout = 30 * time.Second
)

var _ = Describe("Gatekeeper", func() {
	BeforeEach(func() {
		if !useExistingCluster() {
			Skip("Test requires existing cluster. Set environment variable USE_EXISTING_CLUSTER=true and try again.")
		}
	})

	Describe("Install", func() {
		Context("Creating Gatekeeper custom resource", func() {
			It("Should install Gatekeeper", func() {
				ctx := context.Background()
				gatekeeper := &v1alpha1.Gatekeeper{}
				gatekeeper.Namespace = "gatekeeper-system"
				err := sampleGatekeeper(gatekeeper)
				Expect(err).ToNot(HaveOccurred())
				gkDeployment := &appsv1.Deployment{}

				By("Creating Gatekeeper resource", func() {
					Expect(K8sClient.Create(ctx, gatekeeper)).Should(Succeed())
				})

				By("Checking gatekeeper-controller-manager readiness", func() {
					gkName := types.NamespacedName{
						Namespace: "gatekeeper-system",
						Name:      "gatekeeper-controller-manager",
					}

					Eventually(func() (int32, error) {
						return getDeployment(ctx, gkName, gkDeployment)
					}, waitTimeout, pollInterval).Should(Equal(*gatekeeper.Spec.Webhook.Replicas))
				})

				By("Checking gatekeeper-audit readiness", func() {
					gkName := types.NamespacedName{
						Namespace: "gatekeeper-system",
						Name:      "gatekeeper-audit",
					}

					Eventually(func() (int32, error) {
						return getDeployment(ctx, gkName, gkDeployment)
					}, waitTimeout, pollInterval).Should(Equal(*gatekeeper.Spec.Audit.Replicas))
				})
			})
		})
	})
})

func sampleGatekeeper(gatekeeper *v1alpha1.Gatekeeper) error {
	f, err := os.Open("../config/samples/operator_v1alpha1_gatekeeper.yaml")
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

func getDeployment(ctx context.Context, name types.NamespacedName,
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
