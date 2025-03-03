// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single internal range.
func LookupInternalRange(ctx *pulumi.Context, args *LookupInternalRangeArgs, opts ...pulumi.InvokeOption) (*LookupInternalRangeResult, error) {
	var rv LookupInternalRangeResult
	err := ctx.Invoke("google-native:networkconnectivity/v1alpha1:getInternalRange", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInternalRangeArgs struct {
	InternalRangeId string  `pulumi:"internalRangeId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
}

type LookupInternalRangeResult struct {
	// Time when the internal range was created.
	CreateTime string `pulumi:"createTime"`
	// A description of this resource.
	Description string `pulumi:"description"`
	// IP range that this internal range defines.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// User-defined labels.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
	Name string `pulumi:"name"`
	// The URL or resource ID of the network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. This can only be specified for a global internal address. Example: - URL: /compute/v1/projects/{project}/global/networks/{resourceId} - ID: network123
	Network string `pulumi:"network"`
	// Optional. Types of resources that are allowed to overlap with the current internal range.
	Overlaps []string `pulumi:"overlaps"`
	// The type of peering set for this internal range.
	Peering string `pulumi:"peering"`
	// An alternative to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
	PrefixLength int `pulumi:"prefixLength"`
	// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
	TargetCidrRange []string `pulumi:"targetCidrRange"`
	// Time when the internal range was updated.
	UpdateTime string `pulumi:"updateTime"`
	// The type of usage set for this internal range.
	Usage string `pulumi:"usage"`
	// The list of resources that refer to this internal range. Resources that use the internal range for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range that is referred to. Can be empty.
	Users []string `pulumi:"users"`
}

func LookupInternalRangeOutput(ctx *pulumi.Context, args LookupInternalRangeOutputArgs, opts ...pulumi.InvokeOption) LookupInternalRangeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInternalRangeResult, error) {
			args := v.(LookupInternalRangeArgs)
			r, err := LookupInternalRange(ctx, &args, opts...)
			var s LookupInternalRangeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInternalRangeResultOutput)
}

type LookupInternalRangeOutputArgs struct {
	InternalRangeId pulumi.StringInput    `pulumi:"internalRangeId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInternalRangeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalRangeArgs)(nil)).Elem()
}

type LookupInternalRangeResultOutput struct{ *pulumi.OutputState }

func (LookupInternalRangeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalRangeResult)(nil)).Elem()
}

func (o LookupInternalRangeResultOutput) ToLookupInternalRangeResultOutput() LookupInternalRangeResultOutput {
	return o
}

func (o LookupInternalRangeResultOutput) ToLookupInternalRangeResultOutputWithContext(ctx context.Context) LookupInternalRangeResultOutput {
	return o
}

// Time when the internal range was created.
func (o LookupInternalRangeResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of this resource.
func (o LookupInternalRangeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.Description }).(pulumi.StringOutput)
}

// IP range that this internal range defines.
func (o LookupInternalRangeResultOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.IpCidrRange }).(pulumi.StringOutput)
}

// User-defined labels.
func (o LookupInternalRangeResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
func (o LookupInternalRangeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL or resource ID of the network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. This can only be specified for a global internal address. Example: - URL: /compute/v1/projects/{project}/global/networks/{resourceId} - ID: network123
func (o LookupInternalRangeResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.Network }).(pulumi.StringOutput)
}

// Optional. Types of resources that are allowed to overlap with the current internal range.
func (o LookupInternalRangeResultOutput) Overlaps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) []string { return v.Overlaps }).(pulumi.StringArrayOutput)
}

// The type of peering set for this internal range.
func (o LookupInternalRangeResultOutput) Peering() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.Peering }).(pulumi.StringOutput)
}

// An alternative to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
func (o LookupInternalRangeResultOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) int { return v.PrefixLength }).(pulumi.IntOutput)
}

// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
func (o LookupInternalRangeResultOutput) TargetCidrRange() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) []string { return v.TargetCidrRange }).(pulumi.StringArrayOutput)
}

// Time when the internal range was updated.
func (o LookupInternalRangeResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// The type of usage set for this internal range.
func (o LookupInternalRangeResultOutput) Usage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) string { return v.Usage }).(pulumi.StringOutput)
}

// The list of resources that refer to this internal range. Resources that use the internal range for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range that is referred to. Can be empty.
func (o LookupInternalRangeResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInternalRangeResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternalRangeResultOutput{})
}
