// Package cai2hcl converts CAI assets to hcl bytes.
package cai2hcl

import (
	"fmt"

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
	f := hclwrite.NewFile()
	rootBody := f.Body()
	for _, asset := range assets {
		if asset == nil {
			continue
		}

		converter, ok := Converters()[asset.Type]
		if !ok {
			options.ErrorLogger.Warn(fmt.Sprintf("type %s is not supported", asset.Type))
			continue
		}

		if asset.Resource != nil && asset.Resource.Data != nil {
			id, val, err := converter.Convert(asset)
			if err != nil {
				return nil, err
			}
			block := rootBody.AppendNewBlock("resource", []string{converter.TFResourceType(), id})
			if err := hclWriteBlock(val, block.Body()); err != nil {
				return nil, err
			}
		}

		if asset.IAMPolicy != nil {
			id, ctyVal, err := converter.ConvertIAM(asset)
			if err != nil {
				return nil, err
			}
			block := rootBody.AppendNewBlock("resource", []string{converter.TFResourceType() + "_iam_policy", id})
			if err := hclWriteBlock(ctyVal, block.Body()); err != nil {
				return nil, err
			}
		}
	}

	t, err := printer.Format(f.Bytes())
	options.ErrorLogger.Debug(string(t))
	return t, err
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
