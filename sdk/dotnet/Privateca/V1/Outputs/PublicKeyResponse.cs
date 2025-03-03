// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    /// <summary>
    /// A PublicKey describes a public key.
    /// </summary>
    [OutputType]
    public sealed class PublicKeyResponse
    {
        /// <summary>
        /// The format of the public key.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field.
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private PublicKeyResponse(
            string format,

            string key)
        {
            Format = format;
            Key = key;
        }
    }
}
