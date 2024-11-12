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

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	BackupPolicy WorkspaceBackupPolicy `json:"backupPolicy,omitempty"`

	ClusterRbacBinding WorkspaceClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	// +kubebuilder:validation:Required
	Clusters WorkspaceClusters `json:"clusters"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces WorkspaceNamespaces `json:"namespaces,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type WorkspaceBackupPolicy struct {
	BackupLocationId string `json:"backupLocationId,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type WorkspaceClusterRbacBinding struct {
	Subjects WorkspaceClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`
}

type WorkspaceClusterRbacBindingSubjects struct {
	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

type WorkspaceClusters struct {
	Uid string `json:"uid,omitempty"`
}

type WorkspaceNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Workspace is the Schema for the workspaces API
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkspaceSpec   `json:"spec,omitempty"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
