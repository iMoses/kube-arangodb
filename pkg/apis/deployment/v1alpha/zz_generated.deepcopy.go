// +build !ignore_autogenerated

//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha

import (
	time "time"

	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	in.CreationTime.DeepCopyInto(&out.CreationTime)
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeployment) DeepCopyInto(out *ArangoDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeployment.
func (in *ArangoDeployment) DeepCopy() *ArangoDeployment {
	if in == nil {
		return nil
	}
	out := new(ArangoDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeploymentList) DeepCopyInto(out *ArangoDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeploymentList.
func (in *ArangoDeploymentList) DeepCopy() *ArangoDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ArangoDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.JWTSecretName != nil {
		in, out := &in.JWTSecretName, &out.JWTSecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
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
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentMode)
			**out = **in
		}
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		if *in == nil {
			*out = nil
		} else {
			*out = new(Environment)
			**out = **in
		}
	}
	if in.StorageEngine != nil {
		in, out := &in.StorageEngine, &out.StorageEngine
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageEngine)
			**out = **in
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PullPolicy)
			**out = **in
		}
	}
	in.RocksDB.DeepCopyInto(&out.RocksDB)
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	in.Sync.DeepCopyInto(&out.Sync)
	in.Single.DeepCopyInto(&out.Single)
	in.Agents.DeepCopyInto(&out.Agents)
	in.DBServers.DeepCopyInto(&out.DBServers)
	in.Coordinators.DeepCopyInto(&out.Coordinators)
	in.SyncMasters.DeepCopyInto(&out.SyncMasters)
	in.SyncWorkers.DeepCopyInto(&out.SyncWorkers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(ImageInfoList, len(*in))
		copy(*out, *in)
	}
	in.Members.DeepCopyInto(&out.Members)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = make(Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AcceptedSpec != nil {
		in, out := &in.AcceptedSpec, &out.AcceptedSpec
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeploymentSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecretHashes != nil {
		in, out := &in.SecretHashes, &out.SecretHashes
		if *in == nil {
			*out = nil
		} else {
			*out = new(SecretHashes)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatusMembers) DeepCopyInto(out *DeploymentStatusMembers) {
	*out = *in
	if in.Single != nil {
		in, out := &in.Single, &out.Single
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DBServers != nil {
		in, out := &in.DBServers, &out.DBServers
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncMasters != nil {
		in, out := &in.SyncMasters, &out.SyncMasters
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncWorkers != nil {
		in, out := &in.SyncWorkers, &out.SyncWorkers
		*out = make(MemberStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatusMembers.
func (in *DeploymentStatusMembers) DeepCopy() *DeploymentStatusMembers {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatusMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageInfo) DeepCopyInto(out *ImageInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageInfo.
func (in *ImageInfo) DeepCopy() *ImageInfo {
	if in == nil {
		return nil
	}
	out := new(ImageInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberStatus) DeepCopyInto(out *MemberStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatus.
func (in *MemberStatus) DeepCopy() *MemberStatus {
	if in == nil {
		return nil
	}
	out := new(MemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.TokenSecretName != nil {
		in, out := &in.TokenSecretName, &out.TokenSecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocksDBEncryptionSpec) DeepCopyInto(out *RocksDBEncryptionSpec) {
	*out = *in
	if in.KeySecretName != nil {
		in, out := &in.KeySecretName, &out.KeySecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocksDBEncryptionSpec.
func (in *RocksDBEncryptionSpec) DeepCopy() *RocksDBEncryptionSpec {
	if in == nil {
		return nil
	}
	out := new(RocksDBEncryptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocksDBSpec) DeepCopyInto(out *RocksDBSpec) {
	*out = *in
	in.Encryption.DeepCopyInto(&out.Encryption)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocksDBSpec.
func (in *RocksDBSpec) DeepCopy() *RocksDBSpec {
	if in == nil {
		return nil
	}
	out := new(RocksDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretHashes) DeepCopyInto(out *SecretHashes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretHashes.
func (in *SecretHashes) DeepCopy() *SecretHashes {
	if in == nil {
		return nil
	}
	out := new(SecretHashes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerGroupSpec) DeepCopyInto(out *ServerGroupSpec) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerGroupSpec.
func (in *ServerGroupSpec) DeepCopy() *ServerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ServerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncSpec) DeepCopyInto(out *SyncSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PullPolicy)
			**out = **in
		}
	}
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncSpec.
func (in *SyncSpec) DeepCopy() *SyncSpec {
	if in == nil {
		return nil
	}
	out := new(SyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
	if in.CASecretName != nil {
		in, out := &in.CASecretName, &out.CASecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.AltNames != nil {
		in, out := &in.AltNames, &out.AltNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		if *in == nil {
			*out = nil
		} else {
			*out = new(time.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}
