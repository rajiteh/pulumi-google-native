// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3.Inputs
{

    /// <summary>
    /// Used with unidirectional glossaries.
    /// </summary>
    public sealed class LanguageCodePairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ISO-639 language code of the input text, for example, "en-US". Expected to be an exact match for GlossaryTerm.language_code.
        /// </summary>
        [Input("sourceLanguageCode", required: true)]
        public Input<string> SourceLanguageCode { get; set; } = null!;

        /// <summary>
        /// The ISO-639 language code for translation output, for example, "zh-CN". Expected to be an exact match for GlossaryTerm.language_code.
        /// </summary>
        [Input("targetLanguageCode", required: true)]
        public Input<string> TargetLanguageCode { get; set; } = null!;

        public LanguageCodePairArgs()
        {
        }
        public static new LanguageCodePairArgs Empty => new LanguageCodePairArgs();
    }
}
