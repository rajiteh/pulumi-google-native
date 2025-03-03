// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a source.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Source struct {
	pulumi.CustomResourceState

	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description pulumi.StringOutput `pulumi:"description"`
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
}

// NewSource registers a new resource with the given unique name, arguments, and options.
func NewSource(ctx *pulumi.Context,
	name string, args *SourceArgs, opts ...pulumi.ResourceOption) (*Source, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Source
	err := ctx.RegisterResource("google-native:securitycenter/v1beta1:Source", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSource gets an existing Source resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceState, opts ...pulumi.ResourceOption) (*Source, error) {
	var resource Source
	err := ctx.ReadResource("google-native:securitycenter/v1beta1:Source", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Source resources.
type sourceState struct {
}

type SourceState struct {
}

func (SourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceState)(nil)).Elem()
}

type sourceArgs struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description *string `pulumi:"description"`
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName *string `pulumi:"displayName"`
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a Source resource.
type SourceArgs struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description pulumi.StringPtrInput
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName pulumi.StringPtrInput
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceArgs)(nil)).Elem()
}

type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(ctx context.Context) SourceOutput
}

func (*Source) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (i *Source) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i *Source) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
func (o SourceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
func (o SourceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
func (o SourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SourceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Source) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceInput)(nil)).Elem(), &Source{})
	pulumi.RegisterOutputType(SourceOutput{})
}
