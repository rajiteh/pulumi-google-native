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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, advanced_machine_features=None, can_ip_forward=None, confidential_instance_config=None, cpu_platform=None, creation_timestamp=None, deletion_protection=None, description=None, disks=None, display_device=None, erase_windows_vss_signature=None, fingerprint=None, guest_accelerators=None, hostname=None, instance_encryption_key=None, key_revocation_action_type=None, kind=None, label_fingerprint=None, labels=None, last_start_timestamp=None, last_stop_timestamp=None, last_suspended_timestamp=None, machine_type=None, metadata=None, min_cpu_platform=None, name=None, network_interfaces=None, network_performance_config=None, params=None, post_key_revocation_action_type=None, preserved_state_size_gb=None, private_ipv6_google_access=None, reservation_affinity=None, resource_policies=None, resource_status=None, satisfies_pzs=None, scheduling=None, secure_tags=None, self_link=None, self_link_with_id=None, service_accounts=None, service_integration_specs=None, shielded_instance_config=None, shielded_instance_integrity_policy=None, shielded_vm_config=None, shielded_vm_integrity_policy=None, source_machine_image=None, source_machine_image_encryption_key=None, start_restricted=None, status=None, status_message=None, tags=None, upcoming_maintenance=None, zone=None):
        if advanced_machine_features and not isinstance(advanced_machine_features, dict):
            raise TypeError("Expected argument 'advanced_machine_features' to be a dict")
        pulumi.set(__self__, "advanced_machine_features", advanced_machine_features)
        if can_ip_forward and not isinstance(can_ip_forward, bool):
            raise TypeError("Expected argument 'can_ip_forward' to be a bool")
        pulumi.set(__self__, "can_ip_forward", can_ip_forward)
        if confidential_instance_config and not isinstance(confidential_instance_config, dict):
            raise TypeError("Expected argument 'confidential_instance_config' to be a dict")
        pulumi.set(__self__, "confidential_instance_config", confidential_instance_config)
        if cpu_platform and not isinstance(cpu_platform, str):
            raise TypeError("Expected argument 'cpu_platform' to be a str")
        pulumi.set(__self__, "cpu_platform", cpu_platform)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disks and not isinstance(disks, list):
            raise TypeError("Expected argument 'disks' to be a list")
        pulumi.set(__self__, "disks", disks)
        if display_device and not isinstance(display_device, dict):
            raise TypeError("Expected argument 'display_device' to be a dict")
        pulumi.set(__self__, "display_device", display_device)
        if erase_windows_vss_signature and not isinstance(erase_windows_vss_signature, bool):
            raise TypeError("Expected argument 'erase_windows_vss_signature' to be a bool")
        pulumi.set(__self__, "erase_windows_vss_signature", erase_windows_vss_signature)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if guest_accelerators and not isinstance(guest_accelerators, list):
            raise TypeError("Expected argument 'guest_accelerators' to be a list")
        pulumi.set(__self__, "guest_accelerators", guest_accelerators)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if instance_encryption_key and not isinstance(instance_encryption_key, dict):
            raise TypeError("Expected argument 'instance_encryption_key' to be a dict")
        pulumi.set(__self__, "instance_encryption_key", instance_encryption_key)
        if key_revocation_action_type and not isinstance(key_revocation_action_type, str):
            raise TypeError("Expected argument 'key_revocation_action_type' to be a str")
        pulumi.set(__self__, "key_revocation_action_type", key_revocation_action_type)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_start_timestamp and not isinstance(last_start_timestamp, str):
            raise TypeError("Expected argument 'last_start_timestamp' to be a str")
        pulumi.set(__self__, "last_start_timestamp", last_start_timestamp)
        if last_stop_timestamp and not isinstance(last_stop_timestamp, str):
            raise TypeError("Expected argument 'last_stop_timestamp' to be a str")
        pulumi.set(__self__, "last_stop_timestamp", last_stop_timestamp)
        if last_suspended_timestamp and not isinstance(last_suspended_timestamp, str):
            raise TypeError("Expected argument 'last_suspended_timestamp' to be a str")
        pulumi.set(__self__, "last_suspended_timestamp", last_suspended_timestamp)
        if machine_type and not isinstance(machine_type, str):
            raise TypeError("Expected argument 'machine_type' to be a str")
        pulumi.set(__self__, "machine_type", machine_type)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if min_cpu_platform and not isinstance(min_cpu_platform, str):
            raise TypeError("Expected argument 'min_cpu_platform' to be a str")
        pulumi.set(__self__, "min_cpu_platform", min_cpu_platform)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        pulumi.set(__self__, "network_interfaces", network_interfaces)
        if network_performance_config and not isinstance(network_performance_config, dict):
            raise TypeError("Expected argument 'network_performance_config' to be a dict")
        pulumi.set(__self__, "network_performance_config", network_performance_config)
        if params and not isinstance(params, dict):
            raise TypeError("Expected argument 'params' to be a dict")
        pulumi.set(__self__, "params", params)
        if post_key_revocation_action_type and not isinstance(post_key_revocation_action_type, str):
            raise TypeError("Expected argument 'post_key_revocation_action_type' to be a str")
        pulumi.set(__self__, "post_key_revocation_action_type", post_key_revocation_action_type)
        if preserved_state_size_gb and not isinstance(preserved_state_size_gb, str):
            raise TypeError("Expected argument 'preserved_state_size_gb' to be a str")
        pulumi.set(__self__, "preserved_state_size_gb", preserved_state_size_gb)
        if private_ipv6_google_access and not isinstance(private_ipv6_google_access, str):
            raise TypeError("Expected argument 'private_ipv6_google_access' to be a str")
        pulumi.set(__self__, "private_ipv6_google_access", private_ipv6_google_access)
        if reservation_affinity and not isinstance(reservation_affinity, dict):
            raise TypeError("Expected argument 'reservation_affinity' to be a dict")
        pulumi.set(__self__, "reservation_affinity", reservation_affinity)
        if resource_policies and not isinstance(resource_policies, list):
            raise TypeError("Expected argument 'resource_policies' to be a list")
        pulumi.set(__self__, "resource_policies", resource_policies)
        if resource_status and not isinstance(resource_status, dict):
            raise TypeError("Expected argument 'resource_status' to be a dict")
        pulumi.set(__self__, "resource_status", resource_status)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if scheduling and not isinstance(scheduling, dict):
            raise TypeError("Expected argument 'scheduling' to be a dict")
        pulumi.set(__self__, "scheduling", scheduling)
        if secure_tags and not isinstance(secure_tags, list):
            raise TypeError("Expected argument 'secure_tags' to be a list")
        pulumi.set(__self__, "secure_tags", secure_tags)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if service_accounts and not isinstance(service_accounts, list):
            raise TypeError("Expected argument 'service_accounts' to be a list")
        pulumi.set(__self__, "service_accounts", service_accounts)
        if service_integration_specs and not isinstance(service_integration_specs, dict):
            raise TypeError("Expected argument 'service_integration_specs' to be a dict")
        pulumi.set(__self__, "service_integration_specs", service_integration_specs)
        if shielded_instance_config and not isinstance(shielded_instance_config, dict):
            raise TypeError("Expected argument 'shielded_instance_config' to be a dict")
        pulumi.set(__self__, "shielded_instance_config", shielded_instance_config)
        if shielded_instance_integrity_policy and not isinstance(shielded_instance_integrity_policy, dict):
            raise TypeError("Expected argument 'shielded_instance_integrity_policy' to be a dict")
        pulumi.set(__self__, "shielded_instance_integrity_policy", shielded_instance_integrity_policy)
        if shielded_vm_config and not isinstance(shielded_vm_config, dict):
            raise TypeError("Expected argument 'shielded_vm_config' to be a dict")
        pulumi.set(__self__, "shielded_vm_config", shielded_vm_config)
        if shielded_vm_integrity_policy and not isinstance(shielded_vm_integrity_policy, dict):
            raise TypeError("Expected argument 'shielded_vm_integrity_policy' to be a dict")
        pulumi.set(__self__, "shielded_vm_integrity_policy", shielded_vm_integrity_policy)
        if source_machine_image and not isinstance(source_machine_image, str):
            raise TypeError("Expected argument 'source_machine_image' to be a str")
        pulumi.set(__self__, "source_machine_image", source_machine_image)
        if source_machine_image_encryption_key and not isinstance(source_machine_image_encryption_key, dict):
            raise TypeError("Expected argument 'source_machine_image_encryption_key' to be a dict")
        pulumi.set(__self__, "source_machine_image_encryption_key", source_machine_image_encryption_key)
        if start_restricted and not isinstance(start_restricted, bool):
            raise TypeError("Expected argument 'start_restricted' to be a bool")
        pulumi.set(__self__, "start_restricted", start_restricted)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if upcoming_maintenance and not isinstance(upcoming_maintenance, dict):
            raise TypeError("Expected argument 'upcoming_maintenance' to be a dict")
        pulumi.set(__self__, "upcoming_maintenance", upcoming_maintenance)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="advancedMachineFeatures")
    def advanced_machine_features(self) -> 'outputs.AdvancedMachineFeaturesResponse':
        """
        Controls for advanced machine-related behavior features.
        """
        return pulumi.get(self, "advanced_machine_features")

    @property
    @pulumi.getter(name="canIpForward")
    def can_ip_forward(self) -> bool:
        """
        Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
        """
        return pulumi.get(self, "can_ip_forward")

    @property
    @pulumi.getter(name="confidentialInstanceConfig")
    def confidential_instance_config(self) -> 'outputs.ConfidentialInstanceConfigResponse':
        return pulumi.get(self, "confidential_instance_config")

    @property
    @pulumi.getter(name="cpuPlatform")
    def cpu_platform(self) -> str:
        """
        The CPU platform used by this instance.
        """
        return pulumi.get(self, "cpu_platform")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> bool:
        """
        Whether the resource should be protected against deletion.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disks(self) -> Sequence['outputs.AttachedDiskResponse']:
        """
        Array of disks associated with this instance. Persistent disks must be created before you can assign them.
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter(name="displayDevice")
    def display_device(self) -> 'outputs.DisplayDeviceResponse':
        """
        Enables display device for the instance.
        """
        return pulumi.get(self, "display_device")

    @property
    @pulumi.getter(name="eraseWindowsVssSignature")
    def erase_windows_vss_signature(self) -> bool:
        """
        Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
        """
        return pulumi.get(self, "erase_windows_vss_signature")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="guestAccelerators")
    def guest_accelerators(self) -> Sequence['outputs.AcceleratorConfigResponse']:
        """
        A list of the type and count of accelerator cards attached to the instance.
        """
        return pulumi.get(self, "guest_accelerators")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="instanceEncryptionKey")
    def instance_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Encrypts suspended data for an instance with a customer-managed encryption key. If you are creating a new instance, this field will encrypt the local SSD and in-memory contents of the instance during the suspend operation. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key during the suspend operation.
        """
        return pulumi.get(self, "instance_encryption_key")

    @property
    @pulumi.getter(name="keyRevocationActionType")
    def key_revocation_action_type(self) -> str:
        """
        KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
        """
        return pulumi.get(self, "key_revocation_action_type")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#instance for instances.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels to apply to this instance. These can be later modified by the setLabels method.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastStartTimestamp")
    def last_start_timestamp(self) -> str:
        """
        Last start timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_start_timestamp")

    @property
    @pulumi.getter(name="lastStopTimestamp")
    def last_stop_timestamp(self) -> str:
        """
        Last stop timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_stop_timestamp")

    @property
    @pulumi.getter(name="lastSuspendedTimestamp")
    def last_suspended_timestamp(self) -> str:
        """
        Last suspended timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_suspended_timestamp")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> str:
        """
        Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
        """
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.MetadataResponse':
        """
        The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="minCpuPlatform")
    def min_cpu_platform(self) -> str:
        """
        Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
        """
        return pulumi.get(self, "min_cpu_platform")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Sequence['outputs.NetworkInterfaceResponse']:
        """
        An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="networkPerformanceConfig")
    def network_performance_config(self) -> 'outputs.NetworkPerformanceConfigResponse':
        return pulumi.get(self, "network_performance_config")

    @property
    @pulumi.getter
    def params(self) -> 'outputs.InstanceParamsResponse':
        """
        Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
        """
        return pulumi.get(self, "params")

    @property
    @pulumi.getter(name="postKeyRevocationActionType")
    def post_key_revocation_action_type(self) -> str:
        """
        PostKeyRevocationActionType of the instance.
        """
        return pulumi.get(self, "post_key_revocation_action_type")

    @property
    @pulumi.getter(name="preservedStateSizeGb")
    def preserved_state_size_gb(self) -> str:
        """
        Total amount of preserved state for SUSPENDED instances. Read-only in the api.
        """
        return pulumi.get(self, "preserved_state_size_gb")

    @property
    @pulumi.getter(name="privateIpv6GoogleAccess")
    def private_ipv6_google_access(self) -> str:
        """
        The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
        """
        return pulumi.get(self, "private_ipv6_google_access")

    @property
    @pulumi.getter(name="reservationAffinity")
    def reservation_affinity(self) -> 'outputs.ReservationAffinityResponse':
        """
        Specifies the reservations that this instance can consume from.
        """
        return pulumi.get(self, "reservation_affinity")

    @property
    @pulumi.getter(name="resourcePolicies")
    def resource_policies(self) -> Sequence[str]:
        """
        Resource policies applied to this instance.
        """
        return pulumi.get(self, "resource_policies")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> 'outputs.ResourceStatusResponse':
        """
        Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter
    def scheduling(self) -> 'outputs.SchedulingResponse':
        """
        Sets the scheduling options for this instance.
        """
        return pulumi.get(self, "scheduling")

    @property
    @pulumi.getter(name="secureTags")
    def secure_tags(self) -> Sequence[str]:
        """
        [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
        """
        return pulumi.get(self, "secure_tags")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> Sequence['outputs.ServiceAccountResponse']:
        """
        A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
        """
        return pulumi.get(self, "service_accounts")

    @property
    @pulumi.getter(name="serviceIntegrationSpecs")
    def service_integration_specs(self) -> Mapping[str, str]:
        """
        Mapping of user-defined keys to specifications for service integrations. Currently only a single key-value pair is supported.
        """
        return pulumi.get(self, "service_integration_specs")

    @property
    @pulumi.getter(name="shieldedInstanceConfig")
    def shielded_instance_config(self) -> 'outputs.ShieldedInstanceConfigResponse':
        return pulumi.get(self, "shielded_instance_config")

    @property
    @pulumi.getter(name="shieldedInstanceIntegrityPolicy")
    def shielded_instance_integrity_policy(self) -> 'outputs.ShieldedInstanceIntegrityPolicyResponse':
        return pulumi.get(self, "shielded_instance_integrity_policy")

    @property
    @pulumi.getter(name="shieldedVmConfig")
    def shielded_vm_config(self) -> 'outputs.ShieldedVmConfigResponse':
        """
        Deprecating, please use shielded_instance_config.
        """
        return pulumi.get(self, "shielded_vm_config")

    @property
    @pulumi.getter(name="shieldedVmIntegrityPolicy")
    def shielded_vm_integrity_policy(self) -> 'outputs.ShieldedVmIntegrityPolicyResponse':
        """
        Deprecating, please use shielded_instance_integrity_policy.
        """
        return pulumi.get(self, "shielded_vm_integrity_policy")

    @property
    @pulumi.getter(name="sourceMachineImage")
    def source_machine_image(self) -> str:
        """
        Source machine image
        """
        return pulumi.get(self, "source_machine_image")

    @property
    @pulumi.getter(name="sourceMachineImageEncryptionKey")
    def source_machine_image_encryption_key(self) -> 'outputs.CustomerEncryptionKeyResponse':
        """
        Source machine image encryption key when creating an instance from a machine image.
        """
        return pulumi.get(self, "source_machine_image_encryption_key")

    @property
    @pulumi.getter(name="startRestricted")
    def start_restricted(self) -> bool:
        """
        Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
        """
        return pulumi.get(self, "start_restricted")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        An optional, human-readable explanation of the status.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def tags(self) -> 'outputs.TagsResponse':
        """
        Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="upcomingMaintenance")
    def upcoming_maintenance(self) -> 'outputs.UpcomingMaintenanceResponse':
        """
        Specifies upcoming maintenance for the instance.
        """
        return pulumi.get(self, "upcoming_maintenance")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "zone")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            advanced_machine_features=self.advanced_machine_features,
            can_ip_forward=self.can_ip_forward,
            confidential_instance_config=self.confidential_instance_config,
            cpu_platform=self.cpu_platform,
            creation_timestamp=self.creation_timestamp,
            deletion_protection=self.deletion_protection,
            description=self.description,
            disks=self.disks,
            display_device=self.display_device,
            erase_windows_vss_signature=self.erase_windows_vss_signature,
            fingerprint=self.fingerprint,
            guest_accelerators=self.guest_accelerators,
            hostname=self.hostname,
            instance_encryption_key=self.instance_encryption_key,
            key_revocation_action_type=self.key_revocation_action_type,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            last_start_timestamp=self.last_start_timestamp,
            last_stop_timestamp=self.last_stop_timestamp,
            last_suspended_timestamp=self.last_suspended_timestamp,
            machine_type=self.machine_type,
            metadata=self.metadata,
            min_cpu_platform=self.min_cpu_platform,
            name=self.name,
            network_interfaces=self.network_interfaces,
            network_performance_config=self.network_performance_config,
            params=self.params,
            post_key_revocation_action_type=self.post_key_revocation_action_type,
            preserved_state_size_gb=self.preserved_state_size_gb,
            private_ipv6_google_access=self.private_ipv6_google_access,
            reservation_affinity=self.reservation_affinity,
            resource_policies=self.resource_policies,
            resource_status=self.resource_status,
            satisfies_pzs=self.satisfies_pzs,
            scheduling=self.scheduling,
            secure_tags=self.secure_tags,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            service_accounts=self.service_accounts,
            service_integration_specs=self.service_integration_specs,
            shielded_instance_config=self.shielded_instance_config,
            shielded_instance_integrity_policy=self.shielded_instance_integrity_policy,
            shielded_vm_config=self.shielded_vm_config,
            shielded_vm_integrity_policy=self.shielded_vm_integrity_policy,
            source_machine_image=self.source_machine_image,
            source_machine_image_encryption_key=self.source_machine_image_encryption_key,
            start_restricted=self.start_restricted,
            status=self.status,
            status_message=self.status_message,
            tags=self.tags,
            upcoming_maintenance=self.upcoming_maintenance,
            zone=self.zone)


