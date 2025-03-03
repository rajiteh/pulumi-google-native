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
    'GetContentResult',
    'AwaitableGetContentResult',
    'get_content',
    'get_content_output',
]

@pulumi.output_type
class GetContentResult:
    def __init__(__self__, create_time=None, data_text=None, description=None, labels=None, name=None, notebook=None, path=None, sql_script=None, uid=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if data_text and not isinstance(data_text, str):
            raise TypeError("Expected argument 'data_text' to be a str")
        pulumi.set(__self__, "data_text", data_text)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notebook and not isinstance(notebook, dict):
            raise TypeError("Expected argument 'notebook' to be a dict")
        pulumi.set(__self__, "notebook", notebook)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if sql_script and not isinstance(sql_script, dict):
            raise TypeError("Expected argument 'sql_script' to be a dict")
        pulumi.set(__self__, "sql_script", sql_script)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Content creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dataText")
    def data_text(self) -> str:
        """
        Content data in string format.
        """
        return pulumi.get(self, "data_text")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the content.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User defined labels for the content.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of the content, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notebook(self) -> 'outputs.GoogleCloudDataplexV1ContentNotebookResponse':
        """
        Notebook related configurations.
        """
        return pulumi.get(self, "notebook")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="sqlScript")
    def sql_script(self) -> 'outputs.GoogleCloudDataplexV1ContentSqlScriptResponse':
        """
        Sql Script related configurations.
        """
        return pulumi.get(self, "sql_script")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System generated globally unique ID for the content. This ID will be different if the content is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the content was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetContentResult(GetContentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContentResult(
            create_time=self.create_time,
            data_text=self.data_text,
            description=self.description,
            labels=self.labels,
            name=self.name,
            notebook=self.notebook,
            path=self.path,
            sql_script=self.sql_script,
            uid=self.uid,
            update_time=self.update_time)


def get_content(content_id: Optional[str] = None,
                lake_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                view: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContentResult:
    """
    Get a content resource.
    """
    __args__ = dict()
    __args__['contentId'] = content_id
    __args__['lakeId'] = lake_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getContent', __args__, opts=opts, typ=GetContentResult).value

    return AwaitableGetContentResult(
        create_time=__ret__.create_time,
        data_text=__ret__.data_text,
        description=__ret__.description,
        labels=__ret__.labels,
        name=__ret__.name,
        notebook=__ret__.notebook,
        path=__ret__.path,
        sql_script=__ret__.sql_script,
        uid=__ret__.uid,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_content)
def get_content_output(content_id: Optional[pulumi.Input[str]] = None,
                       lake_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       view: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContentResult]:
    """
    Get a content resource.
    """
    ...
