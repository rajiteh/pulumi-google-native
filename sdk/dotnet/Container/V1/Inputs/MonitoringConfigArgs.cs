// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// MonitoringConfig is cluster monitoring configuration.
    /// </summary>
    public sealed class MonitoringConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Monitoring components configuration
        /// </summary>
        [Input("componentConfig")]
        public Input<Inputs.MonitoringComponentConfigArgs>? ComponentConfig { get; set; }

        /// <summary>
        /// Enable Google Cloud Managed Service for Prometheus in the cluster.
        /// </summary>
        [Input("managedPrometheusConfig")]
        public Input<Inputs.ManagedPrometheusConfigArgs>? ManagedPrometheusConfig { get; set; }

        public MonitoringConfigArgs()
        {
        }
        public static new MonitoringConfigArgs Empty => new MonitoringConfigArgs();
    }
}
