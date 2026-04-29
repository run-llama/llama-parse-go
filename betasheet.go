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

// BetaSheetService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaSheetService] method instead.
type BetaSheetService struct {
	options []option.RequestOption
}

// NewBetaSheetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaSheetService(opts ...option.RequestOption) (r BetaSheetService) {
	r = BetaSheetService{}
	r.options = opts
	return
}

// Create a spreadsheet parsing job. Experimental: This endpoint is not yet ready
// for production use and is subject to change at any time.
func (r *BetaSheetService) New(ctx context.Context, params BetaSheetNewParams, opts ...option.RequestOption) (res *SheetsJob, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/sheets/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List spreadsheet parsing jobs. Experimental: This endpoint is not yet ready for
// production use and is subject to change at any time.
func (r *BetaSheetService) List(ctx context.Context, query BetaSheetListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[SheetsJob], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/sheets/jobs"
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

// List spreadsheet parsing jobs. Experimental: This endpoint is not yet ready for
// production use and is subject to change at any time.
func (r *BetaSheetService) ListAutoPaging(ctx context.Context, query BetaSheetListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[SheetsJob] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a spreadsheet parsing job and its associated data. Experimental: This
// endpoint is not yet ready for production use and is subject to change at any
// time.
func (r *BetaSheetService) DeleteJob(ctx context.Context, spreadsheetJobID string, body BetaSheetDeleteJobParams, opts ...option.RequestOption) (res *BetaSheetDeleteJobResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if spreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/sheets/jobs/%s", url.PathEscape(spreadsheetJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Get a spreadsheet parsing job.
//
// When include_results=True (default), the response will include extracted regions
// and results if the job is complete, eliminating the need for a separate /results
// call.
//
// Experimental: This endpoint is not yet ready for production use and is subject
// to change at any time.
func (r *BetaSheetService) Get(ctx context.Context, spreadsheetJobID string, query BetaSheetGetParams, opts ...option.RequestOption) (res *SheetsJob, err error) {
	opts = slices.Concat(r.options, opts)
	if spreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/sheets/jobs/%s", url.PathEscape(spreadsheetJobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Generate a presigned URL to download a specific extracted region. Experimental:
// This endpoint is not yet ready for production use and is subject to change at
// any time.
func (r *BetaSheetService) GetResultTable(ctx context.Context, regionType BetaSheetGetResultTableParamsRegionType, params BetaSheetGetResultTableParams, opts ...option.RequestOption) (res *PresignedURL, err error) {
	opts = slices.Concat(r.options, opts)
	if params.SpreadsheetJobID == "" {
		err = errors.New("missing required spreadsheet_job_id parameter")
		return nil, err
	}
	if params.RegionID == "" {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/sheets/jobs/%s/regions/%s/result/%v", url.PathEscape(params.SpreadsheetJobID), url.PathEscape(params.RegionID), regionType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// A spreadsheet parsing job
type SheetsJob struct {
	// The ID of the job
	ID string `json:"id" api:"required"`
	// Configuration for the parsing job
	Config SheetsParsingConfig `json:"config" api:"required"`
	// When the job was created
	CreatedAt string `json:"created_at" api:"required"`
	// The ID of the input file
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// The ID of the project
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// The status of the parsing job
	//
	// Any of "PENDING", "SUCCESS", "ERROR", "PARTIAL_SUCCESS", "CANCELLED".
	Status StatusEnum `json:"status" api:"required"`
	// When the job was last updated
	UpdatedAt string `json:"updated_at" api:"required"`
	// The ID of the user
	UserID string `json:"user_id" api:"required"`
	// Any errors encountered
	Errors []string `json:"errors"`
	// Schema for a file.
	//
	// Deprecated: deprecated
	File File `json:"file" api:"nullable"`
	// All extracted regions (populated when job is complete)
	Regions []SheetsJobRegion `json:"regions"`
	// Whether the job completed successfully
	Success bool `json:"success" api:"nullable"`
	// Metadata for each processed worksheet (populated when job is complete)
	WorksheetMetadata []SheetsJobWorksheetMetadata `json:"worksheet_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		CreatedAt         respjson.Field
		FileID            respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		UserID            respjson.Field
		Errors            respjson.Field
		File              respjson.Field
		Regions           respjson.Field
		Success           respjson.Field
		WorksheetMetadata respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SheetsJob) RawJSON() string { return r.JSON.raw }
func (r *SheetsJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A summary of a single extracted region from a spreadsheet
type SheetsJobRegion struct {
	// Location of the region in the spreadsheet
	Location string `json:"location" api:"required"`
	// Type of the extracted region
	RegionType string `json:"region_type" api:"required"`
	// Worksheet name where region was found
	SheetName string `json:"sheet_name" api:"required"`
	// Generated description for the region
	Description string `json:"description" api:"nullable"`
	// Unique identifier for this region within the file
	RegionID string `json:"region_id"`
	// Generated title for the region
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location    respjson.Field
		RegionType  respjson.Field
		SheetName   respjson.Field
		Description respjson.Field
		RegionID    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SheetsJobRegion) RawJSON() string { return r.JSON.raw }
func (r *SheetsJobRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about a worksheet in a spreadsheet
type SheetsJobWorksheetMetadata struct {
	// Name of the worksheet
	SheetName string `json:"sheet_name" api:"required"`
	// Generated description of the worksheet
	Description string `json:"description" api:"nullable"`
	// Generated title for the worksheet
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SheetName   respjson.Field
		Description respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SheetsJobWorksheetMetadata) RawJSON() string { return r.JSON.raw }
func (r *SheetsJobWorksheetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for spreadsheet parsing and region extraction
type SheetsParsingConfig struct {
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange string `json:"extraction_range" api:"nullable"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables bool `json:"flatten_hierarchical_tables"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata bool `json:"generate_additional_metadata"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells bool `json:"include_hidden_cells"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names" api:"nullable"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization string `json:"specialization" api:"nullable"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity SheetsParsingConfigTableMergeSensitivity `json:"table_merge_sensitivity"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing bool `json:"use_experimental_processing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtractionRange            respjson.Field
		FlattenHierarchicalTables  respjson.Field
		GenerateAdditionalMetadata respjson.Field
		IncludeHiddenCells         respjson.Field
		SheetNames                 respjson.Field
		Specialization             respjson.Field
		TableMergeSensitivity      respjson.Field
		UseExperimentalProcessing  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SheetsParsingConfig) RawJSON() string { return r.JSON.raw }
func (r *SheetsParsingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SheetsParsingConfig to a SheetsParsingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SheetsParsingConfigParam.Overrides()
func (r SheetsParsingConfig) ToParam() SheetsParsingConfigParam {
	return param.Override[SheetsParsingConfigParam](json.RawMessage(r.RawJSON()))
}

// Influences how likely similar-looking regions are merged into a single table.
// Useful for spreadsheets that either have sparse tables (strong merging) or many
// distinct tables close together (weak merging).
type SheetsParsingConfigTableMergeSensitivity string

const (
	SheetsParsingConfigTableMergeSensitivityStrong SheetsParsingConfigTableMergeSensitivity = "strong"
	SheetsParsingConfigTableMergeSensitivityWeak   SheetsParsingConfigTableMergeSensitivity = "weak"
)

// Configuration for spreadsheet parsing and region extraction
type SheetsParsingConfigParam struct {
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange param.Opt[string] `json:"extraction_range,omitzero"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization param.Opt[string] `json:"specialization,omitzero"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables param.Opt[bool] `json:"flatten_hierarchical_tables,omitzero"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata param.Opt[bool] `json:"generate_additional_metadata,omitzero"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells param.Opt[bool] `json:"include_hidden_cells,omitzero"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing param.Opt[bool] `json:"use_experimental_processing,omitzero"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names,omitzero"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity SheetsParsingConfigTableMergeSensitivity `json:"table_merge_sensitivity,omitzero"`
	paramObj
}

func (r SheetsParsingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SheetsParsingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SheetsParsingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSheetDeleteJobResponse = any

type BetaSheetNewParams struct {
	// The ID of the file to parse
	FileID         string            `json:"file_id" api:"required" format:"uuid"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Configuration for the parsing job
	Config SheetsParsingConfigParam `json:"config,omitzero"`
	paramObj
}

func (r BetaSheetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaSheetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaSheetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaSheetNewParams]'s query parameters as `url.Values`.
func (r BetaSheetNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSheetListParams struct {
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
	// Any of "PENDING", "SUCCESS", "ERROR", "PARTIAL_SUCCESS", "CANCELLED".
	Status StatusEnum `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSheetListParams]'s query parameters as `url.Values`.
func (r BetaSheetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSheetDeleteJobParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSheetDeleteJobParams]'s query parameters as
// `url.Values`.
func (r BetaSheetDeleteJobParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSheetGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	IncludeResults param.Opt[bool]   `query:"include_results,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSheetGetParams]'s query parameters as `url.Values`.
func (r BetaSheetGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSheetGetResultTableParams struct {
	SpreadsheetJobID string            `path:"spreadsheet_job_id" api:"required" json:"-"`
	RegionID         string            `path:"region_id" api:"required" json:"-"`
	ExpiresAtSeconds param.Opt[int64]  `query:"expires_at_seconds,omitzero" json:"-"`
	OrganizationID   param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID        param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaSheetGetResultTableParams]'s query parameters as
// `url.Values`.
func (r BetaSheetGetResultTableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSheetGetResultTableParamsRegionType string

const (
	BetaSheetGetResultTableParamsRegionTypeTable        BetaSheetGetResultTableParamsRegionType = "table"
	BetaSheetGetResultTableParamsRegionTypeExtra        BetaSheetGetResultTableParamsRegionType = "extra"
	BetaSheetGetResultTableParamsRegionTypeCellMetadata BetaSheetGetResultTableParamsRegionType = "cell_metadata"
)
