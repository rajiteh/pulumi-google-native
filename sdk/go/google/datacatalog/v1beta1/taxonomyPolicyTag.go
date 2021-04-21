// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a policy tag in the specified taxonomy.
type TaxonomyPolicyTag struct {
	pulumi.CustomResourceState

	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags pulumi.StringArrayOutput `pulumi:"childPolicyTags"`
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Required. User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringOutput `pulumi:"parentPolicyTag"`
}

// NewTaxonomyPolicyTag registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyPolicyTag(ctx *pulumi.Context,
	name string, args *TaxonomyPolicyTagArgs, opts ...pulumi.ResourceOption) (*TaxonomyPolicyTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.PolicyTagsId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTagsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	if args.TaxonomiesId == nil {
		return nil, errors.New("invalid value for required argument 'TaxonomiesId'")
	}
	var resource TaxonomyPolicyTag
	err := ctx.RegisterResource("google-native:datacatalog/v1beta1:TaxonomyPolicyTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyPolicyTag gets an existing TaxonomyPolicyTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyPolicyTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyPolicyTagState, opts ...pulumi.ResourceOption) (*TaxonomyPolicyTag, error) {
	var resource TaxonomyPolicyTag
	err := ctx.ReadResource("google-native:datacatalog/v1beta1:TaxonomyPolicyTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyPolicyTag resources.
type taxonomyPolicyTagState struct {
	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags []string `pulumi:"childPolicyTags"`
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description *string `pulumi:"description"`
	// Required. User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName *string `pulumi:"displayName"`
	// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
	Name *string `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag *string `pulumi:"parentPolicyTag"`
}

type TaxonomyPolicyTagState struct {
	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags pulumi.StringArrayInput
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description pulumi.StringPtrInput
	// Required. User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringPtrInput
	// Resource name of this policy tag, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
	Name pulumi.StringPtrInput
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringPtrInput
}

func (TaxonomyPolicyTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyPolicyTagState)(nil)).Elem()
}

type taxonomyPolicyTagArgs struct {
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description *string `pulumi:"description"`
	// Required. User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName *string `pulumi:"displayName"`
	LocationsId string  `pulumi:"locationsId"`
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag *string `pulumi:"parentPolicyTag"`
	PolicyTagsId    string  `pulumi:"policyTagsId"`
	ProjectsId      string  `pulumi:"projectsId"`
	TaxonomiesId    string  `pulumi:"taxonomiesId"`
}

// The set of arguments for constructing a TaxonomyPolicyTag resource.
type TaxonomyPolicyTagArgs struct {
	// Description of this policy tag. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description. If not set, defaults to an empty description.
	Description pulumi.StringPtrInput
	// Required. User defined name of this policy tag. It must: be unique within the parent taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringPtrInput
	LocationsId pulumi.StringInput
	// Resource name of this policy tag's parent policy tag (e.g. for the "LatLong" policy tag in the example above, this field contains the resource name of the "Geolocation" policy tag). If empty, it means this policy tag is a top level policy tag (e.g. this field is empty for the "Geolocation" policy tag in the example above). If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringPtrInput
	PolicyTagsId    pulumi.StringInput
	ProjectsId      pulumi.StringInput
	TaxonomiesId    pulumi.StringInput
}

func (TaxonomyPolicyTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyPolicyTagArgs)(nil)).Elem()
}

type TaxonomyPolicyTagInput interface {
	pulumi.Input

	ToTaxonomyPolicyTagOutput() TaxonomyPolicyTagOutput
	ToTaxonomyPolicyTagOutputWithContext(ctx context.Context) TaxonomyPolicyTagOutput
}

func (*TaxonomyPolicyTag) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyPolicyTag)(nil))
}

func (i *TaxonomyPolicyTag) ToTaxonomyPolicyTagOutput() TaxonomyPolicyTagOutput {
	return i.ToTaxonomyPolicyTagOutputWithContext(context.Background())
}

func (i *TaxonomyPolicyTag) ToTaxonomyPolicyTagOutputWithContext(ctx context.Context) TaxonomyPolicyTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyPolicyTagOutput)
}

type TaxonomyPolicyTagOutput struct {
	*pulumi.OutputState
}

func (TaxonomyPolicyTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyPolicyTag)(nil))
}

func (o TaxonomyPolicyTagOutput) ToTaxonomyPolicyTagOutput() TaxonomyPolicyTagOutput {
	return o
}

func (o TaxonomyPolicyTagOutput) ToTaxonomyPolicyTagOutputWithContext(ctx context.Context) TaxonomyPolicyTagOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TaxonomyPolicyTagOutput{})
}