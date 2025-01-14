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
    /// [Optional] The categories attached to this field, used for field-level access control.
    /// </summary>
    public sealed class TableFieldSchemaCategoriesArgs : global::Pulumi.ResourceArgs
    {
        [Input("names")]
        private InputList<string>? _names;

        /// <summary>
        /// A list of category resource names. For example, "projects/1/taxonomies/2/categories/3". At most 5 categories are allowed.
        /// </summary>
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        public TableFieldSchemaCategoriesArgs()
        {
        }
        public static new TableFieldSchemaCategoriesArgs Empty => new TableFieldSchemaCategoriesArgs();
    }
}
