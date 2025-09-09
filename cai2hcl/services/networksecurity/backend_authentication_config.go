package networksecurity

import (
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/caiasset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// BackendAuthenticationConfigAssetType is the CAI asset type name.
const BackendAuthenticationConfigAssetType string = "networksecurity.googleapis.com/BackendAuthenticationConfig"

// BackendAuthenticationConfigSchemaName is the TF resource schema name.
const BackendAuthenticationConfigSchemaName string = "google_network_security_backend_authentication_config"

// BackendAuthenticationConfigConverter for networksecurity backend authentication config resource.
type BackendAuthenticationConfigConverter struct {
	name   string
	schema map[string]*schema.Schema
}

// NewBackendAuthenticationConfigConverter returns an HCL converter.
func NewBackendAuthenticationConfigConverter(provider *schema.Provider) common.Converter {
	// NOTE(laurenzk): Hardcoded BAC-schema used to avoid incrementing `terraform-provider-google-beta`
	// version. This is to minimize impact on existing resources. Is removed from v6.
	schema := getBACSchema()

	return &BackendAuthenticationConfigConverter{
		name:   BackendAuthenticationConfigSchemaName,
		schema: schema,
	}
}

// Convert converts CAI assets to HCL resource blocks (Provider version: 7.0.1)
func (c *BackendAuthenticationConfigConverter) Convert(assets []*caiasset.Asset) ([]*common.HCLResourceBlock, error) {
	var blocks []*common.HCLResourceBlock
	var err error

	for _, asset := range assets {
		if asset == nil {
			continue
		} else if asset.Resource == nil || asset.Resource.Data == nil {
			return nil, fmt.Errorf("INVALID_ARGUMENT: Asset resource data is nil")
		} else if asset.Type != BackendAuthenticationConfigAssetType {
			return nil, fmt.Errorf("INVALID_ARGUMENT: Expected asset of type %s, but received %s", BackendAuthenticationConfigAssetType, asset.Type)
		}
		block, errConvert := c.convertResourceData(asset)
		blocks = append(blocks, block)
		if errConvert != nil {
			err = errors.Join(err, errConvert)
		}
	}
	return blocks, err
}

func (c *BackendAuthenticationConfigConverter) convertResourceData(asset *caiasset.Asset) (*common.HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("INVALID_ARGUMENT: Asset resource data is nil")
	}

	hcl, _ := flattenBackendAuthenticationConfig(asset.Resource)

	ctyVal, err := common.MapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}

	resourceName := hcl["name"].(string)
	return &common.HCLResourceBlock{
		Labels: []string{c.name, resourceName},
		Value:  ctyVal,
	}, nil
}

func flattenBackendAuthenticationConfig(resource *caiasset.AssetResource) (map[string]any, error) {
	result := make(map[string]any)

	resourceData := resource.Data

	result["name"] = flattenName(resourceData["name"].(string))
	result["labels"] = resourceData["labels"]
	result["description"] = resourceData["description"]
	result["client_certificate"] = resourceData["clientCertificate"]
	result["trust_config"] = resourceData["trustConfig"]
	result["well_known_roots"] = resourceData["wellKnownRoots"]
	result["project"] = flattenProjectName(resourceData["name"].(string))
	result["location"] = flattenLocation(resourceData["name"].(string))

	return result, nil
}

func getBACSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"client_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"labels": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"location": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "global",
		},
		"trust_config": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"well_known_roots": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"create_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"effective_labels": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"terraform_labels": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"update_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"project": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
	}
}
