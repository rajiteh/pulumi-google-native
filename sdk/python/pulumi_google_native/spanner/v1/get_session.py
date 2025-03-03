# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetSessionResult',
    'AwaitableGetSessionResult',
    'get_session',
    'get_session_output',
]

@pulumi.output_type
class GetSessionResult:
    def __init__(__self__, approximate_last_use_time=None, create_time=None, creator_role=None, labels=None, name=None):
        if approximate_last_use_time and not isinstance(approximate_last_use_time, str):
            raise TypeError("Expected argument 'approximate_last_use_time' to be a str")
        pulumi.set(__self__, "approximate_last_use_time", approximate_last_use_time)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if creator_role and not isinstance(creator_role, str):
            raise TypeError("Expected argument 'creator_role' to be a str")
        pulumi.set(__self__, "creator_role", creator_role)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="approximateLastUseTime")
    def approximate_last_use_time(self) -> str:
        """
        The approximate timestamp when the session is last used. It is typically earlier than the actual last use time.
        """
        return pulumi.get(self, "approximate_last_use_time")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the session is created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="creatorRole")
    def creator_role(self) -> str:
        """
        The database role which created this session.
        """
        return pulumi.get(self, "creator_role")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels for the session. * Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. * Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`. * No more than 64 labels can be associated with a given session. See https://goo.gl/xmQnxf for more information on and examples of labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the session. This is always system-assigned.
        """
        return pulumi.get(self, "name")


class AwaitableGetSessionResult(GetSessionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSessionResult(
            approximate_last_use_time=self.approximate_last_use_time,
            create_time=self.create_time,
            creator_role=self.creator_role,
            labels=self.labels,
            name=self.name)


def get_session(database_id: Optional[str] = None,
                instance_id: Optional[str] = None,
                project: Optional[str] = None,
                session_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSessionResult:
    """
    Gets a session. Returns `NOT_FOUND` if the session does not exist. This is mainly useful for determining whether a session is still alive.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    __args__['instanceId'] = instance_id
    __args__['project'] = project
    __args__['sessionId'] = session_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:spanner/v1:getSession', __args__, opts=opts, typ=GetSessionResult).value

    return AwaitableGetSessionResult(
        approximate_last_use_time=__ret__.approximate_last_use_time,
        create_time=__ret__.create_time,
        creator_role=__ret__.creator_role,
        labels=__ret__.labels,
        name=__ret__.name)


@_utilities.lift_output_func(get_session)
def get_session_output(database_id: Optional[pulumi.Input[str]] = None,
                       instance_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       session_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSessionResult]:
    """
    Gets a session. Returns `NOT_FOUND` if the session does not exist. This is mainly useful for determining whether a session is still alive.
    """
    ...
