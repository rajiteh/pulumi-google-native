// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new UrlList in a given project and location.
type UrlList struct {
	pulumi.CustomResourceState

	// Time when the security policy was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Time when the security policy was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. Short name of the UrlList resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "url_list".
	UrlListId pulumi.StringOutput `pulumi:"urlListId"`
	// FQDNs and URLs.
	Values pulumi.StringArrayOutput `pulumi:"values"`
}

// NewUrlList registers a new resource with the given unique name, arguments, and options.
func NewUrlList(ctx *pulumi.Context,
	name string, args *UrlListArgs, opts ...pulumi.ResourceOption) (*UrlList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlListId == nil {
		return nil, errors.New("invalid value for required argument 'UrlListId'")
	}
	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"urlListId",
	})
	opts = append(opts, replaceOnChanges)
	var resource UrlList
	err := ctx.RegisterResource("google-native:networksecurity/v1beta1:UrlList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUrlList gets an existing UrlList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUrlList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UrlListState, opts ...pulumi.ResourceOption) (*UrlList, error) {
	var resource UrlList
	err := ctx.ReadResource("google-native:networksecurity/v1beta1:UrlList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UrlList resources.
type urlListState struct {
}

type UrlListState struct {
}

func (UrlListState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlListState)(nil)).Elem()
}

type urlListArgs struct {
	// Optional. Free-text description of the resource.
	Description *string `pulumi:"description"`
	Location    *string `pulumi:"location"`
	// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. Short name of the UrlList resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "url_list".
	UrlListId string `pulumi:"urlListId"`
	// FQDNs and URLs.
	Values []string `pulumi:"values"`
}

// The set of arguments for constructing a UrlList resource.
type UrlListArgs struct {
	// Optional. Free-text description of the resource.
	Description pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. Short name of the UrlList resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "url_list".
	UrlListId pulumi.StringInput
	// FQDNs and URLs.
	Values pulumi.StringArrayInput
}

func (UrlListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlListArgs)(nil)).Elem()
}

type UrlListInput interface {
	pulumi.Input

	ToUrlListOutput() UrlListOutput
	ToUrlListOutputWithContext(ctx context.Context) UrlListOutput
}

func (*UrlList) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlList)(nil)).Elem()
}

func (i *UrlList) ToUrlListOutput() UrlListOutput {
	return i.ToUrlListOutputWithContext(context.Background())
}

func (i *UrlList) ToUrlListOutputWithContext(ctx context.Context) UrlListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlListOutput)
}

type UrlListOutput struct{ *pulumi.OutputState }

func (UrlListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlList)(nil)).Elem()
}

func (o UrlListOutput) ToUrlListOutput() UrlListOutput {
	return o
}

func (o UrlListOutput) ToUrlListOutputWithContext(ctx context.Context) UrlListOutput {
	return o
}

// Time when the security policy was created.
func (o UrlListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o UrlListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o UrlListOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource provided by the user. Name is of the form projects/{project}/locations/{location}/urlLists/{url_list} url_list should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o UrlListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UrlListOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Time when the security policy was updated.
func (o UrlListOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Required. Short name of the UrlList resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "url_list".
func (o UrlListOutput) UrlListId() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringOutput { return v.UrlListId }).(pulumi.StringOutput)
}

// FQDNs and URLs.
func (o UrlListOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UrlList) pulumi.StringArrayOutput { return v.Values }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UrlListInput)(nil)).Elem(), &UrlList{})
	pulumi.RegisterOutputType(UrlListOutput{})
}
