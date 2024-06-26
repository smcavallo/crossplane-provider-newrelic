/*
Copyright 2024 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DashboardJSONInitParameters struct {

	// Determines the New Relic account where the dashboard will be created. Defaults to the account associated with the API key used.
	// The New Relic account ID where you want to create the dashboard.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The JSON export of a dashboard. The JSON can be exported from the UI
	// The dashboard's json.
	JSON *string `json:"json,omitempty" tf:"json,omitempty"`
}

type DashboardJSONObservation struct {

	// Determines the New Relic account where the dashboard will be created. Defaults to the account associated with the API key used.
	// The New Relic account ID where you want to create the dashboard.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The unique entity identifier of the dashboard in New Relic.
	// The unique entity identifier of the dashboard in New Relic.
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The JSON export of a dashboard. The JSON can be exported from the UI
	// The dashboard's json.
	JSON *string `json:"json,omitempty" tf:"json,omitempty"`

	// The URL for viewing the dashboard.
	// The URL of the dashboard.
	Permalink *string `json:"permalink,omitempty" tf:"permalink,omitempty"`

	// The date and time when the dashboard was last updated.
	// The date and time when the dashboard was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type DashboardJSONParameters struct {

	// Determines the New Relic account where the dashboard will be created. Defaults to the account associated with the API key used.
	// The New Relic account ID where you want to create the dashboard.
	// +kubebuilder:validation:Optional
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The JSON export of a dashboard. The JSON can be exported from the UI
	// The dashboard's json.
	// +kubebuilder:validation:Optional
	JSON *string `json:"json,omitempty" tf:"json,omitempty"`
}

// DashboardJSONSpec defines the desired state of DashboardJSON
type DashboardJSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardJSONParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DashboardJSONInitParameters `json:"initProvider,omitempty"`
}

// DashboardJSONStatus defines the observed state of DashboardJSON.
type DashboardJSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardJSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DashboardJSON is the Schema for the DashboardJSONs API. Create and manage dashboards from a JSON file.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,newrelic}
type DashboardJSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.json) || (has(self.initProvider) && has(self.initProvider.json))",message="spec.forProvider.json is a required parameter"
	Spec   DashboardJSONSpec   `json:"spec"`
	Status DashboardJSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardJSONList contains a list of DashboardJSONs
type DashboardJSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DashboardJSON `json:"items"`
}

// Repository type metadata.
var (
	DashboardJSON_Kind             = "DashboardJSON"
	DashboardJSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DashboardJSON_Kind}.String()
	DashboardJSON_KindAPIVersion   = DashboardJSON_Kind + "." + CRDGroupVersion.String()
	DashboardJSON_GroupVersionKind = CRDGroupVersion.WithKind(DashboardJSON_Kind)
)

func init() {
	SchemeBuilder.Register(&DashboardJSON{}, &DashboardJSONList{})
}
