// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified VPN gateway.
func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("google-native:compute/beta:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	Project    *string `pulumi:"project"`
	Region     string  `pulumi:"region"`
	VpnGateway string  `pulumi:"vpnGateway"`
}

type LookupVpnGatewayResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used.
	GatewayIpVersion string `pulumi:"gatewayIpVersion"`
	// Type of resource. Always compute#vpnGateway for VPN gateways.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnGateway.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network string `pulumi:"network"`
	// URL of the region where the VPN gateway resides.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The stack type for this VPN gateway to identify the IP protocols that are enabled. Possible values are: IPV4_ONLY, IPV4_IPV6. If not specified, IPV4_ONLY will be used.
	StackType string `pulumi:"stackType"`
	// The list of VPN interfaces associated with this VPN gateway.
	VpnInterfaces []VpnGatewayVpnGatewayInterfaceResponse `pulumi:"vpnInterfaces"`
}

func LookupVpnGatewayOutput(ctx *pulumi.Context, args LookupVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnGatewayResult, error) {
			args := v.(LookupVpnGatewayArgs)
			r, err := LookupVpnGateway(ctx, &args, opts...)
			var s LookupVpnGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnGatewayResultOutput)
}

type LookupVpnGatewayOutputArgs struct {
	Project    pulumi.StringPtrInput `pulumi:"project"`
	Region     pulumi.StringInput    `pulumi:"region"`
	VpnGateway pulumi.StringInput    `pulumi:"vpnGateway"`
}

func (LookupVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayArgs)(nil)).Elem()
}

type LookupVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayResult)(nil)).Elem()
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutput() LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutputWithContext(ctx context.Context) LookupVpnGatewayResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupVpnGatewayResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupVpnGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

// The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used.
func (o LookupVpnGatewayResultOutput) GatewayIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.GatewayIpVersion }).(pulumi.StringOutput)
}

// Type of resource. Always compute#vpnGateway for VPN gateways.
func (o LookupVpnGatewayResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnGateway.
func (o LookupVpnGatewayResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupVpnGatewayResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
func (o LookupVpnGatewayResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Network }).(pulumi.StringOutput)
}

// URL of the region where the VPN gateway resides.
func (o LookupVpnGatewayResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupVpnGatewayResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The stack type for this VPN gateway to identify the IP protocols that are enabled. Possible values are: IPV4_ONLY, IPV4_IPV6. If not specified, IPV4_ONLY will be used.
func (o LookupVpnGatewayResultOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.StackType }).(pulumi.StringOutput)
}

// The list of VPN interfaces associated with this VPN gateway.
func (o LookupVpnGatewayResultOutput) VpnInterfaces() VpnGatewayVpnGatewayInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnGatewayVpnGatewayInterfaceResponse { return v.VpnInterfaces }).(VpnGatewayVpnGatewayInterfaceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnGatewayResultOutput{})
}
