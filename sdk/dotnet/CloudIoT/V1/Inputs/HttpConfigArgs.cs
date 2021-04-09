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
    /// The configuration of the HTTP bridge for a device registry.
    /// </summary>
    public sealed class HttpConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If enabled, allows devices to use DeviceService via the HTTP protocol. Otherwise, any requests to DeviceService will fail for this registry.
        /// </summary>
        [Input("httpEnabledState")]
        public Input<string>? HttpEnabledState { get; set; }

        public HttpConfigArgs()
        {
        }
    }
}