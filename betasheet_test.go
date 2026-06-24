// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

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

func TestBetaSheetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sheets.New(context.TODO(), llamacloudprod.BetaSheetNewParams{
		FileID:         "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Config: llamacloudprod.SheetsParsingConfigParam{
			ExtractionRange:            llamacloudprod.String("extraction_range"),
			FlattenHierarchicalTables:  llamacloudprod.Bool(true),
			GenerateAdditionalMetadata: llamacloudprod.Bool(true),
			IncludeHiddenCells:         llamacloudprod.Bool(true),
			SheetNames:                 []string{"string"},
			Specialization:             llamacloudprod.String("specialization"),
			TableMergeSensitivity:      llamacloudprod.SheetsParsingConfigTableMergeSensitivityStrong,
			UseExperimentalProcessing:  llamacloudprod.Bool(true),
		},
		Configuration: llamacloudprod.SheetsParsingConfigParam{
			ExtractionRange:            llamacloudprod.String("extraction_range"),
			FlattenHierarchicalTables:  llamacloudprod.Bool(true),
			GenerateAdditionalMetadata: llamacloudprod.Bool(true),
			IncludeHiddenCells:         llamacloudprod.Bool(true),
			SheetNames:                 []string{"string"},
			Specialization:             llamacloudprod.String("specialization"),
			TableMergeSensitivity:      llamacloudprod.SheetsParsingConfigTableMergeSensitivityStrong,
			UseExperimentalProcessing:  llamacloudprod.Bool(true),
		},
		ConfigurationID: llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
		WebhookConfigurations: []llamacloudprod.BetaSheetNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"parse.success", "parse.error"},
			WebhookHeaders: map[string]string{
				"Authorization": "Bearer sk-...",
			},
			WebhookOutputFormat: llamacloudprod.String("json"),
			WebhookURL:          llamacloudprod.String("https://example.com/webhooks/llamacloud"),
		}},
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaSheetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sheets.List(context.TODO(), llamacloudprod.BetaSheetListParams{
		ConfigurationID:     llamacloudprod.String("configuration_id"),
		CreatedAtOnOrAfter:  llamacloudprod.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloudprod.Time(time.Now()),
		IncludeResults:      llamacloudprod.Bool(true),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloudprod.Int(0),
		PageToken:           llamacloudprod.String("page_token"),
		ProjectID:           llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloudprod.BetaSheetListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaSheetDeleteJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sheets.DeleteJob(
		context.TODO(),
		"spreadsheet_job_id",
		llamacloudprod.BetaSheetDeleteJobParams{
			OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestBetaSheetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sheets.Get(
		context.TODO(),
		"spreadsheet_job_id",
		llamacloudprod.BetaSheetGetParams{
			Expand:         []string{"string"},
			IncludeResults: llamacloudprod.Bool(true),
			OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestBetaSheetGetResultTableWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Sheets.GetResultTable(
		context.TODO(),
		llamacloudprod.BetaSheetGetResultTableParamsRegionTypeCellMetadata,
		llamacloudprod.BetaSheetGetResultTableParams{
			SpreadsheetJobID: "spreadsheet_job_id",
			RegionID:         "region_id",
			ExpiresAtSeconds: llamacloudprod.Int(0),
			OrganizationID:   llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:        llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
