// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/databasemigrationservice/MigrationJob.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package databasemigrationservice

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DatabaseMigrationServiceMigrationJobAssetType string = "datamigration.googleapis.com/MigrationJob"

func ResourceConverterDatabaseMigrationServiceMigrationJob() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DatabaseMigrationServiceMigrationJobAssetType,
		Convert:   GetDatabaseMigrationServiceMigrationJobCaiObject,
	}
}

func GetDatabaseMigrationServiceMigrationJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//datamigration.googleapis.com/projects/{{project}}/locations/{{location}}/migrationJobs/{{migration_job_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDatabaseMigrationServiceMigrationJobApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DatabaseMigrationServiceMigrationJobAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datamigration/v1/rest",
				DiscoveryName:        "MigrationJob",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDatabaseMigrationServiceMigrationJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDatabaseMigrationServiceMigrationJobDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeProp, err := expandDatabaseMigrationServiceMigrationJobType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	sourceProp, err := expandDatabaseMigrationServiceMigrationJobSource(d.Get("source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceProp)) && (ok || !reflect.DeepEqual(v, sourceProp)) {
		obj["source"] = sourceProp
	}
	destinationProp, err := expandDatabaseMigrationServiceMigrationJobDestination(d.Get("destination"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationProp)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	dumpFlagsProp, err := expandDatabaseMigrationServiceMigrationJobDumpFlags(d.Get("dump_flags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dump_flags"); !tpgresource.IsEmptyValue(reflect.ValueOf(dumpFlagsProp)) && (ok || !reflect.DeepEqual(v, dumpFlagsProp)) {
		obj["dumpFlags"] = dumpFlagsProp
	}
	performanceConfigProp, err := expandDatabaseMigrationServiceMigrationJobPerformanceConfig(d.Get("performance_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("performance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(performanceConfigProp)) && (ok || !reflect.DeepEqual(v, performanceConfigProp)) {
		obj["performanceConfig"] = performanceConfigProp
	}
	dumpPathProp, err := expandDatabaseMigrationServiceMigrationJobDumpPath(d.Get("dump_path"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dump_path"); !tpgresource.IsEmptyValue(reflect.ValueOf(dumpPathProp)) && (ok || !reflect.DeepEqual(v, dumpPathProp)) {
		obj["dumpPath"] = dumpPathProp
	}
	dumpTypeProp, err := expandDatabaseMigrationServiceMigrationJobDumpType(d.Get("dump_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dump_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(dumpTypeProp)) && (ok || !reflect.DeepEqual(v, dumpTypeProp)) {
		obj["dumpType"] = dumpTypeProp
	}
	staticIpConnectivityProp, err := expandDatabaseMigrationServiceMigrationJobStaticIpConnectivity(d.Get("static_ip_connectivity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("static_ip_connectivity"); ok || !reflect.DeepEqual(v, staticIpConnectivityProp) {
		obj["staticIpConnectivity"] = staticIpConnectivityProp
	}
	reverseSshConnectivityProp, err := expandDatabaseMigrationServiceMigrationJobReverseSshConnectivity(d.Get("reverse_ssh_connectivity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reverse_ssh_connectivity"); !tpgresource.IsEmptyValue(reflect.ValueOf(reverseSshConnectivityProp)) && (ok || !reflect.DeepEqual(v, reverseSshConnectivityProp)) {
		obj["reverseSshConnectivity"] = reverseSshConnectivityProp
	}
	vpcPeeringConnectivityProp, err := expandDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity(d.Get("vpc_peering_connectivity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc_peering_connectivity"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcPeeringConnectivityProp)) && (ok || !reflect.DeepEqual(v, vpcPeeringConnectivityProp)) {
		obj["vpcPeeringConnectivity"] = vpcPeeringConnectivityProp
	}
	labelsProp, err := expandDatabaseMigrationServiceMigrationJobEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDatabaseMigrationServiceMigrationJobDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpFlags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDumpFlags, err := expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlags(original["dump_flags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDumpFlags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dumpFlags"] = transformedDumpFlags
	}

	return transformed, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlagsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlagsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlagsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpFlagsDumpFlagsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobPerformanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDumpParallelLevel, err := expandDatabaseMigrationServiceMigrationJobPerformanceConfigDumpParallelLevel(original["dump_parallel_level"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDumpParallelLevel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dumpParallelLevel"] = transformedDumpParallelLevel
	}

	return transformed, nil
}

func expandDatabaseMigrationServiceMigrationJobPerformanceConfigDumpParallelLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobDumpType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobStaticIpConnectivity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandDatabaseMigrationServiceMigrationJobReverseSshConnectivity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVmIp, err := expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVmIp(original["vm_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vmIp"] = transformedVmIp
	}

	transformedVmPort, err := expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVmPort(original["vm_port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vmPort"] = transformedVmPort
	}

	transformedVm, err := expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVm(original["vm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vm"] = transformedVm
	}

	transformedVpc, err := expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVpc(original["vpc"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpc); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpc"] = transformedVpc
	}

	return transformed, nil
}

func expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVmIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVmPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobReverseSshConnectivityVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVpc, err := expandDatabaseMigrationServiceMigrationJobVpcPeeringConnectivityVpc(original["vpc"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpc); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpc"] = transformedVpc
	}

	return transformed, nil
}

func expandDatabaseMigrationServiceMigrationJobVpcPeeringConnectivityVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServiceMigrationJobEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
