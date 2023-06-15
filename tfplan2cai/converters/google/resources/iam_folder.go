package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/resourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

var IamFolderSchema = resourcemanager.IamFolderSchema

func NewFolderIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return resourcemanager.NewFolderIamUpdater(d, config)
}

func FolderIdParseFunc(d *schema.ResourceData, _ *transport_tpg.Config) error {
	return resourcemanager.FolderIdParseFunc(d, nil)
}

func canonicalFolderId(folder string) string {
	return resourcemanager.CanonicalFolderId(folder)
}

// Retrieve the existing IAM Policy for a folder
func getFolderIamPolicyByFolderName(folderName, userAgent string, config *transport_tpg.Config) (*cloudresourcemanager.Policy, error) {
	return resourcemanager.GetFolderIamPolicyByFolderName(folderName, userAgent, config)
}
