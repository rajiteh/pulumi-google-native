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
    'GetSnapshotResult',
    'AwaitableGetSnapshotResult',
    'get_snapshot',
    'get_snapshot_output',
]

@pulumi.output_type
class GetSnapshotResult:
    def __init__(__self__, architecture=None, auto_created=None, chain_name=None, creation_size_bytes=None, creation_timestamp=None, description=None, disk_size_gb=None, download_bytes=None, kind=None, label_fingerprint=None, labels=None, license_codes=None, licenses=None, location_hint=None, name=None, satisfies_pzs=None, self_link=None, snapshot_encryption_key=None, snapshot_type=None, source_disk=None, source_disk_encryption_key=None, source_disk_id=None, source_snapshot_schedule_policy=None, source_snapshot_schedule_policy_id=None, status=None, storage_bytes=None, storage_bytes_status=None, storage_locations=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if auto_created and not isinstance(auto_created, bool):
            raise TypeError("Expected argument 'auto_created' to be a bool")
        pulumi.set(__self__, "auto_created", auto_created)
        if chain_name and not isinstance(chain_name, str):
            raise TypeError("Expected argument 'chain_name' to be a str")
        pulumi.set(__self__, "chain_name", chain_name)
        if creation_size_bytes and not isinstance(creation_size_bytes, str):
            raise TypeError("Expected argument 'creation_size_bytes' to be a str")
        pulumi.set(__self__, "creation_size_bytes", creation_size_bytes)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size_gb and not isinstance(disk_size_gb, str):
            raise TypeError("Expected argument 'disk_size_gb' to be a str")
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if download_bytes and not isinstance(download_bytes, str):
            raise TypeError("Expected argument 'download_bytes' to be a str")
        pulumi.set(__self__, "download_bytes", download_bytes)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if license_codes and not isinstance(license_codes, list):
            raise TypeError("Expected argument 'license_codes' to be a list")
        pulumi.set(__self__, "license_codes", license_codes)
        if licenses and not isinstance(licenses, list):
            raise TypeError("Expected argument 'licenses' to be a list")
        pulumi.set(__self__, "licenses", licenses)
        if location_hint and not isinstance(location_hint, str):
            raise TypeError("Expected argument 'location_hint' to be a str")
        pulumi.set(__self__, "location_hint", location_hint)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if snapshot_encryption_key and not isinstance(snapshot_encryption_key, dict):
            raise TypeError("Expected argument 'snapshot_encryption_key' to be a dict")
        pulumi.set(__self__, "snapshot_encryption_key", snapshot_encryption_key)
        if snapshot_type and not isinstance(snapshot_type, str):
            raise TypeError("Expected argument 'snapshot_type' to be a str")
        pulumi.set(__self__, "snapshot_type", snapshot_type)
        if source_disk and not isinstance(source_disk, str):
            raise TypeError("Expected argument 'source_disk' to be a str")
        pulumi.set(__self__, "source_disk", source_disk)
        if source_disk_encryption_key and not isinstance(source_disk_encryption_key, dict):
            raise TypeError("Expected argument 'source_disk_encryption_key' to be a dict")
        pulumi.set(__self__, "source_disk_encryption_key", source_disk_encryption_key)
        if source_disk_id and not isinstance(source_disk_id, str):
            raise TypeError("Expected argument 'source_disk_id' to be a str")
        pulumi.set(__self__, "source_disk_id", source_disk_id)
        if source_snapshot_schedule_policy and not isinstance(source_snapshot_schedule_policy, str):
            raise TypeError("Expected argument 'source_snapshot_schedule_policy' to be a str")
        pulumi.set(__self__, "source_snapshot_schedule_policy", source_snapshot_schedule_policy)
        if source_snapshot_schedule_policy_id and not isinstance(source_snapshot_schedule_policy_id, str):
            raise TypeError("Expected argument 'source_snapshot_schedule_policy_id' to be a str")
        pulumi.set(__self__, "source_snapshot_schedule_policy_id", source_snapshot_schedule_policy_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if storage_bytes and not isinstance(storage_bytes, str):
            raise TypeError("Expected argument 'storage_bytes' to be a str")
        pulumi.set(__self__, "storage_bytes", storage_bytes)
        if storage_bytes_status and not isinstance(storage_bytes_status, str):
            raise TypeError("Expected argument 'storage_bytes_status' to be a str")
        pulumi.set(__self__, "storage_bytes_status", storage_bytes_status)
        if storage_locations and not isinstance(storage_locations, list):
            raise TypeError("Expected argument 'storage_locations' to be a list")
        pulumi.set(__self__, "storage_locations", storage_locations)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        """
        The architecture of the snapshot. Valid values are ARM64 or X86_64.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="autoCreated")
    def auto_created(self) -> bool:
        """
        Set to true if snapshots are automatically created by applying resource policy on the target disk.
        """
        return pulumi.get(self, "auto_created")

    @property
    @pulumi.getter(name="chainName")
    def chain_name(self) -> str:
        """
        Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
        """
        return pulumi.get(self, "chain_name")

    @property
    @pulumi.getter(name="creationSizeBytes")
    def creation_size_bytes(self) -> str:
        """
        Size in bytes of the snapshot at creation time.
        """
        return pulumi.get(self, "creation_size_bytes")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> str:
        """
        Size of the source disk, specified in GB.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="downloadBytes")
    def download_bytes(self) -> str:
        """
        Number of bytes downloaded to restore a snapshot to a disk.
        """
        return pulumi.get(self, "download_bytes")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#snapshot for Snapshot resources.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="licenseCodes")
    def license_codes(self) -> Sequence[str]:
        """
        Integer license codes indicating which licenses are attached to this snapshot.
        """
        return pulumi.get(self, "license_codes")

    @property
    @pulumi.getter
    def licenses(self) -> Sequence[str]:
        """
        A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
        """
        return pulumi.get(self, "licenses")

    @property
    @pulumi.getter(name="locationHint")
    def location_hint(self) -> str:
        """
        An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
        """
        return pulumi.get(self, "location_hint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="snapshotEncryptionKey")
    def snapshot_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
        """
        return pulumi.get(self, "snapshot_encryption_key")

    @property
    @pulumi.getter(name="snapshotType")
    def snapshot_type(self) -> str:
        """
        Indicates the type of the snapshot.
        """
        return pulumi.get(self, "snapshot_type")

    @property
    @pulumi.getter(name="sourceDisk")
    def source_disk(self) -> str:
        """
        The source disk used to create this snapshot.
        """
        return pulumi.get(self, "source_disk")

    @property
    @pulumi.getter(name="sourceDiskEncryptionKey")
    def source_disk_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        """
        return pulumi.get(self, "source_disk_encryption_key")

    @property
    @pulumi.getter(name="sourceDiskId")
    def source_disk_id(self) -> str:
        """
        The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
        """
        return pulumi.get(self, "source_disk_id")

    @property
    @pulumi.getter(name="sourceSnapshotSchedulePolicy")
    def source_snapshot_schedule_policy(self) -> str:
        """
        URL of the resource policy which created this scheduled snapshot.
        """
        return pulumi.get(self, "source_snapshot_schedule_policy")

    @property
    @pulumi.getter(name="sourceSnapshotSchedulePolicyId")
    def source_snapshot_schedule_policy_id(self) -> str:
        """
        ID of the resource policy which created this scheduled snapshot.
        """
        return pulumi.get(self, "source_snapshot_schedule_policy_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageBytes")
    def storage_bytes(self) -> str:
        """
        A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
        """
        return pulumi.get(self, "storage_bytes")

    @property
    @pulumi.getter(name="storageBytesStatus")
    def storage_bytes_status(self) -> str:
        """
        An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
        """
        return pulumi.get(self, "storage_bytes_status")

    @property
    @pulumi.getter(name="storageLocations")
    def storage_locations(self) -> Sequence[str]:
        """
        Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
        """
        return pulumi.get(self, "storage_locations")


class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            architecture=self.architecture,
            auto_created=self.auto_created,
            chain_name=self.chain_name,
            creation_size_bytes=self.creation_size_bytes,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            disk_size_gb=self.disk_size_gb,
            download_bytes=self.download_bytes,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            license_codes=self.license_codes,
            licenses=self.licenses,
            location_hint=self.location_hint,
            name=self.name,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            snapshot_encryption_key=self.snapshot_encryption_key,
            snapshot_type=self.snapshot_type,
            source_disk=self.source_disk,
            source_disk_encryption_key=self.source_disk_encryption_key,
            source_disk_id=self.source_disk_id,
            source_snapshot_schedule_policy=self.source_snapshot_schedule_policy,
            source_snapshot_schedule_policy_id=self.source_snapshot_schedule_policy_id,
            status=self.status,
            storage_bytes=self.storage_bytes,
            storage_bytes_status=self.storage_bytes_status,
            storage_locations=self.storage_locations)


def get_snapshot(project: Optional[str] = None,
                 snapshot: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotResult:
    """
    Returns the specified Snapshot resource.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['snapshot'] = snapshot
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getSnapshot', __args__, opts=opts, typ=GetSnapshotResult).value

    return AwaitableGetSnapshotResult(
        architecture=__ret__.architecture,
        auto_created=__ret__.auto_created,
        chain_name=__ret__.chain_name,
        creation_size_bytes=__ret__.creation_size_bytes,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        disk_size_gb=__ret__.disk_size_gb,
        download_bytes=__ret__.download_bytes,
        kind=__ret__.kind,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        license_codes=__ret__.license_codes,
        licenses=__ret__.licenses,
        location_hint=__ret__.location_hint,
        name=__ret__.name,
        satisfies_pzs=__ret__.satisfies_pzs,
        self_link=__ret__.self_link,
        snapshot_encryption_key=__ret__.snapshot_encryption_key,
        snapshot_type=__ret__.snapshot_type,
        source_disk=__ret__.source_disk,
        source_disk_encryption_key=__ret__.source_disk_encryption_key,
        source_disk_id=__ret__.source_disk_id,
        source_snapshot_schedule_policy=__ret__.source_snapshot_schedule_policy,
        source_snapshot_schedule_policy_id=__ret__.source_snapshot_schedule_policy_id,
        status=__ret__.status,
        storage_bytes=__ret__.storage_bytes,
        storage_bytes_status=__ret__.storage_bytes_status,
        storage_locations=__ret__.storage_locations)


@_utilities.lift_output_func(get_snapshot)
def get_snapshot_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                        snapshot: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotResult]:
    """
    Returns the specified Snapshot resource.
    """
    ...
