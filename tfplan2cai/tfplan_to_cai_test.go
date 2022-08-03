// Package tfplan2cai converts tfplan to CAI assets.
package tfplan2cai

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert_noPlanJSON(t *testing.T) {
	ctx := context.Background()
	jsonPlan := []byte{}
	options := &Options{}
	assets, err := Convert(ctx, jsonPlan, options)
	assert.Empty(t, assets)
	assert.Error(t, err)
}

func TestConvert_noResourceChanges(t *testing.T) {
	ctx := context.Background()
	f := "../testdata/empty-0.13.7.tfplan.json"
	jsonPlan, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatalf("Error parsing %s: %s", f, err)
	}
	options := &Options{}
	assets, err := Convert(ctx, jsonPlan, options)
	assert.Empty(t, assets)
	assert.Empty(t, err)
}
