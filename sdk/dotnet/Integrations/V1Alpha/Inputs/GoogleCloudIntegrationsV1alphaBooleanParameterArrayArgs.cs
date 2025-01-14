// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// This message only contains a field of boolean array.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaBooleanParameterArrayArgs : global::Pulumi.ResourceArgs
    {
        [Input("booleanValues")]
        private InputList<bool>? _booleanValues;

        /// <summary>
        /// Boolean array.
        /// </summary>
        public InputList<bool> BooleanValues
        {
            get => _booleanValues ?? (_booleanValues = new InputList<bool>());
            set => _booleanValues = value;
        }

        public GoogleCloudIntegrationsV1alphaBooleanParameterArrayArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaBooleanParameterArrayArgs Empty => new GoogleCloudIntegrationsV1alphaBooleanParameterArrayArgs();
    }
}
