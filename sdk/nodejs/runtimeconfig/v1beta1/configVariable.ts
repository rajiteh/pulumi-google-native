// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a variable within the given configuration. You cannot create a variable with a name that is a prefix of an existing variable name, or a name that has an existing variable name as a prefix. To learn more about creating a variable, read the [Setting and Getting Data](/deployment-manager/runtime-configurator/set-and-get-variables) documentation.
 */
export class ConfigVariable extends pulumi.CustomResource {
    /**
     * Get an existing ConfigVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigVariable {
        return new ConfigVariable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp-native:runtimeconfig/v1beta1:ConfigVariable';

    /**
     * Returns true if the given object is an instance of ConfigVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigVariable.__pulumiType;
    }

    /**
     * The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
     */
    public readonly text!: pulumi.Output<string>;
    /**
     * The time of the last variable update. Timestamp will be UTC timestamp.
     */
    public readonly updateTime!: pulumi.Output<string>;
    /**
     * The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ConfigVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigVariableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            if ((!args || args.variablesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variablesId'");
            }
            inputs["configsId"] = args ? args.configsId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["text"] = args ? args.text : undefined;
            inputs["updateTime"] = args ? args.updateTime : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["variablesId"] = args ? args.variablesId : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["text"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConfigVariable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigVariable resource.
 */
export interface ConfigVariableArgs {
    readonly configsId: pulumi.Input<string>;
    /**
     * The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
     */
    readonly name?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
     */
    readonly text?: pulumi.Input<string>;
    /**
     * The time of the last variable update. Timestamp will be UTC timestamp.
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
     */
    readonly value?: pulumi.Input<string>;
    readonly variablesId: pulumi.Input<string>;
}