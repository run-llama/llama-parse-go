// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiform"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// FileService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	return
}

// Upload a file using multipart/form-data.
//
// Set `purpose` to indicate how the file will be used: `user_data`, `parse`,
// `extract`, `classify`, `split`, `sheet`, or `agent_app`.
//
// Returns the created file metadata including its ID for use in subsequent parse,
// extract, or classify operations.
func (r *FileService) New(ctx context.Context, params FileNewParams, opts ...option.RequestOption) (res *FileNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List files with optional filtering and pagination.
//
// Filter by `file_name`, `file_ids`, or `external_file_id`. Supports cursor-based
// pagination and custom ordering.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[FileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/files"
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

// List files with optional filtering and pagination.
//
// Filter by `file_name`, `file_ids`, or `external_file_id`. Supports cursor-based
// pagination and custom ordering.
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[FileListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a file from the project.
func (r *FileService) Delete(ctx context.Context, fileID string, body FileDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Get a presigned URL to download the file content.
func (r *FileService) Get(ctx context.Context, fileID string, query FileGetParams, opts ...option.RequestOption) (res *PresignedURL, err error) {
	opts = slices.Concat(r.options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/files/%s/content", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Query files with flexible filtering and pagination.
//
// **Deprecated**: Use GET /files instead for listing files with query parameters.
//
// Args: request: The query request with filters and pagination project: Validated
// project from dependency
//
// Returns: Paginated response with files
//
// Deprecated: Use the GET /files endpoint instead
func (r *FileService) Query(ctx context.Context, params FileQueryParams, opts ...option.RequestOption) (res *FileQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/files/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Schema for a file.
type File struct {
	// Unique identifier
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// The ID of the project that the file belongs to
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The ID of the data source that the file belongs to
	DataSourceID string `json:"data_source_id" api:"nullable" format:"uuid"`
	// The expiration date for the file. Files past this date can be deleted.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// The ID of the file in the external system
	ExternalFileID string `json:"external_file_id" api:"nullable"`
	// Size of the file in bytes
	FileSize int64 `json:"file_size" api:"nullable"`
	// File type (e.g. pdf, docx, etc.)
	FileType string `json:"file_type" api:"nullable"`
	// The last modified time of the file
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Permission information for the file
	PermissionInfo map[string]*FilePermissionInfoUnion `json:"permission_info" api:"nullable"`
	// The intended purpose of the file (e.g., 'user_data', 'parse', 'extract',
	// 'split', 'classify')
	Purpose string `json:"purpose" api:"nullable"`
	// Resource information for the file
	ResourceInfo map[string]*FileResourceInfoUnion `json:"resource_info" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ProjectID      respjson.Field
		CreatedAt      respjson.Field
		DataSourceID   respjson.Field
		ExpiresAt      respjson.Field
		ExternalFileID respjson.Field
		FileSize       respjson.Field
		FileType       respjson.Field
		LastModifiedAt respjson.Field
		PermissionInfo respjson.Field
		Purpose        respjson.Field
		ResourceInfo   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FilePermissionInfoUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFilePermissionInfoMapItem OfAnyArray OfString OfFloat OfBool]
type FilePermissionInfoUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfFilePermissionInfoMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfFilePermissionInfoMapItem respjson.Field
		OfAnyArray                  respjson.Field
		OfString                    respjson.Field
		OfFloat                     respjson.Field
		OfBool                      respjson.Field
		raw                         string
	} `json:"-"`
}

func (u FilePermissionInfoUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilePermissionInfoUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilePermissionInfoUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilePermissionInfoUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilePermissionInfoUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FilePermissionInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *FilePermissionInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FileResourceInfoUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFileResourceInfoMapItem OfAnyArray OfString OfFloat OfBool]
type FileResourceInfoUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfFileResourceInfoMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfFileResourceInfoMapItem respjson.Field
		OfAnyArray                respjson.Field
		OfString                  respjson.Field
		OfFloat                   respjson.Field
		OfBool                    respjson.Field
		raw                       string
	} `json:"-"`
}

func (u FileResourceInfoUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileResourceInfoUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileResourceInfoUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileResourceInfoUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileResourceInfoUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileResourceInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *FileResourceInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for a presigned URL.
type PresignedURL struct {
	// The time at which the presigned URL expires
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// A presigned URL for IO operations against a private file
	URL string `json:"url" api:"required" format:"uri"`
	// Form fields for a presigned POST request
	FormFields map[string]string `json:"form_fields" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		URL         respjson.Field
		FormFields  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PresignedURL) RawJSON() string { return r.JSON.raw }
func (r *PresignedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An uploaded file.
type FileNewResponse struct {
	// Unique file identifier
	ID string `json:"id" api:"required"`
	// File name including extension
	Name string `json:"name" api:"required"`
	// Project this file belongs to
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// When the file expires and may be automatically removed. Null means no
	// expiration.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Optional ID for correlating with an external system
	ExternalFileID string `json:"external_file_id" api:"nullable"`
	// File extension (pdf, docx, png, etc.)
	FileType string `json:"file_type" api:"nullable"`
	// When the file was last modified (ISO 8601)
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// How the file will be used: user_data, parse, extract, classify, split, sheet, or
	// agent_app
	Purpose string `json:"purpose" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ProjectID      respjson.Field
		DownloadURL    respjson.Field
		ExpiresAt      respjson.Field
		ExternalFileID respjson.Field
		FileType       respjson.Field
		LastModifiedAt respjson.Field
		Purpose        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An uploaded file.
type FileListResponse struct {
	// Unique file identifier
	ID string `json:"id" api:"required"`
	// File name including extension
	Name string `json:"name" api:"required"`
	// Project this file belongs to
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// When the file expires and may be automatically removed. Null means no
	// expiration.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Optional ID for correlating with an external system
	ExternalFileID string `json:"external_file_id" api:"nullable"`
	// File extension (pdf, docx, png, etc.)
	FileType string `json:"file_type" api:"nullable"`
	// When the file was last modified (ISO 8601)
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// How the file will be used: user_data, parse, extract, classify, split, sheet, or
	// agent_app
	Purpose string `json:"purpose" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ProjectID      respjson.Field
		DownloadURL    respjson.Field
		ExpiresAt      respjson.Field
		ExternalFileID respjson.Field
		FileType       respjson.Field
		LastModifiedAt respjson.Field
		Purpose        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of files.
type FileQueryResponse struct {
	// The list of items.
	Items []FileQueryResponseItem `json:"items" api:"required"`
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
func (r FileQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *FileQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An uploaded file.
type FileQueryResponseItem struct {
	// Unique file identifier
	ID string `json:"id" api:"required"`
	// File name including extension
	Name string `json:"name" api:"required"`
	// Project this file belongs to
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// When the file expires and may be automatically removed. Null means no
	// expiration.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Optional ID for correlating with an external system
	ExternalFileID string `json:"external_file_id" api:"nullable"`
	// File extension (pdf, docx, png, etc.)
	FileType string `json:"file_type" api:"nullable"`
	// When the file was last modified (ISO 8601)
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// How the file will be used: user_data, parse, extract, classify, split, sheet, or
	// agent_app
	Purpose string `json:"purpose" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ProjectID      respjson.Field
		DownloadURL    respjson.Field
		ExpiresAt      respjson.Field
		ExternalFileID respjson.Field
		FileType       respjson.Field
		LastModifiedAt respjson.Field
		Purpose        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileQueryResponseItem) RawJSON() string { return r.JSON.raw }
func (r *FileQueryResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileNewParams struct {
	// The file to upload
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// The intended purpose of the file. Valid values: 'user_data', 'parse', 'extract',
	// 'split', 'classify', 'sheet', 'agent_app'. This determines the storage and
	// retention policy for the file.
	Purpose        string            `json:"purpose" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// The ID of the file in the external system
	ExternalFileID param.Opt[string] `json:"external_file_id,omitzero"`
	paramObj
}

func (r FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [FileNewParams]'s query parameters as `url.Values`.
func (r FileNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileListParams struct {
	// Filter by external file ID.
	ExternalFileID param.Opt[string] `query:"external_file_id,omitzero" json:"-"`
	// Filter by file name (exact match).
	FileName param.Opt[string] `query:"file_name,omitzero" json:"-"`
	// A comma-separated list of fields to order by, sorted in ascending order. Use
	// 'field_name desc' to specify descending order.
	OrderBy        param.Opt[string] `query:"order_by,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// The maximum number of items to return. Defaults to 50, maximum is 1000.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// A page token received from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields to expand on each file.
	Expand []string `query:"expand,omitzero" json:"-"`
	// Filter by specific file IDs.
	FileIDs []string `query:"file_ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [FileDeleteParams]'s query parameters as `url.Values`.
func (r FileDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileGetParams struct {
	ExpiresAtSeconds param.Opt[int64]  `query:"expires_at_seconds,omitzero" json:"-"`
	OrganizationID   param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID        param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [FileGetParams]'s query parameters as `url.Values`.
func (r FileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileQueryParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// A comma-separated list of fields to order by, sorted in ascending order. Use
	// 'field_name desc' to specify descending order.
	OrderBy param.Opt[string] `json:"order_by,omitzero"`
	// The maximum number of items to return. The service may return fewer than this
	// value. If unspecified, a default page size will be used. The maximum value is
	// typically 1000; values above this will be coerced to the maximum.
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// A page token, received from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Filter parameters for file queries.
	Filter FileQueryParamsFilter `json:"filter,omitzero"`
	paramObj
}

func (r FileQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow FileQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FileQueryParams]'s query parameters as `url.Values`.
func (r FileQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameters for file queries.
type FileQueryParamsFilter struct {
	// Filter by data source ID
	DataSourceID param.Opt[string] `json:"data_source_id,omitzero" format:"uuid"`
	// Filter by external file ID
	ExternalFileID param.Opt[string] `json:"external_file_id,omitzero"`
	// Filter by file name
	FileName param.Opt[string] `json:"file_name,omitzero"`
	// Filter only manually uploaded files (data_source_id is null)
	OnlyManuallyUploaded param.Opt[bool] `json:"only_manually_uploaded,omitzero"`
	// Filter by project ID
	ProjectID param.Opt[string] `json:"project_id,omitzero" format:"uuid"`
	// Filter by specific file IDs
	FileIDs []string `json:"file_ids,omitzero" format:"uuid"`
	paramObj
}

func (r FileQueryParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow FileQueryParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileQueryParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
