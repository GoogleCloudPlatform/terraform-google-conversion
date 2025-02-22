// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/certificatemanager/CertificateIssuanceConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package certificatemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CertificateManagerCertificateIssuanceConfigAssetType string = "certificatemanager.googleapis.com/CertificateIssuanceConfig"

func ResourceConverterCertificateManagerCertificateIssuanceConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CertificateManagerCertificateIssuanceConfigAssetType,
		Convert:   GetCertificateManagerCertificateIssuanceConfigCaiObject,
	}
}

func GetCertificateManagerCertificateIssuanceConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCertificateManagerCertificateIssuanceConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CertificateManagerCertificateIssuanceConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "CertificateIssuanceConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCertificateManagerCertificateIssuanceConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateIssuanceConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	rotationWindowPercentageProp, err := expandCertificateManagerCertificateIssuanceConfigRotationWindowPercentage(d.Get("rotation_window_percentage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rotation_window_percentage"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationWindowPercentageProp)) && (ok || !reflect.DeepEqual(v, rotationWindowPercentageProp)) {
		obj["rotationWindowPercentage"] = rotationWindowPercentageProp
	}
	keyAlgorithmProp, err := expandCertificateManagerCertificateIssuanceConfigKeyAlgorithm(d.Get("key_algorithm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyAlgorithmProp)) && (ok || !reflect.DeepEqual(v, keyAlgorithmProp)) {
		obj["keyAlgorithm"] = keyAlgorithmProp
	}
	lifetimeProp, err := expandCertificateManagerCertificateIssuanceConfigLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !tpgresource.IsEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	certificateAuthorityConfigProp, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig(d.Get("certificate_authority_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_authority_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateAuthorityConfigProp)) && (ok || !reflect.DeepEqual(v, certificateAuthorityConfigProp)) {
		obj["certificateAuthorityConfig"] = certificateAuthorityConfigProp
	}
	labelsProp, err := expandCertificateManagerCertificateIssuanceConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandCertificateManagerCertificateIssuanceConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigRotationWindowPercentage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigKeyAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigLifetime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateAuthorityServiceConfig, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig(original["certificate_authority_service_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateAuthorityServiceConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateAuthorityServiceConfig"] = transformedCertificateAuthorityServiceConfig
	}

	return transformed, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCaPool, err := expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigCaPool(original["ca_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaPool); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["caPool"] = transformedCaPool
	}

	return transformed, nil
}

func expandCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigCaPool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateIssuanceConfigEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
