// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    [OutputType]
    public sealed class BucketLoggingResponse
    {
        /// <summary>
        /// The destination bucket where the current bucket's logs should be placed.
        /// </summary>
        public readonly string LogBucket;
        /// <summary>
        /// A prefix for log object names.
        /// </summary>
        public readonly string LogObjectPrefix;

        [OutputConstructor]
        private BucketLoggingResponse(
            string logBucket,

            string logObjectPrefix)
        {
            LogBucket = logBucket;
            LogObjectPrefix = logObjectPrefix;
        }
    }
}