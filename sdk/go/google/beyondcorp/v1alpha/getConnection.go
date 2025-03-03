// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Connection.
func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	ConnectionId string  `pulumi:"connectionId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupConnectionResult struct {
	// Address of the remote application endpoint for the BeyondCorp Connection.
	ApplicationEndpoint ApplicationEndpointResponse `pulumi:"applicationEndpoint"`
	// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
	Connectors []string `pulumi:"connectors"`
	// Timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
	DisplayName string `pulumi:"displayName"`
	// Optional. Gateway used by the connection.
	Gateway GatewayResponse `pulumi:"gateway"`
	// Optional. Resource labels to represent user provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique resource name of the connection. The name is ignored when creating a connection.
	Name string `pulumi:"name"`
	// The current state of the connection.
	State string `pulumi:"state"`
	// The type of network connectivity used by the connection.
	Type string `pulumi:"type"`
	// A unique identifier for the instance generated by the system.
	Uid string `pulumi:"uid"`
	// Timestamp when the resource was last modified.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	ConnectionId pulumi.StringInput    `pulumi:"connectionId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

// Address of the remote application endpoint for the BeyondCorp Connection.
func (o LookupConnectionResultOutput) ApplicationEndpoint() ApplicationEndpointResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) ApplicationEndpointResponse { return v.ApplicationEndpoint }).(ApplicationEndpointResponseOutput)
}

// Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
func (o LookupConnectionResultOutput) Connectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectionResult) []string { return v.Connectors }).(pulumi.StringArrayOutput)
}

// Timestamp when the resource was created.
func (o LookupConnectionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
func (o LookupConnectionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Gateway used by the connection.
func (o LookupConnectionResultOutput) Gateway() GatewayResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) GatewayResponse { return v.Gateway }).(GatewayResponseOutput)
}

// Optional. Resource labels to represent user provided metadata.
func (o LookupConnectionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Unique resource name of the connection. The name is ignored when creating a connection.
func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of the connection.
func (o LookupConnectionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of network connectivity used by the connection.
func (o LookupConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

// A unique identifier for the instance generated by the system.
func (o LookupConnectionResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Timestamp when the resource was last modified.
func (o LookupConnectionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
