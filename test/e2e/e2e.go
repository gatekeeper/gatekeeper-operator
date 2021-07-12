/*
Copyright 2021.

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
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	operatorv1alpha1 "github.com/gatekeeper/gatekeeper-operator/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var K8sClient client.Client
var testEnv *envtest.Environment
var affinityPod *corev1.Pod
var affinityNode *corev1.Node

func RunE2ETests(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	cfg, err := testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = operatorv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	K8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(K8sClient).NotTo(BeNil())

	err = extv1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	affinityNode, err = getAffinityNode()
	Expect(err).ToNot(HaveOccurred())

	if affinityNode != nil {
		Expect(labelNode(affinityNode)).Should(Succeed())
		createAffinityPod()
	}
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")

	if affinityNode != nil {
		K8sClient.Delete(ctx, affinityPod, client.PropagationPolicy(v1.DeletePropagationForeground))
		Expect(unlabelNode(affinityNode)).Should(Succeed())
		err := deleteAffinityPod()
		Expect(err).ToNot(HaveOccurred())
	}
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})

func getAffinityNode() (*corev1.Node, error) {
	nodes := &corev1.NodeList{}
	err := K8sClient.List(context.TODO(), nodes)
	if err != nil {
		return nil, err
	}
	// If true, we use a testEnv
	if len(nodes.Items) == 0 {
		return nil, nil
	}
	return &nodes.Items[0], nil
}

func labelNode(node *corev1.Node) error {
	patch := client.MergeFrom(node.DeepCopy())
	node.ObjectMeta.Labels["region"] = "EMEA"
	node.ObjectMeta.Labels["topology.kubernetes.io/zone"] = "test"
	return K8sClient.Patch(context.TODO(), node, patch)
}

func unlabelNode(node *corev1.Node) error {
	patch := client.MergeFrom(node.DeepCopy())
	delete(node.ObjectMeta.Labels, "region")
	delete(node.ObjectMeta.Labels, "topology.kubernetes.io/zone")
	return K8sClient.Patch(context.TODO(), node, patch)
}

func createAffinityPod() {
	affinityPod, err := loadAffinityPodFromFile(*GatekeeperNamespace)
	Expect(err).ToNot(HaveOccurred())
	Expect(K8sClient.Create(ctx, affinityPod)).Should(Succeed())
}

func deleteAffinityPod() error {
	affinityPodFromFile, err := loadAffinityPodFromFile(*GatekeeperNamespace)
	if err != nil {
		return err
	}

	affinityPodName := types.NamespacedName{
		Namespace: affinityPodFromFile.ObjectMeta.Namespace,
		Name:      affinityPodFromFile.ObjectMeta.Name,
	}
	pod := &corev1.Pod{}
	err = K8sClient.Get(ctx, affinityPodName, pod)
	if apierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return K8sClient.Delete(ctx, pod)
}

func loadAffinityPodFromFile(namespace string) (*corev1.Pod, error) {
	f, err := os.Open("../../config/samples/affinity_pod.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	pod := &corev1.Pod{}
	err = decodeYAML(f, pod)
	pod.ObjectMeta.Namespace = namespace
	return pod, err
}
