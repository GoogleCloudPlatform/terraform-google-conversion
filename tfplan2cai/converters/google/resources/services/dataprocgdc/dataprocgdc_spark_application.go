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

package dataprocgdc

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataprocGdcSparkApplicationAssetType string = "dataprocgdc.googleapis.com/SparkApplication"

func ResourceConverterDataprocGdcSparkApplication() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataprocGdcSparkApplicationAssetType,
		Convert:   GetDataprocGdcSparkApplicationCaiObject,
	}
}

func GetDataprocGdcSparkApplicationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataprocgdc.googleapis.com/projects/{{project}}/locations/{{location}}/serviceInstances/{{serviceinstance}}/sparkApplications/{{spark_application_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataprocGdcSparkApplicationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataprocGdcSparkApplicationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataprocgdc/v1/rest",
				DiscoveryName:        "SparkApplication",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataprocGdcSparkApplicationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	pysparkApplicationConfigProp, err := expandDataprocGdcSparkApplicationPysparkApplicationConfig(d.Get("pyspark_application_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pyspark_application_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(pysparkApplicationConfigProp)) && (ok || !reflect.DeepEqual(v, pysparkApplicationConfigProp)) {
		obj["pysparkApplicationConfig"] = pysparkApplicationConfigProp
	}
	sparkApplicationConfigProp, err := expandDataprocGdcSparkApplicationSparkApplicationConfig(d.Get("spark_application_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_application_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkApplicationConfigProp)) && (ok || !reflect.DeepEqual(v, sparkApplicationConfigProp)) {
		obj["sparkApplicationConfig"] = sparkApplicationConfigProp
	}
	sparkRApplicationConfigProp, err := expandDataprocGdcSparkApplicationSparkRApplicationConfig(d.Get("spark_r_application_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_r_application_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkRApplicationConfigProp)) && (ok || !reflect.DeepEqual(v, sparkRApplicationConfigProp)) {
		obj["sparkRApplicationConfig"] = sparkRApplicationConfigProp
	}
	sparkSqlApplicationConfigProp, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfig(d.Get("spark_sql_application_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_sql_application_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkSqlApplicationConfigProp)) && (ok || !reflect.DeepEqual(v, sparkSqlApplicationConfigProp)) {
		obj["sparkSqlApplicationConfig"] = sparkSqlApplicationConfigProp
	}
	displayNameProp, err := expandDataprocGdcSparkApplicationDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	propertiesProp, err := expandDataprocGdcSparkApplicationProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}
	versionProp, err := expandDataprocGdcSparkApplicationVersion(d.Get("version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	applicationEnvironmentProp, err := expandDataprocGdcSparkApplicationApplicationEnvironment(d.Get("application_environment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_environment"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationEnvironmentProp)) && (ok || !reflect.DeepEqual(v, applicationEnvironmentProp)) {
		obj["applicationEnvironment"] = applicationEnvironmentProp
	}
	namespaceProp, err := expandDataprocGdcSparkApplicationNamespace(d.Get("namespace"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("namespace"); !tpgresource.IsEmptyValue(reflect.ValueOf(namespaceProp)) && (ok || !reflect.DeepEqual(v, namespaceProp)) {
		obj["namespace"] = namespaceProp
	}
	dependencyImagesProp, err := expandDataprocGdcSparkApplicationDependencyImages(d.Get("dependency_images"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dependency_images"); !tpgresource.IsEmptyValue(reflect.ValueOf(dependencyImagesProp)) && (ok || !reflect.DeepEqual(v, dependencyImagesProp)) {
		obj["dependencyImages"] = dependencyImagesProp
	}
	labelsProp, err := expandDataprocGdcSparkApplicationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandDataprocGdcSparkApplicationEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMainPythonFileUri, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigMainPythonFileUri(original["main_python_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainPythonFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainPythonFileUri"] = transformedMainPythonFileUri
	}

	transformedArgs, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedPythonFileUris, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigPythonFileUris(original["python_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonFileUris"] = transformedPythonFileUris
	}

	transformedJarFileUris, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	transformedFileUris, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocGdcSparkApplicationPysparkApplicationConfigArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigMainPythonFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigPythonFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationPysparkApplicationConfigArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMainJarFileUri, err := expandDataprocGdcSparkApplicationSparkApplicationConfigMainJarFileUri(original["main_jar_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainJarFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainJarFileUri"] = transformedMainJarFileUri
	}

	transformedMainClass, err := expandDataprocGdcSparkApplicationSparkApplicationConfigMainClass(original["main_class"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainClass"] = transformedMainClass
	}

	transformedArgs, err := expandDataprocGdcSparkApplicationSparkApplicationConfigArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedJarFileUris, err := expandDataprocGdcSparkApplicationSparkApplicationConfigJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	transformedFileUris, err := expandDataprocGdcSparkApplicationSparkApplicationConfigFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocGdcSparkApplicationSparkApplicationConfigArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigMainJarFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigMainClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkApplicationConfigArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkRApplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMainRFileUri, err := expandDataprocGdcSparkApplicationSparkRApplicationConfigMainRFileUri(original["main_r_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainRFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainRFileUri"] = transformedMainRFileUri
	}

	transformedArgs, err := expandDataprocGdcSparkApplicationSparkRApplicationConfigArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedFileUris, err := expandDataprocGdcSparkApplicationSparkRApplicationConfigFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocGdcSparkApplicationSparkRApplicationConfigArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataprocGdcSparkApplicationSparkRApplicationConfigMainRFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkRApplicationConfigArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkRApplicationConfigFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkRApplicationConfigArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQueryFileUri, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryFileUri(original["query_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryFileUri"] = transformedQueryFileUri
	}

	transformedQueryList, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryList(original["query_list"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryList); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryList"] = transformedQueryList
	}

	transformedScriptVariables, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfigScriptVariables(original["script_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScriptVariables); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scriptVariables"] = transformedScriptVariables
	}

	transformedJarFileUris, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfigJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	return transformed, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryList(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQueries, err := expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListQueries(original["queries"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueries); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queries"] = transformedQueries
	}

	return transformed, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListQueries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfigScriptVariables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocGdcSparkApplicationSparkSqlApplicationConfigJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocGdcSparkApplicationVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationApplicationEnvironment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationDependencyImages(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocGdcSparkApplicationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocGdcSparkApplicationEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
