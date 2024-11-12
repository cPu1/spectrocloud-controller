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

// ApplicationProfileSpec defines the desired state of ApplicationProfile
type ApplicationProfileSpec struct {
	Cloud string `json:"cloud,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	Pack ApplicationProfilePack `json:"pack"`

	Tags []string `json:"tags,omitempty"`

	Version string `json:"version,omitempty"`
}

type ApplicationProfilePack struct {
	SourceAppTier string `json:"sourceAppTier,omitempty"`

	Name string `json:"name,omitempty"`

	Properties map[string]string `json:"properties,omitempty"`

	Manifest ApplicationProfilePackManifest `json:"manifest,omitempty"`

	Tag string `json:"tag,omitempty"`

	Type string `json:"type,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Uid string `json:"uid,omitempty"`

	InstallOrder int `json:"installOrder,omitempty"`

	Values string `json:"values,omitempty"`
}

type ApplicationProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

// ApplicationProfileStatus defines the observed state of ApplicationProfile
type ApplicationProfileStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ApplicationProfile is the Schema for the applicationprofiles API
type ApplicationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationProfileSpec   `json:"spec,omitempty"`
	Status ApplicationProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationProfileList contains a list of ApplicationProfile
type ApplicationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationProfile{}, &ApplicationProfileList{})
}
