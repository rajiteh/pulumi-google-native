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
    'GetProxyConfigResult',
    'AwaitableGetProxyConfigResult',
    'get_proxy_config',
    'get_proxy_config_output',
]

@pulumi.output_type
class GetProxyConfigResult:
    def __init__(__self__, authentication_info=None, create_time=None, display_name=None, name=None, proxy_uri=None, routing_info=None, transport_info=None, update_time=None):
        if authentication_info and not isinstance(authentication_info, dict):
            raise TypeError("Expected argument 'authentication_info' to be a dict")
        pulumi.set(__self__, "authentication_info", authentication_info)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if proxy_uri and not isinstance(proxy_uri, str):
            raise TypeError("Expected argument 'proxy_uri' to be a str")
        pulumi.set(__self__, "proxy_uri", proxy_uri)
        if routing_info and not isinstance(routing_info, dict):
            raise TypeError("Expected argument 'routing_info' to be a dict")
        pulumi.set(__self__, "routing_info", routing_info)
        if transport_info and not isinstance(transport_info, dict):
            raise TypeError("Expected argument 'transport_info' to be a dict")
        pulumi.set(__self__, "transport_info", transport_info)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="authenticationInfo")
    def authentication_info(self) -> 'outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaAuthenticationInfoResponse':
        """
        Optional. Information to facilitate Authentication against the proxy server.
        """
        return pulumi.get(self, "authentication_info")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. An arbitrary caller-provided name for the ProxyConfig. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        ProxyConfig resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="proxyUri")
    def proxy_uri(self) -> str:
        """
        The URI of the proxy server.
        """
        return pulumi.get(self, "proxy_uri")

    @property
    @pulumi.getter(name="routingInfo")
    def routing_info(self) -> 'outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaRoutingInfoResponse':
        """
        Routing info to direct traffic to the proxy server.
        """
        return pulumi.get(self, "routing_info")

    @property
    @pulumi.getter(name="transportInfo")
    def transport_info(self) -> 'outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaTransportInfoResponse':
        """
        Transport layer information to verify for the proxy server.
        """
        return pulumi.get(self, "transport_info")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Timestamp when the resource was last modified.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetProxyConfigResult(GetProxyConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProxyConfigResult(
            authentication_info=self.authentication_info,
            create_time=self.create_time,
            display_name=self.display_name,
            name=self.name,
            proxy_uri=self.proxy_uri,
            routing_info=self.routing_info,
            transport_info=self.transport_info,
            update_time=self.update_time)


def get_proxy_config(organization_id: Optional[str] = None,
                     proxy_config_id: Optional[str] = None,
                     tenant_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProxyConfigResult:
    """
    Gets details of a single Tenant.
    """
    __args__ = dict()
    __args__['organizationId'] = organization_id
    __args__['proxyConfigId'] = proxy_config_id
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:beyondcorp/v1alpha:getProxyConfig', __args__, opts=opts, typ=GetProxyConfigResult).value

    return AwaitableGetProxyConfigResult(
        authentication_info=__ret__.authentication_info,
        create_time=__ret__.create_time,
        display_name=__ret__.display_name,
        name=__ret__.name,
        proxy_uri=__ret__.proxy_uri,
        routing_info=__ret__.routing_info,
        transport_info=__ret__.transport_info,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_proxy_config)
def get_proxy_config_output(organization_id: Optional[pulumi.Input[str]] = None,
                            proxy_config_id: Optional[pulumi.Input[str]] = None,
                            tenant_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProxyConfigResult]:
    """
    Gets details of a single Tenant.
    """
    ...
