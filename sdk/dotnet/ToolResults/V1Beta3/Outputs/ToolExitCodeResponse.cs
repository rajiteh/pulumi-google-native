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
    /// Exit code from a tool execution.
    /// </summary>
    [OutputType]
    public sealed class ToolExitCodeResponse
    {
        /// <summary>
        /// Tool execution exit code. A value of 0 means that the execution was successful. - In response: always set - In create/update request: always set
        /// </summary>
        public readonly int Number;

        [OutputConstructor]
        private ToolExitCodeResponse(int number)
        {
            Number = number;
        }
    }
}
