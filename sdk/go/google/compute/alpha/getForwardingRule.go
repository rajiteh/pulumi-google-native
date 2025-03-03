// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified ForwardingRule resource.
func LookupForwardingRule(ctx *pulumi.Context, args *LookupForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupForwardingRuleResult, error) {
	var rv LookupForwardingRuleResult
	err := ctx.Invoke("google-native:compute/alpha:getForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupForwardingRuleArgs struct {
	ForwardingRule string  `pulumi:"forwardingRule"`
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
}

type LookupForwardingRuleResult struct {
	// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal and external protocol forwarding. Set this field to true to allow packets addressed to any port or packets lacking destination port information (for example, UDP fragments after the first fragment) to be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive.
	AllPorts bool `pulumi:"allPorts"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	AllowGlobalAccess bool `pulumi:"allowGlobalAccess"`
	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	AllowPscGlobalAccess bool `pulumi:"allowPscGlobalAccess"`
	// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
	BackendService string `pulumi:"backendService"`
	// The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
	BaseForwardingRule string `pulumi:"baseForwardingRule"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	Fingerprint string `pulumi:"fingerprint"`
	// IP address for which this forwarding rule accepts traffic. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the referenced target or backendService. While creating a forwarding rule, specifying an IPAddress is required under the following circumstances: - When the target is set to targetGrpcProxy and validateForProxyless is set to true, the IPAddress should be set to 0.0.0.0. - When the target is a Private Service Connect Google APIs bundle, you must specify an IPAddress. Otherwise, you can optionally specify an IP address that references an existing static (reserved) IP address resource. When omitted, Google Cloud assigns an ephemeral IP address. Use one of the following formats to specify an IP address while creating a forwarding rule: * IP address number, as in `100.1.2.3` * IPv6 address range, as in `2600:1234::/96` * Full resource URL, as in https://www.googleapis.com/compute/v1/projects/ project_id/regions/region/addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The forwarding rule's target or backendService, and in most cases, also the loadBalancingScheme, determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). When reading an IPAddress, the API always returns the IP address number.
	IpAddress string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
	IpProtocol string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6.
	IpVersion string `pulumi:"ipVersion"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector bool `pulumi:"isMirroringCollector"`
	// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
	LoadBalancingScheme string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	MetadataFilters []MetadataFilterResponse `pulumi:"metadataFilters"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. For Private Service Connect forwarding rules that forward traffic to Google APIs, the forwarding rule name must be a 1-20 characters string with lowercase letters and numbers and must start with a letter.
	Name string `pulumi:"name"`
	// This field is not used for external load balancing. For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If the subnetwork is specified, the network of the subnetwork will be used. If neither subnetwork nor this field is specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	Network string `pulumi:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	NetworkTier string `pulumi:"networkTier"`
	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDnsZone bool `pulumi:"noAutomateDnsZone"`
	// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By backend service-based network load balancers, target pool-based network load balancers, internal proxy load balancers, external proxy load balancers, Traffic Director, external protocol forwarding, and Classic VPN. Some products have restrictions on what ports can be used. See port specifications for details. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. @pattern: \\d+(?:-\\d+)?
	PortRange string `pulumi:"portRange"`
	// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal protocol forwarding. You can specify a list of up to five ports by number, separated by commas. The ports can be contiguous or discontiguous. Only packets addressed to these ports will be forwarded to the backends configured with this forwarding rule. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. The ports, port_range, and allPorts fields are mutually exclusive. @pattern: \\d+(?:-\\d+)?
	Ports []string `pulumi:"ports"`
	// The PSC connection id of the PSC Forwarding Rule.
	PscConnectionId     string `pulumi:"pscConnectionId"`
	PscConnectionStatus string `pulumi:"pscConnectionStatus"`
	// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	ServiceDirectoryRegistrations []ForwardingRuleServiceDirectoryRegistrationResponse `pulumi:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
	ServiceLabel string `pulumi:"serviceLabel"`
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.
	ServiceName string `pulumi:"serviceName"`
	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each source_ip_range entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIpRanges []string `pulumi:"sourceIpRanges"`
	// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
	Subnetwork string `pulumi:"subnetwork"`
	// The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must be in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. - For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). - For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle: - vpc-sc - APIs that support VPC Service Controls. - all-apis - All supported Google APIs. - For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.
	Target string `pulumi:"target"`
}

func LookupForwardingRuleOutput(ctx *pulumi.Context, args LookupForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupForwardingRuleResult, error) {
			args := v.(LookupForwardingRuleArgs)
			r, err := LookupForwardingRule(ctx, &args, opts...)
			var s LookupForwardingRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupForwardingRuleResultOutput)
}

