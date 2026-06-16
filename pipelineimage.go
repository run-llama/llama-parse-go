// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// PipelineImageService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineImageService] method instead.
type PipelineImageService struct {
	options []option.RequestOption
}

// NewPipelineImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineImageService(opts ...option.RequestOption) (r PipelineImageService) {
	r = PipelineImageService{}
	r.options = opts
	return
}

// Get a specific figure from a page of a file.
func (r *PipelineImageService) GetPageFigure(ctx context.Context, figureName string, params PipelineImageGetPageFigureParams, opts ...option.RequestOption) (res *PipelineImageGetPageFigureResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if figureName == "" {
		err = errors.New("missing required figure_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/files/%s/page-figures/%v/%s", params.ID, params.PageIndex, url.PathEscape(figureName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get screenshot of a page from a file.
func (r *PipelineImageService) GetPageScreenshot(ctx context.Context, pageIndex int64, params PipelineImageGetPageScreenshotParams, opts ...option.RequestOption) (res *PipelineImageGetPageScreenshotResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/files/%s/page_screenshots/%v", params.ID, pageIndex)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List metadata for all figures from all pages of a file.
func (r *PipelineImageService) ListPageFigures(ctx context.Context, id string, query PipelineImageListPageFiguresParams, opts ...option.RequestOption) (res *[]PipelineImageListPageFiguresResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/files/%s/page-figures", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List metadata for all screenshots of pages from a file.
func (r *PipelineImageService) ListPageScreenshots(ctx context.Context, id string, query PipelineImageListPageScreenshotsParams, opts ...option.RequestOption) (res *[]PipelineImageListPageScreenshotsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/files/%s/page_screenshots", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PipelineImageGetPageFigureResponse = any

type PipelineImageGetPageScreenshotResponse = any

type PipelineImageListPageFiguresResponse struct {
	// The confidence of the figure
	Confidence float64 `json:"confidence" api:"required"`
	// The name of the figure
	FigureName string `json:"figure_name" api:"required"`
	// The size of the figure in bytes
	FigureSize int64 `json:"figure_size" api:"required"`
	// The ID of the file that the figure was taken from
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// The index of the page for which the figure is taken (0-indexed)
	PageIndex int64 `json:"page_index" api:"required"`
	// Whether the figure is likely to be noise
	IsLikelyNoise bool `json:"is_likely_noise"`
	// Metadata for the figure
	Metadata map[string]any `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence    respjson.Field
		FigureName    respjson.Field
		FigureSize    respjson.Field
		FileID        respjson.Field
		PageIndex     respjson.Field
		IsLikelyNoise respjson.Field
		Metadata      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineImageListPageFiguresResponse) RawJSON() string { return r.JSON.raw }
func (r *PipelineImageListPageFiguresResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineImageListPageScreenshotsResponse struct {
	// The ID of the file that the page screenshot was taken from
	FileID string `json:"file_id" api:"required" format:"uuid"`
	// The size of the image in bytes
	ImageSize int64 `json:"image_size" api:"required"`
	// The index of the page for which the screenshot is taken (0-indexed)
	PageIndex int64 `json:"page_index" api:"required"`
	// Metadata for the screenshot
	Metadata map[string]any `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		ImageSize   respjson.Field
		PageIndex   respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineImageListPageScreenshotsResponse) RawJSON() string { return r.JSON.raw }
func (r *PipelineImageListPageScreenshotsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineImageGetPageFigureParams struct {
	ID             string            `path:"id" api:"required" format:"uuid" json:"-"`
	PageIndex      int64             `path:"page_index" api:"required" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineImageGetPageFigureParams]'s query parameters as
// `url.Values`.
func (r PipelineImageGetPageFigureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineImageGetPageScreenshotParams struct {
	ID             string            `path:"id" api:"required" format:"uuid" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineImageGetPageScreenshotParams]'s query parameters as
// `url.Values`.
func (r PipelineImageGetPageScreenshotParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineImageListPageFiguresParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineImageListPageFiguresParams]'s query parameters as
// `url.Values`.
func (r PipelineImageListPageFiguresParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineImageListPageScreenshotsParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineImageListPageScreenshotsParams]'s query parameters
// as `url.Values`.
func (r PipelineImageListPageScreenshotsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
