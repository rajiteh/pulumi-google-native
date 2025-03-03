// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// The available logging options for the load balancer traffic served by this backend service.
    /// </summary>
    public sealed class BackendServiceLogConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Denotes whether to enable logging for the load balancer traffic served by this backend service. The default value is false.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        [Input("optionalFields")]
        private InputList<string>? _optionalFields;

        /// <summary>
        /// This field can only be specified if logging is enabled for this backend service and "logConfig.optionalMode" was set to CUSTOM. Contains a list of optional fields you want to include in the logs. For example: serverInstance, serverGkeDetails.cluster, serverGkeDetails.pod.podNamespace
        /// </summary>
        public InputList<string> OptionalFields
        {
            get => _optionalFields ?? (_optionalFields = new InputList<string>());
            set => _optionalFields = value;
        }

        /// <summary>
        /// This field can only be specified if logging is enabled for this backend service. Configures whether all, none or a subset of optional fields should be added to the reported logs. One of [INCLUDE_ALL_OPTIONAL, EXCLUDE_ALL_OPTIONAL, CUSTOM]. Default is EXCLUDE_ALL_OPTIONAL.
        /// </summary>
        [Input("optionalMode")]
        public Input<Pulumi.GoogleNative.Compute.V1.BackendServiceLogConfigOptionalMode>? OptionalMode { get; set; }

        /// <summary>
        /// This field can only be specified if logging is enabled for this backend service. The value of the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0.
        /// </summary>
        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        public BackendServiceLogConfigArgs()
        {
        }
        public static new BackendServiceLogConfigArgs Empty => new BackendServiceLogConfigArgs();
    }
}
