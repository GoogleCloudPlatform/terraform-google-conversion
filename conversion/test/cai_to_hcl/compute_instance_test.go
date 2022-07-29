package hcltest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/conversion"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/model"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

func TestComputeInstanceToHCL(t *testing.T) {
	cases := []struct {
		name string
	}{
		{name: "full_compute_instance"},
	}
	for i := range cases {
		// Allocate a variable to make sure test can run in parallel.
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assetFilePath := fmt.Sprintf("../testdata/%s.json", c.name)
			expectedTFFilePath := fmt.Sprintf("../testdata/%s.tf", c.name)
			assetPayload, err := ioutil.ReadFile(assetFilePath)
			if err != nil {
				t.Fatalf("cannot open %s, got: %s", assetFilePath, err)
			}
			want, err := ioutil.ReadFile(expectedTFFilePath)
			if err != nil {
				t.Fatalf("cannot open %s, got: %s", expectedTFFilePath, err)
			}

			var assets []*model.Asset
			if err := json.Unmarshal(assetPayload, &assets); err != nil {
				t.Fatalf("cannot unmarshal: %s", err)
			}

			got, err := conversion.CAIToHCL(assets, &conversion.CAIToHCLOptions{
				ErrorLogger: zap.NewNop(),
			})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(want), string(got)); diff != "" {
				t.Errorf("cmp.Diff() got diff (-want +got): %s", diff)
			}
		})
	}
}
