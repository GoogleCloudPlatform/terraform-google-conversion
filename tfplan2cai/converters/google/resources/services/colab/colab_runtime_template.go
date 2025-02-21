// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/colab/RuntimeTemplate.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package colab

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ColabRuntimeTemplateAssetType string = "aiplatform.googleapis.com/RuntimeTemplate"

func ResourceConverterColabRuntimeTemplate() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ColabRuntimeTemplateAssetType,
		Convert:   GetColabRuntimeTemplateCaiObject,
	}
}

func GetColabRuntimeTemplateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{location}}/notebookRuntimeTemplates/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetColabRuntimeTemplateApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ColabRuntimeTemplateAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "RuntimeTemplate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetColabRuntimeTemplateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandColabRuntimeTemplateName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandColabRuntimeTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandColabRuntimeTemplateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	machineSpecProp, err := expandColabRuntimeTemplateMachineSpec(d.Get("machine_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineSpecProp)) && (ok || !reflect.DeepEqual(v, machineSpecProp)) {
		obj["machineSpec"] = machineSpecProp
	}
	dataPersistentDiskSpecProp, err := expandColabRuntimeTemplateDataPersistentDiskSpec(d.Get("data_persistent_disk_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_persistent_disk_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataPersistentDiskSpecProp)) && (ok || !reflect.DeepEqual(v, dataPersistentDiskSpecProp)) {
		obj["dataPersistentDiskSpec"] = dataPersistentDiskSpecProp
	}
	networkSpecProp, err := expandColabRuntimeTemplateNetworkSpec(d.Get("network_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkSpecProp)) && (ok || !reflect.DeepEqual(v, networkSpecProp)) {
		obj["networkSpec"] = networkSpecProp
	}
	idleShutdownConfigProp, err := expandColabRuntimeTemplateIdleShutdownConfig(d.Get("idle_shutdown_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("idle_shutdown_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(idleShutdownConfigProp)) && (ok || !reflect.DeepEqual(v, idleShutdownConfigProp)) {
		obj["idleShutdownConfig"] = idleShutdownConfigProp
	}
	eucConfigProp, err := expandColabRuntimeTemplateEucConfig(d.Get("euc_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("euc_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(eucConfigProp)) && (ok || !reflect.DeepEqual(v, eucConfigProp)) {
		obj["eucConfig"] = eucConfigProp
	}
	shieldedVmConfigProp, err := expandColabRuntimeTemplateShieldedVmConfig(d.Get("shielded_vm_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shielded_vm_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(shieldedVmConfigProp)) && (ok || !reflect.DeepEqual(v, shieldedVmConfigProp)) {
		obj["shieldedVmConfig"] = shieldedVmConfigProp
	}
	networkTagsProp, err := expandColabRuntimeTemplateNetworkTags(d.Get("network_tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkTagsProp)) && (ok || !reflect.DeepEqual(v, networkTagsProp)) {
		obj["networkTags"] = networkTagsProp
	}
	encryptionSpecProp, err := expandColabRuntimeTemplateEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}
	softwareConfigProp, err := expandColabRuntimeTemplateSoftwareConfig(d.Get("software_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("software_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(softwareConfigProp)) && (ok || !reflect.DeepEqual(v, softwareConfigProp)) {
		obj["softwareConfig"] = softwareConfigProp
	}
	labelsProp, err := expandColabRuntimeTemplateEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandColabRuntimeTemplateName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateMachineSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMachineType, err := expandColabRuntimeTemplateMachineSpecMachineType(original["machine_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["machineType"] = transformedMachineType
	}

	transformedAcceleratorType, err := expandColabRuntimeTemplateMachineSpecAcceleratorType(original["accelerator_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["acceleratorType"] = transformedAcceleratorType
	}

	transformedAcceleratorCount, err := expandColabRuntimeTemplateMachineSpecAcceleratorCount(original["accelerator_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAcceleratorCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["acceleratorCount"] = transformedAcceleratorCount
	}

	return transformed, nil
}

func expandColabRuntimeTemplateMachineSpecMachineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateMachineSpecAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateMachineSpecAcceleratorCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateDataPersistentDiskSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDiskType, err := expandColabRuntimeTemplateDataPersistentDiskSpecDiskType(original["disk_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiskType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diskType"] = transformedDiskType
	}

	transformedDiskSizeGb, err := expandColabRuntimeTemplateDataPersistentDiskSpecDiskSizeGb(original["disk_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiskSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diskSizeGb"] = transformedDiskSizeGb
	}

	return transformed, nil
}

func expandColabRuntimeTemplateDataPersistentDiskSpecDiskType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateDataPersistentDiskSpecDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateNetworkSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableInternetAccess, err := expandColabRuntimeTemplateNetworkSpecEnableInternetAccess(original["enable_internet_access"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableInternetAccess); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableInternetAccess"] = transformedEnableInternetAccess
	}

	transformedNetwork, err := expandColabRuntimeTemplateNetworkSpecNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedSubnetwork, err := expandColabRuntimeTemplateNetworkSpecSubnetwork(original["subnetwork"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnetwork"] = transformedSubnetwork
	}

	return transformed, nil
}

func expandColabRuntimeTemplateNetworkSpecEnableInternetAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateNetworkSpecNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateNetworkSpecSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateIdleShutdownConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdleTimeout, err := expandColabRuntimeTemplateIdleShutdownConfigIdleTimeout(original["idle_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdleTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idleTimeout"] = transformedIdleTimeout
	}

	return transformed, nil
}

func expandColabRuntimeTemplateIdleShutdownConfigIdleTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateEucConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEucDisabled, err := expandColabRuntimeTemplateEucConfigEucDisabled(original["euc_disabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEucDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eucDisabled"] = transformedEucDisabled
	}

	return transformed, nil
}

func expandColabRuntimeTemplateEucConfigEucDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateShieldedVmConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSecureBoot, err := expandColabRuntimeTemplateShieldedVmConfigEnableSecureBoot(original["enable_secure_boot"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSecureBoot); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSecureBoot"] = transformedEnableSecureBoot
	}

	return transformed, nil
}

func expandColabRuntimeTemplateShieldedVmConfigEnableSecureBoot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateNetworkTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateEncryptionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandColabRuntimeTemplateEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandColabRuntimeTemplateEncryptionSpecKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateSoftwareConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnv, err := expandColabRuntimeTemplateSoftwareConfigEnv(original["env"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["env"] = transformedEnv
	}

	transformedPostStartupScriptConfig, err := expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig(original["post_startup_script_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostStartupScriptConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postStartupScriptConfig"] = transformedPostStartupScriptConfig
	}

	return transformed, nil
}

func expandColabRuntimeTemplateSoftwareConfigEnv(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandColabRuntimeTemplateSoftwareConfigEnvName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandColabRuntimeTemplateSoftwareConfigEnvValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandColabRuntimeTemplateSoftwareConfigEnvName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateSoftwareConfigEnvValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPostStartupScript, err := expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScript(original["post_startup_script"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostStartupScript); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postStartupScript"] = transformedPostStartupScript
	}

	transformedPostStartupScriptUrl, err := expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScriptUrl(original["post_startup_script_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostStartupScriptUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postStartupScriptUrl"] = transformedPostStartupScriptUrl
	}

	transformedPostStartupScriptBehavior, err := expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScriptBehavior(original["post_startup_script_behavior"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostStartupScriptBehavior); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postStartupScriptBehavior"] = transformedPostStartupScriptBehavior
	}

	return transformed, nil
}

func expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScript(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScriptUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateSoftwareConfigPostStartupScriptConfigPostStartupScriptBehavior(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeTemplateEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
