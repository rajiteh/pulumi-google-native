# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BinaryAuthorizationEvaluationMode',
    'CloudRunConfigLoadBalancerType',
    'ClusterAutoscalingAutoscalingProfile',
    'ClusterUpdateDesiredDatapathProvider',
    'ClusterUpdateDesiredPrivateIpv6GoogleAccess',
    'ClusterUpdateDesiredStackType',
    'DNSConfigClusterDns',
    'DNSConfigClusterDnsScope',
    'DatabaseEncryptionState',
    'FilterEventTypeItem',
    'GPUSharingConfigGpuSharingStrategy',
    'GatewayAPIConfigChannel',
    'IPAllocationPolicyIpv6AccessType',
    'IPAllocationPolicyStackType',
    'LinuxNodeConfigCgroupMode',
    'LoggingComponentConfigEnableComponentsItem',
    'LoggingVariantConfigVariant',
    'MaintenanceExclusionOptionsScope',
    'MonitoringComponentConfigEnableComponentsItem',
    'NetworkConfigDatapathProvider',
    'NetworkConfigPrivateIpv6GoogleAccess',
    'NetworkPerformanceConfigTotalEgressBandwidthTier',
    'NetworkPolicyProvider',
    'NodePoolAutoscalingLocationPolicy',
    'NodeTaintEffect',
    'PlacementPolicyType',
    'ReleaseChannelChannel',
    'ReservationAffinityConsumeReservationType',
    'SandboxConfigType',
    'StatusConditionCanonicalCode',
    'StatusConditionCode',
    'UpgradeSettingsStrategy',
    'WindowsNodeConfigOsVersion',
    'WorkloadMetadataConfigMode',
]


class BinaryAuthorizationEvaluationMode(str, Enum):
    """
    Mode of operation for binauthz policy evaluation. If unspecified, defaults to DISABLED.
    """
    EVALUATION_MODE_UNSPECIFIED = "EVALUATION_MODE_UNSPECIFIED"
    """
    Default value
    """
    DISABLED = "DISABLED"
    """
    Disable BinaryAuthorization
    """
    PROJECT_SINGLETON_POLICY_ENFORCE = "PROJECT_SINGLETON_POLICY_ENFORCE"
    """
    Enforce Kubernetes admission requests with BinaryAuthorization using the project's singleton policy. This is equivalent to setting the enabled boolean to true.
    """


class CloudRunConfigLoadBalancerType(str, Enum):
    """
    Which load balancer type is installed for Cloud Run.
    """
    LOAD_BALANCER_TYPE_UNSPECIFIED = "LOAD_BALANCER_TYPE_UNSPECIFIED"
    """
    Load balancer type for Cloud Run is unspecified.
    """
    LOAD_BALANCER_TYPE_EXTERNAL = "LOAD_BALANCER_TYPE_EXTERNAL"
    """
    Install external load balancer for Cloud Run.
    """
    LOAD_BALANCER_TYPE_INTERNAL = "LOAD_BALANCER_TYPE_INTERNAL"
    """
    Install internal load balancer for Cloud Run.
    """


class ClusterAutoscalingAutoscalingProfile(str, Enum):
    """
    Defines autoscaling behaviour.
    """
    PROFILE_UNSPECIFIED = "PROFILE_UNSPECIFIED"
    """
    No change to autoscaling configuration.
    """
    OPTIMIZE_UTILIZATION = "OPTIMIZE_UTILIZATION"
    """
    Prioritize optimizing utilization of resources.
    """
    BALANCED = "BALANCED"
    """
    Use default (balanced) autoscaling configuration.
    """


class ClusterUpdateDesiredDatapathProvider(str, Enum):
    """
    The desired datapath provider for the cluster.
    """
    DATAPATH_PROVIDER_UNSPECIFIED = "DATAPATH_PROVIDER_UNSPECIFIED"
    """
    Default value.
    """
    LEGACY_DATAPATH = "LEGACY_DATAPATH"
    """
    Use the IPTables implementation based on kube-proxy.
    """
    ADVANCED_DATAPATH = "ADVANCED_DATAPATH"
    """
    Use the eBPF based GKE Dataplane V2 with additional features. See the [GKE Dataplane V2 documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/dataplane-v2) for more.
    """


