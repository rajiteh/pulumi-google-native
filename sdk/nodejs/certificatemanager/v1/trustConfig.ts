// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new TrustConfig in a given project and location.
 */
export class TrustConfig extends pulumi.CustomResource {
    /**
     * Get an existing TrustConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TrustConfig {
        return new TrustConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:certificatemanager/v1:TrustConfig';

    /**
     * Returns true if the given object is an instance of TrustConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrustConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrustConfig.__pulumiType;
    }

    /**
     * The creation timestamp of a TrustConfig.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * One or more paragraphs of text description of a TrustConfig.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Set of labels associated with a TrustConfig.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*&#47;locations/*&#47;trustConfigs/*`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Required. A user-provided name of the TrustConfig.
     */
    public readonly trustConfigId!: pulumi.Output<string>;
    /**
     * Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
     */
    public readonly trustStores!: pulumi.Output<outputs.certificatemanager.v1.TrustStoreResponse[]>;
    /**
     * The last update timestamp of a TrustConfig.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TrustConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrustConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.trustConfigId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trustConfigId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["trustConfigId"] = args ? args.trustConfigId : undefined;
            resourceInputs["trustStores"] = args ? args.trustStores : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["trustConfigId"] = undefined /*out*/;
            resourceInputs["trustStores"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project", "trustConfigId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TrustConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TrustConfig resource.
 */
export interface TrustConfigArgs {
    /**
     * One or more paragraphs of text description of a TrustConfig.
     */
    description?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Set of labels associated with a TrustConfig.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*&#47;locations/*&#47;trustConfigs/*`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Required. A user-provided name of the TrustConfig.
     */
    trustConfigId: pulumi.Input<string>;
    /**
     * Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
     */
    trustStores?: pulumi.Input<pulumi.Input<inputs.certificatemanager.v1.TrustStoreArgs>[]>;
}
