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
    /// BareMetalAdminDrainingMachine represents the machines that are currently draining.
    /// </summary>
    [OutputType]
    public sealed class BareMetalAdminDrainingMachineResponse
    {
        /// <summary>
        /// Draining machine IP address.
        /// </summary>
        public readonly string NodeIp;
        /// <summary>
        /// The count of pods yet to drain.
        /// </summary>
        public readonly int PodCount;

        [OutputConstructor]
        private BareMetalAdminDrainingMachineResponse(
            string nodeIp,

            int podCount)
        {
            NodeIp = nodeIp;
            PodCount = podCount;
        }
    }
}
