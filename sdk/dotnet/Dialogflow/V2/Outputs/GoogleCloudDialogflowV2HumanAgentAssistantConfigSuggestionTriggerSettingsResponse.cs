// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Dialogflow.V2.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionTriggerSettingsResponse
    {
        /// <summary>
        /// Do not trigger if last utterance is small talk.
        /// </summary>
        public readonly bool NoSmalltalk;
        /// <summary>
        /// Only trigger suggestion if participant role of last utterance is END_USER.
        /// </summary>
        public readonly bool OnlyEndUser;

        [OutputConstructor]
        private GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionTriggerSettingsResponse(
            bool noSmalltalk,

            bool onlyEndUser)
        {
            NoSmalltalk = noSmalltalk;
            OnlyEndUser = onlyEndUser;
        }
    }
}