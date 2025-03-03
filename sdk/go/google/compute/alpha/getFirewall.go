// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified firewall.
func LookupFirewall(ctx *pulumi.Context, args *LookupFirewallArgs, opts ...pulumi.InvokeOption) (*LookupFirewallResult, error) {
	var rv LookupFirewallResult
	err := ctx.Invoke("google-native:compute/alpha:getFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallArgs struct {
	Firewall string  `pulumi:"firewall"`
	Project  *string `pulumi:"project"`
}

type LookupFirewallResult struct {
	// The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
	Allowed []FirewallAllowedItemResponse `pulumi:"allowed"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
	Denied []FirewallDeniedItemResponse `pulumi:"denied"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description string `pulumi:"description"`
	// If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Both IPv4 and IPv6 are supported.
	DestinationRanges []string `pulumi:"destinationRanges"`
	// Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `EGRESS` traffic, you cannot specify the sourceTags fields.
	Direction string `pulumi:"direction"`
	// Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
	Disabled bool `pulumi:"disabled"`
	// Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
	//
	// Deprecated: Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
	EnableLogging bool `pulumi:"enableLogging"`
	// Type of the resource. Always compute#firewall for firewall rules.
	Kind string `pulumi:"kind"`
	// This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
	LogConfig FirewallLogConfigResponse `pulumi:"logConfig"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name string `pulumi:"name"`
	// URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used: global/networks/default If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network - projects/myproject/global/networks/my-network - global/networks/default
	Network string `pulumi:"network"`
	// Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
	Priority int `pulumi:"priority"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Both IPv4 and IPv6 are supported.
	SourceRanges []string `pulumi:"sourceRanges"`
	// If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
	SourceServiceAccounts []string `pulumi:"sourceServiceAccounts"`
	// If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
	SourceTags []string `pulumi:"sourceTags"`
	// A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
	// A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
	TargetTags []string `pulumi:"targetTags"`
}

func LookupFirewallOutput(ctx *pulumi.Context, args LookupFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallResult, error) {
			args := v.(LookupFirewallArgs)
			r, err := LookupFirewall(ctx, &args, opts...)
			var s LookupFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallResultOutput)
}

type LookupFirewallOutputArgs struct {
	Firewall pulumi.StringInput    `pulumi:"firewall"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallArgs)(nil)).Elem()
}

type LookupFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallResult)(nil)).Elem()
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutput() LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutputWithContext(ctx context.Context) LookupFirewallResultOutput {
	return o
}

// The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
func (o LookupFirewallResultOutput) Allowed() FirewallAllowedItemResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []FirewallAllowedItemResponse { return v.Allowed }).(FirewallAllowedItemResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupFirewallResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
func (o LookupFirewallResultOutput) Denied() FirewallDeniedItemResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []FirewallDeniedItemResponse { return v.Denied }).(FirewallDeniedItemResponseArrayOutput)
}

// An optional description of this resource. Provide this field when you create the resource.
func (o LookupFirewallResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Description }).(pulumi.StringOutput)
}

// If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Both IPv4 and IPv6 are supported.
func (o LookupFirewallResultOutput) DestinationRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.DestinationRanges }).(pulumi.StringArrayOutput)
}

// Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `EGRESS` traffic, you cannot specify the sourceTags fields.
func (o LookupFirewallResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Direction }).(pulumi.StringOutput)
}

// Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
func (o LookupFirewallResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFirewallResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

// Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
//
// Deprecated: Deprecated in favor of enable in LogConfig. This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported t Cloud Logging.
func (o LookupFirewallResultOutput) EnableLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFirewallResult) bool { return v.EnableLogging }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#firewall for firewall rules.
func (o LookupFirewallResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Kind }).(pulumi.StringOutput)
}

// This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
func (o LookupFirewallResultOutput) LogConfig() FirewallLogConfigResponseOutput {
	return o.ApplyT(func(v LookupFirewallResult) FirewallLogConfigResponse { return v.LogConfig }).(FirewallLogConfigResponseOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
func (o LookupFirewallResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used: global/networks/default If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network - projects/myproject/global/networks/my-network - global/networks/default
func (o LookupFirewallResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Network }).(pulumi.StringOutput)
}

// Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
func (o LookupFirewallResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Server-defined URL for the resource.
func (o LookupFirewallResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupFirewallResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Both IPv4 and IPv6 are supported.
func (o LookupFirewallResultOutput) SourceRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.SourceRanges }).(pulumi.StringArrayOutput)
}

// If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
func (o LookupFirewallResultOutput) SourceServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.SourceServiceAccounts }).(pulumi.StringArrayOutput)
}

// If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
func (o LookupFirewallResultOutput) SourceTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.SourceTags }).(pulumi.StringArrayOutput)
}

// A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
func (o LookupFirewallResultOutput) TargetServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.TargetServiceAccounts }).(pulumi.StringArrayOutput)
}

// A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
func (o LookupFirewallResultOutput) TargetTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFirewallResult) []string { return v.TargetTags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallResultOutput{})
}
