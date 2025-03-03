// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2.Outputs
{

    /// <summary>
    /// The location of the function source code.
    /// </summary>
    [OutputType]
    public sealed class SourceResponse
    {
        /// <summary>
        /// If provided, get the source from this location in a Cloud Source Repository.
        /// </summary>
        public readonly Outputs.RepoSourceResponse RepoSource;
        /// <summary>
        /// If provided, get the source from this location in Google Cloud Storage.
        /// </summary>
        public readonly Outputs.StorageSourceResponse StorageSource;

        [OutputConstructor]
        private SourceResponse(
            Outputs.RepoSourceResponse repoSource,

            Outputs.StorageSourceResponse storageSource)
        {
            RepoSource = repoSource;
            StorageSource = storageSource;
        }
    }
}
