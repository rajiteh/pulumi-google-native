# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge',
    'GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature',
    'GoogleCloudRecaptchaenterpriseV1WafSettingsWafService',
    'GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference',
    'GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType',
]


class GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge(str, Enum):
    """
    For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if CHALLENGE.
    """
    TESTING_CHALLENGE_UNSPECIFIED = "TESTING_CHALLENGE_UNSPECIFIED"
    """
    Perform the normal risk analysis and return either nocaptcha or a challenge depending on risk and trust factors.
    """
    NOCAPTCHA = "NOCAPTCHA"
    """
    Challenge requests for this key always return a nocaptcha, which does not require a solution.
    """
    UNSOLVABLE_CHALLENGE = "UNSOLVABLE_CHALLENGE"
    """
    Challenge requests for this key always return an unsolvable challenge.
    """


class GoogleCloudRecaptchaenterpriseV1WafSettingsWafFeature(str, Enum):
    """
    Required. The WAF feature for which this key is enabled.
    """
    WAF_FEATURE_UNSPECIFIED = "WAF_FEATURE_UNSPECIFIED"
    """
    Undefined feature.
    """
    CHALLENGE_PAGE = "CHALLENGE_PAGE"
    """
    Redirects suspicious traffic to reCAPTCHA.
    """
    SESSION_TOKEN = "SESSION_TOKEN"
    """
    Use reCAPTCHA session-tokens to protect the whole user session on the site's domain.
    """
    ACTION_TOKEN = "ACTION_TOKEN"
    """
    Use reCAPTCHA action-tokens to protect user actions.
    """
    EXPRESS = "EXPRESS"
    """
    Use reCAPTCHA WAF express protection to protect any content other than web pages, like APIs and IoT devices.
    """


class GoogleCloudRecaptchaenterpriseV1WafSettingsWafService(str, Enum):
    """
    Required. The WAF service that uses this key.
    """
    WAF_SERVICE_UNSPECIFIED = "WAF_SERVICE_UNSPECIFIED"
    """
    Undefined WAF
    """
    CA = "CA"
    """
    Cloud Armor
    """
    FASTLY = "FASTLY"
    """
    Fastly
    """


class GoogleCloudRecaptchaenterpriseV1WebKeySettingsChallengeSecurityPreference(str, Enum):
    """
    Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE.
    """
    CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED = "CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED"
    """
    Default type that indicates this enum hasn't been specified.
    """
    USABILITY = "USABILITY"
    """
    Key tends to show fewer and easier challenges.
    """
    BALANCE = "BALANCE"
    """
    Key tends to show balanced (in amount and difficulty) challenges.
    """
    SECURITY = "SECURITY"
    """
    Key tends to show more and harder challenges.
    """


class GoogleCloudRecaptchaenterpriseV1WebKeySettingsIntegrationType(str, Enum):
    """
    Required. Describes how this key is integrated with the website.
    """
    INTEGRATION_TYPE_UNSPECIFIED = "INTEGRATION_TYPE_UNSPECIFIED"
    """
    Default type that indicates this enum hasn't been specified. This is not a valid IntegrationType, one of the other types must be specified instead.
    """
    SCORE = "SCORE"
    """
    Only used to produce scores. It doesn't display the "I'm not a robot" checkbox and never shows captcha challenges.
    """
    CHECKBOX = "CHECKBOX"
    """
    Displays the "I'm not a robot" checkbox and may show captcha challenges after it is checked.
    """
    INVISIBLE = "INVISIBLE"
    """
    Doesn't display the "I'm not a robot" checkbox, but may show captcha challenges after risk analysis.
    """
