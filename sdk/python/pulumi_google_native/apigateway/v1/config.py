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
from ._inputs import *

__all__ = ['ConfigArgs', 'Config']

@pulumi.input_type
class ConfigArgs:
    def __init__(__self__, *,
                 api_config_id: pulumi.Input[str],
                 api_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_service_account: Optional[pulumi.Input[str]] = None,
                 grpc_services: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_service_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigFileArgs']]]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigOpenApiDocumentArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Config resource.
        :param pulumi.Input[str] api_config_id: Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        :param pulumi.Input[str] display_name: Optional. Display name.
        :param pulumi.Input[str] gateway_service_account: Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        :param pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigGrpcServiceDefinitionArgs']]] grpc_services: Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigFileArgs']]] managed_service_configs: Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        :param pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigOpenApiDocumentArgs']]] openapi_documents: Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        """
        pulumi.set(__self__, "api_config_id", api_config_id)
        pulumi.set(__self__, "api_id", api_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if gateway_service_account is not None:
            pulumi.set(__self__, "gateway_service_account", gateway_service_account)
        if grpc_services is not None:
            pulumi.set(__self__, "grpc_services", grpc_services)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_service_configs is not None:
            pulumi.set(__self__, "managed_service_configs", managed_service_configs)
        if openapi_documents is not None:
            pulumi.set(__self__, "openapi_documents", openapi_documents)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="apiConfigId")
    def api_config_id(self) -> pulumi.Input[str]:
        """
        Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "api_config_id")

    @api_config_id.setter
    def api_config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_config_id", value)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="gatewayServiceAccount")
    def gateway_service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        """
        return pulumi.get(self, "gateway_service_account")

    @gateway_service_account.setter
    def gateway_service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_service_account", value)

    @property
    @pulumi.getter(name="grpcServices")
    def grpc_services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]]:
        """
        Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        """
        return pulumi.get(self, "grpc_services")

    @grpc_services.setter
    def grpc_services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]]):
        pulumi.set(self, "grpc_services", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managedServiceConfigs")
    def managed_service_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigFileArgs']]]]:
        """
        Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        """
        return pulumi.get(self, "managed_service_configs")

    @managed_service_configs.setter
    def managed_service_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigFileArgs']]]]):
        pulumi.set(self, "managed_service_configs", value)

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigOpenApiDocumentArgs']]]]:
        """
        Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        """
        return pulumi.get(self, "openapi_documents")

    @openapi_documents.setter
    def openapi_documents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApigatewayApiConfigOpenApiDocumentArgs']]]]):
        pulumi.set(self, "openapi_documents", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Config(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_service_account: Optional[pulumi.Input[str]] = None,
                 grpc_services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_service_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigFileArgs']]]]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigOpenApiDocumentArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new ApiConfig in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_config_id: Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        :param pulumi.Input[str] display_name: Optional. Display name.
        :param pulumi.Input[str] gateway_service_account: Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]] grpc_services: Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigFileArgs']]]] managed_service_configs: Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigOpenApiDocumentArgs']]]] openapi_documents: Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ApiConfig in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_service_account: Optional[pulumi.Input[str]] = None,
                 grpc_services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigGrpcServiceDefinitionArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_service_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigFileArgs']]]]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApigatewayApiConfigOpenApiDocumentArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigArgs.__new__(ConfigArgs)

            if api_config_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_config_id'")
            __props__.__dict__["api_config_id"] = api_config_id
            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["gateway_service_account"] = gateway_service_account
            __props__.__dict__["grpc_services"] = grpc_services
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["managed_service_configs"] = managed_service_configs
            __props__.__dict__["openapi_documents"] = openapi_documents
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["service_config_id"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["api_config_id", "api_id", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Config, __self__).__init__(
            'google-native:apigateway/v1:Config',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Config':
        """
        Get an existing Config resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConfigArgs.__new__(ConfigArgs)

        __props__.__dict__["api_config_id"] = None
        __props__.__dict__["api_id"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["gateway_service_account"] = None
        __props__.__dict__["grpc_services"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["managed_service_configs"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["openapi_documents"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["service_config_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return Config(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiConfigId")
    def api_config_id(self) -> pulumi.Output[str]:
        """
        Required. Identifier to assign to the API Config. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "api_config_id")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. Display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gatewayServiceAccount")
    def gateway_service_account(self) -> pulumi.Output[str]:
        """
        Immutable. The Google Cloud IAM Service Account that Gateways serving this config should use to authenticate to other services. This may either be the Service Account's email (`{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com`) or its full resource name (`projects/{PROJECT}/accounts/{UNIQUE_ID}`). This is most often used when the service is a GCP resource such as a Cloud Run Service or an IAP-secured service.
        """
        return pulumi.get(self, "gateway_service_account")

    @property
    @pulumi.getter(name="grpcServices")
    def grpc_services(self) -> pulumi.Output[Sequence['outputs.ApigatewayApiConfigGrpcServiceDefinitionResponse']]:
        """
        Optional. gRPC service definition files. If specified, openapi_documents must not be included.
        """
        return pulumi.get(self, "grpc_services")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedServiceConfigs")
    def managed_service_configs(self) -> pulumi.Output[Sequence['outputs.ApigatewayApiConfigFileResponse']]:
        """
        Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
        """
        return pulumi.get(self, "managed_service_configs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the API Config. Format: projects/{project}/locations/global/apis/{api}/configs/{api_config}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> pulumi.Output[Sequence['outputs.ApigatewayApiConfigOpenApiDocumentResponse']]:
        """
        Optional. OpenAPI specification documents. If specified, grpc_services and managed_service_configs must not be included.
        """
        return pulumi.get(self, "openapi_documents")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceConfigId")
    def service_config_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated Service Config ( https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        return pulumi.get(self, "service_config_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the API Config.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")

