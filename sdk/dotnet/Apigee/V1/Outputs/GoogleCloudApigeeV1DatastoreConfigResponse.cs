// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Apigee.V1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudApigeeV1DatastoreConfigResponse
    {
        /// <summary>
        /// Name of the Cloud Storage bucket. Required for `gcs` target_type.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// BigQuery dataset name Required for `bigquery` target_type.
        /// </summary>
        public readonly string DatasetName;
        /// <summary>
        /// Path of Cloud Storage bucket Required for `gcs` target_type.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Required. GCP project in which the datastore exists
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Prefix of BigQuery table Required for `bigquery` target_type.
        /// </summary>
        public readonly string TablePrefix;

        [OutputConstructor]
        private GoogleCloudApigeeV1DatastoreConfigResponse(
            string bucketName,

            string datasetName,

            string path,

            string projectId,

            string tablePrefix)
        {
            BucketName = bucketName;
            DatasetName = datasetName;
            Path = path;
            ProjectId = projectId;
            TablePrefix = tablePrefix;
        }
    }
}