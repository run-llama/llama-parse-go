// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// SplitService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSplitService] method instead.
type SplitService struct {
	options []option.RequestOption
}

// NewSplitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSplitService(opts ...option.RequestOption) (r SplitService) {
	r = SplitService{}
	r.options = opts
	return
}

// Create a document split job.
func (r *SplitService) New(ctx context.Context, params SplitNewParams, opts ...option.RequestOption) (res *SplitNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/split/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List document split jobs.
func (r *SplitService) List(ctx context.Context, query SplitListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[SplitListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/split/jobs"
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

// List document split jobs.
func (r *SplitService) ListAutoPaging(ctx context.Context, query SplitListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[SplitListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a split job and its results.
func (r *SplitService) Delete(ctx context.Context, splitJobID string, body SplitDeleteParams, opts ...option.RequestOption) (res *SplitDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if splitJobID == "" {
		err = errors.New("missing required split_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/split/jobs/%s", url.PathEscape(splitJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Cancel a running split job.
//
// Requests cancellation; the job transitions to CANCELLED asynchronously once
// processing stops. Returns the job, which may still be in its current
// non-terminal state. Jobs already in a terminal state (COMPLETED, FAILED,
// CANCELLED) cannot be cancelled.
func (r *SplitService) Cancel(ctx context.Context, splitJobID string, body SplitCancelParams, opts ...option.RequestOption) (res *SplitCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if splitJobID == "" {
		err = errors.New("missing required split_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/split/jobs/%s/cancel", url.PathEscape(splitJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a document split job.
func (r *SplitService) Get(ctx context.Context, splitJobID string, query SplitGetParams, opts ...option.RequestOption) (res *SplitGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if splitJobID == "" {
		err = errors.New("missing required split_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/split/jobs/%s", url.PathEscape(splitJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A split job.
type SplitNewResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Whether the input was a file or parse job
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType SplitNewResponseDocumentInputType `json:"document_input_type" api:"required"`
	// File ID or parse job ID
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current job status. Valid values are: pending, processing, completed, failed,
	// cancelled.
	Status string `json:"status" api:"required"`
	// User who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Strategy used for splitting.
	SplittingStrategy SplitNewResponseSplittingStrategy `json:"splitting_strategy"`
	// Idempotency key scoped to the project, if one was provided.
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Categories        respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		SplittingStrategy respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job
type SplitNewResponseDocumentInputType string

const (
	SplitNewResponseDocumentInputTypeFileID     SplitNewResponseDocumentInputType = "file_id"
	SplitNewResponseDocumentInputTypeParseJobID SplitNewResponseDocumentInputType = "parse_job_id"
	SplitNewResponseDocumentInputTypeURL        SplitNewResponseDocumentInputType = "url"
)

// Strategy used for splitting.
type SplitNewResponseSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized"`
	// Minimum pages per segment. Shorter segments are merged into an adjacent segment;
	// 1 disables merging.
	MinPagesPerSplit int64 `json:"min_pages_per_split"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowUncategorized respjson.Field
		MinPagesPerSplit   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitNewResponseSplittingStrategy) RawJSON() string { return r.JSON.raw }
func (r *SplitNewResponseSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A split job.
type SplitListResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Whether the input was a file or parse job
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType SplitListResponseDocumentInputType `json:"document_input_type" api:"required"`
	// File ID or parse job ID
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current job status. Valid values are: pending, processing, completed, failed,
	// cancelled.
	Status string `json:"status" api:"required"`
	// User who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Strategy used for splitting.
	SplittingStrategy SplitListResponseSplittingStrategy `json:"splitting_strategy"`
	// Idempotency key scoped to the project, if one was provided.
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Categories        respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		SplittingStrategy respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitListResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job
type SplitListResponseDocumentInputType string

const (
	SplitListResponseDocumentInputTypeFileID     SplitListResponseDocumentInputType = "file_id"
	SplitListResponseDocumentInputTypeParseJobID SplitListResponseDocumentInputType = "parse_job_id"
	SplitListResponseDocumentInputTypeURL        SplitListResponseDocumentInputType = "url"
)

// Strategy used for splitting.
type SplitListResponseSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized"`
	// Minimum pages per segment. Shorter segments are merged into an adjacent segment;
	// 1 disables merging.
	MinPagesPerSplit int64 `json:"min_pages_per_split"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowUncategorized respjson.Field
		MinPagesPerSplit   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitListResponseSplittingStrategy) RawJSON() string { return r.JSON.raw }
func (r *SplitListResponseSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SplitDeleteResponse = any

// A split job.
type SplitCancelResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Whether the input was a file or parse job
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType SplitCancelResponseDocumentInputType `json:"document_input_type" api:"required"`
	// File ID or parse job ID
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current job status. Valid values are: pending, processing, completed, failed,
	// cancelled.
	Status string `json:"status" api:"required"`
	// User who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Strategy used for splitting.
	SplittingStrategy SplitCancelResponseSplittingStrategy `json:"splitting_strategy"`
	// Idempotency key scoped to the project, if one was provided.
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Categories        respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		SplittingStrategy respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job
type SplitCancelResponseDocumentInputType string

const (
	SplitCancelResponseDocumentInputTypeFileID     SplitCancelResponseDocumentInputType = "file_id"
	SplitCancelResponseDocumentInputTypeParseJobID SplitCancelResponseDocumentInputType = "parse_job_id"
	SplitCancelResponseDocumentInputTypeURL        SplitCancelResponseDocumentInputType = "url"
)

// Strategy used for splitting.
type SplitCancelResponseSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized"`
	// Minimum pages per segment. Shorter segments are merged into an adjacent segment;
	// 1 disables merging.
	MinPagesPerSplit int64 `json:"min_pages_per_split"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowUncategorized respjson.Field
		MinPagesPerSplit   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitCancelResponseSplittingStrategy) RawJSON() string { return r.JSON.raw }
func (r *SplitCancelResponseSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A split job.
type SplitGetResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Whether the input was a file or parse job
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType SplitGetResponseDocumentInputType `json:"document_input_type" api:"required"`
	// File ID or parse job ID
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current job status. Valid values are: pending, processing, completed, failed,
	// cancelled.
	Status string `json:"status" api:"required"`
	// User who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Strategy used for splitting.
	SplittingStrategy SplitGetResponseSplittingStrategy `json:"splitting_strategy"`
	// Idempotency key scoped to the project, if one was provided.
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Categories        respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		ConfigurationID   respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		SplittingStrategy respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the input was a file or parse job
type SplitGetResponseDocumentInputType string

const (
	SplitGetResponseDocumentInputTypeFileID     SplitGetResponseDocumentInputType = "file_id"
	SplitGetResponseDocumentInputTypeParseJobID SplitGetResponseDocumentInputType = "parse_job_id"
	SplitGetResponseDocumentInputTypeURL        SplitGetResponseDocumentInputType = "url"
)

// Strategy used for splitting.
type SplitGetResponseSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized"`
	// Minimum pages per segment. Shorter segments are merged into an adjacent segment;
	// 1 disables merging.
	MinPagesPerSplit int64 `json:"min_pages_per_split"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowUncategorized respjson.Field
		MinPagesPerSplit   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitGetResponseSplittingStrategy) RawJSON() string { return r.JSON.raw }
func (r *SplitGetResponseSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SplitNewParams struct {
	// File ID or parse job ID
	FileInput      string            `json:"file_input" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Saved configuration ID
	ConfigurationID param.Opt[string] `json:"configuration_id,omitzero"`
	// Idempotency key scoped to the project. Reusing a key returns the original job;
	// the new request body is ignored.
	TransactionID param.Opt[string] `json:"transaction_id,omitzero"`
	// Split configuration with categories and splitting strategy.
	Configuration SplitNewParamsConfiguration `json:"configuration,omitzero"`
	// IDs of saved webhook configurations to notify for this job.
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []SplitNewParamsWebhookConfiguration `json:"webhook_configurations,omitzero"`
	paramObj
}

func (r SplitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SplitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SplitNewParams]'s query parameters as `url.Values`.
func (r SplitNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Split configuration with categories and splitting strategy.
//
// The property Categories is required.
type SplitNewParamsConfiguration struct {
	// Categories to split documents into.
	Categories []SplitCategoryParam `json:"categories,omitzero" api:"required"`
	// Strategy for splitting documents.
	SplittingStrategy SplitNewParamsConfigurationSplittingStrategy `json:"splitting_strategy,omitzero"`
	paramObj
}

func (r SplitNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow SplitNewParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitNewParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy for splitting documents.
type SplitNewParamsConfigurationSplittingStrategy struct {
	// Minimum pages per segment. Shorter segments are merged into an adjacent segment;
	// 1 disables merging.
	MinPagesPerSplit param.Opt[int64] `json:"min_pages_per_split,omitzero"`
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized,omitzero"`
	paramObj
}

func (r SplitNewParamsConfigurationSplittingStrategy) MarshalJSON() (data []byte, err error) {
	type shadow SplitNewParamsConfigurationSplittingStrategy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitNewParamsConfigurationSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SplitNewParamsConfigurationSplittingStrategy](
		"allow_uncategorized", "forbid", "include", "omit",
	)
}

// Configuration for a single outbound webhook endpoint.
type SplitNewParamsWebhookConfiguration struct {
	// Response format sent to the webhook: 'string' (default) or 'json'
	WebhookOutputFormat param.Opt[string] `json:"webhook_output_format,omitzero"`
	// Shared signing secret used to sign webhook deliveries. When set, each request
	// includes an HMAC-SHA256 signature of the request body in the 'LC-Signature'
	// header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with
	// this secret to verify the delivery is authentic.
	WebhookSigningSecret param.Opt[string] `json:"webhook_signing_secret,omitzero"`
	// URL to receive webhook POST notifications
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events to subscribe to (e.g. 'parse.success', 'extract.error'). If null, all
	// events are delivered.
	//
	// Any of "batch.cancelled", "batch.error", "batch.pending", "batch.running",
	// "batch.success", "classify.cancelled", "classify.error",
	// "classify.partial_success", "classify.pending", "classify.running",
	// "classify.success", "extract.cancelled", "extract.error",
	// "extract.partial_success", "extract.pending", "extract.success",
	// "parse.cancelled", "parse.error", "parse.partial_success", "parse.pending",
	// "parse.running", "parse.success", "sheets.cancelled", "sheets.error",
	// "sheets.partial_success", "sheets.pending", "sheets.success", "split.cancelled",
	// "split.error", "split.pending", "split.processing", "split.success",
	// "unmapped_event".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers sent with each webhook request (e.g. auth tokens)
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	paramObj
}

func (r SplitNewParamsWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow SplitNewParamsWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitNewParamsWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SplitListParams struct {
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize            param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	PageToken           param.Opt[string]    `query:"page_token,omitzero" json:"-"`
	ProjectID           param.Opt[string]    `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by job status (pending, processing, completed, failed, cancelled)
	//
	// Any of "cancelled", "completed", "failed", "pending", "processing".
	Status SplitListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SplitListParams]'s query parameters as `url.Values`.
func (r SplitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status (pending, processing, completed, failed, cancelled)
type SplitListParamsStatus string

const (
	SplitListParamsStatusCancelled  SplitListParamsStatus = "cancelled"
	SplitListParamsStatusCompleted  SplitListParamsStatus = "completed"
	SplitListParamsStatusFailed     SplitListParamsStatus = "failed"
	SplitListParamsStatusPending    SplitListParamsStatus = "pending"
	SplitListParamsStatusProcessing SplitListParamsStatus = "processing"
)

type SplitDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SplitDeleteParams]'s query parameters as `url.Values`.
func (r SplitDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SplitCancelParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SplitCancelParams]'s query parameters as `url.Values`.
func (r SplitCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SplitGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SplitGetParams]'s query parameters as `url.Values`.
func (r SplitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
