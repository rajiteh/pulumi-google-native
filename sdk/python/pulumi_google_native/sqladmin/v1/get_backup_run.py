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

__all__ = [
    'GetBackupRunResult',
    'AwaitableGetBackupRunResult',
    'get_backup_run',
    'get_backup_run_output',
]

@pulumi.output_type
class GetBackupRunResult:
    def __init__(__self__, backup_kind=None, description=None, disk_encryption_configuration=None, disk_encryption_status=None, end_time=None, enqueued_time=None, error=None, instance=None, kind=None, location=None, self_link=None, start_time=None, status=None, time_zone=None, type=None, window_start_time=None):
        if backup_kind and not isinstance(backup_kind, str):
            raise TypeError("Expected argument 'backup_kind' to be a str")
        pulumi.set(__self__, "backup_kind", backup_kind)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_encryption_configuration and not isinstance(disk_encryption_configuration, dict):
            raise TypeError("Expected argument 'disk_encryption_configuration' to be a dict")
        pulumi.set(__self__, "disk_encryption_configuration", disk_encryption_configuration)
        if disk_encryption_status and not isinstance(disk_encryption_status, dict):
            raise TypeError("Expected argument 'disk_encryption_status' to be a dict")
        pulumi.set(__self__, "disk_encryption_status", disk_encryption_status)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if enqueued_time and not isinstance(enqueued_time, str):
            raise TypeError("Expected argument 'enqueued_time' to be a str")
        pulumi.set(__self__, "enqueued_time", enqueued_time)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        pulumi.set(__self__, "instance", instance)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if window_start_time and not isinstance(window_start_time, str):
            raise TypeError("Expected argument 'window_start_time' to be a str")
        pulumi.set(__self__, "window_start_time", window_start_time)

    @property
    @pulumi.getter(name="backupKind")
    def backup_kind(self) -> str:
        """
        Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
        """
        return pulumi.get(self, "backup_kind")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of this run, only applicable to on-demand backups.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionConfiguration")
    def disk_encryption_configuration(self) -> 'outputs.DiskEncryptionConfigurationResponse':
        """
        Encryption configuration specific to a backup.
        """
        return pulumi.get(self, "disk_encryption_configuration")

    @property
    @pulumi.getter(name="diskEncryptionStatus")
    def disk_encryption_status(self) -> 'outputs.DiskEncryptionStatusResponse':
        """
        Encryption status specific to a backup.
        """
        return pulumi.get(self, "disk_encryption_status")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="enqueuedTime")
    def enqueued_time(self) -> str:
        """
        The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "enqueued_time")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.OperationErrorResponse':
        """
        Information about why the backup operation failed. This is only present if the run has the FAILED status.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def instance(self) -> str:
        """
        Name of the database instance.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        This is always `sql#backupRun`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Location of the backups.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of this run.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        """
        Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="windowStartTime")
    def window_start_time(self) -> str:
        """
        The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "window_start_time")


class AwaitableGetBackupRunResult(GetBackupRunResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupRunResult(
            backup_kind=self.backup_kind,
            description=self.description,
            disk_encryption_configuration=self.disk_encryption_configuration,
            disk_encryption_status=self.disk_encryption_status,
            end_time=self.end_time,
            enqueued_time=self.enqueued_time,
            error=self.error,
            instance=self.instance,
            kind=self.kind,
            location=self.location,
            self_link=self.self_link,
            start_time=self.start_time,
            status=self.status,
            time_zone=self.time_zone,
            type=self.type,
            window_start_time=self.window_start_time)


def get_backup_run(id: Optional[str] = None,
                   instance: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupRunResult:
    """
    Retrieves a resource containing information about a backup run.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['instance'] = instance
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:sqladmin/v1:getBackupRun', __args__, opts=opts, typ=GetBackupRunResult).value

    return AwaitableGetBackupRunResult(
        backup_kind=__ret__.backup_kind,
        description=__ret__.description,
        disk_encryption_configuration=__ret__.disk_encryption_configuration,
        disk_encryption_status=__ret__.disk_encryption_status,
        end_time=__ret__.end_time,
        enqueued_time=__ret__.enqueued_time,
        error=__ret__.error,
        instance=__ret__.instance,
        kind=__ret__.kind,
        location=__ret__.location,
        self_link=__ret__.self_link,
        start_time=__ret__.start_time,
        status=__ret__.status,
        time_zone=__ret__.time_zone,
        type=__ret__.type,
        window_start_time=__ret__.window_start_time)


@_utilities.lift_output_func(get_backup_run)
def get_backup_run_output(id: Optional[pulumi.Input[str]] = None,
                          instance: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupRunResult]:
    """
    Retrieves a resource containing information about a backup run.
    """
    ...
