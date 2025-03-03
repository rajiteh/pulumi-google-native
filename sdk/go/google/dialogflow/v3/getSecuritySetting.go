// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified SecuritySettings. The returned settings may be stale by up to 1 minute.
func LookupSecuritySetting(ctx *pulumi.Context, args *LookupSecuritySettingArgs, opts ...pulumi.InvokeOption) (*LookupSecuritySettingResult, error) {
	var rv LookupSecuritySettingResult
	err := ctx.Invoke("google-native:dialogflow/v3:getSecuritySetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecuritySettingArgs struct {
	Location          string  `pulumi:"location"`
	Project           *string `pulumi:"project"`
	SecuritySettingId string  `pulumi:"securitySettingId"`
}

type LookupSecuritySettingResult struct {
	// Controls audio export settings for post-conversation analytics when ingesting audio to conversations via Participants.AnalyzeContent or Participants.StreamingAnalyzeContent. If retention_strategy is set to REMOVE_AFTER_CONVERSATION or audio_export_settings.gcs_bucket is empty, audio export is disabled. If audio export is enabled, audio is recorded and saved to audio_export_settings.gcs_bucket, subject to retention policy of audio_export_settings.gcs_bucket. This setting won't effect audio input for implicit sessions via Sessions.DetectIntent or Sessions.StreamingDetectIntent.
	AudioExportSettings GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponse `pulumi:"audioExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
	DeidentifyTemplate string `pulumi:"deidentifyTemplate"`
	// The human-readable name of the security settings, unique within the location.
	DisplayName string `pulumi:"displayName"`
	// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
	InsightsExportSettings GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponse `pulumi:"insightsExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
	InspectTemplate string `pulumi:"inspectTemplate"`
	// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
	Name string `pulumi:"name"`
	// List of types of data to remove when retention settings triggers purge.
	PurgeDataTypes []string `pulumi:"purgeDataTypes"`
	// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
	RedactionScope string `pulumi:"redactionScope"`
	// Strategy that defines how we do redaction.
	RedactionStrategy string `pulumi:"redactionStrategy"`
	// Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
	RetentionWindowDays int `pulumi:"retentionWindowDays"`
}

func LookupSecuritySettingOutput(ctx *pulumi.Context, args LookupSecuritySettingOutputArgs, opts ...pulumi.InvokeOption) LookupSecuritySettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecuritySettingResult, error) {
			args := v.(LookupSecuritySettingArgs)
			r, err := LookupSecuritySetting(ctx, &args, opts...)
			var s LookupSecuritySettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecuritySettingResultOutput)
}

type LookupSecuritySettingOutputArgs struct {
	Location          pulumi.StringInput    `pulumi:"location"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
	SecuritySettingId pulumi.StringInput    `pulumi:"securitySettingId"`
}

func (LookupSecuritySettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecuritySettingArgs)(nil)).Elem()
}

type LookupSecuritySettingResultOutput struct{ *pulumi.OutputState }

func (LookupSecuritySettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecuritySettingResult)(nil)).Elem()
}

func (o LookupSecuritySettingResultOutput) ToLookupSecuritySettingResultOutput() LookupSecuritySettingResultOutput {
	return o
}

func (o LookupSecuritySettingResultOutput) ToLookupSecuritySettingResultOutputWithContext(ctx context.Context) LookupSecuritySettingResultOutput {
	return o
}

// Controls audio export settings for post-conversation analytics when ingesting audio to conversations via Participants.AnalyzeContent or Participants.StreamingAnalyzeContent. If retention_strategy is set to REMOVE_AFTER_CONVERSATION or audio_export_settings.gcs_bucket is empty, audio export is disabled. If audio export is enabled, audio is recorded and saved to audio_export_settings.gcs_bucket, subject to retention policy of audio_export_settings.gcs_bucket. This setting won't effect audio input for implicit sessions via Sessions.DetectIntent or Sessions.StreamingDetectIntent.
func (o LookupSecuritySettingResultOutput) AudioExportSettings() GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponseOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponse {
		return v.AudioExportSettings
	}).(GoogleCloudDialogflowCxV3SecuritySettingsAudioExportSettingsResponseOutput)
}

// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
func (o LookupSecuritySettingResultOutput) DeidentifyTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.DeidentifyTemplate }).(pulumi.StringOutput)
}

// The human-readable name of the security settings, unique within the location.
func (o LookupSecuritySettingResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
func (o LookupSecuritySettingResultOutput) InsightsExportSettings() GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponseOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponse {
		return v.InsightsExportSettings
	}).(GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponseOutput)
}

// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
func (o LookupSecuritySettingResultOutput) InspectTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.InspectTemplate }).(pulumi.StringOutput)
}

// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
func (o LookupSecuritySettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of types of data to remove when retention settings triggers purge.
func (o LookupSecuritySettingResultOutput) PurgeDataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) []string { return v.PurgeDataTypes }).(pulumi.StringArrayOutput)
}

// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
func (o LookupSecuritySettingResultOutput) RedactionScope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.RedactionScope }).(pulumi.StringOutput)
}

// Strategy that defines how we do redaction.
func (o LookupSecuritySettingResultOutput) RedactionStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) string { return v.RedactionStrategy }).(pulumi.StringOutput)
}

// Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
func (o LookupSecuritySettingResultOutput) RetentionWindowDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecuritySettingResult) int { return v.RetentionWindowDays }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecuritySettingResultOutput{})
}
