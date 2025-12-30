package common

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// HclWriteBlocks prints HCLResourceBlock objects as string.
func HclWriteBlocks(blocks []*HCLResourceBlock) ([]byte, error) {
	f := hclwrite.NewFile()
	rootBody := f.Body()

	variableMap := make(map[string]struct{})

	for _, resourceBlock := range blocks {
		hclBlock := rootBody.AppendNewBlock("resource", resourceBlock.Labels)
		if err := hclWriteBlock(resourceBlock.Value, hclBlock.Body(), variableMap); err != nil {
			return nil, err
		}
	}

	varNames := make([]string, 0, len(variableMap))
	for name := range variableMap {
		varNames = append(varNames, name)
	}
	sort.Strings(varNames)

	for _, name := range varNames {
		rootBody.AppendNewline()
		varBlock := rootBody.AppendNewBlock("variable", []string{name})

		// type = string
		t := hcl.Traversal{
			hcl.TraverseRoot{Name: "string"},
		}
		varBlock.Body().SetAttributeTraversal("type", t)
	}

	return f.Bytes(), nil
}

func hclWriteBlock(val cty.Value, body *hclwrite.Body, variableMap map[string]struct{}) error {
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
			if err := hclWriteBlock(objVal, newBlock.Body(), variableMap); err != nil {
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
					if err := hclWriteBlock(listVal, subBlock.Body(), variableMap); err != nil {
						return err
					}
				}
				continue
			}
			fallthrough
		default:
			if objValType == cty.String {
				strVal := objVal.AsString()
				if strVal == "" {
					continue
				}

				if strings.HasPrefix(strVal, "unknown.") {
					varName := strings.TrimPrefix(strVal, "unknown.")
					variableMap[varName] = struct{}{}

					t := hcl.Traversal{
						hcl.TraverseRoot{Name: "var"},
						hcl.TraverseAttr{Name: varName},
					}
					body.SetAttributeTraversal(objKey.AsString(), t)
					continue
				}
			}

			if objValType.FriendlyName() == "string" && objVal.AsString() == "" {
				continue
			}
			body.SetAttributeValue(objKey.AsString(), objVal)
		}
	}
	return nil
}
