// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a view over log entries in a log bucket. A bucket may contain a maximum of 30 views.
// Auto-naming is currently not supported for this resource.
type BillingAccountBucketView struct {
	pulumi.CustomResourceState

	BillingAccountId pulumi.StringOutput `pulumi:"billingAccountId"`
	BucketId         pulumi.StringOutput `pulumi:"bucketId"`
	// The creation timestamp of the view.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this view.
	Description pulumi.StringOutput `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter   pulumi.StringOutput `pulumi:"filter"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update timestamp of the view.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	ViewId pulumi.StringOutput `pulumi:"viewId"`
}

// NewBillingAccountBucketView registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountBucketView(ctx *pulumi.Context,
	name string, args *BillingAccountBucketViewArgs, opts ...pulumi.ResourceOption) (*BillingAccountBucketView, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"billingAccountId",
		"bucketId",
		"location",
		"viewId",
	})
	opts = append(opts, replaceOnChanges)
	var resource BillingAccountBucketView
	err := ctx.RegisterResource("google-native:logging/v2:BillingAccountBucketView", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountBucketView gets an existing BillingAccountBucketView resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountBucketView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountBucketViewState, opts ...pulumi.ResourceOption) (*BillingAccountBucketView, error) {
	var resource BillingAccountBucketView
	err := ctx.ReadResource("google-native:logging/v2:BillingAccountBucketView", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountBucketView resources.
type billingAccountBucketViewState struct {
}

type BillingAccountBucketViewState struct {
}

func (BillingAccountBucketViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketViewState)(nil)).Elem()
}

type billingAccountBucketViewArgs struct {
	BillingAccountId string `pulumi:"billingAccountId"`
	BucketId         string `pulumi:"bucketId"`
	// Describes this view.
	Description *string `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter   *string `pulumi:"filter"`
	Location *string `pulumi:"location"`
	// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
	Name *string `pulumi:"name"`
	// Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	ViewId string `pulumi:"viewId"`
}

// The set of arguments for constructing a BillingAccountBucketView resource.
type BillingAccountBucketViewArgs struct {
	BillingAccountId pulumi.StringInput
	BucketId         pulumi.StringInput
	// Describes this view.
	Description pulumi.StringPtrInput
	// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter   pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
	Name pulumi.StringPtrInput
	// Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	ViewId pulumi.StringInput
}

func (BillingAccountBucketViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketViewArgs)(nil)).Elem()
}

type BillingAccountBucketViewInput interface {
	pulumi.Input

	ToBillingAccountBucketViewOutput() BillingAccountBucketViewOutput
	ToBillingAccountBucketViewOutputWithContext(ctx context.Context) BillingAccountBucketViewOutput
}

func (*BillingAccountBucketView) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketView)(nil)).Elem()
}

func (i *BillingAccountBucketView) ToBillingAccountBucketViewOutput() BillingAccountBucketViewOutput {
	return i.ToBillingAccountBucketViewOutputWithContext(context.Background())
}

func (i *BillingAccountBucketView) ToBillingAccountBucketViewOutputWithContext(ctx context.Context) BillingAccountBucketViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBucketViewOutput)
}

type BillingAccountBucketViewOutput struct{ *pulumi.OutputState }

func (BillingAccountBucketViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketView)(nil)).Elem()
}

func (o BillingAccountBucketViewOutput) ToBillingAccountBucketViewOutput() BillingAccountBucketViewOutput {
	return o
}

func (o BillingAccountBucketViewOutput) ToBillingAccountBucketViewOutputWithContext(ctx context.Context) BillingAccountBucketViewOutput {
	return o
}

func (o BillingAccountBucketViewOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

func (o BillingAccountBucketViewOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The creation timestamp of the view.
func (o BillingAccountBucketViewOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this view.
func (o BillingAccountBucketViewOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
func (o BillingAccountBucketViewOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o BillingAccountBucketViewOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
func (o BillingAccountBucketViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The last update timestamp of the view.
func (o BillingAccountBucketViewOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
func (o BillingAccountBucketViewOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountBucketView) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountBucketViewInput)(nil)).Elem(), &BillingAccountBucketView{})
	pulumi.RegisterOutputType(BillingAccountBucketViewOutput{})
}
