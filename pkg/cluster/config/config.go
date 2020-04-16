// Copyright 2020 The Lokomotive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config handles go structs for loading an HCL configuration.
package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/kinvolk/lokomotive/pkg/backend"
	"github.com/kinvolk/lokomotive/pkg/components"
)

// LokomotiveConfig struct defines the user provide configuration
type LokomotiveConfig struct {
	Cluster    *ClusterConfig
	Platform   Platform
	Metadata   *Metadata
	Controller *ControllerConfig
	Backend    backend.Backend
	Flatcar    *FlatcarConfig
	Network    *NetworkConfig
	Components map[string]components.Component
}

// Validate validates the user provided configuration
func (lc *LokomotiveConfig) Validate() hcl.Diagnostics {
	var diagnostics hcl.Diagnostics

	if lc.Platform != nil {
		diagnostics = append(diagnostics, lc.Cluster.Validate()...)
		diagnostics = append(diagnostics, lc.Platform.Validate()...)
		diagnostics = append(diagnostics, lc.Flatcar.Validate()...)
		diagnostics = append(diagnostics, lc.Controller.Validate()...)
		diagnostics = append(diagnostics, lc.Metadata.Validate()...)
		diagnostics = append(diagnostics, lc.Network.Validate()...)
	}

	return diagnostics
}
