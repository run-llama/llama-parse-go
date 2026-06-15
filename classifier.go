// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"github.com/run-llama/llama-parse-go/option"
)

// ClassifierService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClassifierService] method instead.
type ClassifierService struct {
	options []option.RequestOption
	Jobs    ClassifierJobService
}

// NewClassifierService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClassifierService(opts ...option.RequestOption) (r ClassifierService) {
	r = ClassifierService{}
	r.options = opts
	r.Jobs = NewClassifierJobService(opts...)
	return
}
