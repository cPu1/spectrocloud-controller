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

// ClusterImportSpec defines the desired state of ClusterImport
type ClusterImportSpec struct {

	// +kubebuilder:validation:Required
	Cloud string `json:"cloud"`

	ClusterProfile ClusterImportClusterProfile `json:"clusterProfile,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterImportClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterImportClusterProfilePack `json:"pack,omitempty"`
}

type ClusterImportClusterProfilePack struct {
	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterImportClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`
}

type ClusterImportClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

// ClusterImportStatus defines the observed state of ClusterImport
type ClusterImportStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	ClusterImportManifest string `json:"clusterImportManifest,omitempty"`

	ClusterImportManifestApplyCommand string `json:"clusterImportManifestApplyCommand,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterImport is the Schema for the clusterimports API
type ClusterImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterImportSpec   `json:"spec,omitempty"`
	Status ClusterImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterImportList contains a list of ClusterImport
type ClusterImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterImport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterImport{}, &ClusterImportList{})
}
