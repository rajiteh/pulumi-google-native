// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified address resource. Gets a list of available addresses by making a list() request.
func LookupGlobalAddress(ctx *pulumi.Context, args *LookupGlobalAddressArgs, opts ...pulumi.InvokeOption) (*LookupGlobalAddressResult, error) {
	var rv LookupGlobalAddressResult
	err := ctx.Invoke("google-native:compute/v1:getGlobalAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalAddressArgs struct {
	Address string `pulumi:"address"`
	Project string `pulumi:"project"`
}

type LookupGlobalAddressResult struct {
	// The static IP address represented by this resource.
	Address string `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType string `pulumi:"addressType"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description string `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion string `pulumi:"ipVersion"`
	// [Output Only] Type of the resource. Always compute#address for addresses.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network string `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer.
	//
	// If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier string `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength int `pulumi:"prefixLength"`
	// The purpose of this resource, which can be one of the following values:
	// - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// - `DNS_RESOLVER` for a DNS resolver address in a subnetwork
	// - `VPC_PEERING` for addresses that are reserved for VPC peer networks.
	// - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT.
	// - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration. These addresses are regional resources.
	Purpose string `pulumi:"purpose"`
	// [Output Only] The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. This field is not applicable to global addresses.
	Region string `pulumi:"region"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// [Output Only] The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
	Status string `pulumi:"status"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork string `pulumi:"subnetwork"`
	// [Output Only] The URLs of the resources that are using this address.
	Users []string `pulumi:"users"`
}