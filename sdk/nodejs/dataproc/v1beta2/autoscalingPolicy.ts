// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates new autoscaling policy.
 * Auto-naming is currently not supported for this resource.
 */
export class AutoscalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AutoscalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AutoscalingPolicy {
        return new AutoscalingPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1beta2:AutoscalingPolicy';

    /**
     * Returns true if the given object is an instance of AutoscalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoscalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoscalingPolicy.__pulumiType;
    }

    public readonly basicAlgorithm!: pulumi.Output<outputs.dataproc.v1beta2.BasicAutoscalingAlgorithmResponse>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. Describes how the autoscaler will operate for secondary workers.
     */
    public readonly secondaryWorkerConfig!: pulumi.Output<outputs.dataproc.v1beta2.InstanceGroupAutoscalingPolicyConfigResponse>;
    /**
     * Describes how the autoscaler will operate for primary workers.
     */
    public readonly workerConfig!: pulumi.Output<outputs.dataproc.v1beta2.InstanceGroupAutoscalingPolicyConfigResponse>;

    /**
     * Create a AutoscalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoscalingPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'id'");
            }
            if ((!args || args.workerConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerConfig'");
            }
            resourceInputs["basicAlgorithm"] = args ? args.basicAlgorithm : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["secondaryWorkerConfig"] = args ? args.secondaryWorkerConfig : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["basicAlgorithm"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["secondaryWorkerConfig"] = undefined /*out*/;
            resourceInputs["workerConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AutoscalingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AutoscalingPolicy resource.
 */
export interface AutoscalingPolicyArgs {
    basicAlgorithm?: pulumi.Input<inputs.dataproc.v1beta2.BasicAutoscalingAlgorithmArgs>;
    /**
     * The policy id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
     */
    id: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Describes how the autoscaler will operate for secondary workers.
     */
    secondaryWorkerConfig?: pulumi.Input<inputs.dataproc.v1beta2.InstanceGroupAutoscalingPolicyConfigArgs>;
    /**
     * Describes how the autoscaler will operate for primary workers.
     */
    workerConfig: pulumi.Input<inputs.dataproc.v1beta2.InstanceGroupAutoscalingPolicyConfigArgs>;
}
