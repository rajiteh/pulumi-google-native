// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudTasks.V2Beta2.Outputs
{

    [OutputType]
    public sealed class TaskStatusResponse
    {
        /// <summary>
        /// The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
        /// </summary>
        public readonly int AttemptDispatchCount;
        /// <summary>
        /// The number of attempts which have received a response. This field is not calculated for pull tasks.
        /// </summary>
        public readonly int AttemptResponseCount;
        /// <summary>
        /// The status of the task's first attempt. Only dispatch_time will be set. The other AttemptStatus information is not retained by Cloud Tasks. This field is not calculated for pull tasks.
        /// </summary>
        public readonly Outputs.AttemptStatusResponse FirstAttemptStatus;
        /// <summary>
        /// The status of the task's last attempt. This field is not calculated for pull tasks.
        /// </summary>
        public readonly Outputs.AttemptStatusResponse LastAttemptStatus;

        [OutputConstructor]
        private TaskStatusResponse(
            int attemptDispatchCount,

            int attemptResponseCount,

            Outputs.AttemptStatusResponse firstAttemptStatus,

            Outputs.AttemptStatusResponse lastAttemptStatus)
        {
            AttemptDispatchCount = attemptDispatchCount;
            AttemptResponseCount = attemptResponseCount;
            FirstAttemptStatus = firstAttemptStatus;
            LastAttemptStatus = lastAttemptStatus;
        }
    }
}