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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CiscoVPNInterfaceIPSecFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoVPNInterfaceIPSecFeatureTemplateDataSource{}
)

func NewCiscoVPNInterfaceIPSecFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoVPNInterfaceIPSecFeatureTemplateDataSource{}
}

type CiscoVPNInterfaceIPSecFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoVPNInterfaceIPSecFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_vpn_interface_ipsec_feature_template"
}

func (d *CiscoVPNInterfaceIPSecFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco VPN Interface IPSec feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface name: IPsec when present",
				Computed:            true,
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "Interface description",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv4 address",
				Computed:            true,
			},
			"ip_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_source": schema.StringAttribute{
				MarkdownDescription: "Tunnel source IP Address",
				Computed:            true,
			},
			"tunnel_source_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_source_interface": schema.StringAttribute{
				MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
				Computed:            true,
			},
			"tunnel_source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_destination": schema.StringAttribute{
				MarkdownDescription: "Tunnel destination IP address",
				Computed:            true,
			},
			"tunnel_destination_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"application": schema.StringAttribute{
				MarkdownDescription: "Enable Application Tunnel Type",
				Computed:            true,
			},
			"application_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss_adjust": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_adjust_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: "Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)",
				Computed:            true,
			},
			"clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU <68..9216>, in bytes",
				Computed:            true,
			},
			"mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dead_peer_detection_interval": schema.Int64Attribute{
				MarkdownDescription: "IKE keepalive interval (seconds)",
				Computed:            true,
			},
			"dead_peer_detection_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dead_peer_detection_retries": schema.Int64Attribute{
				MarkdownDescription: "IKE keepalive retries",
				Computed:            true,
			},
			"dead_peer_detection_retries_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_version": schema.Int64Attribute{
				MarkdownDescription: "IKE Version <1..2>",
				Computed:            true,
			},
			"ike_mode": schema.StringAttribute{
				MarkdownDescription: "IKE integrity protocol",
				Computed:            true,
			},
			"ike_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: "IKE rekey interval <60..86400> seconds",
				Computed:            true,
			},
			"ike_rekey_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_ciphersuite": schema.StringAttribute{
				MarkdownDescription: "IKE identity the IKE preshared secret belongs to",
				Computed:            true,
			},
			"ike_ciphersuite_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_group": schema.StringAttribute{
				MarkdownDescription: "IKE Diffie Hellman Groups",
				Computed:            true,
			},
			"ike_group_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: "Use preshared key to authenticate IKE peer",
				Computed:            true,
			},
			"ike_pre_shared_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_pre_shared_key_local_id": schema.StringAttribute{
				MarkdownDescription: "IKE ID for the local endpoint. Input IPv4 address, domain name, or email address",
				Computed:            true,
			},
			"ike_pre_shared_key_local_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ike_pre_shared_key_remote_id": schema.StringAttribute{
				MarkdownDescription: "IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address",
				Computed:            true,
			},
			"ike_pre_shared_key_remote_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: "IPsec rekey interval <300..1209600> seconds",
				Computed:            true,
			},
			"ipsec_rekey_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_replay_window": schema.Int64Attribute{
				MarkdownDescription: "Replay window size 32..8192 (must be a power of 2)",
				Computed:            true,
			},
			"ipsec_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_ciphersuite": schema.StringAttribute{
				MarkdownDescription: "IPsec(ESP) encryption and integrity protocol",
				Computed:            true,
			},
			"ipsec_ciphersuite_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_perfect_forward_secrecy": schema.StringAttribute{
				MarkdownDescription: "IPsec perfect forward secrecy settings",
				Computed:            true,
			},
			"ipsec_perfect_forward_secrecy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tracker": schema.ListAttribute{
				MarkdownDescription: "Enable tracker for this interface",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tracker_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_route_via": schema.StringAttribute{
				MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
				Computed:            true,
			},
			"tunnel_route_via_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *CiscoVPNInterfaceIPSecFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *CiscoVPNInterfaceIPSecFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoVPNInterfaceIPSec

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/feature/object/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
