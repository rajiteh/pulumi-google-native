// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches the representation of an existing Policy.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("google-native:dns/v1beta2:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	ClientOperationId *string `pulumi:"clientOperationId"`
	Policy            string  `pulumi:"policy"`
	Project           *string `pulumi:"project"`
}

type LookupPolicyResult struct {
	// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigResponse `pulumi:"alternativeNameServerConfig"`
	// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
	Description string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
	EnableInboundForwarding bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	EnableLogging bool   `pulumi:"enableLogging"`
	Kind          string `pulumi:"kind"`
	// User-assigned name for this policy.
	Name string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	Networks []PolicyNetworkResponse `pulumi:"networks"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	ClientOperationId pulumi.StringPtrInput `pulumi:"clientOperationId"`
	Policy            pulumi.StringInput    `pulumi:"policy"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

// Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
func (o LookupPolicyResultOutput) AlternativeNameServerConfig() PolicyAlternativeNameServerConfigResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) PolicyAlternativeNameServerConfigResponse {
		return v.AlternativeNameServerConfig
	}).(PolicyAlternativeNameServerConfigResponseOutput)
}

// A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
func (o LookupPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
func (o LookupPolicyResultOutput) EnableInboundForwarding() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicyResult) bool { return v.EnableInboundForwarding }).(pulumi.BoolOutput)
}

// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
func (o LookupPolicyResultOutput) EnableLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicyResult) bool { return v.EnableLogging }).(pulumi.BoolOutput)
}

func (o LookupPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// User-assigned name for this policy.
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of network names specifying networks to which this policy is applied.
func (o LookupPolicyResultOutput) Networks() PolicyNetworkResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []PolicyNetworkResponse { return v.Networks }).(PolicyNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
