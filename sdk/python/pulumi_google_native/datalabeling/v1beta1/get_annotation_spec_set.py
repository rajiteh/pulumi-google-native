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
    'GetAnnotationSpecSetResult',
    'AwaitableGetAnnotationSpecSetResult',
    'get_annotation_spec_set',
    'get_annotation_spec_set_output',
]

@pulumi.output_type
class GetAnnotationSpecSetResult:
    def __init__(__self__, annotation_specs=None, blocking_resources=None, description=None, display_name=None, name=None):
        if annotation_specs and not isinstance(annotation_specs, list):
            raise TypeError("Expected argument 'annotation_specs' to be a list")
        pulumi.set(__self__, "annotation_specs", annotation_specs)
        if blocking_resources and not isinstance(blocking_resources, list):
            raise TypeError("Expected argument 'blocking_resources' to be a list")
        pulumi.set(__self__, "blocking_resources", blocking_resources)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="annotationSpecs")
    def annotation_specs(self) -> Sequence['outputs.GoogleCloudDatalabelingV1beta1AnnotationSpecResponse']:
        """
        The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
        """
        return pulumi.get(self, "annotation_specs")

    @property
    @pulumi.getter(name="blockingResources")
    def blocking_resources(self) -> Sequence[str]:
        """
        The names of any related resources that are blocking changes to the annotation spec set.
        """
        return pulumi.get(self, "blocking_resources")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The AnnotationSpecSet resource name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
        """
        return pulumi.get(self, "name")


class AwaitableGetAnnotationSpecSetResult(GetAnnotationSpecSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAnnotationSpecSetResult(
            annotation_specs=self.annotation_specs,
            blocking_resources=self.blocking_resources,
            description=self.description,
            display_name=self.display_name,
            name=self.name)


def get_annotation_spec_set(annotation_spec_set_id: Optional[str] = None,
                            project: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAnnotationSpecSetResult:
    """
    Gets an annotation spec set by resource name.
    """
    __args__ = dict()
    __args__['annotationSpecSetId'] = annotation_spec_set_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datalabeling/v1beta1:getAnnotationSpecSet', __args__, opts=opts, typ=GetAnnotationSpecSetResult).value

    return AwaitableGetAnnotationSpecSetResult(
        annotation_specs=__ret__.annotation_specs,
        blocking_resources=__ret__.blocking_resources,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name)


@_utilities.lift_output_func(get_annotation_spec_set)
def get_annotation_spec_set_output(annotation_spec_set_id: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAnnotationSpecSetResult]:
    """
    Gets an annotation spec set by resource name.
    """
    ...
