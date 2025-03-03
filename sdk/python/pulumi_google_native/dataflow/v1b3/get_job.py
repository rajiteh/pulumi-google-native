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
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    def __init__(__self__, client_request_id=None, create_time=None, created_from_snapshot_id=None, current_state=None, current_state_time=None, environment=None, execution_info=None, job_metadata=None, labels=None, location=None, name=None, pipeline_description=None, project=None, replace_job_id=None, replaced_by_job_id=None, requested_state=None, runtime_updatable_params=None, satisfies_pzs=None, stage_states=None, start_time=None, steps=None, steps_location=None, temp_files=None, transform_name_mapping=None, type=None):
        if client_request_id and not isinstance(client_request_id, str):
            raise TypeError("Expected argument 'client_request_id' to be a str")
        pulumi.set(__self__, "client_request_id", client_request_id)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if created_from_snapshot_id and not isinstance(created_from_snapshot_id, str):
            raise TypeError("Expected argument 'created_from_snapshot_id' to be a str")
        pulumi.set(__self__, "created_from_snapshot_id", created_from_snapshot_id)
        if current_state and not isinstance(current_state, str):
            raise TypeError("Expected argument 'current_state' to be a str")
        pulumi.set(__self__, "current_state", current_state)
        if current_state_time and not isinstance(current_state_time, str):
            raise TypeError("Expected argument 'current_state_time' to be a str")
        pulumi.set(__self__, "current_state_time", current_state_time)
        if environment and not isinstance(environment, dict):
            raise TypeError("Expected argument 'environment' to be a dict")
        pulumi.set(__self__, "environment", environment)
        if execution_info and not isinstance(execution_info, dict):
            raise TypeError("Expected argument 'execution_info' to be a dict")
        if execution_info is not None:
            warnings.warn("""Deprecated.""", DeprecationWarning)
            pulumi.log.warn("""execution_info is deprecated: Deprecated.""")

        pulumi.set(__self__, "execution_info", execution_info)
        if job_metadata and not isinstance(job_metadata, dict):
            raise TypeError("Expected argument 'job_metadata' to be a dict")
        pulumi.set(__self__, "job_metadata", job_metadata)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pipeline_description and not isinstance(pipeline_description, dict):
            raise TypeError("Expected argument 'pipeline_description' to be a dict")
        pulumi.set(__self__, "pipeline_description", pipeline_description)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if replace_job_id and not isinstance(replace_job_id, str):
            raise TypeError("Expected argument 'replace_job_id' to be a str")
        pulumi.set(__self__, "replace_job_id", replace_job_id)
        if replaced_by_job_id and not isinstance(replaced_by_job_id, str):
            raise TypeError("Expected argument 'replaced_by_job_id' to be a str")
        pulumi.set(__self__, "replaced_by_job_id", replaced_by_job_id)
        if requested_state and not isinstance(requested_state, str):
            raise TypeError("Expected argument 'requested_state' to be a str")
        pulumi.set(__self__, "requested_state", requested_state)
        if runtime_updatable_params and not isinstance(runtime_updatable_params, dict):
            raise TypeError("Expected argument 'runtime_updatable_params' to be a dict")
        pulumi.set(__self__, "runtime_updatable_params", runtime_updatable_params)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if stage_states and not isinstance(stage_states, list):
            raise TypeError("Expected argument 'stage_states' to be a list")
        pulumi.set(__self__, "stage_states", stage_states)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if steps and not isinstance(steps, list):
            raise TypeError("Expected argument 'steps' to be a list")
        pulumi.set(__self__, "steps", steps)
        if steps_location and not isinstance(steps_location, str):
            raise TypeError("Expected argument 'steps_location' to be a str")
        pulumi.set(__self__, "steps_location", steps_location)
        if temp_files and not isinstance(temp_files, list):
            raise TypeError("Expected argument 'temp_files' to be a list")
        pulumi.set(__self__, "temp_files", temp_files)
        if transform_name_mapping and not isinstance(transform_name_mapping, dict):
            raise TypeError("Expected argument 'transform_name_mapping' to be a dict")
        pulumi.set(__self__, "transform_name_mapping", transform_name_mapping)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clientRequestId")
    def client_request_id(self) -> str:
        """
        The client's unique identifier of the job, re-used across retried attempts. If this field is set, the service will ensure its uniqueness. The request to create a job will fail if the service has knowledge of a previously submitted job with the same client's ID and job name. The caller may use this field to ensure idempotence of job creation across retried attempts to create a job. By default, the field is empty and, in that case, the service ignores it.
        """
        return pulumi.get(self, "client_request_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the job was initially created. Immutable and set by the Cloud Dataflow service.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="createdFromSnapshotId")
    def created_from_snapshot_id(self) -> str:
        """
        If this is specified, the job's initial state is populated from the given snapshot.
        """
        return pulumi.get(self, "created_from_snapshot_id")

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> str:
        """
        The current state of the job. Jobs are created in the `JOB_STATE_STOPPED` state unless otherwise specified. A job in the `JOB_STATE_RUNNING` state may asynchronously enter a terminal state. After a job has reached a terminal state, no further state updates may be made. This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
        """
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter(name="currentStateTime")
    def current_state_time(self) -> str:
        """
        The timestamp associated with the current state.
        """
        return pulumi.get(self, "current_state_time")

    @property
    @pulumi.getter
    def environment(self) -> 'outputs.EnvironmentResponse':
        """
        The environment for the job.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="executionInfo")
    def execution_info(self) -> 'outputs.JobExecutionInfoResponse':
        """
        Deprecated.
        """
        return pulumi.get(self, "execution_info")

    @property
    @pulumi.getter(name="jobMetadata")
    def job_metadata(self) -> 'outputs.JobMetadataResponse':
        """
        This field is populated by the Dataflow service to support filtering jobs by the metadata values provided here. Populated for ListJobs and all GetJob views SUMMARY and higher.
        """
        return pulumi.get(self, "job_metadata")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        User-defined labels for this job. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: * Keys must conform to regexp: \\p{Ll}\\p{Lo}{0,62} * Values must conform to regexp: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63} * Both keys and values are additionally constrained to be <= 128 bytes in size.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints) that contains this job.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The user-specified Cloud Dataflow job name. Only one Job with a given name can exist in a project within one region at any given time. Jobs in different regions can have the same name. If a caller attempts to create a Job with the same name as an already-existing Job, the attempt returns the existing Job. The name must match the regular expression `[a-z]([-a-z0-9]{0,1022}[a-z0-9])?`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pipelineDescription")
    def pipeline_description(self) -> 'outputs.PipelineDescriptionResponse':
        """
        Preliminary field: The format of this data may change at any time. A description of the user pipeline and stages through which it is executed. Created by Cloud Dataflow service. Only retrieved with JOB_VIEW_DESCRIPTION or JOB_VIEW_ALL.
        """
        return pulumi.get(self, "pipeline_description")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the Cloud Platform project that the job belongs to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="replaceJobId")
    def replace_job_id(self) -> str:
        """
        If this job is an update of an existing job, this field is the job ID of the job it replaced. When sending a `CreateJobRequest`, you can update a job by specifying it here. The job named here is stopped, and its intermediate state is transferred to this job.
        """
        return pulumi.get(self, "replace_job_id")

    @property
    @pulumi.getter(name="replacedByJobId")
    def replaced_by_job_id(self) -> str:
        """
        If another job is an update of this job (and thus, this job is in `JOB_STATE_UPDATED`), this field contains the ID of that job.
        """
        return pulumi.get(self, "replaced_by_job_id")

    @property
    @pulumi.getter(name="requestedState")
    def requested_state(self) -> str:
        """
        The job's requested state. `UpdateJob` may be used to switch between the `JOB_STATE_STOPPED` and `JOB_STATE_RUNNING` states, by setting requested_state. `UpdateJob` may also be used to directly set a job's requested state to `JOB_STATE_CANCELLED` or `JOB_STATE_DONE`, irrevocably terminating the job if it has not already reached a terminal state.
        """
        return pulumi.get(self, "requested_state")

    @property
    @pulumi.getter(name="runtimeUpdatableParams")
    def runtime_updatable_params(self) -> 'outputs.RuntimeUpdatableParamsResponse':
        """
        This field may ONLY be modified at runtime using the projects.jobs.update method to adjust job behavior. This field has no effect when specified at job creation.
        """
        return pulumi.get(self, "runtime_updatable_params")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use. This field is set only in responses from the server; it is ignored if it is set in any requests.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="stageStates")
    def stage_states(self) -> Sequence['outputs.ExecutionStageStateResponse']:
        """
        This field may be mutated by the Cloud Dataflow service; callers cannot mutate it.
        """
        return pulumi.get(self, "stage_states")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        The timestamp when the job was started (transitioned to JOB_STATE_PENDING). Flexible resource scheduling jobs are started with some delay after job creation, so start_time is unset before start and is updated when the job is started by the Cloud Dataflow service. For other jobs, start_time always equals to create_time and is immutable and set by the Cloud Dataflow service.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def steps(self) -> Sequence['outputs.StepResponse']:
        """
        Exactly one of step or steps_location should be specified. The top-level steps that constitute the entire job. Only retrieved with JOB_VIEW_ALL.
        """
        return pulumi.get(self, "steps")

    @property
    @pulumi.getter(name="stepsLocation")
    def steps_location(self) -> str:
        """
        The Cloud Storage location where the steps are stored.
        """
        return pulumi.get(self, "steps_location")

    @property
    @pulumi.getter(name="tempFiles")
    def temp_files(self) -> Sequence[str]:
        """
        A set of files the system should be aware of that are used for temporary storage. These temporary files will be removed on job completion. No duplicates are allowed. No file patterns are supported. The supported files are: Google Cloud Storage: storage.googleapis.com/{bucket}/{object} bucket.storage.googleapis.com/{object}
        """
        return pulumi.get(self, "temp_files")

    @property
    @pulumi.getter(name="transformNameMapping")
    def transform_name_mapping(self) -> Mapping[str, str]:
        """
        The map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job.
        """
        return pulumi.get(self, "transform_name_mapping")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of Cloud Dataflow job.
        """
        return pulumi.get(self, "type")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            client_request_id=self.client_request_id,
            create_time=self.create_time,
            created_from_snapshot_id=self.created_from_snapshot_id,
            current_state=self.current_state,
            current_state_time=self.current_state_time,
            environment=self.environment,
            execution_info=self.execution_info,
            job_metadata=self.job_metadata,
            labels=self.labels,
            location=self.location,
            name=self.name,
            pipeline_description=self.pipeline_description,
            project=self.project,
            replace_job_id=self.replace_job_id,
            replaced_by_job_id=self.replaced_by_job_id,
            requested_state=self.requested_state,
            runtime_updatable_params=self.runtime_updatable_params,
            satisfies_pzs=self.satisfies_pzs,
            stage_states=self.stage_states,
            start_time=self.start_time,
            steps=self.steps,
            steps_location=self.steps_location,
            temp_files=self.temp_files,
            transform_name_mapping=self.transform_name_mapping,
            type=self.type)


