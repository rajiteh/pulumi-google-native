// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ServiceLevelObjective by name.
func LookupServiceLevelObjective(ctx *pulumi.Context, args *LookupServiceLevelObjectiveArgs, opts ...pulumi.InvokeOption) (*LookupServiceLevelObjectiveResult, error) {
	var rv LookupServiceLevelObjectiveResult
	err := ctx.Invoke("google-native:monitoring/v3:getServiceLevelObjective", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceLevelObjectiveArgs struct {
	ServiceId               string  `pulumi:"serviceId"`
	ServiceLevelObjectiveId string  `pulumi:"serviceLevelObjectiveId"`
	V3Id                    string  `pulumi:"v3Id"`
	V3Id1                   string  `pulumi:"v3Id1"`
	View                    *string `pulumi:"view"`
}

type LookupServiceLevelObjectiveResult struct {
	// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
	CalendarPeriod string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
	Goal float64 `pulumi:"goal"`
	// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name string `pulumi:"name"`
	// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod string `pulumi:"rollingPeriod"`
	// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
	ServiceLevelIndicator ServiceLevelIndicatorResponse `pulumi:"serviceLevelIndicator"`
	// Labels which have been used to annotate the service-level objective. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
	UserLabels map[string]string `pulumi:"userLabels"`
}

func LookupServiceLevelObjectiveOutput(ctx *pulumi.Context, args LookupServiceLevelObjectiveOutputArgs, opts ...pulumi.InvokeOption) LookupServiceLevelObjectiveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceLevelObjectiveResult, error) {
			args := v.(LookupServiceLevelObjectiveArgs)
			r, err := LookupServiceLevelObjective(ctx, &args, opts...)
			var s LookupServiceLevelObjectiveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceLevelObjectiveResultOutput)
}

type LookupServiceLevelObjectiveOutputArgs struct {
	ServiceId               pulumi.StringInput    `pulumi:"serviceId"`
	ServiceLevelObjectiveId pulumi.StringInput    `pulumi:"serviceLevelObjectiveId"`
	V3Id                    pulumi.StringInput    `pulumi:"v3Id"`
	V3Id1                   pulumi.StringInput    `pulumi:"v3Id1"`
	View                    pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupServiceLevelObjectiveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLevelObjectiveArgs)(nil)).Elem()
}

type LookupServiceLevelObjectiveResultOutput struct{ *pulumi.OutputState }

func (LookupServiceLevelObjectiveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLevelObjectiveResult)(nil)).Elem()
}

func (o LookupServiceLevelObjectiveResultOutput) ToLookupServiceLevelObjectiveResultOutput() LookupServiceLevelObjectiveResultOutput {
	return o
}

func (o LookupServiceLevelObjectiveResultOutput) ToLookupServiceLevelObjectiveResultOutputWithContext(ctx context.Context) LookupServiceLevelObjectiveResultOutput {
	return o
}

// A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
func (o LookupServiceLevelObjectiveResultOutput) CalendarPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) string { return v.CalendarPeriod }).(pulumi.StringOutput)
}

// Name used for UI elements listing this SLO.
func (o LookupServiceLevelObjectiveResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
func (o LookupServiceLevelObjectiveResultOutput) Goal() pulumi.Float64Output {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) float64 { return v.Goal }).(pulumi.Float64Output)
}

// Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
func (o LookupServiceLevelObjectiveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) string { return v.Name }).(pulumi.StringOutput)
}

// A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
func (o LookupServiceLevelObjectiveResultOutput) RollingPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) string { return v.RollingPeriod }).(pulumi.StringOutput)
}

// The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
func (o LookupServiceLevelObjectiveResultOutput) ServiceLevelIndicator() ServiceLevelIndicatorResponseOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) ServiceLevelIndicatorResponse {
		return v.ServiceLevelIndicator
	}).(ServiceLevelIndicatorResponseOutput)
}

// Labels which have been used to annotate the service-level objective. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
func (o LookupServiceLevelObjectiveResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceLevelObjectiveResultOutput{})
}
