// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Configuration of Fast Socket feature.
    /// </summary>
    [OutputType]
    public sealed class FastSocketResponse
    {
        /// <summary>
        /// Whether Fast Socket features are enabled in the node pool.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private FastSocketResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
