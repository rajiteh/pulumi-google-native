// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// Disk encryption configuration for an instance.
    /// </summary>
    [OutputType]
    public sealed class DiskEncryptionConfigurationResponse
    {
        /// <summary>
        /// This is always `sql#diskEncryptionConfiguration`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource name of KMS key for disk encryption
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private DiskEncryptionConfigurationResponse(
            string kind,

            string kmsKeyName)
        {
            Kind = kind;
            KmsKeyName = kmsKeyName;
        }
    }
}
