// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1CertInfoResponse
    {
        /// <summary>
        /// X.509 basic constraints extension.
        /// </summary>
        public readonly string BasicConstraints;
        /// <summary>
        /// X.509 `notAfter` validity period in milliseconds since epoch.
        /// </summary>
        public readonly string ExpiryDate;
        /// <summary>
        /// Flag that specifies whether the certificate is valid. Flag is set to `Yes` if the certificate is valid, `No` if expired, or `Not yet` if not yet valid.
        /// </summary>
        public readonly string IsValid;
        /// <summary>
        /// X.509 issuer.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Public key component of the X.509 subject public key info.
        /// </summary>
        public readonly string PublicKey;
        /// <summary>
        /// X.509 serial number.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// X.509 signatureAlgorithm.
        /// </summary>
        public readonly string SigAlgName;
        /// <summary>
        /// X.509 subject.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// X.509 subject alternative names (SANs) extension.
        /// </summary>
        public readonly ImmutableArray<string> SubjectAlternativeNames;
        /// <summary>
        /// X.509 `notBefore` validity period in milliseconds since epoch.
        /// </summary>
        public readonly string ValidFrom;
        /// <summary>
        /// X.509 version.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GoogleCloudApigeeV1CertInfoResponse(
            string basicConstraints,

            string expiryDate,

            string isValid,

            string issuer,

            string publicKey,

            string serialNumber,

            string sigAlgName,

            string subject,

            ImmutableArray<string> subjectAlternativeNames,

            string validFrom,

            int version)
        {
            BasicConstraints = basicConstraints;
            ExpiryDate = expiryDate;
            IsValid = isValid;
            Issuer = issuer;
            PublicKey = publicKey;
            SerialNumber = serialNumber;
            SigAlgName = sigAlgName;
            Subject = subject;
            SubjectAlternativeNames = subjectAlternativeNames;
            ValidFrom = validFrom;
            Version = version;
        }
    }
}