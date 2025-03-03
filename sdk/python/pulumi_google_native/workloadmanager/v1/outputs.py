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
    'GceInstanceFilterResponse',
    'ResourceFilterResponse',
    'ResourceStatusResponse',
]

@pulumi.output_type
class GceInstanceFilterResponse(dict):
    """
    Message describing compute engine instance filter
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceAccounts":
            suggest = "service_accounts"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GceInstanceFilterResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GceInstanceFilterResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GceInstanceFilterResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 service_accounts: Sequence[str]):
        """
        Message describing compute engine instance filter
        :param Sequence[str] service_accounts: Service account of compute engine
        """
        pulumi.set(__self__, "service_accounts", service_accounts)

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> Sequence[str]:
        """
        Service account of compute engine
        """
        return pulumi.get(self, "service_accounts")


@pulumi.output_type
class ResourceFilterResponse(dict):
    """
    Message describing resource filters
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "gceInstanceFilter":
            suggest = "gce_instance_filter"
        elif key == "inclusionLabels":
            suggest = "inclusion_labels"
        elif key == "resourceIdPatterns":
            suggest = "resource_id_patterns"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceFilterResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceFilterResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceFilterResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gce_instance_filter: 'outputs.GceInstanceFilterResponse',
                 inclusion_labels: Mapping[str, str],
                 resource_id_patterns: Sequence[str],
                 scopes: Sequence[str]):
        """
        Message describing resource filters
        :param 'GceInstanceFilterResponse' gce_instance_filter: Filter compute engine resource
        :param Mapping[str, str] inclusion_labels: The label used for filter resource
        :param Sequence[str] resource_id_patterns: The id pattern for filter resource
        :param Sequence[str] scopes: The scopes of evaluation resource
        """
        pulumi.set(__self__, "gce_instance_filter", gce_instance_filter)
        pulumi.set(__self__, "inclusion_labels", inclusion_labels)
        pulumi.set(__self__, "resource_id_patterns", resource_id_patterns)
        pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="gceInstanceFilter")
    def gce_instance_filter(self) -> 'outputs.GceInstanceFilterResponse':
        """
        Filter compute engine resource
        """
        return pulumi.get(self, "gce_instance_filter")

    @property
    @pulumi.getter(name="inclusionLabels")
    def inclusion_labels(self) -> Mapping[str, str]:
        """
        The label used for filter resource
        """
        return pulumi.get(self, "inclusion_labels")

    @property
    @pulumi.getter(name="resourceIdPatterns")
    def resource_id_patterns(self) -> Sequence[str]:
        """
        The id pattern for filter resource
        """
        return pulumi.get(self, "resource_id_patterns")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        The scopes of evaluation resource
        """
        return pulumi.get(self, "scopes")


@pulumi.output_type
class ResourceStatusResponse(dict):
    """
    Message describing resource status
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "rulesNewerVersions":
            suggest = "rules_newer_versions"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceStatusResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceStatusResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceStatusResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 rules_newer_versions: Sequence[str],
                 state: str):
        """
        Message describing resource status
        :param Sequence[str] rules_newer_versions: the new version of rule id if exists
        :param str state: State of the resource
        """
        pulumi.set(__self__, "rules_newer_versions", rules_newer_versions)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="rulesNewerVersions")
    def rules_newer_versions(self) -> Sequence[str]:
        """
        the new version of rule id if exists
        """
        return pulumi.get(self, "rules_newer_versions")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the resource
        """
        return pulumi.get(self, "state")


