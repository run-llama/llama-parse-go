// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/testutil"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
)

func TestPipelineFileNew(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipelines.Files.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileNewParams{
			Body: []llamacloudprod.PipelineFileNewParamsBody{{
				FileID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				CustomMetadata: map[string]*llamacloudprod.PipelineFileNewParamsBodyCustomMetadataUnion{
					"foo": {
						OfAnyMap: map[string]any{
							"foo": "bar",
						},
					},
				},
			}},
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineFileUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipelines.Files.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileUpdateParams{
			PipelineID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			CustomMetadata: map[string]*llamacloudprod.PipelineFileUpdateParamsCustomMetadataUnion{
				"foo": {
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
			},
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineFileListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipelines.Files.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileListParams{
			DataSourceID:         llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			FileNameContains:     llamacloudprod.String("file_name_contains"),
			Limit:                llamacloudprod.Int(0),
			Offset:               llamacloudprod.Int(0),
			OnlyManuallyUploaded: llamacloudprod.Bool(true),
			OrderBy:              llamacloudprod.String("order_by"),
			Statuses:             []string{"NOT_STARTED", "IN_PROGRESS"},
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineFileDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Pipelines.Files.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileDeleteParams{
			PipelineID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineFileGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipelines.Files.GetStatus(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileGetStatusParams{
			PipelineID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineFileGetStatusCountsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipelines.Files.GetStatusCounts(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineFileGetStatusCountsParams{
			DataSourceID:         llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			OnlyManuallyUploaded: llamacloudprod.Bool(true),
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
