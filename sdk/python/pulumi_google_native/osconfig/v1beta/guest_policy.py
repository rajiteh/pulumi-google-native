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

__all__ = ['GuestPolicyArgs', 'GuestPolicy']

@pulumi.input_type
class GuestPolicyArgs:
    def __init__(__self__, *,
                 assignment: pulumi.Input['AssignmentArgs'],
                 guest_policy_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_repositories: Optional[pulumi.Input[Sequence[pulumi.Input['PackageRepositoryArgs']]]] = None,
                 packages: Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recipes: Optional[pulumi.Input[Sequence[pulumi.Input['SoftwareRecipeArgs']]]] = None):
        """
        The set of arguments for constructing a GuestPolicy resource.
        :param pulumi.Input['AssignmentArgs'] assignment: Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        :param pulumi.Input[str] guest_policy_id: Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input[str] description: Description of the guest policy. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this guest policy. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[str] name: Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        :param pulumi.Input[Sequence[pulumi.Input['PackageRepositoryArgs']]] package_repositories: A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        :param pulumi.Input[Sequence[pulumi.Input['PackageArgs']]] packages: The software packages to be managed by this policy.
        :param pulumi.Input[Sequence[pulumi.Input['SoftwareRecipeArgs']]] recipes: A list of Recipes to install on the VM instance.
        """
        pulumi.set(__self__, "assignment", assignment)
        pulumi.set(__self__, "guest_policy_id", guest_policy_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if package_repositories is not None:
            pulumi.set(__self__, "package_repositories", package_repositories)
        if packages is not None:
            pulumi.set(__self__, "packages", packages)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if recipes is not None:
            pulumi.set(__self__, "recipes", recipes)

    @property
    @pulumi.getter
    def assignment(self) -> pulumi.Input['AssignmentArgs']:
        """
        Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        """
        return pulumi.get(self, "assignment")

    @assignment.setter
    def assignment(self, value: pulumi.Input['AssignmentArgs']):
        pulumi.set(self, "assignment", value)

    @property
    @pulumi.getter(name="guestPolicyId")
    def guest_policy_id(self) -> pulumi.Input[str]:
        """
        Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "guest_policy_id")

    @guest_policy_id.setter
    def guest_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "guest_policy_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the guest policy. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        The etag for this guest policy. If this is provided on update, it must match the server's etag.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="packageRepositories")
    def package_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PackageRepositoryArgs']]]]:
        """
        A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        """
        return pulumi.get(self, "package_repositories")

    @package_repositories.setter
    def package_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PackageRepositoryArgs']]]]):
        pulumi.set(self, "package_repositories", value)

    @property
    @pulumi.getter
    def packages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]]:
        """
        The software packages to be managed by this policy.
        """
        return pulumi.get(self, "packages")

    @packages.setter
    def packages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PackageArgs']]]]):
        pulumi.set(self, "packages", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def recipes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SoftwareRecipeArgs']]]]:
        """
        A list of Recipes to install on the VM instance.
        """
        return pulumi.get(self, "recipes")

    @recipes.setter
    def recipes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SoftwareRecipeArgs']]]]):
        pulumi.set(self, "recipes", value)


class GuestPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignment: Optional[pulumi.Input[pulumi.InputType['AssignmentArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 guest_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageRepositoryArgs']]]]] = None,
                 packages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recipes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SoftwareRecipeArgs']]]]] = None,
                 __props__=None):
        """
        Create an OS Config guest policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AssignmentArgs']] assignment: Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        :param pulumi.Input[str] description: Description of the guest policy. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] etag: The etag for this guest policy. If this is provided on update, it must match the server's etag.
        :param pulumi.Input[str] guest_policy_id: Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input[str] name: Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageRepositoryArgs']]]] package_repositories: A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageArgs']]]] packages: The software packages to be managed by this policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SoftwareRecipeArgs']]]] recipes: A list of Recipes to install on the VM instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GuestPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create an OS Config guest policy.

        :param str resource_name: The name of the resource.
        :param GuestPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GuestPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignment: Optional[pulumi.Input[pulumi.InputType['AssignmentArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 guest_policy_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageRepositoryArgs']]]]] = None,
                 packages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PackageArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recipes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SoftwareRecipeArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GuestPolicyArgs.__new__(GuestPolicyArgs)

            if assignment is None and not opts.urn:
                raise TypeError("Missing required property 'assignment'")
            __props__.__dict__["assignment"] = assignment
            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            if guest_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'guest_policy_id'")
            __props__.__dict__["guest_policy_id"] = guest_policy_id
            __props__.__dict__["name"] = name
            __props__.__dict__["package_repositories"] = package_repositories
            __props__.__dict__["packages"] = packages
            __props__.__dict__["project"] = project
            __props__.__dict__["recipes"] = recipes
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["guest_policy_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(GuestPolicy, __self__).__init__(
            'google-native:osconfig/v1beta:GuestPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GuestPolicy':
        """
        Get an existing GuestPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GuestPolicyArgs.__new__(GuestPolicyArgs)

        __props__.__dict__["assignment"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["guest_policy_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["package_repositories"] = None
        __props__.__dict__["packages"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["recipes"] = None
        __props__.__dict__["update_time"] = None
        return GuestPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def assignment(self) -> pulumi.Output['outputs.AssignmentResponse']:
        """
        Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        """
        return pulumi.get(self, "assignment")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time this guest policy was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the guest policy. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The etag for this guest policy. If this is provided on update, it must match the server's etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="guestPolicyId")
    def guest_policy_id(self) -> pulumi.Output[str]:
        """
        Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "guest_policy_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageRepositories")
    def package_repositories(self) -> pulumi.Output[Sequence['outputs.PackageRepositoryResponse']]:
        """
        A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        """
        return pulumi.get(self, "package_repositories")

    @property
    @pulumi.getter
    def packages(self) -> pulumi.Output[Sequence['outputs.PackageResponse']]:
        """
        The software packages to be managed by this policy.
        """
        return pulumi.get(self, "packages")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def recipes(self) -> pulumi.Output[Sequence['outputs.SoftwareRecipeResponse']]:
        """
        A list of Recipes to install on the VM instance.
        """
        return pulumi.get(self, "recipes")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last time this guest policy was updated.
        """
        return pulumi.get(self, "update_time")

