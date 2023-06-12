package google

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/bigquery"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var IamBigqueryDatasetSchema = bigquery.IamBigqueryDatasetSchema

func NewBigqueryDatasetIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return bigquery.NewBigqueryDatasetIamUpdater(d, config)
}

func BigqueryDatasetIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	return bigquery.BigqueryDatasetIdParseFunc(d, config)
}
