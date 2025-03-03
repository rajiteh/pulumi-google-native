# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['AccessPolicyArgs', 'AccessPolicy']

@pulumi.input_type
class AccessPolicyArgs:
    def __init__(__self__, *,
                 parent: pulumi.Input[str],
                 title: pulumi.Input[str],
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AccessPolicy resource.
        :param pulumi.Input[str] parent: The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        :param pulumi.Input[str] title: Human readable title. Does not affect behavior.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        """
        pulumi.set(__self__, "parent", parent)
        pulumi.set(__self__, "title", title)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Human readable title. Does not affect behavior.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


class AccessPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an access policy. This method fails if the organization already has an access policy. The long-running operation has a successful status after the access policy propagates to long-lasting storage. Syntactic and basic semantic errors are returned in `metadata` as a BadRequest proto.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent: The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        :param pulumi.Input[str] title: Human readable title. Does not affect behavior.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an access policy. This method fails if the organization already has an access policy. The long-running operation has a successful status after the access policy propagates to long-lasting storage. Syntactic and basic semantic errors are returned in `metadata` as a BadRequest proto.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param AccessPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPolicyArgs.__new__(AccessPolicyArgs)

            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["scopes"] = scopes
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
        super(AccessPolicy, __self__).__init__(
            'google-native:accesscontextmanager/v1:AccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessPolicy':
        """
        Get an existing AccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccessPolicyArgs.__new__(AccessPolicyArgs)

        __props__.__dict__["etag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parent"] = None
        __props__.__dict__["scopes"] = None
        __props__.__dict__["title"] = None
        return AccessPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the `AccessPolicy`. Format: `accessPolicies/{access_policy}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        The scopes of a policy define which resources an ACM policy can restrict, and where ACM resources can be referenced. For example, a policy with scopes=["folders/123"] has the following behavior: - vpcsc perimeters can only restrict projects within folders/123 - access levels can only be referenced by resources within folders/123. If empty, there are no limitations on which resources can be restricted by an ACM policy, and there are no limitations on where ACM resources can be referenced. Only one policy can include a given scope (attempting to create a second policy which includes "folders/123" will result in an error). Currently, scopes cannot be modified after a policy is created. Currently, policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Human readable title. Does not affect behavior.
        """
        return pulumi.get(self, "title")

