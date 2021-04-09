// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.DLP.V2.Inputs
{

    /// <summary>
    /// Message defining a list of words or phrases to search for in the data.
    /// </summary>
    public sealed class GooglePrivacyDlpV2WordListArgs : Pulumi.ResourceArgs
    {
        [Input("words")]
        private InputList<string>? _words;

        /// <summary>
        /// Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]
        /// </summary>
        public InputList<string> Words
        {
            get => _words ?? (_words = new InputList<string>());
            set => _words = value;
        }

        public GooglePrivacyDlpV2WordListArgs()
        {
        }
    }
}