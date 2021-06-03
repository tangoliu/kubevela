// +build !ignore_autogenerated

/*
Copyright 2021 The KubeVela Authors.

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

package common

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRolloutStatus) DeepCopyInto(out *AppRolloutStatus) {
	*out = *in
	in.RolloutStatus.DeepCopyInto(&out.RolloutStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRolloutStatus.
func (in *AppRolloutStatus) DeepCopy() *AppRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(AppRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	in.Rollout.DeepCopyInto(&out.Rollout)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]v1alpha1.TypedReference, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ApplicationComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceTracker != nil {
		in, out := &in.ResourceTracker, &out.ResourceTracker
		*out = new(v1alpha1.TypedReference)
		**out = **in
	}
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = make([]WorkflowStepStatus, len(*in))
		copy(*out, *in)
	}
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationComponentStatus) DeepCopyInto(out *ApplicationComponentStatus) {
	*out = *in
	out.WorkloadDefinition = in.WorkloadDefinition
	if in.Traits != nil {
		in, out := &in.Traits, &out.Traits
		*out = make([]ApplicationTraitStatus, len(*in))
		copy(*out, *in)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]v1alpha1.TypedReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationComponentStatus.
func (in *ApplicationComponentStatus) DeepCopy() *ApplicationComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationTraitStatus) DeepCopyInto(out *ApplicationTraitStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationTraitStatus.
func (in *ApplicationTraitStatus) DeepCopy() *ApplicationTraitStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationTraitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CUE) DeepCopyInto(out *CUE) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CUE.
func (in *CUE) DeepCopy() *CUE {
	if in == nil {
		return nil
	}
	out := new(CUE)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChildResourceKind) DeepCopyInto(out *ChildResourceKind) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChildResourceKind.
func (in *ChildResourceKind) DeepCopy() *ChildResourceKind {
	if in == nil {
		return nil
	}
	out := new(ChildResourceKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionReference) DeepCopyInto(out *DefinitionReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionReference.
func (in *DefinitionReference) DeepCopy() *DefinitionReference {
	if in == nil {
		return nil
	}
	out := new(DefinitionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Helm) DeepCopyInto(out *Helm) {
	*out = *in
	in.Release.DeepCopyInto(&out.Release)
	in.Repository.DeepCopyInto(&out.Repository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Helm.
func (in *Helm) DeepCopy() *Helm {
	if in == nil {
		return nil
	}
	out := new(Helm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kube) DeepCopyInto(out *Kube) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]KubeParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kube.
func (in *Kube) DeepCopy() *Kube {
	if in == nil {
		return nil
	}
	out := new(Kube)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeParameter) DeepCopyInto(out *KubeParameter) {
	*out = *in
	if in.FieldPaths != nil {
		in, out := &in.FieldPaths, &out.FieldPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeParameter.
func (in *KubeParameter) DeepCopy() *KubeParameter {
	if in == nil {
		return nil
	}
	out := new(KubeParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawComponent) DeepCopyInto(out *RawComponent) {
	*out = *in
	in.Raw.DeepCopyInto(&out.Raw)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawComponent.
func (in *RawComponent) DeepCopy() *RawComponent {
	if in == nil {
		return nil
	}
	out := new(RawComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Revision) DeepCopyInto(out *Revision) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Revision.
func (in *Revision) DeepCopy() *Revision {
	if in == nil {
		return nil
	}
	out := new(Revision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schematic) DeepCopyInto(out *Schematic) {
	*out = *in
	if in.KUBE != nil {
		in, out := &in.KUBE, &out.KUBE
		*out = new(Kube)
		(*in).DeepCopyInto(*out)
	}
	if in.CUE != nil {
		in, out := &in.CUE, &out.CUE
		*out = new(CUE)
		**out = **in
	}
	if in.HELM != nil {
		in, out := &in.HELM, &out.HELM
		*out = new(Helm)
		(*in).DeepCopyInto(*out)
	}
	if in.Terraform != nil {
		in, out := &in.Terraform, &out.Terraform
		*out = new(Terraform)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schematic.
func (in *Schematic) DeepCopy() *Schematic {
	if in == nil {
		return nil
	}
	out := new(Schematic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terraform) DeepCopyInto(out *Terraform) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terraform.
func (in *Terraform) DeepCopy() *Terraform {
	if in == nil {
		return nil
	}
	out := new(Terraform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStepStatus) DeepCopyInto(out *WorkflowStepStatus) {
	*out = *in
	out.ResourceRef = in.ResourceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStepStatus.
func (in *WorkflowStepStatus) DeepCopy() *WorkflowStepStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowStepStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadGVK) DeepCopyInto(out *WorkloadGVK) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGVK.
func (in *WorkloadGVK) DeepCopy() *WorkloadGVK {
	if in == nil {
		return nil
	}
	out := new(WorkloadGVK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadTypeDescriptor) DeepCopyInto(out *WorkloadTypeDescriptor) {
	*out = *in
	out.Definition = in.Definition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadTypeDescriptor.
func (in *WorkloadTypeDescriptor) DeepCopy() *WorkloadTypeDescriptor {
	if in == nil {
		return nil
	}
	out := new(WorkloadTypeDescriptor)
	in.DeepCopyInto(out)
	return out
}