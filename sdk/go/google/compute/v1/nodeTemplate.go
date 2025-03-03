// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a NodeTemplate resource in the specified project using the data included in the request.
type NodeTemplate struct {
	pulumi.CustomResourceState

	Accelerators AcceleratorConfigResponseArrayOutput `pulumi:"accelerators"`
	// CPU overcommit.
	CpuOvercommitType pulumi.StringOutput `pulumi:"cpuOvercommitType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput          `pulumi:"description"`
	Disks       LocalDiskResponseArrayOutput `pulumi:"disks"`
	// The type of the resource. Always compute#nodeTemplate for node templates.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapOutput `pulumi:"nodeAffinityLabels"`
	// The node type to use for nodes group that are created from this template.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// Do not use. Instead, use the node_type property.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityResponseOutput `pulumi:"nodeTypeFlexibility"`
	Project             pulumi.StringOutput                           `pulumi:"project"`
	Region              pulumi.StringOutput                           `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
	ServerBinding ServerBindingResponseOutput `pulumi:"serverBinding"`
	// The status of the node template. One of the following values: CREATING, READY, and DELETING.
	Status pulumi.StringOutput `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
}

// NewNodeTemplate registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplate(ctx *pulumi.Context,
	name string, args *NodeTemplateArgs, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"region",
	})
	opts = append(opts, replaceOnChanges)
	var resource NodeTemplate
	err := ctx.RegisterResource("google-native:compute/v1:NodeTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeTemplate gets an existing NodeTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTemplateState, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	var resource NodeTemplate
	err := ctx.ReadResource("google-native:compute/v1:NodeTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeTemplate resources.
type nodeTemplateState struct {
}

type NodeTemplateState struct {
}

func (NodeTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateState)(nil)).Elem()
}

type nodeTemplateArgs struct {
	Accelerators []AcceleratorConfig `pulumi:"accelerators"`
	// CPU overcommit.
	CpuOvercommitType *NodeTemplateCpuOvercommitType `pulumi:"cpuOvercommitType"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string     `pulumi:"description"`
	Disks       []LocalDisk `pulumi:"disks"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// The node type to use for nodes group that are created from this template.
	NodeType *string `pulumi:"nodeType"`
	// Do not use. Instead, use the node_type property.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	Project             *string                          `pulumi:"project"`
	Region              string                           `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
	ServerBinding *ServerBinding `pulumi:"serverBinding"`
}

// The set of arguments for constructing a NodeTemplate resource.
type NodeTemplateArgs struct {
	Accelerators AcceleratorConfigArrayInput
	// CPU overcommit.
	CpuOvercommitType NodeTemplateCpuOvercommitTypePtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	Disks       LocalDiskArrayInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// The node type to use for nodes group that are created from this template.
	NodeType pulumi.StringPtrInput
	// Do not use. Instead, use the node_type property.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	Project             pulumi.StringPtrInput
	Region              pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
	ServerBinding ServerBindingPtrInput
}

func (NodeTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateArgs)(nil)).Elem()
}

type NodeTemplateInput interface {
	pulumi.Input

	ToNodeTemplateOutput() NodeTemplateOutput
	ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput
}

func (*NodeTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTemplate)(nil)).Elem()
}

func (i *NodeTemplate) ToNodeTemplateOutput() NodeTemplateOutput {
	return i.ToNodeTemplateOutputWithContext(context.Background())
}

func (i *NodeTemplate) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateOutput)
}

type NodeTemplateOutput struct{ *pulumi.OutputState }

func (NodeTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTemplate)(nil)).Elem()
}

func (o NodeTemplateOutput) ToNodeTemplateOutput() NodeTemplateOutput {
	return o
}

func (o NodeTemplateOutput) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return o
}

func (o NodeTemplateOutput) Accelerators() AcceleratorConfigResponseArrayOutput {
	return o.ApplyT(func(v *NodeTemplate) AcceleratorConfigResponseArrayOutput { return v.Accelerators }).(AcceleratorConfigResponseArrayOutput)
}

// CPU overcommit.
func (o NodeTemplateOutput) CpuOvercommitType() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.CpuOvercommitType }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o NodeTemplateOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o NodeTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o NodeTemplateOutput) Disks() LocalDiskResponseArrayOutput {
	return o.ApplyT(func(v *NodeTemplate) LocalDiskResponseArrayOutput { return v.Disks }).(LocalDiskResponseArrayOutput)
}

// The type of the resource. Always compute#nodeTemplate for node templates.
func (o NodeTemplateOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o NodeTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Labels to use for node affinity, which will be used in instance scheduling.
func (o NodeTemplateOutput) NodeAffinityLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringMapOutput { return v.NodeAffinityLabels }).(pulumi.StringMapOutput)
}

// The node type to use for nodes group that are created from this template.
func (o NodeTemplateOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// Do not use. Instead, use the node_type property.
func (o NodeTemplateOutput) NodeTypeFlexibility() NodeTemplateNodeTypeFlexibilityResponseOutput {
	return o.ApplyT(func(v *NodeTemplate) NodeTemplateNodeTypeFlexibilityResponseOutput { return v.NodeTypeFlexibility }).(NodeTemplateNodeTypeFlexibilityResponseOutput)
}

func (o NodeTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NodeTemplateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o NodeTemplateOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o NodeTemplateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
func (o NodeTemplateOutput) ServerBinding() ServerBindingResponseOutput {
	return o.ApplyT(func(v *NodeTemplate) ServerBindingResponseOutput { return v.ServerBinding }).(ServerBindingResponseOutput)
}

// The status of the node template. One of the following values: CREATING, READY, and DELETING.
func (o NodeTemplateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// An optional, human-readable explanation of the status.
func (o NodeTemplateOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeTemplateInput)(nil)).Elem(), &NodeTemplate{})
	pulumi.RegisterOutputType(NodeTemplateOutput{})
}
