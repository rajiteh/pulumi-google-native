// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets the details of a group.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:migrationcenter/v1alpha1:getGroup", {
        "groupId": args.groupId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetGroupArgs {
    groupId: string;
    location: string;
    project?: string;
}

export interface GetGroupResult {
    /**
     * The timestamp when the group was created.
     */
    readonly createTime: string;
    /**
     * The description of the resource.
     */
    readonly description: string;
    /**
     * User-friendly display name.
     */
    readonly displayName: string;
    /**
     * Labels as key value pairs.
     */
    readonly labels: {[key: string]: string};
    /**
     * The name of the group.
     */
    readonly name: string;
    /**
     * The timestamp when the group was last updated.
     */
    readonly updateTime: string;
}
/**
 * Gets the details of a group.
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply((a: any) => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    groupId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
