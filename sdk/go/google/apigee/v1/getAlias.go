// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an alias.
func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("google-native:apigee/v1:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAliasArgs struct {
	AliasId        string `pulumi:"aliasId"`
	EnvironmentId  string `pulumi:"environmentId"`
	KeystoreId     string `pulumi:"keystoreId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupAliasResult struct {
	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias string `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo GoogleCloudApigeeV1CertificateResponse `pulumi:"certsInfo"`
	// Type of alias.
	Type string `pulumi:"type"`
}

func LookupAliasOutput(ctx *pulumi.Context, args LookupAliasOutputArgs, opts ...pulumi.InvokeOption) LookupAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAliasResult, error) {
			args := v.(LookupAliasArgs)
			r, err := LookupAlias(ctx, &args, opts...)
			var s LookupAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAliasResultOutput)
}

type LookupAliasOutputArgs struct {
	AliasId        pulumi.StringInput `pulumi:"aliasId"`
	EnvironmentId  pulumi.StringInput `pulumi:"environmentId"`
	KeystoreId     pulumi.StringInput `pulumi:"keystoreId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasArgs)(nil)).Elem()
}

type LookupAliasResultOutput struct{ *pulumi.OutputState }

func (LookupAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasResult)(nil)).Elem()
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutput() LookupAliasResultOutput {
	return o
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutputWithContext(ctx context.Context) LookupAliasResultOutput {
	return o
}

// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
func (o LookupAliasResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAliasResult) string { return v.Alias }).(pulumi.StringOutput)
}

// Chain of certificates under this alias.
func (o LookupAliasResultOutput) CertsInfo() GoogleCloudApigeeV1CertificateResponseOutput {
	return o.ApplyT(func(v LookupAliasResult) GoogleCloudApigeeV1CertificateResponse { return v.CertsInfo }).(GoogleCloudApigeeV1CertificateResponseOutput)
}

// Type of alias.
func (o LookupAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAliasResultOutput{})
}
