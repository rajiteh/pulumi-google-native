// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for Cloud TPU.
    /// </summary>
    [OutputType]
    public sealed class TpuConfigResponse
    {
        /// <summary>
        /// Whether Cloud TPU integration is enabled or not.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// IPv4 CIDR block reserved for Cloud TPU in the VPC.
        /// </summary>
        public readonly string Ipv4CidrBlock;
        /// <summary>
        /// Whether to use service networking for Cloud TPU or not.
        /// </summary>
        public readonly bool UseServiceNetworking;

        [OutputConstructor]
        private TpuConfigResponse(
            bool enabled,

            string ipv4CidrBlock,

            bool useServiceNetworking)
        {
            Enabled = enabled;
            Ipv4CidrBlock = ipv4CidrBlock;
            UseServiceNetworking = useServiceNetworking;
        }
    }
}
