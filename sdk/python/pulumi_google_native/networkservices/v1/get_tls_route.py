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
    'GetTlsRouteResult',
    'AwaitableGetTlsRouteResult',
    'get_tls_route',
    'get_tls_route_output',
]

@pulumi.output_type
class GetTlsRouteResult:
    def __init__(__self__, create_time=None, description=None, gateways=None, meshes=None, name=None, rules=None, self_link=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if gateways and not isinstance(gateways, list):
            raise TypeError("Expected argument 'gateways' to be a list")
        pulumi.set(__self__, "gateways", gateways)
        if meshes and not isinstance(meshes, list):
            raise TypeError("Expected argument 'meshes' to be a list")
        pulumi.set(__self__, "meshes", meshes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def gateways(self) -> Sequence[str]:
        """
        Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        """
        return pulumi.get(self, "gateways")

    @property
    @pulumi.getter
    def meshes(self) -> Sequence[str]:
        """
        Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        """
        return pulumi.get(self, "meshes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.TlsRouteRouteRuleResponse']:
        """
        Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL of this resource
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp when the resource was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetTlsRouteResult(GetTlsRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTlsRouteResult(
            create_time=self.create_time,
            description=self.description,
            gateways=self.gateways,
            meshes=self.meshes,
            name=self.name,
            rules=self.rules,
            self_link=self.self_link,
            update_time=self.update_time)


def get_tls_route(location: Optional[str] = None,
                  project: Optional[str] = None,
                  tls_route_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTlsRouteResult:
    """
    Gets details of a single TlsRoute.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['tlsRouteId'] = tls_route_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:networkservices/v1:getTlsRoute', __args__, opts=opts, typ=GetTlsRouteResult).value

    return AwaitableGetTlsRouteResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        gateways=__ret__.gateways,
        meshes=__ret__.meshes,
        name=__ret__.name,
        rules=__ret__.rules,
        self_link=__ret__.self_link,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_tls_route)
def get_tls_route_output(location: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         tls_route_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTlsRouteResult]:
    """
    Gets details of a single TlsRoute.
    """
    ...