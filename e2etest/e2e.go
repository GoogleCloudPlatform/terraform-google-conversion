package e2etest

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai"
	"github.com/GoogleCloudPlatform/terraform-validator/converters/google"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/hcl/printer"
)

type testData struct {
	TFVersion string
	Provider  map[string]string
	Project   map[string]string
	Time      map[string]string
	OrgID     string
	FolderID  string
	Ancestry  string
}

// initTestData initializes the variables used for testing. As tests rely on
// environment variables, the parsing of those are only done once.
func initTestData() *testData {
	credentials := getTestCredsFromEnv()
	project := getTestProjectFromEnv()
	org := getTestOrgFromEnv(nil)
	billingAccount := getTestBillingAccountFromEnv(nil)
	folder, ok := os.LookupEnv("TEST_FOLDER_ID")
	if !ok {
		log.Printf("Missing required env var TEST_FOLDER_ID. Default (%s) will be used.", defaultFolder)
		folder = defaultFolder
	}
	ancestry, ok := os.LookupEnv("TEST_ANCESTRY")
	if !ok {
		log.Printf("Missing required env var TEST_ANCESTRY. Default (%s) will be used.", defaultAncestry)
		ancestry = defaultAncestry
	}
	providerVersion := defaultProviderVersion
	//As time is not information in terraform resource data, time is fixed for testing purposes
	fixedTime := time.Date(2021, time.April, 14, 15, 16, 17, 0, time.UTC)
	return &testData{
		TFVersion: "0.12",
		Provider: map[string]string{
			"version":     providerVersion,
			"project":     project,
			"credentials": credentials,
		},
		Time: map[string]string{
			"RFC3339Nano": fixedTime.Format(time.RFC3339Nano),
		},
		Project: map[string]string{
			"Name":               "My Project Name",
			"ProjectId":          "my-project-id",
			"BillingAccountName": billingAccount,
			"Number":             "1234567890",
		},
		OrgID:    org,
		FolderID: folder,
		Ancestry: ancestry,
	}
}

func roundtripTest(t *testing.T, name string, tmpDir string, data *testData) {
	dir, err := ioutil.TempDir(tmpDir, "terraform")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// tf file for credential and versions
	generateHeaderFile(t, filepath.Join(dir, "header.tf"), data)

	// fill template to generate the tf file
	generateTestFiles(t, "../testdata", dir, name+".tf", data)

	// terraform init + terraform plan
	terraformWorkflow(t, dir, name)

	// convert from tf json plan to assets
	tfJSONPlanPath := filepath.Join(dir, name+".tfplan.json")
	jsonPlan, err := ioutil.ReadFile(tfJSONPlanPath)
	if err != nil {
		t.Fatalf("cannot read %q, got: %s", tfJSONPlanPath, err)
	}
	gotAssets, err := tfplan2cai.Convert(context.Background(),
		jsonPlan,
		&tfplan2cai.Options{
			DefaultProject: data.Provider["project"],
		},
	)
	if err != nil {
		t.Fatalf("tfplan2cai.Convert() = %s, want = nil", err)
	}

	// convert from assets to hcl
	var assetsInput []*google.Asset
	for ix := range gotAssets {
		assetsInput = append(assetsInput, &gotAssets[ix])
	}
	gotTFPlanBytes, err := cai2hcl.Convert(assetsInput, &cai2hcl.Options{})
	if err != nil {
		t.Fatalf("cai2hcl.Convert() = %s, want = nil", err)
	}

	// compare results
	tfFilePath := filepath.Join(dir, name+".tf")
	tfBytes, err := ioutil.ReadFile(tfFilePath)
	if err != nil {
		t.Fatalf("Error parsing %s: %s", tfFilePath, err)
	}
	wantTFPlanBytes, err := printer.Format(tfBytes)
	if err != nil {
		t.Fatalf("Error format %s: %s", tfFilePath, err)
	}
	if diff := cmp.Diff(string(wantTFPlanBytes), string(gotTFPlanBytes)); diff != "" {
		t.Fatalf("want = %v, got = %v, diff = %s", string(wantTFPlanBytes), string(gotTFPlanBytes), diff)
	}
}

func terraformWorkflow(t *testing.T, dir, name string) {
	terraformInit(t, "terraform", dir)
	terraformPlan(t, "terraform", dir, name+".tfplan")
	payload := terraformShow(t, "terraform", dir, name+".tfplan")
	saveFile(t, dir, name+".tfplan.json", payload)
}

