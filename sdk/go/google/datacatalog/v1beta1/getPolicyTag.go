// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a policy tag.
func LookupPolicyTag(ctx *pulumi.Context, args *LookupPolicyTagArgs, opts ...pulumi.InvokeOption) (*LookupPolicyTagResult, error) {
	var rv LookupPolicyTagResult
	err := ctx.Invoke("google-native:datacatalog/v1beta1:getPolicyTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyTagArgs struct {
	Location    string  `pulumi:"location"`
	PolicyTagId string  `pulumi:"policyTagId"`
	Project     *string `pulumi:"project"`
	TaxonomyId  string  `pulumi:"taxonomyId"`
}

type LookupPolicyTagResult struct {
	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags []string `pulumi:"childPolicyTags"`
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description string `pulumi:"description"`
	// User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName string `pulumi:"displayName"`
	// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
	Name string `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag string `pulumi:"parentPolicyTag"`
}

func LookupPolicyTagOutput(ctx *pulumi.Context, args LookupPolicyTagOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyTagResult, error) {
			args := v.(LookupPolicyTagArgs)
			r, err := LookupPolicyTag(ctx, &args, opts...)
			var s LookupPolicyTagResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyTagResultOutput)
}

type LookupPolicyTagOutputArgs struct {
	Location    pulumi.StringInput    `pulumi:"location"`
	PolicyTagId pulumi.StringInput    `pulumi:"policyTagId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	TaxonomyId  pulumi.StringInput    `pulumi:"taxonomyId"`
}

func (LookupPolicyTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyTagArgs)(nil)).Elem()
}

type LookupPolicyTagResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyTagResult)(nil)).Elem()
}

func (o LookupPolicyTagResultOutput) ToLookupPolicyTagResultOutput() LookupPolicyTagResultOutput {
	return o
}

func (o LookupPolicyTagResultOutput) ToLookupPolicyTagResultOutputWithContext(ctx context.Context) LookupPolicyTagResultOutput {
	return o
}

// Resource names of child policy tags of this policy tag.
func (o LookupPolicyTagResultOutput) ChildPolicyTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyTagResult) []string { return v.ChildPolicyTags }).(pulumi.StringArrayOutput)
}

// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
func (o LookupPolicyTagResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyTagResult) string { return v.Description }).(pulumi.StringOutput)
}

// User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
func (o LookupPolicyTagResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyTagResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
func (o LookupPolicyTagResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyTagResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
func (o LookupPolicyTagResultOutput) ParentPolicyTag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyTagResult) string { return v.ParentPolicyTag }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyTagResultOutput{})
}
