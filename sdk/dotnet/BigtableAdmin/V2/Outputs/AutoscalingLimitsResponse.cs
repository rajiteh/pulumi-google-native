// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// Limits for the number of nodes a Cluster can autoscale up/down to.
    /// </summary>
    [OutputType]
    public sealed class AutoscalingLimitsResponse
    {
        /// <summary>
        /// Maximum number of nodes to scale up to.
        /// </summary>
        public readonly int MaxServeNodes;
        /// <summary>
        /// Minimum number of nodes to scale down to.
        /// </summary>
        public readonly int MinServeNodes;

        [OutputConstructor]
        private AutoscalingLimitsResponse(
            int maxServeNodes,

            int minServeNodes)
        {
            MaxServeNodes = maxServeNodes;
            MinServeNodes = minServeNodes;
        }
    }
}
