// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/pubsublite/Subscription.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package pubsublite

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const PubsubLiteSubscriptionAssetType string = "pubsublite.googleapis.com/Subscription"

func ResourceConverterPubsubLiteSubscription() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: PubsubLiteSubscriptionAssetType,
		Convert:   GetPubsubLiteSubscriptionCaiObject,
	}
}

func GetPubsubLiteSubscriptionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//pubsublite.googleapis.com/projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetPubsubLiteSubscriptionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: PubsubLiteSubscriptionAssetType,
			Resource: &cai.AssetResource{
				Version:              "admin",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsublite/admin/rest",
				DiscoveryName:        "Subscription",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetPubsubLiteSubscriptionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	topicProp, err := expandPubsubLiteSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	deliveryConfigProp, err := expandPubsubLiteSubscriptionDeliveryConfig(d.Get("delivery_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("delivery_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(deliveryConfigProp)) && (ok || !reflect.DeepEqual(v, deliveryConfigProp)) {
		obj["deliveryConfig"] = deliveryConfigProp
	}

	return resourcePubsubLiteSubscriptionEncoder(d, config, obj)
}

func resourcePubsubLiteSubscriptionEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return nil, err
	}

	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	// API Endpoint requires region in the URL. We infer it from the zone.

	region := tpgresource.GetRegionFromZone(zone)

	if region == "" {
		return nil, fmt.Errorf("invalid zone %q, unable to infer region from zone", zone)
	}

	return obj, nil
}

func expandPubsubLiteSubscriptionTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return "", err
	}

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return nil, err
	}

	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile(`projects\/(.*)\/locations\/(.*)\/topics\/(.*)`)
	match := re.FindStringSubmatch(topic)
	if len(match) == 4 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/locations/%s/topics/%s", project, zone, topic)
		if err := d.Set("topic", fullTopic); err != nil {
			return nil, fmt.Errorf("Error setting topic: %s", err)
		}
		return fullTopic, nil
	}
}

func expandPubsubLiteSubscriptionDeliveryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDeliveryRequirement, err := expandPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(original["delivery_requirement"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeliveryRequirement); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deliveryRequirement"] = transformedDeliveryRequirement
	}

	return transformed, nil
}

func expandPubsubLiteSubscriptionDeliveryConfigDeliveryRequirement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
