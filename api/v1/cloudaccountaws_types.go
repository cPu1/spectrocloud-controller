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

// CloudaccountAwsSpec defines the desired state of CloudaccountAws
type CloudaccountAwsSpec struct {
	ARN string `json:"arn,omitempty"`

	AwsAccessKey string `json:"awsAccessKey,omitempty"`

	AwsSecretKey string `json:"awsSecretKey,omitempty"`

	Context string `json:"context,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Partition string `json:"partition,omitempty"`

	PolicyArns []string `json:"policyArns,omitempty"`

	PrivateCloudGatewayId string `json:"privateCloudGatewayId,omitempty"`

	Type string `json:"type,omitempty"`
}

// CloudaccountAwsStatus defines the observed state of CloudaccountAws
type CloudaccountAwsStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudaccountAws is the Schema for the cloudaccountaws API
type CloudaccountAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudaccountAwsSpec   `json:"spec,omitempty"`
	Status CloudaccountAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudaccountAwsList contains a list of CloudaccountAws
type CloudaccountAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudaccountAws `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudaccountAws{}, &CloudaccountAwsList{})
}
