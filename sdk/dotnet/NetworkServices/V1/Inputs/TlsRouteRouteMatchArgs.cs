// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1.Inputs
{

    /// <summary>
    /// RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "AND"ed for evaluation. If no routeMatch field is specified, this rule will unconditionally match traffic.
    /// </summary>
    public sealed class TlsRouteRouteMatchArgs : Pulumi.ResourceArgs
    {
        [Input("alpn")]
        private InputList<string>? _alpn;

        /// <summary>
        /// Optional. ALPN (Application-Layer Protocol Negotiation) to match against. Examples: "http/1.1", "h2". At least one of sni_host and alpn is required. Up to 5 alpns across all matches can be set.
        /// </summary>
        public InputList<string> Alpn
        {
            get => _alpn ?? (_alpn = new InputList<string>());
            set => _alpn = value;
        }

        [Input("sniHost")]
        private InputList<string>? _sniHost;

        /// <summary>
        /// Optional. SNI (server name indicator) to match against. SNI will be matched against all wildcard domains, i.e. www.example.com will be first matched against www.example.com, then *.example.com, then *.com. Partial wildcards are not supported, and values like *w.example.com are invalid. At least one of sni_host and alpn is required. Up to 5 sni hosts across all matches can be set.
        /// </summary>
        public InputList<string> SniHost
        {
            get => _sniHost ?? (_sniHost = new InputList<string>());
            set => _sniHost = value;
        }

        public TlsRouteRouteMatchArgs()
        {
        }
    }
}