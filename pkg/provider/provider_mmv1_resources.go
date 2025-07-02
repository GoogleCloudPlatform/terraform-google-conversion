package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/services/bigquery"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/services/resourcemanager"
)

var handwrittenTfplan2caiResources = map[string]*schema.Resource{
	// ####### START handwritten resources ###########
	"google_compute_instance": compute.ResourceComputeInstance(),
	"google_project":          resourcemanager.ResourceGoogleProject(),
	// ####### END handwritten resources ###########
}

// Generated resources: 7
var generatedResources = map[string]*schema.Resource{
	"google_bigquery_dataset":             bigquery.ResourceBigQueryDataset(),
	"google_compute_address":              compute.ResourceComputeAddress(),
	"google_compute_autoscaler":           compute.ResourceComputeAutoscaler(),
	"google_compute_backend_bucket":       compute.ResourceComputeBackendBucket(),
	"google_compute_disk":                 compute.ResourceComputeDisk(),
	"google_compute_external_vpn_gateway": compute.ResourceComputeExternalVpnGateway(),
	"google_compute_region_autoscaler":    compute.ResourceComputeRegionAutoscaler(),
}
