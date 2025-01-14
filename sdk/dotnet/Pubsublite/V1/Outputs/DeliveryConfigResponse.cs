// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsublite.V1.Outputs
{

    /// <summary>
    /// The settings for a subscription's message delivery.
    /// </summary>
    [OutputType]
    public sealed class DeliveryConfigResponse
    {
        /// <summary>
        /// The DeliveryRequirement for this subscription.
        /// </summary>
        public readonly string DeliveryRequirement;

        [OutputConstructor]
        private DeliveryConfigResponse(string deliveryRequirement)
        {
            DeliveryRequirement = deliveryRequirement;
        }
    }
}
