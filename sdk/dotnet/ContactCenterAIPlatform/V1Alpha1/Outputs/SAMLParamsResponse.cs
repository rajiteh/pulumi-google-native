// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContactCenterAIPlatform.V1Alpha1.Outputs
{

    /// <summary>
    /// Message storing SAML params to enable Google as IDP.
    /// </summary>
    [OutputType]
    public sealed class SAMLParamsResponse
    {
        /// <summary>
        /// SAML certificate
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Entity id URL
        /// </summary>
        public readonly string EntityId;
        /// <summary>
        /// Single sign-on URL
        /// </summary>
        public readonly string SsoUri;
        /// <summary>
        /// Email address of the first admin users.
        /// </summary>
        public readonly string UserEmail;

        [OutputConstructor]
        private SAMLParamsResponse(
            string certificate,

            string entityId,

            string ssoUri,

            string userEmail)
        {
            Certificate = certificate;
            EntityId = entityId;
            SsoUri = ssoUri;
            UserEmail = userEmail;
        }
    }
}
