// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/PublicDelegatedPrefix.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputePublicDelegatedPrefixAssetType string = "compute.googleapis.com/PublicDelegatedPrefix"

func ResourceConverterComputePublicDelegatedPrefix() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputePublicDelegatedPrefixAssetType,
		Convert:   GetComputePublicDelegatedPrefixCaiObject,
	}
}

func GetComputePublicDelegatedPrefixCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputePublicDelegatedPrefixApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputePublicDelegatedPrefixAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "PublicDelegatedPrefix",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputePublicDelegatedPrefixApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicDelegatedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	isLiveMigrationProp, err := expandComputePublicDelegatedPrefixIsLiveMigration(d.Get("is_live_migration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_live_migration"); !tpgresource.IsEmptyValue(reflect.ValueOf(isLiveMigrationProp)) && (ok || !reflect.DeepEqual(v, isLiveMigrationProp)) {
		obj["isLiveMigration"] = isLiveMigrationProp
	}
	nameProp, err := expandComputePublicDelegatedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	parentPrefixProp, err := expandComputePublicDelegatedPrefixParentPrefix(d.Get("parent_prefix"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_prefix"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentPrefixProp)) && (ok || !reflect.DeepEqual(v, parentPrefixProp)) {
		obj["parentPrefix"] = parentPrefixProp
	}
	modeProp, err := expandComputePublicDelegatedPrefixMode(d.Get("mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(modeProp)) && (ok || !reflect.DeepEqual(v, modeProp)) {
		obj["mode"] = modeProp
	}
	allocatablePrefixLengthProp, err := expandComputePublicDelegatedPrefixAllocatablePrefixLength(d.Get("allocatable_prefix_length"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allocatable_prefix_length"); !tpgresource.IsEmptyValue(reflect.ValueOf(allocatablePrefixLengthProp)) && (ok || !reflect.DeepEqual(v, allocatablePrefixLengthProp)) {
		obj["allocatablePrefixLength"] = allocatablePrefixLengthProp
	}
	ipCidrRangeProp, err := expandComputePublicDelegatedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	return obj, nil
}

func expandComputePublicDelegatedPrefixDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIsLiveMigration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixParentPrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixAllocatablePrefixLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
