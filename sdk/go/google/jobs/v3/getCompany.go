// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves specified company.
func LookupCompany(ctx *pulumi.Context, args *LookupCompanyArgs, opts ...pulumi.InvokeOption) (*LookupCompanyResult, error) {
	var rv LookupCompanyResult
	err := ctx.Invoke("google-native:jobs/v3:getCompany", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCompanyArgs struct {
	CompanyId string  `pulumi:"companyId"`
	Project   *string `pulumi:"project"`
}

type LookupCompanyResult struct {
	// Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri string `pulumi:"careerSiteUri"`
	// Derived details about the company.
	DerivedInfo CompanyDerivedInfoResponse `pulumi:"derivedInfo"`
	// The display name of the company, for example, "Google LLC".
	DisplayName string `pulumi:"displayName"`
	// Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText string `pulumi:"eeoText"`
	// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId string `pulumi:"externalId"`
	// Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress string `pulumi:"headquartersAddress"`
	// Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency bool `pulumi:"hiringAgency"`
	// Optional. A URI that hosts the employer's company logo.
	ImageUri string `pulumi:"imageUri"`
	// Optional. This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
	//
	// Deprecated: Optional. This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
	KeywordSearchableJobCustomAttributes []string `pulumi:"keywordSearchableJobCustomAttributes"`
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
	Name string `pulumi:"name"`
	// Optional. The employer's company size.
	Size string `pulumi:"size"`
	// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
	Suspended bool `pulumi:"suspended"`
	// Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri string `pulumi:"websiteUri"`
}

func LookupCompanyOutput(ctx *pulumi.Context, args LookupCompanyOutputArgs, opts ...pulumi.InvokeOption) LookupCompanyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCompanyResult, error) {
			args := v.(LookupCompanyArgs)
			r, err := LookupCompany(ctx, &args, opts...)
			var s LookupCompanyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCompanyResultOutput)
}

type LookupCompanyOutputArgs struct {
	CompanyId pulumi.StringInput    `pulumi:"companyId"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCompanyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCompanyArgs)(nil)).Elem()
}

type LookupCompanyResultOutput struct{ *pulumi.OutputState }

func (LookupCompanyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCompanyResult)(nil)).Elem()
}

func (o LookupCompanyResultOutput) ToLookupCompanyResultOutput() LookupCompanyResultOutput {
	return o
}

func (o LookupCompanyResultOutput) ToLookupCompanyResultOutputWithContext(ctx context.Context) LookupCompanyResultOutput {
	return o
}

// Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
func (o LookupCompanyResultOutput) CareerSiteUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.CareerSiteUri }).(pulumi.StringOutput)
}

// Derived details about the company.
func (o LookupCompanyResultOutput) DerivedInfo() CompanyDerivedInfoResponseOutput {
	return o.ApplyT(func(v LookupCompanyResult) CompanyDerivedInfoResponse { return v.DerivedInfo }).(CompanyDerivedInfoResponseOutput)
}

// The display name of the company, for example, "Google LLC".
func (o LookupCompanyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
func (o LookupCompanyResultOutput) EeoText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.EeoText }).(pulumi.StringOutput)
}

// Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
func (o LookupCompanyResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
func (o LookupCompanyResultOutput) HeadquartersAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.HeadquartersAddress }).(pulumi.StringOutput)
}

// Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
func (o LookupCompanyResultOutput) HiringAgency() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCompanyResult) bool { return v.HiringAgency }).(pulumi.BoolOutput)
}

// Optional. A URI that hosts the employer's company logo.
func (o LookupCompanyResultOutput) ImageUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.ImageUri }).(pulumi.StringOutput)
}

// Optional. This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
//
// Deprecated: Optional. This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
func (o LookupCompanyResultOutput) KeywordSearchableJobCustomAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCompanyResult) []string { return v.KeywordSearchableJobCustomAttributes }).(pulumi.StringArrayOutput)
}

// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
func (o LookupCompanyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. The employer's company size.
func (o LookupCompanyResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.Size }).(pulumi.StringOutput)
}

// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
func (o LookupCompanyResultOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCompanyResult) bool { return v.Suspended }).(pulumi.BoolOutput)
}

// Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
func (o LookupCompanyResultOutput) WebsiteUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompanyResult) string { return v.WebsiteUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCompanyResultOutput{})
}
