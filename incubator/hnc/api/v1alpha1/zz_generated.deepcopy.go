// +build !ignore_autogenerated

/*

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
func (in *AffectedObject) DeepCopyInto(out *AffectedObject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AffectedObject.
func (in *AffectedObject) DeepCopy() *AffectedObject {
	if in == nil {
		return nil
	}
	out := new(AffectedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Affects != nil {
		in, out := &in.Affects, &out.Affects
		*out = make([]AffectedObject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HNCConfiguration) DeepCopyInto(out *HNCConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HNCConfiguration.
func (in *HNCConfiguration) DeepCopy() *HNCConfiguration {
	if in == nil {
		return nil
	}
	out := new(HNCConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HNCConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HNCConfigurationCondition) DeepCopyInto(out *HNCConfigurationCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HNCConfigurationCondition.
func (in *HNCConfigurationCondition) DeepCopy() *HNCConfigurationCondition {
	if in == nil {
		return nil
	}
	out := new(HNCConfigurationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HNCConfigurationList) DeepCopyInto(out *HNCConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HNCConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HNCConfigurationList.
func (in *HNCConfigurationList) DeepCopy() *HNCConfigurationList {
	if in == nil {
		return nil
	}
	out := new(HNCConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HNCConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HNCConfigurationSpec) DeepCopyInto(out *HNCConfigurationSpec) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]TypeSynchronizationSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HNCConfigurationSpec.
func (in *HNCConfigurationSpec) DeepCopy() *HNCConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(HNCConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HNCConfigurationStatus) DeepCopyInto(out *HNCConfigurationStatus) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]TypeSynchronizationStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HNCConfigurationCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HNCConfigurationStatus.
func (in *HNCConfigurationStatus) DeepCopy() *HNCConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(HNCConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchicalNamespace) DeepCopyInto(out *HierarchicalNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchicalNamespace.
func (in *HierarchicalNamespace) DeepCopy() *HierarchicalNamespace {
	if in == nil {
		return nil
	}
	out := new(HierarchicalNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HierarchicalNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchicalNamespaceList) DeepCopyInto(out *HierarchicalNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HierarchicalNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchicalNamespaceList.
func (in *HierarchicalNamespaceList) DeepCopy() *HierarchicalNamespaceList {
	if in == nil {
		return nil
	}
	out := new(HierarchicalNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HierarchicalNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchicalNamespaceStatus) DeepCopyInto(out *HierarchicalNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchicalNamespaceStatus.
func (in *HierarchicalNamespaceStatus) DeepCopy() *HierarchicalNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(HierarchicalNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchyConfiguration) DeepCopyInto(out *HierarchyConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchyConfiguration.
func (in *HierarchyConfiguration) DeepCopy() *HierarchyConfiguration {
	if in == nil {
		return nil
	}
	out := new(HierarchyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HierarchyConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchyConfigurationList) DeepCopyInto(out *HierarchyConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HierarchyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchyConfigurationList.
func (in *HierarchyConfigurationList) DeepCopy() *HierarchyConfigurationList {
	if in == nil {
		return nil
	}
	out := new(HierarchyConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HierarchyConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchyConfigurationSpec) DeepCopyInto(out *HierarchyConfigurationSpec) {
	*out = *in
	if in.RequiredChildren != nil {
		in, out := &in.RequiredChildren, &out.RequiredChildren
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchyConfigurationSpec.
func (in *HierarchyConfigurationSpec) DeepCopy() *HierarchyConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(HierarchyConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HierarchyConfigurationStatus) DeepCopyInto(out *HierarchyConfigurationStatus) {
	*out = *in
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HierarchyConfigurationStatus.
func (in *HierarchyConfigurationStatus) DeepCopy() *HierarchyConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(HierarchyConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeSynchronizationSpec) DeepCopyInto(out *TypeSynchronizationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeSynchronizationSpec.
func (in *TypeSynchronizationSpec) DeepCopy() *TypeSynchronizationSpec {
	if in == nil {
		return nil
	}
	out := new(TypeSynchronizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeSynchronizationStatus) DeepCopyInto(out *TypeSynchronizationStatus) {
	*out = *in
	if in.NumPropagatedObjects != nil {
		in, out := &in.NumPropagatedObjects, &out.NumPropagatedObjects
		*out = new(int32)
		**out = **in
	}
	if in.NumSourceObjects != nil {
		in, out := &in.NumSourceObjects, &out.NumSourceObjects
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeSynchronizationStatus.
func (in *TypeSynchronizationStatus) DeepCopy() *TypeSynchronizationStatus {
	if in == nil {
		return nil
	}
	out := new(TypeSynchronizationStatus)
	in.DeepCopyInto(out)
	return out
}
