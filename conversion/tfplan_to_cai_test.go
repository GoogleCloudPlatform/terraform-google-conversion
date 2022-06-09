package conversion

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTFPlanToCAI_noPlanJSON(t *testing.T) {
	jsonPlan := []byte{}
	options := &TFPlanToCAIOptions{}
	assets, err := TFPlanToCAI(jsonPlan, options)
	assert.Empty(t, assets)
	assert.Error(t, err)
}

func TestTFPlanToCAI_noResourceChanges(t *testing.T) {
	f := "../testdata/empty-0.13.7.tfplan.json"
	jsonPlan, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatalf("Error parsing %s: %s", f, err)
	}
	options := &TFPlanToCAIOptions{}
	assets, err := TFPlanToCAI(jsonPlan, options)
	assert.Empty(t, assets)
	assert.Empty(t, err)
}
