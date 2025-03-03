// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// The basic ingress config for ClientGateways.
    /// </summary>
    [OutputType]
    public sealed class ConfigResponse
    {
        /// <summary>
        /// The settings used to configure basic ClientGateways.
        /// </summary>
        public readonly ImmutableArray<Outputs.DestinationRouteResponse> DestinationRoutes;
        /// <summary>
        /// Immutable. The transport protocol used between the client and the server.
        /// </summary>
        public readonly string TransportProtocol;

        [OutputConstructor]
        private ConfigResponse(
            ImmutableArray<Outputs.DestinationRouteResponse> destinationRoutes,

            string transportProtocol)
        {
            DestinationRoutes = destinationRoutes;
            TransportProtocol = transportProtocol;
        }
    }
}
