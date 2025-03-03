// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Outputs
{

    /// <summary>
    /// A Docker container.
    /// </summary>
    [OutputType]
    public sealed class ContainerResponse
    {
        /// <summary>
        /// Arguments passed to the entrypoint.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// If set, overrides the default ENTRYPOINT specified by the image.
        /// </summary>
        public readonly ImmutableArray<string> Command;
        /// <summary>
        /// Environment variables passed to the container's entrypoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Env;
        /// <summary>
        /// Docker image defining the container. This image must be accessible by the service account specified in the workstation configuration.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// If set, overrides the USER specified in the image with the given uid.
        /// </summary>
        public readonly int RunAsUser;
        /// <summary>
        /// If set, overrides the default DIR specified by the image.
        /// </summary>
        public readonly string WorkingDir;

        [OutputConstructor]
        private ContainerResponse(
            ImmutableArray<string> args,

            ImmutableArray<string> command,

            ImmutableDictionary<string, string> env,

            string image,

            int runAsUser,

            string workingDir)
        {
            Args = args;
            Command = command;
            Env = env;
            Image = image;
            RunAsUser = runAsUser;
            WorkingDir = workingDir;
        }
    }
}
