// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package vertexai

import (
	"reflect"
	"regexp"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VertexAIFeaturestoreEntitytypeFeatureAssetType string = "aiplatform.googleapis.com/FeaturestoreEntitytypeFeature"

func ResourceConverterVertexAIFeaturestoreEntitytypeFeature() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VertexAIFeaturestoreEntitytypeFeatureAssetType,
		Convert:   GetVertexAIFeaturestoreEntitytypeFeatureCaiObject,
	}
}

func GetVertexAIFeaturestoreEntitytypeFeatureCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/{{entitytype}}/features/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVertexAIFeaturestoreEntitytypeFeatureApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VertexAIFeaturestoreEntitytypeFeatureAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "FeaturestoreEntitytypeFeature",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVertexAIFeaturestoreEntitytypeFeatureApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandVertexAIFeaturestoreEntitytypeFeatureDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	valueTypeProp, err := expandVertexAIFeaturestoreEntitytypeFeatureValueType(d.Get("value_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(valueTypeProp)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeFeatureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourceVertexAIFeaturestoreEntitytypeFeatureEncoder(d, config, obj)
}

func resourceVertexAIFeaturestoreEntitytypeFeatureEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := d.GetOk("entitytype"); ok {
		re := regexp.MustCompile("^projects/(.+)/locations/(.+)/featurestores/(.+)/entityTypes/(.+)$")
		if parts := re.FindStringSubmatch(v.(string)); parts != nil {
			d.Set("region", parts[2])
		}
	}

	return obj, nil
}

func expandVertexAIFeaturestoreEntitytypeFeatureDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeFeatureValueType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeFeatureEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
