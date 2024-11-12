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

// ClusterOpenstackSpec defines the desired state of ClusterOpenstack
type ClusterOpenstackSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterOpenstackBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterOpenstackCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterOpenstackClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterOpenstackClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterOpenstackHostConfig `json:"hostConfig,omitempty"`

	LocationConfig ClusterOpenstackLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterOpenstackMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterOpenstackNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterOpenstackScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterOpenstackBackupPolicy struct {
	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterOpenstackCloudConfig struct {
	DnsServers []string `json:"dnsServers,omitempty"`

	SubnetCidr string `json:"subnetCidr,omitempty"`

	Domain string `json:"domain,omitempty"`

	Region string `json:"region,omitempty"`

	Project string `json:"project,omitempty"`

	SshKey string `json:"sshKey,omitempty"`

	NetworkId string `json:"networkId,omitempty"`

	SubnetId string `json:"subnetId,omitempty"`
}

type ClusterOpenstackClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterOpenstackClusterProfilePack `json:"pack,omitempty"`
}

type ClusterOpenstackClusterProfilePack struct {
	Manifest ClusterOpenstackClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`
}

type ClusterOpenstackClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterOpenstackClusterRbacBinding struct {
	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterOpenstackClusterRbacBindingSubjects `json:"subjects,omitempty"`

	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterOpenstackClusterRbacBindingSubjects struct {
	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type ClusterOpenstackHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterOpenstackLocationConfig struct {
	CountryCode string `json:"countryCode,omitempty"`

	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`
}

type ClusterOpenstackMachinePool struct {
	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Node ClusterOpenstackMachinePoolNode `json:"node,omitempty"`

	Count int `json:"count,omitempty"`

	Azs []string `json:"azs,omitempty"`

	SubnetId string `json:"subnetId,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	Taints ClusterOpenstackMachinePoolTaints `json:"taints,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	Name string `json:"name,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`
}

type ClusterOpenstackMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterOpenstackMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterOpenstackNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterOpenstackScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterOpenstackStatus defines the observed state of ClusterOpenstack
type ClusterOpenstackStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterOpenstack is the Schema for the clusteropenstacks API
type ClusterOpenstack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterOpenstackSpec   `json:"spec,omitempty"`
	Status ClusterOpenstackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterOpenstackList contains a list of ClusterOpenstack
type ClusterOpenstackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterOpenstack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterOpenstack{}, &ClusterOpenstackList{})
}
