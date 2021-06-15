// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns details of a `WorkerPool`.
 */
export function getWorkerPool(args: GetWorkerPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkerPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:cloudbuild/v1beta1:getWorkerPool", {
        "location": args.location,
        "project": args.project,
        "workerPoolId": args.workerPoolId,
    }, opts);
}

export interface GetWorkerPoolArgs {
    location: string;
    project: string;
    workerPoolId: string;
}

export interface GetWorkerPoolResult {
    /**
     * Time at which the request to create the `WorkerPool` was received.
     */
    readonly createTime: string;
    /**
     * Time at which the request to delete the `WorkerPool` was received.
     */
    readonly deleteTime: string;
    /**
     * The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
     */
    readonly name: string;
    /**
     * Network configuration for the `WorkerPool`.
     */
    readonly networkConfig: outputs.cloudbuild.v1beta1.NetworkConfigResponse;
    /**
     * `WorkerPool` state.
     */
    readonly state: string;
    /**
     * Time at which the request to update the `WorkerPool` was received.
     */
    readonly updateTime: string;
    /**
     * Worker configuration for the `WorkerPool`.
     */
    readonly workerConfig: outputs.cloudbuild.v1beta1.WorkerConfigResponse;
}