// Package cai2hcl converts CAI assets to hcl bytes.
package cai2hcl

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"go.uber.org/zap"
)

// Options is a struct for options so that adding new options does not
// require updating function signatures all along the pipe.
type Options struct {
	ErrorLogger *zap.Logger
}

// Convert converts Asset into HCL.
func Convert(assets []*caiasset.Asset, options *Options) ([]byte, error) {
	if options == nil || options.ErrorLogger == nil {
		return nil, fmt.Errorf("logger is not initialized")
	}

	// Group resources from the same tf resource type for convert.
	// tf -> cai has 1:N mappings occasionally
	// The following resource has 1:2 mappings:
	//   google_project
	//   google_bigtable_instance
	groups := groupAssets(assets)

	f := hclwrite.NewFile()
	rootBody := f.Body()
	for _, v := range groups {
		converter, ok := Converters()[v[0].Type]
		if !ok {
			options.ErrorLogger.Warn(fmt.Sprintf("type %s is not supported", v[0].Type))
			continue
		}

		id, val, iamVal, err := converter.Convert(v...)
		if err != nil {
			return nil, err
		}
		if val != cty.NilVal {
			block := rootBody.AppendNewBlock("resource", []string{converter.TFResourceType(), id})
			if err := hclWriteBlock(val, block.Body()); err != nil {
				return nil, err
			}
		}
		if iamVal != cty.NilVal {
			block := rootBody.AppendNewBlock("resource", []string{converter.TFResourceType() + "_iam_policy", id + "_iam_policy"})
			if err := hclWriteBlock(iamVal, block.Body()); err != nil {
				return nil, err
			}
		}
	}

	t, err := printer.Format(f.Bytes())
	options.ErrorLogger.Debug(string(t))
	return t, err
}

func groupAssets(assets []*caiasset.Asset) map[string][]*caiasset.Asset {
	groups := make(map[string][]*caiasset.Asset)
	var exceptions []*caiasset.Asset
	for _, asset := range assets {
		if asset == nil || asset.Type == "" || asset.Name == "" {
			continue
		}
		if asset.Type == "bigtableadmin.googleapis.com/Cluster" || asset.Type == "cloudbilling.googleapis.com/ProjectBillingInfo" {
			exceptions = append(exceptions, asset)
		} else {
			groups[asset.Name] = []*caiasset.Asset{asset}
		}
	}
	for _, asset := range exceptions {
		switch asset.Type {
		case "cloudbilling.googleapis.com/ProjectBillingInfo":
			project := parseFieldValue(asset.Name, "projects")
			projectAssetName := fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%s", project)
			if _, ok := groups[projectAssetName]; ok {
				groups[projectAssetName] = append(groups[projectAssetName], asset)
			}
		case "bigtableadmin.googleapis.com/Cluster":
			instanceName := strings.Split(asset.Name, "/clusters/")[0]
			if _, ok := groups[instanceName]; ok {
				groups[instanceName] = append(groups[instanceName], asset)
			}
		}
	}
	return groups
}

func hclWriteBlock(val cty.Value, body *hclwrite.Body) error {
	if val.IsNull() {
		return nil
	}
	if !val.Type().IsObjectType() {
		return fmt.Errorf("expect object type only, but type = %s", val.Type().FriendlyName())
	}
	it := val.ElementIterator()
	for it.Next() {
		objKey, objVal := it.Element()
		if objVal.IsNull() {
			continue
		}
		objValType := objVal.Type()
		switch {
		case objValType.IsObjectType():
			newBlock := body.AppendNewBlock(objKey.AsString(), nil)
			if err := hclWriteBlock(objVal, newBlock.Body()); err != nil {
				return err
			}
		case objValType.IsCollectionType():
			if objVal.LengthInt() == 0 {
				continue
			}
			// Presumes map should not contain object type.
			if !objValType.IsMapType() && objValType.ElementType().IsObjectType() {
				listIterator := objVal.ElementIterator()
				for listIterator.Next() {
					_, listVal := listIterator.Element()
					subBlock := body.AppendNewBlock(objKey.AsString(), nil)
					if err := hclWriteBlock(listVal, subBlock.Body()); err != nil {
						return err
					}
				}
				continue
			}
			fallthrough
		default:
			if objValType.FriendlyName() == "string" && objVal.AsString() == "" {
				continue
			}
			body.SetAttributeValue(objKey.AsString(), objVal)
		}
	}
	return nil
}
