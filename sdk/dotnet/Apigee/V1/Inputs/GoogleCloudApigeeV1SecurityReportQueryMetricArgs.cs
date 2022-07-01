// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Metric of the Query
    /// </summary>
    public sealed class GoogleCloudApigeeV1SecurityReportQueryMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Aggregation function: avg, min, max, or sum.
        /// </summary>
        [Input("aggregationFunction")]
        public Input<string>? AggregationFunction { get; set; }

        /// <summary>
        /// Alias for the metric. Alias will be used to replace metric name in query results.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Metric name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// One of `+`, `-`, `/`, `%`, `*`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// Operand value should be provided when operator is set.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GoogleCloudApigeeV1SecurityReportQueryMetricArgs()
        {
        }
    }
}