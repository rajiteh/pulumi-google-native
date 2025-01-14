// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// ChildRollouts job composition
    /// </summary>
    [OutputType]
    public sealed class ChildRolloutJobsResponse
    {
        /// <summary>
        /// List of AdvanceChildRolloutJobs
        /// </summary>
        public readonly ImmutableArray<Outputs.JobResponse> AdvanceRolloutJobs;
        /// <summary>
        /// List of CreateChildRolloutJobs
        /// </summary>
        public readonly ImmutableArray<Outputs.JobResponse> CreateRolloutJobs;

        [OutputConstructor]
        private ChildRolloutJobsResponse(
            ImmutableArray<Outputs.JobResponse> advanceRolloutJobs,

            ImmutableArray<Outputs.JobResponse> createRolloutJobs)
        {
            AdvanceRolloutJobs = advanceRolloutJobs;
            CreateRolloutJobs = createRolloutJobs;
        }
    }
}
