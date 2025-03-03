// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds a new Feature.
// Auto-naming is currently not supported for this resource.
type Feature struct {
	pulumi.CustomResourceState

	// When the Feature resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// When the Feature resource was deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The ID of the feature to create.
	FeatureId pulumi.StringPtrOutput `pulumi:"featureId"`
	// Optional. Feature configuration applicable to all memberships of the fleet.
	FleetDefaultMemberConfig CommonFleetDefaultMemberConfigSpecResponseOutput `pulumi:"fleetDefaultMemberConfig"`
	// Labels for this Feature.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	MembershipSpecs pulumi.StringMapOutput `pulumi:"membershipSpecs"`
	// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
	MembershipStates pulumi.StringMapOutput `pulumi:"membershipStates"`
	// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// State of the Feature resource itself.
	ResourceState FeatureResourceStateResponseOutput `pulumi:"resourceState"`
	// Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	ScopeSpecs pulumi.StringMapOutput `pulumi:"scopeSpecs"`
	// Scope-specific Feature status. If this Feature does report any per-Scope status, this field may be unused. The keys indicate which Scope the state is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project.
	ScopeStates pulumi.StringMapOutput `pulumi:"scopeStates"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec CommonFeatureSpecResponseOutput `pulumi:"spec"`
	// The Hub-wide Feature state.
	State CommonFeatureStateResponseOutput `pulumi:"state"`
	// When the Feature resource was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFeature registers a new resource with the given unique name, arguments, and options.
func NewFeature(ctx *pulumi.Context,
	name string, args *FeatureArgs, opts ...pulumi.ResourceOption) (*Feature, error) {
	if args == nil {
		args = &FeatureArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Feature
	err := ctx.RegisterResource("google-native:gkehub/v1beta:Feature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeature gets an existing Feature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureState, opts ...pulumi.ResourceOption) (*Feature, error) {
	var resource Feature
	err := ctx.ReadResource("google-native:gkehub/v1beta:Feature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Feature resources.
type featureState struct {
}

type FeatureState struct {
}

func (FeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureState)(nil)).Elem()
}

type featureArgs struct {
	// The ID of the feature to create.
	FeatureId *string `pulumi:"featureId"`
	// Optional. Feature configuration applicable to all memberships of the fleet.
	FleetDefaultMemberConfig *CommonFleetDefaultMemberConfigSpec `pulumi:"fleetDefaultMemberConfig"`
	// Labels for this Feature.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	MembershipSpecs map[string]string `pulumi:"membershipSpecs"`
	Project         *string           `pulumi:"project"`
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	ScopeSpecs map[string]string `pulumi:"scopeSpecs"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec *CommonFeatureSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Feature resource.
type FeatureArgs struct {
	// The ID of the feature to create.
	FeatureId pulumi.StringPtrInput
	// Optional. Feature configuration applicable to all memberships of the fleet.
	FleetDefaultMemberConfig CommonFleetDefaultMemberConfigSpecPtrInput
	// Labels for this Feature.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	MembershipSpecs pulumi.StringMapInput
	Project         pulumi.StringPtrInput
	// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	ScopeSpecs pulumi.StringMapInput
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec CommonFeatureSpecPtrInput
}

func (FeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureArgs)(nil)).Elem()
}

type FeatureInput interface {
	pulumi.Input

	ToFeatureOutput() FeatureOutput
	ToFeatureOutputWithContext(ctx context.Context) FeatureOutput
}

func (*Feature) ElementType() reflect.Type {
	return reflect.TypeOf((**Feature)(nil)).Elem()
}

func (i *Feature) ToFeatureOutput() FeatureOutput {
	return i.ToFeatureOutputWithContext(context.Background())
}

func (i *Feature) ToFeatureOutputWithContext(ctx context.Context) FeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureOutput)
}

type FeatureOutput struct{ *pulumi.OutputState }

func (FeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Feature)(nil)).Elem()
}

func (o FeatureOutput) ToFeatureOutput() FeatureOutput {
	return o
}

func (o FeatureOutput) ToFeatureOutputWithContext(ctx context.Context) FeatureOutput {
	return o
}

// When the Feature resource was created.
func (o FeatureOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// When the Feature resource was deleted.
func (o FeatureOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// The ID of the feature to create.
func (o FeatureOutput) FeatureId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringPtrOutput { return v.FeatureId }).(pulumi.StringPtrOutput)
}

// Optional. Feature configuration applicable to all memberships of the fleet.
func (o FeatureOutput) FleetDefaultMemberConfig() CommonFleetDefaultMemberConfigSpecResponseOutput {
	return o.ApplyT(func(v *Feature) CommonFleetDefaultMemberConfigSpecResponseOutput { return v.FleetDefaultMemberConfig }).(CommonFleetDefaultMemberConfigSpecResponseOutput)
}

// Labels for this Feature.
func (o FeatureOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o FeatureOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
func (o FeatureOutput) MembershipSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.MembershipSpecs }).(pulumi.StringMapOutput)
}

// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
func (o FeatureOutput) MembershipStates() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.MembershipStates }).(pulumi.StringMapOutput)
}

// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
func (o FeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FeatureOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o FeatureOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// State of the Feature resource itself.
func (o FeatureOutput) ResourceState() FeatureResourceStateResponseOutput {
	return o.ApplyT(func(v *Feature) FeatureResourceStateResponseOutput { return v.ResourceState }).(FeatureResourceStateResponseOutput)
}

// Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
func (o FeatureOutput) ScopeSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.ScopeSpecs }).(pulumi.StringMapOutput)
}

// Scope-specific Feature status. If this Feature does report any per-Scope status, this field may be unused. The keys indicate which Scope the state is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project.
func (o FeatureOutput) ScopeStates() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.ScopeStates }).(pulumi.StringMapOutput)
}

// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
func (o FeatureOutput) Spec() CommonFeatureSpecResponseOutput {
	return o.ApplyT(func(v *Feature) CommonFeatureSpecResponseOutput { return v.Spec }).(CommonFeatureSpecResponseOutput)
}

// The Hub-wide Feature state.
func (o FeatureOutput) State() CommonFeatureStateResponseOutput {
	return o.ApplyT(func(v *Feature) CommonFeatureStateResponseOutput { return v.State }).(CommonFeatureStateResponseOutput)
}

// When the Feature resource was last updated.
func (o FeatureOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureInput)(nil)).Elem(), &Feature{})
	pulumi.RegisterOutputType(FeatureOutput{})
}
