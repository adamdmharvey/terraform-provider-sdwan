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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ACLPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &ACLPolicyDefinitionDataSource{}
)

func NewACLPolicyDefinitionDataSource() datasource.DataSource {
	return &ACLPolicyDefinitionDataSource{}
}

type ACLPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *ACLPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_policy_definition"
}

func (d *ACLPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the ACL Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default action, either `accept` or `drop`",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "List of ACL sequences",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence ID",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "IP version, either `ipv4` or `ipv6`",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action, either `accept` or `drop`",
							Computed:            true,
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
										Computed:            true,
									},
									"dscp": schema.Int64Attribute{
										MarkdownDescription: "DSCP value",
										Computed:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "Source IP prefix",
										Computed:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: "Destination IP prefix",
										Computed:            true,
									},
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: "Class map ID",
										Computed:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: "Class map version",
										Computed:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: "Packet length",
										Computed:            true,
									},
									"priority": schema.StringAttribute{
										MarkdownDescription: "PLP - priority",
										Computed:            true,
									},
									"source_port": schema.Int64Attribute{
										MarkdownDescription: "Source port",
										Computed:            true,
									},
									"destination_port": schema.Int64Attribute{
										MarkdownDescription: "Destination port",
										Computed:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Source data prefix list ID",
										Computed:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Source data prefix list version",
										Computed:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Destination data prefix list ID",
										Computed:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Destination data prefix list version",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Single value (0-255) or multiple values separated by spaces",
										Computed:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: "TCP parameters",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of action entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
										Computed:            true,
									},
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: "Class map ID",
										Computed:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: "Class map version",
										Computed:            true,
									},
									"counter_name": schema.StringAttribute{
										MarkdownDescription: "Counter name",
										Computed:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: "Enable logging",
										Computed:            true,
									},
									"mirror_id": schema.StringAttribute{
										MarkdownDescription: "Mirror ID",
										Computed:            true,
									},
									"mirror_version": schema.Int64Attribute{
										MarkdownDescription: "Mirror version",
										Computed:            true,
									},
									"policer_id": schema.StringAttribute{
										MarkdownDescription: "Policer ID",
										Computed:            true,
									},
									"policer_version": schema.Int64Attribute{
										MarkdownDescription: "Policer version",
										Computed:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: "List of set parameters",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of set parameter",
													Computed:            true,
												},
												"dscp": schema.Int64Attribute{
													MarkdownDescription: "DSCP value",
													Computed:            true,
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: "Next hop IP",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ACLPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *ACLPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ACLPolicyDefinition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/definition/acl/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
