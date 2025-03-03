// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Outputs
{

    /// <summary>
    /// An OS policy resource is used to define the desired state configuration and provides a specific functionality like installing/removing packages, executing a script etc. The system ensures that resources are always in their desired state by taking necessary actions if they have drifted from their desired state.
    /// </summary>
    [OutputType]
    public sealed class OSPolicyResourceResponse
    {
        /// <summary>
        /// Exec resource
        /// </summary>
        public readonly Outputs.OSPolicyResourceExecResourceResponse Exec;
        /// <summary>
        /// File resource
        /// </summary>
        public readonly Outputs.OSPolicyResourceFileResourceResponse File;
        /// <summary>
        /// Package resource
        /// </summary>
        public readonly Outputs.OSPolicyResourcePackageResourceResponse Pkg;
        /// <summary>
        /// Package repository resource
        /// </summary>
        public readonly Outputs.OSPolicyResourceRepositoryResourceResponse Repository;

        [OutputConstructor]
        private OSPolicyResourceResponse(
            Outputs.OSPolicyResourceExecResourceResponse exec,

            Outputs.OSPolicyResourceFileResourceResponse file,

            Outputs.OSPolicyResourcePackageResourceResponse pkg,

            Outputs.OSPolicyResourceRepositoryResourceResponse repository)
        {
            Exec = exec;
            File = file;
            Pkg = pkg;
            Repository = repository;
        }
    }
}
