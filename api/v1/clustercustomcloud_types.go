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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterCustomCloudSpec defines the desired state of ClusterCustomCloud
type ClusterCustomCloudSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterCustomCloudBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	Cloud string `json:"cloud"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterCustomCloudCloudConfig `json:"cloudConfig"`

	ClusterProfile ClusterCustomCloudClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterCustomCloudClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	LocationConfig ClusterCustomCloudLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterCustomCloudMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterCustomCloudNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ScanPolicy ClusterCustomCloudScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterCustomCloudBackupPolicy struct {
	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`
}

type ClusterCustomCloudCloudConfig struct {
	Values string `json:"values,omitempty"`
}

type ClusterCustomCloudClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterCustomCloudClusterProfilePack `json:"pack,omitempty"`
}

type ClusterCustomCloudClusterProfilePack struct {
	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterCustomCloudClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

type ClusterCustomCloudClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterCustomCloudClusterRbacBinding struct {
	Subjects ClusterCustomCloudClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`
}

type ClusterCustomCloudClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterCustomCloudLocationConfig struct {
	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`
}

type ClusterCustomCloudMachinePool struct {
	Name string `json:"name,omitempty"`

	Count int `json:"count,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	NodePoolConfig string `json:"nodePoolConfig,omitempty"`
}

type ClusterCustomCloudNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterCustomCloudScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterCustomCloudStatus defines the observed state of ClusterCustomCloud
type ClusterCustomCloudStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterCustomCloud is the Schema for the clustercustomclouds API
type ClusterCustomCloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterCustomCloudSpec   `json:"spec,omitempty"`
	Status ClusterCustomCloudStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterCustomCloudList contains a list of ClusterCustomCloud
type ClusterCustomCloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCustomCloud `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterCustomCloud{}, &ClusterCustomCloudList{})
}
