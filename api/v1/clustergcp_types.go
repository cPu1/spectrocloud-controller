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

// ClusterGcpSpec defines the desired state of ClusterGcp
type ClusterGcpSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterGcpBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterGcpCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterGcpClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterGcpClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterGcpHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterGcpMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterGcpNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterGcpScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterGcpBackupPolicy struct {
	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`
}

type ClusterGcpCloudConfig struct {
	Network string `json:"network,omitempty"`

	Project string `json:"project,omitempty"`

	Region string `json:"region,omitempty"`
}

type ClusterGcpClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterGcpClusterProfilePack `json:"pack,omitempty"`
}

type ClusterGcpClusterProfilePack struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterGcpClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`
}

type ClusterGcpClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterGcpClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterGcpClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterGcpClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterGcpHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterGcpMachinePool struct {
	Name string `json:"name,omitempty"`

	Count int `json:"count,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	Azs []string `json:"azs,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	Taints ClusterGcpMachinePoolTaints `json:"taints,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Node ClusterGcpMachinePoolNode `json:"node,omitempty"`
}

type ClusterGcpMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterGcpMachinePoolTaints struct {
	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`

	Key string `json:"key,omitempty"`
}

type ClusterGcpNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterGcpScanPolicy struct {
	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`

	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`
}

// ClusterGcpStatus defines the observed state of ClusterGcp
type ClusterGcpStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterGcpLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterGcpLocationConfig struct {
	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterGcp is the Schema for the clustergcps API
type ClusterGcp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterGcpSpec   `json:"spec,omitempty"`
	Status ClusterGcpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterGcpList contains a list of ClusterGcp
type ClusterGcpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterGcp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterGcp{}, &ClusterGcpList{})
}
