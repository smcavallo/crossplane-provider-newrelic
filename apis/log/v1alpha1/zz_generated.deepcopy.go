//go:build !ignore_autogenerated

/*
Copyright 2024 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRule) DeepCopyInto(out *ParsingRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRule.
func (in *ParsingRule) DeepCopy() *ParsingRule {
	if in == nil {
		return nil
	}
	out := new(ParsingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParsingRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleInitParameters) DeepCopyInto(out *ParsingRuleInitParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(float64)
		**out = **in
	}
	if in.Attribute != nil {
		in, out := &in.Attribute, &out.Attribute
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Grok != nil {
		in, out := &in.Grok, &out.Grok
		*out = new(string)
		**out = **in
	}
	if in.Lucene != nil {
		in, out := &in.Lucene, &out.Lucene
		*out = new(string)
		**out = **in
	}
	if in.Matched != nil {
		in, out := &in.Matched, &out.Matched
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleInitParameters.
func (in *ParsingRuleInitParameters) DeepCopy() *ParsingRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleList) DeepCopyInto(out *ParsingRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParsingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleList.
func (in *ParsingRuleList) DeepCopy() *ParsingRuleList {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParsingRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleObservation) DeepCopyInto(out *ParsingRuleObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(float64)
		**out = **in
	}
	if in.Attribute != nil {
		in, out := &in.Attribute, &out.Attribute
		*out = new(string)
		**out = **in
	}
	if in.Deleted != nil {
		in, out := &in.Deleted, &out.Deleted
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Grok != nil {
		in, out := &in.Grok, &out.Grok
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Lucene != nil {
		in, out := &in.Lucene, &out.Lucene
		*out = new(string)
		**out = **in
	}
	if in.Matched != nil {
		in, out := &in.Matched, &out.Matched
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleObservation.
func (in *ParsingRuleObservation) DeepCopy() *ParsingRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleParameters) DeepCopyInto(out *ParsingRuleParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(float64)
		**out = **in
	}
	if in.Attribute != nil {
		in, out := &in.Attribute, &out.Attribute
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Grok != nil {
		in, out := &in.Grok, &out.Grok
		*out = new(string)
		**out = **in
	}
	if in.Lucene != nil {
		in, out := &in.Lucene, &out.Lucene
		*out = new(string)
		**out = **in
	}
	if in.Matched != nil {
		in, out := &in.Matched, &out.Matched
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleParameters.
func (in *ParsingRuleParameters) DeepCopy() *ParsingRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleSpec) DeepCopyInto(out *ParsingRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleSpec.
func (in *ParsingRuleSpec) DeepCopy() *ParsingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParsingRuleStatus) DeepCopyInto(out *ParsingRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParsingRuleStatus.
func (in *ParsingRuleStatus) DeepCopy() *ParsingRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ParsingRuleStatus)
	in.DeepCopyInto(out)
	return out
}