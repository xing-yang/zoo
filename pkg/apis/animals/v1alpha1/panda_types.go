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

// PandaSpec defines the desired state of Panda
type PandaSpec struct {
	BirthDay    string `json:"birthDay" protobuf:"bytes,1,opt,name=birthDay"`
	BirthPlace  string `json:"birthPlace" protobuf:"bytes,2,opt,name=birthPlace"`
	BirthWeight int32  `json:"birthWeight" protobuf:"varint,3,opt,name=birthWeight"`
}

// PandaStatus defines the observed state of Panda
type PandaStatus struct {
	Age               int32 `json:"age" protobuf:"varint,1,opt,name=age"`
	Weight            int32 `json:"weight" protobuf:"varint,2,opt,name=weight"`
	BambooConsumption int32 `json:"bambooConsumption,omitempty" protobuf:"varint,3,opt,name=bambooConsumption"`
}

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
