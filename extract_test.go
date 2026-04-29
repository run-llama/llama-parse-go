// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/testutil"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
)

func TestExtractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.New(context.TODO(), llamacloudprod.ExtractNewParams{
		ExtractV2JobCreate: llamacloudprod.ExtractV2JobCreateParam{
			FileInput: "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			Configuration: llamacloudprod.ExtractConfigurationParam{
				DataSchema: map[string]*llamacloudprod.ExtractConfigurationDataSchemaUnionParam{
					"properties": {
						OfAnyMap: map[string]any{
							"vendor_name":  "bar",
							"total_amount": "bar",
						},
					},
					"required": {
						OfAnyArray: []any{"vendor_name", "total_amount"},
					},
					"type": {
						OfString: llamacloudprod.String("object"),
					},
				},
				CiteSources:      llamacloudprod.Bool(true),
				ConfidenceScores: llamacloudprod.Bool(true),
				ExtractVersion:   llamacloudprod.String("latest"),
				ExtractionTarget: llamacloudprod.ExtractConfigurationExtractionTargetPerDoc,
				MaxPages:         llamacloudprod.Int(10),
				ParseConfigID:    llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
				ParseTier:        llamacloudprod.String("fast"),
				SystemPrompt:     llamacloudprod.String("Extract all monetary values in USD. If a currency is not specified, assume USD."),
				TargetPages:      llamacloudprod.String("1,3,5-7"),
				Tier:             llamacloudprod.ExtractConfigurationTierCostEffective,
			},
			ConfigurationID: llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
			WebhookConfigurations: []llamacloudprod.ExtractV2JobCreateWebhookConfigurationParam{{
				WebhookEvents: []string{"parse.success", "parse.error"},
				WebhookHeaders: map[string]string{
					"Authorization": "Bearer sk-...",
				},
				WebhookOutputFormat: llamacloudprod.String("json"),
				WebhookURL:          llamacloudprod.String("https://example.com/webhooks/llamacloud"),
			}},
		},
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractListWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.List(context.TODO(), llamacloudprod.ExtractListParams{
		ConfigurationID:     llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
		CreatedAtOnOrAfter:  llamacloudprod.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloudprod.Time(time.Now()),
		DocumentInputType:   llamacloudprod.String("document_input_type"),
		DocumentInputValue:  llamacloudprod.String("document_input_value"),
		Expand:              []string{"string"},
		FileInput:           llamacloudprod.String("file_input"),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloudprod.Int(0),
		PageToken:           llamacloudprod.String("page_token"),
		ProjectID:           llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloudprod.ExtractListParamsStatusPending,
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.Delete(
		context.TODO(),
		"job_id",
		llamacloudprod.ExtractDeleteParams{
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

func TestExtractGenerateSchemaWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.GenerateSchema(context.TODO(), llamacloudprod.ExtractGenerateSchemaParams{
		ExtractV2SchemaGenerateRequest: llamacloudprod.ExtractV2SchemaGenerateRequestParam{
			DataSchema: map[string]*llamacloudprod.ExtractV2SchemaGenerateRequestDataSchemaUnionParam{
				"foo": {
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
			},
			FileID: llamacloudprod.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			Name:   llamacloudprod.String("invoice_extraction"),
			Prompt: llamacloudprod.String("Extract vendor name, invoice number, date, line items with descriptions and amounts, and total amount from invoices."),
		},
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.Get(
		context.TODO(),
		"job_id",
		llamacloudprod.ExtractGetParams{
			Expand:         []string{"string"},
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

func TestExtractValidateSchema(t *testing.T) {
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
	_, err := client.Extract.ValidateSchema(context.TODO(), llamacloudprod.ExtractValidateSchemaParams{
		ExtractV2SchemaValidateRequest: llamacloudprod.ExtractV2SchemaValidateRequestParam{
			DataSchema: map[string]*llamacloudprod.ExtractV2SchemaValidateRequestDataSchemaUnionParam{
				"properties": {
					OfAnyMap: map[string]any{
						"vendor_name":    "bar",
						"invoice_number": "bar",
						"total_amount":   "bar",
						"line_items":     "bar",
					},
				},
				"required": {
					OfAnyArray: []any{"vendor_name", "invoice_number", "total_amount"},
				},
				"type": {
					OfString: llamacloudprod.String("object"),
				},
			},
		},
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
