// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
    /// </summary>
    public sealed class ManualScalingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of instances to assign to the service at the start. This number can later be altered by using the Modules API (https://cloud.google.com/appengine/docs/python/modules/functions) set_num_instances() function.
        /// </summary>
        [Input("instances")]
        public Input<int>? Instances { get; set; }

        public ManualScalingArgs()
        {
        }
        public static new ManualScalingArgs Empty => new ManualScalingArgs();
    }
}
