// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves a conversation participant.
 */
export function getParticipant(args: GetParticipantArgs, opts?: pulumi.InvokeOptions): Promise<GetParticipantResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dialogflow/v2beta1:getParticipant", {
        "conversationId": args.conversationId,
        "location": args.location,
        "participantId": args.participantId,
        "project": args.project,
    }, opts);
}

export interface GetParticipantArgs {
    conversationId: string;
    location: string;
    participantId: string;
    project?: string;
}

export interface GetParticipantResult {
    /**
     * Optional. Key-value filters on the metadata of documents returned by article suggestion. If specified, article suggestion only returns suggested documents that match all filters in their Document.metadata. Multiple values for a metadata key should be concatenated by comma. For example, filters to match all documents that have 'US' or 'CA' in their market metadata values and 'agent' in their user metadata values will be ``` documents_metadata_filters { key: "market" value: "US,CA" } documents_metadata_filters { key: "user" value: "agent" } ```
     */
    readonly documentsMetadataFilters: {[key: string]: string};
    /**
     * Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
     */
    readonly name: string;
    /**
     * Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
     */
    readonly obfuscatedExternalUserId: string;
    /**
     * Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
     */
    readonly role: string;
}
/**
 * Retrieves a conversation participant.
 */
export function getParticipantOutput(args: GetParticipantOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetParticipantResult> {
    return pulumi.output(args).apply((a: any) => getParticipant(a, opts))
}

export interface GetParticipantOutputArgs {
    conversationId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    participantId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
