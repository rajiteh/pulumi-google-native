// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a view on a log bucket..
func LookupOrganizationBucketView(ctx *pulumi.Context, args *LookupOrganizationBucketViewArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationBucketViewResult, error) {
	var rv LookupOrganizationBucketViewResult
	err := ctx.Invoke("google-native:logging/v2:getOrganizationBucketView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationBucketViewArgs struct {
	BucketId       string `pulumi:"bucketId"`
	Location       string `pulumi:"location"`
	OrganizationId string `pulumi:"organizationId"`
	ViewId         string `pulumi:"viewId"`
}

type LookupOrganizationBucketViewResult struct {
	// The creation timestamp of the view.
	CreateTime string `pulumi:"createTime"`
	// Describes this view.
	Description string `pulumi:"description"`
	// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
	Filter string `pulumi:"filter"`
	// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
	Name string `pulumi:"name"`
	// The last update timestamp of the view.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupOrganizationBucketViewOutput(ctx *pulumi.Context, args LookupOrganizationBucketViewOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationBucketViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationBucketViewResult, error) {
			args := v.(LookupOrganizationBucketViewArgs)
			r, err := LookupOrganizationBucketView(ctx, &args, opts...)
			var s LookupOrganizationBucketViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationBucketViewResultOutput)
}

type LookupOrganizationBucketViewOutputArgs struct {
	BucketId       pulumi.StringInput `pulumi:"bucketId"`
	Location       pulumi.StringInput `pulumi:"location"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
	ViewId         pulumi.StringInput `pulumi:"viewId"`
}

func (LookupOrganizationBucketViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationBucketViewArgs)(nil)).Elem()
}

type LookupOrganizationBucketViewResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationBucketViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationBucketViewResult)(nil)).Elem()
}

func (o LookupOrganizationBucketViewResultOutput) ToLookupOrganizationBucketViewResultOutput() LookupOrganizationBucketViewResultOutput {
	return o
}

func (o LookupOrganizationBucketViewResultOutput) ToLookupOrganizationBucketViewResultOutputWithContext(ctx context.Context) LookupOrganizationBucketViewResultOutput {
	return o
}

// The creation timestamp of the view.
func (o LookupOrganizationBucketViewResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBucketViewResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this view.
func (o LookupOrganizationBucketViewResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBucketViewResult) string { return v.Description }).(pulumi.StringOutput)
}

// Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
func (o LookupOrganizationBucketViewResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBucketViewResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
func (o LookupOrganizationBucketViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBucketViewResult) string { return v.Name }).(pulumi.StringOutput)
}

// The last update timestamp of the view.
func (o LookupOrganizationBucketViewResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationBucketViewResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationBucketViewResultOutput{})
}
