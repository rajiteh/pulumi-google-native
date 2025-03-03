// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1.Outputs
{

    /// <summary>
    /// Represents an OpenId Connect 1.0 identity provider.
    /// </summary>
    [OutputType]
    public sealed class GoogleIamAdminV1WorkforcePoolProviderOidcResponse
    {
        /// <summary>
        /// The client ID. Must match the audience claim of the JWT issued by the identity provider.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
        /// </summary>
        public readonly string IssuerUri;
        /// <summary>
        /// Configuration for web single sign-on for the OIDC provider. Here, web sign-in refers to console sign-in and gcloud sign-in through the browser.
        /// </summary>
        public readonly Outputs.GoogleIamAdminV1WorkforcePoolProviderOidcWebSsoConfigResponse WebSsoConfig;

        [OutputConstructor]
        private GoogleIamAdminV1WorkforcePoolProviderOidcResponse(
            string clientId,

            string issuerUri,

            Outputs.GoogleIamAdminV1WorkforcePoolProviderOidcWebSsoConfigResponse webSsoConfig)
        {
            ClientId = clientId;
            IssuerUri = issuerUri;
            WebSsoConfig = webSsoConfig;
        }
    }
}
