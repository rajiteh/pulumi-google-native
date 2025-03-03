// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Inputs
{

    /// <summary>
    /// VolumeMount describes a mounting of a Volume within a container.
    /// </summary>
    public sealed class VolumeMountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path within the container at which the volume should be mounted. Must not contain ':'.
        /// </summary>
        [Input("mountPath", required: true)]
        public Input<string> MountPath { get; set; } = null!;

        /// <summary>
        /// The name of the volume. There must be a corresponding Volume with the same name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Only true is accepted for Secret Volumes. Defaults to true for Secrets Volumes.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
        /// </summary>
        [Input("subPath")]
        public Input<string>? SubPath { get; set; }

        public VolumeMountArgs()
        {
        }
        public static new VolumeMountArgs Empty => new VolumeMountArgs();
    }
}
