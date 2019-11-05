package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-aws/test"
	"sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/cluster-api/test/framework/management/kind"
	"sigs.k8s.io/cluster-api/test/framework/providers/cabpk"
	"sigs.k8s.io/cluster-api/test/framework/providers/capi"

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
	corev1 "sigs.k8s.io/cluster-api/api/v1alpha2"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha2"
)

func TestCAPA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CAPA Suite")
}

var mgmt framework.ManagementCluster

var _ = BeforeSuite(func() {
	var err error

	// Set up the provider component generators
	core := &capi.Generator{
		Version: "v0.2.7",
	}
	infra := &test.Generator{
		Version: "v0.4.4",
	}
	bootstrap := &cabpk.Generator{
		Version: "v0.1.5",
	}
	scheme := runtime.NewScheme()
	Expect(v1.AddToScheme(scheme)).To(Succeed())
	Expect(corev1.AddToScheme(scheme)).To(Succeed())
	Expect(bootstrapv1.AddToScheme(scheme)).To(Succeed())
	Expect(infrav1.AddToScheme(scheme)).To(Succeed())

	// Create the management cluster
	mgmt, err = kind.NewCluster("mgmt", scheme)

	Expect(err).NotTo(HaveOccurred())
	Expect(mgmt).NotTo(BeNil())
	framework.InstallComponents(mgmt, core, infra, bootstrap)
	// TODO: wait for components to be ready
})

var _ = AfterSuite(func() {
	Expect(mgmt.Teardown()).NotTo(HaveOccurred())
})
