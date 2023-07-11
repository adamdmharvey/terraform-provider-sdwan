// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Route struct {
	Id            types.String     `tfsdk:"id"`
	Version       types.Int64      `tfsdk:"version"`
	Type          types.String     `tfsdk:"type"`
	Name          types.String     `tfsdk:"name"`
	Description   types.String     `tfsdk:"description"`
	DefaultAction types.String     `tfsdk:"default_action"`
	Sequences     []RouteSequences `tfsdk:"sequences"`
}

type RouteSequences struct {
	Id            types.Int64                   `tfsdk:"id"`
	IpType        types.String                  `tfsdk:"ip_type"`
	Name          types.String                  `tfsdk:"name"`
	BaseAction    types.String                  `tfsdk:"base_action"`
	MatchEntries  []RouteSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []RouteSequencesActionEntries `tfsdk:"action_entries"`
}

type RouteSequencesMatchEntries struct {
	Type                         types.String `tfsdk:"type"`
	PrefixListId                 types.String `tfsdk:"prefix_list_id"`
	PrefixListVersion            types.Int64  `tfsdk:"prefix_list_version"`
	AsPathListId                 types.String `tfsdk:"as_path_list_id"`
	AsPathListVersion            types.Int64  `tfsdk:"as_path_list_version"`
	CommunityListIds             types.List   `tfsdk:"community_list_ids"`
	CommunityListVersions        types.List   `tfsdk:"community_list_versions"`
	CommunityListMatchFlag       types.String `tfsdk:"community_list_match_flag"`
	ExpandedCommunityListId      types.String `tfsdk:"expanded_community_list_id"`
	ExpandedCommunityListVersion types.Int64  `tfsdk:"expanded_community_list_version"`
	ExtendedCommunityListId      types.String `tfsdk:"extended_community_list_id"`
	ExtendedCommunityListVersion types.Int64  `tfsdk:"extended_community_list_version"`
	LocalPreference              types.Int64  `tfsdk:"local_preference"`
	Metric                       types.Int64  `tfsdk:"metric"`
	NextHop                      types.String `tfsdk:"next_hop"`
	Origin                       types.String `tfsdk:"origin"`
	Peer                         types.String `tfsdk:"peer"`
	OmpTag                       types.Int64  `tfsdk:"omp_tag"`
	OspfTag                      types.Int64  `tfsdk:"ospf_tag"`
}
type RouteSequencesActionEntries struct {
	Type                types.String `tfsdk:"type"`
	Aggregator          types.Int64  `tfsdk:"aggregator"`
	AggregatorIpAddress types.String `tfsdk:"aggregator_ip_address"`
	AsPathPrepend       types.String `tfsdk:"as_path_prepend"`
	AsPathExclude       types.String `tfsdk:"as_path_exclude"`
	AtomicAggregate     types.Bool   `tfsdk:"atomic_aggregate"`
	Community           types.String `tfsdk:"community"`
	CommunityAdditive   types.Bool   `tfsdk:"community_additive"`
	LocalPreference     types.Int64  `tfsdk:"local_preference"`
	Metric              types.Int64  `tfsdk:"metric"`
	Weight              types.Int64  `tfsdk:"weight"`
	MetricType          types.String `tfsdk:"metric_type"`
	NextHop             types.String `tfsdk:"next_hop"`
	OmpTag              types.Int64  `tfsdk:"omp_tag"`
	OspfTag             types.Int64  `tfsdk:"ospf_tag"`
	Origin              types.String `tfsdk:"origin"`
	Originator          types.String `tfsdk:"originator"`
}

func (data Route) getType() string {
	return "vedgeRoute"
}

