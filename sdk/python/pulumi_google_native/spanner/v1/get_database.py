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
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    def __init__(__self__, create_time=None, database_dialect=None, default_leader=None, earliest_version_time=None, enable_drop_protection=None, encryption_config=None, encryption_info=None, name=None, reconciling=None, restore_info=None, state=None, version_retention_period=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if database_dialect and not isinstance(database_dialect, str):
            raise TypeError("Expected argument 'database_dialect' to be a str")
        pulumi.set(__self__, "database_dialect", database_dialect)
        if default_leader and not isinstance(default_leader, str):
            raise TypeError("Expected argument 'default_leader' to be a str")
        pulumi.set(__self__, "default_leader", default_leader)
        if earliest_version_time and not isinstance(earliest_version_time, str):
            raise TypeError("Expected argument 'earliest_version_time' to be a str")
        pulumi.set(__self__, "earliest_version_time", earliest_version_time)
        if enable_drop_protection and not isinstance(enable_drop_protection, bool):
            raise TypeError("Expected argument 'enable_drop_protection' to be a bool")
        pulumi.set(__self__, "enable_drop_protection", enable_drop_protection)
        if encryption_config and not isinstance(encryption_config, dict):
            raise TypeError("Expected argument 'encryption_config' to be a dict")
        pulumi.set(__self__, "encryption_config", encryption_config)
        if encryption_info and not isinstance(encryption_info, list):
            raise TypeError("Expected argument 'encryption_info' to be a list")
        pulumi.set(__self__, "encryption_info", encryption_info)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if reconciling and not isinstance(reconciling, bool):
            raise TypeError("Expected argument 'reconciling' to be a bool")
        pulumi.set(__self__, "reconciling", reconciling)
        if restore_info and not isinstance(restore_info, dict):
            raise TypeError("Expected argument 'restore_info' to be a dict")
        pulumi.set(__self__, "restore_info", restore_info)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if version_retention_period and not isinstance(version_retention_period, str):
            raise TypeError("Expected argument 'version_retention_period' to be a str")
        pulumi.set(__self__, "version_retention_period", version_retention_period)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        If exists, the time at which the database creation started.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="databaseDialect")
    def database_dialect(self) -> str:
        """
        The dialect of the Cloud Spanner Database.
        """
        return pulumi.get(self, "database_dialect")

    @property
    @pulumi.getter(name="defaultLeader")
    def default_leader(self) -> str:
        """
        The read-write region which contains the database's leader replicas. This is the same as the value of default_leader database option set using DatabaseAdmin.CreateDatabase or DatabaseAdmin.UpdateDatabaseDdl. If not explicitly set, this is empty.
        """
        return pulumi.get(self, "default_leader")

    @property
    @pulumi.getter(name="earliestVersionTime")
    def earliest_version_time(self) -> str:
        """
        Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        """
        return pulumi.get(self, "earliest_version_time")

    @property
    @pulumi.getter(name="enableDropProtection")
    def enable_drop_protection(self) -> bool:
        """
        Whether drop protection is enabled for this database. Defaults to false, if not set.
        """
        return pulumi.get(self, "enable_drop_protection")

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> 'outputs.EncryptionConfigResponse':
        """
        For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
        """
        return pulumi.get(self, "encryption_config")

    @property
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> Sequence['outputs.EncryptionInfoResponse']:
        """
        For databases that are using customer managed encryption, this field contains the encryption information for the database, such as all Cloud KMS key versions that are in use. The `encryption_status' field inside of each `EncryptionInfo` is not populated. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
        """
        return pulumi.get(self, "encryption_info")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def reconciling(self) -> bool:
        """
        If true, the database is being updated. If false, there are no ongoing update operations for the database.
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter(name="restoreInfo")
    def restore_info(self) -> 'outputs.RestoreInfoResponse':
        """
        Applicable only for restored databases. Contains information about the restore source.
        """
        return pulumi.get(self, "restore_info")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current database state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="versionRetentionPeriod")
    def version_retention_period(self) -> str:
        """
        The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
        """
        return pulumi.get(self, "version_retention_period")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            create_time=self.create_time,
            database_dialect=self.database_dialect,
            default_leader=self.default_leader,
            earliest_version_time=self.earliest_version_time,
            enable_drop_protection=self.enable_drop_protection,
            encryption_config=self.encryption_config,
            encryption_info=self.encryption_info,
            name=self.name,
            reconciling=self.reconciling,
            restore_info=self.restore_info,
            state=self.state,
            version_retention_period=self.version_retention_period)


def get_database(database_id: Optional[str] = None,
                 instance_id: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Gets the state of a Cloud Spanner database.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    __args__['instanceId'] = instance_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:spanner/v1:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        create_time=__ret__.create_time,
        database_dialect=__ret__.database_dialect,
        default_leader=__ret__.default_leader,
        earliest_version_time=__ret__.earliest_version_time,
        enable_drop_protection=__ret__.enable_drop_protection,
        encryption_config=__ret__.encryption_config,
        encryption_info=__ret__.encryption_info,
        name=__ret__.name,
        reconciling=__ret__.reconciling,
        restore_info=__ret__.restore_info,
        state=__ret__.state,
        version_retention_period=__ret__.version_retention_period)


@_utilities.lift_output_func(get_database)
def get_database_output(database_id: Optional[pulumi.Input[str]] = None,
                        instance_id: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Gets the state of a Cloud Spanner database.
    """
    ...
