// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.File.V1
{
    public static class GetInstance
    {
        /// <summary>
        /// Gets the details of a specific instance.
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("google-native:file/v1:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a specific instance.
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("google-native:file/v1:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// The time when the instance was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the instance (2048 characters or less).
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// File system shares on the instance. For this version, only a single file share is supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.FileShareConfigResponse> FileShares;
        /// <summary>
        /// KMS key name used for data encryption.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// Resource labels to represent user provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the instance, in the format `projects/{project}/locations/{location}/instances/{instance}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// VPC networks to which the instance is connected. For this version, only a single network is supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkConfigResponse> Networks;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// The instance state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Additional information about the instance state, if available.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// Field indicates all the reasons the instance is in "SUSPENDED" state.
        /// </summary>
        public readonly ImmutableArray<string> SuspensionReasons;
        /// <summary>
        /// The service tier of the instance.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private GetInstanceResult(
            string createTime,

            string description,

            string etag,

            ImmutableArray<Outputs.FileShareConfigResponse> fileShares,

            string kmsKeyName,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.NetworkConfigResponse> networks,

            bool satisfiesPzs,

            string state,

            string statusMessage,

            ImmutableArray<string> suspensionReasons,

            string tier)
        {
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            FileShares = fileShares;
            KmsKeyName = kmsKeyName;
            Labels = labels;
            Name = name;
            Networks = networks;
            SatisfiesPzs = satisfiesPzs;
            State = state;
            StatusMessage = statusMessage;
            SuspensionReasons = suspensionReasons;
            Tier = tier;
        }
    }
}
