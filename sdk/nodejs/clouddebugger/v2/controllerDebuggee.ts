// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Registers the debuggee with the controller service. All agents attached to the same application must call this method with exactly the same request content to get back the same stable `debuggee_id`. Agents should call this method again whenever `google.rpc.Code.NOT_FOUND` is returned from any controller method. This protocol allows the controller service to disable debuggees, recover from data loss, or change the `debuggee_id` format. Agents must handle `debuggee_id` value changing upon re-registration.
 */
export class ControllerDebuggee extends pulumi.CustomResource {
    /**
     * Get an existing ControllerDebuggee resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ControllerDebuggee {
        return new ControllerDebuggee(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp-native:clouddebugger/v2:ControllerDebuggee';

    /**
     * Returns true if the given object is an instance of ControllerDebuggee.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ControllerDebuggee {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ControllerDebuggee.__pulumiType;
    }


    /**
     * Create a ControllerDebuggee resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ControllerDebuggeeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["debuggee"] = args ? args.debuggee : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ControllerDebuggee.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ControllerDebuggee resource.
 */
export interface ControllerDebuggeeArgs {
    /**
     * Required. Debuggee information to register. The fields `project`, `uniquifier`, `description` and `agent_version` of the debuggee must be set.
     */
    readonly debuggee?: pulumi.Input<inputs.clouddebugger.v2.Debuggee>;
}