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
from ._enums import *
from ._inputs import *

__all__ = ['EntryArgs', 'Entry']

@pulumi.input_type
class EntryArgs:
    def __init__(__self__, *,
                 entry_group_id: pulumi.Input[str],
                 entry_id: pulumi.Input[str],
                 bigquery_date_sharded_spec: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']] = None,
                 bigquery_table_spec: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gcs_fileset_spec: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']] = None,
                 linked_resource: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1SchemaArgs']] = None,
                 type: Optional[pulumi.Input['EntryType']] = None,
                 user_specified_system: Optional[pulumi.Input[str]] = None,
                 user_specified_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Entry resource.
        :param pulumi.Input[str] entry_id: Required. The id of the entry to create.
        :param pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs'] bigquery_date_sharded_spec: Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        :param pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs'] bigquery_table_spec: Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        :param pulumi.Input[str] description: Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        :param pulumi.Input[str] display_name: Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        :param pulumi.Input['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs'] gcs_fileset_spec: Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        :param pulumi.Input[str] linked_resource: The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        :param pulumi.Input['GoogleCloudDatacatalogV1beta1SchemaArgs'] schema: Schema of the entry. An entry might not have any schema attached to it.
        :param pulumi.Input['EntryType'] type: The type of the entry. Only used for Entries with types in the EntryType enum.
        :param pulumi.Input[str] user_specified_system: This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        :param pulumi.Input[str] user_specified_type: Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        """
        pulumi.set(__self__, "entry_group_id", entry_group_id)
        pulumi.set(__self__, "entry_id", entry_id)
        if bigquery_date_sharded_spec is not None:
            pulumi.set(__self__, "bigquery_date_sharded_spec", bigquery_date_sharded_spec)
        if bigquery_table_spec is not None:
            pulumi.set(__self__, "bigquery_table_spec", bigquery_table_spec)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if gcs_fileset_spec is not None:
            pulumi.set(__self__, "gcs_fileset_spec", gcs_fileset_spec)
        if linked_resource is not None:
            pulumi.set(__self__, "linked_resource", linked_resource)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_specified_system is not None:
            pulumi.set(__self__, "user_specified_system", user_specified_system)
        if user_specified_type is not None:
            pulumi.set(__self__, "user_specified_type", user_specified_type)

    @property
    @pulumi.getter(name="entryGroupId")
    def entry_group_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "entry_group_id")

    @entry_group_id.setter
    def entry_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry_group_id", value)

    @property
    @pulumi.getter(name="entryId")
    def entry_id(self) -> pulumi.Input[str]:
        """
        Required. The id of the entry to create.
        """
        return pulumi.get(self, "entry_id")

    @entry_id.setter
    def entry_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry_id", value)

    @property
    @pulumi.getter(name="bigqueryDateShardedSpec")
    def bigquery_date_sharded_spec(self) -> Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']]:
        """
        Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        """
        return pulumi.get(self, "bigquery_date_sharded_spec")

    @bigquery_date_sharded_spec.setter
    def bigquery_date_sharded_spec(self, value: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']]):
        pulumi.set(self, "bigquery_date_sharded_spec", value)

    @property
    @pulumi.getter(name="bigqueryTableSpec")
    def bigquery_table_spec(self) -> Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']]:
        """
        Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        """
        return pulumi.get(self, "bigquery_table_spec")

    @bigquery_table_spec.setter
    def bigquery_table_spec(self, value: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']]):
        pulumi.set(self, "bigquery_table_spec", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="gcsFilesetSpec")
    def gcs_fileset_spec(self) -> Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']]:
        """
        Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        """
        return pulumi.get(self, "gcs_fileset_spec")

    @gcs_fileset_spec.setter
    def gcs_fileset_spec(self, value: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']]):
        pulumi.set(self, "gcs_fileset_spec", value)

    @property
    @pulumi.getter(name="linkedResource")
    def linked_resource(self) -> Optional[pulumi.Input[str]]:
        """
        The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        """
        return pulumi.get(self, "linked_resource")

    @linked_resource.setter
    def linked_resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "linked_resource", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1SchemaArgs']]:
        """
        Schema of the entry. An entry might not have any schema attached to it.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input['GoogleCloudDatacatalogV1beta1SchemaArgs']]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['EntryType']]:
        """
        The type of the entry. Only used for Entries with types in the EntryType enum.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['EntryType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userSpecifiedSystem")
    def user_specified_system(self) -> Optional[pulumi.Input[str]]:
        """
        This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        """
        return pulumi.get(self, "user_specified_system")

    @user_specified_system.setter
    def user_specified_system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_specified_system", value)

    @property
    @pulumi.getter(name="userSpecifiedType")
    def user_specified_type(self) -> Optional[pulumi.Input[str]]:
        """
        Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        """
        return pulumi.get(self, "user_specified_type")

    @user_specified_type.setter
    def user_specified_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_specified_type", value)


class Entry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_date_sharded_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']]] = None,
                 bigquery_table_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 entry_group_id: Optional[pulumi.Input[str]] = None,
                 entry_id: Optional[pulumi.Input[str]] = None,
                 gcs_fileset_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']]] = None,
                 linked_resource: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1SchemaArgs']]] = None,
                 type: Optional[pulumi.Input['EntryType']] = None,
                 user_specified_system: Optional[pulumi.Input[str]] = None,
                 user_specified_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an entry. Only entries of 'FILESET' type or user-specified type can be created. Users should enable the Data Catalog API in the project identified by the `parent` parameter (see [Data Catalog Resource Project] (https://cloud.google.com/data-catalog/docs/concepts/resource-project) for more information). A maximum of 100,000 entries may be created per entry group.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']] bigquery_date_sharded_spec: Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']] bigquery_table_spec: Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        :param pulumi.Input[str] description: Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        :param pulumi.Input[str] display_name: Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        :param pulumi.Input[str] entry_id: Required. The id of the entry to create.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']] gcs_fileset_spec: Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        :param pulumi.Input[str] linked_resource: The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1SchemaArgs']] schema: Schema of the entry. An entry might not have any schema attached to it.
        :param pulumi.Input['EntryType'] type: The type of the entry. Only used for Entries with types in the EntryType enum.
        :param pulumi.Input[str] user_specified_system: This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        :param pulumi.Input[str] user_specified_type: Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an entry. Only entries of 'FILESET' type or user-specified type can be created. Users should enable the Data Catalog API in the project identified by the `parent` parameter (see [Data Catalog Resource Project] (https://cloud.google.com/data-catalog/docs/concepts/resource-project) for more information). A maximum of 100,000 entries may be created per entry group.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param EntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bigquery_date_sharded_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecArgs']]] = None,
                 bigquery_table_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1BigQueryTableSpecArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 entry_group_id: Optional[pulumi.Input[str]] = None,
                 entry_id: Optional[pulumi.Input[str]] = None,
                 gcs_fileset_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1GcsFilesetSpecArgs']]] = None,
                 linked_resource: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatacatalogV1beta1SchemaArgs']]] = None,
                 type: Optional[pulumi.Input['EntryType']] = None,
                 user_specified_system: Optional[pulumi.Input[str]] = None,
                 user_specified_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EntryArgs.__new__(EntryArgs)

            __props__.__dict__["bigquery_date_sharded_spec"] = bigquery_date_sharded_spec
            __props__.__dict__["bigquery_table_spec"] = bigquery_table_spec
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if entry_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'entry_group_id'")
            __props__.__dict__["entry_group_id"] = entry_group_id
            if entry_id is None and not opts.urn:
                raise TypeError("Missing required property 'entry_id'")
            __props__.__dict__["entry_id"] = entry_id
            __props__.__dict__["gcs_fileset_spec"] = gcs_fileset_spec
            __props__.__dict__["linked_resource"] = linked_resource
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["schema"] = schema
            __props__.__dict__["type"] = type
            __props__.__dict__["user_specified_system"] = user_specified_system
            __props__.__dict__["user_specified_type"] = user_specified_type
            __props__.__dict__["integrated_system"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["source_system_timestamps"] = None
            __props__.__dict__["usage_signal"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["entry_group_id", "entry_id", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Entry, __self__).__init__(
            'google-native:datacatalog/v1beta1:Entry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Entry':
        """
        Get an existing Entry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EntryArgs.__new__(EntryArgs)

        __props__.__dict__["bigquery_date_sharded_spec"] = None
        __props__.__dict__["bigquery_table_spec"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["entry_group_id"] = None
        __props__.__dict__["entry_id"] = None
        __props__.__dict__["gcs_fileset_spec"] = None
        __props__.__dict__["integrated_system"] = None
        __props__.__dict__["linked_resource"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["schema"] = None
        __props__.__dict__["source_system_timestamps"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["usage_signal"] = None
        __props__.__dict__["user_specified_system"] = None
        __props__.__dict__["user_specified_type"] = None
        return Entry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bigqueryDateShardedSpec")
    def bigquery_date_sharded_spec(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse']:
        """
        Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        """
        return pulumi.get(self, "bigquery_date_sharded_spec")

    @property
    @pulumi.getter(name="bigqueryTableSpec")
    def bigquery_table_spec(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse']:
        """
        Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
        """
        return pulumi.get(self, "bigquery_table_spec")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="entryGroupId")
    def entry_group_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "entry_group_id")

    @property
    @pulumi.getter(name="entryId")
    def entry_id(self) -> pulumi.Output[str]:
        """
        Required. The id of the entry to create.
        """
        return pulumi.get(self, "entry_id")

    @property
    @pulumi.getter(name="gcsFilesetSpec")
    def gcs_fileset_spec(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse']:
        """
        Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
        """
        return pulumi.get(self, "gcs_fileset_spec")

    @property
    @pulumi.getter(name="integratedSystem")
    def integrated_system(self) -> pulumi.Output[str]:
        """
        This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        """
        return pulumi.get(self, "integrated_system")

    @property
    @pulumi.getter(name="linkedResource")
    def linked_resource(self) -> pulumi.Output[str]:
        """
        The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
        """
        return pulumi.get(self, "linked_resource")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Data Catalog resource name of the entry in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id} Note that this Entry and its child resources may not actually be stored in the location in this name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1SchemaResponse']:
        """
        Schema of the entry. An entry might not have any schema attached to it.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="sourceSystemTimestamps")
    def source_system_timestamps(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse']:
        """
        Timestamps about the underlying resource, not about this Data Catalog entry. Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty timestamp.
        """
        return pulumi.get(self, "source_system_timestamps")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the entry. Only used for Entries with types in the EntryType enum.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageSignal")
    def usage_signal(self) -> pulumi.Output['outputs.GoogleCloudDatacatalogV1beta1UsageSignalResponse']:
        """
        Statistics on the usage level of the resource.
        """
        return pulumi.get(self, "usage_signal")

    @property
    @pulumi.getter(name="userSpecifiedSystem")
    def user_specified_system(self) -> pulumi.Output[str]:
        """
        This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        """
        return pulumi.get(self, "user_specified_system")

    @property
    @pulumi.getter(name="userSpecifiedType")
    def user_specified_type(self) -> pulumi.Output[str]:
        """
        Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
        """
        return pulumi.get(self, "user_specified_type")

