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
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    def __init__(__self__, api_version=None, kind=None, metadata=None, spec=None, status=None):
        if api_version and not isinstance(api_version, str):
            raise TypeError("Expected argument 'api_version' to be a str")
        pulumi.set(__self__, "api_version", api_version)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if spec and not isinstance(spec, dict):
            raise TypeError("Expected argument 'spec' to be a dict")
        pulumi.set(__self__, "spec", spec)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> str:
        """
        Optional. APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Optional. Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.ObjectMetaResponse':
        """
        Optional. Standard object's metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.JobSpecResponse':
        """
        Optional. Specification of the desired behavior of a job.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.JobStatusResponse':
        """
        Current status of a job.
        """
        return pulumi.get(self, "status")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            api_version=self.api_version,
            kind=self.kind,
            metadata=self.metadata,
            spec=self.spec,
            status=self.status)


def get_job(job_id: Optional[str] = None,
            namespace_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Get information about a job.
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    __args__['namespaceId'] = namespace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:run/v1:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        api_version=__ret__.api_version,
        kind=__ret__.kind,
        metadata=__ret__.metadata,
        spec=__ret__.spec,
        status=__ret__.status)


@_utilities.lift_output_func(get_job)
def get_job_output(job_id: Optional[pulumi.Input[str]] = None,
                   namespace_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobResult]:
    """
    Get information about a job.
    """
    ...
