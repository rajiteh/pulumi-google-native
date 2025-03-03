// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Digest information.
    /// </summary>
    public sealed class DigestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `SHA1`, `SHA512` etc.
        /// </summary>
        [Input("algo")]
        public Input<string>? Algo { get; set; }

        /// <summary>
        /// Value of the digest.
        /// </summary>
        [Input("digestBytes")]
        public Input<string>? DigestBytes { get; set; }

        public DigestArgs()
        {
        }
        public static new DigestArgs Empty => new DigestArgs();
    }
}
