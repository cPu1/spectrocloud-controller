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

// TeamSpec defines the desired state of Team
type TeamSpec struct {

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	ProjectRoleMapping TeamProjectRoleMapping `json:"projectRoleMapping,omitempty"`

	TenantRoleMapping []string `json:"tenantRoleMapping,omitempty"`

	Users []string `json:"users,omitempty"`

	WorkspaceRoleMapping TeamWorkspaceRoleMapping `json:"workspaceRoleMapping,omitempty"`
}

type TeamProjectRoleMapping struct {
	Id string `json:"id,omitempty"`

	Roles []string `json:"roles,omitempty"`
}

type TeamWorkspaceRoleMapping struct {
	Id string `json:"id,omitempty"`

	Workspace TeamWorkspaceRoleMappingWorkspace `json:"workspace,omitempty"`
}

type TeamWorkspaceRoleMappingWorkspace struct {
	Id string `json:"id,omitempty"`

	Roles []string `json:"roles,omitempty"`
}

// TeamStatus defines the observed state of Team
type TeamStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Team is the Schema for the teams API
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamSpec   `json:"spec,omitempty"`
	Status TeamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamList contains a list of Team
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Team{}, &TeamList{})
}
