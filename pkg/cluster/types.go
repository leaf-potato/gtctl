// Copyright 2023 Greptime Team
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

package cluster

import (
	"context"

	greptimedbclusterv1alpha1 "github.com/GreptimeTeam/greptimedb-operator/apis/v1alpha1"
	"github.com/olekukonko/tablewriter"
)

type Operations interface {
	// Get gets the current cluster profile.
	Get(ctx context.Context, options *GetOptions) error

	// List lists all cluster profiles.
	List(ctx context.Context, options *ListOptions) error

	// Scale scales the current cluster according to NewReplicas in ScaleOptions,
	// and refill the OldReplicas in ScaleOptions.
	Scale(ctx context.Context, options *ScaleOptions) error

	// TODO(sh2): implementing Create and Delete
}

type GetOptions struct {
	Namespace string
	Name      string

	// Table view render.
	Table *tablewriter.Table
}

type ListOptions struct {
	GetOptions
}

type ScaleOptions struct {
	NewReplicas   int32
	OldReplicas   int32
	Namespace     string
	Name          string
	ComponentType greptimedbclusterv1alpha1.ComponentKind
}