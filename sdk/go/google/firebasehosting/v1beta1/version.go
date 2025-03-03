// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new version for the specified site.
type Version struct {
	pulumi.CustomResourceState

	// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
	Config ServingConfigResponseOutput `pulumi:"config"`
	// The time at which the version was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Identifies the user who created the version.
	CreateUser ActingUserResponseOutput `pulumi:"createUser"`
	// The time at which the version was `DELETED`.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// Identifies the user who `DELETED` the version.
	DeleteUser ActingUserResponseOutput `pulumi:"deleteUser"`
	// The total number of files associated with the version. This value is calculated after a version is `FINALIZED`.
	FileCount pulumi.StringOutput `pulumi:"fileCount"`
	// The time at which the version was `FINALIZED`.
	FinalizeTime pulumi.StringOutput `pulumi:"finalizeTime"`
	// Identifies the user who `FINALIZED` the version.
	FinalizeUser ActingUserResponseOutput `pulumi:"finalizeUser"`
	// The labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	SiteId  pulumi.StringOutput `pulumi:"siteId"`
	// The self-reported size of the version. This value is used for a pre-emptive quota check for legacy version uploads.
	SizeBytes pulumi.StringPtrOutput `pulumi:"sizeBytes"`
	// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
	Status pulumi.StringOutput `pulumi:"status"`
	// The total stored bytesize of the version. This value is calculated after a version is `FINALIZED`.
	VersionBytes pulumi.StringOutput `pulumi:"versionBytes"`
	// A unique id for the new version. This is was only specified for legacy version creations, and should be blank.
	VersionId pulumi.StringPtrOutput `pulumi:"versionId"`
}

// NewVersion registers a new resource with the given unique name, arguments, and options.
func NewVersion(ctx *pulumi.Context,
	name string, args *VersionArgs, opts ...pulumi.ResourceOption) (*Version, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"siteId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Version
	err := ctx.RegisterResource("google-native:firebasehosting/v1beta1:Version", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVersion gets an existing Version resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VersionState, opts ...pulumi.ResourceOption) (*Version, error) {
	var resource Version
	err := ctx.ReadResource("google-native:firebasehosting/v1beta1:Version", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Version resources.
type versionState struct {
}

type VersionState struct {
}

func (VersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*versionState)(nil)).Elem()
}

type versionArgs struct {
	// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
	Config *ServingConfig `pulumi:"config"`
	// The labels used for extra metadata and/or filtering.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	SiteId  string  `pulumi:"siteId"`
	// The self-reported size of the version. This value is used for a pre-emptive quota check for legacy version uploads.
	SizeBytes *string `pulumi:"sizeBytes"`
	// A unique id for the new version. This is was only specified for legacy version creations, and should be blank.
	VersionId *string `pulumi:"versionId"`
}

// The set of arguments for constructing a Version resource.
type VersionArgs struct {
	// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
	Config ServingConfigPtrInput
	// The labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	SiteId  pulumi.StringInput
	// The self-reported size of the version. This value is used for a pre-emptive quota check for legacy version uploads.
	SizeBytes pulumi.StringPtrInput
	// A unique id for the new version. This is was only specified for legacy version creations, and should be blank.
	VersionId pulumi.StringPtrInput
}

func (VersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*versionArgs)(nil)).Elem()
}

type VersionInput interface {
	pulumi.Input

	ToVersionOutput() VersionOutput
	ToVersionOutputWithContext(ctx context.Context) VersionOutput
}

func (*Version) ElementType() reflect.Type {
	return reflect.TypeOf((**Version)(nil)).Elem()
}

func (i *Version) ToVersionOutput() VersionOutput {
	return i.ToVersionOutputWithContext(context.Background())
}

func (i *Version) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionOutput)
}

type VersionOutput struct{ *pulumi.OutputState }

func (VersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Version)(nil)).Elem()
}

func (o VersionOutput) ToVersionOutput() VersionOutput {
	return o
}

func (o VersionOutput) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return o
}

// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
func (o VersionOutput) Config() ServingConfigResponseOutput {
	return o.ApplyT(func(v *Version) ServingConfigResponseOutput { return v.Config }).(ServingConfigResponseOutput)
}

// The time at which the version was created.
func (o VersionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Identifies the user who created the version.
func (o VersionOutput) CreateUser() ActingUserResponseOutput {
	return o.ApplyT(func(v *Version) ActingUserResponseOutput { return v.CreateUser }).(ActingUserResponseOutput)
}

// The time at which the version was `DELETED`.
func (o VersionOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// Identifies the user who `DELETED` the version.
func (o VersionOutput) DeleteUser() ActingUserResponseOutput {
	return o.ApplyT(func(v *Version) ActingUserResponseOutput { return v.DeleteUser }).(ActingUserResponseOutput)
}

// The total number of files associated with the version. This value is calculated after a version is `FINALIZED`.
func (o VersionOutput) FileCount() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.FileCount }).(pulumi.StringOutput)
}

// The time at which the version was `FINALIZED`.
func (o VersionOutput) FinalizeTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.FinalizeTime }).(pulumi.StringOutput)
}

// Identifies the user who `FINALIZED` the version.
func (o VersionOutput) FinalizeUser() ActingUserResponseOutput {
	return o.ApplyT(func(v *Version) ActingUserResponseOutput { return v.FinalizeUser }).(ActingUserResponseOutput)
}

// The labels used for extra metadata and/or filtering.
func (o VersionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Version) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
func (o VersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VersionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o VersionOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.SiteId }).(pulumi.StringOutput)
}

// The self-reported size of the version. This value is used for a pre-emptive quota check for legacy version uploads.
func (o VersionOutput) SizeBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Version) pulumi.StringPtrOutput { return v.SizeBytes }).(pulumi.StringPtrOutput)
}

// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
func (o VersionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The total stored bytesize of the version. This value is calculated after a version is `FINALIZED`.
func (o VersionOutput) VersionBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *Version) pulumi.StringOutput { return v.VersionBytes }).(pulumi.StringOutput)
}

// A unique id for the new version. This is was only specified for legacy version creations, and should be blank.
func (o VersionOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Version) pulumi.StringPtrOutput { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VersionInput)(nil)).Elem(), &Version{})
	pulumi.RegisterOutputType(VersionOutput{})
}
