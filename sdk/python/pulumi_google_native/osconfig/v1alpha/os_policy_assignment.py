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

__all__ = ['OsPolicyAssignmentArgs', 'OsPolicyAssignment']

@pulumi.input_type
class OsPolicyAssignmentArgs:
    def __init__(__self__, *,
                 instance_filter: pulumi.Input['OSPolicyAssignmentInstanceFilterArgs'],
                 os_policies: pulumi.Input[Sequence[pulumi.Input['OSPolicyArgs']]],
                 os_policy_assignment_id: pulumi.Input[str],
                 rollout: pulumi.Input['OSPolicyAssignmentRolloutArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OsPolicyAssignment resource.
        :param pulumi.Input['OSPolicyAssignmentInstanceFilterArgs'] instance_filter: Filter to select VMs.
        :param pulumi.Input[Sequence[pulumi.Input['OSPolicyArgs']]] os_policies: List of OS policies to be applied to the VMs.
        :param pulumi.Input[str] os_policy_assignment_id: Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input['OSPolicyAssignmentRolloutArgs'] rollout: Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
        :param pulumi.Input[str] description: OS policy assignment description. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[str] name: Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
        """
        pulumi.set(__self__, "instance_filter", instance_filter)
        pulumi.set(__self__, "os_policies", os_policies)
        pulumi.set(__self__, "os_policy_assignment_id", os_policy_assignment_id)
        pulumi.set(__self__, "rollout", rollout)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="instanceFilter")
    def instance_filter(self) -> pulumi.Input['OSPolicyAssignmentInstanceFilterArgs']:
        """
        Filter to select VMs.
        """
        return pulumi.get(self, "instance_filter")

    @instance_filter.setter
    def instance_filter(self, value: pulumi.Input['OSPolicyAssignmentInstanceFilterArgs']):
        pulumi.set(self, "instance_filter", value)

    @property
    @pulumi.getter(name="osPolicies")
    def os_policies(self) -> pulumi.Input[Sequence[pulumi.Input['OSPolicyArgs']]]:
        """
        List of OS policies to be applied to the VMs.
        """
        return pulumi.get(self, "os_policies")

    @os_policies.setter
    def os_policies(self, value: pulumi.Input[Sequence[pulumi.Input['OSPolicyArgs']]]):
        pulumi.set(self, "os_policies", value)

    @property
    @pulumi.getter(name="osPolicyAssignmentId")
    def os_policy_assignment_id(self) -> pulumi.Input[str]:
        """
        Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "os_policy_assignment_id")

    @os_policy_assignment_id.setter
    def os_policy_assignment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "os_policy_assignment_id", value)

    @property
    @pulumi.getter
    def rollout(self) -> pulumi.Input['OSPolicyAssignmentRolloutArgs']:
        """
        Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
        """
        return pulumi.get(self, "rollout")

    @rollout.setter
    def rollout(self, value: pulumi.Input['OSPolicyAssignmentRolloutArgs']):
        pulumi.set(self, "rollout", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        OS policy assignment description. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

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
        Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
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


class OsPolicyAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance_filter: Optional[pulumi.Input[pulumi.InputType['OSPolicyAssignmentInstanceFilterArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OSPolicyArgs']]]]] = None,
                 os_policy_assignment_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rollout: Optional[pulumi.Input[pulumi.InputType['OSPolicyAssignmentRolloutArgs']]] = None,
                 __props__=None):
        """
        Create an OS policy assignment. This method also creates the first revision of the OS policy assignment. This method returns a long running operation (LRO) that contains the rollout details. The rollout can be cancelled by cancelling the LRO. For more information, see [Method: projects.locations.osPolicyAssignments.operations.cancel](https://cloud.google.com/compute/docs/osconfig/rest/v1alpha/projects.locations.osPolicyAssignments.operations/cancel).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: OS policy assignment description. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[pulumi.InputType['OSPolicyAssignmentInstanceFilterArgs']] instance_filter: Filter to select VMs.
        :param pulumi.Input[str] name: Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OSPolicyArgs']]]] os_policies: List of OS policies to be applied to the VMs.
        :param pulumi.Input[str] os_policy_assignment_id: Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input[pulumi.InputType['OSPolicyAssignmentRolloutArgs']] rollout: Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OsPolicyAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create an OS policy assignment. This method also creates the first revision of the OS policy assignment. This method returns a long running operation (LRO) that contains the rollout details. The rollout can be cancelled by cancelling the LRO. For more information, see [Method: projects.locations.osPolicyAssignments.operations.cancel](https://cloud.google.com/compute/docs/osconfig/rest/v1alpha/projects.locations.osPolicyAssignments.operations/cancel).

        :param str resource_name: The name of the resource.
        :param OsPolicyAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OsPolicyAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance_filter: Optional[pulumi.Input[pulumi.InputType['OSPolicyAssignmentInstanceFilterArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OSPolicyArgs']]]]] = None,
                 os_policy_assignment_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rollout: Optional[pulumi.Input[pulumi.InputType['OSPolicyAssignmentRolloutArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OsPolicyAssignmentArgs.__new__(OsPolicyAssignmentArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            if instance_filter is None and not opts.urn:
                raise TypeError("Missing required property 'instance_filter'")
            __props__.__dict__["instance_filter"] = instance_filter
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if os_policies is None and not opts.urn:
                raise TypeError("Missing required property 'os_policies'")
            __props__.__dict__["os_policies"] = os_policies
            if os_policy_assignment_id is None and not opts.urn:
                raise TypeError("Missing required property 'os_policy_assignment_id'")
            __props__.__dict__["os_policy_assignment_id"] = os_policy_assignment_id
            __props__.__dict__["project"] = project
            if rollout is None and not opts.urn:
                raise TypeError("Missing required property 'rollout'")
            __props__.__dict__["rollout"] = rollout
            __props__.__dict__["baseline"] = None
            __props__.__dict__["deleted"] = None
            __props__.__dict__["reconciling"] = None
            __props__.__dict__["revision_create_time"] = None
            __props__.__dict__["revision_id"] = None
            __props__.__dict__["rollout_state"] = None
            __props__.__dict__["uid"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "os_policy_assignment_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(OsPolicyAssignment, __self__).__init__(
            'google-native:osconfig/v1alpha:OsPolicyAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OsPolicyAssignment':
        """
        Get an existing OsPolicyAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OsPolicyAssignmentArgs.__new__(OsPolicyAssignmentArgs)

        __props__.__dict__["baseline"] = None
        __props__.__dict__["deleted"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["instance_filter"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["os_policies"] = None
        __props__.__dict__["os_policy_assignment_id"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["reconciling"] = None
        __props__.__dict__["revision_create_time"] = None
        __props__.__dict__["revision_id"] = None
        __props__.__dict__["rollout"] = None
        __props__.__dict__["rollout_state"] = None
        __props__.__dict__["uid"] = None
        return OsPolicyAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def baseline(self) -> pulumi.Output[bool]:
        """
        Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
        """
        return pulumi.get(self, "baseline")

    @property
    @pulumi.getter
    def deleted(self) -> pulumi.Output[bool]:
        """
        Indicates that this revision deletes the OS policy assignment.
        """
        return pulumi.get(self, "deleted")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        OS policy assignment description. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="instanceFilter")
    def instance_filter(self) -> pulumi.Output['outputs.OSPolicyAssignmentInstanceFilterResponse']:
        """
        Filter to select VMs.
        """
        return pulumi.get(self, "instance_filter")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osPolicies")
    def os_policies(self) -> pulumi.Output[Sequence['outputs.OSPolicyResponse']]:
        """
        List of OS policies to be applied to the VMs.
        """
        return pulumi.get(self, "os_policies")

    @property
    @pulumi.getter(name="osPolicyAssignmentId")
    def os_policy_assignment_id(self) -> pulumi.Output[str]:
        """
        Required. The logical name of the OS policy assignment in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "os_policy_assignment_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def reconciling(self) -> pulumi.Output[bool]:
        """
        Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter(name="revisionCreateTime")
    def revision_create_time(self) -> pulumi.Output[str]:
        """
        The timestamp that the revision was created.
        """
        return pulumi.get(self, "revision_create_time")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> pulumi.Output[str]:
        """
        The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter
    def rollout(self) -> pulumi.Output['outputs.OSPolicyAssignmentRolloutResponse']:
        """
        Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
        """
        return pulumi.get(self, "rollout")

    @property
    @pulumi.getter(name="rolloutState")
    def rollout_state(self) -> pulumi.Output[str]:
        """
        OS policy assignment rollout state
        """
        return pulumi.get(self, "rollout_state")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Server generated unique id for the OS policy assignment resource.
        """
        return pulumi.get(self, "uid")

