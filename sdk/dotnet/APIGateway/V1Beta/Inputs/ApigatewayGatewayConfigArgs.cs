// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta.Inputs
{

    /// <summary>
    /// Configuration settings for Gateways.
    /// </summary>
    public sealed class ApigatewayGatewayConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backend settings that are applied to all backends of the Gateway.
        /// </summary>
        [Input("backendConfig", required: true)]
        public Input<Inputs.ApigatewayBackendConfigArgs> BackendConfig { get; set; } = null!;

        public ApigatewayGatewayConfigArgs()
        {
        }
        public static new ApigatewayGatewayConfigArgs Empty => new ApigatewayGatewayConfigArgs();
    }
}
