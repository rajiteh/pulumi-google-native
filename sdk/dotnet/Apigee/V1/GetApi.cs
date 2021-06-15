// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    public static class GetApi
    {
        /// <summary>
        /// Gets an API proxy including a list of existing revisions.
        /// </summary>
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("google-native:apigee/v1:getApi", args ?? new GetApiArgs(), options.WithVersion());
    }


    public sealed class GetApiArgs : Pulumi.InvokeArgs
    {
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetApiArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// The id of the most recently created revision for this api proxy.
        /// </summary>
        public readonly string LatestRevisionId;
        /// <summary>
        /// Metadata describing the API proxy.
        /// </summary>
        public readonly Outputs.GoogleCloudApigeeV1EntityMetadataResponse MetaData;
        /// <summary>
        /// Name of the API proxy.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of revisons defined for the API proxy.
        /// </summary>
        public readonly ImmutableArray<string> Revision;

        [OutputConstructor]
        private GetApiResult(
            string latestRevisionId,

            Outputs.GoogleCloudApigeeV1EntityMetadataResponse metaData,

            string name,

            ImmutableArray<string> revision)
        {
            LatestRevisionId = latestRevisionId;
            MetaData = metaData;
            Name = name;
            Revision = revision;
        }
    }
}