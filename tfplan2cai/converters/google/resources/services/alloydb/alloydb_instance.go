// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/alloydb/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package alloydb

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AlloydbInstanceAssetType string = "alloydb.googleapis.com/Instance"

func ResourceConverterAlloydbInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AlloydbInstanceAssetType,
		Convert:   GetAlloydbInstanceCaiObject,
	}
}

func GetAlloydbInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//alloydb.googleapis.com/{{cluster}}/instances/{{instance_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAlloydbInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AlloydbInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/alloydb/v1beta/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAlloydbInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandAlloydbInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gceZoneProp, err := expandAlloydbInstanceGceZone(d.Get("gce_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gce_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(gceZoneProp)) && (ok || !reflect.DeepEqual(v, gceZoneProp)) {
		obj["gceZone"] = gceZoneProp
	}
	databaseFlagsProp, err := expandAlloydbInstanceDatabaseFlags(d.Get("database_flags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database_flags"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseFlagsProp)) && (ok || !reflect.DeepEqual(v, databaseFlagsProp)) {
		obj["databaseFlags"] = databaseFlagsProp
	}
	availabilityTypeProp, err := expandAlloydbInstanceAvailabilityType(d.Get("availability_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("availability_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(availabilityTypeProp)) && (ok || !reflect.DeepEqual(v, availabilityTypeProp)) {
		obj["availabilityType"] = availabilityTypeProp
	}
	instanceTypeProp, err := expandAlloydbInstanceInstanceType(d.Get("instance_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceTypeProp)) && (ok || !reflect.DeepEqual(v, instanceTypeProp)) {
		obj["instanceType"] = instanceTypeProp
	}
	queryInsightsConfigProp, err := expandAlloydbInstanceQueryInsightsConfig(d.Get("query_insights_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("query_insights_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(queryInsightsConfigProp)) && (ok || !reflect.DeepEqual(v, queryInsightsConfigProp)) {
		obj["queryInsightsConfig"] = queryInsightsConfigProp
	}
	observabilityConfigProp, err := expandAlloydbInstanceObservabilityConfig(d.Get("observability_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("observability_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(observabilityConfigProp)) && (ok || !reflect.DeepEqual(v, observabilityConfigProp)) {
		obj["observabilityConfig"] = observabilityConfigProp
	}
	readPoolConfigProp, err := expandAlloydbInstanceReadPoolConfig(d.Get("read_pool_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("read_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(readPoolConfigProp)) && (ok || !reflect.DeepEqual(v, readPoolConfigProp)) {
		obj["readPoolConfig"] = readPoolConfigProp
	}
	machineConfigProp, err := expandAlloydbInstanceMachineConfig(d.Get("machine_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("machine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineConfigProp)) && (ok || !reflect.DeepEqual(v, machineConfigProp)) {
		obj["machineConfig"] = machineConfigProp
	}
	clientConnectionConfigProp, err := expandAlloydbInstanceClientConnectionConfig(d.Get("client_connection_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_connection_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientConnectionConfigProp)) && (ok || !reflect.DeepEqual(v, clientConnectionConfigProp)) {
		obj["clientConnectionConfig"] = clientConnectionConfigProp
	}
	pscInstanceConfigProp, err := expandAlloydbInstancePscInstanceConfig(d.Get("psc_instance_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("psc_instance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscInstanceConfigProp)) && (ok || !reflect.DeepEqual(v, pscInstanceConfigProp)) {
		obj["pscInstanceConfig"] = pscInstanceConfigProp
	}
	networkConfigProp, err := expandAlloydbInstanceNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	labelsProp, err := expandAlloydbInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbInstanceEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandAlloydbInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceGceZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceDatabaseFlags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceAvailabilityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceInstanceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceQueryInsightsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQueryStringLength, err := expandAlloydbInstanceQueryInsightsConfigQueryStringLength(original["query_string_length"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryStringLength); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryStringLength"] = transformedQueryStringLength
	}

	transformedRecordApplicationTags, err := expandAlloydbInstanceQueryInsightsConfigRecordApplicationTags(original["record_application_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecordApplicationTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recordApplicationTags"] = transformedRecordApplicationTags
	}

	transformedRecordClientAddress, err := expandAlloydbInstanceQueryInsightsConfigRecordClientAddress(original["record_client_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecordClientAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recordClientAddress"] = transformedRecordClientAddress
	}

	transformedQueryPlansPerMinute, err := expandAlloydbInstanceQueryInsightsConfigQueryPlansPerMinute(original["query_plans_per_minute"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryPlansPerMinute); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryPlansPerMinute"] = transformedQueryPlansPerMinute
	}

	return transformed, nil
}

func expandAlloydbInstanceQueryInsightsConfigQueryStringLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceQueryInsightsConfigRecordApplicationTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceQueryInsightsConfigRecordClientAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceQueryInsightsConfigQueryPlansPerMinute(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandAlloydbInstanceObservabilityConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedPreserveComments, err := expandAlloydbInstanceObservabilityConfigPreserveComments(original["preserve_comments"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreserveComments); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preserveComments"] = transformedPreserveComments
	}

	transformedTrackWaitEvents, err := expandAlloydbInstanceObservabilityConfigTrackWaitEvents(original["track_wait_events"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrackWaitEvents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trackWaitEvents"] = transformedTrackWaitEvents
	}

	transformedTrackWaitEventTypes, err := expandAlloydbInstanceObservabilityConfigTrackWaitEventTypes(original["track_wait_event_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrackWaitEventTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trackWaitEventTypes"] = transformedTrackWaitEventTypes
	}

	transformedMaxQueryStringLength, err := expandAlloydbInstanceObservabilityConfigMaxQueryStringLength(original["max_query_string_length"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxQueryStringLength); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxQueryStringLength"] = transformedMaxQueryStringLength
	}

	transformedRecordApplicationTags, err := expandAlloydbInstanceObservabilityConfigRecordApplicationTags(original["record_application_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecordApplicationTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recordApplicationTags"] = transformedRecordApplicationTags
	}

	transformedQueryPlansPerMinute, err := expandAlloydbInstanceObservabilityConfigQueryPlansPerMinute(original["query_plans_per_minute"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryPlansPerMinute); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryPlansPerMinute"] = transformedQueryPlansPerMinute
	}

	transformedTrackActiveQueries, err := expandAlloydbInstanceObservabilityConfigTrackActiveQueries(original["track_active_queries"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrackActiveQueries); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trackActiveQueries"] = transformedTrackActiveQueries
	}

	return transformed, nil
}

func expandAlloydbInstanceObservabilityConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigPreserveComments(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigTrackWaitEvents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigTrackWaitEventTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigMaxQueryStringLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigRecordApplicationTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigQueryPlansPerMinute(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceObservabilityConfigTrackActiveQueries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceReadPoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeCount, err := expandAlloydbInstanceReadPoolConfigNodeCount(original["node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeCount"] = transformedNodeCount
	}

	return transformed, nil
}

func expandAlloydbInstanceReadPoolConfigNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceMachineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpuCount, err := expandAlloydbInstanceMachineConfigCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	return transformed, nil
}

func expandAlloydbInstanceMachineConfigCpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceClientConnectionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequireConnectors, err := expandAlloydbInstanceClientConnectionConfigRequireConnectors(original["require_connectors"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireConnectors); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["requireConnectors"] = transformedRequireConnectors
	}

	transformedSslConfig, err := expandAlloydbInstanceClientConnectionConfigSslConfig(original["ssl_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sslConfig"] = transformedSslConfig
	}

	return transformed, nil
}

func expandAlloydbInstanceClientConnectionConfigRequireConnectors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceClientConnectionConfigSslConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSslMode, err := expandAlloydbInstanceClientConnectionConfigSslConfigSslMode(original["ssl_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sslMode"] = transformedSslMode
	}

	return transformed, nil
}

func expandAlloydbInstanceClientConnectionConfigSslConfigSslMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstancePscInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAttachmentLink, err := expandAlloydbInstancePscInstanceConfigServiceAttachmentLink(original["service_attachment_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachmentLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachmentLink"] = transformedServiceAttachmentLink
	}

	transformedAllowedConsumerProjects, err := expandAlloydbInstancePscInstanceConfigAllowedConsumerProjects(original["allowed_consumer_projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedConsumerProjects); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedConsumerProjects"] = transformedAllowedConsumerProjects
	}

	transformedPscDnsName, err := expandAlloydbInstancePscInstanceConfigPscDnsName(original["psc_dns_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPscDnsName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pscDnsName"] = transformedPscDnsName
	}

	return transformed, nil
}

func expandAlloydbInstancePscInstanceConfigServiceAttachmentLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstancePscInstanceConfigAllowedConsumerProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstancePscInstanceConfigPscDnsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAuthorizedExternalNetworks, err := expandAlloydbInstanceNetworkConfigAuthorizedExternalNetworks(original["authorized_external_networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthorizedExternalNetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authorizedExternalNetworks"] = transformedAuthorizedExternalNetworks
	}

	transformedEnablePublicIp, err := expandAlloydbInstanceNetworkConfigEnablePublicIp(original["enable_public_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnablePublicIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enablePublicIp"] = transformedEnablePublicIp
	}

	transformedEnableOutboundPublicIp, err := expandAlloydbInstanceNetworkConfigEnableOutboundPublicIp(original["enable_outbound_public_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableOutboundPublicIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableOutboundPublicIp"] = transformedEnableOutboundPublicIp
	}

	return transformed, nil
}

func expandAlloydbInstanceNetworkConfigAuthorizedExternalNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCidrRange, err := expandAlloydbInstanceNetworkConfigAuthorizedExternalNetworksCidrRange(original["cidr_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCidrRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cidrRange"] = transformedCidrRange
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAlloydbInstanceNetworkConfigAuthorizedExternalNetworksCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceNetworkConfigEnablePublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceNetworkConfigEnableOutboundPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbInstanceEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
