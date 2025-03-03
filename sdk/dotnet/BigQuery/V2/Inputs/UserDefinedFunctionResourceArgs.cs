// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    /// <summary>
    /// This is used for defining User Defined Function (UDF) resources only when using legacy SQL. Users of Standard SQL should leverage either DDL (e.g. CREATE [TEMPORARY] FUNCTION ... ) or the Routines API to define UDF resources. For additional information on migrating, see: https://cloud.google.com/bigquery/docs/reference/standard-sql/migrating-from-legacy-sql#differences_in_user-defined_javascript_functions
    /// </summary>
    public sealed class UserDefinedFunctionResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Pick one] An inline resource that contains code for a user-defined function (UDF). Providing a inline code resource is equivalent to providing a URI for a file containing the same code.
        /// </summary>
        [Input("inlineCode")]
        public Input<string>? InlineCode { get; set; }

        /// <summary>
        /// [Pick one] A code resource to load from a Google Cloud Storage URI (gs://bucket/path).
        /// </summary>
        [Input("resourceUri")]
        public Input<string>? ResourceUri { get; set; }

        public UserDefinedFunctionResourceArgs()
        {
        }
        public static new UserDefinedFunctionResourceArgs Empty => new UserDefinedFunctionResourceArgs();
    }
}
