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

// ClusterEdgeNativeSpec defines the desired state of ClusterEdgeNative
type ClusterEdgeNativeSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterEdgeNativeBackupPolicy `json:"backupPolicy,omitempty"`

	CloudAccountId string `json:"cloudAccountId,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterEdgeNativeCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterEdgeNativeClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterEdgeNativeClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterEdgeNativeHostConfig `json:"hostConfig,omitempty"`

	LocationConfig ClusterEdgeNativeLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterEdgeNativeMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterEdgeNativeNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterEdgeNativeScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterEdgeNativeBackupPolicy struct {
	BackupLocationId string `json:"backupLocationId,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterEdgeNativeCloudConfig struct {
	SshKeys []string `json:"sshKeys,omitempty"`

	Vip string `json:"vip,omitempty"`

	OverlayCidrRange string `json:"overlayCidrRange,omitempty"`

	NtpServers []string `json:"ntpServers,omitempty"`
}

type ClusterEdgeNativeClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterEdgeNativeClusterProfilePack `json:"pack,omitempty"`
}

type ClusterEdgeNativeClusterProfilePack struct {
	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterEdgeNativeClusterProfilePackManifest `json:"manifest,omitempty"`
}

type ClusterEdgeNativeClusterProfilePackManifest struct {
	Content string `json:"content,omitempty"`

	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`
}

type ClusterEdgeNativeClusterRbacBinding struct {
	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterEdgeNativeClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterEdgeNativeClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterEdgeNativeHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterEdgeNativeLocationConfig struct {
	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`
}

type ClusterEdgeNativeMachinePool struct {
	Node ClusterEdgeNativeMachinePoolNode `json:"node,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	EdgeHost ClusterEdgeNativeMachinePoolEdgeHost `json:"edgeHost,omitempty"`

	Name string `json:"name,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Taints ClusterEdgeNativeMachinePoolTaints `json:"taints,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`
}

type ClusterEdgeNativeMachinePoolEdgeHost struct {
	DefaultGateway string `json:"defaultGateway,omitempty"`

	SubnetMask string `json:"subnetMask,omitempty"`

	DnsServers []string `json:"dnsServers,omitempty"`

	TwoNodeRole string `json:"twoNodeRole,omitempty"`

	HostName string `json:"hostName,omitempty"`

	HostUid string `json:"hostUid,omitempty"`

	StaticIp string `json:"staticIp,omitempty"`

	NicName string `json:"nicName,omitempty"`
}

type ClusterEdgeNativeMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterEdgeNativeMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterEdgeNativeNamespaces struct {
	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`

	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`
}

type ClusterEdgeNativeScanPolicy struct {
	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`

	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`
}

// ClusterEdgeNativeStatus defines the observed state of ClusterEdgeNative
type ClusterEdgeNativeStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterEdgeNative is the Schema for the clusteredgenatives API
type ClusterEdgeNative struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterEdgeNativeSpec   `json:"spec,omitempty"`
	Status ClusterEdgeNativeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterEdgeNativeList contains a list of ClusterEdgeNative
type ClusterEdgeNativeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEdgeNative `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterEdgeNative{}, &ClusterEdgeNativeList{})
}
