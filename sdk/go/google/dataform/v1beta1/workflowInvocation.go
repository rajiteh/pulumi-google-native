// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new WorkflowInvocation in a given Repository.
// Auto-naming is currently not supported for this resource.
type WorkflowInvocation struct {
	pulumi.CustomResourceState

	// Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
	CompilationResult pulumi.StringOutput `pulumi:"compilationResult"`
	// Immutable. If left unset, a default InvocationConfig will be used.
	InvocationConfig InvocationConfigResponseOutput `pulumi:"invocationConfig"`
	// This workflow invocation's timing details.
	InvocationTiming IntervalResponseOutput `pulumi:"invocationTiming"`
	Location         pulumi.StringOutput    `pulumi:"location"`
	// The workflow invocation's name.
	Name         pulumi.StringOutput `pulumi:"name"`
	Project      pulumi.StringOutput `pulumi:"project"`
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// This workflow invocation's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
	WorkflowConfig pulumi.StringOutput `pulumi:"workflowConfig"`
}

// NewWorkflowInvocation registers a new resource with the given unique name, arguments, and options.
func NewWorkflowInvocation(ctx *pulumi.Context,
	name string, args *WorkflowInvocationArgs, opts ...pulumi.ResourceOption) (*WorkflowInvocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"repositoryId",
	})
	opts = append(opts, replaceOnChanges)
	var resource WorkflowInvocation
	err := ctx.RegisterResource("google-native:dataform/v1beta1:WorkflowInvocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowInvocation gets an existing WorkflowInvocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowInvocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowInvocationState, opts ...pulumi.ResourceOption) (*WorkflowInvocation, error) {
	var resource WorkflowInvocation
	err := ctx.ReadResource("google-native:dataform/v1beta1:WorkflowInvocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowInvocation resources.
type workflowInvocationState struct {
}

type WorkflowInvocationState struct {
}

func (WorkflowInvocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowInvocationState)(nil)).Elem()
}

type workflowInvocationArgs struct {
	// Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
	CompilationResult *string `pulumi:"compilationResult"`
	// Immutable. If left unset, a default InvocationConfig will be used.
	InvocationConfig *InvocationConfig `pulumi:"invocationConfig"`
	Location         *string           `pulumi:"location"`
	Project          *string           `pulumi:"project"`
	RepositoryId     string            `pulumi:"repositoryId"`
	// Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
	WorkflowConfig *string `pulumi:"workflowConfig"`
}

// The set of arguments for constructing a WorkflowInvocation resource.
type WorkflowInvocationArgs struct {
	// Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
	CompilationResult pulumi.StringPtrInput
	// Immutable. If left unset, a default InvocationConfig will be used.
	InvocationConfig InvocationConfigPtrInput
	Location         pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	RepositoryId     pulumi.StringInput
	// Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
	WorkflowConfig pulumi.StringPtrInput
}

func (WorkflowInvocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowInvocationArgs)(nil)).Elem()
}

type WorkflowInvocationInput interface {
	pulumi.Input

	ToWorkflowInvocationOutput() WorkflowInvocationOutput
	ToWorkflowInvocationOutputWithContext(ctx context.Context) WorkflowInvocationOutput
}

func (*WorkflowInvocation) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowInvocation)(nil)).Elem()
}

func (i *WorkflowInvocation) ToWorkflowInvocationOutput() WorkflowInvocationOutput {
	return i.ToWorkflowInvocationOutputWithContext(context.Background())
}

func (i *WorkflowInvocation) ToWorkflowInvocationOutputWithContext(ctx context.Context) WorkflowInvocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowInvocationOutput)
}

type WorkflowInvocationOutput struct{ *pulumi.OutputState }

func (WorkflowInvocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowInvocation)(nil)).Elem()
}

func (o WorkflowInvocationOutput) ToWorkflowInvocationOutput() WorkflowInvocationOutput {
	return o
}

func (o WorkflowInvocationOutput) ToWorkflowInvocationOutputWithContext(ctx context.Context) WorkflowInvocationOutput {
	return o
}

// Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
func (o WorkflowInvocationOutput) CompilationResult() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.CompilationResult }).(pulumi.StringOutput)
}

// Immutable. If left unset, a default InvocationConfig will be used.
func (o WorkflowInvocationOutput) InvocationConfig() InvocationConfigResponseOutput {
	return o.ApplyT(func(v *WorkflowInvocation) InvocationConfigResponseOutput { return v.InvocationConfig }).(InvocationConfigResponseOutput)
}

// This workflow invocation's timing details.
func (o WorkflowInvocationOutput) InvocationTiming() IntervalResponseOutput {
	return o.ApplyT(func(v *WorkflowInvocation) IntervalResponseOutput { return v.InvocationTiming }).(IntervalResponseOutput)
}

func (o WorkflowInvocationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The workflow invocation's name.
func (o WorkflowInvocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkflowInvocationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o WorkflowInvocationOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// This workflow invocation's current state.
func (o WorkflowInvocationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
func (o WorkflowInvocationOutput) WorkflowConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowInvocation) pulumi.StringOutput { return v.WorkflowConfig }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowInvocationInput)(nil)).Elem(), &WorkflowInvocation{})
	pulumi.RegisterOutputType(WorkflowInvocationOutput{})
}
