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

type DataIPv6PrefixListPolicyObject struct {
	Id      types.String                            `tfsdk:"id"`
	Version types.Int64                             `tfsdk:"version"`
	Name    types.String                            `tfsdk:"name"`
	Entries []DataIPv6PrefixListPolicyObjectEntries `tfsdk:"entries"`
}

type DataIPv6PrefixListPolicyObjectEntries struct {
	Prefix types.String `tfsdk:"prefix"`
}

func (data DataIPv6PrefixListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "type", "dataipv6prefix")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6Prefix", item.Prefix.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *DataIPv6PrefixListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]DataIPv6PrefixListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DataIPv6PrefixListPolicyObjectEntries{}
			if cValue := v.Get("ipv6Prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			} else {
				item.Prefix = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

func (data *DataIPv6PrefixListPolicyObject) hasChanges(ctx context.Context, state *DataIPv6PrefixListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].Prefix.Equal(state.Entries[i].Prefix) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
