// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiform"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
)

// PipelineMetadataService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineMetadataService] method instead.
type PipelineMetadataService struct {
	options []option.RequestOption
}

// NewPipelineMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPipelineMetadataService(opts ...option.RequestOption) (r PipelineMetadataService) {
	r = PipelineMetadataService{}
	r.options = opts
	return
}

// Import metadata for a pipeline.
func (r *PipelineMetadataService) New(ctx context.Context, pipelineID string, body PipelineMetadataNewParams, opts ...option.RequestOption) (res *PipelineMetadataNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/metadata", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Delete metadata for all files in a pipeline.
func (r *PipelineMetadataService) DeleteAll(ctx context.Context, pipelineID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/metadata", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PipelineMetadataNewResponse map[string]string

type PipelineMetadataNewParams struct {
	UploadFile io.Reader `json:"upload_file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r PipelineMetadataNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
