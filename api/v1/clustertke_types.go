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

// ClusterTkeSpec defines the desired state of ClusterTke
type ClusterTkeSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterTkeBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterTkeCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterTkeClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterTkeClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterTkeHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterTkeMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterTkeNamespaces `json:"namespaces,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterTkeScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterTkeBackupPolicy struct {
	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	Prefix string `json:"prefix,omitempty"`
}

type ClusterTkeCloudConfig struct {
	EndpointAccess string `json:"endpointAccess,omitempty"`

	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty"`

	SshKeyName string `json:"sshKeyName,omitempty"`

	Region string `json:"region,omitempty"`

	VpcId string `json:"vpcId,omitempty"`

	Azs []string `json:"azs,omitempty"`

	AzSubnets map[string]string `json:"azSubnets,omitempty"`
}

type ClusterTkeClusterProfile struct {
	Pack ClusterTkeClusterProfilePack `json:"pack,omitempty"`

	Id string `json:"id,omitempty"`
}

type ClusterTkeClusterProfilePack struct {
	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterTkeClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`
}

type ClusterTkeClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterTkeClusterRbacBinding struct {
	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterTkeClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterTkeClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterTkeHostConfig struct {
	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`

	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`
}

type ClusterTkeMachinePool struct {
	Max int `json:"max,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	CapacityType string `json:"capacityType,omitempty"`

	Name string `json:"name,omitempty"`

	Taints ClusterTkeMachinePoolTaints `json:"taints,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	Azs []string `json:"azs,omitempty"`

	Count int `json:"count,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Min int `json:"min,omitempty"`

	MaxPrice string `json:"maxPrice,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Node ClusterTkeMachinePoolNode `json:"node,omitempty"`

	AzSubnets map[string]string `json:"azSubnets,omitempty"`
}

type ClusterTkeMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterTkeMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterTkeNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterTkeScanPolicy struct {
	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`

	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`
}

// ClusterTkeStatus defines the observed state of ClusterTke
type ClusterTkeStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterTkeLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterTkeLocationConfig struct {
	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterTke is the Schema for the clustertkes API
type ClusterTke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterTkeSpec   `json:"spec,omitempty"`
	Status ClusterTkeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterTkeList contains a list of ClusterTke
type ClusterTkeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterTke `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterTke{}, &ClusterTkeList{})
}
