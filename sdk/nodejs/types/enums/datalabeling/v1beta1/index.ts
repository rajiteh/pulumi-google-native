// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EvaluationJobState = {
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The job is scheduled to run at the configured interval. You can pause or delete the job. When the job is in this state, it samples prediction input and output from your model version into your BigQuery table as predictions occur.
     */
    Scheduled: "SCHEDULED",
    /**
     * The job is currently running. When the job runs, Data Labeling Service does several things: 1. If you have configured your job to use Data Labeling Service for ground truth labeling, the service creates a Dataset and a labeling task for all data sampled since the last time the job ran. Human labelers provide ground truth labels for your data. Human labeling may take hours, or even days, depending on how much data has been sampled. The job remains in the `RUNNING` state during this time, and it can even be running multiple times in parallel if it gets triggered again (for example 24 hours later) before the earlier run has completed. When human labelers have finished labeling the data, the next step occurs. If you have configured your job to provide your own ground truth labels, Data Labeling Service still creates a Dataset for newly sampled data, but it expects that you have already added ground truth labels to the BigQuery table by this time. The next step occurs immediately. 2. Data Labeling Service creates an Evaluation by comparing your model version's predictions with the ground truth labels. If the job remains in this state for a long time, it continues to sample prediction data into your BigQuery table and will run again at the next interval, even if it causes the job to run multiple times in parallel.
     */
    Running: "RUNNING",
    /**
     * The job is not sampling prediction input and output into your BigQuery table and it will not run according to its schedule. You can resume the job.
     */
    Paused: "PAUSED",
    /**
     * The job has this state right before it is deleted.
     */
    Stopped: "STOPPED",
} as const;

/**
 * Output only. Describes the current state of the job.
 */
export type EvaluationJobState = (typeof EvaluationJobState)[keyof typeof EvaluationJobState];

export const GoogleCloudDatalabelingV1beta1ImageClassificationConfigAnswerAggregationType = {
    StringAggregationTypeUnspecified: "STRING_AGGREGATION_TYPE_UNSPECIFIED",
    /**
     * Majority vote to aggregate answers.
     */
    MajorityVote: "MAJORITY_VOTE",
    /**
     * Unanimous answers will be adopted.
     */
    UnanimousVote: "UNANIMOUS_VOTE",
    /**
     * Preserve all answers by crowd compute.
     */
    NoAggregation: "NO_AGGREGATION",
} as const;

/**
 * Optional. The type of how to aggregate answers.
 */
export type GoogleCloudDatalabelingV1beta1ImageClassificationConfigAnswerAggregationType = (typeof GoogleCloudDatalabelingV1beta1ImageClassificationConfigAnswerAggregationType)[keyof typeof GoogleCloudDatalabelingV1beta1ImageClassificationConfigAnswerAggregationType];

export const GoogleCloudDatalabelingV1beta1InputConfigAnnotationType = {
    AnnotationTypeUnspecified: "ANNOTATION_TYPE_UNSPECIFIED",
    /**
     * Classification annotations in an image. Allowed for continuous evaluation.
     */
    ImageClassificationAnnotation: "IMAGE_CLASSIFICATION_ANNOTATION",
    /**
     * Bounding box annotations in an image. A form of image object detection. Allowed for continuous evaluation.
     */
    ImageBoundingBoxAnnotation: "IMAGE_BOUNDING_BOX_ANNOTATION",
    /**
     * Oriented bounding box. The box does not have to be parallel to horizontal line.
     */
    ImageOrientedBoundingBoxAnnotation: "IMAGE_ORIENTED_BOUNDING_BOX_ANNOTATION",
    /**
     * Bounding poly annotations in an image.
     */
    ImageBoundingPolyAnnotation: "IMAGE_BOUNDING_POLY_ANNOTATION",
    /**
     * Polyline annotations in an image.
     */
    ImagePolylineAnnotation: "IMAGE_POLYLINE_ANNOTATION",
    /**
     * Segmentation annotations in an image.
     */
    ImageSegmentationAnnotation: "IMAGE_SEGMENTATION_ANNOTATION",
    /**
     * Classification annotations in video shots.
     */
    VideoShotsClassificationAnnotation: "VIDEO_SHOTS_CLASSIFICATION_ANNOTATION",
    /**
     * Video object tracking annotation.
     */
    VideoObjectTrackingAnnotation: "VIDEO_OBJECT_TRACKING_ANNOTATION",
    /**
     * Video object detection annotation.
     */
    VideoObjectDetectionAnnotation: "VIDEO_OBJECT_DETECTION_ANNOTATION",
    /**
     * Video event annotation.
     */
    VideoEventAnnotation: "VIDEO_EVENT_ANNOTATION",
    /**
     * Classification for text. Allowed for continuous evaluation.
     */
    TextClassificationAnnotation: "TEXT_CLASSIFICATION_ANNOTATION",
    /**
     * Entity extraction for text.
     */
    TextEntityExtractionAnnotation: "TEXT_ENTITY_EXTRACTION_ANNOTATION",
    /**
     * General classification. Allowed for continuous evaluation.
     */
    GeneralClassificationAnnotation: "GENERAL_CLASSIFICATION_ANNOTATION",
} as const;

/**
 * Optional. The type of annotation to be performed on this data. You must specify this field if you are using this InputConfig in an EvaluationJob.
 */
export type GoogleCloudDatalabelingV1beta1InputConfigAnnotationType = (typeof GoogleCloudDatalabelingV1beta1InputConfigAnnotationType)[keyof typeof GoogleCloudDatalabelingV1beta1InputConfigAnnotationType];

export const GoogleCloudDatalabelingV1beta1InputConfigDataType = {
    /**
     * Data type is unspecified.
     */
    DataTypeUnspecified: "DATA_TYPE_UNSPECIFIED",
    /**
     * Allowed for continuous evaluation.
     */
    Image: "IMAGE",
    /**
     * Video data type.
     */
    Video: "VIDEO",
    /**
     * Allowed for continuous evaluation.
     */
    Text: "TEXT",
    /**
     * Allowed for continuous evaluation.
     */
    GeneralData: "GENERAL_DATA",
} as const;

/**
 * Required. Data type must be specifed when user tries to import data.
 */
export type GoogleCloudDatalabelingV1beta1InputConfigDataType = (typeof GoogleCloudDatalabelingV1beta1InputConfigDataType)[keyof typeof GoogleCloudDatalabelingV1beta1InputConfigDataType];

export const InstructionDataType = {
    /**
     * Data type is unspecified.
     */
    DataTypeUnspecified: "DATA_TYPE_UNSPECIFIED",
    /**
     * Allowed for continuous evaluation.
     */
    Image: "IMAGE",
    /**
     * Video data type.
     */
    Video: "VIDEO",
    /**
     * Allowed for continuous evaluation.
     */
    Text: "TEXT",
    /**
     * Allowed for continuous evaluation.
     */
    GeneralData: "GENERAL_DATA",
} as const;

/**
 * Required. The data type of this instruction.
 */
export type InstructionDataType = (typeof InstructionDataType)[keyof typeof InstructionDataType];