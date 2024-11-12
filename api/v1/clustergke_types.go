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

// ClusterGkeSpec defines the desired state of ClusterGke
type ClusterGkeSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterGkeBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterGkeCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterGkeClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterGkeClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterGkeHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterGkeMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterGkeNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterGkeScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`

	UpdateWorkerPoolInParallel bool `json:"updateWorkerPoolInParallel,omitempty"`
}

type ClusterGkeBackupPolicy struct {
	ExpiryInHour int `json:"expiryInHour,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`
}

type ClusterGkeCloudConfig struct {
	Region string `json:"region,omitempty"`

	Project string `json:"project,omitempty"`
}

type ClusterGkeClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterGkeClusterProfilePack `json:"pack,omitempty"`
}

type ClusterGkeClusterProfilePack struct {
	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterGkeClusterProfilePackManifest `json:"manifest,omitempty"`
}

type ClusterGkeClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterGkeClusterRbacBinding struct {
	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterGkeClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterGkeClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterGkeHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterGkeMachinePool struct {
	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Node ClusterGkeMachinePoolNode `json:"node,omitempty"`

	Taints ClusterGkeMachinePoolTaints `json:"taints,omitempty"`

	Name string `json:"name,omitempty"`

	Count int `json:"count,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`
}

type ClusterGkeMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterGkeMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterGkeNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterGkeScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterGkeStatus defines the observed state of ClusterGke
type ClusterGkeStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterGkeLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterGkeLocationConfig struct {
	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterGke is the Schema for the clustergkes API
type ClusterGke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterGkeSpec   `json:"spec,omitempty"`
	Status ClusterGkeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGkeList contains a list of ClusterGke
type ClusterGkeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGke `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterGke{}, &ClusterGkeList{})
}
