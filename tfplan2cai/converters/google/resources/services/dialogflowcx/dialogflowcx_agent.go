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

package dialogflowcx

import (
	"encoding/json"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXAgentAssetType string = "{{location}}-dialogflow.googleapis.com/Agent"

func ResourceConverterDialogflowCXAgent() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXAgentAssetType,
		Convert:   GetDialogflowCXAgentCaiObject,
	}
}

func GetDialogflowCXAgentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXAgentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXAgentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "Agent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXAgentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultLanguageCodeProp, err := expandDialogflowCXAgentDefaultLanguageCode(d.Get("default_language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_language_code"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultLanguageCodeProp)) && (ok || !reflect.DeepEqual(v, defaultLanguageCodeProp)) {
		obj["defaultLanguageCode"] = defaultLanguageCodeProp
	}
	supportedLanguageCodesProp, err := expandDialogflowCXAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportedLanguageCodesProp)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowCXAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowCXAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowCXAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("avatar_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(avatarUriProp)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	speechToTextSettingsProp, err := expandDialogflowCXAgentSpeechToTextSettings(d.Get("speech_to_text_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("speech_to_text_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(speechToTextSettingsProp)) && (ok || !reflect.DeepEqual(v, speechToTextSettingsProp)) {
		obj["speechToTextSettings"] = speechToTextSettingsProp
	}
	securitySettingsProp, err := expandDialogflowCXAgentSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXAgentEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXAgentEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}
	advancedSettingsProp, err := expandDialogflowCXAgentAdvancedSettings(d.Get("advanced_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("advanced_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(advancedSettingsProp)) && (ok || !reflect.DeepEqual(v, advancedSettingsProp)) {
		obj["advancedSettings"] = advancedSettingsProp
	}
	gitIntegrationSettingsProp, err := expandDialogflowCXAgentGitIntegrationSettings(d.Get("git_integration_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("git_integration_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(gitIntegrationSettingsProp)) && (ok || !reflect.DeepEqual(v, gitIntegrationSettingsProp)) {
		obj["gitIntegrationSettings"] = gitIntegrationSettingsProp
	}
	textToSpeechSettingsProp, err := expandDialogflowCXAgentTextToSpeechSettings(d.Get("text_to_speech_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("text_to_speech_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(textToSpeechSettingsProp)) && (ok || !reflect.DeepEqual(v, textToSpeechSettingsProp)) {
		obj["textToSpeechSettings"] = textToSpeechSettingsProp
	}

	return obj, nil
}

func expandDialogflowCXAgentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSupportedLanguageCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAvatarUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSpeechToTextSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSpeechAdaptation, err := expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(original["enable_speech_adaptation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSpeechAdaptation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSpeechAdaptation"] = transformedEnableSpeechAdaptation
	}

	return transformed, nil
}

func expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSecuritySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableSpellCorrection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAudioExportGcsDestination, err := expandDialogflowCXAgentAdvancedSettingsAudioExportGcsDestination(original["audio_export_gcs_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudioExportGcsDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audioExportGcsDestination"] = transformedAudioExportGcsDestination
	}

	transformedSpeechSettings, err := expandDialogflowCXAgentAdvancedSettingsSpeechSettings(original["speech_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpeechSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["speechSettings"] = transformedSpeechSettings
	}

	transformedDtmfSettings, err := expandDialogflowCXAgentAdvancedSettingsDtmfSettings(original["dtmf_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDtmfSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dtmfSettings"] = transformedDtmfSettings
	}

	transformedLoggingSettings, err := expandDialogflowCXAgentAdvancedSettingsLoggingSettings(original["logging_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoggingSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loggingSettings"] = transformedLoggingSettings
	}

	return transformed, nil
}

func expandDialogflowCXAgentAdvancedSettingsAudioExportGcsDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowCXAgentAdvancedSettingsAudioExportGcsDestinationUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	return transformed, nil
}

