// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Dataflow.V1b3.Outputs
{

    [OutputType]
    public sealed class SpannerIODetailsResponse
    {
        /// <summary>
        /// DatabaseId accessed in the connection.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// InstanceId accessed in the connection.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// ProjectId accessed in the connection.
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private SpannerIODetailsResponse(
            string databaseId,

            string instanceId,

            string projectId)
        {
            DatabaseId = databaseId;
            InstanceId = instanceId;
            ProjectId = projectId;
        }
    }
}