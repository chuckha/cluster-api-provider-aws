package main_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
	corev1 "sigs.k8s.io/cluster-api/api/v1alpha2"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha2"
	"sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
	"sigs.k8s.io/cluster-api/test/framework"
)

var _ = Describe("CAPA", func() {
	Describe("Single Node Cluster", func() {
		namespace := "default"
		region := "us-west-2"
		sshKeyName := "work-laptop"

		infraCluster := &infrav1.AWSCluster{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      "my-cluster",
			},
			Spec: infrav1.AWSClusterSpec{
				Region:     region,
				SSHKeyName: sshKeyName,
			},
		}

		cluster := &corev1.Cluster{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      "my-cluster",
			},
			Spec: corev1.ClusterSpec{
				ClusterNetwork: &corev1.ClusterNetwork{
					Services: &corev1.NetworkRanges{CIDRBlocks: []string{}},
					Pods:     &corev1.NetworkRanges{CIDRBlocks: []string{"192.168.0.0/16"}},
				},
				InfrastructureRef: &v1.ObjectReference{
					APIVersion: infrav1.GroupVersion.String(),
					Kind:       framework.TypeToKind(infraCluster),
					Namespace:  infraCluster.GetNamespace(),
					Name:       infraCluster.GetName(),
				},
			},
		}

		nodeGen := &NodeGenerator{
			sshKeyName: sshKeyName,
		}

		Context("One node cluster", func() {
			It("should create a single node cluster", func() {
				node := nodeGen.GenerateNode("my-cluster")
				input := &framework.OneNodeClusterInput{
					Mgmt:          mgmt,
					Cluster:       cluster,
					InfraCluster:  infraCluster,
					Node:          node,
					CreateTimeout: 3 * time.Minute,
				}

				framework.OneNodeCluster(input)

				cleanupInput := &framework.CleanUpInput{
					Mgmt:          mgmt,
					Cluster:       cluster,
					DeleteTimeout: 5 * time.Minute,
				}

				framework.CleanUp(cleanupInput)
			})
		})

		FContext("Multi-node controlplane cluster", func() {
			It("should create a multi-node controlplane cluster", func() {
				nodes := make([]framework.Node, 3)
				for i := range nodes {
					nodes[i] = nodeGen.GenerateNode("my-cluster")
				}

				input := &framework.MultiNodeControlplaneClusterInput{
					Mgmt:              mgmt,
					Cluster:           cluster,
					InfraCluster:      infraCluster,
					ControlplaneNodes: nodes,
					CreateTimeout:     5 * time.Minute,
				}
				framework.MultiNodeControlPlaneCluster(input)

				cleanupInput := &framework.CleanUpInput{
					Mgmt:          mgmt,
					Cluster:       cluster,
					DeleteTimeout: 10 * time.Minute,
				}

				framework.CleanUp(cleanupInput)
			})
		})
	})
})

type NodeGenerator struct {
	counter    int
	sshKeyName string
}

func (n *NodeGenerator) GenerateNode(clusterName string) framework.Node {
	generatedName := fmt.Sprintf("controlplane-%d", n.counter)
	n.counter += 1

	namespace := "default"
	version := "v1.15.3"

	controlPlaneInstanceType := "t2.medium"
	controlPlaneInstanceProfile := "control-plane.cluster-api-provider-aws.sigs.k8s.io"
	sshKeyName := n.sshKeyName

	infraMachine := &infrav1.AWSMachine{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: infrav1.AWSMachineSpec{
			InstanceType:       controlPlaneInstanceType,
			IAMInstanceProfile: controlPlaneInstanceProfile,
			SSHKeyName:         sshKeyName,
		},
	}

	bootstrapConfig := &bootstrapv1.KubeadmConfig{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: bootstrapv1.KubeadmConfigSpec{
			ClusterConfiguration: &v1beta1.ClusterConfiguration{
				APIServer: v1beta1.APIServer{
					ControlPlaneComponent: v1beta1.ControlPlaneComponent{
						ExtraArgs: map[string]string{
							"cloud-provider": "aws",
						},
					},
				},
				ControllerManager: v1beta1.ControlPlaneComponent{
					ExtraArgs: map[string]string{
						"cloud-provider": "aws",
					},
				},
			},
			InitConfiguration: &v1beta1.InitConfiguration{
				NodeRegistration: v1beta1.NodeRegistrationOptions{
					KubeletExtraArgs: map[string]string{
						"cloud-provider": "aws",
					},
					Name: "{{ ds.meta_data.hostname }}",
				},
			},
			JoinConfiguration: &v1beta1.JoinConfiguration{
				NodeRegistration: v1beta1.NodeRegistrationOptions{
					KubeletExtraArgs: map[string]string{
						"cloud-provider": "aws",
					},
					Name: "{{ ds.meta_data.hostname }}",
				},
			},
		},
	}

	machine := &corev1.Machine{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
			Labels: map[string]string{
				"cluster.x-k8s.io/control-plane": "true",
				"cluster.x-k8s.io/cluster-name":  clusterName,
			},
		},
		Spec: corev1.MachineSpec{
			Bootstrap: corev1.Bootstrap{
				ConfigRef: &v1.ObjectReference{
					APIVersion: bootstrapv1.GroupVersion.String(),
					Kind:       framework.TypeToKind(bootstrapConfig),
					Namespace:  bootstrapConfig.GetNamespace(),
					Name:       bootstrapConfig.GetName(),
				},
			},
			InfrastructureRef: v1.ObjectReference{
				APIVersion: infrav1.GroupVersion.String(),
				Kind:       framework.TypeToKind(infraMachine),
				Namespace:  infraMachine.GetNamespace(),
				Name:       infraMachine.GetName(),
			},
			Version: &version,
		},
	}
	return framework.Node{
		Machine:         machine,
		InfraMachine:    infraMachine,
		BootstrapConfig: bootstrapConfig,
	}
}
