// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// A MetricRange is used when each window is good when the value x of a single TimeSeries satisfies range.min &lt;= x &lt;= range.max. The provided TimeSeries must have ValueType = INT64 or ValueType = DOUBLE and MetricKind = GAUGE.
    /// </summary>
    public sealed class MetricRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Range of values considered "good." For a one-sided range, set one bound to an infinite value.
        /// </summary>
        [Input("range")]
        public Input<Inputs.GoogleMonitoringV3RangeArgs>? Range { get; set; }

        /// <summary>
        /// A monitoring filter (https://cloud.google.com/monitoring/api/v3/filters) specifying the TimeSeries to use for evaluating window quality.
        /// </summary>
        [Input("timeSeries")]
        public Input<string>? TimeSeries { get; set; }

        public MetricRangeArgs()
        {
        }
        public static new MetricRangeArgs Empty => new MetricRangeArgs();
    }
}
