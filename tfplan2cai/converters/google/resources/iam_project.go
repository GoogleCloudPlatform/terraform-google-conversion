package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/resourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

var IamProjectSchema = resourcemanager.IamProjectSchema

func NewProjectIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return resourcemanager.NewProjectIamUpdater(d, config)
}

func ProjectIdParseFunc(d *schema.ResourceData, _ *transport_tpg.Config) error {
	return resourcemanager.ProjectIdParseFunc(d, nil)
}

func compareProjectName(_, old, new string, _ *schema.ResourceData) bool {
	// We can either get "projects/project-id" or "project-id", so strip any prefixes
	return resourcemanager.CompareProjectName("", old, new, nil)
}
