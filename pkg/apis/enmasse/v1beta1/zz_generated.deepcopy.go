// +build !ignore_autogenerated

/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Address) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressList) DeepCopyInto(out *AddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Address, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressList.
func (in *AddressList) DeepCopy() *AddressList {
	if in == nil {
		return nil
	}
	out := new(AddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpace) DeepCopyInto(out *AddressSpace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpace.
func (in *AddressSpace) DeepCopy() *AddressSpace {
	if in == nil {
		return nil
	}
	out := new(AddressSpace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressSpace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpaceList) DeepCopyInto(out *AddressSpaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddressSpace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpaceList.
func (in *AddressSpaceList) DeepCopy() *AddressSpaceList {
	if in == nil {
		return nil
	}
	out := new(AddressSpaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressSpaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpaceSpec) DeepCopyInto(out *AddressSpaceSpec) {
	*out = *in
	if in.AuthenticationService != nil {
		in, out := &in.AuthenticationService, &out.AuthenticationService
		*out = new(AuthenticationService)
		(*in).DeepCopyInto(*out)
	}
	if in.Ednpoints != nil {
		in, out := &in.Ednpoints, &out.Ednpoints
		*out = make([]EndpointSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpaceSpec.
func (in *AddressSpaceSpec) DeepCopy() *AddressSpaceSpec {
	if in == nil {
		return nil
	}
	out := new(AddressSpaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpaceStatus) DeepCopyInto(out *AddressSpaceStatus) {
	*out = *in
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EndpointStatus != nil {
		in, out := &in.EndpointStatus, &out.EndpointStatus
		*out = make([]EndpointStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpaceStatus.
func (in *AddressSpaceStatus) DeepCopy() *AddressSpaceStatus {
	if in == nil {
		return nil
	}
	out := new(AddressSpaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpec) DeepCopyInto(out *AddressSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpec.
func (in *AddressSpec) DeepCopy() *AddressSpec {
	if in == nil {
		return nil
	}
	out := new(AddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationService) DeepCopyInto(out *AuthenticationService) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationService.
func (in *AuthenticationService) DeepCopy() *AuthenticationService {
	if in == nil {
		return nil
	}
	out := new(AuthenticationService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Detail) DeepCopyInto(out *Detail) {
	*out = *in
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
	if in.FloatValue != nil {
		in, out := &in.FloatValue, &out.FloatValue
		*out = new(float64)
		**out = **in
	}
	if in.BooleanValue != nil {
		in, out := &in.BooleanValue, &out.BooleanValue
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Detail.
func (in *Detail) DeepCopy() *Detail {
	if in == nil {
		return nil
	}
	out := new(Detail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(CertificateSpec)
		**out = **in
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(ExposeSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointStatus) DeepCopyInto(out *EndpointStatus) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.ExternalPorts != nil {
		in, out := &in.ExternalPorts, &out.ExternalPorts
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointStatus.
func (in *EndpointStatus) DeepCopy() *EndpointStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposeSpec) DeepCopyInto(out *ExposeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposeSpec.
func (in *ExposeSpec) DeepCopy() *ExposeSpec {
	if in == nil {
		return nil
	}
	out := new(ExposeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}
