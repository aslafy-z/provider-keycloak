//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapper) DeepCopyInto(out *RoleMapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapper.
func (in *RoleMapper) DeepCopy() *RoleMapper {
	if in == nil {
		return nil
	}
	out := new(RoleMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleMapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperInitParameters) DeepCopyInto(out *RoleMapperInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleIDRef != nil {
		in, out := &in.RoleIDRef, &out.RoleIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleIDSelector != nil {
		in, out := &in.RoleIDSelector, &out.RoleIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperInitParameters.
func (in *RoleMapperInitParameters) DeepCopy() *RoleMapperInitParameters {
	if in == nil {
		return nil
	}
	out := new(RoleMapperInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperList) DeepCopyInto(out *RoleMapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperList.
func (in *RoleMapperList) DeepCopy() *RoleMapperList {
	if in == nil {
		return nil
	}
	out := new(RoleMapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleMapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperObservation) DeepCopyInto(out *RoleMapperObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperObservation.
func (in *RoleMapperObservation) DeepCopy() *RoleMapperObservation {
	if in == nil {
		return nil
	}
	out := new(RoleMapperObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperParameters) DeepCopyInto(out *RoleMapperParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.RoleIDRef != nil {
		in, out := &in.RoleIDRef, &out.RoleIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleIDSelector != nil {
		in, out := &in.RoleIDSelector, &out.RoleIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperParameters.
func (in *RoleMapperParameters) DeepCopy() *RoleMapperParameters {
	if in == nil {
		return nil
	}
	out := new(RoleMapperParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperSpec) DeepCopyInto(out *RoleMapperSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperSpec.
func (in *RoleMapperSpec) DeepCopy() *RoleMapperSpec {
	if in == nil {
		return nil
	}
	out := new(RoleMapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapperStatus) DeepCopyInto(out *RoleMapperStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapperStatus.
func (in *RoleMapperStatus) DeepCopy() *RoleMapperStatus {
	if in == nil {
		return nil
	}
	out := new(RoleMapperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapper) DeepCopyInto(out *SamlProtocolMapper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapper.
func (in *SamlProtocolMapper) DeepCopy() *SamlProtocolMapper {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SamlProtocolMapper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperInitParameters) DeepCopyInto(out *SamlProtocolMapperInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.ClientScopeIDRef != nil {
		in, out := &in.ClientScopeIDRef, &out.ClientScopeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeIDSelector != nil {
		in, out := &in.ClientScopeIDSelector, &out.ClientScopeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ProtocolMapper != nil {
		in, out := &in.ProtocolMapper, &out.ProtocolMapper
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperInitParameters.
func (in *SamlProtocolMapperInitParameters) DeepCopy() *SamlProtocolMapperInitParameters {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperList) DeepCopyInto(out *SamlProtocolMapperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SamlProtocolMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperList.
func (in *SamlProtocolMapperList) DeepCopy() *SamlProtocolMapperList {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SamlProtocolMapperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperObservation) DeepCopyInto(out *SamlProtocolMapperObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ProtocolMapper != nil {
		in, out := &in.ProtocolMapper, &out.ProtocolMapper
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperObservation.
func (in *SamlProtocolMapperObservation) DeepCopy() *SamlProtocolMapperObservation {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperParameters) DeepCopyInto(out *SamlProtocolMapperParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientIDRef != nil {
		in, out := &in.ClientIDRef, &out.ClientIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientIDSelector != nil {
		in, out := &in.ClientIDSelector, &out.ClientIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeID != nil {
		in, out := &in.ClientScopeID, &out.ClientScopeID
		*out = new(string)
		**out = **in
	}
	if in.ClientScopeIDRef != nil {
		in, out := &in.ClientScopeIDRef, &out.ClientScopeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientScopeIDSelector != nil {
		in, out := &in.ClientScopeIDSelector, &out.ClientScopeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ProtocolMapper != nil {
		in, out := &in.ProtocolMapper, &out.ProtocolMapper
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RealmIDRef != nil {
		in, out := &in.RealmIDRef, &out.RealmIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmIDSelector != nil {
		in, out := &in.RealmIDSelector, &out.RealmIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperParameters.
func (in *SamlProtocolMapperParameters) DeepCopy() *SamlProtocolMapperParameters {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperSpec) DeepCopyInto(out *SamlProtocolMapperSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperSpec.
func (in *SamlProtocolMapperSpec) DeepCopy() *SamlProtocolMapperSpec {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProtocolMapperStatus) DeepCopyInto(out *SamlProtocolMapperStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProtocolMapperStatus.
func (in *SamlProtocolMapperStatus) DeepCopy() *SamlProtocolMapperStatus {
	if in == nil {
		return nil
	}
	out := new(SamlProtocolMapperStatus)
	in.DeepCopyInto(out)
	return out
}
