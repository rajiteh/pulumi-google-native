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
    'GetDataCollectorResult',
    'AwaitableGetDataCollectorResult',
    'get_data_collector',
    'get_data_collector_output',
]

@pulumi.output_type
class GetDataCollectorResult:
    def __init__(__self__, created_at=None, description=None, last_modified_at=None, name=None, type=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if last_modified_at and not isinstance(last_modified_at, str):
            raise TypeError("Expected argument 'last_modified_at' to be a str")
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time at which the data collector was created in milliseconds since the epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the data collector.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        The time at which the Data Collector was last updated in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        ID of the data collector. Must begin with `dc_`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Immutable. The type of data this data collector will collect.
        """
        return pulumi.get(self, "type")


class AwaitableGetDataCollectorResult(GetDataCollectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataCollectorResult(
            created_at=self.created_at,
            description=self.description,
            last_modified_at=self.last_modified_at,
            name=self.name,
            type=self.type)


def get_data_collector(datacollector_id: Optional[str] = None,
                       organization_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataCollectorResult:
    """
    Gets a data collector.
    """
    __args__ = dict()
    __args__['datacollectorId'] = datacollector_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getDataCollector', __args__, opts=opts, typ=GetDataCollectorResult).value

    return AwaitableGetDataCollectorResult(
        created_at=__ret__.created_at,
        description=__ret__.description,
        last_modified_at=__ret__.last_modified_at,
        name=__ret__.name,
        type=__ret__.type)


@_utilities.lift_output_func(get_data_collector)
def get_data_collector_output(datacollector_id: Optional[pulumi.Input[str]] = None,
                              organization_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataCollectorResult]:
    """
    Gets a data collector.
    """
    ...
