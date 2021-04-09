// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.DLP.V2.Outputs
{

    [OutputType]
    public sealed class GooglePrivacyDlpV2AuxiliaryTableResponse
    {
        /// <summary>
        /// Required. Quasi-identifier columns.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2QuasiIdFieldResponse> QuasiIds;
        /// <summary>
        /// Required. The relative frequency column must contain a floating-point number between 0 and 1 (inclusive). Null values are assumed to be zero.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2FieldIdResponse RelativeFrequency;
        /// <summary>
        /// Required. Auxiliary table location.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2BigQueryTableResponse Table;

        [OutputConstructor]
        private GooglePrivacyDlpV2AuxiliaryTableResponse(
            ImmutableArray<Outputs.GooglePrivacyDlpV2QuasiIdFieldResponse> quasiIds,

            Outputs.GooglePrivacyDlpV2FieldIdResponse relativeFrequency,

            Outputs.GooglePrivacyDlpV2BigQueryTableResponse table)
        {
            QuasiIds = quasiIds;
            RelativeFrequency = relativeFrequency;
            Table = table;
        }
    }
}