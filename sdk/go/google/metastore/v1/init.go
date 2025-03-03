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
	case "google-native:metastore/v1:Backup":
		r = &Backup{}
	case "google-native:metastore/v1:Federation":
		r = &Federation{}
	case "google-native:metastore/v1:FederationIamBinding":
		r = &FederationIamBinding{}
	case "google-native:metastore/v1:FederationIamMember":
		r = &FederationIamMember{}
	case "google-native:metastore/v1:FederationIamPolicy":
		r = &FederationIamPolicy{}
	case "google-native:metastore/v1:MetadataImport":
		r = &MetadataImport{}
	case "google-native:metastore/v1:Service":
		r = &Service{}
	case "google-native:metastore/v1:ServiceBackupIamBinding":
		r = &ServiceBackupIamBinding{}
	case "google-native:metastore/v1:ServiceBackupIamMember":
		r = &ServiceBackupIamMember{}
	case "google-native:metastore/v1:ServiceBackupIamPolicy":
		r = &ServiceBackupIamPolicy{}
	case "google-native:metastore/v1:ServiceIamBinding":
		r = &ServiceIamBinding{}
	case "google-native:metastore/v1:ServiceIamMember":
		r = &ServiceIamMember{}
	case "google-native:metastore/v1:ServiceIamPolicy":
		r = &ServiceIamPolicy{}
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
		"metastore/v1",
		&module{version},
	)
}
