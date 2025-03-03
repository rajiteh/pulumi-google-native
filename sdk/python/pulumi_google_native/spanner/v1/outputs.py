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
    'BackupInfoResponse',
    'BindingResponse',
    'EncryptionConfigResponse',
    'EncryptionInfoResponse',
    'ExprResponse',
    'FreeInstanceMetadataResponse',
    'ReplicaInfoResponse',
    'RestoreInfoResponse',
    'StatusResponse',
]

@pulumi.output_type
class BackupInfoResponse(dict):
    """
    Information about a backup.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "sourceDatabase":
            suggest = "source_database"
        elif key == "versionTime":
            suggest = "version_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BackupInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BackupInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BackupInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup: str,
                 create_time: str,
                 source_database: str,
                 version_time: str):
        """
        Information about a backup.
        :param str backup: Name of the backup.
        :param str create_time: The time the CreateBackup request was received.
        :param str source_database: Name of the database the backup was created from.
        :param str version_time: The backup contains an externally consistent copy of `source_database` at the timestamp specified by `version_time`. If the CreateBackup request did not specify `version_time`, the `version_time` of the backup is equivalent to the `create_time`.
        """
        pulumi.set(__self__, "backup", backup)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "source_database", source_database)
        pulumi.set(__self__, "version_time", version_time)

    @property
    @pulumi.getter
    def backup(self) -> str:
        """
        Name of the backup.
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the CreateBackup request was received.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="sourceDatabase")
    def source_database(self) -> str:
        """
        Name of the database the backup was created from.
        """
        return pulumi.get(self, "source_database")

    @property
    @pulumi.getter(name="versionTime")
    def version_time(self) -> str:
        """
        The backup contains an externally consistent copy of `source_database` at the timestamp specified by `version_time`. If the CreateBackup request did not specify `version_time`, the `version_time` of the backup is equivalent to the `create_time`.
        """
        return pulumi.get(self, "version_time")


@pulumi.output_type
class BindingResponse(dict):
    """
    Associates `members`, or principals, with a `role`.
    """
    def __init__(__self__, *,
                 condition: 'outputs.ExprResponse',
                 members: Sequence[str],
                 role: str):
        """
        Associates `members`, or principals, with a `role`.
        :param 'ExprResponse' condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param Sequence[str] members: Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
        :param str role: Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.ExprResponse':
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")


@pulumi.output_type
class EncryptionConfigResponse(dict):
    """
    Encryption configuration for a Cloud Spanner database.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "kmsKeyName":
            suggest = "kms_key_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EncryptionConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EncryptionConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EncryptionConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kms_key_name: str):
        """
        Encryption configuration for a Cloud Spanner database.
        :param str kms_key_name: The Cloud KMS key to be used for encrypting and decrypting the database. Values are of the form `projects//locations//keyRings//cryptoKeys/`.
        """
        pulumi.set(__self__, "kms_key_name", kms_key_name)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        The Cloud KMS key to be used for encrypting and decrypting the database. Values are of the form `projects//locations//keyRings//cryptoKeys/`.
        """
        return pulumi.get(self, "kms_key_name")


@pulumi.output_type
class EncryptionInfoResponse(dict):
    """
    Encryption information for a Cloud Spanner database or backup.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "encryptionStatus":
            suggest = "encryption_status"
        elif key == "encryptionType":
            suggest = "encryption_type"
        elif key == "kmsKeyVersion":
            suggest = "kms_key_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EncryptionInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EncryptionInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EncryptionInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 encryption_status: 'outputs.StatusResponse',
                 encryption_type: str,
                 kms_key_version: str):
        """
        Encryption information for a Cloud Spanner database or backup.
        :param 'StatusResponse' encryption_status: If present, the status of a recent encrypt/decrypt call on underlying data for this database or backup. Regardless of status, data is always encrypted at rest.
        :param str encryption_type: The type of encryption.
        :param str kms_key_version: A Cloud KMS key version that is being used to protect the database or backup.
        """
        pulumi.set(__self__, "encryption_status", encryption_status)
        pulumi.set(__self__, "encryption_type", encryption_type)
        pulumi.set(__self__, "kms_key_version", kms_key_version)

    @property
    @pulumi.getter(name="encryptionStatus")
    def encryption_status(self) -> 'outputs.StatusResponse':
        """
        If present, the status of a recent encrypt/decrypt call on underlying data for this database or backup. Regardless of status, data is always encrypted at rest.
        """
        return pulumi.get(self, "encryption_status")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> str:
        """
        The type of encryption.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="kmsKeyVersion")
    def kms_key_version(self) -> str:
        """
        A Cloud KMS key version that is being used to protect the database or backup.
        """
        return pulumi.get(self, "kms_key_version")


@pulumi.output_type
class ExprResponse(dict):
    """
    Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
    """
    def __init__(__self__, *,
                 description: str,
                 expression: str,
                 location: str,
                 title: str):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param str description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param str title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")


@pulumi.output_type
class FreeInstanceMetadataResponse(dict):
    """
    Free instance specific metadata that is kept even after an instance has been upgraded for tracking purposes.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expireBehavior":
            suggest = "expire_behavior"
        elif key == "expireTime":
            suggest = "expire_time"
        elif key == "upgradeTime":
            suggest = "upgrade_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FreeInstanceMetadataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FreeInstanceMetadataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FreeInstanceMetadataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expire_behavior: str,
                 expire_time: str,
                 upgrade_time: str):
        """
        Free instance specific metadata that is kept even after an instance has been upgraded for tracking purposes.
        :param str expire_behavior: Specifies the expiration behavior of a free instance. The default of ExpireBehavior is `REMOVE_AFTER_GRACE_PERIOD`. This can be modified during or after creation, and before expiration.
        :param str expire_time: Timestamp after which the instance will either be upgraded or scheduled for deletion after a grace period. ExpireBehavior is used to choose between upgrading or scheduling the free instance for deletion. This timestamp is set during the creation of a free instance.
        :param str upgrade_time: If present, the timestamp at which the free instance was upgraded to a provisioned instance.
        """
        pulumi.set(__self__, "expire_behavior", expire_behavior)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "upgrade_time", upgrade_time)

    @property
    @pulumi.getter(name="expireBehavior")
    def expire_behavior(self) -> str:
        """
        Specifies the expiration behavior of a free instance. The default of ExpireBehavior is `REMOVE_AFTER_GRACE_PERIOD`. This can be modified during or after creation, and before expiration.
        """
        return pulumi.get(self, "expire_behavior")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        Timestamp after which the instance will either be upgraded or scheduled for deletion after a grace period. ExpireBehavior is used to choose between upgrading or scheduling the free instance for deletion. This timestamp is set during the creation of a free instance.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="upgradeTime")
    def upgrade_time(self) -> str:
        """
        If present, the timestamp at which the free instance was upgraded to a provisioned instance.
        """
        return pulumi.get(self, "upgrade_time")


