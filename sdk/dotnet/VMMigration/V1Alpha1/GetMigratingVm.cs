// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    public static class GetMigratingVm
    {
        /// <summary>
        /// Gets details of a single MigratingVm.
        /// </summary>
        public static Task<GetMigratingVmResult> InvokeAsync(GetMigratingVmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMigratingVmResult>("google-native:vmmigration/v1alpha1:getMigratingVm", args ?? new GetMigratingVmArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single MigratingVm.
        /// </summary>
        public static Output<GetMigratingVmResult> Invoke(GetMigratingVmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMigratingVmResult>("google-native:vmmigration/v1alpha1:getMigratingVm", args ?? new GetMigratingVmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMigratingVmArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public string MigratingVmId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetMigratingVmArgs()
        {
        }
        public static new GetMigratingVmArgs Empty => new GetMigratingVmArgs();
    }

    public sealed class GetMigratingVmInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public Input<string> MigratingVmId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetMigratingVmInvokeArgs()
        {
        }
        public static new GetMigratingVmInvokeArgs Empty => new GetMigratingVmInvokeArgs();
    }


    [OutputType]
    public sealed class GetMigratingVmResult
    {
        /// <summary>
        /// Details of the VM from an AWS source.
        /// </summary>
        public readonly Outputs.AwsSourceVmDetailsResponse AwsSourceVmDetails;
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        public readonly Outputs.ComputeEngineTargetDefaultsResponse ComputeEngineTargetDefaults;
        /// <summary>
        /// Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_defaults instead.
        /// </summary>
        public readonly Outputs.TargetVMDetailsResponse ComputeEngineVmDefaults;
        /// <summary>
        /// The time the migrating VM was created (this refers to this resource and not to the time it was installed in the source).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Details of the current running replication cycle.
        /// </summary>
        public readonly Outputs.ReplicationCycleResponse CurrentSyncInfo;
        /// <summary>
        /// Provides details of future CutoverJobs of a MigratingVm. Set to empty when cutover forecast is unavailable.
        /// </summary>
        public readonly Outputs.CutoverForecastResponse CutoverForecast;
        /// <summary>
        /// The description attached to the migrating VM by the user.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name attached to the MigratingVm by the user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Provides details on the state of the Migrating VM in case of an error in replication.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The group this migrating vm is included in, if any. The group is represented by the full path of the appropriate Group resource.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The labels of the migrating VM.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Details of the last replication cycle. This will be updated whenever a replication cycle is finished and is not to be confused with last_sync which is only updated on successful replication cycles.
        /// </summary>
        public readonly Outputs.ReplicationCycleResponse LastReplicationCycle;
        /// <summary>
        /// The most updated snapshot created time in the source that finished replication.
        /// </summary>
        public readonly Outputs.ReplicationSyncResponse LastSync;
        /// <summary>
        /// The identifier of the MigratingVm.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The replication schedule policy.
        /// </summary>
        public readonly Outputs.SchedulePolicyResponse Policy;
        /// <summary>
        /// The recent clone jobs performed on the migrating VM. This field holds the vm's last completed clone job and the vm's running clone job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloneJobResponse> RecentCloneJobs;
        /// <summary>
        /// The recent cutover jobs performed on the migrating VM. This field holds the vm's last completed cutover job and the vm's running cutover job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
        /// </summary>
        public readonly ImmutableArray<Outputs.CutoverJobResponse> RecentCutoverJobs;
        /// <summary>
        /// The unique ID of the VM in the source. The VM's name in vSphere can be changed, so this is not the VM's name but rather its moRef id. This id is of the form vm-.
        /// </summary>
        public readonly string SourceVmId;
        /// <summary>
        /// State of the MigratingVm.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The last time the migrating VM state was updated.
        /// </summary>
        public readonly string StateTime;
        /// <summary>
        /// The default configuration of the target VM that will be created in Google Cloud as a result of the migration. Deprecated: Use compute_engine_target_defaults instead.
        /// </summary>
        public readonly Outputs.TargetVMDetailsResponse TargetDefaults;
        /// <summary>
        /// The last time the migrating VM resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetMigratingVmResult(
            Outputs.AwsSourceVmDetailsResponse awsSourceVmDetails,

            Outputs.ComputeEngineTargetDefaultsResponse computeEngineTargetDefaults,

            Outputs.TargetVMDetailsResponse computeEngineVmDefaults,

            string createTime,

            Outputs.ReplicationCycleResponse currentSyncInfo,

            Outputs.CutoverForecastResponse cutoverForecast,

            string description,

            string displayName,

            Outputs.StatusResponse error,

            string group,

            ImmutableDictionary<string, string> labels,

            Outputs.ReplicationCycleResponse lastReplicationCycle,

            Outputs.ReplicationSyncResponse lastSync,

            string name,

            Outputs.SchedulePolicyResponse policy,

            ImmutableArray<Outputs.CloneJobResponse> recentCloneJobs,

            ImmutableArray<Outputs.CutoverJobResponse> recentCutoverJobs,

            string sourceVmId,

            string state,

            string stateTime,

            Outputs.TargetVMDetailsResponse targetDefaults,

            string updateTime)
        {
            AwsSourceVmDetails = awsSourceVmDetails;
            ComputeEngineTargetDefaults = computeEngineTargetDefaults;
            ComputeEngineVmDefaults = computeEngineVmDefaults;
            CreateTime = createTime;
            CurrentSyncInfo = currentSyncInfo;
            CutoverForecast = cutoverForecast;
            Description = description;
            DisplayName = displayName;
            Error = error;
            Group = group;
            Labels = labels;
            LastReplicationCycle = lastReplicationCycle;
            LastSync = lastSync;
            Name = name;
            Policy = policy;
            RecentCloneJobs = recentCloneJobs;
            RecentCutoverJobs = recentCutoverJobs;
            SourceVmId = sourceVmId;
            State = state;
            StateTime = stateTime;
            TargetDefaults = targetDefaults;
            UpdateTime = updateTime;
        }
    }
}
