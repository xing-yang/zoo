/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PandaSpec defines the desired state of Panda
type PandaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	BirthYear int32 `json:"birthYear" protobuf:"varint,1,opt,name=birthYear"`

	BirthMonth int32 `json:"birthMonth" protobuf:"varint,2,opt,name=birthMonth"`

	BirthDay int32 `json:"birthDay" protobuf:"varint,3,opt,name=birthDay"`

	BirthPlace string `json:"birthPlace" protobuf:"bytes,4,opt,name=birthPlace"`

	MomName string `json:"momName" protobuf:"bytes,5,opt,name=momName"`

	DadName string `json:"dadName" protobuf:"bytes,6,opt,name=dadName"`

	BirthWeight int32 `json:"birthWeight" protobuf:"varint,7,opt,name=birthWeight"`
}

// PandaStatus defines the observed state of Panda
type PandaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Age by number of months
	Age int32 `json:"age" protobuf:"varint,1,opt,name=age"`

	// Weight in pounds
	Weight int32 `json:"weight" protobuf:"varint,2,opt,name=weight"`

	// Daily bamboo consumption
	// +optional
	BambooConsumption *int64 `json:"bambooConsumption,omitempty" protobuf:"varint,3,opt,name=bambooConsumption"`

	// AppetiteScale can be Low, Normal, or High
	// +optional
	AppetiteScale *AppetiteScale `json:"appetiteScale,omitempty" protobuf:"bytes,4,opt,name=appetiteScale"`

	// WeightScale can be Low, Normal, or High
	// +optional
	WeightScale *WeightScale `json:"weightScale,omitempty" protobuf:"bytes,5,opt,name=weightScale"`
}

type AppetiteScale string

const (
	AppetiteLow    AppetiteScale = "Low"
	AppetiteNormal AppetiteScale = "Normal"
	AppetiteHigh   AppetiteScale = "High"
)

type WeightScale string

const (
	WeightLow    WeightScale = "Low"
	WeightNormal WeightScale = "Normal"
	WeightHigh   WeightScale = "High"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Panda is the Schema for the pandas API
// +k8s:openapi-gen=true
type Panda struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PandaSpec   `json:"spec,omitempty"`
	Status PandaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PandaList contains a list of Panda
type PandaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Panda `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Panda{}, &PandaList{})
}
