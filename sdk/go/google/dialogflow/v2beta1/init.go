// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "google-native:dialogflow/v2beta1:Context":
		r = &Context{}
	case "google-native:dialogflow/v2beta1:Conversation":
		r = &Conversation{}
	case "google-native:dialogflow/v2beta1:ConversationProfile":
		r = &ConversationProfile{}
	case "google-native:dialogflow/v2beta1:Document":
		r = &Document{}
	case "google-native:dialogflow/v2beta1:EntityType":
		r = &EntityType{}
	case "google-native:dialogflow/v2beta1:Environment":
		r = &Environment{}
	case "google-native:dialogflow/v2beta1:Intent":
		r = &Intent{}
	case "google-native:dialogflow/v2beta1:KnowledgeBase":
		r = &KnowledgeBase{}
	case "google-native:dialogflow/v2beta1:Participant":
		r = &Participant{}
	case "google-native:dialogflow/v2beta1:SessionEntityType":
		r = &SessionEntityType{}
	case "google-native:dialogflow/v2beta1:Version":
		r = &Version{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"dialogflow/v2beta1",
		&module{version},
	)
}
