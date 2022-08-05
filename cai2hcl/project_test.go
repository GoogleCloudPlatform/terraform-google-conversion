package cai2hcl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

func TestConvertProject(t *testing.T) {
	cases := []struct {
		name string
	}{
		{name: "project_create"},
		{name: "project_iam"},
	}
	for i := range cases {
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

			var assets []*caiasset.Asset
			if err := json.Unmarshal(assetPayload, &assets); err != nil {
				t.Fatalf("cannot unmarshal: %s", err)
			}

			got, err := Convert(assets, &Options{
				ErrorLogger: zap.NewExample(),
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
