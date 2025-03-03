# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetCaseResult',
    'AwaitableGetCaseResult',
    'get_case',
    'get_case_output',
]

@pulumi.output_type
class GetCaseResult:
    def __init__(__self__, classification=None, contact_email=None, create_time=None, creator=None, description=None, display_name=None, escalated=None, language_code=None, name=None, priority=None, severity=None, state=None, subscriber_email_addresses=None, test_case=None, time_zone=None, update_time=None):
        if classification and not isinstance(classification, dict):
            raise TypeError("Expected argument 'classification' to be a dict")
        pulumi.set(__self__, "classification", classification)
        if contact_email and not isinstance(contact_email, str):
            raise TypeError("Expected argument 'contact_email' to be a str")
        pulumi.set(__self__, "contact_email", contact_email)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if creator and not isinstance(creator, dict):
            raise TypeError("Expected argument 'creator' to be a dict")
        pulumi.set(__self__, "creator", creator)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if escalated and not isinstance(escalated, bool):
            raise TypeError("Expected argument 'escalated' to be a bool")
        pulumi.set(__self__, "escalated", escalated)
        if language_code and not isinstance(language_code, str):
            raise TypeError("Expected argument 'language_code' to be a str")
        pulumi.set(__self__, "language_code", language_code)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if subscriber_email_addresses and not isinstance(subscriber_email_addresses, list):
            raise TypeError("Expected argument 'subscriber_email_addresses' to be a list")
        pulumi.set(__self__, "subscriber_email_addresses", subscriber_email_addresses)
        if test_case and not isinstance(test_case, bool):
            raise TypeError("Expected argument 'test_case' to be a bool")
        pulumi.set(__self__, "test_case", test_case)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def classification(self) -> 'outputs.CaseClassificationResponse':
        """
        The issue classification applicable to this case.
        """
        return pulumi.get(self, "classification")

    @property
    @pulumi.getter(name="contactEmail")
    def contact_email(self) -> str:
        """
        A user-supplied email address to send case update notifications for. This should only be used in BYOID flows, where we cannot infer the user's email address directly from their EUCs.
        """
        return pulumi.get(self, "contact_email")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time this case was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def creator(self) -> 'outputs.ActorResponse':
        """
        The user who created the case. Note: The name and email will be obfuscated if the case was created by Google Support.
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A broad description of the issue.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The short summary of the issue reported in this case.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def escalated(self) -> bool:
        """
        Whether the case is currently escalated.
        """
        return pulumi.get(self, "escalated")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> str:
        """
        The language the user has requested to receive support in. This should be a BCP 47 language code (e.g., `"en"`, `"zh-CN"`, `"zh-TW"`, `"ja"`, `"ko"`). If no language or an unsupported language is specified, this field defaults to English (en). Language selection during case creation may affect your available support options. For a list of supported languages and their support working hours, see: https://cloud.google.com/support/docs/language-working-hours
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the case.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> str:
        """
        The priority of this case.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def severity(self) -> str:
        """
        REMOVED. The severity of this case. Use priority instead.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current status of the support case.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subscriberEmailAddresses")
    def subscriber_email_addresses(self) -> Sequence[str]:
        """
        The email addresses to receive updates on this case.
        """
        return pulumi.get(self, "subscriber_email_addresses")

    @property
    @pulumi.getter(name="testCase")
    def test_case(self) -> bool:
        """
        Whether this case was created for internal API testing and should not be acted on by the support team.
        """
        return pulumi.get(self, "test_case")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        """
        The timezone of the user who created the support case. It should be in a format IANA recognizes: https://www.iana.org/time-zones. There is no additional validation done by the API.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time this case was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetCaseResult(GetCaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCaseResult(
            classification=self.classification,
            contact_email=self.contact_email,
            create_time=self.create_time,
            creator=self.creator,
            description=self.description,
            display_name=self.display_name,
            escalated=self.escalated,
            language_code=self.language_code,
            name=self.name,
            priority=self.priority,
            severity=self.severity,
            state=self.state,
            subscriber_email_addresses=self.subscriber_email_addresses,
            test_case=self.test_case,
            time_zone=self.time_zone,
            update_time=self.update_time)


def get_case(case_id: Optional[str] = None,
             v2beta_id1: Optional[str] = None,
             v2betum_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCaseResult:
    """
    Retrieve the specified case.
    """
    __args__ = dict()
    __args__['caseId'] = case_id
    __args__['v2betaId1'] = v2beta_id1
    __args__['v2betumId'] = v2betum_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudsupport/v2beta:getCase', __args__, opts=opts, typ=GetCaseResult).value

    return AwaitableGetCaseResult(
        classification=__ret__.classification,
        contact_email=__ret__.contact_email,
        create_time=__ret__.create_time,
        creator=__ret__.creator,
        description=__ret__.description,
        display_name=__ret__.display_name,
        escalated=__ret__.escalated,
        language_code=__ret__.language_code,
        name=__ret__.name,
        priority=__ret__.priority,
        severity=__ret__.severity,
        state=__ret__.state,
        subscriber_email_addresses=__ret__.subscriber_email_addresses,
        test_case=__ret__.test_case,
        time_zone=__ret__.time_zone,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_case)
def get_case_output(case_id: Optional[pulumi.Input[str]] = None,
                    v2beta_id1: Optional[pulumi.Input[str]] = None,
                    v2betum_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCaseResult]:
    """
    Retrieve the specified case.
    """
    ...
