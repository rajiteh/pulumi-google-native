# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['HistoryExecution']


class HistoryExecution(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 completion_time: Optional[pulumi.Input[pulumi.InputType['TimestampArgs']]] = None,
                 creation_time: Optional[pulumi.Input[pulumi.InputType['TimestampArgs']]] = None,
                 dimension_definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MatrixDimensionDefinitionArgs']]]]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 history_id: Optional[pulumi.Input[str]] = None,
                 outcome: Optional[pulumi.Input[pulumi.InputType['OutcomeArgs']]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[pulumi.InputType['SpecificationArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 test_execution_matrix_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Execution. The returned Execution will have the id set. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing History does not exist

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TimestampArgs']] completion_time: The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        :param pulumi.Input[pulumi.InputType['TimestampArgs']] creation_time: The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MatrixDimensionDefinitionArgs']]]] dimension_definitions: The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
        :param pulumi.Input[str] execution_id: A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
        :param pulumi.Input[pulumi.InputType['OutcomeArgs']] outcome: Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        :param pulumi.Input[pulumi.InputType['SpecificationArgs']] specification: Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
        :param pulumi.Input[str] state: The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
        :param pulumi.Input[str] test_execution_matrix_id: TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['completion_time'] = completion_time
            __props__['creation_time'] = creation_time
            __props__['dimension_definitions'] = dimension_definitions
            if execution_id is None and not opts.urn:
                raise TypeError("Missing required property 'execution_id'")
            __props__['execution_id'] = execution_id
            if history_id is None and not opts.urn:
                raise TypeError("Missing required property 'history_id'")
            __props__['history_id'] = history_id
            __props__['outcome'] = outcome
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['specification'] = specification
            __props__['state'] = state
            __props__['test_execution_matrix_id'] = test_execution_matrix_id
        super(HistoryExecution, __self__).__init__(
            'gcp-native:toolresults/v1beta3:HistoryExecution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HistoryExecution':
        """
        Get an existing HistoryExecution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["completion_time"] = None
        __props__["creation_time"] = None
        __props__["dimension_definitions"] = None
        __props__["execution_id"] = None
        __props__["outcome"] = None
        __props__["specification"] = None
        __props__["state"] = None
        __props__["test_execution_matrix_id"] = None
        return HistoryExecution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="completionTime")
    def completion_time(self) -> pulumi.Output['outputs.TimestampResponse']:
        """
        The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        """
        return pulumi.get(self, "completion_time")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output['outputs.TimestampResponse']:
        """
        The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dimensionDefinitions")
    def dimension_definitions(self) -> pulumi.Output[Sequence['outputs.MatrixDimensionDefinitionResponse']]:
        """
        The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
        """
        return pulumi.get(self, "dimension_definitions")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Output[str]:
        """
        A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter
    def outcome(self) -> pulumi.Output['outputs.OutcomeResponse']:
        """
        Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        """
        return pulumi.get(self, "outcome")

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Output['outputs.SpecificationResponse']:
        """
        Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
        """
        return pulumi.get(self, "specification")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="testExecutionMatrixId")
    def test_execution_matrix_id(self) -> pulumi.Output[str]:
        """
        TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
        """
        return pulumi.get(self, "test_execution_matrix_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
