// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Create security settings in the specified location.
 */
export class SecuritySetting extends pulumi.CustomResource {
    /**
     * Get an existing SecuritySetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecuritySetting {
        return new SecuritySetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:SecuritySetting';

    /**
     * Returns true if the given object is an instance of SecuritySetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecuritySetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecuritySetting.__pulumiType;
    }

    /**
     * Controls audio export settings for post-conversation analytics when ingesting audio to conversations via Participants.AnalyzeContent or Participants.StreamingAnalyzeContent. If retention_strategy is set to REMOVE_AFTER_CONVERSATION or audio_export_settings.gcs_bucket is empty, audio export is disabled. If audio export is enabled, audio is recorded and saved to audio_export_settings.gcs_bucket, subject to retention policy of audio_export_settings.gcs_bucket. This setting won't effect audio input for implicit sessions via Sessions.DetectIntent or Sessions.StreamingDetectIntent.
     */
    public readonly audioExportSettings!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponse>;
    /**
     * [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
     */
    public readonly deidentifyTemplate!: pulumi.Output<string>;
    /**
     * The human-readable name of the security settings, unique within the location.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
     */
    public readonly insightsExportSettings!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponse>;
    /**
     * [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
     */
    public readonly inspectTemplate!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * List of types of data to remove when retention settings triggers purge.
     */
    public readonly purgeDataTypes!: pulumi.Output<string[]>;
    /**
     * Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
     */
    public readonly redactionScope!: pulumi.Output<string>;
    /**
     * Strategy that defines how we do redaction.
     */
    public readonly redactionStrategy!: pulumi.Output<string>;
    /**
     * Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
     */
    public readonly retentionWindowDays!: pulumi.Output<number>;

    /**
     * Create a SecuritySetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecuritySettingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["audioExportSettings"] = args ? args.audioExportSettings : undefined;
            resourceInputs["deidentifyTemplate"] = args ? args.deidentifyTemplate : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["insightsExportSettings"] = args ? args.insightsExportSettings : undefined;
            resourceInputs["inspectTemplate"] = args ? args.inspectTemplate : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["purgeDataTypes"] = args ? args.purgeDataTypes : undefined;
            resourceInputs["redactionScope"] = args ? args.redactionScope : undefined;
            resourceInputs["redactionStrategy"] = args ? args.redactionStrategy : undefined;
            resourceInputs["retentionWindowDays"] = args ? args.retentionWindowDays : undefined;
        } else {
            resourceInputs["audioExportSettings"] = undefined /*out*/;
            resourceInputs["deidentifyTemplate"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["insightsExportSettings"] = undefined /*out*/;
            resourceInputs["inspectTemplate"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["purgeDataTypes"] = undefined /*out*/;
            resourceInputs["redactionScope"] = undefined /*out*/;
            resourceInputs["redactionStrategy"] = undefined /*out*/;
            resourceInputs["retentionWindowDays"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SecuritySetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecuritySetting resource.
 */
export interface SecuritySettingArgs {
    /**
     * Controls audio export settings for post-conversation analytics when ingesting audio to conversations via Participants.AnalyzeContent or Participants.StreamingAnalyzeContent. If retention_strategy is set to REMOVE_AFTER_CONVERSATION or audio_export_settings.gcs_bucket is empty, audio export is disabled. If audio export is enabled, audio is recorded and saved to audio_export_settings.gcs_bucket, subject to retention policy of audio_export_settings.gcs_bucket. This setting won't effect audio input for implicit sessions via Sessions.DetectIntent or Sessions.StreamingDetectIntent.
     */
    audioExportSettings?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsArgs>;
    /**
     * [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
     */
    deidentifyTemplate?: pulumi.Input<string>;
    /**
     * The human-readable name of the security settings, unique within the location.
     */
    displayName: pulumi.Input<string>;
    /**
     * Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
     */
    insightsExportSettings?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsArgs>;
    /**
     * [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
     */
    inspectTemplate?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * List of types of data to remove when retention settings triggers purge.
     */
    purgeDataTypes?: pulumi.Input<pulumi.Input<enums.dialogflow.v3.SecuritySettingPurgeDataTypesItem>[]>;
    /**
     * Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
     */
    redactionScope?: pulumi.Input<enums.dialogflow.v3.SecuritySettingRedactionScope>;
    /**
     * Strategy that defines how we do redaction.
     */
    redactionStrategy?: pulumi.Input<enums.dialogflow.v3.SecuritySettingRedactionStrategy>;
    /**
     * Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
     */
    retentionWindowDays?: pulumi.Input<number>;
}
