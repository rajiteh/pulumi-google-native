// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// Job status.
    /// </summary>
    [OutputType]
    public sealed class JobStatusResponse
    {
        /// <summary>
        /// The duration of time that the Job spent in status RUNNING.
        /// </summary>
        public readonly string RunDuration;
        /// <summary>
        /// Job state
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Job status events
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusEventResponse> StatusEvents;
        /// <summary>
        /// Aggregated task status for each TaskGroup in the Job. The map key is TaskGroup ID.
        /// </summary>
        public readonly ImmutableDictionary<string, string> TaskGroups;

        [OutputConstructor]
        private JobStatusResponse(
            string runDuration,

            string state,

            ImmutableArray<Outputs.StatusEventResponse> statusEvents,

            ImmutableDictionary<string, string> taskGroups)
        {
            RunDuration = runDuration;
            State = state;
            StatusEvents = statusEvents;
            TaskGroups = taskGroups;
        }
    }
}
