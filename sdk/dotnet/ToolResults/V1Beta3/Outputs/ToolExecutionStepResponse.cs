// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Outputs
{

    /// <summary>
    /// Generic tool step to be used for binaries we do not explicitly support. For example: running cp to copy artifacts from one location to another.
    /// </summary>
    [OutputType]
    public sealed class ToolExecutionStepResponse
    {
        /// <summary>
        /// A Tool execution. - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        public readonly Outputs.ToolExecutionResponse ToolExecution;

        [OutputConstructor]
        private ToolExecutionStepResponse(Outputs.ToolExecutionResponse toolExecution)
        {
            ToolExecution = toolExecution;
        }
    }
}
