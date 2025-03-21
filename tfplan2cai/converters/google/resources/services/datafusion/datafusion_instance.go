// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datafusion/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datafusion

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var instanceAcceleratorOptions = []string{
	"delta.default.checkpoint.directory",
	"ui.feature.cdc",
}

func instanceOptionsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Suppress diffs for the options generated by adding an accelerator to a data fusion instance
	for _, option := range instanceAcceleratorOptions {
		if strings.Contains(k, option) && new == "" {
			return true
		}
	}

	// Let diff be determined by options (above)
	if strings.Contains(k, "options.%") {
		return true
	}

	// For other keys, don't suppress diff.
	return false
}

const DataFusionInstanceAssetType string = "datafusion.googleapis.com/Instance"

func ResourceConverterDataFusionInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataFusionInstanceAssetType,
		Convert:   GetDataFusionInstanceCaiObject,
	}
}

func GetDataFusionInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//datafusion.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataFusionInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataFusionInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datafusion/v1beta1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataFusionInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDataFusionInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandDataFusionInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandDataFusionInstanceType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	enableStackdriverLoggingProp, err := expandDataFusionInstanceEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableStackdriverMonitoringProp, err := expandDataFusionInstanceEnableStackdriverMonitoring(d.Get("enable_stackdriver_monitoring"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_monitoring"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverMonitoringProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverMonitoringProp)) {
		obj["enableStackdriverMonitoring"] = enableStackdriverMonitoringProp
	}
	enableRbacProp, err := expandDataFusionInstanceEnableRbac(d.Get("enable_rbac"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_rbac"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableRbacProp)) && (ok || !reflect.DeepEqual(v, enableRbacProp)) {
		obj["enableRbac"] = enableRbacProp
	}
	optionsProp, err := expandDataFusionInstanceOptions(d.Get("options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("options"); !tpgresource.IsEmptyValue(reflect.ValueOf(optionsProp)) && (ok || !reflect.DeepEqual(v, optionsProp)) {
		obj["options"] = optionsProp
	}
	versionProp, err := expandDataFusionInstanceVersion(d.Get("version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	privateInstanceProp, err := expandDataFusionInstancePrivateInstance(d.Get("private_instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateInstanceProp)) && (ok || !reflect.DeepEqual(v, privateInstanceProp)) {
		obj["privateInstance"] = privateInstanceProp
	}
	dataprocServiceAccountProp, err := expandDataFusionInstanceDataprocServiceAccount(d.Get("dataproc_service_account"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataproc_service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataprocServiceAccountProp)) && (ok || !reflect.DeepEqual(v, dataprocServiceAccountProp)) {
		obj["dataprocServiceAccount"] = dataprocServiceAccountProp
	}
	networkConfigProp, err := expandDataFusionInstanceNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	zoneProp, err := expandDataFusionInstanceZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	displayNameProp, err := expandDataFusionInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	cryptoKeyConfigProp, err := expandDataFusionInstanceCryptoKeyConfig(d.Get("crypto_key_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("crypto_key_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cryptoKeyConfigProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyConfigProp)) {
		obj["cryptoKeyConfig"] = cryptoKeyConfigProp
	}
	eventPublishConfigProp, err := expandDataFusionInstanceEventPublishConfig(d.Get("event_publish_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("event_publish_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(eventPublishConfigProp)) && (ok || !reflect.DeepEqual(v, eventPublishConfigProp)) {
		obj["eventPublishConfig"] = eventPublishConfigProp
	}
	acceleratorsProp, err := expandDataFusionInstanceAccelerators(d.Get("accelerators"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("accelerators"); !tpgresource.IsEmptyValue(reflect.ValueOf(acceleratorsProp)) && (ok || !reflect.DeepEqual(v, acceleratorsProp)) {
		obj["accelerators"] = acceleratorsProp
	}
	tagsProp, err := expandDataFusionInstanceTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	labelsProp, err := expandDataFusionInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDataFusionInstanceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandDataFusionInstanceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableStackdriverMonitoring(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEnableRbac(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstancePrivateInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceDataprocServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAllocation, err := expandDataFusionInstanceNetworkConfigIpAllocation(original["ip_allocation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAllocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAllocation"] = transformedIpAllocation
	}

	transformedNetwork, err := expandDataFusionInstanceNetworkConfigNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedConnectionType, err := expandDataFusionInstanceNetworkConfigConnectionType(original["connection_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectionType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["connectionType"] = transformedConnectionType
	}

	transformedPrivateServiceConnectConfig, err := expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfig(original["private_service_connect_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateServiceConnectConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["privateServiceConnectConfig"] = transformedPrivateServiceConnectConfig
	}

	return transformed, nil
}

func expandDataFusionInstanceNetworkConfigIpAllocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigConnectionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkAttachment, err := expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigNetworkAttachment(original["network_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkAttachment"] = transformedNetworkAttachment
	}

	transformedUnreachableCidrBlock, err := expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigUnreachableCidrBlock(original["unreachable_cidr_block"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnreachableCidrBlock); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["unreachableCidrBlock"] = transformedUnreachableCidrBlock
	}

	transformedEffectiveUnreachableCidrBlock, err := expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigEffectiveUnreachableCidrBlock(original["effective_unreachable_cidr_block"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEffectiveUnreachableCidrBlock); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["effectiveUnreachableCidrBlock"] = transformedEffectiveUnreachableCidrBlock
	}

	return transformed, nil
}

func expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigNetworkAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigUnreachableCidrBlock(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceNetworkConfigPrivateServiceConnectConfigEffectiveUnreachableCidrBlock(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceCryptoKeyConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKeyReference, err := expandDataFusionInstanceCryptoKeyConfigKeyReference(original["key_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyReference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keyReference"] = transformedKeyReference
	}

	return transformed, nil
}

func expandDataFusionInstanceCryptoKeyConfigKeyReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEventPublishConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandDataFusionInstanceEventPublishConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTopic, err := expandDataFusionInstanceEventPublishConfigTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandDataFusionInstanceEventPublishConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceEventPublishConfigTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceAccelerators(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAcceleratorType, err := expandDataFusionInstanceAcceleratorsAcceleratorType(original["accelerator_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["acceleratorType"] = transformedAcceleratorType
		}

		transformedState, err := expandDataFusionInstanceAcceleratorsState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataFusionInstanceAcceleratorsAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceAcceleratorsState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataFusionInstanceTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataFusionInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
