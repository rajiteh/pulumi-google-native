// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Response Policy Rule.
 */
export class ResponsePolicyRule extends pulumi.CustomResource {
    /**
     * Get an existing ResponsePolicyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResponsePolicyRule {
        return new ResponsePolicyRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:dns/v1beta2:ResponsePolicyRule';

    /**
     * Returns true if the given object is an instance of ResponsePolicyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResponsePolicyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResponsePolicyRule.__pulumiType;
    }


    /**
     * Create a ResponsePolicyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponsePolicyRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.responsePolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responsePolicy'");
            }
            inputs["behavior"] = args ? args.behavior : undefined;
            inputs["clientOperationId"] = args ? args.clientOperationId : undefined;
            inputs["dnsName"] = args ? args.dnsName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["localData"] = args ? args.localData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["responsePolicy"] = args ? args.responsePolicy : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResponsePolicyRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResponsePolicyRule resource.
 */
export interface ResponsePolicyRuleArgs {
    /**
     * Answer this query with a behavior rather than DNS data.
     */
    readonly behavior?: pulumi.Input<string>;
    /**
     * For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
     */
    readonly clientOperationId?: pulumi.Input<string>;
    /**
     * The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     */
    readonly dnsName?: pulumi.Input<string>;
    readonly kind?: pulumi.Input<string>;
    /**
     * Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     */
    readonly localData?: pulumi.Input<inputs.dns.v1beta2.ResponsePolicyRuleLocalData>;
    /**
     * Identifies the project addressed by this request.
     */
    readonly project: pulumi.Input<string>;
    /**
     * User assigned name of the Response Policy containing the Response Policy Rule.
     */
    readonly responsePolicy: pulumi.Input<string>;
    /**
     * An identifier for this rule. Must be unique with the ResponsePolicy.
     */
    readonly ruleName?: pulumi.Input<string>;
}