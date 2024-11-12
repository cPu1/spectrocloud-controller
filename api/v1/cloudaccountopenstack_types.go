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

// CloudaccountOpenstackSpec defines the desired state of CloudaccountOpenstack
type CloudaccountOpenstackSpec struct {
	CaCertificate string `json:"caCertificate,omitempty"`

	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	DefaultDomain string `json:"defaultDomain"`

	// +kubebuilder:validation:Required
	DefaultProject string `json:"defaultProject"`

	// +kubebuilder:validation:Required
	IdentityEndpoint string `json:"identityEndpoint"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	OpenstackAllowInsecure bool `json:"openstackAllowInsecure,omitempty"`

	// +kubebuilder:validation:Required
	OpenstackPassword string `json:"openstackPassword"`

	// +kubebuilder:validation:Required
	OpenstackUsername string `json:"openstackUsername"`

	// +kubebuilder:validation:Required
	ParentRegion string `json:"parentRegion"`

	// +kubebuilder:validation:Required
	PrivateCloudGatewayId string `json:"privateCloudGatewayId"`
}

// CloudaccountOpenstackStatus defines the observed state of CloudaccountOpenstack
type CloudaccountOpenstackStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountOpenstack is the Schema for the cloudaccountopenstacks API
type CloudaccountOpenstack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountOpenstackSpec   `json:"spec,omitempty"`
	Status CloudaccountOpenstackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountOpenstackList contains a list of CloudaccountOpenstack
type CloudaccountOpenstackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountOpenstack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountOpenstack{}, &CloudaccountOpenstackList{})
}
