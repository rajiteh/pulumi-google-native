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
    'GetTargetVpnGatewayResult',
    'AwaitableGetTargetVpnGatewayResult',
    'get_target_vpn_gateway',
    'get_target_vpn_gateway_output',
]

@pulumi.output_type
class GetTargetVpnGatewayResult:
    def __init__(__self__, creation_timestamp=None, description=None, forwarding_rules=None, kind=None, label_fingerprint=None, labels=None, name=None, network=None, region=None, self_link=None, status=None, tunnels=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if forwarding_rules and not isinstance(forwarding_rules, list):
            raise TypeError("Expected argument 'forwarding_rules' to be a list")
        pulumi.set(__self__, "forwarding_rules", forwarding_rules)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tunnels and not isinstance(tunnels, list):
            raise TypeError("Expected argument 'tunnels' to be a list")
        pulumi.set(__self__, "tunnels", tunnels)

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
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> Sequence[str]:
        """
        A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
        """
        return pulumi.get(self, "forwarding_rules")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of resource. Always compute#targetVpnGateway for target VPN gateways.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this TargetVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a TargetVpnGateway.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
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
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tunnels(self) -> Sequence[str]:
        """
        A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
        """
        return pulumi.get(self, "tunnels")


class AwaitableGetTargetVpnGatewayResult(GetTargetVpnGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetVpnGatewayResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            forwarding_rules=self.forwarding_rules,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            name=self.name,
            network=self.network,
            region=self.region,
            self_link=self.self_link,
            status=self.status,
            tunnels=self.tunnels)


def get_target_vpn_gateway(project: Optional[str] = None,
                           region: Optional[str] = None,
                           target_vpn_gateway: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetVpnGatewayResult:
    """
    Returns the specified target VPN gateway.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    __args__['targetVpnGateway'] = target_vpn_gateway
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getTargetVpnGateway', __args__, opts=opts, typ=GetTargetVpnGatewayResult).value

    return AwaitableGetTargetVpnGatewayResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        forwarding_rules=__ret__.forwarding_rules,
        kind=__ret__.kind,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        name=__ret__.name,
        network=__ret__.network,
        region=__ret__.region,
        self_link=__ret__.self_link,
        status=__ret__.status,
        tunnels=__ret__.tunnels)


@_utilities.lift_output_func(get_target_vpn_gateway)
def get_target_vpn_gateway_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[str]] = None,
                                  target_vpn_gateway: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTargetVpnGatewayResult]:
    """
    Returns the specified target VPN gateway.
    """
    ...
