// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanRewriteRulePolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanRewriteRulePolicyDefinitionConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_rewrite_rule_policy_definition.test", "rules.0.priority", "low"),
					resource.TestCheckResourceAttr("sdwan_rewrite_rule_policy_definition.test", "rules.0.dscp", "16"),
					resource.TestCheckResourceAttr("sdwan_rewrite_rule_policy_definition.test", "rules.0.layer2cos", "1"),
				),
			},
		},
	})
}

func testAccSdwanRewriteRulePolicyDefinitionConfig_all() string {
	return `
	resource "sdwan_class_map_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			queue = 6
		}]
	}

	resource "sdwan_rewrite_rule_policy_definition" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		rules = [{
			class_map_id = sdwan_class_map_policy_object.test.id
			priority = "low"
			dscp = 16
			layer2cos = 1
		}]
	}
	`
}
