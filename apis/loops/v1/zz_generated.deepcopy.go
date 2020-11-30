// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportSecret) DeepCopyInto(out *ImportSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportSecret.
func (in *ImportSecret) DeepCopy() *ImportSecret {
	if in == nil {
		return nil
	}
	out := new(ImportSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Imports) DeepCopyInto(out *Imports) {
	{
		in := &in
		*out = make(Imports, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Imports.
func (in Imports) DeepCopy() Imports {
	if in == nil {
		return nil
	}
	out := new(Imports)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loop) DeepCopyInto(out *Loop) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	in.State.DeepCopyInto(&out.State)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loop.
func (in *Loop) DeepCopy() *Loop {
	if in == nil {
		return nil
	}
	out := new(Loop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Loop) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopImport) DeepCopyInto(out *LoopImport) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopImport.
func (in *LoopImport) DeepCopy() *LoopImport {
	if in == nil {
		return nil
	}
	out := new(LoopImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopList) DeepCopyInto(out *LoopList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Loop, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopList.
func (in *LoopList) DeepCopy() *LoopList {
	if in == nil {
		return nil
	}
	out := new(LoopList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoopList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopSpec) DeepCopyInto(out *LoopSpec) {
	*out = *in
	out.Every = in.Every
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make(Imports, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopSpec.
func (in *LoopSpec) DeepCopy() *LoopSpec {
	if in == nil {
		return nil
	}
	out := new(LoopSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopState) DeepCopyInto(out *LoopState) {
	*out = *in
	if in.UpdateDate != nil {
		in, out := &in.UpdateDate, &out.UpdateDate
		*out = (*in).DeepCopy()
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopState.
func (in *LoopState) DeepCopy() *LoopState {
	if in == nil {
		return nil
	}
	out := new(LoopState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopStatus) DeepCopyInto(out *LoopStatus) {
	*out = *in
	if in.LastExecution != nil {
		in, out := &in.LastExecution, &out.LastExecution
		*out = (*in).DeepCopy()
	}
	if in.LastExecutionSuccess != nil {
		in, out := &in.LastExecutionSuccess, &out.LastExecutionSuccess
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopStatus.
func (in *LoopStatus) DeepCopy() *LoopStatus {
	if in == nil {
		return nil
	}
	out := new(LoopStatus)
	in.DeepCopyInto(out)
	return out
}
