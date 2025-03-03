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
    def __init__(__self__, alternative_location_id=None, auth_enabled=None, authorized_network=None, available_maintenance_versions=None, connect_mode=None, create_time=None, current_location_id=None, customer_managed_key=None, display_name=None, host=None, labels=None, location=None, maintenance_policy=None, maintenance_schedule=None, maintenance_version=None, memory_size_gb=None, name=None, nodes=None, persistence_config=None, persistence_iam_identity=None, port=None, read_endpoint=None, read_endpoint_port=None, read_replicas_mode=None, redis_configs=None, redis_version=None, replica_count=None, reserved_ip_range=None, secondary_ip_range=None, server_ca_certs=None, state=None, status_message=None, suspension_reasons=None, tier=None, transit_encryption_mode=None):
        if alternative_location_id and not isinstance(alternative_location_id, str):
            raise TypeError("Expected argument 'alternative_location_id' to be a str")
        pulumi.set(__self__, "alternative_location_id", alternative_location_id)
        if auth_enabled and not isinstance(auth_enabled, bool):
            raise TypeError("Expected argument 'auth_enabled' to be a bool")
        pulumi.set(__self__, "auth_enabled", auth_enabled)
        if authorized_network and not isinstance(authorized_network, str):
            raise TypeError("Expected argument 'authorized_network' to be a str")
        pulumi.set(__self__, "authorized_network", authorized_network)
        if available_maintenance_versions and not isinstance(available_maintenance_versions, list):
            raise TypeError("Expected argument 'available_maintenance_versions' to be a list")
        pulumi.set(__self__, "available_maintenance_versions", available_maintenance_versions)
        if connect_mode and not isinstance(connect_mode, str):
            raise TypeError("Expected argument 'connect_mode' to be a str")
        pulumi.set(__self__, "connect_mode", connect_mode)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if current_location_id and not isinstance(current_location_id, str):
            raise TypeError("Expected argument 'current_location_id' to be a str")
        pulumi.set(__self__, "current_location_id", current_location_id)
        if customer_managed_key and not isinstance(customer_managed_key, str):
            raise TypeError("Expected argument 'customer_managed_key' to be a str")
        pulumi.set(__self__, "customer_managed_key", customer_managed_key)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maintenance_policy and not isinstance(maintenance_policy, dict):
            raise TypeError("Expected argument 'maintenance_policy' to be a dict")
        pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if maintenance_schedule and not isinstance(maintenance_schedule, dict):
            raise TypeError("Expected argument 'maintenance_schedule' to be a dict")
        pulumi.set(__self__, "maintenance_schedule", maintenance_schedule)
        if maintenance_version and not isinstance(maintenance_version, str):
            raise TypeError("Expected argument 'maintenance_version' to be a str")
        pulumi.set(__self__, "maintenance_version", maintenance_version)
        if memory_size_gb and not isinstance(memory_size_gb, int):
            raise TypeError("Expected argument 'memory_size_gb' to be a int")
        pulumi.set(__self__, "memory_size_gb", memory_size_gb)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if persistence_config and not isinstance(persistence_config, dict):
            raise TypeError("Expected argument 'persistence_config' to be a dict")
        pulumi.set(__self__, "persistence_config", persistence_config)
        if persistence_iam_identity and not isinstance(persistence_iam_identity, str):
            raise TypeError("Expected argument 'persistence_iam_identity' to be a str")
        pulumi.set(__self__, "persistence_iam_identity", persistence_iam_identity)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if read_endpoint and not isinstance(read_endpoint, str):
            raise TypeError("Expected argument 'read_endpoint' to be a str")
        pulumi.set(__self__, "read_endpoint", read_endpoint)
        if read_endpoint_port and not isinstance(read_endpoint_port, int):
            raise TypeError("Expected argument 'read_endpoint_port' to be a int")
        pulumi.set(__self__, "read_endpoint_port", read_endpoint_port)
        if read_replicas_mode and not isinstance(read_replicas_mode, str):
            raise TypeError("Expected argument 'read_replicas_mode' to be a str")
        pulumi.set(__self__, "read_replicas_mode", read_replicas_mode)
        if redis_configs and not isinstance(redis_configs, dict):
            raise TypeError("Expected argument 'redis_configs' to be a dict")
        pulumi.set(__self__, "redis_configs", redis_configs)
        if redis_version and not isinstance(redis_version, str):
            raise TypeError("Expected argument 'redis_version' to be a str")
        pulumi.set(__self__, "redis_version", redis_version)
        if replica_count and not isinstance(replica_count, int):
            raise TypeError("Expected argument 'replica_count' to be a int")
        pulumi.set(__self__, "replica_count", replica_count)
        if reserved_ip_range and not isinstance(reserved_ip_range, str):
            raise TypeError("Expected argument 'reserved_ip_range' to be a str")
        pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)
        if secondary_ip_range and not isinstance(secondary_ip_range, str):
            raise TypeError("Expected argument 'secondary_ip_range' to be a str")
        pulumi.set(__self__, "secondary_ip_range", secondary_ip_range)
        if server_ca_certs and not isinstance(server_ca_certs, list):
            raise TypeError("Expected argument 'server_ca_certs' to be a list")
        pulumi.set(__self__, "server_ca_certs", server_ca_certs)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if suspension_reasons and not isinstance(suspension_reasons, list):
            raise TypeError("Expected argument 'suspension_reasons' to be a list")
        pulumi.set(__self__, "suspension_reasons", suspension_reasons)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)
        if transit_encryption_mode and not isinstance(transit_encryption_mode, str):
            raise TypeError("Expected argument 'transit_encryption_mode' to be a str")
        pulumi.set(__self__, "transit_encryption_mode", transit_encryption_mode)

    @property
    @pulumi.getter(name="alternativeLocationId")
    def alternative_location_id(self) -> str:
        """
        Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
        """
        return pulumi.get(self, "alternative_location_id")

    @property
    @pulumi.getter(name="authEnabled")
    def auth_enabled(self) -> bool:
        """
        Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
        """
        return pulumi.get(self, "auth_enabled")

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> str:
        """
        Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
        """
        return pulumi.get(self, "authorized_network")

    @property
    @pulumi.getter(name="availableMaintenanceVersions")
    def available_maintenance_versions(self) -> Sequence[str]:
        """
        Optional. The available maintenance versions that an instance could update to.
        """
        return pulumi.get(self, "available_maintenance_versions")

    @property
    @pulumi.getter(name="connectMode")
    def connect_mode(self) -> str:
        """
        Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
        """
        return pulumi.get(self, "connect_mode")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="currentLocationId")
    def current_location_id(self) -> str:
        """
        The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
        """
        return pulumi.get(self, "current_location_id")

    @property
    @pulumi.getter(name="customerManagedKey")
    def customer_managed_key(self) -> str:
        """
        Optional. The KMS key reference that the customer provides when trying to create the instance.
        """
        return pulumi.get(self, "customer_managed_key")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        An arbitrary and optional user-provided name for the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Resource labels to represent user provided metadata
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> 'outputs.MaintenancePolicyResponse':
        """
        Optional. The maintenance policy for the instance. If not provided, maintenance events can be performed at any time.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter(name="maintenanceSchedule")
    def maintenance_schedule(self) -> 'outputs.MaintenanceScheduleResponse':
        """
        Date and time of upcoming maintenance events which have been scheduled.
        """
        return pulumi.get(self, "maintenance_schedule")

    @property
    @pulumi.getter(name="maintenanceVersion")
    def maintenance_version(self) -> str:
        """
        Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
        """
        return pulumi.get(self, "maintenance_version")

    @property
    @pulumi.getter(name="memorySizeGb")
    def memory_size_gb(self) -> int:
        """
        Redis memory size in GiB.
        """
        return pulumi.get(self, "memory_size_gb")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.NodeInfoResponse']:
        """
        Info per node.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="persistenceConfig")
    def persistence_config(self) -> 'outputs.PersistenceConfigResponse':
        """
        Optional. Persistence configuration parameters
        """
        return pulumi.get(self, "persistence_config")

    @property
    @pulumi.getter(name="persistenceIamIdentity")
    def persistence_iam_identity(self) -> str:
        """
        Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
        """
        return pulumi.get(self, "persistence_iam_identity")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The port number of the exposed Redis endpoint.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="readEndpoint")
    def read_endpoint(self) -> str:
        """
        Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
        """
        return pulumi.get(self, "read_endpoint")

    @property
    @pulumi.getter(name="readEndpointPort")
    def read_endpoint_port(self) -> int:
        """
        The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
        """
        return pulumi.get(self, "read_endpoint_port")

    @property
    @pulumi.getter(name="readReplicasMode")
    def read_replicas_mode(self) -> str:
        """
        Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
        """
        return pulumi.get(self, "read_replicas_mode")

    @property
    @pulumi.getter(name="redisConfigs")
    def redis_configs(self) -> Mapping[str, str]:
        """
        Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
        """
        return pulumi.get(self, "redis_configs")

    @property
    @pulumi.getter(name="redisVersion")
    def redis_version(self) -> str:
        """
        Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility
        """
        return pulumi.get(self, "redis_version")

    @property
    @pulumi.getter(name="replicaCount")
    def replica_count(self) -> int:
        """
        Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
        """
        return pulumi.get(self, "replica_count")

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> str:
        """
        Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
        """
        return pulumi.get(self, "reserved_ip_range")

    @property
    @pulumi.getter(name="secondaryIpRange")
    def secondary_ip_range(self) -> str:
        """
        Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
        """
        return pulumi.get(self, "secondary_ip_range")

    @property
    @pulumi.getter(name="serverCaCerts")
    def server_ca_certs(self) -> Sequence['outputs.TlsCertificateResponse']:
        """
        List of server CA certificates for the instance.
        """
        return pulumi.get(self, "server_ca_certs")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of this instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> str:
        """
        Additional information about the current status of this instance, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="suspensionReasons")
    def suspension_reasons(self) -> Sequence[str]:
        """
        Optional. reasons that causes instance in "SUSPENDED" state.
        """
        return pulumi.get(self, "suspension_reasons")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The service tier of the instance.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter(name="transitEncryptionMode")
    def transit_encryption_mode(self) -> str:
        """
        Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
        """
        return pulumi.get(self, "transit_encryption_mode")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            alternative_location_id=self.alternative_location_id,
            auth_enabled=self.auth_enabled,
            authorized_network=self.authorized_network,
            available_maintenance_versions=self.available_maintenance_versions,
            connect_mode=self.connect_mode,
            create_time=self.create_time,
            current_location_id=self.current_location_id,
            customer_managed_key=self.customer_managed_key,
            display_name=self.display_name,
            host=self.host,
            labels=self.labels,
            location=self.location,
            maintenance_policy=self.maintenance_policy,
            maintenance_schedule=self.maintenance_schedule,
            maintenance_version=self.maintenance_version,
            memory_size_gb=self.memory_size_gb,
            name=self.name,
            nodes=self.nodes,
            persistence_config=self.persistence_config,
            persistence_iam_identity=self.persistence_iam_identity,
            port=self.port,
            read_endpoint=self.read_endpoint,
            read_endpoint_port=self.read_endpoint_port,
            read_replicas_mode=self.read_replicas_mode,
            redis_configs=self.redis_configs,
            redis_version=self.redis_version,
            replica_count=self.replica_count,
            reserved_ip_range=self.reserved_ip_range,
            secondary_ip_range=self.secondary_ip_range,
            server_ca_certs=self.server_ca_certs,
            state=self.state,
            status_message=self.status_message,
            suspension_reasons=self.suspension_reasons,
            tier=self.tier,
            transit_encryption_mode=self.transit_encryption_mode)


def get_instance(instance_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets the details of a specific Redis instance.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:redis/v1beta1:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        alternative_location_id=__ret__.alternative_location_id,
        auth_enabled=__ret__.auth_enabled,
        authorized_network=__ret__.authorized_network,
        available_maintenance_versions=__ret__.available_maintenance_versions,
        connect_mode=__ret__.connect_mode,
        create_time=__ret__.create_time,
        current_location_id=__ret__.current_location_id,
        customer_managed_key=__ret__.customer_managed_key,
        display_name=__ret__.display_name,
        host=__ret__.host,
        labels=__ret__.labels,
        location=__ret__.location,
        maintenance_policy=__ret__.maintenance_policy,
        maintenance_schedule=__ret__.maintenance_schedule,
        maintenance_version=__ret__.maintenance_version,
        memory_size_gb=__ret__.memory_size_gb,
        name=__ret__.name,
        nodes=__ret__.nodes,
        persistence_config=__ret__.persistence_config,
        persistence_iam_identity=__ret__.persistence_iam_identity,
        port=__ret__.port,
        read_endpoint=__ret__.read_endpoint,
        read_endpoint_port=__ret__.read_endpoint_port,
        read_replicas_mode=__ret__.read_replicas_mode,
        redis_configs=__ret__.redis_configs,
        redis_version=__ret__.redis_version,
        replica_count=__ret__.replica_count,
        reserved_ip_range=__ret__.reserved_ip_range,
        secondary_ip_range=__ret__.secondary_ip_range,
        server_ca_certs=__ret__.server_ca_certs,
        state=__ret__.state,
        status_message=__ret__.status_message,
        suspension_reasons=__ret__.suspension_reasons,
        tier=__ret__.tier,
        transit_encryption_mode=__ret__.transit_encryption_mode)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets the details of a specific Redis instance.
    """
    ...
