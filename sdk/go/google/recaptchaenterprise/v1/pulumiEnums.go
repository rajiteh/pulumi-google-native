// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if CHALLENGE.
type GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge pulumi.String

const (
	// Perform the normal risk analysis and return either nocaptcha or a challenge depending on risk and trust factors.
	GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallengeTestingChallengeUnspecified = GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("TESTING_CHALLENGE_UNSPECIFIED")
	// Challenge requests for this key will always return a nocaptcha, which does not require a solution.
	GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallengeNocaptcha = GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("NOCAPTCHA")
	// Challenge requests for this key will always return an unsolvable challenge.
	GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallengeUnsolvableChallenge = GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("UNSOLVABLE_CHALLENGE")
)

func (GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE.
type GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference pulumi.String

const (
	// Default type that indicates this enum hasn't been specified.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreferenceChallengeSecurityPreferenceUnspecified = GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED")
	// Key tends to show fewer and easier challenges.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreferenceUsability = GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("USABILITY")
	// Key tends to show balanced (in amount and difficulty) challenges.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreferenceBalance = GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("BALANCE")
	// Key tends to show more and harder challenges.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreferenceSecurity = GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("SECURITY")
)

func (GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. Describes how this key is integrated with the website.
type GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType pulumi.String

const (
	// Default type that indicates this enum hasn't been specified. This is not a valid IntegrationType, one of the other types must be specified instead.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationTypeIntegrationTypeUnspecified = GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("INTEGRATION_TYPE_UNSPECIFIED")
	// Only used to produce scores. It doesn't display the "I'm not a robot" checkbox and never shows captcha challenges.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationTypeScore = GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("SCORE")
	// Displays the "I'm not a robot" checkbox and may show captcha challenges after it is checked.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationTypeCheckbox = GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("CHECKBOX")
	// Doesn't display the "I'm not a robot" checkbox, but may show captcha challenges after risk analysis.
	GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationTypeInvisible = GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("INVISIBLE")
)

func (GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}