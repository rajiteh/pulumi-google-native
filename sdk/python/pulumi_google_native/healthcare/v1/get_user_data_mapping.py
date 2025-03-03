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
    'GetUserDataMappingResult',
    'AwaitableGetUserDataMappingResult',
    'get_user_data_mapping',
    'get_user_data_mapping_output',
]

@pulumi.output_type
class GetUserDataMappingResult:
    def __init__(__self__, archive_time=None, archived=None, data_id=None, name=None, resource_attributes=None, user_id=None):
        if archive_time and not isinstance(archive_time, str):
            raise TypeError("Expected argument 'archive_time' to be a str")
        pulumi.set(__self__, "archive_time", archive_time)
        if archived and not isinstance(archived, bool):
            raise TypeError("Expected argument 'archived' to be a bool")
        pulumi.set(__self__, "archived", archived)
        if data_id and not isinstance(data_id, str):
            raise TypeError("Expected argument 'data_id' to be a str")
        pulumi.set(__self__, "data_id", data_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_attributes and not isinstance(resource_attributes, list):
            raise TypeError("Expected argument 'resource_attributes' to be a list")
        pulumi.set(__self__, "resource_attributes", resource_attributes)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="archiveTime")
    def archive_time(self) -> str:
        """
        Indicates the time when this mapping was archived.
        """
        return pulumi.get(self, "archive_time")

    @property
    @pulumi.getter
    def archived(self) -> bool:
        """
        Indicates whether this mapping is archived.
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="dataId")
    def data_id(self) -> str:
        """
        A unique identifier for the mapped resource.
        """
        return pulumi.get(self, "data_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the User data mapping, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/userDataMappings/{user_data_mapping_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceAttributes")
    def resource_attributes(self) -> Sequence['outputs.AttributeResponse']:
        """
        Attributes of the resource. Only explicitly set attributes are displayed here. Attribute definitions with defaults set implicitly apply to these User data mappings. Attributes listed here must be single valued, that is, exactly one value is specified for the field "values" in each Attribute.
        """
        return pulumi.get(self, "resource_attributes")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")


class AwaitableGetUserDataMappingResult(GetUserDataMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserDataMappingResult(
            archive_time=self.archive_time,
            archived=self.archived,
            data_id=self.data_id,
            name=self.name,
            resource_attributes=self.resource_attributes,
            user_id=self.user_id)


def get_user_data_mapping(consent_store_id: Optional[str] = None,
                          dataset_id: Optional[str] = None,
                          location: Optional[str] = None,
                          project: Optional[str] = None,
                          user_data_mapping_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserDataMappingResult:
    """
    Gets the specified User data mapping.
    """
    __args__ = dict()
    __args__['consentStoreId'] = consent_store_id
    __args__['datasetId'] = dataset_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['userDataMappingId'] = user_data_mapping_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:healthcare/v1:getUserDataMapping', __args__, opts=opts, typ=GetUserDataMappingResult).value

    return AwaitableGetUserDataMappingResult(
        archive_time=__ret__.archive_time,
        archived=__ret__.archived,
        data_id=__ret__.data_id,
        name=__ret__.name,
        resource_attributes=__ret__.resource_attributes,
        user_id=__ret__.user_id)


@_utilities.lift_output_func(get_user_data_mapping)
def get_user_data_mapping_output(consent_store_id: Optional[pulumi.Input[str]] = None,
                                 dataset_id: Optional[pulumi.Input[str]] = None,
                                 location: Optional[pulumi.Input[str]] = None,
                                 project: Optional[pulumi.Input[Optional[str]]] = None,
                                 user_data_mapping_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserDataMappingResult]:
    """
    Gets the specified User data mapping.
    """
    ...
