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
	shimjson "github.com/run-llama/llama-parse-go/internal/encoding/json"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// ExtractService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	options []option.RequestOption
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r ExtractService) {
	r = ExtractService{}
	r.options = opts
	return
}

// Create an extraction job.
//
// Extracts structured data from a document using either a saved configuration or
// an inline JSON Schema.
//
// ## Input
//
// Provide exactly one of:
//
// - `configuration_id` — reference a saved extraction config
// - `configuration` — inline configuration with a `data_schema`
//
// ## Document input
//
// Set `file_input` to a file ID (`dfl-...`) or a completed parse job ID
// (`pjb-...`).
//
// The job runs asynchronously. Poll `GET /extract/{job_id}` or register a webhook
// to monitor completion.
func (r *ExtractService) New(ctx context.Context, params ExtractNewParams, opts ...option.RequestOption) (res *ExtractV2Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List extraction jobs with optional filtering and pagination.
//
// Filter by `configuration_id`, `status`, `file_input`, or creation date range.
// Results are returned newest-first. Use `expand=configuration` to include the
// full configuration used, and `expand=extract_metadata` for per-field metadata.
func (r *ExtractService) List(ctx context.Context, query ExtractListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ExtractV2Job], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/extract"
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

// List extraction jobs with optional filtering and pagination.
//
// Filter by `configuration_id`, `status`, `file_input`, or creation date range.
// Results are returned newest-first. Use `expand=configuration` to include the
// full configuration used, and `expand=extract_metadata` for per-field metadata.
func (r *ExtractService) ListAutoPaging(ctx context.Context, query ExtractListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ExtractV2Job] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an extraction job and its results.
func (r *ExtractService) Delete(ctx context.Context, jobID string, body ExtractDeleteParams, opts ...option.RequestOption) (res *ExtractDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Generate a JSON schema and return a product configuration request.
func (r *ExtractService) GenerateSchema(ctx context.Context, params ExtractGenerateSchemaParams, opts ...option.RequestOption) (res *ConfigurationCreate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/extract/schema/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single extraction job by ID.
//
// Returns the job status and results when complete. Use `expand=configuration` to
// include the full configuration used, and `expand=extract_metadata` for per-field
// metadata.
func (r *ExtractService) Get(ctx context.Context, jobID string, query ExtractGetParams, opts ...option.RequestOption) (res *ExtractV2Job, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Validate a JSON schema for extraction.
func (r *ExtractService) ValidateSchema(ctx context.Context, body ExtractValidateSchemaParams, opts ...option.RequestOption) (res *ExtractV2SchemaValidateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/extract/schema/validation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extract configuration combining parse and extract settings.
type ExtractConfiguration struct {
	// JSON Schema defining the fields to extract. Validate with the /schema/validate
	// endpoint first.
	DataSchema map[string]*ExtractConfigurationDataSchemaUnion `json:"data_schema" api:"required"`
	// Include citations in results
	CiteSources bool `json:"cite_sources"`
	// Include confidence scores in results
	ConfidenceScores bool `json:"confidence_scores"`
	// Granularity of extraction: per_doc returns one object per document, per_page
	// returns one object per page, per_table_row returns one object per table row
	//
	// Any of "per_doc", "per_page", "per_table_row".
	ExtractionTarget ExtractConfigurationExtractionTarget `json:"extraction_target"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// Saved parse configuration ID to control how the document is parsed before
	// extraction
	ParseConfigID string `json:"parse_config_id" api:"nullable"`
	// Parse tier to use before extraction. Defaults to the extract tier if not
	// specified.
	ParseTier string `json:"parse_tier" api:"nullable"`
	// Custom system prompt to guide extraction behavior
	SystemPrompt string `json:"system_prompt" api:"nullable"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
	//
	// Any of "cost_effective", "agentic".
	Tier ExtractConfigurationTier `json:"tier"`
	// Use 'latest' for the latest release for the selected tier or a date string
	// (YYYY-MM-DD format) to pin to the nearest release at or before that date.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSchema       respjson.Field
		CiteSources      respjson.Field
		ConfidenceScores respjson.Field
		ExtractionTarget respjson.Field
		MaxPages         respjson.Field
		ParseConfigID    respjson.Field
		ParseTier        respjson.Field
		SystemPrompt     respjson.Field
		TargetPages      respjson.Field
		Tier             respjson.Field
		Version          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ExtractConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExtractConfiguration to a ExtractConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExtractConfigurationParam.Overrides()
func (r ExtractConfiguration) ToParam() ExtractConfigurationParam {
	return param.Override[ExtractConfigurationParam](json.RawMessage(r.RawJSON()))
}

// ExtractConfigurationDataSchemaUnion contains all possible properties and values
// from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractConfigurationDataSchemaMapItem OfAnyArray OfString
// OfFloat OfBool]
type ExtractConfigurationDataSchemaUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractConfigurationDataSchemaMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractConfigurationDataSchemaMapItem respjson.Field
		OfAnyArray                              respjson.Field
		OfString                                respjson.Field
		OfFloat                                 respjson.Field
		OfBool                                  respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ExtractConfigurationDataSchemaUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractConfigurationDataSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractConfigurationDataSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractConfigurationDataSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractConfigurationDataSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractConfigurationDataSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractConfigurationDataSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granularity of extraction: per_doc returns one object per document, per_page
// returns one object per page, per_table_row returns one object per table row
type ExtractConfigurationExtractionTarget string

const (
	ExtractConfigurationExtractionTargetPerDoc      ExtractConfigurationExtractionTarget = "per_doc"
	ExtractConfigurationExtractionTargetPerPage     ExtractConfigurationExtractionTarget = "per_page"
	ExtractConfigurationExtractionTargetPerTableRow ExtractConfigurationExtractionTarget = "per_table_row"
)

// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
type ExtractConfigurationTier string

const (
	ExtractConfigurationTierCostEffective ExtractConfigurationTier = "cost_effective"
	ExtractConfigurationTierAgentic       ExtractConfigurationTier = "agentic"
)

// Extract configuration combining parse and extract settings.
//
// The property DataSchema is required.
type ExtractConfigurationParam struct {
	// JSON Schema defining the fields to extract. Validate with the /schema/validate
	// endpoint first.
	DataSchema map[string]*ExtractConfigurationDataSchemaUnionParam `json:"data_schema,omitzero" api:"required"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// Saved parse configuration ID to control how the document is parsed before
	// extraction
	ParseConfigID param.Opt[string] `json:"parse_config_id,omitzero"`
	// Parse tier to use before extraction. Defaults to the extract tier if not
	// specified.
	ParseTier param.Opt[string] `json:"parse_tier,omitzero"`
	// Custom system prompt to guide extraction behavior
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	// Include citations in results
	CiteSources param.Opt[bool] `json:"cite_sources,omitzero"`
	// Include confidence scores in results
	ConfidenceScores param.Opt[bool] `json:"confidence_scores,omitzero"`
	// Use 'latest' for the latest release for the selected tier or a date string
	// (YYYY-MM-DD format) to pin to the nearest release at or before that date.
	Version param.Opt[string] `json:"version,omitzero"`
	// Granularity of extraction: per_doc returns one object per document, per_page
	// returns one object per page, per_table_row returns one object per table row
	//
	// Any of "per_doc", "per_page", "per_table_row".
	ExtractionTarget ExtractConfigurationExtractionTarget `json:"extraction_target,omitzero"`
	// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
	//
	// Any of "cost_effective", "agentic".
	Tier ExtractConfigurationTier `json:"tier,omitzero"`
	paramObj
}

func (r ExtractConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtractConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractConfigurationDataSchemaUnionParam struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractConfigurationDataSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *ExtractConfigurationDataSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Extraction metadata.
type ExtractJobMetadata struct {
	// Metadata for extracted fields including document, page, and row level info.
	FieldMetadata ExtractedFieldMetadata `json:"field_metadata" api:"nullable"`
	// Reference to the ParseJob ID used for parsing
	ParseJobID string `json:"parse_job_id" api:"nullable"`
	// Parse tier used for parsing the document
	ParseTier string `json:"parse_tier" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldMetadata respjson.Field
		ParseJobID    respjson.Field
		ParseTier     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extraction usage metrics.
type ExtractJobUsage struct {
	// Number of pages extracted
	NumPagesExtracted int64 `json:"num_pages_extracted" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumPagesExtracted respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractJobUsage) RawJSON() string { return r.JSON.raw }
func (r *ExtractJobUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An extraction job.
type ExtractV2Job struct {
	// Unique job identifier (job_id)
	ID string `json:"id" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// File ID or parse job ID that was extracted
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status.
	//
	// - `PENDING` — queued, not yet started
	// - `RUNNING` — actively processing
	// - `COMPLETED` — finished successfully
	// - `FAILED` — terminated with an error
	// - `CANCELLED` — cancelled by user
	Status string `json:"status" api:"required"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Extract configuration combining parse and extract settings.
	Configuration ExtractConfiguration `json:"configuration" api:"nullable"`
	// Saved extract configuration ID used for this job, if any
	ConfigurationID string `json:"configuration_id" api:"nullable"`
	// Error details when status is FAILED
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Extraction metadata.
	ExtractMetadata ExtractJobMetadata `json:"extract_metadata" api:"nullable"`
	// Extracted data conforming to the data_schema. Returns a single object for
	// per_doc, or an array for per_page / per_table_row.
	ExtractResult ExtractV2JobExtractResultUnion `json:"extract_result" api:"nullable"`
	// Job-level metadata.
	Metadata ExtractV2JobMetadata `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FileInput       respjson.Field
		ProjectID       respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Configuration   respjson.Field
		ConfigurationID respjson.Field
		ErrorMessage    respjson.Field
		ExtractMetadata respjson.Field
		ExtractResult   respjson.Field
		Metadata        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractV2Job) RawJSON() string { return r.JSON.raw }
func (r *ExtractV2Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractV2JobExtractResultUnion contains all possible properties and values from
// [map[string]*ExtractV2JobExtractResultMapItemUnion],
// [[]map[string]*ExtractV2JobExtractResultArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractV2JobExtractResultMapItemMapItem OfAnyArray OfString
// OfFloat OfBool OfMapOfExtractV2JobExtractResultArrayItemMap]
type ExtractV2JobExtractResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractV2JobExtractResultMapItemMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]map[string]\*ExtractV2JobExtractResultArrayItemUnion] instead of an object.
	OfMapOfExtractV2JobExtractResultArrayItemMap []map[string]*ExtractV2JobExtractResultArrayItemUnion `json:",inline"`
	JSON                                         struct {
		OfExtractV2JobExtractResultMapItemMapItem    respjson.Field
		OfAnyArray                                   respjson.Field
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfMapOfExtractV2JobExtractResultArrayItemMap respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u ExtractV2JobExtractResultUnion) AsExtractV2JobExtractResultMapMap() (v map[string]*ExtractV2JobExtractResultMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultUnion) AsMapOfExtractV2JobExtractResultArrayItemMap() (v []map[string]*ExtractV2JobExtractResultArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractV2JobExtractResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractV2JobExtractResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractV2JobExtractResultMapItemUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractV2JobExtractResultMapItemMapItem OfAnyArray OfString
// OfFloat OfBool]
type ExtractV2JobExtractResultMapItemUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractV2JobExtractResultMapItemMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractV2JobExtractResultMapItemMapItem respjson.Field
		OfAnyArray                                respjson.Field
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u ExtractV2JobExtractResultMapItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultMapItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractV2JobExtractResultMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractV2JobExtractResultMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractV2JobExtractResultArrayItemUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractV2JobExtractResultArrayItemMapItem OfAnyArray OfString
// OfFloat OfBool]
type ExtractV2JobExtractResultArrayItemUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractV2JobExtractResultArrayItemMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractV2JobExtractResultArrayItemMapItem respjson.Field
		OfAnyArray                                  respjson.Field
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u ExtractV2JobExtractResultArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2JobExtractResultArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractV2JobExtractResultArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractV2JobExtractResultArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job-level metadata.
type ExtractV2JobMetadata struct {
	// Extraction usage metrics.
	Usage       ExtractJobUsage `json:"usage" api:"nullable"`
	ExtraFields map[string]any  `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractV2JobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractV2JobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create an extraction job. Provide configuration_id or inline
// configuration.
//
// The property FileInput is required.
type ExtractV2JobCreateParam struct {
	// File ID or parse job ID to extract from
	FileInput string `json:"file_input" api:"required"`
	// Saved configuration ID
	ConfigurationID param.Opt[string] `json:"configuration_id,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []ExtractV2JobCreateWebhookConfigurationParam `json:"webhook_configurations,omitzero"`
	// Extract configuration combining parse and extract settings.
	Configuration ExtractConfigurationParam `json:"configuration,omitzero"`
	paramObj
}

func (r ExtractV2JobCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtractV2JobCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractV2JobCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a single outbound webhook endpoint.
type ExtractV2JobCreateWebhookConfigurationParam struct {
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

func (r ExtractV2JobCreateWebhookConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtractV2JobCreateWebhookConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractV2JobCreateWebhookConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of extraction jobs.
type ExtractV2JobQueryResponse struct {
	// The list of items.
	Items []ExtractV2Job `json:"items" api:"required"`
	// A token, which can be sent as page_token to retrieve the next page. If this
	// field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// The total number of items available. This is only populated when specifically
	// requested. The value may be an estimate and can be used for display purposes
	// only.
	TotalSize int64 `json:"total_size" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items         respjson.Field
		NextPageToken respjson.Field
		TotalSize     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractV2JobQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractV2JobQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request schema for generating an extraction schema.
type ExtractV2SchemaGenerateRequestParam struct {
	// Optional file ID to analyze for schema generation
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// Name for the generated configuration (auto-generated if omitted)
	Name param.Opt[string] `json:"name,omitzero"`
	// Natural language description of the data structure to extract
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Optional schema to validate, refine, or extend
	DataSchema map[string]*ExtractV2SchemaGenerateRequestDataSchemaUnionParam `json:"data_schema,omitzero"`
	paramObj
}

func (r ExtractV2SchemaGenerateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtractV2SchemaGenerateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractV2SchemaGenerateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractV2SchemaGenerateRequestDataSchemaUnionParam struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractV2SchemaGenerateRequestDataSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *ExtractV2SchemaGenerateRequestDataSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Request schema for validating an extraction schema.
//
// The property DataSchema is required.
type ExtractV2SchemaValidateRequestParam struct {
	// JSON Schema to validate for use with extract jobs
	DataSchema map[string]*ExtractV2SchemaValidateRequestDataSchemaUnionParam `json:"data_schema,omitzero" api:"required"`
	paramObj
}

func (r ExtractV2SchemaValidateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtractV2SchemaValidateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractV2SchemaValidateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractV2SchemaValidateRequestDataSchemaUnionParam struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractV2SchemaValidateRequestDataSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *ExtractV2SchemaValidateRequestDataSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Response schema for schema validation.
type ExtractV2SchemaValidateResponse struct {
	// Validated JSON Schema, ready for use in extract jobs
	DataSchema map[string]*ExtractV2SchemaValidateResponseDataSchemaUnion `json:"data_schema" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSchema  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractV2SchemaValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractV2SchemaValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractV2SchemaValidateResponseDataSchemaUnion contains all possible properties
// and values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractV2SchemaValidateResponseDataSchemaMapItem OfAnyArray
// OfString OfFloat OfBool]
type ExtractV2SchemaValidateResponseDataSchemaUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractV2SchemaValidateResponseDataSchemaMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractV2SchemaValidateResponseDataSchemaMapItem respjson.Field
		OfAnyArray                                         respjson.Field
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u ExtractV2SchemaValidateResponseDataSchemaUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2SchemaValidateResponseDataSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2SchemaValidateResponseDataSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2SchemaValidateResponseDataSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2SchemaValidateResponseDataSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractV2SchemaValidateResponseDataSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractV2SchemaValidateResponseDataSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for extracted fields including document, page, and row level info.
type ExtractedFieldMetadata struct {
	// Per-field metadata keyed by field name from your schema. Scalar fields (e.g.
	// `vendor`) map to a FieldMetadataEntry with citation and confidence. Array fields
	// (e.g. `items`) map to a list where each element contains per-sub-field
	// FieldMetadataEntry objects, indexed by array position. Nested objects contain
	// sub-field entries recursively.
	DocumentMetadata map[string]*ExtractedFieldMetadataDocumentMetadataUnion `json:"document_metadata" api:"nullable"`
	// Per-page metadata when extraction_target is per_page
	PageMetadata []map[string]*ExtractedFieldMetadataPageMetadataUnion `json:"page_metadata" api:"nullable"`
	// Per-row metadata when extraction_target is per_table_row
	RowMetadata []map[string]*ExtractedFieldMetadataRowMetadataUnion `json:"row_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentMetadata respjson.Field
		PageMetadata     respjson.Field
		RowMetadata      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractedFieldMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractedFieldMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractedFieldMetadataDocumentMetadataUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractedFieldMetadataDocumentMetadataMapItem OfAnyArray
// OfString OfFloat OfBool]
type ExtractedFieldMetadataDocumentMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractedFieldMetadataDocumentMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractedFieldMetadataDocumentMetadataMapItem respjson.Field
		OfAnyArray                                      respjson.Field
		OfString                                        respjson.Field
		OfFloat                                         respjson.Field
		OfBool                                          respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u ExtractedFieldMetadataDocumentMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataDocumentMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataDocumentMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataDocumentMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataDocumentMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractedFieldMetadataDocumentMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractedFieldMetadataDocumentMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractedFieldMetadataPageMetadataUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractedFieldMetadataPageMetadataMapItem OfAnyArray OfString
// OfFloat OfBool]
type ExtractedFieldMetadataPageMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractedFieldMetadataPageMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractedFieldMetadataPageMetadataMapItem respjson.Field
		OfAnyArray                                  respjson.Field
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u ExtractedFieldMetadataPageMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataPageMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataPageMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataPageMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataPageMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractedFieldMetadataPageMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractedFieldMetadataPageMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractedFieldMetadataRowMetadataUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractedFieldMetadataRowMetadataMapItem OfAnyArray OfString
// OfFloat OfBool]
type ExtractedFieldMetadataRowMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractedFieldMetadataRowMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractedFieldMetadataRowMetadataMapItem respjson.Field
		OfAnyArray                                 respjson.Field
		OfString                                   respjson.Field
		OfFloat                                    respjson.Field
		OfBool                                     respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u ExtractedFieldMetadataRowMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataRowMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataRowMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataRowMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractedFieldMetadataRowMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractedFieldMetadataRowMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractedFieldMetadataRowMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractDeleteResponse = any

type ExtractNewParams struct {
	// Request to create an extraction job. Provide configuration_id or inline
	// configuration.
	ExtractV2JobCreate ExtractV2JobCreateParam
	OrganizationID     param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID          param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExtractNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExtractV2JobCreate)
}
func (r *ExtractNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractNewParams]'s query parameters as `url.Values`.
func (r ExtractNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractListParams struct {
	// Filter by configuration ID
	ConfigurationID param.Opt[string] `query:"configuration_id,omitzero" json:"-"`
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	// Filter by document input type (file_id or parse_job_id)
	DocumentInputType param.Opt[string] `query:"document_input_type,omitzero" json:"-"`
	// Deprecated: use file_input instead
	DocumentInputValue param.Opt[string] `query:"document_input_value,omitzero" json:"-"`
	// Filter by file input value
	FileInput      param.Opt[string] `query:"file_input,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by status
	//
	// Any of "PENDING", "THROTTLED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status ExtractListParamsStatus `query:"status,omitzero" json:"-"`
	// Additional fields to include: configuration, extract_metadata
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractListParams]'s query parameters as `url.Values`.
func (r ExtractListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status
type ExtractListParamsStatus string

const (
	ExtractListParamsStatusPending   ExtractListParamsStatus = "PENDING"
	ExtractListParamsStatusThrottled ExtractListParamsStatus = "THROTTLED"
	ExtractListParamsStatusRunning   ExtractListParamsStatus = "RUNNING"
	ExtractListParamsStatusCompleted ExtractListParamsStatus = "COMPLETED"
	ExtractListParamsStatusFailed    ExtractListParamsStatus = "FAILED"
	ExtractListParamsStatusCancelled ExtractListParamsStatus = "CANCELLED"
)

type ExtractDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractDeleteParams]'s query parameters as `url.Values`.
func (r ExtractDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractGenerateSchemaParams struct {
	// Request schema for generating an extraction schema.
	ExtractV2SchemaGenerateRequest ExtractV2SchemaGenerateRequestParam
	OrganizationID                 param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID                      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExtractGenerateSchemaParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExtractV2SchemaGenerateRequest)
}
func (r *ExtractGenerateSchemaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractGenerateSchemaParams]'s query parameters as
// `url.Values`.
func (r ExtractGenerateSchemaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Additional fields to include: configuration, extract_metadata
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractGetParams]'s query parameters as `url.Values`.
func (r ExtractGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractValidateSchemaParams struct {
	// Request schema for validating an extraction schema.
	ExtractV2SchemaValidateRequest ExtractV2SchemaValidateRequestParam
	paramObj
}

func (r ExtractValidateSchemaParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExtractV2SchemaValidateRequest)
}
func (r *ExtractValidateSchemaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
