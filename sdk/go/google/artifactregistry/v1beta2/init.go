// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

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
	case "google-native:artifactregistry/v1beta2:Repository":
		r = &Repository{}
	case "google-native:artifactregistry/v1beta2:RepositoryIamBinding":
		r = &RepositoryIamBinding{}
	case "google-native:artifactregistry/v1beta2:RepositoryIamMember":
		r = &RepositoryIamMember{}
	case "google-native:artifactregistry/v1beta2:RepositoryIamPolicy":
		r = &RepositoryIamPolicy{}
	case "google-native:artifactregistry/v1beta2:Tag":
		r = &Tag{}
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
		"artifactregistry/v1beta2",
		&module{version},
	)
}
