// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WorkloadManager.V1.Outputs
{

    /// <summary>
    /// Message describing resource status
    /// </summary>
    [OutputType]
    public sealed class ResourceStatusResponse
    {
        /// <summary>
        /// the new version of rule id if exists
        /// </summary>
        public readonly ImmutableArray<string> RulesNewerVersions;
        /// <summary>
        /// State of the resource
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private ResourceStatusResponse(
            ImmutableArray<string> rulesNewerVersions,

            string state)
        {
            RulesNewerVersions = rulesNewerVersions;
            State = state;
        }
    }
}
