// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Outputs
{

    /// <summary>
    /// The content of an HL7v2 message in a structured format as specified by a schema.
    /// </summary>
    [OutputType]
    public sealed class SchematizedDataResponse
    {
        /// <summary>
        /// JSON output of the parser.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The error output of the parser.
        /// </summary>
        public readonly string Error;

        [OutputConstructor]
        private SchematizedDataResponse(
            string data,

            string error)
        {
            Data = data;
            Error = error;
        }
    }
}
