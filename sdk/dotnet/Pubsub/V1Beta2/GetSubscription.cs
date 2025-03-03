// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1Beta2
{
    public static class GetSubscription
    {
        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Task<GetSubscriptionResult> InvokeAsync(GetSubscriptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("google-native:pubsub/v1beta2:getSubscription", args ?? new GetSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the configuration details of a subscription.
        /// </summary>
        public static Output<GetSubscriptionResult> Invoke(GetSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriptionResult>("google-native:pubsub/v1beta2:getSubscription", args ?? new GetSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("subscriptionId", required: true)]
        public string SubscriptionId { get; set; } = null!;

        public GetSubscriptionArgs()
        {
        }
        public static new GetSubscriptionArgs Empty => new GetSubscriptionArgs();
    }

    public sealed class GetSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        public GetSubscriptionInvokeArgs()
        {
        }
        public static new GetSubscriptionInvokeArgs Empty => new GetSubscriptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriptionResult
    {
        /// <summary>
        /// This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using pull. The maximum custom deadline you can specify is 600 seconds (10 minutes). For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message. If this parameter is 0, a default value of 10 seconds is used.
        /// </summary>
        public readonly int AckDeadlineSeconds;
        /// <summary>
        /// The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If push delivery is used with this subscription, this field is used to configure it. An empty `pushConfig` signifies that the subscriber will pull and ack messages using API methods.
        /// </summary>
        public readonly Outputs.PushConfigResponse PushConfig;
        /// <summary>
        /// The name of the topic from which this subscription is receiving messages. The value of this field will be `_deleted-topic_` if the topic has been deleted.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private GetSubscriptionResult(
            int ackDeadlineSeconds,

            string name,

            Outputs.PushConfigResponse pushConfig,

            string topic)
        {
            AckDeadlineSeconds = ackDeadlineSeconds;
            Name = name;
            PushConfig = pushConfig;
            Topic = topic;
        }
    }
}
