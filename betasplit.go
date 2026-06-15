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

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// BetaSplitService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaSplitService] method instead.
type BetaSplitService struct {
	options []option.RequestOption
}

// NewBetaSplitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaSplitService(opts ...option.RequestOption) (r BetaSplitService) {
	r = BetaSplitService{}
	r.options = opts
	return
}

// Create a document split job.
func (r *BetaSplitService) New(ctx context.Context, params BetaSplitNewParams, opts ...option.RequestOption) (res *BetaSplitNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/split/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List document split jobs.
func (r *BetaSplitService) List(ctx context.Context, query BetaSplitListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BetaSplitListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/split/jobs"
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
func (r *BetaSplitService) ListAutoPaging(ctx context.Context, query BetaSplitListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BetaSplitListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a document split job.
func (r *BetaSplitService) Get(ctx context.Context, splitJobID string, query BetaSplitGetParams, opts ...option.RequestOption) (res *BetaSplitGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if splitJobID == "" {
		err = errors.New("missing required split_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/split/jobs/%s", url.PathEscape(splitJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Category definition for document splitting.
type SplitCategory struct {
	// Name of the category.
	Name string `json:"name" api:"required"`
	// Optional description of what content belongs in this category.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitCategory) RawJSON() string { return r.JSON.raw }
func (r *SplitCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SplitCategory to a SplitCategoryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SplitCategoryParam.Overrides()
func (r SplitCategory) ToParam() SplitCategoryParam {
	return param.Override[SplitCategoryParam](json.RawMessage(r.RawJSON()))
}

// Category definition for document splitting.
//
// The property Name is required.
type SplitCategoryParam struct {
	// Name of the category.
	Name string `json:"name" api:"required"`
	// Optional description of what content belongs in this category.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r SplitCategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow SplitCategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitCategoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document input specification for beta API.
type SplitDocumentInput struct {
	// Type of document input. Valid values are: file_id
	Type string `json:"type" api:"required"`
	// Document identifier.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitDocumentInput) RawJSON() string { return r.JSON.raw }
func (r *SplitDocumentInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SplitDocumentInput to a SplitDocumentInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SplitDocumentInputParam.Overrides()
func (r SplitDocumentInput) ToParam() SplitDocumentInputParam {
	return param.Override[SplitDocumentInputParam](json.RawMessage(r.RawJSON()))
}

// Document input specification for beta API.
//
// The properties Type, Value are required.
type SplitDocumentInputParam struct {
	// Type of document input. Valid values are: file_id
	Type string `json:"type" api:"required"`
	// Document identifier.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r SplitDocumentInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SplitDocumentInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitDocumentInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a completed split job.
type SplitResultResponse struct {
	// List of document segments.
	Segments []SplitSegmentResponse `json:"segments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitResultResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A segment of the split document.
type SplitSegmentResponse struct {
	// Category name this split belongs to.
	Category string `json:"category" api:"required"`
	// Categorical confidence level. Valid values are: high, medium, low.
	ConfidenceCategory string `json:"confidence_category" api:"required"`
	// 1-indexed page numbers in this split.
	Pages []int64 `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category           respjson.Field
		ConfidenceCategory respjson.Field
		Pages              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitSegmentResponse) RawJSON() string { return r.JSON.raw }
func (r *SplitSegmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta response — uses nested document_input object.
type BetaSplitNewResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Document that was split.
	DocumentInput SplitDocumentInput `json:"document_input" api:"required"`
	// Project ID this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current status of the job. Valid values are: pending, processing, completed,
	// failed, cancelled.
	Status string `json:"status" api:"required"`
	// User ID who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Categories      respjson.Field
		DocumentInput   respjson.Field
		ProjectID       respjson.Field
		Status          respjson.Field
		UserID          respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		ErrorMessage    respjson.Field
		Result          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSplitNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSplitNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta response — uses nested document_input object.
type BetaSplitListResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Document that was split.
	DocumentInput SplitDocumentInput `json:"document_input" api:"required"`
	// Project ID this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current status of the job. Valid values are: pending, processing, completed,
	// failed, cancelled.
	Status string `json:"status" api:"required"`
	// User ID who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Categories      respjson.Field
		DocumentInput   respjson.Field
		ProjectID       respjson.Field
		Status          respjson.Field
		UserID          respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		ErrorMessage    respjson.Field
		Result          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSplitListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSplitListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beta response — uses nested document_input object.
type BetaSplitGetResponse struct {
	// Unique identifier for the split job.
	ID string `json:"id" api:"required"`
	// Categories used for splitting.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Document that was split.
	DocumentInput SplitDocumentInput `json:"document_input" api:"required"`
	// Project ID this job belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Current status of the job. Valid values are: pending, processing, completed,
	// failed, cancelled.
	Status string `json:"status" api:"required"`
	// User ID who created this job.
	UserID string `json:"user_id" api:"required"`
	// Split configuration ID used for this job.
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if the job failed.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a completed split job.
	Result SplitResultResponse `json:"result" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Categories      respjson.Field
		DocumentInput   respjson.Field
		ProjectID       respjson.Field
		Status          respjson.Field
		UserID          respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		ErrorMessage    respjson.Field
		Result          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSplitGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSplitGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSplitNewParams struct {
	// Document to be split.
	DocumentInput  SplitDocumentInputParam `json:"document_input,omitzero" api:"required"`
	OrganizationID param.Opt[string]       `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]       `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Saved split configuration ID.
	ConfigurationID param.Opt[string] `json:"configuration_id,omitzero"`
	// Split configuration with categories and splitting strategy.
	Configuration BetaSplitNewParamsConfiguration `json:"configuration,omitzero"`
	paramObj
}

func (r BetaSplitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaSplitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaSplitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaSplitNewParams]'s query parameters as `url.Values`.
func (r BetaSplitNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Split configuration with categories and splitting strategy.
//
// The property Categories is required.
type BetaSplitNewParamsConfiguration struct {
	// Categories to split documents into.
	Categories []SplitCategoryParam `json:"categories,omitzero" api:"required"`
	// Strategy for splitting documents.
	SplittingStrategy BetaSplitNewParamsConfigurationSplittingStrategy `json:"splitting_strategy,omitzero"`
	paramObj
}

func (r BetaSplitNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow BetaSplitNewParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaSplitNewParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy for splitting documents.
type BetaSplitNewParamsConfigurationSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "include", "forbid", "omit".
	AllowUncategorized string `json:"allow_uncategorized,omitzero"`
	paramObj
}

func (r BetaSplitNewParamsConfigurationSplittingStrategy) MarshalJSON() (data []byte, err error) {
	type shadow BetaSplitNewParamsConfigurationSplittingStrategy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaSplitNewParamsConfigurationSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaSplitNewParamsConfigurationSplittingStrategy](
		"allow_uncategorized", "include", "forbid", "omit",
	)
}

type BetaSplitListParams struct {
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
	// Any of "pending", "processing", "completed", "failed", "cancelled".
	Status BetaSplitListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSplitListParams]'s query parameters as `url.Values`.
func (r BetaSplitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status (pending, processing, completed, failed, cancelled)
type BetaSplitListParamsStatus string

const (
	BetaSplitListParamsStatusPending    BetaSplitListParamsStatus = "pending"
	BetaSplitListParamsStatusProcessing BetaSplitListParamsStatus = "processing"
	BetaSplitListParamsStatusCompleted  BetaSplitListParamsStatus = "completed"
	BetaSplitListParamsStatusFailed     BetaSplitListParamsStatus = "failed"
	BetaSplitListParamsStatusCancelled  BetaSplitListParamsStatus = "cancelled"
)

type BetaSplitGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSplitGetParams]'s query parameters as `url.Values`.
func (r BetaSplitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
