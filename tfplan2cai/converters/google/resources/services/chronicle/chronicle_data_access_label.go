// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/DataAccessLabel.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package chronicle

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ChronicleDataAccessLabelAssetType string = "{{location}}-chronicle.googleapis.com/DataAccessLabel"

func ResourceConverterChronicleDataAccessLabel() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ChronicleDataAccessLabelAssetType,
		Convert:   GetChronicleDataAccessLabelCaiObject,
	}
}

func GetChronicleDataAccessLabelCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-chronicle.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessLabels/{{data_access_label_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetChronicleDataAccessLabelApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ChronicleDataAccessLabelAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-chronicle/v1beta/rest",
				DiscoveryName:        "DataAccessLabel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetChronicleDataAccessLabelApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandChronicleDataAccessLabelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	udmQueryProp, err := expandChronicleDataAccessLabelUdmQuery(d.Get("udm_query"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("udm_query"); !tpgresource.IsEmptyValue(reflect.ValueOf(udmQueryProp)) && (ok || !reflect.DeepEqual(v, udmQueryProp)) {
		obj["udmQuery"] = udmQueryProp
	}

	return obj, nil
}

func expandChronicleDataAccessLabelDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessLabelUdmQuery(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
