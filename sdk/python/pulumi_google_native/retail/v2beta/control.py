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

__all__ = ['ControlArgs', 'Control']

@pulumi.input_type
class ControlArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 control_id: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 solution_types: pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]],
                 facet_spec: Optional[pulumi.Input['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input['GoogleCloudRetailV2betaRuleArgs']] = None,
                 search_solution_use_case: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]]] = None):
        """
        The set of arguments for constructing a Control resource.
        :param pulumi.Input[str] control_id: Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        :param pulumi.Input[str] display_name: The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        :param pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]] solution_types: Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        :param pulumi.Input['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs'] facet_spec: A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
        :param pulumi.Input[str] name: Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
        :param pulumi.Input['GoogleCloudRetailV2betaRuleArgs'] rule: A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        :param pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]] search_solution_use_case: Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "control_id", control_id)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "solution_types", solution_types)
        if facet_spec is not None:
            pulumi.set(__self__, "facet_spec", facet_spec)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rule is not None:
            pulumi.set(__self__, "rule", rule)
        if search_solution_use_case is not None:
            pulumi.set(__self__, "search_solution_use_case", search_solution_use_case)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="controlId")
    def control_id(self) -> pulumi.Input[str]:
        """
        Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        """
        return pulumi.get(self, "control_id")

    @control_id.setter
    def control_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "control_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="solutionTypes")
    def solution_types(self) -> pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]]:
        """
        Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "solution_types")

    @solution_types.setter
    def solution_types(self, value: pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]]):
        pulumi.set(self, "solution_types", value)

    @property
    @pulumi.getter(name="facetSpec")
    def facet_spec(self) -> Optional[pulumi.Input['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']]:
        """
        A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
        """
        return pulumi.get(self, "facet_spec")

    @facet_spec.setter
    def facet_spec(self, value: Optional[pulumi.Input['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']]):
        pulumi.set(self, "facet_spec", value)

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
        Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
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
    @pulumi.getter
    def rule(self) -> Optional[pulumi.Input['GoogleCloudRetailV2betaRuleArgs']]:
        """
        A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: Optional[pulumi.Input['GoogleCloudRetailV2betaRuleArgs']]):
        pulumi.set(self, "rule", value)

    @property
    @pulumi.getter(name="searchSolutionUseCase")
    def search_solution_use_case(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]]]:
        """
        Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
        """
        return pulumi.get(self, "search_solution_use_case")

    @search_solution_use_case.setter
    def search_solution_use_case(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]]]):
        pulumi.set(self, "search_solution_use_case", value)


class Control(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 control_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 facet_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaRuleArgs']]] = None,
                 search_solution_use_case: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]]] = None,
                 solution_types: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]]] = None,
                 __props__=None):
        """
        Creates a Control. If the Control to create already exists, an ALREADY_EXISTS error is returned.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] control_id: Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        :param pulumi.Input[str] display_name: The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']] facet_spec: A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
        :param pulumi.Input[str] name: Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
        :param pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaRuleArgs']] rule: A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        :param pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]] search_solution_use_case: Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
        :param pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]] solution_types: Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ControlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Control. If the Control to create already exists, an ALREADY_EXISTS error is returned.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ControlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ControlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 control_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 facet_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaSearchRequestFacetSpecArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRetailV2betaRuleArgs']]] = None,
                 search_solution_use_case: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSearchSolutionUseCaseItem']]]] = None,
                 solution_types: Optional[pulumi.Input[Sequence[pulumi.Input['ControlSolutionTypesItem']]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ControlArgs.__new__(ControlArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            if control_id is None and not opts.urn:
                raise TypeError("Missing required property 'control_id'")
            __props__.__dict__["control_id"] = control_id
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["facet_spec"] = facet_spec
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["rule"] = rule
            __props__.__dict__["search_solution_use_case"] = search_solution_use_case
            if solution_types is None and not opts.urn:
                raise TypeError("Missing required property 'solution_types'")
            __props__.__dict__["solution_types"] = solution_types
            __props__.__dict__["associated_serving_config_ids"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["catalog_id", "control_id", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Control, __self__).__init__(
            'google-native:retail/v2beta:Control',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Control':
        """
        Get an existing Control resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ControlArgs.__new__(ControlArgs)

        __props__.__dict__["associated_serving_config_ids"] = None
        __props__.__dict__["catalog_id"] = None
        __props__.__dict__["control_id"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["facet_spec"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["rule"] = None
        __props__.__dict__["search_solution_use_case"] = None
        __props__.__dict__["solution_types"] = None
        return Control(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associatedServingConfigIds")
    def associated_serving_config_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
        """
        return pulumi.get(self, "associated_serving_config_ids")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="controlId")
    def control_id(self) -> pulumi.Output[str]:
        """
        Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        """
        return pulumi.get(self, "control_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="facetSpec")
    def facet_spec(self) -> pulumi.Output['outputs.GoogleCloudRetailV2betaSearchRequestFacetSpecResponse']:
        """
        A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
        """
        return pulumi.get(self, "facet_spec")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Output['outputs.GoogleCloudRetailV2betaRuleResponse']:
        """
        A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        """
        return pulumi.get(self, "rule")

    @property
    @pulumi.getter(name="searchSolutionUseCase")
    def search_solution_use_case(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
        """
        return pulumi.get(self, "search_solution_use_case")

    @property
    @pulumi.getter(name="solutionTypes")
    def solution_types(self) -> pulumi.Output[Sequence[str]]:
        """
        Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        """
        return pulumi.get(self, "solution_types")

