// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/eventarc/Enrollment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package eventarc

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const EventarcEnrollmentAssetType string = "eventarc.googleapis.com/Enrollment"

func ResourceConverterEventarcEnrollment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: EventarcEnrollmentAssetType,
		Convert:   GetEventarcEnrollmentCaiObject,
	}
}

func GetEventarcEnrollmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//eventarc.googleapis.com/projects/{{project}}/locations/{{location}}/enrollments/{{enrollment_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetEventarcEnrollmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: EventarcEnrollmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/eventarc/v1/rest",
				DiscoveryName:        "Enrollment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetEventarcEnrollmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandEventarcEnrollmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	messageBusProp, err := expandEventarcEnrollmentMessageBus(d.Get("message_bus"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_bus"); !tpgresource.IsEmptyValue(reflect.ValueOf(messageBusProp)) && (ok || !reflect.DeepEqual(v, messageBusProp)) {
		obj["messageBus"] = messageBusProp
	}
	celMatchProp, err := expandEventarcEnrollmentCelMatch(d.Get("cel_match"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cel_match"); !tpgresource.IsEmptyValue(reflect.ValueOf(celMatchProp)) && (ok || !reflect.DeepEqual(v, celMatchProp)) {
		obj["celMatch"] = celMatchProp
	}
	destinationProp, err := expandEventarcEnrollmentDestination(d.Get("destination"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationProp)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	labelsProp, err := expandEventarcEnrollmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandEventarcEnrollmentEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandEventarcEnrollmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcEnrollmentMessageBus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcEnrollmentCelMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcEnrollmentDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEventarcEnrollmentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandEventarcEnrollmentEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
