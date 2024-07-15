//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLBPodBinding) DeepCopyInto(out *CLBPodBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLBPodBinding.
func (in *CLBPodBinding) DeepCopy() *CLBPodBinding {
	if in == nil {
		return nil
	}
	out := new(CLBPodBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CLBPodBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLBPodBindingList) DeepCopyInto(out *CLBPodBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CLBPodBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLBPodBindingList.
func (in *CLBPodBindingList) DeepCopy() *CLBPodBindingList {
	if in == nil {
		return nil
	}
	out := new(CLBPodBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CLBPodBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLBPodBindingSpec) DeepCopyInto(out *CLBPodBindingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLBPodBindingSpec.
func (in *CLBPodBindingSpec) DeepCopy() *CLBPodBindingSpec {
	if in == nil {
		return nil
	}
	out := new(CLBPodBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLBPodBindingStatus) DeepCopyInto(out *CLBPodBindingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLBPodBindingStatus.
func (in *CLBPodBindingStatus) DeepCopy() *CLBPodBindingStatus {
	if in == nil {
		return nil
	}
	out := new(CLBPodBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedCLBService) DeepCopyInto(out *DedicatedCLBService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedCLBService.
func (in *DedicatedCLBService) DeepCopy() *DedicatedCLBService {
	if in == nil {
		return nil
	}
	out := new(DedicatedCLBService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedCLBService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedCLBServiceList) DeepCopyInto(out *DedicatedCLBServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DedicatedCLBService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedCLBServiceList.
func (in *DedicatedCLBServiceList) DeepCopy() *DedicatedCLBServiceList {
	if in == nil {
		return nil
	}
	out := new(DedicatedCLBServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedCLBServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedCLBServiceSpec) DeepCopyInto(out *DedicatedCLBServiceSpec) {
	*out = *in
	if in.ExistedLbIds != nil {
		in, out := &in.ExistedLbIds, &out.ExistedLbIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedCLBServiceSpec.
func (in *DedicatedCLBServiceSpec) DeepCopy() *DedicatedCLBServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DedicatedCLBServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedCLBServiceStatus) DeepCopyInto(out *DedicatedCLBServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedCLBServiceStatus.
func (in *DedicatedCLBServiceStatus) DeepCopy() *DedicatedCLBServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DedicatedCLBServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedNatgwService) DeepCopyInto(out *DedicatedNatgwService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedNatgwService.
func (in *DedicatedNatgwService) DeepCopy() *DedicatedNatgwService {
	if in == nil {
		return nil
	}
	out := new(DedicatedNatgwService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedNatgwService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedNatgwServiceList) DeepCopyInto(out *DedicatedNatgwServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DedicatedNatgwService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedNatgwServiceList.
func (in *DedicatedNatgwServiceList) DeepCopy() *DedicatedNatgwServiceList {
	if in == nil {
		return nil
	}
	out := new(DedicatedNatgwServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DedicatedNatgwServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedNatgwServiceSpec) DeepCopyInto(out *DedicatedNatgwServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedNatgwServiceSpec.
func (in *DedicatedNatgwServiceSpec) DeepCopy() *DedicatedNatgwServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DedicatedNatgwServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DedicatedNatgwServiceStatus) DeepCopyInto(out *DedicatedNatgwServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DedicatedNatgwServiceStatus.
func (in *DedicatedNatgwServiceStatus) DeepCopy() *DedicatedNatgwServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DedicatedNatgwServiceStatus)
	in.DeepCopyInto(out)
	return out
}