// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a single ChannelConnection.
func LookupChannelConnection(ctx *pulumi.Context, args *LookupChannelConnectionArgs, opts ...pulumi.InvokeOption) (*LookupChannelConnectionResult, error) {
	var rv LookupChannelConnectionResult
	err := ctx.Invoke("google-native:eventarc/v1:getChannelConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelConnectionArgs struct {
	ChannelConnectionId string  `pulumi:"channelConnectionId"`
	Location            string  `pulumi:"location"`
	Project             *string `pulumi:"project"`
}

type LookupChannelConnectionResult struct {
	// Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
	ActivationToken string `pulumi:"activationToken"`
	// The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
	Channel string `pulumi:"channel"`
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// The name of the connection.
	Name string `pulumi:"name"`
	// Server assigned ID of the resource. The server guarantees uniqueness and immutability until deleted.
	Uid string `pulumi:"uid"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupChannelConnectionOutput(ctx *pulumi.Context, args LookupChannelConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupChannelConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelConnectionResult, error) {
			args := v.(LookupChannelConnectionArgs)
			r, err := LookupChannelConnection(ctx, &args, opts...)
			var s LookupChannelConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelConnectionResultOutput)
}

type LookupChannelConnectionOutputArgs struct {
	ChannelConnectionId pulumi.StringInput    `pulumi:"channelConnectionId"`
	Location            pulumi.StringInput    `pulumi:"location"`
	Project             pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupChannelConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelConnectionArgs)(nil)).Elem()
}

type LookupChannelConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupChannelConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelConnectionResult)(nil)).Elem()
}

func (o LookupChannelConnectionResultOutput) ToLookupChannelConnectionResultOutput() LookupChannelConnectionResultOutput {
	return o
}

func (o LookupChannelConnectionResultOutput) ToLookupChannelConnectionResultOutputWithContext(ctx context.Context) LookupChannelConnectionResultOutput {
	return o
}

// Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
func (o LookupChannelConnectionResultOutput) ActivationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.ActivationToken }).(pulumi.StringOutput)
}

// The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
func (o LookupChannelConnectionResultOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.Channel }).(pulumi.StringOutput)
}

// The creation time.
func (o LookupChannelConnectionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The name of the connection.
func (o LookupChannelConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Server assigned ID of the resource. The server guarantees uniqueness and immutability until deleted.
func (o LookupChannelConnectionResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The last-modified time.
func (o LookupChannelConnectionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelConnectionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelConnectionResultOutput{})
}
