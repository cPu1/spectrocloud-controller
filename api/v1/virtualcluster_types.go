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

// VirtualClusterSpec defines the desired state of VirtualCluster
type VirtualClusterSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy VirtualClusterBackupPolicy `json:"backupPolicy,omitempty"`

	CloudConfig VirtualClusterCloudConfig `json:"cloudConfig,omitempty"`

	ClusterGroupUid string `json:"clusterGroupUid,omitempty"`

	ClusterProfile VirtualClusterClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding VirtualClusterClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostClusterUid string `json:"hostClusterUid,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces VirtualClusterNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseCluster bool `json:"pauseCluster,omitempty"`

	Resources VirtualClusterResources `json:"resources,omitempty"`

	ScanPolicy VirtualClusterScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type VirtualClusterBackupPolicy struct {
	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type VirtualClusterCloudConfig struct {
	ChartName string `json:"chartName,omitempty"`

	ChartRepo string `json:"chartRepo,omitempty"`

	ChartVersion string `json:"chartVersion,omitempty"`

	ChartValues string `json:"chartValues,omitempty"`

	K8sVersion string `json:"k8sVersion,omitempty"`
}

type VirtualClusterClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack VirtualClusterClusterProfilePack `json:"pack,omitempty"`
}

type VirtualClusterClusterProfilePack struct {
	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest VirtualClusterClusterProfilePackManifest `json:"manifest,omitempty"`
}

type VirtualClusterClusterProfilePackManifest struct {
	Content string `json:"content,omitempty"`

	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`
}

type VirtualClusterClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects VirtualClusterClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type VirtualClusterClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type VirtualClusterNamespaces struct {
	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`

	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`
}

type VirtualClusterResources struct {
	MaxStorageInGb int `json:"maxStorageInGb,omitempty"`

	MinCpu int `json:"minCpu,omitempty"`

	MinMemInMb int `json:"minMemInMb,omitempty"`

	MinStorageInGb int `json:"minStorageInGb,omitempty"`

	MaxCpu int `json:"maxCpu,omitempty"`

	MaxMemInMb int `json:"maxMemInMb,omitempty"`
}

type VirtualClusterScanPolicy struct {
	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`

	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`
}

// VirtualClusterStatus defines the observed state of VirtualCluster
type VirtualClusterStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig VirtualClusterLocationConfig `json:"locationConfig,omitempty"`
}

type VirtualClusterLocationConfig struct {
	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualCluster is the Schema for the virtualclusters API
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpec   `json:"spec,omitempty"`
	Status VirtualClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualClusterList contains a list of VirtualCluster
type VirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualCluster{}, &VirtualClusterList{})
}
