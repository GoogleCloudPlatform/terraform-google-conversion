package caitest

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/conversion"

	"github.com/stretchr/testify/assert"
)

func TestTFPlanToCAI_noPlanJSON(t *testing.T) {
	ctx := context.Background()
	jsonPlan := []byte{}
	options := &conversion.TFPlanToCAIOptions{}
	assets, err := conversion.TFPlanToCAI(ctx, jsonPlan, options)
	assert.Empty(t, assets)
	assert.Error(t, err)
}

func TestTFPlanToCAI_noResourceChanges(t *testing.T) {
	ctx := context.Background()
	f := "../testdata/empty-0.13.7.tfplan.json"
	jsonPlan, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatalf("Error parsing %s: %s", f, err)
	}
	options := &conversion.TFPlanToCAIOptions{}
	assets, err := conversion.TFPlanToCAI(ctx, jsonPlan, options)
	assert.Empty(t, assets)
	assert.Empty(t, err)
}
