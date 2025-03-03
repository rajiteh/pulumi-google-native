// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    /// <summary>
    /// Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
    /// </summary>
    public sealed class BucketObjectCustomerEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption algorithm.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// SHA256 hash value of the encryption key.
        /// </summary>
        [Input("keySha256")]
        public Input<string>? KeySha256 { get; set; }

        public BucketObjectCustomerEncryptionArgs()
        {
        }
        public static new BucketObjectCustomerEncryptionArgs Empty => new BucketObjectCustomerEncryptionArgs();
    }
}
