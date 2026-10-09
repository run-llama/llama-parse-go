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

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// ChatService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.options = opts
	return
}

// Create a chat session, optionally bound to indexes (locked after the first
// message).
func (r *ChatService) New(ctx context.Context, params ChatNewParams, opts ...option.RequestOption) (res *ChatNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a full session by ID, including its event history.
func (r *ChatService) Get(ctx context.Context, sessionID string, query ChatGetParams, opts ...option.RequestOption) (res *ChatGetResponse, err error) {
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
func (r *ChatService) List(ctx context.Context, query ChatListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ChatListResponse], err error) {
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
func (r *ChatService) ListAutoPaging(ctx context.Context, query ChatListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ChatListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a session.
func (r *ChatService) Delete(ctx context.Context, sessionID string, body ChatDeleteParams, opts ...option.RequestOption) (err error) {
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
func (r *ChatService) GetSummary(ctx context.Context, sessionID string, query ChatGetSummaryParams, opts ...option.RequestOption) (res *ChatGetSummaryResponse, err error) {
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
func (r *ChatService) Stream(ctx context.Context, sessionID string, params ChatStreamParams, opts ...option.RequestOption) (res *ChatStreamResponse, err error) {
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
type ChatNewResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// What this chat's share link grants: read_only (transcript only) or query
	// (viewers may ask new questions).
	//
	// Any of "query", "read_only".
	SharedAccess ChatNewResponseSharedAccess `json:"shared_access" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata ChatNewResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		SharedAccess   respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this chat's share link grants: read_only (transcript only) or query
// (viewers may ask new questions).
type ChatNewResponseSharedAccess string

const (
	ChatNewResponseSharedAccessQuery    ChatNewResponseSharedAccess = "query"
	ChatNewResponseSharedAccessReadOnly ChatNewResponseSharedAccess = "read_only"
)

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type ChatNewResponseJobMetadata struct {
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
func (r ChatNewResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChatNewResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full chat session including its complete event history.
type ChatGetResponse struct {
	// Ordered list of events that make up the conversation history.
	Events []ChatGetResponseEventUnion `json:"events" api:"required"`
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// What this chat's share link grants: read_only (transcript only) or query
	// (viewers may ask new questions).
	//
	// Any of "query", "read_only".
	SharedAccess ChatGetResponseSharedAccess `json:"shared_access" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata ChatGetResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events         respjson.Field
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		SharedAccess   respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChatGetResponseEventUnion contains all possible properties and values from
// [ChatGetResponseEventStop], [ChatGetResponseEventTextDelta],
// [ChatGetResponseEventText], [ChatGetResponseEventThinkingDelta],
// [ChatGetResponseEventThinking], [ChatGetResponseEventToolCall],
// [ChatGetResponseEventToolResult], [ChatGetResponseEventUserInput].
//
// Use the [ChatGetResponseEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChatGetResponseEventUnion struct {
	// This field is from variant [ChatGetResponseEventStop].
	Error string `json:"error"`
	// This field is from variant [ChatGetResponseEventStop].
	IsError bool `json:"is_error"`
	// This field is from variant [ChatGetResponseEventStop].
	Usage ChatGetResponseEventStopUsage `json:"usage"`
	// This field is from variant [ChatGetResponseEventStop].
	SkippedIndexIDs []string `json:"skipped_index_ids"`
	// Any of "stop", "text_delta", "text", "thinking_delta", "thinking", "tool_call",
	// "tool_result", "user_input".
	Type    string `json:"type"`
	Content string `json:"content"`
	// This field is from variant [ChatGetResponseEventToolCall].
	Arguments map[string]any `json:"arguments"`
	CallID    string         `json:"call_id"`
	Name      string         `json:"name"`
	// This field is from variant [ChatGetResponseEventToolResult].
	Result any `json:"result"`
	// This field is from variant [ChatGetResponseEventToolResult].
	ImageAttachment ChatGetResponseEventToolResultImageAttachment `json:"image_attachment"`
	JSON            struct {
		Error           respjson.Field
		IsError         respjson.Field
		Usage           respjson.Field
		SkippedIndexIDs respjson.Field
		Type            respjson.Field
		Content         respjson.Field
		Arguments       respjson.Field
		CallID          respjson.Field
		Name            respjson.Field
		Result          respjson.Field
		ImageAttachment respjson.Field
		raw             string
	} `json:"-"`
}

// anyChatGetResponseEvent is implemented by each variant of
// [ChatGetResponseEventUnion] to add type safety for the return type of
// [ChatGetResponseEventUnion.AsAny]
type anyChatGetResponseEvent interface {
	implChatGetResponseEventUnion()
}

func (ChatGetResponseEventStop) implChatGetResponseEventUnion()          {}
func (ChatGetResponseEventTextDelta) implChatGetResponseEventUnion()     {}
func (ChatGetResponseEventText) implChatGetResponseEventUnion()          {}
func (ChatGetResponseEventThinkingDelta) implChatGetResponseEventUnion() {}
func (ChatGetResponseEventThinking) implChatGetResponseEventUnion()      {}
func (ChatGetResponseEventToolCall) implChatGetResponseEventUnion()      {}
func (ChatGetResponseEventToolResult) implChatGetResponseEventUnion()    {}
func (ChatGetResponseEventUserInput) implChatGetResponseEventUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChatGetResponseEventUnion.AsAny().(type) {
//	case llamacloud.ChatGetResponseEventStop:
//	case llamacloud.ChatGetResponseEventTextDelta:
//	case llamacloud.ChatGetResponseEventText:
//	case llamacloud.ChatGetResponseEventThinkingDelta:
//	case llamacloud.ChatGetResponseEventThinking:
//	case llamacloud.ChatGetResponseEventToolCall:
//	case llamacloud.ChatGetResponseEventToolResult:
//	case llamacloud.ChatGetResponseEventUserInput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChatGetResponseEventUnion) AsAny() anyChatGetResponseEvent {
	switch u.Type {
	case "stop":
		return u.AsStop()
	case "text_delta":
		return u.AsTextDelta()
	case "text":
		return u.AsText()
	case "thinking_delta":
		return u.AsThinkingDelta()
	case "thinking":
		return u.AsThinking()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "user_input":
		return u.AsUserInput()
	}
	return nil
}

func (u ChatGetResponseEventUnion) AsStop() (v ChatGetResponseEventStop) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsTextDelta() (v ChatGetResponseEventTextDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsText() (v ChatGetResponseEventText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsThinkingDelta() (v ChatGetResponseEventThinkingDelta) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsThinking() (v ChatGetResponseEventThinking) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsToolCall() (v ChatGetResponseEventToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsToolResult() (v ChatGetResponseEventToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChatGetResponseEventUnion) AsUserInput() (v ChatGetResponseEventUserInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChatGetResponseEventUnion) RawJSON() string { return u.JSON.raw }

func (r *ChatGetResponseEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventStop struct {
	Error   string                        `json:"error" api:"required"`
	IsError bool                          `json:"is_error" api:"required"`
	Usage   ChatGetResponseEventStopUsage `json:"usage" api:"required"`
	// Requested indexes this turn could not query.
	SkippedIndexIDs []string `json:"skipped_index_ids"`
	// Any of "stop".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error           respjson.Field
		IsError         respjson.Field
		Usage           respjson.Field
		SkippedIndexIDs respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetResponseEventStop) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventStopUsage struct {
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
func (r ChatGetResponseEventStopUsage) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventStopUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventTextDelta struct {
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
func (r ChatGetResponseEventTextDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventTextDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventText struct {
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
func (r ChatGetResponseEventText) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventThinkingDelta struct {
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
func (r ChatGetResponseEventThinkingDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventThinkingDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventThinking struct {
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
func (r ChatGetResponseEventThinking) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventThinking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventToolCall struct {
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
func (r ChatGetResponseEventToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventToolResult struct {
	CallID string `json:"call_id" api:"required"`
	Name   string `json:"name" api:"required"`
	Result any    `json:"result" api:"required"`
	// Coordinates for lazily resolving a page screenshot presigned URL.
	ImageAttachment ChatGetResponseEventToolResultImageAttachment `json:"image_attachment" api:"nullable"`
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
func (r ChatGetResponseEventToolResult) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Coordinates for lazily resolving a page screenshot presigned URL.
type ChatGetResponseEventToolResultImageAttachment struct {
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
func (r ChatGetResponseEventToolResultImageAttachment) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventToolResultImageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatGetResponseEventUserInput struct {
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
func (r ChatGetResponseEventUserInput) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseEventUserInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this chat's share link grants: read_only (transcript only) or query
// (viewers may ask new questions).
type ChatGetResponseSharedAccess string

const (
	ChatGetResponseSharedAccessQuery    ChatGetResponseSharedAccess = "query"
	ChatGetResponseSharedAccessReadOnly ChatGetResponseSharedAccess = "read_only"
)

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type ChatGetResponseJobMetadata struct {
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
func (r ChatGetResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChatGetResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a chat session, including its title and last run metadata.
type ChatListResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// What this chat's share link grants: read_only (transcript only) or query
	// (viewers may ask new questions).
	//
	// Any of "query", "read_only".
	SharedAccess ChatListResponseSharedAccess `json:"shared_access" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata ChatListResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		SharedAccess   respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatListResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this chat's share link grants: read_only (transcript only) or query
// (viewers may ask new questions).
type ChatListResponseSharedAccess string

const (
	ChatListResponseSharedAccessQuery    ChatListResponseSharedAccess = "query"
	ChatListResponseSharedAccessReadOnly ChatListResponseSharedAccess = "read_only"
)

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type ChatListResponseJobMetadata struct {
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
func (r ChatListResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChatListResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a chat session, including its title and last run metadata.
type ChatGetSummaryResponse struct {
	// ISO-format timestamp showing when the session was last updated.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Unique session identifier.
	SessionID string `json:"session_id" api:"required"`
	// What this chat's share link grants: read_only (transcript only) or query
	// (viewers may ask new questions).
	//
	// Any of "query", "read_only".
	SharedAccess ChatGetSummaryResponseSharedAccess `json:"shared_access" api:"required"`
	// Auto-generated title derived from the first user message.
	GeneratedTitle string `json:"generated_title" api:"nullable"`
	// Indexes this session is bound to. Null on unbound sessions.
	IndexIDs []string `json:"index_ids" api:"nullable"`
	// Token usage and status from the most recent run. Null if the session has not
	// been run yet.
	JobMetadata ChatGetSummaryResponseJobMetadata `json:"job_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdatedAt  respjson.Field
		SessionID      respjson.Field
		SharedAccess   respjson.Field
		GeneratedTitle respjson.Field
		IndexIDs       respjson.Field
		JobMetadata    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What this chat's share link grants: read_only (transcript only) or query
// (viewers may ask new questions).
type ChatGetSummaryResponseSharedAccess string

const (
	ChatGetSummaryResponseSharedAccessQuery    ChatGetSummaryResponseSharedAccess = "query"
	ChatGetSummaryResponseSharedAccessReadOnly ChatGetSummaryResponseSharedAccess = "read_only"
)

// Token usage and status from the most recent run. Null if the session has not
// been run yet.
type ChatGetSummaryResponseJobMetadata struct {
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
func (r ChatGetSummaryResponseJobMetadata) RawJSON() string { return r.JSON.raw }
func (r *ChatGetSummaryResponseJobMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatStreamResponse = any

type ChatNewParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Indexes this session will retrieve from. Once set and the first message has been
	// sent, the source set is locked for the session's lifetime. Leave null to create
	// an unbound session.
	IndexIDs []string `json:"index_ids,omitzero"`
	// What this chat's share link grants: read_only (transcript only) or query
	// (viewers may ask new questions). Null follows the deployment default.
	//
	// Any of "query", "read_only".
	SharedAccess ChatNewParamsSharedAccess `json:"shared_access,omitzero"`
	paramObj
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ChatNewParams]'s query parameters as `url.Values`.
func (r ChatNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// What this chat's share link grants: read_only (transcript only) or query
// (viewers may ask new questions). Null follows the deployment default.
type ChatNewParamsSharedAccess string

const (
	ChatNewParamsSharedAccessQuery    ChatNewParamsSharedAccess = "query"
	ChatNewParamsSharedAccessReadOnly ChatNewParamsSharedAccess = "read_only"
)

type ChatGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ChatGetParams]'s query parameters as `url.Values`.
func (r ChatGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ChatListParams]'s query parameters as `url.Values`.
func (r ChatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ChatDeleteParams]'s query parameters as `url.Values`.
func (r ChatDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatGetSummaryParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ChatGetSummaryParams]'s query parameters as `url.Values`.
func (r ChatGetSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChatStreamParams struct {
	// Indexes to retrieve data from.
	IndexIDs []string `json:"index_ids,omitzero" api:"required"`
	// User message for this chat turn.
	Prompt         string            `json:"prompt" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fail the turn if any requested index cannot be queried.
	RequireAllIndexes param.Opt[bool] `json:"require_all_indexes,omitzero"`
	paramObj
}

func (r ChatStreamParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatStreamParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatStreamParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ChatStreamParams]'s query parameters as `url.Values`.
func (r ChatStreamParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
