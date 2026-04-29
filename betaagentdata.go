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

// BetaAgentDataService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaAgentDataService] method instead.
type BetaAgentDataService struct {
	options []option.RequestOption
}

// NewBetaAgentDataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaAgentDataService(opts ...option.RequestOption) (r BetaAgentDataService) {
	r = BetaAgentDataService{}
	r.options = opts
	return
}

// Create new agent data.
func (r *BetaAgentDataService) New(ctx context.Context, params BetaAgentDataNewParams, opts ...option.RequestOption) (res *AgentData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/agent-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update agent data by ID (overwrites).
func (r *BetaAgentDataService) Update(ctx context.Context, itemID string, params BetaAgentDataUpdateParams, opts ...option.RequestOption) (res *AgentData, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/agent-data/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Delete agent data by ID.
func (r *BetaAgentDataService) Delete(ctx context.Context, itemID string, body BetaAgentDataDeleteParams, opts ...option.RequestOption) (res *BetaAgentDataDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/agent-data/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Aggregate agent data with grouping and optional counting/first item retrieval.
func (r *BetaAgentDataService) Aggregate(ctx context.Context, params BetaAgentDataAggregateParams, opts ...option.RequestOption) (res *pagination.PaginatedCursorPost[BetaAgentDataAggregateResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/agent-data/:aggregate"
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

// Aggregate agent data with grouping and optional counting/first item retrieval.
func (r *BetaAgentDataService) AggregateAutoPaging(ctx context.Context, params BetaAgentDataAggregateParams, opts ...option.RequestOption) *pagination.PaginatedCursorPostAutoPager[BetaAgentDataAggregateResponse] {
	return pagination.NewPaginatedCursorPostAutoPager(r.Aggregate(ctx, params, opts...))
}

// Bulk delete agent data by query (deployment_name, collection, optional filters).
func (r *BetaAgentDataService) DeleteByQuery(ctx context.Context, params BetaAgentDataDeleteByQueryParams, opts ...option.RequestOption) (res *BetaAgentDataDeleteByQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/agent-data/:delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get agent data by ID.
func (r *BetaAgentDataService) Get(ctx context.Context, itemID string, query BetaAgentDataGetParams, opts ...option.RequestOption) (res *AgentData, err error) {
	opts = slices.Concat(r.options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/agent-data/%s", url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search agent data with filtering, sorting, and pagination.
func (r *BetaAgentDataService) Search(ctx context.Context, params BetaAgentDataSearchParams, opts ...option.RequestOption) (res *pagination.PaginatedCursorPost[AgentData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/agent-data/:search"
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

// Search agent data with filtering, sorting, and pagination.
func (r *BetaAgentDataService) SearchAutoPaging(ctx context.Context, params BetaAgentDataSearchParams, opts ...option.RequestOption) *pagination.PaginatedCursorPostAutoPager[AgentData] {
	return pagination.NewPaginatedCursorPostAutoPager(r.Search(ctx, params, opts...))
}

// API Result for a single agent data item
type AgentData struct {
	Data           map[string]any `json:"data" api:"required"`
	DeploymentName string         `json:"deployment_name" api:"required"`
	ID             string         `json:"id" api:"nullable"`
	Collection     string         `json:"collection"`
	CreatedAt      time.Time      `json:"created_at" api:"nullable" format:"date-time"`
	ProjectID      string         `json:"project_id" api:"nullable"`
	UpdatedAt      time.Time      `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data           respjson.Field
		DeploymentName respjson.Field
		ID             respjson.Field
		Collection     respjson.Field
		CreatedAt      respjson.Field
		ProjectID      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentData) RawJSON() string { return r.JSON.raw }
func (r *AgentData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaAgentDataDeleteResponse map[string]string

// API Result for a single group in the aggregate response
type BetaAgentDataAggregateResponse struct {
	GroupKey  map[string]any `json:"group_key" api:"required"`
	Count     int64          `json:"count" api:"nullable"`
	FirstItem map[string]any `json:"first_item" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupKey    respjson.Field
		Count       respjson.Field
		FirstItem   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaAgentDataAggregateResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaAgentDataAggregateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response for bulk delete operation
type BetaAgentDataDeleteByQueryResponse struct {
	DeletedCount int64 `json:"deleted_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletedCount respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaAgentDataDeleteByQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaAgentDataDeleteByQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaAgentDataNewParams struct {
	Data           map[string]any    `json:"data,omitzero" api:"required"`
	DeploymentName string            `json:"deployment_name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	Collection     param.Opt[string] `json:"collection,omitzero"`
	paramObj
}

func (r BetaAgentDataNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaAgentDataNewParams]'s query parameters as `url.Values`.
func (r BetaAgentDataNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaAgentDataUpdateParams struct {
	Data           map[string]any    `json:"data,omitzero" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BetaAgentDataUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaAgentDataUpdateParams]'s query parameters as
// `url.Values`.
func (r BetaAgentDataUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaAgentDataDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaAgentDataDeleteParams]'s query parameters as
// `url.Values`.
func (r BetaAgentDataDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaAgentDataAggregateParams struct {
	// The agent deployment's name to aggregate data for
	DeploymentName string            `json:"deployment_name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Whether to count the number of items in each group
	Count param.Opt[bool] `json:"count,omitzero"`
	// Whether to return the first item in each group (Sorted by created_at)
	First param.Opt[bool] `json:"first,omitzero"`
	// The offset to start from. If not provided, the first page is returned
	Offset param.Opt[int64] `json:"offset,omitzero"`
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
	// The logical agent data collection to aggregate data for
	Collection param.Opt[string] `json:"collection,omitzero"`
	// A filter object or expression that filters resources listed in the response.
	Filter map[string]BetaAgentDataAggregateParamsFilter `json:"filter,omitzero"`
	// The fields to group by. If empty, the entire dataset is grouped on. e.g. if left
	// out, can be used for simple count operations
	GroupBy []string `json:"group_by,omitzero"`
	paramObj
}

func (r BetaAgentDataAggregateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataAggregateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataAggregateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaAgentDataAggregateParams]'s query parameters as
// `url.Values`.
func (r BetaAgentDataAggregateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// API request model for a filter comparison operation.
type BetaAgentDataAggregateParamsFilter struct {
	Eq       BetaAgentDataAggregateParamsFilterEqUnion         `json:"eq,omitzero" format:"date-time"`
	Gt       BetaAgentDataAggregateParamsFilterGtUnion         `json:"gt,omitzero" format:"date-time"`
	Gte      BetaAgentDataAggregateParamsFilterGteUnion        `json:"gte,omitzero" format:"date-time"`
	Lt       BetaAgentDataAggregateParamsFilterLtUnion         `json:"lt,omitzero" format:"date-time"`
	Lte      BetaAgentDataAggregateParamsFilterLteUnion        `json:"lte,omitzero" format:"date-time"`
	Ne       BetaAgentDataAggregateParamsFilterNeUnion         `json:"ne,omitzero" format:"date-time"`
	Excludes []*BetaAgentDataAggregateParamsFilterExcludeUnion `json:"excludes,omitzero" format:"date-time"`
	Includes []*BetaAgentDataAggregateParamsFilterIncludeUnion `json:"includes,omitzero" format:"date-time"`
	paramObj
}

func (r BetaAgentDataAggregateParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataAggregateParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataAggregateParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterEqUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterEqUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterEqUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterExcludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterGtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterGtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterGtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterGteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterGteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterGteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterIncludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterIncludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterIncludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterLtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterLtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterLtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterLteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterLteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterLteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataAggregateParamsFilterNeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataAggregateParamsFilterNeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataAggregateParamsFilterNeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BetaAgentDataDeleteByQueryParams struct {
	// The agent deployment's name to delete data for
	DeploymentName string            `json:"deployment_name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// The logical agent data collection to delete from
	Collection param.Opt[string] `json:"collection,omitzero"`
	// Optional filters to select which items to delete
	Filter map[string]BetaAgentDataDeleteByQueryParamsFilter `json:"filter,omitzero"`
	paramObj
}

func (r BetaAgentDataDeleteByQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataDeleteByQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataDeleteByQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaAgentDataDeleteByQueryParams]'s query parameters as
// `url.Values`.
func (r BetaAgentDataDeleteByQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// API request model for a filter comparison operation.
type BetaAgentDataDeleteByQueryParamsFilter struct {
	Eq       BetaAgentDataDeleteByQueryParamsFilterEqUnion         `json:"eq,omitzero" format:"date-time"`
	Gt       BetaAgentDataDeleteByQueryParamsFilterGtUnion         `json:"gt,omitzero" format:"date-time"`
	Gte      BetaAgentDataDeleteByQueryParamsFilterGteUnion        `json:"gte,omitzero" format:"date-time"`
	Lt       BetaAgentDataDeleteByQueryParamsFilterLtUnion         `json:"lt,omitzero" format:"date-time"`
	Lte      BetaAgentDataDeleteByQueryParamsFilterLteUnion        `json:"lte,omitzero" format:"date-time"`
	Ne       BetaAgentDataDeleteByQueryParamsFilterNeUnion         `json:"ne,omitzero" format:"date-time"`
	Excludes []*BetaAgentDataDeleteByQueryParamsFilterExcludeUnion `json:"excludes,omitzero" format:"date-time"`
	Includes []*BetaAgentDataDeleteByQueryParamsFilterIncludeUnion `json:"includes,omitzero" format:"date-time"`
	paramObj
}

func (r BetaAgentDataDeleteByQueryParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataDeleteByQueryParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataDeleteByQueryParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterEqUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterEqUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterEqUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterExcludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterGtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterGtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterGtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterGteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterGteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterGteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterIncludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterIncludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterIncludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterLtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterLtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterLtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterLteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterLteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterLteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataDeleteByQueryParamsFilterNeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataDeleteByQueryParamsFilterNeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataDeleteByQueryParamsFilterNeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type BetaAgentDataGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaAgentDataGetParams]'s query parameters as `url.Values`.
func (r BetaAgentDataGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaAgentDataSearchParams struct {
	// The agent deployment's name to search within
	DeploymentName string            `json:"deployment_name" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// The offset to start from. If not provided, the first page is returned
	Offset param.Opt[int64] `json:"offset,omitzero"`
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
	// The logical agent data collection to search within
	Collection param.Opt[string] `json:"collection,omitzero"`
	// Whether to include the total number of items in the response
	IncludeTotal param.Opt[bool] `json:"include_total,omitzero"`
	// A filter object or expression that filters resources listed in the response.
	Filter map[string]BetaAgentDataSearchParamsFilter `json:"filter,omitzero"`
	paramObj
}

func (r BetaAgentDataSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaAgentDataSearchParams]'s query parameters as
// `url.Values`.
func (r BetaAgentDataSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// API request model for a filter comparison operation.
type BetaAgentDataSearchParamsFilter struct {
	Eq       BetaAgentDataSearchParamsFilterEqUnion         `json:"eq,omitzero" format:"date-time"`
	Gt       BetaAgentDataSearchParamsFilterGtUnion         `json:"gt,omitzero" format:"date-time"`
	Gte      BetaAgentDataSearchParamsFilterGteUnion        `json:"gte,omitzero" format:"date-time"`
	Lt       BetaAgentDataSearchParamsFilterLtUnion         `json:"lt,omitzero" format:"date-time"`
	Lte      BetaAgentDataSearchParamsFilterLteUnion        `json:"lte,omitzero" format:"date-time"`
	Ne       BetaAgentDataSearchParamsFilterNeUnion         `json:"ne,omitzero" format:"date-time"`
	Excludes []*BetaAgentDataSearchParamsFilterExcludeUnion `json:"excludes,omitzero" format:"date-time"`
	Includes []*BetaAgentDataSearchParamsFilterIncludeUnion `json:"includes,omitzero" format:"date-time"`
	paramObj
}

func (r BetaAgentDataSearchParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow BetaAgentDataSearchParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaAgentDataSearchParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterEqUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterEqUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterEqUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterExcludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterExcludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterExcludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterGtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterGtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterGtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterGteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterGteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterGteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterIncludeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterIncludeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterIncludeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterLtUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterLtUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterLtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterLteUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterLteUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterLteUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaAgentDataSearchParamsFilterNeUnion struct {
	OfFloat  param.Opt[float64]   `json:",omitzero,inline"`
	OfString param.Opt[string]    `json:",omitzero,inline"`
	OfTime   param.Opt[time.Time] `json:",omitzero,inline"`
	paramUnion
}

func (u BetaAgentDataSearchParamsFilterNeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfTime)
}
func (u *BetaAgentDataSearchParamsFilterNeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
