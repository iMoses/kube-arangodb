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
// Author Ewout Prangsma
//

package v1alpha

import "github.com/pkg/errors"

// DeploymentState is a strongly typed state of a deployment
type DeploymentState string

const (
	// DeploymentStateNone indicates that the state is not set yet
	DeploymentStateNone DeploymentState = ""
	// DeploymentStateCreating indicates that the deployment is being created
	DeploymentStateCreating DeploymentState = "Creating"
	// DeploymentStateRunning indicates that all servers are running and reachable
	DeploymentStateRunning DeploymentState = "Running"
	// DeploymentStateScaling indicates that servers are being added or removed
	DeploymentStateScaling DeploymentState = "Scaling"
	// DeploymentStateUpgrading indicates that a version upgrade is in progress
	DeploymentStateUpgrading DeploymentState = "Upgrading"
	// DeploymentStateFailed indicates that a deployment is in a failed state
	// from which automatic recovery is impossible. Inspect `Reason` for more info.
	DeploymentStateFailed DeploymentState = "Failed"
)

// IsFailed returns true if given state is DeploymentStateFailed
func (cs DeploymentState) IsFailed() bool {
	return cs == DeploymentStateFailed
}

// DeploymentStatus contains the status part of a Cluster resource.
type DeploymentStatus struct {
	// State holds the current state of the deployment
	State DeploymentState `json:"state"`
	// Reason contains a human readable reason for reaching the current state (can be empty)
	Reason string `json:"reason,omitempty"` // Reason for current state

	// ServiceName holds the name of the Service a client can use (inside the k8s cluster)
	// to access ArangoDB.
	ServiceName string `json:"serviceName,omitempty"`
	// SyncServiceName holds the name of the Service a client can use (inside the k8s cluster)
	// to access syncmasters (only set when dc2dc synchronization is enabled).
	SyncServiceName string `json:"syncServiceName,omitempty"`

	// Members holds the status for all members in all server groups
	Members DeploymentStatusMembers `json:"members"`
}

// DeploymentStatusMembers holds the member status of all server groups
type DeploymentStatusMembers struct {
	Single       MemberStatusList `json:"single,omitempty"`
	Agents       MemberStatusList `json:"agents,omitempty"`
	DBServers    MemberStatusList `json:"dbservers,omitempty"`
	Coordinators MemberStatusList `json:"coordinators,omitempty"`
	SyncMasters  MemberStatusList `json:"syncmasters,omitempty"`
	SyncWorkers  MemberStatusList `json:"syncworkers,omitempty"`
}

// ContainsID returns true if the given set of members contains a member with given ID.
func (ds DeploymentStatusMembers) ContainsID(id string) bool {
	return ds.Single.ContainsID(id) ||
		ds.Agents.ContainsID(id) ||
		ds.DBServers.ContainsID(id) ||
		ds.Coordinators.ContainsID(id) ||
		ds.SyncMasters.ContainsID(id) ||
		ds.SyncWorkers.ContainsID(id)
}

// MemberStatusList is a list of MemberStatus entries
type MemberStatusList []MemberStatus

// ContainsID returns true if the given list contains a member with given ID.
func (l MemberStatusList) ContainsID(id string) bool {
	for _, x := range l {
		if x.ID == id {
			return true
		}
	}
	return false
}

// Add a member to the list.
// Returns an AlreadyExistsError if the ID of the given member already exists.
func (l *MemberStatusList) Add(m MemberStatus) error {
	src := *l
	for _, x := range src {
		if x.ID == m.ID {
			return maskAny(errors.Wrapf(AlreadyExistsError, "Member '%s' already exists", m.ID))
		}
	}
	*l = append(src, m)
	return nil
}

// Update a member in the list.
// Returns a NotFoundError if the ID of the given member cannot be found.
func (l MemberStatusList) Update(m MemberStatus) error {
	for i, x := range l {
		if x.ID == m.ID {
			l[i] = m
			return nil
		}
	}
	return maskAny(errors.Wrapf(NotFoundError, "Member '%s' is not a member", m.ID))
}

// RemoveByID a member with given ID from the list.
// Returns a NotFoundError if the ID of the given member cannot be found.
func (l *MemberStatusList) RemoveByID(id string) error {
	src := *l
	for i, x := range src {
		if x.ID == id {
			*l = append(src[:i], src[i+1:]...)
			return nil
		}
	}
	return maskAny(errors.Wrapf(NotFoundError, "Member '%s' is not a member", id))
}

// MemberState is a strongly typed state of a deployment member
type MemberState string

const (
	// MemberStateNone indicates that the state is not set yet
	MemberStateNone MemberState = ""
	// MemberStateCreating indicates that the member is in the process of being created
	MemberStateCreating MemberState = "Creating"
	// MemberStateReady indicates that the member is running and reachable
	MemberStateReady MemberState = "Ready"
	// MemberStateCleanOut indicates that a dbserver is in the process of being cleaned out
	MemberStateCleanOut MemberState = "CleanOut"
	// MemberStateShuttingDown indicates that a member is shutting down
	MemberStateShuttingDown MemberState = "ShuttingDown"
)

// MemberStatus holds the current status of a single member (server)
type MemberStatus struct {
	// ID holds the unique ID of the member.
	// This id is also used within the ArangoDB cluster to identify this server.
	ID string `json:"id"`
	// State holds the current state of this member
	State MemberState `json:"state"`
	// PersistentVolumeClaimName holds the name of the persistent volume claim used for this member (if any).
	PersistentVolumeClaimName string `json:"persistentVolumeClaimName,omitempty"`
	// PodName holds the name of the Pod that currently runs this member
	PodName string `json:"podName,omitempty"`
}