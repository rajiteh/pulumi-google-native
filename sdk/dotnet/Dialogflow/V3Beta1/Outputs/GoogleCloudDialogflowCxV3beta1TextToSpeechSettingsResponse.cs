// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Settings related to speech generating.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1TextToSpeechSettingsResponse
    {
        /// <summary>
        /// Configuration of how speech should be synthesized, mapping from language (https://dialogflow.com/docs/reference/language) to SynthesizeSpeechConfig.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SynthesizeSpeechConfigs;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1TextToSpeechSettingsResponse(ImmutableDictionary<string, string> synthesizeSpeechConfigs)
        {
            SynthesizeSpeechConfigs = synthesizeSpeechConfigs;
        }
    }
}