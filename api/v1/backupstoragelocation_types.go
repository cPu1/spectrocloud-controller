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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupStorageLocationSpec defines the desired state of BackupStorageLocation
type BackupStorageLocationSpec struct {

	// +kubebuilder:validation:Required
	BucketName string `json:"bucketName"`

	CaCert string `json:"caCert,omitempty"`

	Context string `json:"context,omitempty"`

	// +kubebuilder:validation:Required
	IsDefault bool `json:"isDefault"`

	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	Region string `json:"region"`

	// +kubebuilder:validation:Required
	S3 BackupStorageLocationS3 `json:"s3"`
}

type BackupStorageLocationS3 struct {
	CredentialType string `json:"credentialType,omitempty"`

	AccessKey string `json:"accessKey,omitempty"`

	SecretKey string `json:"secretKey,omitempty"`

	Arn string `json:"arn,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	S3Url string `json:"s3Url,omitempty"`

	S3ForcePathStyle bool `json:"s3ForcePathStyle,omitempty"`
}

// BackupStorageLocationStatus defines the observed state of BackupStorageLocation
type BackupStorageLocationStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BackupStorageLocation is the Schema for the backupstoragelocations API
type BackupStorageLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupStorageLocationSpec   `json:"spec,omitempty"`
	Status BackupStorageLocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupStorageLocationList contains a list of BackupStorageLocation
type BackupStorageLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupStorageLocation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupStorageLocation{}, &BackupStorageLocationList{})
}
