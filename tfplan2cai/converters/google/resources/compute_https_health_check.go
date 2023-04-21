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

package google

import (
	"reflect"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeHttpsHealthCheckAssetType string = "compute.googleapis.com/HttpsHealthCheck"

func resourceConverterComputeHttpsHealthCheck() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeHttpsHealthCheckAssetType,
		Convert:   GetComputeHttpsHealthCheckCaiObject,
	}
}

func GetComputeHttpsHealthCheckCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeHttpsHealthCheckApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeHttpsHealthCheckAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "HttpsHealthCheck",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeHttpsHealthCheckApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHttpsHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHttpsHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHttpsHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	hostProp, err := expandComputeHttpsHealthCheckHost(d.Get("host"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("host"); !isEmptyValue(reflect.ValueOf(hostProp)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	nameProp, err := expandComputeHttpsHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeHttpsHealthCheckPort(d.Get("port"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	requestPathProp, err := expandComputeHttpsHealthCheckRequestPath(d.Get("request_path"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("request_path"); !isEmptyValue(reflect.ValueOf(requestPathProp)) && (ok || !reflect.DeepEqual(v, requestPathProp)) {
		obj["requestPath"] = requestPathProp
	}
	timeoutSecProp, err := expandComputeHttpsHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHttpsHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}

	return obj, nil
}

func expandComputeHttpsHealthCheckCheckIntervalSec(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHealthyThreshold(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHost(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckPort(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckTimeoutSec(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckUnhealthyThreshold(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
