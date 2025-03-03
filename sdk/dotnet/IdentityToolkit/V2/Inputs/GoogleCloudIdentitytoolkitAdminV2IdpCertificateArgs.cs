// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2.Inputs
{

    /// <summary>
    /// The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
    /// </summary>
    public sealed class GoogleCloudIdentitytoolkitAdminV2IdpCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The x509 certificate
        /// </summary>
        [Input("x509Certificate")]
        public Input<string>? X509Certificate { get; set; }

        public GoogleCloudIdentitytoolkitAdminV2IdpCertificateArgs()
        {
        }
        public static new GoogleCloudIdentitytoolkitAdminV2IdpCertificateArgs Empty => new GoogleCloudIdentitytoolkitAdminV2IdpCertificateArgs();
    }
}
