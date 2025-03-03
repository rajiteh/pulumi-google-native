# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'AppEngineHttpTargetArgs',
    'AppEngineRoutingArgs',
    'HttpTargetArgs',
    'OAuthTokenArgs',
    'OidcTokenArgs',
    'PubsubTargetArgs',
    'RetryConfigArgs',
]

@pulumi.input_type
class AppEngineHttpTargetArgs:
    def __init__(__self__, *,
                 app_engine_routing: Optional[pulumi.Input['AppEngineRoutingArgs']] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 http_method: Optional[pulumi.Input['AppEngineHttpTargetHttpMethod']] = None,
                 relative_uri: Optional[pulumi.Input[str]] = None):
        """
        App Engine target. The job will be pushed to a job handler by means of an HTTP request via an http_method such as HTTP POST, HTTP GET, etc. The job is acknowledged by means of an HTTP response code in the range [200 - 299]. Error 503 is considered an App Engine system error instead of an application error. Requests returning error 503 will be retried regardless of retry configuration and not counted against retry counts. Any other response code, or a failure to receive a response before the deadline, constitutes a failed attempt.
        :param pulumi.Input['AppEngineRoutingArgs'] app_engine_routing: App Engine Routing setting for the job.
        :param pulumi.Input[str] body: Body. HTTP request body. A request body is allowed only if the HTTP method is POST or PUT. It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: HTTP request headers. This map contains the header field names and values. Headers can be set when the job is created. Cloud Scheduler sets some headers to default values: * `User-Agent`: By default, this header is `"AppEngine-Google; (+http://code.google.com/appengine)"`. This header can be modified, but Cloud Scheduler will append `"AppEngine-Google; (+http://code.google.com/appengine)"` to the modified `User-Agent`. * `X-CloudScheduler`: This header will be set to true. * `X-CloudScheduler-JobName`: This header will contain the job name. * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in the unix-cron format, this header will contain the job schedule time in RFC3339 UTC "Zulu" format. If the job has an body, Cloud Scheduler sets the following headers: * `Content-Type`: By default, the `Content-Type` header is set to `"application/octet-stream"`. The default can be overridden by explictly setting `Content-Type` to a particular media type when the job is created. For example, `Content-Type` can be set to `"application/json"`. * `Content-Length`: This is computed by Cloud Scheduler. This value is output only. It cannot be changed. The headers below are output only. They cannot be set or overridden: * `X-Google-*`: For Google internal use only. * `X-AppEngine-*`: For Google internal use only. In addition, some App Engine headers, which contain job-specific information, are also be sent to the job handler.
        :param pulumi.Input['AppEngineHttpTargetHttpMethod'] http_method: The HTTP method to use for the request. PATCH and OPTIONS are not permitted.
        :param pulumi.Input[str] relative_uri: The relative URI. The relative URL must begin with "/" and must be a valid HTTP relative URL. It can contain a path, query string arguments, and `#` fragments. If the relative URL is empty, then the root path "/" will be used. No spaces are allowed, and the maximum length allowed is 2083 characters.
        """
        if app_engine_routing is not None:
            pulumi.set(__self__, "app_engine_routing", app_engine_routing)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if relative_uri is not None:
            pulumi.set(__self__, "relative_uri", relative_uri)

    @property
    @pulumi.getter(name="appEngineRouting")
    def app_engine_routing(self) -> Optional[pulumi.Input['AppEngineRoutingArgs']]:
        """
        App Engine Routing setting for the job.
        """
        return pulumi.get(self, "app_engine_routing")

    @app_engine_routing.setter
    def app_engine_routing(self, value: Optional[pulumi.Input['AppEngineRoutingArgs']]):
        pulumi.set(self, "app_engine_routing", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        Body. HTTP request body. A request body is allowed only if the HTTP method is POST or PUT. It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        HTTP request headers. This map contains the header field names and values. Headers can be set when the job is created. Cloud Scheduler sets some headers to default values: * `User-Agent`: By default, this header is `"AppEngine-Google; (+http://code.google.com/appengine)"`. This header can be modified, but Cloud Scheduler will append `"AppEngine-Google; (+http://code.google.com/appengine)"` to the modified `User-Agent`. * `X-CloudScheduler`: This header will be set to true. * `X-CloudScheduler-JobName`: This header will contain the job name. * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in the unix-cron format, this header will contain the job schedule time in RFC3339 UTC "Zulu" format. If the job has an body, Cloud Scheduler sets the following headers: * `Content-Type`: By default, the `Content-Type` header is set to `"application/octet-stream"`. The default can be overridden by explictly setting `Content-Type` to a particular media type when the job is created. For example, `Content-Type` can be set to `"application/json"`. * `Content-Length`: This is computed by Cloud Scheduler. This value is output only. It cannot be changed. The headers below are output only. They cannot be set or overridden: * `X-Google-*`: For Google internal use only. * `X-AppEngine-*`: For Google internal use only. In addition, some App Engine headers, which contain job-specific information, are also be sent to the job handler.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input['AppEngineHttpTargetHttpMethod']]:
        """
        The HTTP method to use for the request. PATCH and OPTIONS are not permitted.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input['AppEngineHttpTargetHttpMethod']]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="relativeUri")
    def relative_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The relative URI. The relative URL must begin with "/" and must be a valid HTTP relative URL. It can contain a path, query string arguments, and `#` fragments. If the relative URL is empty, then the root path "/" will be used. No spaces are allowed, and the maximum length allowed is 2083 characters.
        """
        return pulumi.get(self, "relative_uri")

    @relative_uri.setter
    def relative_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "relative_uri", value)


