// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Autoscaler.yaml
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeAutoscalerAssetType string = "compute.googleapis.com/Autoscaler"

func ResourceConverterComputeAutoscaler() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeAutoscalerAssetType,
		Convert:   GetComputeAutoscalerCaiObject,
	}
}

func GetComputeAutoscalerCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeAutoscalerApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeAutoscalerAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "Autoscaler",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeAutoscalerApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	zoneProp, err := expandComputeAutoscalerZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	return obj, nil
}

func expandComputeAutoscalerName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandComputeAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["minNumReplicas"] = transformedMinReplicas
	}

	transformedMaxReplicas, err := expandComputeAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["maxNumReplicas"] = transformedMaxReplicas
	}

	transformedCooldownPeriod, err := expandComputeAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["coolDownPeriodSec"] = transformedCooldownPeriod
	}

	transformedMode, err := expandComputeAutoscalerAutoscalingPolicyMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedScaleDownControl, err := expandComputeAutoscalerAutoscalingPolicyScaleDownControl(original["scale_down_control"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownControl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownControl"] = transformedScaleDownControl
	}

	transformedScaleInControl, err := expandComputeAutoscalerAutoscalingPolicyScaleInControl(original["scale_in_control"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleInControl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleInControl"] = transformedScaleInControl
	}

	transformedCpuUtilization, err := expandComputeAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuUtilization"] = transformedCpuUtilization
	}

	transformedMetric, err := expandComputeAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetric); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customMetricUtilizations"] = transformedMetric
	}

	transformedLoadBalancingUtilization, err := expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoadBalancingUtilization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization
	}

	transformedScalingSchedules, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedules(original["scaling_schedules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScalingSchedules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scalingSchedules"] = transformedScalingSchedules
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleDownControl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxScaledDownReplicas, err := expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas(original["max_scaled_down_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxScaledDownReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxScaledDownReplicas"] = transformedMaxScaledDownReplicas
	}

	transformedTimeWindowSec, err := expandComputeAutoscalerAutoscalingPolicyScaleDownControlTimeWindowSec(original["time_window_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeWindowSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeWindowSec"] = transformedTimeWindowSec
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixed, err := expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasFixed(original["fixed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixed"] = transformedFixed
	}

	transformedPercent, err := expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasFixed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleDownControlTimeWindowSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleInControl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxScaledInReplicas, err := expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(original["max_scaled_in_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxScaledInReplicas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxScaledInReplicas"] = transformedMaxScaledInReplicas
	}

	transformedTimeWindowSec, err := expandComputeAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(original["time_window_sec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeWindowSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeWindowSec"] = transformedTimeWindowSec
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixed, err := expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(original["fixed"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixed"] = transformedFixed
	}

	transformedPercent, err := expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasFixed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScaleInControlTimeWindowSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	transformedPredictiveMethod, err := expandComputeAutoscalerAutoscalingPolicyCpuUtilizationPredictiveMethod(original["predictive_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredictiveMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predictiveMethod"] = transformedPredictiveMethod
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyCpuUtilizationPredictiveMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetric(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["metric"] = transformedName
		}

		transformedSingleInstanceAssignment, err := expandComputeAutoscalerAutoscalingPolicyMetricSingleInstanceAssignment(original["single_instance_assignment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSingleInstanceAssignment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["singleInstanceAssignment"] = transformedSingleInstanceAssignment
		}

		transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["utilizationTarget"] = transformedTarget
		}

		transformedType, err := expandComputeAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["utilizationTargetType"] = transformedType
		}

		transformedFilter, err := expandComputeAutoscalerAutoscalingPolicyMetricFilter(original["filter"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["filter"] = transformedFilter
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricSingleInstanceAssignment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinRequiredReplicas, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesMinRequiredReplicas(original["min_required_replicas"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["minRequiredReplicas"] = transformedMinRequiredReplicas
		}

		transformedSchedule, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesSchedule(original["schedule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["schedule"] = transformedSchedule
		}

		transformedTimeZone, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesTimeZone(original["time_zone"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["timeZone"] = transformedTimeZone
		}

		transformedDurationSec, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDurationSec(original["duration_sec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDurationSec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["durationSec"] = transformedDurationSec
		}

		transformedDisabled, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDisabled(original["disabled"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["disabled"] = transformedDisabled
		}

		transformedDescription, err := expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedName, err := tpgresource.ExpandString(original["name"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedName] = transformed
	}
	return m, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesMinRequiredReplicas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDurationSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyScalingSchedulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	}
	f, err := tpgresource.ParseZonalFieldValue("instanceGroupManagers", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target: %s", err)
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+f.RelativeLink())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func expandComputeAutoscalerZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