class ClusterUpdateDesiredPrivateIpv6GoogleAccess(str, Enum):
    """
    The desired state of IPv6 connectivity to Google Services.
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED = "PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED"
    """
    Default value. Same as DISABLED
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED = "PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED"
    """
    No private access to or from Google Services
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE = "PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE"
    """
    Enables private IPv6 access to Google Services from GKE
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL = "PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL"
    """
    Enables private IPv6 access to and from Google Services
    """


class ClusterUpdateDesiredStackType(str, Enum):
    """
    The desired stack type of the cluster. If a stack type is provided and does not match the current stack type of the cluster, update will attempt to change the stack type to the new type.
    """
    STACK_TYPE_UNSPECIFIED = "STACK_TYPE_UNSPECIFIED"
    """
    Default value, will be defaulted as IPV4 only
    """
    IPV4 = "IPV4"
    """
    Cluster is IPV4 only
    """
    IPV4_IPV6 = "IPV4_IPV6"
    """
    Cluster can use both IPv4 and IPv6
    """


class DNSConfigClusterDns(str, Enum):
    """
    cluster_dns indicates which in-cluster DNS provider should be used.
    """
    PROVIDER_UNSPECIFIED = "PROVIDER_UNSPECIFIED"
    """
    Default value
    """
    PLATFORM_DEFAULT = "PLATFORM_DEFAULT"
    """
    Use GKE default DNS provider(kube-dns) for DNS resolution.
    """
    CLOUD_DNS = "CLOUD_DNS"
    """
    Use CloudDNS for DNS resolution.
    """


class DNSConfigClusterDnsScope(str, Enum):
    """
    cluster_dns_scope indicates the scope of access to cluster DNS records.
    """
    DNS_SCOPE_UNSPECIFIED = "DNS_SCOPE_UNSPECIFIED"
    """
    Default value, will be inferred as cluster scope.
    """
    CLUSTER_SCOPE = "CLUSTER_SCOPE"
    """
    DNS records are accessible from within the cluster.
    """
    VPC_SCOPE = "VPC_SCOPE"
    """
    DNS records are accessible from within the VPC.
    """


class DatabaseEncryptionState(str, Enum):
    """
    The desired state of etcd encryption.
    """
    UNKNOWN = "UNKNOWN"
    """
    Should never be set
    """
    ENCRYPTED = "ENCRYPTED"
    """
    Secrets in etcd are encrypted.
    """
    DECRYPTED = "DECRYPTED"
    """
    Secrets in etcd are stored in plain text (at etcd level) - this is unrelated to Compute Engine level full disk encryption.
    """


class FilterEventTypeItem(str, Enum):
    EVENT_TYPE_UNSPECIFIED = "EVENT_TYPE_UNSPECIFIED"
    """
    Not set, will be ignored.
    """
    UPGRADE_AVAILABLE_EVENT = "UPGRADE_AVAILABLE_EVENT"
    """
    Corresponds with UpgradeAvailableEvent.
    """
    UPGRADE_EVENT = "UPGRADE_EVENT"
    """
    Corresponds with UpgradeEvent.
    """
    SECURITY_BULLETIN_EVENT = "SECURITY_BULLETIN_EVENT"
    """
    Corresponds with SecurityBulletinEvent.
    """


class GPUSharingConfigGpuSharingStrategy(str, Enum):
    """
    The type of GPU sharing strategy to enable on the GPU node.
    """
    GPU_SHARING_STRATEGY_UNSPECIFIED = "GPU_SHARING_STRATEGY_UNSPECIFIED"
    """
    Default value.
    """
    TIME_SHARING = "TIME_SHARING"
    """
    GPUs are time-shared between containers.
    """


class GatewayAPIConfigChannel(str, Enum):
    """
    The Gateway API release channel to use for Gateway API.
    """
    CHANNEL_UNSPECIFIED = "CHANNEL_UNSPECIFIED"
    """
    Default value.
    """
    CHANNEL_DISABLED = "CHANNEL_DISABLED"
    """
    Gateway API support is disabled
    """
    CHANNEL_EXPERIMENTAL = "CHANNEL_EXPERIMENTAL"
    """
    Gateway API support is enabled, experimental CRDs are installed
    """
    CHANNEL_STANDARD = "CHANNEL_STANDARD"
    """
    Gateway API support is enabled, standard CRDs are installed
    """


