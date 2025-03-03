// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Api.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("google-native:apigateway/v1:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId    string  `pulumi:"apiId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupApiResult struct {
	// Created time.
	CreateTime string `pulumi:"createTime"`
	// Optional. Display name.
	DisplayName string `pulumi:"displayName"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels map[string]string `pulumi:"labels"`
	// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService string `pulumi:"managedService"`
	// Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
	Name string `pulumi:"name"`
	// State of the API.
	State string `pulumi:"state"`
	// Updated time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupApiOutput(ctx *pulumi.Context, args LookupApiOutputArgs, opts ...pulumi.InvokeOption) LookupApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiResult, error) {
			args := v.(LookupApiArgs)
			r, err := LookupApi(ctx, &args, opts...)
			var s LookupApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiResultOutput)
}

type LookupApiOutputArgs struct {
	ApiId    pulumi.StringInput    `pulumi:"apiId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiArgs)(nil)).Elem()
}

type LookupApiResultOutput struct{ *pulumi.OutputState }

func (LookupApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiResult)(nil)).Elem()
}

func (o LookupApiResultOutput) ToLookupApiResultOutput() LookupApiResultOutput {
	return o
}

func (o LookupApiResultOutput) ToLookupApiResultOutputWithContext(ctx context.Context) LookupApiResultOutput {
	return o
}

// Created time.
func (o LookupApiResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Display name.
func (o LookupApiResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
func (o LookupApiResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed). If not specified, a new Service will automatically be created in the same project as this API.
func (o LookupApiResultOutput) ManagedService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.ManagedService }).(pulumi.StringOutput)
}

// Resource name of the API. Format: projects/{project}/locations/global/apis/{api}
func (o LookupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the API.
func (o LookupApiResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.State }).(pulumi.StringOutput)
}

// Updated time.
func (o LookupApiResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
