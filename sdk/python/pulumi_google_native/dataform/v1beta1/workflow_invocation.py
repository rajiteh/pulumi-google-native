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

__all__ = ['WorkflowInvocationArgs', 'WorkflowInvocation']

@pulumi.input_type
class WorkflowInvocationArgs:
    def __init__(__self__, *,
                 repository_id: pulumi.Input[str],
                 compilation_result: Optional[pulumi.Input[str]] = None,
                 invocation_config: Optional[pulumi.Input['InvocationConfigArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workflow_config: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WorkflowInvocation resource.
        :param pulumi.Input[str] compilation_result: Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
        :param pulumi.Input['InvocationConfigArgs'] invocation_config: Immutable. If left unset, a default InvocationConfig will be used.
        :param pulumi.Input[str] workflow_config: Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
        """
        pulumi.set(__self__, "repository_id", repository_id)
        if compilation_result is not None:
            pulumi.set(__self__, "compilation_result", compilation_result)
        if invocation_config is not None:
            pulumi.set(__self__, "invocation_config", invocation_config)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if workflow_config is not None:
            pulumi.set(__self__, "workflow_config", workflow_config)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="compilationResult")
    def compilation_result(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
        """
        return pulumi.get(self, "compilation_result")

    @compilation_result.setter
    def compilation_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compilation_result", value)

    @property
    @pulumi.getter(name="invocationConfig")
    def invocation_config(self) -> Optional[pulumi.Input['InvocationConfigArgs']]:
        """
        Immutable. If left unset, a default InvocationConfig will be used.
        """
        return pulumi.get(self, "invocation_config")

    @invocation_config.setter
    def invocation_config(self, value: Optional[pulumi.Input['InvocationConfigArgs']]):
        pulumi.set(self, "invocation_config", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="workflowConfig")
    def workflow_config(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
        """
        return pulumi.get(self, "workflow_config")

    @workflow_config.setter
    def workflow_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workflow_config", value)


class WorkflowInvocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compilation_result: Optional[pulumi.Input[str]] = None,
                 invocation_config: Optional[pulumi.Input[pulumi.InputType['InvocationConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 workflow_config: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new WorkflowInvocation in a given Repository.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compilation_result: Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
        :param pulumi.Input[pulumi.InputType['InvocationConfigArgs']] invocation_config: Immutable. If left unset, a default InvocationConfig will be used.
        :param pulumi.Input[str] workflow_config: Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkflowInvocationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new WorkflowInvocation in a given Repository.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param WorkflowInvocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkflowInvocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compilation_result: Optional[pulumi.Input[str]] = None,
                 invocation_config: Optional[pulumi.Input[pulumi.InputType['InvocationConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 workflow_config: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkflowInvocationArgs.__new__(WorkflowInvocationArgs)

            __props__.__dict__["compilation_result"] = compilation_result
            __props__.__dict__["invocation_config"] = invocation_config
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            if repository_id is None and not opts.urn:
                raise TypeError("Missing required property 'repository_id'")
            __props__.__dict__["repository_id"] = repository_id
            __props__.__dict__["workflow_config"] = workflow_config
            __props__.__dict__["invocation_timing"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project", "repository_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(WorkflowInvocation, __self__).__init__(
            'google-native:dataform/v1beta1:WorkflowInvocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkflowInvocation':
        """
        Get an existing WorkflowInvocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkflowInvocationArgs.__new__(WorkflowInvocationArgs)

        __props__.__dict__["compilation_result"] = None
        __props__.__dict__["invocation_config"] = None
        __props__.__dict__["invocation_timing"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["repository_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["workflow_config"] = None
        return WorkflowInvocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compilationResult")
    def compilation_result(self) -> pulumi.Output[str]:
        """
        Immutable. The name of the compilation result to compile. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
        """
        return pulumi.get(self, "compilation_result")

    @property
    @pulumi.getter(name="invocationConfig")
    def invocation_config(self) -> pulumi.Output['outputs.InvocationConfigResponse']:
        """
        Immutable. If left unset, a default InvocationConfig will be used.
        """
        return pulumi.get(self, "invocation_config")

    @property
    @pulumi.getter(name="invocationTiming")
    def invocation_timing(self) -> pulumi.Output['outputs.IntervalResponse']:
        """
        This workflow invocation's timing details.
        """
        return pulumi.get(self, "invocation_timing")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The workflow invocation's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        This workflow invocation's current state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workflowConfig")
    def workflow_config(self) -> pulumi.Output[str]:
        """
        Immutable. The name of the workflow config to invoke. Must be in the format `projects/*/locations/*/repositories/*/workflowConfigs/*`.
        """
        return pulumi.get(self, "workflow_config")
