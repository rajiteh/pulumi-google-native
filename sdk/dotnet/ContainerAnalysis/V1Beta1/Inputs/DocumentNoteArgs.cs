// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// DocumentNote represents an SPDX Document Creation Information section: https://spdx.github.io/spdx-spec/v2.3/document-creation-information/
    /// </summary>
    public sealed class DocumentNoteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Compliance with the SPDX specification includes populating the SPDX fields therein with data related to such fields ("SPDX-Metadata")
        /// </summary>
        [Input("dataLicence")]
        public Input<string>? DataLicence { get; set; }

        /// <summary>
        /// Provide a reference number that can be used to understand how to parse and interpret the rest of the file
        /// </summary>
        [Input("spdxVersion")]
        public Input<string>? SpdxVersion { get; set; }

        public DocumentNoteArgs()
        {
        }
        public static new DocumentNoteArgs Empty => new DocumentNoteArgs();
    }
}
