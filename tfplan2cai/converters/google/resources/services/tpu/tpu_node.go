// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/tpu/Node.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package tpu

import (
	"context"
	"fmt"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// compareTpuNodeSchedulingConfig diff suppresses for the default
// scheduling, i.e. if preemptible is false, the API may either return no
// schedulingConfig or an empty schedulingConfig.
func compareTpuNodeSchedulingConfig(k, old, new string, d *schema.ResourceData) bool {
	if k == "scheduling_config.0.preemptible" {
		return old == "" && new == "false"
	}
	if k == "scheduling_config.#" {
		o, n := d.GetChange("scheduling_config.0.preemptible")
		return o.(bool) == n.(bool)
	}
	return false
}

func tpuNodeCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	old, new := diff.GetChange("network")
	config := meta.(*transport_tpg.Config)

	networkLinkRegex := regexp.MustCompile("projects/(.+)/global/networks/(.+)")

	var pid string

	if networkLinkRegex.MatchString(new.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(new.(string))
		pid = parts[1]
	}

	if pid == "" {
		return nil
	}

	project, err := config.NewResourceManagerClient(config.UserAgent).Projects.Get(pid).Do()
	if err != nil {
		return fmt.Errorf("Failed to retrieve project, pid: %s, err: %s", pid, err)
	}

	if networkLinkRegex.MatchString(old.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(old.(string))
		i, err := tpgresource.StringToFixed64(parts[1])
		if err == nil {
			if project.ProjectNumber == i {
				if err := diff.SetNew("network", old); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}

const TPUNodeAssetType string = "tpu.googleapis.com/Node"

func ResourceConverterTPUNode() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: TPUNodeAssetType,
		Convert:   GetTPUNodeCaiObject,
	}
}

func GetTPUNodeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//tpu.googleapis.com/projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetTPUNodeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: TPUNodeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/tpu/v1/rest",
				DiscoveryName:        "Node",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetTPUNodeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandTPUNodeName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandTPUNodeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	acceleratorTypeProp, err := expandTPUNodeAcceleratorType(d.Get("accelerator_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerator_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(acceleratorTypeProp)) && (ok || !reflect.DeepEqual(v, acceleratorTypeProp)) {
		obj["acceleratorType"] = acceleratorTypeProp
	}
	tensorflowVersionProp, err := expandTPUNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tensorflow_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(tensorflowVersionProp)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
		obj["tensorflowVersion"] = tensorflowVersionProp
	}
	networkProp, err := expandTPUNodeNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	cidrBlockProp, err := expandTPUNodeCidrBlock(d.Get("cidr_block"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cidr_block"); !tpgresource.IsEmptyValue(reflect.ValueOf(cidrBlockProp)) && (ok || !reflect.DeepEqual(v, cidrBlockProp)) {
		obj["cidrBlock"] = cidrBlockProp
	}
	useServiceNetworkingProp, err := expandTPUNodeUseServiceNetworking(d.Get("use_service_networking"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("use_service_networking"); !tpgresource.IsEmptyValue(reflect.ValueOf(useServiceNetworkingProp)) && (ok || !reflect.DeepEqual(v, useServiceNetworkingProp)) {
		obj["useServiceNetworking"] = useServiceNetworkingProp
	}
	schedulingConfigProp, err := expandTPUNodeSchedulingConfig(d.Get("scheduling_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scheduling_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(schedulingConfigProp)) && (ok || !reflect.DeepEqual(v, schedulingConfigProp)) {
		obj["schedulingConfig"] = schedulingConfigProp
	}
	labelsProp, err := expandTPUNodeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandTPUNodeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeTensorflowVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeCidrBlock(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeUseServiceNetworking(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeSchedulingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreemptible, err := expandTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreemptible); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preemptible"] = transformedPreemptible
	}

	return transformed, nil
}

func expandTPUNodeSchedulingConfigPreemptible(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
