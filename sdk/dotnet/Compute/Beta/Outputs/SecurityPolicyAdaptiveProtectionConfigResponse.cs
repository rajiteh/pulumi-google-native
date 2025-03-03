// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Configuration options for Cloud Armor Adaptive Protection (CAAP).
    /// </summary>
    [OutputType]
    public sealed class SecurityPolicyAdaptiveProtectionConfigResponse
    {
        public readonly Outputs.SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigResponse AutoDeployConfig;
        /// <summary>
        /// If set to true, enables Cloud Armor Machine Learning.
        /// </summary>
        public readonly Outputs.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigResponse Layer7DdosDefenseConfig;

        [OutputConstructor]
        private SecurityPolicyAdaptiveProtectionConfigResponse(
            Outputs.SecurityPolicyAdaptiveProtectionConfigAutoDeployConfigResponse autoDeployConfig,

            Outputs.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigResponse layer7DdosDefenseConfig)
        {
            AutoDeployConfig = autoDeployConfig;
            Layer7DdosDefenseConfig = layer7DdosDefenseConfig;
        }
    }
}
