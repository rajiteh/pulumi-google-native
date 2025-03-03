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
    /// A TaskGroup contains one or multiple Tasks that share the same Runnable but with different runtime parameters.
    /// </summary>
    [OutputType]
    public sealed class TaskGroupResponse
    {
        /// <summary>
        /// TaskGroup name. The system generates this field based on parent Job name. For example: "projects/123456/locations/us-west1/jobs/job01/taskGroups/group01".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Max number of tasks that can run in parallel. Default to min(task_count, 1000). Field parallelism must be 1 if the scheduling_policy is IN_ORDER.
        /// </summary>
        public readonly string Parallelism;
        /// <summary>
        /// When true, Batch will configure SSH to allow passwordless login between VMs running the Batch tasks in the same TaskGroup.
        /// </summary>
        public readonly bool PermissiveSsh;
        /// <summary>
        /// When true, Batch will populate a file with a list of all VMs assigned to the TaskGroup and set the BATCH_HOSTS_FILE environment variable to the path of that file. Defaults to false.
        /// </summary>
        public readonly bool RequireHostsFile;
        /// <summary>
        /// Number of Tasks in the TaskGroup. Default is 1.
        /// </summary>
        public readonly string TaskCount;
        /// <summary>
        /// Max number of tasks that can be run on a VM at the same time. If not specified, the system will decide a value based on available compute resources on a VM and task requirements.
        /// </summary>
        public readonly string TaskCountPerNode;
        /// <summary>
        /// An array of environment variable mappings, which are passed to Tasks with matching indices. If task_environments is used then task_count should not be specified in the request (and will be ignored). Task count will be the length of task_environments. Tasks get a BATCH_TASK_INDEX and BATCH_TASK_COUNT environment variable, in addition to any environment variables set in task_environments, specifying the number of Tasks in the Task's parent TaskGroup, and the specific Task's index in the TaskGroup (0 through BATCH_TASK_COUNT - 1). task_environments supports up to 200 entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvironmentResponse> TaskEnvironments;
        /// <summary>
        /// Tasks in the group share the same task spec.
        /// </summary>
        public readonly Outputs.TaskSpecResponse TaskSpec;

        [OutputConstructor]
        private TaskGroupResponse(
            string name,

            string parallelism,

            bool permissiveSsh,

            bool requireHostsFile,

            string taskCount,

            string taskCountPerNode,

            ImmutableArray<Outputs.EnvironmentResponse> taskEnvironments,

            Outputs.TaskSpecResponse taskSpec)
        {
            Name = name;
            Parallelism = parallelism;
            PermissiveSsh = permissiveSsh;
            RequireHostsFile = requireHostsFile;
            TaskCount = taskCount;
            TaskCountPerNode = taskCountPerNode;
            TaskEnvironments = taskEnvironments;
            TaskSpec = taskSpec;
        }
    }
}
