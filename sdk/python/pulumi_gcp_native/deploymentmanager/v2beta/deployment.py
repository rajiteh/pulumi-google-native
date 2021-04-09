# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Deployment']


class Deployment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 insert_time: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeploymentLabelEntryArgs']]]]] = None,
                 manifest: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation: Optional[pulumi.Input[pulumi.InputType['OperationArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[pulumi.InputType['TargetConfigurationArgs']]] = None,
                 update: Optional[pulumi.Input[pulumi.InputType['DeploymentUpdateArgs']]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a deployment and all of the resources described by the deployment manifest.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional user-provided description of the deployment.
        :param pulumi.Input[str] fingerprint: Provides a fingerprint to use in requests to modify a deployment, such as `update()`, `stop()`, and `cancelPreview()` requests. A fingerprint is a randomly generated value that must be provided with `update()`, `stop()`, and `cancelPreview()` requests to perform optimistic locking. This ensures optimistic concurrency so that only one request happens at a time. The fingerprint is initially generated by Deployment Manager and changes after every request to modify data. To get the latest fingerprint value, perform a `get()` request to a deployment.
        :param pulumi.Input[str] insert_time: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeploymentLabelEntryArgs']]]] labels: Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        :param pulumi.Input[str] manifest: URL of the manifest representing the last manifest that was successfully deployed. If no manifest has been successfully deployed, this field will be absent.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[pulumi.InputType['OperationArgs']] operation: The Operation that most recently ran, or is currently running, on this deployment.
        :param pulumi.Input[str] self_link: Server defined URL for the resource.
        :param pulumi.Input[pulumi.InputType['TargetConfigurationArgs']] target: [Input Only] The parameters that define your deployment, including the deployment configuration and relevant templates.
        :param pulumi.Input[pulumi.InputType['DeploymentUpdateArgs']] update: If Deployment Manager is currently updating or previewing an update to this deployment, the updated configuration appears here.
        :param pulumi.Input[str] update_time: Update timestamp in RFC3339 text format.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if deployment is None and not opts.urn:
                raise TypeError("Missing required property 'deployment'")
            __props__['deployment'] = deployment
            __props__['description'] = description
            __props__['fingerprint'] = fingerprint
            __props__['id'] = id
            __props__['insert_time'] = insert_time
            __props__['labels'] = labels
            __props__['manifest'] = manifest
            __props__['name'] = name
            __props__['operation'] = operation
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['self_link'] = self_link
            __props__['target'] = target
            __props__['update'] = update
            __props__['update_time'] = update_time
        super(Deployment, __self__).__init__(
            'gcp-native:deploymentmanager/v2beta:Deployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Deployment':
        """
        Get an existing Deployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = None
        __props__["fingerprint"] = None
        __props__["insert_time"] = None
        __props__["labels"] = None
        __props__["manifest"] = None
        __props__["name"] = None
        __props__["operation"] = None
        __props__["self_link"] = None
        __props__["target"] = None
        __props__["update"] = None
        __props__["update_time"] = None
        return Deployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional user-provided description of the deployment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Provides a fingerprint to use in requests to modify a deployment, such as `update()`, `stop()`, and `cancelPreview()` requests. A fingerprint is a randomly generated value that must be provided with `update()`, `stop()`, and `cancelPreview()` requests to perform optimistic locking. This ensures optimistic concurrency so that only one request happens at a time. The fingerprint is initially generated by Deployment Manager and changes after every request to modify data. To get the latest fingerprint value, perform a `get()` request to a deployment.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="insertTime")
    def insert_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "insert_time")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Sequence['outputs.DeploymentLabelEntryResponse']]:
        """
        Map of One Platform labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def manifest(self) -> pulumi.Output[str]:
        """
        URL of the manifest representing the last manifest that was successfully deployed. If no manifest has been successfully deployed, this field will be absent.
        """
        return pulumi.get(self, "manifest")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operation(self) -> pulumi.Output['outputs.OperationResponse']:
        """
        The Operation that most recently ran, or is currently running, on this deployment.
        """
        return pulumi.get(self, "operation")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output['outputs.TargetConfigurationResponse']:
        """
        [Input Only] The parameters that define your deployment, including the deployment configuration and relevant templates.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def update(self) -> pulumi.Output['outputs.DeploymentUpdateResponse']:
        """
        If Deployment Manager is currently updating or previewing an update to this deployment, the updated configuration appears here.
        """
        return pulumi.get(self, "update")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Update timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "update_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
