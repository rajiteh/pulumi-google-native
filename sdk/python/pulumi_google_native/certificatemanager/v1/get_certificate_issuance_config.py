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
    'GetCertificateIssuanceConfigResult',
    'AwaitableGetCertificateIssuanceConfigResult',
    'get_certificate_issuance_config',
    'get_certificate_issuance_config_output',
]

@pulumi.output_type
class GetCertificateIssuanceConfigResult:
    def __init__(__self__, certificate_authority_config=None, create_time=None, description=None, key_algorithm=None, labels=None, lifetime=None, name=None, rotation_window_percentage=None, update_time=None):
        if certificate_authority_config and not isinstance(certificate_authority_config, dict):
            raise TypeError("Expected argument 'certificate_authority_config' to be a dict")
        pulumi.set(__self__, "certificate_authority_config", certificate_authority_config)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if key_algorithm and not isinstance(key_algorithm, str):
            raise TypeError("Expected argument 'key_algorithm' to be a str")
        pulumi.set(__self__, "key_algorithm", key_algorithm)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if lifetime and not isinstance(lifetime, str):
            raise TypeError("Expected argument 'lifetime' to be a str")
        pulumi.set(__self__, "lifetime", lifetime)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rotation_window_percentage and not isinstance(rotation_window_percentage, int):
            raise TypeError("Expected argument 'rotation_window_percentage' to be a int")
        pulumi.set(__self__, "rotation_window_percentage", rotation_window_percentage)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="certificateAuthorityConfig")
    def certificate_authority_config(self) -> 'outputs.CertificateAuthorityConfigResponse':
        """
        The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
        """
        return pulumi.get(self, "certificate_authority_config")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of a CertificateIssuanceConfig.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        One or more paragraphs of text description of a CertificateIssuanceConfig.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> str:
        """
        The key algorithm to use when generating the private key.
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Set of labels associated with a CertificateIssuanceConfig.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def lifetime(self) -> str:
        """
        Workload certificate lifetime requested.
        """
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rotationWindowPercentage")
    def rotation_window_percentage(self) -> int:
        """
        Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
        """
        return pulumi.get(self, "rotation_window_percentage")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of a CertificateIssuanceConfig.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetCertificateIssuanceConfigResult(GetCertificateIssuanceConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateIssuanceConfigResult(
            certificate_authority_config=self.certificate_authority_config,
            create_time=self.create_time,
            description=self.description,
            key_algorithm=self.key_algorithm,
            labels=self.labels,
            lifetime=self.lifetime,
            name=self.name,
            rotation_window_percentage=self.rotation_window_percentage,
            update_time=self.update_time)


def get_certificate_issuance_config(certificate_issuance_config_id: Optional[str] = None,
                                    location: Optional[str] = None,
                                    project: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateIssuanceConfigResult:
    """
    Gets details of a single CertificateIssuanceConfig.
    """
    __args__ = dict()
    __args__['certificateIssuanceConfigId'] = certificate_issuance_config_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:certificatemanager/v1:getCertificateIssuanceConfig', __args__, opts=opts, typ=GetCertificateIssuanceConfigResult).value

    return AwaitableGetCertificateIssuanceConfigResult(
        certificate_authority_config=__ret__.certificate_authority_config,
        create_time=__ret__.create_time,
        description=__ret__.description,
        key_algorithm=__ret__.key_algorithm,
        labels=__ret__.labels,
        lifetime=__ret__.lifetime,
        name=__ret__.name,
        rotation_window_percentage=__ret__.rotation_window_percentage,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_certificate_issuance_config)
def get_certificate_issuance_config_output(certificate_issuance_config_id: Optional[pulumi.Input[str]] = None,
                                           location: Optional[pulumi.Input[str]] = None,
                                           project: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCertificateIssuanceConfigResult]:
    """
    Gets details of a single CertificateIssuanceConfig.
    """
    ...