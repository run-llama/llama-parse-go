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
)

// RetrieverQueryService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetrieverQueryService] method instead.
type RetrieverQueryService struct {
	options []option.RequestOption
}

// NewRetrieverQueryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRetrieverQueryService(opts ...option.RequestOption) (r RetrieverQueryService) {
	r = RetrieverQueryService{}
	r.options = opts
	return
}

// Retrieve data using a Retriever.
func (r *RetrieverQueryService) Search(ctx context.Context, retrieverID string, params RetrieverQuerySearchParams, opts ...option.RequestOption) (res *CompositeRetrievalResult, err error) {
	opts = slices.Concat(r.options, opts)
	if retrieverID == "" {
		err = errors.New("missing required retriever_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/retrievers/%s/retrieve", retrieverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type RetrieverQuerySearchParams struct {
	// The query to retrieve against.
	Query          string            `json:"query" api:"required"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// (use rerank_config.top_n instead) The number of nodes to retrieve after
	// reranking over retrieved nodes from all retrieval tools.
	RerankTopN param.Opt[int64] `json:"rerank_top_n,omitzero"`
	// The mode of composite retrieval.
	//
	// Any of "routing", "full".
	Mode CompositeRetrievalMode `json:"mode,omitzero"`
	// The rerank configuration for composite retrieval.
	RerankConfig ReRankConfigParam `json:"rerank_config,omitzero"`
	paramObj
}

func (r RetrieverQuerySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow RetrieverQuerySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrieverQuerySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RetrieverQuerySearchParams]'s query parameters as
// `url.Values`.
func (r RetrieverQuerySearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
