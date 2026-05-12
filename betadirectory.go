// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// BetaDirectoryService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaDirectoryService] method instead.
type BetaDirectoryService struct {
	options []option.RequestOption
	Files   BetaDirectoryFileService
}

// NewBetaDirectoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaDirectoryService(opts ...option.RequestOption) (r BetaDirectoryService) {
	r = BetaDirectoryService{}
	r.options = opts
	r.Files = NewBetaDirectoryFileService(opts...)
	return
}

// Create a new directory within the specified project.
func (r *BetaDirectoryService) New(ctx context.Context, params BetaDirectoryNewParams, opts ...option.RequestOption) (res *BetaDirectoryNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/directories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update directory metadata.
func (r *BetaDirectoryService) Update(ctx context.Context, directoryID string, params BetaDirectoryUpdateParams, opts ...option.RequestOption) (res *BetaDirectoryUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s", url.PathEscape(directoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List Directories
func (r *BetaDirectoryService) List(ctx context.Context, query BetaDirectoryListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BetaDirectoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/directories"
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

// List Directories
func (r *BetaDirectoryService) ListAutoPaging(ctx context.Context, query BetaDirectoryListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BetaDirectoryListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete a directory.
func (r *BetaDirectoryService) Delete(ctx context.Context, directoryID string, body BetaDirectoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s", url.PathEscape(directoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieve a directory by its identifier.
func (r *BetaDirectoryService) Get(ctx context.Context, directoryID string, query BetaDirectoryGetParams, opts ...option.RequestOption) (res *BetaDirectoryGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/directories/%s", url.PathEscape(directoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// API response schema for a directory.
type BetaDirectoryNewResponse struct {
	// Unique identifier for the directory.
	ID string `json:"id" api:"required"`
	// Human-readable name for the directory.
	Name string `json:"name" api:"required"`
	// Project the directory belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Optional timestamp of when the directory was deleted. Null if not deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Optional description shown to users.
	Description string `json:"description" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory.
type BetaDirectoryUpdateResponse struct {
	// Unique identifier for the directory.
	ID string `json:"id" api:"required"`
	// Human-readable name for the directory.
	Name string `json:"name" api:"required"`
	// Project the directory belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Optional timestamp of when the directory was deleted. Null if not deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Optional description shown to users.
	Description string `json:"description" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory.
type BetaDirectoryListResponse struct {
	// Unique identifier for the directory.
	ID string `json:"id" api:"required"`
	// Human-readable name for the directory.
	Name string `json:"name" api:"required"`
	// Project the directory belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Optional timestamp of when the directory was deleted. Null if not deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Optional description shown to users.
	Description string `json:"description" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a directory.
type BetaDirectoryGetResponse struct {
	// Unique identifier for the directory.
	ID string `json:"id" api:"required"`
	// Human-readable name for the directory.
	Name string `json:"name" api:"required"`
	// Project the directory belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Optional timestamp of when the directory was deleted. Null if not deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Optional description shown to users.
	Description string `json:"description" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaDirectoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaDirectoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaDirectoryNewParams struct {
	// Human-readable name for the directory.
	Name           string            `json:"name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Optional description shown to users.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BetaDirectoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDirectoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDirectoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaDirectoryNewParams]'s query parameters as `url.Values`.
func (r BetaDirectoryNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryUpdateParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Updated description for the directory.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated name for the directory.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BetaDirectoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaDirectoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaDirectoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaDirectoryUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryListParams struct {
	Name           param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	IncludeDeleted param.Opt[bool]   `query:"include_deleted,omitzero" json:"-"`
	// Any of "user", "index".
	Type BetaDirectoryListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryListParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryListParamsType string

const (
	BetaDirectoryListParamsTypeUser  BetaDirectoryListParamsType = "user"
	BetaDirectoryListParamsTypeIndex BetaDirectoryListParamsType = "index"
)

type BetaDirectoryDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaDirectoryDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaDirectoryGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaDirectoryGetParams]'s query parameters as `url.Values`.
func (r BetaDirectoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
