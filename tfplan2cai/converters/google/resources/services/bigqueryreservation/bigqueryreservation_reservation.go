// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryreservation/Reservation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigqueryreservation

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryReservationReservationAssetType string = "bigqueryreservation.googleapis.com/Reservation"

func ResourceConverterBigqueryReservationReservation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryReservationReservationAssetType,
		Convert:   GetBigqueryReservationReservationCaiObject,
	}
}

func GetBigqueryReservationReservationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//bigqueryreservation.googleapis.com/projects/{{project}}/locations/{{location}}/reservations/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryReservationReservationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryReservationReservationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigqueryreservation/v1/rest",
				DiscoveryName:        "Reservation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryReservationReservationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	slotCapacityProp, err := expandBigqueryReservationReservationSlotCapacity(d.Get("slot_capacity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("slot_capacity"); !tpgresource.IsEmptyValue(reflect.ValueOf(slotCapacityProp)) && (ok || !reflect.DeepEqual(v, slotCapacityProp)) {
		obj["slotCapacity"] = slotCapacityProp
	}
	ignoreIdleSlotsProp, err := expandBigqueryReservationReservationIgnoreIdleSlots(d.Get("ignore_idle_slots"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ignore_idle_slots"); !tpgresource.IsEmptyValue(reflect.ValueOf(ignoreIdleSlotsProp)) && (ok || !reflect.DeepEqual(v, ignoreIdleSlotsProp)) {
		obj["ignoreIdleSlots"] = ignoreIdleSlotsProp
	}
	concurrencyProp, err := expandBigqueryReservationReservationConcurrency(d.Get("concurrency"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("concurrency"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyProp)) && (ok || !reflect.DeepEqual(v, concurrencyProp)) {
		obj["concurrency"] = concurrencyProp
	}
	editionProp, err := expandBigqueryReservationReservationEdition(d.Get("edition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edition"); !tpgresource.IsEmptyValue(reflect.ValueOf(editionProp)) && (ok || !reflect.DeepEqual(v, editionProp)) {
		obj["edition"] = editionProp
	}
	autoscaleProp, err := expandBigqueryReservationReservationAutoscale(d.Get("autoscale"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscale"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscaleProp)) && (ok || !reflect.DeepEqual(v, autoscaleProp)) {
		obj["autoscale"] = autoscaleProp
	}

	return obj, nil
}

func expandBigqueryReservationReservationSlotCapacity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationIgnoreIdleSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationConcurrency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationEdition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscale(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCurrentSlots, err := expandBigqueryReservationReservationAutoscaleCurrentSlots(original["current_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentSlots"] = transformedCurrentSlots
	}

	transformedMaxSlots, err := expandBigqueryReservationReservationAutoscaleMaxSlots(original["max_slots"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxSlots); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxSlots"] = transformedMaxSlots
	}

	return transformed, nil
}

func expandBigqueryReservationReservationAutoscaleCurrentSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAutoscaleMaxSlots(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
