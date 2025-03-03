// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Tenant.
func LookupProxyConfig(ctx *pulumi.Context, args *LookupProxyConfigArgs, opts ...pulumi.InvokeOption) (*LookupProxyConfigResult, error) {
	var rv LookupProxyConfigResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getProxyConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProxyConfigArgs struct {
	OrganizationId string `pulumi:"organizationId"`
	ProxyConfigId  string `pulumi:"proxyConfigId"`
	TenantId       string `pulumi:"tenantId"`
}

type LookupProxyConfigResult struct {
	// Optional. Information to facilitate Authentication against the proxy server.
	AuthenticationInfo GoogleCloudBeyondcorpPartnerservicesV1alphaAuthenticationInfoResponse `pulumi:"authenticationInfo"`
	// Timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. An arbitrary caller-provided name for the ProxyConfig. Cannot exceed 64 characters.
	DisplayName string `pulumi:"displayName"`
	// ProxyConfig resource name.
	Name string `pulumi:"name"`
	// The URI of the proxy server.
	ProxyUri string `pulumi:"proxyUri"`
	// Routing info to direct traffic to the proxy server.
	RoutingInfo GoogleCloudBeyondcorpPartnerservicesV1alphaRoutingInfoResponse `pulumi:"routingInfo"`
	// Transport layer information to verify for the proxy server.
	TransportInfo GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponse `pulumi:"transportInfo"`
	// Timestamp when the resource was last modified.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupProxyConfigOutput(ctx *pulumi.Context, args LookupProxyConfigOutputArgs, opts ...pulumi.InvokeOption) LookupProxyConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProxyConfigResult, error) {
			args := v.(LookupProxyConfigArgs)
			r, err := LookupProxyConfig(ctx, &args, opts...)
			var s LookupProxyConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProxyConfigResultOutput)
}

type LookupProxyConfigOutputArgs struct {
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
	ProxyConfigId  pulumi.StringInput `pulumi:"proxyConfigId"`
	TenantId       pulumi.StringInput `pulumi:"tenantId"`
}

func (LookupProxyConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyConfigArgs)(nil)).Elem()
}

type LookupProxyConfigResultOutput struct{ *pulumi.OutputState }

func (LookupProxyConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyConfigResult)(nil)).Elem()
}

func (o LookupProxyConfigResultOutput) ToLookupProxyConfigResultOutput() LookupProxyConfigResultOutput {
	return o
}

func (o LookupProxyConfigResultOutput) ToLookupProxyConfigResultOutputWithContext(ctx context.Context) LookupProxyConfigResultOutput {
	return o
}

// Optional. Information to facilitate Authentication against the proxy server.
func (o LookupProxyConfigResultOutput) AuthenticationInfo() GoogleCloudBeyondcorpPartnerservicesV1alphaAuthenticationInfoResponseOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) GoogleCloudBeyondcorpPartnerservicesV1alphaAuthenticationInfoResponse {
		return v.AuthenticationInfo
	}).(GoogleCloudBeyondcorpPartnerservicesV1alphaAuthenticationInfoResponseOutput)
}

// Timestamp when the resource was created.
func (o LookupProxyConfigResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary caller-provided name for the ProxyConfig. Cannot exceed 64 characters.
func (o LookupProxyConfigResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// ProxyConfig resource name.
func (o LookupProxyConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URI of the proxy server.
func (o LookupProxyConfigResultOutput) ProxyUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) string { return v.ProxyUri }).(pulumi.StringOutput)
}

// Routing info to direct traffic to the proxy server.
func (o LookupProxyConfigResultOutput) RoutingInfo() GoogleCloudBeyondcorpPartnerservicesV1alphaRoutingInfoResponseOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) GoogleCloudBeyondcorpPartnerservicesV1alphaRoutingInfoResponse {
		return v.RoutingInfo
	}).(GoogleCloudBeyondcorpPartnerservicesV1alphaRoutingInfoResponseOutput)
}

// Transport layer information to verify for the proxy server.
func (o LookupProxyConfigResultOutput) TransportInfo() GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponseOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponse {
		return v.TransportInfo
	}).(GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponseOutput)
}

// Timestamp when the resource was last modified.
func (o LookupProxyConfigResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyConfigResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProxyConfigResultOutput{})
}
