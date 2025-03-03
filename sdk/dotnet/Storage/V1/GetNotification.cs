// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1
{
    public static class GetNotification
    {
        /// <summary>
        /// View a notification configuration.
        /// </summary>
        public static Task<GetNotificationResult> InvokeAsync(GetNotificationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNotificationResult>("google-native:storage/v1:getNotification", args ?? new GetNotificationArgs(), options.WithDefaults());

        /// <summary>
        /// View a notification configuration.
        /// </summary>
        public static Output<GetNotificationResult> Invoke(GetNotificationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNotificationResult>("google-native:storage/v1:getNotification", args ?? new GetNotificationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationArgs : global::Pulumi.InvokeArgs
    {
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        [Input("notification", required: true)]
        public string Notification { get; set; } = null!;

        [Input("userProject")]
        public string? UserProject { get; set; }

        public GetNotificationArgs()
        {
        }
        public static new GetNotificationArgs Empty => new GetNotificationArgs();
    }

    public sealed class GetNotificationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("notification", required: true)]
        public Input<string> Notification { get; set; } = null!;

        [Input("userProject")]
        public Input<string>? UserProject { get; set; }

        public GetNotificationInvokeArgs()
        {
        }
        public static new GetNotificationInvokeArgs Empty => new GetNotificationInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotificationResult
    {
        /// <summary>
        /// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
        /// </summary>
        public readonly ImmutableDictionary<string, string> CustomAttributes;
        /// <summary>
        /// HTTP 1.1 Entity tag for this subscription notification.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
        /// </summary>
        public readonly ImmutableArray<string> EventTypes;
        /// <summary>
        /// The kind of item this is. For notifications, this is always storage#notification.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// If present, only apply this notification configuration to object names that begin with this prefix.
        /// </summary>
        public readonly string ObjectNamePrefix;
        /// <summary>
        /// The desired content of the Payload.
        /// </summary>
        public readonly string PayloadFormat;
        /// <summary>
        /// The canonical URL of this notification.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private GetNotificationResult(
            ImmutableDictionary<string, string> customAttributes,

            string etag,

            ImmutableArray<string> eventTypes,

            string kind,

            string objectNamePrefix,

            string payloadFormat,

            string selfLink,

            string topic)
        {
            CustomAttributes = customAttributes;
            Etag = etag;
            EventTypes = eventTypes;
            Kind = kind;
            ObjectNamePrefix = objectNamePrefix;
            PayloadFormat = payloadFormat;
            SelfLink = selfLink;
            Topic = topic;
        }
    }
}
