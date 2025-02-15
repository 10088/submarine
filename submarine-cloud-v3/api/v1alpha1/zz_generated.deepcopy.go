//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonImage) DeepCopyInto(out *CommonImage) {
	*out = *in
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonImage.
func (in *CommonImage) DeepCopy() *CommonImage {
	if in == nil {
		return nil
	}
	out := new(CommonImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Submarine) DeepCopyInto(out *Submarine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Submarine.
func (in *Submarine) DeepCopy() *Submarine {
	if in == nil {
		return nil
	}
	out := new(Submarine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Submarine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineCommon) DeepCopyInto(out *SubmarineCommon) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineCommon.
func (in *SubmarineCommon) DeepCopy() *SubmarineCommon {
	if in == nil {
		return nil
	}
	out := new(SubmarineCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineDatabaseSpec) DeepCopyInto(out *SubmarineDatabaseSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineDatabaseSpec.
func (in *SubmarineDatabaseSpec) DeepCopy() *SubmarineDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineList) DeepCopyInto(out *SubmarineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Submarine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineList.
func (in *SubmarineList) DeepCopy() *SubmarineList {
	if in == nil {
		return nil
	}
	out := new(SubmarineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubmarineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineMinioSpec) DeepCopyInto(out *SubmarineMinioSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineMinioSpec.
func (in *SubmarineMinioSpec) DeepCopy() *SubmarineMinioSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineMinioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineMlflowSpec) DeepCopyInto(out *SubmarineMlflowSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineMlflowSpec.
func (in *SubmarineMlflowSpec) DeepCopy() *SubmarineMlflowSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineMlflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineServerSpec) DeepCopyInto(out *SubmarineServerSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineServerSpec.
func (in *SubmarineServerSpec) DeepCopy() *SubmarineServerSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineSpec) DeepCopyInto(out *SubmarineSpec) {
	*out = *in
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(SubmarineServerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(SubmarineDatabaseSpec)
		**out = **in
	}
	if in.Virtualservice != nil {
		in, out := &in.Virtualservice, &out.Virtualservice
		*out = new(SubmarineVirtualserviceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tensorboard != nil {
		in, out := &in.Tensorboard, &out.Tensorboard
		*out = new(SubmarineTensorboardSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Mlflow != nil {
		in, out := &in.Mlflow, &out.Mlflow
		*out = new(SubmarineMlflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Minio != nil {
		in, out := &in.Minio, &out.Minio
		*out = new(SubmarineMinioSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(SubmarineCommon)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineSpec.
func (in *SubmarineSpec) DeepCopy() *SubmarineSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineState) DeepCopyInto(out *SubmarineState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineState.
func (in *SubmarineState) DeepCopy() *SubmarineState {
	if in == nil {
		return nil
	}
	out := new(SubmarineState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineStatus) DeepCopyInto(out *SubmarineStatus) {
	*out = *in
	out.SubmarineState = in.SubmarineState
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineStatus.
func (in *SubmarineStatus) DeepCopy() *SubmarineStatus {
	if in == nil {
		return nil
	}
	out := new(SubmarineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineTensorboardSpec) DeepCopyInto(out *SubmarineTensorboardSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineTensorboardSpec.
func (in *SubmarineTensorboardSpec) DeepCopy() *SubmarineTensorboardSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineTensorboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmarineVirtualserviceSpec) DeepCopyInto(out *SubmarineVirtualserviceSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmarineVirtualserviceSpec.
func (in *SubmarineVirtualserviceSpec) DeepCopy() *SubmarineVirtualserviceSpec {
	if in == nil {
		return nil
	}
	out := new(SubmarineVirtualserviceSpec)
	in.DeepCopyInto(out)
	return out
}
