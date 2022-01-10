// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// GcfsConfig contains configurations of Google Container File System.
    /// </summary>
    [OutputType]
    public sealed class GcfsConfigResponse
    {
        /// <summary>
        /// Whether to use GCFS.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private GcfsConfigResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}