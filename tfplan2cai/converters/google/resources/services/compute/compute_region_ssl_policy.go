// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/RegionSslPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func regionSslPolicyCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	profile := diff.Get("profile")
	customFeaturesCount := diff.Get("custom_features.#")

	// Validate that policy configs aren't incompatible during all phases
	// CUSTOM profile demands non-zero custom_features, and other profiles (i.e., not CUSTOM) demand zero custom_features
	if diff.HasChange("profile") || diff.HasChange("custom_features") {
		if profile.(string) == "CUSTOM" && customFeaturesCount.(int) == 0 {
			return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but no custom_features are set.", diff.Get("name"), profile.(string))
		} else if profile.(string) != "CUSTOM" && customFeaturesCount.(int) != 0 {
			return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but using custom_features requires the profile to be CUSTOM.", diff.Get("name"), profile.(string))
		}
		return nil
	}
	return nil
}

const ComputeRegionSslPolicyAssetType string = "compute.googleapis.com/RegionSslPolicy"

func ResourceConverterComputeRegionSslPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionSslPolicyAssetType,
		Convert:   GetComputeRegionSslPolicyCaiObject,
	}
}

func GetComputeRegionSslPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionSslPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionSslPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionSslPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionSslPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	profileProp, err := expandComputeRegionSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(profileProp)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeRegionSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("min_tls_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(minTlsVersionProp)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeRegionSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_features"); ok || !reflect.DeepEqual(v, customFeaturesProp) {
		obj["customFeatures"] = customFeaturesProp
	}
	fingerprintProp, err := expandComputeRegionSslPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	regionProp, err := expandComputeRegionSslPolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionSslPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyMinTlsVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyCustomFeatures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionSslPolicyFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
