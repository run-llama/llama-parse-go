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

func TestClassifyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Classify.New(context.TODO(), llamacloudprod.ClassifyNewParams{
		ClassifyCreateRequest: llamacloudprod.ClassifyCreateRequestParam{
			Configuration: llamacloudprod.ClassifyConfigurationParam{
				Rules: []llamacloudprod.ClassifyConfigurationRuleParam{{
					Description: "contains invoice number, line items, and total amount",
					Type:        "invoice",
				}},
				Mode: llamacloudprod.ClassifyConfigurationModeFast,
				ParsingConfiguration: llamacloudprod.ClassifyConfigurationParsingConfigurationParam{
					Lang:        llamacloudprod.String("en"),
					MaxPages:    llamacloudprod.Int(10),
					TargetPages: llamacloudprod.String("1,3,5-7"),
				},
			},
			ConfigurationID: llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
			FileID:          llamacloudprod.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			FileInput:       llamacloudprod.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			ParseJobID:      llamacloudprod.String("pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			TransactionID:   llamacloudprod.String("tx-unique-idempotency-key"),
			WebhookConfigurations: []llamacloudprod.ClassifyCreateRequestWebhookConfigurationParam{{
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

func TestClassifyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Classify.List(context.TODO(), llamacloudprod.ClassifyListParams{
		ConfigurationID:     llamacloudprod.String("cfg-11111111-2222-3333-4444-555555555555"),
		CreatedAtOnOrAfter:  llamacloudprod.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloudprod.Time(time.Now()),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloudprod.Int(1),
		PageToken:           llamacloudprod.String("page_token"),
		ProjectID:           llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloudprod.ClassifyListParamsStatusCompleted,
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Classify.Get(
		context.TODO(),
		"job_id",
		llamacloudprod.ClassifyGetParams{
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
