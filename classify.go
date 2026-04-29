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
	shimjson "github.com/stainless-sdks/llamacloud-prod-go/internal/encoding/json"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// ClassifyService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClassifyService] method instead.
type ClassifyService struct {
	options []option.RequestOption
}

// NewClassifyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClassifyService(opts ...option.RequestOption) (r ClassifyService) {
	r = ClassifyService{}
	r.options = opts
	return
}

// Create a classify job.
//
// Classifies a document against a set of rules. Set `file_input` to a file ID
// (`dfl-...`) or parse job ID (`pjb-...`), and provide either inline
// `configuration` with rules or a `configuration_id` referencing a saved preset.
//
// Each rule has a `type` (the label to assign) and a `description` (natural
// language criteria). The classifier returns the best matching rule with a
// confidence score.
//
// The job runs asynchronously. Poll `GET /classify/{job_id}` to check status and
// retrieve results.
func (r *ClassifyService) New(ctx context.Context, params ClassifyNewParams, opts ...option.RequestOption) (res *ClassifyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/classify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List classify jobs with optional filtering and pagination.
//
// Filter by `status`, `configuration_id`, specific `job_ids`, or creation date
// range.
func (r *ClassifyService) List(ctx context.Context, query ClassifyListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ClassifyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/classify"
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

// List classify jobs with optional filtering and pagination.
//
// Filter by `status`, `configuration_id`, specific `job_ids`, or creation date
// range.
func (r *ClassifyService) ListAutoPaging(ctx context.Context, query ClassifyListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ClassifyListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a classify job by ID.
//
// Returns the job status, configuration, and classify result when complete. The
// result includes the matched document type, confidence score, and reasoning.
func (r *ClassifyService) Get(ctx context.Context, jobID string, query ClassifyGetParams, opts ...option.RequestOption) (res *ClassifyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/classify/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Configuration for a classify job.
type ClassifyConfiguration struct {
	// Classify rules to evaluate against the document (at least one required)
	Rules []ClassifyConfigurationRule `json:"rules" api:"required"`
	// Classify execution mode
	//
	// Any of "FAST".
	Mode ClassifyConfigurationMode `json:"mode"`
	// Parsing configuration for classify jobs.
	ParsingConfiguration ClassifyConfigurationParsingConfiguration `json:"parsing_configuration" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rules                respjson.Field
		Mode                 respjson.Field
		ParsingConfiguration respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ClassifyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassifyConfiguration to a ClassifyConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassifyConfigurationParam.Overrides()
func (r ClassifyConfiguration) ToParam() ClassifyConfigurationParam {
	return param.Override[ClassifyConfigurationParam](json.RawMessage(r.RawJSON()))
}

// A rule for classifying documents.
type ClassifyConfigurationRule struct {
	// Natural language criteria for matching this rule
	Description string `json:"description" api:"required"`
	// Document type to assign when rule matches
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyConfigurationRule) RawJSON() string { return r.JSON.raw }
func (r *ClassifyConfigurationRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classify execution mode
type ClassifyConfigurationMode string

const (
	ClassifyConfigurationModeFast ClassifyConfigurationMode = "FAST"
)

// Parsing configuration for classify jobs.
type ClassifyConfigurationParsingConfiguration struct {
	// ISO 639-1 language code for the document
	Lang string `json:"lang"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages string `json:"target_pages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lang        respjson.Field
		MaxPages    respjson.Field
		TargetPages respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyConfigurationParsingConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ClassifyConfigurationParsingConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a classify job.
//
// The property Rules is required.
type ClassifyConfigurationParam struct {
	// Classify rules to evaluate against the document (at least one required)
	Rules []ClassifyConfigurationRuleParam `json:"rules,omitzero" api:"required"`
	// Parsing configuration for classify jobs.
	ParsingConfiguration ClassifyConfigurationParsingConfigurationParam `json:"parsing_configuration,omitzero"`
	// Classify execution mode
	//
	// Any of "FAST".
	Mode ClassifyConfigurationMode `json:"mode,omitzero"`
	paramObj
}

func (r ClassifyConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule for classifying documents.
//
// The properties Description, Type are required.
type ClassifyConfigurationRuleParam struct {
	// Natural language criteria for matching this rule
	Description string `json:"description" api:"required"`
	// Document type to assign when rule matches
	Type string `json:"type" api:"required"`
	paramObj
}

func (r ClassifyConfigurationRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyConfigurationRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyConfigurationRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration for classify jobs.
type ClassifyConfigurationParsingConfigurationParam struct {
	// Maximum number of pages to process. Omit for no limit.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	// ISO 639-1 language code for the document
	Lang param.Opt[string] `json:"lang,omitzero"`
	paramObj
}

func (r ClassifyConfigurationParsingConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyConfigurationParsingConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyConfigurationParsingConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create a classify job.
type ClassifyCreateRequestParam struct {
	// Saved configuration ID
	ConfigurationID param.Opt[string] `json:"configuration_id,omitzero"`
	// Deprecated: use file_input instead
	//
	// Deprecated: deprecated
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// File ID or parse job ID to classify
	FileInput param.Opt[string] `json:"file_input,omitzero"`
	// Deprecated: use file_input instead
	//
	// Deprecated: deprecated
	ParseJobID param.Opt[string] `json:"parse_job_id,omitzero"`
	// Idempotency key scoped to the project
	TransactionID param.Opt[string] `json:"transaction_id,omitzero"`
	// Configuration for a classify job.
	Configuration ClassifyConfigurationParam `json:"configuration,omitzero"`
	paramObj
}

func (r ClassifyCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of classifying a document.
type ClassifyResult struct {
	// Confidence score between 0.0 and 1.0
	Confidence float64 `json:"confidence" api:"required"`
	// Why the document matched (or didn't match) the returned rule
	Reasoning string `json:"reasoning" api:"required"`
	// Matched rule type, or null if no rule matched
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Reasoning   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyResult) RawJSON() string { return r.JSON.raw }
func (r *ClassifyResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a classify job.
type ClassifyNewResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Classify configuration used for this job
	Configuration ClassifyConfiguration `json:"configuration" api:"required"`
	// Whether the input was a file or parse job (FILE or PARSE_JOB)
	//
	// Any of "url", "file_id", "parse_job_id".
	DocumentInputType ClassifyNewResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file or parse job
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED".
	Status ClassifyNewResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Product configuration ID
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Associated parse job ID
	ParseJobID string `json:"parse_job_id" api:"nullable"`
	// Result of classifying a document.
	Result ClassifyResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		ParseJobID        respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ClassifyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job (FILE or PARSE_JOB)
type ClassifyNewResponseDocumentInputType string

const (
	ClassifyNewResponseDocumentInputTypeURL        ClassifyNewResponseDocumentInputType = "url"
	ClassifyNewResponseDocumentInputTypeFileID     ClassifyNewResponseDocumentInputType = "file_id"
	ClassifyNewResponseDocumentInputTypeParseJobID ClassifyNewResponseDocumentInputType = "parse_job_id"
)

// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
type ClassifyNewResponseStatus string

const (
	ClassifyNewResponseStatusPending   ClassifyNewResponseStatus = "PENDING"
	ClassifyNewResponseStatusRunning   ClassifyNewResponseStatus = "RUNNING"
	ClassifyNewResponseStatusCompleted ClassifyNewResponseStatus = "COMPLETED"
	ClassifyNewResponseStatusFailed    ClassifyNewResponseStatus = "FAILED"
)

// Response for a classify job.
type ClassifyListResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Classify configuration used for this job
	Configuration ClassifyConfiguration `json:"configuration" api:"required"`
	// Whether the input was a file or parse job (FILE or PARSE_JOB)
	//
	// Any of "url", "file_id", "parse_job_id".
	DocumentInputType ClassifyListResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file or parse job
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED".
	Status ClassifyListResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Product configuration ID
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Associated parse job ID
	ParseJobID string `json:"parse_job_id" api:"nullable"`
	// Result of classifying a document.
	Result ClassifyResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		ParseJobID        respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ClassifyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job (FILE or PARSE_JOB)
type ClassifyListResponseDocumentInputType string

const (
	ClassifyListResponseDocumentInputTypeURL        ClassifyListResponseDocumentInputType = "url"
	ClassifyListResponseDocumentInputTypeFileID     ClassifyListResponseDocumentInputType = "file_id"
	ClassifyListResponseDocumentInputTypeParseJobID ClassifyListResponseDocumentInputType = "parse_job_id"
)

// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
type ClassifyListResponseStatus string

const (
	ClassifyListResponseStatusPending   ClassifyListResponseStatus = "PENDING"
	ClassifyListResponseStatusRunning   ClassifyListResponseStatus = "RUNNING"
	ClassifyListResponseStatusCompleted ClassifyListResponseStatus = "COMPLETED"
	ClassifyListResponseStatusFailed    ClassifyListResponseStatus = "FAILED"
)

// Response for a classify job.
type ClassifyGetResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Classify configuration used for this job
	Configuration ClassifyConfiguration `json:"configuration" api:"required"`
	// Whether the input was a file or parse job (FILE or PARSE_JOB)
	//
	// Any of "url", "file_id", "parse_job_id".
	DocumentInputType ClassifyGetResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file or parse job
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED".
	Status ClassifyGetResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Product configuration ID
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Associated parse job ID
	ParseJobID string `json:"parse_job_id" api:"nullable"`
	// Result of classifying a document.
	Result ClassifyResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		ParseJobID        respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ClassifyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job (FILE or PARSE_JOB)
type ClassifyGetResponseDocumentInputType string

const (
	ClassifyGetResponseDocumentInputTypeURL        ClassifyGetResponseDocumentInputType = "url"
	ClassifyGetResponseDocumentInputTypeFileID     ClassifyGetResponseDocumentInputType = "file_id"
	ClassifyGetResponseDocumentInputTypeParseJobID ClassifyGetResponseDocumentInputType = "parse_job_id"
)

// Current job status: PENDING, RUNNING, COMPLETED, or FAILED
type ClassifyGetResponseStatus string

const (
	ClassifyGetResponseStatusPending   ClassifyGetResponseStatus = "PENDING"
	ClassifyGetResponseStatusRunning   ClassifyGetResponseStatus = "RUNNING"
	ClassifyGetResponseStatusCompleted ClassifyGetResponseStatus = "COMPLETED"
	ClassifyGetResponseStatusFailed    ClassifyGetResponseStatus = "FAILED"
)

type ClassifyNewParams struct {
	// Request to create a classify job.
	ClassifyCreateRequest ClassifyCreateRequestParam
	OrganizationID        param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID             param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ClassifyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ClassifyCreateRequest)
}
func (r *ClassifyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ClassifyNewParams]'s query parameters as `url.Values`.
func (r ClassifyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ClassifyListParams struct {
	// Filter by configuration ID
	ConfigurationID param.Opt[string] `query:"configuration_id,omitzero" json:"-"`
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by job status
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED".
	Status ClassifyListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClassifyListParams]'s query parameters as `url.Values`.
func (r ClassifyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status
type ClassifyListParamsStatus string

const (
	ClassifyListParamsStatusPending   ClassifyListParamsStatus = "PENDING"
	ClassifyListParamsStatusRunning   ClassifyListParamsStatus = "RUNNING"
	ClassifyListParamsStatusCompleted ClassifyListParamsStatus = "COMPLETED"
	ClassifyListParamsStatusFailed    ClassifyListParamsStatus = "FAILED"
)

type ClassifyGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ClassifyGetParams]'s query parameters as `url.Values`.
func (r ClassifyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
