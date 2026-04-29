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

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/llamacloud-prod-go/internal/encoding/json"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// PipelineDocumentService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineDocumentService] method instead.
type PipelineDocumentService struct {
	options []option.RequestOption
}

// NewPipelineDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPipelineDocumentService(opts ...option.RequestOption) (r PipelineDocumentService) {
	r = PipelineDocumentService{}
	r.options = opts
	return
}

// Batch create documents for a pipeline.
func (r *PipelineDocumentService) New(ctx context.Context, pipelineID string, body PipelineDocumentNewParams, opts ...option.RequestOption) (res *[]CloudDocument, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Return a list of documents for a pipeline.
func (r *PipelineDocumentService) List(ctx context.Context, pipelineID string, query PipelineDocumentListParams, opts ...option.RequestOption) (res *pagination.PaginatedCloudDocuments[CloudDocument], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/paginated", pipelineID)
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

// Return a list of documents for a pipeline.
func (r *PipelineDocumentService) ListAutoPaging(ctx context.Context, pipelineID string, query PipelineDocumentListParams, opts ...option.RequestOption) *pagination.PaginatedCloudDocumentsAutoPager[CloudDocument] {
	return pagination.NewPaginatedCloudDocumentsAutoPager(r.List(ctx, pipelineID, query, opts...))
}

// Delete a document from a pipeline. Initiates an async job that will:
//
// 1. Delete vectors from the vector store
// 2. Delete the document from MongoDB after vectors are successfully deleted
func (r *PipelineDocumentService) Delete(ctx context.Context, documentID string, body PipelineDocumentDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return err
	}
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/%s", body.PipelineID, url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Return a single document for a pipeline.
func (r *PipelineDocumentService) Get(ctx context.Context, documentID string, query PipelineDocumentGetParams, opts ...option.RequestOption) (res *CloudDocument, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/%s", query.PipelineID, url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Return a list of chunks for a pipeline document.
func (r *PipelineDocumentService) GetChunks(ctx context.Context, documentID string, query PipelineDocumentGetChunksParams, opts ...option.RequestOption) (res *[]TextNode, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/%s/chunks", query.PipelineID, url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Return a single document for a pipeline.
func (r *PipelineDocumentService) GetStatus(ctx context.Context, documentID string, query PipelineDocumentGetStatusParams, opts ...option.RequestOption) (res *ManagedIngestionStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/%s/status", query.PipelineID, url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sync a specific document for a pipeline.
func (r *PipelineDocumentService) Sync(ctx context.Context, documentID string, body PipelineDocumentSyncParams, opts ...option.RequestOption) (res *PipelineDocumentSyncResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents/%s/sync", body.PipelineID, url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Batch create or update a document for a pipeline.
func (r *PipelineDocumentService) Upsert(ctx context.Context, pipelineID string, body PipelineDocumentUpsertParams, opts ...option.RequestOption) (res *[]CloudDocument, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/documents", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Cloud document stored in S3.
type CloudDocument struct {
	ID                        string         `json:"id" api:"required"`
	Metadata                  map[string]any `json:"metadata" api:"required"`
	Text                      string         `json:"text" api:"required"`
	ExcludedEmbedMetadataKeys []string       `json:"excluded_embed_metadata_keys"`
	ExcludedLlmMetadataKeys   []string       `json:"excluded_llm_metadata_keys"`
	// indices in the CloudDocument.text where a new page begins. e.g. Second page
	// starts at index specified by page_positions[1].
	PagePositions  []int64        `json:"page_positions" api:"nullable"`
	StatusMetadata map[string]any `json:"status_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Metadata                  respjson.Field
		Text                      respjson.Field
		ExcludedEmbedMetadataKeys respjson.Field
		ExcludedLlmMetadataKeys   respjson.Field
		PagePositions             respjson.Field
		StatusMetadata            respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudDocument) RawJSON() string { return r.JSON.raw }
func (r *CloudDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create a new cloud document.
//
// The properties Metadata, Text are required.
type CloudDocumentCreateParam struct {
	Metadata map[string]any    `json:"metadata,omitzero" api:"required"`
	Text     string            `json:"text" api:"required"`
	ID       param.Opt[string] `json:"id,omitzero"`
	// indices in the CloudDocument.text where a new page begins. e.g. Second page
	// starts at index specified by page_positions[1].
	PagePositions             []int64  `json:"page_positions,omitzero"`
	ExcludedEmbedMetadataKeys []string `json:"excluded_embed_metadata_keys,omitzero"`
	ExcludedLlmMetadataKeys   []string `json:"excluded_llm_metadata_keys,omitzero"`
	paramObj
}

func (r CloudDocumentCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudDocumentCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudDocumentCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provided for backward compatibility.
type TextNode struct {
	ClassName string `json:"class_name"`
	// Embedding of the node.
	Embedding []float64 `json:"embedding" api:"nullable"`
	// End char index of the node.
	EndCharIdx int64 `json:"end_char_idx" api:"nullable"`
	// Metadata keys that are excluded from text for the embed model.
	ExcludedEmbedMetadataKeys []string `json:"excluded_embed_metadata_keys"`
	// Metadata keys that are excluded from text for the LLM.
	ExcludedLlmMetadataKeys []string `json:"excluded_llm_metadata_keys"`
	// A flat dictionary of metadata fields
	ExtraInfo map[string]any `json:"extra_info"`
	// Unique ID of the node.
	ID string `json:"id_"`
	// Separator between metadata fields when converting to string.
	MetadataSeperator string `json:"metadata_seperator"`
	// Template for how metadata is formatted, with {key} and {value} placeholders.
	MetadataTemplate string `json:"metadata_template"`
	// MIME type of the node content.
	Mimetype string `json:"mimetype"`
	// A mapping of relationships to other node information.
	Relationships map[string]TextNodeRelationshipUnion `json:"relationships"`
	// Start char index of the node.
	StartCharIdx int64 `json:"start_char_idx" api:"nullable"`
	// Text content of the node.
	Text string `json:"text"`
	// Template for how text is formatted, with {content} and {metadata_str}
	// placeholders.
	TextTemplate string `json:"text_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassName                 respjson.Field
		Embedding                 respjson.Field
		EndCharIdx                respjson.Field
		ExcludedEmbedMetadataKeys respjson.Field
		ExcludedLlmMetadataKeys   respjson.Field
		ExtraInfo                 respjson.Field
		ID                        respjson.Field
		MetadataSeperator         respjson.Field
		MetadataTemplate          respjson.Field
		Mimetype                  respjson.Field
		Relationships             respjson.Field
		StartCharIdx              respjson.Field
		Text                      respjson.Field
		TextTemplate              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextNode) RawJSON() string { return r.JSON.raw }
func (r *TextNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TextNodeRelationshipUnion contains all possible properties and values from
// [TextNodeRelationshipRelatedNodeInfo], [[]TextNodeRelationshipArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTextNodeRelationshipArray]
type TextNodeRelationshipUnion struct {
	// This field will be present if the value is a [[]TextNodeRelationshipArrayItem]
	// instead of an object.
	OfTextNodeRelationshipArray []TextNodeRelationshipArrayItem `json:",inline"`
	// This field is from variant [TextNodeRelationshipRelatedNodeInfo].
	NodeID string `json:"node_id"`
	// This field is from variant [TextNodeRelationshipRelatedNodeInfo].
	ClassName string `json:"class_name"`
	// This field is from variant [TextNodeRelationshipRelatedNodeInfo].
	Hash string `json:"hash"`
	// This field is from variant [TextNodeRelationshipRelatedNodeInfo].
	Metadata map[string]any `json:"metadata"`
	// This field is from variant [TextNodeRelationshipRelatedNodeInfo].
	NodeType string `json:"node_type"`
	JSON     struct {
		OfTextNodeRelationshipArray respjson.Field
		NodeID                      respjson.Field
		ClassName                   respjson.Field
		Hash                        respjson.Field
		Metadata                    respjson.Field
		NodeType                    respjson.Field
		raw                         string
	} `json:"-"`
}

func (u TextNodeRelationshipUnion) AsRelatedNodeInfo() (v TextNodeRelationshipRelatedNodeInfo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextNodeRelationshipUnion) AsTextNodeRelationshipArray() (v []TextNodeRelationshipArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextNodeRelationshipUnion) RawJSON() string { return u.JSON.raw }

func (r *TextNodeRelationshipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextNodeRelationshipRelatedNodeInfo struct {
	NodeID    string         `json:"node_id" api:"required"`
	ClassName string         `json:"class_name"`
	Hash      string         `json:"hash" api:"nullable"`
	Metadata  map[string]any `json:"metadata"`
	NodeType  string         `json:"node_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NodeID      respjson.Field
		ClassName   respjson.Field
		Hash        respjson.Field
		Metadata    respjson.Field
		NodeType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextNodeRelationshipRelatedNodeInfo) RawJSON() string { return r.JSON.raw }
func (r *TextNodeRelationshipRelatedNodeInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextNodeRelationshipArrayItem struct {
	NodeID    string         `json:"node_id" api:"required"`
	ClassName string         `json:"class_name"`
	Hash      string         `json:"hash" api:"nullable"`
	Metadata  map[string]any `json:"metadata"`
	NodeType  string         `json:"node_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NodeID      respjson.Field
		ClassName   respjson.Field
		Hash        respjson.Field
		Metadata    respjson.Field
		NodeType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextNodeRelationshipArrayItem) RawJSON() string { return r.JSON.raw }
func (r *TextNodeRelationshipArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineDocumentSyncResponse = any

type PipelineDocumentNewParams struct {
	Body []CloudDocumentCreateParam
	paramObj
}

func (r PipelineDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PipelineDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineDocumentListParams struct {
	FileID                     param.Opt[string] `query:"file_id,omitzero" format:"uuid" json:"-"`
	OnlyAPIDataSourceDocuments param.Opt[bool]   `query:"only_api_data_source_documents,omitzero" json:"-"`
	OnlyDirectUpload           param.Opt[bool]   `query:"only_direct_upload,omitzero" json:"-"`
	Limit                      param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Skip                       param.Opt[int64]  `query:"skip,omitzero" json:"-"`
	// Any of "cached", "ttl".
	StatusRefreshPolicy PipelineDocumentListParamsStatusRefreshPolicy `query:"status_refresh_policy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineDocumentListParams]'s query parameters as
// `url.Values`.
func (r PipelineDocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineDocumentListParamsStatusRefreshPolicy string

const (
	PipelineDocumentListParamsStatusRefreshPolicyCached PipelineDocumentListParamsStatusRefreshPolicy = "cached"
	PipelineDocumentListParamsStatusRefreshPolicyTtl    PipelineDocumentListParamsStatusRefreshPolicy = "ttl"
)

type PipelineDocumentDeleteParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDocumentGetParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDocumentGetChunksParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDocumentGetStatusParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDocumentSyncParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDocumentUpsertParams struct {
	Body []CloudDocumentCreateParam
	paramObj
}

func (r PipelineDocumentUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PipelineDocumentUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
