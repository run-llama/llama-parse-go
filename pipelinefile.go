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

// PipelineFileService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineFileService] method instead.
type PipelineFileService struct {
	options []option.RequestOption
}

// NewPipelineFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineFileService(opts ...option.RequestOption) (r PipelineFileService) {
	r = PipelineFileService{}
	r.options = opts
	return
}

// Add files to a pipeline.
//
// Deprecated: deprecated
func (r *PipelineFileService) New(ctx context.Context, pipelineID string, body PipelineFileNewParams, opts ...option.RequestOption) (res *[]PipelineFile, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Update a file for a pipeline.
//
// Deprecated: deprecated
func (r *PipelineFileService) Update(ctx context.Context, fileID string, params PipelineFileUpdateParams, opts ...option.RequestOption) (res *PipelineFile, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files/%s", params.PipelineID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List files for a pipeline with optional filtering, sorting, and pagination.
//
// Deprecated: deprecated
func (r *PipelineFileService) List(ctx context.Context, pipelineID string, query PipelineFileListParams, opts ...option.RequestOption) (res *pagination.PaginatedPipelineFiles[PipelineFile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files2", pipelineID)
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

// List files for a pipeline with optional filtering, sorting, and pagination.
//
// Deprecated: deprecated
func (r *PipelineFileService) ListAutoPaging(ctx context.Context, pipelineID string, query PipelineFileListParams, opts ...option.RequestOption) *pagination.PaginatedPipelineFilesAutoPager[PipelineFile] {
	return pagination.NewPaginatedPipelineFilesAutoPager(r.List(ctx, pipelineID, query, opts...))
}

// Delete a file from a pipeline.
//
// Deprecated: deprecated
func (r *PipelineFileService) Delete(ctx context.Context, fileID string, body PipelineFileDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return err
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files/%s", body.PipelineID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get status of a file for a pipeline.
//
// Deprecated: deprecated
func (r *PipelineFileService) GetStatus(ctx context.Context, fileID string, query PipelineFileGetStatusParams, opts ...option.RequestOption) (res *ManagedIngestionStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files/%s/status", query.PipelineID, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get files for a pipeline.
//
// Deprecated: deprecated
func (r *PipelineFileService) GetStatusCounts(ctx context.Context, pipelineID string, query PipelineFileGetStatusCountsParams, opts ...option.RequestOption) (res *PipelineFileGetStatusCountsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/files/status-counts", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A file associated with a pipeline.
type PipelineFile struct {
	// Unique identifier for the pipeline file.
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the pipeline that the file is associated with.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	// Hashes for the configuration of the pipeline.
	ConfigHash map[string]*PipelineFileConfigHashUnion `json:"config_hash" api:"nullable"`
	// When the pipeline file was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Custom metadata for the file.
	CustomMetadata map[string]*PipelineFileCustomMetadataUnion `json:"custom_metadata" api:"nullable"`
	// The ID of the data source that the file belongs to.
	DataSourceID string `json:"data_source_id" api:"nullable" format:"uuid"`
	// The ID of the file in the external system.
	ExternalFileID string `json:"external_file_id" api:"nullable"`
	// The ID of the file.
	FileID string `json:"file_id" api:"nullable" format:"uuid"`
	// Size of the file in bytes.
	FileSize int64 `json:"file_size" api:"nullable"`
	// File type (e.g. pdf, docx, etc.).
	FileType string `json:"file_type" api:"nullable"`
	// The number of pages that have been indexed for this file.
	IndexedPageCount int64 `json:"indexed_page_count" api:"nullable"`
	// The last modified time of the file.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Name of the file.
	Name string `json:"name" api:"nullable"`
	// Permission information for the file.
	PermissionInfo map[string]*PipelineFilePermissionInfoUnion `json:"permission_info" api:"nullable"`
	// The ID of the project that the file belongs to.
	ProjectID string `json:"project_id" api:"nullable" format:"uuid"`
	// Resource information for the file.
	ResourceInfo map[string]*PipelineFileResourceInfoUnion `json:"resource_info" api:"nullable"`
	// Status of the pipeline file.
	//
	// Any of "CANCELLED", "ERROR", "IN_PROGRESS", "NOT_STARTED", "SUCCESS".
	Status PipelineFileStatus `json:"status" api:"nullable"`
	// The last time the status was updated.
	StatusUpdatedAt time.Time `json:"status_updated_at" api:"nullable" format:"date-time"`
	// When the pipeline file was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		PipelineID       respjson.Field
		ConfigHash       respjson.Field
		CreatedAt        respjson.Field
		CustomMetadata   respjson.Field
		DataSourceID     respjson.Field
		ExternalFileID   respjson.Field
		FileID           respjson.Field
		FileSize         respjson.Field
		FileType         respjson.Field
		IndexedPageCount respjson.Field
		LastModifiedAt   respjson.Field
		Name             respjson.Field
		PermissionInfo   respjson.Field
		ProjectID        respjson.Field
		ResourceInfo     respjson.Field
		Status           respjson.Field
		StatusUpdatedAt  respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineFile) RawJSON() string { return r.JSON.raw }
func (r *PipelineFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineFileConfigHashUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineFileConfigHashMapItem OfAnyArray OfString OfFloat
// OfBool]
type PipelineFileConfigHashUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineFileConfigHashMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPipelineFileConfigHashMapItem respjson.Field
		OfAnyArray                      respjson.Field
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		raw                             string
	} `json:"-"`
}

func (u PipelineFileConfigHashUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileConfigHashUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileConfigHashUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileConfigHashUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileConfigHashUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineFileConfigHashUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineFileConfigHashUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineFileCustomMetadataUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineFileCustomMetadataMapItem OfAnyArray OfString OfFloat
// OfBool]
type PipelineFileCustomMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineFileCustomMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPipelineFileCustomMetadataMapItem respjson.Field
		OfAnyArray                          respjson.Field
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u PipelineFileCustomMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileCustomMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileCustomMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileCustomMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileCustomMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineFileCustomMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineFileCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineFilePermissionInfoUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineFilePermissionInfoMapItem OfAnyArray OfString OfFloat
// OfBool]
type PipelineFilePermissionInfoUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineFilePermissionInfoMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPipelineFilePermissionInfoMapItem respjson.Field
		OfAnyArray                          respjson.Field
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u PipelineFilePermissionInfoUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFilePermissionInfoUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFilePermissionInfoUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFilePermissionInfoUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFilePermissionInfoUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineFilePermissionInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineFilePermissionInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineFileResourceInfoUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineFileResourceInfoMapItem OfAnyArray OfString OfFloat
// OfBool]
type PipelineFileResourceInfoUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineFileResourceInfoMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPipelineFileResourceInfoMapItem respjson.Field
		OfAnyArray                        respjson.Field
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfBool                            respjson.Field
		raw                               string
	} `json:"-"`
}

func (u PipelineFileResourceInfoUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileResourceInfoUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileResourceInfoUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileResourceInfoUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineFileResourceInfoUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineFileResourceInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineFileResourceInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the pipeline file.
type PipelineFileStatus string

const (
	PipelineFileStatusCancelled  PipelineFileStatus = "CANCELLED"
	PipelineFileStatusError      PipelineFileStatus = "ERROR"
	PipelineFileStatusInProgress PipelineFileStatus = "IN_PROGRESS"
	PipelineFileStatusNotStarted PipelineFileStatus = "NOT_STARTED"
	PipelineFileStatusSuccess    PipelineFileStatus = "SUCCESS"
)

type PipelineFileGetStatusCountsResponse struct {
	// The counts of files by status
	Counts map[string]int64 `json:"counts" api:"required"`
	// The total number of files
	TotalCount int64 `json:"total_count" api:"required"`
	// The ID of the data source that the files belong to
	DataSourceID string `json:"data_source_id" api:"nullable" format:"uuid"`
	// Whether to only count manually uploaded files
	OnlyManuallyUploaded bool `json:"only_manually_uploaded"`
	// The ID of the pipeline that the files belong to
	PipelineID string `json:"pipeline_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counts               respjson.Field
		TotalCount           respjson.Field
		DataSourceID         respjson.Field
		OnlyManuallyUploaded respjson.Field
		PipelineID           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineFileGetStatusCountsResponse) RawJSON() string { return r.JSON.raw }
func (r *PipelineFileGetStatusCountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineFileNewParams struct {
	Body []PipelineFileNewParamsBody
	paramObj
}

func (r PipelineFileNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PipelineFileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating a file that is associated with a pipeline.
//
// The property FileID is required.
type PipelineFileNewParamsBody struct {
	// The ID of the file
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// Custom metadata for the file
	CustomMetadata map[string]*PipelineFileNewParamsBodyCustomMetadataUnion `json:"custom_metadata,omitzero"`
	paramObj
}

func (r PipelineFileNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PipelineFileNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineFileNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineFileNewParamsBodyCustomMetadataUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineFileNewParamsBodyCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *PipelineFileNewParamsBodyCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PipelineFileUpdateParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	// Custom metadata for the file
	CustomMetadata map[string]*PipelineFileUpdateParamsCustomMetadataUnion `json:"custom_metadata,omitzero"`
	paramObj
}

func (r PipelineFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PipelineFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PipelineFileUpdateParamsCustomMetadataUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u PipelineFileUpdateParamsCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *PipelineFileUpdateParamsCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type PipelineFileListParams struct {
	DataSourceID         param.Opt[string] `query:"data_source_id,omitzero" format:"uuid" json:"-"`
	FileNameContains     param.Opt[string] `query:"file_name_contains,omitzero" json:"-"`
	Limit                param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset               param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy              param.Opt[string] `query:"order_by,omitzero" json:"-"`
	OnlyManuallyUploaded param.Opt[bool]   `query:"only_manually_uploaded,omitzero" json:"-"`
	// Filter by file statuses
	//
	// Any of "CANCELLED", "ERROR", "IN_PROGRESS", "NOT_STARTED", "SUCCESS".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineFileListParams]'s query parameters as `url.Values`.
func (r PipelineFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineFileDeleteParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineFileGetStatusParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineFileGetStatusCountsParams struct {
	DataSourceID         param.Opt[string] `query:"data_source_id,omitzero" format:"uuid" json:"-"`
	OnlyManuallyUploaded param.Opt[bool]   `query:"only_manually_uploaded,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineFileGetStatusCountsParams]'s query parameters as
// `url.Values`.
func (r PipelineFileGetStatusCountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
