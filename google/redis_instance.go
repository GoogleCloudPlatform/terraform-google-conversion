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
)

func GetRedisInstanceCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//redis.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetRedisInstanceApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "redis.googleapis.com/Instance",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/redis/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetRedisInstanceApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("alternative_location_id"); !isEmptyValue(reflect.ValueOf(alternativeLocationIdProp)) && (ok || !reflect.DeepEqual(v, alternativeLocationIdProp)) {
		obj["alternativeLocationId"] = alternativeLocationIdProp
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	connectModeProp, err := expandRedisInstanceConnectMode(d.Get("connect_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connect_mode"); !isEmptyValue(reflect.ValueOf(connectModeProp)) && (ok || !reflect.DeepEqual(v, connectModeProp)) {
		obj["connectMode"] = connectModeProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(memorySizeGbProp)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_version"); !isEmptyValue(reflect.ValueOf(redisVersionProp)) && (ok || !reflect.DeepEqual(v, redisVersionProp)) {
		obj["redisVersion"] = redisVersionProp
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}

	return resourceRedisInstanceEncoder(d, config, obj)
}

func resourceRedisInstanceEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	region, err := getRegionFromSchema("region", "location_id", d, config)
	if err != nil {
		return nil, err
	}
	if err := d.Set("region", region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	return obj, nil
}

func expandRedisInstanceAlternativeLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	fv, err := ParseNetworkFieldValue(v.(string), d, config)
	if err != nil {
		return nil, err
	}
	return fv.RelativeLink(), nil
}

func expandRedisInstanceConnectMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceRedisConfigs(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandRedisInstanceMemorySizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRedisVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
