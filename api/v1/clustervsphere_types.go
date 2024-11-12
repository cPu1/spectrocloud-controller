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

// ClusterVsphereSpec defines the desired state of ClusterVsphere
type ClusterVsphereSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterVsphereBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterVsphereCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterVsphereClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterVsphereClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterVsphereHostConfig `json:"hostConfig,omitempty"`

	LocationConfig ClusterVsphereLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterVsphereMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterVsphereNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterVsphereScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterVsphereBackupPolicy struct {
	BackupLocationId string `json:"backupLocationId,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterVsphereCloudConfig struct {
	HostEndpoint string `json:"hostEndpoint,omitempty"`

	NetworkSearchDomain string `json:"networkSearchDomain,omitempty"`

	Datacenter string `json:"datacenter,omitempty"`

	Folder string `json:"folder,omitempty"`

	ImageTemplateFolder string `json:"imageTemplateFolder,omitempty"`

	SshKey string `json:"sshKey,omitempty"`

	SshKeys []string `json:"sshKeys,omitempty"`

	NetworkType string `json:"networkType,omitempty"`

	NtpServers []string `json:"ntpServers,omitempty"`

	StaticIp bool `json:"staticIp,omitempty"`
}

type ClusterVsphereClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterVsphereClusterProfilePack `json:"pack,omitempty"`
}

type ClusterVsphereClusterProfilePack struct {
	Values string `json:"values,omitempty"`

	Manifest ClusterVsphereClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`
}

type ClusterVsphereClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterVsphereClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterVsphereClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterVsphereClusterRbacBindingSubjects struct {
	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterVsphereHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterVsphereLocationConfig struct {
	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`
}

type ClusterVsphereMachinePool struct {
	UpdateStrategy string `json:"updateStrategy,omitempty"`

	InstanceType ClusterVsphereMachinePoolInstanceType `json:"instanceType,omitempty"`

	Placement ClusterVsphereMachinePoolPlacement `json:"placement,omitempty"`

	Node ClusterVsphereMachinePoolNode `json:"node,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	Count int `json:"count,omitempty"`

	Min int `json:"min,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	Max int `json:"max,omitempty"`

	Name string `json:"name,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Taints ClusterVsphereMachinePoolTaints `json:"taints,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`
}

type ClusterVsphereMachinePoolInstanceType struct {
	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	MemoryMb int `json:"memoryMb,omitempty"`

	Cpu int `json:"cpu,omitempty"`
}

type ClusterVsphereMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterVsphereMachinePoolPlacement struct {
	Id string `json:"id,omitempty"`

	Cluster string `json:"cluster,omitempty"`

	ResourcePool string `json:"resourcePool,omitempty"`

	Datastore string `json:"datastore,omitempty"`

	Network string `json:"network,omitempty"`

	StaticIpPoolId string `json:"staticIpPoolId,omitempty"`
}

type ClusterVsphereMachinePoolTaints struct {
	Effect string `json:"effect,omitempty"`

	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type ClusterVsphereNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterVsphereScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterVsphereStatus defines the observed state of ClusterVsphere
type ClusterVsphereStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterVsphere is the Schema for the clustervspheres API
type ClusterVsphere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterVsphereSpec   `json:"spec,omitempty"`
	Status ClusterVsphereStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterVsphereList contains a list of ClusterVsphere
type ClusterVsphereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVsphere `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterVsphere{}, &ClusterVsphereList{})
}
