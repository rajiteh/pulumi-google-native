// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Endpoint.
 */
export function getEndpoint(args: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:ids/v1:getEndpoint", {
        "endpointId": args.endpointId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEndpointArgs {
    endpointId: string;
    location: string;
    project?: string;
}

export interface GetEndpointResult {
    /**
     * The create time timestamp.
     */
    readonly createTime: string;
    /**
     * User-provided description of the endpoint
     */
    readonly description: string;
    /**
     * The fully qualified URL of the endpoint's ILB Forwarding Rule.
     */
    readonly endpointForwardingRule: string;
    /**
     * The IP address of the IDS Endpoint's ILB.
     */
    readonly endpointIp: string;
    /**
     * The labels of the endpoint.
     */
    readonly labels: {[key: string]: string};
    /**
     * The name of the endpoint.
     */
    readonly name: string;
    /**
     * The fully qualified URL of the network to which the IDS Endpoint is attached.
     */
    readonly network: string;
    /**
     * Lowest threat severity that this endpoint will alert on.
     */
    readonly severity: string;
    /**
     * Current state of the endpoint.
     */
    readonly state: string;
    /**
     * Whether the endpoint should report traffic logs in addition to threat logs.
     */
    readonly trafficLogs: boolean;
    /**
     * The update time timestamp.
     */
    readonly updateTime: string;
}

export function getEndpointOutput(args: GetEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointResult> {
    return pulumi.output(args).apply(a => getEndpoint(a, opts))
}

export interface GetEndpointOutputArgs {
    endpointId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}