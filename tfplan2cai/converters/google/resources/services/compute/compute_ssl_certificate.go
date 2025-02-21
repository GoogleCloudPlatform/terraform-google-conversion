// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/SslCertificate.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeSslCertificateAssetType string = "compute.googleapis.com/SslCertificate"

func ResourceConverterComputeSslCertificate() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeSslCertificateAssetType,
		Convert:   GetComputeSslCertificateCaiObject,
	}
}

func GetComputeSslCertificateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeSslCertificateApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeSslCertificateAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "SslCertificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeSslCertificateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	certificateProp, err := expandComputeSslCertificateCertificate(d.Get("certificate"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateProp)) && (ok || !reflect.DeepEqual(v, certificateProp)) {
		obj["certificate"] = certificateProp
	}
	descriptionProp, err := expandComputeSslCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeSslCertificateName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	privateKeyProp, err := expandComputeSslCertificatePrivateKey(d.Get("private_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateKeyProp)) && (ok || !reflect.DeepEqual(v, privateKeyProp)) {
		obj["privateKey"] = privateKeyProp
	}

	return obj, nil
}

func expandComputeSslCertificateCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslCertificateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslCertificateName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	var certName string
	if v, ok := d.GetOk("name"); ok {
		certName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		prefix := v.(string)
		if len(prefix) > 37 {
			certName = tpgresource.ReducedPrefixedUniqueId(prefix)
		} else {
			certName = id.PrefixedUniqueId(prefix)
		}
	} else {
		certName = id.UniqueId()
	}

	// We need to get the {{name}} into schema to set the ID using tpgresource.ReplaceVars
	if err := d.Set("name", certName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return certName, nil
}

func expandComputeSslCertificatePrivateKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
