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

__all__ = [
    'GetExecutionResult',
    'AwaitableGetExecutionResult',
    'get_execution',
    'get_execution_output',
]

@pulumi.output_type
class GetExecutionResult:
    def __init__(__self__, argument=None, call_log_level=None, end_time=None, error=None, name=None, result=None, start_time=None, state=None, status=None, workflow_revision_id=None):
        if argument and not isinstance(argument, str):
            raise TypeError("Expected argument 'argument' to be a str")
        pulumi.set(__self__, "argument", argument)
        if call_log_level and not isinstance(call_log_level, str):
            raise TypeError("Expected argument 'call_log_level' to be a str")
        pulumi.set(__self__, "call_log_level", call_log_level)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if workflow_revision_id and not isinstance(workflow_revision_id, str):
            raise TypeError("Expected argument 'workflow_revision_id' to be a str")
        pulumi.set(__self__, "workflow_revision_id", workflow_revision_id)

    @property
    @pulumi.getter
    def argument(self) -> str:
        """
        Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\\"firstName\\":\\"FIRST\\",\\"lastName\\":\\"LAST\\"}"}'`
        """
        return pulumi.get(self, "argument")

    @property
    @pulumi.getter(name="callLogLevel")
    def call_log_level(self) -> str:
        """
        The call logging level associated to this execution.
        """
        return pulumi.get(self, "call_log_level")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Marks the end of execution, successful or not.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.ErrorResponse':
        """
        The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def result(self) -> str:
        """
        Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Marks the beginning of execution.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the execution.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.StatusResponse':
        """
        Status tracks the current steps and progress data of this execution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="workflowRevisionId")
    def workflow_revision_id(self) -> str:
        """
        Revision of the workflow this execution is using.
        """
        return pulumi.get(self, "workflow_revision_id")


class AwaitableGetExecutionResult(GetExecutionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExecutionResult(
            argument=self.argument,
            call_log_level=self.call_log_level,
            end_time=self.end_time,
            error=self.error,
            name=self.name,
            result=self.result,
            start_time=self.start_time,
            state=self.state,
            status=self.status,
            workflow_revision_id=self.workflow_revision_id)


def get_execution(execution_id: Optional[str] = None,
                  location: Optional[str] = None,
                  project: Optional[str] = None,
                  view: Optional[str] = None,
                  workflow_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExecutionResult:
    """
    Returns an execution of the given name.
    """
    __args__ = dict()
    __args__['executionId'] = execution_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    __args__['workflowId'] = workflow_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:workflowexecutions/v1beta:getExecution', __args__, opts=opts, typ=GetExecutionResult).value

    return AwaitableGetExecutionResult(
        argument=__ret__.argument,
        call_log_level=__ret__.call_log_level,
        end_time=__ret__.end_time,
        error=__ret__.error,
        name=__ret__.name,
        result=__ret__.result,
        start_time=__ret__.start_time,
        state=__ret__.state,
        status=__ret__.status,
        workflow_revision_id=__ret__.workflow_revision_id)


@_utilities.lift_output_func(get_execution)
def get_execution_output(execution_id: Optional[pulumi.Input[str]] = None,
                         location: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         view: Optional[pulumi.Input[Optional[str]]] = None,
                         workflow_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetExecutionResult]:
    """
    Returns an execution of the given name.
    """
    ...
