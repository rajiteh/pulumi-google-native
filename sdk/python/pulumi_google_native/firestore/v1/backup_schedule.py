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

__all__ = ['BackupScheduleArgs', 'BackupSchedule']

@pulumi.input_type
class BackupScheduleArgs:
    def __init__(__self__, *,
                 database_id: pulumi.Input[str],
                 daily_recurrence: Optional[pulumi.Input['GoogleFirestoreAdminV1DailyRecurrenceArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']] = None):
        """
        The set of arguments for constructing a BackupSchedule resource.
        :param pulumi.Input['GoogleFirestoreAdminV1DailyRecurrenceArgs'] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to the creation time of the backup should the backup be deleted, i.e. keep backups for 7 days.
        :param pulumi.Input['GoogleFirestoreAdminV1WeeklyRecurrenceArgs'] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
        """
        pulumi.set(__self__, "database_id", database_id)
        if daily_recurrence is not None:
            pulumi.set(__self__, "daily_recurrence", daily_recurrence)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if retention is not None:
            pulumi.set(__self__, "retention", retention)
        if weekly_recurrence is not None:
            pulumi.set(__self__, "weekly_recurrence", weekly_recurrence)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> Optional[pulumi.Input['GoogleFirestoreAdminV1DailyRecurrenceArgs']]:
        """
        For a schedule that runs daily at a specified time.
        """
        return pulumi.get(self, "daily_recurrence")

    @daily_recurrence.setter
    def daily_recurrence(self, value: Optional[pulumi.Input['GoogleFirestoreAdminV1DailyRecurrenceArgs']]):
        pulumi.set(self, "daily_recurrence", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def retention(self) -> Optional[pulumi.Input[str]]:
        """
        At what relative time in the future, compared to the creation time of the backup should the backup be deleted, i.e. keep backups for 7 days.
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> Optional[pulumi.Input['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']]:
        """
        For a schedule that runs weekly on a specific day and time.
        """
        return pulumi.get(self, "weekly_recurrence")

    @weekly_recurrence.setter
    def weekly_recurrence(self, value: Optional[pulumi.Input['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']]):
        pulumi.set(self, "weekly_recurrence", value)


class BackupSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1DailyRecurrenceArgs']]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']]] = None,
                 __props__=None):
        """
        Creates a backup schedule on a database. At most two backup schedules can be configured on a database, one daily backup schedule with retention up to 7 days and one weekly backup schedule with retention up to 14 weeks.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1DailyRecurrenceArgs']] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to the creation time of the backup should the backup be deleted, i.e. keep backups for 7 days.
        :param pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backup schedule on a database. At most two backup schedules can be configured on a database, one daily backup schedule with retention up to 7 days and one weekly backup schedule with retention up to 14 weeks.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param BackupScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1DailyRecurrenceArgs']]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1WeeklyRecurrenceArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupScheduleArgs.__new__(BackupScheduleArgs)

            __props__.__dict__["daily_recurrence"] = daily_recurrence
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["project"] = project
            __props__.__dict__["retention"] = retention
            __props__.__dict__["weekly_recurrence"] = weekly_recurrence
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["database_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(BackupSchedule, __self__).__init__(
            'google-native:firestore/v1:BackupSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackupSchedule':
        """
        Get an existing BackupSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BackupScheduleArgs.__new__(BackupScheduleArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["daily_recurrence"] = None
        __props__.__dict__["database_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["retention"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["weekly_recurrence"] = None
        return BackupSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp at which this backup schedule was created and effective since. No backups will be created for this schedule before this time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> pulumi.Output['outputs.GoogleFirestoreAdminV1DailyRecurrenceResponse']:
        """
        For a schedule that runs daily at a specified time.
        """
        return pulumi.get(self, "daily_recurrence")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique backup schedule identifier across all locations and databases for the given project. This will be auto-assigned. Format is `projects/{project}/databases/{database}/backupSchedules/{backup_schedule}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Output[str]:
        """
        At what relative time in the future, compared to the creation time of the backup should the backup be deleted, i.e. keep backups for 7 days.
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp at which this backup schedule was most recently updated. When a backup schedule is first created, this is the same as create_time.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> pulumi.Output['outputs.GoogleFirestoreAdminV1WeeklyRecurrenceResponse']:
        """
        For a schedule that runs weekly on a specific day and time.
        """
        return pulumi.get(self, "weekly_recurrence")
