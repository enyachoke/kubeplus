// +build !ignore_autogenerated

/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cond) DeepCopyInto(out *Cond) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cond.
func (in *Cond) DeepCopy() *Cond {
	if in == nil {
		return nil
	}
	out := new(Cond)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependsOn) DeepCopyInto(out *DependsOn) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependsOn.
func (in *DependsOn) DeepCopy() *DependsOn {
	if in == nil {
		return nil
	}
	out := new(DependsOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limits) DeepCopyInto(out *Limits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in *Limits) DeepCopy() *Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mon) DeepCopyInto(out *Mon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mon.
func (in *Mon) DeepCopy() *Mon {
	if in == nil {
		return nil
	}
	out := new(Mon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewResource) DeepCopyInto(out *NewResource) {
	*out = *in
	out.Resource = in.Resource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewResource.
func (in *NewResource) DeepCopy() *NewResource {
	if in == nil {
		return nil
	}
	out := new(NewResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pol) DeepCopyInto(out *Pol) {
	*out = *in
	out.PolicyResources = in.PolicyResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pol.
func (in *Pol) DeepCopy() *Pol {
	if in == nil {
		return nil
	}
	out := new(Pol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyResources) DeepCopyInto(out *PolicyResources) {
	*out = *in
	out.Limits = in.Limits
	out.Requests = in.Requests
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyResources.
func (in *PolicyResources) DeepCopy() *PolicyResources {
	if in == nil {
		return nil
	}
	out := new(PolicyResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Res) DeepCopyInto(out *Res) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Res.
func (in *Res) DeepCopy() *Res {
	if in == nil {
		return nil
	}
	out := new(Res)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceComposition) DeepCopyInto(out *ResourceComposition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceComposition.
func (in *ResourceComposition) DeepCopy() *ResourceComposition {
	if in == nil {
		return nil
	}
	out := new(ResourceComposition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceComposition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCompositionList) DeepCopyInto(out *ResourceCompositionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceComposition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCompositionList.
func (in *ResourceCompositionList) DeepCopy() *ResourceCompositionList {
	if in == nil {
		return nil
	}
	out := new(ResourceCompositionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceCompositionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCompositionSpec) DeepCopyInto(out *ResourceCompositionSpec) {
	*out = *in
	out.NewResource = in.NewResource
	in.ResPolicy.DeepCopyInto(&out.ResPolicy)
	in.ResMonitor.DeepCopyInto(&out.ResMonitor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCompositionSpec.
func (in *ResourceCompositionSpec) DeepCopy() *ResourceCompositionSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceCompositionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCompositionStatus) DeepCopyInto(out *ResourceCompositionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCompositionStatus.
func (in *ResourceCompositionStatus) DeepCopy() *ResourceCompositionStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceCompositionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEvent) DeepCopyInto(out *ResourceEvent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEvent.
func (in *ResourceEvent) DeepCopy() *ResourceEvent {
	if in == nil {
		return nil
	}
	out := new(ResourceEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceEvent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEventList) DeepCopyInto(out *ResourceEventList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEventList.
func (in *ResourceEventList) DeepCopy() *ResourceEventList {
	if in == nil {
		return nil
	}
	out := new(ResourceEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceEventList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEventSpec) DeepCopyInto(out *ResourceEventSpec) {
	*out = *in
	out.Resource = in.Resource
	out.Condition = in.Condition
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEventSpec.
func (in *ResourceEventSpec) DeepCopy() *ResourceEventSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceEventSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceEventStatus) DeepCopyInto(out *ResourceEventStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceEventStatus.
func (in *ResourceEventStatus) DeepCopy() *ResourceEventStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceEventStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMonitor) DeepCopyInto(out *ResourceMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMonitor.
func (in *ResourceMonitor) DeepCopy() *ResourceMonitor {
	if in == nil {
		return nil
	}
	out := new(ResourceMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMonitorList) DeepCopyInto(out *ResourceMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMonitorList.
func (in *ResourceMonitorList) DeepCopy() *ResourceMonitorList {
	if in == nil {
		return nil
	}
	out := new(ResourceMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMonitorSpec) DeepCopyInto(out *ResourceMonitorSpec) {
	*out = *in
	out.Resource = in.Resource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMonitorSpec.
func (in *ResourceMonitorSpec) DeepCopy() *ResourceMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMonitorStatus) DeepCopyInto(out *ResourceMonitorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMonitorStatus.
func (in *ResourceMonitorStatus) DeepCopy() *ResourceMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicy) DeepCopyInto(out *ResourcePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicy.
func (in *ResourcePolicy) DeepCopy() *ResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyList) DeepCopyInto(out *ResourcePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourcePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyList.
func (in *ResourcePolicyList) DeepCopy() *ResourcePolicyList {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicySpec) DeepCopyInto(out *ResourcePolicySpec) {
	*out = *in
	out.Resource = in.Resource
	out.Policy = in.Policy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicySpec.
func (in *ResourcePolicySpec) DeepCopy() *ResourcePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyStatus) DeepCopyInto(out *ResourcePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyStatus.
func (in *ResourcePolicyStatus) DeepCopy() *ResourcePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackElements) DeepCopyInto(out *StackElements) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]DependsOn, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackElements.
func (in *StackElements) DeepCopy() *StackElements {
	if in == nil {
		return nil
	}
	out := new(StackElements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Values) DeepCopyInto(out *Values) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Values.
func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	out := new(Values)
	in.DeepCopyInto(out)
	return out
}
