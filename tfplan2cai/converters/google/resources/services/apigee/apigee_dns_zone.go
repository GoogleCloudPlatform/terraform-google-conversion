// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/DnsZone.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeDnsZoneAssetType string = "apigee.googleapis.com/DnsZone"

func ResourceConverterApigeeDnsZone() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeDnsZoneAssetType,
		Convert:   GetApigeeDnsZoneCaiObject,
	}
}

func GetApigeeDnsZoneCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{org_id}}/dnsZones/{{dns_zone_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeDnsZoneApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeDnsZoneAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "DnsZone",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeDnsZoneApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	domainProp, err := expandApigeeDnsZoneDomain(d.Get("domain"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainProp)) && (ok || !reflect.DeepEqual(v, domainProp)) {
		obj["domain"] = domainProp
	}
	descriptionProp, err := expandApigeeDnsZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	peeringConfigProp, err := expandApigeeDnsZonePeeringConfig(d.Get("peering_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(peeringConfigProp)) && (ok || !reflect.DeepEqual(v, peeringConfigProp)) {
		obj["peeringConfig"] = peeringConfigProp
	}

	return obj, nil
}

func expandApigeeDnsZoneDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDnsZoneDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDnsZonePeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetProjectId, err := expandApigeeDnsZonePeeringConfigTargetProjectId(original["target_project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetProjectId"] = transformedTargetProjectId
	}

	transformedTargetNetworkId, err := expandApigeeDnsZonePeeringConfigTargetNetworkId(original["target_network_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNetworkId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetNetworkId"] = transformedTargetNetworkId
	}

	return transformed, nil
}

func expandApigeeDnsZonePeeringConfigTargetProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeDnsZonePeeringConfigTargetNetworkId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
