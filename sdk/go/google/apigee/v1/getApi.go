// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an API proxy including a list of existing revisions.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("google-native:apigee/v1:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId          string `pulumi:"apiId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupApiResult struct {
	// The type of the API proxy.
	ApiProxyType string `pulumi:"apiProxyType"`
	// User labels applied to this API Proxy.
	Labels map[string]string `pulumi:"labels"`
	// The id of the most recently created revision for this api proxy.
	LatestRevisionId string `pulumi:"latestRevisionId"`
	// Metadata describing the API proxy.
	MetaData GoogleCloudApigeeV1EntityMetadataResponse `pulumi:"metaData"`
	// Name of the API proxy.
	Name string `pulumi:"name"`
	// Whether this proxy is read-only. A read-only proxy cannot have new revisions created through calls to CreateApiProxyRevision. A proxy is read-only if it was generated by an archive.
	ReadOnly bool `pulumi:"readOnly"`
	// List of revisions defined for the API proxy.
	Revision []string `pulumi:"revision"`
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
	ApiId          pulumi.StringInput `pulumi:"apiId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
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

// The type of the API proxy.
func (o LookupApiResultOutput) ApiProxyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.ApiProxyType }).(pulumi.StringOutput)
}

// User labels applied to this API Proxy.
func (o LookupApiResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApiResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The id of the most recently created revision for this api proxy.
func (o LookupApiResultOutput) LatestRevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.LatestRevisionId }).(pulumi.StringOutput)
}

// Metadata describing the API proxy.
func (o LookupApiResultOutput) MetaData() GoogleCloudApigeeV1EntityMetadataResponseOutput {
	return o.ApplyT(func(v LookupApiResult) GoogleCloudApigeeV1EntityMetadataResponse { return v.MetaData }).(GoogleCloudApigeeV1EntityMetadataResponseOutput)
}

// Name of the API proxy.
func (o LookupApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether this proxy is read-only. A read-only proxy cannot have new revisions created through calls to CreateApiProxyRevision. A proxy is read-only if it was generated by an archive.
func (o LookupApiResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

// List of revisions defined for the API proxy.
func (o LookupApiResultOutput) Revision() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiResult) []string { return v.Revision }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiResultOutput{})
}
