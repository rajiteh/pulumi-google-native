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
    'GetRegionTargetHttpProxyResult',
    'AwaitableGetRegionTargetHttpProxyResult',
    'get_region_target_http_proxy',
    'get_region_target_http_proxy_output',
]

@pulumi.output_type
class GetRegionTargetHttpProxyResult:
    def __init__(__self__, creation_timestamp=None, description=None, fingerprint=None, http_filters=None, kind=None, name=None, proxy_bind=None, region=None, self_link=None, url_map=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if http_filters and not isinstance(http_filters, list):
            raise TypeError("Expected argument 'http_filters' to be a list")
        pulumi.set(__self__, "http_filters", http_filters)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if proxy_bind and not isinstance(proxy_bind, bool):
            raise TypeError("Expected argument 'proxy_bind' to be a bool")
        pulumi.set(__self__, "proxy_bind", proxy_bind)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if url_map and not isinstance(url_map, str):
            raise TypeError("Expected argument 'url_map' to be a str")
        pulumi.set(__self__, "url_map", url_map)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
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
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="httpFilters")
    def http_filters(self) -> Sequence[str]:
        """
        URLs to networkservices.HttpFilter resources enabled for xDS clients using this configuration. For example, https://networkservices.googleapis.com/v1alpha1/projects/project/locations/ locationhttpFilters/httpFilter Only filters that handle outbound connection and stream events may be specified. These filters work in conjunction with a default set of HTTP filters that may already be configured by Traffic Director. Traffic Director will determine the final location of these filters within xDS configuration based on the name of the HTTP filter. If Traffic Director positions multiple filters at the same location, those filters will be in the same order as specified in this list. httpFilters only applies for loadbalancers with loadBalancingScheme set to INTERNAL_SELF_MANAGED. See ForwardingRule for more details.
        """
        return pulumi.get(self, "http_filters")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="proxyBind")
    def proxy_bind(self) -> bool:
        """
        This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
        """
        return pulumi.get(self, "proxy_bind")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
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
    @pulumi.getter(name="urlMap")
    def url_map(self) -> str:
        """
        URL to the UrlMap resource that defines the mapping from URL to the BackendService.
        """
        return pulumi.get(self, "url_map")


class AwaitableGetRegionTargetHttpProxyResult(GetRegionTargetHttpProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionTargetHttpProxyResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            fingerprint=self.fingerprint,
            http_filters=self.http_filters,
            kind=self.kind,
            name=self.name,
            proxy_bind=self.proxy_bind,
            region=self.region,
            self_link=self.self_link,
            url_map=self.url_map)


def get_region_target_http_proxy(project: Optional[str] = None,
                                 region: Optional[str] = None,
                                 target_http_proxy: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionTargetHttpProxyResult:
    """
    Returns the specified TargetHttpProxy resource in the specified region.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['targetHttpProxy'] = target_http_proxy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getRegionTargetHttpProxy', __args__, opts=opts, typ=GetRegionTargetHttpProxyResult).value

    return AwaitableGetRegionTargetHttpProxyResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        fingerprint=__ret__.fingerprint,
        http_filters=__ret__.http_filters,
        kind=__ret__.kind,
        name=__ret__.name,
        proxy_bind=__ret__.proxy_bind,
        region=__ret__.region,
        self_link=__ret__.self_link,
        url_map=__ret__.url_map)


@_utilities.lift_output_func(get_region_target_http_proxy)
def get_region_target_http_proxy_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                        region: Optional[pulumi.Input[str]] = None,
                                        target_http_proxy: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionTargetHttpProxyResult]:
    """
    Returns the specified TargetHttpProxy resource in the specified region.
    """
    ...
