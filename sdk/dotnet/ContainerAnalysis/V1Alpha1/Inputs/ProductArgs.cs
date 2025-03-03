// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Product contains information about a product and how to uniquely identify it.
    /// </summary>
    public sealed class ProductArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Token that identifies a product so that it can be referred to from other parts in the document. There is no predefined format as long as it uniquely identifies a group in the context of the current document.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Helps in identifying the underlying product.
        /// </summary>
        [Input("identifierHelper")]
        public Input<Inputs.IdentifierHelperArgs>? IdentifierHelper { get; set; }

        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ProductArgs()
        {
        }
        public static new ProductArgs Empty => new ProductArgs();
    }
}
