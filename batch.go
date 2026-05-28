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

// BatchService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	options []option.RequestOption
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r BatchService) {
	r = BatchService{}
	r.options = opts
	return
}

// Create a batch over a source directory and start processing asynchronously.
func (r *BatchService) New(ctx context.Context, params BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List batches for the current project.
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BatchListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/batches"
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

// List batches for the current project.
func (r *BatchService) ListAutoPaging(ctx context.Context, query BatchListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BatchListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a batch by ID.
func (r *BatchService) Get(ctx context.Context, batchID string, query BatchGetParams, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/batches/%s", url.PathEscape(batchID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A top-level batch.
//
// Example: { "id": "bat-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "project_id":
// "prj-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "source_directory_id":
// "dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "config": { "job": { "type":
// "parse_v2", "configuration_id": "cfg-PARSE_AGENTIC" } }, "status": "COMPLETED",
// "results": [ { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null } ] }
//
// Batch-level `FAILED` means the orchestration failed and cannot provide a
// reliable per-file result set. `results` is only populated when explicitly
// requested with `expand=results` and may be `null` while a batch is still
// running.
type BatchNewResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Batch configuration snapshot.
	Config BatchNewResponseConfig `json:"config" api:"required"`
	// Project this batch belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Directory being processed.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// Current batch status.
	//
	// Any of "PENDING", "THROTTLED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status BatchNewResponseStatus `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Expanded per-file result mappings. Null unless requested with expand=results, or
	// while the batch is still running.
	Results []BatchNewResponseResult `json:"results" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		Status            respjson.Field
		CreatedAt         respjson.Field
		Results           respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch configuration snapshot.
type BatchNewResponseConfig struct {
	// Job to create for each file in the source directory.
	Job BatchNewResponseConfigJob `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseConfig) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job to create for each file in the source directory.
type BatchNewResponseConfigJob struct {
	// Product configuration ID or built-in preset ID matching the job type.
	ConfigurationID string `json:"configuration_id" api:"required"`
	// Product job type to run for each source directory file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchNewResponseConfigJobType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigurationID respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseConfigJob) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseConfigJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product job type to run for each source directory file.
type BatchNewResponseConfigJobType string

const (
	BatchNewResponseConfigJobTypeParseV2   BatchNewResponseConfigJobType = "parse_v2"
	BatchNewResponseConfigJobTypeExtractV2 BatchNewResponseConfigJobType = "extract_v2"
)

// Current batch status.
type BatchNewResponseStatus string

const (
	BatchNewResponseStatusPending   BatchNewResponseStatus = "PENDING"
	BatchNewResponseStatusThrottled BatchNewResponseStatus = "THROTTLED"
	BatchNewResponseStatusRunning   BatchNewResponseStatus = "RUNNING"
	BatchNewResponseStatusCompleted BatchNewResponseStatus = "COMPLETED"
	BatchNewResponseStatusFailed    BatchNewResponseStatus = "FAILED"
	BatchNewResponseStatusCancelled BatchNewResponseStatus = "CANCELLED"
)

// Result projection for one source directory file in a batch.
//
// Example: { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null }
//
// This is a projection of directory-sync state, not a separate child resource that
// callers need to create. The source directory file ID is the stable correlation
// key. Underlying job progress and failures should be resolved through the
// referenced product job endpoint.
type BatchNewResponseResult struct {
	// Source directory file processed by this batch.
	SourceDirectoryFileID string `json:"source_directory_file_id" api:"required"`
	// Batch-level mapping error if the system could not create or associate a job for
	// this source file.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Reference to a job produced by a batch.
	//
	// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
	// }
	JobReference BatchNewResponseResultJobReference `json:"job_reference" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceDirectoryFileID respjson.Field
		ErrorMessage          respjson.Field
		JobReference          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to a job produced by a batch.
//
// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
// }
type BatchNewResponseResultJobReference struct {
	// Job ID, such as a parse job ID.
	ID string `json:"id" api:"required"`
	// Type of job produced for the file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchNewResponseResultJobReferenceType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchNewResponseResultJobReference) RawJSON() string { return r.JSON.raw }
func (r *BatchNewResponseResultJobReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of job produced for the file.
type BatchNewResponseResultJobReferenceType string

const (
	BatchNewResponseResultJobReferenceTypeParseV2   BatchNewResponseResultJobReferenceType = "parse_v2"
	BatchNewResponseResultJobReferenceTypeExtractV2 BatchNewResponseResultJobReferenceType = "extract_v2"
)

// A top-level batch.
//
// Example: { "id": "bat-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "project_id":
// "prj-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "source_directory_id":
// "dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "config": { "job": { "type":
// "parse_v2", "configuration_id": "cfg-PARSE_AGENTIC" } }, "status": "COMPLETED",
// "results": [ { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null } ] }
//
// Batch-level `FAILED` means the orchestration failed and cannot provide a
// reliable per-file result set. `results` is only populated when explicitly
// requested with `expand=results` and may be `null` while a batch is still
// running.
type BatchListResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Batch configuration snapshot.
	Config BatchListResponseConfig `json:"config" api:"required"`
	// Project this batch belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Directory being processed.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// Current batch status.
	//
	// Any of "PENDING", "THROTTLED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status BatchListResponseStatus `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Expanded per-file result mappings. Null unless requested with expand=results, or
	// while the batch is still running.
	Results []BatchListResponseResult `json:"results" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		Status            respjson.Field
		CreatedAt         respjson.Field
		Results           respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch configuration snapshot.
type BatchListResponseConfig struct {
	// Job to create for each file in the source directory.
	Job BatchListResponseConfigJob `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseConfig) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job to create for each file in the source directory.
type BatchListResponseConfigJob struct {
	// Product configuration ID or built-in preset ID matching the job type.
	ConfigurationID string `json:"configuration_id" api:"required"`
	// Product job type to run for each source directory file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchListResponseConfigJobType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigurationID respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseConfigJob) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseConfigJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product job type to run for each source directory file.
type BatchListResponseConfigJobType string

const (
	BatchListResponseConfigJobTypeParseV2   BatchListResponseConfigJobType = "parse_v2"
	BatchListResponseConfigJobTypeExtractV2 BatchListResponseConfigJobType = "extract_v2"
)

// Current batch status.
type BatchListResponseStatus string

const (
	BatchListResponseStatusPending   BatchListResponseStatus = "PENDING"
	BatchListResponseStatusThrottled BatchListResponseStatus = "THROTTLED"
	BatchListResponseStatusRunning   BatchListResponseStatus = "RUNNING"
	BatchListResponseStatusCompleted BatchListResponseStatus = "COMPLETED"
	BatchListResponseStatusFailed    BatchListResponseStatus = "FAILED"
	BatchListResponseStatusCancelled BatchListResponseStatus = "CANCELLED"
)

// Result projection for one source directory file in a batch.
//
// Example: { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null }
//
// This is a projection of directory-sync state, not a separate child resource that
// callers need to create. The source directory file ID is the stable correlation
// key. Underlying job progress and failures should be resolved through the
// referenced product job endpoint.
type BatchListResponseResult struct {
	// Source directory file processed by this batch.
	SourceDirectoryFileID string `json:"source_directory_file_id" api:"required"`
	// Batch-level mapping error if the system could not create or associate a job for
	// this source file.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Reference to a job produced by a batch.
	//
	// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
	// }
	JobReference BatchListResponseResultJobReference `json:"job_reference" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceDirectoryFileID respjson.Field
		ErrorMessage          respjson.Field
		JobReference          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to a job produced by a batch.
//
// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
// }
type BatchListResponseResultJobReference struct {
	// Job ID, such as a parse job ID.
	ID string `json:"id" api:"required"`
	// Type of job produced for the file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchListResponseResultJobReferenceType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchListResponseResultJobReference) RawJSON() string { return r.JSON.raw }
func (r *BatchListResponseResultJobReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of job produced for the file.
type BatchListResponseResultJobReferenceType string

const (
	BatchListResponseResultJobReferenceTypeParseV2   BatchListResponseResultJobReferenceType = "parse_v2"
	BatchListResponseResultJobReferenceTypeExtractV2 BatchListResponseResultJobReferenceType = "extract_v2"
)

// A top-level batch.
//
// Example: { "id": "bat-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "project_id":
// "prj-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "source_directory_id":
// "dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "config": { "job": { "type":
// "parse_v2", "configuration_id": "cfg-PARSE_AGENTIC" } }, "status": "COMPLETED",
// "results": [ { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null } ] }
//
// Batch-level `FAILED` means the orchestration failed and cannot provide a
// reliable per-file result set. `results` is only populated when explicitly
// requested with `expand=results` and may be `null` while a batch is still
// running.
type BatchGetResponse struct {
	// Unique identifier
	ID string `json:"id" api:"required"`
	// Batch configuration snapshot.
	Config BatchGetResponseConfig `json:"config" api:"required"`
	// Project this batch belongs to.
	ProjectID string `json:"project_id" api:"required"`
	// Directory being processed.
	SourceDirectoryID string `json:"source_directory_id" api:"required"`
	// Current batch status.
	//
	// Any of "PENDING", "THROTTLED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status BatchGetResponseStatus `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Expanded per-file result mappings. Null unless requested with expand=results, or
	// while the batch is still running.
	Results []BatchGetResponseResult `json:"results" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		ProjectID         respjson.Field
		SourceDirectoryID respjson.Field
		Status            respjson.Field
		CreatedAt         respjson.Field
		Results           respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch configuration snapshot.
type BatchGetResponseConfig struct {
	// Job to create for each file in the source directory.
	Job BatchGetResponseConfigJob `json:"job" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseConfig) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job to create for each file in the source directory.
type BatchGetResponseConfigJob struct {
	// Product configuration ID or built-in preset ID matching the job type.
	ConfigurationID string `json:"configuration_id" api:"required"`
	// Product job type to run for each source directory file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchGetResponseConfigJobType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigurationID respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseConfigJob) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseConfigJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product job type to run for each source directory file.
type BatchGetResponseConfigJobType string

const (
	BatchGetResponseConfigJobTypeParseV2   BatchGetResponseConfigJobType = "parse_v2"
	BatchGetResponseConfigJobTypeExtractV2 BatchGetResponseConfigJobType = "extract_v2"
)

// Current batch status.
type BatchGetResponseStatus string

const (
	BatchGetResponseStatusPending   BatchGetResponseStatus = "PENDING"
	BatchGetResponseStatusThrottled BatchGetResponseStatus = "THROTTLED"
	BatchGetResponseStatusRunning   BatchGetResponseStatus = "RUNNING"
	BatchGetResponseStatusCompleted BatchGetResponseStatus = "COMPLETED"
	BatchGetResponseStatusFailed    BatchGetResponseStatus = "FAILED"
	BatchGetResponseStatusCancelled BatchGetResponseStatus = "CANCELLED"
)

// Result projection for one source directory file in a batch.
//
// Example: { "source_directory_file_id":
// "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "job_reference": { "type":
// "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee" }, "error_message":
// null }
//
// This is a projection of directory-sync state, not a separate child resource that
// callers need to create. The source directory file ID is the stable correlation
// key. Underlying job progress and failures should be resolved through the
// referenced product job endpoint.
type BatchGetResponseResult struct {
	// Source directory file processed by this batch.
	SourceDirectoryFileID string `json:"source_directory_file_id" api:"required"`
	// Batch-level mapping error if the system could not create or associate a job for
	// this source file.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Reference to a job produced by a batch.
	//
	// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
	// }
	JobReference BatchGetResponseResultJobReference `json:"job_reference" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceDirectoryFileID respjson.Field
		ErrorMessage          respjson.Field
		JobReference          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to a job produced by a batch.
//
// Example: { "type": "parse_v2", "id": "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
// }
type BatchGetResponseResultJobReference struct {
	// Job ID, such as a parse job ID.
	ID string `json:"id" api:"required"`
	// Type of job produced for the file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchGetResponseResultJobReferenceType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchGetResponseResultJobReference) RawJSON() string { return r.JSON.raw }
func (r *BatchGetResponseResultJobReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of job produced for the file.
type BatchGetResponseResultJobReferenceType string

const (
	BatchGetResponseResultJobReferenceTypeParseV2   BatchGetResponseResultJobReferenceType = "parse_v2"
	BatchGetResponseResultJobReferenceTypeExtractV2 BatchGetResponseResultJobReferenceType = "extract_v2"
)

type BatchNewParams struct {
	// Batch configuration snapshot to apply to this source directory.
	Config BatchNewParamsConfig `json:"config,omitzero" api:"required"`
	// Directory whose files should be processed.
	SourceDirectoryID string            `json:"source_directory_id" api:"required"`
	OrganizationID    param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID         param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BatchNewParams]'s query parameters as `url.Values`.
func (r BatchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Batch configuration snapshot to apply to this source directory.
//
// The property Job is required.
type BatchNewParamsConfig struct {
	// Job to create for each file in the source directory.
	Job BatchNewParamsConfigJob `json:"job,omitzero" api:"required"`
	paramObj
}

func (r BatchNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job to create for each file in the source directory.
//
// The properties ConfigurationID, Type are required.
type BatchNewParamsConfigJob struct {
	// Product configuration ID or built-in preset ID matching the job type.
	ConfigurationID string `json:"configuration_id" api:"required"`
	// Product job type to run for each source directory file.
	//
	// Any of "parse_v2", "extract_v2".
	Type BatchNewParamsConfigJobType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r BatchNewParamsConfigJob) MarshalJSON() (data []byte, err error) {
	type shadow BatchNewParamsConfigJob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchNewParamsConfigJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product job type to run for each source directory file.
type BatchNewParamsConfigJobType string

const (
	BatchNewParamsConfigJobTypeParseV2   BatchNewParamsConfigJobType = "parse_v2"
	BatchNewParamsConfigJobTypeExtractV2 BatchNewParamsConfigJobType = "extract_v2"
)

type BatchListParams struct {
	CreatedAtOnOrAfter  param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize            param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	PageToken           param.Opt[string]    `query:"page_token,omitzero" json:"-"`
	ProjectID           param.Opt[string]    `query:"project_id,omitzero" format:"uuid" json:"-"`
	SourceDirectoryID   param.Opt[string]    `query:"source_directory_id,omitzero" json:"-"`
	// Any of "PENDING", "THROTTLED", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status BatchListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchListParamsStatus string

const (
	BatchListParamsStatusPending   BatchListParamsStatus = "PENDING"
	BatchListParamsStatusThrottled BatchListParamsStatus = "THROTTLED"
	BatchListParamsStatusRunning   BatchListParamsStatus = "RUNNING"
	BatchListParamsStatusCompleted BatchListParamsStatus = "COMPLETED"
	BatchListParamsStatusFailed    BatchListParamsStatus = "FAILED"
	BatchListParamsStatusCancelled BatchListParamsStatus = "CANCELLED"
)

type BatchGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields to expand. Supported value: results.
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BatchGetParams]'s query parameters as `url.Values`.
func (r BatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
