// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class QueryTimelineSampleResponse
    {
        /// <summary>
        /// Total number of units currently being processed by workers. This does not correspond directly to slot usage. This is the largest value observed since the last sample.
        /// </summary>
        public readonly string ActiveUnits;
        /// <summary>
        /// Total parallel units of work completed by this query.
        /// </summary>
        public readonly string CompletedUnits;
        /// <summary>
        /// Milliseconds elapsed since the start of query execution.
        /// </summary>
        public readonly string ElapsedMs;
        /// <summary>
        /// Units of work that can be scheduled immediately. Providing additional slots for these units of work will speed up the query, provided no other query in the reservation needs additional slots.
        /// </summary>
        public readonly string EstimatedRunnableUnits;
        /// <summary>
        /// Total units of work remaining for the query. This number can be revised (increased or decreased) while the query is running.
        /// </summary>
        public readonly string PendingUnits;
        /// <summary>
        /// Cumulative slot-ms consumed by the query.
        /// </summary>
        public readonly string TotalSlotMs;

        [OutputConstructor]
        private QueryTimelineSampleResponse(
            string activeUnits,

            string completedUnits,

            string elapsedMs,

            string estimatedRunnableUnits,

            string pendingUnits,

            string totalSlotMs)
        {
            ActiveUnits = activeUnits;
            CompletedUnits = completedUnits;
            ElapsedMs = elapsedMs;
            EstimatedRunnableUnits = estimatedRunnableUnits;
            PendingUnits = pendingUnits;
            TotalSlotMs = totalSlotMs;
        }
    }
}
