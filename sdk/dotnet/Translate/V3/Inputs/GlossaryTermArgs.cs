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
    /// Represents a single glossary term
    /// </summary>
    public sealed class GlossaryTermArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The language for this glossary term.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// The text for the glossary term.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public GlossaryTermArgs()
        {
        }
        public static new GlossaryTermArgs Empty => new GlossaryTermArgs();
    }
}
