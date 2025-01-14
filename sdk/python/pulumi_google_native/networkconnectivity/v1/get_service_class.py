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
    'GetServiceClassResult',
    'AwaitableGetServiceClassResult',
    'get_service_class',
    'get_service_class_output',
]

@pulumi.output_type
class GetServiceClassResult:
    def __init__(__self__, create_time=None, description=None, labels=None, name=None, service_class=None, service_connection_maps=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_class and not isinstance(service_class, str):
            raise TypeError("Expected argument 'service_class' to be a str")
        pulumi.set(__self__, "service_class", service_class)
        if service_connection_maps and not isinstance(service_connection_maps, list):
            raise TypeError("Expected argument 'service_connection_maps' to be a list")
        pulumi.set(__self__, "service_connection_maps", service_connection_maps)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the ServiceClass was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        User-defined labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceClass")
    def service_class(self) -> str:
        """
        The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
        """
        return pulumi.get(self, "service_class")

    @property
    @pulumi.getter(name="serviceConnectionMaps")
    def service_connection_maps(self) -> Sequence[str]:
        """
        URIs of all Service Connection Maps using this service class.
        """
        return pulumi.get(self, "service_connection_maps")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time when the ServiceClass was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetServiceClassResult(GetServiceClassResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceClassResult(
            create_time=self.create_time,
            description=self.description,
            labels=self.labels,
            name=self.name,
            service_class=self.service_class,
            service_connection_maps=self.service_connection_maps,
            update_time=self.update_time)


def get_service_class(location: Optional[str] = None,
                      project: Optional[str] = None,
                      service_class_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceClassResult:
    """
    Gets details of a single ServiceClass.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['serviceClassId'] = service_class_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:networkconnectivity/v1:getServiceClass', __args__, opts=opts, typ=GetServiceClassResult).value

    return AwaitableGetServiceClassResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        labels=__ret__.labels,
        name=__ret__.name,
        service_class=__ret__.service_class,
        service_connection_maps=__ret__.service_connection_maps,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_service_class)
def get_service_class_output(location: Optional[pulumi.Input[str]] = None,
                             project: Optional[pulumi.Input[Optional[str]]] = None,
                             service_class_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceClassResult]:
    """
    Gets details of a single ServiceClass.
    """
    ...
