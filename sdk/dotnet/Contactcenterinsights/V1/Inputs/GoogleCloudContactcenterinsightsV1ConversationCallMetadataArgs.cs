// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Inputs
{

    /// <summary>
    /// Call-specific metadata.
    /// </summary>
    public sealed class GoogleCloudContactcenterinsightsV1ConversationCallMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The audio channel that contains the agent.
        /// </summary>
        [Input("agentChannel")]
        public Input<int>? AgentChannel { get; set; }

        /// <summary>
        /// The audio channel that contains the customer.
        /// </summary>
        [Input("customerChannel")]
        public Input<int>? CustomerChannel { get; set; }

        public GoogleCloudContactcenterinsightsV1ConversationCallMetadataArgs()
        {
        }
        public static new GoogleCloudContactcenterinsightsV1ConversationCallMetadataArgs Empty => new GoogleCloudContactcenterinsightsV1ConversationCallMetadataArgs();
    }
}
