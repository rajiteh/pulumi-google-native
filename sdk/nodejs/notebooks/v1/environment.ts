// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new Environment.
 * Auto-naming is currently not supported for this resource.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:notebooks/v1:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Use a container image to start the notebook instance.
     */
    public readonly containerImage!: pulumi.Output<outputs.notebooks.v1.ContainerImageResponse>;
    /**
     * The time at which this environment was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A brief description of this environment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name of this environment for the UI.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Required. User-defined unique ID of this environment. The `environment_id` must be 1 to 63 characters long and contain only lowercase letters, numeric characters, and dashes. The first character must be a lowercase letter and the last character cannot be a dash.
     */
    public readonly environmentId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of this environment. Format: `projects/{project_id}/locations/{location}/environments/{environment_id}`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
     */
    public readonly postStartupScript!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     */
    public readonly vmImage!: pulumi.Output<outputs.notebooks.v1.VmImageResponse>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            resourceInputs["containerImage"] = args ? args.containerImage : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["postStartupScript"] = args ? args.postStartupScript : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["vmImage"] = args ? args.vmImage : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["containerImage"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["postStartupScript"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["vmImage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["environmentId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Use a container image to start the notebook instance.
     */
    containerImage?: pulumi.Input<inputs.notebooks.v1.ContainerImageArgs>;
    /**
     * A brief description of this environment.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name of this environment for the UI.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Required. User-defined unique ID of this environment. The `environment_id` must be 1 to 63 characters long and contain only lowercase letters, numeric characters, and dashes. The first character must be a lowercase letter and the last character cannot be a dash.
     */
    environmentId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
     */
    postStartupScript?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     */
    vmImage?: pulumi.Input<inputs.notebooks.v1.VmImageArgs>;
}
