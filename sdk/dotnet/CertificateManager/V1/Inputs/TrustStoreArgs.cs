// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1.Inputs
{

    /// <summary>
    /// Defines a trust store.
    /// </summary>
    public sealed class TrustStoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("intermediateCas")]
        private InputList<Inputs.IntermediateCAArgs>? _intermediateCas;

        /// <summary>
        /// Set of intermediate CA certificates used for the path building phase of chain validation. The field is currently not supported if TrustConfig is used for the workload certificate feature.
        /// </summary>
        public InputList<Inputs.IntermediateCAArgs> IntermediateCas
        {
            get => _intermediateCas ?? (_intermediateCas = new InputList<Inputs.IntermediateCAArgs>());
            set => _intermediateCas = value;
        }

        [Input("trustAnchors")]
        private InputList<Inputs.TrustAnchorArgs>? _trustAnchors;

        /// <summary>
        /// List of Trust Anchors to be used while performing validation against a given TrustStore.
        /// </summary>
        public InputList<Inputs.TrustAnchorArgs> TrustAnchors
        {
            get => _trustAnchors ?? (_trustAnchors = new InputList<Inputs.TrustAnchorArgs>());
            set => _trustAnchors = value;
        }

        public TrustStoreArgs()
        {
        }
        public static new TrustStoreArgs Empty => new TrustStoreArgs();
    }
}