// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Jobs.V4
{
    /// <summary>
    /// Creates a new job. Typically, the job becomes searchable within 10 seconds, but it may take up to 5 minutes.
    /// </summary>
    [GoogleNativeResourceType("google-native:jobs/v4:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the same company, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
        /// </summary>
        [Output("addresses")]
        public Output<ImmutableArray<string>> Addresses { get; private set; } = null!;

        /// <summary>
        /// Job application information.
        /// </summary>
        [Output("applicationInfo")]
        public Output<Outputs.ApplicationInfoResponse> ApplicationInfo { get; private set; } = null!;

        /// <summary>
        /// The resource name of the company listing the job. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For example, "projects/foo/tenants/bar/companies/baz".
        /// </summary>
        [Output("company")]
        public Output<string> Company { get; private set; } = null!;

        /// <summary>
        /// Display name of the company listing the job.
        /// </summary>
        [Output("companyDisplayName")]
        public Output<string> CompanyDisplayName { get; private set; } = null!;

        /// <summary>
        /// Job compensation information (a.k.a. "pay rate") i.e., the compensation that will paid to the employee.
        /// </summary>
        [Output("compensationInfo")]
        public Output<Outputs.CompensationInfoResponse> CompensationInfo { get; private set; } = null!;

        /// <summary>
        /// A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: `a-zA-Z*`. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The desired education degrees for the job, such as Bachelors, Masters.
        /// </summary>
        [Output("degreeTypes")]
        public Output<ImmutableArray<string>> DegreeTypes { get; private set; } = null!;

        /// <summary>
        /// The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
        /// </summary>
        [Output("department")]
        public Output<string> Department { get; private set; } = null!;

        /// <summary>
        /// Derived details about the job posting.
        /// </summary>
        [Output("derivedInfo")]
        public Output<Outputs.JobDerivedInfoResponse> DerivedInfo { get; private set; } = null!;

        /// <summary>
        /// The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The employment type(s) of a job, for example, full time or part time.
        /// </summary>
        [Output("employmentTypes")]
        public Output<ImmutableArray<string>> EmploymentTypes { get; private set; } = null!;

        /// <summary>
        /// A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Output("incentives")]
        public Output<string> Incentives { get; private set; } = null!;

        /// <summary>
        /// The benefits included with the job.
        /// </summary>
        [Output("jobBenefits")]
        public Output<ImmutableArray<string>> JobBenefits { get; private set; } = null!;

        /// <summary>
        /// The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        [Output("jobEndTime")]
        public Output<string> JobEndTime { get; private set; } = null!;

        /// <summary>
        /// The experience level associated with the job, such as "Entry Level".
        /// </summary>
        [Output("jobLevel")]
        public Output<string> JobLevel { get; private set; } = null!;

        /// <summary>
        /// The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        [Output("jobStartTime")]
        public Output<string> JobStartTime { get; private set; } = null!;

        /// <summary>
        /// The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
        /// </summary>
        [Output("languageCode")]
        public Output<string> LanguageCode { get; private set; } = null!;

        /// <summary>
        /// Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}". For example, "projects/foo/tenants/bar/jobs/baz". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this job posting was created.
        /// </summary>
        [Output("postingCreateTime")]
        public Output<string> PostingCreateTime { get; private set; } = null!;

        /// <summary>
        /// Strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be listed by the ListJobs API, but it can be retrieved with the GetJob API or updated with the UpdateJob API or deleted with the DeleteJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum number of open jobs over previous 7 days. If this threshold is exceeded, expired jobs are cleaned out in order of earliest expire time. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. If the timestamp is before the instant request is made, the job is treated as expired immediately on creation. This kind of job can not be updated. And when creating a job with past timestamp, the posting_publish_time must be set before posting_expire_time. The purpose of this feature is to allow other objects, such as Application, to refer a job that didn't exist in the system prior to becoming expired. If you want to modify a job that was expired on creation, delete it and create a new one. If this value isn't provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value isn't provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include job_end_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
        /// </summary>
        [Output("postingExpireTime")]
        public Output<string> PostingExpireTime { get; private set; } = null!;

        /// <summary>
        /// The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
        /// </summary>
        [Output("postingPublishTime")]
        public Output<string> PostingPublishTime { get; private set; } = null!;

        /// <summary>
        /// The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
        /// </summary>
        [Output("postingRegion")]
        public Output<string> PostingRegion { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this job posting was last updated.
        /// </summary>
        [Output("postingUpdateTime")]
        public Output<string> PostingUpdateTime { get; private set; } = null!;

        /// <summary>
        /// Options for job processing.
        /// </summary>
        [Output("processingOptions")]
        public Output<Outputs.ProcessingOptionsResponse> ProcessingOptions { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue &gt;0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
        /// </summary>
        [Output("promotionValue")]
        public Output<int> PromotionValue { get; private set; } = null!;

        /// <summary>
        /// A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Output("qualifications")]
        public Output<string> Qualifications { get; private set; } = null!;

        /// <summary>
        /// The requisition ID, also referred to as the posting ID, is assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job isn't allowed to be created if there is another job with the same company, language_code and requisition_id. The maximum number of allowed characters is 255.
        /// </summary>
        [Output("requisitionId")]
        public Output<string> RequisitionId { get; private set; } = null!;

        /// <summary>
        /// A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Output("responsibilities")]
        public Output<string> Responsibilities { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("google-native:jobs/v4:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:jobs/v4:Job", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                    "tenantId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Job(name, id, options);
        }
    }

    public sealed class JobArgs : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;

        /// <summary>
        /// Strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the same company, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        /// <summary>
        /// Job application information.
        /// </summary>
        [Input("applicationInfo")]
        public Input<Inputs.ApplicationInfoArgs>? ApplicationInfo { get; set; }

        /// <summary>
        /// The resource name of the company listing the job. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For example, "projects/foo/tenants/bar/companies/baz".
        /// </summary>
        [Input("company", required: true)]
        public Input<string> Company { get; set; } = null!;

        /// <summary>
        /// Job compensation information (a.k.a. "pay rate") i.e., the compensation that will paid to the employee.
        /// </summary>
        [Input("compensationInfo")]
        public Input<Inputs.CompensationInfoArgs>? CompensationInfo { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: `a-zA-Z*`. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        [Input("degreeTypes")]
        private InputList<Pulumi.GoogleNative.Jobs.V4.JobDegreeTypesItem>? _degreeTypes;

        /// <summary>
        /// The desired education degrees for the job, such as Bachelors, Masters.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Jobs.V4.JobDegreeTypesItem> DegreeTypes
        {
            get => _degreeTypes ?? (_degreeTypes = new InputList<Pulumi.GoogleNative.Jobs.V4.JobDegreeTypesItem>());
            set => _degreeTypes = value;
        }

        /// <summary>
        /// The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
        /// </summary>
        [Input("department")]
        public Input<string>? Department { get; set; }

        /// <summary>
        /// The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("employmentTypes")]
        private InputList<Pulumi.GoogleNative.Jobs.V4.JobEmploymentTypesItem>? _employmentTypes;

        /// <summary>
        /// The employment type(s) of a job, for example, full time or part time.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Jobs.V4.JobEmploymentTypesItem> EmploymentTypes
        {
            get => _employmentTypes ?? (_employmentTypes = new InputList<Pulumi.GoogleNative.Jobs.V4.JobEmploymentTypesItem>());
            set => _employmentTypes = value;
        }

        /// <summary>
        /// A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Input("incentives")]
        public Input<string>? Incentives { get; set; }

        [Input("jobBenefits")]
        private InputList<Pulumi.GoogleNative.Jobs.V4.JobJobBenefitsItem>? _jobBenefits;

        /// <summary>
        /// The benefits included with the job.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Jobs.V4.JobJobBenefitsItem> JobBenefits
        {
            get => _jobBenefits ?? (_jobBenefits = new InputList<Pulumi.GoogleNative.Jobs.V4.JobJobBenefitsItem>());
            set => _jobBenefits = value;
        }

        /// <summary>
        /// The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        [Input("jobEndTime")]
        public Input<string>? JobEndTime { get; set; }

        /// <summary>
        /// The experience level associated with the job, such as "Entry Level".
        /// </summary>
        [Input("jobLevel")]
        public Input<Pulumi.GoogleNative.Jobs.V4.JobJobLevel>? JobLevel { get; set; }

        /// <summary>
        /// The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
        /// </summary>
        [Input("jobStartTime")]
        public Input<string>? JobStartTime { get; set; }

        /// <summary>
        /// The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}". For example, "projects/foo/tenants/bar/jobs/baz". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be listed by the ListJobs API, but it can be retrieved with the GetJob API or updated with the UpdateJob API or deleted with the DeleteJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum number of open jobs over previous 7 days. If this threshold is exceeded, expired jobs are cleaned out in order of earliest expire time. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. If the timestamp is before the instant request is made, the job is treated as expired immediately on creation. This kind of job can not be updated. And when creating a job with past timestamp, the posting_publish_time must be set before posting_expire_time. The purpose of this feature is to allow other objects, such as Application, to refer a job that didn't exist in the system prior to becoming expired. If you want to modify a job that was expired on creation, delete it and create a new one. If this value isn't provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value isn't provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include job_end_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
        /// </summary>
        [Input("postingExpireTime")]
        public Input<string>? PostingExpireTime { get; set; }

        /// <summary>
        /// The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
        /// </summary>
        [Input("postingPublishTime")]
        public Input<string>? PostingPublishTime { get; set; }

        /// <summary>
        /// The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
        /// </summary>
        [Input("postingRegion")]
        public Input<Pulumi.GoogleNative.Jobs.V4.JobPostingRegion>? PostingRegion { get; set; }

        /// <summary>
        /// Options for job processing.
        /// </summary>
        [Input("processingOptions")]
        public Input<Inputs.ProcessingOptionsArgs>? ProcessingOptions { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue &gt;0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
        /// </summary>
        [Input("promotionValue")]
        public Input<int>? PromotionValue { get; set; }

        /// <summary>
        /// A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Input("qualifications")]
        public Input<string>? Qualifications { get; set; }

        /// <summary>
        /// The requisition ID, also referred to as the posting ID, is assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job isn't allowed to be created if there is another job with the same company, language_code and requisition_id. The maximum number of allowed characters is 255.
        /// </summary>
        [Input("requisitionId", required: true)]
        public Input<string> RequisitionId { get; set; } = null!;

        /// <summary>
        /// A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
        /// </summary>
        [Input("responsibilities")]
        public Input<string>? Responsibilities { get; set; }

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        /// <summary>
        /// The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.GoogleNative.Jobs.V4.JobVisibility>? Visibility { get; set; }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }
}
