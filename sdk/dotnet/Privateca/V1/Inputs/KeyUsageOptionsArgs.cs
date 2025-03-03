// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// KeyUsage.KeyUsageOptions corresponds to the key usage values described in https://tools.ietf.org/html/rfc5280#section-4.2.1.3.
    /// </summary>
    public sealed class KeyUsageOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key may be used to sign certificates.
        /// </summary>
        [Input("certSign")]
        public Input<bool>? CertSign { get; set; }

        /// <summary>
        /// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
        /// </summary>
        [Input("contentCommitment")]
        public Input<bool>? ContentCommitment { get; set; }

        /// <summary>
        /// The key may be used sign certificate revocation lists.
        /// </summary>
        [Input("crlSign")]
        public Input<bool>? CrlSign { get; set; }

        /// <summary>
        /// The key may be used to encipher data.
        /// </summary>
        [Input("dataEncipherment")]
        public Input<bool>? DataEncipherment { get; set; }

        /// <summary>
        /// The key may be used to decipher only.
        /// </summary>
        [Input("decipherOnly")]
        public Input<bool>? DecipherOnly { get; set; }

        /// <summary>
        /// The key may be used for digital signatures.
        /// </summary>
        [Input("digitalSignature")]
        public Input<bool>? DigitalSignature { get; set; }

        /// <summary>
        /// The key may be used to encipher only.
        /// </summary>
        [Input("encipherOnly")]
        public Input<bool>? EncipherOnly { get; set; }

        /// <summary>
        /// The key may be used in a key agreement protocol.
        /// </summary>
        [Input("keyAgreement")]
        public Input<bool>? KeyAgreement { get; set; }

        /// <summary>
        /// The key may be used to encipher other keys.
        /// </summary>
        [Input("keyEncipherment")]
        public Input<bool>? KeyEncipherment { get; set; }

        public KeyUsageOptionsArgs()
        {
        }
        public static new KeyUsageOptionsArgs Empty => new KeyUsageOptionsArgs();
    }
}
