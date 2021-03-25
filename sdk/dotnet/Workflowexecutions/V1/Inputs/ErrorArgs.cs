// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Workflowexecutions.V1.Inputs
{

    /// <summary>
    /// Error describes why the execution was abnormally terminated.
    /// </summary>
    public sealed class ErrorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable stack trace string.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Error message and data returned represented as a JSON string.
        /// </summary>
        [Input("payload")]
        public Input<string>? Payload { get; set; }

        /// <summary>
        /// Stack trace with detailed information of where error was generated.
        /// </summary>
        [Input("stackTrace")]
        public Input<Inputs.StackTraceArgs>? StackTrace { get; set; }

        public ErrorArgs()
        {
        }
    }
}