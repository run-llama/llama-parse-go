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

	"github.com/run-llama/llama-parse-go/internal/apiform"
	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// BetaDirectoryFileService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaDirectoryFileService] method instead.
type BetaDirectoryFileService struct {
	options []option.RequestOption
}

// NewBetaDirectoryFileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaDirectoryFileService(opts ...option.RequestOption) (r BetaDirectoryFileService) {
	r = BetaDirectoryFileService{}
	r.options = opts
	return
}

// Update directory-file metadata by `directory_file_id`; set `directory_id` to
// move the file to a different directory. To resolve from `unique_id`, list with a
// filter first.
func (r *BetaDirectoryFileService) Update(ctx context.Context, directoryFileID string, params BetaDirectoryFileUpdateParams, opts ...option.RequestOption) (res *BetaDirectoryFileUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.DirectoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	if directoryFileID == "" {
		err = errors.New("missing required directory_file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files/%s", url.PathEscape(params.DirectoryID), url.PathEscape(directoryFileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List all files within the specified directory with optional filtering and
// pagination.
func (r *BetaDirectoryFileService) List(ctx context.Context, directoryID string, query BetaDirectoryFileListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BetaDirectoryFileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files", url.PathEscape(directoryID))
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

// List all files within the specified directory with optional filtering and
// pagination.
func (r *BetaDirectoryFileService) ListAutoPaging(ctx context.Context, directoryID string, query BetaDirectoryFileListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BetaDirectoryFileListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, directoryID, query, opts...))
}

// Delete a directory file by `directory_file_id`; to resolve from `unique_id`,
// list with a filter first.
func (r *BetaDirectoryFileService) Delete(ctx context.Context, directoryFileID string, params BetaDirectoryFileDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.DirectoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return err
	}
	if directoryFileID == "" {
		err = errors.New("missing required directory_file_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files/%s", url.PathEscape(params.DirectoryID), url.PathEscape(directoryFileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Create a new file within the specified directory; the directory must exist in
// the project and `file_id` must reference an existing file.
func (r *BetaDirectoryFileService) Add(ctx context.Context, directoryID string, params BetaDirectoryFileAddParams, opts ...option.RequestOption) (res *BetaDirectoryFileAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files", url.PathEscape(directoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a directory file by `directory_file_id`; to look up by `unique_id`, use the
// list endpoint with a filter.
func (r *BetaDirectoryFileService) Get(ctx context.Context, directoryFileID string, params BetaDirectoryFileGetParams, opts ...option.RequestOption) (res *BetaDirectoryFileGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.DirectoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	if directoryFileID == "" {
		err = errors.New("missing required directory_file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files/%s", url.PathEscape(params.DirectoryID), url.PathEscape(directoryFileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Upload a file and create its directory entry in one call; `unique_id` /
// `display_name` default to values derived from file metadata.
func (r *BetaDirectoryFileService) Upload(ctx context.Context, directoryID string, params BetaDirectoryFileUploadParams, opts ...option.RequestOption) (res *BetaDirectoryFileUploadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s/files/upload", url.PathEscape(directoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// API response schema for a directory file.
type BetaDirectoryFileUpdateResponse struct {
	// Unique identifier for the directory file.
	ID string `json:"id" api:"required"`
	// Directory the file belongs to.
	DirectoryID string `json:"directory_id" api:"required"`
	// Display name for the file.
	DisplayName string `json:"display_name" api:"required"`
	// Project the directory file belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Unique identifier for the file in the directory
	UniqueID string `json:"unique_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Soft delete marker when the file is removed upstream or by user action.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// File ID for the storage location.
	FileID string `json:"file_id" api:"nullable"`
	// Merged metadata from all sources. Higher-priority sources override lower.
	Metadata map[string]BetaDirectoryFileUpdateResponseMetadataUnion `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DirectoryID respjson.Field
		DisplayName respjson.Field
		ProjectID   respjson.Field
		UniqueID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		DownloadURL respjson.Field
		FileID      respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryFileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryFileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDirectoryFileUpdateResponseMetadataUnion contains all possible properties
// and values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type BetaDirectoryFileUpdateResponseMetadataUnion struct {
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

func (u BetaDirectoryFileUpdateResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUpdateResponseMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUpdateResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUpdateResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUpdateResponseMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDirectoryFileUpdateResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDirectoryFileUpdateResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory file.
type BetaDirectoryFileListResponse struct {
	// Unique identifier for the directory file.
	ID string `json:"id" api:"required"`
	// Directory the file belongs to.
	DirectoryID string `json:"directory_id" api:"required"`
	// Display name for the file.
	DisplayName string `json:"display_name" api:"required"`
	// Project the directory file belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Unique identifier for the file in the directory
	UniqueID string `json:"unique_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Soft delete marker when the file is removed upstream or by user action.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// File ID for the storage location.
	FileID string `json:"file_id" api:"nullable"`
	// Merged metadata from all sources. Higher-priority sources override lower.
	Metadata map[string]BetaDirectoryFileListResponseMetadataUnion `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DirectoryID respjson.Field
		DisplayName respjson.Field
		ProjectID   respjson.Field
		UniqueID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		DownloadURL respjson.Field
		FileID      respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDirectoryFileListResponseMetadataUnion contains all possible properties and
// values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type BetaDirectoryFileListResponseMetadataUnion struct {
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

func (u BetaDirectoryFileListResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileListResponseMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileListResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileListResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileListResponseMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDirectoryFileListResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDirectoryFileListResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory file.
type BetaDirectoryFileAddResponse struct {
	// Unique identifier for the directory file.
	ID string `json:"id" api:"required"`
	// Directory the file belongs to.
	DirectoryID string `json:"directory_id" api:"required"`
	// Display name for the file.
	DisplayName string `json:"display_name" api:"required"`
	// Project the directory file belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Unique identifier for the file in the directory
	UniqueID string `json:"unique_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Soft delete marker when the file is removed upstream or by user action.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// File ID for the storage location.
	FileID string `json:"file_id" api:"nullable"`
	// Merged metadata from all sources. Higher-priority sources override lower.
	Metadata map[string]BetaDirectoryFileAddResponseMetadataUnion `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DirectoryID respjson.Field
		DisplayName respjson.Field
		ProjectID   respjson.Field
		UniqueID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		DownloadURL respjson.Field
		FileID      respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryFileAddResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryFileAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDirectoryFileAddResponseMetadataUnion contains all possible properties and
// values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type BetaDirectoryFileAddResponseMetadataUnion struct {
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

func (u BetaDirectoryFileAddResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileAddResponseMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileAddResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileAddResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileAddResponseMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDirectoryFileAddResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDirectoryFileAddResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory file.
type BetaDirectoryFileGetResponse struct {
	// Unique identifier for the directory file.
	ID string `json:"id" api:"required"`
	// Directory the file belongs to.
	DirectoryID string `json:"directory_id" api:"required"`
	// Display name for the file.
	DisplayName string `json:"display_name" api:"required"`
	// Project the directory file belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Unique identifier for the file in the directory
	UniqueID string `json:"unique_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Soft delete marker when the file is removed upstream or by user action.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// File ID for the storage location.
	FileID string `json:"file_id" api:"nullable"`
	// Merged metadata from all sources. Higher-priority sources override lower.
	Metadata map[string]BetaDirectoryFileGetResponseMetadataUnion `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DirectoryID respjson.Field
		DisplayName respjson.Field
		ProjectID   respjson.Field
		UniqueID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		DownloadURL respjson.Field
		FileID      respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryFileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryFileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDirectoryFileGetResponseMetadataUnion contains all possible properties and
// values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type BetaDirectoryFileGetResponseMetadataUnion struct {
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

func (u BetaDirectoryFileGetResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileGetResponseMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileGetResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileGetResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileGetResponseMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDirectoryFileGetResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDirectoryFileGetResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory file.
type BetaDirectoryFileUploadResponse struct {
	// Unique identifier for the directory file.
	ID string `json:"id" api:"required"`
	// Directory the file belongs to.
	DirectoryID string `json:"directory_id" api:"required"`
	// Display name for the file.
	DisplayName string `json:"display_name" api:"required"`
	// Project the directory file belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Unique identifier for the file in the directory
	UniqueID string `json:"unique_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Soft delete marker when the file is removed upstream or by user action.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Schema for a presigned URL.
	DownloadURL PresignedURL `json:"download_url" api:"nullable"`
	// File ID for the storage location.
	FileID string `json:"file_id" api:"nullable"`
	// Merged metadata from all sources. Higher-priority sources override lower.
	Metadata map[string]BetaDirectoryFileUploadResponseMetadataUnion `json:"metadata"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DirectoryID respjson.Field
		DisplayName respjson.Field
		ProjectID   respjson.Field
		UniqueID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		DownloadURL respjson.Field
		FileID      respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryFileUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryFileUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaDirectoryFileUploadResponseMetadataUnion contains all possible properties
// and values from [string], [int64], [float64], [bool], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfFloat OfBool OfStringArray]
type BetaDirectoryFileUploadResponseMetadataUnion struct {
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

func (u BetaDirectoryFileUploadResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUploadResponseMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUploadResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUploadResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaDirectoryFileUploadResponseMetadataUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaDirectoryFileUploadResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaDirectoryFileUploadResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaDirectoryFileUpdateParams struct {
	DirectoryID    string            `path:"directory_id" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Updated display name.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Move file to a different directory.
	TargetDirectoryID param.Opt[string] `json:"target_directory_id,omitzero"`
	// Updated unique identifier.
	UniqueID param.Opt[string] `json:"unique_id,omitzero"`
	// User-defined metadata key-value pairs. Replaces the user metadata layer.
	Metadata map[string]BetaDirectoryFileUpdateParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r BetaDirectoryFileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDirectoryFileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDirectoryFileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaDirectoryFileUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaDirectoryFileUpdateParamsMetadataUnion struct {
	OfString      param.Opt[string]  `json:",omitzero,inline"`
	OfInt         param.Opt[int64]   `json:",omitzero,inline"`
	OfFloat       param.Opt[float64] `json:",omitzero,inline"`
	OfBool        param.Opt[bool]    `json:",omitzero,inline"`
	OfStringArray []string           `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDirectoryFileUpdateParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfInt,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray)
}
func (u *BetaDirectoryFileUpdateParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BetaDirectoryFileListParams struct {
	DisplayName         param.Opt[string] `query:"display_name,omitzero" json:"-"`
	DisplayNameContains param.Opt[string] `query:"display_name_contains,omitzero" json:"-"`
	FileID              param.Opt[string] `query:"file_id,omitzero" json:"-"`
	OrganizationID      param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize            param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken           param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID           param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	UniqueID            param.Opt[string] `query:"unique_id,omitzero" json:"-"`
	// Include items updated at or after this timestamp (inclusive)
	UpdatedAtOnOrAfter param.Opt[time.Time] `query:"updated_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items updated at or before this timestamp (inclusive)
	UpdatedAtOnOrBefore param.Opt[time.Time] `query:"updated_at_on_or_before,omitzero" format:"date-time" json:"-"`
	IncludeDeleted      param.Opt[bool]      `query:"include_deleted,omitzero" json:"-"`
	// Fields to expand on each directory file.
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryFileListParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryFileDeleteParams struct {
	DirectoryID    string            `path:"directory_id" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryFileDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryFileAddParams struct {
	// File ID for the storage location (required).
	FileID         string            `json:"file_id" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Display name for the file. If not provided, will use the file's name.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Unique identifier for the file in the directory. If not provided, will use the
	// file's external_file_id or name.
	UniqueID param.Opt[string] `json:"unique_id,omitzero"`
	// User-defined metadata key-value pairs to associate with the file.
	Metadata map[string]BetaDirectoryFileAddParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r BetaDirectoryFileAddParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDirectoryFileAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDirectoryFileAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaDirectoryFileAddParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileAddParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaDirectoryFileAddParamsMetadataUnion struct {
	OfString      param.Opt[string]  `json:",omitzero,inline"`
	OfInt         param.Opt[int64]   `json:",omitzero,inline"`
	OfFloat       param.Opt[float64] `json:",omitzero,inline"`
	OfBool        param.Opt[bool]    `json:",omitzero,inline"`
	OfStringArray []string           `json:",omitzero,inline"`
	paramUnion
}

func (u BetaDirectoryFileAddParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfInt,
		u.OfFloat,
		u.OfBool,
		u.OfStringArray)
}
func (u *BetaDirectoryFileAddParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BetaDirectoryFileGetParams struct {
	DirectoryID    string            `path:"directory_id" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields to expand.
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryFileGetParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryFileUploadParams struct {
	UploadFile     io.Reader         `json:"upload_file,omitzero" api:"required" format:"binary"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	DisplayName    param.Opt[string] `json:"display_name,omitzero"`
	ExternalFileID param.Opt[string] `json:"external_file_id,omitzero"`
	// User metadata as a JSON object string.
	Metadata param.Opt[string] `json:"metadata,omitzero"`
	UniqueID param.Opt[string] `json:"unique_id,omitzero"`
	paramObj
}

func (r BetaDirectoryFileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [BetaDirectoryFileUploadParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryFileUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
