// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ArtifactRegistry.V1.Outputs
{

    /// <summary>
    /// Artifact policy configuration for the repository contents.
    /// </summary>
    [OutputType]
    public sealed class UpstreamPolicyResponse
    {
        /// <summary>
        /// Entries with a greater priority value take precedence in the pull order.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// A reference to the repository resource, for example: "projects/p1/locations/us-central1/repositories/repo1".
        /// </summary>
        public readonly string Repository;

        [OutputConstructor]
        private UpstreamPolicyResponse(
            int priority,

            string repository)
        {
            Priority = priority;
            Repository = repository;
        }
    }
}
