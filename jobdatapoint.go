// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"github.com/run-llama/llama-parse-go/option"
)

// JobDataPointService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobDataPointService] method instead.
type JobDataPointService struct {
	options []option.RequestOption
}

// NewJobDataPointService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobDataPointService(opts ...option.RequestOption) (r JobDataPointService) {
	r = JobDataPointService{}
	r.options = opts
	return
}
