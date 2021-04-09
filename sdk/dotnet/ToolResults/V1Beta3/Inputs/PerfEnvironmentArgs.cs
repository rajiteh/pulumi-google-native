// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// Encapsulates performance environment info
    /// </summary>
    public sealed class PerfEnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPU related environment info
        /// </summary>
        [Input("cpuInfo")]
        public Input<Inputs.CPUInfoArgs>? CpuInfo { get; set; }

        /// <summary>
        /// Memory related environment info
        /// </summary>
        [Input("memoryInfo")]
        public Input<Inputs.MemoryInfoArgs>? MemoryInfo { get; set; }

        public PerfEnvironmentArgs()
        {
        }
    }
}