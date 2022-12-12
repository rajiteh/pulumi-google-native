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
from ._inputs import *

__all__ = ['DocumentSchemaArgs', 'DocumentSchema']

@pulumi.input_type
class DocumentSchemaArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 document_is_folder: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 property_definitions: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]] = None):
        """
        The set of arguments for constructing a DocumentSchema resource.
        :param pulumi.Input[str] display_name: Name of the schema given by the user. Must be unique per customer.
        :param pulumi.Input[str] description: Schema description.
        :param pulumi.Input[bool] document_is_folder: Document Type, true refers the document is a folder, otherwise it is a typical document.
        :param pulumi.Input[str] name: The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]] property_definitions: Document details.
        """
        pulumi.set(__self__, "display_name", display_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if document_is_folder is not None:
            pulumi.set(__self__, "document_is_folder", document_is_folder)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if property_definitions is not None:
            pulumi.set(__self__, "property_definitions", property_definitions)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Name of the schema given by the user. Must be unique per customer.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Schema description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="documentIsFolder")
    def document_is_folder(self) -> Optional[pulumi.Input[bool]]:
        """
        Document Type, true refers the document is a folder, otherwise it is a typical document.
        """
        return pulumi.get(self, "document_is_folder")

    @document_is_folder.setter
    def document_is_folder(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "document_is_folder", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="propertyDefinitions")
    def property_definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]]:
        """
        Document details.
        """
        return pulumi.get(self, "property_definitions")

    @property_definitions.setter
    def property_definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]]):
        pulumi.set(self, "property_definitions", value)


class DocumentSchema(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 document_is_folder: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 property_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]]] = None,
                 __props__=None):
        """
        Creates a document schema.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Schema description.
        :param pulumi.Input[str] display_name: Name of the schema given by the user. Must be unique per customer.
        :param pulumi.Input[bool] document_is_folder: Document Type, true refers the document is a folder, otherwise it is a typical document.
        :param pulumi.Input[str] name: The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]] property_definitions: Document details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentSchemaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a document schema.

        :param str resource_name: The name of the resource.
        :param DocumentSchemaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentSchemaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 document_is_folder: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 property_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudContentwarehouseV1PropertyDefinitionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DocumentSchemaArgs.__new__(DocumentSchemaArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["document_is_folder"] = document_is_folder
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["property_definitions"] = property_definitions
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DocumentSchema, __self__).__init__(
            'google-native:contentwarehouse/v1:DocumentSchema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DocumentSchema':
        """
        Get an existing DocumentSchema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DocumentSchemaArgs.__new__(DocumentSchemaArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["document_is_folder"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["property_definitions"] = None
        __props__.__dict__["update_time"] = None
        return DocumentSchema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the document schema is created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Schema description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Name of the schema given by the user. Must be unique per customer.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="documentIsFolder")
    def document_is_folder(self) -> pulumi.Output[bool]:
        """
        Document Type, true refers the document is a folder, otherwise it is a typical document.
        """
        return pulumi.get(self, "document_is_folder")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="propertyDefinitions")
    def property_definitions(self) -> pulumi.Output[Sequence['outputs.GoogleCloudContentwarehouseV1PropertyDefinitionResponse']]:
        """
        Document details.
        """
        return pulumi.get(self, "property_definitions")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the document schema is last updated.
        """
        return pulumi.get(self, "update_time")
