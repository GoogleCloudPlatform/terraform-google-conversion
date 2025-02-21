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

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeKeystoresAliasesSelfSignedCertAssetType string = "apigee.googleapis.com/KeystoresAliasesSelfSignedCert"

func ResourceConverterApigeeKeystoresAliasesSelfSignedCert() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeKeystoresAliasesSelfSignedCertAssetType,
		Convert:   GetApigeeKeystoresAliasesSelfSignedCertCaiObject,
	}
}

func GetApigeeKeystoresAliasesSelfSignedCertCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeKeystoresAliasesSelfSignedCertApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeKeystoresAliasesSelfSignedCertAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "KeystoresAliasesSelfSignedCert",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeKeystoresAliasesSelfSignedCertApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	aliasProp, err := expandApigeeKeystoresAliasesSelfSignedCertAlias(d.Get("alias"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("alias"); !tpgresource.IsEmptyValue(reflect.ValueOf(aliasProp)) && (ok || !reflect.DeepEqual(v, aliasProp)) {
		obj["alias"] = aliasProp
	}
	subjectAlternativeDnsNamesProp, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(d.Get("subject_alternative_dns_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subject_alternative_dns_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(subjectAlternativeDnsNamesProp)) && (ok || !reflect.DeepEqual(v, subjectAlternativeDnsNamesProp)) {
		obj["subjectAlternativeDnsNames"] = subjectAlternativeDnsNamesProp
	}
	keySizeProp, err := expandApigeeKeystoresAliasesSelfSignedCertKeySize(d.Get("key_size"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_size"); !tpgresource.IsEmptyValue(reflect.ValueOf(keySizeProp)) && (ok || !reflect.DeepEqual(v, keySizeProp)) {
		obj["keySize"] = keySizeProp
	}
	sigAlgProp, err := expandApigeeKeystoresAliasesSelfSignedCertSigAlg(d.Get("sig_alg"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sig_alg"); !tpgresource.IsEmptyValue(reflect.ValueOf(sigAlgProp)) && (ok || !reflect.DeepEqual(v, sigAlgProp)) {
		obj["sigAlg"] = sigAlgProp
	}
	subjectProp, err := expandApigeeKeystoresAliasesSelfSignedCertSubject(d.Get("subject"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subject"); !tpgresource.IsEmptyValue(reflect.ValueOf(subjectProp)) && (ok || !reflect.DeepEqual(v, subjectProp)) {
		obj["subject"] = subjectProp
	}
	certValidityInDaysProp, err := expandApigeeKeystoresAliasesSelfSignedCertCertValidityInDays(d.Get("cert_validity_in_days"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cert_validity_in_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(certValidityInDaysProp)) && (ok || !reflect.DeepEqual(v, certValidityInDaysProp)) {
		obj["certValidityInDays"] = certValidityInDaysProp
	}

	return obj, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertAlias(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubjectAlternativeName, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(original["subject_alternative_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAlternativeName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectAlternativeName"] = transformedSubjectAlternativeName
	}

	return transformed, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesSubjectAlternativeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertKeySize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSigAlg(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedState, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedLocality, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedOrg, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectOrg(original["org"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrg); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["org"] = transformedOrg
	}

	transformedOrgUnit, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectOrgUnit(original["org_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrgUnit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["orgUnit"] = transformedOrgUnit
	}

	transformedCommonName, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	transformedEmail, err := expandApigeeKeystoresAliasesSelfSignedCertSubjectEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	return transformed, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectCountryCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectOrg(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectOrgUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectCommonName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertSubjectEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeKeystoresAliasesSelfSignedCertCertValidityInDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
