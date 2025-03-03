// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get an existing environment.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:composer/v1:getEnvironment", {
        "environmentId": args.environmentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEnvironmentArgs {
    environmentId: string;
    location: string;
    project?: string;
}

export interface GetEnvironmentResult {
    /**
     * Configuration parameters for this environment.
     */
    readonly config: outputs.composer.v1.EnvironmentConfigResponse;
    /**
     * The time at which this environment was created.
     */
    readonly createTime: string;
    /**
     * Optional. User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \p{Ll}\p{Lo}{0,62} * Values must conform to regexp: [\p{Ll}\p{Lo}\p{N}_-]{0,63} * Both keys and values are additionally constrained to be <= 128 bytes in size.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the environment, in the form: "projects/{projectId}/locations/{locationId}/environments/{environmentId}" EnvironmentId must start with a lowercase letter followed by up to 63 lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
     */
    readonly name: string;
    /**
     * The current state of the environment.
     */
    readonly state: string;
    /**
     * The time at which this environment was last modified.
     */
    readonly updateTime: string;
    /**
     * The UUID (Universally Unique IDentifier) associated with this environment. This value is generated when the environment is created.
     */
    readonly uuid: string;
}
/**
 * Get an existing environment.
 */
export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply((a: any) => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
