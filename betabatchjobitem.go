// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"encoding/json"
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

// BetaBatchJobItemService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaBatchJobItemService] method instead.
type BetaBatchJobItemService struct {
	options []option.RequestOption
}

// NewBetaBatchJobItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaBatchJobItemService(opts ...option.RequestOption) (r BetaBatchJobItemService) {
	r = BetaBatchJobItemService{}
	r.options = opts
	return
}

// List items in a batch job with optional status filtering.
//
// Useful for finding failed items, viewing completed items, or debugging
// processing issues.
func (r *BetaBatchJobItemService) List(ctx context.Context, jobID string, query BetaBatchJobItemListParams, opts ...option.RequestOption) (res *pagination.PaginatedBatchItems[BetaBatchJobItemListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/batch-processing/%s/items", url.PathEscape(jobID))
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

// List items in a batch job with optional status filtering.
//
// Useful for finding failed items, viewing completed items, or debugging
// processing issues.
func (r *BetaBatchJobItemService) ListAutoPaging(ctx context.Context, jobID string, query BetaBatchJobItemListParams, opts ...option.RequestOption) *pagination.PaginatedBatchItemsAutoPager[BetaBatchJobItemListResponse] {
	return pagination.NewPaginatedBatchItemsAutoPager(r.List(ctx, jobID, query, opts...))
}

// Get all processing results for a specific item.
//
// Returns the complete processing history for an item including what operations
// were performed, parameters used, and where outputs are stored. Optionally filter
// by `job_type`.
func (r *BetaBatchJobItemService) GetProcessingResults(ctx context.Context, itemID string, query BetaBatchJobItemGetProcessingResultsParams, opts ...option.RequestOption) (res *BetaBatchJobItemGetProcessingResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/batch-processing/items/%s/processing-results", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Detailed information about an item in a batch job.
type BetaBatchJobItemListResponse struct {
	// ID of the item
	ItemID string `json:"item_id" api:"required"`
	// Name of the item
	ItemName string `json:"item_name" api:"required"`
	// Processing status of this item
	//
	// Any of "pending", "processing", "completed", "failed", "skipped", "cancelled".
	Status BetaBatchJobItemListResponseStatus `json:"status" api:"required"`
	// When processing completed for this item
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	EffectiveAt time.Time `json:"effective_at" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Job ID for the underlying processing job (links to parse/extract job results)
	JobID string `json:"job_id" api:"nullable"`
	// The job record ID associated with this status, if any.
	JobRecordID string `json:"job_record_id" api:"nullable"`
	// Reason item was skipped (e.g., 'already_processed', 'size_limit_exceeded')
	SkipReason string `json:"skip_reason" api:"nullable"`
	// When processing started for this item
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID       respjson.Field
		ItemName     respjson.Field
		Status       respjson.Field
		CompletedAt  respjson.Field
		EffectiveAt  respjson.Field
		ErrorMessage respjson.Field
		JobID        respjson.Field
		JobRecordID  respjson.Field
		SkipReason   respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchJobItemListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Processing status of this item
type BetaBatchJobItemListResponseStatus string

const (
	BetaBatchJobItemListResponseStatusPending    BetaBatchJobItemListResponseStatus = "pending"
	BetaBatchJobItemListResponseStatusProcessing BetaBatchJobItemListResponseStatus = "processing"
	BetaBatchJobItemListResponseStatusCompleted  BetaBatchJobItemListResponseStatus = "completed"
	BetaBatchJobItemListResponseStatusFailed     BetaBatchJobItemListResponseStatus = "failed"
	BetaBatchJobItemListResponseStatusSkipped    BetaBatchJobItemListResponseStatus = "skipped"
	BetaBatchJobItemListResponseStatusCancelled  BetaBatchJobItemListResponseStatus = "cancelled"
)

// Response containing all processing results for an item.
type BetaBatchJobItemGetProcessingResultsResponse struct {
	// ID of the source item
	ItemID string `json:"item_id" api:"required"`
	// Name of the source item
	ItemName string `json:"item_name" api:"required"`
	// List of all processing operations performed on this item
	ProcessingResults []BetaBatchJobItemGetProcessingResultsResponseProcessingResult `json:"processing_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID            respjson.Field
		ItemName          respjson.Field
		ProcessingResults respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemGetProcessingResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaBatchJobItemGetProcessingResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A processing result with lineage information.
type BetaBatchJobItemGetProcessingResultsResponseProcessingResult struct {
	// Source item that was processed
	ItemID string `json:"item_id" api:"required"`
	// Job configuration used for processing
	JobConfig BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion `json:"job_config" api:"required"`
	// Type of processing performed
	//
	// Any of "parse", "extract", "classify".
	JobType string `json:"job_type" api:"required"`
	// Location of the processing output
	OutputS3Path string `json:"output_s3_path" api:"required"`
	// Content hash of the job configuration for dedup
	ParametersHash string `json:"parameters_hash" api:"required"`
	// When this processing occurred
	ProcessedAt time.Time `json:"processed_at" api:"required" format:"date-time"`
	// Unique identifier for this result
	ResultID string `json:"result_id" api:"required"`
	// Metadata about processing output.
	//
	// Currently empty - will be populated with job-type-specific metadata fields in
	// the future.
	OutputMetadata any `json:"output_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID         respjson.Field
		JobConfig      respjson.Field
		JobType        respjson.Field
		OutputS3Path   respjson.Field
		ParametersHash respjson.Field
		ProcessedAt    respjson.Field
		ResultID       respjson.Field
		OutputMetadata respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemGetProcessingResultsResponseProcessingResult) RawJSON() string {
	return r.JSON.raw
}
func (r *BetaBatchJobItemGetProcessingResultsResponseProcessingResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion
// contains all possible properties and values from
// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate],
// [ClassifyJob].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion struct {
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	CorrelationID string `json:"correlation_id"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	JobName string `json:"job_name"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	Parameters BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParameters `json:"parameters"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	ParentJobExecutionID string `json:"parent_job_execution_id"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	Partitions map[string]string `json:"partitions"`
	ProjectID  string            `json:"project_id"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	// This field is from variant
	// [BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant [ClassifyJob].
	ID string `json:"id"`
	// This field is from variant [ClassifyJob].
	Rules []ClassifierRule `json:"rules"`
	// This field is from variant [ClassifyJob].
	Status StatusEnum `json:"status"`
	// This field is from variant [ClassifyJob].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ClassifyJob].
	EffectiveAt time.Time `json:"effective_at"`
	// This field is from variant [ClassifyJob].
	ErrorMessage string `json:"error_message"`
	// This field is from variant [ClassifyJob].
	JobRecordID string `json:"job_record_id"`
	// This field is from variant [ClassifyJob].
	Mode ClassifyJobMode `json:"mode"`
	// This field is from variant [ClassifyJob].
	ParsingConfiguration ClassifyParsingConfiguration `json:"parsing_configuration"`
	// This field is from variant [ClassifyJob].
	UpdatedAt time.Time `json:"updated_at"`
	JSON      struct {
		CorrelationID        respjson.Field
		JobName              respjson.Field
		Parameters           respjson.Field
		ParentJobExecutionID respjson.Field
		Partitions           respjson.Field
		ProjectID            respjson.Field
		SessionID            respjson.Field
		UserID               respjson.Field
		WebhookURL           respjson.Field
		ID                   respjson.Field
		Rules                respjson.Field
		Status               respjson.Field
		CreatedAt            respjson.Field
		EffectiveAt          respjson.Field
		ErrorMessage         respjson.Field
		JobRecordID          respjson.Field
		Mode                 respjson.Field
		ParsingConfiguration respjson.Field
		UpdatedAt            respjson.Field
		raw                  string
	} `json:"-"`
}

func (u BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion) AsBatchParseJobRecordCreate() (v BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion) AsClassifyJob() (v ClassifyJob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate struct {
	// The correlation ID for this job. Used for tracking the job across services.
	CorrelationID string `json:"correlation_id" api:"nullable" format:"uuid"`
	// Any of "parse_raw_file_job".
	JobName string `json:"job_name"`
	// Generic parse job configuration for batch processing.
	//
	// This model contains the parsing configuration that applies to all files in a
	// batch, but excludes file-specific fields like file_name, file_id, etc. Those
	// file-specific fields are populated from DirectoryFile data when creating
	// individual ParseJobRecordCreate instances for each file.
	//
	// The fields in this model should be generic settings that apply uniformly to all
	// files being processed in the batch.
	Parameters BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParameters `json:"parameters" api:"nullable"`
	// The ID of the parent job execution.
	ParentJobExecutionID string `json:"parent_job_execution_id" api:"nullable" format:"uuid"`
	// The partitions for this execution. Used for determining where to save job
	// output.
	Partitions map[string]string `json:"partitions" format:"uuid"`
	// The ID of the project this job belongs to.
	ProjectID string `json:"project_id" api:"nullable" format:"uuid"`
	// The upstream request ID that created this job. Used for tracking the job across
	// services.
	SessionID string `json:"session_id" api:"nullable" format:"uuid"`
	// The ID of the user that created this job
	UserID string `json:"user_id" api:"nullable"`
	// The URL that needs to be called at the end of the parsing job.
	WebhookURL string `json:"webhook_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorrelationID        respjson.Field
		JobName              respjson.Field
		Parameters           respjson.Field
		ParentJobExecutionID respjson.Field
		Partitions           respjson.Field
		ProjectID            respjson.Field
		SessionID            respjson.Field
		UserID               respjson.Field
		WebhookURL           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate) RawJSON() string {
	return r.JSON.raw
}
func (r *BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParameters struct {
	AdaptiveLongTable                 bool    `json:"adaptive_long_table" api:"nullable"`
	AggressiveTableExtraction         bool    `json:"aggressive_table_extraction" api:"nullable"`
	AnnotateLinks                     bool    `json:"annotate_links" api:"nullable"`
	AutoMode                          bool    `json:"auto_mode" api:"nullable"`
	AutoModeConfigurationJson         string  `json:"auto_mode_configuration_json" api:"nullable"`
	AutoModeTriggerOnImageInPage      bool    `json:"auto_mode_trigger_on_image_in_page" api:"nullable"`
	AutoModeTriggerOnRegexpInPage     string  `json:"auto_mode_trigger_on_regexp_in_page" api:"nullable"`
	AutoModeTriggerOnTableInPage      bool    `json:"auto_mode_trigger_on_table_in_page" api:"nullable"`
	AutoModeTriggerOnTextInPage       string  `json:"auto_mode_trigger_on_text_in_page" api:"nullable"`
	AzureOpenAIAPIVersion             string  `json:"azure_openai_api_version" api:"nullable"`
	AzureOpenAIDeploymentName         string  `json:"azure_openai_deployment_name" api:"nullable"`
	AzureOpenAIEndpoint               string  `json:"azure_openai_endpoint" api:"nullable"`
	AzureOpenAIKey                    string  `json:"azure_openai_key" api:"nullable"`
	BboxBottom                        float64 `json:"bbox_bottom" api:"nullable"`
	BboxLeft                          float64 `json:"bbox_left" api:"nullable"`
	BboxRight                         float64 `json:"bbox_right" api:"nullable"`
	BboxTop                           float64 `json:"bbox_top" api:"nullable"`
	BoundingBox                       string  `json:"bounding_box" api:"nullable"`
	CompactMarkdownTable              bool    `json:"compact_markdown_table" api:"nullable"`
	ComplementalFormattingInstruction string  `json:"complemental_formatting_instruction" api:"nullable"`
	ContentGuidelineInstruction       string  `json:"content_guideline_instruction" api:"nullable"`
	ContinuousMode                    bool    `json:"continuous_mode" api:"nullable"`
	// The custom metadata to attach to the documents.
	CustomMetadata                           map[string]any `json:"custom_metadata" api:"nullable"`
	DisableImageExtraction                   bool           `json:"disable_image_extraction" api:"nullable"`
	DisableOcr                               bool           `json:"disable_ocr" api:"nullable"`
	DisableReconstruction                    bool           `json:"disable_reconstruction" api:"nullable"`
	DoNotCache                               bool           `json:"do_not_cache" api:"nullable"`
	DoNotUnrollColumns                       bool           `json:"do_not_unroll_columns" api:"nullable"`
	EnableCostOptimizer                      bool           `json:"enable_cost_optimizer" api:"nullable"`
	ExtractCharts                            bool           `json:"extract_charts" api:"nullable"`
	ExtractLayout                            bool           `json:"extract_layout" api:"nullable"`
	ExtractPrintedPageNumber                 bool           `json:"extract_printed_page_number" api:"nullable"`
	FastMode                                 bool           `json:"fast_mode" api:"nullable"`
	FormattingInstruction                    string         `json:"formatting_instruction" api:"nullable"`
	Gpt4oAPIKey                              string         `json:"gpt4o_api_key" api:"nullable"`
	Gpt4oMode                                bool           `json:"gpt4o_mode" api:"nullable"`
	GuessXlsxSheetName                       bool           `json:"guess_xlsx_sheet_name" api:"nullable"`
	HideFooters                              bool           `json:"hide_footers" api:"nullable"`
	HideHeaders                              bool           `json:"hide_headers" api:"nullable"`
	HighResOcr                               bool           `json:"high_res_ocr" api:"nullable"`
	HTMLMakeAllElementsVisible               bool           `json:"html_make_all_elements_visible" api:"nullable"`
	HTMLRemoveFixedElements                  bool           `json:"html_remove_fixed_elements" api:"nullable"`
	HTMLRemoveNavigationElements             bool           `json:"html_remove_navigation_elements" api:"nullable"`
	HTTPProxy                                string         `json:"http_proxy" api:"nullable"`
	IgnoreDocumentElementsForLayoutDetection bool           `json:"ignore_document_elements_for_layout_detection" api:"nullable"`
	// Any of "screenshot", "embedded", "layout".
	ImagesToSave           []string `json:"images_to_save" api:"nullable"`
	InlineImagesInMarkdown bool     `json:"inline_images_in_markdown" api:"nullable"`
	InputS3Path            string   `json:"input_s3_path" api:"nullable"`
	// The region for the input S3 bucket.
	InputS3Region                       string  `json:"input_s3_region" api:"nullable"`
	InputURL                            string  `json:"input_url" api:"nullable"`
	InternalIsScreenshotJob             bool    `json:"internal_is_screenshot_job" api:"nullable"`
	InvalidateCache                     bool    `json:"invalidate_cache" api:"nullable"`
	IsFormattingInstruction             bool    `json:"is_formatting_instruction" api:"nullable"`
	JobTimeoutExtraTimePerPageInSeconds float64 `json:"job_timeout_extra_time_per_page_in_seconds" api:"nullable"`
	JobTimeoutInSeconds                 float64 `json:"job_timeout_in_seconds" api:"nullable"`
	KeepPageSeparatorWhenMergingTables  bool    `json:"keep_page_separator_when_merging_tables" api:"nullable"`
	// The language.
	Lang                                  string             `json:"lang"`
	Languages                             []ParsingLanguages `json:"languages"`
	LayoutAware                           bool               `json:"layout_aware" api:"nullable"`
	LineLevelBoundingBox                  bool               `json:"line_level_bounding_box" api:"nullable"`
	MarkdownTableMultilineHeaderSeparator string             `json:"markdown_table_multiline_header_separator" api:"nullable"`
	MaxPages                              int64              `json:"max_pages" api:"nullable"`
	MaxPagesEnforced                      int64              `json:"max_pages_enforced" api:"nullable"`
	MergeTablesAcrossPagesInMarkdown      bool               `json:"merge_tables_across_pages_in_markdown" api:"nullable"`
	Model                                 string             `json:"model" api:"nullable"`
	OutlinedTableExtraction               bool               `json:"outlined_table_extraction" api:"nullable"`
	OutputPdfOfDocument                   bool               `json:"output_pdf_of_document" api:"nullable"`
	// If specified, llamaParse will save the output to the specified path. All output
	// file will use this 'prefix' should be a valid s3:// url
	OutputS3PathPrefix string `json:"output_s3_path_prefix" api:"nullable"`
	// The region for the output S3 bucket.
	OutputS3Region     string `json:"output_s3_region" api:"nullable"`
	OutputTablesAsHTML bool   `json:"output_tables_as_HTML" api:"nullable"`
	// The output bucket.
	OutputBucket       string  `json:"outputBucket" api:"nullable"`
	PageErrorTolerance float64 `json:"page_error_tolerance" api:"nullable"`
	PageFooterPrefix   string  `json:"page_footer_prefix" api:"nullable"`
	PageFooterSuffix   string  `json:"page_footer_suffix" api:"nullable"`
	PageHeaderPrefix   string  `json:"page_header_prefix" api:"nullable"`
	PageHeaderSuffix   string  `json:"page_header_suffix" api:"nullable"`
	PagePrefix         string  `json:"page_prefix" api:"nullable"`
	PageSeparator      string  `json:"page_separator" api:"nullable"`
	PageSuffix         string  `json:"page_suffix" api:"nullable"`
	// Enum for representing the mode of parsing to be used.
	//
	// Any of "parse_page_without_llm", "parse_page_with_llm", "parse_page_with_lvm",
	// "parse_page_with_agent", "parse_page_with_layout_agent",
	// "parse_document_with_llm", "parse_document_with_lvm",
	// "parse_document_with_agent".
	ParseMode          ParsingMode `json:"parse_mode" api:"nullable"`
	ParsingInstruction string      `json:"parsing_instruction" api:"nullable"`
	// The pipeline ID.
	PipelineID                         string `json:"pipeline_id" api:"nullable"`
	PreciseBoundingBox                 bool   `json:"precise_bounding_box" api:"nullable"`
	PremiumMode                        bool   `json:"premium_mode" api:"nullable"`
	PresentationOutOfBoundsContent     bool   `json:"presentation_out_of_bounds_content" api:"nullable"`
	PresentationSkipEmbeddedData       bool   `json:"presentation_skip_embedded_data" api:"nullable"`
	PreserveLayoutAlignmentAcrossPages bool   `json:"preserve_layout_alignment_across_pages" api:"nullable"`
	PreserveVerySmallText              bool   `json:"preserve_very_small_text" api:"nullable"`
	Preset                             string `json:"preset" api:"nullable"`
	// The priority for the request. This field may be ignored or overwritten depending
	// on the organization tier.
	//
	// Any of "low", "medium", "high", "critical".
	Priority         string `json:"priority" api:"nullable"`
	ProjectID        string `json:"project_id" api:"nullable"`
	RemoveHiddenText bool   `json:"remove_hidden_text" api:"nullable"`
	// Enum for representing the different available page error handling modes.
	//
	// Any of "raw_text", "blank_page", "error_message".
	ReplaceFailedPageMode                   FailPageMode `json:"replace_failed_page_mode" api:"nullable"`
	ReplaceFailedPageWithErrorMessagePrefix string       `json:"replace_failed_page_with_error_message_prefix" api:"nullable"`
	ReplaceFailedPageWithErrorMessageSuffix string       `json:"replace_failed_page_with_error_message_suffix" api:"nullable"`
	// The resource info about the file
	ResourceInfo                       map[string]any `json:"resource_info" api:"nullable"`
	SaveImages                         bool           `json:"save_images" api:"nullable"`
	SkipDiagonalText                   bool           `json:"skip_diagonal_text" api:"nullable"`
	SpecializedChartParsingAgentic     bool           `json:"specialized_chart_parsing_agentic" api:"nullable"`
	SpecializedChartParsingEfficient   bool           `json:"specialized_chart_parsing_efficient" api:"nullable"`
	SpecializedChartParsingPlus        bool           `json:"specialized_chart_parsing_plus" api:"nullable"`
	SpecializedImageParsing            bool           `json:"specialized_image_parsing" api:"nullable"`
	SpreadsheetExtractSubTables        bool           `json:"spreadsheet_extract_sub_tables" api:"nullable"`
	SpreadsheetForceFormulaComputation bool           `json:"spreadsheet_force_formula_computation" api:"nullable"`
	SpreadsheetIncludeHiddenSheets     bool           `json:"spreadsheet_include_hidden_sheets" api:"nullable"`
	StrictModeBuggyFont                bool           `json:"strict_mode_buggy_font" api:"nullable"`
	StrictModeImageExtraction          bool           `json:"strict_mode_image_extraction" api:"nullable"`
	StrictModeImageOcr                 bool           `json:"strict_mode_image_ocr" api:"nullable"`
	StrictModeReconstruction           bool           `json:"strict_mode_reconstruction" api:"nullable"`
	StructuredOutput                   bool           `json:"structured_output" api:"nullable"`
	StructuredOutputJsonSchema         string         `json:"structured_output_json_schema" api:"nullable"`
	StructuredOutputJsonSchemaName     string         `json:"structured_output_json_schema_name" api:"nullable"`
	SystemPrompt                       string         `json:"system_prompt" api:"nullable"`
	SystemPromptAppend                 string         `json:"system_prompt_append" api:"nullable"`
	TakeScreenshot                     bool           `json:"take_screenshot" api:"nullable"`
	TargetPages                        string         `json:"target_pages" api:"nullable"`
	Tier                               string         `json:"tier" api:"nullable"`
	// Any of "parse".
	Type                      string `json:"type"`
	UseVendorMultimodalModel  bool   `json:"use_vendor_multimodal_model" api:"nullable"`
	UserPrompt                string `json:"user_prompt" api:"nullable"`
	VendorMultimodalAPIKey    string `json:"vendor_multimodal_api_key" api:"nullable"`
	VendorMultimodalModelName string `json:"vendor_multimodal_model_name" api:"nullable"`
	Version                   string `json:"version" api:"nullable"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration `json:"webhook_configurations" api:"nullable"`
	WebhookURL            string                                                                                                                         `json:"webhook_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveLongTable                        respjson.Field
		AggressiveTableExtraction                respjson.Field
		AnnotateLinks                            respjson.Field
		AutoMode                                 respjson.Field
		AutoModeConfigurationJson                respjson.Field
		AutoModeTriggerOnImageInPage             respjson.Field
		AutoModeTriggerOnRegexpInPage            respjson.Field
		AutoModeTriggerOnTableInPage             respjson.Field
		AutoModeTriggerOnTextInPage              respjson.Field
		AzureOpenAIAPIVersion                    respjson.Field
		AzureOpenAIDeploymentName                respjson.Field
		AzureOpenAIEndpoint                      respjson.Field
		AzureOpenAIKey                           respjson.Field
		BboxBottom                               respjson.Field
		BboxLeft                                 respjson.Field
		BboxRight                                respjson.Field
		BboxTop                                  respjson.Field
		BoundingBox                              respjson.Field
		CompactMarkdownTable                     respjson.Field
		ComplementalFormattingInstruction        respjson.Field
		ContentGuidelineInstruction              respjson.Field
		ContinuousMode                           respjson.Field
		CustomMetadata                           respjson.Field
		DisableImageExtraction                   respjson.Field
		DisableOcr                               respjson.Field
		DisableReconstruction                    respjson.Field
		DoNotCache                               respjson.Field
		DoNotUnrollColumns                       respjson.Field
		EnableCostOptimizer                      respjson.Field
		ExtractCharts                            respjson.Field
		ExtractLayout                            respjson.Field
		ExtractPrintedPageNumber                 respjson.Field
		FastMode                                 respjson.Field
		FormattingInstruction                    respjson.Field
		Gpt4oAPIKey                              respjson.Field
		Gpt4oMode                                respjson.Field
		GuessXlsxSheetName                       respjson.Field
		HideFooters                              respjson.Field
		HideHeaders                              respjson.Field
		HighResOcr                               respjson.Field
		HTMLMakeAllElementsVisible               respjson.Field
		HTMLRemoveFixedElements                  respjson.Field
		HTMLRemoveNavigationElements             respjson.Field
		HTTPProxy                                respjson.Field
		IgnoreDocumentElementsForLayoutDetection respjson.Field
		ImagesToSave                             respjson.Field
		InlineImagesInMarkdown                   respjson.Field
		InputS3Path                              respjson.Field
		InputS3Region                            respjson.Field
		InputURL                                 respjson.Field
		InternalIsScreenshotJob                  respjson.Field
		InvalidateCache                          respjson.Field
		IsFormattingInstruction                  respjson.Field
		JobTimeoutExtraTimePerPageInSeconds      respjson.Field
		JobTimeoutInSeconds                      respjson.Field
		KeepPageSeparatorWhenMergingTables       respjson.Field
		Lang                                     respjson.Field
		Languages                                respjson.Field
		LayoutAware                              respjson.Field
		LineLevelBoundingBox                     respjson.Field
		MarkdownTableMultilineHeaderSeparator    respjson.Field
		MaxPages                                 respjson.Field
		MaxPagesEnforced                         respjson.Field
		MergeTablesAcrossPagesInMarkdown         respjson.Field
		Model                                    respjson.Field
		OutlinedTableExtraction                  respjson.Field
		OutputPdfOfDocument                      respjson.Field
		OutputS3PathPrefix                       respjson.Field
		OutputS3Region                           respjson.Field
		OutputTablesAsHTML                       respjson.Field
		OutputBucket                             respjson.Field
		PageErrorTolerance                       respjson.Field
		PageFooterPrefix                         respjson.Field
		PageFooterSuffix                         respjson.Field
		PageHeaderPrefix                         respjson.Field
		PageHeaderSuffix                         respjson.Field
		PagePrefix                               respjson.Field
		PageSeparator                            respjson.Field
		PageSuffix                               respjson.Field
		ParseMode                                respjson.Field
		ParsingInstruction                       respjson.Field
		PipelineID                               respjson.Field
		PreciseBoundingBox                       respjson.Field
		PremiumMode                              respjson.Field
		PresentationOutOfBoundsContent           respjson.Field
		PresentationSkipEmbeddedData             respjson.Field
		PreserveLayoutAlignmentAcrossPages       respjson.Field
		PreserveVerySmallText                    respjson.Field
		Preset                                   respjson.Field
		Priority                                 respjson.Field
		ProjectID                                respjson.Field
		RemoveHiddenText                         respjson.Field
		ReplaceFailedPageMode                    respjson.Field
		ReplaceFailedPageWithErrorMessagePrefix  respjson.Field
		ReplaceFailedPageWithErrorMessageSuffix  respjson.Field
		ResourceInfo                             respjson.Field
		SaveImages                               respjson.Field
		SkipDiagonalText                         respjson.Field
		SpecializedChartParsingAgentic           respjson.Field
		SpecializedChartParsingEfficient         respjson.Field
		SpecializedChartParsingPlus              respjson.Field
		SpecializedImageParsing                  respjson.Field
		SpreadsheetExtractSubTables              respjson.Field
		SpreadsheetForceFormulaComputation       respjson.Field
		SpreadsheetIncludeHiddenSheets           respjson.Field
		StrictModeBuggyFont                      respjson.Field
		StrictModeImageExtraction                respjson.Field
		StrictModeImageOcr                       respjson.Field
		StrictModeReconstruction                 respjson.Field
		StructuredOutput                         respjson.Field
		StructuredOutputJsonSchema               respjson.Field
		StructuredOutputJsonSchemaName           respjson.Field
		SystemPrompt                             respjson.Field
		SystemPromptAppend                       respjson.Field
		TakeScreenshot                           respjson.Field
		TargetPages                              respjson.Field
		Tier                                     respjson.Field
		Type                                     respjson.Field
		UseVendorMultimodalModel                 respjson.Field
		UserPrompt                               respjson.Field
		VendorMultimodalAPIKey                   respjson.Field
		VendorMultimodalModelName                respjson.Field
		Version                                  respjson.Field
		WebhookConfigurations                    respjson.Field
		WebhookURL                               respjson.Field
		ExtraFields                              map[string]respjson.Field
		raw                                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a single outbound webhook endpoint.
type BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration struct {
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
	WebhookEvents []string `json:"webhook_events" api:"nullable"`
	// Custom HTTP headers sent with each webhook request (e.g. auth tokens)
	WebhookHeaders map[string]string `json:"webhook_headers" api:"nullable"`
	// Response format sent to the webhook: 'string' (default) or 'json'
	WebhookOutputFormat string `json:"webhook_output_format" api:"nullable"`
	// URL to receive webhook POST notifications
	WebhookURL string `json:"webhook_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WebhookEvents       respjson.Field
		WebhookHeaders      respjson.Field
		WebhookOutputFormat respjson.Field
		WebhookURL          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *BetaBatchJobItemGetProcessingResultsResponseProcessingResultJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaBatchJobItemListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter items by status
	//
	// Any of "pending", "processing", "completed", "failed", "skipped", "cancelled".
	Status BetaBatchJobItemListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaBatchJobItemListParams]'s query parameters as
// `url.Values`.
func (r BetaBatchJobItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter items by status
type BetaBatchJobItemListParamsStatus string

const (
	BetaBatchJobItemListParamsStatusPending    BetaBatchJobItemListParamsStatus = "pending"
	BetaBatchJobItemListParamsStatusProcessing BetaBatchJobItemListParamsStatus = "processing"
	BetaBatchJobItemListParamsStatusCompleted  BetaBatchJobItemListParamsStatus = "completed"
	BetaBatchJobItemListParamsStatusFailed     BetaBatchJobItemListParamsStatus = "failed"
	BetaBatchJobItemListParamsStatusSkipped    BetaBatchJobItemListParamsStatus = "skipped"
	BetaBatchJobItemListParamsStatusCancelled  BetaBatchJobItemListParamsStatus = "cancelled"
)

type BetaBatchJobItemGetProcessingResultsParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter results by job type
	//
	// Any of "parse", "extract", "classify".
	JobType BetaBatchJobItemGetProcessingResultsParamsJobType `query:"job_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaBatchJobItemGetProcessingResultsParams]'s query
// parameters as `url.Values`.
func (r BetaBatchJobItemGetProcessingResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter results by job type
type BetaBatchJobItemGetProcessingResultsParamsJobType string

const (
	BetaBatchJobItemGetProcessingResultsParamsJobTypeParse    BetaBatchJobItemGetProcessingResultsParamsJobType = "parse"
	BetaBatchJobItemGetProcessingResultsParamsJobTypeExtract  BetaBatchJobItemGetProcessingResultsParamsJobType = "extract"
	BetaBatchJobItemGetProcessingResultsParamsJobTypeClassify BetaBatchJobItemGetProcessingResultsParamsJobType = "classify"
)
