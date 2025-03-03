// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// DataQualityRuleResult provides a more detailed, per-rule view of the results.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataQualityRuleResultResponse
    {
        /// <summary>
        /// The number of rows a rule was evaluated against. This field is only valid for ColumnMap type rules.Evaluated count can be configured to either include all rows (default) - with null rows automatically failing rule evaluation, or exclude null rows from the evaluated_count, by setting ignore_nulls = true.
        /// </summary>
        public readonly string EvaluatedCount;
        /// <summary>
        /// The query to find rows that did not pass this rule. Only applies to ColumnMap and RowCondition rules.
        /// </summary>
        public readonly string FailingRowsQuery;
        /// <summary>
        /// The number of rows with null values in the specified column.
        /// </summary>
        public readonly string NullCount;
        /// <summary>
        /// The ratio of passed_count / evaluated_count. This field is only valid for ColumnMap type rules.
        /// </summary>
        public readonly double PassRatio;
        /// <summary>
        /// Whether the rule passed or failed.
        /// </summary>
        public readonly bool Passed;
        /// <summary>
        /// The number of rows which passed a rule evaluation. This field is only valid for ColumnMap type rules.
        /// </summary>
        public readonly string PassedCount;
        /// <summary>
        /// The rule specified in the DataQualitySpec, as is.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1DataQualityRuleResponse Rule;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataQualityRuleResultResponse(
            string evaluatedCount,

            string failingRowsQuery,

            string nullCount,

            double passRatio,

            bool passed,

            string passedCount,

            Outputs.GoogleCloudDataplexV1DataQualityRuleResponse rule)
        {
            EvaluatedCount = evaluatedCount;
            FailingRowsQuery = failingRowsQuery;
            NullCount = nullCount;
            PassRatio = passRatio;
            Passed = passed;
            PassedCount = passedCount;
            Rule = rule;
        }
    }
}
