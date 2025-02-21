// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securityscanner/ScanConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securityscanner

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityScannerScanConfigAssetType string = "websecurityscanner.googleapis.com/ScanConfig"

func ResourceConverterSecurityScannerScanConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityScannerScanConfigAssetType,
		Convert:   GetSecurityScannerScanConfigCaiObject,
	}
}

func GetSecurityScannerScanConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//websecurityscanner.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityScannerScanConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityScannerScanConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/websecurityscanner/v1beta/rest",
				DiscoveryName:        "ScanConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityScannerScanConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityScannerScanConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	maxQpsProp, err := expandSecurityScannerScanConfigMaxQps(d.Get("max_qps"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("max_qps"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxQpsProp)) && (ok || !reflect.DeepEqual(v, maxQpsProp)) {
		obj["maxQps"] = maxQpsProp
	}
	startingUrlsProp, err := expandSecurityScannerScanConfigStartingUrls(d.Get("starting_urls"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("starting_urls"); !tpgresource.IsEmptyValue(reflect.ValueOf(startingUrlsProp)) && (ok || !reflect.DeepEqual(v, startingUrlsProp)) {
		obj["startingUrls"] = startingUrlsProp
	}
	authenticationProp, err := expandSecurityScannerScanConfigAuthentication(d.Get("authentication"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authentication"); !tpgresource.IsEmptyValue(reflect.ValueOf(authenticationProp)) && (ok || !reflect.DeepEqual(v, authenticationProp)) {
		obj["authentication"] = authenticationProp
	}
	userAgentProp, err := expandSecurityScannerScanConfigUserAgent(d.Get("user_agent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_agent"); !tpgresource.IsEmptyValue(reflect.ValueOf(userAgentProp)) && (ok || !reflect.DeepEqual(v, userAgentProp)) {
		obj["userAgent"] = userAgentProp
	}
	blacklistPatternsProp, err := expandSecurityScannerScanConfigBlacklistPatterns(d.Get("blacklist_patterns"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("blacklist_patterns"); !tpgresource.IsEmptyValue(reflect.ValueOf(blacklistPatternsProp)) && (ok || !reflect.DeepEqual(v, blacklistPatternsProp)) {
		obj["blacklistPatterns"] = blacklistPatternsProp
	}
	scheduleProp, err := expandSecurityScannerScanConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	targetPlatformsProp, err := expandSecurityScannerScanConfigTargetPlatforms(d.Get("target_platforms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_platforms"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetPlatformsProp)) && (ok || !reflect.DeepEqual(v, targetPlatformsProp)) {
		obj["targetPlatforms"] = targetPlatformsProp
	}
	exportToSecurityCommandCenterProp, err := expandSecurityScannerScanConfigExportToSecurityCommandCenter(d.Get("export_to_security_command_center"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_to_security_command_center"); !tpgresource.IsEmptyValue(reflect.ValueOf(exportToSecurityCommandCenterProp)) && (ok || !reflect.DeepEqual(v, exportToSecurityCommandCenterProp)) {
		obj["exportToSecurityCommandCenter"] = exportToSecurityCommandCenterProp
	}

	return obj, nil
}

func expandSecurityScannerScanConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigMaxQps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigStartingUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthentication(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGoogleAccount, err := expandSecurityScannerScanConfigAuthenticationGoogleAccount(original["google_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGoogleAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["googleAccount"] = transformedGoogleAccount
	}

	transformedCustomAccount, err := expandSecurityScannerScanConfigAuthenticationCustomAccount(original["custom_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customAccount"] = transformedCustomAccount
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandSecurityScannerScanConfigAuthenticationGoogleAccountUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandSecurityScannerScanConfigAuthenticationGoogleAccountPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccountUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccountPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandSecurityScannerScanConfigAuthenticationCustomAccountUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandSecurityScannerScanConfigAuthenticationCustomAccountPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedLoginUrl, err := expandSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(original["login_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoginUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loginUrl"] = transformedLoginUrl
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigUserAgent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigBlacklistPatterns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScheduleTime, err := expandSecurityScannerScanConfigScheduleScheduleTime(original["schedule_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScheduleTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scheduleTime"] = transformedScheduleTime
	}

	transformedIntervalDurationDays, err := expandSecurityScannerScanConfigScheduleIntervalDurationDays(original["interval_duration_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIntervalDurationDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["intervalDurationDays"] = transformedIntervalDurationDays
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigScheduleScheduleTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigScheduleIntervalDurationDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigTargetPlatforms(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigExportToSecurityCommandCenter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
