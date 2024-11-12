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

// ClusterAksSpec defines the desired state of ClusterAks
type ClusterAksSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterAksBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterAksCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterAksClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterAksClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterAksHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterAksMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterAksNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterAksScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterAksBackupPolicy struct {
	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	Prefix string `json:"prefix,omitempty"`
}

type ClusterAksCloudConfig struct {
	ControlPlaneSubnetName string `json:"controlPlaneSubnetName,omitempty"`

	ControlPlaneCidr string `json:"controlPlaneCidr,omitempty"`

	ControlPlaneSubnetSecurityGroupName string `json:"controlPlaneSubnetSecurityGroupName,omitempty"`

	VnetName string `json:"vnetName,omitempty"`

	VnetCidrBlock string `json:"vnetCidrBlock,omitempty"`

	WorkerSubnetName string `json:"workerSubnetName,omitempty"`

	Region string `json:"region,omitempty"`

	ResourceGroup string `json:"resourceGroup,omitempty"`

	SshKey string `json:"sshKey,omitempty"`

	VnetResourceGroup string `json:"vnetResourceGroup,omitempty"`

	WorkerSubnetSecurityGroupName string `json:"workerSubnetSecurityGroupName,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	WorkerCidr string `json:"workerCidr,omitempty"`

	PrivateCluster bool `json:"privateCluster,omitempty"`
}

type ClusterAksClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterAksClusterProfilePack `json:"pack,omitempty"`
}

type ClusterAksClusterProfilePack struct {
	Manifest ClusterAksClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`
}

type ClusterAksClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterAksClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterAksClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterAksClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterAksHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterAksMachinePool struct {
	InstanceType string `json:"instanceType,omitempty"`

	Max int `json:"max,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	Name string `json:"name,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Node ClusterAksMachinePoolNode `json:"node,omitempty"`

	Count int `json:"count,omitempty"`

	StorageAccountType string `json:"storageAccountType,omitempty"`

	Taints ClusterAksMachinePoolTaints `json:"taints,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Min int `json:"min,omitempty"`

	IsSystemNodePool bool `json:"isSystemNodePool,omitempty"`
}

type ClusterAksMachinePoolNode struct {
	Action string `json:"action,omitempty"`

	NodeId string `json:"nodeId,omitempty"`
}

type ClusterAksMachinePoolTaints struct {
	Effect string `json:"effect,omitempty"`

	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type ClusterAksNamespaces struct {
	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`

	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`
}

type ClusterAksScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterAksStatus defines the observed state of ClusterAks
type ClusterAksStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterAksLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterAksLocationConfig struct {
	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterAks is the Schema for the clusteraks API
type ClusterAks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAksSpec   `json:"spec,omitempty"`
	Status ClusterAksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterAksList contains a list of ClusterAks
type ClusterAksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAks `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterAks{}, &ClusterAksList{})
}
