// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified NetworkAttachment resource in the given scope.
func LookupNetworkAttachment(ctx *pulumi.Context, args *LookupNetworkAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupNetworkAttachmentResult, error) {
	var rv LookupNetworkAttachmentResult
	err := ctx.Invoke("google-native:compute/v1:getNetworkAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkAttachmentArgs struct {
	NetworkAttachment string  `pulumi:"networkAttachment"`
	Project           *string `pulumi:"project"`
	Region            string  `pulumi:"region"`
}

type LookupNetworkAttachmentResult struct {
	// An array of connections for all the producers connected to this network attachment.
	ConnectionEndpoints  []NetworkAttachmentConnectedEndpointResponse `pulumi:"connectionEndpoints"`
	ConnectionPreference string                                       `pulumi:"connectionPreference"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
	Fingerprint string `pulumi:"fingerprint"`
	// Type of the resource.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
	Network string `pulumi:"network"`
	// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerAcceptLists []string `pulumi:"producerAcceptLists"`
	// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerRejectLists []string `pulumi:"producerRejectLists"`
	// URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	Subnetworks []string `pulumi:"subnetworks"`
}

func LookupNetworkAttachmentOutput(ctx *pulumi.Context, args LookupNetworkAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkAttachmentResult, error) {
			args := v.(LookupNetworkAttachmentArgs)
			r, err := LookupNetworkAttachment(ctx, &args, opts...)
			var s LookupNetworkAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkAttachmentResultOutput)
}

type LookupNetworkAttachmentOutputArgs struct {
	NetworkAttachment pulumi.StringInput    `pulumi:"networkAttachment"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
	Region            pulumi.StringInput    `pulumi:"region"`
}

func (LookupNetworkAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAttachmentArgs)(nil)).Elem()
}

type LookupNetworkAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkAttachmentResult)(nil)).Elem()
}

func (o LookupNetworkAttachmentResultOutput) ToLookupNetworkAttachmentResultOutput() LookupNetworkAttachmentResultOutput {
	return o
}

func (o LookupNetworkAttachmentResultOutput) ToLookupNetworkAttachmentResultOutputWithContext(ctx context.Context) LookupNetworkAttachmentResultOutput {
	return o
}

// An array of connections for all the producers connected to this network attachment.
func (o LookupNetworkAttachmentResultOutput) ConnectionEndpoints() NetworkAttachmentConnectedEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) []NetworkAttachmentConnectedEndpointResponse {
		return v.ConnectionEndpoints
	}).(NetworkAttachmentConnectedEndpointResponseArrayOutput)
}

func (o LookupNetworkAttachmentResultOutput) ConnectionPreference() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.ConnectionPreference }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupNetworkAttachmentResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupNetworkAttachmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. An up-to-date fingerprint must be provided in order to patch.
func (o LookupNetworkAttachmentResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupNetworkAttachmentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupNetworkAttachmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network which the Network Attachment belongs to. Practically it is inferred by fetching the network of the first subnetwork associated. Because it is required that all the subnetworks must be from the same network, it is assured that the Network Attachment belongs to the same network as all the subnetworks.
func (o LookupNetworkAttachmentResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Network }).(pulumi.StringOutput)
}

// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
func (o LookupNetworkAttachmentResultOutput) ProducerAcceptLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) []string { return v.ProducerAcceptLists }).(pulumi.StringArrayOutput)
}

// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
func (o LookupNetworkAttachmentResultOutput) ProducerRejectLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) []string { return v.ProducerRejectLists }).(pulumi.StringArrayOutput)
}

// URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupNetworkAttachmentResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupNetworkAttachmentResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource's resource id.
func (o LookupNetworkAttachmentResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
func (o LookupNetworkAttachmentResultOutput) Subnetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkAttachmentResult) []string { return v.Subnetworks }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkAttachmentResultOutput{})
}
