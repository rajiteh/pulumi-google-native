// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// SAML IDP (identity provider) configuration.
    /// </summary>
    [OutputType]
    public sealed class SamlIdpConfigResponse
    {
        /// <summary>
        /// The **Change Password URL** of the identity provider. Users will be sent to this URL when changing their passwords at `myaccount.google.com`. This takes precedence over the change password URL configured at customer-level. Must use `HTTPS`.
        /// </summary>
        public readonly string ChangePasswordUri;
        /// <summary>
        /// The SAML **Entity ID** of the identity provider.
        /// </summary>
        public readonly string EntityId;
        /// <summary>
        /// The **Logout Redirect URL** (sign-out page URL) of the identity provider. When a user clicks the sign-out link on a Google page, they will be redirected to this URL. This is a pure redirect with no attached SAML `LogoutRequest` i.e. SAML single logout is not supported. Must use `HTTPS`.
        /// </summary>
        public readonly string LogoutRedirectUri;
        /// <summary>
        /// The `SingleSignOnService` endpoint location (sign-in page URL) of the identity provider. This is the URL where the `AuthnRequest` will be sent. Must use `HTTPS`. Assumed to accept the `HTTP-Redirect` binding.
        /// </summary>
        public readonly string SingleSignOnServiceUri;

        [OutputConstructor]
        private SamlIdpConfigResponse(
            string changePasswordUri,

            string entityId,

            string logoutRedirectUri,

            string singleSignOnServiceUri)
        {
            ChangePasswordUri = changePasswordUri;
            EntityId = entityId;
            LogoutRedirectUri = logoutRedirectUri;
            SingleSignOnServiceUri = singleSignOnServiceUri;
        }
    }
}
