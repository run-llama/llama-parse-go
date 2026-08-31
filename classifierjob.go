// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"github.com/run-llama/llama-parse-go/option"
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
