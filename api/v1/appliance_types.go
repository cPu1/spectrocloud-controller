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

// ApplianceSpec defines the desired state of Appliance
type ApplianceSpec struct {
	PairingKey string `json:"pairingKey,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Uid string `json:"uid"`

	Wait bool `json:"wait,omitempty"`
}

// ApplianceStatus defines the observed state of Appliance
type ApplianceStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Appliance is the Schema for the appliances API
type Appliance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplianceSpec   `json:"spec,omitempty"`
	Status ApplianceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplianceList contains a list of Appliance
type ApplianceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Appliance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Appliance{}, &ApplianceList{})
}
