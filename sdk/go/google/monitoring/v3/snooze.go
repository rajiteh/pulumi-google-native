// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Snooze that will prevent alerts, which match the provided criteria, from being opened. The Snooze applies for a specific time interval.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Snooze struct {
	pulumi.CustomResourceState

	// This defines the criteria for applying the Snooze. See Criteria for more information.
	Criteria CriteriaResponseOutput `pulumi:"criteria"`
	// A display name for the Snooze. This can be, at most, 512 unicode characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The Snooze will be active from interval.start_time through interval.end_time. interval.start_time cannot be in the past. There is a 15 second clock skew to account for the time it takes for a request to reach the API from the UI.
	Interval TimeIntervalResponseOutput `pulumi:"interval"`
	// The name of the Snooze. The format is: projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID] The ID of the Snooze will be generated by the system.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewSnooze registers a new resource with the given unique name, arguments, and options.
func NewSnooze(ctx *pulumi.Context,
	name string, args *SnoozeArgs, opts ...pulumi.ResourceOption) (*Snooze, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Interval == nil {
		return nil, errors.New("invalid value for required argument 'Interval'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Snooze
	err := ctx.RegisterResource("google-native:monitoring/v3:Snooze", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnooze gets an existing Snooze resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnooze(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnoozeState, opts ...pulumi.ResourceOption) (*Snooze, error) {
	var resource Snooze
	err := ctx.ReadResource("google-native:monitoring/v3:Snooze", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snooze resources.
type snoozeState struct {
}

type SnoozeState struct {
}

func (SnoozeState) ElementType() reflect.Type {
	return reflect.TypeOf((*snoozeState)(nil)).Elem()
}

type snoozeArgs struct {
	// This defines the criteria for applying the Snooze. See Criteria for more information.
	Criteria Criteria `pulumi:"criteria"`
	// A display name for the Snooze. This can be, at most, 512 unicode characters.
	DisplayName string `pulumi:"displayName"`
	// The Snooze will be active from interval.start_time through interval.end_time. interval.start_time cannot be in the past. There is a 15 second clock skew to account for the time it takes for a request to reach the API from the UI.
	Interval TimeInterval `pulumi:"interval"`
	// The name of the Snooze. The format is: projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID] The ID of the Snooze will be generated by the system.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Snooze resource.
type SnoozeArgs struct {
	// This defines the criteria for applying the Snooze. See Criteria for more information.
	Criteria CriteriaInput
	// A display name for the Snooze. This can be, at most, 512 unicode characters.
	DisplayName pulumi.StringInput
	// The Snooze will be active from interval.start_time through interval.end_time. interval.start_time cannot be in the past. There is a 15 second clock skew to account for the time it takes for a request to reach the API from the UI.
	Interval TimeIntervalInput
	// The name of the Snooze. The format is: projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID] The ID of the Snooze will be generated by the system.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (SnoozeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snoozeArgs)(nil)).Elem()
}

type SnoozeInput interface {
	pulumi.Input

	ToSnoozeOutput() SnoozeOutput
	ToSnoozeOutputWithContext(ctx context.Context) SnoozeOutput
}

func (*Snooze) ElementType() reflect.Type {
	return reflect.TypeOf((**Snooze)(nil)).Elem()
}

func (i *Snooze) ToSnoozeOutput() SnoozeOutput {
	return i.ToSnoozeOutputWithContext(context.Background())
}

func (i *Snooze) ToSnoozeOutputWithContext(ctx context.Context) SnoozeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnoozeOutput)
}

type SnoozeOutput struct{ *pulumi.OutputState }

func (SnoozeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snooze)(nil)).Elem()
}

func (o SnoozeOutput) ToSnoozeOutput() SnoozeOutput {
	return o
}

func (o SnoozeOutput) ToSnoozeOutputWithContext(ctx context.Context) SnoozeOutput {
	return o
}

// This defines the criteria for applying the Snooze. See Criteria for more information.
func (o SnoozeOutput) Criteria() CriteriaResponseOutput {
	return o.ApplyT(func(v *Snooze) CriteriaResponseOutput { return v.Criteria }).(CriteriaResponseOutput)
}

// A display name for the Snooze. This can be, at most, 512 unicode characters.
func (o SnoozeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snooze) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The Snooze will be active from interval.start_time through interval.end_time. interval.start_time cannot be in the past. There is a 15 second clock skew to account for the time it takes for a request to reach the API from the UI.
func (o SnoozeOutput) Interval() TimeIntervalResponseOutput {
	return o.ApplyT(func(v *Snooze) TimeIntervalResponseOutput { return v.Interval }).(TimeIntervalResponseOutput)
}

// The name of the Snooze. The format is: projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID] The ID of the Snooze will be generated by the system.
func (o SnoozeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snooze) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SnoozeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Snooze) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnoozeInput)(nil)).Elem(), &Snooze{})
	pulumi.RegisterOutputType(SnoozeOutput{})
}
