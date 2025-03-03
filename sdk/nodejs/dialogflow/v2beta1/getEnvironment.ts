// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified agent environment.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dialogflow/v2beta1:getEnvironment", {
        "environmentId": args.environmentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEnvironmentArgs {
    environmentId: string;
    location: string;
    project?: string;
}

export interface GetEnvironmentResult {
    /**
     * Optional. The agent version loaded into this environment. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
     */
    readonly agentVersion: string;
    /**
     * Optional. The developer-provided description for this environment. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description: string;
    /**
     * Optional. The fulfillment settings to use for this environment.
     */
    readonly fulfillment: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1FulfillmentResponse;
    /**
     * The unique identifier of this agent environment. Supported formats: - `projects//agent/environments/` - `projects//locations//agent/environments/`
     */
    readonly name: string;
    /**
     * The state of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
     */
    readonly state: string;
    /**
     * Optional. Text to speech settings for this environment.
     */
    readonly textToSpeechSettings: outputs.dialogflow.v2beta1.GoogleCloudDialogflowV2beta1TextToSpeechSettingsResponse;
    /**
     * The last update time of this environment. This field is read-only, i.e., it cannot be set by create and update methods.
     */
    readonly updateTime: string;
}
/**
 * Retrieves the specified agent environment.
 */
export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply((a: any) => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    environmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
