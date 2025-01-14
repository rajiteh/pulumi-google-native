// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// Specifies how to handle de-identification of image pixels.
    /// </summary>
    public sealed class ImageConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalInfoTypes")]
        private InputList<string>? _additionalInfoTypes;

        /// <summary>
        /// Additional InfoTypes to redact in the images in addition to those used by `text_redaction_mode`. Can only be used when `text_redaction_mode` is set to `REDACT_SENSITIVE_TEXT`, `REDACT_SENSITIVE_TEXT_CLEAN_DESCRIPTORS` or `TEXT_REDACTION_MODE_UNSPECIFIED`.
        /// </summary>
        public InputList<string> AdditionalInfoTypes
        {
            get => _additionalInfoTypes ?? (_additionalInfoTypes = new InputList<string>());
            set => _additionalInfoTypes = value;
        }

        [Input("excludeInfoTypes")]
        private InputList<string>? _excludeInfoTypes;

        /// <summary>
        /// InfoTypes to skip redacting, overriding those used by `text_redaction_mode`. Can only be used when `text_redaction_mode` is set to `REDACT_SENSITIVE_TEXT` or `REDACT_SENSITIVE_TEXT_CLEAN_DESCRIPTORS`.
        /// </summary>
        public InputList<string> ExcludeInfoTypes
        {
            get => _excludeInfoTypes ?? (_excludeInfoTypes = new InputList<string>());
            set => _excludeInfoTypes = value;
        }

        /// <summary>
        /// Determines how to redact text from image.
        /// </summary>
        [Input("textRedactionMode")]
        public Input<Pulumi.GoogleNative.Healthcare.V1Beta1.ImageConfigTextRedactionMode>? TextRedactionMode { get; set; }

        public ImageConfigArgs()
        {
        }
        public static new ImageConfigArgs Empty => new ImageConfigArgs();
    }
}
