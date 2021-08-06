// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new ResourceRecordSet.
 */
export class ResourceRecordSet extends pulumi.CustomResource {
    /**
     * Get an existing ResourceRecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourceRecordSet {
        return new ResourceRecordSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dns/v1:ResourceRecordSet';

    /**
     * Returns true if the given object is an instance of ResourceRecordSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceRecordSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceRecordSet.__pulumiType;
    }

    public readonly kind!: pulumi.Output<string>;
    /**
     * For example, www.example.com.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
     */
    public readonly rrdatas!: pulumi.Output<string[]>;
    /**
     * As defined in RFC 4034 (section 3.2).
     */
    public readonly signatureRrdatas!: pulumi.Output<string[]>;
    /**
     * Number of seconds that this ResourceRecordSet can be cached by resolvers.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The identifier of a supported record type. See the list of Supported DNS record types.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ResourceRecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceRecordSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.managedZone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedZone'");
            }
            inputs["clientOperationId"] = args ? args.clientOperationId : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managedZone"] = args ? args.managedZone : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rrdatas"] = args ? args.rrdatas : undefined;
            inputs["signatureRrdatas"] = args ? args.signatureRrdatas : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
        } else {
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["rrdatas"] = undefined /*out*/;
            inputs["signatureRrdatas"] = undefined /*out*/;
            inputs["ttl"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResourceRecordSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourceRecordSet resource.
 */
export interface ResourceRecordSetArgs {
    clientOperationId?: pulumi.Input<string>;
    kind?: pulumi.Input<string>;
    managedZone: pulumi.Input<string>;
    /**
     * For example, www.example.com.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
     */
    rrdatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * As defined in RFC 4034 (section 3.2).
     */
    signatureRrdatas?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of seconds that this ResourceRecordSet can be cached by resolvers.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The identifier of a supported record type. See the list of Supported DNS record types.
     */
    type?: pulumi.Input<string>;
}