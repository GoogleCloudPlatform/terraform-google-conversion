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

package colab

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ColabRuntimeAssetType string = "aiplatform.googleapis.com/Runtime"

func ResourceConverterColabRuntime() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ColabRuntimeAssetType,
		Convert:   GetColabRuntimeCaiObject,
	}
}

func GetColabRuntimeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetColabRuntimeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ColabRuntimeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "Runtime",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetColabRuntimeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	notebookRuntimeTemplateRefProp, err := expandColabRuntimeNotebookRuntimeTemplateRef(d.Get("notebook_runtime_template_ref"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notebook_runtime_template_ref"); !tpgresource.IsEmptyValue(reflect.ValueOf(notebookRuntimeTemplateRefProp)) && (ok || !reflect.DeepEqual(v, notebookRuntimeTemplateRefProp)) {
		obj["notebookRuntimeTemplateRef"] = notebookRuntimeTemplateRefProp
	}
	runtimeUserProp, err := expandColabRuntimeRuntimeUser(d.Get("runtime_user"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_user"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeUserProp)) && (ok || !reflect.DeepEqual(v, runtimeUserProp)) {
		obj["runtimeUser"] = runtimeUserProp
	}
	displayNameProp, err := expandColabRuntimeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandColabRuntimeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return resourceColabRuntimeEncoder(d, config, obj)
}

func resourceColabRuntimeEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["notebookRuntimeTemplate"], _ = d.GetOk("notebook_runtime_template_ref.0.notebook_runtime_template")

	delete(obj, "notebookRuntimeTemplateRef")

	newObj["notebookRuntime"] = obj
	return newObj, nil
}

func expandColabRuntimeNotebookRuntimeTemplateRef(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNotebookRuntimeTemplate, err := expandColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(original["notebook_runtime_template"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNotebookRuntimeTemplate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["notebookRuntimeTemplate"] = transformedNotebookRuntimeTemplate
	}

	return transformed, nil
}

func expandColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeRuntimeUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
