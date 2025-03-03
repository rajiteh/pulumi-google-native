// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Custom information type provided by the user. Used to find domain-specific sensitive information configurable to the data in question.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2CustomInfoTypeResponse
    {
        /// <summary>
        /// Set of detection rules to apply to all findings of this CustomInfoType. Rules are applied in order that they are specified. Not supported for the `surrogate_type` CustomInfoType.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2DetectionRuleResponse> DetectionRules;
        /// <summary>
        /// A list of phrases to detect as a CustomInfoType.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2DictionaryResponse Dictionary;
        /// <summary>
        /// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching.
        /// </summary>
        public readonly string ExclusionType;
        /// <summary>
        /// CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing infoTypes and that infoType is specified in `InspectContent.info_types` field. Specifying the latter adds findings to the one detected by the system. If built-in info type is not specified in `InspectContent.info_types` list then the name is treated as a custom info type.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2InfoTypeResponse InfoType;
        /// <summary>
        /// Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria specified by the rule. Defaults to `VERY_LIKELY` if not specified.
        /// </summary>
        public readonly string Likelihood;
        /// <summary>
        /// Regular expression based CustomInfoType.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2RegexResponse Regex;
        /// <summary>
        /// Sensitivity for this CustomInfoType. If this CustomInfoType extends an existing InfoType, the sensitivity here will take precedent over that of the original InfoType. If unset for a CustomInfoType, it will default to HIGH. This only applies to data profiling.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2SensitivityScoreResponse SensitivityScore;
        /// <summary>
        /// Load an existing `StoredInfoType` resource for use in `InspectDataSource`. Not currently supported in `InspectContent`.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2StoredTypeResponse StoredType;
        /// <summary>
        /// Message for detecting output from deidentification transformations that support reversing.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2SurrogateTypeResponse SurrogateType;

        [OutputConstructor]
        private GooglePrivacyDlpV2CustomInfoTypeResponse(
            ImmutableArray<Outputs.GooglePrivacyDlpV2DetectionRuleResponse> detectionRules,

            Outputs.GooglePrivacyDlpV2DictionaryResponse dictionary,

            string exclusionType,

            Outputs.GooglePrivacyDlpV2InfoTypeResponse infoType,

            string likelihood,

            Outputs.GooglePrivacyDlpV2RegexResponse regex,

            Outputs.GooglePrivacyDlpV2SensitivityScoreResponse sensitivityScore,

            Outputs.GooglePrivacyDlpV2StoredTypeResponse storedType,

            Outputs.GooglePrivacyDlpV2SurrogateTypeResponse surrogateType)
        {
            DetectionRules = detectionRules;
            Dictionary = dictionary;
            ExclusionType = exclusionType;
            InfoType = infoType;
            Likelihood = likelihood;
            Regex = regex;
            SensitivityScore = sensitivityScore;
            StoredType = storedType;
            SurrogateType = surrogateType;
        }
    }
}
