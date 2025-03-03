// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Labels that can be used to filter Apigee metrics.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1CanaryEvaluationMetricLabelsResponse
    {
        /// <summary>
        /// The environment ID associated with the metrics.
        /// </summary>
        public readonly string Env;
        /// <summary>
        /// The instance ID associated with the metrics. In Apigee Hybrid, the value is configured during installation.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The location associated with the metrics.
        /// </summary>
        public readonly string Location;

        [OutputConstructor]
        private GoogleCloudApigeeV1CanaryEvaluationMetricLabelsResponse(
            string env,

            string instanceId,

            string location)
        {
            Env = env;
            InstanceId = instanceId;
            Location = location;
        }
    }
}
