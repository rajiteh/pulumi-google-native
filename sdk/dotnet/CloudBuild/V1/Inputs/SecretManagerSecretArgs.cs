// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// Pairs a secret environment variable with a SecretVersion in Secret Manager.
    /// </summary>
    public sealed class SecretManagerSecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Environment variable name to associate with the secret. Secret environment variables must be unique across all of a build's secrets, and must be used by at least one build step.
        /// </summary>
        [Input("env")]
        public Input<string>? Env { get; set; }

        /// <summary>
        /// Resource name of the SecretVersion. In format: projects/*/secrets/*/versions/*
        /// </summary>
        [Input("versionName")]
        public Input<string>? VersionName { get; set; }

        public SecretManagerSecretArgs()
        {
        }
        public static new SecretManagerSecretArgs Empty => new SecretManagerSecretArgs();
    }
}
