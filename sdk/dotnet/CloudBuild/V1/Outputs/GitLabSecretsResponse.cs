// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// GitLabSecrets represents the secrets in Secret Manager for a GitLab integration.
    /// </summary>
    [OutputType]
    public sealed class GitLabSecretsResponse
    {
        /// <summary>
        /// The resource name for the api access token’s secret version
        /// </summary>
        public readonly string ApiAccessTokenVersion;
        /// <summary>
        /// Immutable. API Key that will be attached to webhook requests from GitLab to Cloud Build.
        /// </summary>
        public readonly string ApiKeyVersion;
        /// <summary>
        /// The resource name for the read access token’s secret version
        /// </summary>
        public readonly string ReadAccessTokenVersion;
        /// <summary>
        /// Immutable. The resource name for the webhook secret’s secret version. Once this field has been set, it cannot be changed. If you need to change it, please create another GitLabConfig.
        /// </summary>
        public readonly string WebhookSecretVersion;

        [OutputConstructor]
        private GitLabSecretsResponse(
            string apiAccessTokenVersion,

            string apiKeyVersion,

            string readAccessTokenVersion,

            string webhookSecretVersion)
        {
            ApiAccessTokenVersion = apiAccessTokenVersion;
            ApiKeyVersion = apiKeyVersion;
            ReadAccessTokenVersion = readAccessTokenVersion;
            WebhookSecretVersion = webhookSecretVersion;
        }
    }
}