func terraformInit(t *testing.T, executable, dir string) {
	terraformExec(t, executable, dir, "init", "-input=false")
}

func terraformPlan(t *testing.T, executable, dir, tfplan string) {
	terraformExec(t, executable, dir, "plan", "-input=false", "-refresh=false", "-out", tfplan)
}

func terraformShow(t *testing.T, executable, dir, tfplan string) []byte {
	return terraformExec(t, executable, dir, "show", "--json", tfplan)
}

func terraformExec(t *testing.T, executable, dir string, args ...string) []byte {
	cmd := exec.Command(executable, args...)
	cmd.Env = []string{"HOME=" + filepath.Join(dir, "fakehome")}
	cmd.Dir = dir
	wantError := false
	payload, _ := run(t, cmd, wantError)
	return payload
}

func saveFile(t *testing.T, dir, filename string, payload []byte) {
	fullpath := filepath.Join(dir, filename)
	f, err := os.Create(fullpath)
	if err != nil {
		t.Fatalf("error while creating file %s, error %v", fullpath, err)
	}
	_, err = f.Write(payload)
	if err != nil {
		t.Fatalf("error while writing to file %s, error %v", fullpath, err)
	}
}

// run a command and call t.Fatal on non-zero exit.
func run(t *testing.T, cmd *exec.Cmd, wantError bool) ([]byte, []byte) {
	var stderr, stdout bytes.Buffer
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	err := cmd.Run()
	if gotError := (err != nil); gotError != wantError {
		t.Fatalf("running %s: \nerror=%v \nstderr=%s \nstdout=%s", cmdToString(cmd), err, stderr.String(), stdout.String())
	}
	// Print env, stdout and stderr if verbose flag is used.
	if len(cmd.Env) != 0 {
		t.Logf("=== Environment Variable of %s ===", cmdToString(cmd))
		t.Log(strings.Join(cmd.Env, "\n"))
	}
	if stdout.String() != "" {
		t.Logf("=== STDOUT of %s ===", cmdToString(cmd))
		t.Log(stdout.String())
	}
	if stderr.String() != "" {
		t.Logf("=== STDERR of %s ===", cmdToString(cmd))
		t.Log(stderr.String())
	}
	return stdout.Bytes(), stderr.Bytes()
}

// cmdToString clones the logic of https://golang.org/pkg/os/exec/#Cmd.String.
func cmdToString(c *exec.Cmd) string {
	// report the exact executable path (plus args)
	b := new(strings.Builder)
	b.WriteString(c.Path)
	for _, a := range c.Args[1:] {
		b.WriteByte(' ')
		b.WriteString(a)
	}
	return b.String()
}

func generateTestFiles(t *testing.T, sourceDir string, targetDir string, fileName string, data interface{}) {
	funcMap := template.FuncMap{
		"pastLastSlash": func(s string) string {
			split := strings.Split(s, "/")
			return split[len(split)-1]
		},
	}
	tmpls, err := template.New("").Funcs(funcMap).
		ParseGlob(filepath.Join(sourceDir, fileName))
	if err != nil {
		t.Fatalf("generateTestFiles: %v", err)
	}
	for _, tmpl := range tmpls.Templates() {
		if tmpl.Name() == "" {
			continue // Skip base template.
		}
		path := filepath.Join(targetDir, tmpl.Name())
		f, err := os.Create(path)
		if err != nil {
			t.Fatalf("creating terraform file %v: %v", path, err)
		}
		if err := tmpl.Execute(f, data); err != nil {
			t.Fatalf("templating terraform file %v: %v", path, err)
		}
		if err := f.Close(); err != nil {
			t.Fatalf("closing file %v: %v", path, err)
		}
		t.Logf("Successfully created file %v", path)
	}
}

func generateHeaderFile(t *testing.T, path string, data interface{}) {
	t.Helper()
	headerTemplate := `
	terraform {
		required_providers {
			google = {
				source = "hashicorp/google"
				version = "~> {{.Provider.version}}"
			}
		}
	}

	provider "google" {
		{{if .Provider.credentials }}credentials = "{{.Provider.credentials}}"{{end}}
	}
	`
	tmpl, err := template.New("header").Parse(headerTemplate)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("creating terraform file %v: %v", path, err)
	}
	if err := tmpl.Execute(f, data); err != nil {
		t.Fatalf("templating terraform file %v: %v", path, err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("closing file %v: %v", path, err)
	}
}
