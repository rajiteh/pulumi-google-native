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
    /// Evaluates whether each column value is contained by a specified set.
    /// </summary>
    public sealed class GoogleCloudDataplexV1DataQualityRuleSetExpectationArgs : global::Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Expected values for the column value.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GoogleCloudDataplexV1DataQualityRuleSetExpectationArgs()
        {
        }
        public static new GoogleCloudDataplexV1DataQualityRuleSetExpectationArgs Empty => new GoogleCloudDataplexV1DataQualityRuleSetExpectationArgs();
    }
}
