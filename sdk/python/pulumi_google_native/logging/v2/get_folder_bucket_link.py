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
    'GetFolderBucketLinkResult',
    'AwaitableGetFolderBucketLinkResult',
    'get_folder_bucket_link',
    'get_folder_bucket_link_output',
]

@pulumi.output_type
class GetFolderBucketLinkResult:
    def __init__(__self__, bigquery_dataset=None, create_time=None, description=None, lifecycle_state=None, name=None):
        if bigquery_dataset and not isinstance(bigquery_dataset, dict):
            raise TypeError("Expected argument 'bigquery_dataset' to be a dict")
        pulumi.set(__self__, "bigquery_dataset", bigquery_dataset)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if lifecycle_state and not isinstance(lifecycle_state, str):
            raise TypeError("Expected argument 'lifecycle_state' to be a str")
        pulumi.set(__self__, "lifecycle_state", lifecycle_state)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="bigqueryDataset")
    def bigquery_dataset(self) -> 'outputs.BigQueryDatasetResponse':
        """
        The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
        """
        return pulumi.get(self, "bigquery_dataset")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of the link.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Describes this link.The maximum length of the description is 8000 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> str:
        """
        The resource lifecycle state.
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
        """
        return pulumi.get(self, "name")


class AwaitableGetFolderBucketLinkResult(GetFolderBucketLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFolderBucketLinkResult(
            bigquery_dataset=self.bigquery_dataset,
            create_time=self.create_time,
            description=self.description,
            lifecycle_state=self.lifecycle_state,
            name=self.name)


def get_folder_bucket_link(bucket_id: Optional[str] = None,
                           folder_id: Optional[str] = None,
                           link_id: Optional[str] = None,
                           location: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFolderBucketLinkResult:
    """
    Gets a link.
    """
    __args__ = dict()
    __args__['bucketId'] = bucket_id
    __args__['folderId'] = folder_id
    __args__['linkId'] = link_id
    __args__['location'] = location
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:logging/v2:getFolderBucketLink', __args__, opts=opts, typ=GetFolderBucketLinkResult).value

    return AwaitableGetFolderBucketLinkResult(
        bigquery_dataset=__ret__.bigquery_dataset,
        create_time=__ret__.create_time,
        description=__ret__.description,
        lifecycle_state=__ret__.lifecycle_state,
        name=__ret__.name)


@_utilities.lift_output_func(get_folder_bucket_link)
def get_folder_bucket_link_output(bucket_id: Optional[pulumi.Input[str]] = None,
                                  folder_id: Optional[pulumi.Input[str]] = None,
                                  link_id: Optional[pulumi.Input[str]] = None,
                                  location: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFolderBucketLinkResult]:
    """
    Gets a link.
    """
    ...
