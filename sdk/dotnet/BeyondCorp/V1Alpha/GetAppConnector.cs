// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha
{
    public static class GetAppConnector
    {
        /// <summary>
        /// Gets details of a single AppConnector.
        /// </summary>
        public static Task<GetAppConnectorResult> InvokeAsync(GetAppConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppConnectorResult>("google-native:beyondcorp/v1alpha:getAppConnector", args ?? new GetAppConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single AppConnector.
        /// </summary>
        public static Output<GetAppConnectorResult> Invoke(GetAppConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppConnectorResult>("google-native:beyondcorp/v1alpha:getAppConnector", args ?? new GetAppConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppConnectorArgs : global::Pulumi.InvokeArgs
    {
        [Input("appConnectorId", required: true)]
        public string AppConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAppConnectorArgs()
        {
        }
        public static new GetAppConnectorArgs Empty => new GetAppConnectorArgs();
    }

    public sealed class GetAppConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appConnectorId", required: true)]
        public Input<string> AppConnectorId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAppConnectorInvokeArgs()
        {
        }
        public static new GetAppConnectorInvokeArgs Empty => new GetAppConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppConnectorResult
    {
        /// <summary>
        /// Timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. An arbitrary user-provided name for the AppConnector. Cannot exceed 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Resource labels to represent user provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Unique resource name of the AppConnector. The name is ignored when creating a AppConnector.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Principal information about the Identity of the AppConnector.
        /// </summary>
        public readonly Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponse PrincipalInfo;
        /// <summary>
        /// Optional. Resource info of the connector.
        /// </summary>
        public readonly Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse ResourceInfo;
        /// <summary>
        /// The current state of the AppConnector.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A unique identifier for the instance generated by the system.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Timestamp when the resource was last modified.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAppConnectorResult(
            string createTime,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaAppConnectorPrincipalInfoResponse principalInfo,

            Outputs.GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoResponse resourceInfo,

            string state,

            string uid,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            PrincipalInfo = principalInfo;
            ResourceInfo = resourceInfo;
            State = state;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
