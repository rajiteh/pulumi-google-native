// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// A rule captures data quality intent about a data source.
    /// </summary>
    public sealed class GoogleCloudDataplexV1DataQualityRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The unnested column which this rule is evaluated against.
        /// </summary>
        [Input("column")]
        public Input<string>? Column { get; set; }

        /// <summary>
        /// The dimension a rule belongs to. Results are also aggregated at the dimension level. Supported dimensions are "COMPLETENESS", "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS", "INTEGRITY"
        /// </summary>
        [Input("dimension", required: true)]
        public Input<string> Dimension { get; set; } = null!;

        /// <summary>
        /// Optional. Rows with null values will automatically fail a rule, unless ignore_null is true. In that case, such null rows are trivially considered passing.Only applicable to ColumnMap rules.
        /// </summary>
        [Input("ignoreNull")]
        public Input<bool>? IgnoreNull { get; set; }

        /// <summary>
        /// ColumnMap rule which evaluates whether each column value is null.
        /// </summary>
        [Input("nonNullExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleNonNullExpectationArgs>? NonNullExpectation { get; set; }

        /// <summary>
        /// ColumnMap rule which evaluates whether each column value lies between a specified range.
        /// </summary>
        [Input("rangeExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleRangeExpectationArgs>? RangeExpectation { get; set; }

        /// <summary>
        /// ColumnMap rule which evaluates whether each column value matches a specified regex.
        /// </summary>
        [Input("regexExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleRegexExpectationArgs>? RegexExpectation { get; set; }

        /// <summary>
        /// Table rule which evaluates whether each row passes the specified condition.
        /// </summary>
        [Input("rowConditionExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleRowConditionExpectationArgs>? RowConditionExpectation { get; set; }

        /// <summary>
        /// ColumnMap rule which evaluates whether each column value is contained by a specified set.
        /// </summary>
        [Input("setExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleSetExpectationArgs>? SetExpectation { get; set; }

        /// <summary>
        /// ColumnAggregate rule which evaluates whether the column aggregate statistic lies between a specified range.
        /// </summary>
        [Input("statisticRangeExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleStatisticRangeExpectationArgs>? StatisticRangeExpectation { get; set; }

        /// <summary>
        /// Table rule which evaluates whether the provided expression is true.
        /// </summary>
        [Input("tableConditionExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleTableConditionExpectationArgs>? TableConditionExpectation { get; set; }

        /// <summary>
        /// Optional. The minimum ratio of passing_rows / total_rows required to pass this rule, with a range of 0.0, 1.0.0 indicates default value (i.e. 1.0).
        /// </summary>
        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        /// <summary>
        /// ColumnAggregate rule which evaluates whether the column has duplicates.
        /// </summary>
        [Input("uniquenessExpectation")]
        public Input<Inputs.GoogleCloudDataplexV1DataQualityRuleUniquenessExpectationArgs>? UniquenessExpectation { get; set; }

        public GoogleCloudDataplexV1DataQualityRuleArgs()
        {
        }
        public static new GoogleCloudDataplexV1DataQualityRuleArgs Empty => new GoogleCloudDataplexV1DataQualityRuleArgs();
    }
}
