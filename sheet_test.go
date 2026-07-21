// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestSheetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sheets.New(context.TODO(), llamacloud.SheetNewParams{
		FileID:         "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Config: llamacloud.SheetsParsingConfigParam{
			ExtractionRange:            llamacloud.String("extraction_range"),
			FlattenHierarchicalTables:  llamacloud.Bool(true),
			GenerateAdditionalMetadata: llamacloud.Bool(true),
			IncludeHiddenCells:         llamacloud.Bool(true),
			SheetNames:                 []string{"string"},
			Specialization:             llamacloud.String("specialization"),
			TableMergeSensitivity:      llamacloud.SheetsParsingConfigTableMergeSensitivityStrong,
			UseExperimentalProcessing:  llamacloud.Bool(true),
		},
		Configuration: llamacloud.SheetsParsingConfigParam{
			ExtractionRange:            llamacloud.String("extraction_range"),
			FlattenHierarchicalTables:  llamacloud.Bool(true),
			GenerateAdditionalMetadata: llamacloud.Bool(true),
			IncludeHiddenCells:         llamacloud.Bool(true),
			SheetNames:                 []string{"string"},
			Specialization:             llamacloud.String("specialization"),
			TableMergeSensitivity:      llamacloud.SheetsParsingConfigTableMergeSensitivityStrong,
			UseExperimentalProcessing:  llamacloud.Bool(true),
		},
		ConfigurationID: llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
		WebhookConfigurations: []llamacloud.SheetNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"parse.success", "parse.error"},
			WebhookHeaders: map[string]string{
				"Authorization": "Bearer sk-...",
			},
			WebhookOutputFormat:  llamacloud.String("json"),
			WebhookSigningSecret: llamacloud.String("whsec_..."),
			WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
		}},
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSheetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sheets.List(context.TODO(), llamacloud.SheetListParams{
		ConfigurationID:     llamacloud.String("configuration_id"),
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		IncludeResults:      llamacloud.Bool(true),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(0),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloud.SheetListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSheetDeleteJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Sheets.DeleteJob(
		context.TODO(),
		"spreadsheet_job_id",
		llamacloud.SheetDeleteJobParams{
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

func TestSheetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sheets.Get(
		context.TODO(),
		"spreadsheet_job_id",
		llamacloud.SheetGetParams{
			Expand:         []string{"string"},
			IncludeResults: llamacloud.Bool(true),
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

func TestSheetGetResultTableWithOptionalParams(t *testing.T) {
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
	_, err := client.Sheets.GetResultTable(
		context.TODO(),
		llamacloud.SheetGetResultTableParamsRegionTypeCellMetadata,
		llamacloud.SheetGetResultTableParams{
			SpreadsheetJobID: "spreadsheet_job_id",
			RegionID:         "region_id",
			ExpiresAtSeconds: llamacloud.Int(0),
			OrganizationID:   llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:        llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
