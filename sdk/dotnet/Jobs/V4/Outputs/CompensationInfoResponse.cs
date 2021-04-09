// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Jobs.V4.Outputs
{

    [OutputType]
    public sealed class CompensationInfoResponse
    {
        /// <summary>
        /// Annualized base compensation range. Computed as base compensation entry's CompensationEntry.amount times CompensationEntry.expected_units_per_year. See CompensationEntry for explanation on compensation annualization.
        /// </summary>
        public readonly Outputs.CompensationRangeResponse AnnualizedBaseCompensationRange;
        /// <summary>
        /// Annualized total compensation range. Computed as all compensation entries' CompensationEntry.amount times CompensationEntry.expected_units_per_year. See CompensationEntry for explanation on compensation annualization.
        /// </summary>
        public readonly Outputs.CompensationRangeResponse AnnualizedTotalCompensationRange;
        /// <summary>
        /// Job compensation information. At most one entry can be of type CompensationInfo.CompensationType.BASE, which is referred as **base compensation entry** for the job.
        /// </summary>
        public readonly ImmutableArray<Outputs.CompensationEntryResponse> Entries;

        [OutputConstructor]
        private CompensationInfoResponse(
            Outputs.CompensationRangeResponse annualizedBaseCompensationRange,

            Outputs.CompensationRangeResponse annualizedTotalCompensationRange,

            ImmutableArray<Outputs.CompensationEntryResponse> entries)
        {
            AnnualizedBaseCompensationRange = annualizedBaseCompensationRange;
            AnnualizedTotalCompensationRange = annualizedTotalCompensationRange;
            Entries = entries;
        }
    }
}