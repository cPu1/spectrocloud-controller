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

// PrivatecloudgatewayDnsMapSpec defines the desired state of PrivatecloudgatewayDnsMap
type PrivatecloudgatewayDnsMapSpec struct {

	// +kubebuilder:validation:Required
	DataCenter string `json:"dataCenter"`

	// +kubebuilder:validation:Required
	Network string `json:"network"`

	// +kubebuilder:validation:Required
	PrivateCloudGatewayId string `json:"privateCloudGatewayId"`

	// +kubebuilder:validation:Required
	SearchDomainName string `json:"searchDomainName"`
}

// PrivatecloudgatewayDnsMapStatus defines the observed state of PrivatecloudgatewayDnsMap
type PrivatecloudgatewayDnsMapStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PrivatecloudgatewayDnsMap is the Schema for the privatecloudgatewaydnsmaps API
type PrivatecloudgatewayDnsMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivatecloudgatewayDnsMapSpec   `json:"spec,omitempty"`
	Status PrivatecloudgatewayDnsMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivatecloudgatewayDnsMapList contains a list of PrivatecloudgatewayDnsMap
type PrivatecloudgatewayDnsMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivatecloudgatewayDnsMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivatecloudgatewayDnsMap{}, &PrivatecloudgatewayDnsMapList{})
}
