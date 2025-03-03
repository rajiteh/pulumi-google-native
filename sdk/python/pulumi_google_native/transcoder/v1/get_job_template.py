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
    'GetJobTemplateResult',
    'AwaitableGetJobTemplateResult',
    'get_job_template',
    'get_job_template_output',
]

@pulumi.output_type
class GetJobTemplateResult:
    def __init__(__self__, config=None, labels=None, name=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def config(self) -> 'outputs.JobConfigResponse':
        """
        The configuration for this template.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels associated with this job template. You can use these to organize and group your job templates.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the job template. Format: `projects/{project_number}/locations/{location}/jobTemplates/{job_template}`
        """
        return pulumi.get(self, "name")


class AwaitableGetJobTemplateResult(GetJobTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobTemplateResult(
            config=self.config,
            labels=self.labels,
            name=self.name)


def get_job_template(job_template_id: Optional[str] = None,
                     location: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobTemplateResult:
    """
    Returns the job template data.
    """
    __args__ = dict()
    __args__['jobTemplateId'] = job_template_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:transcoder/v1:getJobTemplate', __args__, opts=opts, typ=GetJobTemplateResult).value

    return AwaitableGetJobTemplateResult(
        config=__ret__.config,
        labels=__ret__.labels,
        name=__ret__.name)


@_utilities.lift_output_func(get_job_template)
def get_job_template_output(job_template_id: Optional[pulumi.Input[str]] = None,
                            location: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobTemplateResult]:
    """
    Returns the job template data.
    """
    ...
