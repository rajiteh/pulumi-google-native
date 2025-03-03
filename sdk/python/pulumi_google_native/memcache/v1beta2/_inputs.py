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
    'GoogleCloudMemcacheV1beta2MaintenancePolicyArgs',
    'InstanceMessageArgs',
    'MemcacheParametersArgs',
    'NodeConfigArgs',
    'TimeOfDayArgs',
    'WeeklyMaintenanceWindowArgs',
]

@pulumi.input_type
class GoogleCloudMemcacheV1beta2MaintenancePolicyArgs:
    def __init__(__self__, *,
                 weekly_maintenance_window: pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]],
                 description: Optional[pulumi.Input[str]] = None):
        """
        Maintenance policy per instance.
        :param pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]] weekly_maintenance_window: Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_maintenance_windows is expected to be one.
        :param pulumi.Input[str] description: Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        """
        pulumi.set(__self__, "weekly_maintenance_window", weekly_maintenance_window)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="weeklyMaintenanceWindow")
    def weekly_maintenance_window(self) -> pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]]:
        """
        Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_maintenance_windows is expected to be one.
        """
        return pulumi.get(self, "weekly_maintenance_window")

    @weekly_maintenance_window.setter
    def weekly_maintenance_window(self, value: pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]]):
        pulumi.set(self, "weekly_maintenance_window", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class InstanceMessageArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input['InstanceMessageCode']] = None,
                 message: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input['InstanceMessageCode'] code: A code that correspond to one type of user-facing message.
        :param pulumi.Input[str] message: Message on memcached instance which will be exposed to users.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input['InstanceMessageCode']]:
        """
        A code that correspond to one type of user-facing message.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input['InstanceMessageCode']]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        Message on memcached instance which will be exposed to users.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)


@pulumi.input_type
class MemcacheParametersArgs:
    def __init__(__self__, *,
                 params: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] params: User defined set of parameters to use in the memcached process.
        """
        if params is not None:
            pulumi.set(__self__, "params", params)

    @property
    @pulumi.getter
    def params(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User defined set of parameters to use in the memcached process.
        """
        return pulumi.get(self, "params")

    @params.setter
    def params(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "params", value)


@pulumi.input_type
class NodeConfigArgs:
    def __init__(__self__, *,
                 cpu_count: pulumi.Input[int],
                 memory_size_mb: pulumi.Input[int]):
        """
        Configuration for a Memcached Node.
        :param pulumi.Input[int] cpu_count: Number of cpus per Memcached node.
        :param pulumi.Input[int] memory_size_mb: Memory size in MiB for each Memcached node.
        """
        pulumi.set(__self__, "cpu_count", cpu_count)
        pulumi.set(__self__, "memory_size_mb", memory_size_mb)

    @property
    @pulumi.getter(name="cpuCount")
    def cpu_count(self) -> pulumi.Input[int]:
        """
        Number of cpus per Memcached node.
        """
        return pulumi.get(self, "cpu_count")

    @cpu_count.setter
    def cpu_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "cpu_count", value)

    @property
    @pulumi.getter(name="memorySizeMb")
    def memory_size_mb(self) -> pulumi.Input[int]:
        """
        Memory size in MiB for each Memcached node.
        """
        return pulumi.get(self, "memory_size_mb")

    @memory_size_mb.setter
    def memory_size_mb(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory_size_mb", value)


@pulumi.input_type
class TimeOfDayArgs:
    def __init__(__self__, *,
                 hours: Optional[pulumi.Input[int]] = None,
                 minutes: Optional[pulumi.Input[int]] = None,
                 nanos: Optional[pulumi.Input[int]] = None,
                 seconds: Optional[pulumi.Input[int]] = None):
        """
        Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
        :param pulumi.Input[int] hours: Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
        :param pulumi.Input[int] minutes: Minutes of hour of day. Must be from 0 to 59.
        :param pulumi.Input[int] nanos: Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        :param pulumi.Input[int] seconds: Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
        """
        if hours is not None:
            pulumi.set(__self__, "hours", hours)
        if minutes is not None:
            pulumi.set(__self__, "minutes", minutes)
        if nanos is not None:
            pulumi.set(__self__, "nanos", nanos)
        if seconds is not None:
            pulumi.set(__self__, "seconds", seconds)

    @property
    @pulumi.getter
    def hours(self) -> Optional[pulumi.Input[int]]:
        """
        Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
        """
        return pulumi.get(self, "hours")

    @hours.setter
    def hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hours", value)

    @property
    @pulumi.getter
    def minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Minutes of hour of day. Must be from 0 to 59.
        """
        return pulumi.get(self, "minutes")

    @minutes.setter
    def minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minutes", value)

    @property
    @pulumi.getter
    def nanos(self) -> Optional[pulumi.Input[int]]:
        """
        Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        """
        return pulumi.get(self, "nanos")

    @nanos.setter
    def nanos(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nanos", value)

    @property
    @pulumi.getter
    def seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
        """
        return pulumi.get(self, "seconds")

    @seconds.setter
    def seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seconds", value)


@pulumi.input_type
class WeeklyMaintenanceWindowArgs:
    def __init__(__self__, *,
                 day: pulumi.Input['WeeklyMaintenanceWindowDay'],
                 duration: pulumi.Input[str],
                 start_time: pulumi.Input['TimeOfDayArgs']):
        """
        Time window specified for weekly operations.
        :param pulumi.Input['WeeklyMaintenanceWindowDay'] day: Allows to define schedule that runs specified day of the week.
        :param pulumi.Input[str] duration: Duration of the time window.
        :param pulumi.Input['TimeOfDayArgs'] start_time: Start time of the window in UTC.
        """
        pulumi.set(__self__, "day", day)
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter
    def day(self) -> pulumi.Input['WeeklyMaintenanceWindowDay']:
        """
        Allows to define schedule that runs specified day of the week.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: pulumi.Input['WeeklyMaintenanceWindowDay']):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[str]:
        """
        Duration of the time window.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[str]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input['TimeOfDayArgs']:
        """
        Start time of the window in UTC.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input['TimeOfDayArgs']):
        pulumi.set(self, "start_time", value)


