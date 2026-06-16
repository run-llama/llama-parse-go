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
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
	"github.com/run-llama/llama-parse-go/shared"
)

// PipelineService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	options     []option.RequestOption
	Sync        PipelineSyncService
	DataSources PipelineDataSourceService
	Images      PipelineImageService
	Files       PipelineFileService
	Metadata    PipelineMetadataService
	Documents   PipelineDocumentService
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r PipelineService) {
	r = PipelineService{}
	r.options = opts
	r.Sync = NewPipelineSyncService(opts...)
	r.DataSources = NewPipelineDataSourceService(opts...)
	r.Images = NewPipelineImageService(opts...)
	r.Files = NewPipelineFileService(opts...)
	r.Metadata = NewPipelineMetadataService(opts...)
	r.Documents = NewPipelineDocumentService(opts...)
	return
}

// Create a new managed ingestion pipeline.
//
// A pipeline connects data sources to a vector store for RAG. After creation, call
// `POST /pipelines/{id}/sync` to start ingesting documents.
func (r *PipelineService) New(ctx context.Context, params PipelineNewParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/pipelines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update an existing pipeline's configuration.
func (r *PipelineService) Update(ctx context.Context, pipelineID string, body PipelineUpdateParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Search for pipelines by name, type, or project.
func (r *PipelineService) List(ctx context.Context, query PipelineListParams, opts ...option.RequestOption) (res *[]Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/pipelines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a pipeline and all associated resources.
//
// Removes pipeline files, data sources, and vector store data. This operation is
// irreversible.
func (r *PipelineService) Delete(ctx context.Context, pipelineID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a pipeline by ID.
func (r *PipelineService) Get(ctx context.Context, pipelineID string, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the ingestion status of a managed pipeline.
//
// Returns document counts, sync progress, and the last effective timestamp. Only
// available for managed pipelines.
func (r *PipelineService) GetStatus(ctx context.Context, pipelineID string, query PipelineGetStatusParams, opts ...option.RequestOption) (res *ManagedIngestionStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/status", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Run a retrieval query against a managed pipeline.
//
// Searches the pipeline's vector store using the provided query and retrieval
// parameters. Supports dense, sparse, and hybrid search modes with configurable
// top-k and reranking.
func (r *PipelineService) RunSearch(ctx context.Context, pipelineID string, params PipelineRunSearchParams, opts ...option.RequestOption) (res *PipelineRunSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/retrieve", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Upsert a pipeline.
//
// Updates the pipeline if one with the same name and project already exists,
// otherwise creates a new one.
func (r *PipelineService) Upsert(ctx context.Context, params PipelineUpsertParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/pipelines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type AdvancedModeTransformConfig struct {
	// Configuration for the chunking.
	ChunkingConfig AdvancedModeTransformConfigChunkingConfigUnion `json:"chunking_config"`
	// Any of "advanced".
	Mode AdvancedModeTransformConfigMode `json:"mode"`
	// Configuration for the segmentation.
	SegmentationConfig AdvancedModeTransformConfigSegmentationConfigUnion `json:"segmentation_config"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkingConfig     respjson.Field
		Mode               respjson.Field
		SegmentationConfig respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfig) RawJSON() string { return r.JSON.raw }
func (r *AdvancedModeTransformConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AdvancedModeTransformConfig to a
// AdvancedModeTransformConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AdvancedModeTransformConfigParam.Overrides()
func (r AdvancedModeTransformConfig) ToParam() AdvancedModeTransformConfigParam {
	return param.Override[AdvancedModeTransformConfigParam](json.RawMessage(r.RawJSON()))
}

// AdvancedModeTransformConfigChunkingConfigUnion contains all possible properties
// and values from [AdvancedModeTransformConfigChunkingConfigNoneChunkingConfig],
// [AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfig],
// [AdvancedModeTransformConfigChunkingConfigTokenChunkingConfig],
// [AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig],
// [AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AdvancedModeTransformConfigChunkingConfigUnion struct {
	Mode         string `json:"mode"`
	ChunkOverlap int64  `json:"chunk_overlap"`
	ChunkSize    int64  `json:"chunk_size"`
	Separator    string `json:"separator"`
	// This field is from variant
	// [AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig].
	ParagraphSeparator string `json:"paragraph_separator"`
	// This field is from variant
	// [AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig].
	BreakpointPercentileThreshold int64 `json:"breakpoint_percentile_threshold"`
	// This field is from variant
	// [AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig].
	BufferSize int64 `json:"buffer_size"`
	JSON       struct {
		Mode                          respjson.Field
		ChunkOverlap                  respjson.Field
		ChunkSize                     respjson.Field
		Separator                     respjson.Field
		ParagraphSeparator            respjson.Field
		BreakpointPercentileThreshold respjson.Field
		BufferSize                    respjson.Field
		raw                           string
	} `json:"-"`
}

func (u AdvancedModeTransformConfigChunkingConfigUnion) AsNoneChunkingConfig() (v AdvancedModeTransformConfigChunkingConfigNoneChunkingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigChunkingConfigUnion) AsCharacterChunkingConfig() (v AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigChunkingConfigUnion) AsTokenChunkingConfig() (v AdvancedModeTransformConfigChunkingConfigTokenChunkingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigChunkingConfigUnion) AsSentenceChunkingConfig() (v AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigChunkingConfigUnion) AsSemanticChunkingConfig() (v AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AdvancedModeTransformConfigChunkingConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *AdvancedModeTransformConfigChunkingConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigChunkingConfigNoneChunkingConfig struct {
	// Any of "none".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigChunkingConfigNoneChunkingConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigChunkingConfigNoneChunkingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfig struct {
	ChunkOverlap int64 `json:"chunk_overlap"`
	ChunkSize    int64 `json:"chunk_size"`
	// Any of "character".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlap respjson.Field
		ChunkSize    respjson.Field
		Mode         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigChunkingConfigTokenChunkingConfig struct {
	ChunkOverlap int64 `json:"chunk_overlap"`
	ChunkSize    int64 `json:"chunk_size"`
	// Any of "token".
	Mode      string `json:"mode"`
	Separator string `json:"separator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlap respjson.Field
		ChunkSize    respjson.Field
		Mode         respjson.Field
		Separator    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigChunkingConfigTokenChunkingConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigChunkingConfigTokenChunkingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig struct {
	ChunkOverlap int64 `json:"chunk_overlap"`
	ChunkSize    int64 `json:"chunk_size"`
	// Any of "sentence".
	Mode               string `json:"mode"`
	ParagraphSeparator string `json:"paragraph_separator"`
	Separator          string `json:"separator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlap       respjson.Field
		ChunkSize          respjson.Field
		Mode               respjson.Field
		ParagraphSeparator respjson.Field
		Separator          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig struct {
	BreakpointPercentileThreshold int64 `json:"breakpoint_percentile_threshold"`
	BufferSize                    int64 `json:"buffer_size"`
	// Any of "semantic".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BreakpointPercentileThreshold respjson.Field
		BufferSize                    respjson.Field
		Mode                          respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigMode string

const (
	AdvancedModeTransformConfigModeAdvanced AdvancedModeTransformConfigMode = "advanced"
)

// AdvancedModeTransformConfigSegmentationConfigUnion contains all possible
// properties and values from
// [AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfig],
// [AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig],
// [AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AdvancedModeTransformConfigSegmentationConfigUnion struct {
	Mode string `json:"mode"`
	// This field is from variant
	// [AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig].
	PageSeparator string `json:"page_separator"`
	JSON          struct {
		Mode          respjson.Field
		PageSeparator respjson.Field
		raw           string
	} `json:"-"`
}

func (u AdvancedModeTransformConfigSegmentationConfigUnion) AsNoneSegmentationConfig() (v AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigSegmentationConfigUnion) AsPageSegmentationConfig() (v AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AdvancedModeTransformConfigSegmentationConfigUnion) AsElementSegmentationConfig() (v AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AdvancedModeTransformConfigSegmentationConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *AdvancedModeTransformConfigSegmentationConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfig struct {
	// Any of "none".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig struct {
	// Any of "page".
	Mode          string `json:"mode"`
	PageSeparator string `json:"page_separator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode          respjson.Field
		PageSeparator respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfig struct {
	// Any of "element".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdvancedModeTransformConfigParam struct {
	// Configuration for the chunking.
	ChunkingConfig AdvancedModeTransformConfigChunkingConfigUnionParam `json:"chunking_config,omitzero"`
	// Any of "advanced".
	Mode AdvancedModeTransformConfigMode `json:"mode,omitzero"`
	// Configuration for the segmentation.
	SegmentationConfig AdvancedModeTransformConfigSegmentationConfigUnionParam `json:"segmentation_config,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdvancedModeTransformConfigChunkingConfigUnionParam struct {
	OfNoneChunkingConfig      *AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam      `json:",omitzero,inline"`
	OfCharacterChunkingConfig *AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam `json:",omitzero,inline"`
	OfTokenChunkingConfig     *AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam     `json:",omitzero,inline"`
	OfSentenceChunkingConfig  *AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam  `json:",omitzero,inline"`
	OfSemanticChunkingConfig  *AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam  `json:",omitzero,inline"`
	paramUnion
}

func (u AdvancedModeTransformConfigChunkingConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoneChunkingConfig,
		u.OfCharacterChunkingConfig,
		u.OfTokenChunkingConfig,
		u.OfSentenceChunkingConfig,
		u.OfSemanticChunkingConfig)
}
func (u *AdvancedModeTransformConfigChunkingConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam struct {
	// Any of "none".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigChunkingConfigNoneChunkingConfigParam](
		"mode", "none",
	)
}

type AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam struct {
	ChunkOverlap param.Opt[int64] `json:"chunk_overlap,omitzero"`
	ChunkSize    param.Opt[int64] `json:"chunk_size,omitzero"`
	// Any of "character".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigChunkingConfigCharacterChunkingConfigParam](
		"mode", "character",
	)
}

type AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam struct {
	ChunkOverlap param.Opt[int64]  `json:"chunk_overlap,omitzero"`
	ChunkSize    param.Opt[int64]  `json:"chunk_size,omitzero"`
	Separator    param.Opt[string] `json:"separator,omitzero"`
	// Any of "token".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigChunkingConfigTokenChunkingConfigParam](
		"mode", "token",
	)
}

type AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam struct {
	ChunkOverlap       param.Opt[int64]  `json:"chunk_overlap,omitzero"`
	ChunkSize          param.Opt[int64]  `json:"chunk_size,omitzero"`
	ParagraphSeparator param.Opt[string] `json:"paragraph_separator,omitzero"`
	Separator          param.Opt[string] `json:"separator,omitzero"`
	// Any of "sentence".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigChunkingConfigSentenceChunkingConfigParam](
		"mode", "sentence",
	)
}

type AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam struct {
	BreakpointPercentileThreshold param.Opt[int64] `json:"breakpoint_percentile_threshold,omitzero"`
	BufferSize                    param.Opt[int64] `json:"buffer_size,omitzero"`
	// Any of "semantic".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigChunkingConfigSemanticChunkingConfigParam](
		"mode", "semantic",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdvancedModeTransformConfigSegmentationConfigUnionParam struct {
	OfNoneSegmentationConfig    *AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam    `json:",omitzero,inline"`
	OfPageSegmentationConfig    *AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam    `json:",omitzero,inline"`
	OfElementSegmentationConfig *AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u AdvancedModeTransformConfigSegmentationConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoneSegmentationConfig, u.OfPageSegmentationConfig, u.OfElementSegmentationConfig)
}
func (u *AdvancedModeTransformConfigSegmentationConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam struct {
	// Any of "none".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigSegmentationConfigNoneSegmentationConfigParam](
		"mode", "none",
	)
}

type AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam struct {
	PageSeparator param.Opt[string] `json:"page_separator,omitzero"`
	// Any of "page".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigSegmentationConfigPageSegmentationConfigParam](
		"mode", "page",
	)
}

type AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam struct {
	// Any of "element".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AdvancedModeTransformConfigSegmentationConfigElementSegmentationConfigParam](
		"mode", "element",
	)
}

type AutoTransformConfig struct {
	// Chunk overlap for the transformation.
	ChunkOverlap int64 `json:"chunk_overlap"`
	// Chunk size for the transformation.
	ChunkSize int64 `json:"chunk_size"`
	// Any of "auto".
	Mode AutoTransformConfigMode `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkOverlap respjson.Field
		ChunkSize    respjson.Field
		Mode         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransformConfig) RawJSON() string { return r.JSON.raw }
func (r *AutoTransformConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AutoTransformConfig to a AutoTransformConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AutoTransformConfigParam.Overrides()
func (r AutoTransformConfig) ToParam() AutoTransformConfigParam {
	return param.Override[AutoTransformConfigParam](json.RawMessage(r.RawJSON()))
}

type AutoTransformConfigMode string

const (
	AutoTransformConfigModeAuto AutoTransformConfigMode = "auto"
)

type AutoTransformConfigParam struct {
	// Chunk overlap for the transformation.
	ChunkOverlap param.Opt[int64] `json:"chunk_overlap,omitzero"`
	// Chunk size for the transformation.
	ChunkSize param.Opt[int64] `json:"chunk_size,omitzero"`
	// Any of "auto".
	Mode AutoTransformConfigMode `json:"mode,omitzero"`
	paramObj
}

func (r AutoTransformConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoTransformConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoTransformConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AzureOpenAIEmbedding struct {
	// Additional kwargs for the OpenAI API.
	AdditionalKwargs map[string]any `json:"additional_kwargs"`
	// The base URL for Azure deployment.
	APIBase string `json:"api_base"`
	// The OpenAI API key.
	APIKey string `json:"api_key" api:"nullable"`
	// The version for Azure OpenAI API.
	APIVersion string `json:"api_version"`
	// The Azure deployment to use.
	AzureDeployment string `json:"azure_deployment" api:"nullable"`
	// The Azure endpoint to use.
	AzureEndpoint string `json:"azure_endpoint" api:"nullable"`
	ClassName     string `json:"class_name"`
	// The default headers for API requests.
	DefaultHeaders map[string]string `json:"default_headers" api:"nullable"`
	// The number of dimensions on the output embedding vectors. Works only with v3
	// embedding models.
	Dimensions int64 `json:"dimensions" api:"nullable"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// Maximum number of retries.
	MaxRetries int64 `json:"max_retries"`
	// The name of the OpenAI embedding model.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// Reuse the OpenAI client between requests. When doing anything with large volumes
	// of async API calls, setting this to false can improve stability.
	ReuseClient bool `json:"reuse_client"`
	// Timeout for each request.
	Timeout float64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalKwargs respjson.Field
		APIBase          respjson.Field
		APIKey           respjson.Field
		APIVersion       respjson.Field
		AzureDeployment  respjson.Field
		AzureEndpoint    respjson.Field
		ClassName        respjson.Field
		DefaultHeaders   respjson.Field
		Dimensions       respjson.Field
		EmbedBatchSize   respjson.Field
		MaxRetries       respjson.Field
		ModelName        respjson.Field
		NumWorkers       respjson.Field
		ReuseClient      respjson.Field
		Timeout          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AzureOpenAIEmbedding) RawJSON() string { return r.JSON.raw }
func (r *AzureOpenAIEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AzureOpenAIEmbedding to a AzureOpenAIEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AzureOpenAIEmbeddingParam.Overrides()
func (r AzureOpenAIEmbedding) ToParam() AzureOpenAIEmbeddingParam {
	return param.Override[AzureOpenAIEmbeddingParam](json.RawMessage(r.RawJSON()))
}

type AzureOpenAIEmbeddingParam struct {
	// The OpenAI API key.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// The Azure deployment to use.
	AzureDeployment param.Opt[string] `json:"azure_deployment,omitzero"`
	// The Azure endpoint to use.
	AzureEndpoint param.Opt[string] `json:"azure_endpoint,omitzero"`
	// The number of dimensions on the output embedding vectors. Works only with v3
	// embedding models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64] `json:"num_workers,omitzero"`
	// The base URL for Azure deployment.
	APIBase param.Opt[string] `json:"api_base,omitzero"`
	// The version for Azure OpenAI API.
	APIVersion param.Opt[string] `json:"api_version,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// Maximum number of retries.
	MaxRetries param.Opt[int64] `json:"max_retries,omitzero"`
	// The name of the OpenAI embedding model.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// Reuse the OpenAI client between requests. When doing anything with large volumes
	// of async API calls, setting this to false can improve stability.
	ReuseClient param.Opt[bool] `json:"reuse_client,omitzero"`
	// Timeout for each request.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// The default headers for API requests.
	DefaultHeaders map[string]string `json:"default_headers,omitzero"`
	// Additional kwargs for the OpenAI API.
	AdditionalKwargs map[string]any `json:"additional_kwargs,omitzero"`
	paramObj
}

func (r AzureOpenAIEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow AzureOpenAIEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AzureOpenAIEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AzureOpenAIEmbeddingConfig struct {
	// Configuration for the Azure OpenAI embedding model.
	Component AzureOpenAIEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "AZURE_EMBEDDING".
	Type AzureOpenAIEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AzureOpenAIEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *AzureOpenAIEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AzureOpenAIEmbeddingConfig to a
// AzureOpenAIEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AzureOpenAIEmbeddingConfigParam.Overrides()
func (r AzureOpenAIEmbeddingConfig) ToParam() AzureOpenAIEmbeddingConfigParam {
	return param.Override[AzureOpenAIEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type AzureOpenAIEmbeddingConfigType string

const (
	AzureOpenAIEmbeddingConfigTypeAzureEmbedding AzureOpenAIEmbeddingConfigType = "AZURE_EMBEDDING"
)

type AzureOpenAIEmbeddingConfigParam struct {
	// Configuration for the Azure OpenAI embedding model.
	Component AzureOpenAIEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "AZURE_EMBEDDING".
	Type AzureOpenAIEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r AzureOpenAIEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AzureOpenAIEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AzureOpenAIEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BedrockEmbedding struct {
	// Additional kwargs for the bedrock client.
	AdditionalKwargs map[string]any `json:"additional_kwargs"`
	// AWS Access Key ID to use
	AwsAccessKeyID string `json:"aws_access_key_id" api:"nullable"`
	// AWS Secret Access Key to use
	AwsSecretAccessKey string `json:"aws_secret_access_key" api:"nullable"`
	// AWS Session Token to use
	AwsSessionToken string `json:"aws_session_token" api:"nullable"`
	ClassName       string `json:"class_name"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// The maximum number of API retries.
	MaxRetries int64 `json:"max_retries"`
	// The modelId of the Bedrock model to use.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// The name of aws profile to use. If not given, then the default profile is used.
	ProfileName string `json:"profile_name" api:"nullable"`
	// AWS region name to use. Uses region configured in AWS CLI if not passed
	RegionName string `json:"region_name" api:"nullable"`
	// The timeout for the Bedrock API request in seconds. It will be used for both
	// connect and read timeouts.
	Timeout float64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalKwargs   respjson.Field
		AwsAccessKeyID     respjson.Field
		AwsSecretAccessKey respjson.Field
		AwsSessionToken    respjson.Field
		ClassName          respjson.Field
		EmbedBatchSize     respjson.Field
		MaxRetries         respjson.Field
		ModelName          respjson.Field
		NumWorkers         respjson.Field
		ProfileName        respjson.Field
		RegionName         respjson.Field
		Timeout            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BedrockEmbedding) RawJSON() string { return r.JSON.raw }
func (r *BedrockEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BedrockEmbedding to a BedrockEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BedrockEmbeddingParam.Overrides()
func (r BedrockEmbedding) ToParam() BedrockEmbeddingParam {
	return param.Override[BedrockEmbeddingParam](json.RawMessage(r.RawJSON()))
}

type BedrockEmbeddingParam struct {
	// AWS Access Key ID to use
	AwsAccessKeyID param.Opt[string] `json:"aws_access_key_id,omitzero"`
	// AWS Secret Access Key to use
	AwsSecretAccessKey param.Opt[string] `json:"aws_secret_access_key,omitzero"`
	// AWS Session Token to use
	AwsSessionToken param.Opt[string] `json:"aws_session_token,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64] `json:"num_workers,omitzero"`
	// The name of aws profile to use. If not given, then the default profile is used.
	ProfileName param.Opt[string] `json:"profile_name,omitzero"`
	// AWS region name to use. Uses region configured in AWS CLI if not passed
	RegionName param.Opt[string] `json:"region_name,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// The maximum number of API retries.
	MaxRetries param.Opt[int64] `json:"max_retries,omitzero"`
	// The modelId of the Bedrock model to use.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// The timeout for the Bedrock API request in seconds. It will be used for both
	// connect and read timeouts.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Additional kwargs for the bedrock client.
	AdditionalKwargs map[string]any `json:"additional_kwargs,omitzero"`
	paramObj
}

func (r BedrockEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow BedrockEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BedrockEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BedrockEmbeddingConfig struct {
	// Configuration for the Bedrock embedding model.
	Component BedrockEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "BEDROCK_EMBEDDING".
	Type BedrockEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BedrockEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *BedrockEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BedrockEmbeddingConfig to a BedrockEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BedrockEmbeddingConfigParam.Overrides()
func (r BedrockEmbeddingConfig) ToParam() BedrockEmbeddingConfigParam {
	return param.Override[BedrockEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type BedrockEmbeddingConfigType string

const (
	BedrockEmbeddingConfigTypeBedrockEmbedding BedrockEmbeddingConfigType = "BEDROCK_EMBEDDING"
)

type BedrockEmbeddingConfigParam struct {
	// Configuration for the Bedrock embedding model.
	Component BedrockEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "BEDROCK_EMBEDDING".
	Type BedrockEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r BedrockEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow BedrockEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BedrockEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CohereEmbedding struct {
	// The Cohere API key.
	APIKey    string `json:"api_key" api:"required"`
	ClassName string `json:"class_name"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// Embedding type. If not provided float embedding_type is used when needed.
	EmbeddingType string `json:"embedding_type"`
	// Model Input type. If not provided, search_document and search_query are used
	// when needed.
	InputType string `json:"input_type" api:"nullable"`
	// The modelId of the Cohere model to use.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// Truncation type - START/ END/ NONE
	Truncate string `json:"truncate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey         respjson.Field
		ClassName      respjson.Field
		EmbedBatchSize respjson.Field
		EmbeddingType  respjson.Field
		InputType      respjson.Field
		ModelName      respjson.Field
		NumWorkers     respjson.Field
		Truncate       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CohereEmbedding) RawJSON() string { return r.JSON.raw }
func (r *CohereEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CohereEmbedding to a CohereEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CohereEmbeddingParam.Overrides()
func (r CohereEmbedding) ToParam() CohereEmbeddingParam {
	return param.Override[CohereEmbeddingParam](json.RawMessage(r.RawJSON()))
}

// The property APIKey is required.
type CohereEmbeddingParam struct {
	// The Cohere API key.
	APIKey param.Opt[string] `json:"api_key,omitzero" api:"required"`
	// Model Input type. If not provided, search_document and search_query are used
	// when needed.
	InputType param.Opt[string] `json:"input_type,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64]  `json:"num_workers,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// Embedding type. If not provided float embedding_type is used when needed.
	EmbeddingType param.Opt[string] `json:"embedding_type,omitzero"`
	// The modelId of the Cohere model to use.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// Truncation type - START/ END/ NONE
	Truncate param.Opt[string] `json:"truncate,omitzero"`
	paramObj
}

func (r CohereEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow CohereEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CohereEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CohereEmbeddingConfig struct {
	// Configuration for the Cohere embedding model.
	Component CohereEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "COHERE_EMBEDDING".
	Type CohereEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CohereEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *CohereEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CohereEmbeddingConfig to a CohereEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CohereEmbeddingConfigParam.Overrides()
func (r CohereEmbeddingConfig) ToParam() CohereEmbeddingConfigParam {
	return param.Override[CohereEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type CohereEmbeddingConfigType string

const (
	CohereEmbeddingConfigTypeCohereEmbedding CohereEmbeddingConfigType = "COHERE_EMBEDDING"
)

type CohereEmbeddingConfigParam struct {
	// Configuration for the Cohere embedding model.
	Component CohereEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "COHERE_EMBEDDING".
	Type CohereEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r CohereEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CohereEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CohereEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating a data sink.
//
// The properties Component, Name, SinkType are required.
type DataSinkCreateParam struct {
	// Component that implements the data sink
	Component DataSinkCreateComponentUnionParam `json:"component,omitzero" api:"required"`
	// The name of the data sink.
	Name string `json:"name" api:"required"`
	// Any of "PINECONE", "POSTGRES", "QDRANT", "AZUREAI_SEARCH", "MONGODB_ATLAS",
	// "MILVUS", "ASTRA_DB".
	SinkType DataSinkCreateSinkType `json:"sink_type,omitzero" api:"required"`
	paramObj
}

func (r DataSinkCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow DataSinkCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSinkCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSinkCreateComponentUnionParam struct {
	OfAnyMap                        map[string]any                             `json:",omitzero,inline"`
	OfCloudPineconeVectorStore      *shared.CloudPineconeVectorStoreParam      `json:",omitzero,inline"`
	OfCloudPostgresVectorStore      *shared.CloudPostgresVectorStoreParam      `json:",omitzero,inline"`
	OfCloudQdrantVectorStore        *shared.CloudQdrantVectorStoreParam        `json:",omitzero,inline"`
	OfCloudAzureAISearchVectorStore *shared.CloudAzureAISearchVectorStoreParam `json:",omitzero,inline"`
	OfCloudMongoDBAtlasVectorSearch *shared.CloudMongoDBAtlasVectorSearchParam `json:",omitzero,inline"`
	OfCloudMilvusVectorStore        *shared.CloudMilvusVectorStoreParam        `json:",omitzero,inline"`
	OfCloudAstraDBVectorStore       *shared.CloudAstraDBVectorStoreParam       `json:",omitzero,inline"`
	paramUnion
}

func (u DataSinkCreateComponentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfCloudPineconeVectorStore,
		u.OfCloudPostgresVectorStore,
		u.OfCloudQdrantVectorStore,
		u.OfCloudAzureAISearchVectorStore,
		u.OfCloudMongoDBAtlasVectorSearch,
		u.OfCloudMilvusVectorStore,
		u.OfCloudAstraDBVectorStore)
}
func (u *DataSinkCreateComponentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type DataSinkCreateSinkType string

const (
	DataSinkCreateSinkTypePinecone      DataSinkCreateSinkType = "PINECONE"
	DataSinkCreateSinkTypePostgres      DataSinkCreateSinkType = "POSTGRES"
	DataSinkCreateSinkTypeQdrant        DataSinkCreateSinkType = "QDRANT"
	DataSinkCreateSinkTypeAzureaiSearch DataSinkCreateSinkType = "AZUREAI_SEARCH"
	DataSinkCreateSinkTypeMongoDBAtlas  DataSinkCreateSinkType = "MONGODB_ATLAS"
	DataSinkCreateSinkTypeMilvus        DataSinkCreateSinkType = "MILVUS"
	DataSinkCreateSinkTypeAstraDB       DataSinkCreateSinkType = "ASTRA_DB"
)

type GeminiEmbedding struct {
	// API base to access the model. Defaults to None.
	APIBase string `json:"api_base" api:"nullable"`
	// API key to access the model. Defaults to None.
	APIKey    string `json:"api_key" api:"nullable"`
	ClassName string `json:"class_name"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// The modelId of the Gemini model to use.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// Optional reduced dimension for output embeddings. Supported by
	// models/text-embedding-004 and newer (e.g. gemini-embedding-001). Not supported
	// by models/embedding-001.
	OutputDimensionality int64 `json:"output_dimensionality" api:"nullable"`
	// The task for embedding model.
	TaskType string `json:"task_type" api:"nullable"`
	// Title is only applicable for retrieval_document tasks, and is used to represent
	// a document title. For other tasks, title is invalid.
	Title string `json:"title" api:"nullable"`
	// Transport to access the model. Defaults to None.
	Transport string `json:"transport" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIBase              respjson.Field
		APIKey               respjson.Field
		ClassName            respjson.Field
		EmbedBatchSize       respjson.Field
		ModelName            respjson.Field
		NumWorkers           respjson.Field
		OutputDimensionality respjson.Field
		TaskType             respjson.Field
		Title                respjson.Field
		Transport            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeminiEmbedding) RawJSON() string { return r.JSON.raw }
func (r *GeminiEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GeminiEmbedding to a GeminiEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GeminiEmbeddingParam.Overrides()
func (r GeminiEmbedding) ToParam() GeminiEmbeddingParam {
	return param.Override[GeminiEmbeddingParam](json.RawMessage(r.RawJSON()))
}

type GeminiEmbeddingParam struct {
	// API base to access the model. Defaults to None.
	APIBase param.Opt[string] `json:"api_base,omitzero"`
	// API key to access the model. Defaults to None.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64] `json:"num_workers,omitzero"`
	// Optional reduced dimension for output embeddings. Supported by
	// models/text-embedding-004 and newer (e.g. gemini-embedding-001). Not supported
	// by models/embedding-001.
	OutputDimensionality param.Opt[int64] `json:"output_dimensionality,omitzero"`
	// The task for embedding model.
	TaskType param.Opt[string] `json:"task_type,omitzero"`
	// Title is only applicable for retrieval_document tasks, and is used to represent
	// a document title. For other tasks, title is invalid.
	Title param.Opt[string] `json:"title,omitzero"`
	// Transport to access the model. Defaults to None.
	Transport param.Opt[string] `json:"transport,omitzero"`
	ClassName param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// The modelId of the Gemini model to use.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	paramObj
}

func (r GeminiEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow GeminiEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeminiEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeminiEmbeddingConfig struct {
	// Configuration for the Gemini embedding model.
	Component GeminiEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "GEMINI_EMBEDDING".
	Type GeminiEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeminiEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *GeminiEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GeminiEmbeddingConfig to a GeminiEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GeminiEmbeddingConfigParam.Overrides()
func (r GeminiEmbeddingConfig) ToParam() GeminiEmbeddingConfigParam {
	return param.Override[GeminiEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type GeminiEmbeddingConfigType string

const (
	GeminiEmbeddingConfigTypeGeminiEmbedding GeminiEmbeddingConfigType = "GEMINI_EMBEDDING"
)

type GeminiEmbeddingConfigParam struct {
	// Configuration for the Gemini embedding model.
	Component GeminiEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "GEMINI_EMBEDDING".
	Type GeminiEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r GeminiEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow GeminiEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeminiEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HuggingFaceInferenceAPIEmbedding struct {
	// Hugging Face token. Will default to the locally saved token. Pass token=False if
	// you don’t want to send your token to the server.
	Token     HuggingFaceInferenceAPIEmbeddingTokenUnion `json:"token" api:"nullable"`
	ClassName string                                     `json:"class_name"`
	// Additional cookies to send to the server.
	Cookies map[string]string `json:"cookies" api:"nullable"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// Additional headers to send to the server. By default only the authorization and
	// user-agent headers are sent. Values in this dictionary will override the default
	// values.
	Headers map[string]string `json:"headers" api:"nullable"`
	// Hugging Face model name. If None, the task will be used.
	ModelName string `json:"model_name" api:"nullable"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// Enum of possible pooling choices with pooling behaviors.
	//
	// Any of "cls", "mean", "last".
	Pooling HuggingFaceInferenceAPIEmbeddingPooling `json:"pooling" api:"nullable"`
	// Instruction to prepend during query embedding.
	QueryInstruction string `json:"query_instruction" api:"nullable"`
	// Optional task to pick Hugging Face's recommended model, used when model_name is
	// left as default of None.
	Task string `json:"task" api:"nullable"`
	// Instruction to prepend during text embedding.
	TextInstruction string `json:"text_instruction" api:"nullable"`
	// The maximum number of seconds to wait for a response from the server. Loading a
	// new model in Inference API can take up to several minutes. Defaults to None,
	// meaning it will loop until the server is available.
	Timeout float64 `json:"timeout" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token            respjson.Field
		ClassName        respjson.Field
		Cookies          respjson.Field
		EmbedBatchSize   respjson.Field
		Headers          respjson.Field
		ModelName        respjson.Field
		NumWorkers       respjson.Field
		Pooling          respjson.Field
		QueryInstruction respjson.Field
		Task             respjson.Field
		TextInstruction  respjson.Field
		Timeout          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HuggingFaceInferenceAPIEmbedding) RawJSON() string { return r.JSON.raw }
func (r *HuggingFaceInferenceAPIEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HuggingFaceInferenceAPIEmbedding to a
// HuggingFaceInferenceAPIEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HuggingFaceInferenceAPIEmbeddingParam.Overrides()
func (r HuggingFaceInferenceAPIEmbedding) ToParam() HuggingFaceInferenceAPIEmbeddingParam {
	return param.Override[HuggingFaceInferenceAPIEmbeddingParam](json.RawMessage(r.RawJSON()))
}

// HuggingFaceInferenceAPIEmbeddingTokenUnion contains all possible properties and
// values from [string], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type HuggingFaceInferenceAPIEmbeddingTokenUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u HuggingFaceInferenceAPIEmbeddingTokenUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HuggingFaceInferenceAPIEmbeddingTokenUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u HuggingFaceInferenceAPIEmbeddingTokenUnion) RawJSON() string { return u.JSON.raw }

func (r *HuggingFaceInferenceAPIEmbeddingTokenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enum of possible pooling choices with pooling behaviors.
type HuggingFaceInferenceAPIEmbeddingPooling string

const (
	HuggingFaceInferenceAPIEmbeddingPoolingCls  HuggingFaceInferenceAPIEmbeddingPooling = "cls"
	HuggingFaceInferenceAPIEmbeddingPoolingMean HuggingFaceInferenceAPIEmbeddingPooling = "mean"
	HuggingFaceInferenceAPIEmbeddingPoolingLast HuggingFaceInferenceAPIEmbeddingPooling = "last"
)

type HuggingFaceInferenceAPIEmbeddingParam struct {
	// Hugging Face model name. If None, the task will be used.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64] `json:"num_workers,omitzero"`
	// Instruction to prepend during query embedding.
	QueryInstruction param.Opt[string] `json:"query_instruction,omitzero"`
	// Optional task to pick Hugging Face's recommended model, used when model_name is
	// left as default of None.
	Task param.Opt[string] `json:"task,omitzero"`
	// Instruction to prepend during text embedding.
	TextInstruction param.Opt[string] `json:"text_instruction,omitzero"`
	// The maximum number of seconds to wait for a response from the server. Loading a
	// new model in Inference API can take up to several minutes. Defaults to None,
	// meaning it will loop until the server is available.
	Timeout   param.Opt[float64] `json:"timeout,omitzero"`
	ClassName param.Opt[string]  `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// Hugging Face token. Will default to the locally saved token. Pass token=False if
	// you don’t want to send your token to the server.
	Token HuggingFaceInferenceAPIEmbeddingTokenUnionParam `json:"token,omitzero"`
	// Additional cookies to send to the server.
	Cookies map[string]string `json:"cookies,omitzero"`
	// Additional headers to send to the server. By default only the authorization and
	// user-agent headers are sent. Values in this dictionary will override the default
	// values.
	Headers map[string]string `json:"headers,omitzero"`
	// Enum of possible pooling choices with pooling behaviors.
	//
	// Any of "cls", "mean", "last".
	Pooling HuggingFaceInferenceAPIEmbeddingPooling `json:"pooling,omitzero"`
	paramObj
}

func (r HuggingFaceInferenceAPIEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow HuggingFaceInferenceAPIEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HuggingFaceInferenceAPIEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type HuggingFaceInferenceAPIEmbeddingTokenUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u HuggingFaceInferenceAPIEmbeddingTokenUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool)
}
func (u *HuggingFaceInferenceAPIEmbeddingTokenUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type HuggingFaceInferenceAPIEmbeddingConfig struct {
	// Configuration for the HuggingFace Inference API embedding model.
	Component HuggingFaceInferenceAPIEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "HUGGINGFACE_API_EMBEDDING".
	Type HuggingFaceInferenceAPIEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HuggingFaceInferenceAPIEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *HuggingFaceInferenceAPIEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HuggingFaceInferenceAPIEmbeddingConfig to a
// HuggingFaceInferenceAPIEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HuggingFaceInferenceAPIEmbeddingConfigParam.Overrides()
func (r HuggingFaceInferenceAPIEmbeddingConfig) ToParam() HuggingFaceInferenceAPIEmbeddingConfigParam {
	return param.Override[HuggingFaceInferenceAPIEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type HuggingFaceInferenceAPIEmbeddingConfigType string

const (
	HuggingFaceInferenceAPIEmbeddingConfigTypeHuggingfaceAPIEmbedding HuggingFaceInferenceAPIEmbeddingConfigType = "HUGGINGFACE_API_EMBEDDING"
)

type HuggingFaceInferenceAPIEmbeddingConfigParam struct {
	// Configuration for the HuggingFace Inference API embedding model.
	Component HuggingFaceInferenceAPIEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "HUGGINGFACE_API_EMBEDDING".
	Type HuggingFaceInferenceAPIEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r HuggingFaceInferenceAPIEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow HuggingFaceInferenceAPIEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HuggingFaceInferenceAPIEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlamaParseParametersResp struct {
	AdaptiveLongTable                        bool    `json:"adaptive_long_table" api:"nullable"`
	AggressiveTableExtraction                bool    `json:"aggressive_table_extraction" api:"nullable"`
	AnnotateLinks                            bool    `json:"annotate_links" api:"nullable"`
	AutoMode                                 bool    `json:"auto_mode" api:"nullable"`
	AutoModeConfigurationJson                string  `json:"auto_mode_configuration_json" api:"nullable"`
	AutoModeTriggerOnImageInPage             bool    `json:"auto_mode_trigger_on_image_in_page" api:"nullable"`
	AutoModeTriggerOnRegexpInPage            string  `json:"auto_mode_trigger_on_regexp_in_page" api:"nullable"`
	AutoModeTriggerOnTableInPage             bool    `json:"auto_mode_trigger_on_table_in_page" api:"nullable"`
	AutoModeTriggerOnTextInPage              string  `json:"auto_mode_trigger_on_text_in_page" api:"nullable"`
	AzureOpenAIAPIVersion                    string  `json:"azure_openai_api_version" api:"nullable"`
	AzureOpenAIDeploymentName                string  `json:"azure_openai_deployment_name" api:"nullable"`
	AzureOpenAIEndpoint                      string  `json:"azure_openai_endpoint" api:"nullable"`
	AzureOpenAIKey                           string  `json:"azure_openai_key" api:"nullable"`
	BboxBottom                               float64 `json:"bbox_bottom" api:"nullable"`
	BboxLeft                                 float64 `json:"bbox_left" api:"nullable"`
	BboxRight                                float64 `json:"bbox_right" api:"nullable"`
	BboxTop                                  float64 `json:"bbox_top" api:"nullable"`
	BoundingBox                              string  `json:"bounding_box" api:"nullable"`
	CompactMarkdownTable                     bool    `json:"compact_markdown_table" api:"nullable"`
	ComplementalFormattingInstruction        string  `json:"complemental_formatting_instruction" api:"nullable"`
	ContentGuidelineInstruction              string  `json:"content_guideline_instruction" api:"nullable"`
	ContinuousMode                           bool    `json:"continuous_mode" api:"nullable"`
	DisableImageExtraction                   bool    `json:"disable_image_extraction" api:"nullable"`
	DisableOcr                               bool    `json:"disable_ocr" api:"nullable"`
	DisableReconstruction                    bool    `json:"disable_reconstruction" api:"nullable"`
	DoNotCache                               bool    `json:"do_not_cache" api:"nullable"`
	DoNotUnrollColumns                       bool    `json:"do_not_unroll_columns" api:"nullable"`
	EnableCostOptimizer                      bool    `json:"enable_cost_optimizer" api:"nullable"`
	ExtractCharts                            bool    `json:"extract_charts" api:"nullable"`
	ExtractLayout                            bool    `json:"extract_layout" api:"nullable"`
	ExtractPrintedPageNumber                 bool    `json:"extract_printed_page_number" api:"nullable"`
	FastMode                                 bool    `json:"fast_mode" api:"nullable"`
	FormattingInstruction                    string  `json:"formatting_instruction" api:"nullable"`
	Gpt4oAPIKey                              string  `json:"gpt4o_api_key" api:"nullable"`
	Gpt4oMode                                bool    `json:"gpt4o_mode" api:"nullable"`
	GuessXlsxSheetName                       bool    `json:"guess_xlsx_sheet_name" api:"nullable"`
	HideFooters                              bool    `json:"hide_footers" api:"nullable"`
	HideHeaders                              bool    `json:"hide_headers" api:"nullable"`
	HighResOcr                               bool    `json:"high_res_ocr" api:"nullable"`
	HTMLMakeAllElementsVisible               bool    `json:"html_make_all_elements_visible" api:"nullable"`
	HTMLRemoveFixedElements                  bool    `json:"html_remove_fixed_elements" api:"nullable"`
	HTMLRemoveNavigationElements             bool    `json:"html_remove_navigation_elements" api:"nullable"`
	HTTPProxy                                string  `json:"http_proxy" api:"nullable"`
	IgnoreDocumentElementsForLayoutDetection bool    `json:"ignore_document_elements_for_layout_detection" api:"nullable"`
	// Any of "screenshot", "embedded", "layout".
	ImagesToSave                          []string           `json:"images_to_save" api:"nullable"`
	InlineImagesInMarkdown                bool               `json:"inline_images_in_markdown" api:"nullable"`
	InputS3Path                           string             `json:"input_s3_path" api:"nullable"`
	InputS3Region                         string             `json:"input_s3_region" api:"nullable"`
	InputURL                              string             `json:"input_url" api:"nullable"`
	InternalIsScreenshotJob               bool               `json:"internal_is_screenshot_job" api:"nullable"`
	InvalidateCache                       bool               `json:"invalidate_cache" api:"nullable"`
	IsFormattingInstruction               bool               `json:"is_formatting_instruction" api:"nullable"`
	JobTimeoutExtraTimePerPageInSeconds   float64            `json:"job_timeout_extra_time_per_page_in_seconds" api:"nullable"`
	JobTimeoutInSeconds                   float64            `json:"job_timeout_in_seconds" api:"nullable"`
	KeepPageSeparatorWhenMergingTables    bool               `json:"keep_page_separator_when_merging_tables" api:"nullable"`
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
	OutputS3PathPrefix                    string             `json:"output_s3_path_prefix" api:"nullable"`
	OutputS3Region                        string             `json:"output_s3_region" api:"nullable"`
	OutputTablesAsHTML                    bool               `json:"output_tables_as_HTML" api:"nullable"`
	PageErrorTolerance                    float64            `json:"page_error_tolerance" api:"nullable"`
	PageFooterPrefix                      string             `json:"page_footer_prefix" api:"nullable"`
	PageFooterSuffix                      string             `json:"page_footer_suffix" api:"nullable"`
	PageHeaderPrefix                      string             `json:"page_header_prefix" api:"nullable"`
	PageHeaderSuffix                      string             `json:"page_header_suffix" api:"nullable"`
	PagePrefix                            string             `json:"page_prefix" api:"nullable"`
	PageSeparator                         string             `json:"page_separator" api:"nullable"`
	PageSuffix                            string             `json:"page_suffix" api:"nullable"`
	// Enum for representing the mode of parsing to be used.
	//
	// Any of "parse_page_without_llm", "parse_page_with_llm", "parse_page_with_lvm",
	// "parse_page_with_agent", "parse_page_with_layout_agent",
	// "parse_document_with_llm", "parse_document_with_lvm",
	// "parse_document_with_agent".
	ParseMode                          ParsingMode `json:"parse_mode" api:"nullable"`
	ParsingInstruction                 string      `json:"parsing_instruction" api:"nullable"`
	PreciseBoundingBox                 bool        `json:"precise_bounding_box" api:"nullable"`
	PremiumMode                        bool        `json:"premium_mode" api:"nullable"`
	PresentationOutOfBoundsContent     bool        `json:"presentation_out_of_bounds_content" api:"nullable"`
	PresentationSkipEmbeddedData       bool        `json:"presentation_skip_embedded_data" api:"nullable"`
	PreserveLayoutAlignmentAcrossPages bool        `json:"preserve_layout_alignment_across_pages" api:"nullable"`
	PreserveVerySmallText              bool        `json:"preserve_very_small_text" api:"nullable"`
	Preset                             string      `json:"preset" api:"nullable"`
	// The priority for the request. This field may be ignored or overwritten depending
	// on the organization tier.
	//
	// Any of "low", "medium", "high", "critical".
	Priority         LlamaParseParametersPriority `json:"priority" api:"nullable"`
	ProjectID        string                       `json:"project_id" api:"nullable"`
	RemoveHiddenText bool                         `json:"remove_hidden_text" api:"nullable"`
	// Enum for representing the different available page error handling modes.
	//
	// Any of "raw_text", "blank_page", "error_message".
	ReplaceFailedPageMode                   FailPageMode `json:"replace_failed_page_mode" api:"nullable"`
	ReplaceFailedPageWithErrorMessagePrefix string       `json:"replace_failed_page_with_error_message_prefix" api:"nullable"`
	ReplaceFailedPageWithErrorMessageSuffix string       `json:"replace_failed_page_with_error_message_suffix" api:"nullable"`
	SaveImages                              bool         `json:"save_images" api:"nullable"`
	SkipDiagonalText                        bool         `json:"skip_diagonal_text" api:"nullable"`
	SpecializedChartParsingAgentic          bool         `json:"specialized_chart_parsing_agentic" api:"nullable"`
	SpecializedChartParsingEfficient        bool         `json:"specialized_chart_parsing_efficient" api:"nullable"`
	SpecializedChartParsingPlus             bool         `json:"specialized_chart_parsing_plus" api:"nullable"`
	SpecializedImageParsing                 bool         `json:"specialized_image_parsing" api:"nullable"`
	SpreadsheetExtractSubTables             bool         `json:"spreadsheet_extract_sub_tables" api:"nullable"`
	SpreadsheetForceFormulaComputation      bool         `json:"spreadsheet_force_formula_computation" api:"nullable"`
	SpreadsheetIncludeHiddenSheets          bool         `json:"spreadsheet_include_hidden_sheets" api:"nullable"`
	StrictModeBuggyFont                     bool         `json:"strict_mode_buggy_font" api:"nullable"`
	StrictModeImageExtraction               bool         `json:"strict_mode_image_extraction" api:"nullable"`
	StrictModeImageOcr                      bool         `json:"strict_mode_image_ocr" api:"nullable"`
	StrictModeReconstruction                bool         `json:"strict_mode_reconstruction" api:"nullable"`
	StructuredOutput                        bool         `json:"structured_output" api:"nullable"`
	StructuredOutputJsonSchema              string       `json:"structured_output_json_schema" api:"nullable"`
	StructuredOutputJsonSchemaName          string       `json:"structured_output_json_schema_name" api:"nullable"`
	SystemPrompt                            string       `json:"system_prompt" api:"nullable"`
	SystemPromptAppend                      string       `json:"system_prompt_append" api:"nullable"`
	TakeScreenshot                          bool         `json:"take_screenshot" api:"nullable"`
	TargetPages                             string       `json:"target_pages" api:"nullable"`
	Tier                                    string       `json:"tier" api:"nullable"`
	UseVendorMultimodalModel                bool         `json:"use_vendor_multimodal_model" api:"nullable"`
	UserPrompt                              string       `json:"user_prompt" api:"nullable"`
	VendorMultimodalAPIKey                  string       `json:"vendor_multimodal_api_key" api:"nullable"`
	VendorMultimodalModelName               string       `json:"vendor_multimodal_model_name" api:"nullable"`
	Version                                 string       `json:"version" api:"nullable"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []LlamaParseParametersWebhookConfigurationResp `json:"webhook_configurations" api:"nullable"`
	WebhookURL            string                                         `json:"webhook_url" api:"nullable"`
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
func (r LlamaParseParametersResp) RawJSON() string { return r.JSON.raw }
func (r *LlamaParseParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LlamaParseParametersResp to a LlamaParseParameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LlamaParseParameters.Overrides()
func (r LlamaParseParametersResp) ToParam() LlamaParseParameters {
	return param.Override[LlamaParseParameters](json.RawMessage(r.RawJSON()))
}

// The priority for the request. This field may be ignored or overwritten depending
// on the organization tier.
type LlamaParseParametersPriority string

const (
	LlamaParseParametersPriorityLow      LlamaParseParametersPriority = "low"
	LlamaParseParametersPriorityMedium   LlamaParseParametersPriority = "medium"
	LlamaParseParametersPriorityHigh     LlamaParseParametersPriority = "high"
	LlamaParseParametersPriorityCritical LlamaParseParametersPriority = "critical"
)

// Configuration for a single outbound webhook endpoint.
type LlamaParseParametersWebhookConfigurationResp struct {
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
func (r LlamaParseParametersWebhookConfigurationResp) RawJSON() string { return r.JSON.raw }
func (r *LlamaParseParametersWebhookConfigurationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlamaParseParameters struct {
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
	InputS3Region                            param.Opt[string]  `json:"input_s3_region,omitzero"`
	InputURL                                 param.Opt[string]  `json:"input_url,omitzero"`
	InternalIsScreenshotJob                  param.Opt[bool]    `json:"internal_is_screenshot_job,omitzero"`
	InvalidateCache                          param.Opt[bool]    `json:"invalidate_cache,omitzero"`
	IsFormattingInstruction                  param.Opt[bool]    `json:"is_formatting_instruction,omitzero"`
	JobTimeoutExtraTimePerPageInSeconds      param.Opt[float64] `json:"job_timeout_extra_time_per_page_in_seconds,omitzero"`
	JobTimeoutInSeconds                      param.Opt[float64] `json:"job_timeout_in_seconds,omitzero"`
	KeepPageSeparatorWhenMergingTables       param.Opt[bool]    `json:"keep_page_separator_when_merging_tables,omitzero"`
	LayoutAware                              param.Opt[bool]    `json:"layout_aware,omitzero"`
	LineLevelBoundingBox                     param.Opt[bool]    `json:"line_level_bounding_box,omitzero"`
	MarkdownTableMultilineHeaderSeparator    param.Opt[string]  `json:"markdown_table_multiline_header_separator,omitzero"`
	MaxPages                                 param.Opt[int64]   `json:"max_pages,omitzero"`
	MaxPagesEnforced                         param.Opt[int64]   `json:"max_pages_enforced,omitzero"`
	MergeTablesAcrossPagesInMarkdown         param.Opt[bool]    `json:"merge_tables_across_pages_in_markdown,omitzero"`
	Model                                    param.Opt[string]  `json:"model,omitzero"`
	OutlinedTableExtraction                  param.Opt[bool]    `json:"outlined_table_extraction,omitzero"`
	OutputPdfOfDocument                      param.Opt[bool]    `json:"output_pdf_of_document,omitzero"`
	OutputS3PathPrefix                       param.Opt[string]  `json:"output_s3_path_prefix,omitzero"`
	OutputS3Region                           param.Opt[string]  `json:"output_s3_region,omitzero"`
	OutputTablesAsHTML                       param.Opt[bool]    `json:"output_tables_as_HTML,omitzero"`
	PageErrorTolerance                       param.Opt[float64] `json:"page_error_tolerance,omitzero"`
	PageFooterPrefix                         param.Opt[string]  `json:"page_footer_prefix,omitzero"`
	PageFooterSuffix                         param.Opt[string]  `json:"page_footer_suffix,omitzero"`
	PageHeaderPrefix                         param.Opt[string]  `json:"page_header_prefix,omitzero"`
	PageHeaderSuffix                         param.Opt[string]  `json:"page_header_suffix,omitzero"`
	PagePrefix                               param.Opt[string]  `json:"page_prefix,omitzero"`
	PageSeparator                            param.Opt[string]  `json:"page_separator,omitzero"`
	PageSuffix                               param.Opt[string]  `json:"page_suffix,omitzero"`
	ParsingInstruction                       param.Opt[string]  `json:"parsing_instruction,omitzero"`
	PreciseBoundingBox                       param.Opt[bool]    `json:"precise_bounding_box,omitzero"`
	PremiumMode                              param.Opt[bool]    `json:"premium_mode,omitzero"`
	PresentationOutOfBoundsContent           param.Opt[bool]    `json:"presentation_out_of_bounds_content,omitzero"`
	PresentationSkipEmbeddedData             param.Opt[bool]    `json:"presentation_skip_embedded_data,omitzero"`
	PreserveLayoutAlignmentAcrossPages       param.Opt[bool]    `json:"preserve_layout_alignment_across_pages,omitzero"`
	PreserveVerySmallText                    param.Opt[bool]    `json:"preserve_very_small_text,omitzero"`
	Preset                                   param.Opt[string]  `json:"preset,omitzero"`
	ProjectID                                param.Opt[string]  `json:"project_id,omitzero"`
	RemoveHiddenText                         param.Opt[bool]    `json:"remove_hidden_text,omitzero"`
	ReplaceFailedPageWithErrorMessagePrefix  param.Opt[string]  `json:"replace_failed_page_with_error_message_prefix,omitzero"`
	ReplaceFailedPageWithErrorMessageSuffix  param.Opt[string]  `json:"replace_failed_page_with_error_message_suffix,omitzero"`
	SaveImages                               param.Opt[bool]    `json:"save_images,omitzero"`
	SkipDiagonalText                         param.Opt[bool]    `json:"skip_diagonal_text,omitzero"`
	SpecializedChartParsingAgentic           param.Opt[bool]    `json:"specialized_chart_parsing_agentic,omitzero"`
	SpecializedChartParsingEfficient         param.Opt[bool]    `json:"specialized_chart_parsing_efficient,omitzero"`
	SpecializedChartParsingPlus              param.Opt[bool]    `json:"specialized_chart_parsing_plus,omitzero"`
	SpecializedImageParsing                  param.Opt[bool]    `json:"specialized_image_parsing,omitzero"`
	SpreadsheetExtractSubTables              param.Opt[bool]    `json:"spreadsheet_extract_sub_tables,omitzero"`
	SpreadsheetForceFormulaComputation       param.Opt[bool]    `json:"spreadsheet_force_formula_computation,omitzero"`
	SpreadsheetIncludeHiddenSheets           param.Opt[bool]    `json:"spreadsheet_include_hidden_sheets,omitzero"`
	StrictModeBuggyFont                      param.Opt[bool]    `json:"strict_mode_buggy_font,omitzero"`
	StrictModeImageExtraction                param.Opt[bool]    `json:"strict_mode_image_extraction,omitzero"`
	StrictModeImageOcr                       param.Opt[bool]    `json:"strict_mode_image_ocr,omitzero"`
	StrictModeReconstruction                 param.Opt[bool]    `json:"strict_mode_reconstruction,omitzero"`
	StructuredOutput                         param.Opt[bool]    `json:"structured_output,omitzero"`
	StructuredOutputJsonSchema               param.Opt[string]  `json:"structured_output_json_schema,omitzero"`
	StructuredOutputJsonSchemaName           param.Opt[string]  `json:"structured_output_json_schema_name,omitzero"`
	SystemPrompt                             param.Opt[string]  `json:"system_prompt,omitzero"`
	SystemPromptAppend                       param.Opt[string]  `json:"system_prompt_append,omitzero"`
	TakeScreenshot                           param.Opt[bool]    `json:"take_screenshot,omitzero"`
	TargetPages                              param.Opt[string]  `json:"target_pages,omitzero"`
	Tier                                     param.Opt[string]  `json:"tier,omitzero"`
	UseVendorMultimodalModel                 param.Opt[bool]    `json:"use_vendor_multimodal_model,omitzero"`
	UserPrompt                               param.Opt[string]  `json:"user_prompt,omitzero"`
	VendorMultimodalAPIKey                   param.Opt[string]  `json:"vendor_multimodal_api_key,omitzero"`
	VendorMultimodalModelName                param.Opt[string]  `json:"vendor_multimodal_model_name,omitzero"`
	Version                                  param.Opt[string]  `json:"version,omitzero"`
	WebhookURL                               param.Opt[string]  `json:"webhook_url,omitzero"`
	// Any of "screenshot", "embedded", "layout".
	ImagesToSave []string `json:"images_to_save,omitzero"`
	// The priority for the request. This field may be ignored or overwritten depending
	// on the organization tier.
	//
	// Any of "low", "medium", "high", "critical".
	Priority LlamaParseParametersPriority `json:"priority,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []LlamaParseParametersWebhookConfiguration `json:"webhook_configurations,omitzero"`
	Languages             []ParsingLanguages                         `json:"languages,omitzero"`
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
	paramObj
}

func (r LlamaParseParameters) MarshalJSON() (data []byte, err error) {
	type shadow LlamaParseParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlamaParseParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a single outbound webhook endpoint.
type LlamaParseParametersWebhookConfiguration struct {
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

func (r LlamaParseParametersWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow LlamaParseParametersWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlamaParseParametersWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedIngestionStatusResponse struct {
	// Status of the ingestion.
	//
	// Any of "NOT_STARTED", "IN_PROGRESS", "SUCCESS", "ERROR", "PARTIAL_SUCCESS",
	// "CANCELLED".
	Status ManagedIngestionStatusResponseStatus `json:"status" api:"required"`
	// Date of the deployment.
	DeploymentDate time.Time `json:"deployment_date" api:"nullable" format:"date-time"`
	// When the status is effective
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// List of errors that occurred during ingestion.
	Error []ManagedIngestionStatusResponseError `json:"error" api:"nullable"`
	// ID of the latest job.
	JobID string `json:"job_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status         respjson.Field
		DeploymentDate respjson.Field
		EffectiveAt    respjson.Field
		Error          respjson.Field
		JobID          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedIngestionStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedIngestionStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the ingestion.
type ManagedIngestionStatusResponseStatus string

const (
	ManagedIngestionStatusResponseStatusNotStarted     ManagedIngestionStatusResponseStatus = "NOT_STARTED"
	ManagedIngestionStatusResponseStatusInProgress     ManagedIngestionStatusResponseStatus = "IN_PROGRESS"
	ManagedIngestionStatusResponseStatusSuccess        ManagedIngestionStatusResponseStatus = "SUCCESS"
	ManagedIngestionStatusResponseStatusError          ManagedIngestionStatusResponseStatus = "ERROR"
	ManagedIngestionStatusResponseStatusPartialSuccess ManagedIngestionStatusResponseStatus = "PARTIAL_SUCCESS"
	ManagedIngestionStatusResponseStatusCancelled      ManagedIngestionStatusResponseStatus = "CANCELLED"
)

type ManagedIngestionStatusResponseError struct {
	// ID of the job that failed.
	JobID string `json:"job_id" api:"required" format:"uuid"`
	// List of errors that occurred during ingestion.
	Message string `json:"message" api:"required"`
	// Name of the job that failed.
	//
	// Any of "MANAGED_INGESTION", "DATA_SOURCE", "FILE_UPDATER", "PARSE", "TRANSFORM",
	// "INGESTION", "METADATA_UPDATE".
	Step string `json:"step" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Message     respjson.Field
		Step        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedIngestionStatusResponseError) RawJSON() string { return r.JSON.raw }
func (r *ManagedIngestionStatusResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata filters for vector stores.
type MetadataFilters struct {
	Filters []MetadataFiltersFilterUnion `json:"filters" api:"required"`
	// Vector store filter conditions to combine different filters.
	//
	// Any of "and", "or", "not".
	Condition MetadataFiltersCondition `json:"condition" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Condition   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataFilters) RawJSON() string { return r.JSON.raw }
func (r *MetadataFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MetadataFilters to a MetadataFiltersParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MetadataFiltersParam.Overrides()
func (r MetadataFilters) ToParam() MetadataFiltersParam {
	return param.Override[MetadataFiltersParam](json.RawMessage(r.RawJSON()))
}

// MetadataFiltersFilterUnion contains all possible properties and values from
// [MetadataFiltersFilterMetadataFilter], [MetadataFilters].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MetadataFiltersFilterUnion struct {
	// This field is from variant [MetadataFiltersFilterMetadataFilter].
	Key string `json:"key"`
	// This field is from variant [MetadataFiltersFilterMetadataFilter].
	Value MetadataFiltersFilterMetadataFilterValueUnion `json:"value"`
	// This field is from variant [MetadataFiltersFilterMetadataFilter].
	Operator string `json:"operator"`
	// This field is from variant [MetadataFilters].
	Filters []MetadataFiltersFilterUnion `json:"filters"`
	// This field is from variant [MetadataFilters].
	Condition MetadataFiltersCondition `json:"condition"`
	JSON      struct {
		Key       respjson.Field
		Value     respjson.Field
		Operator  respjson.Field
		Filters   respjson.Field
		Condition respjson.Field
		raw       string
	} `json:"-"`
}

func (u MetadataFiltersFilterUnion) AsMetadataFilter() (v MetadataFiltersFilterMetadataFilter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MetadataFiltersFilterUnion) AsMetadataFilters() (v MetadataFilters) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MetadataFiltersFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *MetadataFiltersFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive metadata filter for vector stores to support more operators.
//
// Value uses Strict types, as int, float and str are compatible types and were all
// converted to string before.
//
// See: https://docs.pydantic.dev/latest/usage/types/#strict-types
type MetadataFiltersFilterMetadataFilter struct {
	Key   string                                        `json:"key" api:"required"`
	Value MetadataFiltersFilterMetadataFilterValueUnion `json:"value" api:"required"`
	// Vector store filter operator.
	//
	// Any of "==", ">", "<", "!=", ">=", "<=", "in", "nin", "any", "all",
	// "text_match", "text_match_insensitive", "contains", "is_empty".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataFiltersFilterMetadataFilter) RawJSON() string { return r.JSON.raw }
func (r *MetadataFiltersFilterMetadataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MetadataFiltersFilterMetadataFilterValueUnion contains all possible properties
// and values from [float64], [string], [[]string], [[]float64], [[]int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString OfStringArray OfNumberArray OfIntegerArray]
type MetadataFiltersFilterMetadataFilterValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a [[]float64] instead of an object.
	OfNumberArray []float64 `json:",inline"`
	// This field will be present if the value is a [[]int64] instead of an object.
	OfIntegerArray []int64 `json:",inline"`
	JSON           struct {
		OfFloat        respjson.Field
		OfString       respjson.Field
		OfStringArray  respjson.Field
		OfNumberArray  respjson.Field
		OfIntegerArray respjson.Field
		raw            string
	} `json:"-"`
}

func (u MetadataFiltersFilterMetadataFilterValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MetadataFiltersFilterMetadataFilterValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MetadataFiltersFilterMetadataFilterValueUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MetadataFiltersFilterMetadataFilterValueUnion) AsNumberArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MetadataFiltersFilterMetadataFilterValueUnion) AsIntegerArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MetadataFiltersFilterMetadataFilterValueUnion) RawJSON() string { return u.JSON.raw }

func (r *MetadataFiltersFilterMetadataFilterValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector store filter conditions to combine different filters.
type MetadataFiltersCondition string

const (
	MetadataFiltersConditionAnd MetadataFiltersCondition = "and"
	MetadataFiltersConditionOr  MetadataFiltersCondition = "or"
	MetadataFiltersConditionNot MetadataFiltersCondition = "not"
)

// Metadata filters for vector stores.
//
// The property Filters is required.
type MetadataFiltersParam struct {
	Filters []MetadataFiltersFilterUnionParam `json:"filters,omitzero" api:"required"`
	// Vector store filter conditions to combine different filters.
	//
	// Any of "and", "or", "not".
	Condition MetadataFiltersCondition `json:"condition,omitzero"`
	paramObj
}

func (r MetadataFiltersParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataFiltersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataFiltersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MetadataFiltersFilterUnionParam struct {
	OfMetadataFilter  *MetadataFiltersFilterMetadataFilterParam `json:",omitzero,inline"`
	OfMetadataFilters *MetadataFiltersParam                     `json:",omitzero,inline"`
	paramUnion
}

func (u MetadataFiltersFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMetadataFilter, u.OfMetadataFilters)
}
func (u *MetadataFiltersFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Comprehensive metadata filter for vector stores to support more operators.
//
// Value uses Strict types, as int, float and str are compatible types and were all
// converted to string before.
//
// See: https://docs.pydantic.dev/latest/usage/types/#strict-types
//
// The properties Key, Value are required.
type MetadataFiltersFilterMetadataFilterParam struct {
	Value MetadataFiltersFilterMetadataFilterValueUnionParam `json:"value,omitzero" api:"required"`
	Key   string                                             `json:"key" api:"required"`
	// Vector store filter operator.
	//
	// Any of "==", ">", "<", "!=", ">=", "<=", "in", "nin", "any", "all",
	// "text_match", "text_match_insensitive", "contains", "is_empty".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r MetadataFiltersFilterMetadataFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataFiltersFilterMetadataFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataFiltersFilterMetadataFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MetadataFiltersFilterMetadataFilterParam](
		"operator", "==", ">", "<", "!=", ">=", "<=", "in", "nin", "any", "all", "text_match", "text_match_insensitive", "contains", "is_empty",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MetadataFiltersFilterMetadataFilterValueUnionParam struct {
	OfFloat        param.Opt[float64] `json:",omitzero,inline"`
	OfString       param.Opt[string]  `json:",omitzero,inline"`
	OfStringArray  []string           `json:",omitzero,inline"`
	OfNumberArray  []float64          `json:",omitzero,inline"`
	OfIntegerArray []int64            `json:",omitzero,inline"`
	paramUnion
}

func (u MetadataFiltersFilterMetadataFilterValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat,
		u.OfString,
		u.OfStringArray,
		u.OfNumberArray,
		u.OfIntegerArray)
}
func (u *MetadataFiltersFilterMetadataFilterValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type OpenAIEmbedding struct {
	// Additional kwargs for the OpenAI API.
	AdditionalKwargs map[string]any `json:"additional_kwargs"`
	// The base URL for OpenAI API.
	APIBase string `json:"api_base" api:"nullable"`
	// The OpenAI API key.
	APIKey string `json:"api_key" api:"nullable"`
	// The version for OpenAI API.
	APIVersion string `json:"api_version" api:"nullable"`
	ClassName  string `json:"class_name"`
	// The default headers for API requests.
	DefaultHeaders map[string]string `json:"default_headers" api:"nullable"`
	// The number of dimensions on the output embedding vectors. Works only with v3
	// embedding models.
	Dimensions int64 `json:"dimensions" api:"nullable"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// Maximum number of retries.
	MaxRetries int64 `json:"max_retries"`
	// The name of the OpenAI embedding model.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// Reuse the OpenAI client between requests. When doing anything with large volumes
	// of async API calls, setting this to false can improve stability.
	ReuseClient bool `json:"reuse_client"`
	// Timeout for each request.
	Timeout float64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalKwargs respjson.Field
		APIBase          respjson.Field
		APIKey           respjson.Field
		APIVersion       respjson.Field
		ClassName        respjson.Field
		DefaultHeaders   respjson.Field
		Dimensions       respjson.Field
		EmbedBatchSize   respjson.Field
		MaxRetries       respjson.Field
		ModelName        respjson.Field
		NumWorkers       respjson.Field
		ReuseClient      respjson.Field
		Timeout          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIEmbedding) RawJSON() string { return r.JSON.raw }
func (r *OpenAIEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIEmbedding to a OpenAIEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIEmbeddingParam.Overrides()
func (r OpenAIEmbedding) ToParam() OpenAIEmbeddingParam {
	return param.Override[OpenAIEmbeddingParam](json.RawMessage(r.RawJSON()))
}

type OpenAIEmbeddingParam struct {
	// The base URL for OpenAI API.
	APIBase param.Opt[string] `json:"api_base,omitzero"`
	// The OpenAI API key.
	APIKey param.Opt[string] `json:"api_key,omitzero"`
	// The version for OpenAI API.
	APIVersion param.Opt[string] `json:"api_version,omitzero"`
	// The number of dimensions on the output embedding vectors. Works only with v3
	// embedding models.
	Dimensions param.Opt[int64] `json:"dimensions,omitzero"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64]  `json:"num_workers,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// Maximum number of retries.
	MaxRetries param.Opt[int64] `json:"max_retries,omitzero"`
	// The name of the OpenAI embedding model.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// Reuse the OpenAI client between requests. When doing anything with large volumes
	// of async API calls, setting this to false can improve stability.
	ReuseClient param.Opt[bool] `json:"reuse_client,omitzero"`
	// Timeout for each request.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// The default headers for API requests.
	DefaultHeaders map[string]string `json:"default_headers,omitzero"`
	// Additional kwargs for the OpenAI API.
	AdditionalKwargs map[string]any `json:"additional_kwargs,omitzero"`
	paramObj
}

func (r OpenAIEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIEmbeddingConfig struct {
	// Configuration for the OpenAI embedding model.
	Component OpenAIEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "OPENAI_EMBEDDING".
	Type OpenAIEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *OpenAIEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OpenAIEmbeddingConfig to a OpenAIEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OpenAIEmbeddingConfigParam.Overrides()
func (r OpenAIEmbeddingConfig) ToParam() OpenAIEmbeddingConfigParam {
	return param.Override[OpenAIEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type OpenAIEmbeddingConfigType string

const (
	OpenAIEmbeddingConfigTypeOpenAIEmbedding OpenAIEmbeddingConfigType = "OPENAI_EMBEDDING"
)

type OpenAIEmbeddingConfigParam struct {
	// Configuration for the OpenAI embedding model.
	Component OpenAIEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "OPENAI_EMBEDDING".
	Type OpenAIEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r OpenAIEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow OpenAIEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OpenAIEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page figure metadata with score
type PageFigureNodeWithScore struct {
	Node PageFigureNodeWithScoreNode `json:"node" api:"required"`
	// The score of the figure node
	Score     float64 `json:"score" api:"required"`
	ClassName string  `json:"class_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Node        respjson.Field
		Score       respjson.Field
		ClassName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageFigureNodeWithScore) RawJSON() string { return r.JSON.raw }
func (r *PageFigureNodeWithScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageFigureNodeWithScoreNode struct {
	// The confidence of the figure
	Confidence float64 `json:"confidence" api:"required"`
	// The name of the figure
	FigureName string `json:"figure_name" api:"required"`
	// The size of the figure in bytes
	FigureSize int64 `json:"figure_size" api:"required"`
	// The ID of the file that the figure was taken from
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// The index of the page for which the figure is taken (0-indexed)
	PageIndex int64 `json:"page_index" api:"required"`
	// Whether the figure is likely to be noise
	IsLikelyNoise bool `json:"is_likely_noise"`
	// Metadata for the figure
	Metadata map[string]any `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence    respjson.Field
		FigureName    respjson.Field
		FigureSize    respjson.Field
		FileID        respjson.Field
		PageIndex     respjson.Field
		IsLikelyNoise respjson.Field
		Metadata      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageFigureNodeWithScoreNode) RawJSON() string { return r.JSON.raw }
func (r *PageFigureNodeWithScoreNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page screenshot metadata with score
type PageScreenshotNodeWithScore struct {
	Node PageScreenshotNodeWithScoreNode `json:"node" api:"required"`
	// The score of the screenshot node
	Score     float64 `json:"score" api:"required"`
	ClassName string  `json:"class_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Node        respjson.Field
		Score       respjson.Field
		ClassName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageScreenshotNodeWithScore) RawJSON() string { return r.JSON.raw }
func (r *PageScreenshotNodeWithScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageScreenshotNodeWithScoreNode struct {
	// The ID of the file that the page screenshot was taken from
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// The size of the image in bytes
	ImageSize int64 `json:"image_size" api:"required"`
	// The index of the page for which the screenshot is taken (0-indexed)
	PageIndex int64 `json:"page_index" api:"required"`
	// Metadata for the screenshot
	Metadata map[string]any `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		ImageSize   respjson.Field
		PageIndex   respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageScreenshotNodeWithScoreNode) RawJSON() string { return r.JSON.raw }
func (r *PageScreenshotNodeWithScoreNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for a pipeline.
type Pipeline struct {
	// Unique identifier
	ID              string                       `json:"id" api:"required" format:"uuid"`
	EmbeddingConfig PipelineEmbeddingConfigUnion `json:"embedding_config" api:"required"`
	Name            string                       `json:"name" api:"required"`
	ProjectID       string                       `json:"project_id" api:"required" format:"uuid"`
	// Hashes for the configuration of a pipeline.
	ConfigHash PipelineConfigHash `json:"config_hash" api:"nullable"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Schema for a data sink.
	DataSink DataSink `json:"data_sink" api:"nullable"`
	// Schema for an embedding model config.
	EmbeddingModelConfig PipelineEmbeddingModelConfig `json:"embedding_model_config" api:"nullable"`
	// The ID of the EmbeddingModelConfig this pipeline is using.
	EmbeddingModelConfigID string `json:"embedding_model_config_id" api:"nullable" format:"uuid"`
	// Settings that can be configured for how to use LlamaParse to parse files within
	// a LlamaCloud pipeline.
	LlamaParseParameters LlamaParseParametersResp `json:"llama_parse_parameters" api:"nullable"`
	// The ID of the ManagedPipeline this playground pipeline is linked to.
	ManagedPipelineID string `json:"managed_pipeline_id" api:"nullable" format:"uuid"`
	// Metadata configuration for the pipeline.
	MetadataConfig PipelineMetadataConfig `json:"metadata_config" api:"nullable"`
	// Type of pipeline. Either PLAYGROUND or MANAGED.
	//
	// Any of "PLAYGROUND", "MANAGED".
	PipelineType PipelineType `json:"pipeline_type"`
	// Preset retrieval parameters for the pipeline.
	PresetRetrievalParameters PresetRetrievalParamsResp `json:"preset_retrieval_parameters"`
	// Configuration for sparse embedding models used in hybrid search.
	//
	// This allows users to choose between Splade and BM25 models for sparse retrieval
	// in managed data sinks.
	SparseModelConfig SparseModelConfig `json:"sparse_model_config" api:"nullable"`
	// Status of the pipeline.
	//
	// Any of "CREATED", "DELETING".
	Status PipelineStatus `json:"status" api:"nullable"`
	// Configuration for the transformation.
	TransformConfig PipelineTransformConfigUnion `json:"transform_config"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		EmbeddingConfig           respjson.Field
		Name                      respjson.Field
		ProjectID                 respjson.Field
		ConfigHash                respjson.Field
		CreatedAt                 respjson.Field
		DataSink                  respjson.Field
		EmbeddingModelConfig      respjson.Field
		EmbeddingModelConfigID    respjson.Field
		LlamaParseParameters      respjson.Field
		ManagedPipelineID         respjson.Field
		MetadataConfig            respjson.Field
		PipelineType              respjson.Field
		PresetRetrievalParameters respjson.Field
		SparseModelConfig         respjson.Field
		Status                    respjson.Field
		TransformConfig           respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pipeline) RawJSON() string { return r.JSON.raw }
func (r *Pipeline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineEmbeddingConfigUnion contains all possible properties and values from
// [PipelineEmbeddingConfigManagedOpenAIEmbedding], [AzureOpenAIEmbeddingConfig],
// [CohereEmbeddingConfig], [GeminiEmbeddingConfig],
// [HuggingFaceInferenceAPIEmbeddingConfig], [OpenAIEmbeddingConfig],
// [VertexAIEmbeddingConfig], [BedrockEmbeddingConfig].
//
// Use the [PipelineEmbeddingConfigUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PipelineEmbeddingConfigUnion struct {
	// This field is a union of
	// [PipelineEmbeddingConfigManagedOpenAIEmbeddingComponent],
	// [AzureOpenAIEmbedding], [CohereEmbedding], [GeminiEmbedding],
	// [HuggingFaceInferenceAPIEmbedding], [OpenAIEmbedding], [VertexTextEmbedding],
	// [BedrockEmbedding]
	Component PipelineEmbeddingConfigUnionComponent `json:"component"`
	// Any of "MANAGED_OPENAI_EMBEDDING", "AZURE_EMBEDDING", "COHERE_EMBEDDING",
	// "GEMINI_EMBEDDING", "HUGGINGFACE_API_EMBEDDING", "OPENAI_EMBEDDING",
	// "VERTEXAI_EMBEDDING", "BEDROCK_EMBEDDING".
	Type string `json:"type"`
	JSON struct {
		Component respjson.Field
		Type      respjson.Field
		raw       string
	} `json:"-"`
}

// anyPipelineEmbeddingConfig is implemented by each variant of
// [PipelineEmbeddingConfigUnion] to add type safety for the return type of
// [PipelineEmbeddingConfigUnion.AsAny]
type anyPipelineEmbeddingConfig interface {
	implPipelineEmbeddingConfigUnion()
}

func (PipelineEmbeddingConfigManagedOpenAIEmbedding) implPipelineEmbeddingConfigUnion() {}
func (AzureOpenAIEmbeddingConfig) implPipelineEmbeddingConfigUnion()                    {}
func (CohereEmbeddingConfig) implPipelineEmbeddingConfigUnion()                         {}
func (GeminiEmbeddingConfig) implPipelineEmbeddingConfigUnion()                         {}
func (HuggingFaceInferenceAPIEmbeddingConfig) implPipelineEmbeddingConfigUnion()        {}
func (OpenAIEmbeddingConfig) implPipelineEmbeddingConfigUnion()                         {}
func (VertexAIEmbeddingConfig) implPipelineEmbeddingConfigUnion()                       {}
func (BedrockEmbeddingConfig) implPipelineEmbeddingConfigUnion()                        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PipelineEmbeddingConfigUnion.AsAny().(type) {
//	case llamacloudprod.PipelineEmbeddingConfigManagedOpenAIEmbedding:
//	case llamacloudprod.AzureOpenAIEmbeddingConfig:
//	case llamacloudprod.CohereEmbeddingConfig:
//	case llamacloudprod.GeminiEmbeddingConfig:
//	case llamacloudprod.HuggingFaceInferenceAPIEmbeddingConfig:
//	case llamacloudprod.OpenAIEmbeddingConfig:
//	case llamacloudprod.VertexAIEmbeddingConfig:
//	case llamacloudprod.BedrockEmbeddingConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PipelineEmbeddingConfigUnion) AsAny() anyPipelineEmbeddingConfig {
	switch u.Type {
	case "MANAGED_OPENAI_EMBEDDING":
		return u.AsManagedOpenAIEmbedding()
	case "AZURE_EMBEDDING":
		return u.AsAzureEmbedding()
	case "COHERE_EMBEDDING":
		return u.AsCohereEmbedding()
	case "GEMINI_EMBEDDING":
		return u.AsGeminiEmbedding()
	case "HUGGINGFACE_API_EMBEDDING":
		return u.AsHuggingfaceAPIEmbedding()
	case "OPENAI_EMBEDDING":
		return u.AsOpenAIEmbedding()
	case "VERTEXAI_EMBEDDING":
		return u.AsVertexaiEmbedding()
	case "BEDROCK_EMBEDDING":
		return u.AsBedrockEmbedding()
	}
	return nil
}

func (u PipelineEmbeddingConfigUnion) AsManagedOpenAIEmbedding() (v PipelineEmbeddingConfigManagedOpenAIEmbedding) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsAzureEmbedding() (v AzureOpenAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsCohereEmbedding() (v CohereEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsGeminiEmbedding() (v GeminiEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsHuggingfaceAPIEmbedding() (v HuggingFaceInferenceAPIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsOpenAIEmbedding() (v OpenAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsVertexaiEmbedding() (v VertexAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingConfigUnion) AsBedrockEmbedding() (v BedrockEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineEmbeddingConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineEmbeddingConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineEmbeddingConfigUnionComponent is an implicit subunion of
// [PipelineEmbeddingConfigUnion]. PipelineEmbeddingConfigUnionComponent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PipelineEmbeddingConfigUnion].
type PipelineEmbeddingConfigUnionComponent struct {
	ClassName        string `json:"class_name"`
	EmbedBatchSize   int64  `json:"embed_batch_size"`
	ModelName        string `json:"model_name"`
	NumWorkers       int64  `json:"num_workers"`
	AdditionalKwargs any    `json:"additional_kwargs"`
	APIBase          string `json:"api_base"`
	APIKey           string `json:"api_key"`
	APIVersion       string `json:"api_version"`
	// This field is from variant [AzureOpenAIEmbedding].
	AzureDeployment string `json:"azure_deployment"`
	// This field is from variant [AzureOpenAIEmbedding].
	AzureEndpoint  string  `json:"azure_endpoint"`
	DefaultHeaders string  `json:"default_headers"`
	Dimensions     int64   `json:"dimensions"`
	MaxRetries     int64   `json:"max_retries"`
	ReuseClient    bool    `json:"reuse_client"`
	Timeout        float64 `json:"timeout"`
	// This field is from variant [CohereEmbedding].
	EmbeddingType string `json:"embedding_type"`
	// This field is from variant [CohereEmbedding].
	InputType string `json:"input_type"`
	// This field is from variant [CohereEmbedding].
	Truncate string `json:"truncate"`
	// This field is from variant [GeminiEmbedding].
	OutputDimensionality int64 `json:"output_dimensionality"`
	// This field is from variant [GeminiEmbedding].
	TaskType string `json:"task_type"`
	// This field is from variant [GeminiEmbedding].
	Title string `json:"title"`
	// This field is from variant [GeminiEmbedding].
	Transport string `json:"transport"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Token HuggingFaceInferenceAPIEmbeddingTokenUnion `json:"token"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Cookies map[string]string `json:"cookies"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Headers map[string]string `json:"headers"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Pooling HuggingFaceInferenceAPIEmbeddingPooling `json:"pooling"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	QueryInstruction string `json:"query_instruction"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Task string `json:"task"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	TextInstruction string `json:"text_instruction"`
	// This field is from variant [VertexTextEmbedding].
	ClientEmail string `json:"client_email"`
	// This field is from variant [VertexTextEmbedding].
	Location string `json:"location"`
	// This field is from variant [VertexTextEmbedding].
	PrivateKey string `json:"private_key"`
	// This field is from variant [VertexTextEmbedding].
	PrivateKeyID string `json:"private_key_id"`
	// This field is from variant [VertexTextEmbedding].
	Project string `json:"project"`
	// This field is from variant [VertexTextEmbedding].
	TokenUri string `json:"token_uri"`
	// This field is from variant [VertexTextEmbedding].
	EmbedMode VertexTextEmbeddingEmbedMode `json:"embed_mode"`
	// This field is from variant [BedrockEmbedding].
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// This field is from variant [BedrockEmbedding].
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// This field is from variant [BedrockEmbedding].
	AwsSessionToken string `json:"aws_session_token"`
	// This field is from variant [BedrockEmbedding].
	ProfileName string `json:"profile_name"`
	// This field is from variant [BedrockEmbedding].
	RegionName string `json:"region_name"`
	JSON       struct {
		ClassName            respjson.Field
		EmbedBatchSize       respjson.Field
		ModelName            respjson.Field
		NumWorkers           respjson.Field
		AdditionalKwargs     respjson.Field
		APIBase              respjson.Field
		APIKey               respjson.Field
		APIVersion           respjson.Field
		AzureDeployment      respjson.Field
		AzureEndpoint        respjson.Field
		DefaultHeaders       respjson.Field
		Dimensions           respjson.Field
		MaxRetries           respjson.Field
		ReuseClient          respjson.Field
		Timeout              respjson.Field
		EmbeddingType        respjson.Field
		InputType            respjson.Field
		Truncate             respjson.Field
		OutputDimensionality respjson.Field
		TaskType             respjson.Field
		Title                respjson.Field
		Transport            respjson.Field
		Token                respjson.Field
		Cookies              respjson.Field
		Headers              respjson.Field
		Pooling              respjson.Field
		QueryInstruction     respjson.Field
		Task                 respjson.Field
		TextInstruction      respjson.Field
		ClientEmail          respjson.Field
		Location             respjson.Field
		PrivateKey           respjson.Field
		PrivateKeyID         respjson.Field
		Project              respjson.Field
		TokenUri             respjson.Field
		EmbedMode            respjson.Field
		AwsAccessKeyID       respjson.Field
		AwsSecretAccessKey   respjson.Field
		AwsSessionToken      respjson.Field
		ProfileName          respjson.Field
		RegionName           respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *PipelineEmbeddingConfigUnionComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineEmbeddingConfigManagedOpenAIEmbedding struct {
	// Configuration for the Managed OpenAI embedding model.
	Component PipelineEmbeddingConfigManagedOpenAIEmbeddingComponent `json:"component"`
	// Type of the embedding model.
	//
	// Any of "MANAGED_OPENAI_EMBEDDING".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineEmbeddingConfigManagedOpenAIEmbedding) RawJSON() string { return r.JSON.raw }
func (r *PipelineEmbeddingConfigManagedOpenAIEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the Managed OpenAI embedding model.
type PipelineEmbeddingConfigManagedOpenAIEmbeddingComponent struct {
	ClassName string `json:"class_name"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// The name of the OpenAI embedding model.
	//
	// Any of "openai-text-embedding-3-small".
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName      respjson.Field
		EmbedBatchSize respjson.Field
		ModelName      respjson.Field
		NumWorkers     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineEmbeddingConfigManagedOpenAIEmbeddingComponent) RawJSON() string { return r.JSON.raw }
func (r *PipelineEmbeddingConfigManagedOpenAIEmbeddingComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hashes for the configuration of a pipeline.
type PipelineConfigHash struct {
	// Hash of the embedding config.
	EmbeddingConfigHash string `json:"embedding_config_hash" api:"nullable"`
	// Hash of the llama parse parameters.
	ParsingConfigHash string `json:"parsing_config_hash" api:"nullable"`
	// Hash of the transform config.
	TransformConfigHash string `json:"transform_config_hash" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbeddingConfigHash respjson.Field
		ParsingConfigHash   respjson.Field
		TransformConfigHash respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineConfigHash) RawJSON() string { return r.JSON.raw }
func (r *PipelineConfigHash) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for an embedding model config.
type PipelineEmbeddingModelConfig struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The embedding configuration for the embedding model config.
	EmbeddingConfig PipelineEmbeddingModelConfigEmbeddingConfigUnion `json:"embedding_config" api:"required"`
	// The name of the embedding model config.
	Name      string `json:"name" api:"required"`
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		EmbeddingConfig respjson.Field
		Name            respjson.Field
		ProjectID       respjson.Field
		CreatedAt       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineEmbeddingModelConfig) RawJSON() string { return r.JSON.raw }
func (r *PipelineEmbeddingModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineEmbeddingModelConfigEmbeddingConfigUnion contains all possible
// properties and values from [AzureOpenAIEmbeddingConfig],
// [CohereEmbeddingConfig], [GeminiEmbeddingConfig],
// [HuggingFaceInferenceAPIEmbeddingConfig], [OpenAIEmbeddingConfig],
// [VertexAIEmbeddingConfig], [BedrockEmbeddingConfig].
//
// Use the [PipelineEmbeddingModelConfigEmbeddingConfigUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PipelineEmbeddingModelConfigEmbeddingConfigUnion struct {
	// This field is a union of [AzureOpenAIEmbedding], [CohereEmbedding],
	// [GeminiEmbedding], [HuggingFaceInferenceAPIEmbedding], [OpenAIEmbedding],
	// [VertexTextEmbedding], [BedrockEmbedding]
	Component PipelineEmbeddingModelConfigEmbeddingConfigUnionComponent `json:"component"`
	// Any of "AZURE_EMBEDDING", "COHERE_EMBEDDING", "GEMINI_EMBEDDING",
	// "HUGGINGFACE_API_EMBEDDING", "OPENAI_EMBEDDING", "VERTEXAI_EMBEDDING",
	// "BEDROCK_EMBEDDING".
	Type string `json:"type"`
	JSON struct {
		Component respjson.Field
		Type      respjson.Field
		raw       string
	} `json:"-"`
}

// anyPipelineEmbeddingModelConfigEmbeddingConfig is implemented by each variant of
// [PipelineEmbeddingModelConfigEmbeddingConfigUnion] to add type safety for the
// return type of [PipelineEmbeddingModelConfigEmbeddingConfigUnion.AsAny]
type anyPipelineEmbeddingModelConfigEmbeddingConfig interface {
	implPipelineEmbeddingModelConfigEmbeddingConfigUnion()
}

func (AzureOpenAIEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion() {}
func (CohereEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion()      {}
func (GeminiEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion()      {}
func (HuggingFaceInferenceAPIEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion() {
}
func (OpenAIEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion()   {}
func (VertexAIEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion() {}
func (BedrockEmbeddingConfig) implPipelineEmbeddingModelConfigEmbeddingConfigUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PipelineEmbeddingModelConfigEmbeddingConfigUnion.AsAny().(type) {
//	case llamacloudprod.AzureOpenAIEmbeddingConfig:
//	case llamacloudprod.CohereEmbeddingConfig:
//	case llamacloudprod.GeminiEmbeddingConfig:
//	case llamacloudprod.HuggingFaceInferenceAPIEmbeddingConfig:
//	case llamacloudprod.OpenAIEmbeddingConfig:
//	case llamacloudprod.VertexAIEmbeddingConfig:
//	case llamacloudprod.BedrockEmbeddingConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsAny() anyPipelineEmbeddingModelConfigEmbeddingConfig {
	switch u.Type {
	case "AZURE_EMBEDDING":
		return u.AsAzureEmbedding()
	case "COHERE_EMBEDDING":
		return u.AsCohereEmbedding()
	case "GEMINI_EMBEDDING":
		return u.AsGeminiEmbedding()
	case "HUGGINGFACE_API_EMBEDDING":
		return u.AsHuggingfaceAPIEmbedding()
	case "OPENAI_EMBEDDING":
		return u.AsOpenAIEmbedding()
	case "VERTEXAI_EMBEDDING":
		return u.AsVertexaiEmbedding()
	case "BEDROCK_EMBEDDING":
		return u.AsBedrockEmbedding()
	}
	return nil
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsAzureEmbedding() (v AzureOpenAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsCohereEmbedding() (v CohereEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsGeminiEmbedding() (v GeminiEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsHuggingfaceAPIEmbedding() (v HuggingFaceInferenceAPIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsOpenAIEmbedding() (v OpenAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsVertexaiEmbedding() (v VertexAIEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) AsBedrockEmbedding() (v BedrockEmbeddingConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineEmbeddingModelConfigEmbeddingConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineEmbeddingModelConfigEmbeddingConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineEmbeddingModelConfigEmbeddingConfigUnionComponent is an implicit
// subunion of [PipelineEmbeddingModelConfigEmbeddingConfigUnion].
// PipelineEmbeddingModelConfigEmbeddingConfigUnionComponent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PipelineEmbeddingModelConfigEmbeddingConfigUnion].
type PipelineEmbeddingModelConfigEmbeddingConfigUnionComponent struct {
	AdditionalKwargs any    `json:"additional_kwargs"`
	APIBase          string `json:"api_base"`
	APIKey           string `json:"api_key"`
	APIVersion       string `json:"api_version"`
	// This field is from variant [AzureOpenAIEmbedding].
	AzureDeployment string `json:"azure_deployment"`
	// This field is from variant [AzureOpenAIEmbedding].
	AzureEndpoint  string  `json:"azure_endpoint"`
	ClassName      string  `json:"class_name"`
	DefaultHeaders string  `json:"default_headers"`
	Dimensions     int64   `json:"dimensions"`
	EmbedBatchSize int64   `json:"embed_batch_size"`
	MaxRetries     int64   `json:"max_retries"`
	ModelName      string  `json:"model_name"`
	NumWorkers     int64   `json:"num_workers"`
	ReuseClient    bool    `json:"reuse_client"`
	Timeout        float64 `json:"timeout"`
	// This field is from variant [CohereEmbedding].
	EmbeddingType string `json:"embedding_type"`
	// This field is from variant [CohereEmbedding].
	InputType string `json:"input_type"`
	// This field is from variant [CohereEmbedding].
	Truncate string `json:"truncate"`
	// This field is from variant [GeminiEmbedding].
	OutputDimensionality int64 `json:"output_dimensionality"`
	// This field is from variant [GeminiEmbedding].
	TaskType string `json:"task_type"`
	// This field is from variant [GeminiEmbedding].
	Title string `json:"title"`
	// This field is from variant [GeminiEmbedding].
	Transport string `json:"transport"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Token HuggingFaceInferenceAPIEmbeddingTokenUnion `json:"token"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Cookies map[string]string `json:"cookies"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Headers map[string]string `json:"headers"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Pooling HuggingFaceInferenceAPIEmbeddingPooling `json:"pooling"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	QueryInstruction string `json:"query_instruction"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	Task string `json:"task"`
	// This field is from variant [HuggingFaceInferenceAPIEmbedding].
	TextInstruction string `json:"text_instruction"`
	// This field is from variant [VertexTextEmbedding].
	ClientEmail string `json:"client_email"`
	// This field is from variant [VertexTextEmbedding].
	Location string `json:"location"`
	// This field is from variant [VertexTextEmbedding].
	PrivateKey string `json:"private_key"`
	// This field is from variant [VertexTextEmbedding].
	PrivateKeyID string `json:"private_key_id"`
	// This field is from variant [VertexTextEmbedding].
	Project string `json:"project"`
	// This field is from variant [VertexTextEmbedding].
	TokenUri string `json:"token_uri"`
	// This field is from variant [VertexTextEmbedding].
	EmbedMode VertexTextEmbeddingEmbedMode `json:"embed_mode"`
	// This field is from variant [BedrockEmbedding].
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// This field is from variant [BedrockEmbedding].
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// This field is from variant [BedrockEmbedding].
	AwsSessionToken string `json:"aws_session_token"`
	// This field is from variant [BedrockEmbedding].
	ProfileName string `json:"profile_name"`
	// This field is from variant [BedrockEmbedding].
	RegionName string `json:"region_name"`
	JSON       struct {
		AdditionalKwargs     respjson.Field
		APIBase              respjson.Field
		APIKey               respjson.Field
		APIVersion           respjson.Field
		AzureDeployment      respjson.Field
		AzureEndpoint        respjson.Field
		ClassName            respjson.Field
		DefaultHeaders       respjson.Field
		Dimensions           respjson.Field
		EmbedBatchSize       respjson.Field
		MaxRetries           respjson.Field
		ModelName            respjson.Field
		NumWorkers           respjson.Field
		ReuseClient          respjson.Field
		Timeout              respjson.Field
		EmbeddingType        respjson.Field
		InputType            respjson.Field
		Truncate             respjson.Field
		OutputDimensionality respjson.Field
		TaskType             respjson.Field
		Title                respjson.Field
		Transport            respjson.Field
		Token                respjson.Field
		Cookies              respjson.Field
		Headers              respjson.Field
		Pooling              respjson.Field
		QueryInstruction     respjson.Field
		Task                 respjson.Field
		TextInstruction      respjson.Field
		ClientEmail          respjson.Field
		Location             respjson.Field
		PrivateKey           respjson.Field
		PrivateKeyID         respjson.Field
		Project              respjson.Field
		TokenUri             respjson.Field
		EmbedMode            respjson.Field
		AwsAccessKeyID       respjson.Field
		AwsSecretAccessKey   respjson.Field
		AwsSessionToken      respjson.Field
		ProfileName          respjson.Field
		RegionName           respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *PipelineEmbeddingModelConfigEmbeddingConfigUnionComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the pipeline.
type PipelineStatus string

const (
	PipelineStatusCreated  PipelineStatus = "CREATED"
	PipelineStatusDeleting PipelineStatus = "DELETING"
)

// PipelineTransformConfigUnion contains all possible properties and values from
// [AutoTransformConfig], [AdvancedModeTransformConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PipelineTransformConfigUnion struct {
	// This field is from variant [AutoTransformConfig].
	ChunkOverlap int64 `json:"chunk_overlap"`
	// This field is from variant [AutoTransformConfig].
	ChunkSize int64  `json:"chunk_size"`
	Mode      string `json:"mode"`
	// This field is from variant [AdvancedModeTransformConfig].
	ChunkingConfig AdvancedModeTransformConfigChunkingConfigUnion `json:"chunking_config"`
	// This field is from variant [AdvancedModeTransformConfig].
	SegmentationConfig AdvancedModeTransformConfigSegmentationConfigUnion `json:"segmentation_config"`
	JSON               struct {
		ChunkOverlap       respjson.Field
		ChunkSize          respjson.Field
		Mode               respjson.Field
		ChunkingConfig     respjson.Field
		SegmentationConfig respjson.Field
		raw                string
	} `json:"-"`
}

func (u PipelineTransformConfigUnion) AsAutoTransformConfig() (v AutoTransformConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineTransformConfigUnion) AsAdvancedModeTransformConfig() (v AdvancedModeTransformConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineTransformConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineTransformConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating a pipeline.
//
// The property Name is required.
type PipelineCreateParam struct {
	Name string `json:"name" api:"required"`
	// Data sink ID. When provided instead of data_sink, the data sink will be looked
	// up by ID.
	DataSinkID param.Opt[string] `json:"data_sink_id,omitzero" format:"uuid"`
	// Embedding model config ID. When provided instead of embedding_config, the
	// embedding model config will be looked up by ID.
	EmbeddingModelConfigID param.Opt[string] `json:"embedding_model_config_id,omitzero" format:"uuid"`
	// The ID of the ManagedPipeline this playground pipeline is linked to.
	ManagedPipelineID param.Opt[string] `json:"managed_pipeline_id,omitzero" format:"uuid"`
	// Status of the pipeline deployment.
	Status          param.Opt[string]                       `json:"status,omitzero"`
	EmbeddingConfig PipelineCreateEmbeddingConfigUnionParam `json:"embedding_config,omitzero"`
	// Configuration for the transformation.
	TransformConfig PipelineCreateTransformConfigUnionParam `json:"transform_config,omitzero"`
	// Schema for creating a data sink.
	DataSink DataSinkCreateParam `json:"data_sink,omitzero"`
	// Settings that can be configured for how to use LlamaParse to parse files within
	// a LlamaCloud pipeline.
	LlamaParseParameters LlamaParseParameters `json:"llama_parse_parameters,omitzero"`
	// Metadata configuration for the pipeline.
	MetadataConfig PipelineMetadataConfigParam `json:"metadata_config,omitzero"`
	// Type of pipeline. Either PLAYGROUND or MANAGED.
	//
	// Any of "PLAYGROUND", "MANAGED".
	PipelineType PipelineType `json:"pipeline_type,omitzero"`
	// Preset retrieval parameters for the pipeline.
	PresetRetrievalParameters PresetRetrievalParams `json:"preset_retrieval_parameters,omitzero"`
	// Configuration for sparse embedding models used in hybrid search.
	//
	// This allows users to choose between Splade and BM25 models for sparse retrieval
	// in managed data sinks.
	SparseModelConfig SparseModelConfigParam `json:"sparse_model_config,omitzero"`
	paramObj
}

func (r PipelineCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelineCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineCreateEmbeddingConfigUnionParam struct {
	OfAzureEmbedding          *AzureOpenAIEmbeddingConfigParam             `json:",omitzero,inline"`
	OfCohereEmbedding         *CohereEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfGeminiEmbedding         *GeminiEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfHuggingfaceAPIEmbedding *HuggingFaceInferenceAPIEmbeddingConfigParam `json:",omitzero,inline"`
	OfOpenAIEmbedding         *OpenAIEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfVertexaiEmbedding       *VertexAIEmbeddingConfigParam                `json:",omitzero,inline"`
	OfBedrockEmbedding        *BedrockEmbeddingConfigParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineCreateEmbeddingConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAzureEmbedding,
		u.OfCohereEmbedding,
		u.OfGeminiEmbedding,
		u.OfHuggingfaceAPIEmbedding,
		u.OfOpenAIEmbedding,
		u.OfVertexaiEmbedding,
		u.OfBedrockEmbedding)
}
func (u *PipelineCreateEmbeddingConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PipelineCreateEmbeddingConfigUnionParam](
		"type",
		apijson.Discriminator[AzureOpenAIEmbeddingConfigParam]("AZURE_EMBEDDING"),
		apijson.Discriminator[CohereEmbeddingConfigParam]("COHERE_EMBEDDING"),
		apijson.Discriminator[GeminiEmbeddingConfigParam]("GEMINI_EMBEDDING"),
		apijson.Discriminator[HuggingFaceInferenceAPIEmbeddingConfigParam]("HUGGINGFACE_API_EMBEDDING"),
		apijson.Discriminator[OpenAIEmbeddingConfigParam]("OPENAI_EMBEDDING"),
		apijson.Discriminator[VertexAIEmbeddingConfigParam]("VERTEXAI_EMBEDDING"),
		apijson.Discriminator[BedrockEmbeddingConfigParam]("BEDROCK_EMBEDDING"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineCreateTransformConfigUnionParam struct {
	OfAutoTransformConfig         *AutoTransformConfigParam         `json:",omitzero,inline"`
	OfAdvancedModeTransformConfig *AdvancedModeTransformConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineCreateTransformConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoTransformConfig, u.OfAdvancedModeTransformConfig)
}
func (u *PipelineCreateTransformConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PipelineMetadataConfig struct {
	// List of metadata keys to exclude from embeddings
	ExcludedEmbedMetadataKeys []string `json:"excluded_embed_metadata_keys"`
	// List of metadata keys to exclude from LLM during retrieval
	ExcludedLlmMetadataKeys []string `json:"excluded_llm_metadata_keys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedEmbedMetadataKeys respjson.Field
		ExcludedLlmMetadataKeys   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineMetadataConfig) RawJSON() string { return r.JSON.raw }
func (r *PipelineMetadataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PipelineMetadataConfig to a PipelineMetadataConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PipelineMetadataConfigParam.Overrides()
func (r PipelineMetadataConfig) ToParam() PipelineMetadataConfigParam {
	return param.Override[PipelineMetadataConfigParam](json.RawMessage(r.RawJSON()))
}

type PipelineMetadataConfigParam struct {
	// List of metadata keys to exclude from embeddings
	ExcludedEmbedMetadataKeys []string `json:"excluded_embed_metadata_keys,omitzero"`
	// List of metadata keys to exclude from LLM during retrieval
	ExcludedLlmMetadataKeys []string `json:"excluded_llm_metadata_keys,omitzero"`
	paramObj
}

func (r PipelineMetadataConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PipelineMetadataConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineMetadataConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enum for representing the type of a pipeline
type PipelineType string

const (
	PipelineTypePlayground PipelineType = "PLAYGROUND"
	PipelineTypeManaged    PipelineType = "MANAGED"
)

// Schema for the search params for an retrieval execution that can be preset for a
// pipeline.
type PresetRetrievalParamsResp struct {
	// Alpha value for hybrid retrieval to determine the weights between dense and
	// sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.
	Alpha     float64 `json:"alpha" api:"nullable"`
	ClassName string  `json:"class_name"`
	// Minimum similarity score wrt query for retrieval
	DenseSimilarityCutoff float64 `json:"dense_similarity_cutoff" api:"nullable"`
	// Number of nodes for dense retrieval.
	DenseSimilarityTopK int64 `json:"dense_similarity_top_k" api:"nullable"`
	// Enable reranking for retrieval
	EnableReranking bool `json:"enable_reranking" api:"nullable"`
	// Number of files to retrieve (only for retrieval mode files_via_metadata and
	// files_via_content).
	FilesTopK int64 `json:"files_top_k" api:"nullable"`
	// Number of reranked nodes for returning.
	RerankTopN int64 `json:"rerank_top_n" api:"nullable"`
	// The retrieval mode for the query.
	//
	// Any of "chunks", "files_via_metadata", "files_via_content", "auto_routed".
	RetrievalMode RetrievalMode `json:"retrieval_mode"`
	// Whether to retrieve image nodes.
	//
	// Deprecated: deprecated
	RetrieveImageNodes bool `json:"retrieve_image_nodes"`
	// Whether to retrieve page figure nodes.
	RetrievePageFigureNodes bool `json:"retrieve_page_figure_nodes"`
	// Whether to retrieve page screenshot nodes.
	RetrievePageScreenshotNodes bool `json:"retrieve_page_screenshot_nodes"`
	// Metadata filters for vector stores.
	SearchFilters MetadataFilters `json:"search_filters" api:"nullable"`
	// JSON Schema that will be used to infer search_filters. Omit or leave as null to
	// skip inference.
	SearchFiltersInferenceSchema map[string]*PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp `json:"search_filters_inference_schema" api:"nullable"`
	// Number of nodes for sparse retrieval.
	SparseSimilarityTopK int64 `json:"sparse_similarity_top_k" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha                        respjson.Field
		ClassName                    respjson.Field
		DenseSimilarityCutoff        respjson.Field
		DenseSimilarityTopK          respjson.Field
		EnableReranking              respjson.Field
		FilesTopK                    respjson.Field
		RerankTopN                   respjson.Field
		RetrievalMode                respjson.Field
		RetrieveImageNodes           respjson.Field
		RetrievePageFigureNodes      respjson.Field
		RetrievePageScreenshotNodes  respjson.Field
		SearchFilters                respjson.Field
		SearchFiltersInferenceSchema respjson.Field
		SparseSimilarityTopK         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PresetRetrievalParamsResp) RawJSON() string { return r.JSON.raw }
func (r *PresetRetrievalParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PresetRetrievalParamsResp to a PresetRetrievalParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PresetRetrievalParams.Overrides()
func (r PresetRetrievalParamsResp) ToParam() PresetRetrievalParams {
	return param.Override[PresetRetrievalParams](json.RawMessage(r.RawJSON()))
}

// PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp contains all possible
// properties and values from [map[string]any], [[]any], [string], [float64],
// [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPresetRetrievalsSearchFiltersInferenceSchemaMapItem OfAnyArray
// OfString OfFloat OfBool]
type PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPresetRetrievalsSearchFiltersInferenceSchemaMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPresetRetrievalsSearchFiltersInferenceSchemaMapItem respjson.Field
		OfAnyArray                                            respjson.Field
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *PresetRetrievalParamsSearchFiltersInferenceSchemaUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for the search params for an retrieval execution that can be preset for a
// pipeline.
type PresetRetrievalParams struct {
	// Alpha value for hybrid retrieval to determine the weights between dense and
	// sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Minimum similarity score wrt query for retrieval
	DenseSimilarityCutoff param.Opt[float64] `json:"dense_similarity_cutoff,omitzero"`
	// Number of nodes for dense retrieval.
	DenseSimilarityTopK param.Opt[int64] `json:"dense_similarity_top_k,omitzero"`
	// Enable reranking for retrieval
	EnableReranking param.Opt[bool] `json:"enable_reranking,omitzero"`
	// Number of files to retrieve (only for retrieval mode files_via_metadata and
	// files_via_content).
	FilesTopK param.Opt[int64] `json:"files_top_k,omitzero"`
	// Number of reranked nodes for returning.
	RerankTopN param.Opt[int64] `json:"rerank_top_n,omitzero"`
	// Number of nodes for sparse retrieval.
	SparseSimilarityTopK param.Opt[int64]  `json:"sparse_similarity_top_k,omitzero"`
	ClassName            param.Opt[string] `json:"class_name,omitzero"`
	// Whether to retrieve image nodes.
	//
	// Deprecated: deprecated
	RetrieveImageNodes param.Opt[bool] `json:"retrieve_image_nodes,omitzero"`
	// Whether to retrieve page figure nodes.
	RetrievePageFigureNodes param.Opt[bool] `json:"retrieve_page_figure_nodes,omitzero"`
	// Whether to retrieve page screenshot nodes.
	RetrievePageScreenshotNodes param.Opt[bool] `json:"retrieve_page_screenshot_nodes,omitzero"`
	// JSON Schema that will be used to infer search_filters. Omit or leave as null to
	// skip inference.
	SearchFiltersInferenceSchema map[string]*PresetRetrievalParamsSearchFiltersInferenceSchemaUnion `json:"search_filters_inference_schema,omitzero"`
	// The retrieval mode for the query.
	//
	// Any of "chunks", "files_via_metadata", "files_via_content", "auto_routed".
	RetrievalMode RetrievalMode `json:"retrieval_mode,omitzero"`
	// Metadata filters for vector stores.
	SearchFilters MetadataFiltersParam `json:"search_filters,omitzero"`
	paramObj
}

func (r PresetRetrievalParams) MarshalJSON() (data []byte, err error) {
	type shadow PresetRetrievalParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PresetRetrievalParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PresetRetrievalParamsSearchFiltersInferenceSchemaUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u PresetRetrievalParamsSearchFiltersInferenceSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *PresetRetrievalParamsSearchFiltersInferenceSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type RetrievalMode string

const (
	RetrievalModeChunks           RetrievalMode = "chunks"
	RetrievalModeFilesViaMetadata RetrievalMode = "files_via_metadata"
	RetrievalModeFilesViaContent  RetrievalMode = "files_via_content"
	RetrievalModeAutoRouted       RetrievalMode = "auto_routed"
)

// Configuration for sparse embedding models used in hybrid search.
//
// This allows users to choose between Splade and BM25 models for sparse retrieval
// in managed data sinks.
type SparseModelConfig struct {
	ClassName string `json:"class_name"`
	// The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default
	// for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based
	// on deployment mode (BYOC uses term frequency, Cloud uses Splade).
	//
	// Any of "splade", "bm25", "auto".
	ModelType SparseModelConfigModelType `json:"model_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName   respjson.Field
		ModelType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SparseModelConfig) RawJSON() string { return r.JSON.raw }
func (r *SparseModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SparseModelConfig to a SparseModelConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SparseModelConfigParam.Overrides()
func (r SparseModelConfig) ToParam() SparseModelConfigParam {
	return param.Override[SparseModelConfigParam](json.RawMessage(r.RawJSON()))
}

// The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default
// for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based
// on deployment mode (BYOC uses term frequency, Cloud uses Splade).
type SparseModelConfigModelType string

const (
	SparseModelConfigModelTypeSplade SparseModelConfigModelType = "splade"
	SparseModelConfigModelTypeBm25   SparseModelConfigModelType = "bm25"
	SparseModelConfigModelTypeAuto   SparseModelConfigModelType = "auto"
)

// Configuration for sparse embedding models used in hybrid search.
//
// This allows users to choose between Splade and BM25 models for sparse retrieval
// in managed data sinks.
type SparseModelConfigParam struct {
	ClassName param.Opt[string] `json:"class_name,omitzero"`
	// The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default
	// for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based
	// on deployment mode (BYOC uses term frequency, Cloud uses Splade).
	//
	// Any of "splade", "bm25", "auto".
	ModelType SparseModelConfigModelType `json:"model_type,omitzero"`
	paramObj
}

func (r SparseModelConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SparseModelConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SparseModelConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VertexAIEmbeddingConfig struct {
	// Configuration for the VertexAI embedding model.
	Component VertexTextEmbedding `json:"component"`
	// Type of the embedding model.
	//
	// Any of "VERTEXAI_EMBEDDING".
	Type VertexAIEmbeddingConfigType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Component   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VertexAIEmbeddingConfig) RawJSON() string { return r.JSON.raw }
func (r *VertexAIEmbeddingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VertexAIEmbeddingConfig to a VertexAIEmbeddingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VertexAIEmbeddingConfigParam.Overrides()
func (r VertexAIEmbeddingConfig) ToParam() VertexAIEmbeddingConfigParam {
	return param.Override[VertexAIEmbeddingConfigParam](json.RawMessage(r.RawJSON()))
}

// Type of the embedding model.
type VertexAIEmbeddingConfigType string

const (
	VertexAIEmbeddingConfigTypeVertexaiEmbedding VertexAIEmbeddingConfigType = "VERTEXAI_EMBEDDING"
)

type VertexAIEmbeddingConfigParam struct {
	// Configuration for the VertexAI embedding model.
	Component VertexTextEmbeddingParam `json:"component,omitzero"`
	// Type of the embedding model.
	//
	// Any of "VERTEXAI_EMBEDDING".
	Type VertexAIEmbeddingConfigType `json:"type,omitzero"`
	paramObj
}

func (r VertexAIEmbeddingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow VertexAIEmbeddingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VertexAIEmbeddingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VertexTextEmbedding struct {
	// The client email for the VertexAI credentials.
	ClientEmail string `json:"client_email" api:"required"`
	// The default location to use when making API calls.
	Location string `json:"location" api:"required"`
	// The private key for the VertexAI credentials.
	PrivateKey string `json:"private_key" api:"required"`
	// The private key ID for the VertexAI credentials.
	PrivateKeyID string `json:"private_key_id" api:"required"`
	// The default GCP project to use when making Vertex API calls.
	Project string `json:"project" api:"required"`
	// The token URI for the VertexAI credentials.
	TokenUri string `json:"token_uri" api:"required"`
	// Additional kwargs for the Vertex.
	AdditionalKwargs map[string]any `json:"additional_kwargs"`
	ClassName        string         `json:"class_name"`
	// The batch size for embedding calls.
	EmbedBatchSize int64 `json:"embed_batch_size"`
	// The embedding mode to use.
	//
	// Any of "default", "classification", "clustering", "similarity", "retrieval".
	EmbedMode VertexTextEmbeddingEmbedMode `json:"embed_mode"`
	// The modelId of the VertexAI model to use.
	ModelName string `json:"model_name"`
	// The number of workers to use for async embedding calls.
	NumWorkers int64 `json:"num_workers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientEmail      respjson.Field
		Location         respjson.Field
		PrivateKey       respjson.Field
		PrivateKeyID     respjson.Field
		Project          respjson.Field
		TokenUri         respjson.Field
		AdditionalKwargs respjson.Field
		ClassName        respjson.Field
		EmbedBatchSize   respjson.Field
		EmbedMode        respjson.Field
		ModelName        respjson.Field
		NumWorkers       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VertexTextEmbedding) RawJSON() string { return r.JSON.raw }
func (r *VertexTextEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VertexTextEmbedding to a VertexTextEmbeddingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VertexTextEmbeddingParam.Overrides()
func (r VertexTextEmbedding) ToParam() VertexTextEmbeddingParam {
	return param.Override[VertexTextEmbeddingParam](json.RawMessage(r.RawJSON()))
}

// The embedding mode to use.
type VertexTextEmbeddingEmbedMode string

const (
	VertexTextEmbeddingEmbedModeDefault        VertexTextEmbeddingEmbedMode = "default"
	VertexTextEmbeddingEmbedModeClassification VertexTextEmbeddingEmbedMode = "classification"
	VertexTextEmbeddingEmbedModeClustering     VertexTextEmbeddingEmbedMode = "clustering"
	VertexTextEmbeddingEmbedModeSimilarity     VertexTextEmbeddingEmbedMode = "similarity"
	VertexTextEmbeddingEmbedModeRetrieval      VertexTextEmbeddingEmbedMode = "retrieval"
)

// The properties ClientEmail, Location, PrivateKey, PrivateKeyID, Project,
// TokenUri are required.
type VertexTextEmbeddingParam struct {
	// The client email for the VertexAI credentials.
	ClientEmail param.Opt[string] `json:"client_email,omitzero" api:"required"`
	// The private key for the VertexAI credentials.
	PrivateKey param.Opt[string] `json:"private_key,omitzero" api:"required"`
	// The private key ID for the VertexAI credentials.
	PrivateKeyID param.Opt[string] `json:"private_key_id,omitzero" api:"required"`
	// The token URI for the VertexAI credentials.
	TokenUri param.Opt[string] `json:"token_uri,omitzero" api:"required"`
	// The default location to use when making API calls.
	Location string `json:"location" api:"required"`
	// The default GCP project to use when making Vertex API calls.
	Project string `json:"project" api:"required"`
	// The number of workers to use for async embedding calls.
	NumWorkers param.Opt[int64]  `json:"num_workers,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The batch size for embedding calls.
	EmbedBatchSize param.Opt[int64] `json:"embed_batch_size,omitzero"`
	// The modelId of the VertexAI model to use.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// Additional kwargs for the Vertex.
	AdditionalKwargs map[string]any `json:"additional_kwargs,omitzero"`
	// The embedding mode to use.
	//
	// Any of "default", "classification", "clustering", "similarity", "retrieval".
	EmbedMode VertexTextEmbeddingEmbedMode `json:"embed_mode,omitzero"`
	paramObj
}

func (r VertexTextEmbeddingParam) MarshalJSON() (data []byte, err error) {
	type shadow VertexTextEmbeddingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VertexTextEmbeddingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for the result of an retrieval execution.
type PipelineRunSearchResponse struct {
	// The ID of the pipeline that the query was retrieved against.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	// The nodes retrieved by the pipeline for the given query.
	RetrievalNodes []PipelineRunSearchResponseRetrievalNode `json:"retrieval_nodes" api:"required"`
	ClassName      string                                   `json:"class_name"`
	// The image nodes retrieved by the pipeline for the given query. Deprecated - will
	// soon be replaced with 'page_screenshot_nodes'.
	//
	// Deprecated: deprecated
	ImageNodes []PageScreenshotNodeWithScore `json:"image_nodes"`
	// Metadata filters for vector stores.
	InferredSearchFilters MetadataFilters `json:"inferred_search_filters" api:"nullable"`
	// Metadata associated with the retrieval execution
	Metadata map[string]string `json:"metadata"`
	// The page figure nodes retrieved by the pipeline for the given query.
	PageFigureNodes []PageFigureNodeWithScore `json:"page_figure_nodes"`
	// The end-to-end latency for retrieval and reranking.
	RetrievalLatency map[string]float64 `json:"retrieval_latency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PipelineID            respjson.Field
		RetrievalNodes        respjson.Field
		ClassName             respjson.Field
		ImageNodes            respjson.Field
		InferredSearchFilters respjson.Field
		Metadata              respjson.Field
		PageFigureNodes       respjson.Field
		RetrievalLatency      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineRunSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *PipelineRunSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Same as NodeWithScore but type for node is a TextNode instead of BaseNode.
// FastAPI doesn't accept abstract classes like BaseNode.
type PipelineRunSearchResponseRetrievalNode struct {
	// Provided for backward compatibility.
	Node      TextNode `json:"node" api:"required"`
	ClassName string   `json:"class_name"`
	Score     float64  `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Node        respjson.Field
		ClassName   respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineRunSearchResponseRetrievalNode) RawJSON() string { return r.JSON.raw }
func (r *PipelineRunSearchResponseRetrievalNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineNewParams struct {
	// Schema for creating a pipeline.
	PipelineCreate PipelineCreateParam
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PipelineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineCreate)
}
func (r *PipelineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PipelineNewParams]'s query parameters as `url.Values`.
func (r PipelineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineUpdateParams struct {
	// Data sink ID. When provided instead of data_sink, the data sink will be looked
	// up by ID.
	DataSinkID param.Opt[string] `json:"data_sink_id,omitzero" format:"uuid"`
	// Embedding model config ID. When provided instead of embedding_config, the
	// embedding model config will be looked up by ID.
	EmbeddingModelConfigID param.Opt[string] `json:"embedding_model_config_id,omitzero" format:"uuid"`
	// The ID of the ManagedPipeline this playground pipeline is linked to.
	ManagedPipelineID param.Opt[string] `json:"managed_pipeline_id,omitzero" format:"uuid"`
	Name              param.Opt[string] `json:"name,omitzero"`
	// Status of the pipeline deployment.
	Status          param.Opt[string]                        `json:"status,omitzero"`
	EmbeddingConfig PipelineUpdateParamsEmbeddingConfigUnion `json:"embedding_config,omitzero"`
	// Configuration for the transformation.
	TransformConfig PipelineUpdateParamsTransformConfigUnion `json:"transform_config,omitzero"`
	// Schema for creating a data sink.
	DataSink DataSinkCreateParam `json:"data_sink,omitzero"`
	// Settings that can be configured for how to use LlamaParse to parse files within
	// a LlamaCloud pipeline.
	LlamaParseParameters LlamaParseParameters `json:"llama_parse_parameters,omitzero"`
	// Metadata configuration for the pipeline.
	MetadataConfig PipelineMetadataConfigParam `json:"metadata_config,omitzero"`
	// Schema for the search params for an retrieval execution that can be preset for a
	// pipeline.
	PresetRetrievalParameters PresetRetrievalParams `json:"preset_retrieval_parameters,omitzero"`
	// Configuration for sparse embedding models used in hybrid search.
	//
	// This allows users to choose between Splade and BM25 models for sparse retrieval
	// in managed data sinks.
	SparseModelConfig SparseModelConfigParam `json:"sparse_model_config,omitzero"`
	paramObj
}

func (r PipelineUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PipelineUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineUpdateParamsEmbeddingConfigUnion struct {
	OfAzureEmbedding          *AzureOpenAIEmbeddingConfigParam             `json:",omitzero,inline"`
	OfCohereEmbedding         *CohereEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfGeminiEmbedding         *GeminiEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfHuggingfaceAPIEmbedding *HuggingFaceInferenceAPIEmbeddingConfigParam `json:",omitzero,inline"`
	OfOpenAIEmbedding         *OpenAIEmbeddingConfigParam                  `json:",omitzero,inline"`
	OfVertexaiEmbedding       *VertexAIEmbeddingConfigParam                `json:",omitzero,inline"`
	OfBedrockEmbedding        *BedrockEmbeddingConfigParam                 `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineUpdateParamsEmbeddingConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAzureEmbedding,
		u.OfCohereEmbedding,
		u.OfGeminiEmbedding,
		u.OfHuggingfaceAPIEmbedding,
		u.OfOpenAIEmbedding,
		u.OfVertexaiEmbedding,
		u.OfBedrockEmbedding)
}
func (u *PipelineUpdateParamsEmbeddingConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[PipelineUpdateParamsEmbeddingConfigUnion](
		"type",
		apijson.Discriminator[AzureOpenAIEmbeddingConfigParam]("AZURE_EMBEDDING"),
		apijson.Discriminator[CohereEmbeddingConfigParam]("COHERE_EMBEDDING"),
		apijson.Discriminator[GeminiEmbeddingConfigParam]("GEMINI_EMBEDDING"),
		apijson.Discriminator[HuggingFaceInferenceAPIEmbeddingConfigParam]("HUGGINGFACE_API_EMBEDDING"),
		apijson.Discriminator[OpenAIEmbeddingConfigParam]("OPENAI_EMBEDDING"),
		apijson.Discriminator[VertexAIEmbeddingConfigParam]("VERTEXAI_EMBEDDING"),
		apijson.Discriminator[BedrockEmbeddingConfigParam]("BEDROCK_EMBEDDING"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineUpdateParamsTransformConfigUnion struct {
	OfAutoTransformConfig         *AutoTransformConfigParam         `json:",omitzero,inline"`
	OfAdvancedModeTransformConfig *AdvancedModeTransformConfigParam `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineUpdateParamsTransformConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoTransformConfig, u.OfAdvancedModeTransformConfig)
}
func (u *PipelineUpdateParamsTransformConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PipelineListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PipelineName   param.Opt[string] `query:"pipeline_name,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	ProjectName    param.Opt[string] `query:"project_name,omitzero" json:"-"`
	// Enum for representing the type of a pipeline
	//
	// Any of "PLAYGROUND", "MANAGED".
	PipelineType PipelineType `query:"pipeline_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineListParams]'s query parameters as `url.Values`.
func (r PipelineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineGetStatusParams struct {
	FullDetails param.Opt[bool] `query:"full_details,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineGetStatusParams]'s query parameters as
// `url.Values`.
func (r PipelineGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineRunSearchParams struct {
	// The query to retrieve against.
	Query          string            `json:"query" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Alpha value for hybrid retrieval to determine the weights between dense and
	// sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Minimum similarity score wrt query for retrieval
	DenseSimilarityCutoff param.Opt[float64] `json:"dense_similarity_cutoff,omitzero"`
	// Number of nodes for dense retrieval.
	DenseSimilarityTopK param.Opt[int64] `json:"dense_similarity_top_k,omitzero"`
	// Enable reranking for retrieval
	EnableReranking param.Opt[bool] `json:"enable_reranking,omitzero"`
	// Number of files to retrieve (only for retrieval mode files_via_metadata and
	// files_via_content).
	FilesTopK param.Opt[int64] `json:"files_top_k,omitzero"`
	// Number of reranked nodes for returning.
	RerankTopN param.Opt[int64] `json:"rerank_top_n,omitzero"`
	// Number of nodes for sparse retrieval.
	SparseSimilarityTopK param.Opt[int64]  `json:"sparse_similarity_top_k,omitzero"`
	ClassName            param.Opt[string] `json:"class_name,omitzero"`
	// Whether to retrieve image nodes.
	RetrieveImageNodes param.Opt[bool] `json:"retrieve_image_nodes,omitzero"`
	// Whether to retrieve page figure nodes.
	RetrievePageFigureNodes param.Opt[bool] `json:"retrieve_page_figure_nodes,omitzero"`
	// Whether to retrieve page screenshot nodes.
	RetrievePageScreenshotNodes param.Opt[bool] `json:"retrieve_page_screenshot_nodes,omitzero"`
	// JSON Schema that will be used to infer search_filters. Omit or leave as null to
	// skip inference.
	SearchFiltersInferenceSchema map[string]*PipelineRunSearchParamsSearchFiltersInferenceSchemaUnion `json:"search_filters_inference_schema,omitzero"`
	// The retrieval mode for the query.
	//
	// Any of "chunks", "files_via_metadata", "files_via_content", "auto_routed".
	RetrievalMode RetrievalMode `json:"retrieval_mode,omitzero"`
	// Metadata filters for vector stores.
	SearchFilters MetadataFiltersParam `json:"search_filters,omitzero"`
	paramObj
}

func (r PipelineRunSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow PipelineRunSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineRunSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PipelineRunSearchParams]'s query parameters as
// `url.Values`.
func (r PipelineRunSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineRunSearchParamsSearchFiltersInferenceSchemaUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineRunSearchParamsSearchFiltersInferenceSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *PipelineRunSearchParamsSearchFiltersInferenceSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PipelineUpsertParams struct {
	// Schema for creating a pipeline.
	PipelineCreate PipelineCreateParam
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PipelineUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PipelineCreate)
}
func (r *PipelineUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PipelineUpsertParams]'s query parameters as `url.Values`.
func (r PipelineUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
