// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a Control.
 */
export function getControl(args: GetControlArgs, opts?: pulumi.InvokeOptions): Promise<GetControlResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:retail/v2:getControl", {
        "catalogId": args.catalogId,
        "controlId": args.controlId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetControlArgs {
    catalogId: string;
    controlId: string;
    location: string;
    project?: string;
}

export interface GetControlResult {
    /**
     * List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
     */
    readonly associatedServingConfigIds: string[];
    /**
     * The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
     */
    readonly displayName: string;
    /**
     * Immutable. Fully qualified name `projects/*&#47;locations/global/catalogs/*&#47;controls/*`
     */
    readonly name: string;
    /**
     * A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
     */
    readonly rule: outputs.retail.v2.GoogleCloudRetailV2RuleResponse;
    /**
     * Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
     */
    readonly searchSolutionUseCase: string[];
    /**
     * Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
     */
    readonly solutionTypes: string[];
}
/**
 * Gets a Control.
 */
export function getControlOutput(args: GetControlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetControlResult> {
    return pulumi.output(args).apply((a: any) => getControl(a, opts))
}

export interface GetControlOutputArgs {
    catalogId: pulumi.Input<string>;
    controlId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
