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
    'GetVersionResult',
    'AwaitableGetVersionResult',
    'get_version',
    'get_version_output',
]

@pulumi.output_type
class GetVersionResult:
    def __init__(__self__, create_time=None, description=None, name=None, status=None, version_number=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if version_number and not isinstance(version_number, int):
            raise TypeError("Expected argument 'version_number' to be a int")
        pulumi.set(__self__, "version_number", version_number)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of this version. This field is read-only, i.e., it cannot be set by create and update methods.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The developer-provided description of this version.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of this agent version. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of this version. This field is read-only and cannot be set by create and update methods.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> int:
        """
        The sequential number of this version. This field is read-only which means it cannot be set by create and update methods.
        """
        return pulumi.get(self, "version_number")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            create_time=self.create_time,
            description=self.description,
            name=self.name,
            status=self.status,
            version_number=self.version_number)


def get_version(location: Optional[str] = None,
                project: Optional[str] = None,
                version_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Retrieves the specified agent version.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['versionId'] = version_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v2:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        name=__ret__.name,
        status=__ret__.status,
        version_number=__ret__.version_number)


@_utilities.lift_output_func(get_version)
def get_version_output(location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       version_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Retrieves the specified agent version.
    """
    ...
