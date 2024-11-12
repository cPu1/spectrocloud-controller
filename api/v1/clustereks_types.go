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

// ClusterEksSpec defines the desired state of ClusterEks
type ClusterEksSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterEksBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterEksCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterEksClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterEksClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	FargateProfile ClusterEksFargateProfile `json:"fargateProfile,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterEksHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterEksMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterEksNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterEksScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterEksBackupPolicy struct {
	Prefix string `json:"prefix,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterEksCloudConfig struct {
	PrivateAccessCidrs []string `json:"privateAccessCidrs,omitempty"`

	SshKeyName string `json:"sshKeyName,omitempty"`

	Azs []string `json:"azs,omitempty"`

	AzSubnets map[string]string `json:"azSubnets,omitempty"`

	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty"`

	Region string `json:"region,omitempty"`

	VpcId string `json:"vpcId,omitempty"`

	EndpointAccess string `json:"endpointAccess,omitempty"`

	EncryptionConfigArn string `json:"encryptionConfigArn,omitempty"`
}

type ClusterEksClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterEksClusterProfilePack `json:"pack,omitempty"`
}

type ClusterEksClusterProfilePack struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterEksClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`
}

type ClusterEksClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterEksClusterRbacBinding struct {
	Subjects ClusterEksClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`
}

type ClusterEksClusterRbacBindingSubjects struct {
	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterEksFargateProfile struct {
	Name string `json:"name,omitempty"`

	Subnets []string `json:"subnets,omitempty"`

	AdditionalTags map[string]string `json:"additionalTags,omitempty"`

	Selector ClusterEksFargateProfileSelector `json:"selector,omitempty"`
}

type ClusterEksFargateProfileSelector struct {
	Namespace string `json:"namespace,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}

type ClusterEksHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterEksMachinePool struct {
	Min int `json:"min,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	MaxPrice string `json:"maxPrice,omitempty"`

	Name string `json:"name,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Taints ClusterEksMachinePoolTaints `json:"taints,omitempty"`

	CapacityType string `json:"capacityType,omitempty"`

	Azs []string `json:"azs,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	Count int `json:"count,omitempty"`

	Node ClusterEksMachinePoolNode `json:"node,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Max int `json:"max,omitempty"`

	AzSubnets map[string]string `json:"azSubnets,omitempty"`

	EksLaunchTemplate ClusterEksMachinePoolEksLaunchTemplate `json:"eksLaunchTemplate,omitempty"`
}

type ClusterEksMachinePoolEksLaunchTemplate struct {
	AmiId string `json:"amiId,omitempty"`

	RootVolumeType string `json:"rootVolumeType,omitempty"`

	RootVolumeIops int `json:"rootVolumeIops,omitempty"`

	RootVolumeThroughput int `json:"rootVolumeThroughput,omitempty"`

	AdditionalSecurityGroups []string `json:"additionalSecurityGroups,omitempty"`
}

type ClusterEksMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterEksMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterEksNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterEksScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterEksStatus defines the observed state of ClusterEks
type ClusterEksStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterEksLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterEksLocationConfig struct {
	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterEks is the Schema for the clustereks API
type ClusterEks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterEksSpec   `json:"spec,omitempty"`
	Status ClusterEksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEksList contains a list of ClusterEks
type ClusterEksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEks `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterEks{}, &ClusterEksList{})
}
