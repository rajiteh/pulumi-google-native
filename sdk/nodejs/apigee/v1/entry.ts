// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates key value entries in a key value map scoped to an organization, environment, or API proxy. **Note**: Supported for Apigee hybrid 1.8.x and higher.
 */
export class Entry extends pulumi.CustomResource {
    /**
     * Get an existing Entry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Entry {
        return new Entry(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Entry';

    /**
     * Returns true if the given object is an instance of Entry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entry.__pulumiType;
    }

    public readonly apiId!: pulumi.Output<string>;
    public readonly keyvaluemapId!: pulumi.Output<string>;
    /**
     * Resource URI that can be used to identify the scope of the key value map entries.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Data or payload that is being retrieved and associated with the unique key.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Entry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.keyvaluemapId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyvaluemapId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["keyvaluemapId"] = args ? args.keyvaluemapId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
        } else {
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["keyvaluemapId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["apiId", "keyvaluemapId", "organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Entry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Entry resource.
 */
export interface EntryArgs {
    apiId: pulumi.Input<string>;
    keyvaluemapId: pulumi.Input<string>;
    /**
     * Resource URI that can be used to identify the scope of the key value map entries.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Data or payload that is being retrieved and associated with the unique key.
     */
    value: pulumi.Input<string>;
}
