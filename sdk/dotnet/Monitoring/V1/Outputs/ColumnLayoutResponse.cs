// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Monitoring.V1.Outputs
{

    [OutputType]
    public sealed class ColumnLayoutResponse
    {
        /// <summary>
        /// The columns of content to display.
        /// </summary>
        public readonly ImmutableArray<Outputs.ColumnResponse> Columns;

        [OutputConstructor]
        private ColumnLayoutResponse(ImmutableArray<Outputs.ColumnResponse> columns)
        {
            Columns = columns;
        }
    }
}