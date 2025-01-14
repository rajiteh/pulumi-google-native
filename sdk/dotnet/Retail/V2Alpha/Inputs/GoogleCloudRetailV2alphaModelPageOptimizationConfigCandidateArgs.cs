// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Inputs
{

    /// <summary>
    /// A candidate to consider for a given panel. Currently only ServingConfig are valid candidates.
    /// </summary>
    public sealed class GoogleCloudRetailV2alphaModelPageOptimizationConfigCandidateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This has to be a valid ServingConfig identifier. For example, for a ServingConfig with full name: `projects/*/locations/global/catalogs/default_catalog/servingConfigs/my_candidate_config`, this would be `my_candidate_config`.
        /// </summary>
        [Input("servingConfigId")]
        public Input<string>? ServingConfigId { get; set; }

        public GoogleCloudRetailV2alphaModelPageOptimizationConfigCandidateArgs()
        {
        }
        public static new GoogleCloudRetailV2alphaModelPageOptimizationConfigCandidateArgs Empty => new GoogleCloudRetailV2alphaModelPageOptimizationConfigCandidateArgs();
    }
}
