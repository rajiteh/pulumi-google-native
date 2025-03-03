// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a data collector.
func LookupDataCollector(ctx *pulumi.Context, args *LookupDataCollectorArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectorResult, error) {
	var rv LookupDataCollectorResult
	err := ctx.Invoke("google-native:apigee/v1:getDataCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectorArgs struct {
	DatacollectorId string `pulumi:"datacollectorId"`
	OrganizationId  string `pulumi:"organizationId"`
}

type LookupDataCollectorResult struct {
	// The time at which the data collector was created in milliseconds since the epoch.
	CreatedAt string `pulumi:"createdAt"`
	// A description of the data collector.
	Description string `pulumi:"description"`
	// The time at which the Data Collector was last updated in milliseconds since the epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// ID of the data collector. Must begin with `dc_`.
	Name string `pulumi:"name"`
	// Immutable. The type of data this data collector will collect.
	Type string `pulumi:"type"`
}

func LookupDataCollectorOutput(ctx *pulumi.Context, args LookupDataCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataCollectorResult, error) {
			args := v.(LookupDataCollectorArgs)
			r, err := LookupDataCollector(ctx, &args, opts...)
			var s LookupDataCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataCollectorResultOutput)
}

type LookupDataCollectorOutputArgs struct {
	DatacollectorId pulumi.StringInput `pulumi:"datacollectorId"`
	OrganizationId  pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupDataCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectorArgs)(nil)).Elem()
}

type LookupDataCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectorResult)(nil)).Elem()
}

func (o LookupDataCollectorResultOutput) ToLookupDataCollectorResultOutput() LookupDataCollectorResultOutput {
	return o
}

func (o LookupDataCollectorResultOutput) ToLookupDataCollectorResultOutputWithContext(ctx context.Context) LookupDataCollectorResultOutput {
	return o
}

// The time at which the data collector was created in milliseconds since the epoch.
func (o LookupDataCollectorResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectorResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description of the data collector.
func (o LookupDataCollectorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectorResult) string { return v.Description }).(pulumi.StringOutput)
}

// The time at which the Data Collector was last updated in milliseconds since the epoch.
func (o LookupDataCollectorResultOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectorResult) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// ID of the data collector. Must begin with `dc_`.
func (o LookupDataCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The type of data this data collector will collect.
func (o LookupDataCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectorResultOutput{})
}
