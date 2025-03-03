// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Peering.
 */
export function getPeering(args: GetPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetPeeringResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:managedidentities/v1beta1:getPeering", {
        "peeringId": args.peeringId,
        "project": args.project,
    }, opts);
}

export interface GetPeeringArgs {
    peeringId: string;
    project?: string;
}

export interface GetPeeringResult {
    /**
     * The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
     */
    readonly authorizedNetwork: string;
    /**
     * The time the instance was created.
     */
    readonly createTime: string;
    /**
     * Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
     */
    readonly domainResource: string;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
     */
    readonly name: string;
    /**
     * The current state of this Peering.
     */
    readonly state: string;
    /**
     * Additional information about the current status of this peering, if available.
     */
    readonly statusMessage: string;
    /**
     * Last update time.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single Peering.
 */
export function getPeeringOutput(args: GetPeeringOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPeeringResult> {
    return pulumi.output(args).apply((a: any) => getPeering(a, opts))
}

export interface GetPeeringOutputArgs {
    peeringId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
