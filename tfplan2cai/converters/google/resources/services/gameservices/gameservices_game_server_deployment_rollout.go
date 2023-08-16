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

package gameservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GameServicesGameServerDeploymentRolloutAssetType string = "gameservices.googleapis.com/GameServerDeploymentRollout"

func ResourceConverterGameServicesGameServerDeploymentRollout() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GameServicesGameServerDeploymentRolloutAssetType,
		Convert:   GetGameServicesGameServerDeploymentRolloutCaiObject,
	}
}

func GetGameServicesGameServerDeploymentRolloutCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gameservices.googleapis.com/projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGameServicesGameServerDeploymentRolloutApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GameServicesGameServerDeploymentRolloutAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gameservices/v1beta/rest",
				DiscoveryName:        "GameServerDeploymentRollout",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGameServicesGameServerDeploymentRolloutApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	defaultGameServerConfigProp, err := expandGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(d.Get("default_game_server_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_game_server_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultGameServerConfigProp)) && (ok || !reflect.DeepEqual(v, defaultGameServerConfigProp)) {
		obj["defaultGameServerConfig"] = defaultGameServerConfigProp
	}
	gameServerConfigOverridesProp, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(d.Get("game_server_config_overrides"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("game_server_config_overrides"); !tpgresource.IsEmptyValue(reflect.ValueOf(gameServerConfigOverridesProp)) && (ok || !reflect.DeepEqual(v, gameServerConfigOverridesProp)) {
		obj["gameServerConfigOverrides"] = gameServerConfigOverridesProp
	}

	return obj, nil
}

func expandGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRealmsSelector, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(original["realms_selector"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRealmsSelector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["realmsSelector"] = transformedRealmsSelector
		}

		transformedConfigVersion, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(original["config_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConfigVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["configVersion"] = transformedConfigVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRealms, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(original["realms"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRealms); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["realms"] = transformedRealms
	}

	return transformed, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
