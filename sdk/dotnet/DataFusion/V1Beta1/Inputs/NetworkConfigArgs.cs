// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataFusion.V1Beta1.Inputs
{

    /// <summary>
    /// Network configuration for a Data Fusion instance. These configurations are used for peering with the customer network. Configurations are optional when a public Data Fusion instance is to be created. However, providing these configurations allows several benefits, such as reduced network latency while accessing the customer resources from managed Data Fusion instance nodes, as well as access to the customer on-prem resources.
    /// </summary>
    public sealed class NetworkConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP range in CIDR notation to use for the managed Data Fusion instance nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
        /// </summary>
        [Input("ipAllocation")]
        public Input<string>? IpAllocation { get; set; }

        /// <summary>
        /// Name of the network in the customer project with which the Tenant Project will be peered for executing pipelines. In case of shared VPC where the network resides in another host project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        public NetworkConfigArgs()
        {
        }
        public static new NetworkConfigArgs Empty => new NetworkConfigArgs();
    }
}
