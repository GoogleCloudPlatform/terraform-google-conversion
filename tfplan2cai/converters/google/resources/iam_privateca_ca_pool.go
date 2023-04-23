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

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

var PrivatecaCaPoolIamSchema = map[string]*schema.Schema{
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
	"ca_pool": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type PrivatecaCaPoolIamUpdater struct {
	project  string
	location string
	caPool   string
	d        TerraformResourceData
	Config   *transport_tpg.Config
}

func PrivatecaCaPoolIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("ca_pool"); ok {
		values["ca_pool"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/caPools/(?P<ca_pool>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<ca_pool>[^/]+)", "(?P<location>[^/]+)/(?P<ca_pool>[^/]+)"}, d, config, d.Get("ca_pool").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PrivatecaCaPoolIamUpdater{
		project:  values["project"],
		location: values["location"],
		caPool:   values["ca_pool"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("ca_pool", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting ca_pool: %s", err)
	}

	return u, nil
}

func PrivatecaCaPoolIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/caPools/(?P<ca_pool>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<ca_pool>[^/]+)", "(?P<location>[^/]+)/(?P<ca_pool>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PrivatecaCaPoolIamUpdater{
		project:  values["project"],
		location: values["location"],
		caPool:   values["ca_pool"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("ca_pool", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting ca_pool: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *PrivatecaCaPoolIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyCaPoolUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	url, err = AddQueryParams(url, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", IamPolicyVersion)})
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *PrivatecaCaPoolIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyCaPoolUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *PrivatecaCaPoolIamUpdater) qualifyCaPoolUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{PrivatecaBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/caPools/%s", u.project, u.location, u.caPool), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *PrivatecaCaPoolIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/caPools/%s", u.project, u.location, u.caPool)
}

func (u *PrivatecaCaPoolIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-privateca-capool-%s", u.GetResourceId())
}

func (u *PrivatecaCaPoolIamUpdater) DescribeResource() string {
	return fmt.Sprintf("privateca capool %q", u.GetResourceId())
}
