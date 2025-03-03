// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta.Inputs
{

    /// <summary>
    /// RoutingConfig configures the behaviour of fleet logging feature.
    /// </summary>
    public sealed class FleetObservabilityRoutingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// mode configures the logs routing mode.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.GKEHub.V1Beta.FleetObservabilityRoutingConfigMode>? Mode { get; set; }

        public FleetObservabilityRoutingConfigArgs()
        {
        }
        public static new FleetObservabilityRoutingConfigArgs Empty => new FleetObservabilityRoutingConfigArgs();
    }
}
