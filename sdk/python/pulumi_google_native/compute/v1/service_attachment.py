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

__all__ = ['ServiceAttachmentArgs', 'ServiceAttachment']

@pulumi.input_type
class ServiceAttachmentArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 connection_preference: Optional[pulumi.Input['ServiceAttachmentConnectionPreference']] = None,
                 consumer_accept_lists: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerProjectLimitArgs']]]] = None,
                 consumer_reject_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reconcile_connections: Optional[pulumi.Input[bool]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 target_service: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceAttachment resource.
        :param pulumi.Input['ServiceAttachmentConnectionPreference'] connection_preference: The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerProjectLimitArgs']]] consumer_accept_lists: Projects that are allowed to connect to this service attachment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] consumer_reject_lists: Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
        :param pulumi.Input[bool] enable_proxy_protocol: If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_subnets: An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        :param pulumi.Input[str] producer_forwarding_rule: The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        :param pulumi.Input[bool] reconcile_connections: This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] target_service: The URL of a service serving the endpoint identified by this service attachment.
        """
        pulumi.set(__self__, "region", region)
        if connection_preference is not None:
            pulumi.set(__self__, "connection_preference", connection_preference)
        if consumer_accept_lists is not None:
            pulumi.set(__self__, "consumer_accept_lists", consumer_accept_lists)
        if consumer_reject_lists is not None:
            pulumi.set(__self__, "consumer_reject_lists", consumer_reject_lists)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if enable_proxy_protocol is not None:
            pulumi.set(__self__, "enable_proxy_protocol", enable_proxy_protocol)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat_subnets is not None:
            pulumi.set(__self__, "nat_subnets", nat_subnets)
        if producer_forwarding_rule is not None:
            pulumi.set(__self__, "producer_forwarding_rule", producer_forwarding_rule)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if reconcile_connections is not None:
            pulumi.set(__self__, "reconcile_connections", reconcile_connections)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if target_service is not None:
            pulumi.set(__self__, "target_service", target_service)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> Optional[pulumi.Input['ServiceAttachmentConnectionPreference']]:
        """
        The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        """
        return pulumi.get(self, "connection_preference")

    @connection_preference.setter
    def connection_preference(self, value: Optional[pulumi.Input['ServiceAttachmentConnectionPreference']]):
        pulumi.set(self, "connection_preference", value)

    @property
    @pulumi.getter(name="consumerAcceptLists")
    def consumer_accept_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerProjectLimitArgs']]]]:
        """
        Projects that are allowed to connect to this service attachment.
        """
        return pulumi.get(self, "consumer_accept_lists")

    @consumer_accept_lists.setter
    def consumer_accept_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerProjectLimitArgs']]]]):
        pulumi.set(self, "consumer_accept_lists", value)

    @property
    @pulumi.getter(name="consumerRejectLists")
    def consumer_reject_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        """
        return pulumi.get(self, "consumer_reject_lists")

    @consumer_reject_lists.setter
    def consumer_reject_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "consumer_reject_lists", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @enable_proxy_protocol.setter
    def enable_proxy_protocol(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_proxy_protocol", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="natSubnets")
    def nat_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        """
        return pulumi.get(self, "nat_subnets")

    @nat_subnets.setter
    def nat_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nat_subnets", value)

    @property
    @pulumi.getter(name="producerForwardingRule")
    def producer_forwarding_rule(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "producer_forwarding_rule")

    @producer_forwarding_rule.setter
    def producer_forwarding_rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "producer_forwarding_rule", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="reconcileConnections")
    def reconcile_connections(self) -> Optional[pulumi.Input[bool]]:
        """
        This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
        """
        return pulumi.get(self, "reconcile_connections")

    @reconcile_connections.setter
    def reconcile_connections(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reconcile_connections", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="targetService")
    def target_service(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of a service serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "target_service")

    @target_service.setter
    def target_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_service", value)


class ServiceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_preference: Optional[pulumi.Input['ServiceAttachmentConnectionPreference']] = None,
                 consumer_accept_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerProjectLimitArgs']]]]] = None,
                 consumer_reject_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reconcile_connections: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 target_service: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['ServiceAttachmentConnectionPreference'] connection_preference: The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerProjectLimitArgs']]]] consumer_accept_lists: Projects that are allowed to connect to this service attachment.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] consumer_reject_lists: Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
        :param pulumi.Input[bool] enable_proxy_protocol: If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_subnets: An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        :param pulumi.Input[str] producer_forwarding_rule: The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        :param pulumi.Input[bool] reconcile_connections: This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] target_service: The URL of a service serving the endpoint identified by this service attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param ServiceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_preference: Optional[pulumi.Input['ServiceAttachmentConnectionPreference']] = None,
                 consumer_accept_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerProjectLimitArgs']]]]] = None,
                 consumer_reject_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reconcile_connections: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 target_service: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceAttachmentArgs.__new__(ServiceAttachmentArgs)

            __props__.__dict__["connection_preference"] = connection_preference
            __props__.__dict__["consumer_accept_lists"] = consumer_accept_lists
            __props__.__dict__["consumer_reject_lists"] = consumer_reject_lists
            __props__.__dict__["description"] = description
            __props__.__dict__["domain_names"] = domain_names
            __props__.__dict__["enable_proxy_protocol"] = enable_proxy_protocol
            __props__.__dict__["name"] = name
            __props__.__dict__["nat_subnets"] = nat_subnets
            __props__.__dict__["producer_forwarding_rule"] = producer_forwarding_rule
            __props__.__dict__["project"] = project
            __props__.__dict__["reconcile_connections"] = reconcile_connections
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["target_service"] = target_service
            __props__.__dict__["connected_endpoints"] = None
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["psc_service_attachment_id"] = None
            __props__.__dict__["self_link"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project", "region"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ServiceAttachment, __self__).__init__(
            'google-native:compute/v1:ServiceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceAttachment':
        """
        Get an existing ServiceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceAttachmentArgs.__new__(ServiceAttachmentArgs)

        __props__.__dict__["connected_endpoints"] = None
        __props__.__dict__["connection_preference"] = None
        __props__.__dict__["consumer_accept_lists"] = None
        __props__.__dict__["consumer_reject_lists"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["domain_names"] = None
        __props__.__dict__["enable_proxy_protocol"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nat_subnets"] = None
        __props__.__dict__["producer_forwarding_rule"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["psc_service_attachment_id"] = None
        __props__.__dict__["reconcile_connections"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["target_service"] = None
        return ServiceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectedEndpoints")
    def connected_endpoints(self) -> pulumi.Output[Sequence['outputs.ServiceAttachmentConnectedEndpointResponse']]:
        """
        An array of connections for all the consumers connected to this service attachment.
        """
        return pulumi.get(self, "connected_endpoints")

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> pulumi.Output[str]:
        """
        The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        """
        return pulumi.get(self, "connection_preference")

    @property
    @pulumi.getter(name="consumerAcceptLists")
    def consumer_accept_lists(self) -> pulumi.Output[Sequence['outputs.ServiceAttachmentConsumerProjectLimitResponse']]:
        """
        Projects that are allowed to connect to this service attachment.
        """
        return pulumi.get(self, "consumer_accept_lists")

    @property
    @pulumi.getter(name="consumerRejectLists")
    def consumer_reject_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        """
        return pulumi.get(self, "consumer_reject_lists")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Output[Sequence[str]]:
        """
        If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS. For example, this is a valid domain name: "p.mycompany.com.". Current max number of domain names supported is 1.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> pulumi.Output[bool]:
        """
        If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#serviceAttachment for service attachments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natSubnets")
    def nat_subnets(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        """
        return pulumi.get(self, "nat_subnets")

    @property
    @pulumi.getter(name="producerForwardingRule")
    def producer_forwarding_rule(self) -> pulumi.Output[str]:
        """
        The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "producer_forwarding_rule")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pscServiceAttachmentId")
    def psc_service_attachment_id(self) -> pulumi.Output['outputs.Uint128Response']:
        """
        An 128-bit global unique ID of the PSC service attachment.
        """
        return pulumi.get(self, "psc_service_attachment_id")

    @property
    @pulumi.getter(name="reconcileConnections")
    def reconcile_connections(self) -> pulumi.Output[bool]:
        """
        This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints. - If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified . - If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list. For newly created service attachment, this boolean defaults to true.
        """
        return pulumi.get(self, "reconcile_connections")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="targetService")
    def target_service(self) -> pulumi.Output[str]:
        """
        The URL of a service serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "target_service")

