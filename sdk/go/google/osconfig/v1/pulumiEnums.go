// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// By changing the type to DIST, the patching is performed using `apt-get dist-upgrade` instead.
type AptSettingsType pulumi.String

const (
	// By default, upgrade will be performed.
	AptSettingsTypeTypeUnspecified = AptSettingsType("TYPE_UNSPECIFIED")
	// Runs `apt-get dist-upgrade`.
	AptSettingsTypeDist = AptSettingsType("DIST")
	// Runs `apt-get upgrade`.
	AptSettingsTypeUpgrade = AptSettingsType("UPGRADE")
)

func (AptSettingsType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AptSettingsType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AptSettingsType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AptSettingsType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AptSettingsType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The script interpreter to use to run the script. If no interpreter is specified the script will be executed directly, which will likely only succeed for scripts with [shebang lines] (https://en.wikipedia.org/wiki/Shebang_\(Unix\)).
type ExecStepConfigInterpreter pulumi.String

const (
	// Invalid for a Windows ExecStepConfig. For a Linux ExecStepConfig, the interpreter will be parsed from the shebang line of the script if unspecified.
	ExecStepConfigInterpreterInterpreterUnspecified = ExecStepConfigInterpreter("INTERPRETER_UNSPECIFIED")
	// Indicates that the script is run with `/bin/sh` on Linux and `cmd` on Windows.
	ExecStepConfigInterpreterShell = ExecStepConfigInterpreter("SHELL")
	// Indicates that the file is run with PowerShell flags `-NonInteractive`, `-NoProfile`, and `-ExecutionPolicy Bypass`.
	ExecStepConfigInterpreterPowershell = ExecStepConfigInterpreter("POWERSHELL")
)

func (ExecStepConfigInterpreter) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ExecStepConfigInterpreter) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExecStepConfigInterpreter) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExecStepConfigInterpreter) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExecStepConfigInterpreter) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Post-patch reboot settings.
type PatchConfigRebootConfig pulumi.String

const (
	// The default behavior is DEFAULT.
	PatchConfigRebootConfigRebootConfigUnspecified = PatchConfigRebootConfig("REBOOT_CONFIG_UNSPECIFIED")
	// The agent decides if a reboot is necessary by checking signals such as registry keys on Windows or `/var/run/reboot-required` on APT based systems. On RPM based systems, a set of core system package install times are compared with system boot time.
	PatchConfigRebootConfigDefault = PatchConfigRebootConfig("DEFAULT")
	// Always reboot the machine after the update completes.
	PatchConfigRebootConfigAlways = PatchConfigRebootConfig("ALWAYS")
	// Never reboot the machine after the update completes.
	PatchConfigRebootConfigNever = PatchConfigRebootConfig("NEVER")
)

func (PatchConfigRebootConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PatchConfigRebootConfig) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchConfigRebootConfig) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchConfigRebootConfig) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PatchConfigRebootConfig) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Mode of the patch rollout.
type PatchRolloutMode pulumi.String

const (
	// Mode must be specified.
	PatchRolloutModeModeUnspecified = PatchRolloutMode("MODE_UNSPECIFIED")
	// Patches are applied one zone at a time. The patch job begins in the region with the lowest number of targeted VMs. Within the region, patching begins in the zone with the lowest number of targeted VMs. If multiple regions (or zones within a region) have the same number of targeted VMs, a tie-breaker is achieved by sorting the regions or zones in alphabetical order.
	PatchRolloutModeZoneByZone = PatchRolloutMode("ZONE_BY_ZONE")
	// Patches are applied to VMs in all zones at the same time.
	PatchRolloutModeConcurrentZones = PatchRolloutMode("CONCURRENT_ZONES")
)

