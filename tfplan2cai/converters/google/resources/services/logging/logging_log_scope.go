// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/logging/LogScope.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package logging

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const LoggingLogScopeAssetType string = "logging.googleapis.com/LogScope"

func ResourceConverterLoggingLogScope() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingLogScopeAssetType,
		Convert:   GetLoggingLogScopeCaiObject,
	}
}

func GetLoggingLogScopeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//logging.googleapis.com/{{parent}}/locations/{{location}}/logScopes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetLoggingLogScopeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: LoggingLogScopeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "LogScope",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetLoggingLogScopeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandLoggingLogScopeName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	resourceNamesProp, err := expandLoggingLogScopeResourceNames(d.Get("resource_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceNamesProp)) && (ok || !reflect.DeepEqual(v, resourceNamesProp)) {
		obj["resourceNames"] = resourceNamesProp
	}
	descriptionProp, err := expandLoggingLogScopeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return resourceLoggingLogScopeEncoder(d, config, obj)
}

func resourceLoggingLogScopeEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Extract any empty fields from the bucket field.
	// Extract parent, location from name
	parent := d.Get("parent").(string)
	name := d.Get("name").(string)
	parent, err := tpgresource.ExtractFieldByPattern("parent", parent, name, "((projects|folders|organizations|billingAccounts)/[a-z0-9A-Z-]*)/locations/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting parent field: %s", err)
	}
	location := d.Get("location").(string)
	location, err = tpgresource.ExtractFieldByPattern("location", location, name, "[a-zA-Z]*/[a-z0-9A-Z-]*/locations/([a-z0-9-]*)/logScopes/.*")
	if err != nil {
		return nil, fmt.Errorf("error extracting location field: %s", err)
	}
	// Set parent to the extracted value.
	d.Set("parent", parent)
	// Set all the other fields to their short forms before forming url and setting ID.
	name = tpgresource.GetResourceNameFromSelfLink(name)
	d.Set("location", location)
	d.Set("name", name)
	return obj, nil
}

func expandLoggingLogScopeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogScopeResourceNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingLogScopeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
