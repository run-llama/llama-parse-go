// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// BetaBatchService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaBatchService] method instead.
type BetaBatchService struct {
	options  []option.RequestOption
	JobItems BetaBatchJobItemService
}

// NewBetaBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaBatchService(opts ...option.RequestOption) (r BetaBatchService) {
	r = BetaBatchService{}
	r.options = opts
	r.JobItems = NewBetaBatchJobItemService(opts...)
	return
}

// Create a batch processing job.
//
// Processes files from a directory or a specific list of item IDs. Supports batch
// parsing and classification operations.
//
// Provide either `directory_id` to process all files in a directory, or `item_ids`
// for specific items. The job runs asynchronously — poll `GET /batch/{job_id}` for
// progress.
func (r *BetaBatchService) New(ctx context.Context, params BetaBatchNewParams, opts ...option.RequestOption) (res *BetaBatchNewResponse, err error) {
	if !param.IsOmitted(params.TemporalNamespace) {
		opts = append(opts, option.WithHeader("temporal-namespace", fmt.Sprintf("%v", params.TemporalNamespace.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/batch-processing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List batch processing jobs with optional filtering.
//
// Filter by `directory_id`, `job_type`, or `status`. Results are paginated with
// configurable `limit` and `offset`.
func (r *BetaBatchService) List(ctx context.Context, query BetaBatchListParams, opts ...option.RequestOption) (res *pagination.PaginatedBatchItems[BetaBatchListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/batch-processing"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List batch processing jobs with optional filtering.
//
// Filter by `directory_id`, `job_type`, or `status`. Results are paginated with
// configurable `limit` and `offset`.
func (r *BetaBatchService) ListAutoPaging(ctx context.Context, query BetaBatchListParams, opts ...option.RequestOption) *pagination.PaginatedBatchItemsAutoPager[BetaBatchListResponse] {
	return pagination.NewPaginatedBatchItemsAutoPager(r.List(ctx, query, opts...))
}

// Cancel a running batch processing job.
//
// Stops processing and marks pending items as cancelled. Items currently being
// processed may still complete.
func (r *BetaBatchService) Cancel(ctx context.Context, jobID string, params BetaBatchCancelParams, opts ...option.RequestOption) (res *BetaBatchCancelResponse, err error) {
	if !param.IsOmitted(params.TemporalNamespace) {
		opts = append(opts, option.WithHeader("temporal-namespace", fmt.Sprintf("%v", params.TemporalNamespace.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/batch-processing/%s/cancel", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get detailed status of a batch processing job.
//
// Returns current progress percentage, file counts (total, processed, failed,
// skipped), and timestamps.
func (r *BetaBatchService) GetStatus(ctx context.Context, jobID string, query BetaBatchGetStatusParams, opts ...option.RequestOption) (res *BetaBatchGetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/batch-processing/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response schema for a batch processing job.
type BetaBatchNewResponse struct {
	// Unique identifier for the batch job
	ID string `json:"id" api:"required"`
	// Type of processing operation (parse or classify)
	//
	// Any of "parse", "extract", "classify".
	JobType BetaBatchNewResponseJobType `json:"job_type" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status
	//
	// Any of "pending", "running", "dispatched", "completed", "failed", "cancelled".
	Status BetaBatchNewResponseStatus `json:"status" api:"required"`
	// Total number of items in the job
	TotalItems int64 `json:"total_items" api:"required"`
	// Timestamp when job completed
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Directory being processed
	DirectoryID string    `json:"directory_id" api:"nullable"`
	EffectiveAt time.Time `json:"effective_at" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Number of items that failed processing
	FailedItems int64 `json:"failed_items"`
	// The job record ID associated with this status, if any.
	JobRecordID string `json:"job_record_id" api:"nullable"`
	// Number of items processed so far
	ProcessedItems int64 `json:"processed_items"`
	// Number of items skipped (already processed or size limit)
	SkippedItems int64 `json:"skipped_items"`
	// Timestamp when job processing started
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Async job tracking ID
	WorkflowID string `json:"workflow_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		JobType        respjson.Field
		ProjectID      respjson.Field
		Status         respjson.Field
		TotalItems     respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DirectoryID    respjson.Field
		EffectiveAt    respjson.Field
		ErrorMessage   respjson.Field
		FailedItems    respjson.Field
		JobRecordID    respjson.Field
		ProcessedItems respjson.Field
		SkippedItems   respjson.Field
		StartedAt      respjson.Field
		UpdatedAt      respjson.Field
		WorkflowID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of processing operation (parse or classify)
type BetaBatchNewResponseJobType string

const (
	BetaBatchNewResponseJobTypeParse    BetaBatchNewResponseJobType = "parse"
	BetaBatchNewResponseJobTypeExtract  BetaBatchNewResponseJobType = "extract"
	BetaBatchNewResponseJobTypeClassify BetaBatchNewResponseJobType = "classify"
)

// Current job status
type BetaBatchNewResponseStatus string

const (
	BetaBatchNewResponseStatusPending    BetaBatchNewResponseStatus = "pending"
	BetaBatchNewResponseStatusRunning    BetaBatchNewResponseStatus = "running"
	BetaBatchNewResponseStatusDispatched BetaBatchNewResponseStatus = "dispatched"
	BetaBatchNewResponseStatusCompleted  BetaBatchNewResponseStatus = "completed"
	BetaBatchNewResponseStatusFailed     BetaBatchNewResponseStatus = "failed"
	BetaBatchNewResponseStatusCancelled  BetaBatchNewResponseStatus = "cancelled"
)

// Response schema for a batch processing job.
type BetaBatchListResponse struct {
	// Unique identifier for the batch job
	ID string `json:"id" api:"required"`
	// Type of processing operation (parse or classify)
	//
	// Any of "parse", "extract", "classify".
	JobType BetaBatchListResponseJobType `json:"job_type" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status
	//
	// Any of "pending", "running", "dispatched", "completed", "failed", "cancelled".
	Status BetaBatchListResponseStatus `json:"status" api:"required"`
	// Total number of items in the job
	TotalItems int64 `json:"total_items" api:"required"`
	// Timestamp when job completed
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Directory being processed
	DirectoryID string    `json:"directory_id" api:"nullable"`
	EffectiveAt time.Time `json:"effective_at" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Number of items that failed processing
	FailedItems int64 `json:"failed_items"`
	// The job record ID associated with this status, if any.
	JobRecordID string `json:"job_record_id" api:"nullable"`
	// Number of items processed so far
	ProcessedItems int64 `json:"processed_items"`
	// Number of items skipped (already processed or size limit)
	SkippedItems int64 `json:"skipped_items"`
	// Timestamp when job processing started
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Async job tracking ID
	WorkflowID string `json:"workflow_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		JobType        respjson.Field
		ProjectID      respjson.Field
		Status         respjson.Field
		TotalItems     respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DirectoryID    respjson.Field
		EffectiveAt    respjson.Field
		ErrorMessage   respjson.Field
		FailedItems    respjson.Field
		JobRecordID    respjson.Field
		ProcessedItems respjson.Field
		SkippedItems   respjson.Field
		StartedAt      respjson.Field
		UpdatedAt      respjson.Field
		WorkflowID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of processing operation (parse or classify)
type BetaBatchListResponseJobType string

const (
	BetaBatchListResponseJobTypeParse    BetaBatchListResponseJobType = "parse"
	BetaBatchListResponseJobTypeExtract  BetaBatchListResponseJobType = "extract"
	BetaBatchListResponseJobTypeClassify BetaBatchListResponseJobType = "classify"
)

// Current job status
type BetaBatchListResponseStatus string

const (
	BetaBatchListResponseStatusPending    BetaBatchListResponseStatus = "pending"
	BetaBatchListResponseStatusRunning    BetaBatchListResponseStatus = "running"
	BetaBatchListResponseStatusDispatched BetaBatchListResponseStatus = "dispatched"
	BetaBatchListResponseStatusCompleted  BetaBatchListResponseStatus = "completed"
	BetaBatchListResponseStatusFailed     BetaBatchListResponseStatus = "failed"
	BetaBatchListResponseStatusCancelled  BetaBatchListResponseStatus = "cancelled"
)

// Response after cancelling a batch job.
type BetaBatchCancelResponse struct {
	// ID of the cancelled job
	JobID string `json:"job_id" api:"required"`
	// Confirmation message
	Message string `json:"message" api:"required"`
	// Number of items processed before cancellation
	ProcessedItems int64 `json:"processed_items" api:"required"`
	// New status (should be 'cancelled')
	//
	// Any of "pending", "running", "dispatched", "completed", "failed", "cancelled".
	Status BetaBatchCancelResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID          respjson.Field
		Message        respjson.Field
		ProcessedItems respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New status (should be 'cancelled')
type BetaBatchCancelResponseStatus string

const (
	BetaBatchCancelResponseStatusPending    BetaBatchCancelResponseStatus = "pending"
	BetaBatchCancelResponseStatusRunning    BetaBatchCancelResponseStatus = "running"
	BetaBatchCancelResponseStatusDispatched BetaBatchCancelResponseStatus = "dispatched"
	BetaBatchCancelResponseStatusCompleted  BetaBatchCancelResponseStatus = "completed"
	BetaBatchCancelResponseStatusFailed     BetaBatchCancelResponseStatus = "failed"
	BetaBatchCancelResponseStatusCancelled  BetaBatchCancelResponseStatus = "cancelled"
)

// Detailed status response for a batch processing job.
type BetaBatchGetStatusResponse struct {
	// Response schema for a batch processing job.
	Job BetaBatchGetStatusResponseJob `json:"job" api:"required"`
	// Percentage of items processed (0-100)
	ProgressPercentage float64 `json:"progress_percentage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job                respjson.Field
		ProgressPercentage respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response schema for a batch processing job.
type BetaBatchGetStatusResponseJob struct {
	// Unique identifier for the batch job
	ID string `json:"id" api:"required"`
	// Type of processing operation (parse or classify)
	//
	// Any of "parse", "extract", "classify".
	JobType string `json:"job_type" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status
	//
	// Any of "pending", "running", "dispatched", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Total number of items in the job
	TotalItems int64 `json:"total_items" api:"required"`
	// Timestamp when job completed
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Directory being processed
	DirectoryID string    `json:"directory_id" api:"nullable"`
	EffectiveAt time.Time `json:"effective_at" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Number of items that failed processing
	FailedItems int64 `json:"failed_items"`
	// The job record ID associated with this status, if any.
	JobRecordID string `json:"job_record_id" api:"nullable"`
	// Number of items processed so far
	ProcessedItems int64 `json:"processed_items"`
	// Number of items skipped (already processed or size limit)
	SkippedItems int64 `json:"skipped_items"`
	// Timestamp when job processing started
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Async job tracking ID
	WorkflowID string `json:"workflow_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		JobType        respjson.Field
		ProjectID      respjson.Field
		Status         respjson.Field
		TotalItems     respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DirectoryID    respjson.Field
		EffectiveAt    respjson.Field
		ErrorMessage   respjson.Field
		FailedItems    respjson.Field
		JobRecordID    respjson.Field
		ProcessedItems respjson.Field
		SkippedItems   respjson.Field
		StartedAt      respjson.Field
		UpdatedAt      respjson.Field
		WorkflowID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchGetStatusResponseJob) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchGetStatusResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaBatchNewParams struct {
	// Job configuration — either a parse or classify config
	JobConfig      BetaBatchNewParamsJobConfigUnion `json:"job_config,omitzero" api:"required"`
	OrganizationID param.Opt[string]                `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]                `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Maximum files to process per execution cycle in directory mode. Defaults to
	// page_size.
	ContinueAsNewThreshold param.Opt[int64] `json:"continue_as_new_threshold,omitzero"`
	// ID of the directory containing files to process
	DirectoryID param.Opt[string] `json:"directory_id,omitzero"`
	// Number of files to process per batch when using directory mode
	PageSize          param.Opt[int64]  `json:"page_size,omitzero"`
	TemporalNamespace param.Opt[string] `header:"temporal-namespace,omitzero" json:"-"`
	// List of specific item IDs to process. Either this or directory_id must be
	// provided.
	ItemIDs []string `json:"item_ids,omitzero"`
	paramObj
}

func (r BetaBatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaBatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaBatchNewParams]'s query parameters as `url.Values`.
func (r BetaBatchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaBatchNewParamsJobConfigUnion struct {
	OfBatchParseJobRecordCreate *BetaBatchNewParamsJobConfigBatchParseJobRecordCreate `json:",omitzero,inline"`
	OfClassifyJob               *ClassifyJobParam                                     `json:",omitzero,inline"`
	paramUnion
}

func (u BetaBatchNewParamsJobConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBatchParseJobRecordCreate, u.OfClassifyJob)
}
func (u *BetaBatchNewParamsJobConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Batch-specific parse job record for batch processing.
//
// This model contains the metadata and configuration for a batch parse job, but
// excludes file-specific information. It's used as input to the batch parent
// workflow and combined with DirectoryFile data to create full
// ParseJobRecordCreate instances for each file.
//
// Attributes: job_name: Must be PARSE_RAW_FILE partitions: Partitions for job
// output location parameters: Generic parse configuration (BatchParseJobConfig)
// session_id: Upstream request ID for tracking correlation_id: Correlation ID for
// cross-service tracking parent_job_execution_id: Parent job execution ID if
// nested user_id: User who created the job project_id: Project this job belongs to
// webhook_url: Optional webhook URL for job completion notifications
type BetaBatchNewParamsJobConfigBatchParseJobRecordCreate struct {
	// The correlation ID for this job. Used for tracking the job across services.
	CorrelationID param.Opt[string] `json:"correlation_id,omitzero" format:"uuid"`
	// The ID of the parent job execution.
	ParentJobExecutionID param.Opt[string] `json:"parent_job_execution_id,omitzero" format:"uuid"`
	// The ID of the project this job belongs to.
	ProjectID param.Opt[string] `json:"project_id,omitzero" format:"uuid"`
	// The upstream request ID that created this job. Used for tracking the job across
	// services.
	SessionID param.Opt[string] `json:"session_id,omitzero" format:"uuid"`
	// The ID of the user that created this job
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// The URL that needs to be called at the end of the parsing job.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Generic parse job configuration for batch processing.
	//
	// This model contains the parsing configuration that applies to all files in a
	// batch, but excludes file-specific fields like file_name, file_id, etc. Those
	// file-specific fields are populated from DirectoryFile data when creating
	// individual ParseJobRecordCreate instances for each file.
	//
	// The fields in this model should be generic settings that apply uniformly to all
	// files being processed in the batch.
	Parameters BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters `json:"parameters,omitzero"`
	// Any of "parse_raw_file_job".
	JobName string `json:"job_name,omitzero"`
	// The partitions for this execution. Used for determining where to save job
	// output.
	Partitions map[string]string `json:"partitions,omitzero" format:"uuid"`
	paramObj
}

func (r BetaBatchNewParamsJobConfigBatchParseJobRecordCreate) MarshalJSON() (data []byte, err error) {
	type shadow BetaBatchNewParamsJobConfigBatchParseJobRecordCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaBatchNewParamsJobConfigBatchParseJobRecordCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaBatchNewParamsJobConfigBatchParseJobRecordCreate](
		"job_name", "parse_raw_file_job",
	)
}

// Generic parse job configuration for batch processing.
//
// This model contains the parsing configuration that applies to all files in a
// batch, but excludes file-specific fields like file_name, file_id, etc. Those
// file-specific fields are populated from DirectoryFile data when creating
// individual ParseJobRecordCreate instances for each file.
//
// The fields in this model should be generic settings that apply uniformly to all
// files being processed in the batch.
type BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters struct {
	AdaptiveLongTable                        param.Opt[bool]    `json:"adaptive_long_table,omitzero"`
	AggressiveTableExtraction                param.Opt[bool]    `json:"aggressive_table_extraction,omitzero"`
	AnnotateLinks                            param.Opt[bool]    `json:"annotate_links,omitzero"`
	AutoMode                                 param.Opt[bool]    `json:"auto_mode,omitzero"`
	AutoModeConfigurationJson                param.Opt[string]  `json:"auto_mode_configuration_json,omitzero"`
	AutoModeTriggerOnImageInPage             param.Opt[bool]    `json:"auto_mode_trigger_on_image_in_page,omitzero"`
	AutoModeTriggerOnRegexpInPage            param.Opt[string]  `json:"auto_mode_trigger_on_regexp_in_page,omitzero"`
	AutoModeTriggerOnTableInPage             param.Opt[bool]    `json:"auto_mode_trigger_on_table_in_page,omitzero"`
	AutoModeTriggerOnTextInPage              param.Opt[string]  `json:"auto_mode_trigger_on_text_in_page,omitzero"`
	AzureOpenAIAPIVersion                    param.Opt[string]  `json:"azure_openai_api_version,omitzero"`
	AzureOpenAIDeploymentName                param.Opt[string]  `json:"azure_openai_deployment_name,omitzero"`
	AzureOpenAIEndpoint                      param.Opt[string]  `json:"azure_openai_endpoint,omitzero"`
	AzureOpenAIKey                           param.Opt[string]  `json:"azure_openai_key,omitzero"`
	BboxBottom                               param.Opt[float64] `json:"bbox_bottom,omitzero"`
	BboxLeft                                 param.Opt[float64] `json:"bbox_left,omitzero"`
	BboxRight                                param.Opt[float64] `json:"bbox_right,omitzero"`
	BboxTop                                  param.Opt[float64] `json:"bbox_top,omitzero"`
	BoundingBox                              param.Opt[string]  `json:"bounding_box,omitzero"`
	CompactMarkdownTable                     param.Opt[bool]    `json:"compact_markdown_table,omitzero"`
	ComplementalFormattingInstruction        param.Opt[string]  `json:"complemental_formatting_instruction,omitzero"`
	ContentGuidelineInstruction              param.Opt[string]  `json:"content_guideline_instruction,omitzero"`
	ContinuousMode                           param.Opt[bool]    `json:"continuous_mode,omitzero"`
	DisableImageExtraction                   param.Opt[bool]    `json:"disable_image_extraction,omitzero"`
	DisableOcr                               param.Opt[bool]    `json:"disable_ocr,omitzero"`
	DisableReconstruction                    param.Opt[bool]    `json:"disable_reconstruction,omitzero"`
	DoNotCache                               param.Opt[bool]    `json:"do_not_cache,omitzero"`
	DoNotUnrollColumns                       param.Opt[bool]    `json:"do_not_unroll_columns,omitzero"`
	EnableCostOptimizer                      param.Opt[bool]    `json:"enable_cost_optimizer,omitzero"`
	ExtractCharts                            param.Opt[bool]    `json:"extract_charts,omitzero"`
	ExtractLayout                            param.Opt[bool]    `json:"extract_layout,omitzero"`
	ExtractPrintedPageNumber                 param.Opt[bool]    `json:"extract_printed_page_number,omitzero"`
	FastMode                                 param.Opt[bool]    `json:"fast_mode,omitzero"`
	FormattingInstruction                    param.Opt[string]  `json:"formatting_instruction,omitzero"`
	Gpt4oAPIKey                              param.Opt[string]  `json:"gpt4o_api_key,omitzero"`
	Gpt4oMode                                param.Opt[bool]    `json:"gpt4o_mode,omitzero"`
	GuessXlsxSheetName                       param.Opt[bool]    `json:"guess_xlsx_sheet_name,omitzero"`
	HideFooters                              param.Opt[bool]    `json:"hide_footers,omitzero"`
	HideHeaders                              param.Opt[bool]    `json:"hide_headers,omitzero"`
	HighResOcr                               param.Opt[bool]    `json:"high_res_ocr,omitzero"`
	HTMLMakeAllElementsVisible               param.Opt[bool]    `json:"html_make_all_elements_visible,omitzero"`
	HTMLRemoveFixedElements                  param.Opt[bool]    `json:"html_remove_fixed_elements,omitzero"`
	HTMLRemoveNavigationElements             param.Opt[bool]    `json:"html_remove_navigation_elements,omitzero"`
	HTTPProxy                                param.Opt[string]  `json:"http_proxy,omitzero"`
	IgnoreDocumentElementsForLayoutDetection param.Opt[bool]    `json:"ignore_document_elements_for_layout_detection,omitzero"`
	InlineImagesInMarkdown                   param.Opt[bool]    `json:"inline_images_in_markdown,omitzero"`
	InputS3Path                              param.Opt[string]  `json:"input_s3_path,omitzero"`
	// The region for the input S3 bucket.
	InputS3Region                         param.Opt[string]  `json:"input_s3_region,omitzero"`
	InputURL                              param.Opt[string]  `json:"input_url,omitzero"`
	InternalIsScreenshotJob               param.Opt[bool]    `json:"internal_is_screenshot_job,omitzero"`
	InvalidateCache                       param.Opt[bool]    `json:"invalidate_cache,omitzero"`
	IsFormattingInstruction               param.Opt[bool]    `json:"is_formatting_instruction,omitzero"`
	JobTimeoutExtraTimePerPageInSeconds   param.Opt[float64] `json:"job_timeout_extra_time_per_page_in_seconds,omitzero"`
	JobTimeoutInSeconds                   param.Opt[float64] `json:"job_timeout_in_seconds,omitzero"`
	KeepPageSeparatorWhenMergingTables    param.Opt[bool]    `json:"keep_page_separator_when_merging_tables,omitzero"`
	LayoutAware                           param.Opt[bool]    `json:"layout_aware,omitzero"`
	LineLevelBoundingBox                  param.Opt[bool]    `json:"line_level_bounding_box,omitzero"`
	MarkdownTableMultilineHeaderSeparator param.Opt[string]  `json:"markdown_table_multiline_header_separator,omitzero"`
	MaxPages                              param.Opt[int64]   `json:"max_pages,omitzero"`
	MaxPagesEnforced                      param.Opt[int64]   `json:"max_pages_enforced,omitzero"`
	MergeTablesAcrossPagesInMarkdown      param.Opt[bool]    `json:"merge_tables_across_pages_in_markdown,omitzero"`
	Model                                 param.Opt[string]  `json:"model,omitzero"`
	OutlinedTableExtraction               param.Opt[bool]    `json:"outlined_table_extraction,omitzero"`
	OutputPdfOfDocument                   param.Opt[bool]    `json:"output_pdf_of_document,omitzero"`
	// If specified, llamaParse will save the output to the specified path. All output
	// file will use this 'prefix' should be a valid s3:// url
	OutputS3PathPrefix param.Opt[string] `json:"output_s3_path_prefix,omitzero"`
	// The region for the output S3 bucket.
	OutputS3Region     param.Opt[string] `json:"output_s3_region,omitzero"`
	OutputTablesAsHTML param.Opt[bool]   `json:"output_tables_as_HTML,omitzero"`
	// The output bucket.
	OutputBucket       param.Opt[string]  `json:"outputBucket,omitzero"`
	PageErrorTolerance param.Opt[float64] `json:"page_error_tolerance,omitzero"`
	PageFooterPrefix   param.Opt[string]  `json:"page_footer_prefix,omitzero"`
	PageFooterSuffix   param.Opt[string]  `json:"page_footer_suffix,omitzero"`
	PageHeaderPrefix   param.Opt[string]  `json:"page_header_prefix,omitzero"`
	PageHeaderSuffix   param.Opt[string]  `json:"page_header_suffix,omitzero"`
	PagePrefix         param.Opt[string]  `json:"page_prefix,omitzero"`
	PageSeparator      param.Opt[string]  `json:"page_separator,omitzero"`
	PageSuffix         param.Opt[string]  `json:"page_suffix,omitzero"`
	ParsingInstruction param.Opt[string]  `json:"parsing_instruction,omitzero"`
	// The pipeline ID.
	PipelineID                              param.Opt[string] `json:"pipeline_id,omitzero"`
	PreciseBoundingBox                      param.Opt[bool]   `json:"precise_bounding_box,omitzero"`
	PremiumMode                             param.Opt[bool]   `json:"premium_mode,omitzero"`
	PresentationOutOfBoundsContent          param.Opt[bool]   `json:"presentation_out_of_bounds_content,omitzero"`
	PresentationSkipEmbeddedData            param.Opt[bool]   `json:"presentation_skip_embedded_data,omitzero"`
	PreserveLayoutAlignmentAcrossPages      param.Opt[bool]   `json:"preserve_layout_alignment_across_pages,omitzero"`
	PreserveVerySmallText                   param.Opt[bool]   `json:"preserve_very_small_text,omitzero"`
	Preset                                  param.Opt[string] `json:"preset,omitzero"`
	ProjectID                               param.Opt[string] `json:"project_id,omitzero"`
	RemoveHiddenText                        param.Opt[bool]   `json:"remove_hidden_text,omitzero"`
	ReplaceFailedPageWithErrorMessagePrefix param.Opt[string] `json:"replace_failed_page_with_error_message_prefix,omitzero"`
	ReplaceFailedPageWithErrorMessageSuffix param.Opt[string] `json:"replace_failed_page_with_error_message_suffix,omitzero"`
	SaveImages                              param.Opt[bool]   `json:"save_images,omitzero"`
	SkipDiagonalText                        param.Opt[bool]   `json:"skip_diagonal_text,omitzero"`
	SpecializedChartParsingAgentic          param.Opt[bool]   `json:"specialized_chart_parsing_agentic,omitzero"`
	SpecializedChartParsingEfficient        param.Opt[bool]   `json:"specialized_chart_parsing_efficient,omitzero"`
	SpecializedChartParsingPlus             param.Opt[bool]   `json:"specialized_chart_parsing_plus,omitzero"`
	SpecializedImageParsing                 param.Opt[bool]   `json:"specialized_image_parsing,omitzero"`
	SpreadsheetExtractSubTables             param.Opt[bool]   `json:"spreadsheet_extract_sub_tables,omitzero"`
	SpreadsheetForceFormulaComputation      param.Opt[bool]   `json:"spreadsheet_force_formula_computation,omitzero"`
	SpreadsheetIncludeHiddenSheets          param.Opt[bool]   `json:"spreadsheet_include_hidden_sheets,omitzero"`
	StrictModeBuggyFont                     param.Opt[bool]   `json:"strict_mode_buggy_font,omitzero"`
	StrictModeImageExtraction               param.Opt[bool]   `json:"strict_mode_image_extraction,omitzero"`
	StrictModeImageOcr                      param.Opt[bool]   `json:"strict_mode_image_ocr,omitzero"`
	StrictModeReconstruction                param.Opt[bool]   `json:"strict_mode_reconstruction,omitzero"`
	StructuredOutput                        param.Opt[bool]   `json:"structured_output,omitzero"`
	StructuredOutputJsonSchema              param.Opt[string] `json:"structured_output_json_schema,omitzero"`
	StructuredOutputJsonSchemaName          param.Opt[string] `json:"structured_output_json_schema_name,omitzero"`
	SystemPrompt                            param.Opt[string] `json:"system_prompt,omitzero"`
	SystemPromptAppend                      param.Opt[string] `json:"system_prompt_append,omitzero"`
	TakeScreenshot                          param.Opt[bool]   `json:"take_screenshot,omitzero"`
	TargetPages                             param.Opt[string] `json:"target_pages,omitzero"`
	Tier                                    param.Opt[string] `json:"tier,omitzero"`
	UseVendorMultimodalModel                param.Opt[bool]   `json:"use_vendor_multimodal_model,omitzero"`
	UserPrompt                              param.Opt[string] `json:"user_prompt,omitzero"`
	VendorMultimodalAPIKey                  param.Opt[string] `json:"vendor_multimodal_api_key,omitzero"`
	VendorMultimodalModelName               param.Opt[string] `json:"vendor_multimodal_model_name,omitzero"`
	Version                                 param.Opt[string] `json:"version,omitzero"`
	WebhookURL                              param.Opt[string] `json:"webhook_url,omitzero"`
	// The language.
	Lang param.Opt[string] `json:"lang,omitzero"`
	// The custom metadata to attach to the documents.
	CustomMetadata map[string]any `json:"custom_metadata,omitzero"`
	// Any of "screenshot", "embedded", "layout".
	ImagesToSave []string `json:"images_to_save,omitzero"`
	// The priority for the request. This field may be ignored or overwritten depending
	// on the organization tier.
	//
	// Any of "low", "medium", "high", "critical".
	Priority string `json:"priority,omitzero"`
	// The resource info about the file
	ResourceInfo map[string]any `json:"resource_info,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration `json:"webhook_configurations,omitzero"`
	Languages             []ParsingLanguages                                                                   `json:"languages,omitzero"`
	// Enum for representing the mode of parsing to be used.
	//
	// Any of "parse_page_without_llm", "parse_page_with_llm", "parse_page_with_lvm",
	// "parse_page_with_agent", "parse_page_with_layout_agent",
	// "parse_document_with_llm", "parse_document_with_lvm",
	// "parse_document_with_agent".
	ParseMode ParsingMode `json:"parse_mode,omitzero"`
	// Enum for representing the different available page error handling modes.
	//
	// Any of "raw_text", "blank_page", "error_message".
	ReplaceFailedPageMode FailPageMode `json:"replace_failed_page_mode,omitzero"`
	// Any of "parse".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters) MarshalJSON() (data []byte, err error) {
	type shadow BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters](
		"priority", "low", "medium", "high", "critical",
	)
	apijson.RegisterFieldValidator[BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters](
		"type", "parse",
	)
}

// Configuration for a single outbound webhook endpoint.
type BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration struct {
	// Response format sent to the webhook: 'string' (default) or 'json'
	WebhookOutputFormat param.Opt[string] `json:"webhook_output_format,omitzero"`
	// URL to receive webhook POST notifications
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events to subscribe to (e.g. 'parse.success', 'extract.error'). If null, all
	// events are delivered.
	//
	// Any of "extract.pending", "extract.success", "extract.error",
	// "extract.partial_success", "extract.cancelled", "parse.pending",
	// "parse.running", "parse.success", "parse.error", "parse.partial_success",
	// "parse.cancelled", "classify.pending", "classify.running", "classify.success",
	// "classify.error", "classify.partial_success", "classify.cancelled",
	// "sheets.pending", "sheets.success", "sheets.error", "sheets.partial_success",
	// "sheets.cancelled", "unmapped_event".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers sent with each webhook request (e.g. auth tokens)
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	paramObj
}

func (r BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaBatchListParams struct {
	// Filter by directory ID
	DirectoryID    param.Opt[string] `query:"directory_id,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Maximum number of jobs to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of jobs to skip for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by job type (PARSE, EXTRACT, CLASSIFY)
	//
	// Any of "parse", "extract", "classify".
	JobType BetaBatchListParamsJobType `query:"job_type,omitzero" json:"-"`
	// Filter by job status (PENDING, RUNNING, COMPLETED, FAILED, CANCELLED)
	//
	// Any of "pending", "running", "dispatched", "completed", "failed", "cancelled".
	Status BetaBatchListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaBatchListParams]'s query parameters as `url.Values`.
func (r BetaBatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job type (PARSE, EXTRACT, CLASSIFY)
type BetaBatchListParamsJobType string

const (
	BetaBatchListParamsJobTypeParse    BetaBatchListParamsJobType = "parse"
	BetaBatchListParamsJobTypeExtract  BetaBatchListParamsJobType = "extract"
	BetaBatchListParamsJobTypeClassify BetaBatchListParamsJobType = "classify"
)

// Filter by job status (PENDING, RUNNING, COMPLETED, FAILED, CANCELLED)
type BetaBatchListParamsStatus string

const (
	BetaBatchListParamsStatusPending    BetaBatchListParamsStatus = "pending"
	BetaBatchListParamsStatusRunning    BetaBatchListParamsStatus = "running"
	BetaBatchListParamsStatusDispatched BetaBatchListParamsStatus = "dispatched"
	BetaBatchListParamsStatusCompleted  BetaBatchListParamsStatus = "completed"
	BetaBatchListParamsStatusFailed     BetaBatchListParamsStatus = "failed"
	BetaBatchListParamsStatusCancelled  BetaBatchListParamsStatus = "cancelled"
)

type BetaBatchCancelParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Optional reason for cancelling the job
	Reason            param.Opt[string] `json:"reason,omitzero"`
	TemporalNamespace param.Opt[string] `header:"temporal-namespace,omitzero" json:"-"`
	paramObj
}

func (r BetaBatchCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaBatchCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaBatchCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaBatchCancelParams]'s query parameters as `url.Values`.
func (r BetaBatchCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaBatchGetStatusParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaBatchGetStatusParams]'s query parameters as
// `url.Values`.
func (r BetaBatchGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
