// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"context"
	"encoding/json"
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

// ExtractionAgentService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractionAgentService] method instead.
type ExtractionAgentService struct {
	options []option.RequestOption
}

// NewExtractionAgentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtractionAgentService(opts ...option.RequestOption) (r ExtractionAgentService) {
	r = ExtractionAgentService{}
	r.options = opts
	return
}

// List the extraction agents in a project, newest first.
func (r *ExtractionAgentService) List(ctx context.Context, query ExtractionAgentListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ExtractAgent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/extraction-agents"
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

// List the extraction agents in a project, newest first.
func (r *ExtractionAgentService) ListAutoPaging(ctx context.Context, query ExtractionAgentListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ExtractAgent] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Schema and configuration for an extraction agent.
type ExtractAgent struct {
	// The id of the extraction agent.
	ID string `json:"id" api:"required" format:"uuid"`
	// The configuration parameters for the extraction agent.
	Config ExtractAgentConfig `json:"config" api:"required"`
	// The schema of the data.
	DataSchema map[string]*ExtractAgentDataSchemaUnion `json:"data_schema" api:"required"`
	// The name of the extraction agent.
	Name string `json:"name" api:"required"`
	// The ID of the project that the extraction agent belongs to.
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// The creation time of the extraction agent.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Custom configuration type for the extraction agent. Currently supports
	// 'default'.
	//
	// Any of "default".
	CustomConfiguration ExtractAgentCustomConfiguration `json:"custom_configuration" api:"nullable"`
	// The last update time of the extraction agent.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Config              respjson.Field
		DataSchema          respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		CreatedAt           respjson.Field
		CustomConfiguration respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractAgent) RawJSON() string { return r.JSON.raw }
func (r *ExtractAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration parameters for the extraction agent.
type ExtractAgentConfig struct {
	// The mode to use for chunking the document.
	//
	// Any of "PAGE", "SECTION".
	ChunkMode string `json:"chunk_mode"`
	// Whether to fetch citation bounding boxes for the extraction. Only available in
	// PREMIUM mode. Deprecated: this is now synonymous with cite_sources.
	//
	// Deprecated: deprecated
	CitationBbox bool `json:"citation_bbox"`
	// Whether to cite sources for the extraction.
	CiteSources bool `json:"cite_sources"`
	// Whether to fetch confidence scores for the extraction.
	ConfidenceScores bool `json:"confidence_scores"`
	// The extract model to use for data extraction. If not provided, uses the default
	// for the extraction mode.
	ExtractModel string `json:"extract_model" api:"nullable"`
	// The extraction mode specified (FAST, BALANCED, MULTIMODAL, PREMIUM).
	//
	// Any of "BALANCED", "FAST", "MULTIMODAL", "PREMIUM".
	ExtractionMode string `json:"extraction_mode"`
	// The extraction target specified.
	//
	// Any of "PER_DOC", "PER_PAGE", "PER_TABLE_ROW".
	ExtractionTarget string `json:"extraction_target"`
	// Whether to use high resolution mode for the extraction.
	HighResolutionMode bool `json:"high_resolution_mode"`
	// Whether to invalidate the cache for the extraction.
	InvalidateCache bool `json:"invalidate_cache"`
	// DEPRECATED: Whether to use fast mode for multimodal extraction.
	MultimodalFastMode bool `json:"multimodal_fast_mode"`
	// Number of pages to pass as context on long document extraction.
	NumPagesContext int64 `json:"num_pages_context" api:"nullable"`
	// Comma-separated list of page numbers or ranges to extract from (1-based, e.g.,
	// '1,3,5-7,9' or '1-3,8-10').
	PageRange string `json:"page_range" api:"nullable"`
	// Public model names.
	//
	// Any of "anthropic-haiku-3.5", "anthropic-haiku-4.5", "anthropic-sonnet-3.5",
	// "anthropic-sonnet-3.5-v2", "anthropic-sonnet-3.7", "anthropic-sonnet-4.0",
	// "anthropic-sonnet-4.5", "gemini-2.0-flash", "gemini-2.0-flash-lite",
	// "gemini-2.5-flash", "gemini-2.5-flash-lite", "gemini-2.5-pro", "gemini-3.0-pro",
	// "gemini-3.1-pro", "openai-gpt-4-1", "openai-gpt-4-1-mini",
	// "openai-gpt-4-1-nano", "openai-gpt-4o", "openai-gpt-4o-mini", "openai-gpt-5",
	// "openai-gpt-5-mini", "openai-gpt-5-nano", "openai-text-embedding-3-large",
	// "openai-text-embedding-3-small", "openai-whisper-1".
	ParseModel string `json:"parse_model" api:"nullable"`
	// The priority for the request. This field may be ignored or overwritten depending
	// on the organization tier.
	//
	// Any of "critical", "high", "low", "medium".
	Priority string `json:"priority" api:"nullable"`
	// The system prompt to use for the extraction.
	SystemPrompt string `json:"system_prompt" api:"nullable"`
	// Whether to use reasoning for the extraction.
	UseReasoning bool `json:"use_reasoning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkMode          respjson.Field
		CitationBbox       respjson.Field
		CiteSources        respjson.Field
		ConfidenceScores   respjson.Field
		ExtractModel       respjson.Field
		ExtractionMode     respjson.Field
		ExtractionTarget   respjson.Field
		HighResolutionMode respjson.Field
		InvalidateCache    respjson.Field
		MultimodalFastMode respjson.Field
		NumPagesContext    respjson.Field
		PageRange          respjson.Field
		ParseModel         respjson.Field
		Priority           respjson.Field
		SystemPrompt       respjson.Field
		UseReasoning       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractAgentConfig) RawJSON() string { return r.JSON.raw }
func (r *ExtractAgentConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractAgentDataSchemaUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractAgentDataSchemaMapItem OfAnyArray OfString OfFloat
// OfBool]
type ExtractAgentDataSchemaUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractAgentDataSchemaMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractAgentDataSchemaMapItem respjson.Field
		OfAnyArray                      respjson.Field
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		raw                             string
	} `json:"-"`
}

func (u ExtractAgentDataSchemaUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractAgentDataSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractAgentDataSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractAgentDataSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractAgentDataSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractAgentDataSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractAgentDataSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom configuration type for the extraction agent. Currently supports
// 'default'.
type ExtractAgentCustomConfiguration string

const (
	ExtractAgentCustomConfigurationDefault ExtractAgentCustomConfiguration = "default"
)

type ExtractionAgentListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Cursor from the previous page's `next_page_token`.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Whether to include default agents in the results
	IncludeDefault param.Opt[bool] `query:"include_default,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractionAgentListParams]'s query parameters as
// `url.Values`.
func (r ExtractionAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
