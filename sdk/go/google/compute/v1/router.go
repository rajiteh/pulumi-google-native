// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Router resource in the specified project and region using the data included in the request.
type Router struct {
	pulumi.CustomResourceState

	// BGP information specific to this router.
	Bgp RouterBgpResponseOutput `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers RouterBgpPeerResponseArrayOutput `pulumi:"bgpPeers"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolOutput `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces RouterInterfaceResponseArrayOutput `pulumi:"interfaces"`
	// Type of resource. Always compute#router for routers.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Keys used for MD5 authentication.
	Md5AuthenticationKeys RouterMd5AuthenticationKeyResponseArrayOutput `pulumi:"md5AuthenticationKeys"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats RouterNatResponseArrayOutput `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network pulumi.StringOutput `pulumi:"network"`
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOption) (*Router, error) {
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
	var resource Router
	err := ctx.RegisterResource("google-native:compute/v1:Router", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterState, opts ...pulumi.ResourceOption) (*Router, error) {
	var resource Router
	err := ctx.ReadResource("google-native:compute/v1:Router", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Router resources.
type routerState struct {
}

type RouterState struct {
}

func (RouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerState)(nil)).Elem()
}

type routerArgs struct {
	// BGP information specific to this router.
	Bgp *RouterBgp `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers []RouterBgpPeer `pulumi:"bgpPeers"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces []RouterInterface `pulumi:"interfaces"`
	// Keys used for MD5 authentication.
	Md5AuthenticationKeys []RouterMd5AuthenticationKey `pulumi:"md5AuthenticationKeys"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats []RouterNat `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// BGP information specific to this router.
	Bgp RouterBgpPtrInput
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers RouterBgpPeerArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces RouterInterfaceArrayInput
	// Keys used for MD5 authentication.
	Md5AuthenticationKeys RouterMd5AuthenticationKeyArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A list of NAT services created in this router.
	Nats RouterNatArrayInput
	// URI of the network to which this router belongs.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (RouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerArgs)(nil)).Elem()
}

type RouterInput interface {
	pulumi.Input

	ToRouterOutput() RouterOutput
	ToRouterOutputWithContext(ctx context.Context) RouterOutput
}

func (*Router) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil)).Elem()
}

func (i *Router) ToRouterOutput() RouterOutput {
	return i.ToRouterOutputWithContext(context.Background())
}

func (i *Router) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOutput)
}

type RouterOutput struct{ *pulumi.OutputState }

func (RouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil)).Elem()
}

func (o RouterOutput) ToRouterOutput() RouterOutput {
	return o
}

func (o RouterOutput) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return o
}

// BGP information specific to this router.
func (o RouterOutput) Bgp() RouterBgpResponseOutput {
	return o.ApplyT(func(v *Router) RouterBgpResponseOutput { return v.Bgp }).(RouterBgpResponseOutput)
}

// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
func (o RouterOutput) BgpPeers() RouterBgpPeerResponseArrayOutput {
	return o.ApplyT(func(v *Router) RouterBgpPeerResponseArrayOutput { return v.BgpPeers }).(RouterBgpPeerResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o RouterOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o RouterOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
func (o RouterOutput) EncryptedInterconnectRouter() pulumi.BoolOutput {
	return o.ApplyT(func(v *Router) pulumi.BoolOutput { return v.EncryptedInterconnectRouter }).(pulumi.BoolOutput)
}

// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
func (o RouterOutput) Interfaces() RouterInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *Router) RouterInterfaceResponseArrayOutput { return v.Interfaces }).(RouterInterfaceResponseArrayOutput)
}

// Type of resource. Always compute#router for routers.
func (o RouterOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Keys used for MD5 authentication.
func (o RouterOutput) Md5AuthenticationKeys() RouterMd5AuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v *Router) RouterMd5AuthenticationKeyResponseArrayOutput { return v.Md5AuthenticationKeys }).(RouterMd5AuthenticationKeyResponseArrayOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o RouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of NAT services created in this router.
func (o RouterOutput) Nats() RouterNatResponseArrayOutput {
	return o.ApplyT(func(v *Router) RouterNatResponseArrayOutput { return v.Nats }).(RouterNatResponseArrayOutput)
}

// URI of the network to which this router belongs.
func (o RouterOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o RouterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RouterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o RouterOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Router) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o RouterOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInput)(nil)).Elem(), &Router{})
	pulumi.RegisterOutputType(RouterOutput{})
}
