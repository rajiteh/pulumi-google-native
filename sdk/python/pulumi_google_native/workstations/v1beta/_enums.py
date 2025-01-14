# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'GceRegionalPersistentDiskReclaimPolicy',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class GceRegionalPersistentDiskReclaimPolicy(str, Enum):
    """
    What should happen to the disk after the workstation is deleted. Defaults to DELETE.
    """
    RECLAIM_POLICY_UNSPECIFIED = "RECLAIM_POLICY_UNSPECIFIED"
    """
    Do not use.
    """
    DELETE = "DELETE"
    """
    The persistent disk will be deleted with the workstation.
    """
    RETAIN = "RETAIN"
    """
    The persistent disk will be remain after the workstation is deleted, and the administrator must manually delete the disk.
    """
