// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// NetworkPolicy describes VM instance network configurations.
    /// </summary>
    [OutputType]
    public sealed class NetworkPolicyResponse
    {
        /// <summary>
        /// Network configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponse> NetworkInterfaces;

        [OutputConstructor]
        private NetworkPolicyResponse(ImmutableArray<Outputs.NetworkInterfaceResponse> networkInterfaces)
        {
            NetworkInterfaces = networkInterfaces;
        }
    }
}
