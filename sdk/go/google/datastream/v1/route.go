// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to create a route for a private connectivity configuration in a project and location.
// Auto-naming is currently not supported for this resource.
type Route struct {
	pulumi.CustomResourceState

	// The create time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Destination address for connection
	DestinationAddress pulumi.StringOutput `pulumi:"destinationAddress"`
	// Destination port for connection
	DestinationPort pulumi.IntOutput `pulumi:"destinationPort"`
	// Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Labels.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The resource's name.
	Name                pulumi.StringOutput `pulumi:"name"`
	PrivateConnectionId pulumi.StringOutput `pulumi:"privateConnectionId"`
	Project             pulumi.StringOutput `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Required. The Route identifier.
	RouteId pulumi.StringOutput `pulumi:"routeId"`
	// The update time of the resource.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationAddress == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAddress'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.PrivateConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateConnectionId'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"privateConnectionId",
		"project",
		"routeId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Route
	err := ctx.RegisterResource("google-native:datastream/v1:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("google-native:datastream/v1:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Destination address for connection
	DestinationAddress string `pulumi:"destinationAddress"`
	// Destination port for connection
	DestinationPort *int `pulumi:"destinationPort"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Labels.
	Labels              map[string]string `pulumi:"labels"`
	Location            *string           `pulumi:"location"`
	PrivateConnectionId string            `pulumi:"privateConnectionId"`
	Project             *string           `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Required. The Route identifier.
	RouteId string `pulumi:"routeId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Destination address for connection
	DestinationAddress pulumi.StringInput
	// Destination port for connection
	DestinationPort pulumi.IntPtrInput
	// Display name.
	DisplayName pulumi.StringInput
	// Labels.
	Labels              pulumi.StringMapInput
	Location            pulumi.StringPtrInput
	PrivateConnectionId pulumi.StringInput
	Project             pulumi.StringPtrInput
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Required. The Route identifier.
	RouteId pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

// The create time of the resource.
func (o RouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Destination address for connection
func (o RouteOutput) DestinationAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DestinationAddress }).(pulumi.StringOutput)
}

// Destination port for connection
func (o RouteOutput) DestinationPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Route) pulumi.IntOutput { return v.DestinationPort }).(pulumi.IntOutput)
}

// Display name.
func (o RouteOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Labels.
func (o RouteOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Route) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o RouteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource's name.
func (o RouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RouteOutput) PrivateConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.PrivateConnectionId }).(pulumi.StringOutput)
}

func (o RouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o RouteOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Required. The Route identifier.
func (o RouteOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

// The update time of the resource.
func (o RouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
