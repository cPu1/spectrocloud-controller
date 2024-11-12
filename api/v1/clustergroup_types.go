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

// ClusterGroupSpec defines the desired state of ClusterGroup
type ClusterGroupSpec struct {
	ClusterProfile ClusterGroupClusterProfile `json:"clusterProfile,omitempty"`

	Clusters ClusterGroupClusters `json:"clusters,omitempty"`

	// +kubebuilder:validation:Required
	Config ClusterGroupConfig `json:"config"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterGroupClusterProfile struct {
	Pack ClusterGroupClusterProfilePack `json:"pack,omitempty"`

	Id string `json:"id,omitempty"`
}

type ClusterGroupClusterProfilePack struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterGroupClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`
}

type ClusterGroupClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterGroupClusters struct {
	ClusterUid string `json:"clusterUid,omitempty"`

	HostDns string `json:"hostDns,omitempty"`
}

type ClusterGroupConfig struct {
	MemoryInMb int `json:"memoryInMb,omitempty"`

	StorageInGb int `json:"storageInGb,omitempty"`

	OversubscriptionPercent int `json:"oversubscriptionPercent,omitempty"`

	Values string `json:"values,omitempty"`

	K8sDistribution string `json:"k8sDistribution,omitempty"`

	HostEndpointType string `json:"hostEndpointType,omitempty"`

	CpuMillicore int `json:"cpuMillicore,omitempty"`
}

// ClusterGroupStatus defines the observed state of ClusterGroup
type ClusterGroupStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterGroup is the Schema for the clustergroups API
type ClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterGroupSpec   `json:"spec,omitempty"`
	Status ClusterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGroupList contains a list of ClusterGroup
type ClusterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterGroup{}, &ClusterGroupList{})
}
