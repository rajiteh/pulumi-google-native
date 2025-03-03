// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the resource representation for a job in a project.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("google-native:dataproc/v1:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId   string  `pulumi:"jobId"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
}

type LookupJobResult struct {
	// Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
	Done bool `pulumi:"done"`
	// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	DriverControlFilesUri string `pulumi:"driverControlFilesUri"`
	// A URI pointing to the location of the stdout of the job's driver program.
	DriverOutputResourceUri string `pulumi:"driverOutputResourceUri"`
	// Optional. Driver scheduling configuration.
	DriverSchedulingConfig DriverSchedulingConfigResponse `pulumi:"driverSchedulingConfig"`
	// Optional. Job is a Hadoop job.
	HadoopJob HadoopJobResponse `pulumi:"hadoopJob"`
	// Optional. Job is a Hive job.
	HiveJob HiveJobResponse `pulumi:"hiveJob"`
	// A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
	JobUuid string `pulumi:"jobUuid"`
	// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Job is a Pig job.
	PigJob PigJobResponse `pulumi:"pigJob"`
	// Job information, including how, when, and where to run the job.
	Placement JobPlacementResponse `pulumi:"placement"`
	// Optional. Job is a Presto job.
	PrestoJob PrestoJobResponse `pulumi:"prestoJob"`
	// Optional. Job is a PySpark job.
	PysparkJob PySparkJobResponse `pulumi:"pysparkJob"`
	// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
	Reference JobReferenceResponse `pulumi:"reference"`
	// Optional. Job scheduling configuration.
	Scheduling JobSchedulingResponse `pulumi:"scheduling"`
	// Optional. Job is a Spark job.
	SparkJob SparkJobResponse `pulumi:"sparkJob"`
	// Optional. Job is a SparkR job.
	SparkRJob SparkRJobResponse `pulumi:"sparkRJob"`
	// Optional. Job is a SparkSql job.
	SparkSqlJob SparkSqlJobResponse `pulumi:"sparkSqlJob"`
	// The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
	Status JobStatusResponse `pulumi:"status"`
	// The previous job status.
	StatusHistory []JobStatusResponse `pulumi:"statusHistory"`
	// Optional. Job is a Trino job.
	TrinoJob TrinoJobResponse `pulumi:"trinoJob"`
	// The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
	YarnApplications []YarnApplicationResponse `pulumi:"yarnApplications"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobId   pulumi.StringInput    `pulumi:"jobId"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringInput    `pulumi:"region"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
func (o LookupJobResultOutput) Done() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.Done }).(pulumi.BoolOutput)
}

// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
func (o LookupJobResultOutput) DriverControlFilesUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.DriverControlFilesUri }).(pulumi.StringOutput)
}

// A URI pointing to the location of the stdout of the job's driver program.
func (o LookupJobResultOutput) DriverOutputResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.DriverOutputResourceUri }).(pulumi.StringOutput)
}

// Optional. Driver scheduling configuration.
func (o LookupJobResultOutput) DriverSchedulingConfig() DriverSchedulingConfigResponseOutput {
	return o.ApplyT(func(v LookupJobResult) DriverSchedulingConfigResponse { return v.DriverSchedulingConfig }).(DriverSchedulingConfigResponseOutput)
}

// Optional. Job is a Hadoop job.
func (o LookupJobResultOutput) HadoopJob() HadoopJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) HadoopJobResponse { return v.HadoopJob }).(HadoopJobResponseOutput)
}

// Optional. Job is a Hive job.
func (o LookupJobResultOutput) HiveJob() HiveJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) HiveJobResponse { return v.HiveJob }).(HiveJobResponseOutput)
}

// A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
func (o LookupJobResultOutput) JobUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.JobUuid }).(pulumi.StringOutput)
}

// Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
func (o LookupJobResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Job is a Pig job.
func (o LookupJobResultOutput) PigJob() PigJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) PigJobResponse { return v.PigJob }).(PigJobResponseOutput)
}

// Job information, including how, when, and where to run the job.
func (o LookupJobResultOutput) Placement() JobPlacementResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobPlacementResponse { return v.Placement }).(JobPlacementResponseOutput)
}

// Optional. Job is a Presto job.
func (o LookupJobResultOutput) PrestoJob() PrestoJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) PrestoJobResponse { return v.PrestoJob }).(PrestoJobResponseOutput)
}

// Optional. Job is a PySpark job.
func (o LookupJobResultOutput) PysparkJob() PySparkJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) PySparkJobResponse { return v.PysparkJob }).(PySparkJobResponseOutput)
}

// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
func (o LookupJobResultOutput) Reference() JobReferenceResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobReferenceResponse { return v.Reference }).(JobReferenceResponseOutput)
}

// Optional. Job scheduling configuration.
func (o LookupJobResultOutput) Scheduling() JobSchedulingResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobSchedulingResponse { return v.Scheduling }).(JobSchedulingResponseOutput)
}

// Optional. Job is a Spark job.
func (o LookupJobResultOutput) SparkJob() SparkJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SparkJobResponse { return v.SparkJob }).(SparkJobResponseOutput)
}

// Optional. Job is a SparkR job.
func (o LookupJobResultOutput) SparkRJob() SparkRJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SparkRJobResponse { return v.SparkRJob }).(SparkRJobResponseOutput)
}

// Optional. Job is a SparkSql job.
func (o LookupJobResultOutput) SparkSqlJob() SparkSqlJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SparkSqlJobResponse { return v.SparkSqlJob }).(SparkSqlJobResponseOutput)
}

// The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
func (o LookupJobResultOutput) Status() JobStatusResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobStatusResponse { return v.Status }).(JobStatusResponseOutput)
}

// The previous job status.
func (o LookupJobResultOutput) StatusHistory() JobStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupJobResult) []JobStatusResponse { return v.StatusHistory }).(JobStatusResponseArrayOutput)
}

// Optional. Job is a Trino job.
func (o LookupJobResultOutput) TrinoJob() TrinoJobResponseOutput {
	return o.ApplyT(func(v LookupJobResult) TrinoJobResponse { return v.TrinoJob }).(TrinoJobResponseOutput)
}

// The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
func (o LookupJobResultOutput) YarnApplications() YarnApplicationResponseArrayOutput {
	return o.ApplyT(func(v LookupJobResult) []YarnApplicationResponse { return v.YarnApplications }).(YarnApplicationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
