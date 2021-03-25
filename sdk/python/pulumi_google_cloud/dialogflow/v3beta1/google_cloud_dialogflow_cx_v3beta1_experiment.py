# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['GoogleCloudDialogflowCxV3beta1Experiment']


class GoogleCloudDialogflowCxV3beta1Experiment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 experiment_length: Optional[pulumi.Input[str]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 result: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 variants_history: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Experiment in the specified Environment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creation time of this experiment.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentDefinitionArgs']] definition: The definition of the experiment.
        :param pulumi.Input[str] description: The human-readable description of the experiment.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
        :param pulumi.Input[str] end_time: End time of this experiment.
        :param pulumi.Input[str] experiment_length: Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
        :param pulumi.Input[str] last_update_time: Last update time of this experiment.
        :param pulumi.Input[str] name: The name of the experiment. Format: projects//locations//agents//environments//experiments/..
        :param pulumi.Input[str] parent: Required. The Agent to create an Environment for. Format: `projects//locations//agents//environments/`.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1ExperimentResultArgs']] result: Inference result of the experiment.
        :param pulumi.Input[str] start_time: Start time of this experiment.
        :param pulumi.Input[str] state: The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3beta1VariantsHistoryArgs']]]] variants_history: The history of updates to the experiment variants.
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

            __props__['create_time'] = create_time
            __props__['definition'] = definition
            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['end_time'] = end_time
            __props__['experiment_length'] = experiment_length
            __props__['last_update_time'] = last_update_time
            __props__['name'] = name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['result'] = result
            __props__['start_time'] = start_time
            __props__['state'] = state
            __props__['variants_history'] = variants_history
        super(GoogleCloudDialogflowCxV3beta1Experiment, __self__).__init__(
            'google-cloud:dialogflow/v3beta1:GoogleCloudDialogflowCxV3beta1Experiment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GoogleCloudDialogflowCxV3beta1Experiment':
        """
        Get an existing GoogleCloudDialogflowCxV3beta1Experiment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GoogleCloudDialogflowCxV3beta1Experiment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
