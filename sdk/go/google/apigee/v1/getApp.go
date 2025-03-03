// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the details for a developer app.
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("google-native:apigee/v1:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	AppId          string  `pulumi:"appId"`
	DeveloperId    string  `pulumi:"developerId"`
	Entity         *string `pulumi:"entity"`
	OrganizationId string  `pulumi:"organizationId"`
	Query          *string `pulumi:"query"`
}

type LookupAppResult struct {
	// List of API products associated with the developer app.
	ApiProducts []string `pulumi:"apiProducts"`
	// Developer app family.
	AppFamily string `pulumi:"appFamily"`
	// ID of the developer app.
	AppId string `pulumi:"appId"`
	// List of attributes for the developer app.
	Attributes []GoogleCloudApigeeV1AttributeResponse `pulumi:"attributes"`
	// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
	CallbackUrl string `pulumi:"callbackUrl"`
	// Time the developer app was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Set of credentials for the developer app consisting of the consumer key/secret pairs associated with the API products.
	Credentials []GoogleCloudApigeeV1CredentialResponse `pulumi:"credentials"`
	// ID of the developer.
	DeveloperId string `pulumi:"developerId"`
	// Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
	KeyExpiresIn string `pulumi:"keyExpiresIn"`
	// Time the developer app was modified in milliseconds since epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// Name of the developer app.
	Name string `pulumi:"name"`
	// Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
	Scopes []string `pulumi:"scopes"`
	// Status of the credential. Valid values include `approved` or `revoked`.
	Status string `pulumi:"status"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	AppId          pulumi.StringInput    `pulumi:"appId"`
	DeveloperId    pulumi.StringInput    `pulumi:"developerId"`
	Entity         pulumi.StringPtrInput `pulumi:"entity"`
	OrganizationId pulumi.StringInput    `pulumi:"organizationId"`
	Query          pulumi.StringPtrInput `pulumi:"query"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

// List of API products associated with the developer app.
func (o LookupAppResultOutput) ApiProducts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []string { return v.ApiProducts }).(pulumi.StringArrayOutput)
}

// Developer app family.
func (o LookupAppResultOutput) AppFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.AppFamily }).(pulumi.StringOutput)
}

// ID of the developer app.
func (o LookupAppResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.AppId }).(pulumi.StringOutput)
}

// List of attributes for the developer app.
func (o LookupAppResultOutput) Attributes() GoogleCloudApigeeV1AttributeResponseArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []GoogleCloudApigeeV1AttributeResponse { return v.Attributes }).(GoogleCloudApigeeV1AttributeResponseArrayOutput)
}

// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
func (o LookupAppResultOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

// Time the developer app was created in milliseconds since epoch.
func (o LookupAppResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Set of credentials for the developer app consisting of the consumer key/secret pairs associated with the API products.
func (o LookupAppResultOutput) Credentials() GoogleCloudApigeeV1CredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []GoogleCloudApigeeV1CredentialResponse { return v.Credentials }).(GoogleCloudApigeeV1CredentialResponseArrayOutput)
}

// ID of the developer.
func (o LookupAppResultOutput) DeveloperId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.DeveloperId }).(pulumi.StringOutput)
}

// Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
func (o LookupAppResultOutput) KeyExpiresIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.KeyExpiresIn }).(pulumi.StringOutput)
}

// Time the developer app was modified in milliseconds since epoch.
func (o LookupAppResultOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// Name of the developer app.
func (o LookupAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
func (o LookupAppResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Status of the credential. Valid values include `approved` or `revoked`.
func (o LookupAppResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
