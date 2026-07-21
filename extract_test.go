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

func TestExtractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.New(context.TODO(), llamacloud.ExtractNewParams{
		ExtractV2JobCreate: llamacloud.ExtractV2JobCreateParam{
			FileInput: "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			Configuration: llamacloud.ExtractConfigurationParam{
				DataSchema: map[string]*llamacloud.ExtractConfigurationDataSchemaUnionParam{
					"properties": {
						OfAnyMap: map[string]any{
							"total_amount": "bar",
							"vendor_name":  "bar",
						},
					},
					"required": {
						OfAnyArray: []any{"total_amount", "vendor_name"},
					},
					"type": {
						OfString: llamacloud.String("object"),
					},
				},
				CiteSources:      llamacloud.Bool(true),
				ConfidenceScores: llamacloud.Bool(true),
				ExtractionTarget: llamacloud.ExtractConfigurationExtractionTargetPerDoc,
				MaxPages:         llamacloud.Int(10),
				ParseConfigID:    llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
				ParseTier:        llamacloud.String("fast"),
				SystemPrompt:     llamacloud.String("Extract all monetary values in USD. If a currency is not specified, assume USD."),
				TargetPages:      llamacloud.String("1,3,5-7"),
				Tier:             llamacloud.ExtractConfigurationTierCostEffective,
				Version:          llamacloud.String("latest"),
			},
			ConfigurationID: llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
			WebhookConfigurations: []llamacloud.ExtractV2JobCreateWebhookConfigurationParam{{
				WebhookEvents: []string{"parse.success", "parse.error"},
				WebhookHeaders: map[string]string{
					"Authorization": "Bearer sk-...",
				},
				WebhookOutputFormat:  llamacloud.String("json"),
				WebhookSigningSecret: llamacloud.String("whsec_..."),
				WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
			}},
		},
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract.List(context.TODO(), llamacloud.ExtractListParams{
		ConfigurationID:     llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		DocumentInputType:   llamacloud.String("document_input_type"),
		DocumentInputValue:  llamacloud.String("document_input_value"),
		Expand:              []string{"string"},
		FileInput:           llamacloud.String("file_input"),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(0),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloud.ExtractListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract.Delete(
		context.TODO(),
		"job_id",
		llamacloud.ExtractDeleteParams{
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

func TestExtractGenerateSchemaWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.GenerateSchema(context.TODO(), llamacloud.ExtractGenerateSchemaParams{
		ExtractV2SchemaGenerateRequest: llamacloud.ExtractV2SchemaGenerateRequestParam{
			DataSchema: map[string]*llamacloud.ExtractV2SchemaGenerateRequestDataSchemaUnionParam{
				"foo": {
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
			},
			FileID: llamacloud.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			Name:   llamacloud.String("invoice_extraction"),
			Prompt: llamacloud.String("Extract vendor name, invoice number, date, line items with descriptions and amounts, and total amount from invoices."),
		},
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract.Get(
		context.TODO(),
		"job_id",
		llamacloud.ExtractGetParams{
			Expand:         []string{"string"},
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

func TestExtractValidateSchema(t *testing.T) {
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
	_, err := client.Extract.ValidateSchema(context.TODO(), llamacloud.ExtractValidateSchemaParams{
		ExtractV2SchemaValidateRequest: llamacloud.ExtractV2SchemaValidateRequestParam{
			DataSchema: map[string]*llamacloud.ExtractV2SchemaValidateRequestDataSchemaUnionParam{
				"properties": {
					OfAnyMap: map[string]any{
						"invoice_number": "bar",
						"line_items":     "bar",
						"total_amount":   "bar",
						"vendor_name":    "bar",
					},
				},
				"required": {
					OfAnyArray: []any{"invoice_number", "total_amount", "vendor_name"},
				},
				"type": {
					OfString: llamacloud.String("object"),
				},
			},
		},
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
