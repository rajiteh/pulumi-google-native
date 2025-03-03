// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Inputs
{

    /// <summary>
    /// The basic ingress config for ClientGateways.
    /// </summary>
    public sealed class ConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationRoutes", required: true)]
        private InputList<Inputs.DestinationRouteArgs>? _destinationRoutes;

        /// <summary>
        /// The settings used to configure basic ClientGateways.
        /// </summary>
        public InputList<Inputs.DestinationRouteArgs> DestinationRoutes
        {
            get => _destinationRoutes ?? (_destinationRoutes = new InputList<Inputs.DestinationRouteArgs>());
            set => _destinationRoutes = value;
        }

        /// <summary>
        /// Immutable. The transport protocol used between the client and the server.
        /// </summary>
        [Input("transportProtocol", required: true)]
        public Input<Pulumi.GoogleNative.BeyondCorp.V1Alpha.ConfigTransportProtocol> TransportProtocol { get; set; } = null!;

        public ConfigArgs()
        {
        }
        public static new ConfigArgs Empty => new ConfigArgs();
    }
}
