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
)

// RetrieverService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetrieverService] method instead.
type RetrieverService struct {
	options   []option.RequestOption
	Retriever RetrieverRetrieverService
}

// NewRetrieverService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetrieverService(opts ...option.RequestOption) (r RetrieverService) {
	r = RetrieverService{}
	r.options = opts
	r.Retriever = NewRetrieverRetrieverService(opts...)
	return
}

// Create a new Retriever.
func (r *RetrieverService) New(ctx context.Context, params RetrieverNewParams, opts ...option.RequestOption) (res *Retriever, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrievers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update an existing Retriever.
func (r *RetrieverService) Update(ctx context.Context, retrieverID string, params RetrieverUpdateParams, opts ...option.RequestOption) (res *Retriever, err error) {
	opts = slices.Concat(r.options, opts)
	if retrieverID == "" {
		err = errors.New("missing required retriever_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/retrievers/%s", retrieverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List Retrievers for a project.
func (r *RetrieverService) List(ctx context.Context, query RetrieverListParams, opts ...option.RequestOption) (res *[]Retriever, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrievers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a Retriever by ID.
func (r *RetrieverService) Delete(ctx context.Context, retrieverID string, body RetrieverDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if retrieverID == "" {
		err = errors.New("missing required retriever_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/retrievers/%s", retrieverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Get a Retriever by ID.
func (r *RetrieverService) Get(ctx context.Context, retrieverID string, query RetrieverGetParams, opts ...option.RequestOption) (res *Retriever, err error) {
	opts = slices.Concat(r.options, opts)
	if retrieverID == "" {
		err = errors.New("missing required retriever_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/retrievers/%s", retrieverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve data using specified pipelines without creating a persistent retriever.
func (r *RetrieverService) Search(ctx context.Context, params RetrieverSearchParams, opts ...option.RequestOption) (res *CompositeRetrievalResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrievers/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Upsert a new Retriever.
func (r *RetrieverService) Upsert(ctx context.Context, params RetrieverUpsertParams, opts ...option.RequestOption) (res *Retriever, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/retrievers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Enum for the mode of composite retrieval.
type CompositeRetrievalMode string

const (
	CompositeRetrievalModeRouting CompositeRetrievalMode = "routing"
	CompositeRetrievalModeFull    CompositeRetrievalMode = "full"
)

type CompositeRetrievalResult struct {
	// The image nodes retrieved by the pipeline for the given query. Deprecated - will
	// soon be replaced with 'page_screenshot_nodes'.
	//
	// Deprecated: deprecated
	ImageNodes []PageScreenshotNodeWithScore `json:"image_nodes"`
	// The retrieved nodes from the composite retrieval.
	Nodes []CompositeRetrievalResultNode `json:"nodes"`
	// The page figure nodes retrieved by the pipeline for the given query.
	PageFigureNodes []PageFigureNodeWithScore `json:"page_figure_nodes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageNodes      respjson.Field
		Nodes           respjson.Field
		PageFigureNodes respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompositeRetrievalResult) RawJSON() string { return r.JSON.raw }
func (r *CompositeRetrievalResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompositeRetrievalResultNode struct {
	Node      CompositeRetrievalResultNodeNode `json:"node" api:"required"`
	ClassName string                           `json:"class_name"`
	Score     float64                          `json:"score" api:"nullable"`
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
func (r CompositeRetrievalResultNode) RawJSON() string { return r.JSON.raw }
func (r *CompositeRetrievalResultNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompositeRetrievalResultNodeNode struct {
	// The ID of the retrieved node.
	ID string `json:"id" api:"required" format:"uuid"`
	// The end character index of the retrieved node in the document
	EndCharIdx int64 `json:"end_char_idx" api:"required"`
	// The ID of the pipeline this node was retrieved from.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	// The ID of the retriever this node was retrieved from.
	RetrieverID string `json:"retriever_id" api:"required" format:"uuid"`
	// The name of the retrieval pipeline this node was retrieved from.
	RetrieverPipelineName string `json:"retriever_pipeline_name" api:"required"`
	// The start character index of the retrieved node in the document
	StartCharIdx int64 `json:"start_char_idx" api:"required"`
	// The text of the retrieved node.
	Text string `json:"text" api:"required"`
	// Metadata associated with the retrieved node.
	Metadata map[string]any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		EndCharIdx            respjson.Field
		PipelineID            respjson.Field
		RetrieverID           respjson.Field
		RetrieverPipelineName respjson.Field
		StartCharIdx          respjson.Field
		Text                  respjson.Field
		Metadata              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompositeRetrievalResultNodeNode) RawJSON() string { return r.JSON.raw }
func (r *CompositeRetrievalResultNodeNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReRankConfigParam struct {
	// The number of nodes to retrieve after reranking over retrieved nodes from all
	// retrieval tools.
	TopN param.Opt[int64] `json:"top_n,omitzero"`
	// The type of reranker to use.
	//
	// Any of "system_default", "llm", "cohere", "bedrock", "score", "disabled".
	Type ReRankConfigType `json:"type,omitzero"`
	paramObj
}

func (r ReRankConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ReRankConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReRankConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of reranker to use.
type ReRankConfigType string

const (
	ReRankConfigTypeSystemDefault ReRankConfigType = "system_default"
	ReRankConfigTypeLlm           ReRankConfigType = "llm"
	ReRankConfigTypeCohere        ReRankConfigType = "cohere"
	ReRankConfigTypeBedrock       ReRankConfigType = "bedrock"
	ReRankConfigTypeScore         ReRankConfigType = "score"
	ReRankConfigTypeDisabled      ReRankConfigType = "disabled"
)

// An entity that retrieves context nodes from several sub RetrieverTools.
type Retriever struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// A name for the retriever tool. Will default to the pipeline name if not
	// provided.
	Name string `json:"name" api:"required"`
	// The ID of the project this retriever resides in.
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The pipelines this retriever uses.
	Pipelines []RetrieverPipeline `json:"pipelines"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		CreatedAt   respjson.Field
		Pipelines   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Retriever) RawJSON() string { return r.JSON.raw }
func (r *Retriever) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type RetrieverCreateParam struct {
	// A name for the retriever tool. Will default to the pipeline name if not
	// provided.
	Name string `json:"name" api:"required"`
	// The pipelines this retriever uses.
	Pipelines []RetrieverPipelineParam `json:"pipelines,omitzero"`
	paramObj
}

func (r RetrieverCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrieverCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrieverCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetrieverPipeline struct {
	// A description of the retriever tool.
	Description string `json:"description" api:"required"`
	// A name for the retriever tool. Will default to the pipeline name if not
	// provided.
	Name string `json:"name" api:"required"`
	// The ID of the pipeline this tool uses.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	// Parameters for retrieval configuration.
	PresetRetrievalParameters PresetRetrievalParamsResp `json:"preset_retrieval_parameters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description               respjson.Field
		Name                      respjson.Field
		PipelineID                respjson.Field
		PresetRetrievalParameters respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrieverPipeline) RawJSON() string { return r.JSON.raw }
func (r *RetrieverPipeline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RetrieverPipeline to a RetrieverPipelineParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RetrieverPipelineParam.Overrides()
func (r RetrieverPipeline) ToParam() RetrieverPipelineParam {
	return param.Override[RetrieverPipelineParam](json.RawMessage(r.RawJSON()))
}

// The properties Description, Name, PipelineID are required.
type RetrieverPipelineParam struct {
	// A description of the retriever tool.
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// A name for the retriever tool. Will default to the pipeline name if not
	// provided.
	Name param.Opt[string] `json:"name,omitzero" api:"required"`
	// The ID of the pipeline this tool uses.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	// Parameters for retrieval configuration.
	PresetRetrievalParameters PresetRetrievalParams `json:"preset_retrieval_parameters,omitzero"`
	paramObj
}

func (r RetrieverPipelineParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrieverPipelineParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrieverPipelineParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetrieverNewParams struct {
	RetrieverCreate RetrieverCreateParam
	OrganizationID  param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID       param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RetrieverNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RetrieverCreate)
}
func (r *RetrieverNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrieverNewParams]'s query parameters as `url.Values`.
func (r RetrieverNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverUpdateParams struct {
	// The pipelines this retriever uses.
	Pipelines      []RetrieverPipelineParam `json:"pipelines,omitzero" api:"required"`
	OrganizationID param.Opt[string]        `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]        `query:"project_id,omitzero" format:"uuid" json:"-"`
	// A name for the retriever.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r RetrieverUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrieverUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrieverUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrieverUpdateParams]'s query parameters as `url.Values`.
func (r RetrieverUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverListParams struct {
	Name           param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RetrieverListParams]'s query parameters as `url.Values`.
func (r RetrieverListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RetrieverDeleteParams]'s query parameters as `url.Values`.
func (r RetrieverDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [RetrieverGetParams]'s query parameters as `url.Values`.
func (r RetrieverGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverSearchParams struct {
	// The query to retrieve against.
	Query          string            `json:"query" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// (use rerank_config.top_n instead) The number of nodes to retrieve after
	// reranking over retrieved nodes from all retrieval tools.
	RerankTopN param.Opt[int64] `json:"rerank_top_n,omitzero"`
	// The mode of composite retrieval.
	//
	// Any of "routing", "full".
	Mode CompositeRetrievalMode `json:"mode,omitzero"`
	// The pipelines to use for retrieval.
	Pipelines []RetrieverPipelineParam `json:"pipelines,omitzero"`
	// The rerank configuration for composite retrieval.
	RerankConfig ReRankConfigParam `json:"rerank_config,omitzero"`
	paramObj
}

func (r RetrieverSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrieverSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrieverSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrieverSearchParams]'s query parameters as `url.Values`.
func (r RetrieverSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RetrieverUpsertParams struct {
	RetrieverCreate RetrieverCreateParam
	OrganizationID  param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID       param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r RetrieverUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RetrieverCreate)
}
func (r *RetrieverUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrieverUpsertParams]'s query parameters as `url.Values`.
func (r RetrieverUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
