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
    'GetTopicResult',
    'AwaitableGetTopicResult',
    'get_topic',
    'get_topic_output',
]

@pulumi.output_type
class GetTopicResult:
    def __init__(__self__, kms_key_name=None, labels=None, message_retention_duration=None, message_storage_policy=None, name=None, satisfies_pzs=None, schema_settings=None):
        if kms_key_name and not isinstance(kms_key_name, str):
            raise TypeError("Expected argument 'kms_key_name' to be a str")
        pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if message_retention_duration and not isinstance(message_retention_duration, str):
            raise TypeError("Expected argument 'message_retention_duration' to be a str")
        pulumi.set(__self__, "message_retention_duration", message_retention_duration)
        if message_storage_policy and not isinstance(message_storage_policy, dict):
            raise TypeError("Expected argument 'message_storage_policy' to be a dict")
        pulumi.set(__self__, "message_storage_policy", message_storage_policy)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if schema_settings and not isinstance(schema_settings, dict):
            raise TypeError("Expected argument 'schema_settings' to be a dict")
        pulumi.set(__self__, "schema_settings", schema_settings)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        See [Creating and managing labels] (https://cloud.google.com/pubsub/docs/labels).
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="messageRetentionDuration")
    def message_retention_duration(self) -> str:
        """
        Indicates the minimum duration to retain a message after it is published to the topic. If this field is set, messages published to the topic in the last `message_retention_duration` are always available to subscribers. For instance, it allows any attached subscription to [seek to a timestamp](https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) that is up to `message_retention_duration` in the past. If this field is not set, message retention is controlled by settings on individual subscriptions. Cannot be more than 31 days or less than 10 minutes.
        """
        return pulumi.get(self, "message_retention_duration")

    @property
    @pulumi.getter(name="messageStoragePolicy")
    def message_storage_policy(self) -> 'outputs.MessageStoragePolicyResponse':
        """
        Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect.
        """
        return pulumi.get(self, "message_storage_policy")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the topic. It must have the format `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="schemaSettings")
    def schema_settings(self) -> 'outputs.SchemaSettingsResponse':
        """
        Settings for validating messages published against a schema.
        """
        return pulumi.get(self, "schema_settings")


class AwaitableGetTopicResult(GetTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicResult(
            kms_key_name=self.kms_key_name,
            labels=self.labels,
            message_retention_duration=self.message_retention_duration,
            message_storage_policy=self.message_storage_policy,
            name=self.name,
            satisfies_pzs=self.satisfies_pzs,
            schema_settings=self.schema_settings)


def get_topic(project: Optional[str] = None,
              topic_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicResult:
    """
    Gets the configuration of a topic.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['topicId'] = topic_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:pubsub/v1:getTopic', __args__, opts=opts, typ=GetTopicResult).value

    return AwaitableGetTopicResult(
        kms_key_name=__ret__.kms_key_name,
        labels=__ret__.labels,
        message_retention_duration=__ret__.message_retention_duration,
        message_storage_policy=__ret__.message_storage_policy,
        name=__ret__.name,
        satisfies_pzs=__ret__.satisfies_pzs,
        schema_settings=__ret__.schema_settings)


@_utilities.lift_output_func(get_topic)
def get_topic_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                     topic_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicResult]:
    """
    Gets the configuration of a topic.
    """
    ...
