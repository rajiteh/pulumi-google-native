// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve a default supported Idp configuration for an Identity Toolkit project.
func LookupDefaultSupportedIdpConfig(ctx *pulumi.Context, args *LookupDefaultSupportedIdpConfigArgs, opts ...pulumi.InvokeOption) (*LookupDefaultSupportedIdpConfigResult, error) {
	var rv LookupDefaultSupportedIdpConfigResult
	err := ctx.Invoke("google-native:identitytoolkit/v2:getDefaultSupportedIdpConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultSupportedIdpConfigArgs struct {
	DefaultSupportedIdpConfigId string  `pulumi:"defaultSupportedIdpConfigId"`
	Project                     *string `pulumi:"project"`
	TenantId                    string  `pulumi:"tenantId"`
}

type LookupDefaultSupportedIdpConfigResult struct {
	// Additional config for Apple-based projects.
	AppleSignInConfig GoogleCloudIdentitytoolkitAdminV2AppleSignInConfigResponse `pulumi:"appleSignInConfig"`
	// OAuth client ID.
	ClientId string `pulumi:"clientId"`
	// OAuth client secret.
	ClientSecret string `pulumi:"clientSecret"`
	// True if allows the user to sign in with the provider.
	Enabled bool `pulumi:"enabled"`
	// The name of the DefaultSupportedIdpConfig resource, for example: "projects/my-awesome-project/defaultSupportedIdpConfigs/google.com"
	Name string `pulumi:"name"`
}

func LookupDefaultSupportedIdpConfigOutput(ctx *pulumi.Context, args LookupDefaultSupportedIdpConfigOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultSupportedIdpConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultSupportedIdpConfigResult, error) {
			args := v.(LookupDefaultSupportedIdpConfigArgs)
			r, err := LookupDefaultSupportedIdpConfig(ctx, &args, opts...)
			var s LookupDefaultSupportedIdpConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultSupportedIdpConfigResultOutput)
}

type LookupDefaultSupportedIdpConfigOutputArgs struct {
	DefaultSupportedIdpConfigId pulumi.StringInput    `pulumi:"defaultSupportedIdpConfigId"`
	Project                     pulumi.StringPtrInput `pulumi:"project"`
	TenantId                    pulumi.StringInput    `pulumi:"tenantId"`
}

func (LookupDefaultSupportedIdpConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultSupportedIdpConfigArgs)(nil)).Elem()
}

type LookupDefaultSupportedIdpConfigResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultSupportedIdpConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultSupportedIdpConfigResult)(nil)).Elem()
}

func (o LookupDefaultSupportedIdpConfigResultOutput) ToLookupDefaultSupportedIdpConfigResultOutput() LookupDefaultSupportedIdpConfigResultOutput {
	return o
}

func (o LookupDefaultSupportedIdpConfigResultOutput) ToLookupDefaultSupportedIdpConfigResultOutputWithContext(ctx context.Context) LookupDefaultSupportedIdpConfigResultOutput {
	return o
}

// Additional config for Apple-based projects.
func (o LookupDefaultSupportedIdpConfigResultOutput) AppleSignInConfig() GoogleCloudIdentitytoolkitAdminV2AppleSignInConfigResponseOutput {
	return o.ApplyT(func(v LookupDefaultSupportedIdpConfigResult) GoogleCloudIdentitytoolkitAdminV2AppleSignInConfigResponse {
		return v.AppleSignInConfig
	}).(GoogleCloudIdentitytoolkitAdminV2AppleSignInConfigResponseOutput)
}

// OAuth client ID.
func (o LookupDefaultSupportedIdpConfigResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultSupportedIdpConfigResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// OAuth client secret.
func (o LookupDefaultSupportedIdpConfigResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultSupportedIdpConfigResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// True if allows the user to sign in with the provider.
func (o LookupDefaultSupportedIdpConfigResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDefaultSupportedIdpConfigResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The name of the DefaultSupportedIdpConfig resource, for example: "projects/my-awesome-project/defaultSupportedIdpConfigs/google.com"
func (o LookupDefaultSupportedIdpConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultSupportedIdpConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultSupportedIdpConfigResultOutput{})
}
