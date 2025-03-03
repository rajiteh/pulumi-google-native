// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new service in the specified project using the data included in the request.
type NetworkEdgeSecurityService struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#networkEdgeSecurityService for NetworkEdgeSecurityServices
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicy pulumi.StringOutput `pulumi:"securityPolicy"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
}

// NewNetworkEdgeSecurityService registers a new resource with the given unique name, arguments, and options.
func NewNetworkEdgeSecurityService(ctx *pulumi.Context,
	name string, args *NetworkEdgeSecurityServiceArgs, opts ...pulumi.ResourceOption) (*NetworkEdgeSecurityService, error) {
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
	var resource NetworkEdgeSecurityService
	err := ctx.RegisterResource("google-native:compute/beta:NetworkEdgeSecurityService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkEdgeSecurityService gets an existing NetworkEdgeSecurityService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkEdgeSecurityService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkEdgeSecurityServiceState, opts ...pulumi.ResourceOption) (*NetworkEdgeSecurityService, error) {
	var resource NetworkEdgeSecurityService
	err := ctx.ReadResource("google-native:compute/beta:NetworkEdgeSecurityService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkEdgeSecurityService resources.
type networkEdgeSecurityServiceState struct {
}

type NetworkEdgeSecurityServiceState struct {
}

func (NetworkEdgeSecurityServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEdgeSecurityServiceState)(nil)).Elem()
}

type networkEdgeSecurityServiceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicy *string `pulumi:"securityPolicy"`
}

// The set of arguments for constructing a NetworkEdgeSecurityService resource.
type NetworkEdgeSecurityServiceArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicy pulumi.StringPtrInput
}

func (NetworkEdgeSecurityServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEdgeSecurityServiceArgs)(nil)).Elem()
}

type NetworkEdgeSecurityServiceInput interface {
	pulumi.Input

	ToNetworkEdgeSecurityServiceOutput() NetworkEdgeSecurityServiceOutput
	ToNetworkEdgeSecurityServiceOutputWithContext(ctx context.Context) NetworkEdgeSecurityServiceOutput
}

func (*NetworkEdgeSecurityService) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEdgeSecurityService)(nil)).Elem()
}

func (i *NetworkEdgeSecurityService) ToNetworkEdgeSecurityServiceOutput() NetworkEdgeSecurityServiceOutput {
	return i.ToNetworkEdgeSecurityServiceOutputWithContext(context.Background())
}

func (i *NetworkEdgeSecurityService) ToNetworkEdgeSecurityServiceOutputWithContext(ctx context.Context) NetworkEdgeSecurityServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEdgeSecurityServiceOutput)
}

type NetworkEdgeSecurityServiceOutput struct{ *pulumi.OutputState }

func (NetworkEdgeSecurityServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkEdgeSecurityService)(nil)).Elem()
}

func (o NetworkEdgeSecurityServiceOutput) ToNetworkEdgeSecurityServiceOutput() NetworkEdgeSecurityServiceOutput {
	return o
}

func (o NetworkEdgeSecurityServiceOutput) ToNetworkEdgeSecurityServiceOutputWithContext(ctx context.Context) NetworkEdgeSecurityServiceOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o NetworkEdgeSecurityServiceOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o NetworkEdgeSecurityServiceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
func (o NetworkEdgeSecurityServiceOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#networkEdgeSecurityService for NetworkEdgeSecurityServices
func (o NetworkEdgeSecurityServiceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o NetworkEdgeSecurityServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkEdgeSecurityServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NetworkEdgeSecurityServiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o NetworkEdgeSecurityServiceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The resource URL for the network edge security service associated with this network edge security service.
func (o NetworkEdgeSecurityServiceOutput) SecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.SecurityPolicy }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o NetworkEdgeSecurityServiceOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o NetworkEdgeSecurityServiceOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkEdgeSecurityService) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkEdgeSecurityServiceInput)(nil)).Elem(), &NetworkEdgeSecurityService{})
	pulumi.RegisterOutputType(NetworkEdgeSecurityServiceOutput{})
}
