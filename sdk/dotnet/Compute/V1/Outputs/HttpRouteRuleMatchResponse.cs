// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// HttpRouteRuleMatch specifies a set of criteria for matching requests to an HttpRouteRule. All specified criteria must be satisfied for a match to occur.
    /// </summary>
    [OutputType]
    public sealed class HttpRouteRuleMatchResponse
    {
        /// <summary>
        /// For satisfying the matchRule condition, the path of the request must exactly match the value specified in fullPathMatch after removing any query parameters and anchor that may be part of the original URL. fullPathMatch must be from 1 to 1024 characters. Only one of prefixMatch, fullPathMatch or regexMatch must be specified.
        /// </summary>
        public readonly string FullPathMatch;
        /// <summary>
        /// Specifies a list of header match criteria, all of which must match corresponding headers in the request.
        /// </summary>
        public readonly ImmutableArray<Outputs.HttpHeaderMatchResponse> HeaderMatches;
        /// <summary>
        /// Specifies that prefixMatch and fullPathMatch matches are case sensitive. The default value is false. ignoreCase must not be used with regexMatch. Not supported when the URL map is bound to a target gRPC proxy.
        /// </summary>
        public readonly bool IgnoreCase;
        /// <summary>
        /// Opaque filter criteria used by the load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to the load balancer, xDS clients present node metadata. When there is a match, the relevant routing configuration is made available to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadata filters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here is applied after those specified in ForwardingRule that refers to the UrlMap this HttpRouteRuleMatch belongs to. metadataFilters only applies to load balancers that have loadBalancingScheme set to INTERNAL_SELF_MANAGED. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetadataFilterResponse> MetadataFilters;
        /// <summary>
        /// If specified, the route is a pattern match expression that must match the :path header once the query string is removed. A pattern match allows you to match - The value must be between 1 and 1024 characters - The pattern must start with a leading slash ("/") - There may be no more than 5 operators in pattern Precisely one of prefix_match, full_path_match, regex_match or path_template_match must be set.
        /// </summary>
        public readonly string PathTemplateMatch;
        /// <summary>
        /// For satisfying the matchRule condition, the request's path must begin with the specified prefixMatch. prefixMatch must begin with a /. The value must be from 1 to 1024 characters. Only one of prefixMatch, fullPathMatch or regexMatch must be specified.
        /// </summary>
        public readonly string PrefixMatch;
        /// <summary>
        /// Specifies a list of query parameter match criteria, all of which must match corresponding query parameters in the request. Not supported when the URL map is bound to a target gRPC proxy.
        /// </summary>
        public readonly ImmutableArray<Outputs.HttpQueryParameterMatchResponse> QueryParameterMatches;
        /// <summary>
        /// For satisfying the matchRule condition, the path of the request must satisfy the regular expression specified in regexMatch after removing any query parameters and anchor supplied with the original URL. For more information about regular expression syntax, see Syntax. Only one of prefixMatch, fullPathMatch or regexMatch must be specified. Regular expressions can only be used when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public readonly string RegexMatch;

        [OutputConstructor]
        private HttpRouteRuleMatchResponse(
            string fullPathMatch,

            ImmutableArray<Outputs.HttpHeaderMatchResponse> headerMatches,

            bool ignoreCase,

            ImmutableArray<Outputs.MetadataFilterResponse> metadataFilters,

            string pathTemplateMatch,

            string prefixMatch,

            ImmutableArray<Outputs.HttpQueryParameterMatchResponse> queryParameterMatches,

            string regexMatch)
        {
            FullPathMatch = fullPathMatch;
            HeaderMatches = headerMatches;
            IgnoreCase = ignoreCase;
            MetadataFilters = metadataFilters;
            PathTemplateMatch = pathTemplateMatch;
            PrefixMatch = prefixMatch;
            QueryParameterMatches = queryParameterMatches;
            RegexMatch = regexMatch;
        }
    }
}
