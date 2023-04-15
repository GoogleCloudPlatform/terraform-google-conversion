package cai2hcl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

func assertTestData(fileName string) (err error) {
	assetFilePath := fmt.Sprintf("../testdata/%s.json", fileName)
	expectedTfFilePath := fmt.Sprintf("../testdata/%s.tf", fileName)
	assetPayload, err := ioutil.ReadFile(assetFilePath)
	if err != nil {
		return fmt.Errorf("cannot open %s, got: %s", assetFilePath, err)
	}
	want, err := ioutil.ReadFile(expectedTfFilePath)
	if err != nil {
		return fmt.Errorf("cannot open %s, got: %s", expectedTfFilePath, err)
	}

	var assets []*caiasset.Asset
	if err := json.Unmarshal(assetPayload, &assets); err != nil {
		return fmt.Errorf("cannot unmarshal: %s", err)
	}

	logger, err := zap.NewDevelopment()

	got, err := Convert(assets, &Options{
		ErrorLogger: logger,
	})
	if err != nil {
		return err
	}
	if diff := cmp.Diff(string(want), string(got)); diff != "" {
		return fmt.Errorf("cmp.Diff() got diff (-want +got): %s", diff)
	}

	return nil
}
