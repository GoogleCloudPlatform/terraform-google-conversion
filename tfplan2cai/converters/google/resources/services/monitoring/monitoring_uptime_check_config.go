// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/monitoring/UptimeCheckConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package monitoring

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourceMonitoringUptimeCheckConfigHttpCheckPathDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return old == "/"+new
}

func resourceMonitoringUptimeCheckConfigMonitoredResourceLabelsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// GCP adds the project_id to the labels if unset.
	// We want to suppress the diff if not set in the config.
	if strings.HasSuffix(k, "project_id") && new == "" && old != "" {
		return true
	}
	return false
}

const MonitoringUptimeCheckConfigAssetType string = "monitoring.googleapis.com/UptimeCheckConfig"

func ResourceConverterMonitoringUptimeCheckConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MonitoringUptimeCheckConfigAssetType,
		Convert:   GetMonitoringUptimeCheckConfigCaiObject,
	}
}

func GetMonitoringUptimeCheckConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMonitoringUptimeCheckConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MonitoringUptimeCheckConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "UptimeCheckConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMonitoringUptimeCheckConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringUptimeCheckConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	periodProp, err := expandMonitoringUptimeCheckConfigPeriod(d.Get("period"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("period"); !tpgresource.IsEmptyValue(reflect.ValueOf(periodProp)) && (ok || !reflect.DeepEqual(v, periodProp)) {
		obj["period"] = periodProp
	}
	timeoutProp, err := expandMonitoringUptimeCheckConfigTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeoutProp)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	contentMatchersProp, err := expandMonitoringUptimeCheckConfigContentMatchers(d.Get("content_matchers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("content_matchers"); !tpgresource.IsEmptyValue(reflect.ValueOf(contentMatchersProp)) && (ok || !reflect.DeepEqual(v, contentMatchersProp)) {
		obj["contentMatchers"] = contentMatchersProp
	}
	selectedRegionsProp, err := expandMonitoringUptimeCheckConfigSelectedRegions(d.Get("selected_regions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("selected_regions"); !tpgresource.IsEmptyValue(reflect.ValueOf(selectedRegionsProp)) && (ok || !reflect.DeepEqual(v, selectedRegionsProp)) {
		obj["selectedRegions"] = selectedRegionsProp
	}
	logCheckFailuresProp, err := expandMonitoringUptimeCheckConfigLogCheckFailures(d.Get("log_check_failures"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_check_failures"); !tpgresource.IsEmptyValue(reflect.ValueOf(logCheckFailuresProp)) && (ok || !reflect.DeepEqual(v, logCheckFailuresProp)) {
		obj["logCheckFailures"] = logCheckFailuresProp
	}
	checkerTypeProp, err := expandMonitoringUptimeCheckConfigCheckerType(d.Get("checker_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("checker_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(checkerTypeProp)) && (ok || !reflect.DeepEqual(v, checkerTypeProp)) {
		obj["checkerType"] = checkerTypeProp
	}
	userLabelsProp, err := expandMonitoringUptimeCheckConfigUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(userLabelsProp)) && (ok || !reflect.DeepEqual(v, userLabelsProp)) {
		obj["userLabels"] = userLabelsProp
	}
	httpCheckProp, err := expandMonitoringUptimeCheckConfigHttpCheck(d.Get("http_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpCheckProp)) && (ok || !reflect.DeepEqual(v, httpCheckProp)) {
		obj["httpCheck"] = httpCheckProp
	}
	tcpCheckProp, err := expandMonitoringUptimeCheckConfigTcpCheck(d.Get("tcp_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tcp_check"); !tpgresource.IsEmptyValue(reflect.ValueOf(tcpCheckProp)) && (ok || !reflect.DeepEqual(v, tcpCheckProp)) {
		obj["tcpCheck"] = tcpCheckProp
	}
	resourceGroupProp, err := expandMonitoringUptimeCheckConfigResourceGroup(d.Get("resource_group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceGroupProp)) && (ok || !reflect.DeepEqual(v, resourceGroupProp)) {
		obj["resourceGroup"] = resourceGroupProp
	}
	monitoredResourceProp, err := expandMonitoringUptimeCheckConfigMonitoredResource(d.Get("monitored_resource"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("monitored_resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(monitoredResourceProp)) && (ok || !reflect.DeepEqual(v, monitoredResourceProp)) {
		obj["monitoredResource"] = monitoredResourceProp
	}
	syntheticMonitorProp, err := expandMonitoringUptimeCheckConfigSyntheticMonitor(d.Get("synthetic_monitor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("synthetic_monitor"); !tpgresource.IsEmptyValue(reflect.ValueOf(syntheticMonitorProp)) && (ok || !reflect.DeepEqual(v, syntheticMonitorProp)) {
		obj["syntheticMonitor"] = syntheticMonitorProp
	}

	return resourceMonitoringUptimeCheckConfigEncoder(d, config, obj)
}

func resourceMonitoringUptimeCheckConfigEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// remove passwordWoVersion from the request body
	if v, ok := obj["httpCheck"]; ok {
		httpCheck := v.(map[string]interface{})
		if authInfo, ok := httpCheck["authInfo"].(map[string]interface{}); ok {
			delete(authInfo, "passwordWoVersion")
			if len(authInfo) > 0 {
				httpCheck["authInfo"] = authInfo
			} else {
				delete(httpCheck, "authInfo")
			}
			obj["httpCheck"] = httpCheck
		}
	}

	return obj, nil
}

func expandMonitoringUptimeCheckConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigContentMatchers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedContent, err := expandMonitoringUptimeCheckConfigContentMatchersContent(original["content"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["content"] = transformedContent
		}

		transformedMatcher, err := expandMonitoringUptimeCheckConfigContentMatchersMatcher(original["matcher"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatcher); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["matcher"] = transformedMatcher
		}

		transformedJsonPathMatcher, err := expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcher(original["json_path_matcher"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedJsonPathMatcher); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["jsonPathMatcher"] = transformedJsonPathMatcher
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersContent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedJsonPath, err := expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcherJsonPath(original["json_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJsonPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jsonPath"] = transformedJsonPath
	}

	transformedJsonMatcher, err := expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcherJsonMatcher(original["json_matcher"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJsonMatcher); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jsonMatcher"] = transformedJsonMatcher
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcherJsonPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersJsonPathMatcherJsonMatcher(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigSelectedRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigLogCheckFailures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigCheckerType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigUserLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringUptimeCheckConfigHttpCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequestMethod, err := expandMonitoringUptimeCheckConfigHttpCheckRequestMethod(original["request_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requestMethod"] = transformedRequestMethod
	}

	transformedContentType, err := expandMonitoringUptimeCheckConfigHttpCheckContentType(original["content_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContentType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["contentType"] = transformedContentType
	}

	transformedCustomContentType, err := expandMonitoringUptimeCheckConfigHttpCheckCustomContentType(original["custom_content_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomContentType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customContentType"] = transformedCustomContentType
	}

	transformedAuthInfo, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfo(original["auth_info"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthInfo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authInfo"] = transformedAuthInfo
	}

	transformedServiceAgentAuthentication, err := expandMonitoringUptimeCheckConfigHttpCheckServiceAgentAuthentication(original["service_agent_authentication"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAgentAuthentication); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAgentAuthentication"] = transformedServiceAgentAuthentication
	}

	transformedPort, err := expandMonitoringUptimeCheckConfigHttpCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedHeaders, err := expandMonitoringUptimeCheckConfigHttpCheckHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	transformedPath, err := expandMonitoringUptimeCheckConfigHttpCheckPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedUseSsl, err := expandMonitoringUptimeCheckConfigHttpCheckUseSsl(original["use_ssl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUseSsl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["useSsl"] = transformedUseSsl
	}

	transformedValidateSsl, err := expandMonitoringUptimeCheckConfigHttpCheckValidateSsl(original["validate_ssl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValidateSsl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["validateSsl"] = transformedValidateSsl
	}

	transformedMaskHeaders, err := expandMonitoringUptimeCheckConfigHttpCheckMaskHeaders(original["mask_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaskHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maskHeaders"] = transformedMaskHeaders
	}

	transformedBody, err := expandMonitoringUptimeCheckConfigHttpCheckBody(original["body"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBody); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["body"] = transformedBody
	}

	transformedAcceptedResponseStatusCodes, err := expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodes(original["accepted_response_status_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAcceptedResponseStatusCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["acceptedResponseStatusCodes"] = transformedAcceptedResponseStatusCodes
	}

	transformedPingConfig, err := expandMonitoringUptimeCheckConfigHttpCheckPingConfig(original["ping_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pingConfig"] = transformedPingConfig
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckRequestMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckContentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckCustomContentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPassword, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedPasswordWo, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPasswordWo(original["password_wo"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordWo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordWo"] = transformedPasswordWo
	}

	transformedPasswordWoVersion, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPasswordWoVersion(original["password_wo_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordWoVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordWoVersion"] = transformedPasswordWoVersion
	}

	transformedUsername, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPasswordWo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPasswordWoVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckServiceAgentAuthentication(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandMonitoringUptimeCheckConfigHttpCheckServiceAgentAuthenticationType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckServiceAgentAuthenticationType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckUseSsl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckValidateSsl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckMaskHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckBody(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedStatusValue, err := expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesStatusValue(original["status_value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatusValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["statusValue"] = transformedStatusValue
		}

		transformedStatusClass, err := expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesStatusClass(original["status_class"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatusClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["statusClass"] = transformedStatusClass
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesStatusValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesStatusClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPingsCount, err := expandMonitoringUptimeCheckConfigHttpCheckPingConfigPingsCount(original["pings_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPingsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pingsCount"] = transformedPingsCount
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPingConfigPingsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigTcpCheck(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPort, err := expandMonitoringUptimeCheckConfigTcpCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPingConfig, err := expandMonitoringUptimeCheckConfigTcpCheckPingConfig(original["ping_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pingConfig"] = transformedPingConfig
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigTcpCheckPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigTcpCheckPingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPingsCount, err := expandMonitoringUptimeCheckConfigTcpCheckPingConfigPingsCount(original["pings_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPingsCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pingsCount"] = transformedPingsCount
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigTcpCheckPingConfigPingsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigResourceGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceType, err := expandMonitoringUptimeCheckConfigResourceGroupResourceType(original["resource_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceType"] = transformedResourceType
	}

	transformedGroupId, err := expandMonitoringUptimeCheckConfigResourceGroupGroupId(original["group_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["groupId"] = transformedGroupId
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigResourceGroupResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigResourceGroupGroupId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandMonitoringUptimeCheckConfigMonitoredResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandMonitoringUptimeCheckConfigMonitoredResourceType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedLabels, err := expandMonitoringUptimeCheckConfigMonitoredResourceLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigMonitoredResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigMonitoredResourceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringUptimeCheckConfigSyntheticMonitor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCloudFunctionV2, err := expandMonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2(original["cloud_function_v2"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudFunctionV2); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudFunctionV2"] = transformedCloudFunctionV2
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandMonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2Name(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2Name(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
