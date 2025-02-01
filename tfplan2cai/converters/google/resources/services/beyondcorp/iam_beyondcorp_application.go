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

package beyondcorp

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var BeyondcorpApplicationIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"security_gateways_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"application_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type BeyondcorpApplicationIamUpdater struct {
	project            string
	securityGatewaysId string
	applicationId      string
	d                  tpgresource.TerraformResourceData
	Config             *transport_tpg.Config
}

func BeyondcorpApplicationIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("security_gateways_id"); ok {
		values["security_gateways_id"] = v.(string)
	}

	if v, ok := d.GetOk("application_id"); ok {
		values["application_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/global/securityGateways/(?P<security_gateways_id>[^/]+)/applications/(?P<application_id>[^/]+)", "(?P<project>[^/]+)/(?P<security_gateways_id>[^/]+)/(?P<application_id>[^/]+)", "(?P<security_gateways_id>[^/]+)/(?P<application_id>[^/]+)", "(?P<application_id>[^/]+)"}, d, config, d.Get("application_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BeyondcorpApplicationIamUpdater{
		project:            values["project"],
		securityGatewaysId: values["security_gateways_id"],
		applicationId:      values["application_id"],
		d:                  d,
		Config:             config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("security_gateways_id", u.securityGatewaysId); err != nil {
		return nil, fmt.Errorf("Error setting security_gateways_id: %s", err)
	}
	if err := d.Set("application_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting application_id: %s", err)
	}

	return u, nil
}

func BeyondcorpApplicationIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/global/securityGateways/(?P<security_gateways_id>[^/]+)/applications/(?P<application_id>[^/]+)", "(?P<project>[^/]+)/(?P<security_gateways_id>[^/]+)/(?P<application_id>[^/]+)", "(?P<security_gateways_id>[^/]+)/(?P<application_id>[^/]+)", "(?P<application_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BeyondcorpApplicationIamUpdater{
		project:            values["project"],
		securityGatewaysId: values["security_gateways_id"],
		applicationId:      values["application_id"],
		d:                  d,
		Config:             config,
	}
	if err := d.Set("application_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting application_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BeyondcorpApplicationIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyApplicationUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", tpgiamresource.IamPolicyVersion)})
	if err != nil {
		return nil, err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
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

func (u *BeyondcorpApplicationIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyApplicationUrl("setIamPolicy")
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

func (u *BeyondcorpApplicationIamUpdater) qualifyApplicationUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{BeyondcorpBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", u.project, u.securityGatewaysId, u.applicationId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *BeyondcorpApplicationIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/global/securityGateways/%s/applications/%s", u.project, u.securityGatewaysId, u.applicationId)
}

func (u *BeyondcorpApplicationIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-beyondcorp-application-%s", u.GetResourceId())
}

func (u *BeyondcorpApplicationIamUpdater) DescribeResource() string {
	return fmt.Sprintf("beyondcorp application %q", u.GetResourceId())
}
