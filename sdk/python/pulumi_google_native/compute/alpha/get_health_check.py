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
    'GetHealthCheckResult',
    'AwaitableGetHealthCheckResult',
    'get_health_check',
    'get_health_check_output',
]

@pulumi.output_type
class GetHealthCheckResult:
    def __init__(__self__, check_interval_sec=None, creation_timestamp=None, description=None, grpc_health_check=None, healthy_threshold=None, http2_health_check=None, http_health_check=None, https_health_check=None, kind=None, log_config=None, name=None, region=None, self_link=None, self_link_with_id=None, source_regions=None, ssl_health_check=None, tcp_health_check=None, timeout_sec=None, type=None, udp_health_check=None, unhealthy_threshold=None):
        if check_interval_sec and not isinstance(check_interval_sec, int):
            raise TypeError("Expected argument 'check_interval_sec' to be a int")
        pulumi.set(__self__, "check_interval_sec", check_interval_sec)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if grpc_health_check and not isinstance(grpc_health_check, dict):
            raise TypeError("Expected argument 'grpc_health_check' to be a dict")
        pulumi.set(__self__, "grpc_health_check", grpc_health_check)
        if healthy_threshold and not isinstance(healthy_threshold, int):
            raise TypeError("Expected argument 'healthy_threshold' to be a int")
        pulumi.set(__self__, "healthy_threshold", healthy_threshold)
        if http2_health_check and not isinstance(http2_health_check, dict):
            raise TypeError("Expected argument 'http2_health_check' to be a dict")
        pulumi.set(__self__, "http2_health_check", http2_health_check)
        if http_health_check and not isinstance(http_health_check, dict):
            raise TypeError("Expected argument 'http_health_check' to be a dict")
        pulumi.set(__self__, "http_health_check", http_health_check)
        if https_health_check and not isinstance(https_health_check, dict):
            raise TypeError("Expected argument 'https_health_check' to be a dict")
        pulumi.set(__self__, "https_health_check", https_health_check)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if log_config and not isinstance(log_config, dict):
            raise TypeError("Expected argument 'log_config' to be a dict")
        pulumi.set(__self__, "log_config", log_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if source_regions and not isinstance(source_regions, list):
            raise TypeError("Expected argument 'source_regions' to be a list")
        pulumi.set(__self__, "source_regions", source_regions)
        if ssl_health_check and not isinstance(ssl_health_check, dict):
            raise TypeError("Expected argument 'ssl_health_check' to be a dict")
        pulumi.set(__self__, "ssl_health_check", ssl_health_check)
        if tcp_health_check and not isinstance(tcp_health_check, dict):
            raise TypeError("Expected argument 'tcp_health_check' to be a dict")
        pulumi.set(__self__, "tcp_health_check", tcp_health_check)
        if timeout_sec and not isinstance(timeout_sec, int):
            raise TypeError("Expected argument 'timeout_sec' to be a int")
        pulumi.set(__self__, "timeout_sec", timeout_sec)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if udp_health_check and not isinstance(udp_health_check, dict):
            raise TypeError("Expected argument 'udp_health_check' to be a dict")
        pulumi.set(__self__, "udp_health_check", udp_health_check)
        if unhealthy_threshold and not isinstance(unhealthy_threshold, int):
            raise TypeError("Expected argument 'unhealthy_threshold' to be a int")
        pulumi.set(__self__, "unhealthy_threshold", unhealthy_threshold)

    @property
    @pulumi.getter(name="checkIntervalSec")
    def check_interval_sec(self) -> int:
        """
        How often (in seconds) to send a health check. The default value is 5 seconds.
        """
        return pulumi.get(self, "check_interval_sec")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in 3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="grpcHealthCheck")
    def grpc_health_check(self) -> 'outputs.GRPCHealthCheckResponse':
        return pulumi.get(self, "grpc_health_check")

    @property
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> int:
        """
        A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        """
        return pulumi.get(self, "healthy_threshold")

    @property
    @pulumi.getter(name="http2HealthCheck")
    def http2_health_check(self) -> 'outputs.HTTP2HealthCheckResponse':
        return pulumi.get(self, "http2_health_check")

    @property
    @pulumi.getter(name="httpHealthCheck")
    def http_health_check(self) -> 'outputs.HTTPHealthCheckResponse':
        return pulumi.get(self, "http_health_check")

    @property
    @pulumi.getter(name="httpsHealthCheck")
    def https_health_check(self) -> 'outputs.HTTPSHealthCheckResponse':
        return pulumi.get(self, "https_health_check")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> 'outputs.HealthCheckLogConfigResponse':
        """
        Configure logging on this health check.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region where the health check resides. Not applicable to global health checks.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="sourceRegions")
    def source_regions(self) -> Sequence[str]:
        """
        The list of cloud regions from which health checks are performed. If any regions are specified, then exactly 3 regions should be specified. The region names must be valid names of GCP regions. This can only be set for global health check. If this list is non-empty, then there are restrictions on what other health check fields are supported and what other resources can use this health check: - SSL, HTTP2, and GRPC protocols are not supported. - The TCP request field is not supported. - The proxyHeader field for HTTP, HTTPS, and TCP is not supported. - The checkIntervalSec field must be at least 30. - The health check cannot be used with BackendService nor with managed instance group auto-healing. 
        """
        return pulumi.get(self, "source_regions")

    @property
    @pulumi.getter(name="sslHealthCheck")
    def ssl_health_check(self) -> 'outputs.SSLHealthCheckResponse':
        return pulumi.get(self, "ssl_health_check")

    @property
    @pulumi.getter(name="tcpHealthCheck")
    def tcp_health_check(self) -> 'outputs.TCPHealthCheckResponse':
        return pulumi.get(self, "tcp_health_check")

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> int:
        """
        How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        """
        return pulumi.get(self, "timeout_sec")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check fields must be specified, which must match type field.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="udpHealthCheck")
    def udp_health_check(self) -> 'outputs.UDPHealthCheckResponse':
        return pulumi.get(self, "udp_health_check")

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> int:
        """
        A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        return pulumi.get(self, "unhealthy_threshold")


