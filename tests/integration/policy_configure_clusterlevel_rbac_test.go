// Copyright (c) 2022 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package integration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	policiesv1 "open-cluster-management.io/governance-policy-propagator/api/v1"
	"open-cluster-management.io/governance-policy-propagator/test/utils"
	"github.com/liswang89/policy-openshift-plus/tests/common"
)

var _ = Describe("GRC: Test the policy-configure-clusterlevel-rbac policy", Ordered, Label("policy-collection", "stable"), func() {
	const (
		policyTestName   = "policy-configure-clusterlevel-rbac"
		policyTestNSName = "policies"
	)
	
	It("Checking that " + policyTestName + " exists on the Hub cluster", func() {
	    By("Checking that " + policyTestName + " exists on the Hub cluster")
		rootPolicy := utils.GetWithTimeout(
			clientHubDynamic, common.GvrPolicy, policyTestName, userNamespace, true, defaultTimeoutSeconds,
		)
		Expect(rootPolicy).NotTo(BeNil())
	})

	It("stable/"+policyTestName+" should be created on managed cluster", func() {
		By("Checking the policy on managed cluster in ns " + clusterNamespace)
		managedPolicy := utils.GetWithTimeout(
			clientManagedDynamic,
			common.GvrPolicy,
			userNamespace+"."+policyTestName,
			clusterNamespace,
			true,
			defaultTimeoutSeconds,
		)
		
		Expect(managedPolicy).NotTo(BeNil())
	})

	It("stable/"+policyTestName+" should be Compliant", func() {
		By("Checking if the status of the root policy is Compliant")
		Eventually(
			common.GetComplianceState(clientHubDynamic, userNamespace, policyTestName, clusterNamespace),
			defaultTimeoutSeconds*2,
			1,
		).Should(Equal(policiesv1.Compliant))
	})

	
	AfterAll(func() {
	/*	_, err := utils.KubectlWithOutput(
			"delete", "-f", policyPodURL, "-n", userNamespace, "--kubeconfig="+kubeconfigHub,
		)
		Expect(err).To(BeNil())

		err = clientManaged.CoreV1().Namespaces().Delete(
			context.TODO(), policyPodNSName, metav1.DeleteOptions{},
		)
		Expect(err).To(BeNil())*/
	})
})