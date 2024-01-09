// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type DropRuleInitParameters struct {

	// Account where the drop rule will be put. Defaults to the account associated with the API key used.
	// Account with the NRQL drop rule will be put.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// An action type specifying how to apply the NRQL string (either drop_data, drop_attributes, or  drop_attributes_from_metric_aggregates).
	// The drop rule action (drop_data, drop_attributes, or drop_attributes_from_metric_aggregates).
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The description of the drop rule.
	// Provides additional information about the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A NRQL string that specifies what data types to drop.
	// Explains which data to apply the drop rule to.
	Nrql *string `json:"nrql,omitempty" tf:"nrql,omitempty"`
}

type DropRuleObservation struct {

	// Account where the drop rule will be put. Defaults to the account associated with the API key used.
	// Account with the NRQL drop rule will be put.
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// An action type specifying how to apply the NRQL string (either drop_data, drop_attributes, or  drop_attributes_from_metric_aggregates).
	// The drop rule action (drop_data, drop_attributes, or drop_attributes_from_metric_aggregates).
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The description of the drop rule.
	// Provides additional information about the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A NRQL string that specifies what data types to drop.
	// Explains which data to apply the drop rule to.
	Nrql *string `json:"nrql,omitempty" tf:"nrql,omitempty"`

	// The id, uniquely identifying the rule.
	// The id, uniquely identifying the rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type DropRuleParameters struct {

	// Account where the drop rule will be put. Defaults to the account associated with the API key used.
	// Account with the NRQL drop rule will be put.
	// +kubebuilder:validation:Optional
	AccountID *float64 `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// An action type specifying how to apply the NRQL string (either drop_data, drop_attributes, or  drop_attributes_from_metric_aggregates).
	// The drop rule action (drop_data, drop_attributes, or drop_attributes_from_metric_aggregates).
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The description of the drop rule.
	// Provides additional information about the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A NRQL string that specifies what data types to drop.
	// Explains which data to apply the drop rule to.
	// +kubebuilder:validation:Optional
	Nrql *string `json:"nrql,omitempty" tf:"nrql,omitempty"`
}

// DropRuleSpec defines the desired state of DropRule
type DropRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DropRuleParameters `json:"forProvider"`
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
	InitProvider DropRuleInitParameters `json:"initProvider,omitempty"`
}

// DropRuleStatus defines the observed state of DropRule.
type DropRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DropRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DropRule is the Schema for the DropRules API. Create and manage NRQL Drop Rules.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,newrelic}
type DropRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nrql) || (has(self.initProvider) && has(self.initProvider.nrql))",message="spec.forProvider.nrql is a required parameter"
	Spec   DropRuleSpec   `json:"spec"`
	Status DropRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DropRuleList contains a list of DropRules
type DropRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DropRule `json:"items"`
}

// Repository type metadata.
var (
	DropRule_Kind             = "DropRule"
	DropRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DropRule_Kind}.String()
	DropRule_KindAPIVersion   = DropRule_Kind + "." + CRDGroupVersion.String()
	DropRule_GroupVersionKind = CRDGroupVersion.WithKind(DropRule_Kind)
)

func init() {
	SchemeBuilder.Register(&DropRule{}, &DropRuleList{})
}