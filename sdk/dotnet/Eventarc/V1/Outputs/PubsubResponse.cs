// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1.Outputs
{

    /// <summary>
    /// Represents a Pub/Sub transport.
    /// </summary>
    [OutputType]
    public sealed class PubsubResponse
    {
        /// <summary>
        /// The name of the Pub/Sub subscription created and managed by Eventarc as a transport for the event delivery. Format: `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`.
        /// </summary>
        public readonly string Subscription;
        /// <summary>
        /// Optional. The name of the Pub/Sub topic created and managed by Eventarc as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME}`. You can set an existing topic for triggers of the type `google.cloud.pubsub.topic.v1.messagePublished`. The topic you provide here is not deleted by Eventarc at trigger deletion.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private PubsubResponse(
            string subscription,

            string topic)
        {
            Subscription = subscription;
            Topic = topic;
        }
    }
}
