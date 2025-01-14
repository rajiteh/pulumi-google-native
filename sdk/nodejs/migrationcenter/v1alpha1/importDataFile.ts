// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an import data file.
 * Auto-naming is currently not supported for this resource.
 */
export class ImportDataFile extends pulumi.CustomResource {
    /**
     * Get an existing ImportDataFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ImportDataFile {
        return new ImportDataFile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:migrationcenter/v1alpha1:ImportDataFile';

    /**
     * Returns true if the given object is an instance of ImportDataFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImportDataFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImportDataFile.__pulumiType;
    }

    /**
     * The timestamp when the file was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * User-friendly display name. Maximum length is 63 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The payload format.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Required. The ID of the new data file.
     */
    public readonly importDataFileId!: pulumi.Output<string>;
    public readonly importJobId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the file.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The state of the import data file.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Information about a file that is uploaded to a storage service.
     */
    public readonly uploadFileInfo!: pulumi.Output<outputs.migrationcenter.v1alpha1.UploadFileInfoResponse>;

    /**
     * Create a ImportDataFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImportDataFileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.format === undefined) && !opts.urn) {
                throw new Error("Missing required property 'format'");
            }
            if ((!args || args.importDataFileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'importDataFileId'");
            }
            if ((!args || args.importJobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'importJobId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["importDataFileId"] = args ? args.importDataFileId : undefined;
            resourceInputs["importJobId"] = args ? args.importJobId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["uploadFileInfo"] = args ? args.uploadFileInfo : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["format"] = undefined /*out*/;
            resourceInputs["importDataFileId"] = undefined /*out*/;
            resourceInputs["importJobId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uploadFileInfo"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["importDataFileId", "importJobId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ImportDataFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ImportDataFile resource.
 */
export interface ImportDataFileArgs {
    /**
     * User-friendly display name. Maximum length is 63 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The payload format.
     */
    format: pulumi.Input<enums.migrationcenter.v1alpha1.ImportDataFileFormat>;
    /**
     * Required. The ID of the new data file.
     */
    importDataFileId: pulumi.Input<string>;
    importJobId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Information about a file that is uploaded to a storage service.
     */
    uploadFileInfo?: pulumi.Input<inputs.migrationcenter.v1alpha1.UploadFileInfoArgs>;
}
