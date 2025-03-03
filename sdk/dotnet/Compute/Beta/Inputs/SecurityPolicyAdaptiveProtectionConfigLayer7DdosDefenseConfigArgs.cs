// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Configuration options for L7 DDoS detection. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
    /// </summary>
    public sealed class SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, enables CAAP for L7 DDoS detection. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
        /// </summary>
        [Input("ruleVisibility")]
        public Input<Pulumi.GoogleNative.Compute.Beta.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigRuleVisibility>? RuleVisibility { get; set; }

        public SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs()
        {
        }
        public static new SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs Empty => new SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigArgs();
    }
}
