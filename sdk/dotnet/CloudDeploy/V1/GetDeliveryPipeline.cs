// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1
{
    public static class GetDeliveryPipeline
    {
        /// <summary>
        /// Gets details of a single DeliveryPipeline.
        /// </summary>
        public static Task<GetDeliveryPipelineResult> InvokeAsync(GetDeliveryPipelineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeliveryPipelineResult>("google-native:clouddeploy/v1:getDeliveryPipeline", args ?? new GetDeliveryPipelineArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single DeliveryPipeline.
        /// </summary>
        public static Output<GetDeliveryPipelineResult> Invoke(GetDeliveryPipelineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeliveryPipelineResult>("google-native:clouddeploy/v1:getDeliveryPipeline", args ?? new GetDeliveryPipelineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeliveryPipelineArgs : global::Pulumi.InvokeArgs
    {
        [Input("deliveryPipelineId", required: true)]
        public string DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDeliveryPipelineArgs()
        {
        }
        public static new GetDeliveryPipelineArgs Empty => new GetDeliveryPipelineArgs();
    }

    public sealed class GetDeliveryPipelineInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("deliveryPipelineId", required: true)]
        public Input<string> DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDeliveryPipelineInvokeArgs()
        {
        }
        public static new GetDeliveryPipelineInvokeArgs Empty => new GetDeliveryPipelineInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeliveryPipelineResult
    {
        /// <summary>
        /// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Information around the state of the Delivery Pipeline.
        /// </summary>
        public readonly Outputs.PipelineConditionResponse Condition;
        /// <summary>
        /// Time at which the pipeline was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of the `DeliveryPipeline`. Max length is 255 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
        /// </summary>
        public readonly Outputs.SerialPipelineResponse SerialPipeline;
        /// <summary>
        /// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
        /// </summary>
        public readonly bool Suspended;
        /// <summary>
        /// Unique identifier of the `DeliveryPipeline`.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Most recent time at which the pipeline was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDeliveryPipelineResult(
            ImmutableDictionary<string, string> annotations,

            Outputs.PipelineConditionResponse condition,

            string createTime,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.SerialPipelineResponse serialPipeline,

            bool suspended,

            string uid,

            string updateTime)
        {
            Annotations = annotations;
            Condition = condition;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            SerialPipeline = serialPipeline;
            Suspended = suspended;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}