class IPAllocationPolicyIpv6AccessType(str, Enum):
    """
    The ipv6 access type (internal or external) when create_subnetwork is true
    """
    IPV6_ACCESS_TYPE_UNSPECIFIED = "IPV6_ACCESS_TYPE_UNSPECIFIED"
    """
    Default value, will be defaulted as type external.
    """
    INTERNAL = "INTERNAL"
    """
    Access type internal (all v6 addresses are internal IPs)
    """
    EXTERNAL = "EXTERNAL"
    """
    Access type external (all v6 addresses are external IPs)
    """


class IPAllocationPolicyStackType(str, Enum):
    """
    The IP stack type of the cluster
    """
    STACK_TYPE_UNSPECIFIED = "STACK_TYPE_UNSPECIFIED"
    """
    Default value, will be defaulted as IPV4 only
    """
    IPV4 = "IPV4"
    """
    Cluster is IPV4 only
    """
    IPV4_IPV6 = "IPV4_IPV6"
    """
    Cluster can use both IPv4 and IPv6
    """


class LinuxNodeConfigCgroupMode(str, Enum):
    """
    cgroup_mode specifies the cgroup mode to be used on the node.
    """
    CGROUP_MODE_UNSPECIFIED = "CGROUP_MODE_UNSPECIFIED"
    """
    CGROUP_MODE_UNSPECIFIED is when unspecified cgroup configuration is used. The default for the GKE node OS image will be used.
    """
    CGROUP_MODE_V1 = "CGROUP_MODE_V1"
    """
    CGROUP_MODE_V1 specifies to use cgroupv1 for the cgroup configuration on the node image.
    """
    CGROUP_MODE_V2 = "CGROUP_MODE_V2"
    """
    CGROUP_MODE_V2 specifies to use cgroupv2 for the cgroup configuration on the node image.
    """


class LoggingComponentConfigEnableComponentsItem(str, Enum):
    COMPONENT_UNSPECIFIED = "COMPONENT_UNSPECIFIED"
    """
    Default value. This shouldn't be used.
    """
    SYSTEM_COMPONENTS = "SYSTEM_COMPONENTS"
    """
    system components
    """
    WORKLOADS = "WORKLOADS"
    """
    workloads
    """
    APISERVER = "APISERVER"
    """
    kube-apiserver
    """
    SCHEDULER = "SCHEDULER"
    """
    kube-scheduler
    """
    CONTROLLER_MANAGER = "CONTROLLER_MANAGER"
    """
    kube-controller-manager
    """


class LoggingVariantConfigVariant(str, Enum):
    """
    Logging variant deployed on nodes.
    """
    VARIANT_UNSPECIFIED = "VARIANT_UNSPECIFIED"
    """
    Default value. This shouldn't be used.
    """
    DEFAULT = "DEFAULT"
    """
    default logging variant.
    """
    MAX_THROUGHPUT = "MAX_THROUGHPUT"
    """
    maximum logging throughput variant.
    """


class MaintenanceExclusionOptionsScope(str, Enum):
    """
    Scope specifies the upgrade scope which upgrades are blocked by the exclusion.
    """
    NO_UPGRADES = "NO_UPGRADES"
    """
    NO_UPGRADES excludes all upgrades, including patch upgrades and minor upgrades across control planes and nodes. This is the default exclusion behavior.
    """
    NO_MINOR_UPGRADES = "NO_MINOR_UPGRADES"
    """
    NO_MINOR_UPGRADES excludes all minor upgrades for the cluster, only patches are allowed.
    """
    NO_MINOR_OR_NODE_UPGRADES = "NO_MINOR_OR_NODE_UPGRADES"
    """
    NO_MINOR_OR_NODE_UPGRADES excludes all minor upgrades for the cluster, and also exclude all node pool upgrades. Only control plane patches are allowed.
    """


class MonitoringComponentConfigEnableComponentsItem(str, Enum):
    COMPONENT_UNSPECIFIED = "COMPONENT_UNSPECIFIED"
    """
    Default value. This shouldn't be used.
    """
    SYSTEM_COMPONENTS = "SYSTEM_COMPONENTS"
    """
    system components
    """
    APISERVER = "APISERVER"
    """
    kube-apiserver
    """
    SCHEDULER = "SCHEDULER"
    """
    kube-scheduler
    """
    CONTROLLER_MANAGER = "CONTROLLER_MANAGER"
    """
    kube-controller-manager
    """


