// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// Represents a hardware accelerator request config. Note that the AcceleratorConfig can be used in both Jobs and Versions. Learn more about [accelerators for training](/ml-engine/docs/using-gpus) and [accelerators for online prediction](/ml-engine/docs/machine-types-online-prediction#gpus).
    /// </summary>
    public sealed class GoogleCloudMlV1__AcceleratorConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of accelerators to attach to each machine running the job.
        /// </summary>
        [Input("count")]
        public Input<string>? Count { get; set; }

        /// <summary>
        /// The type of accelerator to use.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Ml.V1.GoogleCloudMlV1__AcceleratorConfigType>? Type { get; set; }

        public GoogleCloudMlV1__AcceleratorConfigArgs()
        {
        }
        public static new GoogleCloudMlV1__AcceleratorConfigArgs Empty => new GoogleCloudMlV1__AcceleratorConfigArgs();
    }
}
