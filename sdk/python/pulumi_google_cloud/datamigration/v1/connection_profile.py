# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['ConnectionProfile']


class ConnectionProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudsql: Optional[pulumi.Input[pulumi.InputType['CloudSqlConnectionProfileArgs']]] = None,
                 connection_profile_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 error: Optional[pulumi.Input[pulumi.InputType['StatusArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mysql: Optional[pulumi.Input[pulumi.InputType['MySqlConnectionProfileArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 postgresql: Optional[pulumi.Input[pulumi.InputType['PostgreSqlConnectionProfileArgs']]] = None,
                 provider: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new connection profile in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CloudSqlConnectionProfileArgs']] cloudsql: A CloudSQL database connection profile.
        :param pulumi.Input[str] connection_profile_id: Required. The connection profile identifier.
        :param pulumi.Input[str] create_time: Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[str] display_name: The connection profile display name.
        :param pulumi.Input[pulumi.InputType['StatusArgs']] error: Output only. The error details in case of state FAILED.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        :param pulumi.Input[pulumi.InputType['MySqlConnectionProfileArgs']] mysql: A MySQL database connection profile.
        :param pulumi.Input[str] name: The name of this connection profile resource in the form of projects/{project}/locations/{location}/instances/{instance}.
        :param pulumi.Input[str] parent: Required. The parent, which owns this collection of connection profiles.
        :param pulumi.Input[pulumi.InputType['PostgreSqlConnectionProfileArgs']] postgresql: A PostgreSQL database connection profile.
        :param pulumi.Input[str] provider: The database provider.
        :param pulumi.Input[str] request_id: A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        :param pulumi.Input[str] state: The current connection profile state (e.g. DRAFT, READY, or FAILED).
        :param pulumi.Input[str] update_time: Output only. The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cloudsql'] = cloudsql
            __props__['connection_profile_id'] = connection_profile_id
            __props__['create_time'] = create_time
            __props__['display_name'] = display_name
            __props__['error'] = error
            __props__['labels'] = labels
            __props__['mysql'] = mysql
            __props__['name'] = name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['postgresql'] = postgresql
            __props__['provider'] = provider
            __props__['request_id'] = request_id
            __props__['state'] = state
            __props__['update_time'] = update_time
        super(ConnectionProfile, __self__).__init__(
            'google-cloud:datamigration/v1:ConnectionProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConnectionProfile':
        """
        Get an existing ConnectionProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ConnectionProfile(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
