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

// CloudaccountMaasSpec defines the desired state of CloudaccountMaas
type CloudaccountMaasSpec struct {
	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	MaasApiEndpoint string `json:"maasApiEndpoint"`

	// +kubebuilder:validation:Required
	MaasApiKey string `json:"maasApiKey"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	PrivateCloudGatewayId string `json:"privateCloudGatewayId"`
}

// CloudaccountMaasStatus defines the observed state of CloudaccountMaas
type CloudaccountMaasStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountMaas is the Schema for the cloudaccountmaas API
type CloudaccountMaas struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountMaasSpec   `json:"spec,omitempty"`
	Status CloudaccountMaasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountMaasList contains a list of CloudaccountMaas
type CloudaccountMaasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountMaas `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountMaas{}, &CloudaccountMaasList{})
}
