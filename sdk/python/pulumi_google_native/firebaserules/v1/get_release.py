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
    'GetReleaseResult',
    'AwaitableGetReleaseResult',
    'get_release',
    'get_release_output',
]

@pulumi.output_type
class GetReleaseResult:
    def __init__(__self__, create_time=None, name=None, ruleset_name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ruleset_name and not isinstance(ruleset_name, str):
            raise TypeError("Expected argument 'ruleset_name' to be a str")
        pulumi.set(__self__, "ruleset_name", ruleset_name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time the release was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Format: `projects/{project_id}/releases/{release_id}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rulesetName")
    def ruleset_name(self) -> str:
        """
        Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        """
        return pulumi.get(self, "ruleset_name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time the release was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetReleaseResult(GetReleaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReleaseResult(
            create_time=self.create_time,
            name=self.name,
            ruleset_name=self.ruleset_name,
            update_time=self.update_time)


def get_release(project: Optional[str] = None,
                release_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReleaseResult:
    """
    Get a `Release` by name.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['releaseId'] = release_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:firebaserules/v1:getRelease', __args__, opts=opts, typ=GetReleaseResult).value

    return AwaitableGetReleaseResult(
        create_time=__ret__.create_time,
        name=__ret__.name,
        ruleset_name=__ret__.ruleset_name,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_release)
def get_release_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                       release_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReleaseResult]:
    """
    Get a `Release` by name.
    """
    ...
