// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/FirewallPolicy.yaml
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

const ComputeFirewallPolicyAssetType string = "compute.googleapis.com/FirewallPolicy"

func ResourceConverterComputeFirewallPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeFirewallPolicyAssetType,
		Convert:   GetComputeFirewallPolicyCaiObject,
	}
}

func GetComputeFirewallPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/locations/global/firewallPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeFirewallPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeFirewallPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "FirewallPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeFirewallPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	shortNameProp, err := expandComputeFirewallPolicyShortName(d.Get("short_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("short_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(shortNameProp)) && (ok || !reflect.DeepEqual(v, shortNameProp)) {
		obj["shortName"] = shortNameProp
	}
	descriptionProp, err := expandComputeFirewallPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentProp, err := expandComputeFirewallPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	fingerprintProp, err := expandComputeFirewallPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	return obj, nil
}

func expandComputeFirewallPolicyShortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
