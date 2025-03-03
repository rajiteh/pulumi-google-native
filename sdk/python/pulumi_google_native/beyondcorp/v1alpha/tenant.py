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

__all__ = ['TenantArgs', 'Tenant']

@pulumi.input_type
class TenantArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 google_group_email: Optional[pulumi.Input[str]] = None,
                 google_group_id: Optional[pulumi.Input[str]] = None,
                 partner_metadata: Optional[pulumi.Input['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Tenant resource.
        :param pulumi.Input[str] display_name: Optional. An arbitrary caller-provided name for the Tenant. Cannot exceed 64 characters.
        :param pulumi.Input[str] google_group_email: Optional. Google group email to which the Tenant is enabled.
        :param pulumi.Input[str] google_group_id: Optional. Google group ID to which the Tenant is enabled.
        :param pulumi.Input['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs'] partner_metadata: Optional. Metadata provided by the Partner associated with Tenant.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        pulumi.set(__self__, "organization_id", organization_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if google_group_email is not None:
            pulumi.set(__self__, "google_group_email", google_group_email)
        if google_group_id is not None:
            pulumi.set(__self__, "google_group_id", google_group_id)
        if partner_metadata is not None:
            pulumi.set(__self__, "partner_metadata", partner_metadata)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. An arbitrary caller-provided name for the Tenant. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="googleGroupEmail")
    def google_group_email(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Google group email to which the Tenant is enabled.
        """
        return pulumi.get(self, "google_group_email")

    @google_group_email.setter
    def google_group_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "google_group_email", value)

    @property
    @pulumi.getter(name="googleGroupId")
    def google_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Google group ID to which the Tenant is enabled.
        """
        return pulumi.get(self, "google_group_id")

    @google_group_id.setter
    def google_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "google_group_id", value)

    @property
    @pulumi.getter(name="partnerMetadata")
    def partner_metadata(self) -> Optional[pulumi.Input['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']]:
        """
        Optional. Metadata provided by the Partner associated with Tenant.
        """
        return pulumi.get(self, "partner_metadata")

    @partner_metadata.setter
    def partner_metadata(self, value: Optional[pulumi.Input['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']]):
        pulumi.set(self, "partner_metadata", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class Tenant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 google_group_email: Optional[pulumi.Input[str]] = None,
                 google_group_id: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 partner_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new BeyondCorp Enterprise tenant in a given organization and can only be called by onboarded BeyondCorp Enterprise partner.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Optional. An arbitrary caller-provided name for the Tenant. Cannot exceed 64 characters.
        :param pulumi.Input[str] google_group_email: Optional. Google group email to which the Tenant is enabled.
        :param pulumi.Input[str] google_group_id: Optional. Google group ID to which the Tenant is enabled.
        :param pulumi.Input[pulumi.InputType['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']] partner_metadata: Optional. Metadata provided by the Partner associated with Tenant.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TenantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new BeyondCorp Enterprise tenant in a given organization and can only be called by onboarded BeyondCorp Enterprise partner.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param TenantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 google_group_email: Optional[pulumi.Input[str]] = None,
                 google_group_id: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 partner_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataArgs']]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TenantArgs.__new__(TenantArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["google_group_email"] = google_group_email
            __props__.__dict__["google_group_id"] = google_group_id
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["partner_metadata"] = partner_metadata
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["organization_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Tenant, __self__).__init__(
            'google-native:beyondcorp/v1alpha:Tenant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tenant':
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TenantArgs.__new__(TenantArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["google_group_email"] = None
        __props__.__dict__["google_group_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization_id"] = None
        __props__.__dict__["partner_metadata"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["update_time"] = None
        return Tenant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. An arbitrary caller-provided name for the Tenant. Cannot exceed 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="googleGroupEmail")
    def google_group_email(self) -> pulumi.Output[str]:
        """
        Optional. Google group email to which the Tenant is enabled.
        """
        return pulumi.get(self, "google_group_email")

    @property
    @pulumi.getter(name="googleGroupId")
    def google_group_id(self) -> pulumi.Output[str]:
        """
        Optional. Google group ID to which the Tenant is enabled.
        """
        return pulumi.get(self, "google_group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique resource name of the Tenant. The name is ignored when creating Tenant.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="partnerMetadata")
    def partner_metadata(self) -> pulumi.Output['outputs.GoogleCloudBeyondcorpPartnerservicesV1alphaPartnerMetadataResponse']:
        """
        Optional. Metadata provided by the Partner associated with Tenant.
        """
        return pulumi.get(self, "partner_metadata")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the resource was last modified.
        """
        return pulumi.get(self, "update_time")

