// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// The zip file information for a zip deployment.
    /// </summary>
    [OutputType]
    public sealed class ZipInfoResponse
    {
        /// <summary>
        /// An estimate of the number of files in a zip for a zip deployment. If set, must be greater than or equal to the actual number of files. Used for optimizing performance; if not provided, deployment may be slow.
        /// </summary>
        public readonly int FilesCount;
        /// <summary>
        /// URL of the zip file to deploy from. Must be a URL to a resource in Google Cloud Storage in the form 'http(s)://storage.googleapis.com//'.
        /// </summary>
        public readonly string SourceUrl;

        [OutputConstructor]
        private ZipInfoResponse(
            int filesCount,

            string sourceUrl)
        {
            FilesCount = filesCount;
            SourceUrl = sourceUrl;
        }
    }
}