class AwaitableGetHealthCheckResult(GetHealthCheckResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHealthCheckResult(
            check_interval_sec=self.check_interval_sec,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            grpc_health_check=self.grpc_health_check,
            healthy_threshold=self.healthy_threshold,
            http2_health_check=self.http2_health_check,
            http_health_check=self.http_health_check,
            https_health_check=self.https_health_check,
            kind=self.kind,
            log_config=self.log_config,
            name=self.name,
            region=self.region,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            source_regions=self.source_regions,
            ssl_health_check=self.ssl_health_check,
            tcp_health_check=self.tcp_health_check,
            timeout_sec=self.timeout_sec,
            type=self.type,
            udp_health_check=self.udp_health_check,
            unhealthy_threshold=self.unhealthy_threshold)


def get_health_check(health_check: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHealthCheckResult:
    """
    Returns the specified HealthCheck resource.
    """
    __args__ = dict()
    __args__['healthCheck'] = health_check
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getHealthCheck', __args__, opts=opts, typ=GetHealthCheckResult).value

    return AwaitableGetHealthCheckResult(
        check_interval_sec=__ret__.check_interval_sec,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        grpc_health_check=__ret__.grpc_health_check,
        healthy_threshold=__ret__.healthy_threshold,
        http2_health_check=__ret__.http2_health_check,
        http_health_check=__ret__.http_health_check,
        https_health_check=__ret__.https_health_check,
        kind=__ret__.kind,
        log_config=__ret__.log_config,
        name=__ret__.name,
        region=__ret__.region,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        source_regions=__ret__.source_regions,
        ssl_health_check=__ret__.ssl_health_check,
        tcp_health_check=__ret__.tcp_health_check,
        timeout_sec=__ret__.timeout_sec,
        type=__ret__.type,
        udp_health_check=__ret__.udp_health_check,
        unhealthy_threshold=__ret__.unhealthy_threshold)


@_utilities.lift_output_func(get_health_check)
def get_health_check_output(health_check: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHealthCheckResult]:
    """
    Returns the specified HealthCheck resource.
    """
    ...
