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

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// BetaChatService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaChatService] method instead.
type BetaChatService struct {
	options []option.RequestOption
}

// NewBetaChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaChatService(opts ...option.RequestOption) (r BetaChatService) {
	r = BetaChatService{}
	r.options = opts
	return
}

// Create a chat session, optionally bound to indexes (locked after the first
// message).
func (r *BetaChatService) New(ctx context.Context, params BetaChatNewParams, opts ...option.RequestOption) (res *BetaChatNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a full session by ID, including its event history.
func (r *BetaChatService) Get(ctx context.Context, sessionID string, query BetaChatGetParams, opts ...option.RequestOption) (res *BetaChatGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/chat/%s", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List all chat sessions for the current project.
func (r *BetaChatService) List(ctx context.Context, query BetaChatListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BetaChatListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/chat"
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

// List all chat sessions for the current project.
func (r *BetaChatService) ListAutoPaging(ctx context.Context, query BetaChatListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BetaChatListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a session.
func (r *BetaChatService) Delete(ctx context.Context, sessionID string, body BetaChatDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/chat/%s", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieve a session summary by ID.
func (r *BetaChatService) GetSummary(ctx context.Context, sessionID string, query BetaChatGetSummaryParams, opts ...option.RequestOption) (res *BetaChatGetSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/chat/%s/summary", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stream agent events for a chat turn as Server-Sent Events.
func (r *BetaChatService) Stream(ctx context.Context, sessionID string, params BetaChatStreamParams, opts ...option.RequestOption) (res *BetaChatStreamResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/chat/%s/messages/stream", url.PathEscape(sessionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Summary of a chat session, including its title and last run metadata.
type BetaChatNewResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata BetaChatNewResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaChatNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type BetaChatNewResponseJobMetadata struct {
	DurationMs        float64  `json:"duration_ms"`
	Error             string   `json:"error" api:"nullable"`
	ExportConfigIDs   []string `json:"export_config_ids" api:"nullable"`
	IsError           bool     `json:"is_error"`
	TotalInputTokens  int64    `json:"total_input_tokens" api:"nullable"`
	TotalOutputTokens int64    `json:"total_output_tokens" api:"nullable"`
	Turns             int64    `json:"turns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		Error             respjson.Field
		ExportConfigIDs   respjson.Field
		IsError           respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		Turns             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatNewResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *BetaChatNewResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full chat session including its complete event history.
type BetaChatGetResponse struct {
	// Ordered list of events that make up the conversation history.
	Events []BetaChatGetResponseEventUnion `json:"events" api:"required"`
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata BetaChatGetResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events         respjson.Field
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BetaChatGetResponseEventUnion contains all possible properties and values from
// [BetaChatGetResponseEventThinkingDelta], [BetaChatGetResponseEventTextDelta],
// [BetaChatGetResponseEventThinking], [BetaChatGetResponseEventText],
// [BetaChatGetResponseEventToolCall], [BetaChatGetResponseEventToolResult],
// [BetaChatGetResponseEventStop], [BetaChatGetResponseEventUserInput].
//
// Use the [BetaChatGetResponseEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BetaChatGetResponseEventUnion struct {
	Content string `json:"content"`
	// Any of "thinking_delta", "text_delta", "thinking", "text", "tool_call",
	// "tool_result", "stop", "user_input".
	Type string `json:"type"`
	// This field is from variant [BetaChatGetResponseEventToolCall].
	Arguments map[string]any `json:"arguments"`
	CallID    string         `json:"call_id"`
	Name      string         `json:"name"`
	// This field is from variant [BetaChatGetResponseEventToolResult].
	Result any `json:"result"`
	// This field is from variant [BetaChatGetResponseEventToolResult].
	ImageAttachment BetaChatGetResponseEventToolResultImageAttachment `json:"image_attachment"`
	// This field is from variant [BetaChatGetResponseEventStop].
	Error string `json:"error"`
	// This field is from variant [BetaChatGetResponseEventStop].
	IsError bool `json:"is_error"`
	// This field is from variant [BetaChatGetResponseEventStop].
	Usage BetaChatGetResponseEventStopUsage `json:"usage"`
	JSON  struct {
		Content         respjson.Field
		Type            respjson.Field
		Arguments       respjson.Field
		CallID          respjson.Field
		Name            respjson.Field
		Result          respjson.Field
		ImageAttachment respjson.Field
		Error           respjson.Field
		IsError         respjson.Field
		Usage           respjson.Field
		raw             string
	} `json:"-"`
}

// anyBetaChatGetResponseEvent is implemented by each variant of
// [BetaChatGetResponseEventUnion] to add type safety for the return type of
// [BetaChatGetResponseEventUnion.AsAny]
type anyBetaChatGetResponseEvent interface {
	implBetaChatGetResponseEventUnion()
}

func (BetaChatGetResponseEventThinkingDelta) implBetaChatGetResponseEventUnion() {}
func (BetaChatGetResponseEventTextDelta) implBetaChatGetResponseEventUnion()     {}
func (BetaChatGetResponseEventThinking) implBetaChatGetResponseEventUnion()      {}
func (BetaChatGetResponseEventText) implBetaChatGetResponseEventUnion()          {}
func (BetaChatGetResponseEventToolCall) implBetaChatGetResponseEventUnion()      {}
func (BetaChatGetResponseEventToolResult) implBetaChatGetResponseEventUnion()    {}
func (BetaChatGetResponseEventStop) implBetaChatGetResponseEventUnion()          {}
func (BetaChatGetResponseEventUserInput) implBetaChatGetResponseEventUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BetaChatGetResponseEventUnion.AsAny().(type) {
//	case llamacloudprod.BetaChatGetResponseEventThinkingDelta:
//	case llamacloudprod.BetaChatGetResponseEventTextDelta:
//	case llamacloudprod.BetaChatGetResponseEventThinking:
//	case llamacloudprod.BetaChatGetResponseEventText:
//	case llamacloudprod.BetaChatGetResponseEventToolCall:
//	case llamacloudprod.BetaChatGetResponseEventToolResult:
//	case llamacloudprod.BetaChatGetResponseEventStop:
//	case llamacloudprod.BetaChatGetResponseEventUserInput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BetaChatGetResponseEventUnion) AsAny() anyBetaChatGetResponseEvent {
	switch u.Type {
	case "thinking_delta":
		return u.AsThinkingDelta()
	case "text_delta":
		return u.AsTextDelta()
	case "thinking":
		return u.AsThinking()
	case "text":
		return u.AsText()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "stop":
		return u.AsStop()
	case "user_input":
		return u.AsUserInput()
	}
	return nil
}

func (u BetaChatGetResponseEventUnion) AsThinkingDelta() (v BetaChatGetResponseEventThinkingDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsTextDelta() (v BetaChatGetResponseEventTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsThinking() (v BetaChatGetResponseEventThinking) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsText() (v BetaChatGetResponseEventText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsToolCall() (v BetaChatGetResponseEventToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsToolResult() (v BetaChatGetResponseEventToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsStop() (v BetaChatGetResponseEventStop) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BetaChatGetResponseEventUnion) AsUserInput() (v BetaChatGetResponseEventUserInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BetaChatGetResponseEventUnion) RawJSON() string { return u.JSON.raw }

func (r *BetaChatGetResponseEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventThinkingDelta struct {
	Content string `json:"content" api:"required"`
	// Any of "thinking_delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventThinkingDelta) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventThinkingDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventTextDelta struct {
	Content string `json:"content" api:"required"`
	// Any of "text_delta".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventTextDelta) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventThinking struct {
	Content string `json:"content" api:"required"`
	// Any of "thinking".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventThinking) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventThinking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventText struct {
	Content string `json:"content" api:"required"`
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventText) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventToolCall struct {
	Arguments map[string]any `json:"arguments" api:"required"`
	CallID    string         `json:"call_id" api:"required"`
	Name      string         `json:"name" api:"required"`
	// Any of "tool_call".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		CallID      respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventToolCall) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventToolResult struct {
	CallID string `json:"call_id" api:"required"`
	Name   string `json:"name" api:"required"`
	Result any    `json:"result" api:"required"`
	// Coordinates for lazily resolving a page screenshot presigned URL.
	ImageAttachment BetaChatGetResponseEventToolResultImageAttachment `json:"image_attachment" api:"nullable"`
	// Any of "tool_result".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID          respjson.Field
		Name            respjson.Field
		Result          respjson.Field
		ImageAttachment respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventToolResult) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coordinates for lazily resolving a page screenshot presigned URL.
type BetaChatGetResponseEventToolResultImageAttachment struct {
	AttachmentName string `json:"attachment_name" api:"required"`
	SourceID       string `json:"source_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentName respjson.Field
		SourceID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventToolResultImageAttachment) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventToolResultImageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventStop struct {
	Error   string                            `json:"error" api:"required"`
	IsError bool                              `json:"is_error" api:"required"`
	Usage   BetaChatGetResponseEventStopUsage `json:"usage" api:"required"`
	// Any of "stop".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		IsError     respjson.Field
		Usage       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventStop) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventStopUsage struct {
	DurationMs        float64 `json:"duration_ms"`
	TotalInputTokens  int64   `json:"total_input_tokens" api:"nullable"`
	TotalOutputTokens int64   `json:"total_output_tokens" api:"nullable"`
	Turns             int64   `json:"turns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		Turns             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventStopUsage) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventStopUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatGetResponseEventUserInput struct {
	Content string `json:"content" api:"required"`
	// Any of "user_input".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseEventUserInput) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseEventUserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type BetaChatGetResponseJobMetadata struct {
	DurationMs        float64  `json:"duration_ms"`
	Error             string   `json:"error" api:"nullable"`
	ExportConfigIDs   []string `json:"export_config_ids" api:"nullable"`
	IsError           bool     `json:"is_error"`
	TotalInputTokens  int64    `json:"total_input_tokens" api:"nullable"`
	TotalOutputTokens int64    `json:"total_output_tokens" api:"nullable"`
	Turns             int64    `json:"turns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		Error             respjson.Field
		ExportConfigIDs   respjson.Field
		IsError           respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		Turns             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a chat session, including its title and last run metadata.
type BetaChatListResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata BetaChatListResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaChatListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type BetaChatListResponseJobMetadata struct {
	DurationMs        float64  `json:"duration_ms"`
	Error             string   `json:"error" api:"nullable"`
	ExportConfigIDs   []string `json:"export_config_ids" api:"nullable"`
	IsError           bool     `json:"is_error"`
	TotalInputTokens  int64    `json:"total_input_tokens" api:"nullable"`
	TotalOutputTokens int64    `json:"total_output_tokens" api:"nullable"`
	Turns             int64    `json:"turns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		Error             respjson.Field
		ExportConfigIDs   respjson.Field
		IsError           respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		Turns             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatListResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *BetaChatListResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a chat session, including its title and last run metadata.
type BetaChatGetSummaryResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata BetaChatGetSummaryResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type BetaChatGetSummaryResponseJobMetadata struct {
	DurationMs        float64  `json:"duration_ms"`
	Error             string   `json:"error" api:"nullable"`
	ExportConfigIDs   []string `json:"export_config_ids" api:"nullable"`
	IsError           bool     `json:"is_error"`
	TotalInputTokens  int64    `json:"total_input_tokens" api:"nullable"`
	TotalOutputTokens int64    `json:"total_output_tokens" api:"nullable"`
	Turns             int64    `json:"turns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		Error             respjson.Field
		ExportConfigIDs   respjson.Field
		IsError           respjson.Field
		TotalInputTokens  respjson.Field
		TotalOutputTokens respjson.Field
		Turns             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaChatGetSummaryResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *BetaChatGetSummaryResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaChatStreamResponse = any

type BetaChatNewParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Indexes this session will retrieve from. Once set and the first message has been
	// sent, the source set is locked for the session's lifetime. Leave null to create
	// an unbound session.
	IndexIDs []string `json:"index_ids,omitzero"`
	paramObj
}

func (r BetaChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaChatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaChatNewParams]'s query parameters as `url.Values`.
func (r BetaChatNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaChatGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaChatGetParams]'s query parameters as `url.Values`.
func (r BetaChatGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaChatListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaChatListParams]'s query parameters as `url.Values`.
func (r BetaChatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaChatDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaChatDeleteParams]'s query parameters as `url.Values`.
func (r BetaChatDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaChatGetSummaryParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaChatGetSummaryParams]'s query parameters as
// `url.Values`.
func (r BetaChatGetSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaChatStreamParams struct {
	// Indexes to retrieve data from.
	IndexIDs []string `json:"index_ids,omitzero" api:"required"`
	// User message for this chat turn.
	Prompt         string            `json:"prompt" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r BetaChatStreamParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaChatStreamParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaChatStreamParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BetaChatStreamParams]'s query parameters as `url.Values`.
func (r BetaChatStreamParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
