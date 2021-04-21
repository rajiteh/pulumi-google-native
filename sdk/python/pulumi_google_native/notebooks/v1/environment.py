# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 environments_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 container_image: Optional[pulumi.Input['ContainerImageArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 post_startup_script: Optional[pulumi.Input[str]] = None,
                 vm_image: Optional[pulumi.Input['VmImageArgs']] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input['ContainerImageArgs'] container_image: Use a container image to start the notebook instance.
        :param pulumi.Input[str] description: A brief description of this environment.
        :param pulumi.Input[str] display_name: Display name of this environment for the UI.
        :param pulumi.Input[str] post_startup_script: Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
        :param pulumi.Input['VmImageArgs'] vm_image: Use a Compute Engine VM image to start the notebook instance.
        """
        pulumi.set(__self__, "environments_id", environments_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if container_image is not None:
            pulumi.set(__self__, "container_image", container_image)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if post_startup_script is not None:
            pulumi.set(__self__, "post_startup_script", post_startup_script)
        if vm_image is not None:
            pulumi.set(__self__, "vm_image", vm_image)

    @property
    @pulumi.getter(name="environmentsId")
    def environments_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environments_id")

    @environments_id.setter
    def environments_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environments_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="containerImage")
    def container_image(self) -> Optional[pulumi.Input['ContainerImageArgs']]:
        """
        Use a container image to start the notebook instance.
        """
        return pulumi.get(self, "container_image")

    @container_image.setter
    def container_image(self, value: Optional[pulumi.Input['ContainerImageArgs']]):
        pulumi.set(self, "container_image", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A brief description of this environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of this environment for the UI.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="postStartupScript")
    def post_startup_script(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
        """
        return pulumi.get(self, "post_startup_script")

    @post_startup_script.setter
    def post_startup_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "post_startup_script", value)

    @property
    @pulumi.getter(name="vmImage")
    def vm_image(self) -> Optional[pulumi.Input['VmImageArgs']]:
        """
        Use a Compute Engine VM image to start the notebook instance.
        """
        return pulumi.get(self, "vm_image")

    @vm_image.setter
    def vm_image(self, value: Optional[pulumi.Input['VmImageArgs']]):
        pulumi.set(self, "vm_image", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_image: Optional[pulumi.Input[pulumi.InputType['ContainerImageArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 post_startup_script: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 vm_image: Optional[pulumi.Input[pulumi.InputType['VmImageArgs']]] = None,
                 __props__=None):
        """
        Creates a new Environment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ContainerImageArgs']] container_image: Use a container image to start the notebook instance.
        :param pulumi.Input[str] description: A brief description of this environment.
        :param pulumi.Input[str] display_name: Display name of this environment for the UI.
        :param pulumi.Input[str] post_startup_script: Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
        :param pulumi.Input[pulumi.InputType['VmImageArgs']] vm_image: Use a Compute Engine VM image to start the notebook instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Environment.

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_image: Optional[pulumi.Input[pulumi.InputType['ContainerImageArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 post_startup_script: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 vm_image: Optional[pulumi.Input[pulumi.InputType['VmImageArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["container_image"] = container_image
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if environments_id is None and not opts.urn:
                raise TypeError("Missing required property 'environments_id'")
            __props__.__dict__["environments_id"] = environments_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["post_startup_script"] = post_startup_script
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["vm_image"] = vm_image
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
        super(Environment, __self__).__init__(
            'google-native:notebooks/v1:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

        __props__.__dict__["container_image"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["post_startup_script"] = None
        __props__.__dict__["vm_image"] = None
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerImage")
    def container_image(self) -> pulumi.Output['outputs.ContainerImageResponse']:
        """
        Use a container image to start the notebook instance.
        """
        return pulumi.get(self, "container_image")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which this environment was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A brief description of this environment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name of this environment for the UI.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of this environment. Format: `projects/{project_id}/locations/{location}/environments/{environment_id}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="postStartupScript")
    def post_startup_script(self) -> pulumi.Output[str]:
        """
        Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
        """
        return pulumi.get(self, "post_startup_script")

    @property
    @pulumi.getter(name="vmImage")
    def vm_image(self) -> pulumi.Output['outputs.VmImageResponse']:
        """
        Use a Compute Engine VM image to start the notebook instance.
        """
        return pulumi.get(self, "vm_image")
