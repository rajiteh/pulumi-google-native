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
    /// DataQualityDimensionResult provides a more detailed, per-dimension view of the results.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataQualityDimensionResultResponse
    {
        /// <summary>
        /// Whether the dimension passed or failed.
        /// </summary>
        public readonly bool Passed;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataQualityDimensionResultResponse(bool passed)
        {
            Passed = passed;
        }
    }
}
