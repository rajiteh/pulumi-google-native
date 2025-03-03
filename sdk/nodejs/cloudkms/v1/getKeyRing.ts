// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns metadata for a given KeyRing.
 */
export function getKeyRing(args: GetKeyRingArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyRingResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudkms/v1:getKeyRing", {
        "keyRingId": args.keyRingId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetKeyRingArgs {
    keyRingId: string;
    location: string;
    project?: string;
}

export interface GetKeyRingResult {
    /**
     * The time at which this KeyRing was created.
     */
    readonly createTime: string;
    /**
     * The resource name for the KeyRing in the format `projects/*&#47;locations/*&#47;keyRings/*`.
     */
    readonly name: string;
}
/**
 * Returns metadata for a given KeyRing.
 */
export function getKeyRingOutput(args: GetKeyRingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeyRingResult> {
    return pulumi.output(args).apply((a: any) => getKeyRing(a, opts))
}

export interface GetKeyRingOutputArgs {
    keyRingId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
