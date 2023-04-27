// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1Alpha1.Outputs
{

    /// <summary>
    /// Describes a collection of data points rendered as a Chart.
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryChartDataResponse
    {
        /// <summary>
        /// Each data point in the chart is represented as a name-value pair with the name being the x-axis label, and the value being the y-axis value.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportSummaryChartDataDataPointResponse> DataPoints;

        [OutputConstructor]
        private ReportSummaryChartDataResponse(ImmutableArray<Outputs.ReportSummaryChartDataDataPointResponse> dataPoints)
        {
            DataPoints = dataPoints;
        }
    }
}