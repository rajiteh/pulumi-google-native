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
    'GetConfigResult',
    'AwaitableGetConfigResult',
    'get_config',
    'get_config_output',
]

@pulumi.output_type
class GetConfigResult:
    def __init__(__self__, create_time=None, display_name=None, gateway_config=None, gateway_service_account=None, grpc_services=None, labels=None, managed_service_configs=None, name=None, openapi_documents=None, service_config_id=None, state=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if gateway_config and not isinstance(gateway_config, dict):
            raise TypeError("Expected argument 'gateway_config' to be a dict")
        pulumi.set(__self__, "gateway_config", gateway_config)
        if gateway_service_account and not isinstance(gateway_service_account, str):
            raise TypeError("Expected argument 'gateway_service_account' to be a str")
        pulumi.set(__self__, "gateway_service_account", gateway_service_account)
        if grpc_services and not isinstance(grpc_services, list):
            raise TypeError("Expected argument 'grpc_services' to be a list")
        pulumi.set(__self__, "grpc_services", grpc_services)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if managed_service_configs and not isinstance(managed_service_configs, list):
            raise TypeError("Expected argument 'managed_service_configs' to be a list")
        pulumi.set(__self__, "managed_service_configs", managed_service_configs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if openapi_documents and not isinstance(openapi_documents, list):
            raise TypeError("Expected argument 'openapi_documents' to be a list")
        pulumi.set(__self__, "openapi_documents", openapi_documents)
        if service_config_id and not isinstance(service_config_id, str):
            raise TypeError("Expected argument 'service_config_id' to be a str")
        pulumi.set(__self__, "service_config_id", service_config_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. Display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> 'outputs.ApigatewayGatewayConfigResponse':
        """
        Immutable. Gateway specific configuration.
        """
        return pulumi.get(self, "gateway_config")

    @property
    @pulumi.getter(name="gatewayServiceAccount")
    def gateway_service_account(self) -> str:
        """
        Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        """
        return pulumi.get(self, "gateway_service_account")

    @property
    @pulumi.getter(name="grpcServices")
    def grpc_services(self) -> Sequence['outputs.ApigatewayApiConfigGrpcServiceDefinitionResponse']:
        """
        Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        """
        return pulumi.get(self, "grpc_services")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="managedServiceConfigs")
    def managed_service_configs(self) -> Sequence['outputs.ApigatewayApiConfigFileResponse']:
        """
        Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        """
        return pulumi.get(self, "managed_service_configs")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the API Config. Format: projects/{project}/locations/global/apis/{api}/configs/{api_config}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> Sequence['outputs.ApigatewayApiConfigOpenApiDocumentResponse']:
        """
        Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        """
        return pulumi.get(self, "openapi_documents")

    @property
    @pulumi.getter(name="serviceConfigId")
    def service_config_id(self) -> str:
        """
        The ID of the associated Service Config ( https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        return pulumi.get(self, "service_config_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the API Config.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetConfigResult(GetConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigResult(
            create_time=self.create_time,
            display_name=self.display_name,
            gateway_config=self.gateway_config,
            gateway_service_account=self.gateway_service_account,
            grpc_services=self.grpc_services,
            labels=self.labels,
            managed_service_configs=self.managed_service_configs,
            name=self.name,
            openapi_documents=self.openapi_documents,
            service_config_id=self.service_config_id,
            state=self.state,
            update_time=self.update_time)


def get_config(api_id: Optional[str] = None,
               config_id: Optional[str] = None,
               location: Optional[str] = None,
               project: Optional[str] = None,
               view: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigResult:
    """
    Gets details of a single ApiConfig.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['configId'] = config_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigateway/v1beta:getConfig', __args__, opts=opts, typ=GetConfigResult).value

    return AwaitableGetConfigResult(
        create_time=__ret__.create_time,
        display_name=__ret__.display_name,
        gateway_config=__ret__.gateway_config,
        gateway_service_account=__ret__.gateway_service_account,
        grpc_services=__ret__.grpc_services,
        labels=__ret__.labels,
        managed_service_configs=__ret__.managed_service_configs,
        name=__ret__.name,
        openapi_documents=__ret__.openapi_documents,
        service_config_id=__ret__.service_config_id,
        state=__ret__.state,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_config)
def get_config_output(api_id: Optional[pulumi.Input[str]] = None,
                      config_id: Optional[pulumi.Input[str]] = None,
                      location: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      view: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigResult]:
    """
    Gets details of a single ApiConfig.
    """
    ...
