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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

var PubsubTopicIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"topic": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type PubsubTopicIamUpdater struct {
	project string
	topic   string
	d       tpgresource.TerraformResourceData
	Config  *transport_tpg.Config
}

func PubsubTopicIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("topic"); ok {
		values["topic"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)", "(?P<project>[^/]+)/(?P<topic>[^/]+)", "(?P<topic>[^/]+)"}, d, config, d.Get("topic").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic:   values["topic"],
		d:       d,
		Config:  config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("topic", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting topic: %s", err)
	}

	return u, nil
}

func PubsubTopicIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/topics/(?P<topic>[^/]+)", "(?P<project>[^/]+)/(?P<topic>[^/]+)", "(?P<topic>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &PubsubTopicIamUpdater{
		project: values["project"],
		topic:   values["topic"],
		d:       d,
		Config:  config,
	}
	if err := d.Set("topic", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting topic: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *PubsubTopicIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTopicUrl("getIamPolicy")
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
		Config:               u.Config,
		Method:               "GET",
		Project:              project,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.PubsubTopicProjectNotReady},
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

func (u *PubsubTopicIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTopicUrl("setIamPolicy")
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
		Config:               u.Config,
		Method:               "POST",
		Project:              project,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              u.d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.PubsubTopicProjectNotReady},
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *PubsubTopicIamUpdater) qualifyTopicUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{PubsubBasePath}}%s:%s", fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *PubsubTopicIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/topics/%s", u.project, u.topic)
}

func (u *PubsubTopicIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-pubsub-topic-%s", u.GetResourceId())
}

func (u *PubsubTopicIamUpdater) DescribeResource() string {
	return fmt.Sprintf("pubsub topic %q", u.GetResourceId())
}
