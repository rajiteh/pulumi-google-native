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
    /// End previous overlay animation from the video. Without AnimationEnd, the overlay object will keep the state of previous animation until the end of the video.
    /// </summary>
    public sealed class AnimationEndArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time to end overlay object, in seconds. Default: 0
        /// </summary>
        [Input("startTimeOffset")]
        public Input<string>? StartTimeOffset { get; set; }

        public AnimationEndArgs()
        {
        }
        public static new AnimationEndArgs Empty => new AnimationEndArgs();
    }
}
