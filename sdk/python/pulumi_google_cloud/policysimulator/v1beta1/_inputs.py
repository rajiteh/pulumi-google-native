# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs',
    'GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryArgs',
    'GoogleTypeDateArgs',
]

@pulumi.input_type
class GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs:
    def __init__(__self__, *,
                 log_source: Optional[pulumi.Input[str]] = None,
                 policy_overlay: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The configuration used for a Replay.
        :param pulumi.Input[str] log_source: The logs to use as input for the Replay.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] policy_overlay: A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        """
        if log_source is not None:
            pulumi.set(__self__, "log_source", log_source)
        if policy_overlay is not None:
            pulumi.set(__self__, "policy_overlay", policy_overlay)

    @property
    @pulumi.getter(name="logSource")
    def log_source(self) -> Optional[pulumi.Input[str]]:
        """
        The logs to use as input for the Replay.
        """
        return pulumi.get(self, "log_source")

    @log_source.setter
    def log_source(self, value: Optional[pulumi.Input[str]]):
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


@pulumi.input_type
class GoogleCloudPolicysimulatorV1beta1ReplayResultsSummaryArgs:
    def __init__(__self__, *,
                 difference_count: Optional[pulumi.Input[int]] = None,
                 error_count: Optional[pulumi.Input[int]] = None,
                 log_count: Optional[pulumi.Input[int]] = None,
                 newest_date: Optional[pulumi.Input['GoogleTypeDateArgs']] = None,
                 oldest_date: Optional[pulumi.Input['GoogleTypeDateArgs']] = None,
                 unchanged_count: Optional[pulumi.Input[int]] = None):
        """
        Summary statistics about the replayed log entries.
        :param pulumi.Input[int] difference_count: The number of replayed log entries with a difference between baseline and simulated policies.
        :param pulumi.Input[int] error_count: The number of log entries that could not be replayed.
        :param pulumi.Input[int] log_count: The total number of log entries replayed.
        :param pulumi.Input['GoogleTypeDateArgs'] newest_date: The date of the newest log entry replayed.
        :param pulumi.Input['GoogleTypeDateArgs'] oldest_date: The date of the oldest log entry replayed.
        :param pulumi.Input[int] unchanged_count: The number of replayed log entries with no difference between baseline and simulated policies.
        """
        if difference_count is not None:
            pulumi.set(__self__, "difference_count", difference_count)
        if error_count is not None:
            pulumi.set(__self__, "error_count", error_count)
        if log_count is not None:
            pulumi.set(__self__, "log_count", log_count)
        if newest_date is not None:
            pulumi.set(__self__, "newest_date", newest_date)
        if oldest_date is not None:
            pulumi.set(__self__, "oldest_date", oldest_date)
        if unchanged_count is not None:
            pulumi.set(__self__, "unchanged_count", unchanged_count)

    @property
    @pulumi.getter(name="differenceCount")
    def difference_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of replayed log entries with a difference between baseline and simulated policies.
        """
        return pulumi.get(self, "difference_count")

    @difference_count.setter
    def difference_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "difference_count", value)

    @property
    @pulumi.getter(name="errorCount")
    def error_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of log entries that could not be replayed.
        """
        return pulumi.get(self, "error_count")

    @error_count.setter
    def error_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "error_count", value)

    @property
    @pulumi.getter(name="logCount")
    def log_count(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of log entries replayed.
        """
        return pulumi.get(self, "log_count")

    @log_count.setter
    def log_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "log_count", value)

    @property
    @pulumi.getter(name="newestDate")
    def newest_date(self) -> Optional[pulumi.Input['GoogleTypeDateArgs']]:
        """
        The date of the newest log entry replayed.
        """
        return pulumi.get(self, "newest_date")

    @newest_date.setter
    def newest_date(self, value: Optional[pulumi.Input['GoogleTypeDateArgs']]):
        pulumi.set(self, "newest_date", value)

    @property
    @pulumi.getter(name="oldestDate")
    def oldest_date(self) -> Optional[pulumi.Input['GoogleTypeDateArgs']]:
        """
        The date of the oldest log entry replayed.
        """
        return pulumi.get(self, "oldest_date")

    @oldest_date.setter
    def oldest_date(self, value: Optional[pulumi.Input['GoogleTypeDateArgs']]):
        pulumi.set(self, "oldest_date", value)

    @property
    @pulumi.getter(name="unchangedCount")
    def unchanged_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of replayed log entries with no difference between baseline and simulated policies.
        """
        return pulumi.get(self, "unchanged_count")

    @unchanged_count.setter
    def unchanged_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unchanged_count", value)


@pulumi.input_type
class GoogleTypeDateArgs:
    def __init__(__self__, *,
                 day: Optional[pulumi.Input[int]] = None,
                 month: Optional[pulumi.Input[int]] = None,
                 year: Optional[pulumi.Input[int]] = None):
        """
        Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date Related types are google.type.TimeOfDay and `google.protobuf.Timestamp`.
        :param pulumi.Input[int] day: Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
        :param pulumi.Input[int] month: Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
        :param pulumi.Input[int] year: Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
        """
        if day is not None:
            pulumi.set(__self__, "day", day)
        if month is not None:
            pulumi.set(__self__, "month", month)
        if year is not None:
            pulumi.set(__self__, "year", year)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input[int]]:
        """
        Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def month(self) -> Optional[pulumi.Input[int]]:
        """
        Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
        """
        return pulumi.get(self, "month")

    @month.setter
    def month(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "month", value)

    @property
    @pulumi.getter
    def year(self) -> Optional[pulumi.Input[int]]:
        """
        Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
        """
        return pulumi.get(self, "year")

    @year.setter
    def year(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "year", value)

