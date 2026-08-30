// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/run-llama/llama-parse-go/internal/apiquery"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/param"
)

// PipelineSyncService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineSyncService] method instead.
type PipelineSyncService struct {
	options []option.RequestOption
}

// NewPipelineSyncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineSyncService(opts ...option.RequestOption) (r PipelineSyncService) {
	r = PipelineSyncService{}
	r.options = opts
	return
}

// Trigger an incremental sync for a managed pipeline.
//
// Processes new and updated documents from data sources and files, then updates
// the index for retrieval.
//
// Deprecated: deprecated
func (r *PipelineSyncService) New(ctx context.Context, pipelineID string, body PipelineSyncNewParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/sync", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cancel all running sync jobs for a pipeline.
//
// Deprecated: deprecated
func (r *PipelineSyncService) Cancel(ctx context.Context, pipelineID string, body PipelineSyncCancelParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/sync/cancel", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PipelineSyncNewParams struct {
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineSyncNewParams]'s query parameters as `url.Values`.
func (r PipelineSyncNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PipelineSyncCancelParams struct {
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PipelineSyncCancelParams]'s query parameters as
// `url.Values`.
func (r PipelineSyncCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
