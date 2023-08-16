// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataproc

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataprocAutoscalingPolicyIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"policy_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type DataprocAutoscalingPolicyIamUpdater struct {
	project  string
	location string
	policyId string
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func DataprocAutoscalingPolicyIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("policy_id"); ok {
		values["policy_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/autoscalingPolicies/(?P<policy_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<policy_id>[^/]+)", "(?P<location>[^/]+)/(?P<policy_id>[^/]+)", "(?P<policy_id>[^/]+)"}, d, config, d.Get("policy_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocAutoscalingPolicyIamUpdater{
		project:  values["project"],
		location: values["location"],
		policyId: values["policy_id"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("policy_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting policy_id: %s", err)
	}

	return u, nil
}

func DataprocAutoscalingPolicyIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/autoscalingPolicies/(?P<policy_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<policy_id>[^/]+)", "(?P<location>[^/]+)/(?P<policy_id>[^/]+)", "(?P<policy_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocAutoscalingPolicyIamUpdater{
		project:  values["project"],
		location: values["location"],
		policyId: values["policy_id"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("policy_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting policy_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataprocAutoscalingPolicyIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAutoscalingPolicyUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataprocAutoscalingPolicyIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAutoscalingPolicyUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataprocAutoscalingPolicyIamUpdater) qualifyAutoscalingPolicyUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataprocBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", u.project, u.location, u.policyId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataprocAutoscalingPolicyIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/autoscalingPolicies/%s", u.project, u.location, u.policyId)
}

func (u *DataprocAutoscalingPolicyIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataproc-autoscalingpolicy-%s", u.GetResourceId())
}

func (u *DataprocAutoscalingPolicyIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataproc autoscalingpolicy %q", u.GetResourceId())
}
