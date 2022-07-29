package conversion

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/conversion/hcl"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/model"

	hashicorpcty "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"go.uber.org/zap"
)

// CAIToHCLOptions is a struct for options so that adding new options does not
// require updating function signatures all along the pipe
type CAIToHCLOptions struct {
	ErrorLogger *zap.Logger
}

// CAIToHCL converts Asset into HCL.
func CAIToHCL(assets []*model.Asset, options *CAIToHCLOptions) ([]byte, error) {
	if options == nil || options.ErrorLogger == nil {
		return nil, fmt.Errorf("logger is not initialized")
	}
	f := hclwrite.NewFile()
	rootBody := f.Body()
	for _, asset := range assets {
		converter, ok := hcl.Converters()[asset.Type]
		if !ok {
			options.ErrorLogger.Warn(fmt.Sprintf("type %s is not supported", asset.Type))
			continue
		}

		id, data, err := converter.Convert(asset)
		if err != nil {
			options.ErrorLogger.Warn(fmt.Sprintf("failed to convert asset: %s", err))
			continue
		}

		block := rootBody.AppendNewBlock("resource", []string{converter.TFResourceName, id})

		val := mapToCtyValWithSchema(data, converter.Resource.Schema)
		if err := hclWriteBlock(val, block.Body()); err != nil {
			return nil, err
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
		// TODO: remove
		// fmt.Println(objKey.AsString(), objVal.Type().FriendlyName(), objVal.Type().HasDynamicTypes())
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
			// TODO: remove
			// fmt.Printf("key = %s, collection type = %s, length = %d\n", objKey.AsString(), objValType.FriendlyName(), objVal.LengthInt())
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

func mapToCtyValWithSchema(m map[string]interface{}, s map[string]*schema.Schema) cty.Value {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(fmt.Errorf("error marshaling map as JSON: %v", err))
	}
	ty, err := hashicorpCtyTypeToZclconfCtyType(schema.InternalMap(s).CoreConfigSchema().ImpliedType())
	if err != nil {
		panic(fmt.Errorf("error casting type: %v", err))
	}
	ret, err := ctyjson.Unmarshal(b, ty)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling JSON as cty.Value: %v", err))
	}
	return ret
}

func hashicorpCtyTypeToZclconfCtyType(t hashicorpcty.Type) (cty.Type, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return cty.NilType, err
	}
	var ret cty.Type
	if err := json.Unmarshal(b, &ret); err != nil {
		return cty.NilType, err
	}
	return ret, nil
}
