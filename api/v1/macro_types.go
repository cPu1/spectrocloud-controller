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

// MacroSpec defines the desired state of Macro
type MacroSpec struct {

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Project string `json:"project,omitempty"`

	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

// MacroStatus defines the observed state of Macro
type MacroStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Macro is the Schema for the macroes API
type Macro struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MacroSpec   `json:"spec,omitempty"`
	Status MacroStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MacroList contains a list of Macro
type MacroList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macro `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Macro{}, &MacroList{})
}
