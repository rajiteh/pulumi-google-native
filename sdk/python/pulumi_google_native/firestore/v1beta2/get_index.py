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
    'GetIndexResult',
    'AwaitableGetIndexResult',
    'get_index',
    'get_index_output',
]

@pulumi.output_type
class GetIndexResult:
    def __init__(__self__, fields=None, name=None, query_scope=None, state=None):
        if fields and not isinstance(fields, list):
            raise TypeError("Expected argument 'fields' to be a list")
        pulumi.set(__self__, "fields", fields)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query_scope and not isinstance(query_scope, str):
            raise TypeError("Expected argument 'query_scope' to be a str")
        pulumi.set(__self__, "query_scope", query_scope)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def fields(self) -> Sequence['outputs.GoogleFirestoreAdminV1beta2IndexFieldResponse']:
        """
        The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> str:
        """
        Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        return pulumi.get(self, "query_scope")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The serving state of the index.
        """
        return pulumi.get(self, "state")


class AwaitableGetIndexResult(GetIndexResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIndexResult(
            fields=self.fields,
            name=self.name,
            query_scope=self.query_scope,
            state=self.state)


def get_index(collection_group_id: Optional[str] = None,
              database_id: Optional[str] = None,
              index_id: Optional[str] = None,
              project: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIndexResult:
    """
    Gets a composite index.
    """
    __args__ = dict()
    __args__['collectionGroupId'] = collection_group_id
    __args__['databaseId'] = database_id
    __args__['indexId'] = index_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:firestore/v1beta2:getIndex', __args__, opts=opts, typ=GetIndexResult).value

    return AwaitableGetIndexResult(
        fields=__ret__.fields,
        name=__ret__.name,
        query_scope=__ret__.query_scope,
        state=__ret__.state)


@_utilities.lift_output_func(get_index)
def get_index_output(collection_group_id: Optional[pulumi.Input[str]] = None,
                     database_id: Optional[pulumi.Input[str]] = None,
                     index_id: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIndexResult]:
    """
    Gets a composite index.
    """
    ...
