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

__all__ = ['PatchDeploymentArgs', 'PatchDeployment']

@pulumi.input_type
class PatchDeploymentArgs:
    def __init__(__self__, *,
                 instance_filter: pulumi.Input['PatchInstanceFilterArgs'],
                 one_time_schedule: pulumi.Input['OneTimeScheduleArgs'],
                 patch_deployment_id: pulumi.Input[str],
                 recurring_schedule: pulumi.Input['RecurringScheduleArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 patch_config: Optional[pulumi.Input['PatchConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rollout: Optional[pulumi.Input['PatchRolloutArgs']] = None):
        """
        The set of arguments for constructing a PatchDeployment resource.
        :param pulumi.Input['PatchInstanceFilterArgs'] instance_filter: VM instances to patch.
        :param pulumi.Input['OneTimeScheduleArgs'] one_time_schedule: Schedule a one-time execution.
        :param pulumi.Input[str] patch_deployment_id: Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input['RecurringScheduleArgs'] recurring_schedule: Schedule recurring executions.
        :param pulumi.Input[str] description: Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] duration: Optional. Duration of the patch. After the duration ends, the patch times out.
        :param pulumi.Input[str] name: Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        :param pulumi.Input['PatchConfigArgs'] patch_config: Optional. Patch configuration that is applied.
        :param pulumi.Input['PatchRolloutArgs'] rollout: Optional. Rollout strategy of the patch job.
        """
        pulumi.set(__self__, "instance_filter", instance_filter)
        pulumi.set(__self__, "one_time_schedule", one_time_schedule)
        pulumi.set(__self__, "patch_deployment_id", patch_deployment_id)
        pulumi.set(__self__, "recurring_schedule", recurring_schedule)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if patch_config is not None:
            pulumi.set(__self__, "patch_config", patch_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rollout is not None:
            pulumi.set(__self__, "rollout", rollout)

    @property
    @pulumi.getter(name="instanceFilter")
    def instance_filter(self) -> pulumi.Input['PatchInstanceFilterArgs']:
        """
        VM instances to patch.
        """
        return pulumi.get(self, "instance_filter")

    @instance_filter.setter
    def instance_filter(self, value: pulumi.Input['PatchInstanceFilterArgs']):
        pulumi.set(self, "instance_filter", value)

    @property
    @pulumi.getter(name="oneTimeSchedule")
    def one_time_schedule(self) -> pulumi.Input['OneTimeScheduleArgs']:
        """
        Schedule a one-time execution.
        """
        return pulumi.get(self, "one_time_schedule")

    @one_time_schedule.setter
    def one_time_schedule(self, value: pulumi.Input['OneTimeScheduleArgs']):
        pulumi.set(self, "one_time_schedule", value)

    @property
    @pulumi.getter(name="patchDeploymentId")
    def patch_deployment_id(self) -> pulumi.Input[str]:
        """
        Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "patch_deployment_id")

    @patch_deployment_id.setter
    def patch_deployment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "patch_deployment_id", value)

    @property
    @pulumi.getter(name="recurringSchedule")
    def recurring_schedule(self) -> pulumi.Input['RecurringScheduleArgs']:
        """
        Schedule recurring executions.
        """
        return pulumi.get(self, "recurring_schedule")

    @recurring_schedule.setter
    def recurring_schedule(self, value: pulumi.Input['RecurringScheduleArgs']):
        pulumi.set(self, "recurring_schedule", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Duration of the patch. After the duration ends, the patch times out.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="patchConfig")
    def patch_config(self) -> Optional[pulumi.Input['PatchConfigArgs']]:
        """
        Optional. Patch configuration that is applied.
        """
        return pulumi.get(self, "patch_config")

    @patch_config.setter
    def patch_config(self, value: Optional[pulumi.Input['PatchConfigArgs']]):
        pulumi.set(self, "patch_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def rollout(self) -> Optional[pulumi.Input['PatchRolloutArgs']]:
        """
        Optional. Rollout strategy of the patch job.
        """
        return pulumi.get(self, "rollout")

    @rollout.setter
    def rollout(self, value: Optional[pulumi.Input['PatchRolloutArgs']]):
        pulumi.set(self, "rollout", value)


class PatchDeployment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 instance_filter: Optional[pulumi.Input[pulumi.InputType['PatchInstanceFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 one_time_schedule: Optional[pulumi.Input[pulumi.InputType['OneTimeScheduleArgs']]] = None,
                 patch_config: Optional[pulumi.Input[pulumi.InputType['PatchConfigArgs']]] = None,
                 patch_deployment_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recurring_schedule: Optional[pulumi.Input[pulumi.InputType['RecurringScheduleArgs']]] = None,
                 rollout: Optional[pulumi.Input[pulumi.InputType['PatchRolloutArgs']]] = None,
                 __props__=None):
        """
        Create an OS Config patch deployment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        :param pulumi.Input[str] duration: Optional. Duration of the patch. After the duration ends, the patch times out.
        :param pulumi.Input[pulumi.InputType['PatchInstanceFilterArgs']] instance_filter: VM instances to patch.
        :param pulumi.Input[str] name: Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        :param pulumi.Input[pulumi.InputType['OneTimeScheduleArgs']] one_time_schedule: Schedule a one-time execution.
        :param pulumi.Input[pulumi.InputType['PatchConfigArgs']] patch_config: Optional. Patch configuration that is applied.
        :param pulumi.Input[str] patch_deployment_id: Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        :param pulumi.Input[pulumi.InputType['RecurringScheduleArgs']] recurring_schedule: Schedule recurring executions.
        :param pulumi.Input[pulumi.InputType['PatchRolloutArgs']] rollout: Optional. Rollout strategy of the patch job.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PatchDeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create an OS Config patch deployment.

        :param str resource_name: The name of the resource.
        :param PatchDeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PatchDeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 instance_filter: Optional[pulumi.Input[pulumi.InputType['PatchInstanceFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 one_time_schedule: Optional[pulumi.Input[pulumi.InputType['OneTimeScheduleArgs']]] = None,
                 patch_config: Optional[pulumi.Input[pulumi.InputType['PatchConfigArgs']]] = None,
                 patch_deployment_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recurring_schedule: Optional[pulumi.Input[pulumi.InputType['RecurringScheduleArgs']]] = None,
                 rollout: Optional[pulumi.Input[pulumi.InputType['PatchRolloutArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PatchDeploymentArgs.__new__(PatchDeploymentArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["duration"] = duration
            if instance_filter is None and not opts.urn:
                raise TypeError("Missing required property 'instance_filter'")
            __props__.__dict__["instance_filter"] = instance_filter
            __props__.__dict__["name"] = name
            if one_time_schedule is None and not opts.urn:
                raise TypeError("Missing required property 'one_time_schedule'")
            __props__.__dict__["one_time_schedule"] = one_time_schedule
            __props__.__dict__["patch_config"] = patch_config
            if patch_deployment_id is None and not opts.urn:
                raise TypeError("Missing required property 'patch_deployment_id'")
            __props__.__dict__["patch_deployment_id"] = patch_deployment_id
            __props__.__dict__["project"] = project
            if recurring_schedule is None and not opts.urn:
                raise TypeError("Missing required property 'recurring_schedule'")
            __props__.__dict__["recurring_schedule"] = recurring_schedule
            __props__.__dict__["rollout"] = rollout
            __props__.__dict__["create_time"] = None
            __props__.__dict__["last_execute_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["patch_deployment_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(PatchDeployment, __self__).__init__(
            'google-native:osconfig/v1beta:PatchDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PatchDeployment':
        """
        Get an existing PatchDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PatchDeploymentArgs.__new__(PatchDeploymentArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["duration"] = None
        __props__.__dict__["instance_filter"] = None
        __props__.__dict__["last_execute_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["one_time_schedule"] = None
        __props__.__dict__["patch_config"] = None
        __props__.__dict__["patch_deployment_id"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["recurring_schedule"] = None
        __props__.__dict__["rollout"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return PatchDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[str]:
        """
        Optional. Duration of the patch. After the duration ends, the patch times out.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="instanceFilter")
    def instance_filter(self) -> pulumi.Output['outputs.PatchInstanceFilterResponse']:
        """
        VM instances to patch.
        """
        return pulumi.get(self, "instance_filter")

    @property
    @pulumi.getter(name="lastExecuteTime")
    def last_execute_time(self) -> pulumi.Output[str]:
        """
        The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "last_execute_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oneTimeSchedule")
    def one_time_schedule(self) -> pulumi.Output['outputs.OneTimeScheduleResponse']:
        """
        Schedule a one-time execution.
        """
        return pulumi.get(self, "one_time_schedule")

    @property
    @pulumi.getter(name="patchConfig")
    def patch_config(self) -> pulumi.Output['outputs.PatchConfigResponse']:
        """
        Optional. Patch configuration that is applied.
        """
        return pulumi.get(self, "patch_config")

    @property
    @pulumi.getter(name="patchDeploymentId")
    def patch_deployment_id(self) -> pulumi.Output[str]:
        """
        Required. A name for the patch deployment in the project. When creating a name the following rules apply: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        """
        return pulumi.get(self, "patch_deployment_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="recurringSchedule")
    def recurring_schedule(self) -> pulumi.Output['outputs.RecurringScheduleResponse']:
        """
        Schedule recurring executions.
        """
        return pulumi.get(self, "recurring_schedule")

    @property
    @pulumi.getter
    def rollout(self) -> pulumi.Output['outputs.PatchRolloutResponse']:
        """
        Optional. Rollout strategy of the patch job.
        """
        return pulumi.get(self, "rollout")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the patch deployment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        """
        return pulumi.get(self, "update_time")

