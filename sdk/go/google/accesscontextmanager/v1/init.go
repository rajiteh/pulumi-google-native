// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-cloud/sdk/go/google-cloud"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "google-cloud:accesscontextmanager/v1:AccessLevel":
		r, err = NewAccessLevel(ctx, name, nil, pulumi.URN_(urn))
	case "google-cloud:accesscontextmanager/v1:AccessPolicy":
		r, err = NewAccessPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "google-cloud:accesscontextmanager/v1:GcpUserAccessBinding":
		r, err = NewGcpUserAccessBinding(ctx, name, nil, pulumi.URN_(urn))
	case "google-cloud:accesscontextmanager/v1:ServicePerimeter":
		r, err = NewServicePerimeter(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := google - cloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"google-cloud",
		"accesscontextmanager/v1",
		&module{version},
	)
}