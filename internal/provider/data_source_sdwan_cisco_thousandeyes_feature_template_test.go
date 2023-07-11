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

func TestAccDataSourceSdwanCiscoThousandEyesFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoThousandEyesFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.instance_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.application_type", "te"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_account_group_token", "1234567"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_vpn", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_agent_ip", "1.1.1.2/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_default_gateway", "1.1.1.255"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_name_server", "10.2.2.2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_hostname", "agent1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_web_proxy_type", "static"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_proxy_host", "3.3.3.3"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_thousandeyes_feature_template.test", "virtual_applications.0.te_proxy_port", "80"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoThousandEyesFeatureTemplateConfig = `

resource "sdwan_cisco_thousandeyes_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  virtual_applications = [{
    instance_id = "1"
    application_type = "te"
    te_account_group_token = "1234567"
    te_vpn = 1
    te_agent_ip = "1.1.1.2/24"
    te_default_gateway = "1.1.1.255"
    te_name_server = "10.2.2.2"
    te_hostname = "agent1"
    te_web_proxy_type = "static"
    te_proxy_host = "3.3.3.3"
    te_proxy_port = 80
  }]
}

data "sdwan_cisco_thousandeyes_feature_template" "test" {
  id = sdwan_cisco_thousandeyes_feature_template.test.id
}
`
