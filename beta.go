// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"github.com/run-llama/llama-parse-go/option"
)

// BetaService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	options     []option.RequestOption
	Indexes     BetaIndexService
	Retrieval   BetaRetrievalService
	Chat        BetaChatService
	AgentData   BetaAgentDataService
	Sheets      BetaSheetService
	Directories BetaDirectoryService
	Batch       BetaBatchService
	Split       BetaSplitService
}

// NewBetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaService(opts ...option.RequestOption) (r BetaService) {
	r = BetaService{}
	r.options = opts
	r.Indexes = NewBetaIndexService(opts...)
	r.Retrieval = NewBetaRetrievalService(opts...)
	r.Chat = NewBetaChatService(opts...)
	r.AgentData = NewBetaAgentDataService(opts...)
	r.Sheets = NewBetaSheetService(opts...)
	r.Directories = NewBetaDirectoryService(opts...)
	r.Batch = NewBetaBatchService(opts...)
	r.Split = NewBetaSplitService(opts...)
	return
}
