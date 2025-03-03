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
    'GetTaskResult',
    'AwaitableGetTaskResult',
    'get_task',
    'get_task_output',
]

@pulumi.output_type
class GetTaskResult:
    def __init__(__self__, app_engine_http_request=None, create_time=None, dispatch_count=None, dispatch_deadline=None, first_attempt=None, http_request=None, last_attempt=None, name=None, pull_message=None, response_count=None, schedule_time=None, view=None):
        if app_engine_http_request and not isinstance(app_engine_http_request, dict):
            raise TypeError("Expected argument 'app_engine_http_request' to be a dict")
        pulumi.set(__self__, "app_engine_http_request", app_engine_http_request)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if dispatch_count and not isinstance(dispatch_count, int):
            raise TypeError("Expected argument 'dispatch_count' to be a int")
        pulumi.set(__self__, "dispatch_count", dispatch_count)
        if dispatch_deadline and not isinstance(dispatch_deadline, str):
            raise TypeError("Expected argument 'dispatch_deadline' to be a str")
        pulumi.set(__self__, "dispatch_deadline", dispatch_deadline)
        if first_attempt and not isinstance(first_attempt, dict):
            raise TypeError("Expected argument 'first_attempt' to be a dict")
        pulumi.set(__self__, "first_attempt", first_attempt)
        if http_request and not isinstance(http_request, dict):
            raise TypeError("Expected argument 'http_request' to be a dict")
        pulumi.set(__self__, "http_request", http_request)
        if last_attempt and not isinstance(last_attempt, dict):
            raise TypeError("Expected argument 'last_attempt' to be a dict")
        pulumi.set(__self__, "last_attempt", last_attempt)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pull_message and not isinstance(pull_message, dict):
            raise TypeError("Expected argument 'pull_message' to be a dict")
        pulumi.set(__self__, "pull_message", pull_message)
        if response_count and not isinstance(response_count, int):
            raise TypeError("Expected argument 'response_count' to be a int")
        pulumi.set(__self__, "response_count", response_count)
        if schedule_time and not isinstance(schedule_time, str):
            raise TypeError("Expected argument 'schedule_time' to be a str")
        pulumi.set(__self__, "schedule_time", schedule_time)
        if view and not isinstance(view, str):
            raise TypeError("Expected argument 'view' to be a str")
        pulumi.set(__self__, "view", view)

    @property
    @pulumi.getter(name="appEngineHttpRequest")
    def app_engine_http_request(self) -> 'outputs.AppEngineHttpRequestResponse':
        """
        HTTP request that is sent to the App Engine app handler. An App Engine task is a task that has AppEngineHttpRequest set.
        """
        return pulumi.get(self, "app_engine_http_request")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time that the task was created. `create_time` will be truncated to the nearest second.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dispatchCount")
    def dispatch_count(self) -> int:
        """
        The number of attempts dispatched. This count includes attempts which have been dispatched but haven't received a response.
        """
        return pulumi.get(self, "dispatch_count")

    @property
    @pulumi.getter(name="dispatchDeadline")
    def dispatch_deadline(self) -> str:
        """
        The deadline for requests sent to the worker. If the worker does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the task according to the RetryConfig. Note that when the request is cancelled, Cloud Tasks will stop listening for the response, but whether the worker stops processing depends on the worker. For example, if the worker is stuck, it may not react to cancelled requests. The default and maximum values depend on the type of request: * For HTTP tasks, the default is 10 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine tasks, 0 indicates that the request has the default deadline. The default deadline depends on the [scaling type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling) of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. Regardless of the task's `dispatch_deadline`, the app handler will not run for longer than than the service's timeout. We recommend setting the `dispatch_deadline` to at most a few seconds more than the app handler's timeout. For more information see [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts). `dispatch_deadline` will be truncated to the nearest millisecond. The deadline is an approximate deadline.
        """
        return pulumi.get(self, "dispatch_deadline")

    @property
    @pulumi.getter(name="firstAttempt")
    def first_attempt(self) -> 'outputs.AttemptResponse':
        """
        The status of the task's first attempt. Only dispatch_time will be set. The other Attempt information is not retained by Cloud Tasks.
        """
        return pulumi.get(self, "first_attempt")

    @property
    @pulumi.getter(name="httpRequest")
    def http_request(self) -> 'outputs.HttpRequestResponse':
        """
        HTTP request that is sent to the task's target. An HTTP task is a task that has HttpRequest set.
        """
        return pulumi.get(self, "http_request")

    @property
    @pulumi.getter(name="lastAttempt")
    def last_attempt(self) -> 'outputs.AttemptResponse':
        """
        The status of the task's last attempt.
        """
        return pulumi.get(self, "last_attempt")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Optionally caller-specified in CreateTask. The task name. The task name must have the following format: `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID` * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the task's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or hyphens (-). The maximum length is 100 characters. * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pullMessage")
    def pull_message(self) -> 'outputs.PullMessageResponse':
        """
        Pull Message contained in a task in a PULL queue type. This payload type cannot be explicitly set through Cloud Tasks API. Its purpose, currently is to provide backward compatibility with App Engine Task Queue [pull](https://cloud.google.com/appengine/docs/standard/java/taskqueue/pull/) queues to provide a way to inspect contents of pull tasks through the CloudTasks.GetTask.
        """
        return pulumi.get(self, "pull_message")

    @property
    @pulumi.getter(name="responseCount")
    def response_count(self) -> int:
        """
        The number of attempts which have received a response.
        """
        return pulumi.get(self, "response_count")

    @property
    @pulumi.getter(name="scheduleTime")
    def schedule_time(self) -> str:
        """
        The time when the task is scheduled to be attempted. For App Engine queues, this is when the task will be attempted or retried. `schedule_time` will be truncated to the nearest microsecond.
        """
        return pulumi.get(self, "schedule_time")

    @property
    @pulumi.getter
    def view(self) -> str:
        """
        The view specifies which subset of the Task has been returned.
        """
        return pulumi.get(self, "view")


