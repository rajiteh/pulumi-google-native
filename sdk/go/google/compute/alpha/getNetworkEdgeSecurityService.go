// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a specified NetworkEdgeSecurityService.
func LookupNetworkEdgeSecurityService(ctx *pulumi.Context, args *LookupNetworkEdgeSecurityServiceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkEdgeSecurityServiceResult, error) {
	var rv LookupNetworkEdgeSecurityServiceResult
	err := ctx.Invoke("google-native:compute/alpha:getNetworkEdgeSecurityService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkEdgeSecurityServiceArgs struct {
	NetworkEdgeSecurityService string  `pulumi:"networkEdgeSecurityService"`
	Project                    *string `pulumi:"project"`
	Region                     string  `pulumi:"region"`
}

type LookupNetworkEdgeSecurityServiceResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
	Fingerprint string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#networkEdgeSecurityService for NetworkEdgeSecurityServices
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// URL of the region where the resource resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicy string `pulumi:"securityPolicy"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
}

func LookupNetworkEdgeSecurityServiceOutput(ctx *pulumi.Context, args LookupNetworkEdgeSecurityServiceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkEdgeSecurityServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkEdgeSecurityServiceResult, error) {
			args := v.(LookupNetworkEdgeSecurityServiceArgs)
			r, err := LookupNetworkEdgeSecurityService(ctx, &args, opts...)
			var s LookupNetworkEdgeSecurityServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkEdgeSecurityServiceResultOutput)
}

type LookupNetworkEdgeSecurityServiceOutputArgs struct {
	NetworkEdgeSecurityService pulumi.StringInput    `pulumi:"networkEdgeSecurityService"`
	Project                    pulumi.StringPtrInput `pulumi:"project"`
	Region                     pulumi.StringInput    `pulumi:"region"`
}

func (LookupNetworkEdgeSecurityServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEdgeSecurityServiceArgs)(nil)).Elem()
}

type LookupNetworkEdgeSecurityServiceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkEdgeSecurityServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEdgeSecurityServiceResult)(nil)).Elem()
}

func (o LookupNetworkEdgeSecurityServiceResultOutput) ToLookupNetworkEdgeSecurityServiceResultOutput() LookupNetworkEdgeSecurityServiceResultOutput {
	return o
}

func (o LookupNetworkEdgeSecurityServiceResultOutput) ToLookupNetworkEdgeSecurityServiceResultOutputWithContext(ctx context.Context) LookupNetworkEdgeSecurityServiceResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupNetworkEdgeSecurityServiceResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupNetworkEdgeSecurityServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService. An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a NetworkEdgeSecurityService.
func (o LookupNetworkEdgeSecurityServiceResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#networkEdgeSecurityService for NetworkEdgeSecurityServices
func (o LookupNetworkEdgeSecurityServiceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupNetworkEdgeSecurityServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL of the region where the resource resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupNetworkEdgeSecurityServiceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.Region }).(pulumi.StringOutput)
}

// The resource URL for the network edge security service associated with this network edge security service.
func (o LookupNetworkEdgeSecurityServiceResultOutput) SecurityPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.SecurityPolicy }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupNetworkEdgeSecurityServiceResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupNetworkEdgeSecurityServiceResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEdgeSecurityServiceResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkEdgeSecurityServiceResultOutput{})
}
