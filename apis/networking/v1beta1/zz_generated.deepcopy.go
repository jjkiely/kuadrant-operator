// +build !ignore_autogenerated

/*
Copyright 2021 Red Hat, Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyAuth) DeepCopyInto(out *APIKeyAuth) {
	*out = *in
	in.APIKeySelector.DeepCopyInto(&out.APIKeySelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyAuth.
func (in *APIKeyAuth) DeepCopy() *APIKeyAuth {
	if in == nil {
		return nil
	}
	out := new(APIKeyAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]*Operation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Operation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.BackendServer != nil {
		in, out := &in.BackendServer, &out.BackendServer
		*out = make([]*BackendServer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BackendServer)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SecurityScheme != nil {
		in, out := &in.SecurityScheme, &out.SecurityScheme
		*out = make([]*SecurityScheme, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityScheme)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendServer) DeepCopyInto(out *BackendServer) {
	*out = *in
	in.ServiceRef.DeepCopyInto(&out.ServiceRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendServer.
func (in *BackendServer) DeepCopy() *BackendServer {
	if in == nil {
		return nil
	}
	out := new(BackendServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectAuth) DeepCopyInto(out *OpenIDConnectAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectAuth.
func (in *OpenIDConnectAuth) DeepCopy() *OpenIDConnectAuth {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]*SecurityParameters, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityParameters)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityParameters) DeepCopyInto(out *SecurityParameters) {
	*out = *in
	if in.APIKeyAuth != nil {
		in, out := &in.APIKeyAuth, &out.APIKeyAuth
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Oauth2 != nil {
		in, out := &in.Oauth2, &out.Oauth2
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityParameters.
func (in *SecurityParameters) DeepCopy() *SecurityParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityScheme) DeepCopyInto(out *SecurityScheme) {
	*out = *in
	if in.APIKeyAuth != nil {
		in, out := &in.APIKeyAuth, &out.APIKeyAuth
		*out = new(APIKeyAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenIDConnectAuth != nil {
		in, out := &in.OpenIDConnectAuth, &out.OpenIDConnectAuth
		*out = new(OpenIDConnectAuth)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityScheme.
func (in *SecurityScheme) DeepCopy() *SecurityScheme {
	if in == nil {
		return nil
	}
	out := new(SecurityScheme)
	in.DeepCopyInto(out)
	return out
}
