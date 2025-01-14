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
    'GetManagedZoneResult',
    'AwaitableGetManagedZoneResult',
    'get_managed_zone',
    'get_managed_zone_output',
]

@pulumi.output_type
class GetManagedZoneResult:
    def __init__(__self__, create_time=None, description=None, dns=None, labels=None, name=None, target_project=None, target_vpc=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dns and not isinstance(dns, str):
            raise TypeError("Expected argument 'dns' to be a str")
        pulumi.set(__self__, "dns", dns)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if target_project and not isinstance(target_project, str):
            raise TypeError("Expected argument 'target_project' to be a str")
        pulumi.set(__self__, "target_project", target_project)
        if target_vpc and not isinstance(target_vpc, str):
            raise TypeError("Expected argument 'target_vpc' to be a str")
        pulumi.set(__self__, "target_vpc", target_vpc)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dns(self) -> str:
        """
        DNS Name of the resource
        """
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the Managed Zone. Format: projects/{project}/locations/global/managedZones/{managed_zone}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="targetProject")
    def target_project(self) -> str:
        """
        The name of the Target Project
        """
        return pulumi.get(self, "target_project")

    @property
    @pulumi.getter(name="targetVpc")
    def target_vpc(self) -> str:
        """
        The name of the Target Project VPC Network
        """
        return pulumi.get(self, "target_vpc")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetManagedZoneResult(GetManagedZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedZoneResult(
            create_time=self.create_time,
            description=self.description,
            dns=self.dns,
            labels=self.labels,
            name=self.name,
            target_project=self.target_project,
            target_vpc=self.target_vpc,
            update_time=self.update_time)


def get_managed_zone(managed_zone: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedZoneResult:
    """
    Gets details of a single ManagedZone.
    """
    __args__ = dict()
    __args__['managedZone'] = managed_zone
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:connectors/v1:getManagedZone', __args__, opts=opts, typ=GetManagedZoneResult).value

    return AwaitableGetManagedZoneResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        dns=__ret__.dns,
        labels=__ret__.labels,
        name=__ret__.name,
        target_project=__ret__.target_project,
        target_vpc=__ret__.target_vpc,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_managed_zone)
def get_managed_zone_output(managed_zone: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetManagedZoneResult]:
    """
    Gets details of a single ManagedZone.
    """
    ...
