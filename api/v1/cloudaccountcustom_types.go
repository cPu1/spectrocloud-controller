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

// CloudaccountCustomSpec defines the desired state of CloudaccountCustom
type CloudaccountCustomSpec struct {

	// +kubebuilder:validation:Required
	Cloud string `json:"cloud"`

	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	Credentials map[string]string `json:"credentials"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	PrivateCloudGatewayId string `json:"privateCloudGatewayId"`
}

// CloudaccountCustomStatus defines the observed state of CloudaccountCustom
type CloudaccountCustomStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountCustom is the Schema for the cloudaccountcustoms API
type CloudaccountCustom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountCustomSpec   `json:"spec,omitempty"`
	Status CloudaccountCustomStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountCustomList contains a list of CloudaccountCustom
type CloudaccountCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountCustom `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountCustom{}, &CloudaccountCustomList{})
}
