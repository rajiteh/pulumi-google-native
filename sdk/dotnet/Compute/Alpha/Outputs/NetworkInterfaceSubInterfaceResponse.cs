// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceSubInterfaceResponse
    {
        /// <summary>
        /// An IPv4 internal IP address to assign to the instance for this subinterface.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// If specified, this subnetwork must belong to the same network as that of the network interface. If not specified the subnet of network interface will be used. If you specify this property, you can specify the subnetwork as a full or partial URL. For example, the following are all valid URLs:  
        /// - https://www.googleapis.com/compute/v1/projects/project/regions/region/subnetworks/subnetwork 
        /// - regions/region/subnetworks/subnetwork
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// VLAN tag. Should match the VLAN(s) supported by the subnetwork to which this subinterface is connecting.
        /// </summary>
        public readonly int Vlan;

        [OutputConstructor]
        private NetworkInterfaceSubInterfaceResponse(
            string ipAddress,

            string subnetwork,

            int vlan)
        {
            IpAddress = ipAddress;
            Subnetwork = subnetwork;
            Vlan = vlan;
        }
    }
}