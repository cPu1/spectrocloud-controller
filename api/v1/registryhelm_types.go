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

// RegistryHelmSpec defines the desired state of RegistryHelm
type RegistryHelmSpec struct {

	// +kubebuilder:validation:Required
	Credentials RegistryHelmCredentials `json:"credentials"`

	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// +kubebuilder:validation:Required
	IsPrivate bool `json:"isPrivate"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type RegistryHelmCredentials struct {
	CredentialType string `json:"credentialType,omitempty"`

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	Token string `json:"token,omitempty"`
}

// RegistryHelmStatus defines the observed state of RegistryHelm
type RegistryHelmStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RegistryHelm is the Schema for the registryhelms API
type RegistryHelm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegistryHelmSpec   `json:"spec,omitempty"`
	Status RegistryHelmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryHelmList contains a list of RegistryHelm
type RegistryHelmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryHelm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RegistryHelm{}, &RegistryHelmList{})
}
