// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Healthcare.V1.Outputs
{

    [OutputType]
    public sealed class ParsedDataResponse
    {
        public readonly ImmutableArray<Outputs.SegmentResponse> Segments;

        [OutputConstructor]
        private ParsedDataResponse(ImmutableArray<Outputs.SegmentResponse> segments)
        {
            Segments = segments;
        }
    }
}