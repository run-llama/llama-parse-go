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

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
)

// SheetService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSheetService] method instead.
type SheetService struct {
	options []option.RequestOption
}

// NewSheetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSheetService(opts ...option.RequestOption) (r SheetService) {
	r = SheetService{}
	r.options = opts
	return
}

// Create a spreadsheet parsing job.
//
// Provide at most one of `configuration` (an inline parsing configuration) or
// `configuration_id` (a saved configuration preset). If neither is provided, a
// default configuration is used. Optionally include `webhook_configurations` to
// receive `sheets.*` status notifications.
func (r *SheetService) New(ctx context.Context, params SheetNewParams, opts ...option.RequestOption) (res *SheetsJob, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/sheets/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List spreadsheet parsing jobs.
func (r *SheetService) List(ctx context.Context, query SheetListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[SheetsJob], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/sheets/jobs"
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

// List spreadsheet parsing jobs.
func (r *SheetService) ListAutoPaging(ctx context.Context, query SheetListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[SheetsJob] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a spreadsheet parsing job and its associated data.
func (r *SheetService) DeleteJob(ctx context.Context, spreadsheetJobID string, body SheetDeleteJobParams, opts ...option.RequestOption) (res *SheetDeleteJobResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if spreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/sheets/jobs/%s", url.PathEscape(spreadsheetJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Get a spreadsheet parsing job. When `include_results=True` (default), embeds
// extracted regions and results if complete, skipping the separate `/results`
// call.
func (r *SheetService) Get(ctx context.Context, spreadsheetJobID string, query SheetGetParams, opts ...option.RequestOption) (res *SheetsJob, err error) {
	opts = slices.Concat(r.options, opts)
	if spreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/sheets/jobs/%s", url.PathEscape(spreadsheetJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Generate a presigned URL to download a specific extracted region.
func (r *SheetService) GetResultTable(ctx context.Context, regionType SheetGetResultTableParamsRegionType, params SheetGetResultTableParams, opts ...option.RequestOption) (res *PresignedURL, err error) {
	opts = slices.Concat(r.options, opts)
	if params.SpreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	if params.RegionID == "" {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/sheets/jobs/%s/regions/%s/result/%v", url.PathEscape(params.SpreadsheetJobID), url.PathEscape(params.RegionID), regionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type SheetDeleteJobResponse = any

type SheetNewParams struct {
	// The ID of the file to parse
	FileID         string            `json:"file_id" api:"required" format:"uuid"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Saved configuration ID
	ConfigurationID param.Opt[string] `json:"configuration_id,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []SheetNewParamsWebhookConfiguration `json:"webhook_configurations,omitzero"`
	// Configuration for spreadsheet parsing and region extraction
	Config SheetsParsingConfigParam `json:"config,omitzero"`
	// Configuration for spreadsheet parsing and region extraction
	Configuration SheetsParsingConfigParam `json:"configuration,omitzero"`
	paramObj
}

func (r SheetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SheetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SheetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SheetNewParams]'s query parameters as `url.Values`.
func (r SheetNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Configuration for a single outbound webhook endpoint.
type SheetNewParamsWebhookConfiguration struct {
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
	// Any of "classify.cancelled", "classify.error", "classify.partial_success",
	// "classify.pending", "classify.running", "classify.success", "extract.cancelled",
	// "extract.error", "extract.partial_success", "extract.pending",
	// "extract.success", "parse.cancelled", "parse.error", "parse.partial_success",
	// "parse.pending", "parse.running", "parse.success", "sheets.cancelled",
	// "sheets.error", "sheets.partial_success", "sheets.pending", "sheets.success",
	// "split.cancelled", "split.error", "split.pending", "split.processing",
	// "split.success", "unmapped_event".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers sent with each webhook request (e.g. auth tokens)
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	paramObj
}

func (r SheetNewParamsWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow SheetNewParamsWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SheetNewParamsWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SheetListParams struct {
	// Filter by saved configuration ID
	ConfigurationID param.Opt[string] `query:"configuration_id,omitzero" json:"-"`
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize            param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	PageToken           param.Opt[string]    `query:"page_token,omitzero" json:"-"`
	ProjectID           param.Opt[string]    `query:"project_id,omitzero" format:"uuid" json:"-"`
	IncludeResults      param.Opt[bool]      `query:"include_results,omitzero" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by job status
	//
	// Any of "CANCELLED", "ERROR", "PARTIAL_SUCCESS", "PENDING", "SUCCESS".
	Status SheetListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SheetListParams]'s query parameters as `url.Values`.
func (r SheetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status
type SheetListParamsStatus string

const (
	SheetListParamsStatusCancelled      SheetListParamsStatus = "CANCELLED"
	SheetListParamsStatusError          SheetListParamsStatus = "ERROR"
	SheetListParamsStatusPartialSuccess SheetListParamsStatus = "PARTIAL_SUCCESS"
	SheetListParamsStatusPending        SheetListParamsStatus = "PENDING"
	SheetListParamsStatusSuccess        SheetListParamsStatus = "SUCCESS"
)

type SheetDeleteJobParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SheetDeleteJobParams]'s query parameters as `url.Values`.
func (r SheetDeleteJobParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SheetGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	IncludeResults param.Opt[bool]   `query:"include_results,omitzero" json:"-"`
	// Optional fields to populate on the response. Valid values:
	// metadata_state_transitions.
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SheetGetParams]'s query parameters as `url.Values`.
func (r SheetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SheetGetResultTableParams struct {
	SpreadsheetJobID string            `path:"spreadsheet_job_id" api:"required" json:"-"`
	RegionID         string            `path:"region_id" api:"required" json:"-"`
	ExpiresAtSeconds param.Opt[int64]  `query:"expires_at_seconds,omitzero" json:"-"`
	OrganizationID   param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID        param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [SheetGetResultTableParams]'s query parameters as
// `url.Values`.
func (r SheetGetResultTableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SheetGetResultTableParamsRegionType string

const (
	SheetGetResultTableParamsRegionTypeCellMetadata SheetGetResultTableParamsRegionType = "cell_metadata"
	SheetGetResultTableParamsRegionTypeExtra        SheetGetResultTableParamsRegionType = "extra"
	SheetGetResultTableParamsRegionTypeTable        SheetGetResultTableParamsRegionType = "table"
)
