// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/TargetHttpProxy.yaml
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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeTargetHttpProxyAssetType string = "compute.googleapis.com/TargetHttpProxy"

func ResourceConverterComputeTargetHttpProxy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeTargetHttpProxyAssetType,
		Convert:   GetComputeTargetHttpProxyCaiObject,
	}
}

func GetComputeTargetHttpProxyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/targetHttpProxies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeTargetHttpProxyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeTargetHttpProxyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "TargetHttpProxy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeTargetHttpProxyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetHttpProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeTargetHttpProxyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	urlMapProp, err := expandComputeTargetHttpProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("url_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	proxyBindProp, err := expandComputeTargetHttpProxyProxyBind(d.Get("proxy_bind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("proxy_bind"); !tpgresource.IsEmptyValue(reflect.ValueOf(proxyBindProp)) && (ok || !reflect.DeepEqual(v, proxyBindProp)) {
		obj["proxyBind"] = proxyBindProp
	}
	httpKeepAliveTimeoutSecProp, err := expandComputeTargetHttpProxyHttpKeepAliveTimeoutSec(d.Get("http_keep_alive_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_keep_alive_timeout_sec"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpKeepAliveTimeoutSecProp)) && (ok || !reflect.DeepEqual(v, httpKeepAliveTimeoutSecProp)) {
		obj["httpKeepAliveTimeoutSec"] = httpKeepAliveTimeoutSecProp
	}

	return obj, nil
}

func expandComputeTargetHttpProxyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpProxyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpProxyUrlMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("urlMaps", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetHttpProxyProxyBind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpProxyHttpKeepAliveTimeoutSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
