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

package kms

import (
	"fmt"
	"reflect"
	"time"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const KMSCryptoKeyAssetType string = "cloudkms.googleapis.com/CryptoKey"

func ResourceConverterKMSCryptoKey() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: KMSCryptoKeyAssetType,
		Convert:   GetKMSCryptoKeyCaiObject,
	}
}

func GetKMSCryptoKeyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudkms.googleapis.com/{{key_ring}}/cryptoKeys/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetKMSCryptoKeyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: KMSCryptoKeyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "CryptoKey",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetKMSCryptoKeyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	purposeProp, err := expandKMSCryptoKeyPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !tpgresource.IsEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	rotationPeriodProp, err := expandKMSCryptoKeyRotationPeriod(d.Get("rotation_period"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation_period"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationPeriodProp)) && (ok || !reflect.DeepEqual(v, rotationPeriodProp)) {
		obj["rotationPeriod"] = rotationPeriodProp
	}
	versionTemplateProp, err := expandKMSCryptoKeyVersionTemplate(d.Get("version_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionTemplateProp)) && (ok || !reflect.DeepEqual(v, versionTemplateProp)) {
		obj["versionTemplate"] = versionTemplateProp
	}
	destroyScheduledDurationProp, err := expandKMSCryptoKeyDestroyScheduledDuration(d.Get("destroy_scheduled_duration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destroy_scheduled_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(destroyScheduledDurationProp)) && (ok || !reflect.DeepEqual(v, destroyScheduledDurationProp)) {
		obj["destroyScheduledDuration"] = destroyScheduledDurationProp
	}
	importOnlyProp, err := expandKMSCryptoKeyImportOnly(d.Get("import_only"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_only"); !tpgresource.IsEmptyValue(reflect.ValueOf(importOnlyProp)) && (ok || !reflect.DeepEqual(v, importOnlyProp)) {
		obj["importOnly"] = importOnlyProp
	}
	labelsProp, err := expandKMSCryptoKeyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourceKMSCryptoKeyEncoder(d, config, obj)
}

func resourceKMSCryptoKeyEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// if rotationPeriod is set, nextRotationTime must also be set.
	if d.Get("rotation_period") != "" {
		rotationPeriod := d.Get("rotation_period").(string)
		nextRotation, err := kmsCryptoKeyNextRotation(time.Now(), rotationPeriod)

		if err != nil {
			return nil, fmt.Errorf("Error setting CryptoKey rotation period: %s", err.Error())
		}

		obj["nextRotationTime"] = nextRotation
	}

	// set to false if it is not true explicitly
	if !(d.Get("skip_initial_version_creation").(bool)) {
		if err := d.Set("skip_initial_version_creation", false); err != nil {
			return nil, fmt.Errorf("Error setting skip_initial_version_creation: %s", err)
		}
	}

	return obj, nil
}

func expandKMSCryptoKeyPurpose(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyRotationPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyVersionTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAlgorithm, err := expandKMSCryptoKeyVersionTemplateAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	transformedProtectionLevel, err := expandKMSCryptoKeyVersionTemplateProtectionLevel(original["protection_level"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProtectionLevel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["protectionLevel"] = transformedProtectionLevel
	}

	return transformed, nil
}

func expandKMSCryptoKeyVersionTemplateAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyVersionTemplateProtectionLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyDestroyScheduledDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyImportOnly(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSCryptoKeyEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
