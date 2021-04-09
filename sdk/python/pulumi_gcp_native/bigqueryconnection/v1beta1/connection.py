# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Connection']


class Connection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_sql: Optional[pulumi.Input[pulumi.InputType['CloudSqlPropertiesArgs']]] = None,
                 connections_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new connection.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CloudSqlPropertiesArgs']] cloud_sql: Cloud SQL properties.
        :param pulumi.Input[str] description: User provided description.
        :param pulumi.Input[str] friendly_name: User provided display name for the connection.
        :param pulumi.Input[str] name: The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
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

            __props__['cloud_sql'] = cloud_sql
            if connections_id is None and not opts.urn:
                raise TypeError("Missing required property 'connections_id'")
            __props__['connections_id'] = connections_id
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__['locations_id'] = locations_id
            __props__['name'] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__['projects_id'] = projects_id
            __props__['creation_time'] = None
            __props__['has_credential'] = None
            __props__['last_modified_time'] = None
        super(Connection, __self__).__init__(
            'gcp-native:bigqueryconnection/v1beta1:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cloud_sql"] = None
        __props__["creation_time"] = None
        __props__["description"] = None
        __props__["friendly_name"] = None
        __props__["has_credential"] = None
        __props__["last_modified_time"] = None
        __props__["name"] = None
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudSql")
    def cloud_sql(self) -> pulumi.Output['outputs.CloudSqlPropertiesResponse']:
        """
        Cloud SQL properties.
        """
        return pulumi.get(self, "cloud_sql")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of the connection.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User provided description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[str]:
        """
        User provided display name for the connection.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="hasCredential")
    def has_credential(self) -> pulumi.Output[bool]:
        """
        True, if credential is configured for this connection.
        """
        return pulumi.get(self, "has_credential")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of the connection.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
