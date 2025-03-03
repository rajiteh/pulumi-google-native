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

__all__ = ['BareMetalAdminClusterArgs', 'BareMetalAdminCluster']

@pulumi.input_type
class BareMetalAdminClusterArgs:
    def __init__(__self__, *,
                 bare_metal_admin_cluster_id: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 bare_metal_version: Optional[pulumi.Input[str]] = None,
                 cluster_operations: Optional[pulumi.Input['BareMetalAdminClusterOperationsConfigArgs']] = None,
                 control_plane: Optional[pulumi.Input['BareMetalAdminControlPlaneConfigArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 load_balancer: Optional[pulumi.Input['BareMetalAdminLoadBalancerConfigArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_config: Optional[pulumi.Input['BareMetalAdminMaintenanceConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['BareMetalAdminNetworkConfigArgs']] = None,
                 node_access_config: Optional[pulumi.Input['BareMetalAdminNodeAccessConfigArgs']] = None,
                 node_config: Optional[pulumi.Input['BareMetalAdminWorkloadNodeConfigArgs']] = None,
                 os_environment_config: Optional[pulumi.Input['BareMetalAdminOsEnvironmentConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input['BareMetalAdminProxyConfigArgs']] = None,
                 security_config: Optional[pulumi.Input['BareMetalAdminSecurityConfigArgs']] = None,
                 storage: Optional[pulumi.Input['BareMetalAdminStorageConfigArgs']] = None):
        """
        The set of arguments for constructing a BareMetalAdminCluster resource.
        :param pulumi.Input[str] bare_metal_admin_cluster_id: Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        :param pulumi.Input[str] bare_metal_version: The Anthos clusters on bare metal version for the bare metal admin cluster.
        :param pulumi.Input['BareMetalAdminClusterOperationsConfigArgs'] cluster_operations: Cluster operations configuration.
        :param pulumi.Input['BareMetalAdminControlPlaneConfigArgs'] control_plane: Control plane configuration.
        :param pulumi.Input[str] description: A human readable description of this bare metal admin cluster.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        :param pulumi.Input['BareMetalAdminLoadBalancerConfigArgs'] load_balancer: Load balancer configuration.
        :param pulumi.Input['BareMetalAdminMaintenanceConfigArgs'] maintenance_config: Maintenance configuration.
        :param pulumi.Input[str] name: Immutable. The bare metal admin cluster resource name.
        :param pulumi.Input['BareMetalAdminNetworkConfigArgs'] network_config: Network configuration.
        :param pulumi.Input['BareMetalAdminNodeAccessConfigArgs'] node_access_config: Node access related configurations.
        :param pulumi.Input['BareMetalAdminWorkloadNodeConfigArgs'] node_config: Workload node configuration.
        :param pulumi.Input['BareMetalAdminOsEnvironmentConfigArgs'] os_environment_config: OS environment related configurations.
        :param pulumi.Input['BareMetalAdminProxyConfigArgs'] proxy: Proxy configuration.
        :param pulumi.Input['BareMetalAdminSecurityConfigArgs'] security_config: Security related configuration.
        :param pulumi.Input['BareMetalAdminStorageConfigArgs'] storage: Storage configuration.
        """
        pulumi.set(__self__, "bare_metal_admin_cluster_id", bare_metal_admin_cluster_id)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if bare_metal_version is not None:
            pulumi.set(__self__, "bare_metal_version", bare_metal_version)
        if cluster_operations is not None:
            pulumi.set(__self__, "cluster_operations", cluster_operations)
        if control_plane is not None:
            pulumi.set(__self__, "control_plane", control_plane)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if load_balancer is not None:
            pulumi.set(__self__, "load_balancer", load_balancer)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if maintenance_config is not None:
            pulumi.set(__self__, "maintenance_config", maintenance_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if node_access_config is not None:
            pulumi.set(__self__, "node_access_config", node_access_config)
        if node_config is not None:
            pulumi.set(__self__, "node_config", node_config)
        if os_environment_config is not None:
            pulumi.set(__self__, "os_environment_config", os_environment_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if proxy is not None:
            pulumi.set(__self__, "proxy", proxy)
        if security_config is not None:
            pulumi.set(__self__, "security_config", security_config)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)

    @property
    @pulumi.getter(name="bareMetalAdminClusterId")
    def bare_metal_admin_cluster_id(self) -> pulumi.Input[str]:
        """
        Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
        """
        return pulumi.get(self, "bare_metal_admin_cluster_id")

    @bare_metal_admin_cluster_id.setter
    def bare_metal_admin_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bare_metal_admin_cluster_id", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="bareMetalVersion")
    def bare_metal_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Anthos clusters on bare metal version for the bare metal admin cluster.
        """
        return pulumi.get(self, "bare_metal_version")

    @bare_metal_version.setter
    def bare_metal_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bare_metal_version", value)

    @property
    @pulumi.getter(name="clusterOperations")
    def cluster_operations(self) -> Optional[pulumi.Input['BareMetalAdminClusterOperationsConfigArgs']]:
        """
        Cluster operations configuration.
        """
        return pulumi.get(self, "cluster_operations")

    @cluster_operations.setter
    def cluster_operations(self, value: Optional[pulumi.Input['BareMetalAdminClusterOperationsConfigArgs']]):
        pulumi.set(self, "cluster_operations", value)

    @property
    @pulumi.getter(name="controlPlane")
    def control_plane(self) -> Optional[pulumi.Input['BareMetalAdminControlPlaneConfigArgs']]:
        """
        Control plane configuration.
        """
        return pulumi.get(self, "control_plane")

    @control_plane.setter
    def control_plane(self, value: Optional[pulumi.Input['BareMetalAdminControlPlaneConfigArgs']]):
        pulumi.set(self, "control_plane", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human readable description of this bare metal admin cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> Optional[pulumi.Input['BareMetalAdminLoadBalancerConfigArgs']]:
        """
        Load balancer configuration.
        """
        return pulumi.get(self, "load_balancer")

    @load_balancer.setter
    def load_balancer(self, value: Optional[pulumi.Input['BareMetalAdminLoadBalancerConfigArgs']]):
        pulumi.set(self, "load_balancer", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maintenanceConfig")
    def maintenance_config(self) -> Optional[pulumi.Input['BareMetalAdminMaintenanceConfigArgs']]:
        """
        Maintenance configuration.
        """
        return pulumi.get(self, "maintenance_config")

    @maintenance_config.setter
    def maintenance_config(self, value: Optional[pulumi.Input['BareMetalAdminMaintenanceConfigArgs']]):
        pulumi.set(self, "maintenance_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The bare metal admin cluster resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['BareMetalAdminNetworkConfigArgs']]:
        """
        Network configuration.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['BareMetalAdminNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter(name="nodeAccessConfig")
    def node_access_config(self) -> Optional[pulumi.Input['BareMetalAdminNodeAccessConfigArgs']]:
        """
        Node access related configurations.
        """
        return pulumi.get(self, "node_access_config")

    @node_access_config.setter
    def node_access_config(self, value: Optional[pulumi.Input['BareMetalAdminNodeAccessConfigArgs']]):
        pulumi.set(self, "node_access_config", value)

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> Optional[pulumi.Input['BareMetalAdminWorkloadNodeConfigArgs']]:
        """
        Workload node configuration.
        """
        return pulumi.get(self, "node_config")

    @node_config.setter
    def node_config(self, value: Optional[pulumi.Input['BareMetalAdminWorkloadNodeConfigArgs']]):
        pulumi.set(self, "node_config", value)

    @property
    @pulumi.getter(name="osEnvironmentConfig")
    def os_environment_config(self) -> Optional[pulumi.Input['BareMetalAdminOsEnvironmentConfigArgs']]:
        """
        OS environment related configurations.
        """
        return pulumi.get(self, "os_environment_config")

    @os_environment_config.setter
    def os_environment_config(self, value: Optional[pulumi.Input['BareMetalAdminOsEnvironmentConfigArgs']]):
        pulumi.set(self, "os_environment_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def proxy(self) -> Optional[pulumi.Input['BareMetalAdminProxyConfigArgs']]:
        """
        Proxy configuration.
        """
        return pulumi.get(self, "proxy")

    @proxy.setter
    def proxy(self, value: Optional[pulumi.Input['BareMetalAdminProxyConfigArgs']]):
        pulumi.set(self, "proxy", value)

    @property
    @pulumi.getter(name="securityConfig")
    def security_config(self) -> Optional[pulumi.Input['BareMetalAdminSecurityConfigArgs']]:
        """
        Security related configuration.
        """
        return pulumi.get(self, "security_config")

    @security_config.setter
    def security_config(self, value: Optional[pulumi.Input['BareMetalAdminSecurityConfigArgs']]):
        pulumi.set(self, "security_config", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input['BareMetalAdminStorageConfigArgs']]:
        """
        Storage configuration.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input['BareMetalAdminStorageConfigArgs']]):
        pulumi.set(self, "storage", value)


class BareMetalAdminCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 bare_metal_admin_cluster_id: Optional[pulumi.Input[str]] = None,
                 bare_metal_version: Optional[pulumi.Input[str]] = None,
                 cluster_operations: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminClusterOperationsConfigArgs']]] = None,
                 control_plane: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminControlPlaneConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 load_balancer: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminLoadBalancerConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminMaintenanceConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminNetworkConfigArgs']]] = None,
                 node_access_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminNodeAccessConfigArgs']]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminWorkloadNodeConfigArgs']]] = None,
                 os_environment_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminOsEnvironmentConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminProxyConfigArgs']]] = None,
                 security_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminSecurityConfigArgs']]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminStorageConfigArgs']]] = None,
                 __props__=None):
        """
        Creates a new bare metal admin cluster in a given project and location. The API needs to be combined with creating a bootstrap cluster to work. See: https://cloud.google.com/anthos/clusters/docs/bare-metal/latest/installing/creating-clusters/create-admin-cluster-api#prepare_bootstrap_environment
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        :param pulumi.Input[str] bare_metal_admin_cluster_id: Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
        :param pulumi.Input[str] bare_metal_version: The Anthos clusters on bare metal version for the bare metal admin cluster.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminClusterOperationsConfigArgs']] cluster_operations: Cluster operations configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminControlPlaneConfigArgs']] control_plane: Control plane configuration.
        :param pulumi.Input[str] description: A human readable description of this bare metal admin cluster.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminLoadBalancerConfigArgs']] load_balancer: Load balancer configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminMaintenanceConfigArgs']] maintenance_config: Maintenance configuration.
        :param pulumi.Input[str] name: Immutable. The bare metal admin cluster resource name.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminNetworkConfigArgs']] network_config: Network configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminNodeAccessConfigArgs']] node_access_config: Node access related configurations.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminWorkloadNodeConfigArgs']] node_config: Workload node configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminOsEnvironmentConfigArgs']] os_environment_config: OS environment related configurations.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminProxyConfigArgs']] proxy: Proxy configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminSecurityConfigArgs']] security_config: Security related configuration.
        :param pulumi.Input[pulumi.InputType['BareMetalAdminStorageConfigArgs']] storage: Storage configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BareMetalAdminClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new bare metal admin cluster in a given project and location. The API needs to be combined with creating a bootstrap cluster to work. See: https://cloud.google.com/anthos/clusters/docs/bare-metal/latest/installing/creating-clusters/create-admin-cluster-api#prepare_bootstrap_environment
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param BareMetalAdminClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BareMetalAdminClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 bare_metal_admin_cluster_id: Optional[pulumi.Input[str]] = None,
                 bare_metal_version: Optional[pulumi.Input[str]] = None,
                 cluster_operations: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminClusterOperationsConfigArgs']]] = None,
                 control_plane: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminControlPlaneConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 load_balancer: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminLoadBalancerConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminMaintenanceConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminNetworkConfigArgs']]] = None,
                 node_access_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminNodeAccessConfigArgs']]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminWorkloadNodeConfigArgs']]] = None,
                 os_environment_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminOsEnvironmentConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminProxyConfigArgs']]] = None,
                 security_config: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminSecurityConfigArgs']]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['BareMetalAdminStorageConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BareMetalAdminClusterArgs.__new__(BareMetalAdminClusterArgs)

            __props__.__dict__["annotations"] = annotations
            if bare_metal_admin_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'bare_metal_admin_cluster_id'")
            __props__.__dict__["bare_metal_admin_cluster_id"] = bare_metal_admin_cluster_id
            __props__.__dict__["bare_metal_version"] = bare_metal_version
            __props__.__dict__["cluster_operations"] = cluster_operations
            __props__.__dict__["control_plane"] = control_plane
            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            __props__.__dict__["load_balancer"] = load_balancer
            __props__.__dict__["location"] = location
            __props__.__dict__["maintenance_config"] = maintenance_config
            __props__.__dict__["name"] = name
            __props__.__dict__["network_config"] = network_config
            __props__.__dict__["node_access_config"] = node_access_config
            __props__.__dict__["node_config"] = node_config
            __props__.__dict__["os_environment_config"] = os_environment_config
            __props__.__dict__["project"] = project
            __props__.__dict__["proxy"] = proxy
            __props__.__dict__["security_config"] = security_config
            __props__.__dict__["storage"] = storage
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["fleet"] = None
            __props__.__dict__["local_name"] = None
            __props__.__dict__["maintenance_status"] = None
            __props__.__dict__["reconciling"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["validation_check"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["bare_metal_admin_cluster_id", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(BareMetalAdminCluster, __self__).__init__(
            'google-native:gkeonprem/v1:BareMetalAdminCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BareMetalAdminCluster':
        """
        Get an existing BareMetalAdminCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BareMetalAdminClusterArgs.__new__(BareMetalAdminClusterArgs)

        __props__.__dict__["annotations"] = None
        __props__.__dict__["bare_metal_admin_cluster_id"] = None
        __props__.__dict__["bare_metal_version"] = None
        __props__.__dict__["cluster_operations"] = None
        __props__.__dict__["control_plane"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["endpoint"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["fleet"] = None
        __props__.__dict__["load_balancer"] = None
        __props__.__dict__["local_name"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["maintenance_config"] = None
        __props__.__dict__["maintenance_status"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_config"] = None
        __props__.__dict__["node_access_config"] = None
        __props__.__dict__["node_config"] = None
        __props__.__dict__["os_environment_config"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["proxy"] = None
        __props__.__dict__["reconciling"] = None
        __props__.__dict__["security_config"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["storage"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["validation_check"] = None
        return BareMetalAdminCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="bareMetalAdminClusterId")
    def bare_metal_admin_cluster_id(self) -> pulumi.Output[str]:
        """
        Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
        """
        return pulumi.get(self, "bare_metal_admin_cluster_id")

    @property
    @pulumi.getter(name="bareMetalVersion")
    def bare_metal_version(self) -> pulumi.Output[str]:
        """
        The Anthos clusters on bare metal version for the bare metal admin cluster.
        """
        return pulumi.get(self, "bare_metal_version")

    @property
    @pulumi.getter(name="clusterOperations")
    def cluster_operations(self) -> pulumi.Output['outputs.BareMetalAdminClusterOperationsConfigResponse']:
        """
        Cluster operations configuration.
        """
        return pulumi.get(self, "cluster_operations")

    @property
    @pulumi.getter(name="controlPlane")
    def control_plane(self) -> pulumi.Output['outputs.BareMetalAdminControlPlaneConfigResponse']:
        """
        Control plane configuration.
        """
        return pulumi.get(self, "control_plane")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which this bare metal admin cluster was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        The time at which this bare metal admin cluster was deleted. If the resource is not deleted, this must be empty
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A human readable description of this bare metal admin cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The IP address name of bare metal admin cluster's API server.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def fleet(self) -> pulumi.Output['outputs.FleetResponse']:
        """
        Fleet configuration for the cluster.
        """
        return pulumi.get(self, "fleet")

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> pulumi.Output['outputs.BareMetalAdminLoadBalancerConfigResponse']:
        """
        Load balancer configuration.
        """
        return pulumi.get(self, "load_balancer")

    @property
    @pulumi.getter(name="localName")
    def local_name(self) -> pulumi.Output[str]:
        """
        The object name of the bare metal cluster custom resource. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
        """
        return pulumi.get(self, "local_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceConfig")
    def maintenance_config(self) -> pulumi.Output['outputs.BareMetalAdminMaintenanceConfigResponse']:
        """
        Maintenance configuration.
        """
        return pulumi.get(self, "maintenance_config")

    @property
    @pulumi.getter(name="maintenanceStatus")
    def maintenance_status(self) -> pulumi.Output['outputs.BareMetalAdminMaintenanceStatusResponse']:
        """
        MaintenanceStatus representing state of maintenance.
        """
        return pulumi.get(self, "maintenance_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The bare metal admin cluster resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output['outputs.BareMetalAdminNetworkConfigResponse']:
        """
        Network configuration.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter(name="nodeAccessConfig")
    def node_access_config(self) -> pulumi.Output['outputs.BareMetalAdminNodeAccessConfigResponse']:
        """
        Node access related configurations.
        """
        return pulumi.get(self, "node_access_config")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Output['outputs.BareMetalAdminWorkloadNodeConfigResponse']:
        """
        Workload node configuration.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="osEnvironmentConfig")
    def os_environment_config(self) -> pulumi.Output['outputs.BareMetalAdminOsEnvironmentConfigResponse']:
        """
        OS environment related configurations.
        """
        return pulumi.get(self, "os_environment_config")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def proxy(self) -> pulumi.Output['outputs.BareMetalAdminProxyConfigResponse']:
        """
        Proxy configuration.
        """
        return pulumi.get(self, "proxy")

    @property
    @pulumi.getter
    def reconciling(self) -> pulumi.Output[bool]:
        """
        If set, there are currently changes in flight to the bare metal Admin Cluster.
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter(name="securityConfig")
    def security_config(self) -> pulumi.Output['outputs.BareMetalAdminSecurityConfigResponse']:
        """
        Security related configuration.
        """
        return pulumi.get(self, "security_config")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the bare metal admin cluster.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.ResourceStatusResponse']:
        """
        ResourceStatus representing detailed cluster status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output['outputs.BareMetalAdminStorageConfigResponse']:
        """
        Storage configuration.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        The unique identifier of the bare metal admin cluster.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time at which this bare metal admin cluster was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="validationCheck")
    def validation_check(self) -> pulumi.Output['outputs.ValidationCheckResponse']:
        """
        ValidationCheck representing the result of the preflight check.
        """
        return pulumi.get(self, "validation_check")

