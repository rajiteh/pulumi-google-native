// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Canary represents the canary deployment strategy.
    /// </summary>
    [OutputType]
    public sealed class CanaryResponse
    {
        /// <summary>
        /// Configures the progressive based deployment for a Target.
        /// </summary>
        public readonly Outputs.CanaryDeploymentResponse CanaryDeployment;
        /// <summary>
        /// Configures the progressive based deployment for a Target, but allows customizing at the phase level where a phase represents each of the percentage deployments.
        /// </summary>
        public readonly Outputs.CustomCanaryDeploymentResponse CustomCanaryDeployment;
        /// <summary>
        /// Optional. Runtime specific configurations for the deployment strategy. The runtime configuration is used to determine how Cloud Deploy will split traffic to enable a progressive deployment.
        /// </summary>
        public readonly Outputs.RuntimeConfigResponse RuntimeConfig;

        [OutputConstructor]
        private CanaryResponse(
            Outputs.CanaryDeploymentResponse canaryDeployment,

            Outputs.CustomCanaryDeploymentResponse customCanaryDeployment,

            Outputs.RuntimeConfigResponse runtimeConfig)
        {
            CanaryDeployment = canaryDeployment;
            CustomCanaryDeployment = customCanaryDeployment;
            RuntimeConfig = runtimeConfig;
        }
    }
}