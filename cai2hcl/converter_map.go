package convert

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/resourcemanager"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta"
)

var ConverterNames = map[string]string{
	resourcemanager.ProjectAssetType:        "google_project",
	resourcemanager.ProjectBillingAssetType: "google_project",
	compute.ComputeInstanceAssetType:        "google_compute_instance",
	compute.ComputeForwardingRuleAssetType:  "google_compute_forwarding_rule",
}

var converterFactories = map[string]func(name string, schema map[string]*schema.Schema) common.Converter{
	"google_project":                 resourcemanager.NewProjectConverter,
	"google_compute_instance":        compute.NewComputeInstanceConverter,
	"google_compute_forwarding_rule": compute.NewComputeForwardingRuleConverter,
}

var ConverterMap map[string]common.Converter

func init() {
	tpgProvider := tpg.Provider()

	ConverterMap = make(map[string]common.Converter, len(converterFactories))
	for name, factory := range converterFactories {
		ConverterMap[name] = factory(name, tpgProvider.ResourcesMap[name].Schema)
	}
}
