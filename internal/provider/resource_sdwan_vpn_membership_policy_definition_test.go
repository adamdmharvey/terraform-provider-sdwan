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

func TestAccSdwanVPNMembershipPolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNMembershipPolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_vpn_membership_policy_definition.test", "name", "Example"),
					resource.TestCheckResourceAttr("sdwan_vpn_membership_policy_definition.test", "description", "My description"),
					resource.TestCheckResourceAttr("sdwan_vpn_membership_policy_definition.test", "sites.0.site_list_id", "e858e1c4-6aa8-4de7-99df-c3adbf80290d"),
				),
			},
		},
	})
}

const testAccSdwanVPNMembershipPolicyDefinitionConfig = `


resource "sdwan_vpn_membership_policy_definition" "test" {
	name = "Example"
	description = "My description"
	sites = [{
		site_list_id = "e858e1c4-6aa8-4de7-99df-c3adbf80290d"
		vpn_list_ids = ["04fcbb0b-efbf-43d2-a04b-847d3a7b104e"]
	}]
}
`
