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

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func GetComputeRegionSslCertificateCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeRegionSslCertificateApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/RegionSslCertificate",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "RegionSslCertificate",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeRegionSslCertificateApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	certificateProp, err := expandComputeRegionSslCertificateCertificate(d.Get("certificate"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate"); !isEmptyValue(reflect.ValueOf(certificateProp)) && (ok || !reflect.DeepEqual(v, certificateProp)) {
		obj["certificate"] = certificateProp
	}
	descriptionProp, err := expandComputeRegionSslCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionSslCertificateName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	privateKeyProp, err := expandComputeRegionSslCertificatePrivateKey(d.Get("private_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_key"); !isEmptyValue(reflect.ValueOf(privateKeyProp)) && (ok || !reflect.DeepEqual(v, privateKeyProp)) {
		obj["privateKey"] = privateKeyProp
	}
	regionProp, err := expandComputeRegionSslCertificateRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionSslCertificateCertificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	var certName string
	if v, ok := d.GetOk("name"); ok {
		certName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		certName = resource.PrefixedUniqueId(v.(string))
	} else {
		certName = resource.UniqueId()
	}

	// We need to get the {{name}} into schema to set the ID using ReplaceVars
	if err := d.Set("name", certName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return certName, nil
}

func expandComputeRegionSslCertificatePrivateKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslCertificateRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
