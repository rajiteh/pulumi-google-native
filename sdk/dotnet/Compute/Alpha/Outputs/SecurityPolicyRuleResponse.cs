// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Represents a rule that describes one or more match conditions along with the action to be taken when traffic matches this condition (allow or deny).
    /// </summary>
    [OutputType]
    public sealed class SecurityPolicyRuleResponse
    {
        /// <summary>
        /// The Action to perform when the rule is matched. The following are the valid actions: - allow: allow access to target. - deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for `STATUS` are 403, 404, and 502. - rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rate_limit_options to be set. - redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR. - throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rate_limit_options to be set for this. 
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The direction in which this rule applies. This field may only be specified when versioned_expr is set to FIREWALL.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules. This field may only be specified when the versioned_expr is set to FIREWALL.
        /// </summary>
        public readonly bool EnableLogging;
        /// <summary>
        /// Optional, additional actions that are performed on headers. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
        /// </summary>
        public readonly Outputs.SecurityPolicyRuleHttpHeaderActionResponse HeaderAction;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#securityPolicyRule for security policy rules
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// </summary>
        public readonly Outputs.SecurityPolicyRuleMatcherResponse Match;
        /// <summary>
        /// A match condition that incoming packets are evaluated against for CLOUD_ARMOR_NETWORK security policies. If it matches, the corresponding 'action' is enforced. The match criteria for a rule consists of built-in match fields (like 'srcIpRanges') and potentially multiple user-defined match fields ('userDefinedFields'). Field values may be extracted directly from the packet or derived from it (e.g. 'srcRegionCodes'). Some fields may not be present in every packet (e.g. 'srcPorts'). A user-defined field is only present if the base header is found in the packet and the entire field is in bounds. Each match field may specify which values can match it, listing one or more ranges, prefixes, or exact values that are considered a match for the field. A field value must be present in order to match a specified match field. If no match values are specified for a match field, then any field value is considered to match it, and it's not required to be present. For strings specifying '*' is also equivalent to match all. For a packet to match a rule, all specified match fields must match the corresponding field values derived from the packet. Example: networkMatch: srcIpRanges: - "192.0.2.0/24" - "198.51.100.0/24" userDefinedFields: - name: "ipv4_fragment_offset" values: - "1-0x1fff" The above match condition matches packets with a source IP in 192.0.2.0/24 or 198.51.100.0/24 and a user-defined field named "ipv4_fragment_offset" with a value between 1 and 0x1fff inclusive.
        /// </summary>
        public readonly Outputs.SecurityPolicyRuleNetworkMatcherResponse NetworkMatch;
        /// <summary>
        /// Preconfigured WAF configuration to be applied for the rule. If the rule does not evaluate preconfigured WAF rules, i.e., if evaluatePreconfiguredWaf() is not used, this field will have no effect.
        /// </summary>
        public readonly Outputs.SecurityPolicyRulePreconfiguredWafConfigResponse PreconfiguredWafConfig;
        /// <summary>
        /// If set to true, the specified action is not enforced.
        /// </summary>
        public readonly bool Preview;
        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for any other actions.
        /// </summary>
        public readonly Outputs.SecurityPolicyRuleRateLimitOptionsResponse RateLimitOptions;
        /// <summary>
        /// Parameters defining the redirect action. Cannot be specified for any other actions. This field is only supported in Global Security Policies of type CLOUD_ARMOR.
        /// </summary>
        public readonly Outputs.SecurityPolicyRuleRedirectOptionsResponse RedirectOptions;
        /// <summary>
        /// This must be specified for redirect actions. Cannot be specified for any other actions.
        /// </summary>
        public readonly string RedirectTarget;
        /// <summary>
        /// The minimum managed protection tier required for this rule. [Deprecated] Use requiredManagedProtectionTiers instead.
        /// </summary>
        public readonly string RuleManagedProtectionTier;
        /// <summary>
        /// Identifier for the rule. This is only unique within the given security policy. This can only be set during rule creation, if rule number is not specified it will be generated by the server.
        /// </summary>
        public readonly string RuleNumber;
        /// <summary>
        /// Calculation of the complexity of a single firewall security policy rule.
        /// </summary>
        public readonly int RuleTupleCount;
        /// <summary>
        /// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule. This field may only be specified when versioned_expr is set to FIREWALL.
        /// </summary>
        public readonly ImmutableArray<string> TargetResources;
        /// <summary>
        /// A list of service accounts indicating the sets of instances that are applied with this rule.
        /// </summary>
        public readonly ImmutableArray<string> TargetServiceAccounts;

        [OutputConstructor]
        private SecurityPolicyRuleResponse(
            string action,

            string description,

            string direction,

            bool enableLogging,

            Outputs.SecurityPolicyRuleHttpHeaderActionResponse headerAction,

            string kind,

            Outputs.SecurityPolicyRuleMatcherResponse match,

            Outputs.SecurityPolicyRuleNetworkMatcherResponse networkMatch,

            Outputs.SecurityPolicyRulePreconfiguredWafConfigResponse preconfiguredWafConfig,

            bool preview,

            int priority,

            Outputs.SecurityPolicyRuleRateLimitOptionsResponse rateLimitOptions,

            Outputs.SecurityPolicyRuleRedirectOptionsResponse redirectOptions,

            string redirectTarget,

            string ruleManagedProtectionTier,

            string ruleNumber,

            int ruleTupleCount,

            ImmutableArray<string> targetResources,

            ImmutableArray<string> targetServiceAccounts)
        {
            Action = action;
            Description = description;
            Direction = direction;
            EnableLogging = enableLogging;
            HeaderAction = headerAction;
            Kind = kind;
            Match = match;
            NetworkMatch = networkMatch;
            PreconfiguredWafConfig = preconfiguredWafConfig;
            Preview = preview;
            Priority = priority;
            RateLimitOptions = rateLimitOptions;
            RedirectOptions = redirectOptions;
            RedirectTarget = redirectTarget;
            RuleManagedProtectionTier = ruleManagedProtectionTier;
            RuleNumber = ruleNumber;
            RuleTupleCount = ruleTupleCount;
            TargetResources = targetResources;
            TargetServiceAccounts = targetServiceAccounts;
        }
    }
}
