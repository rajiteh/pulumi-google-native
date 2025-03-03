// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new domain mapping.
// Auto-naming is currently not supported for this resource.
type DomainMapping struct {
	pulumi.CustomResourceState

	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun pulumi.StringPtrOutput `pulumi:"dryRun"`
	// The kind of resource, in this case "DomainMapping".
	Kind     pulumi.StringOutput `pulumi:"kind"`
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this BuildTemplate.
	Metadata ObjectMetaResponseOutput `pulumi:"metadata"`
	Project  pulumi.StringOutput      `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec DomainMappingSpecResponseOutput `pulumi:"spec"`
	// The current status of the DomainMapping.
	Status DomainMappingStatusResponseOutput `pulumi:"status"`
}

// NewDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewDomainMapping(ctx *pulumi.Context,
	name string, args *DomainMappingArgs, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	if args == nil {
		args = &DomainMappingArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource DomainMapping
	err := ctx.RegisterResource("google-native:run/v1:DomainMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainMapping gets an existing DomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainMappingState, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	var resource DomainMapping
	err := ctx.ReadResource("google-native:run/v1:DomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainMapping resources.
type domainMappingState struct {
}

type DomainMappingState struct {
}

func (DomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingState)(nil)).Elem()
}

type domainMappingArgs struct {
	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion *string `pulumi:"apiVersion"`
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun *string `pulumi:"dryRun"`
	// The kind of resource, in this case "DomainMapping".
	Kind     *string `pulumi:"kind"`
	Location *string `pulumi:"location"`
	// Metadata associated with this BuildTemplate.
	Metadata *ObjectMeta `pulumi:"metadata"`
	Project  *string     `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec *DomainMappingSpec `pulumi:"spec"`
}

// The set of arguments for constructing a DomainMapping resource.
type DomainMappingArgs struct {
	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion pulumi.StringPtrInput
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun pulumi.StringPtrInput
	// The kind of resource, in this case "DomainMapping".
	Kind     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// Metadata associated with this BuildTemplate.
	Metadata ObjectMetaPtrInput
	Project  pulumi.StringPtrInput
	// The spec for this DomainMapping.
	Spec DomainMappingSpecPtrInput
}

func (DomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingArgs)(nil)).Elem()
}

type DomainMappingInput interface {
	pulumi.Input

	ToDomainMappingOutput() DomainMappingOutput
	ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput
}

func (*DomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil)).Elem()
}

func (i *DomainMapping) ToDomainMappingOutput() DomainMappingOutput {
	return i.ToDomainMappingOutputWithContext(context.Background())
}

func (i *DomainMapping) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput)
}

type DomainMappingOutput struct{ *pulumi.OutputState }

func (DomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainMapping)(nil)).Elem()
}

func (o DomainMappingOutput) ToDomainMappingOutput() DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return o
}

// The API version for this call such as "domains.cloudrun.com/v1".
func (o DomainMappingOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
func (o DomainMappingOutput) DryRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringPtrOutput { return v.DryRun }).(pulumi.StringPtrOutput)
}

// The kind of resource, in this case "DomainMapping".
func (o DomainMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DomainMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metadata associated with this BuildTemplate.
func (o DomainMappingOutput) Metadata() ObjectMetaResponseOutput {
	return o.ApplyT(func(v *DomainMapping) ObjectMetaResponseOutput { return v.Metadata }).(ObjectMetaResponseOutput)
}

func (o DomainMappingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainMapping) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The spec for this DomainMapping.
func (o DomainMappingOutput) Spec() DomainMappingSpecResponseOutput {
	return o.ApplyT(func(v *DomainMapping) DomainMappingSpecResponseOutput { return v.Spec }).(DomainMappingSpecResponseOutput)
}

// The current status of the DomainMapping.
func (o DomainMappingOutput) Status() DomainMappingStatusResponseOutput {
	return o.ApplyT(func(v *DomainMapping) DomainMappingStatusResponseOutput { return v.Status }).(DomainMappingStatusResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMappingInput)(nil)).Elem(), &DomainMapping{})
	pulumi.RegisterOutputType(DomainMappingOutput{})
}
