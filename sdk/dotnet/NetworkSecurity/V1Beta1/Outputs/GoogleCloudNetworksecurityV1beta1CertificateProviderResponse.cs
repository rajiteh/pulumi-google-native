// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1.Outputs
{

    /// <summary>
    /// Specification of certificate provider. Defines the mechanism to obtain the certificate and private key for peer to peer authentication.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudNetworksecurityV1beta1CertificateProviderResponse
    {
        /// <summary>
        /// The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
        /// </summary>
        public readonly Outputs.CertificateProviderInstanceResponse CertificateProviderInstance;
        /// <summary>
        /// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
        /// </summary>
        public readonly Outputs.GoogleCloudNetworksecurityV1beta1GrpcEndpointResponse GrpcEndpoint;

        [OutputConstructor]
        private GoogleCloudNetworksecurityV1beta1CertificateProviderResponse(
            Outputs.CertificateProviderInstanceResponse certificateProviderInstance,

            Outputs.GoogleCloudNetworksecurityV1beta1GrpcEndpointResponse grpcEndpoint)
        {
            CertificateProviderInstance = certificateProviderInstance;
            GrpcEndpoint = grpcEndpoint;
        }
    }
}
