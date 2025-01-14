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
    'GetInboundSsoAssignmentResult',
    'AwaitableGetInboundSsoAssignmentResult',
    'get_inbound_sso_assignment',
    'get_inbound_sso_assignment_output',
]

@pulumi.output_type
class GetInboundSsoAssignmentResult:
    def __init__(__self__, customer=None, name=None, rank=None, saml_sso_info=None, sign_in_behavior=None, sso_mode=None, target_group=None, target_org_unit=None):
        if customer and not isinstance(customer, str):
            raise TypeError("Expected argument 'customer' to be a str")
        pulumi.set(__self__, "customer", customer)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rank and not isinstance(rank, int):
            raise TypeError("Expected argument 'rank' to be a int")
        pulumi.set(__self__, "rank", rank)
        if saml_sso_info and not isinstance(saml_sso_info, dict):
            raise TypeError("Expected argument 'saml_sso_info' to be a dict")
        pulumi.set(__self__, "saml_sso_info", saml_sso_info)
        if sign_in_behavior and not isinstance(sign_in_behavior, dict):
            raise TypeError("Expected argument 'sign_in_behavior' to be a dict")
        pulumi.set(__self__, "sign_in_behavior", sign_in_behavior)
        if sso_mode and not isinstance(sso_mode, str):
            raise TypeError("Expected argument 'sso_mode' to be a str")
        pulumi.set(__self__, "sso_mode", sso_mode)
        if target_group and not isinstance(target_group, str):
            raise TypeError("Expected argument 'target_group' to be a str")
        pulumi.set(__self__, "target_group", target_group)
        if target_org_unit and not isinstance(target_org_unit, str):
            raise TypeError("Expected argument 'target_org_unit' to be a str")
        pulumi.set(__self__, "target_org_unit", target_org_unit)

    @property
    @pulumi.getter
    def customer(self) -> str:
        """
        Immutable. The customer. For example: `customers/C0123abc`.
        """
        return pulumi.get(self, "customer")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        [Resource name](https://cloud.google.com/apis/design/resource_names) of the Inbound SSO Assignment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rank(self) -> int:
        """
        Must be zero (which is the default value so it can be omitted) for assignments with `target_org_unit` set and must be greater-than-or-equal-to one for assignments with `target_group` set.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="samlSsoInfo")
    def saml_sso_info(self) -> 'outputs.SamlSsoInfoResponse':
        """
        SAML SSO details. Must be set if and only if `sso_mode` is set to `SAML_SSO`.
        """
        return pulumi.get(self, "saml_sso_info")

    @property
    @pulumi.getter(name="signInBehavior")
    def sign_in_behavior(self) -> 'outputs.SignInBehaviorResponse':
        """
        Assertions about users assigned to an IdP will always be accepted from that IdP. This controls whether/when Google should redirect a user to the IdP. Unset (defaults) is the recommended configuration.
        """
        return pulumi.get(self, "sign_in_behavior")

    @property
    @pulumi.getter(name="ssoMode")
    def sso_mode(self) -> str:
        """
        Inbound SSO behavior.
        """
        return pulumi.get(self, "sso_mode")

    @property
    @pulumi.getter(name="targetGroup")
    def target_group(self) -> str:
        """
        Immutable. Must be of the form `groups/{group}`.
        """
        return pulumi.get(self, "target_group")

    @property
    @pulumi.getter(name="targetOrgUnit")
    def target_org_unit(self) -> str:
        """
        Immutable. Must be of the form `orgUnits/{org_unit}`.
        """
        return pulumi.get(self, "target_org_unit")


class AwaitableGetInboundSsoAssignmentResult(GetInboundSsoAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInboundSsoAssignmentResult(
            customer=self.customer,
            name=self.name,
            rank=self.rank,
            saml_sso_info=self.saml_sso_info,
            sign_in_behavior=self.sign_in_behavior,
            sso_mode=self.sso_mode,
            target_group=self.target_group,
            target_org_unit=self.target_org_unit)


def get_inbound_sso_assignment(inbound_sso_assignment_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInboundSsoAssignmentResult:
    """
    Gets an InboundSsoAssignment.
    """
    __args__ = dict()
    __args__['inboundSsoAssignmentId'] = inbound_sso_assignment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudidentity/v1:getInboundSsoAssignment', __args__, opts=opts, typ=GetInboundSsoAssignmentResult).value

    return AwaitableGetInboundSsoAssignmentResult(
        customer=__ret__.customer,
        name=__ret__.name,
        rank=__ret__.rank,
        saml_sso_info=__ret__.saml_sso_info,
        sign_in_behavior=__ret__.sign_in_behavior,
        sso_mode=__ret__.sso_mode,
        target_group=__ret__.target_group,
        target_org_unit=__ret__.target_org_unit)


@_utilities.lift_output_func(get_inbound_sso_assignment)
def get_inbound_sso_assignment_output(inbound_sso_assignment_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInboundSsoAssignmentResult]:
    """
    Gets an InboundSsoAssignment.
    """
    ...
