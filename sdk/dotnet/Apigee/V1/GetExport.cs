// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetExport
    {
        /// <summary>
        /// Gets the details and status of an analytics export job. If the export job is still in progress, its `state` is set to "running". After the export job has completed successfully, its `state` is set to "completed". If the export job fails, its `state` is set to `failed`.
        /// </summary>
        public static Task<GetExportResult> InvokeAsync(GetExportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExportResult>("google-native:apigee/v1:getExport", args ?? new GetExportArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details and status of an analytics export job. If the export job is still in progress, its `state` is set to "running". After the export job has completed successfully, its `state` is set to "completed". If the export job fails, its `state` is set to `failed`.
        /// </summary>
        public static Output<GetExportResult> Invoke(GetExportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExportResult>("google-native:apigee/v1:getExport", args ?? new GetExportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExportArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("exportId", required: true)]
        public string ExportId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetExportArgs()
        {
        }
        public static new GetExportArgs Empty => new GetExportArgs();
    }

    public sealed class GetExportInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("exportId", required: true)]
        public Input<string> ExportId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetExportInvokeArgs()
        {
        }
        public static new GetExportInvokeArgs Empty => new GetExportInvokeArgs();
    }


    [OutputType]
    public sealed class GetExportResult
    {
        /// <summary>
        /// Time the export job was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// Name of the datastore that is the destination of the export job [datastore]
        /// </summary>
        public readonly string DatastoreName;
        /// <summary>
        /// Description of the export job.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Error is set when export fails
        /// </summary>
        public readonly string Error;
        /// <summary>
        /// Execution time for this export job. If the job is still in progress, it will be set to the amount of time that has elapsed since`created`, in seconds. Else, it will set to (`updated` - `created`), in seconds.
        /// </summary>
        public readonly string ExecutionTime;
        /// <summary>
        /// Display name of the export job.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Self link of the export job. A URI that can be used to retrieve the status of an export job. Example: `/organizations/myorg/environments/myenv/analytics/exports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
        /// </summary>
        public readonly string Self;
        /// <summary>
        /// Status of the export job. Valid values include `enqueued`, `running`, `completed`, and `failed`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Time the export job was last updated.
        /// </summary>
        public readonly string Updated;

        [OutputConstructor]
        private GetExportResult(
            string created,

            string datastoreName,

            string description,

            string error,

            string executionTime,

            string name,

            string self,

            string state,

            string updated)
        {
            Created = created;
            DatastoreName = datastoreName;
            Description = description;
            Error = error;
            ExecutionTime = executionTime;
            Name = name;
            Self = self;
            State = state;
            Updated = updated;
        }
    }
}
