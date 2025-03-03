// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A set of Confidential Instance options.
    /// </summary>
    public sealed class ConfidentialInstanceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the instance should have confidential compute enabled.
        /// </summary>
        [Input("enableConfidentialCompute")]
        public Input<bool>? EnableConfidentialCompute { get; set; }

        public ConfidentialInstanceConfigArgs()
        {
        }
        public static new ConfidentialInstanceConfigArgs Empty => new ConfidentialInstanceConfigArgs();
    }
}
