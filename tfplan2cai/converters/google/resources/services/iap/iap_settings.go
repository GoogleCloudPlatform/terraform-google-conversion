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

package iap

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IapSettingsAssetType string = "iap.googleapis.com/Settings"

func ResourceConverterIapSettings() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IapSettingsAssetType,
		Convert:   GetIapSettingsCaiObject,
	}
}

func GetIapSettingsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iap.googleapis.com/{{name}}:iapSettings")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIapSettingsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IapSettingsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iap/v1/rest",
				DiscoveryName:        "Settings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIapSettingsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandIapSettingsName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	accessSettingsProp, err := expandIapSettingsAccessSettings(d.Get("access_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("access_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessSettingsProp)) && (ok || !reflect.DeepEqual(v, accessSettingsProp)) {
		obj["accessSettings"] = accessSettingsProp
	}
	applicationSettingsProp, err := expandIapSettingsApplicationSettings(d.Get("application_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationSettingsProp)) && (ok || !reflect.DeepEqual(v, applicationSettingsProp)) {
		obj["applicationSettings"] = applicationSettingsProp
	}

	return obj, nil
}

func expandIapSettingsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcipSettings, err := expandIapSettingsAccessSettingsGcipSettings(original["gcip_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcipSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcipSettings"] = transformedGcipSettings
	}

	transformedCorsSettings, err := expandIapSettingsAccessSettingsCorsSettings(original["cors_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCorsSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["corsSettings"] = transformedCorsSettings
	}

	transformedOauthSettings, err := expandIapSettingsAccessSettingsOauthSettings(original["oauth_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauthSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oauthSettings"] = transformedOauthSettings
	}

	transformedReauthSettings, err := expandIapSettingsAccessSettingsReauthSettings(original["reauth_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReauthSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["reauthSettings"] = transformedReauthSettings
	}

	transformedAllowedDomainsSettings, err := expandIapSettingsAccessSettingsAllowedDomainsSettings(original["allowed_domains_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedDomainsSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedDomainsSettings"] = transformedAllowedDomainsSettings
	}

	transformedWorkforceIdentitySettings, err := expandIapSettingsAccessSettingsWorkforceIdentitySettings(original["workforce_identity_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWorkforceIdentitySettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["workforceIdentitySettings"] = transformedWorkforceIdentitySettings
	}

	transformedIdentitySources, err := expandIapSettingsAccessSettingsIdentitySources(original["identity_sources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentitySources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["identitySources"] = transformedIdentitySources
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsGcipSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTenantIds, err := expandIapSettingsAccessSettingsGcipSettingsTenantIds(original["tenant_ids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTenantIds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tenantIds"] = transformedTenantIds
	}

	transformedLoginPageUri, err := expandIapSettingsAccessSettingsGcipSettingsLoginPageUri(original["login_page_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoginPageUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loginPageUri"] = transformedLoginPageUri
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsGcipSettingsTenantIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsGcipSettingsLoginPageUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsCorsSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowHttpOptions, err := expandIapSettingsAccessSettingsCorsSettingsAllowHttpOptions(original["allow_http_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowHttpOptions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowHttpOptions"] = transformedAllowHttpOptions
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsCorsSettingsAllowHttpOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsOauthSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLoginHint, err := expandIapSettingsAccessSettingsOauthSettingsLoginHint(original["login_hint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoginHint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loginHint"] = transformedLoginHint
	}

	transformedProgrammaticClients, err := expandIapSettingsAccessSettingsOauthSettingsProgrammaticClients(original["programmatic_clients"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProgrammaticClients); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["programmaticClients"] = transformedProgrammaticClients
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsOauthSettingsLoginHint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsOauthSettingsProgrammaticClients(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsReauthSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMethod, err := expandIapSettingsAccessSettingsReauthSettingsMethod(original["method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["method"] = transformedMethod
	}

	transformedMaxAge, err := expandIapSettingsAccessSettingsReauthSettingsMaxAge(original["max_age"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxAge); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxAge"] = transformedMaxAge
	}

	transformedPolicyType, err := expandIapSettingsAccessSettingsReauthSettingsPolicyType(original["policy_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicyType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["policyType"] = transformedPolicyType
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsReauthSettingsMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsReauthSettingsMaxAge(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsReauthSettingsPolicyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsAllowedDomainsSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDomains, err := expandIapSettingsAccessSettingsAllowedDomainsSettingsDomains(original["domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomains); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["domains"] = transformedDomains
	}

	transformedEnable, err := expandIapSettingsAccessSettingsAllowedDomainsSettingsEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsAllowedDomainsSettingsDomains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsAllowedDomainsSettingsEnable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWorkforcePools, err := expandIapSettingsAccessSettingsWorkforceIdentitySettingsWorkforcePools(original["workforce_pools"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWorkforcePools); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["workforcePools"] = transformedWorkforcePools
	}

	transformedOauth2, err := expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2(original["oauth2"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOauth2); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["oauth2"] = transformedOauth2
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettingsWorkforcePools(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClientId, err := expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	transformedClientSecret, err := expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientSecret(original["client_secret"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientSecret"] = transformedClientSecret
	}

	transformedClientSecretSha256, err := expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientSecretSha256(original["client_secret_sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientSecretSha256); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientSecretSha256"] = transformedClientSecretSha256
	}

	return transformed, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2ClientSecretSha256(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsAccessSettingsIdentitySources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCsmSettings, err := expandIapSettingsApplicationSettingsCsmSettings(original["csm_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCsmSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["csmSettings"] = transformedCsmSettings
	}

	transformedAccessDeniedPageSettings, err := expandIapSettingsApplicationSettingsAccessDeniedPageSettings(original["access_denied_page_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessDeniedPageSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessDeniedPageSettings"] = transformedAccessDeniedPageSettings
	}

	transformedCookieDomain, err := expandIapSettingsApplicationSettingsCookieDomain(original["cookie_domain"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCookieDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cookieDomain"] = transformedCookieDomain
	}

	transformedAttributePropagationSettings, err := expandIapSettingsApplicationSettingsAttributePropagationSettings(original["attribute_propagation_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributePropagationSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["attributePropagationSettings"] = transformedAttributePropagationSettings
	}

	return transformed, nil
}

func expandIapSettingsApplicationSettingsCsmSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRctokenAud, err := expandIapSettingsApplicationSettingsCsmSettingsRctokenAud(original["rctoken_aud"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRctokenAud); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rctokenAud"] = transformedRctokenAud
	}

	return transformed, nil
}

func expandIapSettingsApplicationSettingsCsmSettingsRctokenAud(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAccessDeniedPageSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAccessDeniedPageUri, err := expandIapSettingsApplicationSettingsAccessDeniedPageSettingsAccessDeniedPageUri(original["access_denied_page_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessDeniedPageUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessDeniedPageUri"] = transformedAccessDeniedPageUri
	}

	transformedGenerateTroubleshootingUri, err := expandIapSettingsApplicationSettingsAccessDeniedPageSettingsGenerateTroubleshootingUri(original["generate_troubleshooting_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGenerateTroubleshootingUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["generateTroubleshootingUri"] = transformedGenerateTroubleshootingUri
	}

	transformedRemediationTokenGenerationEnabled, err := expandIapSettingsApplicationSettingsAccessDeniedPageSettingsRemediationTokenGenerationEnabled(original["remediation_token_generation_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRemediationTokenGenerationEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["remediationTokenGenerationEnabled"] = transformedRemediationTokenGenerationEnabled
	}

	return transformed, nil
}

func expandIapSettingsApplicationSettingsAccessDeniedPageSettingsAccessDeniedPageUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAccessDeniedPageSettingsGenerateTroubleshootingUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAccessDeniedPageSettingsRemediationTokenGenerationEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsCookieDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAttributePropagationSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOutputCredentials, err := expandIapSettingsApplicationSettingsAttributePropagationSettingsOutputCredentials(original["output_credentials"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputCredentials); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["outputCredentials"] = transformedOutputCredentials
	}

	transformedExpression, err := expandIapSettingsApplicationSettingsAttributePropagationSettingsExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedEnable, err := expandIapSettingsApplicationSettingsAttributePropagationSettingsEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandIapSettingsApplicationSettingsAttributePropagationSettingsOutputCredentials(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAttributePropagationSettingsExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapSettingsApplicationSettingsAttributePropagationSettingsEnable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
