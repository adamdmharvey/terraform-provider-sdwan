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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ExpandedCommunityList struct {
	Id      types.String                   `tfsdk:"id"`
	Version types.Int64                    `tfsdk:"version"`
	Name    types.String                   `tfsdk:"name"`
	Entries []ExpandedCommunityListEntries `tfsdk:"entries"`
}

type ExpandedCommunityListEntries struct {
	Community types.String `tfsdk:"community"`
}

func (data ExpandedCommunityList) getType() string {
	return "expandedcommunity"
}

func (data ExpandedCommunityList) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "description", "Desc Not Required")
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "type", "expandedcommunity")
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Community.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "community", item.Community.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *ExpandedCommunityList) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]ExpandedCommunityListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ExpandedCommunityListEntries{}
			if cValue := v.Get("community"); cValue.Exists() {
				item.Community = types.StringValue(cValue.String())
			} else {
				item.Community = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}
