// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha
{
    public static class GetSfdcChannel
    {
        /// <summary>
        /// Gets an sfdc channel. If the channel doesn't exist, Code.NOT_FOUND exception will be thrown.
        /// </summary>
        public static Task<GetSfdcChannelResult> InvokeAsync(GetSfdcChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSfdcChannelResult>("google-native:integrations/v1alpha:getSfdcChannel", args ?? new GetSfdcChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an sfdc channel. If the channel doesn't exist, Code.NOT_FOUND exception will be thrown.
        /// </summary>
        public static Output<GetSfdcChannelResult> Invoke(GetSfdcChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSfdcChannelResult>("google-native:integrations/v1alpha:getSfdcChannel", args ?? new GetSfdcChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSfdcChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sfdcChannelId", required: true)]
        public string SfdcChannelId { get; set; } = null!;

        [Input("sfdcInstanceId", required: true)]
        public string SfdcInstanceId { get; set; } = null!;

        public GetSfdcChannelArgs()
        {
        }
        public static new GetSfdcChannelArgs Empty => new GetSfdcChannelArgs();
    }

    public sealed class GetSfdcChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sfdcChannelId", required: true)]
        public Input<string> SfdcChannelId { get; set; } = null!;

        [Input("sfdcInstanceId", required: true)]
        public Input<string> SfdcInstanceId { get; set; } = null!;

        public GetSfdcChannelInvokeArgs()
        {
        }
        public static new GetSfdcChannelInvokeArgs Empty => new GetSfdcChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetSfdcChannelResult
    {
        /// <summary>
        /// The Channel topic defined by salesforce once an channel is opened
        /// </summary>
        public readonly string ChannelTopic;
        /// <summary>
        /// Time when the channel is created
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Time when the channel was deleted. Empty if not deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// The description for this channel
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Client level unique name/alias to easily reference a channel.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Indicated if a channel has any active integrations referencing it. Set to false when the channel is created, and set to true if there is any integration published with the channel configured in it.
        /// </summary>
        public readonly bool IsActive;
        /// <summary>
        /// Last sfdc messsage replay id for channel
        /// </summary>
        public readonly string LastReplayId;
        /// <summary>
        /// Resource name of the SFDC channel projects/{project}/locations/{location}/sfdcInstances/{sfdc_instance}/sfdcChannels/{sfdc_channel}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Time when the channel was last updated
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetSfdcChannelResult(
            string channelTopic,

            string createTime,

            string deleteTime,

            string description,

            string displayName,

            bool isActive,

            string lastReplayId,

            string name,

            string updateTime)
        {
            ChannelTopic = channelTopic;
            CreateTime = createTime;
            DeleteTime = deleteTime;
            Description = description;
            DisplayName = displayName;
            IsActive = isActive;
            LastReplayId = lastReplayId;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}