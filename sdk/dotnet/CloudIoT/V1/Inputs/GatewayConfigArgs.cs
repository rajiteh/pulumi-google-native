// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudIoT.V1.Inputs
{

    /// <summary>
    /// Gateway-related configuration and state.
    /// </summary>
    public sealed class GatewayConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates how to authorize and/or authenticate devices to access the gateway.
        /// </summary>
        [Input("gatewayAuthMethod")]
        public Input<string>? GatewayAuthMethod { get; set; }

        /// <summary>
        /// Indicates whether the device is a gateway.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// [Output only] The ID of the gateway the device accessed most recently.
        /// </summary>
        [Input("lastAccessedGatewayId")]
        public Input<string>? LastAccessedGatewayId { get; set; }

        /// <summary>
        /// [Output only] The most recent time at which the device accessed the gateway specified in `last_accessed_gateway`.
        /// </summary>
        [Input("lastAccessedGatewayTime")]
        public Input<string>? LastAccessedGatewayTime { get; set; }

        public GatewayConfigArgs()
        {
        }
    }
}