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
var _ resource.Resource = &DeviceACLPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &DeviceACLPolicyDefinitionResource{}

func NewDeviceACLPolicyDefinitionResource() resource.Resource {
	return &DeviceACLPolicyDefinitionResource{}
}

type DeviceACLPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *DeviceACLPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_acl_policy_definition"
}

func (r *DeviceACLPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Device ACL policy definition.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the policy definition",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the policy definition",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The policy defintion type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "drop"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of ACL sequences").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").AddIntegerRangeDescription(1, 65534).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65534),
							},
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP version, either `ipv4` or `ipv6`").AddStringEnumDescription("ipv4", "ipv6").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Optional:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("accept", "drop"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("sourceIp", "destinationIp", "sourcePort", "destinationPort", "sourceDataPrefixList", "destinationDataPrefixList").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("sourceIp", "destinationIp", "sourcePort", "destinationPort", "sourceDataPrefixList", "destinationDataPrefixList"),
										},
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source IP prefix").String,
										Optional:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination IP prefix").String,
										Optional:            true,
									},
									"source_port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source port").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"destination_port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination port").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source data prefix list ID").String,
										Optional:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source data prefix list version").String,
										Optional:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination data prefix list ID").String,
										Optional:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination data prefix list version").String,
										Optional:            true,
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
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("count").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("count"),
										},
									},
									"counter_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter name").String,
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

func (r *DeviceACLPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

func (r *DeviceACLPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceACL

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/policy/definition/deviceaccesspolicy", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)
	plan.Type = types.StringValue(plan.getType())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceACLPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, oldState DeviceACL

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.State.Get(ctx, &oldState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/policy/definition/deviceaccesspolicy/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Failed to find specified resource" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	state.updateVersions(ctx, oldState)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *DeviceACLPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceACL

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
		res, err := r.client.Put("/template/policy/definition/deviceaccesspolicy/"+plan.Id.ValueString(), body)
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

func (r *DeviceACLPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceACL

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/policy/definition/deviceaccesspolicy/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *DeviceACLPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
