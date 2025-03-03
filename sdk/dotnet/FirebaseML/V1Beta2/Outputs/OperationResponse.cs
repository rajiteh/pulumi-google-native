// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseML.V1Beta2.Outputs
{

    /// <summary>
    /// This resource represents a long-running operation that is the result of a network API call.
    /// </summary>
    [OutputType]
    public sealed class OperationResponse
    {
        /// <summary>
        /// If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
        /// </summary>
        public readonly bool Done;
        /// <summary>
        /// The error result of the operation in case of failure or cancellation.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The normal response of the operation in case of success. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Response;

        [OutputConstructor]
        private OperationResponse(
            bool done,

            Outputs.StatusResponse error,

            ImmutableDictionary<string, string> metadata,

            string name,

            ImmutableDictionary<string, string> response)
        {
            Done = done;
            Error = error;
            Metadata = metadata;
            Name = name;
            Response = response;
        }
    }
}
