// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new execution using the latest revision of the given workflow.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Execution struct {
	pulumi.CustomResourceState

	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument pulumi.StringOutput `pulumi:"argument"`
	// The call logging level associated to this execution.
	CallLogLevel pulumi.StringOutput `pulumi:"callLogLevel"`
	// Measures the duration of the execution.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// Marks the end of execution, successful or not.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
	Error ErrorResponseOutput `pulumi:"error"`
	// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
	Result pulumi.StringOutput `pulumi:"result"`
	// Marks the beginning of execution.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Current state of the execution.
	State pulumi.StringOutput `pulumi:"state"`
	// Error regarding the state of the Execution resource. For example, this field will have error details if the Execution data is unavailable due to revoked KMS key permissions.
	StateError StateErrorResponseOutput `pulumi:"stateError"`
	// Status tracks the current steps and progress data of this execution.
	Status     StatusResponseOutput `pulumi:"status"`
	WorkflowId pulumi.StringOutput  `pulumi:"workflowId"`
	// Revision of the workflow this execution is using.
	WorkflowRevisionId pulumi.StringOutput `pulumi:"workflowRevisionId"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkflowId == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"workflowId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Execution
	err := ctx.RegisterResource("google-native:workflowexecutions/v1:Execution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecution gets an existing Execution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionState, opts ...pulumi.ResourceOption) (*Execution, error) {
	var resource Execution
	err := ctx.ReadResource("google-native:workflowexecutions/v1:Execution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Execution resources.
type executionState struct {
}

type ExecutionState struct {
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument *string `pulumi:"argument"`
	// The call logging level associated to this execution.
	CallLogLevel *ExecutionCallLogLevel `pulumi:"callLogLevel"`
	// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
	Labels     map[string]string `pulumi:"labels"`
	Location   *string           `pulumi:"location"`
	Project    *string           `pulumi:"project"`
	WorkflowId string            `pulumi:"workflowId"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument pulumi.StringPtrInput
	// The call logging level associated to this execution.
	CallLogLevel ExecutionCallLogLevelPtrInput
	// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
	Labels     pulumi.StringMapInput
	Location   pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	WorkflowId pulumi.StringInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}

type ExecutionInput interface {
	pulumi.Input

	ToExecutionOutput() ExecutionOutput
	ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput
}

func (*Execution) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (i *Execution) ToExecutionOutput() ExecutionOutput {
	return i.ToExecutionOutputWithContext(context.Background())
}

func (i *Execution) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionOutput)
}

type ExecutionOutput struct{ *pulumi.OutputState }

func (ExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (o ExecutionOutput) ToExecutionOutput() ExecutionOutput {
	return o
}

func (o ExecutionOutput) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return o
}

// Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
func (o ExecutionOutput) Argument() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Argument }).(pulumi.StringOutput)
}

// The call logging level associated to this execution.
func (o ExecutionOutput) CallLogLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.CallLogLevel }).(pulumi.StringOutput)
}

// Measures the duration of the execution.
func (o ExecutionOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

// Marks the end of execution, successful or not.
func (o ExecutionOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
func (o ExecutionOutput) Error() ErrorResponseOutput {
	return o.ApplyT(func(v *Execution) ErrorResponseOutput { return v.Error }).(ErrorResponseOutput)
}

// Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
func (o ExecutionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o ExecutionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
func (o ExecutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExecutionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
func (o ExecutionOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

// Marks the beginning of execution.
func (o ExecutionOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Current state of the execution.
func (o ExecutionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Error regarding the state of the Execution resource. For example, this field will have error details if the Execution data is unavailable due to revoked KMS key permissions.
func (o ExecutionOutput) StateError() StateErrorResponseOutput {
	return o.ApplyT(func(v *Execution) StateErrorResponseOutput { return v.StateError }).(StateErrorResponseOutput)
}

// Status tracks the current steps and progress data of this execution.
func (o ExecutionOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v *Execution) StatusResponseOutput { return v.Status }).(StatusResponseOutput)
}

func (o ExecutionOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

// Revision of the workflow this execution is using.
func (o ExecutionOutput) WorkflowRevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.WorkflowRevisionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionInput)(nil)).Elem(), &Execution{})
	pulumi.RegisterOutputType(ExecutionOutput{})
}
