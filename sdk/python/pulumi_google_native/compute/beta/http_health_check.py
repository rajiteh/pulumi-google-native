# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['HttpHealthCheckArgs', 'HttpHealthCheck']

@pulumi.input_type
class HttpHealthCheckArgs:
    def __init__(__self__, *,
                 check_interval_sec: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthy_threshold: Optional[pulumi.Input[int]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 request_path: Optional[pulumi.Input[str]] = None,
                 timeout_sec: Optional[pulumi.Input[int]] = None,
                 unhealthy_threshold: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a HttpHealthCheck resource.
        :param pulumi.Input[int] check_interval_sec: How often (in seconds) to send a health check. The default value is 5 seconds.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[int] healthy_threshold: A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        :param pulumi.Input[str] host: The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[int] port: The TCP port number for the HTTP health check request. The default value is 80.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] request_path: The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
        :param pulumi.Input[int] timeout_sec: How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        :param pulumi.Input[int] unhealthy_threshold: A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        if check_interval_sec is not None:
            pulumi.set(__self__, "check_interval_sec", check_interval_sec)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if healthy_threshold is not None:
            pulumi.set(__self__, "healthy_threshold", healthy_threshold)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if request_path is not None:
            pulumi.set(__self__, "request_path", request_path)
        if timeout_sec is not None:
            pulumi.set(__self__, "timeout_sec", timeout_sec)
        if unhealthy_threshold is not None:
            pulumi.set(__self__, "unhealthy_threshold", unhealthy_threshold)

    @property
    @pulumi.getter(name="checkIntervalSec")
    def check_interval_sec(self) -> Optional[pulumi.Input[int]]:
        """
        How often (in seconds) to send a health check. The default value is 5 seconds.
        """
        return pulumi.get(self, "check_interval_sec")

    @check_interval_sec.setter
    def check_interval_sec(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "check_interval_sec", value)

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
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        """
        return pulumi.get(self, "healthy_threshold")

    @healthy_threshold.setter
    def healthy_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "healthy_threshold", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

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
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The TCP port number for the HTTP health check request. The default value is 80.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

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
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="requestPath")
    def request_path(self) -> Optional[pulumi.Input[str]]:
        """
        The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
        """
        return pulumi.get(self, "request_path")

    @request_path.setter
    def request_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_path", value)

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> Optional[pulumi.Input[int]]:
        """
        How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        """
        return pulumi.get(self, "timeout_sec")

    @timeout_sec.setter
    def timeout_sec(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_sec", value)

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        return pulumi.get(self, "unhealthy_threshold")

    @unhealthy_threshold.setter
    def unhealthy_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unhealthy_threshold", value)


class HttpHealthCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_interval_sec: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthy_threshold: Optional[pulumi.Input[int]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 request_path: Optional[pulumi.Input[str]] = None,
                 timeout_sec: Optional[pulumi.Input[int]] = None,
                 unhealthy_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a HttpHealthCheck resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] check_interval_sec: How often (in seconds) to send a health check. The default value is 5 seconds.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[int] healthy_threshold: A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        :param pulumi.Input[str] host: The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[int] port: The TCP port number for the HTTP health check request. The default value is 80.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] request_path: The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
        :param pulumi.Input[int] timeout_sec: How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        :param pulumi.Input[int] unhealthy_threshold: A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[HttpHealthCheckArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a HttpHealthCheck resource in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param HttpHealthCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HttpHealthCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_interval_sec: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 healthy_threshold: Optional[pulumi.Input[int]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 request_path: Optional[pulumi.Input[str]] = None,
                 timeout_sec: Optional[pulumi.Input[int]] = None,
                 unhealthy_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HttpHealthCheckArgs.__new__(HttpHealthCheckArgs)

            __props__.__dict__["check_interval_sec"] = check_interval_sec
            __props__.__dict__["description"] = description
            __props__.__dict__["healthy_threshold"] = healthy_threshold
            __props__.__dict__["host"] = host
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["request_path"] = request_path
            __props__.__dict__["timeout_sec"] = timeout_sec
            __props__.__dict__["unhealthy_threshold"] = unhealthy_threshold
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["self_link"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(HttpHealthCheck, __self__).__init__(
            'google-native:compute/beta:HttpHealthCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HttpHealthCheck':
        """
        Get an existing HttpHealthCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HttpHealthCheckArgs.__new__(HttpHealthCheckArgs)

        __props__.__dict__["check_interval_sec"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["healthy_threshold"] = None
        __props__.__dict__["host"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["request_path"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["timeout_sec"] = None
        __props__.__dict__["unhealthy_threshold"] = None
        return HttpHealthCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="checkIntervalSec")
    def check_interval_sec(self) -> pulumi.Output[int]:
        """
        How often (in seconds) to send a health check. The default value is 5 seconds.
        """
        return pulumi.get(self, "check_interval_sec")

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
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> pulumi.Output[int]:
        """
        A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        """
        return pulumi.get(self, "healthy_threshold")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#httpHealthCheck for HTTP health checks.
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
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The TCP port number for the HTTP health check request. The default value is 80.
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
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="requestPath")
    def request_path(self) -> pulumi.Output[str]:
        """
        The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
        """
        return pulumi.get(self, "request_path")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> pulumi.Output[int]:
        """
        How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        """
        return pulumi.get(self, "timeout_sec")

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> pulumi.Output[int]:
        """
        A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        """
        return pulumi.get(self, "unhealthy_threshold")