def get_instance(instance: Optional[str] = None,
                 project: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Returns the specified Instance resource.
    """
    __args__ = dict()
    __args__['instance'] = instance
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        advanced_machine_features=__ret__.advanced_machine_features,
        can_ip_forward=__ret__.can_ip_forward,
        confidential_instance_config=__ret__.confidential_instance_config,
        cpu_platform=__ret__.cpu_platform,
        creation_timestamp=__ret__.creation_timestamp,
        deletion_protection=__ret__.deletion_protection,
        description=__ret__.description,
        disks=__ret__.disks,
        display_device=__ret__.display_device,
        erase_windows_vss_signature=__ret__.erase_windows_vss_signature,
        fingerprint=__ret__.fingerprint,
        guest_accelerators=__ret__.guest_accelerators,
        hostname=__ret__.hostname,
        instance_encryption_key=__ret__.instance_encryption_key,
        key_revocation_action_type=__ret__.key_revocation_action_type,
        kind=__ret__.kind,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        last_start_timestamp=__ret__.last_start_timestamp,
        last_stop_timestamp=__ret__.last_stop_timestamp,
        last_suspended_timestamp=__ret__.last_suspended_timestamp,
        machine_type=__ret__.machine_type,
        metadata=__ret__.metadata,
        min_cpu_platform=__ret__.min_cpu_platform,
        name=__ret__.name,
        network_interfaces=__ret__.network_interfaces,
        network_performance_config=__ret__.network_performance_config,
        params=__ret__.params,
        post_key_revocation_action_type=__ret__.post_key_revocation_action_type,
        preserved_state_size_gb=__ret__.preserved_state_size_gb,
        private_ipv6_google_access=__ret__.private_ipv6_google_access,
        reservation_affinity=__ret__.reservation_affinity,
        resource_policies=__ret__.resource_policies,
        resource_status=__ret__.resource_status,
        satisfies_pzs=__ret__.satisfies_pzs,
        scheduling=__ret__.scheduling,
        secure_tags=__ret__.secure_tags,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        service_accounts=__ret__.service_accounts,
        service_integration_specs=__ret__.service_integration_specs,
        shielded_instance_config=__ret__.shielded_instance_config,
        shielded_instance_integrity_policy=__ret__.shielded_instance_integrity_policy,
        shielded_vm_config=__ret__.shielded_vm_config,
        shielded_vm_integrity_policy=__ret__.shielded_vm_integrity_policy,
        source_machine_image=__ret__.source_machine_image,
        source_machine_image_encryption_key=__ret__.source_machine_image_encryption_key,
        start_restricted=__ret__.start_restricted,
        status=__ret__.status,
        status_message=__ret__.status_message,
        tags=__ret__.tags,
        upcoming_maintenance=__ret__.upcoming_maintenance,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Returns the specified Instance resource.
    """
    ...
