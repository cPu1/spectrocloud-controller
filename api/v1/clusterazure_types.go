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

// ClusterAzureSpec defines the desired state of ClusterAzure
type ClusterAzureSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterAzureBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterAzureCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterAzureClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterAzureClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterAzureHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterAzureMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterAzureNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterAzureScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterAzureBackupPolicy struct {
	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterAzureCloudConfig struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"`

	VirtualNetworkCidrBlock string `json:"virtualNetworkCidrBlock,omitempty"`

	WorkerNodeSubnet ClusterAzureCloudConfigWorkerNodeSubnet `json:"workerNodeSubnet,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty"`

	Region string `json:"region,omitempty"`

	SshKey string `json:"sshKey,omitempty"`

	StorageAccountName string `json:"storageAccountName,omitempty"`

	ResourceGroup string `json:"resourceGroup,omitempty"`

	ContainerName string `json:"containerName,omitempty"`

	NetworkResourceGroup string `json:"networkResourceGroup,omitempty"`

	ControlPlaneSubnet ClusterAzureCloudConfigControlPlaneSubnet `json:"controlPlaneSubnet,omitempty"`
}

type ClusterAzureCloudConfigControlPlaneSubnet struct {
	SecurityGroupName string `json:"securityGroupName,omitempty"`

	Name string `json:"name,omitempty"`

	CidrBlock string `json:"cidrBlock,omitempty"`
}

type ClusterAzureCloudConfigWorkerNodeSubnet struct {
	Name string `json:"name,omitempty"`

	CidrBlock string `json:"cidrBlock,omitempty"`

	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type ClusterAzureClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterAzureClusterProfilePack `json:"pack,omitempty"`
}

type ClusterAzureClusterProfilePack struct {
	Values string `json:"values,omitempty"`

	Manifest ClusterAzureClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`
}

type ClusterAzureClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterAzureClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterAzureClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterAzureClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterAzureHostConfig struct {
	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`

	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`
}

type ClusterAzureMachinePool struct {
	Taints ClusterAzureMachinePoolTaints `json:"taints,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	Azs []string `json:"azs,omitempty"`

	IsSystemNodePool bool `json:"isSystemNodePool,omitempty"`

	Node ClusterAzureMachinePoolNode `json:"node,omitempty"`

	Name string `json:"name,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	OsType string `json:"osType,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Count int `json:"count,omitempty"`

	Disk ClusterAzureMachinePoolDisk `json:"disk,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`
}

type ClusterAzureMachinePoolDisk struct {
	SizeGb int `json:"sizeGb,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterAzureMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterAzureMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterAzureNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterAzureScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterAzureStatus defines the observed state of ClusterAzure
type ClusterAzureStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterAzureLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterAzureLocationConfig struct {
	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterAzure is the Schema for the clusterazures API
type ClusterAzure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAzureSpec   `json:"spec,omitempty"`
	Status ClusterAzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterAzureList contains a list of ClusterAzure
type ClusterAzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAzure `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterAzure{}, &ClusterAzureList{})
}
