/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSClusterProviderStatus contains the status fields
// relevant to AWS in the cluster object.
// +k8s:openapi-gen=true
type AWSClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Region  string   `json:"region"`
	Network Network  `json:"network"`
	Bastion Instance `json:"bastion"`

	// CACertificate is a PEM encoded CA Certificate for the control plane nodes.
	CACertificate []byte

	// CAPrivateKey is a PEM encoded PKCS1 CA PrivateKey for the control plane nodes.
	CAPrivateKey []byte
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSClusterProviderStatusList contains a list of AWSClusterProviderStatus
type AWSClusterProviderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSClusterProviderStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSClusterProviderStatus{}, &AWSClusterProviderStatusList{})
}
