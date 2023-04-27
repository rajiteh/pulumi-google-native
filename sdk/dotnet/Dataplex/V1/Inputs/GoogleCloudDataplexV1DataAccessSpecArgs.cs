// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// DataAccessSpec holds the access control configuration to be enforced on data stored within resources (eg: rows, columns in BigQuery Tables). When associated with data, the data is only accessible to principals explicitly granted access through the DataAccessSpec. Principals with access to the containing resource are not implicitly granted access.
    /// </summary>
    public sealed class GoogleCloudDataplexV1DataAccessSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("readers")]
        private InputList<string>? _readers;

        /// <summary>
        /// Optional. The format of strings follows the pattern followed by IAM in the bindings. user:{email}, serviceAccount:{email} group:{email}. The set of principals to be granted reader role on data stored within resources.
        /// </summary>
        public InputList<string> Readers
        {
            get => _readers ?? (_readers = new InputList<string>());
            set => _readers = value;
        }

        public GoogleCloudDataplexV1DataAccessSpecArgs()
        {
        }
        public static new GoogleCloudDataplexV1DataAccessSpecArgs Empty => new GoogleCloudDataplexV1DataAccessSpecArgs();
    }
}