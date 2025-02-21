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

package dialogflowcx

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXWebhookAssetType string = "{{location}}-dialogflow.googleapis.com/Webhook"

func ResourceConverterDialogflowCXWebhook() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXWebhookAssetType,
		Convert:   GetDialogflowCXWebhookCaiObject,
	}
}

func GetDialogflowCXWebhookCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/{{parent}}/webhooks/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXWebhookApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXWebhookAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "Webhook",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXWebhookApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXWebhookDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	timeoutProp, err := expandDialogflowCXWebhookTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeoutProp)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	disabledProp, err := expandDialogflowCXWebhookDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	genericWebServiceProp, err := expandDialogflowCXWebhookGenericWebService(d.Get("generic_web_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("generic_web_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(genericWebServiceProp)) && (ok || !reflect.DeepEqual(v, genericWebServiceProp)) {
		obj["genericWebService"] = genericWebServiceProp
	}
	serviceDirectoryProp, err := expandDialogflowCXWebhookServiceDirectory(d.Get("service_directory"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_directory"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceDirectoryProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryProp)) {
		obj["serviceDirectory"] = serviceDirectoryProp
	}
	securitySettingsProp, err := expandDialogflowCXWebhookSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXWebhookEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXWebhookEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	return obj, nil
}

func expandDialogflowCXWebhookDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookGenericWebService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowCXWebhookGenericWebServiceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedRequestHeaders, err := expandDialogflowCXWebhookGenericWebServiceRequestHeaders(original["request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestHeaders"] = transformedRequestHeaders
	}

	transformedAllowedCaCerts, err := expandDialogflowCXWebhookGenericWebServiceAllowedCaCerts(original["allowed_ca_certs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedCaCerts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedCaCerts"] = transformedAllowedCaCerts
	}

	return transformed, nil
}

func expandDialogflowCXWebhookGenericWebServiceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookGenericWebServiceRequestHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXWebhookGenericWebServiceAllowedCaCerts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandDialogflowCXWebhookServiceDirectoryService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedGenericWebService, err := expandDialogflowCXWebhookServiceDirectoryGenericWebService(original["generic_web_service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGenericWebService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["genericWebService"] = transformedGenericWebService
	}

	return transformed, nil
}

func expandDialogflowCXWebhookServiceDirectoryService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedRequestHeaders, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(original["request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestHeaders"] = transformedRequestHeaders
	}

	transformedAllowedCaCerts, err := expandDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(original["allowed_ca_certs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedCaCerts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedCaCerts"] = transformedAllowedCaCerts
	}

	return transformed, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceRequestHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXWebhookServiceDirectoryGenericWebServiceAllowedCaCerts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookSecuritySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXWebhookEnableSpellCorrection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
