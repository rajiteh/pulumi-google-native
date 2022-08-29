// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Outputs
{

    /// <summary>
    /// A single target dataset to which all data will be streamed.
    /// </summary>
    [OutputType]
    public sealed class SingleTargetDatasetResponse
    {
        public readonly string DatasetId;

        [OutputConstructor]
        private SingleTargetDatasetResponse(string datasetId)
        {
            DatasetId = datasetId;
        }
    }
}