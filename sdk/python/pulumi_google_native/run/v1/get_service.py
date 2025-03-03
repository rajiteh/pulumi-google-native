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
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
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
        The API version for this call. It must be "serving.knative.dev/v1".
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of resource. It must be "Service".
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.ObjectMetaResponse':
        """
        Metadata associated with this Service, including name, namespace, labels, and annotations. In Cloud Run, annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted, and the accepted annotations will be different depending on the resource type. The following Cloud Run-specific annotations are accepted in Service.metadata.annotations. * `run.googleapis.com/binary-authorization-breakglass` * `run.googleapis.com/binary-authorization` * `run.googleapis.com/client-name` * `run.googleapis.com/custom-audiences` * `run.googleapis.com/description` * `run.googleapis.com/gc-traffic-tags` * `run.googleapis.com/ingress` * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.ServiceSpecResponse':
        """
        Holds the desired state of the Service (from the client).
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.ServiceStatusResponse':
        """
        Communicates the system-controlled state of the Service.
        """
        return pulumi.get(self, "status")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            api_version=self.api_version,
            kind=self.kind,
            metadata=self.metadata,
            spec=self.spec,
            status=self.status)


def get_service(location: Optional[str] = None,
                project: Optional[str] = None,
                service_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Gets information about a service.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:run/v1:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        api_version=__ret__.api_version,
        kind=__ret__.kind,
        metadata=__ret__.metadata,
        spec=__ret__.spec,
        status=__ret__.status)


@_utilities.lift_output_func(get_service)
def get_service_output(location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       service_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    Gets information about a service.
    """
    ...
