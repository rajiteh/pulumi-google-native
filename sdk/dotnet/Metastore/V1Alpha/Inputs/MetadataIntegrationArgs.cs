// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha.Inputs
{

    /// <summary>
    /// Specifies how metastore metadata should be integrated with external services.
    /// </summary>
    public sealed class MetadataIntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The integration config for the Data Catalog service.
        /// </summary>
        [Input("dataCatalogConfig")]
        public Input<Inputs.DataCatalogConfigArgs>? DataCatalogConfig { get; set; }

        /// <summary>
        /// The integration config for the Dataplex service.
        /// </summary>
        [Input("dataplexConfig")]
        public Input<Inputs.DataplexConfigArgs>? DataplexConfig { get; set; }

        public MetadataIntegrationArgs()
        {
        }
        public static new MetadataIntegrationArgs Empty => new MetadataIntegrationArgs();
    }
}
