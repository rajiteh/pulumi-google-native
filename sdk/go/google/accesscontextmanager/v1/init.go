// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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
	case "google-native:accesscontextmanager/v1:AccessLevel":
		r = &AccessLevel{}
	case "google-native:accesscontextmanager/v1:AccessPolicy":
		r = &AccessPolicy{}
	case "google-native:accesscontextmanager/v1:AccessPolicyIamBinding":
		r = &AccessPolicyIamBinding{}
	case "google-native:accesscontextmanager/v1:AccessPolicyIamMember":
		r = &AccessPolicyIamMember{}
	case "google-native:accesscontextmanager/v1:AccessPolicyIamPolicy":
		r = &AccessPolicyIamPolicy{}
	case "google-native:accesscontextmanager/v1:AuthorizedOrgsDesc":
		r = &AuthorizedOrgsDesc{}
	case "google-native:accesscontextmanager/v1:GcpUserAccessBinding":
		r = &GcpUserAccessBinding{}
	case "google-native:accesscontextmanager/v1:ServicePerimeter":
		r = &ServicePerimeter{}
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
		"accesscontextmanager/v1",
		&module{version},
	)
}
