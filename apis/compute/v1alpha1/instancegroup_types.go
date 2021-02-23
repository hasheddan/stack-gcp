/*
Copyright 2020 The Crossplane Authors.

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
	compute "google.golang.org/api/compute/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

var _ = compute.InstanceGroup{}

// InstanceGroupParameters define the desired state of a Google Compute Engine VPC
// InstanceGroup. Most fields map directly to a InstanceGroup:
// https://cloud.google.com/compute/docs/reference/rest/v1/networks
type InstanceGroupParameters struct {
	// Description: An optional description of this resource. Provide this
	// property when you create the resource.
	Description string `json:"description,omitempty"`

	// NamedPorts: Assigns a name to a port number. For example: {name:
	// "http", port: 80}
	//
	// This allows the system to reference ports by the assigned name
	// instead of a port number. Named ports can also contain multiple
	// ports. For example: [{name: "http", port: 80},{name: "http", port:
	// 8080}]
	//
	// Named ports apply to all instances in this instance group.
	NamedPorts []*NamedPort `json:"namedPorts,omitempty"`

	// Network: The URL of the network to which all instances in the
	// instance group belong.
	// +optional
	// +immutable
	Network *string `json:"network,omitempty"`

	// NetworkRef references a Network and retrieves its URI
	// +optional
	// +immutable
	NetworkRef *xpv1.Reference `json:"networkRef,omitempty"`

	// NetworkSelector selects a reference to a Network
	// +optional
	// +immutable
	NetworkSelector *xpv1.Selector `json:"networkSelector,omitempty"`
}

// NamedPort is the named port. For example: .
type NamedPort struct {
	// Name: The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name string `json:"name,omitempty"`

	// Port: The port number, which can be a value between 1 and 65535.
	Port int64 `json:"port,omitempty"`
}

// A InstanceGroupObservation represents the observed state of a Google Compute Engine
// VPC InstanceGroup.
type InstanceGroupObservation struct {
	// CreationTimestamp: [Output Only] The creation timestamp for this
	// instance group in RFC3339 text format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Fingerprint: [Output Only] The fingerprint of the named ports. The
	// system uses this fingerprint to detect conflicts when multiple users
	// change the named ports concurrently.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Id: [Output Only] A unique identifier for this instance group,
	// generated by the server.
	ID uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] The resource type, which is always
	// compute#instanceGroup for instance groups.
	Kind string `json:"kind,omitempty"`

	// Region: [Output Only] The URL of the region where the instance group
	// is located (for regional resources).
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] The URL for this instance group. The server
	// generates this URL.
	SelfLink string `json:"selfLink,omitempty"`

	// Size: [Output Only] The total number of instances in the instance
	// group.
	Size int64 `json:"size,omitempty"`

	// Subnetwork: [Output Only] The URL of the subnetwork to which all
	// instances in the instance group belong.
	Subnetwork string `json:"subnetwork,omitempty"`

	// Zone: [Output Only] The URL of the zone where the instance group is
	// located (for zonal resources).
	Zone string `json:"zone,omitempty"`
}

// A InstanceGroupSpec defines the desired state of a InstanceGroup.
type InstanceGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InstanceGroupParameters `json:"forProvider"`
}

// A InstanceGroupStatus represents the observed state of a InstanceGroup.
type InstanceGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InstanceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A InstanceGroup is a managed resource that represents a Google Compute Engine VPC
// InstanceGroup.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceGroupSpec   `json:"spec"`
	Status InstanceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupList contains a list of InstanceGroup.
type InstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceGroup `json:"items"`
}
