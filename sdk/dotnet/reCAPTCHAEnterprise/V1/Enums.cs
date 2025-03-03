// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1
{
    /// <summary>
    /// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if CHALLENGE.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge : IEquatable<GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge>
    {
        private readonly string _value;

        private GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Perform the normal risk analysis and return either nocaptcha or a challenge depending on risk and trust factors.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge TestingChallengeUnspecified { get; } = new GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("TESTING_CHALLENGE_UNSPECIFIED");
        /// <summary>
        /// Challenge requests for this key always return a nocaptcha, which does not require a solution.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge Nocaptcha { get; } = new GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("NOCAPTCHA");
        /// <summary>
        /// Challenge requests for this key always return an unsolvable challenge.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge UnsolvableChallenge { get; } = new GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge("UNSOLVABLE_CHALLENGE");

        public static bool operator ==(GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge left, GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge left, GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge other && Equals(other);
        public bool Equals(GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The WAF feature for which this key is enabled.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature : IEquatable<GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature>
    {
        private readonly string _value;

        private GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Undefined feature.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature WafFeatureUnspecified { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature("WAF_FEATURE_UNSPECIFIED");
        /// <summary>
        /// Redirects suspicious traffic to reCAPTCHA.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature ChallengePage { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature("CHALLENGE_PAGE");
        /// <summary>
        /// Use reCAPTCHA session-tokens to protect the whole user session on the site's domain.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature SessionToken { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature("SESSION_TOKEN");
        /// <summary>
        /// Use reCAPTCHA action-tokens to protect user actions.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature ActionToken { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature("ACTION_TOKEN");
        /// <summary>
        /// Use reCAPTCHA WAF express protection to protect any content other than web pages, like APIs and IoT devices.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature Express { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature("EXPRESS");

        public static bool operator ==(GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature left, GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature left, GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature other && Equals(other);
        public bool Equals(GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The WAF service that uses this key.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecaptchaenterpriseV1WafSettingsWafService : IEquatable<GoogleCloudRecaptchaenterpriseV1WafSettingsWafService>
    {
        private readonly string _value;

        private GoogleCloudRecaptchaenterpriseV1WafSettingsWafService(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Undefined WAF
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafService WafServiceUnspecified { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafService("WAF_SERVICE_UNSPECIFIED");
        /// <summary>
        /// Cloud Armor
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafService Ca { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafService("CA");
        /// <summary>
        /// Fastly
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WafSettingsWafService Fastly { get; } = new GoogleCloudRecaptchaenterpriseV1WafSettingsWafService("FASTLY");

        public static bool operator ==(GoogleCloudRecaptchaenterpriseV1WafSettingsWafService left, GoogleCloudRecaptchaenterpriseV1WafSettingsWafService right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecaptchaenterpriseV1WafSettingsWafService left, GoogleCloudRecaptchaenterpriseV1WafSettingsWafService right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecaptchaenterpriseV1WafSettingsWafService value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecaptchaenterpriseV1WafSettingsWafService other && Equals(other);
        public bool Equals(GoogleCloudRecaptchaenterpriseV1WafSettingsWafService other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference : IEquatable<GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference>
    {
        private readonly string _value;

        private GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default type that indicates this enum hasn't been specified.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference ChallengeSecurityPreferenceUnspecified { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED");
        /// <summary>
        /// Key tends to show fewer and easier challenges.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference Usability { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("USABILITY");
        /// <summary>
        /// Key tends to show balanced (in amount and difficulty) challenges.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference Balance { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("BALANCE");
        /// <summary>
        /// Key tends to show more and harder challenges.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference Security { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference("SECURITY");

        public static bool operator ==(GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference left, GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference left, GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference other && Equals(other);
        public bool Equals(GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. Describes how this key is integrated with the website.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType : IEquatable<GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType>
    {
        private readonly string _value;

        private GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default type that indicates this enum hasn't been specified. This is not a valid IntegrationType, one of the other types must be specified instead.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType IntegrationTypeUnspecified { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("INTEGRATION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Only used to produce scores. It doesn't display the "I'm not a robot" checkbox and never shows captcha challenges.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType Score { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("SCORE");
        /// <summary>
        /// Displays the "I'm not a robot" checkbox and may show captcha challenges after it is checked.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType Checkbox { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("CHECKBOX");
        /// <summary>
        /// Doesn't display the "I'm not a robot" checkbox, but may show captcha challenges after risk analysis.
        /// </summary>
        public static GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType Invisible { get; } = new GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType("INVISIBLE");

        public static bool operator ==(GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType left, GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType left, GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType other && Equals(other);
        public bool Equals(GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
