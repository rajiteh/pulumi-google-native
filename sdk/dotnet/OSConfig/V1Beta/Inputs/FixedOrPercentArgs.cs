// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Message encapsulating a value that can be either absolute ("fixed") or relative ("percent") to a value.
    /// </summary>
    public sealed class FixedOrPercentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a fixed value.
        /// </summary>
        [Input("fixed")]
        public Input<int>? Fixed { get; set; }

        /// <summary>
        /// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
        /// </summary>
        [Input("percent")]
        public Input<int>? Percent { get; set; }

        public FixedOrPercentArgs()
        {
        }
        public static new FixedOrPercentArgs Empty => new FixedOrPercentArgs();
    }
}
