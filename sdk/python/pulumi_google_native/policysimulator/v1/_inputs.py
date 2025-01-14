# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'GoogleCloudPolicysimulatorV1ReplayConfigArgs',
]

@pulumi.input_type
class GoogleCloudPolicysimulatorV1ReplayConfigArgs:
    def __init__(__self__, *,
                 log_source: Optional[pulumi.Input['GoogleCloudPolicysimulatorV1ReplayConfigLogSource']] = None,
                 policy_overlay: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The configuration used for a Replay.
        :param pulumi.Input['GoogleCloudPolicysimulatorV1ReplayConfigLogSource'] log_source: The logs to use as input for the Replay.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] policy_overlay: A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        """
        if log_source is not None:
            pulumi.set(__self__, "log_source", log_source)
        if policy_overlay is not None:
            pulumi.set(__self__, "policy_overlay", policy_overlay)

    @property
    @pulumi.getter(name="logSource")
    def log_source(self) -> Optional[pulumi.Input['GoogleCloudPolicysimulatorV1ReplayConfigLogSource']]:
        """
        The logs to use as input for the Replay.
        """
        return pulumi.get(self, "log_source")

    @log_source.setter
    def log_source(self, value: Optional[pulumi.Input['GoogleCloudPolicysimulatorV1ReplayConfigLogSource']]):
        pulumi.set(self, "log_source", value)

    @property
    @pulumi.getter(name="policyOverlay")
    def policy_overlay(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        """
        return pulumi.get(self, "policy_overlay")

    @policy_overlay.setter
    def policy_overlay(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "policy_overlay", value)


