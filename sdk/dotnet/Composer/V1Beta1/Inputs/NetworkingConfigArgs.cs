// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration options for networking connections in the Composer 2 environment.
    /// </summary>
    public sealed class NetworkingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Indicates the user requested specifc connection type between Tenant and Customer projects. You cannot set networking connection type in public IP environment.
        /// </summary>
        [Input("connectionType")]
        public Input<Pulumi.GoogleNative.Composer.V1Beta1.NetworkingConfigConnectionType>? ConnectionType { get; set; }

        public NetworkingConfigArgs()
        {
        }
        public static new NetworkingConfigArgs Empty => new NetworkingConfigArgs();
    }
}
