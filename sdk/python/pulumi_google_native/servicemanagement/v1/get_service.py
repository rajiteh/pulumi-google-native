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
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    def __init__(__self__, producer_project_id=None, service_name=None):
        if producer_project_id and not isinstance(producer_project_id, str):
            raise TypeError("Expected argument 'producer_project_id' to be a str")
        pulumi.set(__self__, "producer_project_id", producer_project_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="producerProjectId")
    def producer_project_id(self) -> str:
        """
        ID of the project that produces and owns this service.
        """
        return pulumi.get(self, "producer_project_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The name of the service. See the [overview](https://cloud.google.com/service-infrastructure/docs/overview) for naming requirements.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            producer_project_id=self.producer_project_id,
            service_name=self.service_name)


def get_service(service_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Gets a managed service. Authentication is required unless the service is public.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:servicemanagement/v1:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        producer_project_id=__ret__.producer_project_id,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_service)
def get_service_output(service_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    Gets a managed service. Authentication is required unless the service is public.
    """
    ...
