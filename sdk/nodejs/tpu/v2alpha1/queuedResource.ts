// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a QueuedResource TPU instance.
 * Auto-naming is currently not supported for this resource.
 */
export class QueuedResource extends pulumi.CustomResource {
    /**
     * Get an existing QueuedResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): QueuedResource {
        return new QueuedResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:tpu/v2alpha1:QueuedResource';

    /**
     * Returns true if the given object is an instance of QueuedResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueuedResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueuedResource.__pulumiType;
    }

    /**
     * The BestEffort tier.
     */
    public readonly bestEffort!: pulumi.Output<outputs.tpu.v2alpha1.BestEffortResponse>;
    /**
     * The Guaranteed tier
     */
    public readonly guaranteed!: pulumi.Output<outputs.tpu.v2alpha1.GuaranteedResponse>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The name of the QueuedResource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The unqualified resource name. Should follow the `^[A-Za-z0-9_.~+%-]+$` regex format.
     */
    public readonly queuedResourceId!: pulumi.Output<string | undefined>;
    /**
     * The queueing policy of the QueuedRequest.
     */
    public readonly queueingPolicy!: pulumi.Output<outputs.tpu.v2alpha1.QueueingPolicyResponse>;
    /**
     * Idempotent request UUID.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Name of the reservation in which the resource should be provisioned. Format: projects/{project}/locations/{zone}/reservations/{reservation}
     */
    public readonly reservationName!: pulumi.Output<string>;
    /**
     * State of the QueuedResource request.
     */
    public /*out*/ readonly state!: pulumi.Output<outputs.tpu.v2alpha1.QueuedResourceStateResponse>;
    /**
     * Defines a TPU resource.
     */
    public readonly tpu!: pulumi.Output<outputs.tpu.v2alpha1.TpuResponse>;

    /**
     * Create a QueuedResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QueuedResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["bestEffort"] = args ? args.bestEffort : undefined;
            resourceInputs["guaranteed"] = args ? args.guaranteed : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["queuedResourceId"] = args ? args.queuedResourceId : undefined;
            resourceInputs["queueingPolicy"] = args ? args.queueingPolicy : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["reservationName"] = args ? args.reservationName : undefined;
            resourceInputs["tpu"] = args ? args.tpu : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["bestEffort"] = undefined /*out*/;
            resourceInputs["guaranteed"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["queuedResourceId"] = undefined /*out*/;
            resourceInputs["queueingPolicy"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["reservationName"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tpu"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(QueuedResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a QueuedResource resource.
 */
export interface QueuedResourceArgs {
    /**
     * The BestEffort tier.
     */
    bestEffort?: pulumi.Input<inputs.tpu.v2alpha1.BestEffortArgs>;
    /**
     * The Guaranteed tier
     */
    guaranteed?: pulumi.Input<inputs.tpu.v2alpha1.GuaranteedArgs>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The unqualified resource name. Should follow the `^[A-Za-z0-9_.~+%-]+$` regex format.
     */
    queuedResourceId?: pulumi.Input<string>;
    /**
     * The queueing policy of the QueuedRequest.
     */
    queueingPolicy?: pulumi.Input<inputs.tpu.v2alpha1.QueueingPolicyArgs>;
    /**
     * Idempotent request UUID.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Name of the reservation in which the resource should be provisioned. Format: projects/{project}/locations/{zone}/reservations/{reservation}
     */
    reservationName?: pulumi.Input<string>;
    /**
     * Defines a TPU resource.
     */
    tpu?: pulumi.Input<inputs.tpu.v2alpha1.TpuArgs>;
}
