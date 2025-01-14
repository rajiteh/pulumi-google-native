// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Additional details for quota exceeded error for resource quota.
    /// </summary>
    [OutputType]
    public sealed class QuotaExceededInfoResponse
    {
        /// <summary>
        /// The map holding related quota dimensions.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Dimensions;
        /// <summary>
        /// Current effective quota limit. The limit's unit depends on the quota type or metric.
        /// </summary>
        public readonly double Limit;
        /// <summary>
        /// The name of the quota limit.
        /// </summary>
        public readonly string LimitName;
        /// <summary>
        /// The Compute Engine quota metric name.
        /// </summary>
        public readonly string MetricName;

        [OutputConstructor]
        private QuotaExceededInfoResponse(
            ImmutableDictionary<string, string> dimensions,

            double limit,

            string limitName,

            string metricName)
        {
            Dimensions = dimensions;
            Limit = limit;
            LimitName = limitName;
            MetricName = metricName;
        }
    }
}
