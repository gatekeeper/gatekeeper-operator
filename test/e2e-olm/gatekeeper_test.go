// Copyright (c) 2020 Red Hat, Inc.

package e2eolm

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const OperatorResourcesYaml string = "./resources/operator.yaml"

var _ = Describe("Test gatekeeper", func() {
	Describe("Test gatekeeper operator", func() {
		It("gatekeeper operator resources should be created on managed", func() {
			By("Creating resources on cluster")
			Kubectl("apply", "-f", OperatorResourcesYaml)
			crd := GetClusterLevelWithTimeout(clientManagedDynamic, gvrCRD, "gatekeepers.operator.gatekeeper.sh", true, 360)
			Expect(crd).NotTo(BeNil())
			Kubectl("apply", "-f", "./resources/cr.yaml")
		})
		It("should create gatekeeper pods on managed cluster", func() {
			By("Checking number of pods in gatekeeper-system ns")
			ListWithTimeoutByNamespace(clientManagedDynamic, gvrPod, metav1.ListOptions{}, "gatekeeper-system", 6, true, 240)
		})
	})
})