class NetworkConfigDatapathProvider(str, Enum):
    """
    The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.
    """
    DATAPATH_PROVIDER_UNSPECIFIED = "DATAPATH_PROVIDER_UNSPECIFIED"
    """
    Default value.
    """
    LEGACY_DATAPATH = "LEGACY_DATAPATH"
    """
    Use the IPTables implementation based on kube-proxy.
    """
    ADVANCED_DATAPATH = "ADVANCED_DATAPATH"
    """
    Use the eBPF based GKE Dataplane V2 with additional features. See the [GKE Dataplane V2 documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/dataplane-v2) for more.
    """


class NetworkConfigPrivateIpv6GoogleAccess(str, Enum):
    """
    The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4)
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED = "PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED"
    """
    Default value. Same as DISABLED
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED = "PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED"
    """
    No private access to or from Google Services
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE = "PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE"
    """
    Enables private IPv6 access to Google Services from GKE
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL = "PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL"
    """
    Enables private IPv6 access to and from Google Services
    """


class NetworkPerformanceConfigTotalEgressBandwidthTier(str, Enum):
    """
    Specifies the total network bandwidth tier for the NodePool.
    """
    TIER_UNSPECIFIED = "TIER_UNSPECIFIED"
    """
    Default value
    """
    TIER1 = "TIER_1"
    """
    Higher bandwidth, actual values based on VM size.
    """


class NetworkPolicyProvider(str, Enum):
    """
    The selected network policy provider.
    """
    PROVIDER_UNSPECIFIED = "PROVIDER_UNSPECIFIED"
    """
    Not set
    """
    CALICO = "CALICO"
    """
    Tigera (Calico Felix).
    """


class NodePoolAutoscalingLocationPolicy(str, Enum):
    """
    Location policy used when scaling up a nodepool.
    """
    LOCATION_POLICY_UNSPECIFIED = "LOCATION_POLICY_UNSPECIFIED"
    """
    Not set.
    """
    BALANCED = "BALANCED"
    """
    BALANCED is a best effort policy that aims to balance the sizes of different zones.
    """
    ANY = "ANY"
    """
    ANY policy picks zones that have the highest capacity available.
    """


class NodeTaintEffect(str, Enum):
    """
    Effect for taint.
    """
    EFFECT_UNSPECIFIED = "EFFECT_UNSPECIFIED"
    """
    Not set
    """
    NO_SCHEDULE = "NO_SCHEDULE"
    """
    NoSchedule
    """
    PREFER_NO_SCHEDULE = "PREFER_NO_SCHEDULE"
    """
    PreferNoSchedule
    """
    NO_EXECUTE = "NO_EXECUTE"
    """
    NoExecute
    """


class PlacementPolicyType(str, Enum):
    """
    The type of placement.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    TYPE_UNSPECIFIED specifies no requirements on nodes placement.
    """
    COMPACT = "COMPACT"
    """
    COMPACT specifies node placement in the same availability domain to ensure low communication latency.
    """


class ReleaseChannelChannel(str, Enum):
    """
    channel specifies which release channel the cluster is subscribed to.
    """
    UNSPECIFIED = "UNSPECIFIED"
    """
    No channel specified.
    """
    RAPID = "RAPID"
    """
    RAPID channel is offered on an early access basis for customers who want to test new releases. WARNING: Versions available in the RAPID Channel may be subject to unresolved issues with no known workaround and are not subject to any SLAs.
    """
    REGULAR = "REGULAR"
    """
    Clusters subscribed to REGULAR receive versions that are considered GA quality. REGULAR is intended for production users who want to take advantage of new features.
    """
    STABLE = "STABLE"
    """
    Clusters subscribed to STABLE receive versions that are known to be stable and reliable in production.
    """


class ReservationAffinityConsumeReservationType(str, Enum):
    """
    Corresponds to the type of reservation consumption.
    """
    UNSPECIFIED = "UNSPECIFIED"
    """
    Default value. This should not be used.
    """
    NO_RESERVATION = "NO_RESERVATION"
    """
    Do not consume from any reserved capacity.
    """
    ANY_RESERVATION = "ANY_RESERVATION"
    """
    Consume any reservation available.
    """
    SPECIFIC_RESERVATION = "SPECIFIC_RESERVATION"
    """
    Must consume from a specific reservation. Must specify key value fields for specifying the reservations.
    """


