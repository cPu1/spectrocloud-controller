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

// ClusterProfileSpec defines the desired state of ClusterProfile
type ClusterProfileSpec struct {
	Cloud string `json:"cloud,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Pack ClusterProfilePack `json:"pack,omitempty"`

	ProfileVariables ClusterProfileProfileVariables `json:"profileVariables,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Type string `json:"type,omitempty"`

	Version string `json:"version,omitempty"`
}

type ClusterProfilePack struct {
	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterProfileProfileVariables struct {
	Variable ClusterProfileProfileVariablesVariable `json:"variable,omitempty"`
}

type ClusterProfileProfileVariablesVariable struct {
	Regex string `json:"regex,omitempty"`

	Required bool `json:"required,omitempty"`

	Immutable bool `json:"immutable,omitempty"`

	IsSensitive bool `json:"isSensitive,omitempty"`

	Hidden bool `json:"hidden,omitempty"`

	Name string `json:"name,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	Format string `json:"format,omitempty"`

	Description string `json:"description,omitempty"`

	DefaultValue string `json:"defaultValue,omitempty"`
}

// ClusterProfileStatus defines the observed state of ClusterProfile
type ClusterProfileStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterProfile is the Schema for the clusterprofiles API
type ClusterProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterProfileSpec   `json:"spec,omitempty"`
	Status ClusterProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterProfileList contains a list of ClusterProfile
type ClusterProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterProfile{}, &ClusterProfileList{})
}
