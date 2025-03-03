// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Mesh.
 */
export function getMesh(args: GetMeshArgs, opts?: pulumi.InvokeOptions): Promise<GetMeshResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:networkservices/v1:getMesh", {
        "location": args.location,
        "meshId": args.meshId,
        "project": args.project,
    }, opts);
}

export interface GetMeshArgs {
    location: string;
    meshId: string;
    project?: string;
}

export interface GetMeshResult {
    /**
     * The timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    readonly description: string;
    /**
     * Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This will is applicable only for sidecar proxy deployments.
     */
    readonly interceptionPort: number;
    /**
     * Optional. Set of label tags associated with the Mesh resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the Mesh resource. It matches pattern `projects/*&#47;locations/global/meshes/`.
     */
    readonly name: string;
    /**
     * Server-defined URL of this resource
     */
    readonly selfLink: string;
    /**
     * The timestamp when the resource was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single Mesh.
 */
export function getMeshOutput(args: GetMeshOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMeshResult> {
    return pulumi.output(args).apply((a: any) => getMesh(a, opts))
}

export interface GetMeshOutputArgs {
    location: pulumi.Input<string>;
    meshId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
