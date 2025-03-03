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
    'GetConsentResult',
    'AwaitableGetConsentResult',
    'get_consent',
    'get_consent_output',
]

@pulumi.output_type
class GetConsentResult:
    def __init__(__self__, consent_artifact=None, expire_time=None, metadata=None, name=None, policies=None, revision_create_time=None, revision_id=None, state=None, ttl=None, user_id=None):
        if consent_artifact and not isinstance(consent_artifact, str):
            raise TypeError("Expected argument 'consent_artifact' to be a str")
        pulumi.set(__self__, "consent_artifact", consent_artifact)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if revision_create_time and not isinstance(revision_create_time, str):
            raise TypeError("Expected argument 'revision_create_time' to be a str")
        pulumi.set(__self__, "revision_create_time", revision_create_time)
        if revision_id and not isinstance(revision_id, str):
            raise TypeError("Expected argument 'revision_id' to be a str")
        pulumi.set(__self__, "revision_id", revision_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if ttl and not isinstance(ttl, str):
            raise TypeError("Expected argument 'ttl' to be a str")
        pulumi.set(__self__, "ttl", ttl)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="consentArtifact")
    def consent_artifact(self) -> str:
        """
        The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        """
        return pulumi.get(self, "consent_artifact")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        Timestamp in UTC of when this Consent is considered expired.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GoogleCloudHealthcareV1beta1ConsentPolicyResponse']:
        """
        Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="revisionCreateTime")
    def revision_create_time(self) -> str:
        """
        The timestamp that the revision was created.
        """
        return pulumi.get(self, "revision_create_time")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> str:
        """
        The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Indicates the current state of this Consent.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def ttl(self) -> str:
        """
        Input only. The time to live for this Consent from when it is created.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")


class AwaitableGetConsentResult(GetConsentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsentResult(
            consent_artifact=self.consent_artifact,
            expire_time=self.expire_time,
            metadata=self.metadata,
            name=self.name,
            policies=self.policies,
            revision_create_time=self.revision_create_time,
            revision_id=self.revision_id,
            state=self.state,
            ttl=self.ttl,
            user_id=self.user_id)


def get_consent(consent_id: Optional[str] = None,
                consent_store_id: Optional[str] = None,
                dataset_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsentResult:
    """
    Gets the specified revision of a Consent, or the latest revision if `revision_id` is not specified in the resource name.
    """
    __args__ = dict()
    __args__['consentId'] = consent_id
    __args__['consentStoreId'] = consent_store_id
    __args__['datasetId'] = dataset_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:healthcare/v1beta1:getConsent', __args__, opts=opts, typ=GetConsentResult).value

    return AwaitableGetConsentResult(
        consent_artifact=__ret__.consent_artifact,
        expire_time=__ret__.expire_time,
        metadata=__ret__.metadata,
        name=__ret__.name,
        policies=__ret__.policies,
        revision_create_time=__ret__.revision_create_time,
        revision_id=__ret__.revision_id,
        state=__ret__.state,
        ttl=__ret__.ttl,
        user_id=__ret__.user_id)


@_utilities.lift_output_func(get_consent)
def get_consent_output(consent_id: Optional[pulumi.Input[str]] = None,
                       consent_store_id: Optional[pulumi.Input[str]] = None,
                       dataset_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConsentResult]:
    """
    Gets the specified revision of a Consent, or the latest revision if `revision_id` is not specified in the resource name.
    """
    ...
