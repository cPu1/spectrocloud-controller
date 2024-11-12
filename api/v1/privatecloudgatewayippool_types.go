/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivatecloudgatewayIppoolSpec defines the desired state of PrivatecloudgatewayIppool
type PrivatecloudgatewayIppoolSpec struct {

	// +kubebuilder:validation:Required
	Gateway string `json:"gateway"`

	IpEndRange string `json:"ipEndRange,omitempty"`

	IpStartRange string `json:"ipStartRange,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	NameserverAddresses []string `json:"nameserverAddresses,omitempty"`

	NameserverSearchSuffix []string `json:"nameserverSearchSuffix,omitempty"`

	// +kubebuilder:validation:Required
	NetworkType string `json:"networkType"`

	// +kubebuilder:validation:Required
	Prefix int `json:"prefix"`

	// +kubebuilder:validation:Required
	PrivateCloudGatewayId string `json:"privateCloudGatewayId"`

	RestrictToSingleCluster bool `json:"restrictToSingleCluster,omitempty"`

	SubnetCidr string `json:"subnetCidr,omitempty"`
}

// PrivatecloudgatewayIppoolStatus defines the observed state of PrivatecloudgatewayIppool
type PrivatecloudgatewayIppoolStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PrivatecloudgatewayIppool is the Schema for the privatecloudgatewayippools API
type PrivatecloudgatewayIppool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivatecloudgatewayIppoolSpec   `json:"spec,omitempty"`
	Status PrivatecloudgatewayIppoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivatecloudgatewayIppoolList contains a list of PrivatecloudgatewayIppool
type PrivatecloudgatewayIppoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivatecloudgatewayIppool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivatecloudgatewayIppool{}, &PrivatecloudgatewayIppoolList{})
}
