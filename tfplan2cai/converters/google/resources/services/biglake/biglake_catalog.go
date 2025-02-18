// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/biglake/Catalog.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package biglake

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BiglakeCatalogAssetType string = "biglake.googleapis.com/Catalog"

func ResourceConverterBiglakeCatalog() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BiglakeCatalogAssetType,
		Convert:   GetBiglakeCatalogCaiObject,
	}
}

func GetBiglakeCatalogCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//biglake.googleapis.com/projects/{{project}}/locations/{{location}}/catalogs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBiglakeCatalogApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BiglakeCatalogAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/biglake/v1/rest",
				DiscoveryName:        "Catalog",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBiglakeCatalogApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})

	return obj, nil
}
