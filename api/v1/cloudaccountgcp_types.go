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

// CloudaccountGcpSpec defines the desired state of CloudaccountGcp
type CloudaccountGcpSpec struct {
	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	GcpJsonCredentials string `json:"gcpJsonCredentials"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

// CloudaccountGcpStatus defines the observed state of CloudaccountGcp
type CloudaccountGcpStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountGcp is the Schema for the cloudaccountgcps API
type CloudaccountGcp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountGcpSpec   `json:"spec,omitempty"`
	Status CloudaccountGcpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountGcpList contains a list of CloudaccountGcp
type CloudaccountGcpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountGcp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountGcp{}, &CloudaccountGcpList{})
}
