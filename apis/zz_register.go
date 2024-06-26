/*
Copyright 2024 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/alert/v1alpha1"
	v1alpha1dashboard "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/dashboard/v1alpha1"
	v1alpha1log "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/log/v1alpha1"
	v1alpha1nrql "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/nrql/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/v1alpha1"
	v1beta1 "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/v1beta1"
	v1alpha1workflow "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/workflow/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1dashboard.SchemeBuilder.AddToScheme,
		v1alpha1log.SchemeBuilder.AddToScheme,
		v1alpha1nrql.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1workflow.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
