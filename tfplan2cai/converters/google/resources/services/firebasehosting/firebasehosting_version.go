// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebasehosting/Version.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebasehosting

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseHostingVersionAssetType string = "firebasehosting.googleapis.com/Version"

func ResourceConverterFirebaseHostingVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseHostingVersionAssetType,
		Convert:   GetFirebaseHostingVersionCaiObject,
	}
}

func GetFirebaseHostingVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebasehosting.googleapis.com/sites/{{site_id}}/versions/{{version_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseHostingVersionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseHostingVersionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebasehosting/v1beta1/rest",
				DiscoveryName:        "Version",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseHostingVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	configProp, err := expandFirebaseHostingVersionConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}

	return obj, nil
}

func expandFirebaseHostingVersionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRewrites, err := expandFirebaseHostingVersionConfigRewrites(original["rewrites"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRewrites); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rewrites"] = transformedRewrites
	}

	transformedRedirects, err := expandFirebaseHostingVersionConfigRedirects(original["redirects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRedirects); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["redirects"] = transformedRedirects
	}

	transformedHeaders, err := expandFirebaseHostingVersionConfigHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	return transformed, nil
}

func expandFirebaseHostingVersionConfigRewrites(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGlob, err := expandFirebaseHostingVersionConfigRewritesGlob(original["glob"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGlob); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["glob"] = transformedGlob
		}

		transformedRegex, err := expandFirebaseHostingVersionConfigRewritesRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedPath, err := expandFirebaseHostingVersionConfigRewritesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedFunction, err := expandFirebaseHostingVersionConfigRewritesFunction(original["function"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["function"] = transformedFunction
		}

		transformedRun, err := expandFirebaseHostingVersionConfigRewritesRun(original["run"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRun); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["run"] = transformedRun
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirebaseHostingVersionConfigRewritesGlob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesFunction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRun(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceId, err := expandFirebaseHostingVersionConfigRewritesRunServiceId(original["service_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceId"] = transformedServiceId
	}

	transformedRegion, err := expandFirebaseHostingVersionConfigRewritesRunRegion(original["region"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["region"] = transformedRegion
	}

	return transformed, nil
}

func expandFirebaseHostingVersionConfigRewritesRunServiceId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRunRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGlob, err := expandFirebaseHostingVersionConfigRedirectsGlob(original["glob"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGlob); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["glob"] = transformedGlob
		}

		transformedRegex, err := expandFirebaseHostingVersionConfigRedirectsRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedStatusCode, err := expandFirebaseHostingVersionConfigRedirectsStatusCode(original["status_code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatusCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["statusCode"] = transformedStatusCode
		}

		transformedLocation, err := expandFirebaseHostingVersionConfigRedirectsLocation(original["location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["location"] = transformedLocation
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirebaseHostingVersionConfigRedirectsGlob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsStatusCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGlob, err := expandFirebaseHostingVersionConfigHeadersGlob(original["glob"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGlob); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["glob"] = transformedGlob
		}

		transformedRegex, err := expandFirebaseHostingVersionConfigHeadersRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedHeaders, err := expandFirebaseHostingVersionConfigHeadersHeaders(original["headers"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["headers"] = transformedHeaders
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirebaseHostingVersionConfigHeadersGlob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigHeadersRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigHeadersHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
