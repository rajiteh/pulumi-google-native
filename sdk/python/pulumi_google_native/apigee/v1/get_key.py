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
    'GetKeyResult',
    'AwaitableGetKeyResult',
    'get_key',
    'get_key_output',
]

@pulumi.output_type
class GetKeyResult:
    def __init__(__self__, api_products=None, attributes=None, consumer_key=None, consumer_secret=None, expires_at=None, expires_in_seconds=None, issued_at=None, scopes=None, status=None):
        if api_products and not isinstance(api_products, list):
            raise TypeError("Expected argument 'api_products' to be a list")
        pulumi.set(__self__, "api_products", api_products)
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if consumer_key and not isinstance(consumer_key, str):
            raise TypeError("Expected argument 'consumer_key' to be a str")
        pulumi.set(__self__, "consumer_key", consumer_key)
        if consumer_secret and not isinstance(consumer_secret, str):
            raise TypeError("Expected argument 'consumer_secret' to be a str")
        pulumi.set(__self__, "consumer_secret", consumer_secret)
        if expires_at and not isinstance(expires_at, str):
            raise TypeError("Expected argument 'expires_at' to be a str")
        pulumi.set(__self__, "expires_at", expires_at)
        if expires_in_seconds and not isinstance(expires_in_seconds, str):
            raise TypeError("Expected argument 'expires_in_seconds' to be a str")
        pulumi.set(__self__, "expires_in_seconds", expires_in_seconds)
        if issued_at and not isinstance(issued_at, str):
            raise TypeError("Expected argument 'issued_at' to be a str")
        pulumi.set(__self__, "issued_at", issued_at)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="apiProducts")
    def api_products(self) -> Sequence[Any]:
        """
        List of API products for which the credential can be used. **Note**: Do not specify the list of API products when creating a consumer key and secret for a developer app. Instead, use the UpdateDeveloperAppKey API to make the association after the consumer key and secret are created.
        """
        return pulumi.get(self, "api_products")

    @property
    @pulumi.getter
    def attributes(self) -> Sequence['outputs.GoogleCloudApigeeV1AttributeResponse']:
        """
        List of attributes associated with the credential.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="consumerKey")
    def consumer_key(self) -> str:
        """
        Consumer key.
        """
        return pulumi.get(self, "consumer_key")

    @property
    @pulumi.getter(name="consumerSecret")
    def consumer_secret(self) -> str:
        """
        Secret key.
        """
        return pulumi.get(self, "consumer_secret")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> str:
        """
        Time the developer app expires in milliseconds since epoch.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="expiresInSeconds")
    def expires_in_seconds(self) -> str:
        """
        Input only. Expiration time, in seconds, for the consumer key. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
        """
        return pulumi.get(self, "expires_in_seconds")

    @property
    @pulumi.getter(name="issuedAt")
    def issued_at(self) -> str:
        """
        Time the developer app was created in milliseconds since epoch.
        """
        return pulumi.get(self, "issued_at")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        Scopes to apply to the app. The specified scope names must already be defined for the API product that you associate with the app.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the credential. Valid values include `approved` or `revoked`.
        """
        return pulumi.get(self, "status")


class AwaitableGetKeyResult(GetKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyResult(
            api_products=self.api_products,
            attributes=self.attributes,
            consumer_key=self.consumer_key,
            consumer_secret=self.consumer_secret,
            expires_at=self.expires_at,
            expires_in_seconds=self.expires_in_seconds,
            issued_at=self.issued_at,
            scopes=self.scopes,
            status=self.status)


def get_key(app_id: Optional[str] = None,
            developer_id: Optional[str] = None,
            key_id: Optional[str] = None,
            organization_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyResult:
    """
    Gets details for a consumer key for a developer app, including the key and secret value, associated API products, and other information.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['developerId'] = developer_id
    __args__['keyId'] = key_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getKey', __args__, opts=opts, typ=GetKeyResult).value

    return AwaitableGetKeyResult(
        api_products=__ret__.api_products,
        attributes=__ret__.attributes,
        consumer_key=__ret__.consumer_key,
        consumer_secret=__ret__.consumer_secret,
        expires_at=__ret__.expires_at,
        expires_in_seconds=__ret__.expires_in_seconds,
        issued_at=__ret__.issued_at,
        scopes=__ret__.scopes,
        status=__ret__.status)


@_utilities.lift_output_func(get_key)
def get_key_output(app_id: Optional[pulumi.Input[str]] = None,
                   developer_id: Optional[pulumi.Input[str]] = None,
                   key_id: Optional[pulumi.Input[str]] = None,
                   organization_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyResult]:
    """
    Gets details for a consumer key for a developer app, including the key and secret value, associated API products, and other information.
    """
    ...
