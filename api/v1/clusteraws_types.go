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

// ClusterAwsSpec defines the desired state of ClusterAws
type ClusterAwsSpec struct {
	ApplySetting string `json:"applySetting,omitempty"`

	BackupPolicy ClusterAwsBackupPolicy `json:"backupPolicy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountId string `json:"cloudAccountId"`

	// +kubebuilder:validation:Required
	CloudConfig ClusterAwsCloudConfig `json:"cloudConfig"`

	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	ClusterProfile ClusterAwsClusterProfile `json:"clusterProfile,omitempty"`

	ClusterRbacBinding ClusterAwsClusterRbacBinding `json:"clusterRbacBinding,omitempty"`

	Context string `json:"context,omitempty"`

	Description string `json:"description,omitempty"`

	ForceDelete bool `json:"forceDelete,omitempty"`

	ForceDeleteDelay int `json:"forceDeleteDelay,omitempty"`

	HostConfig ClusterAwsHostConfig `json:"hostConfig,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool ClusterAwsMachinePool `json:"machinePool"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	Namespaces ClusterAwsNamespaces `json:"namespaces,omitempty"`

	OsPatchAfter string `json:"osPatchAfter,omitempty"`

	OsPatchOnBoot bool `json:"osPatchOnBoot,omitempty"`

	OsPatchSchedule string `json:"osPatchSchedule,omitempty"`

	PauseAgentUpgrades string `json:"pauseAgentUpgrades,omitempty"`

	ReviewRepaveState string `json:"reviewRepaveState,omitempty"`

	ScanPolicy ClusterAwsScanPolicy `json:"scanPolicy,omitempty"`

	SkipCompletion bool `json:"skipCompletion,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

type ClusterAwsBackupPolicy struct {
	Prefix string `json:"prefix,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	ExpiryInHour int `json:"expiryInHour,omitempty"`

	BackupLocationId string `json:"backupLocationId,omitempty"`

	IncludeDisks bool `json:"includeDisks,omitempty"`

	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	Namespaces []string `json:"namespaces,omitempty"`

	ClusterUids []string `json:"clusterUids,omitempty"`

	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

type ClusterAwsCloudConfig struct {
	SshKeyName string `json:"sshKeyName,omitempty"`

	Region string `json:"region,omitempty"`

	VpcId string `json:"vpcId,omitempty"`

	ControlPlaneLb string `json:"controlPlaneLb,omitempty"`
}

type ClusterAwsClusterProfile struct {
	Pack ClusterAwsClusterProfilePack `json:"pack,omitempty"`

	Id string `json:"id,omitempty"`
}

type ClusterAwsClusterProfilePack struct {
	Manifest ClusterAwsClusterProfilePackManifest `json:"manifest,omitempty"`

	Uid string `json:"uid,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`

	RegistryUid string `json:"registryUid,omitempty"`

	Tag string `json:"tag,omitempty"`

	Values string `json:"values,omitempty"`
}

type ClusterAwsClusterProfilePackManifest struct {
	Name string `json:"name,omitempty"`

	Content string `json:"content,omitempty"`

	Uid string `json:"uid,omitempty"`
}

type ClusterAwsClusterRbacBinding struct {
	Type string `json:"type,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Role map[string]string `json:"role,omitempty"`

	Subjects ClusterAwsClusterRbacBindingSubjects `json:"subjects,omitempty"`
}

type ClusterAwsClusterRbacBindingSubjects struct {
	Namespace string `json:"namespace,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

type ClusterAwsHostConfig struct {
	HostEndpointType string `json:"hostEndpointType,omitempty"`

	IngressHost string `json:"ingressHost,omitempty"`

	ExternalTrafficPolicy string `json:"externalTrafficPolicy,omitempty"`

	LoadBalancerSourceRanges string `json:"loadBalancerSourceRanges,omitempty"`
}

type ClusterAwsMachinePool struct {
	Max int `json:"max,omitempty"`

	Azs []string `json:"azs,omitempty"`

	InstanceType string `json:"instanceType,omitempty"`

	Min int `json:"min,omitempty"`

	Node ClusterAwsMachinePoolNode `json:"node,omitempty"`

	ControlPlane bool `json:"controlPlane,omitempty"`

	UpdateStrategy string `json:"updateStrategy,omitempty"`

	Name string `json:"name,omitempty"`

	MaxPrice string `json:"maxPrice,omitempty"`

	AzSubnets map[string]string `json:"azSubnets,omitempty"`

	AdditionalSecurityGroups []string `json:"additionalSecurityGroups,omitempty"`

	Taints ClusterAwsMachinePoolTaints `json:"taints,omitempty"`

	ControlPlaneAsWorker bool `json:"controlPlaneAsWorker,omitempty"`

	NodeRepaveInterval int `json:"nodeRepaveInterval,omitempty"`

	CapacityType string `json:"capacityType,omitempty"`

	DiskSizeGb int `json:"diskSizeGb,omitempty"`

	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	Count int `json:"count,omitempty"`
}

type ClusterAwsMachinePoolNode struct {
	NodeId string `json:"nodeId,omitempty"`

	Action string `json:"action,omitempty"`
}

type ClusterAwsMachinePoolTaints struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Effect string `json:"effect,omitempty"`
}

type ClusterAwsNamespaces struct {
	Name string `json:"name,omitempty"`

	ResourceAllocation map[string]string `json:"resourceAllocation,omitempty"`

	ImagesBlacklist []string `json:"imagesBlacklist,omitempty"`
}

type ClusterAwsScanPolicy struct {
	ConformanceScanSchedule string `json:"conformanceScanSchedule,omitempty"`

	ConfigurationScanSchedule string `json:"configurationScanSchedule,omitempty"`

	PenetrationScanSchedule string `json:"penetrationScanSchedule,omitempty"`
}

// ClusterAwsStatus defines the observed state of ClusterAws
type ClusterAwsStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	AdminKubeConfig string `json:"adminKubeConfig,omitempty"`

	CloudConfigId string `json:"cloudConfigId,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`

	LocationConfig ClusterAwsLocationConfig `json:"locationConfig,omitempty"`
}

type ClusterAwsLocationConfig struct {
	CountryName string `json:"countryName,omitempty"`

	RegionCode string `json:"regionCode,omitempty"`

	RegionName string `json:"regionName,omitempty"`

	Latitude resource.Quantity `json:"latitude,omitempty"`

	Longitude resource.Quantity `json:"longitude,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterAws is the Schema for the clusteraws API
type ClusterAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAwsSpec   `json:"spec,omitempty"`
	Status ClusterAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterAwsList contains a list of ClusterAws
type ClusterAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAws `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterAws{}, &ClusterAwsList{})
}