func (PatchRolloutMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e PatchRolloutMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchRolloutMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PatchRolloutMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PatchRolloutMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. The frequency unit of this recurring schedule.
type RecurringScheduleFrequency pulumi.String

const (
	// Invalid. A frequency must be specified.
	RecurringScheduleFrequencyFrequencyUnspecified = RecurringScheduleFrequency("FREQUENCY_UNSPECIFIED")
	// Indicates that the frequency should be expressed in terms of weeks.
	RecurringScheduleFrequencyWeekly = RecurringScheduleFrequency("WEEKLY")
	// Indicates that the frequency should be expressed in terms of months.
	RecurringScheduleFrequencyMonthly = RecurringScheduleFrequency("MONTHLY")
)

func (RecurringScheduleFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RecurringScheduleFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurringScheduleFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurringScheduleFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurringScheduleFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. A day of the week.
type WeekDayOfMonthDayOfWeek pulumi.String

const (
	// The day of the week is unspecified.
	WeekDayOfMonthDayOfWeekDayOfWeekUnspecified = WeekDayOfMonthDayOfWeek("DAY_OF_WEEK_UNSPECIFIED")
	// Monday
	WeekDayOfMonthDayOfWeekMonday = WeekDayOfMonthDayOfWeek("MONDAY")
	// Tuesday
	WeekDayOfMonthDayOfWeekTuesday = WeekDayOfMonthDayOfWeek("TUESDAY")
	// Wednesday
	WeekDayOfMonthDayOfWeekWednesday = WeekDayOfMonthDayOfWeek("WEDNESDAY")
	// Thursday
	WeekDayOfMonthDayOfWeekThursday = WeekDayOfMonthDayOfWeek("THURSDAY")
	// Friday
	WeekDayOfMonthDayOfWeekFriday = WeekDayOfMonthDayOfWeek("FRIDAY")
	// Saturday
	WeekDayOfMonthDayOfWeekSaturday = WeekDayOfMonthDayOfWeek("SATURDAY")
	// Sunday
	WeekDayOfMonthDayOfWeekSunday = WeekDayOfMonthDayOfWeek("SUNDAY")
)

func (WeekDayOfMonthDayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WeekDayOfMonthDayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekDayOfMonthDayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekDayOfMonthDayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeekDayOfMonthDayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. Day of the week.
type WeeklyScheduleDayOfWeek pulumi.String

const (
	// The day of the week is unspecified.
	WeeklyScheduleDayOfWeekDayOfWeekUnspecified = WeeklyScheduleDayOfWeek("DAY_OF_WEEK_UNSPECIFIED")
	// Monday
	WeeklyScheduleDayOfWeekMonday = WeeklyScheduleDayOfWeek("MONDAY")
	// Tuesday
	WeeklyScheduleDayOfWeekTuesday = WeeklyScheduleDayOfWeek("TUESDAY")
	// Wednesday
	WeeklyScheduleDayOfWeekWednesday = WeeklyScheduleDayOfWeek("WEDNESDAY")
	// Thursday
	WeeklyScheduleDayOfWeekThursday = WeeklyScheduleDayOfWeek("THURSDAY")
	// Friday
	WeeklyScheduleDayOfWeekFriday = WeeklyScheduleDayOfWeek("FRIDAY")
	// Saturday
	WeeklyScheduleDayOfWeekSaturday = WeeklyScheduleDayOfWeek("SATURDAY")
	// Sunday
	WeeklyScheduleDayOfWeekSunday = WeeklyScheduleDayOfWeek("SUNDAY")
)

func (WeeklyScheduleDayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WeeklyScheduleDayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeeklyScheduleDayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeeklyScheduleDayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeeklyScheduleDayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WindowsUpdateSettingsClassificationsItem pulumi.String

const (
	// Invalid. If classifications are included, they must be specified.
	WindowsUpdateSettingsClassificationsItemClassificationUnspecified = WindowsUpdateSettingsClassificationsItem("CLASSIFICATION_UNSPECIFIED")
	// "A widely released fix for a specific problem that addresses a critical, non-security-related bug." [1]
	WindowsUpdateSettingsClassificationsItemCritical = WindowsUpdateSettingsClassificationsItem("CRITICAL")
	// "A widely released fix for a product-specific, security-related vulnerability. Security vulnerabilities are rated by their severity. The severity rating is indicated in the Microsoft security bulletin as critical, important, moderate, or low." [1]
	WindowsUpdateSettingsClassificationsItemSecurity = WindowsUpdateSettingsClassificationsItem("SECURITY")
	// "A widely released and frequent software update that contains additions to a product's definition database. Definition databases are often used to detect objects that have specific attributes, such as malicious code, phishing websites, or junk mail." [1]
	WindowsUpdateSettingsClassificationsItemDefinition = WindowsUpdateSettingsClassificationsItem("DEFINITION")
	// "Software that controls the input and output of a device." [1]
	WindowsUpdateSettingsClassificationsItemDriver = WindowsUpdateSettingsClassificationsItem("DRIVER")
	// "New product functionality that is first distributed outside the context of a product release and that is typically included in the next full product release." [1]
	WindowsUpdateSettingsClassificationsItemFeaturePack = WindowsUpdateSettingsClassificationsItem("FEATURE_PACK")
	// "A tested, cumulative set of all hotfixes, security updates, critical updates, and updates. Additionally, service packs may contain additional fixes for problems that are found internally since the release of the product. Service packs my also contain a limited number of customer-requested design changes or features." [1]
	WindowsUpdateSettingsClassificationsItemServicePack = WindowsUpdateSettingsClassificationsItem("SERVICE_PACK")
	// "A utility or feature that helps complete a task or set of tasks." [1]
	WindowsUpdateSettingsClassificationsItemTool = WindowsUpdateSettingsClassificationsItem("TOOL")
	// "A tested, cumulative set of hotfixes, security updates, critical updates, and updates that are packaged together for easy deployment. A rollup generally targets a specific area, such as security, or a component of a product, such as Internet Information Services (IIS)." [1]
	WindowsUpdateSettingsClassificationsItemUpdateRollup = WindowsUpdateSettingsClassificationsItem("UPDATE_ROLLUP")
	// "A widely released fix for a specific problem. An update addresses a noncritical, non-security-related bug." [1]
	WindowsUpdateSettingsClassificationsItemUpdate = WindowsUpdateSettingsClassificationsItem("UPDATE")
)

func (WindowsUpdateSettingsClassificationsItem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e WindowsUpdateSettingsClassificationsItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WindowsUpdateSettingsClassificationsItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WindowsUpdateSettingsClassificationsItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WindowsUpdateSettingsClassificationsItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// WindowsUpdateSettingsClassificationsItemArrayInput is an input type that accepts WindowsUpdateSettingsClassificationsItemArray and WindowsUpdateSettingsClassificationsItemArrayOutput values.
// You can construct a concrete instance of `WindowsUpdateSettingsClassificationsItemArrayInput` via:
//
//          WindowsUpdateSettingsClassificationsItemArray{ WindowsUpdateSettingsClassificationsItemArgs{...} }
type WindowsUpdateSettingsClassificationsItemArrayInput interface {
	pulumi.Input

	ToWindowsUpdateSettingsClassificationsItemArrayOutput() WindowsUpdateSettingsClassificationsItemArrayOutput
	ToWindowsUpdateSettingsClassificationsItemArrayOutputWithContext(context.Context) WindowsUpdateSettingsClassificationsItemArrayOutput
}

type WindowsUpdateSettingsClassificationsItemArray []WindowsUpdateSettingsClassificationsItem

func (WindowsUpdateSettingsClassificationsItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsUpdateSettingsClassificationsItem)(nil)).Elem()
}

func (i WindowsUpdateSettingsClassificationsItemArray) ToWindowsUpdateSettingsClassificationsItemArrayOutput() WindowsUpdateSettingsClassificationsItemArrayOutput {
	return i.ToWindowsUpdateSettingsClassificationsItemArrayOutputWithContext(context.Background())
}

func (i WindowsUpdateSettingsClassificationsItemArray) ToWindowsUpdateSettingsClassificationsItemArrayOutputWithContext(ctx context.Context) WindowsUpdateSettingsClassificationsItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsUpdateSettingsClassificationsItemArrayOutput)
}

type WindowsUpdateSettingsClassificationsItemArrayOutput struct{ *pulumi.OutputState }

func (WindowsUpdateSettingsClassificationsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsUpdateSettingsClassificationsItem)(nil)).Elem()
}

func (o WindowsUpdateSettingsClassificationsItemArrayOutput) ToWindowsUpdateSettingsClassificationsItemArrayOutput() WindowsUpdateSettingsClassificationsItemArrayOutput {
	return o
}

func (o WindowsUpdateSettingsClassificationsItemArrayOutput) ToWindowsUpdateSettingsClassificationsItemArrayOutputWithContext(ctx context.Context) WindowsUpdateSettingsClassificationsItemArrayOutput {
	return o
}

func (o WindowsUpdateSettingsClassificationsItemArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]WindowsUpdateSettingsClassificationsItem)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}