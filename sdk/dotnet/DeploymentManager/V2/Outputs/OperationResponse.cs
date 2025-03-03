// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2.Outputs
{

    /// <summary>
    /// Represents an Operation resource. Google Compute Engine has three Operation resources: * [Global](/compute/docs/reference/rest/{$api_version}/globalOperations) * [Regional](/compute/docs/reference/rest/{$api_version}/regionOperations) * [Zonal](/compute/docs/reference/rest/{$api_version}/zoneOperations) You can use an operation resource to manage asynchronous API requests. For more information, read Handling API responses. Operations can be global, regional or zonal. - For global operations, use the `globalOperations` resource. - For regional operations, use the `regionOperations` resource. - For zonal operations, use the `zonalOperations` resource. For more information, read Global, Regional, and Zonal Resources.
    /// </summary>
    [OutputType]
    public sealed class OperationResponse
    {
        /// <summary>
        /// The value of `requestId` if you provided it in the request. Not present otherwise.
        /// </summary>
        public readonly string ClientOperationId;
        /// <summary>
        /// [Deprecated] This field is deprecated.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// A textual description of the operation, which is set when the operation is created.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The time that this operation was completed. This value is in RFC3339 text format.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// If errors are generated during processing of the operation, this field will be populated.
        /// </summary>
        public readonly Outputs.OperationErrorResponse Error;
        /// <summary>
        /// If the operation fails, this field contains the HTTP error message that was returned, such as `NOT FOUND`.
        /// </summary>
        public readonly string HttpErrorMessage;
        /// <summary>
        /// If the operation fails, this field contains the HTTP error status code that was returned. For example, a `404` means the resource was not found.
        /// </summary>
        public readonly int HttpErrorStatusCode;
        /// <summary>
        /// The time that this operation was requested. This value is in RFC3339 text format.
        /// </summary>
        public readonly string InsertTime;
        /// <summary>
        /// Type of the resource. Always `compute#operation` for Operation resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the operation.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An ID that represents a group of operations, such as when a group of operations results from a `bulkInsert` API request.
        /// </summary>
        public readonly string OperationGroupId;
        /// <summary>
        /// The type of operation, such as `insert`, `update`, or `delete`, and so on.
        /// </summary>
        public readonly string OperationType;
        /// <summary>
        /// An optional progress indicator that ranges from 0 to 100. There is no requirement that this be linear or support any granularity of operations. This should not be used to guess when the operation will be complete. This number should monotonically increase as the operation progresses.
        /// </summary>
        public readonly int Progress;
        /// <summary>
        /// The URL of the region where the operation resides. Only applicable when performing regional operations.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The time that this operation was started by the server. This value is in RFC3339 text format.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The status of the operation, which can be one of the following: `PENDING`, `RUNNING`, or `DONE`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An optional textual description of the current status of the operation.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// The unique target ID, which identifies a specific incarnation of the target resource.
        /// </summary>
        public readonly string TargetId;
        /// <summary>
        /// The URL of the resource that the operation modifies. For operations related to creating a snapshot, this points to the persistent disk that the snapshot was created from.
        /// </summary>
        public readonly string TargetLink;
        /// <summary>
        /// User who requested the operation, for example: `user@example.com`.
        /// </summary>
        public readonly string User;
        /// <summary>
        /// If warning messages are generated during processing of the operation, this field will be populated.
        /// </summary>
        public readonly ImmutableArray<Outputs.OperationWarningsItemResponse> Warnings;
        /// <summary>
        /// The URL of the zone where the operation resides. Only applicable when performing per-zone operations.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private OperationResponse(
            string clientOperationId,

            string creationTimestamp,

            string description,

            string endTime,

            Outputs.OperationErrorResponse error,

            string httpErrorMessage,

            int httpErrorStatusCode,

            string insertTime,

            string kind,

            string name,

            string operationGroupId,

            string operationType,

            int progress,

            string region,

            string selfLink,

            string startTime,

            string status,

            string statusMessage,

            string targetId,

            string targetLink,

            string user,

            ImmutableArray<Outputs.OperationWarningsItemResponse> warnings,

            string zone)
        {
            ClientOperationId = clientOperationId;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EndTime = endTime;
            Error = error;
            HttpErrorMessage = httpErrorMessage;
            HttpErrorStatusCode = httpErrorStatusCode;
            InsertTime = insertTime;
            Kind = kind;
            Name = name;
            OperationGroupId = operationGroupId;
            OperationType = operationType;
            Progress = progress;
            Region = region;
            SelfLink = selfLink;
            StartTime = startTime;
            Status = status;
            StatusMessage = statusMessage;
            TargetId = targetId;
            TargetLink = targetLink;
            User = user;
            Warnings = warnings;
            Zone = zone;
        }
    }
}
