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
from ._enums import *
from ._inputs import *

__all__ = ['RegionCommitmentArgs', 'RegionCommitment']

@pulumi.input_type
class RegionCommitmentArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input['LicenseResourceCommitmentArgs']] = None,
                 merge_source_commitments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]] = None,
                 split_source_commitment: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None):
        """
        The set of arguments for constructing a RegionCommitment resource.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        :param pulumi.Input['RegionCommitmentCategory'] category: The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input['LicenseResourceCommitmentArgs'] license_resource: The license specification required as part of a license commitment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] merge_source_commitments: List of source commitments to be merged into a new commitment.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input['RegionCommitmentPlan'] plan: The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]] reservations: List of reservations in this commitment.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]] resources: A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        :param pulumi.Input[str] split_source_commitment: Source commitment to be split into a new commitment.
        :param pulumi.Input['RegionCommitmentType'] type: The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        pulumi.set(__self__, "region", region)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if license_resource is not None:
            pulumi.set(__self__, "license_resource", license_resource)
        if merge_source_commitments is not None:
            pulumi.set(__self__, "merge_source_commitments", merge_source_commitments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if reservations is not None:
            pulumi.set(__self__, "reservations", reservations)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if split_source_commitment is not None:
            pulumi.set(__self__, "split_source_commitment", split_source_commitment)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input['RegionCommitmentCategory']]:
        """
        The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input['RegionCommitmentCategory']]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="licenseResource")
    def license_resource(self) -> Optional[pulumi.Input['LicenseResourceCommitmentArgs']]:
        """
        The license specification required as part of a license commitment.
        """
        return pulumi.get(self, "license_resource")

    @license_resource.setter
    def license_resource(self, value: Optional[pulumi.Input['LicenseResourceCommitmentArgs']]):
        pulumi.set(self, "license_resource", value)

    @property
    @pulumi.getter(name="mergeSourceCommitments")
    def merge_source_commitments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of source commitments to be merged into a new commitment.
        """
        return pulumi.get(self, "merge_source_commitments")

    @merge_source_commitments.setter
    def merge_source_commitments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "merge_source_commitments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['RegionCommitmentPlan']]:
        """
        The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['RegionCommitmentPlan']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def reservations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]]:
        """
        List of reservations in this commitment.
        """
        return pulumi.get(self, "reservations")

    @reservations.setter
    def reservations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReservationArgs']]]]):
        pulumi.set(self, "reservations", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]]:
        """
        A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceCommitmentArgs']]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter(name="splitSourceCommitment")
    def split_source_commitment(self) -> Optional[pulumi.Input[str]]:
        """
        Source commitment to be split into a new commitment.
        """
        return pulumi.get(self, "split_source_commitment")

    @split_source_commitment.setter
    def split_source_commitment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "split_source_commitment", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['RegionCommitmentType']]:
        """
        The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['RegionCommitmentType']]):
        pulumi.set(self, "type", value)


