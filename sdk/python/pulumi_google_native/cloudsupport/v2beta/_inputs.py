# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'ActorArgs',
    'CaseClassificationArgs',
]

@pulumi.input_type
class ActorArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None):
        """
        An object containing information about the effective user and authenticated principal responsible for an action.
        :param pulumi.Input[str] display_name: The name to display for the actor. If not provided, it is inferred from credentials supplied during case creation. When an email is provided, a display name must also be provided. This will be obfuscated if the user is a Google Support agent.
        :param pulumi.Input[str] email: The email address of the actor. If not provided, it is inferred from credentials supplied during case creation. If the authenticated principal does not have an email address, one must be provided. When a name is provided, an email must also be provided. This will be obfuscated if the user is a Google Support agent.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email is not None:
            pulumi.set(__self__, "email", email)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to display for the actor. If not provided, it is inferred from credentials supplied during case creation. When an email is provided, a display name must also be provided. This will be obfuscated if the user is a Google Support agent.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address of the actor. If not provided, it is inferred from credentials supplied during case creation. If the authenticated principal does not have an email address, one must be provided. When a name is provided, an email must also be provided. This will be obfuscated if the user is a Google Support agent.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)


@pulumi.input_type
class CaseClassificationArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None):
        """
        A classification object with a product type and value.
        :param pulumi.Input[str] display_name: The display name of the classification.
        :param pulumi.Input[str] id: The unique ID for a classification. Must be specified for case creation. To retrieve valid classification IDs for case creation, use `caseClassifications.search`.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the classification.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for a classification. Must be specified for case creation. To retrieve valid classification IDs for case creation, use `caseClassifications.search`.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


