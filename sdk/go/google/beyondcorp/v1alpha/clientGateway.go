// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ClientGateway in a given project and location.
type ClientGateway struct {
	pulumi.CustomResourceState

	// The client connector service name that the client gateway is associated to. Client Connector Services, named as follows: `projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
	ClientConnectorService pulumi.StringOutput `pulumi:"clientConnectorService"`
	// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ClientGatewayId pulumi.StringPtrOutput `pulumi:"clientGatewayId"`
	// [Output only] Create time stamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	Location   pulumi.StringOutput `pulumi:"location"`
	// name of resource. The name is ignored during creation.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The operational state of the gateway.
	State pulumi.StringOutput `pulumi:"state"`
	// [Output only] Update time stamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewClientGateway registers a new resource with the given unique name, arguments, and options.
func NewClientGateway(ctx *pulumi.Context,
	name string, args *ClientGatewayArgs, opts ...pulumi.ResourceOption) (*ClientGateway, error) {
	if args == nil {
		args = &ClientGatewayArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource ClientGateway
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:ClientGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientGateway gets an existing ClientGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientGatewayState, opts ...pulumi.ResourceOption) (*ClientGateway, error) {
	var resource ClientGateway
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:ClientGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientGateway resources.
type clientGatewayState struct {
}

type ClientGatewayState struct {
}

func (ClientGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGatewayState)(nil)).Elem()
}

type clientGatewayArgs struct {
	// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ClientGatewayId *string `pulumi:"clientGatewayId"`
	Location        *string `pulumi:"location"`
	// name of resource. The name is ignored during creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a ClientGateway resource.
type ClientGatewayArgs struct {
	// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
	ClientGatewayId pulumi.StringPtrInput
	Location        pulumi.StringPtrInput
	// name of resource. The name is ignored during creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (ClientGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGatewayArgs)(nil)).Elem()
}

type ClientGatewayInput interface {
	pulumi.Input

	ToClientGatewayOutput() ClientGatewayOutput
	ToClientGatewayOutputWithContext(ctx context.Context) ClientGatewayOutput
}

func (*ClientGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGateway)(nil)).Elem()
}

func (i *ClientGateway) ToClientGatewayOutput() ClientGatewayOutput {
	return i.ToClientGatewayOutputWithContext(context.Background())
}

func (i *ClientGateway) ToClientGatewayOutputWithContext(ctx context.Context) ClientGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGatewayOutput)
}

type ClientGatewayOutput struct{ *pulumi.OutputState }

func (ClientGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGateway)(nil)).Elem()
}

func (o ClientGatewayOutput) ToClientGatewayOutput() ClientGatewayOutput {
	return o
}

func (o ClientGatewayOutput) ToClientGatewayOutputWithContext(ctx context.Context) ClientGatewayOutput {
	return o
}

// The client connector service name that the client gateway is associated to. Client Connector Services, named as follows: `projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
func (o ClientGatewayOutput) ClientConnectorService() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.ClientConnectorService }).(pulumi.StringOutput)
}

// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
func (o ClientGatewayOutput) ClientGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringPtrOutput { return v.ClientGatewayId }).(pulumi.StringPtrOutput)
}

// [Output only] Create time stamp.
func (o ClientGatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o ClientGatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// name of resource. The name is ignored during creation.
func (o ClientGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o ClientGatewayOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The operational state of the gateway.
func (o ClientGatewayOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// [Output only] Update time stamp.
func (o ClientGatewayOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientGateway) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientGatewayInput)(nil)).Elem(), &ClientGateway{})
	pulumi.RegisterOutputType(ClientGatewayOutput{})
}
