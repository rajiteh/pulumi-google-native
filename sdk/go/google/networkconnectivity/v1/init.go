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
	case "google-native:networkconnectivity/v1:Hub":
		r = &Hub{}
	case "google-native:networkconnectivity/v1:HubIamBinding":
		r = &HubIamBinding{}
	case "google-native:networkconnectivity/v1:HubIamMember":
		r = &HubIamMember{}
	case "google-native:networkconnectivity/v1:HubIamPolicy":
		r = &HubIamPolicy{}
	case "google-native:networkconnectivity/v1:InternalRange":
		r = &InternalRange{}
	case "google-native:networkconnectivity/v1:PolicyBasedRouteIamBinding":
		r = &PolicyBasedRouteIamBinding{}
	case "google-native:networkconnectivity/v1:PolicyBasedRouteIamMember":
		r = &PolicyBasedRouteIamMember{}
	case "google-native:networkconnectivity/v1:PolicyBasedRouteIamPolicy":
		r = &PolicyBasedRouteIamPolicy{}
	case "google-native:networkconnectivity/v1:ServiceClass":
		r = &ServiceClass{}
	case "google-native:networkconnectivity/v1:ServiceClassIamBinding":
		r = &ServiceClassIamBinding{}
	case "google-native:networkconnectivity/v1:ServiceClassIamMember":
		r = &ServiceClassIamMember{}
	case "google-native:networkconnectivity/v1:ServiceClassIamPolicy":
		r = &ServiceClassIamPolicy{}
	case "google-native:networkconnectivity/v1:ServiceConnectionMap":
		r = &ServiceConnectionMap{}
	case "google-native:networkconnectivity/v1:ServiceConnectionMapIamBinding":
		r = &ServiceConnectionMapIamBinding{}
	case "google-native:networkconnectivity/v1:ServiceConnectionMapIamMember":
		r = &ServiceConnectionMapIamMember{}
	case "google-native:networkconnectivity/v1:ServiceConnectionMapIamPolicy":
		r = &ServiceConnectionMapIamPolicy{}
	case "google-native:networkconnectivity/v1:ServiceConnectionPolicy":
		r = &ServiceConnectionPolicy{}
	case "google-native:networkconnectivity/v1:ServiceConnectionPolicyIamBinding":
		r = &ServiceConnectionPolicyIamBinding{}
	case "google-native:networkconnectivity/v1:ServiceConnectionPolicyIamMember":
		r = &ServiceConnectionPolicyIamMember{}
	case "google-native:networkconnectivity/v1:ServiceConnectionPolicyIamPolicy":
		r = &ServiceConnectionPolicyIamPolicy{}
	case "google-native:networkconnectivity/v1:ServiceConnectionToken":
		r = &ServiceConnectionToken{}
	case "google-native:networkconnectivity/v1:Spoke":
		r = &Spoke{}
	case "google-native:networkconnectivity/v1:SpokeIamBinding":
		r = &SpokeIamBinding{}
	case "google-native:networkconnectivity/v1:SpokeIamMember":
		r = &SpokeIamMember{}
	case "google-native:networkconnectivity/v1:SpokeIamPolicy":
		r = &SpokeIamPolicy{}
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
		"networkconnectivity/v1",
		&module{version},
	)
}
