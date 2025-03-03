// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DiscoveryEngine.V1Beta
{
    /// <summary>
    /// Creates a Document.
    /// </summary>
    [GoogleNativeResourceType("google-native:discoveryengine/v1beta:Document")]
    public partial class Document : global::Pulumi.CustomResource
    {
        [Output("branchId")]
        public Output<string> BranchId { get; private set; } = null!;

        [Output("collectionId")]
        public Output<string> CollectionId { get; private set; } = null!;

        [Output("dataStoreId")]
        public Output<string> DataStoreId { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the Document, which will become the final component of the Document.name. If the caller does not have permission to create the Document, regardless of whether or not it exists, a `PERMISSION_DENIED` error is returned. This field must be unique among all Documents with the same parent. Otherwise, an `ALREADY_EXISTS` error is returned. This field must conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
        /// </summary>
        [Output("documentId")]
        public Output<string> DocumentId { get; private set; } = null!;

        /// <summary>
        /// The JSON string representation of the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        /// </summary>
        [Output("jsonData")]
        public Output<string> JsonData { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The full resource name of the document. Format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The identifier of the parent document. Currently supports at most two level document hierarchy. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
        /// </summary>
        [Output("parentDocumentId")]
        public Output<string> ParentDocumentId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The identifier of the schema located in the same data store.
        /// </summary>
        [Output("schemaId")]
        public Output<string> SchemaId { get; private set; } = null!;

        /// <summary>
        /// The structured JSON data for the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        /// </summary>
        [Output("structData")]
        public Output<ImmutableDictionary<string, string>> StructData { get; private set; } = null!;


        /// <summary>
        /// Create a Document resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Document(string name, DocumentArgs args, CustomResourceOptions? options = null)
            : base("google-native:discoveryengine/v1beta:Document", name, args ?? new DocumentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Document(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:discoveryengine/v1beta:Document", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "branchId",
                    "collectionId",
                    "dataStoreId",
                    "documentId",
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
        /// Get an existing Document resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Document Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Document(name, id, options);
        }
    }

    public sealed class DocumentArgs : global::Pulumi.ResourceArgs
    {
        [Input("branchId", required: true)]
        public Input<string> BranchId { get; set; } = null!;

        [Input("collectionId", required: true)]
        public Input<string> CollectionId { get; set; } = null!;

        [Input("dataStoreId", required: true)]
        public Input<string> DataStoreId { get; set; } = null!;

        /// <summary>
        /// Required. The ID to use for the Document, which will become the final component of the Document.name. If the caller does not have permission to create the Document, regardless of whether or not it exists, a `PERMISSION_DENIED` error is returned. This field must be unique among all Documents with the same parent. Otherwise, an `ALREADY_EXISTS` error is returned. This field must conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
        /// </summary>
        [Input("documentId", required: true)]
        public Input<string> DocumentId { get; set; } = null!;

        /// <summary>
        /// Immutable. The identifier of the document. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The JSON string representation of the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        /// </summary>
        [Input("jsonData")]
        public Input<string>? JsonData { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The full resource name of the document. Format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The identifier of the parent document. Currently supports at most two level document hierarchy. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
        /// </summary>
        [Input("parentDocumentId")]
        public Input<string>? ParentDocumentId { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The identifier of the schema located in the same data store.
        /// </summary>
        [Input("schemaId")]
        public Input<string>? SchemaId { get; set; }

        [Input("structData")]
        private InputMap<string>? _structData;

        /// <summary>
        /// The structured JSON data for the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        /// </summary>
        public InputMap<string> StructData
        {
            get => _structData ?? (_structData = new InputMap<string>());
            set => _structData = value;
        }

        public DocumentArgs()
        {
        }
        public static new DocumentArgs Empty => new DocumentArgs();
    }
}
