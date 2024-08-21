package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &resourcenamesDataSource{}
)

// NewResourcenamesDataSource is a helper function to simplify the provider implementation.
func NewResourcenamesDataSource() datasource.DataSource {
	return &resourcenamesDataSource{}
}

// resourcenamesDataSource is the data source implementation.
type resourcenamesDataSource struct{}

// Metadata returns the data source type name.
func (d *resourcenamesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resourcenames"
}

// Schema defines the schema for the data source.
func (d *resourcenamesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bu": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the business unit token.",
			},
			"purpose": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the purpose token.",
			},
			"geography": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the geography token.",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the region token.",
			},
			"client": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the client token.",
			},
			"program": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the program token.",
			},
			"environment": schema.StringAttribute{
				Optional:    true,
				Description: "The value of the environment token.",
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Description: "The integer index to append to the end of the resource name.",
			},
			"subscription": schema.StringAttribute{
				Computed:    true,
				Description: "The computed name for a Subscription.",
			},
			"resourcegroup": schema.StringAttribute{
				Computed:    true,
				Description: "The computed name for a Resource Group.",
			},
		},
	}
}

type resourcenamesDataModel struct {
	Bu            types.String `tfsdk:"bu"`
	Purpose       types.String `tfsdk:"purpose"`
	Geography     types.String `tfsdk:"geography"`
	Region        types.String `tfsdk:"region"`
	Client        types.String `tfsdk:"client"`
	Program       types.String `tfsdk:"program"`
	Environment   types.String `tfsdk:"environment"`
	Index         types.Int64  `tfsdk:"index"`
	Subscription  types.String `tfsdk:"subscription"`
	Resourcegroup types.String `tfsdk:"resourcegroup"`
}

// Read refreshes the Terraform state with the latest data.
func (d *resourcenamesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state resourcenamesDataModel
	req.Config.Get(ctx, &state)

	bu := state.Bu.ValueStringPointer()
	purpose := state.Purpose.ValueStringPointer()
	geography := state.Geography.ValueStringPointer()
	// region := state.Region.ValueStringPointer()
	client := state.Client.ValueStringPointer()
	program := state.Program.ValueStringPointer()
	environment := state.Environment.ValueStringPointer()
	index := state.Index.ValueInt64()

	state.Subscription = types.StringValue(joinNonNilStrings("-", bu, purpose, geography, client, program, environment) + "-sub-" + fmt.Sprintf("%03d", index))
	state.Resourcegroup = types.StringValue(joinNonNilStrings("-", bu, purpose, geography, client, program, environment) + "-rg-" + fmt.Sprintf("%03d", index))

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func joinNonNilStrings(delimiter string, strs ...*string) string {
	var filtered []string
	for _, str := range strs {
		if str != nil && *str != "" {
			filtered = append(filtered, *str)
		}
	}
	return strings.Join(filtered, delimiter)
}
