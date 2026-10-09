// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud

import (
	"github.com/run-llama/llama-parse-go/option"
)

// AlphaService contains methods and other services that help with interacting with
// the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaService] method instead.
type AlphaService struct {
	options []option.RequestOption
	Verify  AlphaVerifyService
}

// NewAlphaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlphaService(opts ...option.RequestOption) (r AlphaService) {
	r = AlphaService{}
	r.options = opts
	r.Verify = NewAlphaVerifyService(opts...)
	return
}
