---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_traffic_data_policy_definition Resource - terraform-provider-sdwan"
subcategory: "Centralized Policies"
description: |-
  This resource can manage a Traffic Data Policy Definition .
---

# sdwan_traffic_data_policy_definition (Resource)

This resource can manage a Traffic Data Policy Definition .

## Example Usage

```terraform
resource "sdwan_traffic_data_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id      = 1
      name    = "Seq1"
      type    = "applicationFirewall"
      ip_type = "ipv4"
      match_entries = [
        {
          type                = "appList"
          application_list_id = "e3aad846-abb9-425f-aaa8-9ed17b9c8d7c"
        }
      ]
      action_entries = [
        {
          type = "log"
          log  = true
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition
- `sequences` (Attributes List) List of sequences (see [below for nested schema](#nestedatt--sequences))

### Optional

- `default_action` (String) Default action, either `accept` or `drop`
  - Choices: `accept`, `drop`

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

<a id="nestedatt--sequences"></a>
### Nested Schema for `sequences`

Required:

- `id` (Number) Sequence ID
- `name` (String) Sequence name
- `type` (String) Sequence type
  - Choices: `applicationFirewall`, `qos`, `serviceChaining`, `trafficEngineering`, `data`

Optional:

- `action_entries` (Attributes List) List of action entries (see [below for nested schema](#nestedatt--sequences--action_entries))
- `ip_type` (String) Sequence IP type, either `ipv4`, `ipv6` or `all`
  - Choices: `ipv4`, `ipv6`, `all`
- `match_entries` (Attributes List) List of match entries (see [below for nested schema](#nestedatt--sequences--match_entries))

<a id="nestedatt--sequences--action_entries"></a>
### Nested Schema for `sequences.action_entries`

Required:

- `type` (String) Type of action entry
  - Choices: `cflowd`, `count`, `dreOptimization`, `fallbackToRouting`, `log`, `lossProtect`, `lossProtectFec`, `nat`, `redirectDns`, `serviceNodeGroup`, `set`, `sig`, `tcpOptimization`

Optional:

- `cflowd` (Boolean) Enable cflowd
- `counter` (String) Counter name
- `dre_optimization` (Boolean) Enable DRE optimization
- `fallback_to_routing` (Boolean) Enable fallback to routing
- `log` (Boolean) Enable logging
- `loss_correction` (String) Loss correction
  - Choices: `fecAdaptive`, `fecAlways`, `packetDuplication`
- `loss_correction_fec` (String) Loss correction FEC
  - Choices: `fecAdaptive`, `fecAlways`, `packetDuplication`
- `loss_correction_fec_threshold` (Number) Loss correction FEC threshold
  - Range: `1`-`5`
- `nat_parameters` (Attributes List) List of NAT parameters (see [below for nested schema](#nestedatt--sequences--action_entries--nat_parameters))
- `nat_pool` (String) NAT pool
  - Choices: `pool`
- `nat_pool_id` (Number) NAT pool ID
  - Range: `1`-`31`
- `redirect_dns` (String) Redirect DNS
  - Choices: `dnsType`, `ipAddress`
- `redirect_dns_address` (String) Redirect DNS IP address
- `redirect_dns_type` (String) Redirect DNS type
  - Choices: `host`, `umbrella`
- `secure_internet_gateway` (Boolean) Enable secure internet gateway
- `service_node_group` (String) Service node group
- `set_parameters` (Attributes List) List of set parameters (see [below for nested schema](#nestedatt--sequences--action_entries--set_parameters))
- `tcp_optimization` (Boolean) Enable TCP optimization

<a id="nestedatt--sequences--action_entries--nat_parameters"></a>
### Nested Schema for `sequences.action_entries.nat_parameters`

Required:

- `type` (String) Type of NAT parameter
  - Choices: `useVpn`, `fallback`

Optional:

- `fallback` (Boolean) Fallback
- `vpn_id` (Number) DSCP


<a id="nestedatt--sequences--action_entries--set_parameters"></a>
### Nested Schema for `sequences.action_entries.set_parameters`

Required:

- `type` (String) Type of set parameter
  - Choices: `dscp`, `forwardingClass`, `localTlocList`, `nextHop`, `nextHopLoose`, `policer`, `preferredColorGroup`, `tlocList`, `tloc`, `service`, `vpn`

Optional:

- `dscp` (Number) DSCP
  - Range: `0`-`63`
- `forwarding_class` (String) Forwarding class
- `local_tloc_list_color` (String) Local TLOC list color
- `local_tloc_list_encap` (String) Local TLOC list encapsulation
  - Choices: `ipsec`, `gre`
- `local_tloc_list_restrict` (Boolean) Local TLOC list restrict
- `next_hop` (String) Next hop IP
- `next_hop_loose` (Boolean) Use routing table entry to forward the packet in case Next-hop is not available
- `policer_list_id` (String) Policer list ID
- `policer_list_version` (Number) Policer list version
- `preferred_color_group_list` (String) Preferred color group list ID
- `preferred_color_group_list_version` (Number) Preferred color group list version
- `service_tloc_color` (String) Service TLOC color
- `service_tloc_encapsulation` (String) Service TLOC encapsulation
  - Choices: `ipsec`, `gre`
- `service_tloc_ip` (String) Service TLOC IP address
- `service_tloc_list_id` (String) Service TLOC list ID
- `service_tloc_list_version` (Number) Service TLOC list version
- `service_type` (String) Service type
  - Choices: `FW`, `IDP`, `IDS`, `netsvc1`, `netsvc2`, `netsvc3`, `netsvc4`, `netsvc5`
- `service_vpn_id` (Number) Service VPN ID
  - Range: `0`-`65536`
- `tloc_color` (String) TLOC color
- `tloc_encapsulation` (String) TLOC encapsulation
  - Choices: `ipsec`, `gre`
- `tloc_ip` (String) TLOC IP address
- `tloc_list_id` (String) TLOC list ID
- `tloc_list_version` (Number) TLOC list version
- `vpn_id` (Number) DSCP
  - Range: `0`-`65530`



<a id="nestedatt--sequences--match_entries"></a>
### Nested Schema for `sequences.match_entries`

Required:

- `type` (String) Type of match entry
  - Choices: `appList`, `dnsAppList`, `dns`, `dscp`, `packetLength`, `plp`, `protocol`, `sourceDataPrefixList`, `sourceIp`, `sourcePort`, `destinationDataPrefixList`, `destinationIp`, `destinationRegion`, `destinationPort`, `tcp`, `trafficTo`

Optional:

- `application_list_id` (String) Application list ID
- `application_list_version` (Number) Application list version
- `destination_data_prefix_list_id` (String) Destination Data Prefix list ID
- `destination_data_prefix_list_version` (Number) Destination Data Prefix list version
- `destination_ip` (String) Destination IP
- `destination_port` (Number) Destination port
  - Range: `0`-`65535`
- `destination_region` (String) Destination region
  - Choices: `primary-region`, `secondary-region`, `other-region`
- `dns` (String) DNS request or response
  - Choices: `request`, `response`
- `dns_application_list_id` (String) DNS Application list ID
- `dns_application_list_version` (Number) DNS Application list version
- `dscp` (Number) DSCP value
  - Range: `0`-`63`
- `packet_length` (Number) Packet length
  - Range: `0`-`65535`
- `plp` (String) PLP
  - Choices: `low`, `high`
- `protocol` (Number) IP Protocol
  - Range: `0`-`255`
- `source_data_prefix_list_id` (String) Source Data Prefix list ID
- `source_data_prefix_list_version` (Number) Source Data Prefix list version
- `source_ip` (String) Source IP
- `source_port` (Number) Source port
  - Range: `0`-`65535`
- `tcp` (String) TCP flags
  - Choices: `syn`
- `traffic_to` (String) Traffic to
  - Choices: `access`, `core`, `service`

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_traffic_data_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