func (data Route) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	body, _ = sjson.Set(body, "type", "vedgeRoute")
	path := ""
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, path+"defaultAction.type", data.DefaultAction.ValueString())
	}
	if len(data.Sequences) > 0 {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "sequenceType", "vedgeRoute")
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if len(item.MatchEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.PrefixListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.PrefixListId.ValueString())
					}
					if !childItem.AsPathListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.AsPathListId.ValueString())
					}
					if !childItem.CommunityListIds.IsNull() {
						var values []string
						childItem.CommunityListIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "refs", values)
					}
					if !childItem.CommunityListMatchFlag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "matchFlag", childItem.CommunityListMatchFlag.ValueString())
					}
					if !childItem.ExpandedCommunityListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ExpandedCommunityListId.ValueString())
					}
					if !childItem.ExtendedCommunityListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ExtendedCommunityListId.ValueString())
					}
					if !childItem.LocalPreference.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.LocalPreference.ValueInt64()))
					}
					if !childItem.Metric.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Metric.ValueInt64()))
					}
					if !childItem.NextHop.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.NextHop.ValueString())
					}
					if !childItem.Origin.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Origin.ValueString())
					}
					if !childItem.Peer.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Peer.ValueString())
					}
					if !childItem.OmpTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OmpTag.ValueInt64()))
					}
					if !childItem.OspfTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OspfTag.ValueInt64()))
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			itemBody, _ = sjson.Set(itemBody, "actions.0.type", "set")
			if len(item.ActionEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "actions.0.parameter", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.Aggregator.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.aggregator", fmt.Sprint(childItem.Aggregator.ValueInt64()))
					}
					if !childItem.AggregatorIpAddress.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.ipAddress", childItem.AggregatorIpAddress.ValueString())
					}
					if !childItem.AsPathPrepend.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.prepend", childItem.AsPathPrepend.ValueString())
					}
					if !childItem.AsPathExclude.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.exclude", childItem.AsPathExclude.ValueString())
					}
					if !childItem.AtomicAggregate.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.AtomicAggregate.ValueBool()))
					}
					if !childItem.Community.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Community.ValueString())
					}
					if !childItem.CommunityAdditive.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.CommunityAdditive.ValueBool()))
					}
					if !childItem.LocalPreference.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.LocalPreference.ValueInt64()))
					}
					if !childItem.Metric.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Metric.ValueInt64()))
					}
					if !childItem.Weight.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Weight.ValueInt64()))
					}
					if !childItem.MetricType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.MetricType.ValueString())
					}
					if !childItem.NextHop.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.NextHop.ValueString())
					}
					if !childItem.OmpTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OmpTag.ValueInt64()))
					}
					if !childItem.OspfTag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OspfTag.ValueInt64()))
					}
					if !childItem.Origin.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Origin.ValueString())
					}
					if !childItem.Originator.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Originator.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.0.parameter.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

