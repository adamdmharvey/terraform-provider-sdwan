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
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CustomControlTopologyPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &CustomControlTopologyPolicyDefinitionResource{}

func NewCustomControlTopologyPolicyDefinitionResource() resource.Resource {
	return &CustomControlTopologyPolicyDefinitionResource{}
}

type CustomControlTopologyPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CustomControlTopologyPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_control_topology_policy_definition"
}

func (r *CustomControlTopologyPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Custom Control Topology Policy Definition .").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition").String,
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action, either `accept` or `reject`").AddStringEnumDescription("accept", "reject").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "reject"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of sequences").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence type, either `route` or `tloc`").AddStringEnumDescription("route", "tloc").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("route", "tloc"),
							},
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence IP type, either `ipv4`, `ipv6` or `all`").AddStringEnumDescription("ipv4", "ipv6", "all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6", "all"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("colorList", "community", "expandedCommunity", "ompTag", "origin", "originator", "preference", "siteList", "pathType", "tlocList", "vpnList", "prefixList", "vpn", "tloc", "siteId", "carrier", "domainId", "groupId").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("colorList", "community", "expandedCommunity", "ompTag", "origin", "originator", "preference", "siteList", "pathType", "tlocList", "vpnList", "prefixList", "vpn", "tloc", "siteId", "carrier", "domainId", "groupId"),
										},
									},
									"color_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Color list ID").String,
										Optional:            true,
									},
									"color_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Color list version").String,
										Optional:            true,
									},
									"community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list ID").String,
										Optional:            true,
									},
									"community_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list version").String,
										Optional:            true,
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Expanded community list ID").String,
										Optional:            true,
									},
									"expanded_community_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Expanded community list version").String,
										Optional:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("OMP tag").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Origin").AddStringEnumDescription("igp", "egp", "incomplete").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("igp", "egp", "incomplete"),
										},
									},
									"originator": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Originator IP").String,
										Optional:            true,
									},
									"preference": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Preference").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"site_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Site list ID").String,
										Optional:            true,
									},
									"site_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Site list version").String,
										Optional:            true,
									},
									"path_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Path type").AddStringEnumDescription("hierarchical-path", "direct-path", "transport-gateway-path").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("hierarchical-path", "direct-path", "transport-gateway-path"),
										},
									},
									"tloc_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TLOC list ID").String,
										Optional:            true,
									},
									"tloc_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("TLOC list version").String,
										Optional:            true,
									},
									"vpn_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("VPN list ID").String,
										Optional:            true,
									},
									"vpn_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VPN list version").String,
										Optional:            true,
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list ID").String,
										Optional:            true,
									},
									"prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list version").String,
										Optional:            true,
									},
									"vpn_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VPN ID").AddIntegerRangeDescription(0, 65536).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65536),
										},
									},
									"tloc_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TLOC IP address").String,
										Optional:            true,
									},
									"tloc_color": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TLOC color").String,
										Optional:            true,
									},
									"tloc_encapsulation": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TLOC encapsulation").AddStringEnumDescription("ipsec", "gre").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ipsec", "gre"),
										},
									},
									"site_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Site ID").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"carrier": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Carrier").AddStringEnumDescription("default", "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("default", "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8"),
										},
									},
									"domain_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Domain ID").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"group_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Group ID").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of action entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("set", "exportTo").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("set", "exportTo"),
										},
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of set parameters").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of set parameter").AddStringEnumDescription("tlocList", "tloc", "tlocAction", "preference", "ompTag", "community", "communityAdditive", "service").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("tlocList", "tloc", "tlocAction", "preference", "ompTag", "community", "communityAdditive", "service"),
													},
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC list ID").String,
													Optional:            true,
												},
												"tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC list version").String,
													Optional:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC IP address").String,
													Optional:            true,
												},
												"tloc_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC color").String,
													Optional:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC encapsulation").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
												"tloc_action": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC action").AddStringEnumDescription("strict", "primary", "backup", "ecmp").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("strict", "primary", "backup", "ecmp"),
													},
												},
												"preference": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Preference").AddIntegerRangeDescription(0, 4294967295).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 4294967295),
													},
												},
												"omp_tag": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("OMP tag").AddIntegerRangeDescription(0, 4294967295).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 4294967295),
													},
												},
												"community": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Community value, e.g. `1000:10000` or `internet` or `local-AS`").String,
													Optional:            true,
												},
												"community_additive": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Community additive").String,
													Optional:            true,
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service type").AddStringEnumDescription("FW", "IDP", "IDS", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "netsvc5").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("FW", "IDP", "IDS", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "netsvc5"),
													},
												},
												"service_vpn_id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service VPN ID").AddIntegerRangeDescription(0, 65536).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 65536),
													},
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC list ID").String,
													Optional:            true,
												},
												"service_tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC list version").String,
													Optional:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC IP address").String,
													Optional:            true,
												},
												"service_tloc_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC color").String,
													Optional:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC encapsulation").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
											},
										},
									},
									"export_to_vpn_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Export to VPN list ID").String,
										Optional:            true,
									},
									"export_to_vpn_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Export to VPN list version").String,
										Optional:            true,
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

func (r *CustomControlTopologyPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

func (r *CustomControlTopologyPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CustomControlTopologyPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/policy/definition/control/", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomControlTopologyPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CustomControlTopologyPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/policy/definition/control/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Failed to find specified resource" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomControlTopologyPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CustomControlTopologyPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put("/template/policy/definition/control/"+plan.Id.ValueString(), body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomControlTopologyPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CustomControlTopologyPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/policy/definition/control/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *CustomControlTopologyPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
