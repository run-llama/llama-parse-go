// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// RetrievalService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetrievalService] method instead.
type RetrievalService struct {
	options []option.RequestOption
}

// NewRetrievalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetrievalService(opts ...option.RequestOption) (r RetrievalService) {
	r = RetrievalService{}
	r.options = opts
	return
}

// Retrieve relevant chunks via hybrid search (vector + full-text), with filtering
// on built-in or user-defined metadata.
func (r *RetrievalService) Get(ctx context.Context, params RetrievalGetParams, opts ...option.RequestOption) (res *RetrievalGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrieval/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Search for files by name.
func (r *RetrievalService) Find(ctx context.Context, params RetrievalFindParams, opts ...option.RequestOption) (res *pagination.PaginatedCursorPost[RetrievalFindResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/retrieval/files/find"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Search for files by name.
func (r *RetrievalService) FindAutoPaging(ctx context.Context, params RetrievalFindParams, opts ...option.RequestOption) *pagination.PaginatedCursorPostAutoPager[RetrievalFindResponse] {
	return pagination.NewPaginatedCursorPostAutoPager(r.Find(ctx, params, opts...))
}

// Grep within a file's parsed content using a regex pattern.
func (r *RetrievalService) Grep(ctx context.Context, params RetrievalGrepParams, opts ...option.RequestOption) (res *pagination.PaginatedCursorPost[RetrievalGrepResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/retrieval/files/grep"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Grep within a file's parsed content using a regex pattern.
func (r *RetrievalService) GrepAutoPaging(ctx context.Context, params RetrievalGrepParams, opts ...option.RequestOption) *pagination.PaginatedCursorPostAutoPager[RetrievalGrepResponse] {
	return pagination.NewPaginatedCursorPostAutoPager(r.Grep(ctx, params, opts...))
}

// Read the parsed text content of a specific file.
func (r *RetrievalService) Read(ctx context.Context, params RetrievalReadParams, opts ...option.RequestOption) (res *RetrievalReadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrieval/files/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response containing retrieval results.
type RetrievalGetResponse struct {
	// Ordered list of retrieved chunks.
	Results []RetrievalGetResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RetrievalGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single retrieval result.
type RetrievalGetResponseResult struct {
	// Text content of the retrieved chunk.
	Content string `json:"content" api:"required"`
	// User-defined metadata associated with the chunk.
	Metadata map[string]RetrievalGetResponseResultMetadataUnion `json:"metadata" api:"nullable"`
	// Relevance score from the reranker, if reranking was applied.
	RerankScore float64 `json:"rerank_score" api:"nullable"`
	// Hybrid search relevance score.
	Score float64 `json:"score" api:"nullable"`
	// Built-in fields stored for every exported chunk.
	StaticFields RetrievalGetResponseResultStaticFields `json:"static_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		Metadata     respjson.Field
		RerankScore  respjson.Field
		Score        respjson.Field
		StaticFields respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *RetrievalGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RetrievalGetResponseResultMetadataUnion contains all possible properties and
// values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type RetrievalGetResponseResultMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfInt         respjson.Field
		OfFloat       respjson.Field
		OfBool        respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u RetrievalGetResponseResultMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RetrievalGetResponseResultMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RetrievalGetResponseResultMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RetrievalGetResponseResultMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RetrievalGetResponseResultMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RetrievalGetResponseResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *RetrievalGetResponseResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Built-in fields stored for every exported chunk.
type RetrievalGetResponseResultStaticFields struct {
	// Attachments associated with the chunk
	Attachments []RetrievalGetResponseResultStaticFieldsAttachment `json:"attachments"`
	// End character offset of the chunk.
	ChunkEndChar int64 `json:"chunk_end_char" api:"nullable"`
	// Index of the chunk within the file.
	ChunkIndex int64 `json:"chunk_index" api:"nullable"`
	// Start character offset of the chunk.
	ChunkStartChar int64 `json:"chunk_start_char" api:"nullable"`
	// Token count of the chunk.
	ChunkTokenCount int64 `json:"chunk_token_count" api:"nullable"`
	// Last page number covered by this chunk.
	PageRangeEnd int64 `json:"page_range_end" api:"nullable"`
	// First page number covered by this chunk.
	PageRangeStart int64 `json:"page_range_start" api:"nullable"`
	// ID of the parsed file.
	ParsedDirectoryFileID string `json:"parsed_directory_file_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments           respjson.Field
		ChunkEndChar          respjson.Field
		ChunkIndex            respjson.Field
		ChunkStartChar        respjson.Field
		ChunkTokenCount       respjson.Field
		PageRangeEnd          respjson.Field
		PageRangeStart        respjson.Field
		ParsedDirectoryFileID respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalGetResponseResultStaticFields) RawJSON() string { return r.JSON.raw }
func (r *RetrievalGetResponseResultStaticFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to a file attachment, retrievable via
// `GET /api/v1/beta/attachments/{attachment_name}?source_id=...`.
type RetrievalGetResponseResultStaticFieldsAttachment struct {
	// Attachment-relative path, e.g. 'screenshots/page_7.jpg'.
	AttachmentName string `json:"attachment_name" api:"required"`
	// File ID to pass as source_id when fetching the attachment.
	SourceID string `json:"source_id" api:"required"`
	// Attachment kind, e.g. 'screenshot', 'items'.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentName respjson.Field
		SourceID       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalGetResponseResultStaticFieldsAttachment) RawJSON() string { return r.JSON.raw }
func (r *RetrievalGetResponseResultStaticFieldsAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file returned by find.
type RetrievalFindResponse struct {
	// ID of the file.
	FileID string `json:"file_id" api:"required"`
	// Display name of the file.
	FileName string `json:"file_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		FileName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalFindResponse) RawJSON() string { return r.JSON.raw }
func (r *RetrievalFindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single grep match within a file.
type RetrievalGrepResponse struct {
	// Matched text content.
	Content string `json:"content" api:"required"`
	// End character offset of the match.
	EndChar int64 `json:"end_char" api:"required"`
	// Start character offset of the match.
	StartChar int64 `json:"start_char" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		EndChar     respjson.Field
		StartChar   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalGrepResponse) RawJSON() string { return r.JSON.raw }
func (r *RetrievalGrepResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File read result.
type RetrievalReadResponse struct {
	// Parsed text content of the file.
	Content string `json:"content" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalReadResponse) RawJSON() string { return r.JSON.raw }
func (r *RetrievalReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetrievalGetParams struct {
	// ID of the index to retrieve against.
	IndexID string `json:"index_id" api:"required"`
	// Natural-language query to retrieve relevant chunks.
	Query          string            `json:"query" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Weight of the full-text search pipeline (0-1).
	FullTextPipelineWeight param.Opt[float64] `json:"full_text_pipeline_weight,omitzero"`
	// Number of candidates for approximate nearest neighbor search.
	NumCandidates param.Opt[int64] `json:"num_candidates,omitzero"`
	// Minimum score threshold for returned results.
	ScoreThreshold param.Opt[float64] `json:"score_threshold,omitzero"`
	// Maximum number of results to return. Values above 500 are capped at 500.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Weight of the vector search pipeline (0-1).
	VectorPipelineWeight param.Opt[float64] `json:"vector_pipeline_weight,omitzero"`
	// Filters on user-defined metadata fields.
	CustomFilters map[string]*RetrievalGetParamsCustomFilterUnion `json:"custom_filters,omitzero"`
	// Filters on built-in document fields (page range, chunk index, etc.).
	StaticFilters RetrievalGetParamsStaticFilters `json:"static_filters,omitzero"`
	// Reranking configuration applied after hybrid search. Enabled by default.
	Rerank RetrievalGetParamsRerank `json:"rerank,omitzero"`
	paramObj
}

func (r RetrievalGetParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrievalGetParams]'s query parameters as `url.Values`.
func (r RetrievalGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RetrievalGetParamsCustomFilterUnion struct {
	OfValueFilter                    *RetrievalGetParamsCustomFilterValueFilter `json:",omitzero,inline"`
	OfRetrievalGetsCustomFilterArray []RetrievalGetParamsCustomFilterArrayItem  `json:",omitzero,inline"`
	paramUnion
}

func (u RetrievalGetParamsCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfValueFilter, u.OfRetrievalGetsCustomFilterArray)
}
func (u *RetrievalGetParamsCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Filter on a single metadata field value.
//
// The properties Operator, Value are required.
type RetrievalGetParamsCustomFilterValueFilter struct {
	// Any of "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin".
	Operator string                                              `json:"operator,omitzero" api:"required"`
	Value    RetrievalGetParamsCustomFilterValueFilterValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r RetrievalGetParamsCustomFilterValueFilter) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParamsCustomFilterValueFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParamsCustomFilterValueFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RetrievalGetParamsCustomFilterValueFilter](
		"operator", "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RetrievalGetParamsCustomFilterValueFilterValueUnion struct {
	OfString                                         param.Opt[string]                                              `json:",omitzero,inline"`
	OfBool                                           param.Opt[bool]                                                `json:",omitzero,inline"`
	OfFloat                                          param.Opt[float64]                                             `json:",omitzero,inline"`
	OfRetrievalGetsCustomFilterValueFilterValueArray []RetrievalGetParamsCustomFilterValueFilterValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u RetrievalGetParamsCustomFilterValueFilterValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool, u.OfFloat, u.OfRetrievalGetsCustomFilterValueFilterValueArray)
}
func (u *RetrievalGetParamsCustomFilterValueFilterValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RetrievalGetParamsCustomFilterValueFilterValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u RetrievalGetParamsCustomFilterValueFilterValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool, u.OfFloat)
}
func (u *RetrievalGetParamsCustomFilterValueFilterValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// One bound of a numeric range filter on a metadata field.
//
// The properties Operator, Value are required.
type RetrievalGetParamsCustomFilterArrayItem struct {
	// Any of "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin".
	Operator string                                            `json:"operator,omitzero" api:"required"`
	Value    RetrievalGetParamsCustomFilterArrayItemValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r RetrievalGetParamsCustomFilterArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParamsCustomFilterArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParamsCustomFilterArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RetrievalGetParamsCustomFilterArrayItem](
		"operator", "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RetrievalGetParamsCustomFilterArrayItemValueUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u RetrievalGetParamsCustomFilterArrayItemValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *RetrievalGetParamsCustomFilterArrayItemValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Reranking configuration applied after hybrid search. Enabled by default.
type RetrievalGetParamsRerank struct {
	// Number of results to return after reranking.
	TopN param.Opt[int64] `json:"top_n,omitzero"`
	// Set to false to disable reranking.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r RetrievalGetParamsRerank) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParamsRerank
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParamsRerank) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filters on built-in document fields (page range, chunk index, etc.).
type RetrievalGetParamsStaticFilters struct {
	// Filter on a string field.
	ParsedDirectoryFileID RetrievalGetParamsStaticFiltersParsedDirectoryFileID `json:"parsed_directory_file_id,omitzero"`
	paramObj
}

func (r RetrievalGetParamsStaticFilters) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParamsStaticFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParamsStaticFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter on a string field.
//
// The properties Operator, Value are required.
type RetrievalGetParamsStaticFiltersParsedDirectoryFileID struct {
	// Any of "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin".
	Operator string                                                         `json:"operator,omitzero" api:"required"`
	Value    RetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion `json:"value,omitzero" api:"required"`
	paramObj
}

func (r RetrievalGetParamsStaticFiltersParsedDirectoryFileID) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGetParamsStaticFiltersParsedDirectoryFileID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGetParamsStaticFiltersParsedDirectoryFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RetrievalGetParamsStaticFiltersParsedDirectoryFileID](
		"operator", "eq", "gt", "gte", "in", "lt", "lte", "ne", "nin",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u RetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *RetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type RetrievalFindParams struct {
	// ID of the index to search within.
	IndexID        string            `json:"index_id" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Exact file name to match.
	FileName param.Opt[string] `json:"file_name,omitzero"`
	// Substring match on file name (case-insensitive).
	FileNameContains param.Opt[string] `json:"file_name_contains,omitzero"`
	// The maximum number of items to return. The service may return fewer than this
	// value. If unspecified, a default page size will be used. The maximum value is
	// typically 1000; values above this will be coerced to the maximum.
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// A page token, received from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	paramObj
}

func (r RetrievalFindParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalFindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalFindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrievalFindParams]'s query parameters as `url.Values`.
func (r RetrievalFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrievalGrepParams struct {
	// ID of the file to grep.
	FileID string `json:"file_id" api:"required"`
	// ID of the index the file belongs to.
	IndexID string `json:"index_id" api:"required"`
	// Regex pattern to search for.
	Pattern        string            `json:"pattern" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Number of characters of context to include before and after the matched pattern
	// in the content field of the response
	ContextChars param.Opt[int64] `json:"context_chars,omitzero"`
	// The maximum number of items to return. The service may return fewer than this
	// value. If unspecified, a default page size will be used. The maximum value is
	// typically 1000; values above this will be coerced to the maximum.
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// A page token, received from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	paramObj
}

func (r RetrievalGrepParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalGrepParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalGrepParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrievalGrepParams]'s query parameters as `url.Values`.
func (r RetrievalGrepParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrievalReadParams struct {
	// ID of the file to read.
	FileID string `json:"file_id" api:"required"`
	// ID of the index the file belongs to.
	IndexID        string            `json:"index_id" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Maximum number of characters to read from the offset.
	MaxLength param.Opt[int64] `json:"max_length,omitzero"`
	// Starting character offset.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	paramObj
}

func (r RetrievalReadParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalReadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrievalReadParams]'s query parameters as `url.Values`.
func (r RetrievalReadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
