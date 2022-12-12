// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a document.
 */
export class Document extends pulumi.CustomResource {
    /**
     * Get an existing Document resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Document {
        return new Document(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:contentwarehouse/v1:Document';

    /**
     * Returns true if the given object is an instance of Document.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Document {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Document.__pulumiType;
    }

    /**
     * If true, makes the document visible to asynchronous policies and rules.
     */
    public readonly asyncEnabled!: pulumi.Output<boolean>;
    /**
     * Document AI format to save the structured content, including OCR.
     */
    public readonly cloudAiDocument!: pulumi.Output<outputs.contentwarehouse.v1.GoogleCloudDocumentaiV1DocumentResponse>;
    /**
     * The time when the document is created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user who creates the document.
     */
    public readonly creator!: pulumi.Output<string>;
    /**
     * Display name of the document given by the user. This name will be displayed in the UI. Customer can populate this field with the name of the document. This differs from the 'title' field as 'title' is optional and stores the top heading in the document.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Uri to display the document, for example, in the UI.
     */
    public readonly displayUri!: pulumi.Output<string>;
    /**
     * The Document schema name. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
     */
    public readonly documentSchemaName!: pulumi.Output<string>;
    /**
     * Raw document content.
     */
    public readonly inlineRawDocument!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the document. Format: projects/{project_number}/locations/{location}/documents/{document_id}. The name is ignored when creating a document.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Other document format, such as PPTX, XLXS
     */
    public readonly plainText!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * List of values that are user supplied metadata.
     */
    public readonly properties!: pulumi.Output<outputs.contentwarehouse.v1.GoogleCloudContentwarehouseV1PropertyResponse[]>;
    /**
     * This is used when DocAI was not used to load the document and parsing/ extracting is needed for the inline_raw_document. For example, if inline_raw_document is the byte representation of a PDF file, then this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
     */
    public readonly rawDocumentFileType!: pulumi.Output<string>;
    /**
     * Raw document file in Cloud Storage path.
     */
    public readonly rawDocumentPath!: pulumi.Output<string>;
    /**
     * The reference ID set by customers. Must be unique per project and location.
     */
    public readonly referenceId!: pulumi.Output<string>;
    /**
     * A path linked to structured content file.
     */
    public readonly structuredContentUri!: pulumi.Output<string>;
    /**
     * If true, text extraction will not be performed.
     */
    public readonly textExtractionDisabled!: pulumi.Output<boolean>;
    /**
     * Title that describes the document. This is usually present in the top section of the document, and is a mandatory field for the question-answering feature.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The time when the document is last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The user who lastly updates the document.
     */
    public readonly updater!: pulumi.Output<string>;

