// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// The specification for personalization.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaSearchRequestPersonalizationSpecResponse
    {
        /// <summary>
        /// Defaults to Mode.AUTO.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaSearchRequestPersonalizationSpecResponse(string mode)
        {
            Mode = mode;
        }
    }
}
