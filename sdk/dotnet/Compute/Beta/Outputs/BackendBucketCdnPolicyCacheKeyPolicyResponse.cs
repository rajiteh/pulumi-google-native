// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Message containing what to include in the cache key for a request for Cloud CDN.
    /// </summary>
    [OutputType]
    public sealed class BackendBucketCdnPolicyCacheKeyPolicyResponse
    {
        /// <summary>
        /// Allows HTTP request headers (by name) to be used in the cache key.
        /// </summary>
        public readonly ImmutableArray<string> IncludeHttpHeaders;
        /// <summary>
        /// Names of query string parameters to include in cache keys. Default parameters are always included. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public readonly ImmutableArray<string> QueryStringWhitelist;

        [OutputConstructor]
        private BackendBucketCdnPolicyCacheKeyPolicyResponse(
            ImmutableArray<string> includeHttpHeaders,

            ImmutableArray<string> queryStringWhitelist)
        {
            IncludeHttpHeaders = includeHttpHeaders;
            QueryStringWhitelist = queryStringWhitelist;
        }
    }
}
