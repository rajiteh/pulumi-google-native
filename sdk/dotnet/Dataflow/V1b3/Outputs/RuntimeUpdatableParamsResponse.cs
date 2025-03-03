// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Additional job parameters that can only be updated during runtime using the projects.jobs.update method. These fields have no effect when specified during job creation.
    /// </summary>
    [OutputType]
    public sealed class RuntimeUpdatableParamsResponse
    {
        /// <summary>
        /// The maximum number of workers to cap autoscaling at. This field is currently only supported for Streaming Engine jobs.
        /// </summary>
        public readonly int MaxNumWorkers;
        /// <summary>
        /// The minimum number of workers to scale down to. This field is currently only supported for Streaming Engine jobs.
        /// </summary>
        public readonly int MinNumWorkers;

        [OutputConstructor]
        private RuntimeUpdatableParamsResponse(
            int maxNumWorkers,

            int minNumWorkers)
        {
            MaxNumWorkers = maxNumWorkers;
            MinNumWorkers = minNumWorkers;
        }
    }
}
