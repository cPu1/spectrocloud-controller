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

// AlertSpec defines the desired state of Alert
type AlertSpec struct {
	AlertAllUsers bool `json:"alertAllUsers,omitempty"`

	// +kubebuilder:validation:Required
	Component string `json:"component"`

	CreatedBy string `json:"createdBy,omitempty"`

	Http AlertHttp `json:"http,omitempty"`

	Identifiers []string `json:"identifiers,omitempty"`

	// +kubebuilder:validation:Required
	IsActive bool `json:"isActive"`

	// +kubebuilder:validation:Required
	Project string `json:"project"`

	Status AlertSpecStatus `json:"status,omitempty"`

	// +kubebuilder:validation:Required
	Type string `json:"type"`
}

type AlertHttp struct {
	Method string `json:"method,omitempty"`

	Url string `json:"url,omitempty"`

	Body string `json:"body,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`
}

type AlertSpecStatus struct {
	Time string `json:"time,omitempty"`

	IsSucceeded bool `json:"isSucceeded,omitempty"`

	Message string `json:"message,omitempty"`
}

// AlertStatus defines the observed state of Alert
type AlertStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Alert is the Schema for the alerts API
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertSpec   `json:"spec,omitempty"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alert
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
