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

// SshKeySpec defines the desired state of SshKey
type SshKeySpec struct {
	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	SshKey string `json:"sshKey"`
}

// SshKeyStatus defines the observed state of SshKey
type SshKeyStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SshKey is the Schema for the sshkeys API
type SshKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SshKeySpec   `json:"spec,omitempty"`
	Status SshKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SshKeyList contains a list of SshKey
type SshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SshKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SshKey{}, &SshKeyList{})
}
