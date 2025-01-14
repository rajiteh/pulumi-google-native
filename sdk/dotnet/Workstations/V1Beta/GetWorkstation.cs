// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta
{
    public static class GetWorkstation
    {
        /// <summary>
        /// Returns the requested workstation.
        /// </summary>
        public static Task<GetWorkstationResult> InvokeAsync(GetWorkstationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkstationResult>("google-native:workstations/v1beta:getWorkstation", args ?? new GetWorkstationArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the requested workstation.
        /// </summary>
        public static Output<GetWorkstationResult> Invoke(GetWorkstationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkstationResult>("google-native:workstations/v1beta:getWorkstation", args ?? new GetWorkstationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkstationArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("workstationClusterId", required: true)]
        public string WorkstationClusterId { get; set; } = null!;

        [Input("workstationConfigId", required: true)]
        public string WorkstationConfigId { get; set; } = null!;

        [Input("workstationId", required: true)]
        public string WorkstationId { get; set; } = null!;

        public GetWorkstationArgs()
        {
        }
        public static new GetWorkstationArgs Empty => new GetWorkstationArgs();
    }

    public sealed class GetWorkstationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("workstationClusterId", required: true)]
        public Input<string> WorkstationClusterId { get; set; } = null!;

        [Input("workstationConfigId", required: true)]
        public Input<string> WorkstationConfigId { get; set; } = null!;

        [Input("workstationId", required: true)]
        public Input<string> WorkstationId { get; set; } = null!;

        public GetWorkstationInvokeArgs()
        {
        }
        public static new GetWorkstationInvokeArgs Empty => new GetWorkstationInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkstationResult
    {
        /// <summary>
        /// Client-specified annotations.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Time when this resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Time when this resource was soft-deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// Human-readable name for this resource.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Environment variables passed to the workstation container's entrypoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Env;
        /// <summary>
        /// Checksum computed by the server. May be sent on update and delete requests to make sure that the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Full name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether this resource is currently being updated to match its intended state.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// Current state of the workstation.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A system-assigned unique identified for this resource.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Time when this resource was most recently updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetWorkstationResult(
            ImmutableDictionary<string, string> annotations,

            string createTime,

            string deleteTime,

            string displayName,

            ImmutableDictionary<string, string> env,

            string etag,

            string host,

            ImmutableDictionary<string, string> labels,

            string name,

            bool reconciling,

            string state,

            string uid,

            string updateTime)
        {
            Annotations = annotations;
            CreateTime = createTime;
            DeleteTime = deleteTime;
            DisplayName = displayName;
            Env = env;
            Etag = etag;
            Host = host;
            Labels = labels;
            Name = name;
            Reconciling = reconciling;
            State = state;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
