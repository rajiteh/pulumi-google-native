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
    'GetDefaultSupportedIdpConfigResult',
    'AwaitableGetDefaultSupportedIdpConfigResult',
    'get_default_supported_idp_config',
    'get_default_supported_idp_config_output',
]

@pulumi.output_type
class GetDefaultSupportedIdpConfigResult:
    def __init__(__self__, apple_sign_in_config=None, client_id=None, client_secret=None, enabled=None, name=None):
        if apple_sign_in_config and not isinstance(apple_sign_in_config, dict):
            raise TypeError("Expected argument 'apple_sign_in_config' to be a dict")
        pulumi.set(__self__, "apple_sign_in_config", apple_sign_in_config)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="appleSignInConfig")
    def apple_sign_in_config(self) -> 'outputs.GoogleCloudIdentitytoolkitAdminV2AppleSignInConfigResponse':
        """
        Additional config for Apple-based projects.
        """
        return pulumi.get(self, "apple_sign_in_config")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        OAuth client ID.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        """
        OAuth client secret.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        True if allows the user to sign in with the provider.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the DefaultSupportedIdpConfig resource, for example: "projects/my-awesome-project/defaultSupportedIdpConfigs/google.com"
        """
        return pulumi.get(self, "name")


class AwaitableGetDefaultSupportedIdpConfigResult(GetDefaultSupportedIdpConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultSupportedIdpConfigResult(
            apple_sign_in_config=self.apple_sign_in_config,
            client_id=self.client_id,
            client_secret=self.client_secret,
            enabled=self.enabled,
            name=self.name)


def get_default_supported_idp_config(default_supported_idp_config_id: Optional[str] = None,
                                     project: Optional[str] = None,
                                     tenant_id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultSupportedIdpConfigResult:
    """
    Retrieve a default supported Idp configuration for an Identity Toolkit project.
    """
    __args__ = dict()
    __args__['defaultSupportedIdpConfigId'] = default_supported_idp_config_id
    __args__['project'] = project
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:identitytoolkit/v2:getDefaultSupportedIdpConfig', __args__, opts=opts, typ=GetDefaultSupportedIdpConfigResult).value

    return AwaitableGetDefaultSupportedIdpConfigResult(
        apple_sign_in_config=__ret__.apple_sign_in_config,
        client_id=__ret__.client_id,
        client_secret=__ret__.client_secret,
        enabled=__ret__.enabled,
        name=__ret__.name)


@_utilities.lift_output_func(get_default_supported_idp_config)
def get_default_supported_idp_config_output(default_supported_idp_config_id: Optional[pulumi.Input[str]] = None,
                                            project: Optional[pulumi.Input[Optional[str]]] = None,
                                            tenant_id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultSupportedIdpConfigResult]:
    """
    Retrieve a default supported Idp configuration for an Identity Toolkit project.
    """
    ...
