/*
Copyright 2024 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this ProviderConfigUsageList.
func (p *ProviderConfigUsageList) GetItems() []resource.ProviderConfigUsage {
	items := make([]resource.ProviderConfigUsage, len(p.Items))
	for i := range p.Items {
		items[i] = &p.Items[i]
	}
	return items
}
