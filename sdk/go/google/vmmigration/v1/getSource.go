// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Source.
func LookupSource(ctx *pulumi.Context, args *LookupSourceArgs, opts ...pulumi.InvokeOption) (*LookupSourceResult, error) {
	var rv LookupSourceResult
	err := ctx.Invoke("google-native:vmmigration/v1:getSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceArgs struct {
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	SourceId string  `pulumi:"sourceId"`
}

type LookupSourceResult struct {
	// AWS type source details.
	Aws AwsSourceDetailsResponse `pulumi:"aws"`
	// The create time timestamp.
	CreateTime string `pulumi:"createTime"`
	// User-provided description of the source.
	Description string `pulumi:"description"`
	// The labels of the source.
	Labels map[string]string `pulumi:"labels"`
	// The Source name.
	Name string `pulumi:"name"`
	// The update time timestamp.
	UpdateTime string `pulumi:"updateTime"`
	// Vmware type source details.
	Vmware VmwareSourceDetailsResponse `pulumi:"vmware"`
}

func LookupSourceOutput(ctx *pulumi.Context, args LookupSourceOutputArgs, opts ...pulumi.InvokeOption) LookupSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceResult, error) {
			args := v.(LookupSourceArgs)
			r, err := LookupSource(ctx, &args, opts...)
			var s LookupSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceResultOutput)
}

type LookupSourceOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
	SourceId pulumi.StringInput    `pulumi:"sourceId"`
}

func (LookupSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceArgs)(nil)).Elem()
}

type LookupSourceResultOutput struct{ *pulumi.OutputState }

func (LookupSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceResult)(nil)).Elem()
}

func (o LookupSourceResultOutput) ToLookupSourceResultOutput() LookupSourceResultOutput {
	return o
}

func (o LookupSourceResultOutput) ToLookupSourceResultOutputWithContext(ctx context.Context) LookupSourceResultOutput {
	return o
}

// AWS type source details.
func (o LookupSourceResultOutput) Aws() AwsSourceDetailsResponseOutput {
	return o.ApplyT(func(v LookupSourceResult) AwsSourceDetailsResponse { return v.Aws }).(AwsSourceDetailsResponseOutput)
}

// The create time timestamp.
func (o LookupSourceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// User-provided description of the source.
func (o LookupSourceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The labels of the source.
func (o LookupSourceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSourceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The Source name.
func (o LookupSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The update time timestamp.
func (o LookupSourceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Vmware type source details.
func (o LookupSourceResultOutput) Vmware() VmwareSourceDetailsResponseOutput {
	return o.ApplyT(func(v LookupSourceResult) VmwareSourceDetailsResponse { return v.Vmware }).(VmwareSourceDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceResultOutput{})
}
