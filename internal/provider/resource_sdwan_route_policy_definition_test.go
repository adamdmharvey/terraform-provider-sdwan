// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanRoutePolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanRoutePolicyDefinitionConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "default_action", "reject"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.id", "10"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.ip_type", "ipv4"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.name", "Sequence 10"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.base_action", "accept"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.match_entries.0.type", "metric"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.match_entries.0.metric", "100"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.type", "aggregator"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator", "10"),
					resource.TestCheckResourceAttr("sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator_ip_address", "10.1.2.3"),
				),
			},
		},
	})
}

func testAccSdwanRoutePolicyDefinitionConfig_all() string {
	return `
	resource "sdwan_route_policy_definition" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		default_action = "reject"
		sequences = [{
			id = 10
			ip_type = "ipv4"
			name = "Sequence 10"
			base_action = "accept"
			match_entries = [{
				type = "metric"
				metric = 100
			}]
			action_entries = [{
				type = "aggregator"
				aggregator = 10
				aggregator_ip_address = "10.1.2.3"
			}]
		}]
	}
	`
}
