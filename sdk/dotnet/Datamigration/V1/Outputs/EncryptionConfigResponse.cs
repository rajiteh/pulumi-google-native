// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Outputs
{

    /// <summary>
    /// EncryptionConfig describes the encryption config of a cluster that is encrypted with a CMEK (customer-managed encryption key).
    /// </summary>
    [OutputType]
    public sealed class EncryptionConfigResponse
    {
        /// <summary>
        /// The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private EncryptionConfigResponse(string kmsKeyName)
        {
            KmsKeyName = kmsKeyName;
        }
    }
}
