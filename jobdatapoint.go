// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"context"
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

// JobDataPointService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobDataPointService] method instead.
type JobDataPointService struct {
	options []option.RequestOption
}

// NewJobDataPointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobDataPointService(opts ...option.RequestOption) (r JobDataPointService) {
	r = JobDataPointService{}
	r.options = opts
	return
}

// Returns paginated job data points for the current project.
func (r *JobDataPointService) List(ctx context.Context, query JobDataPointListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[JobDataPoint], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/job-data-points"
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

// Returns paginated job data points for the current project.
func (r *JobDataPointService) ListAutoPaging(ctx context.Context, query JobDataPointListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[JobDataPoint] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// A job data point.
type JobDataPoint struct {
	// Job ID.
	ID string `json:"id" api:"required"`
	// Created timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Custom tag.
	CustomTag string `json:"custom_tag" api:"required"`
	// Project ID.
	ProjectID string `json:"project_id" api:"required"`
	// Job status.
	Status string `json:"status" api:"required"`
	// Updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Error message, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Job state transition timestamps.
	StateTransitions JobDataPointStateTransitions `json:"state_transitions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		CustomTag        respjson.Field
		ProjectID        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ErrorMessage     respjson.Field
		StateTransitions respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobDataPoint) RawJSON() string { return r.JSON.raw }
func (r *JobDataPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job state transition timestamps.
type JobDataPointStateTransitions struct {
	CancelledAt time.Time `json:"cancelled_at" api:"nullable" format:"date-time"`
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	FailedAt    time.Time `json:"failed_at" api:"nullable" format:"date-time"`
	PendingAt   time.Time `json:"pending_at" api:"nullable" format:"date-time"`
	RunningAt   time.Time `json:"running_at" api:"nullable" format:"date-time"`
	ThrottledAt time.Time `json:"throttled_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CancelledAt respjson.Field
		CompletedAt respjson.Field
		FailedAt    respjson.Field
		PendingAt   respjson.Field
		RunningAt   respjson.Field
		ThrottledAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobDataPointStateTransitions) RawJSON() string { return r.JSON.raw }
func (r *JobDataPointStateTransitions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobDataPointListParams struct {
	// Job type to query.
	//
	// Any of "classify", "extract", "parse".
	JobType JobDataPointListParamsJobType `query:"job_type,omitzero" api:"required" json:"-"`
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Cursor token for the next page.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Hours of history to include.
	Hours param.Opt[int64] `query:"hours,omitzero" json:"-"`
	// Filter by status.
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobDataPointListParams]'s query parameters as `url.Values`.
func (r JobDataPointListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Job type to query.
type JobDataPointListParamsJobType string

const (
	JobDataPointListParamsJobTypeClassify JobDataPointListParamsJobType = "classify"
	JobDataPointListParamsJobTypeExtract  JobDataPointListParamsJobType = "extract"
	JobDataPointListParamsJobTypeParse    JobDataPointListParamsJobType = "parse"
)
