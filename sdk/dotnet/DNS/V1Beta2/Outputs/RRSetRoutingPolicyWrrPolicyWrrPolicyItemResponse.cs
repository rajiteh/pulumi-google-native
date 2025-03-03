// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    /// <summary>
    /// A routing block which contains the routing information for one WRR item.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse
    {
        /// <summary>
        /// endpoints that need to be health checked before making the routing decision. The unhealthy endpoints will be omitted from the result. If all endpoints within a buckete are unhealthy, we'll choose a different bucket (sampled w.r.t. its weight) for responding. Note that if DNSSEC is enabled for this zone, only one of rrdata or health_checked_targets can be set.
        /// </summary>
        public readonly Outputs.RRSetRoutingPolicyHealthCheckTargetsResponse HealthCheckedTargets;
        public readonly string Kind;
        public readonly ImmutableArray<string> Rrdatas;
        /// <summary>
        /// DNSSEC generated signatures for all the rrdata within this item. Note that if health checked targets are provided for DNSSEC enabled zones, there's a restriction of 1 ip per item. .
        /// </summary>
        public readonly ImmutableArray<string> SignatureRrdatas;
        /// <summary>
        /// The weight corresponding to this subset of rrdata. When multiple WeightedRoundRobinPolicyItems are configured, the probability of returning an rrset is proportional to its weight relative to the sum of weights configured for all items. This weight should be non-negative.
        /// </summary>
        public readonly double Weight;

        [OutputConstructor]
        private RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse(
            Outputs.RRSetRoutingPolicyHealthCheckTargetsResponse healthCheckedTargets,

            string kind,

            ImmutableArray<string> rrdatas,

            ImmutableArray<string> signatureRrdatas,

            double weight)
        {
            HealthCheckedTargets = healthCheckedTargets;
            Kind = kind;
            Rrdatas = rrdatas;
            SignatureRrdatas = signatureRrdatas;
            Weight = weight;
        }
    }
}
