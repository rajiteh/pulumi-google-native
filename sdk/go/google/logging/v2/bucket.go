// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a log bucket that can be used to store log entries. After a bucket has been created, the bucket's location cannot be changed.
// Auto-naming is currently not supported for this resource.
type Bucket struct {
	pulumi.CustomResourceState

	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled pulumi.BoolOutput `pulumi:"analyticsEnabled"`
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings CmekSettingsResponseOutput `pulumi:"cmekSettings"`
	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of indexed fields and related configuration data.
	IndexConfigs IndexConfigResponseArrayOutput `pulumi:"indexConfigs"`
	// The bucket lifecycle state.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	Location       pulumi.StringOutput `pulumi:"location"`
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields pulumi.StringArrayOutput `pulumi:"restrictedFields"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The last update timestamp of the bucket.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucketId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Bucket
	err := ctx.RegisterResource("google-native:logging/v2:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("google-native:logging/v2:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
}

type BucketState struct {
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled *bool `pulumi:"analyticsEnabled"`
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *CmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// A list of indexed fields and related configuration data.
	IndexConfigs []IndexConfig `pulumi:"indexConfigs"`
	Location     *string       `pulumi:"location"`
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked  *bool   `pulumi:"locked"`
	Project *string `pulumi:"project"`
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields []string `pulumi:"restrictedFields"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
	AnalyticsEnabled pulumi.BoolPtrInput
	// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
	BucketId pulumi.StringInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
	CmekSettings CmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// A list of indexed fields and related configuration data.
	IndexConfigs IndexConfigArrayInput
	Location     pulumi.StringPtrInput
	// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked  pulumi.BoolPtrInput
	Project pulumi.StringPtrInput
	// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
	RestrictedFields pulumi.StringArrayInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// Whether log analytics is enabled for this bucket.Once enabled, log analytics features cannot be disabled.
func (o BucketOutput) AnalyticsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolOutput { return v.AnalyticsEnabled }).(pulumi.BoolOutput)
}

// Required. A client-assigned identifier such as "my-bucket". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
func (o BucketOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed.
func (o BucketOutput) CmekSettings() CmekSettingsResponseOutput {
	return o.ApplyT(func(v *Bucket) CmekSettingsResponseOutput { return v.CmekSettings }).(CmekSettingsResponseOutput)
}

// The creation timestamp of the bucket. This is not set for any of the default buckets.
func (o BucketOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this bucket.
func (o BucketOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A list of indexed fields and related configuration data.
func (o BucketOutput) IndexConfigs() IndexConfigResponseArrayOutput {
	return o.ApplyT(func(v *Bucket) IndexConfigResponseArrayOutput { return v.IndexConfigs }).(IndexConfigResponseArrayOutput)
}

// The bucket lifecycle state.
func (o BucketOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

func (o BucketOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Whether the bucket is locked.The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
func (o BucketOutput) Locked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolOutput { return v.Locked }).(pulumi.BoolOutput)
}

// The resource name of the bucket.For example:projects/my-project/locations/global/buckets/my-bucketFor a list of supported locations, see Supported Regions (https://cloud.google.com/logging/docs/region-support)For the location of global it is unspecified where log entries are actually stored.After a bucket has been created, the location cannot be changed.
func (o BucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BucketOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Log entry field paths that are denied access in this bucket.The following fields and their children are eligible: textPayload, jsonPayload, protoPayload, httpRequest, labels, sourceLocation.Restricting a repeated field will restrict all values. Adding a parent will block all child fields. (e.g. foo.bar will block foo.bar.baz)
func (o BucketOutput) RestrictedFields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringArrayOutput { return v.RestrictedFields }).(pulumi.StringArrayOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
func (o BucketOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Bucket) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

// The last update timestamp of the bucket.
func (o BucketOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterOutputType(BucketOutput{})
}
