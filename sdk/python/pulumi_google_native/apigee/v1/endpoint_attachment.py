# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['EndpointAttachmentArgs', 'EndpointAttachment']

@pulumi.input_type
class EndpointAttachmentArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointAttachment resource.
        :param pulumi.Input[str] endpoint_attachment_id: ID to use for the endpoint attachment. ID must start with a lowercase letter followed by up to 31 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. The minimum length is 2.
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] name: Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        pulumi.set(__self__, "organization_id", organization_id)
        if endpoint_attachment_id is not None:
            pulumi.set(__self__, "endpoint_attachment_id", endpoint_attachment_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_attachment is not None:
            pulumi.set(__self__, "service_attachment", service_attachment)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="endpointAttachmentId")
    def endpoint_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID to use for the endpoint attachment. ID must start with a lowercase letter followed by up to 31 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. The minimum length is 2.
        """
        return pulumi.get(self, "endpoint_attachment_id")

    @endpoint_attachment_id.setter
    def endpoint_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_attachment_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the endpoint attachment.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> Optional[pulumi.Input[str]]:
        """
        Format: projects/*/regions/*/serviceAttachments/*
        """
        return pulumi.get(self, "service_attachment")

    @service_attachment.setter
    def service_attachment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_attachment", value)


class EndpointAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_attachment_id: ID to use for the endpoint attachment. ID must start with a lowercase letter followed by up to 31 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. The minimum length is 2.
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] name: Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param EndpointAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointAttachmentArgs.__new__(EndpointAttachmentArgs)

            __props__.__dict__["endpoint_attachment_id"] = endpoint_attachment_id
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["service_attachment"] = service_attachment
            __props__.__dict__["connection_state"] = None
            __props__.__dict__["host"] = None
            __props__.__dict__["state"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["organization_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(EndpointAttachment, __self__).__init__(
            'google-native:apigee/v1:EndpointAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EndpointAttachment':
        """
        Get an existing EndpointAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointAttachmentArgs.__new__(EndpointAttachmentArgs)

        __props__.__dict__["connection_state"] = None
        __props__.__dict__["endpoint_attachment_id"] = None
        __props__.__dict__["host"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization_id"] = None
        __props__.__dict__["service_attachment"] = None
        __props__.__dict__["state"] = None
        return EndpointAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionState")
    def connection_state(self) -> pulumi.Output[str]:
        """
        State of the endpoint attachment connection to the service attachment.
        """
        return pulumi.get(self, "connection_state")

    @property
    @pulumi.getter(name="endpointAttachmentId")
    def endpoint_attachment_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID to use for the endpoint attachment. ID must start with a lowercase letter followed by up to 31 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. The minimum length is 2.
        """
        return pulumi.get(self, "endpoint_attachment_id")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Host that can be used in either the HTTP target endpoint directly or as the host in target server.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location of the endpoint attachment.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the endpoint attachment. Use the following structure in your request: `organizations/{org}/endpointAttachments/{endpoint_attachment}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> pulumi.Output[str]:
        """
        Format: projects/*/regions/*/serviceAttachments/*
        """
        return pulumi.get(self, "service_attachment")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the endpoint attachment. Values other than `ACTIVE` mean the resource is not ready to use.
        """
        return pulumi.get(self, "state")

