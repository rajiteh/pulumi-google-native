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
    'GetSslPolicyResult',
    'AwaitableGetSslPolicyResult',
    'get_ssl_policy',
    'get_ssl_policy_output',
]

@pulumi.output_type
class GetSslPolicyResult:
    def __init__(__self__, creation_timestamp=None, custom_features=None, description=None, enabled_features=None, fingerprint=None, kind=None, min_tls_version=None, name=None, profile=None, region=None, self_link=None, warnings=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_features and not isinstance(custom_features, list):
            raise TypeError("Expected argument 'custom_features' to be a list")
        pulumi.set(__self__, "custom_features", custom_features)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled_features and not isinstance(enabled_features, list):
            raise TypeError("Expected argument 'enabled_features' to be a list")
        pulumi.set(__self__, "enabled_features", enabled_features)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if min_tls_version and not isinstance(min_tls_version, str):
            raise TypeError("Expected argument 'min_tls_version' to be a str")
        pulumi.set(__self__, "min_tls_version", min_tls_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if profile and not isinstance(profile, str):
            raise TypeError("Expected argument 'profile' to be a str")
        pulumi.set(__self__, "profile", profile)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if warnings and not isinstance(warnings, list):
            raise TypeError("Expected argument 'warnings' to be a list")
        pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> Sequence[str]:
        """
        A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        """
        return pulumi.get(self, "custom_features")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enabledFeatures")
    def enabled_features(self) -> Sequence[str]:
        """
        The list of features enabled in the SSL policy.
        """
        return pulumi.get(self, "enabled_features")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> str:
        """
        The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        """
        return pulumi.get(self, "min_tls_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def profile(self) -> str:
        """
        Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional SSL policy resides. This field is not applicable to global SSL policies.
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
    def warnings(self) -> Sequence['outputs.SslPolicyWarningsItemResponse']:
        """
        If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        """
        return pulumi.get(self, "warnings")


class AwaitableGetSslPolicyResult(GetSslPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSslPolicyResult(
            creation_timestamp=self.creation_timestamp,
            custom_features=self.custom_features,
            description=self.description,
            enabled_features=self.enabled_features,
            fingerprint=self.fingerprint,
            kind=self.kind,
            min_tls_version=self.min_tls_version,
            name=self.name,
            profile=self.profile,
            region=self.region,
            self_link=self.self_link,
            warnings=self.warnings)


def get_ssl_policy(project: Optional[str] = None,
                   ssl_policy: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSslPolicyResult:
    """
    Lists all of the ordered rules present in a single specified policy.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['sslPolicy'] = ssl_policy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getSslPolicy', __args__, opts=opts, typ=GetSslPolicyResult).value

    return AwaitableGetSslPolicyResult(
        creation_timestamp=__ret__.creation_timestamp,
        custom_features=__ret__.custom_features,
        description=__ret__.description,
        enabled_features=__ret__.enabled_features,
        fingerprint=__ret__.fingerprint,
        kind=__ret__.kind,
        min_tls_version=__ret__.min_tls_version,
        name=__ret__.name,
        profile=__ret__.profile,
        region=__ret__.region,
        self_link=__ret__.self_link,
        warnings=__ret__.warnings)


@_utilities.lift_output_func(get_ssl_policy)
def get_ssl_policy_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                          ssl_policy: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSslPolicyResult]:
    """
    Lists all of the ordered rules present in a single specified policy.
    """
    ...
