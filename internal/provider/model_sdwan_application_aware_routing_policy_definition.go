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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ApplicationAwareRoutingPolicyDefinition struct {
	Id          types.String                                       `tfsdk:"id"`
	Version     types.Int64                                        `tfsdk:"version"`
	Name        types.String                                       `tfsdk:"name"`
	Description types.String                                       `tfsdk:"description"`
	Sequences   []ApplicationAwareRoutingPolicyDefinitionSequences `tfsdk:"sequences"`
}

type ApplicationAwareRoutingPolicyDefinitionSequences struct {
	Id            types.Int64                                                     `tfsdk:"id"`
	Name          types.String                                                    `tfsdk:"name"`
	IpType        types.String                                                    `tfsdk:"ip_type"`
	MatchEntries  []ApplicationAwareRoutingPolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []ApplicationAwareRoutingPolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type ApplicationAwareRoutingPolicyDefinitionSequencesMatchEntries struct {
	Type                             types.String `tfsdk:"type"`
	ApplicationListId                types.String `tfsdk:"application_list_id"`
	ApplicationListVersion           types.Int64  `tfsdk:"application_list_version"`
	DnsApplicationListId             types.String `tfsdk:"dns_application_list_id"`
	DnsApplicationListVersion        types.Int64  `tfsdk:"dns_application_list_version"`
	Dns                              types.String `tfsdk:"dns"`
	Dscp                             types.Int64  `tfsdk:"dscp"`
	Plp                              types.String `tfsdk:"plp"`
	Protocol                         types.Int64  `tfsdk:"protocol"`
	SourceDataPrefixListId           types.String `tfsdk:"source_data_prefix_list_id"`
	SourceDataPrefixListVersion      types.Int64  `tfsdk:"source_data_prefix_list_version"`
	SourceIp                         types.String `tfsdk:"source_ip"`
	SourcePort                       types.Int64  `tfsdk:"source_port"`
	DestinationDataPrefixListId      types.String `tfsdk:"destination_data_prefix_list_id"`
	DestinationDataPrefixListVersion types.Int64  `tfsdk:"destination_data_prefix_list_version"`
	DestinationIp                    types.String `tfsdk:"destination_ip"`
	DestinationPort                  types.Int64  `tfsdk:"destination_port"`
	DestinationRegion                types.String `tfsdk:"destination_region"`
	TrafficTo                        types.String `tfsdk:"traffic_to"`
}
type ApplicationAwareRoutingPolicyDefinitionSequencesActionEntries struct {
	Type                    types.String                                                                      `tfsdk:"type"`
	BackupSlaPreferredColor types.String                                                                      `tfsdk:"backup_sla_preferred_color"`
	Counter                 types.String                                                                      `tfsdk:"counter"`
	Log                     types.Bool                                                                        `tfsdk:"log"`
	CloudSla                types.Bool                                                                        `tfsdk:"cloud_sla"`
	SlaClassParameters      []ApplicationAwareRoutingPolicyDefinitionSequencesActionEntriesSlaClassParameters `tfsdk:"sla_class_parameters"`
}

type ApplicationAwareRoutingPolicyDefinitionSequencesActionEntriesSlaClassParameters struct {
	Type                           types.String `tfsdk:"type"`
	SlaClassList                   types.String `tfsdk:"sla_class_list"`
	SlaClassListVersion            types.Int64  `tfsdk:"sla_class_list_version"`
	PreferredColorGroupList        types.String `tfsdk:"preferred_color_group_list"`
	PreferredColorGroupListVersion types.Int64  `tfsdk:"preferred_color_group_list_version"`
	PreferredColor                 types.String `tfsdk:"preferred_color"`
}

func (data ApplicationAwareRoutingPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "type", "appRoute")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Sequences) > 0 {
		body, _ = sjson.Set(body, "sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "sequenceType", "appRoute")
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if len(item.MatchEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.ApplicationListId.IsNull() && childItem.Type.ValueString() == "appList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ApplicationListId.ValueString())
					}
					if !childItem.DnsApplicationListId.IsNull() && childItem.Type.ValueString() == "dnsAppList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DnsApplicationListId.ValueString())
					}
					if !childItem.Dns.IsNull() && childItem.Type.ValueString() == "dns" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Dns.ValueString())
					}
					if !childItem.Dscp.IsNull() && childItem.Type.ValueString() == "dscp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Dscp.ValueInt64()))
					}
					if !childItem.Plp.IsNull() && childItem.Type.ValueString() == "plp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Plp.ValueString())
					}
					if !childItem.Protocol.IsNull() && childItem.Type.ValueString() == "protocol" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Protocol.ValueInt64()))
					}
					if !childItem.SourceDataPrefixListId.IsNull() && childItem.Type.ValueString() == "sourceDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.SourceDataPrefixListId.ValueString())
					}
					if !childItem.SourceIp.IsNull() && childItem.Type.ValueString() == "sourceIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourceIp.ValueString())
					}
					if !childItem.SourcePort.IsNull() && childItem.Type.ValueString() == "sourcePort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.SourcePort.ValueInt64()))
					}
					if !childItem.DestinationDataPrefixListId.IsNull() && childItem.Type.ValueString() == "destinationDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DestinationDataPrefixListId.ValueString())
					}
					if !childItem.DestinationIp.IsNull() && childItem.Type.ValueString() == "destinationIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationIp.ValueString())
					}
					if !childItem.DestinationPort.IsNull() && childItem.Type.ValueString() == "destinationPort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.DestinationPort.ValueInt64()))
					}
					if !childItem.DestinationRegion.IsNull() && childItem.Type.ValueString() == "destinationRegion" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationRegion.ValueString())
					}
					if !childItem.TrafficTo.IsNull() && childItem.Type.ValueString() == "trafficTo" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.TrafficTo.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if len(item.ActionEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.BackupSlaPreferredColor.IsNull() && childItem.Type.ValueString() == "backupSlaPreferredColor" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.BackupSlaPreferredColor.ValueString())
					}
					if !childItem.Counter.IsNull() && childItem.Type.ValueString() == "count" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.Counter.ValueString())
					}
					if !childItem.Log.IsNull() && childItem.Type.ValueString() == "log" {
						if false && childItem.Log.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.Log.ValueBool())
						}
					}
					if !childItem.CloudSla.IsNull() && childItem.Type.ValueString() == "cloudSaas" {
						if true && childItem.CloudSla.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", "")
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.CloudSla.ValueBool())
						}
					}
					if len(childItem.SlaClassParameters) > 0 && childItem.Type.ValueString() == "slaClass" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", []interface{}{})
						for _, childChildItem := range childItem.SlaClassParameters {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "field", childChildItem.Type.ValueString())
							}
							if !childChildItem.SlaClassList.IsNull() && childChildItem.Type.ValueString() == "name" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.SlaClassList.ValueString())
							}
							if !childChildItem.PreferredColorGroupList.IsNull() && childChildItem.Type.ValueString() == "preferredColorGroup" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.PreferredColorGroupList.ValueString())
							}
							if !childChildItem.PreferredColor.IsNull() && childChildItem.Type.ValueString() == "preferredColor" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.PreferredColor.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "parameter.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

func (data *ApplicationAwareRoutingPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("sequences"); value.Exists() {
		data.Sequences = make([]ApplicationAwareRoutingPolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationAwareRoutingPolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() {
				item.MatchEntries = make([]ApplicationAwareRoutingPolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationAwareRoutingPolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "appList" {
						cItem.ApplicationListId = types.StringValue(ccValue.String())
					} else {
						cItem.ApplicationListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "dnsAppList" {
						cItem.DnsApplicationListId = types.StringValue(ccValue.String())
					} else {
						cItem.DnsApplicationListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "dns" {
						cItem.Dns = types.StringValue(ccValue.String())
					} else {
						cItem.Dns = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "dscp" {
						cItem.Dscp = types.Int64Value(ccValue.Int())
					} else {
						cItem.Dscp = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "plp" {
						cItem.Plp = types.StringValue(ccValue.String())
					} else {
						cItem.Plp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "protocol" {
						cItem.Protocol = types.Int64Value(ccValue.Int())
					} else {
						cItem.Protocol = types.Int64Null()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "sourceDataPrefixList" {
						cItem.SourceDataPrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.SourceDataPrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourceIp" {
						cItem.SourceIp = types.StringValue(ccValue.String())
					} else {
						cItem.SourceIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourcePort" {
						cItem.SourcePort = types.Int64Value(ccValue.Int())
					} else {
						cItem.SourcePort = types.Int64Null()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "destinationDataPrefixList" {
						cItem.DestinationDataPrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationDataPrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationIp" {
						cItem.DestinationIp = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationPort" {
						cItem.DestinationPort = types.Int64Value(ccValue.Int())
					} else {
						cItem.DestinationPort = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationRegion" {
						cItem.DestinationRegion = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationRegion = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "trafficTo" {
						cItem.TrafficTo = types.StringValue(ccValue.String())
					} else {
						cItem.TrafficTo = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() {
				item.ActionEntries = make([]ApplicationAwareRoutingPolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ApplicationAwareRoutingPolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "backupSlaPreferredColor" {
						cItem.BackupSlaPreferredColor = types.StringValue(ccValue.String())
					} else {
						cItem.BackupSlaPreferredColor = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "count" {
						cItem.Counter = types.StringValue(ccValue.String())
					} else {
						cItem.Counter = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "log" {
						if false && ccValue.String() == "" {
							cItem.Log = types.BoolValue(true)
						} else {
							cItem.Log = types.BoolValue(ccValue.Bool())
						}
					} else {
						cItem.Log = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "cloudSaas" {
						if true && ccValue.String() == "" {
							cItem.CloudSla = types.BoolValue(true)
						} else {
							cItem.CloudSla = types.BoolValue(ccValue.Bool())
						}
					} else {
						cItem.CloudSla = types.BoolNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "slaClass" {
						cItem.SlaClassParameters = make([]ApplicationAwareRoutingPolicyDefinitionSequencesActionEntriesSlaClassParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ApplicationAwareRoutingPolicyDefinitionSequencesActionEntriesSlaClassParameters{}
							if cccValue := ccv.Get("field"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "name" {
								ccItem.SlaClassList = types.StringValue(cccValue.String())
							} else {
								ccItem.SlaClassList = types.StringNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "preferredColorGroup" {
								ccItem.PreferredColorGroupList = types.StringValue(cccValue.String())
							} else {
								ccItem.PreferredColorGroupList = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "preferredColor" {
								ccItem.PreferredColor = types.StringValue(cccValue.String())
							} else {
								ccItem.PreferredColor = types.StringNull()
							}
							cItem.SlaClassParameters = append(cItem.SlaClassParameters, ccItem)
							return true
						})
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
	data.updateVersions(ctx, &state)
}

func (data *ApplicationAwareRoutingPolicyDefinition) hasChanges(ctx context.Context, state *ApplicationAwareRoutingPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].IpType.Equal(state.Sequences[i].IpType) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ApplicationListId.Equal(state.Sequences[i].MatchEntries[ii].ApplicationListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DnsApplicationListId.Equal(state.Sequences[i].MatchEntries[ii].DnsApplicationListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Dns.Equal(state.Sequences[i].MatchEntries[ii].Dns) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Dscp.Equal(state.Sequences[i].MatchEntries[ii].Dscp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Plp.Equal(state.Sequences[i].MatchEntries[ii].Plp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Protocol.Equal(state.Sequences[i].MatchEntries[ii].Protocol) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceDataPrefixListId.Equal(state.Sequences[i].MatchEntries[ii].SourceDataPrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceIp.Equal(state.Sequences[i].MatchEntries[ii].SourceIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourcePort.Equal(state.Sequences[i].MatchEntries[ii].SourcePort) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListId.Equal(state.Sequences[i].MatchEntries[ii].DestinationDataPrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationIp.Equal(state.Sequences[i].MatchEntries[ii].DestinationIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationPort.Equal(state.Sequences[i].MatchEntries[ii].DestinationPort) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationRegion.Equal(state.Sequences[i].MatchEntries[ii].DestinationRegion) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TrafficTo.Equal(state.Sequences[i].MatchEntries[ii].TrafficTo) {
						hasChanges = true
					}
				}
			}
			if len(data.Sequences[i].ActionEntries) != len(state.Sequences[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].ActionEntries {
					if !data.Sequences[i].ActionEntries[ii].Type.Equal(state.Sequences[i].ActionEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].BackupSlaPreferredColor.Equal(state.Sequences[i].ActionEntries[ii].BackupSlaPreferredColor) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Counter.Equal(state.Sequences[i].ActionEntries[ii].Counter) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Log.Equal(state.Sequences[i].ActionEntries[ii].Log) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CloudSla.Equal(state.Sequences[i].ActionEntries[ii].CloudSla) {
						hasChanges = true
					}
					if len(data.Sequences[i].ActionEntries[ii].SlaClassParameters) != len(state.Sequences[i].ActionEntries[ii].SlaClassParameters) {
						hasChanges = true
					} else {
						for iii := range data.Sequences[i].ActionEntries[ii].SlaClassParameters {
							if !data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].Type.Equal(state.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].Type) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].SlaClassList.Equal(state.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].SlaClassList) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColorGroupList.Equal(state.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColorGroupList) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColor.Equal(state.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColor) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

func (data *ApplicationAwareRoutingPolicyDefinition) updateVersions(ctx context.Context, state *ApplicationAwareRoutingPolicyDefinition) {
	for i := range data.Sequences {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].Id.ValueInt64()), fmt.Sprintf("%v", data.Sequences[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Sequences {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[j].Id.ValueInt64()), fmt.Sprintf("%v", state.Sequences[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Sequences[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ApplicationListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ApplicationListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ApplicationListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DnsApplicationListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DnsApplicationListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DnsApplicationListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].SourceDataPrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].SourceDataPrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].SourceDataPrefixListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DestinationDataPrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DestinationDataPrefixListVersion = types.Int64Null()
			}
		}
		for ii := range data.Sequences[i].ActionEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].ActionEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			for iii := range data.Sequences[i].ActionEntries[ii].SlaClassParameters {
				ccDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].Type.ValueString())}
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.Sequences[stateIndex].ActionEntries[cStateIndex].SlaClassParameters {
						ccStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[cStateIndex].SlaClassParameters[jjj].Type.ValueString())}
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].SlaClassListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SlaClassParameters[ccStateIndex].SlaClassListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].SlaClassListVersion = types.Int64Null()
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColorGroupListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SlaClassParameters[ccStateIndex].PreferredColorGroupListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SlaClassParameters[iii].PreferredColorGroupListVersion = types.Int64Null()
				}
			}
		}
	}
}
