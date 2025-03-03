// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Binds the resources in an API proxy or remote service with the allowed REST methods and associated quota enforcement.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1OperationConfigResponse
    {
        /// <summary>
        /// Name of the API proxy or remote service with which the resources, methods, and quota are associated.
        /// </summary>
        public readonly string ApiSource;
        /// <summary>
        /// Custom attributes associated with the operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> Attributes;
        /// <summary>
        /// List of resource/method pairs for the API proxy or remote service to which quota will applied. **Note**: Currently, you can specify only a single resource/method pair. The call will fail if more than one resource/method pair is provided.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudApigeeV1OperationResponse> Operations;
        /// <summary>
        /// Quota parameters to be enforced for the resources, methods, and API source combination. If none are specified, quota enforcement will not be done.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1QuotaResponse Quota;

        [OutputConstructor]
        private GoogleCloudApigeeV1OperationConfigResponse(
            string apiSource,

            ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse> attributes,

            ImmutableArray<Outputs.GoogleCloudApigeeV1OperationResponse> operations,

            Outputs.GoogleCloudApigeeV1QuotaResponse quota)
        {
            ApiSource = apiSource;
            Attributes = attributes;
            Operations = operations;
            Quota = quota;
        }
    }
}
