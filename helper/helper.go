package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func DecodeJSON(data map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	return nil
}

func ConvertServiceScope(url string) string {
	// This is a convenience map of short names used by the gcloud tool
	// to the GCE auth endpoints they alias to.
	scopeMap := map[string]string{
		"https://www.googleapis.com/auth/bigquery":                    "bigquery",
		"https://www.googleapis.com/auth/cloud-platform":              "cloud-platform",
		"https://www.googleapis.com/auth/source.full_control":         "cloud-source-repos",
		"https://www.googleapis.com/auth/source.read_only":            "cloud-source-repos-ro",
		"https://www.googleapis.com/auth/compute.readonly":            "compute-ro",
		"https://www.googleapis.com/auth/compute":                     "compute-rw",
		"https://www.googleapis.com/auth/datastore":                   "datastore",
		"https://www.googleapis.com/auth/logging.write":               "logging-write",
		"https://www.googleapis.com/auth/monitoring":                  "monitoring",
		"https://www.googleapis.com/auth/monitoring.read":             "monitoring-read",
		"https://www.googleapis.com/auth/monitoring.write":            "monitoring-write",
		"https://www.googleapis.com/auth/pubsub":                      "pubsub",
		"https://www.googleapis.com/auth/servicecontrol":              "service-control",
		"https://www.googleapis.com/auth/service.management.readonly": "service-management",
		"https://www.googleapis.com/auth/sqlservice":                  "sql",
		"https://www.googleapis.com/auth/sqlservice.admin":            "sql-admin",
		"https://www.googleapis.com/auth/devstorage.full_control":     "storage-full",
		"https://www.googleapis.com/auth/devstorage.read_only":        "storage-ro",
		"https://www.googleapis.com/auth/devstorage.read_write":       "storage-rw",
		"https://www.googleapis.com/auth/taskqueue":                   "taskqueue",
		"https://www.googleapis.com/auth/trace.append":                "trace",
		"https://www.googleapis.com/auth/cloud.useraccounts.readonly": "useraccounts-ro",
		"https://www.googleapis.com/auth/cloud.useraccounts":          "useraccounts-rw",
		"https://www.googleapis.com/auth/userinfo.email":              "userinfo-email",
	}

	if scope, ok := scopeMap[url]; ok {
		return scope
	}
	return url
}

func ParseFieldValue(str string, field string) string {
	strList := strings.Split(str, "/")
	for ix, item := range strList {
		if item == field && ix+1 < len(strList) {
			return strList[ix+1]
		}
	}
	return ""
}

// InstanceStateToMap converts state into a map[string]interface{}, using the schema as defined in r
// to coerce values to the appropriate type.
func InstanceStateToMap(r *schema.Resource, state *terraform.InstanceState) map[string]interface{} {
	ctyType := ctyTypeForResource(r)
	stateVal, err := state.AttrsAsObjectValue(ctyType)
	if err != nil {
		panic(fmt.Errorf("error parsing instance state as cty.Value: %v", err))
	}
	return CtyValToMap(stateVal, ctyType)
}

func MapToInstanceState(r *schema.Resource, m map[string]interface{}) *terraform.InstanceState {
	state := terraform.NewInstanceStateShimmedFromValue(MapToCtyVal(m, r.CoreConfigSchema().ImpliedType()), r.SchemaVersion)
	// Patch the schema version as a string; this seems to be a bug in the underlying Terraform code.
	state.Meta["schema_version"] = fmt.Sprintf("%v", r.SchemaVersion)
	return state
}

func ResourceConfigToMap(config *terraform.ResourceConfig) map[string]interface{} {
	return config.Raw
}

func MapToResourceConfig(r *schema.Resource, m map[string]interface{}) *terraform.ResourceConfig {
	schema := r.CoreConfigSchema()
	return terraform.NewResourceConfigShimmed(MapToCtyVal(m, schema.ImpliedType()), schema)
}

func CtyValToMap(val cty.Value, t cty.Type) map[string]interface{} {
	b, err := ctyjson.Marshal(val, t)
	if err != nil {
		panic(fmt.Errorf("error marshaling cty.Value as JSON: %v", err))
	}
	var ret map[string]interface{}
	if err := json.Unmarshal(b, &ret); err != nil {
		panic(fmt.Errorf("error unmarshaling JSON as map[string]interface{}: %v", err))
	}
	return ret
}

func MapToCtyVal(m map[string]interface{}, t cty.Type) cty.Value {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(fmt.Errorf("error marshaling map as JSON: %v", err))
	}
	ret, err := ctyjson.Unmarshal(b, t)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling JSON as cty.Value: %v", err))
	}
	return ret
}

func MapToCtyValWithSchema(m map[string]interface{}, s map[string]*schema.Schema) cty.Value {
	b, err := json.Marshal(&m)
	if err != nil {
		panic(fmt.Errorf("error marshaling map as JSON: %v", err))
	}
	ret, err := ctyjson.Unmarshal(b, schema.InternalMap(s).CoreConfigSchema().ImpliedType())
	if err != nil {
		panic(fmt.Errorf("error unmarshaling JSON as cty.Value: %v", err))
	}
	return ret
}

func ctyTypeForResource(r *schema.Resource) cty.Type {
	return r.CoreConfigSchema().ImpliedType()
}
