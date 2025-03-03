# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetSourceResult',
    'AwaitableGetSourceResult',
    'get_source',
    'get_source_output',
]

@pulumi.output_type
class GetSourceResult:
    def __init__(__self__, canonical_name=None, description=None, display_name=None, name=None):
        if canonical_name and not isinstance(canonical_name, str):
            raise TypeError("Expected argument 'canonical_name' to be a str")
        pulumi.set(__self__, "canonical_name", canonical_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="canonicalName")
    def canonical_name(self) -> str:
        """
        The canonical name of the finding. It's either "organizations/{organization_id}/sources/{source_id}", "folders/{folder_id}/sources/{source_id}" or "projects/{project_number}/sources/{source_id}", depending on the closest CRM ancestor of the resource associated with the finding.
        """
        return pulumi.get(self, "canonical_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated or insecure libraries."
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
        """
        return pulumi.get(self, "name")


class AwaitableGetSourceResult(GetSourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourceResult(
            canonical_name=self.canonical_name,
            description=self.description,
            display_name=self.display_name,
            name=self.name)


def get_source(organization_id: Optional[str] = None,
               source_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourceResult:
    """
    Gets a source.
    """
    __args__ = dict()
    __args__['organizationId'] = organization_id
    __args__['sourceId'] = source_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:securitycenter/v1:getSource', __args__, opts=opts, typ=GetSourceResult).value

    return AwaitableGetSourceResult(
        canonical_name=__ret__.canonical_name,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name)


@_utilities.lift_output_func(get_source)
def get_source_output(organization_id: Optional[pulumi.Input[str]] = None,
                      source_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSourceResult]:
    """
    Gets a source.
    """
    ...
