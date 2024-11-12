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

// ClusterProfileImportSpec defines the desired state of ClusterProfileImport
type ClusterProfileImportSpec struct {
	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	ImportFile string `json:"importFile"`
}

// ClusterProfileImportStatus defines the observed state of ClusterProfileImport
type ClusterProfileImportStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterProfileImport is the Schema for the clusterprofileimports API
type ClusterProfileImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterProfileImportSpec   `json:"spec,omitempty"`
	Status ClusterProfileImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterProfileImportList contains a list of ClusterProfileImport
type ClusterProfileImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProfileImport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterProfileImport{}, &ClusterProfileImportList{})
}
