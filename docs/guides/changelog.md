---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.2.3 (unreleased)

- Add `sdwan_centralized_policy` resource and data source
- Add `sdwan_activate_centralized_policy` resource

## 0.2.2

- Add `sdwan_application_aware_routing_policy_definition` resource and data source
- Add `sdwan_cflowd_policy_definition` resource and data source
- Add `sdwan_traffic_data_policy_definition` resource and data source
- Add `set_parameters` attribute to `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Remove `dscp` and `next_hop` attributes of `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `layer2cos` attribute of `sdwan_rewrite_rule_policy_definition` resource and data source to `layer2_cos`
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_class_map_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_app_probe_class_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_mirror_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_policer_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_preferred_color_group_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_sla_class_policy_object` resource and data source
- Add `app_probe_class_version` attribute to `sdwan_sla_class_policy_object` resource and data source

## 0.2.1

- Add `sdwan_cisco_ospf_feature_template` resource and data source
- Add `sdwan_cisco_vpn_interface_ipsec_feature_template` resource and data source
- Add `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source
- Add `sdwan_hub_and_spoke_topology_policy_definition` resource and data source
- Add `sdwan_mesh_topology_policy_definition` resource and data source
- Add `sdwan_custom_control_topology_policy_definition` resource and data source
- Add `sdwan_vpn_membership_policy_definition` resource and data source

## 0.2.0

- BREAKING CHANGE: Completely revamped the provider based on `github.com/netascode/terraform-provider-sdwan` codebase, replacing all existing resources and data sources

## 0.1.0 (July 23, 2021)

- Initial Release

