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

// CloudaccountTencentSpec defines the desired state of CloudaccountTencent
type CloudaccountTencentSpec struct {
	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	TencentSecretId string `json:"tencentSecretId,omitempty"`

	TencentSecretKey string `json:"tencentSecretKey,omitempty"`
}

// CloudaccountTencentStatus defines the observed state of CloudaccountTencent
type CloudaccountTencentStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountTencent is the Schema for the cloudaccounttencents API
type CloudaccountTencent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountTencentSpec   `json:"spec,omitempty"`
	Status CloudaccountTencentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountTencentList contains a list of CloudaccountTencent
type CloudaccountTencentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountTencent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountTencent{}, &CloudaccountTencentList{})
}
