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

// IndexService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndexService] method instead.
type IndexService struct {
	options []option.RequestOption
}

// NewIndexService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIndexService(opts ...option.RequestOption) (r IndexService) {
	r = IndexService{}
	r.options = opts
	return
}

// Create a searchable index over a source directory.
func (r *IndexService) New(ctx context.Context, params IndexNewParams, opts ...option.RequestOption) (res *IndexNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/indexes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List indexes for the current project.
func (r *IndexService) List(ctx context.Context, query IndexListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[IndexListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/indexes"
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

// List indexes for the current project.
func (r *IndexService) ListAutoPaging(ctx context.Context, query IndexListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[IndexListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an index.
func (r *IndexService) Delete(ctx context.Context, indexID string, body IndexDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if indexID == "" {
		err = errors.New("missing required index_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/indexes/%s", url.PathEscape(indexID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Cancel the running sync for an index. Returns 409 if no sync is running.
func (r *IndexService) CancelSync(ctx context.Context, indexID string, body IndexCancelSyncParams, opts ...option.RequestOption) (res *IndexCancelSyncResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if indexID == "" {
		err = errors.New("missing required index_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/indexes/%s/sync/cancel", url.PathEscape(indexID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an index by ID.
func (r *IndexService) Get(ctx context.Context, indexID string, query IndexGetParams, opts ...option.RequestOption) (res *IndexGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if indexID == "" {
		err = errors.New("missing required index_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/indexes/%s", url.PathEscape(indexID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Trigger a sync and export for an existing index, re-parsing changed files and
// exporting updated chunks.
func (r *IndexService) Sync(ctx context.Context, indexID string, body IndexSyncParams, opts ...option.RequestOption) (res *IndexSyncResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if indexID == "" {
		err = errors.New("missing required index_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/indexes/%s/sync", url.PathEscape(indexID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A searchable index over a directory of documents.
type IndexNewResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// ID of the export configuration.
	ExportConfigID string `json:"export_config_id" api:"required"`
	// Index name.
	Name string `json:"name" api:"required"`
	// ID of the output directory holding the indexed files.
	OutputDirectoryID string `json:"output_directory_id" api:"required"`
	// Project this index belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// ID of the source directory.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// ID of the sync configuration.
	SyncConfigID string `json:"sync_config_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Index description.
	Description string `json:"description" api:"nullable"`
	// Last export time.
	LastExportedAt time.Time `json:"last_exported_at" api:"nullable" format:"date-time"`
	// Last sync time.
	LastSyncedAt time.Time `json:"last_synced_at" api:"nullable" format:"date-time"`
	// Build state and diagnostic info.
	Metadata map[string]any `json:"metadata"`
	// Whether the index is syncing its source or exporting the result. Requires
	// `expand=sync_in_progress`.
	SyncInProgress bool `json:"sync_in_progress" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ExportConfigID    respjson.Field
		Name              respjson.Field
		OutputDirectoryID respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		SyncConfigID      respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		LastExportedAt    respjson.Field
		LastSyncedAt      respjson.Field
		Metadata          respjson.Field
		SyncInProgress    respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexNewResponse) RawJSON() string { return r.JSON.raw }
func (r *IndexNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A searchable index over a directory of documents.
type IndexListResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// ID of the export configuration.
	ExportConfigID string `json:"export_config_id" api:"required"`
	// Index name.
	Name string `json:"name" api:"required"`
	// ID of the output directory holding the indexed files.
	OutputDirectoryID string `json:"output_directory_id" api:"required"`
	// Project this index belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// ID of the source directory.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// ID of the sync configuration.
	SyncConfigID string `json:"sync_config_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Index description.
	Description string `json:"description" api:"nullable"`
	// Last export time.
	LastExportedAt time.Time `json:"last_exported_at" api:"nullable" format:"date-time"`
	// Last sync time.
	LastSyncedAt time.Time `json:"last_synced_at" api:"nullable" format:"date-time"`
	// Build state and diagnostic info.
	Metadata map[string]any `json:"metadata"`
	// Whether the index is syncing its source or exporting the result. Requires
	// `expand=sync_in_progress`.
	SyncInProgress bool `json:"sync_in_progress" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ExportConfigID    respjson.Field
		Name              respjson.Field
		OutputDirectoryID respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		SyncConfigID      respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		LastExportedAt    respjson.Field
		LastSyncedAt      respjson.Field
		Metadata          respjson.Field
		SyncInProgress    respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexListResponse) RawJSON() string { return r.JSON.raw }
func (r *IndexListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndexCancelSyncResponse = any

// A searchable index over a directory of documents.
type IndexGetResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// ID of the export configuration.
	ExportConfigID string `json:"export_config_id" api:"required"`
	// Index name.
	Name string `json:"name" api:"required"`
	// ID of the output directory holding the indexed files.
	OutputDirectoryID string `json:"output_directory_id" api:"required"`
	// Project this index belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// ID of the source directory.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// ID of the sync configuration.
	SyncConfigID string `json:"sync_config_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Index description.
	Description string `json:"description" api:"nullable"`
	// Last export time.
	LastExportedAt time.Time `json:"last_exported_at" api:"nullable" format:"date-time"`
	// Last sync time.
	LastSyncedAt time.Time `json:"last_synced_at" api:"nullable" format:"date-time"`
	// Build state and diagnostic info.
	Metadata map[string]any `json:"metadata"`
	// Whether the index is syncing its source or exporting the result. Requires
	// `expand=sync_in_progress`.
	SyncInProgress bool `json:"sync_in_progress" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ExportConfigID    respjson.Field
		Name              respjson.Field
		OutputDirectoryID respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		SyncConfigID      respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		LastExportedAt    respjson.Field
		LastSyncedAt      respjson.Field
		Metadata          respjson.Field
		SyncInProgress    respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IndexGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IndexGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IndexSyncResponse = any

type IndexNewParams struct {
	// ID of the source directory containing your documents.
	SourceDirectoryID string            `json:"source_directory_id" api:"required"`
	OrganizationID    param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID         param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Optional description of the index.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional display name for the index. If omitted, the index is named after the
	// source directory.
	Name param.Opt[string] `json:"name,omitzero"`
	// How often to re-run the sync. One of: manual, daily, on_source_change. Defaults
	// to manual.
	SyncFrequency param.Opt[string] `json:"sync_frequency,omitzero"`
	// Product configurations for syncing. Omit to use a default parse configuration.
	// Include an explicit entry per product type (e.g. parse, extract) to override the
	// default.
	Products []IndexNewParamsProduct `json:"products,omitzero"`
	// Attachment kinds to store alongside parsed output. Each entry must be one of:
	// screenshots, items. For example, ['screenshots'] renders and stores per-page
	// screenshots; ['items'] stores structured items with bounding boxes. Omit or pass
	// an empty list to skip attachments.
	StoreAttachments []string `json:"store_attachments,omitzero"`
	// Vector export destination for the index. 'DEFAULT' exports to the managed vector
	// DB destination resolved from configuration. 'DISABLED' skips vector export — the
	// export destination falls back to 'Download'.
	//
	// Any of "DEFAULT", "DISABLED".
	VectorTarget IndexNewParamsVectorTarget `json:"vector_target,omitzero"`
	paramObj
}

func (r IndexNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IndexNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndexNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IndexNewParams]'s query parameters as `url.Values`.
func (r IndexNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A product configuration to include in an index's sync.
//
// Structurally mirrors `directory_sync.SyncProductEntryRequest` but is a distinct
// class so the Index API surface stays SDK-gen-isolated from directory-sync
// internals. Translation between the two happens in `index/api_utils.py`.
//
// The properties ProductConfigID, ProductType are required.
type IndexNewParamsProduct struct {
	// ID of the product configuration.
	ProductConfigID string `json:"product_config_id" api:"required"`
	// Product type. One of: parse, extract.
	ProductType string `json:"product_type" api:"required"`
	paramObj
}

func (r IndexNewParamsProduct) MarshalJSON() (data []byte, err error) {
	type shadow IndexNewParamsProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IndexNewParamsProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector export destination for the index. 'DEFAULT' exports to the managed vector
// DB destination resolved from configuration. 'DISABLED' skips vector export — the
// export destination falls back to 'Download'.
type IndexNewParamsVectorTarget string

const (
	IndexNewParamsVectorTargetDefault  IndexNewParamsVectorTarget = "DEFAULT"
	IndexNewParamsVectorTargetDisabled IndexNewParamsVectorTarget = "DISABLED"
)

type IndexListParams struct {
	OrganizationID    param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize          param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken         param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID         param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	SourceDirectoryID param.Opt[string] `query:"source_directory_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndexListParams]'s query parameters as `url.Values`.
func (r IndexListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndexDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [IndexDeleteParams]'s query parameters as `url.Values`.
func (r IndexDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndexCancelSyncParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [IndexCancelSyncParams]'s query parameters as `url.Values`.
func (r IndexCancelSyncParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndexGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields to expand. Supported value: sync_in_progress.
	//
	// Any of "sync_in_progress".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IndexGetParams]'s query parameters as `url.Values`.
func (r IndexGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IndexSyncParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [IndexSyncParams]'s query parameters as `url.Values`.
func (r IndexSyncParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
