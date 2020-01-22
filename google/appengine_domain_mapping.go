// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import "reflect"

func GetAppEngineDomainMappingCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//appengine.googleapis.com/apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetAppEngineDomainMappingApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "appengine.googleapis.com/DomainMapping",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest",
				DiscoveryName:        "DomainMapping",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetAppEngineDomainMappingApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	sslSettingsProp, err := expandAppEngineDomainMappingSslSettings(d.Get("ssl_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_settings"); !isEmptyValue(reflect.ValueOf(sslSettingsProp)) && (ok || !reflect.DeepEqual(v, sslSettingsProp)) {
		obj["sslSettings"] = sslSettingsProp
	}
	idProp, err := expandAppEngineDomainMappingDomainName(d.Get("domain_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain_name"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}

	return obj, nil
}

func expandAppEngineDomainMappingSslSettings(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateId, err := expandAppEngineDomainMappingSslSettingsCertificateId(original["certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateId); val.IsValid() && !isEmptyValue(val) {
		transformed["certificateId"] = transformedCertificateId
	}

	transformedSslManagementType, err := expandAppEngineDomainMappingSslSettingsSslManagementType(original["ssl_management_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslManagementType); val.IsValid() && !isEmptyValue(val) {
		transformed["sslManagementType"] = transformedSslManagementType
	}

	transformedPendingManagedCertificateId, err := expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(original["pending_managed_certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPendingManagedCertificateId); val.IsValid() && !isEmptyValue(val) {
		transformed["pendingManagedCertificateId"] = transformedPendingManagedCertificateId
	}

	return transformed, nil
}

func expandAppEngineDomainMappingSslSettingsCertificateId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsSslManagementType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingDomainName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
