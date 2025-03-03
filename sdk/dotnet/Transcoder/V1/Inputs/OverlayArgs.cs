// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Overlay configuration.
    /// </summary>
    public sealed class OverlayArgs : global::Pulumi.ResourceArgs
    {
        [Input("animations")]
        private InputList<Inputs.AnimationArgs>? _animations;

        /// <summary>
        /// List of Animations. The list should be chronological, without any time overlap.
        /// </summary>
        public InputList<Inputs.AnimationArgs> Animations
        {
            get => _animations ?? (_animations = new InputList<Inputs.AnimationArgs>());
            set => _animations = value;
        }

        /// <summary>
        /// Image overlay.
        /// </summary>
        [Input("image")]
        public Input<Inputs.ImageArgs>? Image { get; set; }

        public OverlayArgs()
        {
        }
        public static new OverlayArgs Empty => new OverlayArgs();
    }
}
