/*
Copyright 2023.

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
	"github.com/fluxcd/pkg/apis/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AKS configures how AKS clusters are reflected.
type AKS struct {
	// SubscriptionID is the Azure subscription ID
	// +required
	SubscriptionID string `json:"subscriptionID"`
}

// CAPI defines the desired state of CAPI
type CAPI struct {
	// Current Cluster name indicating the management cluster
	// used to avoid choosing the cluster the controller is running in.
	CurrentClusterRef Cluster `json:"currentClusterRef,omitempty"`
}

type Cluster struct {
	// Name is the name of the cluster
	// +required
	Name string `json:"name"`
}

// String returns the string representation of the Cluster.
func (c Cluster) String() string {
	return c.Name
}

// EKS configures how AKS clusters are reflected.
type EKS struct {
	// Region is the AWS region
	// +required
	Region string `json:"region"`
}

// AutomatedClusterDiscoverySpec defines the desired state of
// AutomatedClusterDiscovery.
type AutomatedClusterDiscoverySpec struct {
	// Type is the provider type
	// +kubebuilder:validation:Enum=aks;eks;capi
	Type string `json:"type"`

	// If DisableTags is true, labels will not be applied to the generated
	// Clusters from the tags on the upstream Clusters.
	// +optional
	DisableTags bool `json:"disableTags"`

	// AKS configures discovery of AKS clusters from Azure. Must be provided if
	// the type is aks.
	AKS *AKS `json:"aks,omitempty"`

	// CAPI configures discovery of CAPI clusters
	CAPI *CAPI `json:"capi,omitempty"`

	// EKS configures discovery of EKS clusters from AWS. Must be provided if
	// the type is eks.
	EKS *EKS `json:"eks,omitempty"`

	// The interval at which to run the discovery
	// +required
	Interval metav1.Duration `json:"interval"`

	// Suspend tells the controller to suspend the reconciliation of this
	// AutomatedClusterDiscovery.
	// +optional
	Suspend bool `json:"suspend,omitempty"`

	// Labels to add to all generated resources.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`
	// Annotations to add to all generated resources.
	CommonAnnotations map[string]string `json:"commonAnnotations,omitempty"`
}

// AutomatedClusterDiscoveryStatus defines the observed state of AutomatedClusterDiscovery
type AutomatedClusterDiscoveryStatus struct {
	meta.ReconcileRequestStatus `json:",inline"`

	// ObservedGeneration is the last observed generation of the
	// object.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions holds the conditions for the AutomatedClusterDiscovery
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Inventory contains the list of Kubernetes resource object references that
	// have been successfully applied
	// +optional
	Inventory *ResourceInventory `json:"inventory,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName="acd"
//+kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message",description=""

// AutomatedClusterDiscovery is the Schema for the automatedclusterdiscoveries API
type AutomatedClusterDiscovery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AutomatedClusterDiscoverySpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status AutomatedClusterDiscoveryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AutomatedClusterDiscoveryList contains a list of AutomatedClusterDiscovery
type AutomatedClusterDiscoveryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutomatedClusterDiscovery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AutomatedClusterDiscovery{}, &AutomatedClusterDiscoveryList{})
}
