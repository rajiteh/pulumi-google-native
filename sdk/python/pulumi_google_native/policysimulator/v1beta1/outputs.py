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

__all__ = [
    'GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse',
    'GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse',
    'GoogleTypeDateResponse',
]

@pulumi.output_type
class GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse(dict):
    """
    The configuration used for a Replay.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "logSource":
            suggest = "log_source"
        elif key == "policyOverlay":
            suggest = "policy_overlay"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudPolicysimulatorV1beta1ReplayConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 log_source: str,
                 policy_overlay: Mapping[str, str]):
        """
        The configuration used for a Replay.
        :param str log_source: The logs to use as input for the Replay.
        :param Mapping[str, str] policy_overlay: A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        """
        pulumi.set(__self__, "log_source", log_source)
        pulumi.set(__self__, "policy_overlay", policy_overlay)

    @property
    @pulumi.getter(name="logSource")
    def log_source(self) -> str:
        """
        The logs to use as input for the Replay.
        """
        return pulumi.get(self, "log_source")

    @property
    @pulumi.getter(name="policyOverlay")
    def policy_overlay(self) -> Mapping[str, str]:
        """
        A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        """
        return pulumi.get(self, "policy_overlay")


@pulumi.output_type
class GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse(dict):
    """
    Summary statistics about the replayed log entries.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "differenceCount":
            suggest = "difference_count"
        elif key == "errorCount":
            suggest = "error_count"
        elif key == "logCount":
            suggest = "log_count"
        elif key == "newestDate":
            suggest = "newest_date"
        elif key == "oldestDate":
            suggest = "oldest_date"
        elif key == "unchangedCount":
            suggest = "unchanged_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 difference_count: int,
                 error_count: int,
                 log_count: int,
                 newest_date: 'outputs.GoogleTypeDateResponse',
                 oldest_date: 'outputs.GoogleTypeDateResponse',
                 unchanged_count: int):
        """
        Summary statistics about the replayed log entries.
        :param int difference_count: The number of replayed log entries with a difference between baseline and simulated policies.
        :param int error_count: The number of log entries that could not be replayed.
        :param int log_count: The total number of log entries replayed.
        :param 'GoogleTypeDateResponse' newest_date: The date of the newest log entry replayed.
        :param 'GoogleTypeDateResponse' oldest_date: The date of the oldest log entry replayed.
        :param int unchanged_count: The number of replayed log entries with no difference between baseline and simulated policies.
        """
        pulumi.set(__self__, "difference_count", difference_count)
        pulumi.set(__self__, "error_count", error_count)
        pulumi.set(__self__, "log_count", log_count)
        pulumi.set(__self__, "newest_date", newest_date)
        pulumi.set(__self__, "oldest_date", oldest_date)
        pulumi.set(__self__, "unchanged_count", unchanged_count)

    @property
    @pulumi.getter(name="differenceCount")
    def difference_count(self) -> int:
        """
        The number of replayed log entries with a difference between baseline and simulated policies.
        """
        return pulumi.get(self, "difference_count")

    @property
    @pulumi.getter(name="errorCount")
    def error_count(self) -> int:
        """
        The number of log entries that could not be replayed.
        """
        return pulumi.get(self, "error_count")

    @property
    @pulumi.getter(name="logCount")
    def log_count(self) -> int:
        """
        The total number of log entries replayed.
        """
        return pulumi.get(self, "log_count")

    @property
    @pulumi.getter(name="newestDate")
    def newest_date(self) -> 'outputs.GoogleTypeDateResponse':
        """
        The date of the newest log entry replayed.
        """
        return pulumi.get(self, "newest_date")

    @property
    @pulumi.getter(name="oldestDate")
    def oldest_date(self) -> 'outputs.GoogleTypeDateResponse':
        """
        The date of the oldest log entry replayed.
        """
        return pulumi.get(self, "oldest_date")

    @property
    @pulumi.getter(name="unchangedCount")
    def unchanged_count(self) -> int:
        """
        The number of replayed log entries with no difference between baseline and simulated policies.
        """
        return pulumi.get(self, "unchanged_count")


@pulumi.output_type
class GoogleTypeDateResponse(dict):
    """
    Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
    """
    def __init__(__self__, *,
                 day: int,
                 month: int,
                 year: int):
        """
        Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
        :param int day: Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
        :param int month: Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
        :param int year: Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
        """
        pulumi.set(__self__, "day", day)
        pulumi.set(__self__, "month", month)
        pulumi.set(__self__, "year", year)

    @property
    @pulumi.getter
    def day(self) -> int:
        """
        Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def month(self) -> int:
        """
        Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
        """
        return pulumi.get(self, "month")

    @property
    @pulumi.getter
    def year(self) -> int:
        """
        Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
        """
        return pulumi.get(self, "year")


