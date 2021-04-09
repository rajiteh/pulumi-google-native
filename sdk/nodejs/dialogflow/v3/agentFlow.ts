// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a flow in the specified agent.
 */
export class AgentFlow extends pulumi.CustomResource {
    /**
     * Get an existing AgentFlow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AgentFlow {
        return new AgentFlow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp-native:dialogflow/v3:AgentFlow';

    /**
     * Returns true if the given object is an instance of AgentFlow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AgentFlow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AgentFlow.__pulumiType;
    }

    /**
     * The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Required. The human-readable name of the flow.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     */
    public readonly eventHandlers!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3EventHandlerResponse[]>;
    /**
     * The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * NLU related settings of the flow.
     */
    public readonly nluSettings!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3NluSettingsResponse>;
    /**
     * A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    public readonly transitionRouteGroups!: pulumi.Output<string[]>;
    /**
     * A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
     */
    public readonly transitionRoutes!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRouteResponse[]>;

    /**
     * Create a AgentFlow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentFlowArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentsId'");
            }
            if ((!args || args.flowsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowsId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["agentsId"] = args ? args.agentsId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["eventHandlers"] = args ? args.eventHandlers : undefined;
            inputs["flowsId"] = args ? args.flowsId : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nluSettings"] = args ? args.nluSettings : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["transitionRouteGroups"] = args ? args.transitionRouteGroups : undefined;
            inputs["transitionRoutes"] = args ? args.transitionRoutes : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["eventHandlers"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nluSettings"] = undefined /*out*/;
            inputs["transitionRouteGroups"] = undefined /*out*/;
            inputs["transitionRoutes"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AgentFlow.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AgentFlow resource.
 */
export interface AgentFlowArgs {
    readonly agentsId: pulumi.Input<string>;
    /**
     * The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Required. The human-readable name of the flow.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
     */
    readonly eventHandlers?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3EventHandler>[]>;
    readonly flowsId: pulumi.Input<string>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * NLU related settings of the flow.
     */
    readonly nluSettings?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3NluSettings>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
     */
    readonly transitionRouteGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
     */
    readonly transitionRoutes?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3TransitionRoute>[]>;
}