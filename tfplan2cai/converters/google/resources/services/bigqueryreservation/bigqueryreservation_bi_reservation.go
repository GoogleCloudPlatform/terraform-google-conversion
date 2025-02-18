// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryreservation/BiReservation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigqueryreservation

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryReservationBiReservationAssetType string = "bigqueryreservation.googleapis.com/BiReservation"

func ResourceConverterBigqueryReservationBiReservation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryReservationBiReservationAssetType,
		Convert:   GetBigqueryReservationBiReservationCaiObject,
	}
}

func GetBigqueryReservationBiReservationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//bigqueryreservation.googleapis.com/projects/{{project}}/locations/{{location}}/biReservation")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryReservationBiReservationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryReservationBiReservationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigqueryreservation/v1/rest",
				DiscoveryName:        "BiReservation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryReservationBiReservationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	sizeProp, err := expandBigqueryReservationBiReservationSize(d.Get("size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("size"); !tpgresource.IsEmptyValue(reflect.ValueOf(sizeProp)) && (ok || !reflect.DeepEqual(v, sizeProp)) {
		obj["size"] = sizeProp
	}
	preferredTablesProp, err := expandBigqueryReservationBiReservationPreferredTables(d.Get("preferred_tables"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("preferred_tables"); !tpgresource.IsEmptyValue(reflect.ValueOf(preferredTablesProp)) && (ok || !reflect.DeepEqual(v, preferredTablesProp)) {
		obj["preferredTables"] = preferredTablesProp
	}

	return obj, nil
}

func expandBigqueryReservationBiReservationSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectId, err := expandBigqueryReservationBiReservationPreferredTablesProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedDatasetId, err := expandBigqueryReservationBiReservationPreferredTablesDatasetId(original["dataset_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["datasetId"] = transformedDatasetId
		}

		transformedTableId, err := expandBigqueryReservationBiReservationPreferredTablesTableId(original["table_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tableId"] = transformedTableId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigqueryReservationBiReservationPreferredTablesProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTablesDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationBiReservationPreferredTablesTableId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
