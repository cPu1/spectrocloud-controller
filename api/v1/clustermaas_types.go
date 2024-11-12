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

// ClusterMaasSpec defines the desired state of ClusterMaas
type ClusterMaasSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterMaasBackupPolicy `json:"backupPolicy,omitempty"`

	CloudAccountId string `json:"cloudAccountId,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterMaasCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterMaasClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterMaasClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterMaasHostConfig `json:"hostConfig,omitempty"`

	LocationConfig ClusterMaasLocationConfig `json:"locationConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterMaasMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterMaasNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterMaasScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterMaasBackupPolicy struct {
	BackupLocationId string `json:"backupLocationId,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`
}

type ClusterMaasCloudConfig struct {
	Domain string `json:"domain,omitempty"`
}

type ClusterMaasClusterProfile struct {
	Id string `json:"id,omitempty"`

	Pack ClusterMaasClusterProfilePack `json:"pack,omitempty"`
}

type ClusterMaasClusterProfilePack struct {
	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`

	Manifest ClusterMaasClusterProfilePackManifest `json:"manifest,omitempty"`
}

type ClusterMaasClusterProfilePackManifest struct {
	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`
}

type ClusterMaasClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterMaasClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterMaasClusterRbacBindingSubjects struct {
	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

type ClusterMaasHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterMaasLocationConfig struct {
	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`
}

type ClusterMaasMachinePool struct {
	InstanceType ClusterMaasMachinePoolInstanceType `json:"instanceType,omitempty"`

	Node ClusterMaasMachinePoolNode `json:"node,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	Name string `json:"name,omitempty"`

	Count int `json:"count,omitempty"`

	Max int `json:"max,omitempty"`

	Azs []string `json:"azs,omitempty"`

	Placement ClusterMaasMachinePoolPlacement `json:"placement,omitempty"`

	Taints ClusterMaasMachinePoolTaints `json:"taints,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Min int `json:"min,omitempty"`

	NodeTags []string `json:"nodeTags,omitempty"`
}

type ClusterMaasMachinePoolInstanceType struct {
	MinMemoryMb int `json:"minMemoryMb,omitempty"`

	MinCpu int `json:"minCpu,omitempty"`
}

type ClusterMaasMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterMaasMachinePoolPlacement struct {
	Id string `json:"id,omitempty"`

	ResourcePool string `json:"resourcePool,omitempty"`
}

type ClusterMaasMachinePoolTaints struct {
	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`

	Key string `json:"key,omitempty"`
}

type ClusterMaasNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterMaasScanPolicy struct {
	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`

	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`
}

// ClusterMaasStatus defines the observed state of ClusterMaas
type ClusterMaasStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterMaas is the Schema for the clustermaas API
type ClusterMaas struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterMaasSpec   `json:"spec,omitempty"`
	Status ClusterMaasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterMaasList contains a list of ClusterMaas
type ClusterMaasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMaas `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterMaas{}, &ClusterMaasList{})
}