class AwaitableGetTaskResult(GetTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTaskResult(
            app_engine_http_request=self.app_engine_http_request,
            create_time=self.create_time,
            dispatch_count=self.dispatch_count,
            dispatch_deadline=self.dispatch_deadline,
            first_attempt=self.first_attempt,
            http_request=self.http_request,
            last_attempt=self.last_attempt,
            name=self.name,
            pull_message=self.pull_message,
            response_count=self.response_count,
            schedule_time=self.schedule_time,
            view=self.view)


def get_task(location: Optional[str] = None,
             project: Optional[str] = None,
             queue_id: Optional[str] = None,
             response_view: Optional[str] = None,
             task_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTaskResult:
    """
    Gets a task.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['queueId'] = queue_id
    __args__['responseView'] = response_view
    __args__['taskId'] = task_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudtasks/v2beta3:getTask', __args__, opts=opts, typ=GetTaskResult).value

    return AwaitableGetTaskResult(
        app_engine_http_request=__ret__.app_engine_http_request,
        create_time=__ret__.create_time,
        dispatch_count=__ret__.dispatch_count,
        dispatch_deadline=__ret__.dispatch_deadline,
        first_attempt=__ret__.first_attempt,
        http_request=__ret__.http_request,
        last_attempt=__ret__.last_attempt,
        name=__ret__.name,
        pull_message=__ret__.pull_message,
        response_count=__ret__.response_count,
        schedule_time=__ret__.schedule_time,
        view=__ret__.view)


@_utilities.lift_output_func(get_task)
def get_task_output(location: Optional[pulumi.Input[str]] = None,
                    project: Optional[pulumi.Input[Optional[str]]] = None,
                    queue_id: Optional[pulumi.Input[str]] = None,
                    response_view: Optional[pulumi.Input[Optional[str]]] = None,
                    task_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTaskResult]:
    """
    Gets a task.
    """
    ...
