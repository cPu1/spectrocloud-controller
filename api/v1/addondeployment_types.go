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

// AddonDeploymentSpec defines the desired state of AddonDeployment
type AddonDeploymentSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	ClusterProfile AddonDeploymentClusterProfile `json:"clusterProfile,omitempty"`

	// +kubebuilder:validation:Required
	ClusterUid string `json:"clusterUid"`

	Context string `json:"context,omitempty"`
}

type AddonDeploymentClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack AddonDeploymentClusterProfilePack `json:"pack,omitempty"`
}

type AddonDeploymentClusterProfilePack struct {
	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest AddonDeploymentClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

type AddonDeploymentClusterProfilePackManifest struct {
	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`

	Uid string `json:"uid,omitempty"`
}

// AddonDeploymentStatus defines the observed state of AddonDeployment
type AddonDeploymentStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AddonDeployment is the Schema for the addondeployments API
type AddonDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddonDeploymentSpec   `json:"spec,omitempty"`
	Status AddonDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonDeploymentList contains a list of AddonDeployment
type AddonDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AddonDeployment{}, &AddonDeploymentList{})
}
