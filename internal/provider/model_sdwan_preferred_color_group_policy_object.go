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

type PreferredColorGroupPolicyObject struct {
	Id                       types.String `tfsdk:"id"`
	Version                  types.Int64  `tfsdk:"version"`
	Name                     types.String `tfsdk:"name"`
	PrimaryColorPreference   types.String `tfsdk:"primary_color_preference"`
	PrimaryPathPreference    types.String `tfsdk:"primary_path_preference"`
	SecondaryColorPreference types.String `tfsdk:"secondary_color_preference"`
	SecondaryPathPreference  types.String `tfsdk:"secondary_path_preference"`
	TertiaryColorPreference  types.String `tfsdk:"tertiary_color_preference"`
	TertiaryPathPreference   types.String `tfsdk:"tertiary_path_preference"`
}

func (data PreferredColorGroupPolicyObject) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "type", "preferredColorGroup")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PrimaryColorPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.primaryPreference.colorPreference", data.PrimaryColorPreference.ValueString())
	}
	if !data.PrimaryPathPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.primaryPreference.pathPreference", data.PrimaryPathPreference.ValueString())
	}
	if !data.SecondaryColorPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.secondaryPreference.colorPreference", data.SecondaryColorPreference.ValueString())
	}
	if !data.SecondaryPathPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.secondaryPreference.pathPreference", data.SecondaryPathPreference.ValueString())
	}
	if !data.TertiaryColorPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.tertiaryPreference.colorPreference", data.TertiaryColorPreference.ValueString())
	}
	if !data.TertiaryPathPreference.IsNull() {
		body, _ = sjson.Set(body, "entries.0.tertiaryPreference.pathPreference", data.TertiaryPathPreference.ValueString())
	}
	return body
}

func (data *PreferredColorGroupPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries.0.primaryPreference.colorPreference"); value.Exists() {
		data.PrimaryColorPreference = types.StringValue(value.String())
	} else {
		data.PrimaryColorPreference = types.StringNull()
	}
	if value := res.Get("entries.0.primaryPreference.pathPreference"); value.Exists() {
		data.PrimaryPathPreference = types.StringValue(value.String())
	} else {
		data.PrimaryPathPreference = types.StringNull()
	}
	if value := res.Get("entries.0.secondaryPreference.colorPreference"); value.Exists() {
		data.SecondaryColorPreference = types.StringValue(value.String())
	} else {
		data.SecondaryColorPreference = types.StringNull()
	}
	if value := res.Get("entries.0.secondaryPreference.pathPreference"); value.Exists() {
		data.SecondaryPathPreference = types.StringValue(value.String())
	} else {
		data.SecondaryPathPreference = types.StringNull()
	}
	if value := res.Get("entries.0.tertiaryPreference.colorPreference"); value.Exists() {
		data.TertiaryColorPreference = types.StringValue(value.String())
	} else {
		data.TertiaryColorPreference = types.StringNull()
	}
	if value := res.Get("entries.0.tertiaryPreference.pathPreference"); value.Exists() {
		data.TertiaryPathPreference = types.StringValue(value.String())
	} else {
		data.TertiaryPathPreference = types.StringNull()
	}
}

func (data *PreferredColorGroupPolicyObject) hasChanges(ctx context.Context, state *PreferredColorGroupPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.PrimaryColorPreference.Equal(state.PrimaryColorPreference) {
		hasChanges = true
	}
	if !data.PrimaryPathPreference.Equal(state.PrimaryPathPreference) {
		hasChanges = true
	}
	if !data.SecondaryColorPreference.Equal(state.SecondaryColorPreference) {
		hasChanges = true
	}
	if !data.SecondaryPathPreference.Equal(state.SecondaryPathPreference) {
		hasChanges = true
	}
	if !data.TertiaryColorPreference.Equal(state.TertiaryColorPreference) {
		hasChanges = true
	}
	if !data.TertiaryPathPreference.Equal(state.TertiaryPathPreference) {
		hasChanges = true
	}
	return hasChanges
}
