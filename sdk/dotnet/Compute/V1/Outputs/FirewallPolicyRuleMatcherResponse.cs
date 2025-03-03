// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Represents a match condition that incoming traffic is evaluated against. Exactly one field must be specified.
    /// </summary>
    [OutputType]
    public sealed class FirewallPolicyRuleMatcherResponse
    {
        /// <summary>
        /// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.
        /// </summary>
        public readonly ImmutableArray<string> DestAddressGroups;
        /// <summary>
        /// Fully Qualified Domain Name (FQDN) which should be matched against traffic destination. Maximum number of destination fqdn allowed is 100.
        /// </summary>
        public readonly ImmutableArray<string> DestFqdns;
        /// <summary>
        /// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
        /// </summary>
        public readonly ImmutableArray<string> DestIpRanges;
        /// <summary>
        /// Region codes whose IP addresses will be used to match for destination of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of dest region codes allowed is 5000.
        /// </summary>
        public readonly ImmutableArray<string> DestRegionCodes;
        /// <summary>
        /// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
        /// </summary>
        public readonly ImmutableArray<string> DestThreatIntelligences;
        /// <summary>
        /// Pairs of IP protocols and ports that the rule should match.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyRuleMatcherLayer4ConfigResponse> Layer4Configs;
        /// <summary>
        /// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.
        /// </summary>
        public readonly ImmutableArray<string> SrcAddressGroups;
        /// <summary>
        /// Fully Qualified Domain Name (FQDN) which should be matched against traffic source. Maximum number of source fqdn allowed is 100.
        /// </summary>
        public readonly ImmutableArray<string> SrcFqdns;
        /// <summary>
        /// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
        /// </summary>
        public readonly ImmutableArray<string> SrcIpRanges;
        /// <summary>
        /// Region codes whose IP addresses will be used to match for source of traffic. Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of source region codes allowed is 5000.
        /// </summary>
        public readonly ImmutableArray<string> SrcRegionCodes;
        /// <summary>
        /// List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the srcSecureTag are INEFFECTIVE, and there is no srcIpRange, this rule will be ignored. Maximum number of source tag values allowed is 256.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyRuleSecureTagResponse> SrcSecureTags;
        /// <summary>
        /// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
        /// </summary>
        public readonly ImmutableArray<string> SrcThreatIntelligences;

        [OutputConstructor]
        private FirewallPolicyRuleMatcherResponse(
            ImmutableArray<string> destAddressGroups,

            ImmutableArray<string> destFqdns,

            ImmutableArray<string> destIpRanges,

            ImmutableArray<string> destRegionCodes,

            ImmutableArray<string> destThreatIntelligences,

            ImmutableArray<Outputs.FirewallPolicyRuleMatcherLayer4ConfigResponse> layer4Configs,

            ImmutableArray<string> srcAddressGroups,

            ImmutableArray<string> srcFqdns,

            ImmutableArray<string> srcIpRanges,

            ImmutableArray<string> srcRegionCodes,

            ImmutableArray<Outputs.FirewallPolicyRuleSecureTagResponse> srcSecureTags,

            ImmutableArray<string> srcThreatIntelligences)
        {
            DestAddressGroups = destAddressGroups;
            DestFqdns = destFqdns;
            DestIpRanges = destIpRanges;
            DestRegionCodes = destRegionCodes;
            DestThreatIntelligences = destThreatIntelligences;
            Layer4Configs = layer4Configs;
            SrcAddressGroups = srcAddressGroups;
            SrcFqdns = srcFqdns;
            SrcIpRanges = srcIpRanges;
            SrcRegionCodes = srcRegionCodes;
            SrcSecureTags = srcSecureTags;
            SrcThreatIntelligences = srcThreatIntelligences;
        }
    }
}
