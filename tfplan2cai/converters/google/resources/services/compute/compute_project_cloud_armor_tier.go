// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/ProjectCloudArmorTier.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeProjectCloudArmorTierAssetType string = "compute.googleapis.com/ProjectCloudArmorTier"

func ResourceConverterComputeProjectCloudArmorTier() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeProjectCloudArmorTierAssetType,
		Convert:   GetComputeProjectCloudArmorTierCaiObject,
	}
}

func GetComputeProjectCloudArmorTierCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeProjectCloudArmorTierApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeProjectCloudArmorTierAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "ProjectCloudArmorTier",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeProjectCloudArmorTierApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	cloudArmorTierProp, err := expandComputeProjectCloudArmorTierCloudArmorTier(d.Get("cloud_armor_tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_armor_tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudArmorTierProp)) && (ok || !reflect.DeepEqual(v, cloudArmorTierProp)) {
		obj["cloudArmorTier"] = cloudArmorTierProp
	}

	return obj, nil
}

func expandComputeProjectCloudArmorTierCloudArmorTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