@pulumi.output_type
class ReplicaInfoResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultLeaderLocation":
            suggest = "default_leader_location"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ReplicaInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ReplicaInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ReplicaInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_leader_location: bool,
                 location: str,
                 type: str):
        """
        :param bool default_leader_location: If true, this location is designated as the default leader location where leader replicas are placed. See the [region types documentation](https://cloud.google.com/spanner/docs/instances#region_types) for more details.
        :param str location: The location of the serving resources, e.g. "us-central1".
        :param str type: The type of replica.
        """
        pulumi.set(__self__, "default_leader_location", default_leader_location)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultLeaderLocation")
    def default_leader_location(self) -> bool:
        """
        If true, this location is designated as the default leader location where leader replicas are placed. See the [region types documentation](https://cloud.google.com/spanner/docs/instances#region_types) for more details.
        """
        return pulumi.get(self, "default_leader_location")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the serving resources, e.g. "us-central1".
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of replica.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class RestoreInfoResponse(dict):
    """
    Information about the database restore.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backupInfo":
            suggest = "backup_info"
        elif key == "sourceType":
            suggest = "source_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RestoreInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RestoreInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RestoreInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup_info: 'outputs.BackupInfoResponse',
                 source_type: str):
        """
        Information about the database restore.
        :param 'BackupInfoResponse' backup_info: Information about the backup used to restore the database. The backup may no longer exist.
        :param str source_type: The type of the restore source.
        """
        pulumi.set(__self__, "backup_info", backup_info)
        pulumi.set(__self__, "source_type", source_type)

    @property
    @pulumi.getter(name="backupInfo")
    def backup_info(self) -> 'outputs.BackupInfoResponse':
        """
        Information about the backup used to restore the database. The backup may no longer exist.
        """
        return pulumi.get(self, "backup_info")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        The type of the restore source.
        """
        return pulumi.get(self, "source_type")


@pulumi.output_type
class StatusResponse(dict):
    """
    The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
    """
    def __init__(__self__, *,
                 code: int,
                 details: Sequence[Mapping[str, str]],
                 message: str):
        """
        The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
        :param int code: The status code, which should be an enum value of google.rpc.Code.
        :param Sequence[Mapping[str, str]] details: A list of messages that carry the error details. There is a common set of message types for APIs to use.
        :param str message: A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> int:
        """
        The status code, which should be an enum value of google.rpc.Code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def details(self) -> Sequence[Mapping[str, str]]:
        """
        A list of messages that carry the error details. There is a common set of message types for APIs to use.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
        """
        return pulumi.get(self, "message")


