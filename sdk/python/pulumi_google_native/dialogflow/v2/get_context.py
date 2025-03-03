# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetContextResult',
    'AwaitableGetContextResult',
    'get_context',
    'get_context_output',
]

@pulumi.output_type
class GetContextResult:
    def __init__(__self__, lifespan_count=None, name=None, parameters=None):
        if lifespan_count and not isinstance(lifespan_count, int):
            raise TypeError("Expected argument 'lifespan_count' to be a int")
        pulumi.set(__self__, "lifespan_count", lifespan_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="lifespanCount")
    def lifespan_count(self) -> int:
        """
        Optional. The number of conversational query requests after which the context expires. The default is `0`. If set to `0`, the context expires immediately. Contexts expire automatically after 20 minutes if there are no matching queries.
        """
        return pulumi.get(self, "lifespan_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of the context. Format: `projects//agent/sessions//contexts/`, or `projects//agent/environments//users//sessions//contexts/`. The `Context ID` is always converted to lowercase, may only contain characters in a-zA-Z0-9_-% and may be at most 250 bytes long. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. The following context names are reserved for internal use by Dialogflow. You should not use these contexts or create contexts with these names: * `__system_counters__` * `*_id_dialog_context` * `*_dialog_params_size`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Mapping[str, str]:
        """
        Optional. The collection of parameters associated with this context. Depending on your protocol or client library language, this is a map, associative array, symbol table, dictionary, or JSON object composed of a collection of (MapKey, MapValue) pairs: - MapKey type: string - MapKey value: parameter name - MapValue type: - If parameter's entity type is a composite entity: map - Else: depending on parameter value type, could be one of string, number, boolean, null, list or map - MapValue value: - If parameter's entity type is a composite entity: map from composite entity property names to property values - Else: parameter value
        """
        return pulumi.get(self, "parameters")


class AwaitableGetContextResult(GetContextResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContextResult(
            lifespan_count=self.lifespan_count,
            name=self.name,
            parameters=self.parameters)


def get_context(context_id: Optional[str] = None,
                environment_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                session_id: Optional[str] = None,
                user_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContextResult:
    """
    Retrieves the specified context.
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['environmentId'] = environment_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['sessionId'] = session_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v2:getContext', __args__, opts=opts, typ=GetContextResult).value

    return AwaitableGetContextResult(
        lifespan_count=__ret__.lifespan_count,
        name=__ret__.name,
        parameters=__ret__.parameters)


@_utilities.lift_output_func(get_context)
def get_context_output(context_id: Optional[pulumi.Input[str]] = None,
                       environment_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       session_id: Optional[pulumi.Input[str]] = None,
                       user_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContextResult]:
    """
    Retrieves the specified context.
    """
    ...
