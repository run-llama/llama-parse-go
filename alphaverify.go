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

// AlphaVerifyService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaVerifyService] method instead.
type AlphaVerifyService struct {
	options []option.RequestOption
}

// NewAlphaVerifyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaVerifyService(opts ...option.RequestOption) (r AlphaVerifyService) {
	r = AlphaVerifyService{}
	r.options = opts
	return
}

// Create a Verify job.
//
// Analyzes a document for signs of doctoring (splicing, copy-move, AI generation,
// metadata tampering, ...). Set `file_input` to a file ID (`dfl-...`). Optionally
// provide a `configuration` object to control the semantic agent.
//
// The job runs asynchronously. Poll `GET /verify/{job_id}` with `expand=result` to
// check status and retrieve results.
func (r *AlphaVerifyService) New(ctx context.Context, params AlphaVerifyNewParams, opts ...option.RequestOption) (res *AlphaVerifyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/alpha/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List Verify jobs with optional filtering and pagination.
//
// Filter by `status`, specific `job_ids`, or creation date range.
func (r *AlphaVerifyService) List(ctx context.Context, query AlphaVerifyListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[AlphaVerifyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/alpha/verify"
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

// List Verify jobs with optional filtering and pagination.
//
// Filter by `status`, specific `job_ids`, or creation date range.
func (r *AlphaVerifyService) ListAutoPaging(ctx context.Context, query AlphaVerifyListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[AlphaVerifyListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Cancel a running Verify job.
//
// Stops processing and marks the job as CANCELLED. Returns the updated job. Jobs
// already in a terminal state (COMPLETED, FAILED, CANCELLED) cannot be cancelled.
func (r *AlphaVerifyService) Cancel(ctx context.Context, jobID string, body AlphaVerifyCancelParams, opts ...option.RequestOption) (res *AlphaVerifyCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/alpha/verify/%s/cancel", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a Verify job by ID.
//
// Returns the job status and configuration. Pass `expand=result` to include the
// Verify result (overall score, verdict, confidence, composite scores, and suspect
// regions) when the job is complete.
//
// Raw per-signal detail is available via `GET /verify/{job_id}/details`.
func (r *AlphaVerifyService) Get(ctx context.Context, jobID string, query AlphaVerifyGetParams, opts ...option.RequestOption) (res *AlphaVerifyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/alpha/verify/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the raw per-signal detail for a completed Verify job.
//
// Forensic drill-down behind the simplified result: the full evidence list,
// per-family sub-scores, raw localized regions, and per-page forensic heatmap
// overlays (presigned image URLs).
func (r *AlphaVerifyService) GetDetails(ctx context.Context, jobID string, query AlphaVerifyGetDetailsParams, opts ...option.RequestOption) (res *AlphaVerifyGetDetailsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/alpha/verify/%s/details", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response for a Verify job.
type AlphaVerifyNewResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Verify configuration used for this job
	Configuration AlphaVerifyNewResponseConfiguration `json:"configuration" api:"required"`
	// Type of the document input (FILE)
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType AlphaVerifyNewResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "CANCELLED", "COMPLETED", "FAILED", "PENDING", "RUNNING".
	Status AlphaVerifyNewResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a Verify (doctored-document) analysis.
	//
	// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
	// forensic heatmaps) is available separately via the job's details endpoint.
	Result AlphaVerifyNewResponseResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verify configuration used for this job
type AlphaVerifyNewResponseConfiguration struct {
	// Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all
	// pages. Ignored for non-PDF inputs.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata,
	// content integrity, container structure, pixel statistics); 'agentic' (default)
	// runs the full pipeline including the learned detectors and the semantic review
	// pass.
	//
	// Any of "agentic", "fast".
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetPages respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the document input (FILE)
type AlphaVerifyNewResponseDocumentInputType string

const (
	AlphaVerifyNewResponseDocumentInputTypeFileID     AlphaVerifyNewResponseDocumentInputType = "file_id"
	AlphaVerifyNewResponseDocumentInputTypeParseJobID AlphaVerifyNewResponseDocumentInputType = "parse_job_id"
	AlphaVerifyNewResponseDocumentInputTypeURL        AlphaVerifyNewResponseDocumentInputType = "url"
)

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type AlphaVerifyNewResponseStatus string

const (
	AlphaVerifyNewResponseStatusCancelled AlphaVerifyNewResponseStatus = "CANCELLED"
	AlphaVerifyNewResponseStatusCompleted AlphaVerifyNewResponseStatus = "COMPLETED"
	AlphaVerifyNewResponseStatusFailed    AlphaVerifyNewResponseStatus = "FAILED"
	AlphaVerifyNewResponseStatusPending   AlphaVerifyNewResponseStatus = "PENDING"
	AlphaVerifyNewResponseStatusRunning   AlphaVerifyNewResponseStatus = "RUNNING"
)

// Result of a Verify (doctored-document) analysis.
//
// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
// forensic heatmaps) is available separately via the job's details endpoint.
type AlphaVerifyNewResponseResult struct {
	// Version of the detector that produced the result
	DetectorVersion string `json:"detector_version" api:"required"`
	// Overall doctoring likelihood (0 to 1)
	OverallScore float64 `json:"overall_score" api:"required"`
	// Overall verdict for the document
	//
	// Any of "AUTHENTIC", "DOCTORED", "LIKELY_DOCTORED", "NO_STRONG_SIGNAL",
	// "SUSPICIOUS".
	Verdict string `json:"verdict" api:"required"`
	// Composite scores, each answering one question about the document
	CompositeScores AlphaVerifyNewResponseResultCompositeScores `json:"composite_scores"`
	// Confidence in the verdict (0 to 1): how firmly the detected signals support the
	// verdict bucket, independent of the doctoring likelihood itself
	Confidence float64 `json:"confidence"`
	// Error detail when the analysis could not complete
	Error string `json:"error" api:"nullable"`
	// Number of analysed pages (1 for images/docx)
	PageCount int64 `json:"page_count"`
	// Rendered pixel size per page, so region bboxes can be scaled onto the page
	PageDimensions []AlphaVerifyNewResponseResultPageDimension `json:"page_dimensions"`
	// Explanation of the verdict
	Reasoning string `json:"reasoning"`
	// Regions that led to the suspected fraud, ranked most-suspect first, each with an
	// explanation of what makes it suspect
	SuspectRegions []AlphaVerifyNewResponseResultSuspectRegion `json:"suspect_regions"`
	// Likelihood (0 to 1) that the document is wholly generated or fabricated rather
	// than a capture of a real document. Null for jobs completed before this score was
	// introduced
	SyntheticScore float64 `json:"synthetic_score" api:"nullable"`
	// Likelihood (0 to 1) that a real captured document was locally edited — a genuine
	// capture with regions altered after the fact. Null for jobs completed before this
	// score was introduced
	TamperingScore float64 `json:"tampering_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectorVersion respjson.Field
		OverallScore    respjson.Field
		Verdict         respjson.Field
		CompositeScores respjson.Field
		Confidence      respjson.Field
		Error           respjson.Field
		PageCount       respjson.Field
		PageDimensions  respjson.Field
		Reasoning       respjson.Field
		SuspectRegions  respjson.Field
		SyntheticScore  respjson.Field
		TamperingScore  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Composite scores, each answering one question about the document
type AlphaVerifyNewResponseResultCompositeScores struct {
	// Was this content synthesized by a generative model?
	AIGenerated AlphaVerifyNewResponseResultCompositeScoresAIGenerated `json:"ai_generated"`
	// Does the document's content agree with itself (checksums, arithmetic,
	// machine-readable zones)?
	DocumentCoherence AlphaVerifyNewResponseResultCompositeScoresDocumentCoherence `json:"document_coherence"`
	// Does the file's provenance / toolchain history look suspicious? Advisory:
	// individually weak workflow-hygiene signals
	DocumentMetadata AlphaVerifyNewResponseResultCompositeScoresDocumentMetadata `json:"document_metadata"`
	// Has this asset (or its template) been seen in fraud before?
	KnownFraud AlphaVerifyNewResponseResultCompositeScoresKnownFraud `json:"known_fraud"`
	// Was this document altered after creation (splice, retype, redact, inpaint)?
	ManuallyEdited AlphaVerifyNewResponseResultCompositeScoresManuallyEdited `json:"manually_edited"`
	// Was the document captured through a channel that destroys forensic evidence
	// (photo of a screen, print-then-rescan)?
	Recapture AlphaVerifyNewResponseResultCompositeScoresRecapture `json:"recapture"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerated       respjson.Field
		DocumentCoherence respjson.Field
		DocumentMetadata  respjson.Field
		KnownFraud        respjson.Field
		ManuallyEdited    respjson.Field
		Recapture         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScores) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultCompositeScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this content synthesized by a generative model?
type AlphaVerifyNewResponseResultCompositeScoresAIGenerated struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresAIGenerated) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultCompositeScoresAIGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the document's content agree with itself (checksums, arithmetic,
// machine-readable zones)?
type AlphaVerifyNewResponseResultCompositeScoresDocumentCoherence struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresDocumentCoherence) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyNewResponseResultCompositeScoresDocumentCoherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the file's provenance / toolchain history look suspicious? Advisory:
// individually weak workflow-hygiene signals
type AlphaVerifyNewResponseResultCompositeScoresDocumentMetadata struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresDocumentMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyNewResponseResultCompositeScoresDocumentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Has this asset (or its template) been seen in fraud before?
type AlphaVerifyNewResponseResultCompositeScoresKnownFraud struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresKnownFraud) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultCompositeScoresKnownFraud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this document altered after creation (splice, retype, redact, inpaint)?
type AlphaVerifyNewResponseResultCompositeScoresManuallyEdited struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresManuallyEdited) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyNewResponseResultCompositeScoresManuallyEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was the document captured through a channel that destroys forensic evidence
// (photo of a screen, print-then-rescan)?
type AlphaVerifyNewResponseResultCompositeScoresRecapture struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultCompositeScoresRecapture) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultCompositeScoresRecapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered pixel size of a page — the coordinate space region bboxes use, so the
// UI can scale the suspect-region overlay onto the displayed page.
type AlphaVerifyNewResponseResultPageDimension struct {
	// Rendered page height in pixels
	Height int64 `json:"height" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Rendered page width in pixels
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Page        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultPageDimension) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultPageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A region that led to the suspected fraud, with why it is suspect.
//
// A curated, high-signal subset of `regions`: reviewer-dismissed candidates are
// dropped and the remainder is ranked by suspicion, so consumers can act on
// `verdict` + `confidence` + this list without reading the raw signals.
type AlphaVerifyNewResponseResultSuspectRegion struct {
	// Region bounding box as [x, y, w, h] in page-render pixels
	Bbox []int64 `json:"bbox" api:"required"`
	// Human-readable explanation of what makes this region suspect
	Explanation string `json:"explanation" api:"required"`
	// Kind of anomaly detected in this region
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Suspicion score for this region (0 to 1)
	Score float64 `json:"score" api:"required"`
	// Detector that flagged this region
	Source string `json:"source" api:"required"`
	// Whether this region is part of the small set of decisive evidence behind the
	// verdict — the boxes a reviewer should look at first
	Primary bool `json:"primary"`
	// Automated reviewer verdict for this region (confirmed, dismissed, unsure, or
	// empty). A dismissed region can still be surfaced when it is the only place to
	// look; this label says how to read it
	Review string `json:"review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bbox        respjson.Field
		Explanation respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Primary     respjson.Field
		Review      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyNewResponseResultSuspectRegion) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyNewResponseResultSuspectRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a Verify job.
type AlphaVerifyListResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Verify configuration used for this job
	Configuration AlphaVerifyListResponseConfiguration `json:"configuration" api:"required"`
	// Type of the document input (FILE)
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType AlphaVerifyListResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "CANCELLED", "COMPLETED", "FAILED", "PENDING", "RUNNING".
	Status AlphaVerifyListResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a Verify (doctored-document) analysis.
	//
	// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
	// forensic heatmaps) is available separately via the job's details endpoint.
	Result AlphaVerifyListResponseResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verify configuration used for this job
type AlphaVerifyListResponseConfiguration struct {
	// Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all
	// pages. Ignored for non-PDF inputs.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata,
	// content integrity, container structure, pixel statistics); 'agentic' (default)
	// runs the full pipeline including the learned detectors and the semantic review
	// pass.
	//
	// Any of "agentic", "fast".
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetPages respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the document input (FILE)
type AlphaVerifyListResponseDocumentInputType string

const (
	AlphaVerifyListResponseDocumentInputTypeFileID     AlphaVerifyListResponseDocumentInputType = "file_id"
	AlphaVerifyListResponseDocumentInputTypeParseJobID AlphaVerifyListResponseDocumentInputType = "parse_job_id"
	AlphaVerifyListResponseDocumentInputTypeURL        AlphaVerifyListResponseDocumentInputType = "url"
)

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type AlphaVerifyListResponseStatus string

const (
	AlphaVerifyListResponseStatusCancelled AlphaVerifyListResponseStatus = "CANCELLED"
	AlphaVerifyListResponseStatusCompleted AlphaVerifyListResponseStatus = "COMPLETED"
	AlphaVerifyListResponseStatusFailed    AlphaVerifyListResponseStatus = "FAILED"
	AlphaVerifyListResponseStatusPending   AlphaVerifyListResponseStatus = "PENDING"
	AlphaVerifyListResponseStatusRunning   AlphaVerifyListResponseStatus = "RUNNING"
)

// Result of a Verify (doctored-document) analysis.
//
// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
// forensic heatmaps) is available separately via the job's details endpoint.
type AlphaVerifyListResponseResult struct {
	// Version of the detector that produced the result
	DetectorVersion string `json:"detector_version" api:"required"`
	// Overall doctoring likelihood (0 to 1)
	OverallScore float64 `json:"overall_score" api:"required"`
	// Overall verdict for the document
	//
	// Any of "AUTHENTIC", "DOCTORED", "LIKELY_DOCTORED", "NO_STRONG_SIGNAL",
	// "SUSPICIOUS".
	Verdict string `json:"verdict" api:"required"`
	// Composite scores, each answering one question about the document
	CompositeScores AlphaVerifyListResponseResultCompositeScores `json:"composite_scores"`
	// Confidence in the verdict (0 to 1): how firmly the detected signals support the
	// verdict bucket, independent of the doctoring likelihood itself
	Confidence float64 `json:"confidence"`
	// Error detail when the analysis could not complete
	Error string `json:"error" api:"nullable"`
	// Number of analysed pages (1 for images/docx)
	PageCount int64 `json:"page_count"`
	// Rendered pixel size per page, so region bboxes can be scaled onto the page
	PageDimensions []AlphaVerifyListResponseResultPageDimension `json:"page_dimensions"`
	// Explanation of the verdict
	Reasoning string `json:"reasoning"`
	// Regions that led to the suspected fraud, ranked most-suspect first, each with an
	// explanation of what makes it suspect
	SuspectRegions []AlphaVerifyListResponseResultSuspectRegion `json:"suspect_regions"`
	// Likelihood (0 to 1) that the document is wholly generated or fabricated rather
	// than a capture of a real document. Null for jobs completed before this score was
	// introduced
	SyntheticScore float64 `json:"synthetic_score" api:"nullable"`
	// Likelihood (0 to 1) that a real captured document was locally edited — a genuine
	// capture with regions altered after the fact. Null for jobs completed before this
	// score was introduced
	TamperingScore float64 `json:"tampering_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectorVersion respjson.Field
		OverallScore    respjson.Field
		Verdict         respjson.Field
		CompositeScores respjson.Field
		Confidence      respjson.Field
		Error           respjson.Field
		PageCount       respjson.Field
		PageDimensions  respjson.Field
		Reasoning       respjson.Field
		SuspectRegions  respjson.Field
		SyntheticScore  respjson.Field
		TamperingScore  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Composite scores, each answering one question about the document
type AlphaVerifyListResponseResultCompositeScores struct {
	// Was this content synthesized by a generative model?
	AIGenerated AlphaVerifyListResponseResultCompositeScoresAIGenerated `json:"ai_generated"`
	// Does the document's content agree with itself (checksums, arithmetic,
	// machine-readable zones)?
	DocumentCoherence AlphaVerifyListResponseResultCompositeScoresDocumentCoherence `json:"document_coherence"`
	// Does the file's provenance / toolchain history look suspicious? Advisory:
	// individually weak workflow-hygiene signals
	DocumentMetadata AlphaVerifyListResponseResultCompositeScoresDocumentMetadata `json:"document_metadata"`
	// Has this asset (or its template) been seen in fraud before?
	KnownFraud AlphaVerifyListResponseResultCompositeScoresKnownFraud `json:"known_fraud"`
	// Was this document altered after creation (splice, retype, redact, inpaint)?
	ManuallyEdited AlphaVerifyListResponseResultCompositeScoresManuallyEdited `json:"manually_edited"`
	// Was the document captured through a channel that destroys forensic evidence
	// (photo of a screen, print-then-rescan)?
	Recapture AlphaVerifyListResponseResultCompositeScoresRecapture `json:"recapture"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerated       respjson.Field
		DocumentCoherence respjson.Field
		DocumentMetadata  respjson.Field
		KnownFraud        respjson.Field
		ManuallyEdited    respjson.Field
		Recapture         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScores) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultCompositeScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this content synthesized by a generative model?
type AlphaVerifyListResponseResultCompositeScoresAIGenerated struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresAIGenerated) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultCompositeScoresAIGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the document's content agree with itself (checksums, arithmetic,
// machine-readable zones)?
type AlphaVerifyListResponseResultCompositeScoresDocumentCoherence struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresDocumentCoherence) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyListResponseResultCompositeScoresDocumentCoherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the file's provenance / toolchain history look suspicious? Advisory:
// individually weak workflow-hygiene signals
type AlphaVerifyListResponseResultCompositeScoresDocumentMetadata struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresDocumentMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyListResponseResultCompositeScoresDocumentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Has this asset (or its template) been seen in fraud before?
type AlphaVerifyListResponseResultCompositeScoresKnownFraud struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresKnownFraud) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultCompositeScoresKnownFraud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this document altered after creation (splice, retype, redact, inpaint)?
type AlphaVerifyListResponseResultCompositeScoresManuallyEdited struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresManuallyEdited) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyListResponseResultCompositeScoresManuallyEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was the document captured through a channel that destroys forensic evidence
// (photo of a screen, print-then-rescan)?
type AlphaVerifyListResponseResultCompositeScoresRecapture struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultCompositeScoresRecapture) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultCompositeScoresRecapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered pixel size of a page — the coordinate space region bboxes use, so the
// UI can scale the suspect-region overlay onto the displayed page.
type AlphaVerifyListResponseResultPageDimension struct {
	// Rendered page height in pixels
	Height int64 `json:"height" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Rendered page width in pixels
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Page        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultPageDimension) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultPageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A region that led to the suspected fraud, with why it is suspect.
//
// A curated, high-signal subset of `regions`: reviewer-dismissed candidates are
// dropped and the remainder is ranked by suspicion, so consumers can act on
// `verdict` + `confidence` + this list without reading the raw signals.
type AlphaVerifyListResponseResultSuspectRegion struct {
	// Region bounding box as [x, y, w, h] in page-render pixels
	Bbox []int64 `json:"bbox" api:"required"`
	// Human-readable explanation of what makes this region suspect
	Explanation string `json:"explanation" api:"required"`
	// Kind of anomaly detected in this region
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Suspicion score for this region (0 to 1)
	Score float64 `json:"score" api:"required"`
	// Detector that flagged this region
	Source string `json:"source" api:"required"`
	// Whether this region is part of the small set of decisive evidence behind the
	// verdict — the boxes a reviewer should look at first
	Primary bool `json:"primary"`
	// Automated reviewer verdict for this region (confirmed, dismissed, unsure, or
	// empty). A dismissed region can still be surfaced when it is the only place to
	// look; this label says how to read it
	Review string `json:"review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bbox        respjson.Field
		Explanation respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Primary     respjson.Field
		Review      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyListResponseResultSuspectRegion) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyListResponseResultSuspectRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a Verify job.
type AlphaVerifyCancelResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Verify configuration used for this job
	Configuration AlphaVerifyCancelResponseConfiguration `json:"configuration" api:"required"`
	// Type of the document input (FILE)
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType AlphaVerifyCancelResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "CANCELLED", "COMPLETED", "FAILED", "PENDING", "RUNNING".
	Status AlphaVerifyCancelResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a Verify (doctored-document) analysis.
	//
	// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
	// forensic heatmaps) is available separately via the job's details endpoint.
	Result AlphaVerifyCancelResponseResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verify configuration used for this job
type AlphaVerifyCancelResponseConfiguration struct {
	// Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all
	// pages. Ignored for non-PDF inputs.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata,
	// content integrity, container structure, pixel statistics); 'agentic' (default)
	// runs the full pipeline including the learned detectors and the semantic review
	// pass.
	//
	// Any of "agentic", "fast".
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetPages respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the document input (FILE)
type AlphaVerifyCancelResponseDocumentInputType string

const (
	AlphaVerifyCancelResponseDocumentInputTypeFileID     AlphaVerifyCancelResponseDocumentInputType = "file_id"
	AlphaVerifyCancelResponseDocumentInputTypeParseJobID AlphaVerifyCancelResponseDocumentInputType = "parse_job_id"
	AlphaVerifyCancelResponseDocumentInputTypeURL        AlphaVerifyCancelResponseDocumentInputType = "url"
)

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type AlphaVerifyCancelResponseStatus string

const (
	AlphaVerifyCancelResponseStatusCancelled AlphaVerifyCancelResponseStatus = "CANCELLED"
	AlphaVerifyCancelResponseStatusCompleted AlphaVerifyCancelResponseStatus = "COMPLETED"
	AlphaVerifyCancelResponseStatusFailed    AlphaVerifyCancelResponseStatus = "FAILED"
	AlphaVerifyCancelResponseStatusPending   AlphaVerifyCancelResponseStatus = "PENDING"
	AlphaVerifyCancelResponseStatusRunning   AlphaVerifyCancelResponseStatus = "RUNNING"
)

// Result of a Verify (doctored-document) analysis.
//
// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
// forensic heatmaps) is available separately via the job's details endpoint.
type AlphaVerifyCancelResponseResult struct {
	// Version of the detector that produced the result
	DetectorVersion string `json:"detector_version" api:"required"`
	// Overall doctoring likelihood (0 to 1)
	OverallScore float64 `json:"overall_score" api:"required"`
	// Overall verdict for the document
	//
	// Any of "AUTHENTIC", "DOCTORED", "LIKELY_DOCTORED", "NO_STRONG_SIGNAL",
	// "SUSPICIOUS".
	Verdict string `json:"verdict" api:"required"`
	// Composite scores, each answering one question about the document
	CompositeScores AlphaVerifyCancelResponseResultCompositeScores `json:"composite_scores"`
	// Confidence in the verdict (0 to 1): how firmly the detected signals support the
	// verdict bucket, independent of the doctoring likelihood itself
	Confidence float64 `json:"confidence"`
	// Error detail when the analysis could not complete
	Error string `json:"error" api:"nullable"`
	// Number of analysed pages (1 for images/docx)
	PageCount int64 `json:"page_count"`
	// Rendered pixel size per page, so region bboxes can be scaled onto the page
	PageDimensions []AlphaVerifyCancelResponseResultPageDimension `json:"page_dimensions"`
	// Explanation of the verdict
	Reasoning string `json:"reasoning"`
	// Regions that led to the suspected fraud, ranked most-suspect first, each with an
	// explanation of what makes it suspect
	SuspectRegions []AlphaVerifyCancelResponseResultSuspectRegion `json:"suspect_regions"`
	// Likelihood (0 to 1) that the document is wholly generated or fabricated rather
	// than a capture of a real document. Null for jobs completed before this score was
	// introduced
	SyntheticScore float64 `json:"synthetic_score" api:"nullable"`
	// Likelihood (0 to 1) that a real captured document was locally edited — a genuine
	// capture with regions altered after the fact. Null for jobs completed before this
	// score was introduced
	TamperingScore float64 `json:"tampering_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectorVersion respjson.Field
		OverallScore    respjson.Field
		Verdict         respjson.Field
		CompositeScores respjson.Field
		Confidence      respjson.Field
		Error           respjson.Field
		PageCount       respjson.Field
		PageDimensions  respjson.Field
		Reasoning       respjson.Field
		SuspectRegions  respjson.Field
		SyntheticScore  respjson.Field
		TamperingScore  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Composite scores, each answering one question about the document
type AlphaVerifyCancelResponseResultCompositeScores struct {
	// Was this content synthesized by a generative model?
	AIGenerated AlphaVerifyCancelResponseResultCompositeScoresAIGenerated `json:"ai_generated"`
	// Does the document's content agree with itself (checksums, arithmetic,
	// machine-readable zones)?
	DocumentCoherence AlphaVerifyCancelResponseResultCompositeScoresDocumentCoherence `json:"document_coherence"`
	// Does the file's provenance / toolchain history look suspicious? Advisory:
	// individually weak workflow-hygiene signals
	DocumentMetadata AlphaVerifyCancelResponseResultCompositeScoresDocumentMetadata `json:"document_metadata"`
	// Has this asset (or its template) been seen in fraud before?
	KnownFraud AlphaVerifyCancelResponseResultCompositeScoresKnownFraud `json:"known_fraud"`
	// Was this document altered after creation (splice, retype, redact, inpaint)?
	ManuallyEdited AlphaVerifyCancelResponseResultCompositeScoresManuallyEdited `json:"manually_edited"`
	// Was the document captured through a channel that destroys forensic evidence
	// (photo of a screen, print-then-rescan)?
	Recapture AlphaVerifyCancelResponseResultCompositeScoresRecapture `json:"recapture"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerated       respjson.Field
		DocumentCoherence respjson.Field
		DocumentMetadata  respjson.Field
		KnownFraud        respjson.Field
		ManuallyEdited    respjson.Field
		Recapture         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScores) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResultCompositeScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this content synthesized by a generative model?
type AlphaVerifyCancelResponseResultCompositeScoresAIGenerated struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresAIGenerated) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyCancelResponseResultCompositeScoresAIGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the document's content agree with itself (checksums, arithmetic,
// machine-readable zones)?
type AlphaVerifyCancelResponseResultCompositeScoresDocumentCoherence struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresDocumentCoherence) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyCancelResponseResultCompositeScoresDocumentCoherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the file's provenance / toolchain history look suspicious? Advisory:
// individually weak workflow-hygiene signals
type AlphaVerifyCancelResponseResultCompositeScoresDocumentMetadata struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresDocumentMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyCancelResponseResultCompositeScoresDocumentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Has this asset (or its template) been seen in fraud before?
type AlphaVerifyCancelResponseResultCompositeScoresKnownFraud struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresKnownFraud) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResultCompositeScoresKnownFraud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this document altered after creation (splice, retype, redact, inpaint)?
type AlphaVerifyCancelResponseResultCompositeScoresManuallyEdited struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresManuallyEdited) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyCancelResponseResultCompositeScoresManuallyEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was the document captured through a channel that destroys forensic evidence
// (photo of a screen, print-then-rescan)?
type AlphaVerifyCancelResponseResultCompositeScoresRecapture struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultCompositeScoresRecapture) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResultCompositeScoresRecapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered pixel size of a page — the coordinate space region bboxes use, so the
// UI can scale the suspect-region overlay onto the displayed page.
type AlphaVerifyCancelResponseResultPageDimension struct {
	// Rendered page height in pixels
	Height int64 `json:"height" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Rendered page width in pixels
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Page        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultPageDimension) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResultPageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A region that led to the suspected fraud, with why it is suspect.
//
// A curated, high-signal subset of `regions`: reviewer-dismissed candidates are
// dropped and the remainder is ranked by suspicion, so consumers can act on
// `verdict` + `confidence` + this list without reading the raw signals.
type AlphaVerifyCancelResponseResultSuspectRegion struct {
	// Region bounding box as [x, y, w, h] in page-render pixels
	Bbox []int64 `json:"bbox" api:"required"`
	// Human-readable explanation of what makes this region suspect
	Explanation string `json:"explanation" api:"required"`
	// Kind of anomaly detected in this region
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Suspicion score for this region (0 to 1)
	Score float64 `json:"score" api:"required"`
	// Detector that flagged this region
	Source string `json:"source" api:"required"`
	// Whether this region is part of the small set of decisive evidence behind the
	// verdict — the boxes a reviewer should look at first
	Primary bool `json:"primary"`
	// Automated reviewer verdict for this region (confirmed, dismissed, unsure, or
	// empty). A dismissed region can still be surfaced when it is the only place to
	// look; this label says how to read it
	Review string `json:"review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bbox        respjson.Field
		Explanation respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Primary     respjson.Field
		Review      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyCancelResponseResultSuspectRegion) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyCancelResponseResultSuspectRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for a Verify job.
type AlphaVerifyGetResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Verify configuration used for this job
	Configuration AlphaVerifyGetResponseConfiguration `json:"configuration" api:"required"`
	// Type of the document input (FILE)
	//
	// Any of "file_id", "parse_job_id", "url".
	DocumentInputType AlphaVerifyGetResponseDocumentInputType `json:"document_input_type" api:"required"`
	// ID of the input file
	FileInput string `json:"file_input" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "CANCELLED", "COMPLETED", "FAILED", "PENDING", "RUNNING".
	Status AlphaVerifyGetResponseStatus `json:"status" api:"required"`
	// User who created this job
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message if job failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Result of a Verify (doctored-document) analysis.
	//
	// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
	// forensic heatmaps) is available separately via the job's details endpoint.
	Result AlphaVerifyGetResponseResult `json:"result" api:"nullable"`
	// Idempotency key
	TransactionID string `json:"transaction_id" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Configuration     respjson.Field
		DocumentInputType respjson.Field
		FileInput         respjson.Field
		ProjectID         respjson.Field
		Status            respjson.Field
		UserID            respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		Result            respjson.Field
		TransactionID     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verify configuration used for this job
type AlphaVerifyGetResponseConfiguration struct {
	// Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all
	// pages. Ignored for non-PDF inputs.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata,
	// content integrity, container structure, pixel statistics); 'agentic' (default)
	// runs the full pipeline including the learned detectors and the semantic review
	// pass.
	//
	// Any of "agentic", "fast".
	Tier string `json:"tier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetPages respjson.Field
		Tier        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the document input (FILE)
type AlphaVerifyGetResponseDocumentInputType string

const (
	AlphaVerifyGetResponseDocumentInputTypeFileID     AlphaVerifyGetResponseDocumentInputType = "file_id"
	AlphaVerifyGetResponseDocumentInputTypeParseJobID AlphaVerifyGetResponseDocumentInputType = "parse_job_id"
	AlphaVerifyGetResponseDocumentInputTypeURL        AlphaVerifyGetResponseDocumentInputType = "url"
)

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type AlphaVerifyGetResponseStatus string

const (
	AlphaVerifyGetResponseStatusCancelled AlphaVerifyGetResponseStatus = "CANCELLED"
	AlphaVerifyGetResponseStatusCompleted AlphaVerifyGetResponseStatus = "COMPLETED"
	AlphaVerifyGetResponseStatusFailed    AlphaVerifyGetResponseStatus = "FAILED"
	AlphaVerifyGetResponseStatusPending   AlphaVerifyGetResponseStatus = "PENDING"
	AlphaVerifyGetResponseStatusRunning   AlphaVerifyGetResponseStatus = "RUNNING"
)

// Result of a Verify (doctored-document) analysis.
//
// Raw per-signal detail (evidence list, per-family sub-scores, raw regions,
// forensic heatmaps) is available separately via the job's details endpoint.
type AlphaVerifyGetResponseResult struct {
	// Version of the detector that produced the result
	DetectorVersion string `json:"detector_version" api:"required"`
	// Overall doctoring likelihood (0 to 1)
	OverallScore float64 `json:"overall_score" api:"required"`
	// Overall verdict for the document
	//
	// Any of "AUTHENTIC", "DOCTORED", "LIKELY_DOCTORED", "NO_STRONG_SIGNAL",
	// "SUSPICIOUS".
	Verdict string `json:"verdict" api:"required"`
	// Composite scores, each answering one question about the document
	CompositeScores AlphaVerifyGetResponseResultCompositeScores `json:"composite_scores"`
	// Confidence in the verdict (0 to 1): how firmly the detected signals support the
	// verdict bucket, independent of the doctoring likelihood itself
	Confidence float64 `json:"confidence"`
	// Error detail when the analysis could not complete
	Error string `json:"error" api:"nullable"`
	// Number of analysed pages (1 for images/docx)
	PageCount int64 `json:"page_count"`
	// Rendered pixel size per page, so region bboxes can be scaled onto the page
	PageDimensions []AlphaVerifyGetResponseResultPageDimension `json:"page_dimensions"`
	// Explanation of the verdict
	Reasoning string `json:"reasoning"`
	// Regions that led to the suspected fraud, ranked most-suspect first, each with an
	// explanation of what makes it suspect
	SuspectRegions []AlphaVerifyGetResponseResultSuspectRegion `json:"suspect_regions"`
	// Likelihood (0 to 1) that the document is wholly generated or fabricated rather
	// than a capture of a real document. Null for jobs completed before this score was
	// introduced
	SyntheticScore float64 `json:"synthetic_score" api:"nullable"`
	// Likelihood (0 to 1) that a real captured document was locally edited — a genuine
	// capture with regions altered after the fact. Null for jobs completed before this
	// score was introduced
	TamperingScore float64 `json:"tampering_score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectorVersion respjson.Field
		OverallScore    respjson.Field
		Verdict         respjson.Field
		CompositeScores respjson.Field
		Confidence      respjson.Field
		Error           respjson.Field
		PageCount       respjson.Field
		PageDimensions  respjson.Field
		Reasoning       respjson.Field
		SuspectRegions  respjson.Field
		SyntheticScore  respjson.Field
		TamperingScore  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Composite scores, each answering one question about the document
type AlphaVerifyGetResponseResultCompositeScores struct {
	// Was this content synthesized by a generative model?
	AIGenerated AlphaVerifyGetResponseResultCompositeScoresAIGenerated `json:"ai_generated"`
	// Does the document's content agree with itself (checksums, arithmetic,
	// machine-readable zones)?
	DocumentCoherence AlphaVerifyGetResponseResultCompositeScoresDocumentCoherence `json:"document_coherence"`
	// Does the file's provenance / toolchain history look suspicious? Advisory:
	// individually weak workflow-hygiene signals
	DocumentMetadata AlphaVerifyGetResponseResultCompositeScoresDocumentMetadata `json:"document_metadata"`
	// Has this asset (or its template) been seen in fraud before?
	KnownFraud AlphaVerifyGetResponseResultCompositeScoresKnownFraud `json:"known_fraud"`
	// Was this document altered after creation (splice, retype, redact, inpaint)?
	ManuallyEdited AlphaVerifyGetResponseResultCompositeScoresManuallyEdited `json:"manually_edited"`
	// Was the document captured through a channel that destroys forensic evidence
	// (photo of a screen, print-then-rescan)?
	Recapture AlphaVerifyGetResponseResultCompositeScoresRecapture `json:"recapture"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerated       respjson.Field
		DocumentCoherence respjson.Field
		DocumentMetadata  respjson.Field
		KnownFraud        respjson.Field
		ManuallyEdited    respjson.Field
		Recapture         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScores) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultCompositeScores) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this content synthesized by a generative model?
type AlphaVerifyGetResponseResultCompositeScoresAIGenerated struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresAIGenerated) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultCompositeScoresAIGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the document's content agree with itself (checksums, arithmetic,
// machine-readable zones)?
type AlphaVerifyGetResponseResultCompositeScoresDocumentCoherence struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresDocumentCoherence) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyGetResponseResultCompositeScoresDocumentCoherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Does the file's provenance / toolchain history look suspicious? Advisory:
// individually weak workflow-hygiene signals
type AlphaVerifyGetResponseResultCompositeScoresDocumentMetadata struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresDocumentMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyGetResponseResultCompositeScoresDocumentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Has this asset (or its template) been seen in fraud before?
type AlphaVerifyGetResponseResultCompositeScoresKnownFraud struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresKnownFraud) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultCompositeScoresKnownFraud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was this document altered after creation (splice, retype, redact, inpaint)?
type AlphaVerifyGetResponseResultCompositeScoresManuallyEdited struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresManuallyEdited) RawJSON() string {
	return r.JSON.raw
}
func (r *AlphaVerifyGetResponseResultCompositeScoresManuallyEdited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Was the document captured through a channel that destroys forensic evidence
// (photo of a screen, print-then-rescan)?
type AlphaVerifyGetResponseResultCompositeScoresRecapture struct {
	// Whether the checks feeding this composite ran on this document. When false the
	// document was not checked for this — not cleared of it
	Applicable bool `json:"applicable"`
	// Score (0 to 1); null when the composite was not applicable
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Applicable  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultCompositeScoresRecapture) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultCompositeScoresRecapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered pixel size of a page — the coordinate space region bboxes use, so the
// UI can scale the suspect-region overlay onto the displayed page.
type AlphaVerifyGetResponseResultPageDimension struct {
	// Rendered page height in pixels
	Height int64 `json:"height" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Rendered page width in pixels
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Page        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultPageDimension) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultPageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A region that led to the suspected fraud, with why it is suspect.
//
// A curated, high-signal subset of `regions`: reviewer-dismissed candidates are
// dropped and the remainder is ranked by suspicion, so consumers can act on
// `verdict` + `confidence` + this list without reading the raw signals.
type AlphaVerifyGetResponseResultSuspectRegion struct {
	// Region bounding box as [x, y, w, h] in page-render pixels
	Bbox []int64 `json:"bbox" api:"required"`
	// Human-readable explanation of what makes this region suspect
	Explanation string `json:"explanation" api:"required"`
	// Kind of anomaly detected in this region
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Suspicion score for this region (0 to 1)
	Score float64 `json:"score" api:"required"`
	// Detector that flagged this region
	Source string `json:"source" api:"required"`
	// Whether this region is part of the small set of decisive evidence behind the
	// verdict — the boxes a reviewer should look at first
	Primary bool `json:"primary"`
	// Automated reviewer verdict for this region (confirmed, dismissed, unsure, or
	// empty). A dismissed region can still be surfaced when it is the only place to
	// look; this label says how to read it
	Review string `json:"review"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bbox        respjson.Field
		Explanation respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Primary     respjson.Field
		Review      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetResponseResultSuspectRegion) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetResponseResultSuspectRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw per-signal detail for a completed Verify job.
//
// Forensic drill-down behind the simplified result: the full evidence list,
// per-family sub-scores, raw localized regions, and heatmap overlays.
type AlphaVerifyGetDetailsResponse struct {
	// ID of the Verify job
	JobID string `json:"job_id" api:"required"`
	// Checks that could not run on this job (with the reason). A check listed here
	// produced no findings because it could not run, not because the document is clean
	DegradedTools []AlphaVerifyGetDetailsResponseDegradedTool `json:"degraded_tools"`
	// Evidence items produced by detection tools
	Evidence []AlphaVerifyGetDetailsResponseEvidence `json:"evidence"`
	// Per-page forensic heatmap overlays as presigned image URLs
	Heatmaps []AlphaVerifyGetDetailsResponseHeatmap `json:"heatmaps"`
	// Rendered pixel size per page, so region bboxes can be scaled onto the page
	PageDimensions []AlphaVerifyGetDetailsResponsePageDimension `json:"page_dimensions"`
	// Suspicious regions localized on rendered pages
	Regions []AlphaVerifyGetDetailsResponseRegion `json:"regions"`
	// Per-family scores (metadata, ai_generation, splicing, copy_move, compression,
	// noise, coherence, pdf_structure)
	SubScores map[string]float64 `json:"sub_scores"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID          respjson.Field
		DegradedTools  respjson.Field
		Evidence       respjson.Field
		Heatmaps       respjson.Field
		PageDimensions respjson.Field
		Regions        respjson.Field
		SubScores      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A check that was attempted but could not run on this job.
type AlphaVerifyGetDetailsResponseDegradedTool struct {
	// Name of the check
	Tool string `json:"tool" api:"required"`
	// Why the check could not run
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tool        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponseDegradedTool) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponseDegradedTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single piece of evidence produced by a detection tool.
type AlphaVerifyGetDetailsResponseEvidence struct {
	// Machine-readable evidence code
	Code string `json:"code" api:"required"`
	// Human-readable evidence detail
	Detail string `json:"detail" api:"required"`
	// Signal family (e.g. metadata, splicing, compression)
	Family string `json:"family" api:"required"`
	// Evidence strength score
	Score float64 `json:"score" api:"required"`
	// Tool that produced this evidence
	Tool string `json:"tool" api:"required"`
	// Tool-specific structured payload
	Data map[string]any `json:"data"`
	// Whether this is hard (conclusive) evidence
	Hard bool `json:"hard"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Family      respjson.Field
		Score       respjson.Field
		Tool        respjson.Field
		Data        respjson.Field
		Hard        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponseEvidence) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponseEvidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A per-page forensic heatmap overlay, as a presigned image URL.
type AlphaVerifyGetDetailsResponseHeatmap struct {
	// The time at which the presigned URL expires
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Producing signal, e.g. double_compression, ela, noise
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Presigned URL to the heatmap PNG (page overlay)
	URL string `json:"url" api:"required"`
	// Producing tool's max score (for ranking)
	Score float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		URL         respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponseHeatmap) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponseHeatmap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rendered pixel size of a page — the coordinate space region bboxes use, so the
// UI can scale the suspect-region overlay onto the displayed page.
type AlphaVerifyGetDetailsResponsePageDimension struct {
	// Rendered page height in pixels
	Height int64 `json:"height" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Rendered page width in pixels
	Width int64 `json:"width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Page        respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponsePageDimension) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponsePageDimension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A suspicious region localized on a rendered page.
type AlphaVerifyGetDetailsResponseRegion struct {
	// Region bounding box as [x, y, w, h] in page-render pixels
	Bbox []int64 `json:"bbox" api:"required"`
	// Human-readable detail about the region
	Detail string `json:"detail" api:"required"`
	// Kind of anomaly detected in this region
	Kind string `json:"kind" api:"required"`
	// 0-based page index (0 for standalone images)
	Page int64 `json:"page" api:"required"`
	// Region-level doctoring likelihood score
	Score float64 `json:"score" api:"required"`
	// Detector/tool that produced this region
	Source string `json:"source" api:"required"`
	// Whether this region is part of the small set of decisive evidence behind the
	// verdict — the boxes a reviewer should look at first
	Primary bool `json:"primary"`
	// Review status/verdict for this region
	Review string `json:"review"`
	// Free-form review note for this region
	ReviewNote string `json:"review_note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bbox        respjson.Field
		Detail      respjson.Field
		Kind        respjson.Field
		Page        respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		Primary     respjson.Field
		Review      respjson.Field
		ReviewNote  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaVerifyGetDetailsResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *AlphaVerifyGetDetailsResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaVerifyNewParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Deprecated: use file_input instead
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// File ID of the document to analyze
	FileInput param.Opt[string] `json:"file_input,omitzero"`
	// Idempotency key scoped to the project. Reusing a key returns the original job;
	// the new request body is ignored.
	TransactionID param.Opt[string] `json:"transaction_id,omitzero"`
	// Configuration for a Verify job.
	Configuration AlphaVerifyNewParamsConfiguration `json:"configuration,omitzero"`
	// IDs of saved webhook configurations to notify for this job.
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids,omitzero"`
	// Outbound webhook endpoints to notify on job status changes
	WebhookConfigurations []AlphaVerifyNewParamsWebhookConfiguration `json:"webhook_configurations,omitzero"`
	paramObj
}

func (r AlphaVerifyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaVerifyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaVerifyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AlphaVerifyNewParams]'s query parameters as `url.Values`.
func (r AlphaVerifyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Configuration for a Verify job.
type AlphaVerifyNewParamsConfiguration struct {
	// Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all
	// pages. Ignored for non-PDF inputs.
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	// Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata,
	// content integrity, container structure, pixel statistics); 'agentic' (default)
	// runs the full pipeline including the learned detectors and the semantic review
	// pass.
	//
	// Any of "agentic", "fast".
	Tier string `json:"tier,omitzero"`
	paramObj
}

func (r AlphaVerifyNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow AlphaVerifyNewParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaVerifyNewParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaVerifyNewParamsConfiguration](
		"tier", "agentic", "fast",
	)
}

// Configuration for a single outbound webhook endpoint.
type AlphaVerifyNewParamsWebhookConfiguration struct {
	// Response format sent to the webhook: 'string' (default) or 'json'
	WebhookOutputFormat param.Opt[string] `json:"webhook_output_format,omitzero"`
	// Shared signing secret used to sign webhook deliveries. When set, each request
	// includes an HMAC-SHA256 signature of the request body in the 'LC-Signature'
	// header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with
	// this secret to verify the delivery is authentic.
	WebhookSigningSecret param.Opt[string] `json:"webhook_signing_secret,omitzero"`
	// URL to receive webhook POST notifications
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events to subscribe to (e.g. 'parse.success', 'extract.error'). If null, all
	// events are delivered.
	//
	// Any of "batch.cancelled", "batch.error", "batch.pending", "batch.running",
	// "batch.success", "classify.cancelled", "classify.error",
	// "classify.partial_success", "classify.pending", "classify.running",
	// "classify.success", "extract.cancelled", "extract.error",
	// "extract.partial_success", "extract.pending", "extract.success",
	// "parse.cancelled", "parse.error", "parse.partial_success", "parse.pending",
	// "parse.running", "parse.success", "sheets.cancelled", "sheets.error",
	// "sheets.partial_success", "sheets.pending", "sheets.success", "split.cancelled",
	// "split.error", "split.pending", "split.processing", "split.success",
	// "unmapped_event", "verify.cancelled", "verify.error", "verify.pending",
	// "verify.running", "verify.success".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers sent with each webhook request (e.g. auth tokens)
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	paramObj
}

func (r AlphaVerifyNewParamsWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow AlphaVerifyNewParamsWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaVerifyNewParamsWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaVerifyListParams struct {
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by job status
	//
	// Any of "CANCELLED", "COMPLETED", "FAILED", "PENDING", "RUNNING".
	Status AlphaVerifyListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional fields to include (e.g. `result`).
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaVerifyListParams]'s query parameters as `url.Values`.
func (r AlphaVerifyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status
type AlphaVerifyListParamsStatus string

const (
	AlphaVerifyListParamsStatusCancelled AlphaVerifyListParamsStatus = "CANCELLED"
	AlphaVerifyListParamsStatusCompleted AlphaVerifyListParamsStatus = "COMPLETED"
	AlphaVerifyListParamsStatusFailed    AlphaVerifyListParamsStatus = "FAILED"
	AlphaVerifyListParamsStatusPending   AlphaVerifyListParamsStatus = "PENDING"
	AlphaVerifyListParamsStatusRunning   AlphaVerifyListParamsStatus = "RUNNING"
)

type AlphaVerifyCancelParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaVerifyCancelParams]'s query parameters as
// `url.Values`.
func (r AlphaVerifyCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlphaVerifyGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Optional fields to include (e.g. `result`).
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaVerifyGetParams]'s query parameters as `url.Values`.
func (r AlphaVerifyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlphaVerifyGetDetailsParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaVerifyGetDetailsParams]'s query parameters as
// `url.Values`.
func (r AlphaVerifyGetDetailsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
