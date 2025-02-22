// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dialogflowcx/Intent.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dialogflowcx

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXIntentAssetType string = "{{location}}-dialogflow.googleapis.com/Intent"

func ResourceConverterDialogflowCXIntent() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXIntentAssetType,
		Convert:   GetDialogflowCXIntentCaiObject,
	}
}

func GetDialogflowCXIntentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/{{parent}}/intents/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXIntentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXIntentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "Intent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXIntentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXIntentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	trainingPhrasesProp, err := expandDialogflowCXIntentTrainingPhrases(d.Get("training_phrases"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("training_phrases"); !tpgresource.IsEmptyValue(reflect.ValueOf(trainingPhrasesProp)) && (ok || !reflect.DeepEqual(v, trainingPhrasesProp)) {
		obj["trainingPhrases"] = trainingPhrasesProp
	}
	parametersProp, err := expandDialogflowCXIntentParameters(d.Get("parameters"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parameters"); !tpgresource.IsEmptyValue(reflect.ValueOf(parametersProp)) && (ok || !reflect.DeepEqual(v, parametersProp)) {
		obj["parameters"] = parametersProp
	}
	priorityProp, err := expandDialogflowCXIntentPriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	isFallbackProp, err := expandDialogflowCXIntentIsFallback(d.Get("is_fallback"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_fallback"); !tpgresource.IsEmptyValue(reflect.ValueOf(isFallbackProp)) && (ok || !reflect.DeepEqual(v, isFallbackProp)) {
		obj["isFallback"] = isFallbackProp
	}
	descriptionProp, err := expandDialogflowCXIntentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandDialogflowCXIntentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	languageCodeProp, err := expandDialogflowCXIntentLanguageCode(d.Get("language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("language_code"); !tpgresource.IsEmptyValue(reflect.ValueOf(languageCodeProp)) && (ok || !reflect.DeepEqual(v, languageCodeProp)) {
		obj["languageCode"] = languageCodeProp
	}

	return obj, nil
}

func expandDialogflowCXIntentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandDialogflowCXIntentTrainingPhrasesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedParts, err := expandDialogflowCXIntentTrainingPhrasesParts(original["parts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["parts"] = transformedParts
		}

		transformedRepeatCount, err := expandDialogflowCXIntentTrainingPhrasesRepeatCount(original["repeat_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRepeatCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["repeatCount"] = transformedRepeatCount
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentTrainingPhrasesId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesParts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXIntentTrainingPhrasesPartsText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["text"] = transformedText
		}

		transformedParameterId, err := expandDialogflowCXIntentTrainingPhrasesPartsParameterId(original["parameter_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParameterId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["parameterId"] = transformedParameterId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentTrainingPhrasesPartsText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesPartsParameterId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentTrainingPhrasesRepeatCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandDialogflowCXIntentParametersId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedEntityType, err := expandDialogflowCXIntentParametersEntityType(original["entity_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEntityType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["entityType"] = transformedEntityType
		}

		transformedIsList, err := expandDialogflowCXIntentParametersIsList(original["is_list"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsList); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["isList"] = transformedIsList
		}

		transformedRedact, err := expandDialogflowCXIntentParametersRedact(original["redact"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRedact); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["redact"] = transformedRedact
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXIntentParametersId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersEntityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersIsList(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentParametersRedact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentPriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentIsFallback(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXIntentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXIntentLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
