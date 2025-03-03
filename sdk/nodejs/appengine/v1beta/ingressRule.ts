// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a firewall rule for the application.
 * Auto-naming is currently not supported for this resource.
 */
export class IngressRule extends pulumi.CustomResource {
    /**
     * Get an existing IngressRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IngressRule {
        return new IngressRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:appengine/v1beta:IngressRule';

    /**
     * Returns true if the given object is an instance of IngressRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IngressRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IngressRule.__pulumiType;
    }

    /**
     * The action to take on matched requests.
     */
    public readonly action!: pulumi.Output<string>;
    public readonly appId!: pulumi.Output<string>;
    /**
     * An optional string description of this rule. This field has a maximum length of 400 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
     */
    public readonly sourceRange!: pulumi.Output<string>;

    /**
     * Create a IngressRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IngressRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["sourceRange"] = args ? args.sourceRange : undefined;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["sourceRange"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["appId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(IngressRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IngressRule resource.
 */
export interface IngressRuleArgs {
    /**
     * The action to take on matched requests.
     */
    action?: pulumi.Input<enums.appengine.v1beta.IngressRuleAction>;
    appId: pulumi.Input<string>;
    /**
     * An optional string description of this rule. This field has a maximum length of 400 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
     */
    priority?: pulumi.Input<number>;
    /**
     * IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
     */
    sourceRange?: pulumi.Input<string>;
}