class RegionCommitment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']]] = None,
                 merge_source_commitments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]]] = None,
                 split_source_commitment: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None,
                 __props__=None):
        """
        Creates a commitment in the specified project using the data included in the request.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        :param pulumi.Input['RegionCommitmentCategory'] category: The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']] license_resource: The license specification required as part of a license commitment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] merge_source_commitments: List of source commitments to be merged into a new commitment.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input['RegionCommitmentPlan'] plan: The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]] reservations: List of reservations in this commitment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]] resources: A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        :param pulumi.Input[str] split_source_commitment: Source commitment to be split into a new commitment.
        :param pulumi.Input['RegionCommitmentType'] type: The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionCommitmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a commitment in the specified project using the data included in the request.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param RegionCommitmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionCommitmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 category: Optional[pulumi.Input['RegionCommitmentCategory']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 license_resource: Optional[pulumi.Input[pulumi.InputType['LicenseResourceCommitmentArgs']]] = None,
                 merge_source_commitments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['RegionCommitmentPlan']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reservations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReservationArgs']]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceCommitmentArgs']]]]] = None,
                 split_source_commitment: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['RegionCommitmentType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegionCommitmentArgs.__new__(RegionCommitmentArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            __props__.__dict__["category"] = category
            __props__.__dict__["description"] = description
            __props__.__dict__["license_resource"] = license_resource
            __props__.__dict__["merge_source_commitments"] = merge_source_commitments
            __props__.__dict__["name"] = name
            __props__.__dict__["plan"] = plan
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["reservations"] = reservations
            __props__.__dict__["resources"] = resources
            __props__.__dict__["split_source_commitment"] = split_source_commitment
            __props__.__dict__["type"] = type
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["end_timestamp"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["start_timestamp"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_message"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project", "region"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(RegionCommitment, __self__).__init__(
            'google-native:compute/beta:RegionCommitment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionCommitment':
        """
        Get an existing RegionCommitment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionCommitmentArgs.__new__(RegionCommitmentArgs)

        __props__.__dict__["auto_renew"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["end_timestamp"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["license_resource"] = None
        __props__.__dict__["merge_source_commitments"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["plan"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["reservations"] = None
        __props__.__dict__["resources"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["split_source_commitment"] = None
        __props__.__dict__["start_timestamp"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["type"] = None
        return RegionCommitment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[bool]:
        """
        Specifies whether to enable automatic renewal for the commitment. The default value is false if not specified. The field can be updated until the day of the commitment expiration at 12:00am PST. If the field is set to true, the commitment will be automatically renewed for either one or three years according to the terms of the existing commitment.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        The category of the commitment. Category MACHINE specifies commitments composed of machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE specifies commitments composed of software licenses, listed in licenseResources. Note that only MACHINE commitments should have a Type specified.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTimestamp")
    def end_timestamp(self) -> pulumi.Output[str]:
        """
        Commitment end time in RFC3339 text format.
        """
        return pulumi.get(self, "end_timestamp")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#commitment for commitments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="licenseResource")
    def license_resource(self) -> pulumi.Output['outputs.LicenseResourceCommitmentResponse']:
        """
        The license specification required as part of a license commitment.
        """
        return pulumi.get(self, "license_resource")

    @property
    @pulumi.getter(name="mergeSourceCommitments")
    def merge_source_commitments(self) -> pulumi.Output[Sequence[str]]:
        """
        List of source commitments to be merged into a new commitment.
        """
        return pulumi.get(self, "merge_source_commitments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[str]:
        """
        The plan for this commitment, which determines duration and discount rate. The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years).
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def reservations(self) -> pulumi.Output[Sequence['outputs.ReservationResponse']]:
        """
        List of reservations in this commitment.
        """
        return pulumi.get(self, "reservations")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Sequence['outputs.ResourceCommitmentResponse']]:
        """
        A list of commitment amounts for particular resources. Note that VCPU and MEMORY resource commitments must occur together.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="splitSourceCommitment")
    def split_source_commitment(self) -> pulumi.Output[str]:
        """
        Source commitment to be split into a new commitment.
        """
        return pulumi.get(self, "split_source_commitment")

    @property
    @pulumi.getter(name="startTimestamp")
    def start_timestamp(self) -> pulumi.Output[str]:
        """
        Commitment start time in RFC3339 text format.
        """
        return pulumi.get(self, "start_timestamp")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the commitment with regards to eventual expiration (each commitment has an end date defined). One of the following values: NOT_YET_ACTIVE, ACTIVE, EXPIRED.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        """
        An optional, human-readable explanation of the status.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of commitment, which affects the discount rate and the eligible resources. Type MEMORY_OPTIMIZED specifies a commitment that will only apply to memory optimized machines. Type ACCELERATOR_OPTIMIZED specifies a commitment that will only apply to accelerator optimized machines.
        """
        return pulumi.get(self, "type")

