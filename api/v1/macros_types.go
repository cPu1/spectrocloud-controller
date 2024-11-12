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

// MacrosSpec defines the desired state of Macros
type MacrosSpec struct {

	// +kubebuilder:validation:Required
	Macros map[string]string `json:"macros"`

	Project string `json:"project,omitempty"`
}

// MacrosStatus defines the observed state of Macros
type MacrosStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Macros is the Schema for the macros API
type Macros struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MacrosSpec   `json:"spec,omitempty"`
	Status MacrosStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MacrosList contains a list of Macros
type MacrosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macros `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Macros{}, &MacrosList{})
}
