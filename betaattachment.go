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

// BetaAttachmentService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaAttachmentService] method instead.
type BetaAttachmentService struct {
	options []option.RequestOption
}

// NewBetaAttachmentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaAttachmentService(opts ...option.RequestOption) (r BetaAttachmentService) {
	r = BetaAttachmentService{}
	r.options = opts
	return
}

// List the attachments associated with a file (e.g. per-page screenshots).
func (r *BetaAttachmentService) List(ctx context.Context, query BetaAttachmentListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[BetaAttachmentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/attachments"
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

// List the attachments associated with a file (e.g. per-page screenshots).
func (r *BetaAttachmentService) ListAutoPaging(ctx context.Context, query BetaAttachmentListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[BetaAttachmentListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Return a presigned download URL for a specific attachment.
func (r *BetaAttachmentService) Get(ctx context.Context, attachmentName string, query BetaAttachmentGetParams, opts ...option.RequestOption) (res *PresignedURL, err error) {
	opts = slices.Concat(r.options, opts)
	if attachmentName == "" {
		err = errors.New("missing required attachment_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/attachments/%s", url.PathEscape(attachmentName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Metadata for a single file attachment.
type BetaAttachmentListResponse struct {
	// Name of the attachment
	Name string `json:"name" api:"required"`
	// Size of the attachment in bytes
	Size int64 `json:"size" api:"required"`
	// When the attachment was last modified
	LastModified time.Time `json:"last_modified" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		Size         respjson.Field
		LastModified respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaAttachmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaAttachmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaAttachmentListParams struct {
	// File UUID or directory file ID (dfl-...).
	SourceID       string            `query:"source_id" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PageToken      param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaAttachmentListParams]'s query parameters as
// `url.Values`.
func (r BetaAttachmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaAttachmentGetParams struct {
	// File UUID or directory file ID (dfl-...).
	SourceID       string            `query:"source_id" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [BetaAttachmentGetParams]'s query parameters as
// `url.Values`.
func (r BetaAttachmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
