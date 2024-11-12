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

// RegistryOciSpec defines the desired state of RegistryOci
type RegistryOciSpec struct {

	// +kubebuilder:validation:Required
	Credentials RegistryOciCredentials `json:"credentials"`

	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// +kubebuilder:validation:Required
	IsPrivate bool `json:"isPrivate"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	ProviderType string `json:"providerType,omitempty"`

	// +kubebuilder:validation:Required
	Type string `json:"type"`
}

type RegistryOciCredentials struct {
	Arn string `json:"arn,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	CredentialType string `json:"credentialType,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	SecretKey string `json:"secretKey,omitempty"`
}

// RegistryOciStatus defines the observed state of RegistryOci
type RegistryOciStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RegistryOci is the Schema for the registryocis API
type RegistryOci struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistryOciSpec   `json:"spec,omitempty"`
	Status RegistryOciStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryOciList contains a list of RegistryOci
type RegistryOciList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryOci `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RegistryOci{}, &RegistryOciList{})
}
