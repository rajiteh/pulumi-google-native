// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a link.
func LookupLink(ctx *pulumi.Context, args *LookupLinkArgs, opts ...pulumi.InvokeOption) (*LookupLinkResult, error) {
	var rv LookupLinkResult
	err := ctx.Invoke("google-native:logging/v2:getLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkArgs struct {
	BucketId string  `pulumi:"bucketId"`
	LinkId   string  `pulumi:"linkId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupLinkResult struct {
	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset BigQueryDatasetResponse `pulumi:"bigqueryDataset"`
	// The creation timestamp of the link.
	CreateTime string `pulumi:"createTime"`
	// Describes this link.The maximum length of the description is 8000 characters.
	Description string `pulumi:"description"`
	// The resource lifecycle state.
	LifecycleState string `pulumi:"lifecycleState"`
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name string `pulumi:"name"`
}

func LookupLinkOutput(ctx *pulumi.Context, args LookupLinkOutputArgs, opts ...pulumi.InvokeOption) LookupLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkResult, error) {
			args := v.(LookupLinkArgs)
			r, err := LookupLink(ctx, &args, opts...)
			var s LookupLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkResultOutput)
}

type LookupLinkOutputArgs struct {
	BucketId pulumi.StringInput    `pulumi:"bucketId"`
	LinkId   pulumi.StringInput    `pulumi:"linkId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkArgs)(nil)).Elem()
}

type LookupLinkResultOutput struct{ *pulumi.OutputState }

func (LookupLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkResult)(nil)).Elem()
}

func (o LookupLinkResultOutput) ToLookupLinkResultOutput() LookupLinkResultOutput {
	return o
}

func (o LookupLinkResultOutput) ToLookupLinkResultOutputWithContext(ctx context.Context) LookupLinkResultOutput {
	return o
}

// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
func (o LookupLinkResultOutput) BigqueryDataset() BigQueryDatasetResponseOutput {
	return o.ApplyT(func(v LookupLinkResult) BigQueryDatasetResponse { return v.BigqueryDataset }).(BigQueryDatasetResponseOutput)
}

// The creation timestamp of the link.
func (o LookupLinkResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this link.The maximum length of the description is 8000 characters.
func (o LookupLinkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.Description }).(pulumi.StringOutput)
}

// The resource lifecycle state.
func (o LookupLinkResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
func (o LookupLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkResultOutput{})
}