def get_job(job_id: Optional[str] = None,
            location: Optional[str] = None,
            project: Optional[str] = None,
            view: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Gets the state of the specified Cloud Dataflow job. To get the state of a job, we recommend using `projects.locations.jobs.get` with a [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using `projects.jobs.get` is not recommended, as you can only get the state of jobs that are running in `us-central1`.
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataflow/v1b3:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        client_request_id=__ret__.client_request_id,
        create_time=__ret__.create_time,
        created_from_snapshot_id=__ret__.created_from_snapshot_id,
        current_state=__ret__.current_state,
        current_state_time=__ret__.current_state_time,
        environment=__ret__.environment,
        execution_info=__ret__.execution_info,
        job_metadata=__ret__.job_metadata,
        labels=__ret__.labels,
        location=__ret__.location,
        name=__ret__.name,
        pipeline_description=__ret__.pipeline_description,
        project=__ret__.project,
        replace_job_id=__ret__.replace_job_id,
        replaced_by_job_id=__ret__.replaced_by_job_id,
        requested_state=__ret__.requested_state,
        runtime_updatable_params=__ret__.runtime_updatable_params,
        satisfies_pzs=__ret__.satisfies_pzs,
        stage_states=__ret__.stage_states,
        start_time=__ret__.start_time,
        steps=__ret__.steps,
        steps_location=__ret__.steps_location,
        temp_files=__ret__.temp_files,
        transform_name_mapping=__ret__.transform_name_mapping,
        type=__ret__.type)


@_utilities.lift_output_func(get_job)
def get_job_output(job_id: Optional[pulumi.Input[str]] = None,
                   location: Optional[pulumi.Input[str]] = None,
                   project: Optional[pulumi.Input[Optional[str]]] = None,
                   view: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobResult]:
    """
    Gets the state of the specified Cloud Dataflow job. To get the state of a job, we recommend using `projects.locations.jobs.get` with a [regional endpoint] (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints). Using `projects.jobs.get` is not recommended, as you can only get the state of jobs that are running in `us-central1`.
    """
    ...
