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

// ClusterEdgeVsphereSpec defines the desired state of ClusterEdgeVsphere
type ClusterEdgeVsphereSpec struct {
	BackupPolicy ClusterEdgeVsphereBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterEdgeVsphereCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterEdgeVsphereClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterEdgeVsphereClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	EdgeHostUid string `json:"edgeHostUid"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterEdgeVsphereHostConfig `json:"hostConfig,omitempty"`

	LocationConfig ClusterEdgeVsphereLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterEdgeVsphereMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterEdgeVsphereNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterEdgeVsphereScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterEdgeVsphereBackupPolicy struct {
	Prefix string `json:"prefix,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`
}

type ClusterEdgeVsphereCloudConfig struct {
	Folder string `json:"folder,omitempty"`

	ImageTemplateFolder string `json:"imageTemplateFolder,omitempty"`

	SshKey string `json:"sshKey,omitempty"`

	Vip string `json:"vip,omitempty"`

	StaticIp bool `json:"staticIp,omitempty"`

	Datacenter string `json:"datacenter,omitempty"`

	SshKeys []string `json:"sshKeys,omitempty"`

	NetworkType string `json:"networkType,omitempty"`

	NetworkSearchDomain string `json:"networkSearchDomain,omitempty"`
}

type ClusterEdgeVsphereClusterProfile struct {
	Pack ClusterEdgeVsphereClusterProfilePack `json:"pack,omitempty"`

	Id string `json:"id,omitempty"`
}

type ClusterEdgeVsphereClusterProfilePack struct {
	Values string `json:"values,omitempty"`

	Manifest ClusterEdgeVsphereClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`
}

type ClusterEdgeVsphereClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterEdgeVsphereClusterRbacBinding struct {
	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterEdgeVsphereClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterEdgeVsphereClusterRbacBindingSubjects struct {
	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`
}

type ClusterEdgeVsphereHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterEdgeVsphereLocationConfig struct {
	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`
}

type ClusterEdgeVsphereMachinePool struct {
	InstanceType ClusterEdgeVsphereMachinePoolInstanceType `json:"instanceType,omitempty"`

	Name string `json:"name,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Node ClusterEdgeVsphereMachinePoolNode `json:"node,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	Count int `json:"count,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Taints ClusterEdgeVsphereMachinePoolTaints `json:"taints,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	Placement ClusterEdgeVsphereMachinePoolPlacement `json:"placement,omitempty"`
}

type ClusterEdgeVsphereMachinePoolInstanceType struct {
	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	MemoryMb int `json:"memoryMb,omitempty"`

	Cpu int `json:"cpu,omitempty"`
}

type ClusterEdgeVsphereMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterEdgeVsphereMachinePoolPlacement struct {
	Id string `json:"id,omitempty"`

	Cluster string `json:"cluster,omitempty"`

	ResourcePool string `json:"resourcePool,omitempty"`

	Datastore string `json:"datastore,omitempty"`

	Network string `json:"network,omitempty"`

	StaticIpPoolId string `json:"staticIpPoolId,omitempty"`
}

type ClusterEdgeVsphereMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterEdgeVsphereNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterEdgeVsphereScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterEdgeVsphereStatus defines the observed state of ClusterEdgeVsphere
type ClusterEdgeVsphereStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterEdgeVsphere is the Schema for the clusteredgevspheres API
type ClusterEdgeVsphere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterEdgeVsphereSpec   `json:"spec,omitempty"`
	Status ClusterEdgeVsphereStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEdgeVsphereList contains a list of ClusterEdgeVsphere
type ClusterEdgeVsphereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEdgeVsphere `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterEdgeVsphere{}, &ClusterEdgeVsphereList{})
}