class SandboxConfigType(str, Enum):
    """
    Type of the sandbox to use for the node.
    """
    UNSPECIFIED = "UNSPECIFIED"
    """
    Default value. This should not be used.
    """
    GVISOR = "GVISOR"
    """
    Run sandbox using gvisor.
    """


class StatusConditionCanonicalCode(str, Enum):
    """
    Canonical code of the condition.
    """
    OK = "OK"
    """
    Not an error; returned on success. HTTP Mapping: 200 OK
    """
    CANCELLED = "CANCELLED"
    """
    The operation was cancelled, typically by the caller. HTTP Mapping: 499 Client Closed Request
    """
    UNKNOWN = "UNKNOWN"
    """
    Unknown error. For example, this error may be returned when a `Status` value received from another address space belongs to an error space that is not known in this address space. Also errors raised by APIs that do not return enough error information may be converted to this error. HTTP Mapping: 500 Internal Server Error
    """
    INVALID_ARGUMENT = "INVALID_ARGUMENT"
    """
    The client specified an invalid argument. Note that this differs from `FAILED_PRECONDITION`. `INVALID_ARGUMENT` indicates arguments that are problematic regardless of the state of the system (e.g., a malformed file name). HTTP Mapping: 400 Bad Request
    """
    DEADLINE_EXCEEDED = "DEADLINE_EXCEEDED"
    """
    The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully. For example, a successful response from a server could have been delayed long enough for the deadline to expire. HTTP Mapping: 504 Gateway Timeout
    """
    NOT_FOUND = "NOT_FOUND"
    """
    Some requested entity (e.g., file or directory) was not found. Note to server developers: if a request is denied for an entire class of users, such as gradual feature rollout or undocumented allowlist, `NOT_FOUND` may be used. If a request is denied for some users within a class of users, such as user-based access control, `PERMISSION_DENIED` must be used. HTTP Mapping: 404 Not Found
    """
    ALREADY_EXISTS = "ALREADY_EXISTS"
    """
    The entity that a client attempted to create (e.g., file or directory) already exists. HTTP Mapping: 409 Conflict
    """
    PERMISSION_DENIED = "PERMISSION_DENIED"
    """
    The caller does not have permission to execute the specified operation. `PERMISSION_DENIED` must not be used for rejections caused by exhausting some resource (use `RESOURCE_EXHAUSTED` instead for those errors). `PERMISSION_DENIED` must not be used if the caller can not be identified (use `UNAUTHENTICATED` instead for those errors). This error code does not imply the request is valid or the requested entity exists or satisfies other pre-conditions. HTTP Mapping: 403 Forbidden
    """
    UNAUTHENTICATED = "UNAUTHENTICATED"
    """
    The request does not have valid authentication credentials for the operation. HTTP Mapping: 401 Unauthorized
    """
    RESOURCE_EXHAUSTED = "RESOURCE_EXHAUSTED"
    """
    Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space. HTTP Mapping: 429 Too Many Requests
    """
    FAILED_PRECONDITION = "FAILED_PRECONDITION"
    """
    The operation was rejected because the system is not in a state required for the operation's execution. For example, the directory to be deleted is non-empty, an rmdir operation is applied to a non-directory, etc. Service implementors can use the following guidelines to decide between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`: (a) Use `UNAVAILABLE` if the client can retry just the failing call. (b) Use `ABORTED` if the client should retry at a higher level. For example, when a client-specified test-and-set fails, indicating the client should restart a read-modify-write sequence. (c) Use `FAILED_PRECONDITION` if the client should not retry until the system state has been explicitly fixed. For example, if an "rmdir" fails because the directory is non-empty, `FAILED_PRECONDITION` should be returned since the client should not retry unless the files are deleted from the directory. HTTP Mapping: 400 Bad Request
    """
    ABORTED = "ABORTED"
    """
    The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort. See the guidelines above for deciding between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`. HTTP Mapping: 409 Conflict
    """
    OUT_OF_RANGE = "OUT_OF_RANGE"
    """
    The operation was attempted past the valid range. E.g., seeking or reading past end-of-file. Unlike `INVALID_ARGUMENT`, this error indicates a problem that may be fixed if the system state changes. For example, a 32-bit file system will generate `INVALID_ARGUMENT` if asked to read at an offset that is not in the range [0,2^32-1], but it will generate `OUT_OF_RANGE` if asked to read from an offset past the current file size. There is a fair bit of overlap between `FAILED_PRECONDITION` and `OUT_OF_RANGE`. We recommend using `OUT_OF_RANGE` (the more specific error) when it applies so that callers who are iterating through a space can easily look for an `OUT_OF_RANGE` error to detect when they are done. HTTP Mapping: 400 Bad Request
    """
    UNIMPLEMENTED = "UNIMPLEMENTED"
    """
    The operation is not implemented or is not supported/enabled in this service. HTTP Mapping: 501 Not Implemented
    """
    INTERNAL = "INTERNAL"
    """
    Internal errors. This means that some invariants expected by the underlying system have been broken. This error code is reserved for serious errors. HTTP Mapping: 500 Internal Server Error
    """
    UNAVAILABLE = "UNAVAILABLE"
    """
    The service is currently unavailable. This is most likely a transient condition, which can be corrected by retrying with a backoff. Note that it is not always safe to retry non-idempotent operations. See the guidelines above for deciding between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`. HTTP Mapping: 503 Service Unavailable
    """
    DATA_LOSS = "DATA_LOSS"
    """
    Unrecoverable data loss or corruption. HTTP Mapping: 500 Internal Server Error
    """


