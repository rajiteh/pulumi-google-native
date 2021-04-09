// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// A signature object consists of the KeyID used and the signature itself.
    /// </summary>
    public sealed class GrafeasV1beta1IntotoSignatureArgs : Pulumi.ResourceArgs
    {
        [Input("keyid")]
        public Input<string>? Keyid { get; set; }

        [Input("sig")]
        public Input<string>? Sig { get; set; }

        public GrafeasV1beta1IntotoSignatureArgs()
        {
        }
    }
}