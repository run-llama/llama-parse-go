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

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// ClassifierJobService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClassifierJobService] method instead.
type ClassifierJobService struct {
	options []option.RequestOption
}

// NewClassifierJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClassifierJobService(opts ...option.RequestOption) (r ClassifierJobService) {
	r = ClassifierJobService{}
	r.options = opts
	return
}

// Create a classify job. Experimental: not production-ready and subject to change.
//
// Deprecated: Please use `client.classify.create()`
func (r *ClassifierJobService) New(ctx context.Context, params ClassifierJobNewParams, opts ...option.RequestOption) (res *ClassifyJob, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/classifier/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List classify jobs. Experimental: not production-ready and subject to change.
//
// Deprecated: Please use `client.classify.list()`
func (r *ClassifierJobService) List(ctx context.Context, query ClassifierJobListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ClassifyJob], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/classifier/jobs"
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

// List classify jobs. Experimental: not production-ready and subject to change.
//
// Deprecated: Please use `client.classify.list()`
func (r *ClassifierJobService) ListAutoPaging(ctx context.Context, query ClassifierJobListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ClassifyJob] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Get a classify job. Experimental: not production-ready and subject to change.
//
// Deprecated: Please use `client.classify.get()`
func (r *ClassifierJobService) Get(ctx context.Context, classifyJobID string, query ClassifierJobGetParams, opts ...option.RequestOption) (res *ClassifyJob, err error) {
	opts = slices.Concat(r.options, opts)
	if classifyJobID == "" {
		err = errors.New("missing required classify_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/classifier/jobs/%s", classifyJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the results of a classify job. Experimental: not production-ready and
// subject to change.
//
// Deprecated: Please use `client.classify.get()`
func (r *ClassifierJobService) GetResults(ctx context.Context, classifyJobID string, query ClassifierJobGetResultsParams, opts ...option.RequestOption) (res *ClassifierJobGetResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if classifyJobID == "" {
		err = errors.New("missing required classify_job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/classifier/jobs/%s/results", classifyJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A rule for classifying documents - v0 simplified version.
//
// This represents a single classification rule that will be applied to documents.
// All rules are content-based and use natural language descriptions.
type ClassifierRule struct {
	// Natural language description of what to classify. Be specific about the content
	// characteristics that identify this document type.
	Description string `json:"description" api:"required"`
	// The document type to assign when this rule matches (e.g., 'invoice', 'receipt',
	// 'contract')
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifierRule) RawJSON() string { return r.JSON.raw }
func (r *ClassifierRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassifierRule to a ClassifierRuleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassifierRuleParam.Overrides()
func (r ClassifierRule) ToParam() ClassifierRuleParam {
	return param.Override[ClassifierRuleParam](json.RawMessage(r.RawJSON()))
}

// A rule for classifying documents - v0 simplified version.
//
// This represents a single classification rule that will be applied to documents.
// All rules are content-based and use natural language descriptions.
//
// The properties Description, Type are required.
type ClassifierRuleParam struct {
	// Natural language description of what to classify. Be specific about the content
	// characteristics that identify this document type.
	Description string `json:"description" api:"required"`
	// The document type to assign when this rule matches (e.g., 'invoice', 'receipt',
	// 'contract')
	Type string `json:"type" api:"required"`
	paramObj
}

func (r ClassifierRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifierRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifierRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A classify job.
type ClassifyJob struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the project
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// The rules to classify the files
	Rules []ClassifierRule `json:"rules" api:"required"`
	// The status of the classify job
	//
	// Any of "PENDING", "SUCCESS", "ERROR", "PARTIAL_SUCCESS", "CANCELLED".
	Status StatusEnum `json:"status" api:"required"`
	// The ID of the user
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt   time.Time `json:"created_at" api:"nullable" format:"date-time"`
	EffectiveAt time.Time `json:"effective_at" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// The job record ID associated with this status, if any.
	JobRecordID string `json:"job_record_id" api:"nullable"`
	// The classification mode to use
	//
	// Any of "FAST", "MULTIMODAL".
	Mode ClassifyJobMode `json:"mode"`
	// The configuration for the parsing job
	ParsingConfiguration ClassifyParsingConfiguration `json:"parsing_configuration"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ProjectID            respjson.Field
		Rules                respjson.Field
		Status               respjson.Field
		UserID               respjson.Field
		CreatedAt            respjson.Field
		EffectiveAt          respjson.Field
		ErrorMessage         respjson.Field
		JobRecordID          respjson.Field
		Mode                 respjson.Field
		ParsingConfiguration respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyJob) RawJSON() string { return r.JSON.raw }
func (r *ClassifyJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassifyJob to a ClassifyJobParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassifyJobParam.Overrides()
func (r ClassifyJob) ToParam() ClassifyJobParam {
	return param.Override[ClassifyJobParam](json.RawMessage(r.RawJSON()))
}

// The classification mode to use
type ClassifyJobMode string

const (
	ClassifyJobModeFast       ClassifyJobMode = "FAST"
	ClassifyJobModeMultimodal ClassifyJobMode = "MULTIMODAL"
)

// A classify job.
//
// The properties ID, ProjectID, Rules, Status, UserID are required.
type ClassifyJobParam struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the project
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// The rules to classify the files
	Rules []ClassifierRuleParam `json:"rules,omitzero" api:"required"`
	// The status of the classify job
	//
	// Any of "PENDING", "SUCCESS", "ERROR", "PARTIAL_SUCCESS", "CANCELLED".
	Status StatusEnum `json:"status,omitzero" api:"required"`
	// The ID of the user
	UserID string `json:"user_id" api:"required"`
	// Creation datetime
	CreatedAt param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	// Error message for the latest job attempt, if any.
	ErrorMessage param.Opt[string] `json:"error_message,omitzero"`
	// The job record ID associated with this status, if any.
	JobRecordID param.Opt[string] `json:"job_record_id,omitzero"`
	// Update datetime
	UpdatedAt   param.Opt[time.Time] `json:"updated_at,omitzero" format:"date-time"`
	EffectiveAt param.Opt[time.Time] `json:"effective_at,omitzero" format:"date-time"`
	// The classification mode to use
	//
	// Any of "FAST", "MULTIMODAL".
	Mode ClassifyJobMode `json:"mode,omitzero"`
	// The configuration for the parsing job
	ParsingConfiguration ClassifyParsingConfigurationParam `json:"parsing_configuration,omitzero"`
	paramObj
}

func (r ClassifyJobParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyJobParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyJobParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration for a classify job.
type ClassifyParsingConfiguration struct {
	// The language to parse the files in
	//
	// Any of "af", "az", "bs", "cs", "cy", "da", "de", "en", "es", "et", "fr", "ga",
	// "hr", "hu", "id", "is", "it", "ku", "la", "lt", "lv", "mi", "ms", "mt", "nl",
	// "no", "oc", "pi", "pl", "pt", "ro", "rs_latin", "sk", "sl", "sq", "sv", "sw",
	// "tl", "tr", "uz", "vi", "ar", "fa", "ug", "ur", "bn", "as", "mni", "ru",
	// "rs_cyrillic", "be", "bg", "uk", "mn", "abq", "ady", "kbd", "ava", "dar", "inh",
	// "che", "lbe", "lez", "tab", "tjk", "hi", "mr", "ne", "bh", "mai", "ang", "bho",
	// "mah", "sck", "new", "gom", "sa", "bgc", "th", "ch_sim", "ch_tra", "ja", "ko",
	// "ta", "te", "kn".
	Lang ParsingLanguages `json:"lang"`
	// The maximum number of pages to parse
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// The pages to target for parsing (0-indexed, so first page is at 0)
	TargetPages []int64 `json:"target_pages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lang        respjson.Field
		MaxPages    respjson.Field
		TargetPages respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyParsingConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ClassifyParsingConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassifyParsingConfiguration to a
// ClassifyParsingConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassifyParsingConfigurationParam.Overrides()
func (r ClassifyParsingConfiguration) ToParam() ClassifyParsingConfigurationParam {
	return param.Override[ClassifyParsingConfigurationParam](json.RawMessage(r.RawJSON()))
}

// Parsing configuration for a classify job.
type ClassifyParsingConfigurationParam struct {
	// The maximum number of pages to parse
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// The pages to target for parsing (0-indexed, so first page is at 0)
	TargetPages []int64 `json:"target_pages,omitzero"`
	// The language to parse the files in
	//
	// Any of "af", "az", "bs", "cs", "cy", "da", "de", "en", "es", "et", "fr", "ga",
	// "hr", "hu", "id", "is", "it", "ku", "la", "lt", "lv", "mi", "ms", "mt", "nl",
	// "no", "oc", "pi", "pl", "pt", "ro", "rs_latin", "sk", "sl", "sq", "sv", "sw",
	// "tl", "tr", "uz", "vi", "ar", "fa", "ug", "ur", "bn", "as", "mni", "ru",
	// "rs_cyrillic", "be", "bg", "uk", "mn", "abq", "ady", "kbd", "ava", "dar", "inh",
	// "che", "lbe", "lez", "tab", "tjk", "hi", "mr", "ne", "bh", "mai", "ang", "bho",
	// "mah", "sck", "new", "gom", "sa", "bgc", "th", "ch_sim", "ch_tra", "ja", "ko",
	// "ta", "te", "kn".
	Lang ParsingLanguages `json:"lang,omitzero"`
	paramObj
}

func (r ClassifyParsingConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyParsingConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyParsingConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the classify endpoint following AIP-132 pagination standard.
type ClassifierJobGetResultsResponse struct {
	// The list of items.
	Items []ClassifierJobGetResultsResponseItem `json:"items" api:"required"`
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
func (r ClassifierJobGetResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ClassifierJobGetResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file classification.
type ClassifierJobGetResultsResponseItem struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the classify job
	ClassifyJobID string `json:"classify_job_id" api:"required" format:"uuid"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// The ID of the classified file
	FileID string `json:"file_id" api:"nullable" format:"uuid"`
	// Result of classifying a single file.
	Result ClassifierJobGetResultsResponseItemResult `json:"result" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ClassifyJobID respjson.Field
		CreatedAt     respjson.Field
		FileID        respjson.Field
		Result        respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifierJobGetResultsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ClassifierJobGetResultsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of classifying a single file.
type ClassifierJobGetResultsResponseItemResult struct {
	// Confidence score of the classification (0.0-1.0)
	Confidence float64 `json:"confidence" api:"required"`
	// Step-by-step explanation of why this classification was chosen and the
	// confidence score assigned
	Reasoning string `json:"reasoning" api:"required"`
	// The document type that best matches, or null if no match.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Reasoning   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifierJobGetResultsResponseItemResult) RawJSON() string { return r.JSON.raw }
func (r *ClassifierJobGetResultsResponseItemResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassifierJobNewParams struct {
	// The IDs of the files to classify
	FileIDs []string `json:"file_ids,omitzero" api:"required" format:"uuid"`
	// The rules to classify the files
	Rules          []ClassifierRuleParam `json:"rules,omitzero" api:"required"`
	OrganizationID param.Opt[string]     `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]     `query:"project_id,omitzero" format:"uuid" json:"-"`
	// The classification mode to use
	//
	// Any of "FAST", "MULTIMODAL".
	Mode ClassifierJobNewParamsMode `json:"mode,omitzero"`
	// The configuration for the parsing job
	ParsingConfiguration ClassifyParsingConfigurationParam `json:"parsing_configuration,omitzero"`
	// List of webhook configurations for notifications
	WebhookConfigurations []ClassifierJobNewParamsWebhookConfiguration `json:"webhook_configurations,omitzero"`
	paramObj
}

func (r ClassifierJobNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClassifierJobNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifierJobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ClassifierJobNewParams]'s query parameters as `url.Values`.
func (r ClassifierJobNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The classification mode to use
type ClassifierJobNewParamsMode string

const (
	ClassifierJobNewParamsModeFast       ClassifierJobNewParamsMode = "FAST"
	ClassifierJobNewParamsModeMultimodal ClassifierJobNewParamsMode = "MULTIMODAL"
)

// Webhook configuration for receiving parsing job notifications.
//
// Webhooks are called when specified events occur during job processing. Configure
// multiple webhook configurations to send to different endpoints.
type ClassifierJobNewParamsWebhookConfiguration struct {
	// HTTPS URL to receive webhook POST requests. Must be publicly accessible
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events that trigger this webhook. Options: 'parse.success' (job completed),
	// 'parse.failure' (job failed), 'parse.partial' (some pages failed). If not
	// specified, webhook fires for all events
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers to include in webhook requests. Use for authentication
	// tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}
	WebhookHeaders map[string]any `json:"webhook_headers,omitzero"`
	paramObj
}

func (r ClassifierJobNewParamsWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ClassifierJobNewParamsWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifierJobNewParamsWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassifierJobListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ClassifierJobListParams]'s query parameters as
// `url.Values`.
func (r ClassifierJobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ClassifierJobGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ClassifierJobGetParams]'s query parameters as `url.Values`.
func (r ClassifierJobGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ClassifierJobGetResultsParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ClassifierJobGetResultsParams]'s query parameters as
// `url.Values`.
func (r ClassifierJobGetResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
