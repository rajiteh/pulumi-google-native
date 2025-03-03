// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// Publisher contains information about the publisher of this Note.
    /// </summary>
    [OutputType]
    public sealed class PublisherResponse
    {
        /// <summary>
        /// Provides information about the authority of the issuing party to release the document, in particular, the party's constituency and responsibilities or other obligations.
        /// </summary>
        public readonly string IssuingAuthority;
        /// <summary>
        /// Name of the publisher. Examples: 'Google', 'Google Cloud Platform'.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The context or namespace. Contains a URL which is under control of the issuing party and can be used as a globally unique identifier for that issuing party. Example: https://csaf.io
        /// </summary>
        public readonly string PublisherNamespace;

        [OutputConstructor]
        private PublisherResponse(
            string issuingAuthority,

            string name,

            string publisherNamespace)
        {
            IssuingAuthority = issuingAuthority;
            Name = name;
            PublisherNamespace = publisherNamespace;
        }
    }
}
