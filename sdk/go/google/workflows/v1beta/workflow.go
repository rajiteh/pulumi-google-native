// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new workflow. If a workflow with the specified name already exists in the specified project and location, the long running operation will return ALREADY_EXISTS error.
type Workflow struct {
	pulumi.CustomResourceState

	// The timestamp of when the workflow was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description pulumi.StringOutput `pulumi:"description"`
	// Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores and dashes. Label keys must start with a letter. International characters are allowed.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The resource name of the workflow. Format: projects/{project}/locations/{location}/workflows/{workflow}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The timestamp that the latest revision of the workflow was created.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// The revision of the workflow. A new revision of a workflow is created as a result of updating the following properties of a workflow: - Service account - Workflow code to be executed The format is "000001-a4d", where the first 6 characters define the zero-padded revision ordinal number. They are followed by a hyphen and 3 hexadecimal random characters.
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. Format: projects/{project}/serviceAccounts/{account} or {account} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. The `{account}` value can be the `email` address or the `unique_id` of the service account. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Workflow code to be executed. The size limit is 128KB.
	SourceContents pulumi.StringOutput `pulumi:"sourceContents"`
	// State of the workflow deployment.
	State pulumi.StringOutput `pulumi:"state"`
	// The last update timestamp of the workflow.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. The ID of the workflow to be created. It has to fulfill the following requirements: * Must contain only letters, numbers, underscores and hyphens. * Must start with a letter. * Must be between 1-64 characters. * Must end with a number or a letter. * Must be unique within the customer project and location.
	WorkflowId pulumi.StringOutput `pulumi:"workflowId"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
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
	var resource Workflow
	err := ctx.RegisterResource("google-native:workflows/v1beta:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("google-native:workflows/v1beta:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description *string `pulumi:"description"`
	// Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores and dashes. Label keys must start with a letter. International characters are allowed.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The resource name of the workflow. Format: projects/{project}/locations/{location}/workflows/{workflow}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. Format: projects/{project}/serviceAccounts/{account} or {account} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. The `{account}` value can be the `email` address or the `unique_id` of the service account. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Workflow code to be executed. The size limit is 128KB.
	SourceContents *string `pulumi:"sourceContents"`
	// Required. The ID of the workflow to be created. It has to fulfill the following requirements: * Must contain only letters, numbers, underscores and hyphens. * Must start with a letter. * Must be between 1-64 characters. * Must end with a number or a letter. * Must be unique within the customer project and location.
	WorkflowId string `pulumi:"workflowId"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description pulumi.StringPtrInput
	// Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores and dashes. Label keys must start with a letter. International characters are allowed.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The resource name of the workflow. Format: projects/{project}/locations/{location}/workflows/{workflow}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. Format: projects/{project}/serviceAccounts/{account} or {account} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. The `{account}` value can be the `email` address or the `unique_id` of the service account. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount pulumi.StringPtrInput
	// Workflow code to be executed. The size limit is 128KB.
	SourceContents pulumi.StringPtrInput
	// Required. The ID of the workflow to be created. It has to fulfill the following requirements: * Must contain only letters, numbers, underscores and hyphens. * Must start with a letter. * Must be between 1-64 characters. * Must end with a number or a letter. * Must be unique within the customer project and location.
	WorkflowId pulumi.StringInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

// The timestamp of when the workflow was created.
func (o WorkflowOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
func (o WorkflowOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores and dashes. Label keys must start with a letter. International characters are allowed.
func (o WorkflowOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o WorkflowOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the workflow. Format: projects/{project}/locations/{location}/workflows/{workflow}
func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The timestamp that the latest revision of the workflow was created.
func (o WorkflowOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// The revision of the workflow. A new revision of a workflow is created as a result of updating the following properties of a workflow: - Service account - Workflow code to be executed The format is "000001-a4d", where the first 6 characters define the zero-padded revision ordinal number. They are followed by a hyphen and 3 hexadecimal random characters.
func (o WorkflowOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. Format: projects/{project}/serviceAccounts/{account} or {account} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. The `{account}` value can be the `email` address or the `unique_id` of the service account. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
func (o WorkflowOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Workflow code to be executed. The size limit is 128KB.
func (o WorkflowOutput) SourceContents() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.SourceContents }).(pulumi.StringOutput)
}

// State of the workflow deployment.
func (o WorkflowOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The last update timestamp of the workflow.
func (o WorkflowOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Required. The ID of the workflow to be created. It has to fulfill the following requirements: * Must contain only letters, numbers, underscores and hyphens. * Must start with a letter. * Must be between 1-64 characters. * Must end with a number or a letter. * Must be unique within the customer project and location.
func (o WorkflowOutput) WorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.WorkflowId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowInput)(nil)).Elem(), &Workflow{})
	pulumi.RegisterOutputType(WorkflowOutput{})
}
