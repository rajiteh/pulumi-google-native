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
from ._inputs import *

__all__ = ['TrustConfigArgs', 'TrustConfig']

@pulumi.input_type
class TrustConfigArgs:
    def __init__(__self__, *,
                 trust_config_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 trust_stores: Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreArgs']]]] = None):
        """
        The set of arguments for constructing a TrustConfig resource.
        :param pulumi.Input[str] trust_config_id: Required. A user-provided name of the TrustConfig.
        :param pulumi.Input[str] description: One or more paragraphs of text description of a TrustConfig.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels associated with a TrustConfig.
        :param pulumi.Input[str] name: A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
        :param pulumi.Input[Sequence[pulumi.Input['TrustStoreArgs']]] trust_stores: Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
        """
        pulumi.set(__self__, "trust_config_id", trust_config_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if trust_stores is not None:
            pulumi.set(__self__, "trust_stores", trust_stores)

    @property
    @pulumi.getter(name="trustConfigId")
    def trust_config_id(self) -> pulumi.Input[str]:
        """
        Required. A user-provided name of the TrustConfig.
        """
        return pulumi.get(self, "trust_config_id")

    @trust_config_id.setter
    def trust_config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_config_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        One or more paragraphs of text description of a TrustConfig.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of labels associated with a TrustConfig.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="trustStores")
    def trust_stores(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreArgs']]]]:
        """
        Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
        """
        return pulumi.get(self, "trust_stores")

    @trust_stores.setter
    def trust_stores(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreArgs']]]]):
        pulumi.set(self, "trust_stores", value)


class TrustConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 trust_config_id: Optional[pulumi.Input[str]] = None,
                 trust_stores: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustStoreArgs']]]]] = None,
                 __props__=None):
        """
        Creates a new TrustConfig in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: One or more paragraphs of text description of a TrustConfig.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels associated with a TrustConfig.
        :param pulumi.Input[str] name: A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
        :param pulumi.Input[str] trust_config_id: Required. A user-provided name of the TrustConfig.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustStoreArgs']]]] trust_stores: Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrustConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new TrustConfig in a given project and location.

        :param str resource_name: The name of the resource.
        :param TrustConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrustConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 trust_config_id: Optional[pulumi.Input[str]] = None,
                 trust_stores: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustStoreArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrustConfigArgs.__new__(TrustConfigArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if trust_config_id is None and not opts.urn:
                raise TypeError("Missing required property 'trust_config_id'")
            __props__.__dict__["trust_config_id"] = trust_config_id
            __props__.__dict__["trust_stores"] = trust_stores
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project", "trust_config_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(TrustConfig, __self__).__init__(
            'google-native:certificatemanager/v1:TrustConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TrustConfig':
        """
        Get an existing TrustConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrustConfigArgs.__new__(TrustConfigArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["trust_config_id"] = None
        __props__.__dict__["trust_stores"] = None
        __props__.__dict__["update_time"] = None
        return TrustConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of a TrustConfig.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        One or more paragraphs of text description of a TrustConfig.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Set of labels associated with a TrustConfig.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="trustConfigId")
    def trust_config_id(self) -> pulumi.Output[str]:
        """
        Required. A user-provided name of the TrustConfig.
        """
        return pulumi.get(self, "trust_config_id")

    @property
    @pulumi.getter(name="trustStores")
    def trust_stores(self) -> pulumi.Output[Sequence['outputs.TrustStoreResponse']]:
        """
        Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
        """
        return pulumi.get(self, "trust_stores")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of a TrustConfig.
        """
        return pulumi.get(self, "update_time")