func (data *Route) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	path := ""
	if value := res.Get(path + "defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get(path + "sequences"); value.Exists() {
		data.Sequences = make([]RouteSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RouteSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() {
				item.MatchEntries = make([]RouteSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := RouteSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); cItem.Type.ValueString() == "address" && ccValue.Exists() {
						cItem.PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); cItem.Type.ValueString() == "asPath" && ccValue.Exists() {
						cItem.AsPathListId = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathListId = types.StringNull()
					}
					if ccValue := cv.Get("refs"); cItem.Type.ValueString() == "advancedCommunity" && ccValue.Exists() {
						cItem.CommunityListIds = helpers.GetStringList(ccValue.Array())
					} else {
						cItem.CommunityListIds = types.ListNull(types.StringType)
					}
					if ccValue := cv.Get("matchFlag"); cItem.Type.ValueString() == "advancedCommunity" && ccValue.Exists() {
						cItem.CommunityListMatchFlag = types.StringValue(ccValue.String())
					} else {
						cItem.CommunityListMatchFlag = types.StringNull()
					}
					if ccValue := cv.Get("ref"); cItem.Type.ValueString() == "expandedCommunity" && ccValue.Exists() {
						cItem.ExpandedCommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExpandedCommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); cItem.Type.ValueString() == "extCommunity" && ccValue.Exists() {
						cItem.ExtendedCommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExtendedCommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "localPreference" && ccValue.Exists() {
						cItem.LocalPreference = types.Int64Value(ccValue.Int())
					} else {
						cItem.LocalPreference = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "metric" && ccValue.Exists() {
						cItem.Metric = types.Int64Value(ccValue.Int())
					} else {
						cItem.Metric = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "nextHop" && ccValue.Exists() {
						cItem.NextHop = types.StringValue(ccValue.String())
					} else {
						cItem.NextHop = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "origin" && ccValue.Exists() {
						cItem.Origin = types.StringValue(ccValue.String())
					} else {
						cItem.Origin = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "peer" && ccValue.Exists() {
						cItem.Peer = types.StringValue(ccValue.String())
					} else {
						cItem.Peer = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "ompTag" && ccValue.Exists() {
						cItem.OmpTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OmpTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "ospfTag" && ccValue.Exists() {
						cItem.OspfTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OspfTag = types.Int64Null()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions.0.parameter"); cValue.Exists() {
				item.ActionEntries = make([]RouteSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := RouteSequencesActionEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("value.aggregator"); cItem.Type.ValueString() == "aggregator" && ccValue.Exists() {
						cItem.Aggregator = types.Int64Value(ccValue.Int())
					} else {
						cItem.Aggregator = types.Int64Null()
					}
					if ccValue := cv.Get("value.ipAddress"); cItem.Type.ValueString() == "aggregator" && ccValue.Exists() {
						cItem.AggregatorIpAddress = types.StringValue(ccValue.String())
					} else {
						cItem.AggregatorIpAddress = types.StringNull()
					}
					if ccValue := cv.Get("value.prepend"); cItem.Type.ValueString() == "asPath" && ccValue.Exists() {
						cItem.AsPathPrepend = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathPrepend = types.StringNull()
					}
					if ccValue := cv.Get("value.exclude"); cItem.Type.ValueString() == "asPath" && ccValue.Exists() {
						cItem.AsPathExclude = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathExclude = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "community" && ccValue.Exists() {
						cItem.Community = types.StringValue(ccValue.String())
					} else {
						cItem.Community = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "communityAdditive" && ccValue.Exists() {
						cItem.CommunityAdditive = types.BoolValue(ccValue.Bool())
					} else {
						cItem.CommunityAdditive = types.BoolNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "localPreference" && ccValue.Exists() {
						cItem.LocalPreference = types.Int64Value(ccValue.Int())
					} else {
						cItem.LocalPreference = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "metric" && ccValue.Exists() {
						cItem.Metric = types.Int64Value(ccValue.Int())
					} else {
						cItem.Metric = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "weight" && ccValue.Exists() {
						cItem.Weight = types.Int64Value(ccValue.Int())
					} else {
						cItem.Weight = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "metrictype" && ccValue.Exists() {
						cItem.MetricType = types.StringValue(ccValue.String())
					} else {
						cItem.MetricType = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "nextHop" && ccValue.Exists() {
						cItem.NextHop = types.StringValue(ccValue.String())
					} else {
						cItem.NextHop = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "ompTag" && ccValue.Exists() {
						cItem.OmpTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OmpTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "ospfTag" && ccValue.Exists() {
						cItem.OspfTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OspfTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "origin" && ccValue.Exists() {
						cItem.Origin = types.StringValue(ccValue.String())
					} else {
						cItem.Origin = types.StringNull()
					}
					if ccValue := cv.Get("value"); cItem.Type.ValueString() == "originator" && ccValue.Exists() {
						cItem.Originator = types.StringValue(ccValue.String())
					} else {
						cItem.Originator = types.StringNull()
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

func (data *Route) hasChanges(ctx context.Context, state *Route) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].IpType.Equal(state.Sequences[i].IpType) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].BaseAction.Equal(state.Sequences[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].AsPathListId.Equal(state.Sequences[i].MatchEntries[ii].AsPathListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListIds.Equal(state.Sequences[i].MatchEntries[ii].CommunityListIds) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListMatchFlag.Equal(state.Sequences[i].MatchEntries[ii].CommunityListMatchFlag) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExpandedCommunityListId.Equal(state.Sequences[i].MatchEntries[ii].ExpandedCommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExtendedCommunityListId.Equal(state.Sequences[i].MatchEntries[ii].ExtendedCommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].LocalPreference.Equal(state.Sequences[i].MatchEntries[ii].LocalPreference) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Metric.Equal(state.Sequences[i].MatchEntries[ii].Metric) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].NextHop.Equal(state.Sequences[i].MatchEntries[ii].NextHop) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Origin.Equal(state.Sequences[i].MatchEntries[ii].Origin) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Peer.Equal(state.Sequences[i].MatchEntries[ii].Peer) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].OmpTag.Equal(state.Sequences[i].MatchEntries[ii].OmpTag) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].OspfTag.Equal(state.Sequences[i].MatchEntries[ii].OspfTag) {
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
					if !data.Sequences[i].ActionEntries[ii].Aggregator.Equal(state.Sequences[i].ActionEntries[ii].Aggregator) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AggregatorIpAddress.Equal(state.Sequences[i].ActionEntries[ii].AggregatorIpAddress) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AsPathPrepend.Equal(state.Sequences[i].ActionEntries[ii].AsPathPrepend) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AsPathExclude.Equal(state.Sequences[i].ActionEntries[ii].AsPathExclude) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AtomicAggregate.Equal(state.Sequences[i].ActionEntries[ii].AtomicAggregate) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Community.Equal(state.Sequences[i].ActionEntries[ii].Community) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CommunityAdditive.Equal(state.Sequences[i].ActionEntries[ii].CommunityAdditive) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LocalPreference.Equal(state.Sequences[i].ActionEntries[ii].LocalPreference) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Metric.Equal(state.Sequences[i].ActionEntries[ii].Metric) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Weight.Equal(state.Sequences[i].ActionEntries[ii].Weight) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].MetricType.Equal(state.Sequences[i].ActionEntries[ii].MetricType) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].NextHop.Equal(state.Sequences[i].ActionEntries[ii].NextHop) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].OmpTag.Equal(state.Sequences[i].ActionEntries[ii].OmpTag) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].OspfTag.Equal(state.Sequences[i].ActionEntries[ii].OspfTag) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Origin.Equal(state.Sequences[i].ActionEntries[ii].Origin) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Originator.Equal(state.Sequences[i].ActionEntries[ii].Originator) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

func (data *Route) getMatchPrefixListVersion(ctx context.Context, name string, id int64) types.Int64 {
	for _, item := range data.Sequences {
		if item.Name.ValueString() == name && item.Id.ValueInt64() == id {
			for _, cItem := range item.MatchEntries {
				if cItem.Type.ValueString() == "address" {
					return cItem.PrefixListVersion
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *Route) getMatchAsPathListVersion(ctx context.Context, name string, id int64) types.Int64 {
	for _, item := range data.Sequences {
		if item.Name.ValueString() == name && item.Id.ValueInt64() == id {
			for _, cItem := range item.MatchEntries {
				if cItem.Type.ValueString() == "asPath" {
					return cItem.AsPathListVersion
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *Route) getMatchCommunityListVersions(ctx context.Context, name string, id int64) types.List {
	for _, item := range data.Sequences {
		if item.Name.ValueString() == name && item.Id.ValueInt64() == id {
			for _, cItem := range item.MatchEntries {
				if cItem.Type.ValueString() == "advancedCommunity" {
					return cItem.CommunityListVersions
				}
			}
		}
	}
	return types.ListNull(types.StringType)
}

func (data *Route) getMatchExpandedCommunityListVersion(ctx context.Context, name string, id int64) types.Int64 {
	for _, item := range data.Sequences {
		if item.Name.ValueString() == name && item.Id.ValueInt64() == id {
			for _, cItem := range item.MatchEntries {
				if cItem.Type.ValueString() == "expandedCommunity" {
					return cItem.ExpandedCommunityListVersion
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *Route) getMatchExtendedCommunityListVersion(ctx context.Context, name string, id int64) types.Int64 {
	for _, item := range data.Sequences {
		if item.Name.ValueString() == name && item.Id.ValueInt64() == id {
			for _, cItem := range item.MatchEntries {
				if cItem.Type.ValueString() == "extCommunity" {
					return cItem.ExtendedCommunityListVersion
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *Route) updateVersions(ctx context.Context, state Route) {
	for s := range data.Sequences {
		id := data.Sequences[s].Id.ValueInt64()
		name := data.Sequences[s].Name.ValueString()
		for m := range data.Sequences[s].MatchEntries {
			t := data.Sequences[s].MatchEntries[m].Type.ValueString()
			data.Sequences[s].MatchEntries[m].CommunityListVersions = types.ListNull(types.StringType)
			if t == "address" {
				data.Sequences[s].MatchEntries[m].PrefixListVersion = state.getMatchPrefixListVersion(ctx, name, id)
			} else if t == "asPath" {
				data.Sequences[s].MatchEntries[m].AsPathListVersion = state.getMatchAsPathListVersion(ctx, name, id)
			} else if t == "advancedCommunity" {
				data.Sequences[s].MatchEntries[m].CommunityListVersions = state.getMatchCommunityListVersions(ctx, name, id)
			} else if t == "expandedCommunity" {
				data.Sequences[s].MatchEntries[m].ExpandedCommunityListVersion = state.getMatchExpandedCommunityListVersion(ctx, name, id)
			} else if t == "extCommunity" {
				data.Sequences[s].MatchEntries[m].ExtendedCommunityListVersion = state.getMatchExtendedCommunityListVersion(ctx, name, id)
			}
		}
	}
}
