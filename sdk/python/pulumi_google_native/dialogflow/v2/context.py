# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ContextArgs', 'Context']

@pulumi.input_type
class ContextArgs:
    def __init__(__self__, *,
                 environment_id: pulumi.Input[str],
                 name: pulumi.Input[str],
                 session_id: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 lifespan_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Context resource.
        :param pulumi.Input[str] name: The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        :param pulumi.Input[int] lifespan_count: Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        """
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "session_id", session_id)
        pulumi.set(__self__, "user_id", user_id)
        if lifespan_count is not None:
            pulumi.set(__self__, "lifespan_count", lifespan_count)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "session_id")

    @session_id.setter
    def session_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "session_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="lifespanCount")
    def lifespan_count(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        """
        return pulumi.get(self, "lifespan_count")

    @lifespan_count.setter
    def lifespan_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lifespan_count", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Context(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 lifespan_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a context. If the specified context already exists, overrides the context.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] lifespan_count: Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        :param pulumi.Input[str] name: The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContextArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a context. If the specified context already exists, overrides the context.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ContextArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContextArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 lifespan_count: Optional[pulumi.Input[int]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContextArgs.__new__(ContextArgs)

            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["lifespan_count"] = lifespan_count
            __props__.__dict__["location"] = location
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["project"] = project
            if session_id is None and not opts.urn:
                raise TypeError("Missing required property 'session_id'")
            __props__.__dict__["session_id"] = session_id
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["environment_id", "location", "project", "session_id", "user_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Context, __self__).__init__(
            'google-native:dialogflow/v2:Context',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Context':
        """
        Get an existing Context resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ContextArgs.__new__(ContextArgs)

        __props__.__dict__["environment_id"] = None
        __props__.__dict__["lifespan_count"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["session_id"] = None
        __props__.__dict__["user_id"] = None
        return Context(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="lifespanCount")
    def lifespan_count(self) -> pulumi.Output[int]:
        """
        Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        """
        return pulumi.get(self, "lifespan_count")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "session_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_id")

