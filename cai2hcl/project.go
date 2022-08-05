package cai2hcl

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tpg "github.com/hashicorp/terraform-provider-google/google"
	"github.com/zclconf/go-cty/cty"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// ProjectAssetType is the CAI asset type name for project.
const ProjectAssetType string = "cloudresourcemanager.googleapis.com/Project"

// ProjectConverter for compute project resource.
type ProjectConverter struct {
	Resource *tfschema.Resource
}

// NewProjectConverter returns an HCL converter for compute project.
func NewProjectConverter() *ProjectConverter {
	return &ProjectConverter{
		Resource: tpg.Provider().ResourcesMap["google_project"],
	}
}

// TFResourceType returns terraform resource type.
func (c *ProjectConverter) TFResourceType() string {
	return "google_project"
}

// Convert converts asset resource data.
func (c *ProjectConverter) Convert(assets ...*caiasset.Asset) (id string, data cty.Value, iam cty.Value, err error) {
	if len(assets) == 0 || assets[0] == nil {
		return "", cty.NilVal, cty.NilVal, fmt.Errorf("asset does not provide enough data for conversion")
	}

	id = parseFieldValue(assets[0].Name, "projects")
	var projectAsset *caiasset.Asset
	var billingAsset *caiasset.Asset
	for _, asset := range assets {
		if asset.Type == "cloudbilling.googleapis.com/ProjectBillingInfo" {
			billingAsset = asset
		} else {
			projectAsset = asset
		}
	}

	hclData, err := c.convertResourceData(projectAsset)
	if err != nil {
		return "", cty.NilVal, cty.NilVal, err
	}
	var val cty.Value
	if hclData != nil {
		billingAccount := c.convertBilling(billingAsset)
		hclData["billing_account"] = billingAccount
		val, err = mapToCtyValWithSchema(hclData, c.Resource.Schema)
		if err != nil {
			return "", cty.NilVal, cty.NilVal, err
		}
	}

	iamVal, err := c.convertIAM(projectAsset)
	if err != nil {
		return "", cty.NilVal, cty.NilVal, err
	}

	return id, val, iamVal, nil
}

func (c *ProjectConverter) convertIAM(asset *caiasset.Asset) (cty.Value, error) {
	if asset == nil {
		return cty.NilVal, fmt.Errorf("asset does not provide enough data for conversion")
	}
	if asset.IAMPolicy == nil {
		return cty.NilVal, nil
	}
	project := parseFieldValue(asset.Name, "projects")
	policyData, err := json.Marshal(asset.IAMPolicy)
	if err != nil {
		return cty.NilVal, err
	}

	return cty.ObjectVal(
		map[string]cty.Value{
			"project":     cty.StringVal(project),
			"policy_data": cty.StringVal(string(policyData)),
		},
	), nil
}

func (c *ProjectConverter) convertBilling(asset *caiasset.Asset) string {
	if asset != nil && asset.Resource != nil && asset.Resource.Data != nil {
		return strings.TrimPrefix(asset.Resource.Data["billingAccountName"].(string), "billingAccounts/")
	}
	return ""
}

func (c *ProjectConverter) convertResourceData(asset *caiasset.Asset) (map[string]interface{}, error) {
	if asset == nil {
		return nil, fmt.Errorf("asset does not provide enough data for conversion")
	}
	if asset.Resource == nil || asset.Resource.Data == nil {
		return nil, nil
	}
	var project *cloudresourcemanager.Project
	if err := decodeJSON(asset.Resource.Data, &project); err != nil {
		return nil, err
	}

	hclData := make(map[string]interface{})
	hclData["name"] = project.Name
	hclData["project_id"] = project.ProjectId
	hclData["labels"] = project.Labels
	if strings.Contains(asset.Resource.Parent, "folders/") {
		hclData["folder_id"] = parseFieldValue(asset.Resource.Parent, "folders")
	} else if strings.Contains(asset.Resource.Parent, "organizations/") {
		hclData["org_id"] = parseFieldValue(asset.Resource.Parent, "organizations")
	}

	return hclData, nil
}
