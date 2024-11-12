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

// FilterSpec defines the desired state of Filter
type FilterSpec struct {

	// +kubebuilder:validation:Required
	Metadata FilterMetadata `json:"metadata"`

	// +kubebuilder:validation:Required
	Spec FilterSpecSpec `json:"spec"`
}

type FilterMetadata struct {
	Name string `json:"name,omitempty"`
}

type FilterSpecSpec struct {
	FilterGroup FilterSpecSpecFilterGroup `json:"filterGroup,omitempty"`
}

type FilterSpecSpecFilterGroup struct {
	Conjunction string `json:"conjunction,omitempty"`

	Filters FilterSpecSpecFilterGroupFilters `json:"filters,omitempty"`
}

type FilterSpecSpecFilterGroupFilters struct {
	Operator string `json:"operator,omitempty"`

	Values []string `json:"values,omitempty"`

	Key string `json:"key,omitempty"`

	Negation bool `json:"negation,omitempty"`
}

// FilterStatus defines the observed state of Filter
type FilterStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Filter is the Schema for the filters API
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilterSpec   `json:"spec,omitempty"`
	Status FilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FilterList contains a list of Filter
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Filter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Filter{}, &FilterList{})
}
