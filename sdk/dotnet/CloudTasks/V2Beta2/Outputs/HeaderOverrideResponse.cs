// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2.Outputs
{

    /// <summary>
    /// Wraps the Header object.
    /// </summary>
    [OutputType]
    public sealed class HeaderOverrideResponse
    {
        /// <summary>
        /// header embodying a key and a value.
        /// </summary>
        public readonly Outputs.HeaderResponse Header;

        [OutputConstructor]
        private HeaderOverrideResponse(Outputs.HeaderResponse header)
        {
            Header = header;
        }
    }
}