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
    'GetBitbucketServerConfigResult',
    'AwaitableGetBitbucketServerConfigResult',
    'get_bitbucket_server_config',
    'get_bitbucket_server_config_output',
]

@pulumi.output_type
class GetBitbucketServerConfigResult:
    def __init__(__self__, api_key=None, connected_repositories=None, create_time=None, host_uri=None, name=None, peered_network=None, secrets=None, ssl_ca=None, username=None, webhook_key=None):
        if api_key and not isinstance(api_key, str):
            raise TypeError("Expected argument 'api_key' to be a str")
        pulumi.set(__self__, "api_key", api_key)
        if connected_repositories and not isinstance(connected_repositories, list):
            raise TypeError("Expected argument 'connected_repositories' to be a list")
        pulumi.set(__self__, "connected_repositories", connected_repositories)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if host_uri and not isinstance(host_uri, str):
            raise TypeError("Expected argument 'host_uri' to be a str")
        pulumi.set(__self__, "host_uri", host_uri)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peered_network and not isinstance(peered_network, str):
            raise TypeError("Expected argument 'peered_network' to be a str")
        pulumi.set(__self__, "peered_network", peered_network)
        if secrets and not isinstance(secrets, dict):
            raise TypeError("Expected argument 'secrets' to be a dict")
        pulumi.set(__self__, "secrets", secrets)
        if ssl_ca and not isinstance(ssl_ca, str):
            raise TypeError("Expected argument 'ssl_ca' to be a str")
        pulumi.set(__self__, "ssl_ca", ssl_ca)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if webhook_key and not isinstance(webhook_key, str):
            raise TypeError("Expected argument 'webhook_key' to be a str")
        pulumi.set(__self__, "webhook_key", webhook_key)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="connectedRepositories")
    def connected_repositories(self) -> Sequence['outputs.BitbucketServerRepositoryIdResponse']:
        """
        Connected Bitbucket Server repositories for this config.
        """
        return pulumi.get(self, "connected_repositories")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the config was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="hostUri")
    def host_uri(self) -> str:
        """
        Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed. If you need to change it, please create another BitbucketServerConfig.
        """
        return pulumi.get(self, "host_uri")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the config.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeredNetwork")
    def peered_network(self) -> str:
        """
        Optional. The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection. This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty, no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format `projects/{project}/global/networks/{network}`, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        """
        return pulumi.get(self, "peered_network")

    @property
    @pulumi.getter
    def secrets(self) -> 'outputs.BitbucketServerSecretsResponse':
        """
        Secret Manager secrets needed by the config.
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter(name="sslCa")
    def ssl_ca(self) -> str:
        """
        Optional. SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
        """
        return pulumi.get(self, "ssl_ca")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Username of the account Cloud Build will use on Bitbucket Server.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="webhookKey")
    def webhook_key(self) -> str:
        """
        UUID included in webhook requests. The UUID is used to look up the corresponding config.
        """
        return pulumi.get(self, "webhook_key")


class AwaitableGetBitbucketServerConfigResult(GetBitbucketServerConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBitbucketServerConfigResult(
            api_key=self.api_key,
            connected_repositories=self.connected_repositories,
            create_time=self.create_time,
            host_uri=self.host_uri,
            name=self.name,
            peered_network=self.peered_network,
            secrets=self.secrets,
            ssl_ca=self.ssl_ca,
            username=self.username,
            webhook_key=self.webhook_key)


def get_bitbucket_server_config(bitbucket_server_config_id: Optional[str] = None,
                                location: Optional[str] = None,
                                project: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBitbucketServerConfigResult:
    """
    Retrieve a `BitbucketServerConfig`. This API is experimental.
    """
    __args__ = dict()
    __args__['bitbucketServerConfigId'] = bitbucket_server_config_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudbuild/v1:getBitbucketServerConfig', __args__, opts=opts, typ=GetBitbucketServerConfigResult).value

    return AwaitableGetBitbucketServerConfigResult(
        api_key=__ret__.api_key,
        connected_repositories=__ret__.connected_repositories,
        create_time=__ret__.create_time,
        host_uri=__ret__.host_uri,
        name=__ret__.name,
        peered_network=__ret__.peered_network,
        secrets=__ret__.secrets,
        ssl_ca=__ret__.ssl_ca,
        username=__ret__.username,
        webhook_key=__ret__.webhook_key)


@_utilities.lift_output_func(get_bitbucket_server_config)
def get_bitbucket_server_config_output(bitbucket_server_config_id: Optional[pulumi.Input[str]] = None,
                                       location: Optional[pulumi.Input[str]] = None,
                                       project: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBitbucketServerConfigResult]:
    """
    Retrieve a `BitbucketServerConfig`. This API is experimental.
    """
    ...
