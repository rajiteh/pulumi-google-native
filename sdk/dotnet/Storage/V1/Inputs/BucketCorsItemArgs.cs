// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Inputs
{

    public sealed class BucketCorsItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value, in seconds, to return in the  Access-Control-Max-Age header used in preflight responses.
        /// </summary>
        [Input("maxAgeSeconds")]
        public Input<int>? MaxAgeSeconds { get; set; }

        [Input("method")]
        private InputList<string>? _method;

        /// <summary>
        /// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
        /// </summary>
        public InputList<string> Method
        {
            get => _method ?? (_method = new InputList<string>());
            set => _method = value;
        }

        [Input("origin")]
        private InputList<string>? _origin;

        /// <summary>
        /// The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
        /// </summary>
        public InputList<string> Origin
        {
            get => _origin ?? (_origin = new InputList<string>());
            set => _origin = value;
        }

        [Input("responseHeader")]
        private InputList<string>? _responseHeader;

        /// <summary>
        /// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
        /// </summary>
        public InputList<string> ResponseHeader
        {
            get => _responseHeader ?? (_responseHeader = new InputList<string>());
            set => _responseHeader = value;
        }

        public BucketCorsItemArgs()
        {
        }
        public static new BucketCorsItemArgs Empty => new BucketCorsItemArgs();
    }
}
