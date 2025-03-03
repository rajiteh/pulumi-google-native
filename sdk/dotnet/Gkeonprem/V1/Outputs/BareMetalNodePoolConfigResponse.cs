// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// BareMetalNodePoolConfig describes the configuration of all nodes within a given bare metal node pool.
    /// </summary>
    [OutputType]
    public sealed class BareMetalNodePoolConfigResponse
    {
        /// <summary>
        /// The modifiable kubelet configurations for the baremetal machines.
        /// </summary>
        public readonly Outputs.BareMetalKubeletConfigResponse KubeletConfig;
        /// <summary>
        /// The labels assigned to nodes of this node pool. An object containing a list of key/value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The list of machine addresses in the bare metal node pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.BareMetalNodeConfigResponse> NodeConfigs;
        /// <summary>
        /// Specifies the nodes operating system (default: LINUX).
        /// </summary>
        public readonly string OperatingSystem;
        /// <summary>
        /// The initial taints assigned to nodes of this node pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeTaintResponse> Taints;

        [OutputConstructor]
        private BareMetalNodePoolConfigResponse(
            Outputs.BareMetalKubeletConfigResponse kubeletConfig,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.BareMetalNodeConfigResponse> nodeConfigs,

            string operatingSystem,

            ImmutableArray<Outputs.NodeTaintResponse> taints)
        {
            KubeletConfig = kubeletConfig;
            Labels = labels;
            NodeConfigs = nodeConfigs;
            OperatingSystem = operatingSystem;
            Taints = taints;
        }
    }
}
