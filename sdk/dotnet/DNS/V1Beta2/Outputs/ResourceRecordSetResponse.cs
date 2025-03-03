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
    /// A unit of data that is returned by the DNS servers.
    /// </summary>
    [OutputType]
    public sealed class ResourceRecordSetResponse
    {
        public readonly string Kind;
        /// <summary>
        /// For example, www.example.com.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
        /// </summary>
        public readonly Outputs.RRSetRoutingPolicyResponse RoutingPolicy;
        /// <summary>
        /// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        /// </summary>
        public readonly ImmutableArray<string> Rrdatas;
        /// <summary>
        /// As defined in RFC 4034 (section 3.2).
        /// </summary>
        public readonly ImmutableArray<string> SignatureRrdatas;
        /// <summary>
        /// Number of seconds that this ResourceRecordSet can be cached by resolvers.
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// The identifier of a supported record type. See the list of Supported DNS record types.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ResourceRecordSetResponse(
            string kind,

            string name,

            Outputs.RRSetRoutingPolicyResponse routingPolicy,

            ImmutableArray<string> rrdatas,

            ImmutableArray<string> signatureRrdatas,

            int ttl,

            string type)
        {
            Kind = kind;
            Name = name;
            RoutingPolicy = routingPolicy;
            Rrdatas = rrdatas;
            SignatureRrdatas = signatureRrdatas;
            Ttl = ttl;
            Type = type;
        }
    }
}
