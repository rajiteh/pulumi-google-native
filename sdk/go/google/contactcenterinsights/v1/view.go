// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a view.
type View struct {
	pulumi.CustomResourceState

	// The time at which this view was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The human-readable display name of the view.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The most recent time at which the view was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// String with specific view properties, must be non-empty.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		args = &ViewArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource View
	err := ctx.RegisterResource("google-native:contactcenterinsights/v1:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("google-native:contactcenterinsights/v1:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
}

type ViewState struct {
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	// The human-readable display name of the view.
	DisplayName *string `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	// Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// String with specific view properties, must be non-empty.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// The human-readable display name of the view.
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// String with specific view properties, must be non-empty.
	Value pulumi.StringPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

// The time at which this view was created.
func (o ViewOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The human-readable display name of the view.
func (o ViewOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ViewOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ViewOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The most recent time at which the view was updated.
func (o ViewOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// String with specific view properties, must be non-empty.
func (o ViewOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewInput)(nil)).Elem(), &View{})
	pulumi.RegisterOutputType(ViewOutput{})
}
