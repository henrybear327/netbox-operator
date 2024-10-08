/*
Copyright 2024 Swisscom (Schweiz) AG.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PrefixClaimSpec defines the desired state of PrefixClaim
type PrefixClaimSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Format=cidr
	//+kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field 'parentPrefix' is immutable"
	ParentPrefix string `json:"parentPrefix"`

	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Pattern=`^\/[0-9]|[1-9][0-9]|1[01][0-9]|12[0-8]$`
	//+kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field 'prefixLength' is immutable"
	PrefixLength string `json:"prefixLength"`

	Site string `json:"site,omitempty"`

	//+kubebuilder:validation:XValidation:rule="self == oldSelf",message="Field 'tenant' is immutable"
	Tenant string `json:"tenant,omitempty"`

	Description string `json:"description,omitempty"`

	Comments string `json:"comments,omitempty"`

	CustomFields map[string]string `json:"customFields,omitempty"`

	PreserveInNetbox bool `json:"preserveInNetbox,omitempty"`
}

// PrefixClaimStatus defines the observed state of PrefixClaim
type PrefixClaimStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Prefix status: container, active, reserved , deprecated
	Prefix     string             `json:"prefix,omitempty"`
	PrefixName string             `json:"prefixName,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Prefix",type=string,JSONPath=`.status.prefix`
// +kubebuilder:printcolumn:name="PrefixAssigned",type=string,JSONPath=`.status.conditions[?(@.type=="PrefixAssigned")].status`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:resource:shortName=pxc
// PrefixClaim is the Schema for the prefixclaims API
type PrefixClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrefixClaimSpec   `json:"spec,omitempty"`
	Status PrefixClaimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrefixClaimList contains a list of PrefixClaim
type PrefixClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrefixClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrefixClaim{}, &PrefixClaimList{})
}

var ConditionPrefixClaimReadyTrue = metav1.Condition{
	Type:    "Ready",
	Status:  "True",
	Reason:  "PrefixClaimResourceReady",
	Message: "PrefixClaim Resource is ready",
}

var ConditionPrefixClaimReadyFalse = metav1.Condition{
	Type:    "Ready",
	Status:  "False",
	Reason:  "PrefixClaimResourceNotReady",
	Message: "PrefixClaim Resource is not ready",
}

var ConditionPrefixAssignedTrue = metav1.Condition{
	Type:    "PrefixAssigned",
	Status:  "True",
	Reason:  "PrefixCRCreated",
	Message: "New prefix fetched from NetBox and prefix CR was created",
}

var ConditionPrefixAssignedFalse = metav1.Condition{
	Type:    "PrefixAssigned",
	Status:  "False",
	Reason:  "PrefixCRNotCreated",
	Message: "Failed to fetch new Prefix from NetBox",
}
