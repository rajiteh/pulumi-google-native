// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an authorized orgs desc based on the resource name.
func LookupAuthorizedOrgsDesc(ctx *pulumi.Context, args *LookupAuthorizedOrgsDescArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizedOrgsDescResult, error) {
	var rv LookupAuthorizedOrgsDescResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1:getAuthorizedOrgsDesc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizedOrgsDescArgs struct {
	AccessPolicyId       string `pulumi:"accessPolicyId"`
	AuthorizedOrgsDescId string `pulumi:"authorizedOrgsDescId"`
}

type LookupAuthorizedOrgsDescResult struct {
	// The asset type of this authorized orgs desc. Valid values are `ASSET_TYPE_DEVICE`, and `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	AssetType string `pulumi:"assetType"`
	// The direction of the authorization relationship between this organization and the organizations listed in the `orgs` field. The valid values for this field include the following: `AUTHORIZATION_DIRECTION_FROM`: Allows this organization to evaluate traffic in the organizations listed in the `orgs` field. `AUTHORIZATION_DIRECTION_TO`: Allows the organizations listed in the `orgs` field to evaluate the traffic in this organization. For the authorization relationship to take effect, all of the organizations must authorize and specify the appropriate relationship direction. For example, if organization A authorized organization B and C to evaluate its traffic, by specifying `AUTHORIZATION_DIRECTION_TO` as the authorization direction, organizations B and C must specify `AUTHORIZATION_DIRECTION_FROM` as the authorization direction in their `AuthorizedOrgsDesc` resource.
	AuthorizationDirection string `pulumi:"authorizationDirection"`
	// A granular control type for authorization levels. Valid value is `AUTHORIZATION_TYPE_TRUST`.
	AuthorizationType string `pulumi:"authorizationType"`
	// Resource name for the `AuthorizedOrgsDesc`. Format: `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`. The `authorized_orgs_desc` component must begin with a letter, followed by alphanumeric characters or `_`. After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	Name string `pulumi:"name"`
	// The list of organization ids in this AuthorizedOrgsDesc. Format: `organizations/` Example: `organizations/123456`
	Orgs []string `pulumi:"orgs"`
}

func LookupAuthorizedOrgsDescOutput(ctx *pulumi.Context, args LookupAuthorizedOrgsDescOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizedOrgsDescResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizedOrgsDescResult, error) {
			args := v.(LookupAuthorizedOrgsDescArgs)
			r, err := LookupAuthorizedOrgsDesc(ctx, &args, opts...)
			var s LookupAuthorizedOrgsDescResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizedOrgsDescResultOutput)
}

type LookupAuthorizedOrgsDescOutputArgs struct {
	AccessPolicyId       pulumi.StringInput `pulumi:"accessPolicyId"`
	AuthorizedOrgsDescId pulumi.StringInput `pulumi:"authorizedOrgsDescId"`
}

func (LookupAuthorizedOrgsDescOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedOrgsDescArgs)(nil)).Elem()
}

type LookupAuthorizedOrgsDescResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizedOrgsDescResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedOrgsDescResult)(nil)).Elem()
}

func (o LookupAuthorizedOrgsDescResultOutput) ToLookupAuthorizedOrgsDescResultOutput() LookupAuthorizedOrgsDescResultOutput {
	return o
}

func (o LookupAuthorizedOrgsDescResultOutput) ToLookupAuthorizedOrgsDescResultOutputWithContext(ctx context.Context) LookupAuthorizedOrgsDescResultOutput {
	return o
}

// The asset type of this authorized orgs desc. Valid values are `ASSET_TYPE_DEVICE`, and `ASSET_TYPE_CREDENTIAL_STRENGTH`.
func (o LookupAuthorizedOrgsDescResultOutput) AssetType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AssetType }).(pulumi.StringOutput)
}

// The direction of the authorization relationship between this organization and the organizations listed in the `orgs` field. The valid values for this field include the following: `AUTHORIZATION_DIRECTION_FROM`: Allows this organization to evaluate traffic in the organizations listed in the `orgs` field. `AUTHORIZATION_DIRECTION_TO`: Allows the organizations listed in the `orgs` field to evaluate the traffic in this organization. For the authorization relationship to take effect, all of the organizations must authorize and specify the appropriate relationship direction. For example, if organization A authorized organization B and C to evaluate its traffic, by specifying `AUTHORIZATION_DIRECTION_TO` as the authorization direction, organizations B and C must specify `AUTHORIZATION_DIRECTION_FROM` as the authorization direction in their `AuthorizedOrgsDesc` resource.
func (o LookupAuthorizedOrgsDescResultOutput) AuthorizationDirection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AuthorizationDirection }).(pulumi.StringOutput)
}

// A granular control type for authorization levels. Valid value is `AUTHORIZATION_TYPE_TRUST`.
func (o LookupAuthorizedOrgsDescResultOutput) AuthorizationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AuthorizationType }).(pulumi.StringOutput)
}

// Resource name for the `AuthorizedOrgsDesc`. Format: `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`. The `authorized_orgs_desc` component must begin with a letter, followed by alphanumeric characters or `_`. After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
func (o LookupAuthorizedOrgsDescResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of organization ids in this AuthorizedOrgsDesc. Format: `organizations/` Example: `organizations/123456`
func (o LookupAuthorizedOrgsDescResultOutput) Orgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) []string { return v.Orgs }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizedOrgsDescResultOutput{})
}
