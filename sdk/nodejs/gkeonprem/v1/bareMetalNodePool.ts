// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new bare metal node pool in a given project, location and Bare Metal cluster.
 */
export class BareMetalNodePool extends pulumi.CustomResource {
    /**
     * Get an existing BareMetalNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BareMetalNodePool {
        return new BareMetalNodePool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gkeonprem/v1:BareMetalNodePool';

    /**
     * Returns true if the given object is an instance of BareMetalNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BareMetalNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BareMetalNodePool.__pulumiType;
    }

    /**
     * Annotations on the bare metal node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    public readonly bareMetalClusterId!: pulumi.Output<string>;
    /**
     * The ID to use for the node pool, which will become the final component of the node pool's resource name. This value must be up to 63 characters, and valid characters are /a-z-/. The value must not be permitted to be a UUID (or UUID-like: anything matching /^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$/i).
     */
    public readonly bareMetalNodePoolId!: pulumi.Output<string | undefined>;
    /**
     * The time at which this bare metal node pool was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time at which this bare metal node pool was deleted. If the resource is not deleted, this must be empty
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * The display name for the bare metal node pool.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
     */
    public readonly etag!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The bare metal node pool resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Node pool configuration.
     */
    public readonly nodePoolConfig!: pulumi.Output<outputs.gkeonprem.v1.BareMetalNodePoolConfigResponse>;
    public readonly project!: pulumi.Output<string>;
    /**
     * If set, there are currently changes in flight to the bare metal node pool.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * The current state of the bare metal node pool.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * ResourceStatus representing the detailed node pool status.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.gkeonprem.v1.ResourceStatusResponse>;
    /**
     * The unique identifier of the bare metal node pool.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time at which this bare metal node pool was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a BareMetalNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BareMetalNodePoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bareMetalClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bareMetalClusterId'");
            }
            if ((!args || args.nodePoolConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodePoolConfig'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["bareMetalClusterId"] = args ? args.bareMetalClusterId : undefined;
            resourceInputs["bareMetalNodePoolId"] = args ? args.bareMetalNodePoolId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodePoolConfig"] = args ? args.nodePoolConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["bareMetalClusterId"] = undefined /*out*/;
            resourceInputs["bareMetalNodePoolId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodePoolConfig"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["bareMetalClusterId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(BareMetalNodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BareMetalNodePool resource.
 */
export interface BareMetalNodePoolArgs {
    /**
     * Annotations on the bare metal node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    bareMetalClusterId: pulumi.Input<string>;
    /**
     * The ID to use for the node pool, which will become the final component of the node pool's resource name. This value must be up to 63 characters, and valid characters are /a-z-/. The value must not be permitted to be a UUID (or UUID-like: anything matching /^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$/i).
     */
    bareMetalNodePoolId?: pulumi.Input<string>;
    /**
     * The display name for the bare metal node pool.
     */
    displayName?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
     */
    etag?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The bare metal node pool resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Node pool configuration.
     */
    nodePoolConfig: pulumi.Input<inputs.gkeonprem.v1.BareMetalNodePoolConfigArgs>;
    project?: pulumi.Input<string>;
}