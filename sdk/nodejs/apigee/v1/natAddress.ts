// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a NAT address. The address is created in the RESERVED state and a static external IP address will be provisioned. At this time, the instance will not use this IP address for Internet egress traffic. The address can be activated for use once any required firewall IP whitelisting has been completed. **Note:** Not supported for Apigee hybrid.
 */
export class NatAddress extends pulumi.CustomResource {
    /**
     * Get an existing NatAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NatAddress {
        return new NatAddress(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:NatAddress';

    /**
     * Returns true if the given object is an instance of NatAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatAddress.__pulumiType;
    }

    /**
     * The static IPV4 address.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Required. Resource ID of the NAT address.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * State of the nat address.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a NatAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatAddressArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["ipAddress"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["ipAddress"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NatAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NatAddress resource.
 */
export interface NatAddressArgs {
    instanceId: pulumi.Input<string>;
    /**
     * Required. Resource ID of the NAT address.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}