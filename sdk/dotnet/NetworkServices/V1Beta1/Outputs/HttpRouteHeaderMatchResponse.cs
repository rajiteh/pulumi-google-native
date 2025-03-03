// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Outputs
{

    /// <summary>
    /// Specifies how to select a route rule based on HTTP request headers.
    /// </summary>
    [OutputType]
    public sealed class HttpRouteHeaderMatchResponse
    {
        /// <summary>
        /// The value of the header should match exactly the content of exact_match.
        /// </summary>
        public readonly string ExactMatch;
        /// <summary>
        /// The name of the HTTP header to match against.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// If specified, the match result will be inverted before checking. Default value is set to false.
        /// </summary>
        public readonly bool InvertMatch;
        /// <summary>
        /// The value of the header must start with the contents of prefix_match.
        /// </summary>
        public readonly string PrefixMatch;
        /// <summary>
        /// A header with header_name must exist. The match takes place whether or not the header has a value.
        /// </summary>
        public readonly bool PresentMatch;
        /// <summary>
        /// If specified, the rule will match if the request header value is within the range.
        /// </summary>
        public readonly Outputs.HttpRouteHeaderMatchIntegerRangeResponse RangeMatch;
        /// <summary>
        /// The value of the header must match the regular expression specified in regex_match. For regular expression grammar, please see: https://github.com/google/re2/wiki/Syntax
        /// </summary>
        public readonly string RegexMatch;
        /// <summary>
        /// The value of the header must end with the contents of suffix_match.
        /// </summary>
        public readonly string SuffixMatch;

        [OutputConstructor]
        private HttpRouteHeaderMatchResponse(
            string exactMatch,

            string header,

            bool invertMatch,

            string prefixMatch,

            bool presentMatch,

            Outputs.HttpRouteHeaderMatchIntegerRangeResponse rangeMatch,

            string regexMatch,

            string suffixMatch)
        {
            ExactMatch = exactMatch;
            Header = header;
            InvertMatch = invertMatch;
            PrefixMatch = prefixMatch;
            PresentMatch = presentMatch;
            RangeMatch = rangeMatch;
            RegexMatch = regexMatch;
            SuffixMatch = suffixMatch;
        }
    }
}
