# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ManagedZoneArgs', 'ManagedZone']

@pulumi.input_type
class ManagedZoneArgs:
    def __init__(__self__, *,
                 dns: pulumi.Input[str],
                 managed_zone_id: pulumi.Input[str],
                 target_project: pulumi.Input[str],
                 target_vpc: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ManagedZone resource.
        :param pulumi.Input[str] dns: DNS Name of the resource
        :param pulumi.Input[str] managed_zone_id: Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        :param pulumi.Input[str] target_project: The name of the Target Project
        :param pulumi.Input[str] target_vpc: The name of the Target Project VPC Network
        :param pulumi.Input[str] description: Optional. Description of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        pulumi.set(__self__, "dns", dns)
        pulumi.set(__self__, "managed_zone_id", managed_zone_id)
        pulumi.set(__self__, "target_project", target_project)
        pulumi.set(__self__, "target_vpc", target_vpc)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def dns(self) -> pulumi.Input[str]:
        """
        DNS Name of the resource
        """
        return pulumi.get(self, "dns")

    @dns.setter
    def dns(self, value: pulumi.Input[str]):
        pulumi.set(self, "dns", value)

    @property
    @pulumi.getter(name="managedZoneId")
    def managed_zone_id(self) -> pulumi.Input[str]:
        """
        Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "managed_zone_id")

    @managed_zone_id.setter
    def managed_zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_zone_id", value)

    @property
    @pulumi.getter(name="targetProject")
    def target_project(self) -> pulumi.Input[str]:
        """
        The name of the Target Project
        """
        return pulumi.get(self, "target_project")

    @target_project.setter
    def target_project(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_project", value)

    @property
    @pulumi.getter(name="targetVpc")
    def target_vpc(self) -> pulumi.Input[str]:
        """
        The name of the Target Project VPC Network
        """
        return pulumi.get(self, "target_vpc")

    @target_vpc.setter
    def target_vpc(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_vpc", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class ManagedZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_zone_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 target_project: Optional[pulumi.Input[str]] = None,
                 target_vpc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new ManagedZone in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. Description of the resource.
        :param pulumi.Input[str] dns: DNS Name of the resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[str] managed_zone_id: Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        :param pulumi.Input[str] target_project: The name of the Target Project
        :param pulumi.Input[str] target_vpc: The name of the Target Project VPC Network
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ManagedZone in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ManagedZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_zone_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 target_project: Optional[pulumi.Input[str]] = None,
                 target_vpc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagedZoneArgs.__new__(ManagedZoneArgs)

            __props__.__dict__["description"] = description
            if dns is None and not opts.urn:
                raise TypeError("Missing required property 'dns'")
            __props__.__dict__["dns"] = dns
            __props__.__dict__["labels"] = labels
            if managed_zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'managed_zone_id'")
            __props__.__dict__["managed_zone_id"] = managed_zone_id
            __props__.__dict__["project"] = project
            if target_project is None and not opts.urn:
                raise TypeError("Missing required property 'target_project'")
            __props__.__dict__["target_project"] = target_project
            if target_vpc is None and not opts.urn:
                raise TypeError("Missing required property 'target_vpc'")
            __props__.__dict__["target_vpc"] = target_vpc
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["managed_zone_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ManagedZone, __self__).__init__(
            'google-native:connectors/v1:ManagedZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagedZone':
        """
        Get an existing ManagedZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ManagedZoneArgs.__new__(ManagedZoneArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["dns"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["managed_zone_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["target_project"] = None
        __props__.__dict__["target_vpc"] = None
        __props__.__dict__["update_time"] = None
        return ManagedZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dns(self) -> pulumi.Output[str]:
        """
        DNS Name of the resource
        """
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="managedZoneId")
    def managed_zone_id(self) -> pulumi.Output[str]:
        """
        Required. Identifier to assign to the ManagedZone. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "managed_zone_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the Managed Zone. Format: projects/{project}/locations/global/managedZones/{managed_zone}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="targetProject")
    def target_project(self) -> pulumi.Output[str]:
        """
        The name of the Target Project
        """
        return pulumi.get(self, "target_project")

    @property
    @pulumi.getter(name="targetVpc")
    def target_vpc(self) -> pulumi.Output[str]:
        """
        The name of the Target Project VPC Network
        """
        return pulumi.get(self, "target_vpc")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")
