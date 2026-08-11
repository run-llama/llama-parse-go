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

// V2ProjectService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ProjectService] method instead.
type V2ProjectService struct {
	options []option.RequestOption
}

// NewV2ProjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ProjectService(opts ...option.RequestOption) (r V2ProjectService) {
	r = V2ProjectService{}
	r.options = opts
	return
}

// List projects in an organization. Requires `organization_id` or a project-scoped
// API key.
func (r *V2ProjectService) List(ctx context.Context, query V2ProjectListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[V2ProjectListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/projects"
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

// List projects in an organization. Requires `organization_id` or a project-scoped
// API key.
func (r *V2ProjectService) ListAutoPaging(ctx context.Context, query V2ProjectListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[V2ProjectListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a project by ID.
func (r *V2ProjectService) Get(ctx context.Context, projectID string, query V2ProjectGetParams, opts ...option.RequestOption) (res *V2ProjectGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// API response schema for a project.
type V2ProjectListResponse struct {
	// The project's unique identifier.
	ID string `json:"id" api:"required"`
	// The project's display name.
	Name string `json:"name" api:"required"`
	// The organization the project belongs to.
	OrganizationID string `json:"organization_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Whether this project is the default project for its organization.
	IsDefault bool `json:"is_default"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		CreatedAt      respjson.Field
		IsDefault      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API response schema for a project.
type V2ProjectGetResponse struct {
	// The project's unique identifier.
	ID string `json:"id" api:"required"`
	// The project's display name.
	Name string `json:"name" api:"required"`
	// The organization the project belongs to.
	OrganizationID string `json:"organization_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Whether this project is the default project for its organization.
	IsDefault bool `json:"is_default"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		CreatedAt      respjson.Field
		IsDefault      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ProjectListParams struct {
	Name           param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2ProjectListParams]'s query parameters as `url.Values`.
func (r V2ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2ProjectGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V2ProjectGetParams]'s query parameters as `url.Values`.
func (r V2ProjectGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
