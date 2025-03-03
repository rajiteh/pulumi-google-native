// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new entity row of the specified entity type in the external system. The field values for creating the row are contained in the body of the request. The response message contains a `Entity` message object returned as a response by the external system.
// Auto-naming is currently not supported for this resource.
type Entity struct {
	pulumi.CustomResourceState

	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	EntityTypeId pulumi.StringOutput `pulumi:"entityTypeId"`
	// Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
	Fields   pulumi.StringMapOutput `pulumi:"fields"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Resource name of the Entity. Format: projects/{project}/locations/{location}/connections/{connection}/entityTypes/{type}/entities/{id}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewEntity registers a new resource with the given unique name, arguments, and options.
func NewEntity(ctx *pulumi.Context,
	name string, args *EntityArgs, opts ...pulumi.ResourceOption) (*Entity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.EntityTypeId == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"connectionId",
		"entityTypeId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Entity
	err := ctx.RegisterResource("google-native:connectors/v2:Entity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntity gets an existing Entity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityState, opts ...pulumi.ResourceOption) (*Entity, error) {
	var resource Entity
	err := ctx.ReadResource("google-native:connectors/v2:Entity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entity resources.
type entityState struct {
}

type EntityState struct {
}

func (EntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityState)(nil)).Elem()
}

type entityArgs struct {
	ConnectionId string `pulumi:"connectionId"`
	EntityTypeId string `pulumi:"entityTypeId"`
	// Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
	Fields   map[string]string `pulumi:"fields"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
}

// The set of arguments for constructing a Entity resource.
type EntityArgs struct {
	ConnectionId pulumi.StringInput
	EntityTypeId pulumi.StringInput
	// Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
	Fields   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
}

func (EntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityArgs)(nil)).Elem()
}

type EntityInput interface {
	pulumi.Input

	ToEntityOutput() EntityOutput
	ToEntityOutputWithContext(ctx context.Context) EntityOutput
}

func (*Entity) ElementType() reflect.Type {
	return reflect.TypeOf((**Entity)(nil)).Elem()
}

func (i *Entity) ToEntityOutput() EntityOutput {
	return i.ToEntityOutputWithContext(context.Background())
}

func (i *Entity) ToEntityOutputWithContext(ctx context.Context) EntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityOutput)
}

type EntityOutput struct{ *pulumi.OutputState }

func (EntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Entity)(nil)).Elem()
}

func (o EntityOutput) ToEntityOutput() EntityOutput {
	return o
}

func (o EntityOutput) ToEntityOutputWithContext(ctx context.Context) EntityOutput {
	return o
}

func (o EntityOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

func (o EntityOutput) EntityTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringOutput { return v.EntityTypeId }).(pulumi.StringOutput)
}

// Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
func (o EntityOutput) Fields() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringMapOutput { return v.Fields }).(pulumi.StringMapOutput)
}

func (o EntityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name of the Entity. Format: projects/{project}/locations/{location}/connections/{connection}/entityTypes/{type}/entities/{id}
func (o EntityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EntityOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Entity) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInput)(nil)).Elem(), &Entity{})
	pulumi.RegisterOutputType(EntityOutput{})
}