func expandDialogflowCXAgentAdvancedSettingsAudioExportGcsDestinationUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsSpeechSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEndpointerSensitivity, err := expandDialogflowCXAgentAdvancedSettingsSpeechSettingsEndpointerSensitivity(original["endpointer_sensitivity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndpointerSensitivity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endpointerSensitivity"] = transformedEndpointerSensitivity
	}

	transformedNoSpeechTimeout, err := expandDialogflowCXAgentAdvancedSettingsSpeechSettingsNoSpeechTimeout(original["no_speech_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoSpeechTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["noSpeechTimeout"] = transformedNoSpeechTimeout
	}

	transformedUseTimeoutBasedEndpointing, err := expandDialogflowCXAgentAdvancedSettingsSpeechSettingsUseTimeoutBasedEndpointing(original["use_timeout_based_endpointing"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUseTimeoutBasedEndpointing); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["useTimeoutBasedEndpointing"] = transformedUseTimeoutBasedEndpointing
	}

	transformedModels, err := expandDialogflowCXAgentAdvancedSettingsSpeechSettingsModels(original["models"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedModels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["models"] = transformedModels
	}

	return transformed, nil
}

func expandDialogflowCXAgentAdvancedSettingsSpeechSettingsEndpointerSensitivity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsSpeechSettingsNoSpeechTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsSpeechSettingsUseTimeoutBasedEndpointing(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsSpeechSettingsModels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDialogflowCXAgentAdvancedSettingsDtmfSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandDialogflowCXAgentAdvancedSettingsDtmfSettingsEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedMaxDigits, err := expandDialogflowCXAgentAdvancedSettingsDtmfSettingsMaxDigits(original["max_digits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDigits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDigits"] = transformedMaxDigits
	}

	transformedFinishDigit, err := expandDialogflowCXAgentAdvancedSettingsDtmfSettingsFinishDigit(original["finish_digit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFinishDigit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["finishDigit"] = transformedFinishDigit
	}

	return transformed, nil
}

func expandDialogflowCXAgentAdvancedSettingsDtmfSettingsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsDtmfSettingsMaxDigits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsDtmfSettingsFinishDigit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsLoggingSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableStackdriverLogging, err := expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableStackdriverLogging(original["enable_stackdriver_logging"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableStackdriverLogging); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableStackdriverLogging"] = transformedEnableStackdriverLogging
	}

	transformedEnableInteractionLogging, err := expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableInteractionLogging(original["enable_interaction_logging"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableInteractionLogging); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableInteractionLogging"] = transformedEnableInteractionLogging
	}

	transformedEnableConsentBasedRedaction, err := expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableConsentBasedRedaction(original["enable_consent_based_redaction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableConsentBasedRedaction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableConsentBasedRedaction"] = transformedEnableConsentBasedRedaction
	}

	return transformed, nil
}

func expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableInteractionLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAdvancedSettingsLoggingSettingsEnableConsentBasedRedaction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentGitIntegrationSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGithubSettings, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettings(original["github_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGithubSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["githubSettings"] = transformedGithubSettings
	}

	return transformed, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisplayName, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsDisplayName(original["display_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	transformedRepositoryUri, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsRepositoryUri(original["repository_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepositoryUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repositoryUri"] = transformedRepositoryUri
	}

	transformedTrackingBranch, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsTrackingBranch(original["tracking_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrackingBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trackingBranch"] = transformedTrackingBranch
	}

	transformedAccessToken, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsAccessToken(original["access_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessToken); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessToken"] = transformedAccessToken
	}

	transformedBranches, err := expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsBranches(original["branches"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["branches"] = transformedBranches
	}

	return transformed, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsRepositoryUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsTrackingBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsAccessToken(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentGitIntegrationSettingsGithubSettingsBranches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentTextToSpeechSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSynthesizeSpeechConfigs, err := expandDialogflowCXAgentTextToSpeechSettingsSynthesizeSpeechConfigs(original["synthesize_speech_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSynthesizeSpeechConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["synthesizeSpeechConfigs"] = transformedSynthesizeSpeechConfigs
	}

	return transformed, nil
}

func expandDialogflowCXAgentTextToSpeechSettingsSynthesizeSpeechConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}