type LookupForwardingRuleOutputArgs struct {
	ForwardingRule pulumi.StringInput    `pulumi:"forwardingRule"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
}

func (LookupForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleArgs)(nil)).Elem()
}

type LookupForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleResult)(nil)).Elem()
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutput() LookupForwardingRuleResultOutput {
	return o
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutputWithContext(ctx context.Context) LookupForwardingRuleResultOutput {
	return o
}

// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal and external protocol forwarding. Set this field to true to allow packets addressed to any port or packets lacking destination port information (for example, UDP fragments after the first fragment) to be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive.
func (o LookupForwardingRuleResultOutput) AllPorts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) bool { return v.AllPorts }).(pulumi.BoolOutput)
}

// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
func (o LookupForwardingRuleResultOutput) AllowGlobalAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) bool { return v.AllowGlobalAccess }).(pulumi.BoolOutput)
}

// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
func (o LookupForwardingRuleResultOutput) AllowPscGlobalAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) bool { return v.AllowPscGlobalAccess }).(pulumi.BoolOutput)
}

// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
func (o LookupForwardingRuleResultOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.BackendService }).(pulumi.StringOutput)
}

// The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
func (o LookupForwardingRuleResultOutput) BaseForwardingRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.BaseForwardingRule }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupForwardingRuleResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupForwardingRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
func (o LookupForwardingRuleResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// IP address for which this forwarding rule accepts traffic. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the referenced target or backendService. While creating a forwarding rule, specifying an IPAddress is required under the following circumstances: - When the target is set to targetGrpcProxy and validateForProxyless is set to true, the IPAddress should be set to 0.0.0.0. - When the target is a Private Service Connect Google APIs bundle, you must specify an IPAddress. Otherwise, you can optionally specify an IP address that references an existing static (reserved) IP address resource. When omitted, Google Cloud assigns an ephemeral IP address. Use one of the following formats to specify an IP address while creating a forwarding rule: * IP address number, as in `100.1.2.3` * IPv6 address range, as in `2600:1234::/96` * Full resource URL, as in https://www.googleapis.com/compute/v1/projects/ project_id/regions/region/addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The forwarding rule's target or backendService, and in most cases, also the loadBalancingScheme, determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). When reading an IPAddress, the API always returns the IP address number.
func (o LookupForwardingRuleResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
func (o LookupForwardingRuleResultOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.IpProtocol }).(pulumi.StringOutput)
}

// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6.
func (o LookupForwardingRuleResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
func (o LookupForwardingRuleResultOutput) IsMirroringCollector() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) bool { return v.IsMirroringCollector }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
func (o LookupForwardingRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
func (o LookupForwardingRuleResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupForwardingRuleResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
func (o LookupForwardingRuleResultOutput) LoadBalancingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.LoadBalancingScheme }).(pulumi.StringOutput)
}

// Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o LookupForwardingRuleResultOutput) MetadataFilters() MetadataFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []MetadataFilterResponse { return v.MetadataFilters }).(MetadataFilterResponseArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. For Private Service Connect forwarding rules that forward traffic to Google APIs, the forwarding rule name must be a 1-20 characters string with lowercase letters and numbers and must start with a letter.
func (o LookupForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// This field is not used for external load balancing. For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If the subnetwork is specified, the network of the subnetwork will be used. If neither subnetwork nor this field is specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
func (o LookupForwardingRuleResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Network }).(pulumi.StringOutput)
}

// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
func (o LookupForwardingRuleResultOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.NetworkTier }).(pulumi.StringOutput)
}

// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
func (o LookupForwardingRuleResultOutput) NoAutomateDnsZone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) bool { return v.NoAutomateDnsZone }).(pulumi.BoolOutput)
}

// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By backend service-based network load balancers, target pool-based network load balancers, internal proxy load balancers, external proxy load balancers, Traffic Director, external protocol forwarding, and Classic VPN. Some products have restrictions on what ports can be used. See port specifications for details. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The ports, port_range, and allPorts fields are mutually exclusive. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. @pattern: \\d+(?:-\\d+)?
func (o LookupForwardingRuleResultOutput) PortRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.PortRange }).(pulumi.StringOutput)
}

// This field can only be used: - If IPProtocol is one of TCP, UDP, or SCTP. - By internal TCP/UDP load balancers, backend service-based network load balancers, and internal protocol forwarding. You can specify a list of up to five ports by number, separated by commas. The ports can be contiguous or discontiguous. Only packets addressed to these ports will be forwarded to the backends configured with this forwarding rule. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot share any values defined in ports. The ports, port_range, and allPorts fields are mutually exclusive. @pattern: \\d+(?:-\\d+)?
func (o LookupForwardingRuleResultOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

// The PSC connection id of the PSC Forwarding Rule.
func (o LookupForwardingRuleResultOutput) PscConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.PscConnectionId }).(pulumi.StringOutput)
}

func (o LookupForwardingRuleResultOutput) PscConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.PscConnectionStatus }).(pulumi.StringOutput)
}

// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupForwardingRuleResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupForwardingRuleResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupForwardingRuleResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
func (o LookupForwardingRuleResultOutput) ServiceDirectoryRegistrations() ForwardingRuleServiceDirectoryRegistrationResponseArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []ForwardingRuleServiceDirectoryRegistrationResponse {
		return v.ServiceDirectoryRegistrations
	}).(ForwardingRuleServiceDirectoryRegistrationResponseArrayOutput)
}

// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
func (o LookupForwardingRuleResultOutput) ServiceLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.ServiceLabel }).(pulumi.StringOutput)
}

// The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.
func (o LookupForwardingRuleResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each source_ip_range entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
func (o LookupForwardingRuleResultOutput) SourceIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []string { return v.SourceIpRanges }).(pulumi.StringArrayOutput)
}

// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
func (o LookupForwardingRuleResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must be in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. - For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). - For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle: - vpc-sc - APIs that support VPC Service Controls. - all-apis - All supported Google APIs. - For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.
func (o LookupForwardingRuleResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Target }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupForwardingRuleResultOutput{})
}
