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
	shimjson "github.com/run-llama/llama-parse-go/internal/encoding/json"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
	"github.com/run-llama/llama-parse-go/shared/constant"
)

// WebhookConfigService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookConfigService] method instead.
type WebhookConfigService struct {
	options []option.RequestOption
}

// NewWebhookConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookConfigService(opts ...option.RequestOption) (r WebhookConfigService) {
	r = WebhookConfigService{}
	r.options = opts
	return
}

// Create a reusable webhook configuration for the current project.
func (r *WebhookConfigService) New(ctx context.Context, params WebhookConfigNewParams, opts ...option.RequestOption) (res *WebhookConfigResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/webhook-configs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single webhook configuration by ID.
func (r *WebhookConfigService) Get(ctx context.Context, configID string, query WebhookConfigGetParams, opts ...option.RequestOption) (res *WebhookConfigResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/webhook-configs/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update a webhook configuration. Only fields present in the request change.
func (r *WebhookConfigService) Update(ctx context.Context, configID string, params WebhookConfigUpdateParams, opts ...option.RequestOption) (res *WebhookConfigResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/webhook-configs/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List the webhook configurations for the current project, newest first.
func (r *WebhookConfigService) List(ctx context.Context, query WebhookConfigListParams, opts ...option.RequestOption) (res *[]WebhookConfigResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/webhook-configs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a webhook configuration.
func (r *WebhookConfigService) Delete(ctx context.Context, configID string, body WebhookConfigDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/webhook-configs/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Request to create a stored webhook configuration.
//
// The owning tenant is taken from the request context (e.g. the project in the
// path), not the body.
//
// The property WebhookURL is required.
type WebhookConfigCreateParam struct {
	// URL to receive webhook POST notifications.
	WebhookURL string `json:"webhook_url" api:"required"`
	// Shared secret used to sign deliveries to this endpoint. Write-only: it is never
	// returned in responses.
	WebhookSigningSecret param.Opt[string] `json:"webhook_signing_secret,omitzero"`
	// Events to subscribe to. If null, all events are delivered.
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
	// "unmapped_event".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers sent with each webhook request.
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	// Response format sent to the webhook: 'string' (default) or 'json'.
	//
	// Any of "json", "string".
	WebhookOutputFormat WebhookConfigCreateWebhookOutputFormat `json:"webhook_output_format,omitzero"`
	paramObj
}

func (r WebhookConfigCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookConfigCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookConfigCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response format sent to the webhook: 'string' (default) or 'json'.
type WebhookConfigCreateWebhookOutputFormat string

const (
	WebhookConfigCreateWebhookOutputFormatJson   WebhookConfigCreateWebhookOutputFormat = "json"
	WebhookConfigCreateWebhookOutputFormatString WebhookConfigCreateWebhookOutputFormat = "string"
)

// A stored webhook configuration. The signing secret is never included.
type WebhookConfigResponse struct {
	// Unique identifier for the webhook configuration.
	ID string `json:"id" api:"required"`
	// Whether a signing secret is configured for this endpoint.
	HasSecret bool `json:"has_secret" api:"required"`
	// Owner tenant ID.
	TenantID string `json:"tenant_id" api:"required"`
	// Owner tenant type.
	TenantType constant.Project `json:"tenant_type" default:"project"`
	// URL that receives webhook POST notifications.
	WebhookURL string `json:"webhook_url" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Subscribed events (null = all events).
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
	// "unmapped_event".
	WebhookEvents []string `json:"webhook_events" api:"nullable"`
	// Custom HTTP headers sent with each request.
	WebhookHeaders map[string]string `json:"webhook_headers" api:"nullable"`
	// Response format sent to the webhook.
	//
	// Any of "json", "string".
	WebhookOutputFormat WebhookConfigResponseWebhookOutputFormat `json:"webhook_output_format" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		HasSecret           respjson.Field
		TenantID            respjson.Field
		TenantType          respjson.Field
		WebhookURL          respjson.Field
		CreatedAt           respjson.Field
		UpdatedAt           respjson.Field
		WebhookEvents       respjson.Field
		WebhookHeaders      respjson.Field
		WebhookOutputFormat respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response format sent to the webhook.
type WebhookConfigResponseWebhookOutputFormat string

const (
	WebhookConfigResponseWebhookOutputFormatJson   WebhookConfigResponseWebhookOutputFormat = "json"
	WebhookConfigResponseWebhookOutputFormatString WebhookConfigResponseWebhookOutputFormat = "string"
)

type WebhookConfigNewParams struct {
	// Request to create a stored webhook configuration.
	//
	// The owning tenant is taken from the request context (e.g. the project in the
	// path), not the body.
	WebhookConfigCreate WebhookConfigCreateParam
	OrganizationID      param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID           param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r WebhookConfigNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WebhookConfigCreate)
}
func (r *WebhookConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookConfigNewParams]'s query parameters as `url.Values`.
func (r WebhookConfigNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookConfigGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookConfigGetParams]'s query parameters as `url.Values`.
func (r WebhookConfigGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookConfigUpdateParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Updated signing secret (write-only). Send to rotate the secret.
	WebhookSigningSecret param.Opt[string] `json:"webhook_signing_secret,omitzero"`
	// Updated webhook URL.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Updated event subscriptions.
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
	// "unmapped_event".
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Updated headers.
	WebhookHeaders map[string]string `json:"webhook_headers,omitzero"`
	// Updated output format.
	//
	// Any of "json", "string".
	WebhookOutputFormat WebhookConfigUpdateParamsWebhookOutputFormat `json:"webhook_output_format,omitzero"`
	paramObj
}

func (r WebhookConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookConfigUpdateParams]'s query parameters as
// `url.Values`.
func (r WebhookConfigUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Updated output format.
type WebhookConfigUpdateParamsWebhookOutputFormat string

const (
	WebhookConfigUpdateParamsWebhookOutputFormatJson   WebhookConfigUpdateParamsWebhookOutputFormat = "json"
	WebhookConfigUpdateParamsWebhookOutputFormatString WebhookConfigUpdateParamsWebhookOutputFormat = "string"
)

type WebhookConfigListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookConfigListParams]'s query parameters as
// `url.Values`.
func (r WebhookConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookConfigDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookConfigDeleteParams]'s query parameters as
// `url.Values`.
func (r WebhookConfigDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
