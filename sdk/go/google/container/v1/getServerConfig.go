// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns configuration info about the Google Kubernetes Engine service.
func GetServerConfig(ctx *pulumi.Context, args *GetServerConfigArgs, opts ...pulumi.InvokeOption) (*GetServerConfigResult, error) {
	var rv GetServerConfigResult
	err := ctx.Invoke("google-native:container/v1:getServerConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetServerConfigArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	ProjectId *string `pulumi:"projectId"`
	Zone      *string `pulumi:"zone"`
}

type GetServerConfigResult struct {
	// List of release channel configurations.
	Channels []ReleaseChannelConfigResponse `pulumi:"channels"`
	// Version of Kubernetes the service deploys by default.
	DefaultClusterVersion string `pulumi:"defaultClusterVersion"`
	// Default image type.
	DefaultImageType string `pulumi:"defaultImageType"`
	// List of valid image types.
	ValidImageTypes []string `pulumi:"validImageTypes"`
	// List of valid master versions, in descending order.
	ValidMasterVersions []string `pulumi:"validMasterVersions"`
	// List of valid node upgrade target versions, in descending order.
	ValidNodeVersions []string `pulumi:"validNodeVersions"`
}

func GetServerConfigOutput(ctx *pulumi.Context, args GetServerConfigOutputArgs, opts ...pulumi.InvokeOption) GetServerConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerConfigResult, error) {
			args := v.(GetServerConfigArgs)
			r, err := GetServerConfig(ctx, &args, opts...)
			var s GetServerConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerConfigResultOutput)
}

type GetServerConfigOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	Zone      pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetServerConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerConfigArgs)(nil)).Elem()
}

type GetServerConfigResultOutput struct{ *pulumi.OutputState }

func (GetServerConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerConfigResult)(nil)).Elem()
}

func (o GetServerConfigResultOutput) ToGetServerConfigResultOutput() GetServerConfigResultOutput {
	return o
}

func (o GetServerConfigResultOutput) ToGetServerConfigResultOutputWithContext(ctx context.Context) GetServerConfigResultOutput {
	return o
}

// List of release channel configurations.
func (o GetServerConfigResultOutput) Channels() ReleaseChannelConfigResponseArrayOutput {
	return o.ApplyT(func(v GetServerConfigResult) []ReleaseChannelConfigResponse { return v.Channels }).(ReleaseChannelConfigResponseArrayOutput)
}

// Version of Kubernetes the service deploys by default.
func (o GetServerConfigResultOutput) DefaultClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerConfigResult) string { return v.DefaultClusterVersion }).(pulumi.StringOutput)
}

// Default image type.
func (o GetServerConfigResultOutput) DefaultImageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerConfigResult) string { return v.DefaultImageType }).(pulumi.StringOutput)
}

// List of valid image types.
func (o GetServerConfigResultOutput) ValidImageTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerConfigResult) []string { return v.ValidImageTypes }).(pulumi.StringArrayOutput)
}

// List of valid master versions, in descending order.
func (o GetServerConfigResultOutput) ValidMasterVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerConfigResult) []string { return v.ValidMasterVersions }).(pulumi.StringArrayOutput)
}

// List of valid node upgrade target versions, in descending order.
func (o GetServerConfigResultOutput) ValidNodeVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerConfigResult) []string { return v.ValidNodeVersions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerConfigResultOutput{})
}
