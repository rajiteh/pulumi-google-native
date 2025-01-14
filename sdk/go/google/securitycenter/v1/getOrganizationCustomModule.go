// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a SecurityHealthAnalyticsCustomModule.
func LookupOrganizationCustomModule(ctx *pulumi.Context, args *LookupOrganizationCustomModuleArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationCustomModuleResult, error) {
	var rv LookupOrganizationCustomModuleResult
	err := ctx.Invoke("google-native:securitycenter/v1:getOrganizationCustomModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationCustomModuleArgs struct {
	CustomModuleId string `pulumi:"customModuleId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupOrganizationCustomModuleResult struct {
	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
	AncestorModule string `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig GoogleCloudSecuritycenterV1CustomConfigResponse `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName string `pulumi:"displayName"`
	// The enablement state of the custom module.
	EnablementState string `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor string `pulumi:"lastEditor"`
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name string `pulumi:"name"`
	// The time at which the custom module was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupOrganizationCustomModuleOutput(ctx *pulumi.Context, args LookupOrganizationCustomModuleOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationCustomModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationCustomModuleResult, error) {
			args := v.(LookupOrganizationCustomModuleArgs)
			r, err := LookupOrganizationCustomModule(ctx, &args, opts...)
			var s LookupOrganizationCustomModuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationCustomModuleResultOutput)
}

type LookupOrganizationCustomModuleOutputArgs struct {
	CustomModuleId pulumi.StringInput `pulumi:"customModuleId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupOrganizationCustomModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationCustomModuleArgs)(nil)).Elem()
}

type LookupOrganizationCustomModuleResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationCustomModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationCustomModuleResult)(nil)).Elem()
}

func (o LookupOrganizationCustomModuleResultOutput) ToLookupOrganizationCustomModuleResultOutput() LookupOrganizationCustomModuleResultOutput {
	return o
}

func (o LookupOrganizationCustomModuleResultOutput) ToLookupOrganizationCustomModuleResultOutputWithContext(ctx context.Context) LookupOrganizationCustomModuleResultOutput {
	return o
}

// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
func (o LookupOrganizationCustomModuleResultOutput) AncestorModule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.AncestorModule }).(pulumi.StringOutput)
}

// The user specified custom configuration for the module.
func (o LookupOrganizationCustomModuleResultOutput) CustomConfig() GoogleCloudSecuritycenterV1CustomConfigResponseOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) GoogleCloudSecuritycenterV1CustomConfigResponse {
		return v.CustomConfig
	}).(GoogleCloudSecuritycenterV1CustomConfigResponseOutput)
}

// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
func (o LookupOrganizationCustomModuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The enablement state of the custom module.
func (o LookupOrganizationCustomModuleResultOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.EnablementState }).(pulumi.StringOutput)
}

// The editor that last updated the custom module.
func (o LookupOrganizationCustomModuleResultOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.LastEditor }).(pulumi.StringOutput)
}

// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
func (o LookupOrganizationCustomModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The time at which the custom module was last updated.
func (o LookupOrganizationCustomModuleResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomModuleResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationCustomModuleResultOutput{})
}
