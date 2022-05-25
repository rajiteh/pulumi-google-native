// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Outputs
{

    /// <summary>
    /// Configuration for repricing a Google bill over a period of time.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudChannelV1RepricingConfigResponse
    {
        /// <summary>
        /// Information about the adjustment.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1RepricingAdjustmentResponse Adjustment;
        /// <summary>
        /// Applies the repricing configuration at the channel partner level. This is the only supported value for ChannelPartnerRepricingConfig.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1RepricingConfigChannelPartnerGranularityResponse ChannelPartnerGranularity;
        /// <summary>
        /// The YearMonth when these adjustments activate. The Day field needs to be "0" since we only accept YearMonth repricing boundaries.
        /// </summary>
        public readonly Outputs.GoogleTypeDateResponse EffectiveInvoiceMonth;
        /// <summary>
        /// Applies the repricing configuration at the entitlement level. This is the only supported value for CustomerRepricingConfig.
        /// </summary>
        public readonly Outputs.GoogleCloudChannelV1RepricingConfigEntitlementGranularityResponse EntitlementGranularity;
        /// <summary>
        /// The RebillingBasis to use for this bill. Specifies the relative cost based on repricing costs you will apply.
        /// </summary>
        public readonly string RebillingBasis;

        [OutputConstructor]
        private GoogleCloudChannelV1RepricingConfigResponse(
            Outputs.GoogleCloudChannelV1RepricingAdjustmentResponse adjustment,

            Outputs.GoogleCloudChannelV1RepricingConfigChannelPartnerGranularityResponse channelPartnerGranularity,

            Outputs.GoogleTypeDateResponse effectiveInvoiceMonth,

            Outputs.GoogleCloudChannelV1RepricingConfigEntitlementGranularityResponse entitlementGranularity,

            string rebillingBasis)
        {
            Adjustment = adjustment;
            ChannelPartnerGranularity = channelPartnerGranularity;
            EffectiveInvoiceMonth = effectiveInvoiceMonth;
            EntitlementGranularity = entitlementGranularity;
            RebillingBasis = rebillingBasis;
        }
    }
}