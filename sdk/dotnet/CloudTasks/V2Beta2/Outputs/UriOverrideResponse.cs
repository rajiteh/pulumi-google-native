// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2.Outputs
{

    /// <summary>
    /// Uri Override. When specified, all the HTTP tasks inside the queue will be partially or fully overridden depending on the configured values.
    /// </summary>
    [OutputType]
    public sealed class UriOverrideResponse
    {
        /// <summary>
        /// Host override. When specified, replaces the host part of the task URL. For example, if the task URL is "https://www.google.com," and host value is set to "example.net", the overridden URI will be changed to "https://example.net." Host value cannot be an empty string (INVALID_ARGUMENT).
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// URI path. When specified, replaces the existing path of the task URL. Setting the path value to an empty string clears the URI path segment.
        /// </summary>
        public readonly Outputs.PathOverrideResponse PathOverride;
        /// <summary>
        /// Port override. When specified, replaces the port part of the task URI. For instance, for a URI http://www.google.com/foo and port=123, the overridden URI becomes http://www.google.com:123/foo. Note that the port value must be a positive integer. Setting the port to 0 (Zero) clears the URI port.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// URI Query. When specified, replaces the query part of the task URI. Setting the query value to an empty string clears the URI query segment.
        /// </summary>
        public readonly Outputs.QueryOverrideResponse QueryOverride;
        /// <summary>
        /// Scheme override. When specified, the task URI scheme is replaced by the provided value (HTTP or HTTPS).
        /// </summary>
        public readonly string Scheme;
        /// <summary>
        /// URI Override Enforce Mode When specified, determines the Target UriOverride mode. If not specified, it defaults to ALWAYS.
        /// </summary>
        public readonly string UriOverrideEnforceMode;

        [OutputConstructor]
        private UriOverrideResponse(
            string host,

            Outputs.PathOverrideResponse pathOverride,

            string port,

            Outputs.QueryOverrideResponse queryOverride,

            string scheme,

            string uriOverrideEnforceMode)
        {
            Host = host;
            PathOverride = pathOverride;
            Port = port;
            QueryOverride = queryOverride;
            Scheme = scheme;
            UriOverrideEnforceMode = uriOverrideEnforceMode;
        }
    }
}
