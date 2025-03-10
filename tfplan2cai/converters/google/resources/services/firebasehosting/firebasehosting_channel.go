// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebasehosting/Channel.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebasehosting

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseHostingChannelAssetType string = "firebasehosting.googleapis.com/Channel"

func ResourceConverterFirebaseHostingChannel() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseHostingChannelAssetType,
		Convert:   GetFirebaseHostingChannelCaiObject,
	}
}

func GetFirebaseHostingChannelCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebasehosting.googleapis.com/sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseHostingChannelApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseHostingChannelAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebasehosting/v1beta1/rest",
				DiscoveryName:        "Channel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseHostingChannelApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	retainedReleaseCountProp, err := expandFirebaseHostingChannelRetainedReleaseCount(d.Get("retained_release_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retained_release_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(retainedReleaseCountProp)) && (ok || !reflect.DeepEqual(v, retainedReleaseCountProp)) {
		obj["retainedReleaseCount"] = retainedReleaseCountProp
	}
	expireTimeProp, err := expandFirebaseHostingChannelExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandFirebaseHostingChannelTtl(d.Get("ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	labelsProp, err := expandFirebaseHostingChannelEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandFirebaseHostingChannelRetainedReleaseCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingChannelEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
