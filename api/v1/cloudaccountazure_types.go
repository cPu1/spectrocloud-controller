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

// CloudaccountAzureSpec defines the desired state of CloudaccountAzure
type CloudaccountAzureSpec struct {

	// +kubebuilder:validation:Required
	AzureClientId string `json:"azureClientId"`

	// +kubebuilder:validation:Required
	AzureClientSecret string `json:"azureClientSecret"`

	// +kubebuilder:validation:Required
	AzureTenantId string `json:"azureTenantId"`

	Cloud string `json:"cloud,omitempty"`

	Context string `json:"context,omitempty"`

	DisablePropertiesRequest bool `json:"disablePropertiesRequest,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	PrivateCloudGatewayId string `json:"privateCloudGatewayId,omitempty"`

	TenantName string `json:"tenantName,omitempty"`
}

// CloudaccountAzureStatus defines the observed state of CloudaccountAzure
type CloudaccountAzureStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountAzure is the Schema for the cloudaccountazures API
type CloudaccountAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountAzureSpec   `json:"spec,omitempty"`
	Status CloudaccountAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountAzureList contains a list of CloudaccountAzure
type CloudaccountAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountAzure `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountAzure{}, &CloudaccountAzureList{})
}
