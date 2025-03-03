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
    /// ApplicationEndpoint represents a remote application endpoint.
    /// </summary>
    [OutputType]
    public sealed class ApplicationEndpointResponse
    {
        /// <summary>
        /// Hostname or IP address of the remote application endpoint.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Port of the remote application endpoint.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private ApplicationEndpointResponse(
            string host,

            int port)
        {
            Host = host;
            Port = port;
        }
    }
}
