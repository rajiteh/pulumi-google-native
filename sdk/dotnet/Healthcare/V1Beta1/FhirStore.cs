// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    /// <summary>
    /// Creates a new FHIR store within the parent dataset.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1beta1:FhirStore")]
    public partial class FhirStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        /// </summary>
        [Output("complexDataTypeReferenceParsing")]
        public Output<string> ComplexDataTypeReferenceParsing { get; private set; } = null!;

        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        /// </summary>
        [Output("defaultSearchHandlingStrict")]
        public Output<bool> DefaultSearchHandlingStrict { get; private set; } = null!;

        /// <summary>
        /// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        /// </summary>
        [Output("disableReferentialIntegrity")]
        public Output<bool> DisableReferentialIntegrity { get; private set; } = null!;

        /// <summary>
        /// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        /// </summary>
        [Output("disableResourceVersioning")]
        public Output<bool> DisableResourceVersioning { get; private set; } = null!;

        /// <summary>
        /// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        /// </summary>
        [Output("enableUpdateCreate")]
        public Output<bool> EnableUpdateCreate { get; private set; } = null!;

        /// <summary>
        /// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
        /// </summary>
        [Output("fhirStoreId")]
        public Output<string?> FhirStoreId { get; private set; } = null!;

        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Deprecated. Use `notification_configs` instead. If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.NotificationConfigResponse> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies where and whether to send notifications upon changes to a Fhir store.
        /// </summary>
        [Output("notificationConfigs")]
        public Output<ImmutableArray<Outputs.FhirNotificationConfigResponse>> NotificationConfigs { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Configuration for how FHIR resources can be searched.
        /// </summary>
        [Output("searchConfig")]
        public Output<Outputs.SearchConfigResponse> SearchConfig { get; private set; } = null!;

        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// </summary>
        [Output("streamConfigs")]
        public Output<ImmutableArray<Outputs.StreamConfigResponse>> StreamConfigs { get; private set; } = null!;

        /// <summary>
        /// Configuration for how to validate incoming FHIR resources against configured profiles.
        /// </summary>
        [Output("validationConfig")]
        public Output<Outputs.ValidationConfigResponse> ValidationConfig { get; private set; } = null!;

        /// <summary>
        /// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a FhirStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FhirStore(string name, FhirStoreArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:FhirStore", name, args ?? new FhirStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FhirStore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:FhirStore", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "datasetId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FhirStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FhirStore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FhirStore(name, id, options);
        }
    }

    public sealed class FhirStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        /// </summary>
        [Input("complexDataTypeReferenceParsing")]
        public Input<Pulumi.GoogleNative.Healthcare.V1Beta1.FhirStoreComplexDataTypeReferenceParsing>? ComplexDataTypeReferenceParsing { get; set; }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        /// </summary>
        [Input("defaultSearchHandlingStrict")]
        public Input<bool>? DefaultSearchHandlingStrict { get; set; }

        /// <summary>
        /// Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        /// </summary>
        [Input("disableReferentialIntegrity")]
        public Input<bool>? DisableReferentialIntegrity { get; set; }

        /// <summary>
        /// Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        /// </summary>
        [Input("disableResourceVersioning")]
        public Input<bool>? DisableResourceVersioning { get; set; }

        /// <summary>
        /// Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        /// </summary>
        [Input("enableUpdateCreate")]
        public Input<bool>? EnableUpdateCreate { get; set; }

        /// <summary>
        /// The ID of the FHIR store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
        /// </summary>
        [Input("fhirStoreId")]
        public Input<string>? FhirStoreId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Deprecated. Use `notification_configs` instead. If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.NotificationConfigArgs>? NotificationConfig { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.FhirNotificationConfigArgs>? _notificationConfigs;

        /// <summary>
        /// Specifies where and whether to send notifications upon changes to a Fhir store.
        /// </summary>
        public InputList<Inputs.FhirNotificationConfigArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.FhirNotificationConfigArgs>());
            set => _notificationConfigs = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configuration for how FHIR resources can be searched.
        /// </summary>
        [Input("searchConfig")]
        public Input<Inputs.SearchConfigArgs>? SearchConfig { get; set; }

        [Input("streamConfigs")]
        private InputList<Inputs.StreamConfigArgs>? _streamConfigs;

        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// </summary>
        public InputList<Inputs.StreamConfigArgs> StreamConfigs
        {
            get => _streamConfigs ?? (_streamConfigs = new InputList<Inputs.StreamConfigArgs>());
            set => _streamConfigs = value;
        }

        /// <summary>
        /// Configuration for how to validate incoming FHIR resources against configured profiles.
        /// </summary>
        [Input("validationConfig")]
        public Input<Inputs.ValidationConfigArgs>? ValidationConfig { get; set; }

        /// <summary>
        /// Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        /// </summary>
        [Input("version")]
        public Input<Pulumi.GoogleNative.Healthcare.V1Beta1.FhirStoreVersion>? Version { get; set; }

        public FhirStoreArgs()
        {
        }
        public static new FhirStoreArgs Empty => new FhirStoreArgs();
    }
}
