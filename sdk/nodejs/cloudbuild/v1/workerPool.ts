// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a `WorkerPool`.
 */
export class WorkerPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkerPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkerPool {
        return new WorkerPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudbuild/v1:WorkerPool';

    /**
     * Returns true if the given object is an instance of WorkerPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkerPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerPool.__pulumiType;
    }

    /**
     * User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Time at which the request to create the `WorkerPool` was received.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time at which the request to delete the `WorkerPool` was received.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Private Pool using a v1 configuration.
     */
    public readonly privatePoolV1Config!: pulumi.Output<outputs.cloudbuild.v1.PrivatePoolV1ConfigResponse>;
    /**
     * `WorkerPool` state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A unique identifier for the `WorkerPool`.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Time at which the request to update the `WorkerPool` was received.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a WorkerPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkerPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.workerPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerPoolId'");
            }
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["privatePoolV1Config"] = args ? args.privatePoolV1Config : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["validateOnly"] = args ? args.validateOnly : undefined;
            inputs["workerPoolId"] = args ? args.workerPoolId : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["deleteTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["annotations"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["deleteTime"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privatePoolV1Config"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["uid"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkerPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkerPool resource.
 */
export interface WorkerPoolArgs {
    /**
     * User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
     */
    displayName?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    /**
     * Private Pool using a v1 configuration.
     */
    privatePoolV1Config?: pulumi.Input<inputs.cloudbuild.v1.PrivatePoolV1ConfigArgs>;
    project: pulumi.Input<string>;
    validateOnly?: pulumi.Input<string>;
    workerPoolId: pulumi.Input<string>;
}