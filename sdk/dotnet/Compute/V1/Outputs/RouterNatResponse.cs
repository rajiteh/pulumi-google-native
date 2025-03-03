// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Represents a Nat resource. It enables the VMs within the specified subnetworks to access Internet without external IP addresses. It specifies a list of subnetworks (and the ranges within) that want to use NAT. Customers can also provide the external IPs that would be used for NAT. GCP would auto-allocate ephemeral IPs if no external IPs are provided.
    /// </summary>
    [OutputType]
    public sealed class RouterNatResponse
    {
        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be valid static external IPs that have been assigned to the NAT. These IPs should be used for updating/patching a NAT only.
        /// </summary>
        public readonly ImmutableArray<string> DrainNatIps;
        /// <summary>
        /// Enable Dynamic Port Allocation. If not specified, it is disabled by default. If set to true, - Dynamic Port Allocation will be enabled on this NAT config. - enableEndpointIndependentMapping cannot be set to true. - If minPorts is set, minPortsPerVm must be set to a power of two greater than or equal to 32. If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config. 
        /// </summary>
        public readonly bool EnableDynamicPortAllocation;
        public readonly bool EnableEndpointIndependentMapping;
        /// <summary>
        /// List of NAT-ted endpoint types supported by the Nat Gateway. If the list is empty, then it will be equivalent to include ENDPOINT_TYPE_VM
        /// </summary>
        public readonly ImmutableArray<string> EndpointTypes;
        /// <summary>
        /// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
        /// </summary>
        public readonly int IcmpIdleTimeoutSec;
        /// <summary>
        /// Configure logging on this NAT.
        /// </summary>
        public readonly Outputs.RouterNatLogConfigResponse LogConfig;
        /// <summary>
        /// Maximum number of ports allocated to a VM from this NAT config when Dynamic Port Allocation is enabled. If Dynamic Port Allocation is not enabled, this field has no effect. If Dynamic Port Allocation is enabled, and this field is set, it must be set to a power of two greater than minPortsPerVm, or 64 if minPortsPerVm is not set. If Dynamic Port Allocation is enabled and this field is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
        /// </summary>
        public readonly int MaxPortsPerVm;
        /// <summary>
        /// Minimum number of ports allocated to a VM from this NAT config. If not set, a default number of ports is allocated to a VM. This is rounded up to the nearest power of 2. For example, if the value of this field is 50, at least 64 ports are allocated to a VM.
        /// </summary>
        public readonly int MinPortsPerVm;
        /// <summary>
        /// Unique name of this Nat service. The name must be 1-63 characters long and comply with RFC1035.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specify the NatIpAllocateOption, which can take one of the following values: - MANUAL_ONLY: Uses only Nat IP addresses provided by customers. When there are not enough specified Nat IPs, the Nat service fails for new VMs. - AUTO_ONLY: Nat IPs are allocated by Google Cloud Platform; customers can't specify any Nat IPs. When choosing AUTO_ONLY, then nat_ip should be empty. 
        /// </summary>
        public readonly string NatIpAllocateOption;
        /// <summary>
        /// A list of URLs of the IP resources used for this Nat service. These IP addresses must be valid static external IP addresses assigned to the project.
        /// </summary>
        public readonly ImmutableArray<string> NatIps;
        /// <summary>
        /// A list of rules associated with this NAT.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterNatRuleResponse> Rules;
        /// <summary>
        /// Specify the Nat option, which can take one of the following values: - ALL_SUBNETWORKS_ALL_IP_RANGES: All of the IP ranges in every Subnetwork are allowed to Nat. - ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES: All of the primary IP ranges in every Subnetwork are allowed to Nat. - LIST_OF_SUBNETWORKS: A list of Subnetworks are allowed to Nat (specified in the field subnetwork below) The default is SUBNETWORK_IP_RANGE_TO_NAT_OPTION_UNSPECIFIED. Note that if this field contains ALL_SUBNETWORKS_ALL_IP_RANGES then there should not be any other Router.Nat section in any Router for this network in this region.
        /// </summary>
        public readonly string SourceSubnetworkIpRangesToNat;
        /// <summary>
        /// A list of Subnetwork resources whose traffic should be translated by NAT Gateway. It is used only when LIST_OF_SUBNETWORKS is selected for the SubnetworkIpRangeToNatOption above.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterNatSubnetworkToNatResponse> Subnetworks;
        /// <summary>
        /// Timeout (in seconds) for TCP established connections. Defaults to 1200s if not set.
        /// </summary>
        public readonly int TcpEstablishedIdleTimeoutSec;
        /// <summary>
        /// Timeout (in seconds) for TCP connections that are in TIME_WAIT state. Defaults to 120s if not set.
        /// </summary>
        public readonly int TcpTimeWaitTimeoutSec;
        /// <summary>
        /// Timeout (in seconds) for TCP transitory connections. Defaults to 30s if not set.
        /// </summary>
        public readonly int TcpTransitoryIdleTimeoutSec;
        /// <summary>
        /// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
        /// </summary>
        public readonly int UdpIdleTimeoutSec;

        [OutputConstructor]
        private RouterNatResponse(
            ImmutableArray<string> drainNatIps,

            bool enableDynamicPortAllocation,

            bool enableEndpointIndependentMapping,

            ImmutableArray<string> endpointTypes,

            int icmpIdleTimeoutSec,

            Outputs.RouterNatLogConfigResponse logConfig,

            int maxPortsPerVm,

            int minPortsPerVm,

            string name,

            string natIpAllocateOption,

            ImmutableArray<string> natIps,

            ImmutableArray<Outputs.RouterNatRuleResponse> rules,

            string sourceSubnetworkIpRangesToNat,

            ImmutableArray<Outputs.RouterNatSubnetworkToNatResponse> subnetworks,

            int tcpEstablishedIdleTimeoutSec,

            int tcpTimeWaitTimeoutSec,

            int tcpTransitoryIdleTimeoutSec,

            int udpIdleTimeoutSec)
        {
            DrainNatIps = drainNatIps;
            EnableDynamicPortAllocation = enableDynamicPortAllocation;
            EnableEndpointIndependentMapping = enableEndpointIndependentMapping;
            EndpointTypes = endpointTypes;
            IcmpIdleTimeoutSec = icmpIdleTimeoutSec;
            LogConfig = logConfig;
            MaxPortsPerVm = maxPortsPerVm;
            MinPortsPerVm = minPortsPerVm;
            Name = name;
            NatIpAllocateOption = natIpAllocateOption;
            NatIps = natIps;
            Rules = rules;
            SourceSubnetworkIpRangesToNat = sourceSubnetworkIpRangesToNat;
            Subnetworks = subnetworks;
            TcpEstablishedIdleTimeoutSec = tcpEstablishedIdleTimeoutSec;
            TcpTimeWaitTimeoutSec = tcpTimeWaitTimeoutSec;
            TcpTransitoryIdleTimeoutSec = tcpTransitoryIdleTimeoutSec;
            UdpIdleTimeoutSec = udpIdleTimeoutSec;
        }
    }
}