    /**
     * Create a Document resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["asyncEnabled"] = args ? args.asyncEnabled : undefined;
            resourceInputs["cloudAiDocument"] = args ? args.cloudAiDocument : undefined;
            resourceInputs["cloudAiDocumentOption"] = args ? args.cloudAiDocumentOption : undefined;
            resourceInputs["createMask"] = args ? args.createMask : undefined;
            resourceInputs["creator"] = args ? args.creator : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["displayUri"] = args ? args.displayUri : undefined;
            resourceInputs["documentSchemaName"] = args ? args.documentSchemaName : undefined;
            resourceInputs["inlineRawDocument"] = args ? args.inlineRawDocument : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["plainText"] = args ? args.plainText : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["rawDocumentFileType"] = args ? args.rawDocumentFileType : undefined;
            resourceInputs["rawDocumentPath"] = args ? args.rawDocumentPath : undefined;
            resourceInputs["referenceId"] = args ? args.referenceId : undefined;
            resourceInputs["requestMetadata"] = args ? args.requestMetadata : undefined;
            resourceInputs["structuredContentUri"] = args ? args.structuredContentUri : undefined;
            resourceInputs["textExtractionDisabled"] = args ? args.textExtractionDisabled : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["updater"] = args ? args.updater : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["asyncEnabled"] = undefined /*out*/;
            resourceInputs["cloudAiDocument"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["creator"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["displayUri"] = undefined /*out*/;
            resourceInputs["documentSchemaName"] = undefined /*out*/;
            resourceInputs["inlineRawDocument"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["plainText"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["rawDocumentFileType"] = undefined /*out*/;
            resourceInputs["rawDocumentPath"] = undefined /*out*/;
            resourceInputs["referenceId"] = undefined /*out*/;
            resourceInputs["structuredContentUri"] = undefined /*out*/;
            resourceInputs["textExtractionDisabled"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["updater"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Document.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Document resource.
 */
export interface DocumentArgs {
    /**
     * If true, makes the document visible to asynchronous policies and rules.
     */
    asyncEnabled?: pulumi.Input<boolean>;
    /**
     * Document AI format to save the structured content, including OCR.
     */
    cloudAiDocument?: pulumi.Input<inputs.contentwarehouse.v1.GoogleCloudDocumentaiV1DocumentArgs>;
    /**
     * Request Option for processing Cloud AI Document in CW Document.
     */
    cloudAiDocumentOption?: pulumi.Input<inputs.contentwarehouse.v1.GoogleCloudContentwarehouseV1CloudAIDocumentOptionArgs>;
    /**
     * Field mask for creating Document fields. If mask path is empty, it means all fields are masked. For the `FieldMask` definition, see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
     */
    createMask?: pulumi.Input<string>;
    /**
     * The user who creates the document.
     */
    creator?: pulumi.Input<string>;
    /**
     * Display name of the document given by the user. This name will be displayed in the UI. Customer can populate this field with the name of the document. This differs from the 'title' field as 'title' is optional and stores the top heading in the document.
     */
    displayName: pulumi.Input<string>;
    /**
     * Uri to display the document, for example, in the UI.
     */
    displayUri?: pulumi.Input<string>;
    /**
     * The Document schema name. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
     */
    documentSchemaName?: pulumi.Input<string>;
    /**
     * Raw document content.
     */
    inlineRawDocument?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the document. Format: projects/{project_number}/locations/{location}/documents/{document_id}. The name is ignored when creating a document.
     */
    name?: pulumi.Input<string>;
    /**
     * Other document format, such as PPTX, XLXS
     */
    plainText?: pulumi.Input<string>;
    /**
     * Default document policy during creation. Conditions defined in the policy will be ignored.
     */
    policy?: pulumi.Input<inputs.contentwarehouse.v1.GoogleIamV1PolicyArgs>;
    project?: pulumi.Input<string>;
    /**
     * List of values that are user supplied metadata.
     */
    properties?: pulumi.Input<pulumi.Input<inputs.contentwarehouse.v1.GoogleCloudContentwarehouseV1PropertyArgs>[]>;
    /**
     * This is used when DocAI was not used to load the document and parsing/ extracting is needed for the inline_raw_document. For example, if inline_raw_document is the byte representation of a PDF file, then this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
     */
    rawDocumentFileType?: pulumi.Input<enums.contentwarehouse.v1.DocumentRawDocumentFileType>;
    /**
     * Raw document file in Cloud Storage path.
     */
    rawDocumentPath?: pulumi.Input<string>;
    /**
     * The reference ID set by customers. Must be unique per project and location.
     */
    referenceId?: pulumi.Input<string>;
    /**
     * The meta information collected about the end user, used to enforce access control for the service.
     */
    requestMetadata?: pulumi.Input<inputs.contentwarehouse.v1.GoogleCloudContentwarehouseV1RequestMetadataArgs>;
    /**
     * A path linked to structured content file.
     */
    structuredContentUri?: pulumi.Input<string>;
    /**
     * If true, text extraction will not be performed.
     */
    textExtractionDisabled?: pulumi.Input<boolean>;
    /**
     * Title that describes the document. This is usually present in the top section of the document, and is a mandatory field for the question-answering feature.
     */
    title?: pulumi.Input<string>;
    /**
     * The user who lastly updates the document.
     */
    updater?: pulumi.Input<string>;
}