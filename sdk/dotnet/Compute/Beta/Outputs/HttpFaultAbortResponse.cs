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
    /// Specification for how requests are aborted as part of fault injection.
    /// </summary>
    [OutputType]
    public sealed class HttpFaultAbortResponse
    {
        /// <summary>
        /// The HTTP status code used to abort the request. The value must be from 200 to 599 inclusive. For gRPC protocol, the gRPC status code is mapped to HTTP status code according to this mapping table. HTTP status 200 is mapped to gRPC status UNKNOWN. Injecting an OK status is currently not supported by Traffic Director.
        /// </summary>
        public readonly int HttpStatus;
        /// <summary>
        /// The percentage of traffic for connections, operations, or requests that is aborted as part of fault injection. The value must be from 0.0 to 100.0 inclusive.
        /// </summary>
        public readonly double Percentage;

        [OutputConstructor]
        private HttpFaultAbortResponse(
            int httpStatus,

            double percentage)
        {
            HttpStatus = httpStatus;
            Percentage = percentage;
        }
    }
}
