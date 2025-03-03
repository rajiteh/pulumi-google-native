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
    'GetBuildResult',
    'AwaitableGetBuildResult',
    'get_build',
    'get_build_output',
]

@pulumi.output_type
class GetBuildResult:
    def __init__(__self__, approval=None, artifacts=None, available_secrets=None, build_trigger_id=None, create_time=None, failure_info=None, finish_time=None, images=None, log_url=None, logs_bucket=None, name=None, options=None, project=None, queue_ttl=None, results=None, secrets=None, service_account=None, source=None, source_provenance=None, start_time=None, status=None, status_detail=None, steps=None, substitutions=None, tags=None, timeout=None, timing=None, warnings=None):
        if approval and not isinstance(approval, dict):
            raise TypeError("Expected argument 'approval' to be a dict")
        pulumi.set(__self__, "approval", approval)
        if artifacts and not isinstance(artifacts, dict):
            raise TypeError("Expected argument 'artifacts' to be a dict")
        pulumi.set(__self__, "artifacts", artifacts)
        if available_secrets and not isinstance(available_secrets, dict):
            raise TypeError("Expected argument 'available_secrets' to be a dict")
        pulumi.set(__self__, "available_secrets", available_secrets)
        if build_trigger_id and not isinstance(build_trigger_id, str):
            raise TypeError("Expected argument 'build_trigger_id' to be a str")
        pulumi.set(__self__, "build_trigger_id", build_trigger_id)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if failure_info and not isinstance(failure_info, dict):
            raise TypeError("Expected argument 'failure_info' to be a dict")
        pulumi.set(__self__, "failure_info", failure_info)
        if finish_time and not isinstance(finish_time, str):
            raise TypeError("Expected argument 'finish_time' to be a str")
        pulumi.set(__self__, "finish_time", finish_time)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if log_url and not isinstance(log_url, str):
            raise TypeError("Expected argument 'log_url' to be a str")
        pulumi.set(__self__, "log_url", log_url)
        if logs_bucket and not isinstance(logs_bucket, str):
            raise TypeError("Expected argument 'logs_bucket' to be a str")
        pulumi.set(__self__, "logs_bucket", logs_bucket)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if queue_ttl and not isinstance(queue_ttl, str):
            raise TypeError("Expected argument 'queue_ttl' to be a str")
        pulumi.set(__self__, "queue_ttl", queue_ttl)
        if results and not isinstance(results, dict):
            raise TypeError("Expected argument 'results' to be a dict")
        pulumi.set(__self__, "results", results)
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        pulumi.set(__self__, "secrets", secrets)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)
        if source_provenance and not isinstance(source_provenance, dict):
            raise TypeError("Expected argument 'source_provenance' to be a dict")
        pulumi.set(__self__, "source_provenance", source_provenance)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_detail and not isinstance(status_detail, str):
            raise TypeError("Expected argument 'status_detail' to be a str")
        pulumi.set(__self__, "status_detail", status_detail)
        if steps and not isinstance(steps, list):
            raise TypeError("Expected argument 'steps' to be a list")
        pulumi.set(__self__, "steps", steps)
        if substitutions and not isinstance(substitutions, dict):
            raise TypeError("Expected argument 'substitutions' to be a dict")
        pulumi.set(__self__, "substitutions", substitutions)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if timeout and not isinstance(timeout, str):
            raise TypeError("Expected argument 'timeout' to be a str")
        pulumi.set(__self__, "timeout", timeout)
        if timing and not isinstance(timing, dict):
            raise TypeError("Expected argument 'timing' to be a dict")
        pulumi.set(__self__, "timing", timing)
        if warnings and not isinstance(warnings, list):
            raise TypeError("Expected argument 'warnings' to be a list")
        pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter
    def approval(self) -> 'outputs.BuildApprovalResponse':
        """
        Describes this build's approval configuration, status, and result.
        """
        return pulumi.get(self, "approval")

    @property
    @pulumi.getter
    def artifacts(self) -> 'outputs.ArtifactsResponse':
        """
        Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
        """
        return pulumi.get(self, "artifacts")

    @property
    @pulumi.getter(name="availableSecrets")
    def available_secrets(self) -> 'outputs.SecretsResponse':
        """
        Secrets and secret environment variables.
        """
        return pulumi.get(self, "available_secrets")

    @property
    @pulumi.getter(name="buildTriggerId")
    def build_trigger_id(self) -> str:
        """
        The ID of the `BuildTrigger` that triggered this build, if it was triggered automatically.
        """
        return pulumi.get(self, "build_trigger_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time at which the request to create the build was received.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="failureInfo")
    def failure_info(self) -> 'outputs.FailureInfoResponse':
        """
        Contains information about the build when status=FAILURE.
        """
        return pulumi.get(self, "failure_info")

    @property
    @pulumi.getter(name="finishTime")
    def finish_time(self) -> str:
        """
        Time at which execution of the build was finished. The difference between finish_time and start_time is the duration of the build's execution.
        """
        return pulumi.get(self, "finish_time")

    @property
    @pulumi.getter
    def images(self) -> Sequence[str]:
        """
        A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the `Build` resource's results field. If any of the images fail to be pushed, the build status is marked `FAILURE`.
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="logUrl")
    def log_url(self) -> str:
        """
        URL to logs for this build in Google Cloud Console.
        """
        return pulumi.get(self, "log_url")

    @property
    @pulumi.getter(name="logsBucket")
    def logs_bucket(self) -> str:
        """
        Google Cloud Storage bucket where logs should be written (see [Bucket Name Requirements](https://cloud.google.com/storage/docs/bucket-naming#requirements)). Logs file names will be of the format `${logs_bucket}/log-${build_id}.txt`.
        """
        return pulumi.get(self, "logs_bucket")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The 'Build' name with format: `projects/{project}/locations/{location}/builds/{build}`, where {build} is a unique identifier generated by the service.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> 'outputs.BuildOptionsResponse':
        """
        Special options for this build.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        ID of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="queueTtl")
    def queue_ttl(self) -> str:
        """
        TTL in queue for this build. If provided and the build is enqueued longer than this value, the build will expire and the build status will be `EXPIRED`. The TTL starts ticking from create_time.
        """
        return pulumi.get(self, "queue_ttl")

    @property
    @pulumi.getter
    def results(self) -> 'outputs.ResultsResponse':
        """
        Results of the build.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter
    def secrets(self) -> Sequence['outputs.SecretResponse']:
        """
        Secrets to decrypt using Cloud Key Management Service. Note: Secret Manager is the recommended technique for managing sensitive data with Cloud Build. Use `available_secrets` to configure builds to access secrets from Secret Manager. For instructions, see: https://cloud.google.com/cloud-build/docs/securing-builds/use-secrets
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        IAM service account whose credentials will be used at build runtime. Must be of the format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. ACCOUNT can be email address or uniqueId of the service account. 
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def source(self) -> 'outputs.SourceResponse':
        """
        The location of the source files to build.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceProvenance")
    def source_provenance(self) -> 'outputs.SourceProvenanceResponse':
        """
        A permanent fixed identifier for source.
        """
        return pulumi.get(self, "source_provenance")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Time at which execution of the build was started.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the build.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDetail")
    def status_detail(self) -> str:
        """
        Customer-readable message about the current status.
        """
        return pulumi.get(self, "status_detail")

    @property
    @pulumi.getter
    def steps(self) -> Sequence['outputs.BuildStepResponse']:
        """
        The operations to be performed on the workspace.
        """
        return pulumi.get(self, "steps")

    @property
    @pulumi.getter
    def substitutions(self) -> Mapping[str, str]:
        """
        Substitutions data for `Build` resource.
        """
        return pulumi.get(self, "substitutions")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags for annotation of a `Build`. These are not docker tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> str:
        """
        Amount of time that this build should be allowed to run, to second granularity. If this amount of time elapses, work on the build will cease and the build status will be `TIMEOUT`. `timeout` starts ticking from `startTime`. Default time is 60 minutes.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def timing(self) -> Mapping[str, str]:
        """
        Stores timing information for phases of the build. Valid keys are: * BUILD: time to execute all build steps. * PUSH: time to push all artifacts including docker images and non docker artifacts. * FETCHSOURCE: time to fetch source. * SETUPBUILD: time to set up build. If the build does not specify source or images, these keys will not be included.
        """
        return pulumi.get(self, "timing")

    @property
    @pulumi.getter
    def warnings(self) -> Sequence['outputs.WarningResponse']:
        """
        Non-fatal problems encountered during the execution of the build.
        """
        return pulumi.get(self, "warnings")


class AwaitableGetBuildResult(GetBuildResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBuildResult(
            approval=self.approval,
            artifacts=self.artifacts,
            available_secrets=self.available_secrets,
            build_trigger_id=self.build_trigger_id,
            create_time=self.create_time,
            failure_info=self.failure_info,
            finish_time=self.finish_time,
            images=self.images,
            log_url=self.log_url,
            logs_bucket=self.logs_bucket,
            name=self.name,
            options=self.options,
            project=self.project,
            queue_ttl=self.queue_ttl,
            results=self.results,
            secrets=self.secrets,
            service_account=self.service_account,
            source=self.source,
            source_provenance=self.source_provenance,
            start_time=self.start_time,
            status=self.status,
            status_detail=self.status_detail,
            steps=self.steps,
            substitutions=self.substitutions,
            tags=self.tags,
            timeout=self.timeout,
            timing=self.timing,
            warnings=self.warnings)


def get_build(build_id: Optional[str] = None,
              id: Optional[str] = None,
              location: Optional[str] = None,
              project: Optional[str] = None,
              project_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBuildResult:
    """
    Returns information about a previously requested build. The `Build` that is returned includes its status (such as `SUCCESS`, `FAILURE`, or `WORKING`), and timing information.
    """
    __args__ = dict()
    __args__['buildId'] = build_id
    __args__['id'] = id
    __args__['location'] = location
    __args__['project'] = project
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudbuild/v1:getBuild', __args__, opts=opts, typ=GetBuildResult).value

    return AwaitableGetBuildResult(
        approval=__ret__.approval,
        artifacts=__ret__.artifacts,
        available_secrets=__ret__.available_secrets,
        build_trigger_id=__ret__.build_trigger_id,
        create_time=__ret__.create_time,
        failure_info=__ret__.failure_info,
        finish_time=__ret__.finish_time,
        images=__ret__.images,
        log_url=__ret__.log_url,
        logs_bucket=__ret__.logs_bucket,
        name=__ret__.name,
        options=__ret__.options,
        project=__ret__.project,
        queue_ttl=__ret__.queue_ttl,
        results=__ret__.results,
        secrets=__ret__.secrets,
        service_account=__ret__.service_account,
        source=__ret__.source,
        source_provenance=__ret__.source_provenance,
        start_time=__ret__.start_time,
        status=__ret__.status,
        status_detail=__ret__.status_detail,
        steps=__ret__.steps,
        substitutions=__ret__.substitutions,
        tags=__ret__.tags,
        timeout=__ret__.timeout,
        timing=__ret__.timing,
        warnings=__ret__.warnings)


@_utilities.lift_output_func(get_build)
def get_build_output(build_id: Optional[pulumi.Input[str]] = None,
                     id: Optional[pulumi.Input[str]] = None,
                     location: Optional[pulumi.Input[str]] = None,
                     project: Optional[pulumi.Input[Optional[str]]] = None,
                     project_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBuildResult]:
    """
    Returns information about a previously requested build. The `Build` that is returned includes its status (such as `SUCCESS`, `FAILURE`, or `WORKING`), and timing information.
    """
    ...
