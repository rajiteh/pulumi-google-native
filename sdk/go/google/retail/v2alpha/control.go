// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Control. If the Control to create already exists, an ALREADY_EXISTS error is returned.
// Auto-naming is currently not supported for this resource.
type Control struct {
	pulumi.CustomResourceState

	// List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
	AssociatedServingConfigIds pulumi.StringArrayOutput `pulumi:"associatedServingConfigIds"`
	CatalogId                  pulumi.StringOutput      `pulumi:"catalogId"`
	// Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
	ControlId pulumi.StringOutput `pulumi:"controlId"`
	// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
	FacetSpec GoogleCloudRetailV2alphaSearchRequestFacetSpecResponseOutput `pulumi:"facetSpec"`
	Location  pulumi.StringOutput                                          `pulumi:"location"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
	Rule GoogleCloudRetailV2alphaRuleResponseOutput `pulumi:"rule"`
	// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
	SearchSolutionUseCase pulumi.StringArrayOutput `pulumi:"searchSolutionUseCase"`
	// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
	SolutionTypes pulumi.StringArrayOutput `pulumi:"solutionTypes"`
}

// NewControl registers a new resource with the given unique name, arguments, and options.
func NewControl(ctx *pulumi.Context,
	name string, args *ControlArgs, opts ...pulumi.ResourceOption) (*Control, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.ControlId == nil {
		return nil, errors.New("invalid value for required argument 'ControlId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.SolutionTypes == nil {
		return nil, errors.New("invalid value for required argument 'SolutionTypes'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"catalogId",
		"controlId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Control
	err := ctx.RegisterResource("google-native:retail/v2alpha:Control", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControl gets an existing Control resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlState, opts ...pulumi.ResourceOption) (*Control, error) {
	var resource Control
	err := ctx.ReadResource("google-native:retail/v2alpha:Control", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Control resources.
type controlState struct {
}

type ControlState struct {
}

func (ControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlState)(nil)).Elem()
}

type controlArgs struct {
	CatalogId string `pulumi:"catalogId"`
	// Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
	ControlId string `pulumi:"controlId"`
	// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	DisplayName string `pulumi:"displayName"`
	// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
	FacetSpec *GoogleCloudRetailV2alphaSearchRequestFacetSpec `pulumi:"facetSpec"`
	Location  *string                                         `pulumi:"location"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
	Rule *GoogleCloudRetailV2alphaRule `pulumi:"rule"`
	// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
	SearchSolutionUseCase []ControlSearchSolutionUseCaseItem `pulumi:"searchSolutionUseCase"`
	// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
	SolutionTypes []ControlSolutionTypesItem `pulumi:"solutionTypes"`
}

// The set of arguments for constructing a Control resource.
type ControlArgs struct {
	CatalogId pulumi.StringInput
	// Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
	ControlId pulumi.StringInput
	// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	DisplayName pulumi.StringInput
	// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
	FacetSpec GoogleCloudRetailV2alphaSearchRequestFacetSpecPtrInput
	Location  pulumi.StringPtrInput
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
	Rule GoogleCloudRetailV2alphaRulePtrInput
	// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
	SearchSolutionUseCase ControlSearchSolutionUseCaseItemArrayInput
	// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
	SolutionTypes ControlSolutionTypesItemArrayInput
}

func (ControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlArgs)(nil)).Elem()
}

type ControlInput interface {
	pulumi.Input

	ToControlOutput() ControlOutput
	ToControlOutputWithContext(ctx context.Context) ControlOutput
}

func (*Control) ElementType() reflect.Type {
	return reflect.TypeOf((**Control)(nil)).Elem()
}

func (i *Control) ToControlOutput() ControlOutput {
	return i.ToControlOutputWithContext(context.Background())
}

func (i *Control) ToControlOutputWithContext(ctx context.Context) ControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlOutput)
}

type ControlOutput struct{ *pulumi.OutputState }

func (ControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Control)(nil)).Elem()
}

func (o ControlOutput) ToControlOutput() ControlOutput {
	return o
}

func (o ControlOutput) ToControlOutputWithContext(ctx context.Context) ControlOutput {
	return o
}

// List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
func (o ControlOutput) AssociatedServingConfigIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Control) pulumi.StringArrayOutput { return v.AssociatedServingConfigIds }).(pulumi.StringArrayOutput)
}

func (o ControlOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
func (o ControlOutput) ControlId() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.ControlId }).(pulumi.StringOutput)
}

// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
func (o ControlOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
func (o ControlOutput) FacetSpec() GoogleCloudRetailV2alphaSearchRequestFacetSpecResponseOutput {
	return o.ApplyT(func(v *Control) GoogleCloudRetailV2alphaSearchRequestFacetSpecResponseOutput { return v.FacetSpec }).(GoogleCloudRetailV2alphaSearchRequestFacetSpecResponseOutput)
}

func (o ControlOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
func (o ControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ControlOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Control) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
func (o ControlOutput) Rule() GoogleCloudRetailV2alphaRuleResponseOutput {
	return o.ApplyT(func(v *Control) GoogleCloudRetailV2alphaRuleResponseOutput { return v.Rule }).(GoogleCloudRetailV2alphaRuleResponseOutput)
}

// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
func (o ControlOutput) SearchSolutionUseCase() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Control) pulumi.StringArrayOutput { return v.SearchSolutionUseCase }).(pulumi.StringArrayOutput)
}

// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
func (o ControlOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Control) pulumi.StringArrayOutput { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlInput)(nil)).Elem(), &Control{})
	pulumi.RegisterOutputType(ControlOutput{})
}
