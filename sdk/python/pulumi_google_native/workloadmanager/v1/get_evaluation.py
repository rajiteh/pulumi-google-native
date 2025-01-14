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
    'GetEvaluationResult',
    'AwaitableGetEvaluationResult',
    'get_evaluation',
    'get_evaluation_output',
]

@pulumi.output_type
class GetEvaluationResult:
    def __init__(__self__, create_time=None, description=None, labels=None, name=None, resource_filter=None, resource_status=None, rule_names=None, rule_versions=None, schedule=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_filter and not isinstance(resource_filter, dict):
            raise TypeError("Expected argument 'resource_filter' to be a dict")
        pulumi.set(__self__, "resource_filter", resource_filter)
        if resource_status and not isinstance(resource_status, dict):
            raise TypeError("Expected argument 'resource_status' to be a dict")
        pulumi.set(__self__, "resource_status", resource_status)
        if rule_names and not isinstance(rule_names, list):
            raise TypeError("Expected argument 'rule_names' to be a list")
        pulumi.set(__self__, "rule_names", rule_names)
        if rule_versions and not isinstance(rule_versions, list):
            raise TypeError("Expected argument 'rule_versions' to be a list")
        pulumi.set(__self__, "rule_versions", rule_versions)
        if schedule and not isinstance(schedule, str):
            raise TypeError("Expected argument 'schedule' to be a str")
        pulumi.set(__self__, "schedule", schedule)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        [Output only] Create time stamp
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the Evaluation
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels as key value pairs
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of resource names have the form 'projects/{project_id}/locations/{location_id}/evaluations/{evaluation_id}'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceFilter")
    def resource_filter(self) -> 'outputs.ResourceFilterResponse':
        """
        annotations as key value pairs
        """
        return pulumi.get(self, "resource_filter")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> 'outputs.ResourceStatusResponse':
        """
        [Output only] The updated rule ids if exist.
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="ruleNames")
    def rule_names(self) -> Sequence[str]:
        """
        the name of the rule
        """
        return pulumi.get(self, "rule_names")

    @property
    @pulumi.getter(name="ruleVersions")
    def rule_versions(self) -> Sequence[str]:
        """
        [Output only] The updated rule ids if exist.
        """
        return pulumi.get(self, "rule_versions")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        crontab format schedule for scheduled evaluation, example: 0 */3 * * *
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        [Output only] Update time stamp
        """
        return pulumi.get(self, "update_time")


class AwaitableGetEvaluationResult(GetEvaluationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEvaluationResult(
            create_time=self.create_time,
            description=self.description,
            labels=self.labels,
            name=self.name,
            resource_filter=self.resource_filter,
            resource_status=self.resource_status,
            rule_names=self.rule_names,
            rule_versions=self.rule_versions,
            schedule=self.schedule,
            update_time=self.update_time)


def get_evaluation(evaluation_id: Optional[str] = None,
                   location: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEvaluationResult:
    """
    Gets details of a single Evaluation.
    """
    __args__ = dict()
    __args__['evaluationId'] = evaluation_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:workloadmanager/v1:getEvaluation', __args__, opts=opts, typ=GetEvaluationResult).value

    return AwaitableGetEvaluationResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        labels=__ret__.labels,
        name=__ret__.name,
        resource_filter=__ret__.resource_filter,
        resource_status=__ret__.resource_status,
        rule_names=__ret__.rule_names,
        rule_versions=__ret__.rule_versions,
        schedule=__ret__.schedule,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_evaluation)
def get_evaluation_output(evaluation_id: Optional[pulumi.Input[str]] = None,
                          location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEvaluationResult]:
    """
    Gets details of a single Evaluation.
    """
    ...
