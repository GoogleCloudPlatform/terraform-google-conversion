package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestHclWriteBlocks_VariableReplacement(t *testing.T) {
	// Create a test block with a string value starting with "unknown."
	blocks := []*HCLResourceBlock{
		{
			Labels: []string{"test_resource", "test_name"},
			Value: cty.ObjectVal(map[string]cty.Value{
				"attr1": cty.StringVal("unknown.my_variable"),
				"attr2": cty.StringVal("normal_value"),
				"attr3": cty.ObjectVal(map[string]cty.Value{
					"nested_attr": cty.StringVal("unknown.another_variable"),
				}),
			}),
		},
	}

	output, err := HclWriteBlocks(blocks)
	assert.Nil(t, err)

	hclString := string(output)

	// Check if the variable references are correct
	// Note: hclwrite output format might vary slightly, but we expect var.my_variable
	// We expect "attr1 = var.my_variable" (no quotes around var.my_variable)
	assert.Contains(t, hclString, "attr1 = var.my_variable")
	assert.Contains(t, hclString, "nested_attr = var.another_variable")

	// Check if normal strings are still quoted
	assert.Contains(t, hclString, `attr2 = "normal_value"`)

	// Check if variable blocks are generated
	assert.True(t, strings.Contains(hclString, `variable "my_variable" {`), "missing variable block for my_variable")
	assert.True(t, strings.Contains(hclString, `type = string`), "missing type = string")
	assert.True(t, strings.Contains(hclString, `variable "another_variable" {`), "missing variable block for another_variable")
}
