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
    'GetModelResult',
    'AwaitableGetModelResult',
    'get_model',
    'get_model_output',
]

@pulumi.output_type
class GetModelResult:
    def __init__(__self__, default_version=None, description=None, etag=None, labels=None, name=None, online_prediction_console_logging=None, online_prediction_logging=None, regions=None):
        if default_version and not isinstance(default_version, dict):
            raise TypeError("Expected argument 'default_version' to be a dict")
        pulumi.set(__self__, "default_version", default_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if online_prediction_console_logging and not isinstance(online_prediction_console_logging, bool):
            raise TypeError("Expected argument 'online_prediction_console_logging' to be a bool")
        pulumi.set(__self__, "online_prediction_console_logging", online_prediction_console_logging)
        if online_prediction_logging and not isinstance(online_prediction_logging, bool):
            raise TypeError("Expected argument 'online_prediction_logging' to be a bool")
        pulumi.set(__self__, "online_prediction_logging", online_prediction_logging)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> 'outputs.GoogleCloudMlV1__VersionResponse':
        """
        The default version of the model. This version will be used to handle prediction requests that do not specify a version. You can change the default version by calling projects.models.versions.setDefault.
        """
        return pulumi.get(self, "default_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The description specified for the model when it was created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a model from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform model updates in order to avoid race conditions: An `etag` is returned in the response to `GetModel`, and systems are expected to put that etag in the request to `UpdateModel` to ensure that their change will be applied to the model as intended.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. One or more labels that you can add, to organize your models. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels. Note that this field is not updatable for mls1* models.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name specified for the model when it was created. The model name must be unique within the project it is created in.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onlinePredictionConsoleLogging")
    def online_prediction_console_logging(self) -> bool:
        """
        Optional. If true, online prediction nodes send `stderr` and `stdout` streams to Cloud Logging. These can be more verbose than the standard access logs (see `onlinePredictionLogging`) and can incur higher cost. However, they are helpful for debugging. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high QPS. Estimate your costs before enabling this option. Default is false.
        """
        return pulumi.get(self, "online_prediction_console_logging")

    @property
    @pulumi.getter(name="onlinePredictionLogging")
    def online_prediction_logging(self) -> bool:
        """
        Optional. If true, online prediction access logs are sent to Cloud Logging. These logs are like standard server access logs, containing information like timestamp and latency for each request. Note that [logs may incur a cost](/stackdriver/pricing), especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. Default is false.
        """
        return pulumi.get(self, "online_prediction_logging")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        Optional. The list of regions where the model is going to be deployed. Only one region per model is supported. Defaults to 'us-central1' if nothing is set. See the available regions for AI Platform services. Note: * No matter where a model is deployed, it can always be accessed by users from anywhere, both for online and batch prediction. * The region for a batch prediction job is set by the region field when submitting the batch prediction job and does not take its value from this field.
        """
        return pulumi.get(self, "regions")


class AwaitableGetModelResult(GetModelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelResult(
            default_version=self.default_version,
            description=self.description,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            online_prediction_console_logging=self.online_prediction_console_logging,
            online_prediction_logging=self.online_prediction_logging,
            regions=self.regions)


def get_model(model_id: Optional[str] = None,
              project: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelResult:
    """
    Gets information about a model, including its name, the description (if set), and the default version (if at least one version of the model has been deployed).
    """
    __args__ = dict()
    __args__['modelId'] = model_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:ml/v1:getModel', __args__, opts=opts, typ=GetModelResult).value

    return AwaitableGetModelResult(
        default_version=__ret__.default_version,
        description=__ret__.description,
        etag=__ret__.etag,
        labels=__ret__.labels,
        name=__ret__.name,
        online_prediction_console_logging=__ret__.online_prediction_console_logging,
        online_prediction_logging=__ret__.online_prediction_logging,
        regions=__ret__.regions)


@_utilities.lift_output_func(get_model)
def get_model_output(model_id: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetModelResult]:
    """
    Gets information about a model, including its name, the description (if set), and the default version (if at least one version of the model has been deployed).
    """
    ...
