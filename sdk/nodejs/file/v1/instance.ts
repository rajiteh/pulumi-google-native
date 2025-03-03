// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an instance. When creating from a backup, the capacity of the new instance needs to be equal to or larger than the capacity of the backup (and also equal to or larger than the minimum capacity of the tier).
 * Auto-naming is currently not supported for this resource.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:file/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The time when the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the instance (2048 characters or less).
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * File system shares on the instance. For this version, only a single file share is supported.
     */
    public readonly fileShares!: pulumi.Output<outputs.file.v1.FileShareConfigResponse[]>;
    /**
     * Required. The name of the instance to create. The name must be unique for the specified project and location.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * KMS key name used for data encryption.
     */
    public readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the instance, in the format `projects/{project}/locations/{location}/instances/{instance}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * VPC networks to which the instance is connected. For this version, only a single network is supported.
     */
    public readonly networks!: pulumi.Output<outputs.file.v1.NetworkConfigResponse[]>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * The instance state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the instance state, if available.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Field indicates all the reasons the instance is in "SUSPENDED" state.
     */
    public /*out*/ readonly suspensionReasons!: pulumi.Output<string[]>;
    /**
     * The service tier of the instance.
     */
    public readonly tier!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["fileShares"] = args ? args.fileShares : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["suspensionReasons"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["fileShares"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networks"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["satisfiesPzs"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["suspensionReasons"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["instanceId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The description of the instance (2048 characters or less).
     */
    description?: pulumi.Input<string>;
    /**
     * Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     */
    etag?: pulumi.Input<string>;
    /**
     * File system shares on the instance. For this version, only a single file share is supported.
     */
    fileShares?: pulumi.Input<pulumi.Input<inputs.file.v1.FileShareConfigArgs>[]>;
    /**
     * Required. The name of the instance to create. The name must be unique for the specified project and location.
     */
    instanceId: pulumi.Input<string>;
    /**
     * KMS key name used for data encryption.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * VPC networks to which the instance is connected. For this version, only a single network is supported.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.file.v1.NetworkConfigArgs>[]>;
    project?: pulumi.Input<string>;
    /**
     * The service tier of the instance.
     */
    tier?: pulumi.Input<enums.file.v1.InstanceTier>;
}
