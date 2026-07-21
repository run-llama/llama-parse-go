// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestBetaDirectoryFileUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Directories.Files.Update(
		context.TODO(),
		"directory_file_id",
		llamacloud.BetaDirectoryFileUpdateParams{
			DirectoryID:    "directory_id",
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DisplayName:    llamacloud.String("display_name"),
			Metadata: map[string]llamacloud.BetaDirectoryFileUpdateParamsMetadataUnion{
				"foo": {
					OfString: llamacloud.String("string"),
				},
			},
			TargetDirectoryID: llamacloud.String("target_directory_id"),
			UniqueID:          llamacloud.String("x"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDirectoryFileListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Directories.Files.List(
		context.TODO(),
		"directory_id",
		llamacloud.BetaDirectoryFileListParams{
			DisplayName:         llamacloud.String("display_name"),
			DisplayNameContains: llamacloud.String("display_name_contains"),
			Expand:              []string{"string", "string"},
			FileID:              llamacloud.String("file_id"),
			IncludeDeleted:      llamacloud.Bool(true),
			OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PageSize:            llamacloud.Int(0),
			PageToken:           llamacloud.String("page_token"),
			ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UniqueID:            llamacloud.String("unique_id"),
			UpdatedAtOnOrAfter:  llamacloud.Time(time.Now()),
			UpdatedAtOnOrBefore: llamacloud.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDirectoryFileDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Beta.Directories.Files.Delete(
		context.TODO(),
		"directory_file_id",
		llamacloud.BetaDirectoryFileDeleteParams{
			DirectoryID:    "directory_id",
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDirectoryFileAddWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Directories.Files.Add(
		context.TODO(),
		"directory_id",
		llamacloud.BetaDirectoryFileAddParams{
			FileID:         "file_id",
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DisplayName:    llamacloud.String("display_name"),
			Metadata: map[string]llamacloud.BetaDirectoryFileAddParamsMetadataUnion{
				"foo": {
					OfString: llamacloud.String("string"),
				},
			},
			UniqueID: llamacloud.String("unique_id"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDirectoryFileGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Directories.Files.Get(
		context.TODO(),
		"directory_file_id",
		llamacloud.BetaDirectoryFileGetParams{
			DirectoryID:    "directory_id",
			Expand:         []string{"string", "string"},
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaDirectoryFileUploadWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Directories.Files.Upload(
		context.TODO(),
		"directory_id",
		llamacloud.BetaDirectoryFileUploadParams{
			UploadFile:     io.Reader(bytes.NewBuffer([]byte("Example data"))),
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DisplayName:    llamacloud.String("display_name"),
			ExternalFileID: llamacloud.String("external_file_id"),
			Metadata:       llamacloud.String(`{"source": "web", "priority": 1}`),
			UniqueID:       llamacloud.String("unique_id"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
