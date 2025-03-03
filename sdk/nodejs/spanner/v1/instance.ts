// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an instance and begins preparing it to begin serving. The returned long-running operation can be used to track the progress of preparing the new instance. The instance name is assigned by the caller. If the named instance already exists, `CreateInstance` returns `ALREADY_EXISTS`. Immediately upon completion of this request: * The instance is readable via the API, with all requested attributes but no allocated resources. Its state is `CREATING`. Until completion of the returned operation: * Cancelling the operation renders the instance immediately unreadable via the API. * The instance can be deleted. * All other attempts to modify the instance are rejected. Upon completion of the returned operation: * Billing for all successfully-allocated resources begins (some types may have lower than the requested levels). * Databases can be created in the instance. * The instance's allocated resource levels are readable via the API. * The instance's state becomes `READY`. The returned long-running operation will have a name of the format `/operations/` and can be used to track creation of the instance. The metadata field type is CreateInstanceMetadata. The response field type is Instance, if successful.
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
    public static readonly __pulumiType = 'google-native:spanner/v1:Instance';

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
     * The name of the instance's configuration. Values are of the form `projects//instanceConfigs/`. See also InstanceConfig and ListInstanceConfigs.
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The time at which the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Deprecated. This field is not populated.
     *
     * @deprecated Deprecated. This field is not populated.
     */
    public readonly endpointUris!: pulumi.Output<string[]>;
    /**
     * Free instance metadata. Only populated for free instances.
     */
    public readonly freeInstanceMetadata!: pulumi.Output<outputs.spanner.v1.FreeInstanceMetadataResponse>;
    /**
     * The `InstanceType` of the current instance.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. And so you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. Values are of the form `projects//instances/a-z*[a-z0-9]`. The final segment of the name must be between 2 and 64 characters in length.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of nodes allocated to this instance. At most one of either node_count or processing_units should be present in the message. Users can set the node_count field to specify the target number of nodes allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The number of processing units allocated to this instance. At most one of processing_units or node_count should be present in the message. Users can set the processing_units field to specify the target number of processing units allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    public readonly processingUnits!: pulumi.Output<number>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The current instance state. For CreateInstance, the state must be either omitted or set to `CREATING`. For UpdateInstance, the state must be either omitted or set to `READY`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The time at which the instance was most recently updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

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
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endpointUris"] = args ? args.endpointUris : undefined;
            resourceInputs["freeInstanceMetadata"] = args ? args.freeInstanceMetadata : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["processingUnits"] = args ? args.processingUnits : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["endpointUris"] = undefined /*out*/;
            resourceInputs["freeInstanceMetadata"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodeCount"] = undefined /*out*/;
            resourceInputs["processingUnits"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The name of the instance's configuration. Values are of the form `projects//instanceConfigs/`. See also InstanceConfig and ListInstanceConfigs.
     */
    config: pulumi.Input<string>;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length.
     */
    displayName: pulumi.Input<string>;
    /**
     * Deprecated. This field is not populated.
     *
     * @deprecated Deprecated. This field is not populated.
     */
    endpointUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Free instance metadata. Only populated for free instances.
     */
    freeInstanceMetadata?: pulumi.Input<inputs.spanner.v1.FreeInstanceMetadataArgs>;
    /**
     * The ID of the instance to create. Valid identifiers are of the form `a-z*[a-z0-9]` and must be between 2 and 64 characters in length.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The `InstanceType` of the current instance.
     */
    instanceType?: pulumi.Input<enums.spanner.v1.InstanceInstanceType>;
    /**
     * Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies. Cloud Labels can be used to filter collections of resources. They can be used to control how resource metrics are aggregated. And they can be used as arguments to policy management rules (e.g. route, firewall, load balancing, etc.). * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `a-z{0,62}`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `[a-z0-9_-]{0,63}`. * No more than 64 labels can be associated with a given resource. See https://goo.gl/xmQnxf for more information on and examples of labels. If you plan to use labels in your own code, please note that additional characters may be allowed in the future. And so you are advised to use an internal label representation, such as JSON, which doesn't rely upon specific characters being disallowed. For example, representing labels as the string: name + "_" + value would prove problematic if we were to allow "_" in a future release.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. Values are of the form `projects//instances/a-z*[a-z0-9]`. The final segment of the name must be between 2 and 64 characters in length.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes allocated to this instance. At most one of either node_count or processing_units should be present in the message. Users can set the node_count field to specify the target number of nodes allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The number of processing units allocated to this instance. At most one of processing_units or node_count should be present in the message. Users can set the processing_units field to specify the target number of processing units allocated to the instance. This may be zero in API responses for instances that are not yet in state `READY`. See [the documentation](https://cloud.google.com/spanner/docs/compute-capacity) for more information about nodes and processing units.
     */
    processingUnits?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
}
