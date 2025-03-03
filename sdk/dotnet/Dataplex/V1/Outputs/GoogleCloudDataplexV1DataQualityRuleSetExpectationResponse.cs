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
    /// Evaluates whether each column value is contained by a specified set.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataQualityRuleSetExpectationResponse
    {
        /// <summary>
        /// Expected values for the column value.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataQualityRuleSetExpectationResponse(ImmutableArray<string> values)
        {
            Values = values;
        }
    }
}