class StatusConditionCode(str, Enum):
    """
    Machine-friendly representation of the condition Deprecated. Use canonical_code instead.
    """
    UNKNOWN = "UNKNOWN"
    """
    UNKNOWN indicates a generic condition.
    """
    GCE_STOCKOUT = "GCE_STOCKOUT"
    """
    GCE_STOCKOUT indicates that Google Compute Engine resources are temporarily unavailable.
    """
    GKE_SERVICE_ACCOUNT_DELETED = "GKE_SERVICE_ACCOUNT_DELETED"
    """
    GKE_SERVICE_ACCOUNT_DELETED indicates that the user deleted their robot service account.
    """
    GCE_QUOTA_EXCEEDED = "GCE_QUOTA_EXCEEDED"
    """
    Google Compute Engine quota was exceeded.
    """
    SET_BY_OPERATOR = "SET_BY_OPERATOR"
    """
    Cluster state was manually changed by an SRE due to a system logic error.
    """
    CLOUD_KMS_KEY_ERROR = "CLOUD_KMS_KEY_ERROR"
    """
    Unable to perform an encrypt operation against the CloudKMS key used for etcd level encryption.
    """
    CA_EXPIRING = "CA_EXPIRING"
    """
    Cluster CA is expiring soon.
    """


class UpgradeSettingsStrategy(str, Enum):
    """
    Update strategy of the node pool.
    """
    NODE_POOL_UPDATE_STRATEGY_UNSPECIFIED = "NODE_POOL_UPDATE_STRATEGY_UNSPECIFIED"
    """
    Default value if unset. GKE internally defaults the update strategy to SURGE for unspecified strategies.
    """
    BLUE_GREEN = "BLUE_GREEN"
    """
    blue-green upgrade.
    """
    SURGE = "SURGE"
    """
    SURGE is the traditional way of upgrade a node pool. max_surge and max_unavailable determines the level of upgrade parallelism.
    """


class WindowsNodeConfigOsVersion(str, Enum):
    """
    OSVersion specifies the Windows node config to be used on the node
    """
    OS_VERSION_UNSPECIFIED = "OS_VERSION_UNSPECIFIED"
    """
    When OSVersion is not specified
    """
    OS_VERSION_LTSC2019 = "OS_VERSION_LTSC2019"
    """
    LTSC2019 specifies to use LTSC2019 as the Windows Servercore Base Image
    """
    OS_VERSION_LTSC2022 = "OS_VERSION_LTSC2022"
    """
    LTSC2022 specifies to use LTSC2022 as the Windows Servercore Base Image
    """


class WorkloadMetadataConfigMode(str, Enum):
    """
    Mode is the configuration for how to expose metadata to workloads running on the node pool.
    """
    MODE_UNSPECIFIED = "MODE_UNSPECIFIED"
    """
    Not set.
    """
    GCE_METADATA = "GCE_METADATA"
    """
    Expose all Compute Engine metadata to pods.
    """
    GKE_METADATA = "GKE_METADATA"
    """
    Run the GKE Metadata Server on this node. The GKE Metadata Server exposes a metadata API to workloads that is compatible with the V1 Compute Metadata APIs exposed by the Compute Engine and App Engine Metadata Servers. This feature can only be enabled if Workload Identity is enabled at the cluster level.
    """
