// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a TagValue as a child of the specified TagKey. If a another request with the same parameters is sent while the original request is in process the second request will receive an error. A maximum of 1000 TagValues can exist under a TagKey at any given time.
type TagValue struct {
	pulumi.CustomResourceState

	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Immutable. Resource name for TagValue in the format `tagValues/456`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespaced name of the TagValue. Can be in the form `{organization_id}/{tag_key_short_name}/{tag_value_short_name}` or `{project_id}/{tag_key_short_name}/{tag_value_short_name}` or `{project_number}/{tag_key_short_name}/{tag_value_short_name}`.
	NamespacedName pulumi.StringOutput `pulumi:"namespacedName"`
	// Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTagValue registers a new resource with the given unique name, arguments, and options.
func NewTagValue(ctx *pulumi.Context,
	name string, args *TagValueArgs, opts ...pulumi.ResourceOption) (*TagValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ShortName == nil {
		return nil, errors.New("invalid value for required argument 'ShortName'")
	}
	var resource TagValue
	err := ctx.RegisterResource("google-native:cloudresourcemanager/v3:TagValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagValue gets an existing TagValue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagValueState, opts ...pulumi.ResourceOption) (*TagValue, error) {
	var resource TagValue
	err := ctx.ReadResource("google-native:cloudresourcemanager/v3:TagValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagValue resources.
type tagValueState struct {
}

type TagValueState struct {
}

func (TagValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueState)(nil)).Elem()
}

type tagValueArgs struct {
	// Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
	Description *string `pulumi:"description"`
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
	Etag *string `pulumi:"etag"`
	// Immutable. Resource name for TagValue in the format `tagValues/456`.
	Name *string `pulumi:"name"`
	// Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
	Parent *string `pulumi:"parent"`
	// Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName string `pulumi:"shortName"`
}

// The set of arguments for constructing a TagValue resource.
type TagValueArgs struct {
	// Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
	Description pulumi.StringPtrInput
	// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
	Etag pulumi.StringPtrInput
	// Immutable. Resource name for TagValue in the format `tagValues/456`.
	Name pulumi.StringPtrInput
	// Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
	Parent pulumi.StringPtrInput
	// Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName pulumi.StringInput
}

func (TagValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueArgs)(nil)).Elem()
}

type TagValueInput interface {
	pulumi.Input

	ToTagValueOutput() TagValueOutput
	ToTagValueOutputWithContext(ctx context.Context) TagValueOutput
}

func (*TagValue) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValue)(nil)).Elem()
}

func (i *TagValue) ToTagValueOutput() TagValueOutput {
	return i.ToTagValueOutputWithContext(context.Background())
}

func (i *TagValue) ToTagValueOutputWithContext(ctx context.Context) TagValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueOutput)
}

type TagValueOutput struct{ *pulumi.OutputState }

func (TagValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValue)(nil)).Elem()
}

func (o TagValueOutput) ToTagValueOutput() TagValueOutput {
	return o
}

func (o TagValueOutput) ToTagValueOutputWithContext(ctx context.Context) TagValueOutput {
	return o
}

// Creation time.
func (o TagValueOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
func (o TagValueOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
func (o TagValueOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Immutable. Resource name for TagValue in the format `tagValues/456`.
func (o TagValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespaced name of the TagValue. Can be in the form `{organization_id}/{tag_key_short_name}/{tag_value_short_name}` or `{project_id}/{tag_key_short_name}/{tag_value_short_name}` or `{project_number}/{tag_key_short_name}/{tag_value_short_name}`.
func (o TagValueOutput) NamespacedName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.NamespacedName }).(pulumi.StringOutput)
}

// Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
func (o TagValueOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o TagValueOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

// Update time.
func (o TagValueOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TagValue) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagValueInput)(nil)).Elem(), &TagValue{})
	pulumi.RegisterOutputType(TagValueOutput{})
}
