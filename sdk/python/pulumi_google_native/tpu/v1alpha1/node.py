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

__all__ = ['NodeArgs', 'Node']

@pulumi.input_type
class NodeArgs:
    def __init__(__self__, *,
                 accelerator_type: pulumi.Input[str],
                 tensorflow_version: pulumi.Input[str],
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 health: Optional[pulumi.Input['NodeHealth']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scheduling_config: Optional[pulumi.Input['SchedulingConfigArgs']] = None,
                 use_service_networking: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Node resource.
        :param pulumi.Input[str] accelerator_type: The type of hardware accelerators associated with this node.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[str] cidr_block: The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        :param pulumi.Input[str] description: The user-supplied description of the TPU. Maximum of 512 characters.
        :param pulumi.Input['NodeHealth'] health: The health status of the TPU node.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] network: The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
        :param pulumi.Input[str] node_id: The unqualified resource name.
        :param pulumi.Input[str] request_id: Idempotent request UUID.
        :param pulumi.Input['SchedulingConfigArgs'] scheduling_config: The scheduling options for this node.
        :param pulumi.Input[bool] use_service_networking: Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
        """
        pulumi.set(__self__, "accelerator_type", accelerator_type)
        pulumi.set(__self__, "tensorflow_version", tensorflow_version)
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if health is not None:
            pulumi.set(__self__, "health", health)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if scheduling_config is not None:
            pulumi.set(__self__, "scheduling_config", scheduling_config)
        if use_service_networking is not None:
            pulumi.set(__self__, "use_service_networking", use_service_networking)

    @property
    @pulumi.getter(name="acceleratorType")
    def accelerator_type(self) -> pulumi.Input[str]:
        """
        The type of hardware accelerators associated with this node.
        """
        return pulumi.get(self, "accelerator_type")

    @accelerator_type.setter
    def accelerator_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "accelerator_type", value)

    @property
    @pulumi.getter(name="tensorflowVersion")
    def tensorflow_version(self) -> pulumi.Input[str]:
        """
        The version of Tensorflow running in the Node.
        """
        return pulumi.get(self, "tensorflow_version")

    @tensorflow_version.setter
    def tensorflow_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "tensorflow_version", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The user-supplied description of the TPU. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def health(self) -> Optional[pulumi.Input['NodeHealth']]:
        """
        The health status of the TPU node.
        """
        return pulumi.get(self, "health")

    @health.setter
    def health(self, value: Optional[pulumi.Input['NodeHealth']]):
        pulumi.set(self, "health", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unqualified resource name.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        Idempotent request UUID.
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="schedulingConfig")
    def scheduling_config(self) -> Optional[pulumi.Input['SchedulingConfigArgs']]:
        """
        The scheduling options for this node.
        """
        return pulumi.get(self, "scheduling_config")

    @scheduling_config.setter
    def scheduling_config(self, value: Optional[pulumi.Input['SchedulingConfigArgs']]):
        pulumi.set(self, "scheduling_config", value)

    @property
    @pulumi.getter(name="useServiceNetworking")
    def use_service_networking(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
        """
        return pulumi.get(self, "use_service_networking")

    @use_service_networking.setter
    def use_service_networking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_service_networking", value)


class Node(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_type: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 health: Optional[pulumi.Input['NodeHealth']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scheduling_config: Optional[pulumi.Input[pulumi.InputType['SchedulingConfigArgs']]] = None,
                 tensorflow_version: Optional[pulumi.Input[str]] = None,
                 use_service_networking: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Creates a node.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_type: The type of hardware accelerators associated with this node.
        :param pulumi.Input[str] cidr_block: The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        :param pulumi.Input[str] description: The user-supplied description of the TPU. Maximum of 512 characters.
        :param pulumi.Input['NodeHealth'] health: The health status of the TPU node.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] network: The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
        :param pulumi.Input[str] node_id: The unqualified resource name.
        :param pulumi.Input[str] request_id: Idempotent request UUID.
        :param pulumi.Input[pulumi.InputType['SchedulingConfigArgs']] scheduling_config: The scheduling options for this node.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[bool] use_service_networking: Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a node.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param NodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_type: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 health: Optional[pulumi.Input['NodeHealth']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scheduling_config: Optional[pulumi.Input[pulumi.InputType['SchedulingConfigArgs']]] = None,
                 tensorflow_version: Optional[pulumi.Input[str]] = None,
                 use_service_networking: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeArgs.__new__(NodeArgs)

            if accelerator_type is None and not opts.urn:
                raise TypeError("Missing required property 'accelerator_type'")
            __props__.__dict__["accelerator_type"] = accelerator_type
            __props__.__dict__["cidr_block"] = cidr_block
            __props__.__dict__["description"] = description
            __props__.__dict__["health"] = health
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["network"] = network
            __props__.__dict__["node_id"] = node_id
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["scheduling_config"] = scheduling_config
            if tensorflow_version is None and not opts.urn:
                raise TypeError("Missing required property 'tensorflow_version'")
            __props__.__dict__["tensorflow_version"] = tensorflow_version
            __props__.__dict__["use_service_networking"] = use_service_networking
            __props__.__dict__["api_version"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["health_description"] = None
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["network_endpoints"] = None
            __props__.__dict__["port"] = None
            __props__.__dict__["service_account"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["symptoms"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Node, __self__).__init__(
            'google-native:tpu/v1alpha1:Node',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Node':
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NodeArgs.__new__(NodeArgs)

        __props__.__dict__["accelerator_type"] = None
        __props__.__dict__["api_version"] = None
        __props__.__dict__["cidr_block"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["health"] = None
        __props__.__dict__["health_description"] = None
        __props__.__dict__["ip_address"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network"] = None
        __props__.__dict__["network_endpoints"] = None
        __props__.__dict__["node_id"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["scheduling_config"] = None
        __props__.__dict__["service_account"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["symptoms"] = None
        __props__.__dict__["tensorflow_version"] = None
        __props__.__dict__["use_service_networking"] = None
        return Node(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceleratorType")
    def accelerator_type(self) -> pulumi.Output[str]:
        """
        The type of hardware accelerators associated with this node.
        """
        return pulumi.get(self, "accelerator_type")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        The API version that created this Node.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        """
        The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the node was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The user-supplied description of the TPU. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def health(self) -> pulumi.Output[str]:
        """
        The health status of the TPU node.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter(name="healthDescription")
    def health_description(self) -> pulumi.Output[str]:
        """
        If this field is populated, it contains a description of why the TPU Node is unhealthy.
        """
        return pulumi.get(self, "health_description")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The name of the TPU
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkEndpoints")
    def network_endpoints(self) -> pulumi.Output[Sequence['outputs.NetworkEndpointResponse']]:
        """
        The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
        """
        return pulumi.get(self, "network_endpoints")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Output[Optional[str]]:
        """
        The unqualified resource name.
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        Idempotent request UUID.
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="schedulingConfig")
    def scheduling_config(self) -> pulumi.Output['outputs.SchedulingConfigResponse']:
        """
        The scheduling options for this node.
        """
        return pulumi.get(self, "scheduling_config")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[str]:
        """
        The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state for the TPU Node.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def symptoms(self) -> pulumi.Output[Sequence['outputs.SymptomResponse']]:
        """
        The Symptoms that have occurred to the TPU Node.
        """
        return pulumi.get(self, "symptoms")

    @property
    @pulumi.getter(name="tensorflowVersion")
    def tensorflow_version(self) -> pulumi.Output[str]:
        """
        The version of Tensorflow running in the Node.
        """
        return pulumi.get(self, "tensorflow_version")

    @property
    @pulumi.getter(name="useServiceNetworking")
    def use_service_networking(self) -> pulumi.Output[bool]:
        """
        Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
        """
        return pulumi.get(self, "use_service_networking")