@pulumi.input_type
class AppEngineRoutingArgs:
    def __init__(__self__, *,
                 instance: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        App Engine Routing. For more information about services, versions, and instances see [An Overview of App Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine), [Microservices Architecture on Google App Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine), [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed), and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
        :param pulumi.Input[str] instance: App instance. By default, the job is sent to an instance which is available when the job is attempted. Requests can only be sent to a specific instance if [manual scaling is used in App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?#scaling_types_and_instance_classes). App Engine Flex does not support instances. For more information, see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed) and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
        :param pulumi.Input[str] service: App service. By default, the job is sent to the service which is the default service when the job is attempted.
        :param pulumi.Input[str] version: App version. By default, the job is sent to the version which is the default version when the job is attempted.
        """
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        App instance. By default, the job is sent to an instance which is available when the job is attempted. Requests can only be sent to a specific instance if [manual scaling is used in App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?#scaling_types_and_instance_classes). App Engine Flex does not support instances. For more information, see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed) and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        App service. By default, the job is sent to the service which is the default service when the job is attempted.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        App version. By default, the job is sent to the version which is the default version when the job is attempted.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class HttpTargetArgs:
    def __init__(__self__, *,
                 uri: pulumi.Input[str],
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 http_method: Optional[pulumi.Input['HttpTargetHttpMethod']] = None,
                 oauth_token: Optional[pulumi.Input['OAuthTokenArgs']] = None,
                 oidc_token: Optional[pulumi.Input['OidcTokenArgs']] = None):
        """
        Http target. The job will be pushed to the job handler by means of an HTTP request via an http_method such as HTTP POST, HTTP GET, etc. The job is acknowledged by means of an HTTP response code in the range [200 - 299]. A failure to receive a response constitutes a failed execution. For a redirected request, the response returned by the redirected request is considered.
        :param pulumi.Input[str] uri: The full URI path that the request will be sent to. This string must begin with either "http://" or "https://". Some examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding.
        :param pulumi.Input[str] body: HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: The user can specify HTTP request headers to send with the job's HTTP request. This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas. These headers represent a subset of the headers that will accompany the job's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is below: - Host: This will be computed by Cloud Scheduler and derived from uri. * `Content-Length`: This will be computed by Cloud Scheduler. * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`. * `X-Google-*`: Google internal use only. * `X-AppEngine-*`: Google internal use only. * `X-CloudScheduler`: This header will be set to true. * `X-CloudScheduler-JobName`: This header will contain the job name. * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in the unix-cron format, this header will contain the job schedule time in RFC3339 UTC "Zulu" format. The total size of headers must be less than 80KB.
        :param pulumi.Input['HttpTargetHttpMethod'] http_method: Which HTTP method to use for the request.
        :param pulumi.Input['OAuthTokenArgs'] oauth_token: If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        :param pulumi.Input['OidcTokenArgs'] oidc_token: If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        """
        pulumi.set(__self__, "uri", uri)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if oauth_token is not None:
            pulumi.set(__self__, "oauth_token", oauth_token)
        if oidc_token is not None:
            pulumi.set(__self__, "oidc_token", oidc_token)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        The full URI path that the request will be sent to. This string must begin with either "http://" or "https://". Some examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The user can specify HTTP request headers to send with the job's HTTP request. This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas. These headers represent a subset of the headers that will accompany the job's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is below: - Host: This will be computed by Cloud Scheduler and derived from uri. * `Content-Length`: This will be computed by Cloud Scheduler. * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`. * `X-Google-*`: Google internal use only. * `X-AppEngine-*`: Google internal use only. * `X-CloudScheduler`: This header will be set to true. * `X-CloudScheduler-JobName`: This header will contain the job name. * `X-CloudScheduler-ScheduleTime`: For Cloud Scheduler jobs specified in the unix-cron format, this header will contain the job schedule time in RFC3339 UTC "Zulu" format. The total size of headers must be less than 80KB.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input['HttpTargetHttpMethod']]:
        """
        Which HTTP method to use for the request.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input['HttpTargetHttpMethod']]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="oauthToken")
    def oauth_token(self) -> Optional[pulumi.Input['OAuthTokenArgs']]:
        """
        If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        """
        return pulumi.get(self, "oauth_token")

    @oauth_token.setter
    def oauth_token(self, value: Optional[pulumi.Input['OAuthTokenArgs']]):
        pulumi.set(self, "oauth_token", value)

    @property
    @pulumi.getter(name="oidcToken")
    def oidc_token(self) -> Optional[pulumi.Input['OidcTokenArgs']]:
        """
        If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        """
        return pulumi.get(self, "oidc_token")

    @oidc_token.setter
    def oidc_token(self, value: Optional[pulumi.Input['OidcTokenArgs']]):
        pulumi.set(self, "oidc_token", value)


@pulumi.input_type
class OAuthTokenArgs:
    def __init__(__self__, *,
                 scope: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None):
        """
        Contains information needed for generating an [OAuth token](https://developers.google.com/identity/protocols/OAuth2). This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        :param pulumi.Input[str] scope: OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used.
        :param pulumi.Input[str] service_account_email: [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OAuth token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.
        """
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OAuth token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)


@pulumi.input_type
class OidcTokenArgs:
    def __init__(__self__, *,
                 audience: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None):
        """
        Contains information needed for generating an [OpenID Connect token](https://developers.google.com/identity/protocols/OpenIDConnect). This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        :param pulumi.Input[str] audience: Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.
        :param pulumi.Input[str] service_account_email: [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OIDC token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.
        """
        if audience is not None:
            pulumi.set(__self__, "audience", audience)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)

    @property
    @pulumi.getter
    def audience(self) -> Optional[pulumi.Input[str]]:
        """
        Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audience", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OIDC token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)


@pulumi.input_type
class PubsubTargetArgs:
    def __init__(__self__, *,
                 topic_name: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 data: Optional[pulumi.Input[str]] = None):
        """
        Pub/Sub target. The job will be delivered by publishing a message to the given Pub/Sub topic.
        :param pulumi.Input[str] topic_name: The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        :param pulumi.Input[str] data: The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        """
        pulumi.set(__self__, "topic_name", topic_name)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class RetryConfigArgs:
    def __init__(__self__, *,
                 max_backoff_duration: Optional[pulumi.Input[str]] = None,
                 max_doublings: Optional[pulumi.Input[int]] = None,
                 max_retry_duration: Optional[pulumi.Input[str]] = None,
                 min_backoff_duration: Optional[pulumi.Input[str]] = None,
                 retry_count: Optional[pulumi.Input[int]] = None):
        """
        Settings that determine the retry behavior. By default, if a job does not complete successfully (meaning that an acknowledgement is not received from the handler, then it will be retried with exponential backoff according to the settings in RetryConfig.
        :param pulumi.Input[str] max_backoff_duration: The maximum amount of time to wait before retrying a job after it fails. The default value of this field is 1 hour.
        :param pulumi.Input[int] max_doublings: The time between retries will double `max_doublings` times. A job's retry interval starts at min_backoff_duration, then doubles `max_doublings` times, then increases linearly, and finally retries at intervals of max_backoff_duration up to retry_count times. For example, if min_backoff_duration is 10s, max_backoff_duration is 300s, and `max_doublings` is 3, then the job will first be retried in 10s. The retry interval will double three times, and then increase linearly by 2^3 * 10s. Finally, the job will retry at intervals of max_backoff_duration until the job has been attempted retry_count times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, .... The default value of this field is 5.
        :param pulumi.Input[str] max_retry_duration: The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retry_count, the job will be retried until both limits are reached. The default value for max_retry_duration is zero, which means retry duration is unlimited.
        :param pulumi.Input[str] min_backoff_duration: The minimum amount of time to wait before retrying a job after it fails. The default value of this field is 5 seconds.
        :param pulumi.Input[int] retry_count: The number of attempts that the system will make to run a job using the exponential backoff procedure described by max_doublings. The default value of retry_count is zero. If retry_count is zero, a job attempt will *not* be retried if it fails. Instead the Cloud Scheduler system will wait for the next scheduled execution time. If retry_count is set to a non-zero number then Cloud Scheduler will retry failed attempts, using exponential backoff, retry_count times, or until the next scheduled execution time, whichever comes first. Values greater than 5 and negative values are not allowed.
        """
        if max_backoff_duration is not None:
            pulumi.set(__self__, "max_backoff_duration", max_backoff_duration)
        if max_doublings is not None:
            pulumi.set(__self__, "max_doublings", max_doublings)
        if max_retry_duration is not None:
            pulumi.set(__self__, "max_retry_duration", max_retry_duration)
        if min_backoff_duration is not None:
            pulumi.set(__self__, "min_backoff_duration", min_backoff_duration)
        if retry_count is not None:
            pulumi.set(__self__, "retry_count", retry_count)

    @property
    @pulumi.getter(name="maxBackoffDuration")
    def max_backoff_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum amount of time to wait before retrying a job after it fails. The default value of this field is 1 hour.
        """
        return pulumi.get(self, "max_backoff_duration")

    @max_backoff_duration.setter
    def max_backoff_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_backoff_duration", value)

    @property
    @pulumi.getter(name="maxDoublings")
    def max_doublings(self) -> Optional[pulumi.Input[int]]:
        """
        The time between retries will double `max_doublings` times. A job's retry interval starts at min_backoff_duration, then doubles `max_doublings` times, then increases linearly, and finally retries at intervals of max_backoff_duration up to retry_count times. For example, if min_backoff_duration is 10s, max_backoff_duration is 300s, and `max_doublings` is 3, then the job will first be retried in 10s. The retry interval will double three times, and then increase linearly by 2^3 * 10s. Finally, the job will retry at intervals of max_backoff_duration until the job has been attempted retry_count times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, .... The default value of this field is 5.
        """
        return pulumi.get(self, "max_doublings")

    @max_doublings.setter
    def max_doublings(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_doublings", value)

    @property
    @pulumi.getter(name="maxRetryDuration")
    def max_retry_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retry_count, the job will be retried until both limits are reached. The default value for max_retry_duration is zero, which means retry duration is unlimited.
        """
        return pulumi.get(self, "max_retry_duration")

    @max_retry_duration.setter
    def max_retry_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_retry_duration", value)

    @property
    @pulumi.getter(name="minBackoffDuration")
    def min_backoff_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum amount of time to wait before retrying a job after it fails. The default value of this field is 5 seconds.
        """
        return pulumi.get(self, "min_backoff_duration")

    @min_backoff_duration.setter
    def min_backoff_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_backoff_duration", value)

    @property
    @pulumi.getter(name="retryCount")
    def retry_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of attempts that the system will make to run a job using the exponential backoff procedure described by max_doublings. The default value of retry_count is zero. If retry_count is zero, a job attempt will *not* be retried if it fails. Instead the Cloud Scheduler system will wait for the next scheduled execution time. If retry_count is set to a non-zero number then Cloud Scheduler will retry failed attempts, using exponential backoff, retry_count times, or until the next scheduled execution time, whichever comes first. Values greater than 5 and negative values are not allowed.
        """
        return pulumi.get(self, "retry_count")

    @retry_count.setter
    def retry_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_count", value)


