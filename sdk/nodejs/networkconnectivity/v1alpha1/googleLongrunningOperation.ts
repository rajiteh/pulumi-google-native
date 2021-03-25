// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Spoke in a given project and location.
 */
export class GoogleLongrunningOperation extends pulumi.CustomResource {
    /**
     * Get an existing GoogleLongrunningOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GoogleLongrunningOperation {
        return new GoogleLongrunningOperation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:networkconnectivity/v1alpha1:GoogleLongrunningOperation';

    /**
     * Returns true if the given object is an instance of GoogleLongrunningOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GoogleLongrunningOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GoogleLongrunningOperation.__pulumiType;
    }


    /**
     * Create a GoogleLongrunningOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GoogleLongrunningOperationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["hub"] = args ? args.hub : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["linkedInterconnectAttachments"] = args ? args.linkedInterconnectAttachments : undefined;
            inputs["linkedRouterApplianceInstances"] = args ? args.linkedRouterApplianceInstances : undefined;
            inputs["linkedVpnTunnels"] = args ? args.linkedVpnTunnels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["spokeId"] = args ? args.spokeId : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["uniqueId"] = args ? args.uniqueId : undefined;
            inputs["updateTime"] = args ? args.updateTime : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GoogleLongrunningOperation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GoogleLongrunningOperation resource.
 */
export interface GoogleLongrunningOperationArgs {
    /**
     * The time when the Spoke was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Short description of the spoke resource
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The resource URL of the hub resource that the spoke is attached to
     */
    readonly hub?: pulumi.Input<string>;
    /**
     * User-defined labels.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The URIs of linked interconnect attachment resources
     */
    readonly linkedInterconnectAttachments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URIs of linked Router appliance resources
     */
    readonly linkedRouterApplianceInstances?: pulumi.Input<pulumi.Input<inputs.networkconnectivity.v1alpha1.RouterApplianceInstance>[]>;
    /**
     * The URIs of linked VPN tunnel resources
     */
    readonly linkedVpnTunnels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Immutable. The name of a Spoke resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Required. The parent's resource name of the Spoke.
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    readonly requestId?: pulumi.Input<string>;
    /**
     * Optional. Unique id for the Spoke to create.
     */
    readonly spokeId?: pulumi.Input<string>;
    /**
     * Output only. The current lifecycle state of this Hub.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Output only. Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
     */
    readonly uniqueId?: pulumi.Input<string>;
    /**
     * The time when the Spoke was updated.
     */
    readonly updateTime?: pulumi.Input<string>;
}