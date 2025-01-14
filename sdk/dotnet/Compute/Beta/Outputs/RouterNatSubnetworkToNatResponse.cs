// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Defines the IP ranges that want to use NAT for a subnetwork.
    /// </summary>
    [OutputType]
    public sealed class RouterNatSubnetworkToNatResponse
    {
        /// <summary>
        /// URL for the subnetwork resource that will use NAT.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of the secondary ranges of the Subnetwork that are allowed to use NAT. This can be populated only if "LIST_OF_SECONDARY_IP_RANGES" is one of the values in source_ip_ranges_to_nat.
        /// </summary>
        public readonly ImmutableArray<string> SecondaryIpRangeNames;
        /// <summary>
        /// Specify the options for NAT ranges in the Subnetwork. All options of a single value are valid except NAT_IP_RANGE_OPTION_UNSPECIFIED. The only valid option with multiple values is: ["PRIMARY_IP_RANGE", "LIST_OF_SECONDARY_IP_RANGES"] Default: [ALL_IP_RANGES]
        /// </summary>
        public readonly ImmutableArray<string> SourceIpRangesToNat;

        [OutputConstructor]
        private RouterNatSubnetworkToNatResponse(
            string name,

            ImmutableArray<string> secondaryIpRangeNames,

            ImmutableArray<string> sourceIpRangesToNat)
        {
            Name = name;
            SecondaryIpRangeNames = secondaryIpRangeNames;
            SourceIpRangesToNat = sourceIpRangesToNat;
        }
    }
}
