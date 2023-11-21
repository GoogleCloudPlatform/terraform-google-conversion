// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/google/go-cmp/cmp"
	terraformJSON "github.com/hashicorp/terraform-json"
)

// TestCLI tests the "convert" and "validate" subcommand against a generated .tfplan file.
func TestCLI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
		return
	}

	if os.Getenv("CREATE_TEST_PROJECT") != "" && os.Getenv("TEST_ORG_ID") != "" {
		orgID := os.Getenv("TEST_ORG_ID")
		dir, err := os.MkdirTemp(t.TempDir(), "terraform")
		if err != nil {
			t.Fatalf("os.MkdirTemp = %v", err)
		}
		// Do not use defer since it will execute before t.Parallel()
		t.Cleanup(func() {
			terraformDestroy(t, "terraform", dir, "")
		})
		createTestProject(t, dir, orgID)
	}

	// Test cases for each type of resource is defined here.
	cases := []struct {
		name                 string
		compareConvertOutput compareConvertOutputFunc
	}{
		{name: "example_project_iam_binding", compareConvertOutput: compareMergedIamBindingOutput},
		{name: "example_project_iam_member", compareConvertOutput: compareMergedIamMemberOutput},
	}

	for i := range cases {
		// Allocate a variable to make sure test can run in parallel.
		c := cases[i]

		// Add default convert comparison func if not set
		if c.compareConvertOutput == nil {
			c.compareConvertOutput = compareUnmergedConvertOutput
		}

		// Test both offline and online mode.
		for _, offline := range []bool{true, false} {
			offline := offline
			t.Run(fmt.Sprintf("tf=%s/offline=%t", c.name, offline), func(t *testing.T) {
				t.Parallel()
				// Create a temporary directory for running terraform.
				dir, err := os.MkdirTemp(tmpDir, "terraform")
				if err != nil {
					log.Fatal(err)
				}
				defer os.RemoveAll(dir)

				// Generate the <name>.tf and <name>_assets.json files into the temporary directory.
				generateTestFiles(t, "../testdata/templates", dir, c.name+".tf")
				generateTestFiles(t, "../testdata/templates", dir, c.name+".json")

				// Uses glob matching to match generateTestFiles internals.
				tfstateMatches, err := filepath.Glob(filepath.Join("../testdata/templates", c.name+".tfstate"))
				if err != nil {
					t.Fatalf("malformed glob: %v", err)
				}
				if tfstateMatches != nil {
					generateTestFiles(t, "../testdata/templates", dir, c.name+".tfstate")
					err = os.Rename(
						filepath.Join(dir, c.name+".tfstate"),
						filepath.Join(dir, "terraform.tfstate"),
					)
					if err != nil {
						t.Fatalf("renaming tfstate: %v", err)
					}
				}

				terraformWorkflow(t, dir, c.name)

				t.Run("cmd=convert", func(t *testing.T) {
					testConvertCommand(t, dir, c.name, c.name, offline, true, c.compareConvertOutput)
				})
			})
		}
	}
}

type compareConvertOutputFunc func(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool)

func compareUnmergedConvertOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, actual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}

// For merged IAM members, only consider whether the expected members are present.
func compareMergedIamMemberOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	var normalizedActual []caiasset.Asset
	for i := range expected {
		expectedAsset := expected[i]
		actualAsset := actual[i]

		// Copy actualAsset
		normalizedActualAsset := actualAsset

		expectedBindings := map[string]map[string]struct{}{}
		for _, binding := range expectedAsset.IAMPolicy.Bindings {
			expectedBindings[binding.Role] = map[string]struct{}{}
			for _, member := range binding.Members {
				expectedBindings[binding.Role][member] = struct{}{}
			}
		}

		iamPolicy := caiasset.IAMPolicy{}
		for _, binding := range actualAsset.IAMPolicy.Bindings {
			if expectedMembers, exists := expectedBindings[binding.Role]; exists {
				iamBinding := caiasset.IAMBinding{
					Role: binding.Role,
				}
				for _, member := range binding.Members {
					if _, exists := expectedMembers[member]; exists {
						iamBinding.Members = append(iamBinding.Members, member)
					}
				}
				iamPolicy.Bindings = append(iamPolicy.Bindings, iamBinding)
			}
		}
		normalizedActualAsset.IAMPolicy = &iamPolicy
		normalizedActual = append(normalizedActual, normalizedActualAsset)
	}

	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, normalizedActual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}

// For merged IAM bindings, only consider whether the expected bindings are as expected.
func compareMergedIamBindingOutput(t *testing.T, expected []caiasset.Asset, actual []caiasset.Asset, offline bool) {
	var normalizedActual []caiasset.Asset
	for i := range expected {
		expectedAsset := expected[i]
		actualAsset := actual[i]

		// Copy actualAsset
		normalizedActualAsset := actualAsset

		expectedBindings := map[string]struct{}{}
		for _, binding := range expectedAsset.IAMPolicy.Bindings {
			expectedBindings[binding.Role] = struct{}{}
		}

		iamPolicy := caiasset.IAMPolicy{}
		for _, binding := range actualAsset.IAMPolicy.Bindings {
			if _, exists := expectedBindings[binding.Role]; exists {
				iamPolicy.Bindings = append(iamPolicy.Bindings, binding)
			}
		}
		normalizedActualAsset.IAMPolicy = &iamPolicy
		normalizedActual = append(normalizedActual, normalizedActualAsset)
	}

	expectedAssets := normalizeAssets(t, expected, offline)
	actualAssets := normalizeAssets(t, normalizedActual, offline)
	if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
		t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
	}
}

func createTestProject(t *testing.T, dir string, orgID string) {
	generateTestFiles(t, "./", dir, "create_test_project.tf")
	// any terraform execute failure will trigger t.Fatal
	terraformInit(t, "terraform", dir)
	terraformPlan(t, "terraform", dir, "create_test_project.tfplan")
	terraformApply(t, "terraform", dir, "")
	// terraform show result contains format_version field which is required in unmarshal.
	b := terraformShow(t, "terraform", dir, "")
	var state terraformJSON.State
	err := json.Unmarshal(b, &state)
	if err != nil {
		t.Fatal(err)
	}

	var ok bool
	for _, resource := range state.Values.RootModule.Resources {
		if resource.Type == "google_project" {
			data.FolderID, ok = resource.AttributeValues["folder_id"].(string)
			if !ok {
				t.Fatalf("Failed to get folder ID from value %v", resource.AttributeValues["folder_id"])
			}
			data.Project["project"], ok = resource.AttributeValues["project_id"].(string)
			if !ok {
				t.Fatalf("Failed to get project ID from value %v", resource.AttributeValues["project_id"])
			}
		}
	}
	data.Ancestry = fmt.Sprintf("organizations/%s/folders/%s", orgID, data.FolderID)
	t.Logf("Successfully created folder_id=%v, project_id=%v", data.FolderID, data.Project["project"])
}
