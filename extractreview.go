// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

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
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// ExtractReviewService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractReviewService] method instead.
type ExtractReviewService struct {
	options []option.RequestOption
}

// NewExtractReviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtractReviewService(opts ...option.RequestOption) (r ExtractReviewService) {
	r = ExtractReviewService{}
	r.options = opts
	return
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Turn review on for an extract job.
//
// Fields read below `confidence_threshold`, and required fields left empty, are
// flagged for a person to check. Every setting has a default, so the body can be
// left out. A job still running can be named: it waits for review once it
// completes. A job that failed or was cancelled without a review, or that cannot
// be reviewed, answers 409 with `code: job_not_reviewable`; a field in
// `fields_in_scope` the job's schema does not have answers 400.
//
// Answers 201 with the new review, or 200 with the job's existing review, whose
// settings stay as they are.
func (r *ExtractReviewService) New(ctx context.Context, jobID string, params ExtractReviewNewParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Approve the reviewed result.
//
// While flagged fields nobody checked remain, answers 409 with `code: open_flags`
// and their `paths`, unless `sign_off` confirms them as part of approving. A job
// that has not completed answers 409 with `code: job_not_reviewable`. A `revision`
// other than the review's latest answers 409 with `code: stale_revision` and the
// latest `revision`.
func (r *ExtractReviewService) Approve(ctx context.Context, jobID string, params ExtractReviewApproveParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review/approve", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Confirm, correct, clear, restore or undo fields of a completed job.
//
// The changes apply in order, in one write. A job without a review gets one with
// the default settings, at revision 0. A path that is not a field of the job, or a
// value the field's schema refuses, answers 400 and changes nothing. A job that
// has not completed answers 409 with `code: job_not_reviewable`. A `revision`
// other than the review's latest answers 409 with `code: stale_revision` and the
// latest `revision`.
func (r *ExtractReviewService) ChangeFields(ctx context.Context, jobID string, params ExtractReviewChangeFieldsParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review/fields", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Get an extract job's review: its flags, field actions and decision, with the
// result as the model read it and as reviewed. A job without a review answers 404.
func (r *ExtractReviewService) Get(ctx context.Context, jobID string, query ExtractReviewGetParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// List a review's history, newest first.
func (r *ExtractReviewService) ListEvents(ctx context.Context, jobID string, query ExtractReviewListEventsParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ExtractReviewEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review/events", url.PathEscape(jobID))
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

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// List a review's history, newest first.
func (r *ExtractReviewService) ListEventsAutoPaging(ctx context.Context, jobID string, query ExtractReviewListEventsParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ExtractReviewEvent] {
	return pagination.NewPaginatedCursorAutoPager(r.ListEvents(ctx, jobID, query, opts...))
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Reject the result, with a reason. A job that failed or was cancelled can be
// rejected, which closes a review turned on while it ran. A job still queued or
// running answers 409 with `code: job_not_reviewable`. A `revision` other than the
// review's latest answers 409 with `code: stale_revision` and the latest
// `revision`.
func (r *ExtractReviewService) Reject(ctx context.Context, jobID string, params ExtractReviewRejectParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review/reject", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Beta: answers 403 until extract review is enabled for the caller and their
// organization.
//
// Put an approved or rejected review back to pending. Its field actions stay.
//
// A `revision` other than the review's latest answers 409 with
// `code: stale_revision` and the latest `revision`.
func (r *ExtractReviewService) Reopen(ctx context.Context, jobID string, params ExtractReviewReopenParams, opts ...option.RequestOption) (res *ExtractReview, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/extract/%s/review/reopen", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// An extract job's review: its settings, flags, field actions and decision, with
// the result as read and as reviewed.
type ExtractReview struct {
	// Unique identifier of the review.
	ID string `json:"id" api:"required"`
	// Fields read below this confidence are flagged.
	ConfidenceThreshold float64 `json:"confidence_threshold" api:"required"`
	// When review was turned on.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When it was approved or rejected.
	DecidedAt time.Time `json:"decided_at" api:"required" format:"date-time"`
	// ID of the user who approved or rejected it.
	DecidedBy string `json:"decided_by" api:"required"`
	// Why it was rejected.
	DecidedReason string `json:"decided_reason" api:"required"`
	// The latest action per field, keyed by JSON Pointer.
	FieldActions map[string]ExtractReviewFieldAction `json:"field_actions" api:"required"`
	// JSON Pointers of the fields the threshold applies to; empty means all.
	FieldsInScope []string `json:"fields_in_scope" api:"required"`
	// Fields the review asks a person to check.
	Flags []ExtractReviewFlag `json:"flags" api:"required"`
	// The extract job reviewed.
	JobID string `json:"job_id" api:"required"`
	// JSON Pointers of flagged fields nobody checked.
	OpenFlags []string `json:"open_flags" api:"required"`
	// The result as the model read it.
	OriginalResult map[string]any `json:"original_result" api:"required"`
	// JSON Pointers of fields flagged when empty.
	RequiredFields []string `json:"required_fields" api:"required"`
	// When the job and its document may be deleted.
	RetainUntil time.Time `json:"retain_until" api:"required" format:"date-time"`
	// The result with the field actions.
	ReviewedResult map[string]any `json:"reviewed_result" api:"required"`
	// Bumped by every change; actions name the one they saw.
	Revision int64 `json:"revision" api:"required"`
	// The file the job read, or the file behind the parse job it read; null for a URL.
	SourceFileID string `json:"source_file_id" api:"required"`
	// pending until approved or rejected.
	//
	// Any of "approved", "pending", "rejected".
	Status ExtractReviewStatus `json:"status" api:"required"`
	// When the review last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ConfidenceThreshold respjson.Field
		CreatedAt           respjson.Field
		DecidedAt           respjson.Field
		DecidedBy           respjson.Field
		DecidedReason       respjson.Field
		FieldActions        respjson.Field
		FieldsInScope       respjson.Field
		Flags               respjson.Field
		JobID               respjson.Field
		OpenFlags           respjson.Field
		OriginalResult      respjson.Field
		RequiredFields      respjson.Field
		RetainUntil         respjson.Field
		ReviewedResult      respjson.Field
		Revision            respjson.Field
		SourceFileID        respjson.Field
		Status              respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReview) RawJSON() string { return r.JSON.raw }
func (r *ExtractReview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The latest thing a person did to one field.
type ExtractReviewFieldAction struct {
	// confirm, correct or clear.
	//
	// Any of "clear", "confirm", "correct".
	Action string `json:"action" api:"required"`
	// ID of the user who acted.
	ActorID string `json:"actor_id" api:"required"`
	// When they acted.
	At time.Time `json:"at" api:"required" format:"date-time"`
	// The corrected value, for `correct`.
	Value any `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		ActorID     respjson.Field
		At          respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewFieldAction) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewFieldAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A field the review asks a person to check.
type ExtractReviewFlag struct {
	// Whether someone confirmed, corrected or cleared it.
	Handled bool `json:"handled" api:"required"`
	// JSON Pointer of the field.
	Path string `json:"path" api:"required"`
	// below_threshold, or required_empty for a required field left empty.
	//
	// Any of "below_threshold", "required_empty".
	Reason string `json:"reason" api:"required"`
	// Where in a list the least confident value is.
	Cell ExtractReviewFlagCell `json:"cell" api:"nullable"`
	// The field's confidence; for a list, its least confident value.
	Confidence float64 `json:"confidence" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handled     respjson.Field
		Path        respjson.Field
		Reason      respjson.Field
		Cell        respjson.Field
		Confidence  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewFlag) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewFlag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where in a list the least confident value is.
type ExtractReviewFlagCell struct {
	// Index of the row, from 0.
	Row int64 `json:"row" api:"required"`
	// Property of the row, when rows are objects.
	Column string `json:"column" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Row         respjson.Field
		Column      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewFlagCell) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewFlagCell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// pending until approved or rejected.
type ExtractReviewStatus string

const (
	ExtractReviewStatusApproved ExtractReviewStatus = "approved"
	ExtractReviewStatusPending  ExtractReviewStatus = "pending"
	ExtractReviewStatusRejected ExtractReviewStatus = "rejected"
)

// One entry in a review's history.
type ExtractReviewEvent struct {
	// Unique identifier of the entry.
	ID string `json:"id" api:"required"`
	// ID of the user who acted.
	ActorID string `json:"actor_id" api:"required"`
	// When it happened.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// What happened.
	//
	// Any of "approved", "created", "field_cleared", "field_confirmed",
	// "field_corrected", "field_restored", "field_undone", "flags_signed_off",
	// "rejected", "reopened".
	Kind ExtractReviewEventKind `json:"kind" api:"required"`
	// The settings for `created`, the reason for `rejected`, and the fields for
	// `flags_signed_off`; null for the rest.
	Detail ExtractReviewEventDetailUnion `json:"detail" api:"nullable"`
	// JSON Pointer of the field it changed.
	Path string `json:"path" api:"nullable"`
	// The field's value after the change.
	ValueAfter any `json:"value_after"`
	// The field's value before the change.
	ValueBefore any `json:"value_before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ActorID     respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Detail      respjson.Field
		Path        respjson.Field
		ValueAfter  respjson.Field
		ValueBefore respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewEvent) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What happened.
type ExtractReviewEventKind string

const (
	ExtractReviewEventKindApproved       ExtractReviewEventKind = "approved"
	ExtractReviewEventKindCreated        ExtractReviewEventKind = "created"
	ExtractReviewEventKindFieldCleared   ExtractReviewEventKind = "field_cleared"
	ExtractReviewEventKindFieldConfirmed ExtractReviewEventKind = "field_confirmed"
	ExtractReviewEventKindFieldCorrected ExtractReviewEventKind = "field_corrected"
	ExtractReviewEventKindFieldRestored  ExtractReviewEventKind = "field_restored"
	ExtractReviewEventKindFieldUndone    ExtractReviewEventKind = "field_undone"
	ExtractReviewEventKindFlagsSignedOff ExtractReviewEventKind = "flags_signed_off"
	ExtractReviewEventKindRejected       ExtractReviewEventKind = "rejected"
	ExtractReviewEventKindReopened       ExtractReviewEventKind = "reopened"
)

// ExtractReviewEventDetailUnion contains all possible properties and values from
// [ExtractReviewEventDetailExtractReviewCreatedDetail],
// [ExtractReviewEventDetailExtractReviewRejectedDetail],
// [ExtractReviewEventDetailExtractReviewSignedOffDetail].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExtractReviewEventDetailUnion struct {
	// This field is from variant [ExtractReviewEventDetailExtractReviewCreatedDetail].
	ConfidenceThreshold float64 `json:"confidence_threshold"`
	// This field is from variant [ExtractReviewEventDetailExtractReviewCreatedDetail].
	FieldsInScope []string `json:"fields_in_scope"`
	// This field is from variant [ExtractReviewEventDetailExtractReviewCreatedDetail].
	Origin string `json:"origin"`
	// This field is from variant [ExtractReviewEventDetailExtractReviewCreatedDetail].
	RequiredFields []string `json:"required_fields"`
	// This field is from variant
	// [ExtractReviewEventDetailExtractReviewRejectedDetail].
	Reason string `json:"reason"`
	// This field is from variant
	// [ExtractReviewEventDetailExtractReviewSignedOffDetail].
	Paths []string `json:"paths"`
	JSON  struct {
		ConfidenceThreshold respjson.Field
		FieldsInScope       respjson.Field
		Origin              respjson.Field
		RequiredFields      respjson.Field
		Reason              respjson.Field
		Paths               respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ExtractReviewEventDetailUnion) AsExtractReviewCreatedDetail() (v ExtractReviewEventDetailExtractReviewCreatedDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractReviewEventDetailUnion) AsExtractReviewRejectedDetail() (v ExtractReviewEventDetailExtractReviewRejectedDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractReviewEventDetailUnion) AsExtractReviewSignedOffDetail() (v ExtractReviewEventDetailExtractReviewSignedOffDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractReviewEventDetailUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractReviewEventDetailUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a `created` entry records: where review was turned on, and its settings.
type ExtractReviewEventDetailExtractReviewCreatedDetail struct {
	// Fields read below this confidence are flagged.
	ConfidenceThreshold float64 `json:"confidence_threshold" api:"required"`
	// JSON Pointers of the fields the threshold applies to; empty means all.
	FieldsInScope []string `json:"fields_in_scope" api:"required"`
	// Turned on as the job ran, from History, by a first edit, or by the API.
	//
	// Any of "api", "edit", "history", "run".
	Origin string `json:"origin" api:"required"`
	// JSON Pointers of fields flagged when empty.
	RequiredFields []string `json:"required_fields" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfidenceThreshold respjson.Field
		FieldsInScope       respjson.Field
		Origin              respjson.Field
		RequiredFields      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewEventDetailExtractReviewCreatedDetail) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewEventDetailExtractReviewCreatedDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a `rejected` entry records.
type ExtractReviewEventDetailExtractReviewRejectedDetail struct {
	// Why the result was rejected.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewEventDetailExtractReviewRejectedDetail) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewEventDetailExtractReviewRejectedDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a `flags_signed_off` entry records.
type ExtractReviewEventDetailExtractReviewSignedOffDetail struct {
	// JSON Pointers of the flagged fields confirmed as part of approving.
	Paths []string `json:"paths" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Paths       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractReviewEventDetailExtractReviewSignedOffDetail) RawJSON() string { return r.JSON.raw }
func (r *ExtractReviewEventDetailExtractReviewSignedOffDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractReviewNewParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields read below this confidence are flagged.
	ConfidenceThreshold param.Opt[float64] `json:"confidence_threshold,omitzero"`
	// JSON Pointers of fields flagged when empty; defaults to the schema's required.
	RequiredFields []string `json:"required_fields,omitzero"`
	// JSON Pointers of the fields the threshold applies to; empty means all.
	FieldsInScope []string `json:"fields_in_scope,omitzero"`
	paramObj
}

func (r ExtractReviewNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractReviewNewParams]'s query parameters as `url.Values`.
func (r ExtractReviewNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractReviewApproveParams struct {
	// The review's revision as the caller last read it.
	Revision       int64             `json:"revision" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Confirm the flagged fields nobody checked as part of approving.
	SignOff param.Opt[bool] `json:"sign_off,omitzero"`
	paramObj
}

func (r ExtractReviewApproveParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewApproveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewApproveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractReviewApproveParams]'s query parameters as
// `url.Values`.
func (r ExtractReviewApproveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractReviewChangeFieldsParams struct {
	// Changes to apply, in order.
	Changes []ExtractReviewChangeFieldsParamsChange `json:"changes,omitzero" api:"required"`
	// The review's revision as the caller last read it.
	Revision       int64             `json:"revision" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExtractReviewChangeFieldsParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewChangeFieldsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewChangeFieldsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractReviewChangeFieldsParams]'s query parameters as
// `url.Values`.
func (r ExtractReviewChangeFieldsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One change to one field.
//
// The properties Change, Path are required.
type ExtractReviewChangeFieldsParamsChange struct {
	// confirm, correct to `value`, clear, restore the model's value, or undo.
	//
	// Any of "clear", "confirm", "correct", "restore", "undo".
	Change string `json:"change,omitzero" api:"required"`
	// JSON Pointer of the field.
	Path string `json:"path" api:"required"`
	// The corrected value, for `correct`.
	Value any `json:"value,omitzero"`
	paramObj
}

func (r ExtractReviewChangeFieldsParamsChange) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewChangeFieldsParamsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewChangeFieldsParamsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractReviewChangeFieldsParamsChange](
		"change", "clear", "confirm", "correct", "restore", "undo",
	)
}

type ExtractReviewGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractReviewGetParams]'s query parameters as `url.Values`.
func (r ExtractReviewGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractReviewListEventsParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Token from the previous page.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Events per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractReviewListEventsParams]'s query parameters as
// `url.Values`.
func (r ExtractReviewListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractReviewRejectParams struct {
	// Why the result was rejected.
	Reason string `json:"reason" api:"required"`
	// The review's revision as the caller last read it.
	Revision       int64             `json:"revision" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExtractReviewRejectParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewRejectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewRejectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractReviewRejectParams]'s query parameters as
// `url.Values`.
func (r ExtractReviewRejectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractReviewReopenParams struct {
	// The review's revision as the caller last read it.
	Revision       int64             `json:"revision" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ExtractReviewReopenParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractReviewReopenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractReviewReopenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ExtractReviewReopenParams]'s query parameters as
// `url.Values`.
func (r ExtractReviewReopenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
